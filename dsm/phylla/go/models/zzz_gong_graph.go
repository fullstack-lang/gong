// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Angle0Shape:
		ok = stage.IsStagedAngle0Shape(target)

	case *ArcNormalVectorShape:
		ok = stage.IsStagedArcNormalVectorShape(target)

	case *ArcNormalVectorShapeGrid:
		ok = stage.IsStagedArcNormalVectorShapeGrid(target)

	case *AxesShape:
		ok = stage.IsStagedAxesShape(target)

	case *BaseVectorShape:
		ok = stage.IsStagedBaseVectorShape(target)

	case *BaseVectorShapeGrid:
		ok = stage.IsStagedBaseVectorShapeGrid(target)

	case *ChosenP1P2PairShape:
		ok = stage.IsStagedChosenP1P2PairShape(target)

	case *CircleGridShape:
		ok = stage.IsStagedCircleGridShape(target)

	case *Circumference3DShape:
		ok = stage.IsStagedCircumference3DShape(target)

	case *Clock2DDiagram:
		ok = stage.IsStagedClock2DDiagram(target)

	case *Clock3DDiagram:
		ok = stage.IsStagedClock3DDiagram(target)

	case *ClockAbstract:
		ok = stage.IsStagedClockAbstract(target)

	case *ClockTopCurveShape:
		ok = stage.IsStagedClockTopCurveShape(target)

	case *CutLine3DShape:
		ok = stage.IsStagedCutLine3DShape(target)

	case *EndArcShape:
		ok = stage.IsStagedEndArcShape(target)

	case *EndArcShapeGrid:
		ok = stage.IsStagedEndArcShapeGrid(target)

	case *EndHalfwayArcShape:
		ok = stage.IsStagedEndHalfwayArcShape(target)

	case *EndHalfwayArcShapeGrid:
		ok = stage.IsStagedEndHalfwayArcShapeGrid(target)

	case *ExplanationTextShape:
		ok = stage.IsStagedExplanationTextShape(target)

	case *Eye3DShape:
		ok = stage.IsStagedEye3DShape(target)

	case *EyeCornersSampledPoints3DShape:
		ok = stage.IsStagedEyeCornersSampledPoints3DShape(target)

	case *EyeSampledPoints3DShape:
		ok = stage.IsStagedEyeSampledPoints3DShape(target)

	case *EyeSeatBottomCurveShape:
		ok = stage.IsStagedEyeSeatBottomCurveShape(target)

	case *EyeStoolBottomCurveShape:
		ok = stage.IsStagedEyeStoolBottomCurveShape(target)

	case *EyeVolume3DShape:
		ok = stage.IsStagedEyeVolume3DShape(target)

	case *GridPathShape:
		ok = stage.IsStagedGridPathShape(target)

	case *GrowthCurve2D:
		ok = stage.IsStagedGrowthCurve2D(target)

	case *GrowthCurve2DRibbon:
		ok = stage.IsStagedGrowthCurve2DRibbon(target)

	case *GrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedGrowthCurve2DRibbonEndShape(target)

	case *GrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedGrowthCurve2DRibbonStartShape(target)

	case *GrowthCurveRhombusGridShape:
		ok = stage.IsStagedGrowthCurveRhombusGridShape(target)

	case *GrowthCurveRhombusShape:
		ok = stage.IsStagedGrowthCurveRhombusShape(target)

	case *GrowthVectorShape:
		ok = stage.IsStagedGrowthVectorShape(target)

	case *InitialRhombusGridShape:
		ok = stage.IsStagedInitialRhombusGridShape(target)

	case *InitialRhombusShape:
		ok = stage.IsStagedInitialRhombusShape(target)

	case *Key3DShape:
		ok = stage.IsStagedKey3DShape(target)

	case *KeyHole3DShape:
		ok = stage.IsStagedKeyHole3DShape(target)

	case *KeyHoleShape:
		ok = stage.IsStagedKeyHoleShape(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *MidArcVectorShape:
		ok = stage.IsStagedMidArcVectorShape(target)

	case *MidArcVectorShapeGrid:
		ok = stage.IsStagedMidArcVectorShapeGrid(target)

	case *OriginalPoints3DShape:
		ok = stage.IsStagedOriginalPoints3DShape(target)

	case *ParastichyMCurves3DShape:
		ok = stage.IsStagedParastichyMCurves3DShape(target)

	case *ParastichyNCurves3DShape:
		ok = stage.IsStagedParastichyNCurves3DShape(target)

	case *PartiallyGrowthCurve2DRibbon:
		ok = stage.IsStagedPartiallyGrowthCurve2DRibbon(target)

	case *PartiallyGrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DRibbonEndShape(target)

	case *PartiallyGrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DRibbonStartShape(target)

	case *PartiallyGrowthCurve2DTrajectory:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectory(target)

	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryP1CurveShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryP1P2(target)

	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryP1PointShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryP2CurveShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryP2PointShape(target)

	case *PartiallyGrowthCurve2DTrajectoryShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryShape(target)

	case *PartiallyRotatedSeatBottomCurveShape:
		ok = stage.IsStagedPartiallyRotatedSeatBottomCurveShape(target)

	case *PartiallyRotatedSeatTopCurveShape:
		ok = stage.IsStagedPartiallyRotatedSeatTopCurveShape(target)

	case *PartiallyRotatedTorusShape:
		ok = stage.IsStagedPartiallyRotatedTorusShape(target)

	case *PerpendicularVector:
		ok = stage.IsStagedPerpendicularVector(target)

	case *PerpendicularVectorGrid:
		ok = stage.IsStagedPerpendicularVectorGrid(target)

	case *PerpendicularVectorGridHalfway:
		ok = stage.IsStagedPerpendicularVectorGridHalfway(target)

	case *PerpendicularVectorHalfway:
		ok = stage.IsStagedPerpendicularVectorHalfway(target)

	case *Plant2DDiagram:
		ok = stage.IsStagedPlant2DDiagram(target)

	case *Plant3DDiagram:
		ok = stage.IsStagedPlant3DDiagram(target)

	case *PlantAbstract:
		ok = stage.IsStagedPlantAbstract(target)

	case *PlantCircumferenceShape:
		ok = stage.IsStagedPlantCircumferenceShape(target)

	case *PointsAndLines3DShape:
		ok = stage.IsStagedPointsAndLines3DShape(target)

	case *PxShape:
		ok = stage.IsStagedPxShape(target)

	case *Rendered3DShape:
		ok = stage.IsStagedRendered3DShape(target)

	case *RhombusShape:
		ok = stage.IsStagedRhombusShape(target)

	case *RhombusStuff:
		ok = stage.IsStagedRhombusStuff(target)

	case *RotatedRhombusGridShape:
		ok = stage.IsStagedRotatedRhombusGridShape(target)

	case *RotatedRhombusShape:
		ok = stage.IsStagedRotatedRhombusShape(target)

	case *RotatedSampledPoints3DShape:
		ok = stage.IsStagedRotatedSampledPoints3DShape(target)

	case *RotatedSeatAndLegs3DShape:
		ok = stage.IsStagedRotatedSeatAndLegs3DShape(target)

	case *SampledPoints3DShape:
		ok = stage.IsStagedSampledPoints3DShape(target)

	case *Seat3DShape:
		ok = stage.IsStagedSeat3DShape(target)

	case *SeatAndLegs3DShape:
		ok = stage.IsStagedSeatAndLegs3DShape(target)

	case *SeatBottomCurveShape:
		ok = stage.IsStagedSeatBottomCurveShape(target)

	case *SeatTopCurveShape:
		ok = stage.IsStagedSeatTopCurveShape(target)

	case *ShiftedBottomTopStartArcShape:
		ok = stage.IsStagedShiftedBottomTopStartArcShape(target)

	case *ShiftedBottomTopStartArcShapeGrid:
		ok = stage.IsStagedShiftedBottomTopStartArcShapeGrid(target)

	case *ShiftedLeftGrowthCurve2DRibbon:
		ok = stage.IsStagedShiftedLeftGrowthCurve2DRibbon(target)

	case *ShiftedLeftGrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedShiftedLeftGrowthCurve2DRibbonEndShape(target)

	case *ShiftedLeftGrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedShiftedLeftGrowthCurve2DRibbonStartShape(target)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbon:
		ok = stage.IsStagedShiftedLeftPartiallyGrowthCurve2DRibbon(target)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(target)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(target)

	case *ShiftedLeftStackGrowthCurveEndArcShape:
		ok = stage.IsStagedShiftedLeftStackGrowthCurveEndArcShape(target)

	case *ShiftedLeftStackGrowthCurveStartArcShape:
		ok = stage.IsStagedShiftedLeftStackGrowthCurveStartArcShape(target)

	case *ShiftedLeftStackNormalVector:
		ok = stage.IsStagedShiftedLeftStackNormalVector(target)

	case *ShiftedLeftStackOfGrowthCurve:
		ok = stage.IsStagedShiftedLeftStackOfGrowthCurve(target)

	case *ShiftedLeftStackOfNormalVector:
		ok = stage.IsStagedShiftedLeftStackOfNormalVector(target)

	case *ShiftedRightGrowthCurve2DRibbon:
		ok = stage.IsStagedShiftedRightGrowthCurve2DRibbon(target)

	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedShiftedRightGrowthCurve2DRibbonEndShape(target)

	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedShiftedRightGrowthCurve2DRibbonStartShape(target)

	case *StackGrowthCurve2DEndHalfwayArcShape:
		ok = stage.IsStagedStackGrowthCurve2DEndHalfwayArcShape(target)

	case *StackGrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedStackGrowthCurve2DRibbonEndShape(target)

	case *StackGrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedStackGrowthCurve2DRibbonStartShape(target)

	case *StackGrowthCurve2DStartHalfwayArcShape:
		ok = stage.IsStagedStackGrowthCurve2DStartHalfwayArcShape(target)

	case *StackOfGrowthCurve2D:
		ok = stage.IsStagedStackOfGrowthCurve2D(target)

	case *StackOfGrowthCurve2DByGrowthVector:
		ok = stage.IsStagedStackOfGrowthCurve2DByGrowthVector(target)

	case *StackOfGrowthCurve2DRibbon:
		ok = stage.IsStagedStackOfGrowthCurve2DRibbon(target)

	case *StackOfPartiallyRotatedTorusShape:
		ok = stage.IsStagedStackOfPartiallyRotatedTorusShape(target)

	case *StackOfRotatedGrowthCurve2D:
		ok = stage.IsStagedStackOfRotatedGrowthCurve2D(target)

	case *StackOfRotatedGrowthCurve2DRibbon:
		ok = stage.IsStagedStackOfRotatedGrowthCurve2DRibbon(target)

	case *StackRotatedGrowthCurve2DEndArcShape:
		ok = stage.IsStagedStackRotatedGrowthCurve2DEndArcShape(target)

	case *StackRotatedGrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedStackRotatedGrowthCurve2DRibbonEndShape(target)

	case *StackRotatedGrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedStackRotatedGrowthCurve2DRibbonStartShape(target)

	case *StackRotatedGrowthCurve2DStartArcShape:
		ok = stage.IsStagedStackRotatedGrowthCurve2DStartArcShape(target)

	case *StartArcShape:
		ok = stage.IsStagedStartArcShape(target)

	case *StartArcShapeGrid:
		ok = stage.IsStagedStartArcShapeGrid(target)

	case *StartHalfwayArcShape:
		ok = stage.IsStagedStartHalfwayArcShape(target)

	case *StartHalfwayArcShapeGrid:
		ok = stage.IsStagedStartHalfwayArcShapeGrid(target)

	case *StemCylinder3DShape:
		ok = stage.IsStagedStemCylinder3DShape(target)

	case *Stool2DDiagram:
		ok = stage.IsStagedStool2DDiagram(target)

	case *Stool3DDiagram:
		ok = stage.IsStagedStool3DDiagram(target)

	case *StoolAbstract:
		ok = stage.IsStagedStoolAbstract(target)

	case *TiledFloor3DShape:
		ok = stage.IsStagedTiledFloor3DShape(target)

	case *TopEndArcShape:
		ok = stage.IsStagedTopEndArcShape(target)

	case *TopEndArcShapeGrid:
		ok = stage.IsStagedTopEndArcShapeGrid(target)

	case *TopEndHalfwayArcShape:
		ok = stage.IsStagedTopEndHalfwayArcShape(target)

	case *TopEndHalfwayArcShapeGrid:
		ok = stage.IsStagedTopEndHalfwayArcShapeGrid(target)

	case *TopGrowthCurve2D:
		ok = stage.IsStagedTopGrowthCurve2D(target)

	case *TopMidArcVectorShape:
		ok = stage.IsStagedTopMidArcVectorShape(target)

	case *TopMidArcVectorShapeGrid:
		ok = stage.IsStagedTopMidArcVectorShapeGrid(target)

	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		ok = stage.IsStagedTopStackGrowthCurve2DEndHalfwayArcShape(target)

	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		ok = stage.IsStagedTopStackGrowthCurve2DStartHalfwayArcShape(target)

	case *TopStackOfGrowthCurve2D:
		ok = stage.IsStagedTopStackOfGrowthCurve2D(target)

	case *TopStackOfRotatedGrowthCurve2D:
		ok = stage.IsStagedTopStackOfRotatedGrowthCurve2D(target)

	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		ok = stage.IsStagedTopStackOfRotatedGrowthCurve2DEndArcShape(target)

	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		ok = stage.IsStagedTopStackOfRotatedGrowthCurve2DStartArcShape(target)

	case *TopStartArcShape:
		ok = stage.IsStagedTopStartArcShape(target)

	case *TopStartArcShapeGrid:
		ok = stage.IsStagedTopStartArcShapeGrid(target)

	case *TopStartHalfwayArcShape:
		ok = stage.IsStagedTopStartHalfwayArcShape(target)

	case *TopStartHalfwayArcShapeGrid:
		ok = stage.IsStagedTopStartHalfwayArcShapeGrid(target)

	case *Torus3DShape:
		ok = stage.IsStagedTorus3DShape(target)

	case *TorusEdge3DShape:
		ok = stage.IsStagedTorusEdge3DShape(target)

	case *TorusStackShape:
		ok = stage.IsStagedTorusStackShape(target)

	case *Vase2DDiagram:
		ok = stage.IsStagedVase2DDiagram(target)

	case *Vase3DDiagram:
		ok = stage.IsStagedVase3DDiagram(target)

	case *VaseAbstract:
		ok = stage.IsStagedVaseAbstract(target)

	case *VerticalTorusStackShape:
		ok = stage.IsStagedVerticalTorusStackShape(target)

	case *VolumeKey3DShape:
		ok = stage.IsStagedVolumeKey3DShape(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Angle0Shape:
		ok = stage.IsStagedAngle0Shape(target)

	case *ArcNormalVectorShape:
		ok = stage.IsStagedArcNormalVectorShape(target)

	case *ArcNormalVectorShapeGrid:
		ok = stage.IsStagedArcNormalVectorShapeGrid(target)

	case *AxesShape:
		ok = stage.IsStagedAxesShape(target)

	case *BaseVectorShape:
		ok = stage.IsStagedBaseVectorShape(target)

	case *BaseVectorShapeGrid:
		ok = stage.IsStagedBaseVectorShapeGrid(target)

	case *ChosenP1P2PairShape:
		ok = stage.IsStagedChosenP1P2PairShape(target)

	case *CircleGridShape:
		ok = stage.IsStagedCircleGridShape(target)

	case *Circumference3DShape:
		ok = stage.IsStagedCircumference3DShape(target)

	case *Clock2DDiagram:
		ok = stage.IsStagedClock2DDiagram(target)

	case *Clock3DDiagram:
		ok = stage.IsStagedClock3DDiagram(target)

	case *ClockAbstract:
		ok = stage.IsStagedClockAbstract(target)

	case *ClockTopCurveShape:
		ok = stage.IsStagedClockTopCurveShape(target)

	case *CutLine3DShape:
		ok = stage.IsStagedCutLine3DShape(target)

	case *EndArcShape:
		ok = stage.IsStagedEndArcShape(target)

	case *EndArcShapeGrid:
		ok = stage.IsStagedEndArcShapeGrid(target)

	case *EndHalfwayArcShape:
		ok = stage.IsStagedEndHalfwayArcShape(target)

	case *EndHalfwayArcShapeGrid:
		ok = stage.IsStagedEndHalfwayArcShapeGrid(target)

	case *ExplanationTextShape:
		ok = stage.IsStagedExplanationTextShape(target)

	case *Eye3DShape:
		ok = stage.IsStagedEye3DShape(target)

	case *EyeCornersSampledPoints3DShape:
		ok = stage.IsStagedEyeCornersSampledPoints3DShape(target)

	case *EyeSampledPoints3DShape:
		ok = stage.IsStagedEyeSampledPoints3DShape(target)

	case *EyeSeatBottomCurveShape:
		ok = stage.IsStagedEyeSeatBottomCurveShape(target)

	case *EyeStoolBottomCurveShape:
		ok = stage.IsStagedEyeStoolBottomCurveShape(target)

	case *EyeVolume3DShape:
		ok = stage.IsStagedEyeVolume3DShape(target)

	case *GridPathShape:
		ok = stage.IsStagedGridPathShape(target)

	case *GrowthCurve2D:
		ok = stage.IsStagedGrowthCurve2D(target)

	case *GrowthCurve2DRibbon:
		ok = stage.IsStagedGrowthCurve2DRibbon(target)

	case *GrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedGrowthCurve2DRibbonEndShape(target)

	case *GrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedGrowthCurve2DRibbonStartShape(target)

	case *GrowthCurveRhombusGridShape:
		ok = stage.IsStagedGrowthCurveRhombusGridShape(target)

	case *GrowthCurveRhombusShape:
		ok = stage.IsStagedGrowthCurveRhombusShape(target)

	case *GrowthVectorShape:
		ok = stage.IsStagedGrowthVectorShape(target)

	case *InitialRhombusGridShape:
		ok = stage.IsStagedInitialRhombusGridShape(target)

	case *InitialRhombusShape:
		ok = stage.IsStagedInitialRhombusShape(target)

	case *Key3DShape:
		ok = stage.IsStagedKey3DShape(target)

	case *KeyHole3DShape:
		ok = stage.IsStagedKeyHole3DShape(target)

	case *KeyHoleShape:
		ok = stage.IsStagedKeyHoleShape(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *MidArcVectorShape:
		ok = stage.IsStagedMidArcVectorShape(target)

	case *MidArcVectorShapeGrid:
		ok = stage.IsStagedMidArcVectorShapeGrid(target)

	case *OriginalPoints3DShape:
		ok = stage.IsStagedOriginalPoints3DShape(target)

	case *ParastichyMCurves3DShape:
		ok = stage.IsStagedParastichyMCurves3DShape(target)

	case *ParastichyNCurves3DShape:
		ok = stage.IsStagedParastichyNCurves3DShape(target)

	case *PartiallyGrowthCurve2DRibbon:
		ok = stage.IsStagedPartiallyGrowthCurve2DRibbon(target)

	case *PartiallyGrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DRibbonEndShape(target)

	case *PartiallyGrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DRibbonStartShape(target)

	case *PartiallyGrowthCurve2DTrajectory:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectory(target)

	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryP1CurveShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryP1P2(target)

	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryP1PointShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryP2CurveShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryP2PointShape(target)

	case *PartiallyGrowthCurve2DTrajectoryShape:
		ok = stage.IsStagedPartiallyGrowthCurve2DTrajectoryShape(target)

	case *PartiallyRotatedSeatBottomCurveShape:
		ok = stage.IsStagedPartiallyRotatedSeatBottomCurveShape(target)

	case *PartiallyRotatedSeatTopCurveShape:
		ok = stage.IsStagedPartiallyRotatedSeatTopCurveShape(target)

	case *PartiallyRotatedTorusShape:
		ok = stage.IsStagedPartiallyRotatedTorusShape(target)

	case *PerpendicularVector:
		ok = stage.IsStagedPerpendicularVector(target)

	case *PerpendicularVectorGrid:
		ok = stage.IsStagedPerpendicularVectorGrid(target)

	case *PerpendicularVectorGridHalfway:
		ok = stage.IsStagedPerpendicularVectorGridHalfway(target)

	case *PerpendicularVectorHalfway:
		ok = stage.IsStagedPerpendicularVectorHalfway(target)

	case *Plant2DDiagram:
		ok = stage.IsStagedPlant2DDiagram(target)

	case *Plant3DDiagram:
		ok = stage.IsStagedPlant3DDiagram(target)

	case *PlantAbstract:
		ok = stage.IsStagedPlantAbstract(target)

	case *PlantCircumferenceShape:
		ok = stage.IsStagedPlantCircumferenceShape(target)

	case *PointsAndLines3DShape:
		ok = stage.IsStagedPointsAndLines3DShape(target)

	case *PxShape:
		ok = stage.IsStagedPxShape(target)

	case *Rendered3DShape:
		ok = stage.IsStagedRendered3DShape(target)

	case *RhombusShape:
		ok = stage.IsStagedRhombusShape(target)

	case *RhombusStuff:
		ok = stage.IsStagedRhombusStuff(target)

	case *RotatedRhombusGridShape:
		ok = stage.IsStagedRotatedRhombusGridShape(target)

	case *RotatedRhombusShape:
		ok = stage.IsStagedRotatedRhombusShape(target)

	case *RotatedSampledPoints3DShape:
		ok = stage.IsStagedRotatedSampledPoints3DShape(target)

	case *RotatedSeatAndLegs3DShape:
		ok = stage.IsStagedRotatedSeatAndLegs3DShape(target)

	case *SampledPoints3DShape:
		ok = stage.IsStagedSampledPoints3DShape(target)

	case *Seat3DShape:
		ok = stage.IsStagedSeat3DShape(target)

	case *SeatAndLegs3DShape:
		ok = stage.IsStagedSeatAndLegs3DShape(target)

	case *SeatBottomCurveShape:
		ok = stage.IsStagedSeatBottomCurveShape(target)

	case *SeatTopCurveShape:
		ok = stage.IsStagedSeatTopCurveShape(target)

	case *ShiftedBottomTopStartArcShape:
		ok = stage.IsStagedShiftedBottomTopStartArcShape(target)

	case *ShiftedBottomTopStartArcShapeGrid:
		ok = stage.IsStagedShiftedBottomTopStartArcShapeGrid(target)

	case *ShiftedLeftGrowthCurve2DRibbon:
		ok = stage.IsStagedShiftedLeftGrowthCurve2DRibbon(target)

	case *ShiftedLeftGrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedShiftedLeftGrowthCurve2DRibbonEndShape(target)

	case *ShiftedLeftGrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedShiftedLeftGrowthCurve2DRibbonStartShape(target)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbon:
		ok = stage.IsStagedShiftedLeftPartiallyGrowthCurve2DRibbon(target)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(target)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(target)

	case *ShiftedLeftStackGrowthCurveEndArcShape:
		ok = stage.IsStagedShiftedLeftStackGrowthCurveEndArcShape(target)

	case *ShiftedLeftStackGrowthCurveStartArcShape:
		ok = stage.IsStagedShiftedLeftStackGrowthCurveStartArcShape(target)

	case *ShiftedLeftStackNormalVector:
		ok = stage.IsStagedShiftedLeftStackNormalVector(target)

	case *ShiftedLeftStackOfGrowthCurve:
		ok = stage.IsStagedShiftedLeftStackOfGrowthCurve(target)

	case *ShiftedLeftStackOfNormalVector:
		ok = stage.IsStagedShiftedLeftStackOfNormalVector(target)

	case *ShiftedRightGrowthCurve2DRibbon:
		ok = stage.IsStagedShiftedRightGrowthCurve2DRibbon(target)

	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedShiftedRightGrowthCurve2DRibbonEndShape(target)

	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedShiftedRightGrowthCurve2DRibbonStartShape(target)

	case *StackGrowthCurve2DEndHalfwayArcShape:
		ok = stage.IsStagedStackGrowthCurve2DEndHalfwayArcShape(target)

	case *StackGrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedStackGrowthCurve2DRibbonEndShape(target)

	case *StackGrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedStackGrowthCurve2DRibbonStartShape(target)

	case *StackGrowthCurve2DStartHalfwayArcShape:
		ok = stage.IsStagedStackGrowthCurve2DStartHalfwayArcShape(target)

	case *StackOfGrowthCurve2D:
		ok = stage.IsStagedStackOfGrowthCurve2D(target)

	case *StackOfGrowthCurve2DByGrowthVector:
		ok = stage.IsStagedStackOfGrowthCurve2DByGrowthVector(target)

	case *StackOfGrowthCurve2DRibbon:
		ok = stage.IsStagedStackOfGrowthCurve2DRibbon(target)

	case *StackOfPartiallyRotatedTorusShape:
		ok = stage.IsStagedStackOfPartiallyRotatedTorusShape(target)

	case *StackOfRotatedGrowthCurve2D:
		ok = stage.IsStagedStackOfRotatedGrowthCurve2D(target)

	case *StackOfRotatedGrowthCurve2DRibbon:
		ok = stage.IsStagedStackOfRotatedGrowthCurve2DRibbon(target)

	case *StackRotatedGrowthCurve2DEndArcShape:
		ok = stage.IsStagedStackRotatedGrowthCurve2DEndArcShape(target)

	case *StackRotatedGrowthCurve2DRibbonEndShape:
		ok = stage.IsStagedStackRotatedGrowthCurve2DRibbonEndShape(target)

	case *StackRotatedGrowthCurve2DRibbonStartShape:
		ok = stage.IsStagedStackRotatedGrowthCurve2DRibbonStartShape(target)

	case *StackRotatedGrowthCurve2DStartArcShape:
		ok = stage.IsStagedStackRotatedGrowthCurve2DStartArcShape(target)

	case *StartArcShape:
		ok = stage.IsStagedStartArcShape(target)

	case *StartArcShapeGrid:
		ok = stage.IsStagedStartArcShapeGrid(target)

	case *StartHalfwayArcShape:
		ok = stage.IsStagedStartHalfwayArcShape(target)

	case *StartHalfwayArcShapeGrid:
		ok = stage.IsStagedStartHalfwayArcShapeGrid(target)

	case *StemCylinder3DShape:
		ok = stage.IsStagedStemCylinder3DShape(target)

	case *Stool2DDiagram:
		ok = stage.IsStagedStool2DDiagram(target)

	case *Stool3DDiagram:
		ok = stage.IsStagedStool3DDiagram(target)

	case *StoolAbstract:
		ok = stage.IsStagedStoolAbstract(target)

	case *TiledFloor3DShape:
		ok = stage.IsStagedTiledFloor3DShape(target)

	case *TopEndArcShape:
		ok = stage.IsStagedTopEndArcShape(target)

	case *TopEndArcShapeGrid:
		ok = stage.IsStagedTopEndArcShapeGrid(target)

	case *TopEndHalfwayArcShape:
		ok = stage.IsStagedTopEndHalfwayArcShape(target)

	case *TopEndHalfwayArcShapeGrid:
		ok = stage.IsStagedTopEndHalfwayArcShapeGrid(target)

	case *TopGrowthCurve2D:
		ok = stage.IsStagedTopGrowthCurve2D(target)

	case *TopMidArcVectorShape:
		ok = stage.IsStagedTopMidArcVectorShape(target)

	case *TopMidArcVectorShapeGrid:
		ok = stage.IsStagedTopMidArcVectorShapeGrid(target)

	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		ok = stage.IsStagedTopStackGrowthCurve2DEndHalfwayArcShape(target)

	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		ok = stage.IsStagedTopStackGrowthCurve2DStartHalfwayArcShape(target)

	case *TopStackOfGrowthCurve2D:
		ok = stage.IsStagedTopStackOfGrowthCurve2D(target)

	case *TopStackOfRotatedGrowthCurve2D:
		ok = stage.IsStagedTopStackOfRotatedGrowthCurve2D(target)

	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		ok = stage.IsStagedTopStackOfRotatedGrowthCurve2DEndArcShape(target)

	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		ok = stage.IsStagedTopStackOfRotatedGrowthCurve2DStartArcShape(target)

	case *TopStartArcShape:
		ok = stage.IsStagedTopStartArcShape(target)

	case *TopStartArcShapeGrid:
		ok = stage.IsStagedTopStartArcShapeGrid(target)

	case *TopStartHalfwayArcShape:
		ok = stage.IsStagedTopStartHalfwayArcShape(target)

	case *TopStartHalfwayArcShapeGrid:
		ok = stage.IsStagedTopStartHalfwayArcShapeGrid(target)

	case *Torus3DShape:
		ok = stage.IsStagedTorus3DShape(target)

	case *TorusEdge3DShape:
		ok = stage.IsStagedTorusEdge3DShape(target)

	case *TorusStackShape:
		ok = stage.IsStagedTorusStackShape(target)

	case *Vase2DDiagram:
		ok = stage.IsStagedVase2DDiagram(target)

	case *Vase3DDiagram:
		ok = stage.IsStagedVase3DDiagram(target)

	case *VaseAbstract:
		ok = stage.IsStagedVaseAbstract(target)

	case *VerticalTorusStackShape:
		ok = stage.IsStagedVerticalTorusStackShape(target)

	case *VolumeKey3DShape:
		ok = stage.IsStagedVolumeKey3DShape(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedAngle0Shape(angle0shape *Angle0Shape) (ok bool) {

	_, ok = stage.Angle0Shapes[angle0shape]

	return
}

func (stage *Stage) IsStagedArcNormalVectorShape(arcnormalvectorshape *ArcNormalVectorShape) (ok bool) {

	_, ok = stage.ArcNormalVectorShapes[arcnormalvectorshape]

	return
}

func (stage *Stage) IsStagedArcNormalVectorShapeGrid(arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) (ok bool) {

	_, ok = stage.ArcNormalVectorShapeGrids[arcnormalvectorshapegrid]

	return
}

func (stage *Stage) IsStagedAxesShape(axesshape *AxesShape) (ok bool) {

	_, ok = stage.AxesShapes[axesshape]

	return
}

func (stage *Stage) IsStagedBaseVectorShape(basevectorshape *BaseVectorShape) (ok bool) {

	_, ok = stage.BaseVectorShapes[basevectorshape]

	return
}

func (stage *Stage) IsStagedBaseVectorShapeGrid(basevectorshapegrid *BaseVectorShapeGrid) (ok bool) {

	_, ok = stage.BaseVectorShapeGrids[basevectorshapegrid]

	return
}

func (stage *Stage) IsStagedChosenP1P2PairShape(chosenp1p2pairshape *ChosenP1P2PairShape) (ok bool) {

	_, ok = stage.ChosenP1P2PairShapes[chosenp1p2pairshape]

	return
}

func (stage *Stage) IsStagedCircleGridShape(circlegridshape *CircleGridShape) (ok bool) {

	_, ok = stage.CircleGridShapes[circlegridshape]

	return
}

func (stage *Stage) IsStagedCircumference3DShape(circumference3dshape *Circumference3DShape) (ok bool) {

	_, ok = stage.Circumference3DShapes[circumference3dshape]

	return
}

func (stage *Stage) IsStagedClock2DDiagram(clock2ddiagram *Clock2DDiagram) (ok bool) {

	_, ok = stage.Clock2DDiagrams[clock2ddiagram]

	return
}

func (stage *Stage) IsStagedClock3DDiagram(clock3ddiagram *Clock3DDiagram) (ok bool) {

	_, ok = stage.Clock3DDiagrams[clock3ddiagram]

	return
}

func (stage *Stage) IsStagedClockAbstract(clockabstract *ClockAbstract) (ok bool) {

	_, ok = stage.ClockAbstracts[clockabstract]

	return
}

func (stage *Stage) IsStagedClockTopCurveShape(clocktopcurveshape *ClockTopCurveShape) (ok bool) {

	_, ok = stage.ClockTopCurveShapes[clocktopcurveshape]

	return
}

func (stage *Stage) IsStagedCutLine3DShape(cutline3dshape *CutLine3DShape) (ok bool) {

	_, ok = stage.CutLine3DShapes[cutline3dshape]

	return
}

func (stage *Stage) IsStagedEndArcShape(endarcshape *EndArcShape) (ok bool) {

	_, ok = stage.EndArcShapes[endarcshape]

	return
}

func (stage *Stage) IsStagedEndArcShapeGrid(endarcshapegrid *EndArcShapeGrid) (ok bool) {

	_, ok = stage.EndArcShapeGrids[endarcshapegrid]

	return
}

func (stage *Stage) IsStagedEndHalfwayArcShape(endhalfwayarcshape *EndHalfwayArcShape) (ok bool) {

	_, ok = stage.EndHalfwayArcShapes[endhalfwayarcshape]

	return
}

func (stage *Stage) IsStagedEndHalfwayArcShapeGrid(endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) (ok bool) {

	_, ok = stage.EndHalfwayArcShapeGrids[endhalfwayarcshapegrid]

	return
}

func (stage *Stage) IsStagedExplanationTextShape(explanationtextshape *ExplanationTextShape) (ok bool) {

	_, ok = stage.ExplanationTextShapes[explanationtextshape]

	return
}

func (stage *Stage) IsStagedEye3DShape(eye3dshape *Eye3DShape) (ok bool) {

	_, ok = stage.Eye3DShapes[eye3dshape]

	return
}

func (stage *Stage) IsStagedEyeCornersSampledPoints3DShape(eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) (ok bool) {

	_, ok = stage.EyeCornersSampledPoints3DShapes[eyecornerssampledpoints3dshape]

	return
}

func (stage *Stage) IsStagedEyeSampledPoints3DShape(eyesampledpoints3dshape *EyeSampledPoints3DShape) (ok bool) {

	_, ok = stage.EyeSampledPoints3DShapes[eyesampledpoints3dshape]

	return
}

func (stage *Stage) IsStagedEyeSeatBottomCurveShape(eyeseatbottomcurveshape *EyeSeatBottomCurveShape) (ok bool) {

	_, ok = stage.EyeSeatBottomCurveShapes[eyeseatbottomcurveshape]

	return
}

func (stage *Stage) IsStagedEyeStoolBottomCurveShape(eyestoolbottomcurveshape *EyeStoolBottomCurveShape) (ok bool) {

	_, ok = stage.EyeStoolBottomCurveShapes[eyestoolbottomcurveshape]

	return
}

func (stage *Stage) IsStagedEyeVolume3DShape(eyevolume3dshape *EyeVolume3DShape) (ok bool) {

	_, ok = stage.EyeVolume3DShapes[eyevolume3dshape]

	return
}

func (stage *Stage) IsStagedGridPathShape(gridpathshape *GridPathShape) (ok bool) {

	_, ok = stage.GridPathShapes[gridpathshape]

	return
}

func (stage *Stage) IsStagedGrowthCurve2D(growthcurve2d *GrowthCurve2D) (ok bool) {

	_, ok = stage.GrowthCurve2Ds[growthcurve2d]

	return
}

func (stage *Stage) IsStagedGrowthCurve2DRibbon(growthcurve2dribbon *GrowthCurve2DRibbon) (ok bool) {

	_, ok = stage.GrowthCurve2DRibbons[growthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedGrowthCurve2DRibbonEndShape(growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) (ok bool) {

	_, ok = stage.GrowthCurve2DRibbonEndShapes[growthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedGrowthCurve2DRibbonStartShape(growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) (ok bool) {

	_, ok = stage.GrowthCurve2DRibbonStartShapes[growthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedGrowthCurveRhombusGridShape(growthcurverhombusgridshape *GrowthCurveRhombusGridShape) (ok bool) {

	_, ok = stage.GrowthCurveRhombusGridShapes[growthcurverhombusgridshape]

	return
}

func (stage *Stage) IsStagedGrowthCurveRhombusShape(growthcurverhombusshape *GrowthCurveRhombusShape) (ok bool) {

	_, ok = stage.GrowthCurveRhombusShapes[growthcurverhombusshape]

	return
}

func (stage *Stage) IsStagedGrowthVectorShape(growthvectorshape *GrowthVectorShape) (ok bool) {

	_, ok = stage.GrowthVectorShapes[growthvectorshape]

	return
}

func (stage *Stage) IsStagedInitialRhombusGridShape(initialrhombusgridshape *InitialRhombusGridShape) (ok bool) {

	_, ok = stage.InitialRhombusGridShapes[initialrhombusgridshape]

	return
}

func (stage *Stage) IsStagedInitialRhombusShape(initialrhombusshape *InitialRhombusShape) (ok bool) {

	_, ok = stage.InitialRhombusShapes[initialrhombusshape]

	return
}

func (stage *Stage) IsStagedKey3DShape(key3dshape *Key3DShape) (ok bool) {

	_, ok = stage.Key3DShapes[key3dshape]

	return
}

func (stage *Stage) IsStagedKeyHole3DShape(keyhole3dshape *KeyHole3DShape) (ok bool) {

	_, ok = stage.KeyHole3DShapes[keyhole3dshape]

	return
}

func (stage *Stage) IsStagedKeyHoleShape(keyholeshape *KeyHoleShape) (ok bool) {

	_, ok = stage.KeyHoleShapes[keyholeshape]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedMidArcVectorShape(midarcvectorshape *MidArcVectorShape) (ok bool) {

	_, ok = stage.MidArcVectorShapes[midarcvectorshape]

	return
}

func (stage *Stage) IsStagedMidArcVectorShapeGrid(midarcvectorshapegrid *MidArcVectorShapeGrid) (ok bool) {

	_, ok = stage.MidArcVectorShapeGrids[midarcvectorshapegrid]

	return
}

func (stage *Stage) IsStagedOriginalPoints3DShape(originalpoints3dshape *OriginalPoints3DShape) (ok bool) {

	_, ok = stage.OriginalPoints3DShapes[originalpoints3dshape]

	return
}

func (stage *Stage) IsStagedParastichyMCurves3DShape(parastichymcurves3dshape *ParastichyMCurves3DShape) (ok bool) {

	_, ok = stage.ParastichyMCurves3DShapes[parastichymcurves3dshape]

	return
}

func (stage *Stage) IsStagedParastichyNCurves3DShape(parastichyncurves3dshape *ParastichyNCurves3DShape) (ok bool) {

	_, ok = stage.ParastichyNCurves3DShapes[parastichyncurves3dshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DRibbon(partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DRibbons[partiallygrowthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DRibbonEndShape(partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DRibbonEndShapes[partiallygrowthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DRibbonStartShape(partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DRibbonStartShapes[partiallygrowthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectory(partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectorys[partiallygrowthcurve2dtrajectory]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryP1CurveShape(partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes[partiallygrowthcurve2dtrajectoryp1curveshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryP1P2(partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryP1P2s[partiallygrowthcurve2dtrajectoryp1p2]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes[partiallygrowthcurve2dtrajectoryp1p2pairlineshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryP1PointShape(partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes[partiallygrowthcurve2dtrajectoryp1pointshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryP2CurveShape(partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes[partiallygrowthcurve2dtrajectoryp2curveshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryP2PointShape(partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes[partiallygrowthcurve2dtrajectoryp2pointshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryShape(partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryShapes[partiallygrowthcurve2dtrajectoryshape]

	return
}

func (stage *Stage) IsStagedPartiallyRotatedSeatBottomCurveShape(partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) (ok bool) {

	_, ok = stage.PartiallyRotatedSeatBottomCurveShapes[partiallyrotatedseatbottomcurveshape]

	return
}

func (stage *Stage) IsStagedPartiallyRotatedSeatTopCurveShape(partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) (ok bool) {

	_, ok = stage.PartiallyRotatedSeatTopCurveShapes[partiallyrotatedseattopcurveshape]

	return
}

func (stage *Stage) IsStagedPartiallyRotatedTorusShape(partiallyrotatedtorusshape *PartiallyRotatedTorusShape) (ok bool) {

	_, ok = stage.PartiallyRotatedTorusShapes[partiallyrotatedtorusshape]

	return
}

func (stage *Stage) IsStagedPerpendicularVector(perpendicularvector *PerpendicularVector) (ok bool) {

	_, ok = stage.PerpendicularVectors[perpendicularvector]

	return
}

func (stage *Stage) IsStagedPerpendicularVectorGrid(perpendicularvectorgrid *PerpendicularVectorGrid) (ok bool) {

	_, ok = stage.PerpendicularVectorGrids[perpendicularvectorgrid]

	return
}

func (stage *Stage) IsStagedPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) (ok bool) {

	_, ok = stage.PerpendicularVectorGridHalfways[perpendicularvectorgridhalfway]

	return
}

func (stage *Stage) IsStagedPerpendicularVectorHalfway(perpendicularvectorhalfway *PerpendicularVectorHalfway) (ok bool) {

	_, ok = stage.PerpendicularVectorHalfways[perpendicularvectorhalfway]

	return
}

func (stage *Stage) IsStagedPlant2DDiagram(plant2ddiagram *Plant2DDiagram) (ok bool) {

	_, ok = stage.Plant2DDiagrams[plant2ddiagram]

	return
}

func (stage *Stage) IsStagedPlant3DDiagram(plant3ddiagram *Plant3DDiagram) (ok bool) {

	_, ok = stage.Plant3DDiagrams[plant3ddiagram]

	return
}

func (stage *Stage) IsStagedPlantAbstract(plantabstract *PlantAbstract) (ok bool) {

	_, ok = stage.PlantAbstracts[plantabstract]

	return
}

func (stage *Stage) IsStagedPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape) (ok bool) {

	_, ok = stage.PlantCircumferenceShapes[plantcircumferenceshape]

	return
}

func (stage *Stage) IsStagedPointsAndLines3DShape(pointsandlines3dshape *PointsAndLines3DShape) (ok bool) {

	_, ok = stage.PointsAndLines3DShapes[pointsandlines3dshape]

	return
}

func (stage *Stage) IsStagedPxShape(pxshape *PxShape) (ok bool) {

	_, ok = stage.PxShapes[pxshape]

	return
}

func (stage *Stage) IsStagedRendered3DShape(rendered3dshape *Rendered3DShape) (ok bool) {

	_, ok = stage.Rendered3DShapes[rendered3dshape]

	return
}

func (stage *Stage) IsStagedRhombusShape(rhombusshape *RhombusShape) (ok bool) {

	_, ok = stage.RhombusShapes[rhombusshape]

	return
}

func (stage *Stage) IsStagedRhombusStuff(rhombusstuff *RhombusStuff) (ok bool) {

	_, ok = stage.RhombusStuffs[rhombusstuff]

	return
}

func (stage *Stage) IsStagedRotatedRhombusGridShape(rotatedrhombusgridshape *RotatedRhombusGridShape) (ok bool) {

	_, ok = stage.RotatedRhombusGridShapes[rotatedrhombusgridshape]

	return
}

func (stage *Stage) IsStagedRotatedRhombusShape(rotatedrhombusshape *RotatedRhombusShape) (ok bool) {

	_, ok = stage.RotatedRhombusShapes[rotatedrhombusshape]

	return
}

func (stage *Stage) IsStagedRotatedSampledPoints3DShape(rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) (ok bool) {

	_, ok = stage.RotatedSampledPoints3DShapes[rotatedsampledpoints3dshape]

	return
}

func (stage *Stage) IsStagedRotatedSeatAndLegs3DShape(rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) (ok bool) {

	_, ok = stage.RotatedSeatAndLegs3DShapes[rotatedseatandlegs3dshape]

	return
}

func (stage *Stage) IsStagedSampledPoints3DShape(sampledpoints3dshape *SampledPoints3DShape) (ok bool) {

	_, ok = stage.SampledPoints3DShapes[sampledpoints3dshape]

	return
}

func (stage *Stage) IsStagedSeat3DShape(seat3dshape *Seat3DShape) (ok bool) {

	_, ok = stage.Seat3DShapes[seat3dshape]

	return
}

func (stage *Stage) IsStagedSeatAndLegs3DShape(seatandlegs3dshape *SeatAndLegs3DShape) (ok bool) {

	_, ok = stage.SeatAndLegs3DShapes[seatandlegs3dshape]

	return
}

func (stage *Stage) IsStagedSeatBottomCurveShape(seatbottomcurveshape *SeatBottomCurveShape) (ok bool) {

	_, ok = stage.SeatBottomCurveShapes[seatbottomcurveshape]

	return
}

func (stage *Stage) IsStagedSeatTopCurveShape(seattopcurveshape *SeatTopCurveShape) (ok bool) {

	_, ok = stage.SeatTopCurveShapes[seattopcurveshape]

	return
}

func (stage *Stage) IsStagedShiftedBottomTopStartArcShape(shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) (ok bool) {

	_, ok = stage.ShiftedBottomTopStartArcShapes[shiftedbottomtopstartarcshape]

	return
}

func (stage *Stage) IsStagedShiftedBottomTopStartArcShapeGrid(shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) (ok bool) {

	_, ok = stage.ShiftedBottomTopStartArcShapeGrids[shiftedbottomtopstartarcshapegrid]

	return
}

func (stage *Stage) IsStagedShiftedLeftGrowthCurve2DRibbon(shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) (ok bool) {

	_, ok = stage.ShiftedLeftGrowthCurve2DRibbons[shiftedleftgrowthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedShiftedLeftGrowthCurve2DRibbonEndShape(shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) (ok bool) {

	_, ok = stage.ShiftedLeftGrowthCurve2DRibbonEndShapes[shiftedleftgrowthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedShiftedLeftGrowthCurve2DRibbonStartShape(shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) (ok bool) {

	_, ok = stage.ShiftedLeftGrowthCurve2DRibbonStartShapes[shiftedleftgrowthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedShiftedLeftPartiallyGrowthCurve2DRibbon(shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) (ok bool) {

	_, ok = stage.ShiftedLeftPartiallyGrowthCurve2DRibbons[shiftedleftpartiallygrowthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) (ok bool) {

	_, ok = stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes[shiftedleftpartiallygrowthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) (ok bool) {

	_, ok = stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes[shiftedleftpartiallygrowthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedShiftedLeftStackGrowthCurveEndArcShape(shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) (ok bool) {

	_, ok = stage.ShiftedLeftStackGrowthCurveEndArcShapes[shiftedleftstackgrowthcurveendarcshape]

	return
}

func (stage *Stage) IsStagedShiftedLeftStackGrowthCurveStartArcShape(shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) (ok bool) {

	_, ok = stage.ShiftedLeftStackGrowthCurveStartArcShapes[shiftedleftstackgrowthcurvestartarcshape]

	return
}

func (stage *Stage) IsStagedShiftedLeftStackNormalVector(shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) (ok bool) {

	_, ok = stage.ShiftedLeftStackNormalVectors[shiftedleftstacknormalvector]

	return
}

func (stage *Stage) IsStagedShiftedLeftStackOfGrowthCurve(shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) (ok bool) {

	_, ok = stage.ShiftedLeftStackOfGrowthCurves[shiftedleftstackofgrowthcurve]

	return
}

func (stage *Stage) IsStagedShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) (ok bool) {

	_, ok = stage.ShiftedLeftStackOfNormalVectors[shiftedleftstackofnormalvector]

	return
}

func (stage *Stage) IsStagedShiftedRightGrowthCurve2DRibbon(shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) (ok bool) {

	_, ok = stage.ShiftedRightGrowthCurve2DRibbons[shiftedrightgrowthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedShiftedRightGrowthCurve2DRibbonEndShape(shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) (ok bool) {

	_, ok = stage.ShiftedRightGrowthCurve2DRibbonEndShapes[shiftedrightgrowthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedShiftedRightGrowthCurve2DRibbonStartShape(shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) (ok bool) {

	_, ok = stage.ShiftedRightGrowthCurve2DRibbonStartShapes[shiftedrightgrowthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedStackGrowthCurve2DEndHalfwayArcShape(stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) (ok bool) {

	_, ok = stage.StackGrowthCurve2DEndHalfwayArcShapes[stackgrowthcurve2dendhalfwayarcshape]

	return
}

func (stage *Stage) IsStagedStackGrowthCurve2DRibbonEndShape(stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) (ok bool) {

	_, ok = stage.StackGrowthCurve2DRibbonEndShapes[stackgrowthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedStackGrowthCurve2DRibbonStartShape(stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) (ok bool) {

	_, ok = stage.StackGrowthCurve2DRibbonStartShapes[stackgrowthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedStackGrowthCurve2DStartHalfwayArcShape(stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) (ok bool) {

	_, ok = stage.StackGrowthCurve2DStartHalfwayArcShapes[stackgrowthcurve2dstarthalfwayarcshape]

	return
}

func (stage *Stage) IsStagedStackOfGrowthCurve2D(stackofgrowthcurve2d *StackOfGrowthCurve2D) (ok bool) {

	_, ok = stage.StackOfGrowthCurve2Ds[stackofgrowthcurve2d]

	return
}

func (stage *Stage) IsStagedStackOfGrowthCurve2DByGrowthVector(stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) (ok bool) {

	_, ok = stage.StackOfGrowthCurve2DByGrowthVectors[stackofgrowthcurve2dbygrowthvector]

	return
}

func (stage *Stage) IsStagedStackOfGrowthCurve2DRibbon(stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) (ok bool) {

	_, ok = stage.StackOfGrowthCurve2DRibbons[stackofgrowthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedStackOfPartiallyRotatedTorusShape(stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) (ok bool) {

	_, ok = stage.StackOfPartiallyRotatedTorusShapes[stackofpartiallyrotatedtorusshape]

	return
}

func (stage *Stage) IsStagedStackOfRotatedGrowthCurve2D(stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) (ok bool) {

	_, ok = stage.StackOfRotatedGrowthCurve2Ds[stackofrotatedgrowthcurve2d]

	return
}

func (stage *Stage) IsStagedStackOfRotatedGrowthCurve2DRibbon(stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) (ok bool) {

	_, ok = stage.StackOfRotatedGrowthCurve2DRibbons[stackofrotatedgrowthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedStackRotatedGrowthCurve2DEndArcShape(stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) (ok bool) {

	_, ok = stage.StackRotatedGrowthCurve2DEndArcShapes[stackrotatedgrowthcurve2dendarcshape]

	return
}

func (stage *Stage) IsStagedStackRotatedGrowthCurve2DRibbonEndShape(stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) (ok bool) {

	_, ok = stage.StackRotatedGrowthCurve2DRibbonEndShapes[stackrotatedgrowthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedStackRotatedGrowthCurve2DRibbonStartShape(stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) (ok bool) {

	_, ok = stage.StackRotatedGrowthCurve2DRibbonStartShapes[stackrotatedgrowthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedStackRotatedGrowthCurve2DStartArcShape(stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) (ok bool) {

	_, ok = stage.StackRotatedGrowthCurve2DStartArcShapes[stackrotatedgrowthcurve2dstartarcshape]

	return
}

func (stage *Stage) IsStagedStartArcShape(startarcshape *StartArcShape) (ok bool) {

	_, ok = stage.StartArcShapes[startarcshape]

	return
}

func (stage *Stage) IsStagedStartArcShapeGrid(startarcshapegrid *StartArcShapeGrid) (ok bool) {

	_, ok = stage.StartArcShapeGrids[startarcshapegrid]

	return
}

func (stage *Stage) IsStagedStartHalfwayArcShape(starthalfwayarcshape *StartHalfwayArcShape) (ok bool) {

	_, ok = stage.StartHalfwayArcShapes[starthalfwayarcshape]

	return
}

func (stage *Stage) IsStagedStartHalfwayArcShapeGrid(starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) (ok bool) {

	_, ok = stage.StartHalfwayArcShapeGrids[starthalfwayarcshapegrid]

	return
}

func (stage *Stage) IsStagedStemCylinder3DShape(stemcylinder3dshape *StemCylinder3DShape) (ok bool) {

	_, ok = stage.StemCylinder3DShapes[stemcylinder3dshape]

	return
}

func (stage *Stage) IsStagedStool2DDiagram(stool2ddiagram *Stool2DDiagram) (ok bool) {

	_, ok = stage.Stool2DDiagrams[stool2ddiagram]

	return
}

func (stage *Stage) IsStagedStool3DDiagram(stool3ddiagram *Stool3DDiagram) (ok bool) {

	_, ok = stage.Stool3DDiagrams[stool3ddiagram]

	return
}

func (stage *Stage) IsStagedStoolAbstract(stoolabstract *StoolAbstract) (ok bool) {

	_, ok = stage.StoolAbstracts[stoolabstract]

	return
}

func (stage *Stage) IsStagedTiledFloor3DShape(tiledfloor3dshape *TiledFloor3DShape) (ok bool) {

	_, ok = stage.TiledFloor3DShapes[tiledfloor3dshape]

	return
}

func (stage *Stage) IsStagedTopEndArcShape(topendarcshape *TopEndArcShape) (ok bool) {

	_, ok = stage.TopEndArcShapes[topendarcshape]

	return
}

func (stage *Stage) IsStagedTopEndArcShapeGrid(topendarcshapegrid *TopEndArcShapeGrid) (ok bool) {

	_, ok = stage.TopEndArcShapeGrids[topendarcshapegrid]

	return
}

func (stage *Stage) IsStagedTopEndHalfwayArcShape(topendhalfwayarcshape *TopEndHalfwayArcShape) (ok bool) {

	_, ok = stage.TopEndHalfwayArcShapes[topendhalfwayarcshape]

	return
}

func (stage *Stage) IsStagedTopEndHalfwayArcShapeGrid(topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) (ok bool) {

	_, ok = stage.TopEndHalfwayArcShapeGrids[topendhalfwayarcshapegrid]

	return
}

func (stage *Stage) IsStagedTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D) (ok bool) {

	_, ok = stage.TopGrowthCurve2Ds[topgrowthcurve2d]

	return
}

func (stage *Stage) IsStagedTopMidArcVectorShape(topmidarcvectorshape *TopMidArcVectorShape) (ok bool) {

	_, ok = stage.TopMidArcVectorShapes[topmidarcvectorshape]

	return
}

func (stage *Stage) IsStagedTopMidArcVectorShapeGrid(topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) (ok bool) {

	_, ok = stage.TopMidArcVectorShapeGrids[topmidarcvectorshapegrid]

	return
}

func (stage *Stage) IsStagedTopStackGrowthCurve2DEndHalfwayArcShape(topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) (ok bool) {

	_, ok = stage.TopStackGrowthCurve2DEndHalfwayArcShapes[topstackgrowthcurve2dendhalfwayarcshape]

	return
}

func (stage *Stage) IsStagedTopStackGrowthCurve2DStartHalfwayArcShape(topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) (ok bool) {

	_, ok = stage.TopStackGrowthCurve2DStartHalfwayArcShapes[topstackgrowthcurve2dstarthalfwayarcshape]

	return
}

func (stage *Stage) IsStagedTopStackOfGrowthCurve2D(topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) (ok bool) {

	_, ok = stage.TopStackOfGrowthCurve2Ds[topstackofgrowthcurve2d]

	return
}

func (stage *Stage) IsStagedTopStackOfRotatedGrowthCurve2D(topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) (ok bool) {

	_, ok = stage.TopStackOfRotatedGrowthCurve2Ds[topstackofrotatedgrowthcurve2d]

	return
}

func (stage *Stage) IsStagedTopStackOfRotatedGrowthCurve2DEndArcShape(topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) (ok bool) {

	_, ok = stage.TopStackOfRotatedGrowthCurve2DEndArcShapes[topstackofrotatedgrowthcurve2dendarcshape]

	return
}

func (stage *Stage) IsStagedTopStackOfRotatedGrowthCurve2DStartArcShape(topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) (ok bool) {

	_, ok = stage.TopStackOfRotatedGrowthCurve2DStartArcShapes[topstackofrotatedgrowthcurve2dstartarcshape]

	return
}

func (stage *Stage) IsStagedTopStartArcShape(topstartarcshape *TopStartArcShape) (ok bool) {

	_, ok = stage.TopStartArcShapes[topstartarcshape]

	return
}

func (stage *Stage) IsStagedTopStartArcShapeGrid(topstartarcshapegrid *TopStartArcShapeGrid) (ok bool) {

	_, ok = stage.TopStartArcShapeGrids[topstartarcshapegrid]

	return
}

func (stage *Stage) IsStagedTopStartHalfwayArcShape(topstarthalfwayarcshape *TopStartHalfwayArcShape) (ok bool) {

	_, ok = stage.TopStartHalfwayArcShapes[topstarthalfwayarcshape]

	return
}

func (stage *Stage) IsStagedTopStartHalfwayArcShapeGrid(topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) (ok bool) {

	_, ok = stage.TopStartHalfwayArcShapeGrids[topstarthalfwayarcshapegrid]

	return
}

func (stage *Stage) IsStagedTorus3DShape(torus3dshape *Torus3DShape) (ok bool) {

	_, ok = stage.Torus3DShapes[torus3dshape]

	return
}

func (stage *Stage) IsStagedTorusEdge3DShape(torusedge3dshape *TorusEdge3DShape) (ok bool) {

	_, ok = stage.TorusEdge3DShapes[torusedge3dshape]

	return
}

func (stage *Stage) IsStagedTorusStackShape(torusstackshape *TorusStackShape) (ok bool) {

	_, ok = stage.TorusStackShapes[torusstackshape]

	return
}

func (stage *Stage) IsStagedVase2DDiagram(vase2ddiagram *Vase2DDiagram) (ok bool) {

	_, ok = stage.Vase2DDiagrams[vase2ddiagram]

	return
}

func (stage *Stage) IsStagedVase3DDiagram(vase3ddiagram *Vase3DDiagram) (ok bool) {

	_, ok = stage.Vase3DDiagrams[vase3ddiagram]

	return
}

func (stage *Stage) IsStagedVaseAbstract(vaseabstract *VaseAbstract) (ok bool) {

	_, ok = stage.VaseAbstracts[vaseabstract]

	return
}

func (stage *Stage) IsStagedVerticalTorusStackShape(verticaltorusstackshape *VerticalTorusStackShape) (ok bool) {

	_, ok = stage.VerticalTorusStackShapes[verticaltorusstackshape]

	return
}

func (stage *Stage) IsStagedVolumeKey3DShape(volumekey3dshape *VolumeKey3DShape) (ok bool) {

	_, ok = stage.VolumeKey3DShapes[volumekey3dshape]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Angle0Shape:
		stage.StageBranchAngle0Shape(target)

	case *ArcNormalVectorShape:
		stage.StageBranchArcNormalVectorShape(target)

	case *ArcNormalVectorShapeGrid:
		stage.StageBranchArcNormalVectorShapeGrid(target)

	case *AxesShape:
		stage.StageBranchAxesShape(target)

	case *BaseVectorShape:
		stage.StageBranchBaseVectorShape(target)

	case *BaseVectorShapeGrid:
		stage.StageBranchBaseVectorShapeGrid(target)

	case *ChosenP1P2PairShape:
		stage.StageBranchChosenP1P2PairShape(target)

	case *CircleGridShape:
		stage.StageBranchCircleGridShape(target)

	case *Circumference3DShape:
		stage.StageBranchCircumference3DShape(target)

	case *Clock2DDiagram:
		stage.StageBranchClock2DDiagram(target)

	case *Clock3DDiagram:
		stage.StageBranchClock3DDiagram(target)

	case *ClockAbstract:
		stage.StageBranchClockAbstract(target)

	case *ClockTopCurveShape:
		stage.StageBranchClockTopCurveShape(target)

	case *CutLine3DShape:
		stage.StageBranchCutLine3DShape(target)

	case *EndArcShape:
		stage.StageBranchEndArcShape(target)

	case *EndArcShapeGrid:
		stage.StageBranchEndArcShapeGrid(target)

	case *EndHalfwayArcShape:
		stage.StageBranchEndHalfwayArcShape(target)

	case *EndHalfwayArcShapeGrid:
		stage.StageBranchEndHalfwayArcShapeGrid(target)

	case *ExplanationTextShape:
		stage.StageBranchExplanationTextShape(target)

	case *Eye3DShape:
		stage.StageBranchEye3DShape(target)

	case *EyeCornersSampledPoints3DShape:
		stage.StageBranchEyeCornersSampledPoints3DShape(target)

	case *EyeSampledPoints3DShape:
		stage.StageBranchEyeSampledPoints3DShape(target)

	case *EyeSeatBottomCurveShape:
		stage.StageBranchEyeSeatBottomCurveShape(target)

	case *EyeStoolBottomCurveShape:
		stage.StageBranchEyeStoolBottomCurveShape(target)

	case *EyeVolume3DShape:
		stage.StageBranchEyeVolume3DShape(target)

	case *GridPathShape:
		stage.StageBranchGridPathShape(target)

	case *GrowthCurve2D:
		stage.StageBranchGrowthCurve2D(target)

	case *GrowthCurve2DRibbon:
		stage.StageBranchGrowthCurve2DRibbon(target)

	case *GrowthCurve2DRibbonEndShape:
		stage.StageBranchGrowthCurve2DRibbonEndShape(target)

	case *GrowthCurve2DRibbonStartShape:
		stage.StageBranchGrowthCurve2DRibbonStartShape(target)

	case *GrowthCurveRhombusGridShape:
		stage.StageBranchGrowthCurveRhombusGridShape(target)

	case *GrowthCurveRhombusShape:
		stage.StageBranchGrowthCurveRhombusShape(target)

	case *GrowthVectorShape:
		stage.StageBranchGrowthVectorShape(target)

	case *InitialRhombusGridShape:
		stage.StageBranchInitialRhombusGridShape(target)

	case *InitialRhombusShape:
		stage.StageBranchInitialRhombusShape(target)

	case *Key3DShape:
		stage.StageBranchKey3DShape(target)

	case *KeyHole3DShape:
		stage.StageBranchKeyHole3DShape(target)

	case *KeyHoleShape:
		stage.StageBranchKeyHoleShape(target)

	case *Library:
		stage.StageBranchLibrary(target)

	case *MidArcVectorShape:
		stage.StageBranchMidArcVectorShape(target)

	case *MidArcVectorShapeGrid:
		stage.StageBranchMidArcVectorShapeGrid(target)

	case *OriginalPoints3DShape:
		stage.StageBranchOriginalPoints3DShape(target)

	case *ParastichyMCurves3DShape:
		stage.StageBranchParastichyMCurves3DShape(target)

	case *ParastichyNCurves3DShape:
		stage.StageBranchParastichyNCurves3DShape(target)

	case *PartiallyGrowthCurve2DRibbon:
		stage.StageBranchPartiallyGrowthCurve2DRibbon(target)

	case *PartiallyGrowthCurve2DRibbonEndShape:
		stage.StageBranchPartiallyGrowthCurve2DRibbonEndShape(target)

	case *PartiallyGrowthCurve2DRibbonStartShape:
		stage.StageBranchPartiallyGrowthCurve2DRibbonStartShape(target)

	case *PartiallyGrowthCurve2DTrajectory:
		stage.StageBranchPartiallyGrowthCurve2DTrajectory(target)

	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		stage.StageBranchPartiallyGrowthCurve2DTrajectoryP1CurveShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		stage.StageBranchPartiallyGrowthCurve2DTrajectoryP1P2(target)

	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		stage.StageBranchPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		stage.StageBranchPartiallyGrowthCurve2DTrajectoryP1PointShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		stage.StageBranchPartiallyGrowthCurve2DTrajectoryP2CurveShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		stage.StageBranchPartiallyGrowthCurve2DTrajectoryP2PointShape(target)

	case *PartiallyGrowthCurve2DTrajectoryShape:
		stage.StageBranchPartiallyGrowthCurve2DTrajectoryShape(target)

	case *PartiallyRotatedSeatBottomCurveShape:
		stage.StageBranchPartiallyRotatedSeatBottomCurveShape(target)

	case *PartiallyRotatedSeatTopCurveShape:
		stage.StageBranchPartiallyRotatedSeatTopCurveShape(target)

	case *PartiallyRotatedTorusShape:
		stage.StageBranchPartiallyRotatedTorusShape(target)

	case *PerpendicularVector:
		stage.StageBranchPerpendicularVector(target)

	case *PerpendicularVectorGrid:
		stage.StageBranchPerpendicularVectorGrid(target)

	case *PerpendicularVectorGridHalfway:
		stage.StageBranchPerpendicularVectorGridHalfway(target)

	case *PerpendicularVectorHalfway:
		stage.StageBranchPerpendicularVectorHalfway(target)

	case *Plant2DDiagram:
		stage.StageBranchPlant2DDiagram(target)

	case *Plant3DDiagram:
		stage.StageBranchPlant3DDiagram(target)

	case *PlantAbstract:
		stage.StageBranchPlantAbstract(target)

	case *PlantCircumferenceShape:
		stage.StageBranchPlantCircumferenceShape(target)

	case *PointsAndLines3DShape:
		stage.StageBranchPointsAndLines3DShape(target)

	case *PxShape:
		stage.StageBranchPxShape(target)

	case *Rendered3DShape:
		stage.StageBranchRendered3DShape(target)

	case *RhombusShape:
		stage.StageBranchRhombusShape(target)

	case *RhombusStuff:
		stage.StageBranchRhombusStuff(target)

	case *RotatedRhombusGridShape:
		stage.StageBranchRotatedRhombusGridShape(target)

	case *RotatedRhombusShape:
		stage.StageBranchRotatedRhombusShape(target)

	case *RotatedSampledPoints3DShape:
		stage.StageBranchRotatedSampledPoints3DShape(target)

	case *RotatedSeatAndLegs3DShape:
		stage.StageBranchRotatedSeatAndLegs3DShape(target)

	case *SampledPoints3DShape:
		stage.StageBranchSampledPoints3DShape(target)

	case *Seat3DShape:
		stage.StageBranchSeat3DShape(target)

	case *SeatAndLegs3DShape:
		stage.StageBranchSeatAndLegs3DShape(target)

	case *SeatBottomCurveShape:
		stage.StageBranchSeatBottomCurveShape(target)

	case *SeatTopCurveShape:
		stage.StageBranchSeatTopCurveShape(target)

	case *ShiftedBottomTopStartArcShape:
		stage.StageBranchShiftedBottomTopStartArcShape(target)

	case *ShiftedBottomTopStartArcShapeGrid:
		stage.StageBranchShiftedBottomTopStartArcShapeGrid(target)

	case *ShiftedLeftGrowthCurve2DRibbon:
		stage.StageBranchShiftedLeftGrowthCurve2DRibbon(target)

	case *ShiftedLeftGrowthCurve2DRibbonEndShape:
		stage.StageBranchShiftedLeftGrowthCurve2DRibbonEndShape(target)

	case *ShiftedLeftGrowthCurve2DRibbonStartShape:
		stage.StageBranchShiftedLeftGrowthCurve2DRibbonStartShape(target)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbon:
		stage.StageBranchShiftedLeftPartiallyGrowthCurve2DRibbon(target)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape:
		stage.StageBranchShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(target)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape:
		stage.StageBranchShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(target)

	case *ShiftedLeftStackGrowthCurveEndArcShape:
		stage.StageBranchShiftedLeftStackGrowthCurveEndArcShape(target)

	case *ShiftedLeftStackGrowthCurveStartArcShape:
		stage.StageBranchShiftedLeftStackGrowthCurveStartArcShape(target)

	case *ShiftedLeftStackNormalVector:
		stage.StageBranchShiftedLeftStackNormalVector(target)

	case *ShiftedLeftStackOfGrowthCurve:
		stage.StageBranchShiftedLeftStackOfGrowthCurve(target)

	case *ShiftedLeftStackOfNormalVector:
		stage.StageBranchShiftedLeftStackOfNormalVector(target)

	case *ShiftedRightGrowthCurve2DRibbon:
		stage.StageBranchShiftedRightGrowthCurve2DRibbon(target)

	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		stage.StageBranchShiftedRightGrowthCurve2DRibbonEndShape(target)

	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		stage.StageBranchShiftedRightGrowthCurve2DRibbonStartShape(target)

	case *StackGrowthCurve2DEndHalfwayArcShape:
		stage.StageBranchStackGrowthCurve2DEndHalfwayArcShape(target)

	case *StackGrowthCurve2DRibbonEndShape:
		stage.StageBranchStackGrowthCurve2DRibbonEndShape(target)

	case *StackGrowthCurve2DRibbonStartShape:
		stage.StageBranchStackGrowthCurve2DRibbonStartShape(target)

	case *StackGrowthCurve2DStartHalfwayArcShape:
		stage.StageBranchStackGrowthCurve2DStartHalfwayArcShape(target)

	case *StackOfGrowthCurve2D:
		stage.StageBranchStackOfGrowthCurve2D(target)

	case *StackOfGrowthCurve2DByGrowthVector:
		stage.StageBranchStackOfGrowthCurve2DByGrowthVector(target)

	case *StackOfGrowthCurve2DRibbon:
		stage.StageBranchStackOfGrowthCurve2DRibbon(target)

	case *StackOfPartiallyRotatedTorusShape:
		stage.StageBranchStackOfPartiallyRotatedTorusShape(target)

	case *StackOfRotatedGrowthCurve2D:
		stage.StageBranchStackOfRotatedGrowthCurve2D(target)

	case *StackOfRotatedGrowthCurve2DRibbon:
		stage.StageBranchStackOfRotatedGrowthCurve2DRibbon(target)

	case *StackRotatedGrowthCurve2DEndArcShape:
		stage.StageBranchStackRotatedGrowthCurve2DEndArcShape(target)

	case *StackRotatedGrowthCurve2DRibbonEndShape:
		stage.StageBranchStackRotatedGrowthCurve2DRibbonEndShape(target)

	case *StackRotatedGrowthCurve2DRibbonStartShape:
		stage.StageBranchStackRotatedGrowthCurve2DRibbonStartShape(target)

	case *StackRotatedGrowthCurve2DStartArcShape:
		stage.StageBranchStackRotatedGrowthCurve2DStartArcShape(target)

	case *StartArcShape:
		stage.StageBranchStartArcShape(target)

	case *StartArcShapeGrid:
		stage.StageBranchStartArcShapeGrid(target)

	case *StartHalfwayArcShape:
		stage.StageBranchStartHalfwayArcShape(target)

	case *StartHalfwayArcShapeGrid:
		stage.StageBranchStartHalfwayArcShapeGrid(target)

	case *StemCylinder3DShape:
		stage.StageBranchStemCylinder3DShape(target)

	case *Stool2DDiagram:
		stage.StageBranchStool2DDiagram(target)

	case *Stool3DDiagram:
		stage.StageBranchStool3DDiagram(target)

	case *StoolAbstract:
		stage.StageBranchStoolAbstract(target)

	case *TiledFloor3DShape:
		stage.StageBranchTiledFloor3DShape(target)

	case *TopEndArcShape:
		stage.StageBranchTopEndArcShape(target)

	case *TopEndArcShapeGrid:
		stage.StageBranchTopEndArcShapeGrid(target)

	case *TopEndHalfwayArcShape:
		stage.StageBranchTopEndHalfwayArcShape(target)

	case *TopEndHalfwayArcShapeGrid:
		stage.StageBranchTopEndHalfwayArcShapeGrid(target)

	case *TopGrowthCurve2D:
		stage.StageBranchTopGrowthCurve2D(target)

	case *TopMidArcVectorShape:
		stage.StageBranchTopMidArcVectorShape(target)

	case *TopMidArcVectorShapeGrid:
		stage.StageBranchTopMidArcVectorShapeGrid(target)

	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		stage.StageBranchTopStackGrowthCurve2DEndHalfwayArcShape(target)

	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		stage.StageBranchTopStackGrowthCurve2DStartHalfwayArcShape(target)

	case *TopStackOfGrowthCurve2D:
		stage.StageBranchTopStackOfGrowthCurve2D(target)

	case *TopStackOfRotatedGrowthCurve2D:
		stage.StageBranchTopStackOfRotatedGrowthCurve2D(target)

	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		stage.StageBranchTopStackOfRotatedGrowthCurve2DEndArcShape(target)

	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		stage.StageBranchTopStackOfRotatedGrowthCurve2DStartArcShape(target)

	case *TopStartArcShape:
		stage.StageBranchTopStartArcShape(target)

	case *TopStartArcShapeGrid:
		stage.StageBranchTopStartArcShapeGrid(target)

	case *TopStartHalfwayArcShape:
		stage.StageBranchTopStartHalfwayArcShape(target)

	case *TopStartHalfwayArcShapeGrid:
		stage.StageBranchTopStartHalfwayArcShapeGrid(target)

	case *Torus3DShape:
		stage.StageBranchTorus3DShape(target)

	case *TorusEdge3DShape:
		stage.StageBranchTorusEdge3DShape(target)

	case *TorusStackShape:
		stage.StageBranchTorusStackShape(target)

	case *Vase2DDiagram:
		stage.StageBranchVase2DDiagram(target)

	case *Vase3DDiagram:
		stage.StageBranchVase3DDiagram(target)

	case *VaseAbstract:
		stage.StageBranchVaseAbstract(target)

	case *VerticalTorusStackShape:
		stage.StageBranchVerticalTorusStackShape(target)

	case *VolumeKey3DShape:
		stage.StageBranchVolumeKey3DShape(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchAngle0Shape(angle0shape *Angle0Shape) {

	// check if instance is already staged
	if IsStaged(stage, angle0shape) {
		return
	}

	angle0shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchArcNormalVectorShape(arcnormalvectorshape *ArcNormalVectorShape) {

	// check if instance is already staged
	if IsStaged(stage, arcnormalvectorshape) {
		return
	}

	arcnormalvectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchArcNormalVectorShapeGrid(arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, arcnormalvectorshapegrid) {
		return
	}

	arcnormalvectorshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchAxesShape(axesshape *AxesShape) {

	// check if instance is already staged
	if IsStaged(stage, axesshape) {
		return
	}

	axesshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchBaseVectorShape(basevectorshape *BaseVectorShape) {

	// check if instance is already staged
	if IsStaged(stage, basevectorshape) {
		return
	}

	basevectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchBaseVectorShapeGrid(basevectorshapegrid *BaseVectorShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, basevectorshapegrid) {
		return
	}

	basevectorshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchChosenP1P2PairShape(chosenp1p2pairshape *ChosenP1P2PairShape) {

	// check if instance is already staged
	if IsStaged(stage, chosenp1p2pairshape) {
		return
	}

	chosenp1p2pairshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCircleGridShape(circlegridshape *CircleGridShape) {

	// check if instance is already staged
	if IsStaged(stage, circlegridshape) {
		return
	}

	circlegridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCircumference3DShape(circumference3dshape *Circumference3DShape) {

	// check if instance is already staged
	if IsStaged(stage, circumference3dshape) {
		return
	}

	circumference3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchClock2DDiagram(clock2ddiagram *Clock2DDiagram) {

	// check if instance is already staged
	if IsStaged(stage, clock2ddiagram) {
		return
	}

	clock2ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchClock3DDiagram(clock3ddiagram *Clock3DDiagram) {

	// check if instance is already staged
	if IsStaged(stage, clock3ddiagram) {
		return
	}

	clock3ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if clock3ddiagram.SampledPoints3DShape != nil {
		StageBranch(stage, clock3ddiagram.SampledPoints3DShape)
	}
	if clock3ddiagram.Rendered3DShape != nil {
		StageBranch(stage, clock3ddiagram.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchClockAbstract(clockabstract *ClockAbstract) {

	// check if instance is already staged
	if IsStaged(stage, clockabstract) {
		return
	}

	clockabstract.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchClockTopCurveShape(clocktopcurveshape *ClockTopCurveShape) {

	// check if instance is already staged
	if IsStaged(stage, clocktopcurveshape) {
		return
	}

	clocktopcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCutLine3DShape(cutline3dshape *CutLine3DShape) {

	// check if instance is already staged
	if IsStaged(stage, cutline3dshape) {
		return
	}

	cutline3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEndArcShape(endarcshape *EndArcShape) {

	// check if instance is already staged
	if IsStaged(stage, endarcshape) {
		return
	}

	endarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEndArcShapeGrid(endarcshapegrid *EndArcShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, endarcshapegrid) {
		return
	}

	endarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEndHalfwayArcShape(endhalfwayarcshape *EndHalfwayArcShape) {

	// check if instance is already staged
	if IsStaged(stage, endhalfwayarcshape) {
		return
	}

	endhalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEndHalfwayArcShapeGrid(endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, endhalfwayarcshapegrid) {
		return
	}

	endhalfwayarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchExplanationTextShape(explanationtextshape *ExplanationTextShape) {

	// check if instance is already staged
	if IsStaged(stage, explanationtextshape) {
		return
	}

	explanationtextshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEye3DShape(eye3dshape *Eye3DShape) {

	// check if instance is already staged
	if IsStaged(stage, eye3dshape) {
		return
	}

	eye3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEyeCornersSampledPoints3DShape(eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) {

	// check if instance is already staged
	if IsStaged(stage, eyecornerssampledpoints3dshape) {
		return
	}

	eyecornerssampledpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEyeSampledPoints3DShape(eyesampledpoints3dshape *EyeSampledPoints3DShape) {

	// check if instance is already staged
	if IsStaged(stage, eyesampledpoints3dshape) {
		return
	}

	eyesampledpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEyeSeatBottomCurveShape(eyeseatbottomcurveshape *EyeSeatBottomCurveShape) {

	// check if instance is already staged
	if IsStaged(stage, eyeseatbottomcurveshape) {
		return
	}

	eyeseatbottomcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEyeStoolBottomCurveShape(eyestoolbottomcurveshape *EyeStoolBottomCurveShape) {

	// check if instance is already staged
	if IsStaged(stage, eyestoolbottomcurveshape) {
		return
	}

	eyestoolbottomcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEyeVolume3DShape(eyevolume3dshape *EyeVolume3DShape) {

	// check if instance is already staged
	if IsStaged(stage, eyevolume3dshape) {
		return
	}

	eyevolume3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGridPathShape(gridpathshape *GridPathShape) {

	// check if instance is already staged
	if IsStaged(stage, gridpathshape) {
		return
	}

	gridpathshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGrowthCurve2D(growthcurve2d *GrowthCurve2D) {

	// check if instance is already staged
	if IsStaged(stage, growthcurve2d) {
		return
	}

	growthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGrowthCurve2DRibbon(growthcurve2dribbon *GrowthCurve2DRibbon) {

	// check if instance is already staged
	if IsStaged(stage, growthcurve2dribbon) {
		return
	}

	growthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGrowthCurve2DRibbonEndShape(growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if IsStaged(stage, growthcurve2dribbonendshape) {
		return
	}

	growthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGrowthCurve2DRibbonStartShape(growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if IsStaged(stage, growthcurve2dribbonstartshape) {
		return
	}

	growthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGrowthCurveRhombusGridShape(growthcurverhombusgridshape *GrowthCurveRhombusGridShape) {

	// check if instance is already staged
	if IsStaged(stage, growthcurverhombusgridshape) {
		return
	}

	growthcurverhombusgridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGrowthCurveRhombusShape(growthcurverhombusshape *GrowthCurveRhombusShape) {

	// check if instance is already staged
	if IsStaged(stage, growthcurverhombusshape) {
		return
	}

	growthcurverhombusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGrowthVectorShape(growthvectorshape *GrowthVectorShape) {

	// check if instance is already staged
	if IsStaged(stage, growthvectorshape) {
		return
	}

	growthvectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchInitialRhombusGridShape(initialrhombusgridshape *InitialRhombusGridShape) {

	// check if instance is already staged
	if IsStaged(stage, initialrhombusgridshape) {
		return
	}

	initialrhombusgridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchInitialRhombusShape(initialrhombusshape *InitialRhombusShape) {

	// check if instance is already staged
	if IsStaged(stage, initialrhombusshape) {
		return
	}

	initialrhombusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchKey3DShape(key3dshape *Key3DShape) {

	// check if instance is already staged
	if IsStaged(stage, key3dshape) {
		return
	}

	key3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchKeyHole3DShape(keyhole3dshape *KeyHole3DShape) {

	// check if instance is already staged
	if IsStaged(stage, keyhole3dshape) {
		return
	}

	keyhole3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchKeyHoleShape(keyholeshape *KeyHoleShape) {

	// check if instance is already staged
	if IsStaged(stage, keyholeshape) {
		return
	}

	keyholeshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchLibrary(library *Library) {

	// check if instance is already staged
	if IsStaged(stage, library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		StageBranch(stage, _library)
	}
	for _, _plantabstract := range library.Plants {
		StageBranch(stage, _plantabstract)
	}

}

func (stage *Stage) StageBranchMidArcVectorShape(midarcvectorshape *MidArcVectorShape) {

	// check if instance is already staged
	if IsStaged(stage, midarcvectorshape) {
		return
	}

	midarcvectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMidArcVectorShapeGrid(midarcvectorshapegrid *MidArcVectorShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, midarcvectorshapegrid) {
		return
	}

	midarcvectorshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchOriginalPoints3DShape(originalpoints3dshape *OriginalPoints3DShape) {

	// check if instance is already staged
	if IsStaged(stage, originalpoints3dshape) {
		return
	}

	originalpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchParastichyMCurves3DShape(parastichymcurves3dshape *ParastichyMCurves3DShape) {

	// check if instance is already staged
	if IsStaged(stage, parastichymcurves3dshape) {
		return
	}

	parastichymcurves3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchParastichyNCurves3DShape(parastichyncurves3dshape *ParastichyNCurves3DShape) {

	// check if instance is already staged
	if IsStaged(stage, parastichyncurves3dshape) {
		return
	}

	parastichyncurves3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DRibbon(partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) {

	// check if instance is already staged
	if IsStaged(stage, partiallygrowthcurve2dribbon) {
		return
	}

	partiallygrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DRibbonEndShape(partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if IsStaged(stage, partiallygrowthcurve2dribbonendshape) {
		return
	}

	partiallygrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DRibbonStartShape(partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if IsStaged(stage, partiallygrowthcurve2dribbonstartshape) {
		return
	}

	partiallygrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectory(partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) {

	// check if instance is already staged
	if IsStaged(stage, partiallygrowthcurve2dtrajectory) {
		return
	}

	partiallygrowthcurve2dtrajectory.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryP1CurveShape(partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) {

	// check if instance is already staged
	if IsStaged(stage, partiallygrowthcurve2dtrajectoryp1curveshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1curveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryP1P2(partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) {

	// check if instance is already staged
	if IsStaged(stage, partiallygrowthcurve2dtrajectoryp1p2) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) {

	// check if instance is already staged
	if IsStaged(stage, partiallygrowthcurve2dtrajectoryp1p2pairlineshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2pairlineshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryP1PointShape(partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) {

	// check if instance is already staged
	if IsStaged(stage, partiallygrowthcurve2dtrajectoryp1pointshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1pointshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryP2CurveShape(partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) {

	// check if instance is already staged
	if IsStaged(stage, partiallygrowthcurve2dtrajectoryp2curveshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp2curveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryP2PointShape(partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) {

	// check if instance is already staged
	if IsStaged(stage, partiallygrowthcurve2dtrajectoryp2pointshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp2pointshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryShape(partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) {

	// check if instance is already staged
	if IsStaged(stage, partiallygrowthcurve2dtrajectoryshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyRotatedSeatBottomCurveShape(partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) {

	// check if instance is already staged
	if IsStaged(stage, partiallyrotatedseatbottomcurveshape) {
		return
	}

	partiallyrotatedseatbottomcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyRotatedSeatTopCurveShape(partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) {

	// check if instance is already staged
	if IsStaged(stage, partiallyrotatedseattopcurveshape) {
		return
	}

	partiallyrotatedseattopcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPartiallyRotatedTorusShape(partiallyrotatedtorusshape *PartiallyRotatedTorusShape) {

	// check if instance is already staged
	if IsStaged(stage, partiallyrotatedtorusshape) {
		return
	}

	partiallyrotatedtorusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPerpendicularVector(perpendicularvector *PerpendicularVector) {

	// check if instance is already staged
	if IsStaged(stage, perpendicularvector) {
		return
	}

	perpendicularvector.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPerpendicularVectorGrid(perpendicularvectorgrid *PerpendicularVectorGrid) {

	// check if instance is already staged
	if IsStaged(stage, perpendicularvectorgrid) {
		return
	}

	perpendicularvectorgrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) {

	// check if instance is already staged
	if IsStaged(stage, perpendicularvectorgridhalfway) {
		return
	}

	perpendicularvectorgridhalfway.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPerpendicularVectorHalfway(perpendicularvectorhalfway *PerpendicularVectorHalfway) {

	// check if instance is already staged
	if IsStaged(stage, perpendicularvectorhalfway) {
		return
	}

	perpendicularvectorhalfway.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPlant2DDiagram(plant2ddiagram *Plant2DDiagram) {

	// check if instance is already staged
	if IsStaged(stage, plant2ddiagram) {
		return
	}

	plant2ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPlant3DDiagram(plant3ddiagram *Plant3DDiagram) {

	// check if instance is already staged
	if IsStaged(stage, plant3ddiagram) {
		return
	}

	plant3ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plant3ddiagram.StemCylinder3DShape != nil {
		StageBranch(stage, plant3ddiagram.StemCylinder3DShape)
	}
	if plant3ddiagram.ParastichyNCurves3DShape != nil {
		StageBranch(stage, plant3ddiagram.ParastichyNCurves3DShape)
	}
	if plant3ddiagram.ParastichyMCurves3DShape != nil {
		StageBranch(stage, plant3ddiagram.ParastichyMCurves3DShape)
	}
	if plant3ddiagram.CutLine3DShape != nil {
		StageBranch(stage, plant3ddiagram.CutLine3DShape)
	}
	if plant3ddiagram.Circumference3DShape != nil {
		StageBranch(stage, plant3ddiagram.Circumference3DShape)
	}
	if plant3ddiagram.Rendered3DShape != nil {
		StageBranch(stage, plant3ddiagram.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPlantAbstract(plantabstract *PlantAbstract) {

	// check if instance is already staged
	if IsStaged(stage, plantabstract) {
		return
	}

	plantabstract.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plantabstract.VaseAbstract != nil {
		StageBranch(stage, plantabstract.VaseAbstract)
	}
	if plantabstract.StoolAbstract != nil {
		StageBranch(stage, plantabstract.StoolAbstract)
	}
	if plantabstract.ClockAbstract != nil {
		StageBranch(stage, plantabstract.ClockAbstract)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plant2ddiagram := range plantabstract.Plant2DDiagrams {
		StageBranch(stage, _plant2ddiagram)
	}
	for _, _plant3ddiagram := range plantabstract.Plant3DDiagrams {
		StageBranch(stage, _plant3ddiagram)
	}
	for _, _vase2ddiagram := range plantabstract.Vase2DDiagrams {
		StageBranch(stage, _vase2ddiagram)
	}
	for _, _vase3ddiagram := range plantabstract.Vase3DDiagrams {
		StageBranch(stage, _vase3ddiagram)
	}
	for _, _stool2ddiagram := range plantabstract.Stool2DDiagrams {
		StageBranch(stage, _stool2ddiagram)
	}
	for _, _stool3ddiagram := range plantabstract.Stool3DDiagrams {
		StageBranch(stage, _stool3ddiagram)
	}
	for _, _clock2ddiagram := range plantabstract.Clock2DDiagrams {
		StageBranch(stage, _clock2ddiagram)
	}
	for _, _clock3ddiagram := range plantabstract.Clock3DDiagrams {
		StageBranch(stage, _clock3ddiagram)
	}

}

func (stage *Stage) StageBranchPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape) {

	// check if instance is already staged
	if IsStaged(stage, plantcircumferenceshape) {
		return
	}

	plantcircumferenceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPointsAndLines3DShape(pointsandlines3dshape *PointsAndLines3DShape) {

	// check if instance is already staged
	if IsStaged(stage, pointsandlines3dshape) {
		return
	}

	pointsandlines3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPxShape(pxshape *PxShape) {

	// check if instance is already staged
	if IsStaged(stage, pxshape) {
		return
	}

	pxshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRendered3DShape(rendered3dshape *Rendered3DShape) {

	// check if instance is already staged
	if IsStaged(stage, rendered3dshape) {
		return
	}

	rendered3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRhombusShape(rhombusshape *RhombusShape) {

	// check if instance is already staged
	if IsStaged(stage, rhombusshape) {
		return
	}

	rhombusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRhombusStuff(rhombusstuff *RhombusStuff) {

	// check if instance is already staged
	if IsStaged(stage, rhombusstuff) {
		return
	}

	rhombusstuff.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRotatedRhombusGridShape(rotatedrhombusgridshape *RotatedRhombusGridShape) {

	// check if instance is already staged
	if IsStaged(stage, rotatedrhombusgridshape) {
		return
	}

	rotatedrhombusgridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRotatedRhombusShape(rotatedrhombusshape *RotatedRhombusShape) {

	// check if instance is already staged
	if IsStaged(stage, rotatedrhombusshape) {
		return
	}

	rotatedrhombusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRotatedSampledPoints3DShape(rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) {

	// check if instance is already staged
	if IsStaged(stage, rotatedsampledpoints3dshape) {
		return
	}

	rotatedsampledpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRotatedSeatAndLegs3DShape(rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) {

	// check if instance is already staged
	if IsStaged(stage, rotatedseatandlegs3dshape) {
		return
	}

	rotatedseatandlegs3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSampledPoints3DShape(sampledpoints3dshape *SampledPoints3DShape) {

	// check if instance is already staged
	if IsStaged(stage, sampledpoints3dshape) {
		return
	}

	sampledpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSeat3DShape(seat3dshape *Seat3DShape) {

	// check if instance is already staged
	if IsStaged(stage, seat3dshape) {
		return
	}

	seat3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSeatAndLegs3DShape(seatandlegs3dshape *SeatAndLegs3DShape) {

	// check if instance is already staged
	if IsStaged(stage, seatandlegs3dshape) {
		return
	}

	seatandlegs3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSeatBottomCurveShape(seatbottomcurveshape *SeatBottomCurveShape) {

	// check if instance is already staged
	if IsStaged(stage, seatbottomcurveshape) {
		return
	}

	seatbottomcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSeatTopCurveShape(seattopcurveshape *SeatTopCurveShape) {

	// check if instance is already staged
	if IsStaged(stage, seattopcurveshape) {
		return
	}

	seattopcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedBottomTopStartArcShape(shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) {

	// check if instance is already staged
	if IsStaged(stage, shiftedbottomtopstartarcshape) {
		return
	}

	shiftedbottomtopstartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedBottomTopStartArcShapeGrid(shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, shiftedbottomtopstartarcshapegrid) {
		return
	}

	shiftedbottomtopstartarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedLeftGrowthCurve2DRibbon(shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) {

	// check if instance is already staged
	if IsStaged(stage, shiftedleftgrowthcurve2dribbon) {
		return
	}

	shiftedleftgrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedLeftGrowthCurve2DRibbonEndShape(shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if IsStaged(stage, shiftedleftgrowthcurve2dribbonendshape) {
		return
	}

	shiftedleftgrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedLeftGrowthCurve2DRibbonStartShape(shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if IsStaged(stage, shiftedleftgrowthcurve2dribbonstartshape) {
		return
	}

	shiftedleftgrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedLeftPartiallyGrowthCurve2DRibbon(shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) {

	// check if instance is already staged
	if IsStaged(stage, shiftedleftpartiallygrowthcurve2dribbon) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if IsStaged(stage, shiftedleftpartiallygrowthcurve2dribbonendshape) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if IsStaged(stage, shiftedleftpartiallygrowthcurve2dribbonstartshape) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedLeftStackGrowthCurveEndArcShape(shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) {

	// check if instance is already staged
	if IsStaged(stage, shiftedleftstackgrowthcurveendarcshape) {
		return
	}

	shiftedleftstackgrowthcurveendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedLeftStackGrowthCurveStartArcShape(shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) {

	// check if instance is already staged
	if IsStaged(stage, shiftedleftstackgrowthcurvestartarcshape) {
		return
	}

	shiftedleftstackgrowthcurvestartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedLeftStackNormalVector(shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) {

	// check if instance is already staged
	if IsStaged(stage, shiftedleftstacknormalvector) {
		return
	}

	shiftedleftstacknormalvector.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedLeftStackOfGrowthCurve(shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) {

	// check if instance is already staged
	if IsStaged(stage, shiftedleftstackofgrowthcurve) {
		return
	}

	shiftedleftstackofgrowthcurve.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) {

	// check if instance is already staged
	if IsStaged(stage, shiftedleftstackofnormalvector) {
		return
	}

	shiftedleftstackofnormalvector.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedRightGrowthCurve2DRibbon(shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) {

	// check if instance is already staged
	if IsStaged(stage, shiftedrightgrowthcurve2dribbon) {
		return
	}

	shiftedrightgrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedRightGrowthCurve2DRibbonEndShape(shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if IsStaged(stage, shiftedrightgrowthcurve2dribbonendshape) {
		return
	}

	shiftedrightgrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShiftedRightGrowthCurve2DRibbonStartShape(shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if IsStaged(stage, shiftedrightgrowthcurve2dribbonstartshape) {
		return
	}

	shiftedrightgrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackGrowthCurve2DEndHalfwayArcShape(stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) {

	// check if instance is already staged
	if IsStaged(stage, stackgrowthcurve2dendhalfwayarcshape) {
		return
	}

	stackgrowthcurve2dendhalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackGrowthCurve2DRibbonEndShape(stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if IsStaged(stage, stackgrowthcurve2dribbonendshape) {
		return
	}

	stackgrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackGrowthCurve2DRibbonStartShape(stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if IsStaged(stage, stackgrowthcurve2dribbonstartshape) {
		return
	}

	stackgrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackGrowthCurve2DStartHalfwayArcShape(stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) {

	// check if instance is already staged
	if IsStaged(stage, stackgrowthcurve2dstarthalfwayarcshape) {
		return
	}

	stackgrowthcurve2dstarthalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackOfGrowthCurve2D(stackofgrowthcurve2d *StackOfGrowthCurve2D) {

	// check if instance is already staged
	if IsStaged(stage, stackofgrowthcurve2d) {
		return
	}

	stackofgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackOfGrowthCurve2DByGrowthVector(stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) {

	// check if instance is already staged
	if IsStaged(stage, stackofgrowthcurve2dbygrowthvector) {
		return
	}

	stackofgrowthcurve2dbygrowthvector.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackOfGrowthCurve2DRibbon(stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) {

	// check if instance is already staged
	if IsStaged(stage, stackofgrowthcurve2dribbon) {
		return
	}

	stackofgrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackOfPartiallyRotatedTorusShape(stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) {

	// check if instance is already staged
	if IsStaged(stage, stackofpartiallyrotatedtorusshape) {
		return
	}

	stackofpartiallyrotatedtorusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackOfRotatedGrowthCurve2D(stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) {

	// check if instance is already staged
	if IsStaged(stage, stackofrotatedgrowthcurve2d) {
		return
	}

	stackofrotatedgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackOfRotatedGrowthCurve2DRibbon(stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) {

	// check if instance is already staged
	if IsStaged(stage, stackofrotatedgrowthcurve2dribbon) {
		return
	}

	stackofrotatedgrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackRotatedGrowthCurve2DEndArcShape(stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) {

	// check if instance is already staged
	if IsStaged(stage, stackrotatedgrowthcurve2dendarcshape) {
		return
	}

	stackrotatedgrowthcurve2dendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackRotatedGrowthCurve2DRibbonEndShape(stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if IsStaged(stage, stackrotatedgrowthcurve2dribbonendshape) {
		return
	}

	stackrotatedgrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackRotatedGrowthCurve2DRibbonStartShape(stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if IsStaged(stage, stackrotatedgrowthcurve2dribbonstartshape) {
		return
	}

	stackrotatedgrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackRotatedGrowthCurve2DStartArcShape(stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) {

	// check if instance is already staged
	if IsStaged(stage, stackrotatedgrowthcurve2dstartarcshape) {
		return
	}

	stackrotatedgrowthcurve2dstartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStartArcShape(startarcshape *StartArcShape) {

	// check if instance is already staged
	if IsStaged(stage, startarcshape) {
		return
	}

	startarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStartArcShapeGrid(startarcshapegrid *StartArcShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, startarcshapegrid) {
		return
	}

	startarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStartHalfwayArcShape(starthalfwayarcshape *StartHalfwayArcShape) {

	// check if instance is already staged
	if IsStaged(stage, starthalfwayarcshape) {
		return
	}

	starthalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStartHalfwayArcShapeGrid(starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, starthalfwayarcshapegrid) {
		return
	}

	starthalfwayarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStemCylinder3DShape(stemcylinder3dshape *StemCylinder3DShape) {

	// check if instance is already staged
	if IsStaged(stage, stemcylinder3dshape) {
		return
	}

	stemcylinder3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStool2DDiagram(stool2ddiagram *Stool2DDiagram) {

	// check if instance is already staged
	if IsStaged(stage, stool2ddiagram) {
		return
	}

	stool2ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStool3DDiagram(stool3ddiagram *Stool3DDiagram) {

	// check if instance is already staged
	if IsStaged(stage, stool3ddiagram) {
		return
	}

	stool3ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stool3ddiagram.SampledPoints3DShape != nil {
		StageBranch(stage, stool3ddiagram.SampledPoints3DShape)
	}
	if stool3ddiagram.Rendered3DShape != nil {
		StageBranch(stage, stool3ddiagram.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStoolAbstract(stoolabstract *StoolAbstract) {

	// check if instance is already staged
	if IsStaged(stage, stoolabstract) {
		return
	}

	stoolabstract.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTiledFloor3DShape(tiledfloor3dshape *TiledFloor3DShape) {

	// check if instance is already staged
	if IsStaged(stage, tiledfloor3dshape) {
		return
	}

	tiledfloor3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopEndArcShape(topendarcshape *TopEndArcShape) {

	// check if instance is already staged
	if IsStaged(stage, topendarcshape) {
		return
	}

	topendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopEndArcShapeGrid(topendarcshapegrid *TopEndArcShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, topendarcshapegrid) {
		return
	}

	topendarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopEndHalfwayArcShape(topendhalfwayarcshape *TopEndHalfwayArcShape) {

	// check if instance is already staged
	if IsStaged(stage, topendhalfwayarcshape) {
		return
	}

	topendhalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopEndHalfwayArcShapeGrid(topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, topendhalfwayarcshapegrid) {
		return
	}

	topendhalfwayarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D) {

	// check if instance is already staged
	if IsStaged(stage, topgrowthcurve2d) {
		return
	}

	topgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopMidArcVectorShape(topmidarcvectorshape *TopMidArcVectorShape) {

	// check if instance is already staged
	if IsStaged(stage, topmidarcvectorshape) {
		return
	}

	topmidarcvectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopMidArcVectorShapeGrid(topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, topmidarcvectorshapegrid) {
		return
	}

	topmidarcvectorshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStackGrowthCurve2DEndHalfwayArcShape(topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) {

	// check if instance is already staged
	if IsStaged(stage, topstackgrowthcurve2dendhalfwayarcshape) {
		return
	}

	topstackgrowthcurve2dendhalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStackGrowthCurve2DStartHalfwayArcShape(topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) {

	// check if instance is already staged
	if IsStaged(stage, topstackgrowthcurve2dstarthalfwayarcshape) {
		return
	}

	topstackgrowthcurve2dstarthalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStackOfGrowthCurve2D(topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) {

	// check if instance is already staged
	if IsStaged(stage, topstackofgrowthcurve2d) {
		return
	}

	topstackofgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStackOfRotatedGrowthCurve2D(topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) {

	// check if instance is already staged
	if IsStaged(stage, topstackofrotatedgrowthcurve2d) {
		return
	}

	topstackofrotatedgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStackOfRotatedGrowthCurve2DEndArcShape(topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) {

	// check if instance is already staged
	if IsStaged(stage, topstackofrotatedgrowthcurve2dendarcshape) {
		return
	}

	topstackofrotatedgrowthcurve2dendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStackOfRotatedGrowthCurve2DStartArcShape(topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) {

	// check if instance is already staged
	if IsStaged(stage, topstackofrotatedgrowthcurve2dstartarcshape) {
		return
	}

	topstackofrotatedgrowthcurve2dstartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStartArcShape(topstartarcshape *TopStartArcShape) {

	// check if instance is already staged
	if IsStaged(stage, topstartarcshape) {
		return
	}

	topstartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStartArcShapeGrid(topstartarcshapegrid *TopStartArcShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, topstartarcshapegrid) {
		return
	}

	topstartarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStartHalfwayArcShape(topstarthalfwayarcshape *TopStartHalfwayArcShape) {

	// check if instance is already staged
	if IsStaged(stage, topstarthalfwayarcshape) {
		return
	}

	topstarthalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStartHalfwayArcShapeGrid(topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, topstarthalfwayarcshapegrid) {
		return
	}

	topstarthalfwayarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTorus3DShape(torus3dshape *Torus3DShape) {

	// check if instance is already staged
	if IsStaged(stage, torus3dshape) {
		return
	}

	torus3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTorusEdge3DShape(torusedge3dshape *TorusEdge3DShape) {

	// check if instance is already staged
	if IsStaged(stage, torusedge3dshape) {
		return
	}

	torusedge3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTorusStackShape(torusstackshape *TorusStackShape) {

	// check if instance is already staged
	if IsStaged(stage, torusstackshape) {
		return
	}

	torusstackshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchVase2DDiagram(vase2ddiagram *Vase2DDiagram) {

	// check if instance is already staged
	if IsStaged(stage, vase2ddiagram) {
		return
	}

	vase2ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchVase3DDiagram(vase3ddiagram *Vase3DDiagram) {

	// check if instance is already staged
	if IsStaged(stage, vase3ddiagram) {
		return
	}

	vase3ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if vase3ddiagram.Rendered3DShape != nil {
		StageBranch(stage, vase3ddiagram.Rendered3DShape)
	}
	if vase3ddiagram.SampledPoints3DShape != nil {
		StageBranch(stage, vase3ddiagram.SampledPoints3DShape)
	}
	if vase3ddiagram.OriginalPoints3DShape != nil {
		StageBranch(stage, vase3ddiagram.OriginalPoints3DShape)
	}
	if vase3ddiagram.Angle0Shape != nil {
		StageBranch(stage, vase3ddiagram.Angle0Shape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchVaseAbstract(vaseabstract *VaseAbstract) {

	// check if instance is already staged
	if IsStaged(stage, vaseabstract) {
		return
	}

	vaseabstract.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchVerticalTorusStackShape(verticaltorusstackshape *VerticalTorusStackShape) {

	// check if instance is already staged
	if IsStaged(stage, verticaltorusstackshape) {
		return
	}

	verticaltorusstackshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchVolumeKey3DShape(volumekey3dshape *VolumeKey3DShape) {

	// check if instance is already staged
	if IsStaged(stage, volumekey3dshape) {
		return
	}

	volumekey3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Angle0Shape:
		toT := CopyBranchAngle0Shape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ArcNormalVectorShape:
		toT := CopyBranchArcNormalVectorShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ArcNormalVectorShapeGrid:
		toT := CopyBranchArcNormalVectorShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AxesShape:
		toT := CopyBranchAxesShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BaseVectorShape:
		toT := CopyBranchBaseVectorShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BaseVectorShapeGrid:
		toT := CopyBranchBaseVectorShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ChosenP1P2PairShape:
		toT := CopyBranchChosenP1P2PairShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CircleGridShape:
		toT := CopyBranchCircleGridShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Circumference3DShape:
		toT := CopyBranchCircumference3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Clock2DDiagram:
		toT := CopyBranchClock2DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Clock3DDiagram:
		toT := CopyBranchClock3DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ClockAbstract:
		toT := CopyBranchClockAbstract(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ClockTopCurveShape:
		toT := CopyBranchClockTopCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CutLine3DShape:
		toT := CopyBranchCutLine3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EndArcShape:
		toT := CopyBranchEndArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EndArcShapeGrid:
		toT := CopyBranchEndArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EndHalfwayArcShape:
		toT := CopyBranchEndHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EndHalfwayArcShapeGrid:
		toT := CopyBranchEndHalfwayArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ExplanationTextShape:
		toT := CopyBranchExplanationTextShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Eye3DShape:
		toT := CopyBranchEye3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EyeCornersSampledPoints3DShape:
		toT := CopyBranchEyeCornersSampledPoints3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EyeSampledPoints3DShape:
		toT := CopyBranchEyeSampledPoints3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EyeSeatBottomCurveShape:
		toT := CopyBranchEyeSeatBottomCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EyeStoolBottomCurveShape:
		toT := CopyBranchEyeStoolBottomCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EyeVolume3DShape:
		toT := CopyBranchEyeVolume3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GridPathShape:
		toT := CopyBranchGridPathShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurve2D:
		toT := CopyBranchGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurve2DRibbon:
		toT := CopyBranchGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurve2DRibbonEndShape:
		toT := CopyBranchGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurve2DRibbonStartShape:
		toT := CopyBranchGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurveRhombusGridShape:
		toT := CopyBranchGrowthCurveRhombusGridShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurveRhombusShape:
		toT := CopyBranchGrowthCurveRhombusShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthVectorShape:
		toT := CopyBranchGrowthVectorShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *InitialRhombusGridShape:
		toT := CopyBranchInitialRhombusGridShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *InitialRhombusShape:
		toT := CopyBranchInitialRhombusShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Key3DShape:
		toT := CopyBranchKey3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *KeyHole3DShape:
		toT := CopyBranchKeyHole3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *KeyHoleShape:
		toT := CopyBranchKeyHoleShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MidArcVectorShape:
		toT := CopyBranchMidArcVectorShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MidArcVectorShapeGrid:
		toT := CopyBranchMidArcVectorShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OriginalPoints3DShape:
		toT := CopyBranchOriginalPoints3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParastichyMCurves3DShape:
		toT := CopyBranchParastichyMCurves3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParastichyNCurves3DShape:
		toT := CopyBranchParastichyNCurves3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DRibbon:
		toT := CopyBranchPartiallyGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DRibbonEndShape:
		toT := CopyBranchPartiallyGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DRibbonStartShape:
		toT := CopyBranchPartiallyGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectory:
		toT := CopyBranchPartiallyGrowthCurve2DTrajectory(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		toT := CopyBranchPartiallyGrowthCurve2DTrajectoryP1CurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		toT := CopyBranchPartiallyGrowthCurve2DTrajectoryP1P2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		toT := CopyBranchPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		toT := CopyBranchPartiallyGrowthCurve2DTrajectoryP1PointShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		toT := CopyBranchPartiallyGrowthCurve2DTrajectoryP2CurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		toT := CopyBranchPartiallyGrowthCurve2DTrajectoryP2PointShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryShape:
		toT := CopyBranchPartiallyGrowthCurve2DTrajectoryShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyRotatedSeatBottomCurveShape:
		toT := CopyBranchPartiallyRotatedSeatBottomCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyRotatedSeatTopCurveShape:
		toT := CopyBranchPartiallyRotatedSeatTopCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyRotatedTorusShape:
		toT := CopyBranchPartiallyRotatedTorusShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PerpendicularVector:
		toT := CopyBranchPerpendicularVector(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PerpendicularVectorGrid:
		toT := CopyBranchPerpendicularVectorGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PerpendicularVectorGridHalfway:
		toT := CopyBranchPerpendicularVectorGridHalfway(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PerpendicularVectorHalfway:
		toT := CopyBranchPerpendicularVectorHalfway(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Plant2DDiagram:
		toT := CopyBranchPlant2DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Plant3DDiagram:
		toT := CopyBranchPlant3DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PlantAbstract:
		toT := CopyBranchPlantAbstract(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PlantCircumferenceShape:
		toT := CopyBranchPlantCircumferenceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PointsAndLines3DShape:
		toT := CopyBranchPointsAndLines3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PxShape:
		toT := CopyBranchPxShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Rendered3DShape:
		toT := CopyBranchRendered3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RhombusShape:
		toT := CopyBranchRhombusShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RhombusStuff:
		toT := CopyBranchRhombusStuff(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RotatedRhombusGridShape:
		toT := CopyBranchRotatedRhombusGridShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RotatedRhombusShape:
		toT := CopyBranchRotatedRhombusShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RotatedSampledPoints3DShape:
		toT := CopyBranchRotatedSampledPoints3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RotatedSeatAndLegs3DShape:
		toT := CopyBranchRotatedSeatAndLegs3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SampledPoints3DShape:
		toT := CopyBranchSampledPoints3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Seat3DShape:
		toT := CopyBranchSeat3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SeatAndLegs3DShape:
		toT := CopyBranchSeatAndLegs3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SeatBottomCurveShape:
		toT := CopyBranchSeatBottomCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SeatTopCurveShape:
		toT := CopyBranchSeatTopCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedBottomTopStartArcShape:
		toT := CopyBranchShiftedBottomTopStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedBottomTopStartArcShapeGrid:
		toT := CopyBranchShiftedBottomTopStartArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftGrowthCurve2DRibbon:
		toT := CopyBranchShiftedLeftGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftGrowthCurve2DRibbonEndShape:
		toT := CopyBranchShiftedLeftGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftGrowthCurve2DRibbonStartShape:
		toT := CopyBranchShiftedLeftGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbon:
		toT := CopyBranchShiftedLeftPartiallyGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape:
		toT := CopyBranchShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape:
		toT := CopyBranchShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftStackGrowthCurveEndArcShape:
		toT := CopyBranchShiftedLeftStackGrowthCurveEndArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftStackGrowthCurveStartArcShape:
		toT := CopyBranchShiftedLeftStackGrowthCurveStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftStackNormalVector:
		toT := CopyBranchShiftedLeftStackNormalVector(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftStackOfGrowthCurve:
		toT := CopyBranchShiftedLeftStackOfGrowthCurve(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftStackOfNormalVector:
		toT := CopyBranchShiftedLeftStackOfNormalVector(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedRightGrowthCurve2DRibbon:
		toT := CopyBranchShiftedRightGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		toT := CopyBranchShiftedRightGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		toT := CopyBranchShiftedRightGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackGrowthCurve2DEndHalfwayArcShape:
		toT := CopyBranchStackGrowthCurve2DEndHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackGrowthCurve2DRibbonEndShape:
		toT := CopyBranchStackGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackGrowthCurve2DRibbonStartShape:
		toT := CopyBranchStackGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackGrowthCurve2DStartHalfwayArcShape:
		toT := CopyBranchStackGrowthCurve2DStartHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfGrowthCurve2D:
		toT := CopyBranchStackOfGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfGrowthCurve2DByGrowthVector:
		toT := CopyBranchStackOfGrowthCurve2DByGrowthVector(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfGrowthCurve2DRibbon:
		toT := CopyBranchStackOfGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfPartiallyRotatedTorusShape:
		toT := CopyBranchStackOfPartiallyRotatedTorusShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfRotatedGrowthCurve2D:
		toT := CopyBranchStackOfRotatedGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfRotatedGrowthCurve2DRibbon:
		toT := CopyBranchStackOfRotatedGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackRotatedGrowthCurve2DEndArcShape:
		toT := CopyBranchStackRotatedGrowthCurve2DEndArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackRotatedGrowthCurve2DRibbonEndShape:
		toT := CopyBranchStackRotatedGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackRotatedGrowthCurve2DRibbonStartShape:
		toT := CopyBranchStackRotatedGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackRotatedGrowthCurve2DStartArcShape:
		toT := CopyBranchStackRotatedGrowthCurve2DStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartArcShape:
		toT := CopyBranchStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartArcShapeGrid:
		toT := CopyBranchStartArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartHalfwayArcShape:
		toT := CopyBranchStartHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartHalfwayArcShapeGrid:
		toT := CopyBranchStartHalfwayArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StemCylinder3DShape:
		toT := CopyBranchStemCylinder3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Stool2DDiagram:
		toT := CopyBranchStool2DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Stool3DDiagram:
		toT := CopyBranchStool3DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StoolAbstract:
		toT := CopyBranchStoolAbstract(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TiledFloor3DShape:
		toT := CopyBranchTiledFloor3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopEndArcShape:
		toT := CopyBranchTopEndArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopEndArcShapeGrid:
		toT := CopyBranchTopEndArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopEndHalfwayArcShape:
		toT := CopyBranchTopEndHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopEndHalfwayArcShapeGrid:
		toT := CopyBranchTopEndHalfwayArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopGrowthCurve2D:
		toT := CopyBranchTopGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopMidArcVectorShape:
		toT := CopyBranchTopMidArcVectorShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopMidArcVectorShapeGrid:
		toT := CopyBranchTopMidArcVectorShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		toT := CopyBranchTopStackGrowthCurve2DEndHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		toT := CopyBranchTopStackGrowthCurve2DStartHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackOfGrowthCurve2D:
		toT := CopyBranchTopStackOfGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackOfRotatedGrowthCurve2D:
		toT := CopyBranchTopStackOfRotatedGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		toT := CopyBranchTopStackOfRotatedGrowthCurve2DEndArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		toT := CopyBranchTopStackOfRotatedGrowthCurve2DStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStartArcShape:
		toT := CopyBranchTopStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStartArcShapeGrid:
		toT := CopyBranchTopStartArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStartHalfwayArcShape:
		toT := CopyBranchTopStartHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStartHalfwayArcShapeGrid:
		toT := CopyBranchTopStartHalfwayArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Torus3DShape:
		toT := CopyBranchTorus3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TorusEdge3DShape:
		toT := CopyBranchTorusEdge3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TorusStackShape:
		toT := CopyBranchTorusStackShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Vase2DDiagram:
		toT := CopyBranchVase2DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Vase3DDiagram:
		toT := CopyBranchVase3DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VaseAbstract:
		toT := CopyBranchVaseAbstract(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VerticalTorusStackShape:
		toT := CopyBranchVerticalTorusStackShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VolumeKey3DShape:
		toT := CopyBranchVolumeKey3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAngle0Shape(mapOrigCopy map[any]any, angle0shapeFrom *Angle0Shape) (angle0shapeTo *Angle0Shape) {

	// angle0shapeFrom has already been copied
	if _angle0shapeTo, ok := mapOrigCopy[angle0shapeFrom]; ok {
		angle0shapeTo = _angle0shapeTo.(*Angle0Shape)
		return
	}

	angle0shapeTo = new(Angle0Shape)
	mapOrigCopy[angle0shapeFrom] = angle0shapeTo
	angle0shapeFrom.CopyBasicFields(angle0shapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchArcNormalVectorShape(mapOrigCopy map[any]any, arcnormalvectorshapeFrom *ArcNormalVectorShape) (arcnormalvectorshapeTo *ArcNormalVectorShape) {

	// arcnormalvectorshapeFrom has already been copied
	if _arcnormalvectorshapeTo, ok := mapOrigCopy[arcnormalvectorshapeFrom]; ok {
		arcnormalvectorshapeTo = _arcnormalvectorshapeTo.(*ArcNormalVectorShape)
		return
	}

	arcnormalvectorshapeTo = new(ArcNormalVectorShape)
	mapOrigCopy[arcnormalvectorshapeFrom] = arcnormalvectorshapeTo
	arcnormalvectorshapeFrom.CopyBasicFields(arcnormalvectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchArcNormalVectorShapeGrid(mapOrigCopy map[any]any, arcnormalvectorshapegridFrom *ArcNormalVectorShapeGrid) (arcnormalvectorshapegridTo *ArcNormalVectorShapeGrid) {

	// arcnormalvectorshapegridFrom has already been copied
	if _arcnormalvectorshapegridTo, ok := mapOrigCopy[arcnormalvectorshapegridFrom]; ok {
		arcnormalvectorshapegridTo = _arcnormalvectorshapegridTo.(*ArcNormalVectorShapeGrid)
		return
	}

	arcnormalvectorshapegridTo = new(ArcNormalVectorShapeGrid)
	mapOrigCopy[arcnormalvectorshapegridFrom] = arcnormalvectorshapegridTo
	arcnormalvectorshapegridFrom.CopyBasicFields(arcnormalvectorshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAxesShape(mapOrigCopy map[any]any, axesshapeFrom *AxesShape) (axesshapeTo *AxesShape) {

	// axesshapeFrom has already been copied
	if _axesshapeTo, ok := mapOrigCopy[axesshapeFrom]; ok {
		axesshapeTo = _axesshapeTo.(*AxesShape)
		return
	}

	axesshapeTo = new(AxesShape)
	mapOrigCopy[axesshapeFrom] = axesshapeTo
	axesshapeFrom.CopyBasicFields(axesshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBaseVectorShape(mapOrigCopy map[any]any, basevectorshapeFrom *BaseVectorShape) (basevectorshapeTo *BaseVectorShape) {

	// basevectorshapeFrom has already been copied
	if _basevectorshapeTo, ok := mapOrigCopy[basevectorshapeFrom]; ok {
		basevectorshapeTo = _basevectorshapeTo.(*BaseVectorShape)
		return
	}

	basevectorshapeTo = new(BaseVectorShape)
	mapOrigCopy[basevectorshapeFrom] = basevectorshapeTo
	basevectorshapeFrom.CopyBasicFields(basevectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBaseVectorShapeGrid(mapOrigCopy map[any]any, basevectorshapegridFrom *BaseVectorShapeGrid) (basevectorshapegridTo *BaseVectorShapeGrid) {

	// basevectorshapegridFrom has already been copied
	if _basevectorshapegridTo, ok := mapOrigCopy[basevectorshapegridFrom]; ok {
		basevectorshapegridTo = _basevectorshapegridTo.(*BaseVectorShapeGrid)
		return
	}

	basevectorshapegridTo = new(BaseVectorShapeGrid)
	mapOrigCopy[basevectorshapegridFrom] = basevectorshapegridTo
	basevectorshapegridFrom.CopyBasicFields(basevectorshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchChosenP1P2PairShape(mapOrigCopy map[any]any, chosenp1p2pairshapeFrom *ChosenP1P2PairShape) (chosenp1p2pairshapeTo *ChosenP1P2PairShape) {

	// chosenp1p2pairshapeFrom has already been copied
	if _chosenp1p2pairshapeTo, ok := mapOrigCopy[chosenp1p2pairshapeFrom]; ok {
		chosenp1p2pairshapeTo = _chosenp1p2pairshapeTo.(*ChosenP1P2PairShape)
		return
	}

	chosenp1p2pairshapeTo = new(ChosenP1P2PairShape)
	mapOrigCopy[chosenp1p2pairshapeFrom] = chosenp1p2pairshapeTo
	chosenp1p2pairshapeFrom.CopyBasicFields(chosenp1p2pairshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCircleGridShape(mapOrigCopy map[any]any, circlegridshapeFrom *CircleGridShape) (circlegridshapeTo *CircleGridShape) {

	// circlegridshapeFrom has already been copied
	if _circlegridshapeTo, ok := mapOrigCopy[circlegridshapeFrom]; ok {
		circlegridshapeTo = _circlegridshapeTo.(*CircleGridShape)
		return
	}

	circlegridshapeTo = new(CircleGridShape)
	mapOrigCopy[circlegridshapeFrom] = circlegridshapeTo
	circlegridshapeFrom.CopyBasicFields(circlegridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCircumference3DShape(mapOrigCopy map[any]any, circumference3dshapeFrom *Circumference3DShape) (circumference3dshapeTo *Circumference3DShape) {

	// circumference3dshapeFrom has already been copied
	if _circumference3dshapeTo, ok := mapOrigCopy[circumference3dshapeFrom]; ok {
		circumference3dshapeTo = _circumference3dshapeTo.(*Circumference3DShape)
		return
	}

	circumference3dshapeTo = new(Circumference3DShape)
	mapOrigCopy[circumference3dshapeFrom] = circumference3dshapeTo
	circumference3dshapeFrom.CopyBasicFields(circumference3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchClock2DDiagram(mapOrigCopy map[any]any, clock2ddiagramFrom *Clock2DDiagram) (clock2ddiagramTo *Clock2DDiagram) {

	// clock2ddiagramFrom has already been copied
	if _clock2ddiagramTo, ok := mapOrigCopy[clock2ddiagramFrom]; ok {
		clock2ddiagramTo = _clock2ddiagramTo.(*Clock2DDiagram)
		return
	}

	clock2ddiagramTo = new(Clock2DDiagram)
	mapOrigCopy[clock2ddiagramFrom] = clock2ddiagramTo
	clock2ddiagramFrom.CopyBasicFields(clock2ddiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchClock3DDiagram(mapOrigCopy map[any]any, clock3ddiagramFrom *Clock3DDiagram) (clock3ddiagramTo *Clock3DDiagram) {

	// clock3ddiagramFrom has already been copied
	if _clock3ddiagramTo, ok := mapOrigCopy[clock3ddiagramFrom]; ok {
		clock3ddiagramTo = _clock3ddiagramTo.(*Clock3DDiagram)
		return
	}

	clock3ddiagramTo = new(Clock3DDiagram)
	mapOrigCopy[clock3ddiagramFrom] = clock3ddiagramTo
	clock3ddiagramFrom.CopyBasicFields(clock3ddiagramTo)

	//insertion point for the staging of instances referenced by pointers
	if clock3ddiagramFrom.SampledPoints3DShape != nil {
		clock3ddiagramTo.SampledPoints3DShape = CopyBranchSampledPoints3DShape(mapOrigCopy, clock3ddiagramFrom.SampledPoints3DShape)
	}
	if clock3ddiagramFrom.Rendered3DShape != nil {
		clock3ddiagramTo.Rendered3DShape = CopyBranchRendered3DShape(mapOrigCopy, clock3ddiagramFrom.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchClockAbstract(mapOrigCopy map[any]any, clockabstractFrom *ClockAbstract) (clockabstractTo *ClockAbstract) {

	// clockabstractFrom has already been copied
	if _clockabstractTo, ok := mapOrigCopy[clockabstractFrom]; ok {
		clockabstractTo = _clockabstractTo.(*ClockAbstract)
		return
	}

	clockabstractTo = new(ClockAbstract)
	mapOrigCopy[clockabstractFrom] = clockabstractTo
	clockabstractFrom.CopyBasicFields(clockabstractTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchClockTopCurveShape(mapOrigCopy map[any]any, clocktopcurveshapeFrom *ClockTopCurveShape) (clocktopcurveshapeTo *ClockTopCurveShape) {

	// clocktopcurveshapeFrom has already been copied
	if _clocktopcurveshapeTo, ok := mapOrigCopy[clocktopcurveshapeFrom]; ok {
		clocktopcurveshapeTo = _clocktopcurveshapeTo.(*ClockTopCurveShape)
		return
	}

	clocktopcurveshapeTo = new(ClockTopCurveShape)
	mapOrigCopy[clocktopcurveshapeFrom] = clocktopcurveshapeTo
	clocktopcurveshapeFrom.CopyBasicFields(clocktopcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCutLine3DShape(mapOrigCopy map[any]any, cutline3dshapeFrom *CutLine3DShape) (cutline3dshapeTo *CutLine3DShape) {

	// cutline3dshapeFrom has already been copied
	if _cutline3dshapeTo, ok := mapOrigCopy[cutline3dshapeFrom]; ok {
		cutline3dshapeTo = _cutline3dshapeTo.(*CutLine3DShape)
		return
	}

	cutline3dshapeTo = new(CutLine3DShape)
	mapOrigCopy[cutline3dshapeFrom] = cutline3dshapeTo
	cutline3dshapeFrom.CopyBasicFields(cutline3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEndArcShape(mapOrigCopy map[any]any, endarcshapeFrom *EndArcShape) (endarcshapeTo *EndArcShape) {

	// endarcshapeFrom has already been copied
	if _endarcshapeTo, ok := mapOrigCopy[endarcshapeFrom]; ok {
		endarcshapeTo = _endarcshapeTo.(*EndArcShape)
		return
	}

	endarcshapeTo = new(EndArcShape)
	mapOrigCopy[endarcshapeFrom] = endarcshapeTo
	endarcshapeFrom.CopyBasicFields(endarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEndArcShapeGrid(mapOrigCopy map[any]any, endarcshapegridFrom *EndArcShapeGrid) (endarcshapegridTo *EndArcShapeGrid) {

	// endarcshapegridFrom has already been copied
	if _endarcshapegridTo, ok := mapOrigCopy[endarcshapegridFrom]; ok {
		endarcshapegridTo = _endarcshapegridTo.(*EndArcShapeGrid)
		return
	}

	endarcshapegridTo = new(EndArcShapeGrid)
	mapOrigCopy[endarcshapegridFrom] = endarcshapegridTo
	endarcshapegridFrom.CopyBasicFields(endarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEndHalfwayArcShape(mapOrigCopy map[any]any, endhalfwayarcshapeFrom *EndHalfwayArcShape) (endhalfwayarcshapeTo *EndHalfwayArcShape) {

	// endhalfwayarcshapeFrom has already been copied
	if _endhalfwayarcshapeTo, ok := mapOrigCopy[endhalfwayarcshapeFrom]; ok {
		endhalfwayarcshapeTo = _endhalfwayarcshapeTo.(*EndHalfwayArcShape)
		return
	}

	endhalfwayarcshapeTo = new(EndHalfwayArcShape)
	mapOrigCopy[endhalfwayarcshapeFrom] = endhalfwayarcshapeTo
	endhalfwayarcshapeFrom.CopyBasicFields(endhalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEndHalfwayArcShapeGrid(mapOrigCopy map[any]any, endhalfwayarcshapegridFrom *EndHalfwayArcShapeGrid) (endhalfwayarcshapegridTo *EndHalfwayArcShapeGrid) {

	// endhalfwayarcshapegridFrom has already been copied
	if _endhalfwayarcshapegridTo, ok := mapOrigCopy[endhalfwayarcshapegridFrom]; ok {
		endhalfwayarcshapegridTo = _endhalfwayarcshapegridTo.(*EndHalfwayArcShapeGrid)
		return
	}

	endhalfwayarcshapegridTo = new(EndHalfwayArcShapeGrid)
	mapOrigCopy[endhalfwayarcshapegridFrom] = endhalfwayarcshapegridTo
	endhalfwayarcshapegridFrom.CopyBasicFields(endhalfwayarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchExplanationTextShape(mapOrigCopy map[any]any, explanationtextshapeFrom *ExplanationTextShape) (explanationtextshapeTo *ExplanationTextShape) {

	// explanationtextshapeFrom has already been copied
	if _explanationtextshapeTo, ok := mapOrigCopy[explanationtextshapeFrom]; ok {
		explanationtextshapeTo = _explanationtextshapeTo.(*ExplanationTextShape)
		return
	}

	explanationtextshapeTo = new(ExplanationTextShape)
	mapOrigCopy[explanationtextshapeFrom] = explanationtextshapeTo
	explanationtextshapeFrom.CopyBasicFields(explanationtextshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEye3DShape(mapOrigCopy map[any]any, eye3dshapeFrom *Eye3DShape) (eye3dshapeTo *Eye3DShape) {

	// eye3dshapeFrom has already been copied
	if _eye3dshapeTo, ok := mapOrigCopy[eye3dshapeFrom]; ok {
		eye3dshapeTo = _eye3dshapeTo.(*Eye3DShape)
		return
	}

	eye3dshapeTo = new(Eye3DShape)
	mapOrigCopy[eye3dshapeFrom] = eye3dshapeTo
	eye3dshapeFrom.CopyBasicFields(eye3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEyeCornersSampledPoints3DShape(mapOrigCopy map[any]any, eyecornerssampledpoints3dshapeFrom *EyeCornersSampledPoints3DShape) (eyecornerssampledpoints3dshapeTo *EyeCornersSampledPoints3DShape) {

	// eyecornerssampledpoints3dshapeFrom has already been copied
	if _eyecornerssampledpoints3dshapeTo, ok := mapOrigCopy[eyecornerssampledpoints3dshapeFrom]; ok {
		eyecornerssampledpoints3dshapeTo = _eyecornerssampledpoints3dshapeTo.(*EyeCornersSampledPoints3DShape)
		return
	}

	eyecornerssampledpoints3dshapeTo = new(EyeCornersSampledPoints3DShape)
	mapOrigCopy[eyecornerssampledpoints3dshapeFrom] = eyecornerssampledpoints3dshapeTo
	eyecornerssampledpoints3dshapeFrom.CopyBasicFields(eyecornerssampledpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEyeSampledPoints3DShape(mapOrigCopy map[any]any, eyesampledpoints3dshapeFrom *EyeSampledPoints3DShape) (eyesampledpoints3dshapeTo *EyeSampledPoints3DShape) {

	// eyesampledpoints3dshapeFrom has already been copied
	if _eyesampledpoints3dshapeTo, ok := mapOrigCopy[eyesampledpoints3dshapeFrom]; ok {
		eyesampledpoints3dshapeTo = _eyesampledpoints3dshapeTo.(*EyeSampledPoints3DShape)
		return
	}

	eyesampledpoints3dshapeTo = new(EyeSampledPoints3DShape)
	mapOrigCopy[eyesampledpoints3dshapeFrom] = eyesampledpoints3dshapeTo
	eyesampledpoints3dshapeFrom.CopyBasicFields(eyesampledpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEyeSeatBottomCurveShape(mapOrigCopy map[any]any, eyeseatbottomcurveshapeFrom *EyeSeatBottomCurveShape) (eyeseatbottomcurveshapeTo *EyeSeatBottomCurveShape) {

	// eyeseatbottomcurveshapeFrom has already been copied
	if _eyeseatbottomcurveshapeTo, ok := mapOrigCopy[eyeseatbottomcurveshapeFrom]; ok {
		eyeseatbottomcurveshapeTo = _eyeseatbottomcurveshapeTo.(*EyeSeatBottomCurveShape)
		return
	}

	eyeseatbottomcurveshapeTo = new(EyeSeatBottomCurveShape)
	mapOrigCopy[eyeseatbottomcurveshapeFrom] = eyeseatbottomcurveshapeTo
	eyeseatbottomcurveshapeFrom.CopyBasicFields(eyeseatbottomcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEyeStoolBottomCurveShape(mapOrigCopy map[any]any, eyestoolbottomcurveshapeFrom *EyeStoolBottomCurveShape) (eyestoolbottomcurveshapeTo *EyeStoolBottomCurveShape) {

	// eyestoolbottomcurveshapeFrom has already been copied
	if _eyestoolbottomcurveshapeTo, ok := mapOrigCopy[eyestoolbottomcurveshapeFrom]; ok {
		eyestoolbottomcurveshapeTo = _eyestoolbottomcurveshapeTo.(*EyeStoolBottomCurveShape)
		return
	}

	eyestoolbottomcurveshapeTo = new(EyeStoolBottomCurveShape)
	mapOrigCopy[eyestoolbottomcurveshapeFrom] = eyestoolbottomcurveshapeTo
	eyestoolbottomcurveshapeFrom.CopyBasicFields(eyestoolbottomcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEyeVolume3DShape(mapOrigCopy map[any]any, eyevolume3dshapeFrom *EyeVolume3DShape) (eyevolume3dshapeTo *EyeVolume3DShape) {

	// eyevolume3dshapeFrom has already been copied
	if _eyevolume3dshapeTo, ok := mapOrigCopy[eyevolume3dshapeFrom]; ok {
		eyevolume3dshapeTo = _eyevolume3dshapeTo.(*EyeVolume3DShape)
		return
	}

	eyevolume3dshapeTo = new(EyeVolume3DShape)
	mapOrigCopy[eyevolume3dshapeFrom] = eyevolume3dshapeTo
	eyevolume3dshapeFrom.CopyBasicFields(eyevolume3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGridPathShape(mapOrigCopy map[any]any, gridpathshapeFrom *GridPathShape) (gridpathshapeTo *GridPathShape) {

	// gridpathshapeFrom has already been copied
	if _gridpathshapeTo, ok := mapOrigCopy[gridpathshapeFrom]; ok {
		gridpathshapeTo = _gridpathshapeTo.(*GridPathShape)
		return
	}

	gridpathshapeTo = new(GridPathShape)
	mapOrigCopy[gridpathshapeFrom] = gridpathshapeTo
	gridpathshapeFrom.CopyBasicFields(gridpathshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGrowthCurve2D(mapOrigCopy map[any]any, growthcurve2dFrom *GrowthCurve2D) (growthcurve2dTo *GrowthCurve2D) {

	// growthcurve2dFrom has already been copied
	if _growthcurve2dTo, ok := mapOrigCopy[growthcurve2dFrom]; ok {
		growthcurve2dTo = _growthcurve2dTo.(*GrowthCurve2D)
		return
	}

	growthcurve2dTo = new(GrowthCurve2D)
	mapOrigCopy[growthcurve2dFrom] = growthcurve2dTo
	growthcurve2dFrom.CopyBasicFields(growthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGrowthCurve2DRibbon(mapOrigCopy map[any]any, growthcurve2dribbonFrom *GrowthCurve2DRibbon) (growthcurve2dribbonTo *GrowthCurve2DRibbon) {

	// growthcurve2dribbonFrom has already been copied
	if _growthcurve2dribbonTo, ok := mapOrigCopy[growthcurve2dribbonFrom]; ok {
		growthcurve2dribbonTo = _growthcurve2dribbonTo.(*GrowthCurve2DRibbon)
		return
	}

	growthcurve2dribbonTo = new(GrowthCurve2DRibbon)
	mapOrigCopy[growthcurve2dribbonFrom] = growthcurve2dribbonTo
	growthcurve2dribbonFrom.CopyBasicFields(growthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, growthcurve2dribbonendshapeFrom *GrowthCurve2DRibbonEndShape) (growthcurve2dribbonendshapeTo *GrowthCurve2DRibbonEndShape) {

	// growthcurve2dribbonendshapeFrom has already been copied
	if _growthcurve2dribbonendshapeTo, ok := mapOrigCopy[growthcurve2dribbonendshapeFrom]; ok {
		growthcurve2dribbonendshapeTo = _growthcurve2dribbonendshapeTo.(*GrowthCurve2DRibbonEndShape)
		return
	}

	growthcurve2dribbonendshapeTo = new(GrowthCurve2DRibbonEndShape)
	mapOrigCopy[growthcurve2dribbonendshapeFrom] = growthcurve2dribbonendshapeTo
	growthcurve2dribbonendshapeFrom.CopyBasicFields(growthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, growthcurve2dribbonstartshapeFrom *GrowthCurve2DRibbonStartShape) (growthcurve2dribbonstartshapeTo *GrowthCurve2DRibbonStartShape) {

	// growthcurve2dribbonstartshapeFrom has already been copied
	if _growthcurve2dribbonstartshapeTo, ok := mapOrigCopy[growthcurve2dribbonstartshapeFrom]; ok {
		growthcurve2dribbonstartshapeTo = _growthcurve2dribbonstartshapeTo.(*GrowthCurve2DRibbonStartShape)
		return
	}

	growthcurve2dribbonstartshapeTo = new(GrowthCurve2DRibbonStartShape)
	mapOrigCopy[growthcurve2dribbonstartshapeFrom] = growthcurve2dribbonstartshapeTo
	growthcurve2dribbonstartshapeFrom.CopyBasicFields(growthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGrowthCurveRhombusGridShape(mapOrigCopy map[any]any, growthcurverhombusgridshapeFrom *GrowthCurveRhombusGridShape) (growthcurverhombusgridshapeTo *GrowthCurveRhombusGridShape) {

	// growthcurverhombusgridshapeFrom has already been copied
	if _growthcurverhombusgridshapeTo, ok := mapOrigCopy[growthcurverhombusgridshapeFrom]; ok {
		growthcurverhombusgridshapeTo = _growthcurverhombusgridshapeTo.(*GrowthCurveRhombusGridShape)
		return
	}

	growthcurverhombusgridshapeTo = new(GrowthCurveRhombusGridShape)
	mapOrigCopy[growthcurverhombusgridshapeFrom] = growthcurverhombusgridshapeTo
	growthcurverhombusgridshapeFrom.CopyBasicFields(growthcurverhombusgridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGrowthCurveRhombusShape(mapOrigCopy map[any]any, growthcurverhombusshapeFrom *GrowthCurveRhombusShape) (growthcurverhombusshapeTo *GrowthCurveRhombusShape) {

	// growthcurverhombusshapeFrom has already been copied
	if _growthcurverhombusshapeTo, ok := mapOrigCopy[growthcurverhombusshapeFrom]; ok {
		growthcurverhombusshapeTo = _growthcurverhombusshapeTo.(*GrowthCurveRhombusShape)
		return
	}

	growthcurverhombusshapeTo = new(GrowthCurveRhombusShape)
	mapOrigCopy[growthcurverhombusshapeFrom] = growthcurverhombusshapeTo
	growthcurverhombusshapeFrom.CopyBasicFields(growthcurverhombusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGrowthVectorShape(mapOrigCopy map[any]any, growthvectorshapeFrom *GrowthVectorShape) (growthvectorshapeTo *GrowthVectorShape) {

	// growthvectorshapeFrom has already been copied
	if _growthvectorshapeTo, ok := mapOrigCopy[growthvectorshapeFrom]; ok {
		growthvectorshapeTo = _growthvectorshapeTo.(*GrowthVectorShape)
		return
	}

	growthvectorshapeTo = new(GrowthVectorShape)
	mapOrigCopy[growthvectorshapeFrom] = growthvectorshapeTo
	growthvectorshapeFrom.CopyBasicFields(growthvectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchInitialRhombusGridShape(mapOrigCopy map[any]any, initialrhombusgridshapeFrom *InitialRhombusGridShape) (initialrhombusgridshapeTo *InitialRhombusGridShape) {

	// initialrhombusgridshapeFrom has already been copied
	if _initialrhombusgridshapeTo, ok := mapOrigCopy[initialrhombusgridshapeFrom]; ok {
		initialrhombusgridshapeTo = _initialrhombusgridshapeTo.(*InitialRhombusGridShape)
		return
	}

	initialrhombusgridshapeTo = new(InitialRhombusGridShape)
	mapOrigCopy[initialrhombusgridshapeFrom] = initialrhombusgridshapeTo
	initialrhombusgridshapeFrom.CopyBasicFields(initialrhombusgridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchInitialRhombusShape(mapOrigCopy map[any]any, initialrhombusshapeFrom *InitialRhombusShape) (initialrhombusshapeTo *InitialRhombusShape) {

	// initialrhombusshapeFrom has already been copied
	if _initialrhombusshapeTo, ok := mapOrigCopy[initialrhombusshapeFrom]; ok {
		initialrhombusshapeTo = _initialrhombusshapeTo.(*InitialRhombusShape)
		return
	}

	initialrhombusshapeTo = new(InitialRhombusShape)
	mapOrigCopy[initialrhombusshapeFrom] = initialrhombusshapeTo
	initialrhombusshapeFrom.CopyBasicFields(initialrhombusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchKey3DShape(mapOrigCopy map[any]any, key3dshapeFrom *Key3DShape) (key3dshapeTo *Key3DShape) {

	// key3dshapeFrom has already been copied
	if _key3dshapeTo, ok := mapOrigCopy[key3dshapeFrom]; ok {
		key3dshapeTo = _key3dshapeTo.(*Key3DShape)
		return
	}

	key3dshapeTo = new(Key3DShape)
	mapOrigCopy[key3dshapeFrom] = key3dshapeTo
	key3dshapeFrom.CopyBasicFields(key3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchKeyHole3DShape(mapOrigCopy map[any]any, keyhole3dshapeFrom *KeyHole3DShape) (keyhole3dshapeTo *KeyHole3DShape) {

	// keyhole3dshapeFrom has already been copied
	if _keyhole3dshapeTo, ok := mapOrigCopy[keyhole3dshapeFrom]; ok {
		keyhole3dshapeTo = _keyhole3dshapeTo.(*KeyHole3DShape)
		return
	}

	keyhole3dshapeTo = new(KeyHole3DShape)
	mapOrigCopy[keyhole3dshapeFrom] = keyhole3dshapeTo
	keyhole3dshapeFrom.CopyBasicFields(keyhole3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchKeyHoleShape(mapOrigCopy map[any]any, keyholeshapeFrom *KeyHoleShape) (keyholeshapeTo *KeyHoleShape) {

	// keyholeshapeFrom has already been copied
	if _keyholeshapeTo, ok := mapOrigCopy[keyholeshapeFrom]; ok {
		keyholeshapeTo = _keyholeshapeTo.(*KeyHoleShape)
		return
	}

	keyholeshapeTo = new(KeyHoleShape)
	mapOrigCopy[keyholeshapeFrom] = keyholeshapeTo
	keyholeshapeFrom.CopyBasicFields(keyholeshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {

	// libraryFrom has already been copied
	if _libraryTo, ok := mapOrigCopy[libraryFrom]; ok {
		libraryTo = _libraryTo.(*Library)
		return
	}

	libraryTo = new(Library)
	mapOrigCopy[libraryFrom] = libraryTo
	libraryFrom.CopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, CopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _plantabstract := range libraryFrom.Plants {
		libraryTo.Plants = append(libraryTo.Plants, CopyBranchPlantAbstract(mapOrigCopy, _plantabstract))
	}

	return
}

func CopyBranchMidArcVectorShape(mapOrigCopy map[any]any, midarcvectorshapeFrom *MidArcVectorShape) (midarcvectorshapeTo *MidArcVectorShape) {

	// midarcvectorshapeFrom has already been copied
	if _midarcvectorshapeTo, ok := mapOrigCopy[midarcvectorshapeFrom]; ok {
		midarcvectorshapeTo = _midarcvectorshapeTo.(*MidArcVectorShape)
		return
	}

	midarcvectorshapeTo = new(MidArcVectorShape)
	mapOrigCopy[midarcvectorshapeFrom] = midarcvectorshapeTo
	midarcvectorshapeFrom.CopyBasicFields(midarcvectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMidArcVectorShapeGrid(mapOrigCopy map[any]any, midarcvectorshapegridFrom *MidArcVectorShapeGrid) (midarcvectorshapegridTo *MidArcVectorShapeGrid) {

	// midarcvectorshapegridFrom has already been copied
	if _midarcvectorshapegridTo, ok := mapOrigCopy[midarcvectorshapegridFrom]; ok {
		midarcvectorshapegridTo = _midarcvectorshapegridTo.(*MidArcVectorShapeGrid)
		return
	}

	midarcvectorshapegridTo = new(MidArcVectorShapeGrid)
	mapOrigCopy[midarcvectorshapegridFrom] = midarcvectorshapegridTo
	midarcvectorshapegridFrom.CopyBasicFields(midarcvectorshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchOriginalPoints3DShape(mapOrigCopy map[any]any, originalpoints3dshapeFrom *OriginalPoints3DShape) (originalpoints3dshapeTo *OriginalPoints3DShape) {

	// originalpoints3dshapeFrom has already been copied
	if _originalpoints3dshapeTo, ok := mapOrigCopy[originalpoints3dshapeFrom]; ok {
		originalpoints3dshapeTo = _originalpoints3dshapeTo.(*OriginalPoints3DShape)
		return
	}

	originalpoints3dshapeTo = new(OriginalPoints3DShape)
	mapOrigCopy[originalpoints3dshapeFrom] = originalpoints3dshapeTo
	originalpoints3dshapeFrom.CopyBasicFields(originalpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchParastichyMCurves3DShape(mapOrigCopy map[any]any, parastichymcurves3dshapeFrom *ParastichyMCurves3DShape) (parastichymcurves3dshapeTo *ParastichyMCurves3DShape) {

	// parastichymcurves3dshapeFrom has already been copied
	if _parastichymcurves3dshapeTo, ok := mapOrigCopy[parastichymcurves3dshapeFrom]; ok {
		parastichymcurves3dshapeTo = _parastichymcurves3dshapeTo.(*ParastichyMCurves3DShape)
		return
	}

	parastichymcurves3dshapeTo = new(ParastichyMCurves3DShape)
	mapOrigCopy[parastichymcurves3dshapeFrom] = parastichymcurves3dshapeTo
	parastichymcurves3dshapeFrom.CopyBasicFields(parastichymcurves3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchParastichyNCurves3DShape(mapOrigCopy map[any]any, parastichyncurves3dshapeFrom *ParastichyNCurves3DShape) (parastichyncurves3dshapeTo *ParastichyNCurves3DShape) {

	// parastichyncurves3dshapeFrom has already been copied
	if _parastichyncurves3dshapeTo, ok := mapOrigCopy[parastichyncurves3dshapeFrom]; ok {
		parastichyncurves3dshapeTo = _parastichyncurves3dshapeTo.(*ParastichyNCurves3DShape)
		return
	}

	parastichyncurves3dshapeTo = new(ParastichyNCurves3DShape)
	mapOrigCopy[parastichyncurves3dshapeFrom] = parastichyncurves3dshapeTo
	parastichyncurves3dshapeFrom.CopyBasicFields(parastichyncurves3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyGrowthCurve2DRibbon(mapOrigCopy map[any]any, partiallygrowthcurve2dribbonFrom *PartiallyGrowthCurve2DRibbon) (partiallygrowthcurve2dribbonTo *PartiallyGrowthCurve2DRibbon) {

	// partiallygrowthcurve2dribbonFrom has already been copied
	if _partiallygrowthcurve2dribbonTo, ok := mapOrigCopy[partiallygrowthcurve2dribbonFrom]; ok {
		partiallygrowthcurve2dribbonTo = _partiallygrowthcurve2dribbonTo.(*PartiallyGrowthCurve2DRibbon)
		return
	}

	partiallygrowthcurve2dribbonTo = new(PartiallyGrowthCurve2DRibbon)
	mapOrigCopy[partiallygrowthcurve2dribbonFrom] = partiallygrowthcurve2dribbonTo
	partiallygrowthcurve2dribbonFrom.CopyBasicFields(partiallygrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, partiallygrowthcurve2dribbonendshapeFrom *PartiallyGrowthCurve2DRibbonEndShape) (partiallygrowthcurve2dribbonendshapeTo *PartiallyGrowthCurve2DRibbonEndShape) {

	// partiallygrowthcurve2dribbonendshapeFrom has already been copied
	if _partiallygrowthcurve2dribbonendshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dribbonendshapeFrom]; ok {
		partiallygrowthcurve2dribbonendshapeTo = _partiallygrowthcurve2dribbonendshapeTo.(*PartiallyGrowthCurve2DRibbonEndShape)
		return
	}

	partiallygrowthcurve2dribbonendshapeTo = new(PartiallyGrowthCurve2DRibbonEndShape)
	mapOrigCopy[partiallygrowthcurve2dribbonendshapeFrom] = partiallygrowthcurve2dribbonendshapeTo
	partiallygrowthcurve2dribbonendshapeFrom.CopyBasicFields(partiallygrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, partiallygrowthcurve2dribbonstartshapeFrom *PartiallyGrowthCurve2DRibbonStartShape) (partiallygrowthcurve2dribbonstartshapeTo *PartiallyGrowthCurve2DRibbonStartShape) {

	// partiallygrowthcurve2dribbonstartshapeFrom has already been copied
	if _partiallygrowthcurve2dribbonstartshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dribbonstartshapeFrom]; ok {
		partiallygrowthcurve2dribbonstartshapeTo = _partiallygrowthcurve2dribbonstartshapeTo.(*PartiallyGrowthCurve2DRibbonStartShape)
		return
	}

	partiallygrowthcurve2dribbonstartshapeTo = new(PartiallyGrowthCurve2DRibbonStartShape)
	mapOrigCopy[partiallygrowthcurve2dribbonstartshapeFrom] = partiallygrowthcurve2dribbonstartshapeTo
	partiallygrowthcurve2dribbonstartshapeFrom.CopyBasicFields(partiallygrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyGrowthCurve2DTrajectory(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryFrom *PartiallyGrowthCurve2DTrajectory) (partiallygrowthcurve2dtrajectoryTo *PartiallyGrowthCurve2DTrajectory) {

	// partiallygrowthcurve2dtrajectoryFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryFrom]; ok {
		partiallygrowthcurve2dtrajectoryTo = _partiallygrowthcurve2dtrajectoryTo.(*PartiallyGrowthCurve2DTrajectory)
		return
	}

	partiallygrowthcurve2dtrajectoryTo = new(PartiallyGrowthCurve2DTrajectory)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryFrom] = partiallygrowthcurve2dtrajectoryTo
	partiallygrowthcurve2dtrajectoryFrom.CopyBasicFields(partiallygrowthcurve2dtrajectoryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyGrowthCurve2DTrajectoryP1CurveShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp1curveshapeFrom *PartiallyGrowthCurve2DTrajectoryP1CurveShape) (partiallygrowthcurve2dtrajectoryp1curveshapeTo *PartiallyGrowthCurve2DTrajectoryP1CurveShape) {

	// partiallygrowthcurve2dtrajectoryp1curveshapeFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryp1curveshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryp1curveshapeFrom]; ok {
		partiallygrowthcurve2dtrajectoryp1curveshapeTo = _partiallygrowthcurve2dtrajectoryp1curveshapeTo.(*PartiallyGrowthCurve2DTrajectoryP1CurveShape)
		return
	}

	partiallygrowthcurve2dtrajectoryp1curveshapeTo = new(PartiallyGrowthCurve2DTrajectoryP1CurveShape)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryp1curveshapeFrom] = partiallygrowthcurve2dtrajectoryp1curveshapeTo
	partiallygrowthcurve2dtrajectoryp1curveshapeFrom.CopyBasicFields(partiallygrowthcurve2dtrajectoryp1curveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyGrowthCurve2DTrajectoryP1P2(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp1p2From *PartiallyGrowthCurve2DTrajectoryP1P2) (partiallygrowthcurve2dtrajectoryp1p2To *PartiallyGrowthCurve2DTrajectoryP1P2) {

	// partiallygrowthcurve2dtrajectoryp1p2From has already been copied
	if _partiallygrowthcurve2dtrajectoryp1p2To, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryp1p2From]; ok {
		partiallygrowthcurve2dtrajectoryp1p2To = _partiallygrowthcurve2dtrajectoryp1p2To.(*PartiallyGrowthCurve2DTrajectoryP1P2)
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2To = new(PartiallyGrowthCurve2DTrajectoryP1P2)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryp1p2From] = partiallygrowthcurve2dtrajectoryp1p2To
	partiallygrowthcurve2dtrajectoryp1p2From.CopyBasicFields(partiallygrowthcurve2dtrajectoryp1p2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFrom *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) (partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) {

	// partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFrom]; ok {
		partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo = _partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo.(*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo = new(PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFrom] = partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo
	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFrom.CopyBasicFields(partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyGrowthCurve2DTrajectoryP1PointShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp1pointshapeFrom *PartiallyGrowthCurve2DTrajectoryP1PointShape) (partiallygrowthcurve2dtrajectoryp1pointshapeTo *PartiallyGrowthCurve2DTrajectoryP1PointShape) {

	// partiallygrowthcurve2dtrajectoryp1pointshapeFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryp1pointshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryp1pointshapeFrom]; ok {
		partiallygrowthcurve2dtrajectoryp1pointshapeTo = _partiallygrowthcurve2dtrajectoryp1pointshapeTo.(*PartiallyGrowthCurve2DTrajectoryP1PointShape)
		return
	}

	partiallygrowthcurve2dtrajectoryp1pointshapeTo = new(PartiallyGrowthCurve2DTrajectoryP1PointShape)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryp1pointshapeFrom] = partiallygrowthcurve2dtrajectoryp1pointshapeTo
	partiallygrowthcurve2dtrajectoryp1pointshapeFrom.CopyBasicFields(partiallygrowthcurve2dtrajectoryp1pointshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyGrowthCurve2DTrajectoryP2CurveShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp2curveshapeFrom *PartiallyGrowthCurve2DTrajectoryP2CurveShape) (partiallygrowthcurve2dtrajectoryp2curveshapeTo *PartiallyGrowthCurve2DTrajectoryP2CurveShape) {

	// partiallygrowthcurve2dtrajectoryp2curveshapeFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryp2curveshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryp2curveshapeFrom]; ok {
		partiallygrowthcurve2dtrajectoryp2curveshapeTo = _partiallygrowthcurve2dtrajectoryp2curveshapeTo.(*PartiallyGrowthCurve2DTrajectoryP2CurveShape)
		return
	}

	partiallygrowthcurve2dtrajectoryp2curveshapeTo = new(PartiallyGrowthCurve2DTrajectoryP2CurveShape)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryp2curveshapeFrom] = partiallygrowthcurve2dtrajectoryp2curveshapeTo
	partiallygrowthcurve2dtrajectoryp2curveshapeFrom.CopyBasicFields(partiallygrowthcurve2dtrajectoryp2curveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyGrowthCurve2DTrajectoryP2PointShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp2pointshapeFrom *PartiallyGrowthCurve2DTrajectoryP2PointShape) (partiallygrowthcurve2dtrajectoryp2pointshapeTo *PartiallyGrowthCurve2DTrajectoryP2PointShape) {

	// partiallygrowthcurve2dtrajectoryp2pointshapeFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryp2pointshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryp2pointshapeFrom]; ok {
		partiallygrowthcurve2dtrajectoryp2pointshapeTo = _partiallygrowthcurve2dtrajectoryp2pointshapeTo.(*PartiallyGrowthCurve2DTrajectoryP2PointShape)
		return
	}

	partiallygrowthcurve2dtrajectoryp2pointshapeTo = new(PartiallyGrowthCurve2DTrajectoryP2PointShape)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryp2pointshapeFrom] = partiallygrowthcurve2dtrajectoryp2pointshapeTo
	partiallygrowthcurve2dtrajectoryp2pointshapeFrom.CopyBasicFields(partiallygrowthcurve2dtrajectoryp2pointshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyGrowthCurve2DTrajectoryShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryshapeFrom *PartiallyGrowthCurve2DTrajectoryShape) (partiallygrowthcurve2dtrajectoryshapeTo *PartiallyGrowthCurve2DTrajectoryShape) {

	// partiallygrowthcurve2dtrajectoryshapeFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryshapeFrom]; ok {
		partiallygrowthcurve2dtrajectoryshapeTo = _partiallygrowthcurve2dtrajectoryshapeTo.(*PartiallyGrowthCurve2DTrajectoryShape)
		return
	}

	partiallygrowthcurve2dtrajectoryshapeTo = new(PartiallyGrowthCurve2DTrajectoryShape)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryshapeFrom] = partiallygrowthcurve2dtrajectoryshapeTo
	partiallygrowthcurve2dtrajectoryshapeFrom.CopyBasicFields(partiallygrowthcurve2dtrajectoryshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyRotatedSeatBottomCurveShape(mapOrigCopy map[any]any, partiallyrotatedseatbottomcurveshapeFrom *PartiallyRotatedSeatBottomCurveShape) (partiallyrotatedseatbottomcurveshapeTo *PartiallyRotatedSeatBottomCurveShape) {

	// partiallyrotatedseatbottomcurveshapeFrom has already been copied
	if _partiallyrotatedseatbottomcurveshapeTo, ok := mapOrigCopy[partiallyrotatedseatbottomcurveshapeFrom]; ok {
		partiallyrotatedseatbottomcurveshapeTo = _partiallyrotatedseatbottomcurveshapeTo.(*PartiallyRotatedSeatBottomCurveShape)
		return
	}

	partiallyrotatedseatbottomcurveshapeTo = new(PartiallyRotatedSeatBottomCurveShape)
	mapOrigCopy[partiallyrotatedseatbottomcurveshapeFrom] = partiallyrotatedseatbottomcurveshapeTo
	partiallyrotatedseatbottomcurveshapeFrom.CopyBasicFields(partiallyrotatedseatbottomcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyRotatedSeatTopCurveShape(mapOrigCopy map[any]any, partiallyrotatedseattopcurveshapeFrom *PartiallyRotatedSeatTopCurveShape) (partiallyrotatedseattopcurveshapeTo *PartiallyRotatedSeatTopCurveShape) {

	// partiallyrotatedseattopcurveshapeFrom has already been copied
	if _partiallyrotatedseattopcurveshapeTo, ok := mapOrigCopy[partiallyrotatedseattopcurveshapeFrom]; ok {
		partiallyrotatedseattopcurveshapeTo = _partiallyrotatedseattopcurveshapeTo.(*PartiallyRotatedSeatTopCurveShape)
		return
	}

	partiallyrotatedseattopcurveshapeTo = new(PartiallyRotatedSeatTopCurveShape)
	mapOrigCopy[partiallyrotatedseattopcurveshapeFrom] = partiallyrotatedseattopcurveshapeTo
	partiallyrotatedseattopcurveshapeFrom.CopyBasicFields(partiallyrotatedseattopcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPartiallyRotatedTorusShape(mapOrigCopy map[any]any, partiallyrotatedtorusshapeFrom *PartiallyRotatedTorusShape) (partiallyrotatedtorusshapeTo *PartiallyRotatedTorusShape) {

	// partiallyrotatedtorusshapeFrom has already been copied
	if _partiallyrotatedtorusshapeTo, ok := mapOrigCopy[partiallyrotatedtorusshapeFrom]; ok {
		partiallyrotatedtorusshapeTo = _partiallyrotatedtorusshapeTo.(*PartiallyRotatedTorusShape)
		return
	}

	partiallyrotatedtorusshapeTo = new(PartiallyRotatedTorusShape)
	mapOrigCopy[partiallyrotatedtorusshapeFrom] = partiallyrotatedtorusshapeTo
	partiallyrotatedtorusshapeFrom.CopyBasicFields(partiallyrotatedtorusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPerpendicularVector(mapOrigCopy map[any]any, perpendicularvectorFrom *PerpendicularVector) (perpendicularvectorTo *PerpendicularVector) {

	// perpendicularvectorFrom has already been copied
	if _perpendicularvectorTo, ok := mapOrigCopy[perpendicularvectorFrom]; ok {
		perpendicularvectorTo = _perpendicularvectorTo.(*PerpendicularVector)
		return
	}

	perpendicularvectorTo = new(PerpendicularVector)
	mapOrigCopy[perpendicularvectorFrom] = perpendicularvectorTo
	perpendicularvectorFrom.CopyBasicFields(perpendicularvectorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPerpendicularVectorGrid(mapOrigCopy map[any]any, perpendicularvectorgridFrom *PerpendicularVectorGrid) (perpendicularvectorgridTo *PerpendicularVectorGrid) {

	// perpendicularvectorgridFrom has already been copied
	if _perpendicularvectorgridTo, ok := mapOrigCopy[perpendicularvectorgridFrom]; ok {
		perpendicularvectorgridTo = _perpendicularvectorgridTo.(*PerpendicularVectorGrid)
		return
	}

	perpendicularvectorgridTo = new(PerpendicularVectorGrid)
	mapOrigCopy[perpendicularvectorgridFrom] = perpendicularvectorgridTo
	perpendicularvectorgridFrom.CopyBasicFields(perpendicularvectorgridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPerpendicularVectorGridHalfway(mapOrigCopy map[any]any, perpendicularvectorgridhalfwayFrom *PerpendicularVectorGridHalfway) (perpendicularvectorgridhalfwayTo *PerpendicularVectorGridHalfway) {

	// perpendicularvectorgridhalfwayFrom has already been copied
	if _perpendicularvectorgridhalfwayTo, ok := mapOrigCopy[perpendicularvectorgridhalfwayFrom]; ok {
		perpendicularvectorgridhalfwayTo = _perpendicularvectorgridhalfwayTo.(*PerpendicularVectorGridHalfway)
		return
	}

	perpendicularvectorgridhalfwayTo = new(PerpendicularVectorGridHalfway)
	mapOrigCopy[perpendicularvectorgridhalfwayFrom] = perpendicularvectorgridhalfwayTo
	perpendicularvectorgridhalfwayFrom.CopyBasicFields(perpendicularvectorgridhalfwayTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPerpendicularVectorHalfway(mapOrigCopy map[any]any, perpendicularvectorhalfwayFrom *PerpendicularVectorHalfway) (perpendicularvectorhalfwayTo *PerpendicularVectorHalfway) {

	// perpendicularvectorhalfwayFrom has already been copied
	if _perpendicularvectorhalfwayTo, ok := mapOrigCopy[perpendicularvectorhalfwayFrom]; ok {
		perpendicularvectorhalfwayTo = _perpendicularvectorhalfwayTo.(*PerpendicularVectorHalfway)
		return
	}

	perpendicularvectorhalfwayTo = new(PerpendicularVectorHalfway)
	mapOrigCopy[perpendicularvectorhalfwayFrom] = perpendicularvectorhalfwayTo
	perpendicularvectorhalfwayFrom.CopyBasicFields(perpendicularvectorhalfwayTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPlant2DDiagram(mapOrigCopy map[any]any, plant2ddiagramFrom *Plant2DDiagram) (plant2ddiagramTo *Plant2DDiagram) {

	// plant2ddiagramFrom has already been copied
	if _plant2ddiagramTo, ok := mapOrigCopy[plant2ddiagramFrom]; ok {
		plant2ddiagramTo = _plant2ddiagramTo.(*Plant2DDiagram)
		return
	}

	plant2ddiagramTo = new(Plant2DDiagram)
	mapOrigCopy[plant2ddiagramFrom] = plant2ddiagramTo
	plant2ddiagramFrom.CopyBasicFields(plant2ddiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPlant3DDiagram(mapOrigCopy map[any]any, plant3ddiagramFrom *Plant3DDiagram) (plant3ddiagramTo *Plant3DDiagram) {

	// plant3ddiagramFrom has already been copied
	if _plant3ddiagramTo, ok := mapOrigCopy[plant3ddiagramFrom]; ok {
		plant3ddiagramTo = _plant3ddiagramTo.(*Plant3DDiagram)
		return
	}

	plant3ddiagramTo = new(Plant3DDiagram)
	mapOrigCopy[plant3ddiagramFrom] = plant3ddiagramTo
	plant3ddiagramFrom.CopyBasicFields(plant3ddiagramTo)

	//insertion point for the staging of instances referenced by pointers
	if plant3ddiagramFrom.StemCylinder3DShape != nil {
		plant3ddiagramTo.StemCylinder3DShape = CopyBranchStemCylinder3DShape(mapOrigCopy, plant3ddiagramFrom.StemCylinder3DShape)
	}
	if plant3ddiagramFrom.ParastichyNCurves3DShape != nil {
		plant3ddiagramTo.ParastichyNCurves3DShape = CopyBranchParastichyNCurves3DShape(mapOrigCopy, plant3ddiagramFrom.ParastichyNCurves3DShape)
	}
	if plant3ddiagramFrom.ParastichyMCurves3DShape != nil {
		plant3ddiagramTo.ParastichyMCurves3DShape = CopyBranchParastichyMCurves3DShape(mapOrigCopy, plant3ddiagramFrom.ParastichyMCurves3DShape)
	}
	if plant3ddiagramFrom.CutLine3DShape != nil {
		plant3ddiagramTo.CutLine3DShape = CopyBranchCutLine3DShape(mapOrigCopy, plant3ddiagramFrom.CutLine3DShape)
	}
	if plant3ddiagramFrom.Circumference3DShape != nil {
		plant3ddiagramTo.Circumference3DShape = CopyBranchCircumference3DShape(mapOrigCopy, plant3ddiagramFrom.Circumference3DShape)
	}
	if plant3ddiagramFrom.Rendered3DShape != nil {
		plant3ddiagramTo.Rendered3DShape = CopyBranchRendered3DShape(mapOrigCopy, plant3ddiagramFrom.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPlantAbstract(mapOrigCopy map[any]any, plantabstractFrom *PlantAbstract) (plantabstractTo *PlantAbstract) {

	// plantabstractFrom has already been copied
	if _plantabstractTo, ok := mapOrigCopy[plantabstractFrom]; ok {
		plantabstractTo = _plantabstractTo.(*PlantAbstract)
		return
	}

	plantabstractTo = new(PlantAbstract)
	mapOrigCopy[plantabstractFrom] = plantabstractTo
	plantabstractFrom.CopyBasicFields(plantabstractTo)

	//insertion point for the staging of instances referenced by pointers
	if plantabstractFrom.VaseAbstract != nil {
		plantabstractTo.VaseAbstract = CopyBranchVaseAbstract(mapOrigCopy, plantabstractFrom.VaseAbstract)
	}
	if plantabstractFrom.StoolAbstract != nil {
		plantabstractTo.StoolAbstract = CopyBranchStoolAbstract(mapOrigCopy, plantabstractFrom.StoolAbstract)
	}
	if plantabstractFrom.ClockAbstract != nil {
		plantabstractTo.ClockAbstract = CopyBranchClockAbstract(mapOrigCopy, plantabstractFrom.ClockAbstract)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plant2ddiagram := range plantabstractFrom.Plant2DDiagrams {
		plantabstractTo.Plant2DDiagrams = append(plantabstractTo.Plant2DDiagrams, CopyBranchPlant2DDiagram(mapOrigCopy, _plant2ddiagram))
	}
	for _, _plant3ddiagram := range plantabstractFrom.Plant3DDiagrams {
		plantabstractTo.Plant3DDiagrams = append(plantabstractTo.Plant3DDiagrams, CopyBranchPlant3DDiagram(mapOrigCopy, _plant3ddiagram))
	}
	for _, _vase2ddiagram := range plantabstractFrom.Vase2DDiagrams {
		plantabstractTo.Vase2DDiagrams = append(plantabstractTo.Vase2DDiagrams, CopyBranchVase2DDiagram(mapOrigCopy, _vase2ddiagram))
	}
	for _, _vase3ddiagram := range plantabstractFrom.Vase3DDiagrams {
		plantabstractTo.Vase3DDiagrams = append(plantabstractTo.Vase3DDiagrams, CopyBranchVase3DDiagram(mapOrigCopy, _vase3ddiagram))
	}
	for _, _stool2ddiagram := range plantabstractFrom.Stool2DDiagrams {
		plantabstractTo.Stool2DDiagrams = append(plantabstractTo.Stool2DDiagrams, CopyBranchStool2DDiagram(mapOrigCopy, _stool2ddiagram))
	}
	for _, _stool3ddiagram := range plantabstractFrom.Stool3DDiagrams {
		plantabstractTo.Stool3DDiagrams = append(plantabstractTo.Stool3DDiagrams, CopyBranchStool3DDiagram(mapOrigCopy, _stool3ddiagram))
	}
	for _, _clock2ddiagram := range plantabstractFrom.Clock2DDiagrams {
		plantabstractTo.Clock2DDiagrams = append(plantabstractTo.Clock2DDiagrams, CopyBranchClock2DDiagram(mapOrigCopy, _clock2ddiagram))
	}
	for _, _clock3ddiagram := range plantabstractFrom.Clock3DDiagrams {
		plantabstractTo.Clock3DDiagrams = append(plantabstractTo.Clock3DDiagrams, CopyBranchClock3DDiagram(mapOrigCopy, _clock3ddiagram))
	}

	return
}

func CopyBranchPlantCircumferenceShape(mapOrigCopy map[any]any, plantcircumferenceshapeFrom *PlantCircumferenceShape) (plantcircumferenceshapeTo *PlantCircumferenceShape) {

	// plantcircumferenceshapeFrom has already been copied
	if _plantcircumferenceshapeTo, ok := mapOrigCopy[plantcircumferenceshapeFrom]; ok {
		plantcircumferenceshapeTo = _plantcircumferenceshapeTo.(*PlantCircumferenceShape)
		return
	}

	plantcircumferenceshapeTo = new(PlantCircumferenceShape)
	mapOrigCopy[plantcircumferenceshapeFrom] = plantcircumferenceshapeTo
	plantcircumferenceshapeFrom.CopyBasicFields(plantcircumferenceshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPointsAndLines3DShape(mapOrigCopy map[any]any, pointsandlines3dshapeFrom *PointsAndLines3DShape) (pointsandlines3dshapeTo *PointsAndLines3DShape) {

	// pointsandlines3dshapeFrom has already been copied
	if _pointsandlines3dshapeTo, ok := mapOrigCopy[pointsandlines3dshapeFrom]; ok {
		pointsandlines3dshapeTo = _pointsandlines3dshapeTo.(*PointsAndLines3DShape)
		return
	}

	pointsandlines3dshapeTo = new(PointsAndLines3DShape)
	mapOrigCopy[pointsandlines3dshapeFrom] = pointsandlines3dshapeTo
	pointsandlines3dshapeFrom.CopyBasicFields(pointsandlines3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPxShape(mapOrigCopy map[any]any, pxshapeFrom *PxShape) (pxshapeTo *PxShape) {

	// pxshapeFrom has already been copied
	if _pxshapeTo, ok := mapOrigCopy[pxshapeFrom]; ok {
		pxshapeTo = _pxshapeTo.(*PxShape)
		return
	}

	pxshapeTo = new(PxShape)
	mapOrigCopy[pxshapeFrom] = pxshapeTo
	pxshapeFrom.CopyBasicFields(pxshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRendered3DShape(mapOrigCopy map[any]any, rendered3dshapeFrom *Rendered3DShape) (rendered3dshapeTo *Rendered3DShape) {

	// rendered3dshapeFrom has already been copied
	if _rendered3dshapeTo, ok := mapOrigCopy[rendered3dshapeFrom]; ok {
		rendered3dshapeTo = _rendered3dshapeTo.(*Rendered3DShape)
		return
	}

	rendered3dshapeTo = new(Rendered3DShape)
	mapOrigCopy[rendered3dshapeFrom] = rendered3dshapeTo
	rendered3dshapeFrom.CopyBasicFields(rendered3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRhombusShape(mapOrigCopy map[any]any, rhombusshapeFrom *RhombusShape) (rhombusshapeTo *RhombusShape) {

	// rhombusshapeFrom has already been copied
	if _rhombusshapeTo, ok := mapOrigCopy[rhombusshapeFrom]; ok {
		rhombusshapeTo = _rhombusshapeTo.(*RhombusShape)
		return
	}

	rhombusshapeTo = new(RhombusShape)
	mapOrigCopy[rhombusshapeFrom] = rhombusshapeTo
	rhombusshapeFrom.CopyBasicFields(rhombusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRhombusStuff(mapOrigCopy map[any]any, rhombusstuffFrom *RhombusStuff) (rhombusstuffTo *RhombusStuff) {

	// rhombusstuffFrom has already been copied
	if _rhombusstuffTo, ok := mapOrigCopy[rhombusstuffFrom]; ok {
		rhombusstuffTo = _rhombusstuffTo.(*RhombusStuff)
		return
	}

	rhombusstuffTo = new(RhombusStuff)
	mapOrigCopy[rhombusstuffFrom] = rhombusstuffTo
	rhombusstuffFrom.CopyBasicFields(rhombusstuffTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRotatedRhombusGridShape(mapOrigCopy map[any]any, rotatedrhombusgridshapeFrom *RotatedRhombusGridShape) (rotatedrhombusgridshapeTo *RotatedRhombusGridShape) {

	// rotatedrhombusgridshapeFrom has already been copied
	if _rotatedrhombusgridshapeTo, ok := mapOrigCopy[rotatedrhombusgridshapeFrom]; ok {
		rotatedrhombusgridshapeTo = _rotatedrhombusgridshapeTo.(*RotatedRhombusGridShape)
		return
	}

	rotatedrhombusgridshapeTo = new(RotatedRhombusGridShape)
	mapOrigCopy[rotatedrhombusgridshapeFrom] = rotatedrhombusgridshapeTo
	rotatedrhombusgridshapeFrom.CopyBasicFields(rotatedrhombusgridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRotatedRhombusShape(mapOrigCopy map[any]any, rotatedrhombusshapeFrom *RotatedRhombusShape) (rotatedrhombusshapeTo *RotatedRhombusShape) {

	// rotatedrhombusshapeFrom has already been copied
	if _rotatedrhombusshapeTo, ok := mapOrigCopy[rotatedrhombusshapeFrom]; ok {
		rotatedrhombusshapeTo = _rotatedrhombusshapeTo.(*RotatedRhombusShape)
		return
	}

	rotatedrhombusshapeTo = new(RotatedRhombusShape)
	mapOrigCopy[rotatedrhombusshapeFrom] = rotatedrhombusshapeTo
	rotatedrhombusshapeFrom.CopyBasicFields(rotatedrhombusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRotatedSampledPoints3DShape(mapOrigCopy map[any]any, rotatedsampledpoints3dshapeFrom *RotatedSampledPoints3DShape) (rotatedsampledpoints3dshapeTo *RotatedSampledPoints3DShape) {

	// rotatedsampledpoints3dshapeFrom has already been copied
	if _rotatedsampledpoints3dshapeTo, ok := mapOrigCopy[rotatedsampledpoints3dshapeFrom]; ok {
		rotatedsampledpoints3dshapeTo = _rotatedsampledpoints3dshapeTo.(*RotatedSampledPoints3DShape)
		return
	}

	rotatedsampledpoints3dshapeTo = new(RotatedSampledPoints3DShape)
	mapOrigCopy[rotatedsampledpoints3dshapeFrom] = rotatedsampledpoints3dshapeTo
	rotatedsampledpoints3dshapeFrom.CopyBasicFields(rotatedsampledpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRotatedSeatAndLegs3DShape(mapOrigCopy map[any]any, rotatedseatandlegs3dshapeFrom *RotatedSeatAndLegs3DShape) (rotatedseatandlegs3dshapeTo *RotatedSeatAndLegs3DShape) {

	// rotatedseatandlegs3dshapeFrom has already been copied
	if _rotatedseatandlegs3dshapeTo, ok := mapOrigCopy[rotatedseatandlegs3dshapeFrom]; ok {
		rotatedseatandlegs3dshapeTo = _rotatedseatandlegs3dshapeTo.(*RotatedSeatAndLegs3DShape)
		return
	}

	rotatedseatandlegs3dshapeTo = new(RotatedSeatAndLegs3DShape)
	mapOrigCopy[rotatedseatandlegs3dshapeFrom] = rotatedseatandlegs3dshapeTo
	rotatedseatandlegs3dshapeFrom.CopyBasicFields(rotatedseatandlegs3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSampledPoints3DShape(mapOrigCopy map[any]any, sampledpoints3dshapeFrom *SampledPoints3DShape) (sampledpoints3dshapeTo *SampledPoints3DShape) {

	// sampledpoints3dshapeFrom has already been copied
	if _sampledpoints3dshapeTo, ok := mapOrigCopy[sampledpoints3dshapeFrom]; ok {
		sampledpoints3dshapeTo = _sampledpoints3dshapeTo.(*SampledPoints3DShape)
		return
	}

	sampledpoints3dshapeTo = new(SampledPoints3DShape)
	mapOrigCopy[sampledpoints3dshapeFrom] = sampledpoints3dshapeTo
	sampledpoints3dshapeFrom.CopyBasicFields(sampledpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSeat3DShape(mapOrigCopy map[any]any, seat3dshapeFrom *Seat3DShape) (seat3dshapeTo *Seat3DShape) {

	// seat3dshapeFrom has already been copied
	if _seat3dshapeTo, ok := mapOrigCopy[seat3dshapeFrom]; ok {
		seat3dshapeTo = _seat3dshapeTo.(*Seat3DShape)
		return
	}

	seat3dshapeTo = new(Seat3DShape)
	mapOrigCopy[seat3dshapeFrom] = seat3dshapeTo
	seat3dshapeFrom.CopyBasicFields(seat3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSeatAndLegs3DShape(mapOrigCopy map[any]any, seatandlegs3dshapeFrom *SeatAndLegs3DShape) (seatandlegs3dshapeTo *SeatAndLegs3DShape) {

	// seatandlegs3dshapeFrom has already been copied
	if _seatandlegs3dshapeTo, ok := mapOrigCopy[seatandlegs3dshapeFrom]; ok {
		seatandlegs3dshapeTo = _seatandlegs3dshapeTo.(*SeatAndLegs3DShape)
		return
	}

	seatandlegs3dshapeTo = new(SeatAndLegs3DShape)
	mapOrigCopy[seatandlegs3dshapeFrom] = seatandlegs3dshapeTo
	seatandlegs3dshapeFrom.CopyBasicFields(seatandlegs3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSeatBottomCurveShape(mapOrigCopy map[any]any, seatbottomcurveshapeFrom *SeatBottomCurveShape) (seatbottomcurveshapeTo *SeatBottomCurveShape) {

	// seatbottomcurveshapeFrom has already been copied
	if _seatbottomcurveshapeTo, ok := mapOrigCopy[seatbottomcurveshapeFrom]; ok {
		seatbottomcurveshapeTo = _seatbottomcurveshapeTo.(*SeatBottomCurveShape)
		return
	}

	seatbottomcurveshapeTo = new(SeatBottomCurveShape)
	mapOrigCopy[seatbottomcurveshapeFrom] = seatbottomcurveshapeTo
	seatbottomcurveshapeFrom.CopyBasicFields(seatbottomcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSeatTopCurveShape(mapOrigCopy map[any]any, seattopcurveshapeFrom *SeatTopCurveShape) (seattopcurveshapeTo *SeatTopCurveShape) {

	// seattopcurveshapeFrom has already been copied
	if _seattopcurveshapeTo, ok := mapOrigCopy[seattopcurveshapeFrom]; ok {
		seattopcurveshapeTo = _seattopcurveshapeTo.(*SeatTopCurveShape)
		return
	}

	seattopcurveshapeTo = new(SeatTopCurveShape)
	mapOrigCopy[seattopcurveshapeFrom] = seattopcurveshapeTo
	seattopcurveshapeFrom.CopyBasicFields(seattopcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedBottomTopStartArcShape(mapOrigCopy map[any]any, shiftedbottomtopstartarcshapeFrom *ShiftedBottomTopStartArcShape) (shiftedbottomtopstartarcshapeTo *ShiftedBottomTopStartArcShape) {

	// shiftedbottomtopstartarcshapeFrom has already been copied
	if _shiftedbottomtopstartarcshapeTo, ok := mapOrigCopy[shiftedbottomtopstartarcshapeFrom]; ok {
		shiftedbottomtopstartarcshapeTo = _shiftedbottomtopstartarcshapeTo.(*ShiftedBottomTopStartArcShape)
		return
	}

	shiftedbottomtopstartarcshapeTo = new(ShiftedBottomTopStartArcShape)
	mapOrigCopy[shiftedbottomtopstartarcshapeFrom] = shiftedbottomtopstartarcshapeTo
	shiftedbottomtopstartarcshapeFrom.CopyBasicFields(shiftedbottomtopstartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedBottomTopStartArcShapeGrid(mapOrigCopy map[any]any, shiftedbottomtopstartarcshapegridFrom *ShiftedBottomTopStartArcShapeGrid) (shiftedbottomtopstartarcshapegridTo *ShiftedBottomTopStartArcShapeGrid) {

	// shiftedbottomtopstartarcshapegridFrom has already been copied
	if _shiftedbottomtopstartarcshapegridTo, ok := mapOrigCopy[shiftedbottomtopstartarcshapegridFrom]; ok {
		shiftedbottomtopstartarcshapegridTo = _shiftedbottomtopstartarcshapegridTo.(*ShiftedBottomTopStartArcShapeGrid)
		return
	}

	shiftedbottomtopstartarcshapegridTo = new(ShiftedBottomTopStartArcShapeGrid)
	mapOrigCopy[shiftedbottomtopstartarcshapegridFrom] = shiftedbottomtopstartarcshapegridTo
	shiftedbottomtopstartarcshapegridFrom.CopyBasicFields(shiftedbottomtopstartarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedLeftGrowthCurve2DRibbon(mapOrigCopy map[any]any, shiftedleftgrowthcurve2dribbonFrom *ShiftedLeftGrowthCurve2DRibbon) (shiftedleftgrowthcurve2dribbonTo *ShiftedLeftGrowthCurve2DRibbon) {

	// shiftedleftgrowthcurve2dribbonFrom has already been copied
	if _shiftedleftgrowthcurve2dribbonTo, ok := mapOrigCopy[shiftedleftgrowthcurve2dribbonFrom]; ok {
		shiftedleftgrowthcurve2dribbonTo = _shiftedleftgrowthcurve2dribbonTo.(*ShiftedLeftGrowthCurve2DRibbon)
		return
	}

	shiftedleftgrowthcurve2dribbonTo = new(ShiftedLeftGrowthCurve2DRibbon)
	mapOrigCopy[shiftedleftgrowthcurve2dribbonFrom] = shiftedleftgrowthcurve2dribbonTo
	shiftedleftgrowthcurve2dribbonFrom.CopyBasicFields(shiftedleftgrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedLeftGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, shiftedleftgrowthcurve2dribbonendshapeFrom *ShiftedLeftGrowthCurve2DRibbonEndShape) (shiftedleftgrowthcurve2dribbonendshapeTo *ShiftedLeftGrowthCurve2DRibbonEndShape) {

	// shiftedleftgrowthcurve2dribbonendshapeFrom has already been copied
	if _shiftedleftgrowthcurve2dribbonendshapeTo, ok := mapOrigCopy[shiftedleftgrowthcurve2dribbonendshapeFrom]; ok {
		shiftedleftgrowthcurve2dribbonendshapeTo = _shiftedleftgrowthcurve2dribbonendshapeTo.(*ShiftedLeftGrowthCurve2DRibbonEndShape)
		return
	}

	shiftedleftgrowthcurve2dribbonendshapeTo = new(ShiftedLeftGrowthCurve2DRibbonEndShape)
	mapOrigCopy[shiftedleftgrowthcurve2dribbonendshapeFrom] = shiftedleftgrowthcurve2dribbonendshapeTo
	shiftedleftgrowthcurve2dribbonendshapeFrom.CopyBasicFields(shiftedleftgrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedLeftGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, shiftedleftgrowthcurve2dribbonstartshapeFrom *ShiftedLeftGrowthCurve2DRibbonStartShape) (shiftedleftgrowthcurve2dribbonstartshapeTo *ShiftedLeftGrowthCurve2DRibbonStartShape) {

	// shiftedleftgrowthcurve2dribbonstartshapeFrom has already been copied
	if _shiftedleftgrowthcurve2dribbonstartshapeTo, ok := mapOrigCopy[shiftedleftgrowthcurve2dribbonstartshapeFrom]; ok {
		shiftedleftgrowthcurve2dribbonstartshapeTo = _shiftedleftgrowthcurve2dribbonstartshapeTo.(*ShiftedLeftGrowthCurve2DRibbonStartShape)
		return
	}

	shiftedleftgrowthcurve2dribbonstartshapeTo = new(ShiftedLeftGrowthCurve2DRibbonStartShape)
	mapOrigCopy[shiftedleftgrowthcurve2dribbonstartshapeFrom] = shiftedleftgrowthcurve2dribbonstartshapeTo
	shiftedleftgrowthcurve2dribbonstartshapeFrom.CopyBasicFields(shiftedleftgrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedLeftPartiallyGrowthCurve2DRibbon(mapOrigCopy map[any]any, shiftedleftpartiallygrowthcurve2dribbonFrom *ShiftedLeftPartiallyGrowthCurve2DRibbon) (shiftedleftpartiallygrowthcurve2dribbonTo *ShiftedLeftPartiallyGrowthCurve2DRibbon) {

	// shiftedleftpartiallygrowthcurve2dribbonFrom has already been copied
	if _shiftedleftpartiallygrowthcurve2dribbonTo, ok := mapOrigCopy[shiftedleftpartiallygrowthcurve2dribbonFrom]; ok {
		shiftedleftpartiallygrowthcurve2dribbonTo = _shiftedleftpartiallygrowthcurve2dribbonTo.(*ShiftedLeftPartiallyGrowthCurve2DRibbon)
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonTo = new(ShiftedLeftPartiallyGrowthCurve2DRibbon)
	mapOrigCopy[shiftedleftpartiallygrowthcurve2dribbonFrom] = shiftedleftpartiallygrowthcurve2dribbonTo
	shiftedleftpartiallygrowthcurve2dribbonFrom.CopyBasicFields(shiftedleftpartiallygrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, shiftedleftpartiallygrowthcurve2dribbonendshapeFrom *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) (shiftedleftpartiallygrowthcurve2dribbonendshapeTo *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) {

	// shiftedleftpartiallygrowthcurve2dribbonendshapeFrom has already been copied
	if _shiftedleftpartiallygrowthcurve2dribbonendshapeTo, ok := mapOrigCopy[shiftedleftpartiallygrowthcurve2dribbonendshapeFrom]; ok {
		shiftedleftpartiallygrowthcurve2dribbonendshapeTo = _shiftedleftpartiallygrowthcurve2dribbonendshapeTo.(*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonendshapeTo = new(ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
	mapOrigCopy[shiftedleftpartiallygrowthcurve2dribbonendshapeFrom] = shiftedleftpartiallygrowthcurve2dribbonendshapeTo
	shiftedleftpartiallygrowthcurve2dribbonendshapeFrom.CopyBasicFields(shiftedleftpartiallygrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, shiftedleftpartiallygrowthcurve2dribbonstartshapeFrom *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) (shiftedleftpartiallygrowthcurve2dribbonstartshapeTo *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) {

	// shiftedleftpartiallygrowthcurve2dribbonstartshapeFrom has already been copied
	if _shiftedleftpartiallygrowthcurve2dribbonstartshapeTo, ok := mapOrigCopy[shiftedleftpartiallygrowthcurve2dribbonstartshapeFrom]; ok {
		shiftedleftpartiallygrowthcurve2dribbonstartshapeTo = _shiftedleftpartiallygrowthcurve2dribbonstartshapeTo.(*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonstartshapeTo = new(ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
	mapOrigCopy[shiftedleftpartiallygrowthcurve2dribbonstartshapeFrom] = shiftedleftpartiallygrowthcurve2dribbonstartshapeTo
	shiftedleftpartiallygrowthcurve2dribbonstartshapeFrom.CopyBasicFields(shiftedleftpartiallygrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedLeftStackGrowthCurveEndArcShape(mapOrigCopy map[any]any, shiftedleftstackgrowthcurveendarcshapeFrom *ShiftedLeftStackGrowthCurveEndArcShape) (shiftedleftstackgrowthcurveendarcshapeTo *ShiftedLeftStackGrowthCurveEndArcShape) {

	// shiftedleftstackgrowthcurveendarcshapeFrom has already been copied
	if _shiftedleftstackgrowthcurveendarcshapeTo, ok := mapOrigCopy[shiftedleftstackgrowthcurveendarcshapeFrom]; ok {
		shiftedleftstackgrowthcurveendarcshapeTo = _shiftedleftstackgrowthcurveendarcshapeTo.(*ShiftedLeftStackGrowthCurveEndArcShape)
		return
	}

	shiftedleftstackgrowthcurveendarcshapeTo = new(ShiftedLeftStackGrowthCurveEndArcShape)
	mapOrigCopy[shiftedleftstackgrowthcurveendarcshapeFrom] = shiftedleftstackgrowthcurveendarcshapeTo
	shiftedleftstackgrowthcurveendarcshapeFrom.CopyBasicFields(shiftedleftstackgrowthcurveendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedLeftStackGrowthCurveStartArcShape(mapOrigCopy map[any]any, shiftedleftstackgrowthcurvestartarcshapeFrom *ShiftedLeftStackGrowthCurveStartArcShape) (shiftedleftstackgrowthcurvestartarcshapeTo *ShiftedLeftStackGrowthCurveStartArcShape) {

	// shiftedleftstackgrowthcurvestartarcshapeFrom has already been copied
	if _shiftedleftstackgrowthcurvestartarcshapeTo, ok := mapOrigCopy[shiftedleftstackgrowthcurvestartarcshapeFrom]; ok {
		shiftedleftstackgrowthcurvestartarcshapeTo = _shiftedleftstackgrowthcurvestartarcshapeTo.(*ShiftedLeftStackGrowthCurveStartArcShape)
		return
	}

	shiftedleftstackgrowthcurvestartarcshapeTo = new(ShiftedLeftStackGrowthCurveStartArcShape)
	mapOrigCopy[shiftedleftstackgrowthcurvestartarcshapeFrom] = shiftedleftstackgrowthcurvestartarcshapeTo
	shiftedleftstackgrowthcurvestartarcshapeFrom.CopyBasicFields(shiftedleftstackgrowthcurvestartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedLeftStackNormalVector(mapOrigCopy map[any]any, shiftedleftstacknormalvectorFrom *ShiftedLeftStackNormalVector) (shiftedleftstacknormalvectorTo *ShiftedLeftStackNormalVector) {

	// shiftedleftstacknormalvectorFrom has already been copied
	if _shiftedleftstacknormalvectorTo, ok := mapOrigCopy[shiftedleftstacknormalvectorFrom]; ok {
		shiftedleftstacknormalvectorTo = _shiftedleftstacknormalvectorTo.(*ShiftedLeftStackNormalVector)
		return
	}

	shiftedleftstacknormalvectorTo = new(ShiftedLeftStackNormalVector)
	mapOrigCopy[shiftedleftstacknormalvectorFrom] = shiftedleftstacknormalvectorTo
	shiftedleftstacknormalvectorFrom.CopyBasicFields(shiftedleftstacknormalvectorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedLeftStackOfGrowthCurve(mapOrigCopy map[any]any, shiftedleftstackofgrowthcurveFrom *ShiftedLeftStackOfGrowthCurve) (shiftedleftstackofgrowthcurveTo *ShiftedLeftStackOfGrowthCurve) {

	// shiftedleftstackofgrowthcurveFrom has already been copied
	if _shiftedleftstackofgrowthcurveTo, ok := mapOrigCopy[shiftedleftstackofgrowthcurveFrom]; ok {
		shiftedleftstackofgrowthcurveTo = _shiftedleftstackofgrowthcurveTo.(*ShiftedLeftStackOfGrowthCurve)
		return
	}

	shiftedleftstackofgrowthcurveTo = new(ShiftedLeftStackOfGrowthCurve)
	mapOrigCopy[shiftedleftstackofgrowthcurveFrom] = shiftedleftstackofgrowthcurveTo
	shiftedleftstackofgrowthcurveFrom.CopyBasicFields(shiftedleftstackofgrowthcurveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedLeftStackOfNormalVector(mapOrigCopy map[any]any, shiftedleftstackofnormalvectorFrom *ShiftedLeftStackOfNormalVector) (shiftedleftstackofnormalvectorTo *ShiftedLeftStackOfNormalVector) {

	// shiftedleftstackofnormalvectorFrom has already been copied
	if _shiftedleftstackofnormalvectorTo, ok := mapOrigCopy[shiftedleftstackofnormalvectorFrom]; ok {
		shiftedleftstackofnormalvectorTo = _shiftedleftstackofnormalvectorTo.(*ShiftedLeftStackOfNormalVector)
		return
	}

	shiftedleftstackofnormalvectorTo = new(ShiftedLeftStackOfNormalVector)
	mapOrigCopy[shiftedleftstackofnormalvectorFrom] = shiftedleftstackofnormalvectorTo
	shiftedleftstackofnormalvectorFrom.CopyBasicFields(shiftedleftstackofnormalvectorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedRightGrowthCurve2DRibbon(mapOrigCopy map[any]any, shiftedrightgrowthcurve2dribbonFrom *ShiftedRightGrowthCurve2DRibbon) (shiftedrightgrowthcurve2dribbonTo *ShiftedRightGrowthCurve2DRibbon) {

	// shiftedrightgrowthcurve2dribbonFrom has already been copied
	if _shiftedrightgrowthcurve2dribbonTo, ok := mapOrigCopy[shiftedrightgrowthcurve2dribbonFrom]; ok {
		shiftedrightgrowthcurve2dribbonTo = _shiftedrightgrowthcurve2dribbonTo.(*ShiftedRightGrowthCurve2DRibbon)
		return
	}

	shiftedrightgrowthcurve2dribbonTo = new(ShiftedRightGrowthCurve2DRibbon)
	mapOrigCopy[shiftedrightgrowthcurve2dribbonFrom] = shiftedrightgrowthcurve2dribbonTo
	shiftedrightgrowthcurve2dribbonFrom.CopyBasicFields(shiftedrightgrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedRightGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, shiftedrightgrowthcurve2dribbonendshapeFrom *ShiftedRightGrowthCurve2DRibbonEndShape) (shiftedrightgrowthcurve2dribbonendshapeTo *ShiftedRightGrowthCurve2DRibbonEndShape) {

	// shiftedrightgrowthcurve2dribbonendshapeFrom has already been copied
	if _shiftedrightgrowthcurve2dribbonendshapeTo, ok := mapOrigCopy[shiftedrightgrowthcurve2dribbonendshapeFrom]; ok {
		shiftedrightgrowthcurve2dribbonendshapeTo = _shiftedrightgrowthcurve2dribbonendshapeTo.(*ShiftedRightGrowthCurve2DRibbonEndShape)
		return
	}

	shiftedrightgrowthcurve2dribbonendshapeTo = new(ShiftedRightGrowthCurve2DRibbonEndShape)
	mapOrigCopy[shiftedrightgrowthcurve2dribbonendshapeFrom] = shiftedrightgrowthcurve2dribbonendshapeTo
	shiftedrightgrowthcurve2dribbonendshapeFrom.CopyBasicFields(shiftedrightgrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShiftedRightGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, shiftedrightgrowthcurve2dribbonstartshapeFrom *ShiftedRightGrowthCurve2DRibbonStartShape) (shiftedrightgrowthcurve2dribbonstartshapeTo *ShiftedRightGrowthCurve2DRibbonStartShape) {

	// shiftedrightgrowthcurve2dribbonstartshapeFrom has already been copied
	if _shiftedrightgrowthcurve2dribbonstartshapeTo, ok := mapOrigCopy[shiftedrightgrowthcurve2dribbonstartshapeFrom]; ok {
		shiftedrightgrowthcurve2dribbonstartshapeTo = _shiftedrightgrowthcurve2dribbonstartshapeTo.(*ShiftedRightGrowthCurve2DRibbonStartShape)
		return
	}

	shiftedrightgrowthcurve2dribbonstartshapeTo = new(ShiftedRightGrowthCurve2DRibbonStartShape)
	mapOrigCopy[shiftedrightgrowthcurve2dribbonstartshapeFrom] = shiftedrightgrowthcurve2dribbonstartshapeTo
	shiftedrightgrowthcurve2dribbonstartshapeFrom.CopyBasicFields(shiftedrightgrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackGrowthCurve2DEndHalfwayArcShape(mapOrigCopy map[any]any, stackgrowthcurve2dendhalfwayarcshapeFrom *StackGrowthCurve2DEndHalfwayArcShape) (stackgrowthcurve2dendhalfwayarcshapeTo *StackGrowthCurve2DEndHalfwayArcShape) {

	// stackgrowthcurve2dendhalfwayarcshapeFrom has already been copied
	if _stackgrowthcurve2dendhalfwayarcshapeTo, ok := mapOrigCopy[stackgrowthcurve2dendhalfwayarcshapeFrom]; ok {
		stackgrowthcurve2dendhalfwayarcshapeTo = _stackgrowthcurve2dendhalfwayarcshapeTo.(*StackGrowthCurve2DEndHalfwayArcShape)
		return
	}

	stackgrowthcurve2dendhalfwayarcshapeTo = new(StackGrowthCurve2DEndHalfwayArcShape)
	mapOrigCopy[stackgrowthcurve2dendhalfwayarcshapeFrom] = stackgrowthcurve2dendhalfwayarcshapeTo
	stackgrowthcurve2dendhalfwayarcshapeFrom.CopyBasicFields(stackgrowthcurve2dendhalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, stackgrowthcurve2dribbonendshapeFrom *StackGrowthCurve2DRibbonEndShape) (stackgrowthcurve2dribbonendshapeTo *StackGrowthCurve2DRibbonEndShape) {

	// stackgrowthcurve2dribbonendshapeFrom has already been copied
	if _stackgrowthcurve2dribbonendshapeTo, ok := mapOrigCopy[stackgrowthcurve2dribbonendshapeFrom]; ok {
		stackgrowthcurve2dribbonendshapeTo = _stackgrowthcurve2dribbonendshapeTo.(*StackGrowthCurve2DRibbonEndShape)
		return
	}

	stackgrowthcurve2dribbonendshapeTo = new(StackGrowthCurve2DRibbonEndShape)
	mapOrigCopy[stackgrowthcurve2dribbonendshapeFrom] = stackgrowthcurve2dribbonendshapeTo
	stackgrowthcurve2dribbonendshapeFrom.CopyBasicFields(stackgrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, stackgrowthcurve2dribbonstartshapeFrom *StackGrowthCurve2DRibbonStartShape) (stackgrowthcurve2dribbonstartshapeTo *StackGrowthCurve2DRibbonStartShape) {

	// stackgrowthcurve2dribbonstartshapeFrom has already been copied
	if _stackgrowthcurve2dribbonstartshapeTo, ok := mapOrigCopy[stackgrowthcurve2dribbonstartshapeFrom]; ok {
		stackgrowthcurve2dribbonstartshapeTo = _stackgrowthcurve2dribbonstartshapeTo.(*StackGrowthCurve2DRibbonStartShape)
		return
	}

	stackgrowthcurve2dribbonstartshapeTo = new(StackGrowthCurve2DRibbonStartShape)
	mapOrigCopy[stackgrowthcurve2dribbonstartshapeFrom] = stackgrowthcurve2dribbonstartshapeTo
	stackgrowthcurve2dribbonstartshapeFrom.CopyBasicFields(stackgrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackGrowthCurve2DStartHalfwayArcShape(mapOrigCopy map[any]any, stackgrowthcurve2dstarthalfwayarcshapeFrom *StackGrowthCurve2DStartHalfwayArcShape) (stackgrowthcurve2dstarthalfwayarcshapeTo *StackGrowthCurve2DStartHalfwayArcShape) {

	// stackgrowthcurve2dstarthalfwayarcshapeFrom has already been copied
	if _stackgrowthcurve2dstarthalfwayarcshapeTo, ok := mapOrigCopy[stackgrowthcurve2dstarthalfwayarcshapeFrom]; ok {
		stackgrowthcurve2dstarthalfwayarcshapeTo = _stackgrowthcurve2dstarthalfwayarcshapeTo.(*StackGrowthCurve2DStartHalfwayArcShape)
		return
	}

	stackgrowthcurve2dstarthalfwayarcshapeTo = new(StackGrowthCurve2DStartHalfwayArcShape)
	mapOrigCopy[stackgrowthcurve2dstarthalfwayarcshapeFrom] = stackgrowthcurve2dstarthalfwayarcshapeTo
	stackgrowthcurve2dstarthalfwayarcshapeFrom.CopyBasicFields(stackgrowthcurve2dstarthalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackOfGrowthCurve2D(mapOrigCopy map[any]any, stackofgrowthcurve2dFrom *StackOfGrowthCurve2D) (stackofgrowthcurve2dTo *StackOfGrowthCurve2D) {

	// stackofgrowthcurve2dFrom has already been copied
	if _stackofgrowthcurve2dTo, ok := mapOrigCopy[stackofgrowthcurve2dFrom]; ok {
		stackofgrowthcurve2dTo = _stackofgrowthcurve2dTo.(*StackOfGrowthCurve2D)
		return
	}

	stackofgrowthcurve2dTo = new(StackOfGrowthCurve2D)
	mapOrigCopy[stackofgrowthcurve2dFrom] = stackofgrowthcurve2dTo
	stackofgrowthcurve2dFrom.CopyBasicFields(stackofgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackOfGrowthCurve2DByGrowthVector(mapOrigCopy map[any]any, stackofgrowthcurve2dbygrowthvectorFrom *StackOfGrowthCurve2DByGrowthVector) (stackofgrowthcurve2dbygrowthvectorTo *StackOfGrowthCurve2DByGrowthVector) {

	// stackofgrowthcurve2dbygrowthvectorFrom has already been copied
	if _stackofgrowthcurve2dbygrowthvectorTo, ok := mapOrigCopy[stackofgrowthcurve2dbygrowthvectorFrom]; ok {
		stackofgrowthcurve2dbygrowthvectorTo = _stackofgrowthcurve2dbygrowthvectorTo.(*StackOfGrowthCurve2DByGrowthVector)
		return
	}

	stackofgrowthcurve2dbygrowthvectorTo = new(StackOfGrowthCurve2DByGrowthVector)
	mapOrigCopy[stackofgrowthcurve2dbygrowthvectorFrom] = stackofgrowthcurve2dbygrowthvectorTo
	stackofgrowthcurve2dbygrowthvectorFrom.CopyBasicFields(stackofgrowthcurve2dbygrowthvectorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackOfGrowthCurve2DRibbon(mapOrigCopy map[any]any, stackofgrowthcurve2dribbonFrom *StackOfGrowthCurve2DRibbon) (stackofgrowthcurve2dribbonTo *StackOfGrowthCurve2DRibbon) {

	// stackofgrowthcurve2dribbonFrom has already been copied
	if _stackofgrowthcurve2dribbonTo, ok := mapOrigCopy[stackofgrowthcurve2dribbonFrom]; ok {
		stackofgrowthcurve2dribbonTo = _stackofgrowthcurve2dribbonTo.(*StackOfGrowthCurve2DRibbon)
		return
	}

	stackofgrowthcurve2dribbonTo = new(StackOfGrowthCurve2DRibbon)
	mapOrigCopy[stackofgrowthcurve2dribbonFrom] = stackofgrowthcurve2dribbonTo
	stackofgrowthcurve2dribbonFrom.CopyBasicFields(stackofgrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackOfPartiallyRotatedTorusShape(mapOrigCopy map[any]any, stackofpartiallyrotatedtorusshapeFrom *StackOfPartiallyRotatedTorusShape) (stackofpartiallyrotatedtorusshapeTo *StackOfPartiallyRotatedTorusShape) {

	// stackofpartiallyrotatedtorusshapeFrom has already been copied
	if _stackofpartiallyrotatedtorusshapeTo, ok := mapOrigCopy[stackofpartiallyrotatedtorusshapeFrom]; ok {
		stackofpartiallyrotatedtorusshapeTo = _stackofpartiallyrotatedtorusshapeTo.(*StackOfPartiallyRotatedTorusShape)
		return
	}

	stackofpartiallyrotatedtorusshapeTo = new(StackOfPartiallyRotatedTorusShape)
	mapOrigCopy[stackofpartiallyrotatedtorusshapeFrom] = stackofpartiallyrotatedtorusshapeTo
	stackofpartiallyrotatedtorusshapeFrom.CopyBasicFields(stackofpartiallyrotatedtorusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackOfRotatedGrowthCurve2D(mapOrigCopy map[any]any, stackofrotatedgrowthcurve2dFrom *StackOfRotatedGrowthCurve2D) (stackofrotatedgrowthcurve2dTo *StackOfRotatedGrowthCurve2D) {

	// stackofrotatedgrowthcurve2dFrom has already been copied
	if _stackofrotatedgrowthcurve2dTo, ok := mapOrigCopy[stackofrotatedgrowthcurve2dFrom]; ok {
		stackofrotatedgrowthcurve2dTo = _stackofrotatedgrowthcurve2dTo.(*StackOfRotatedGrowthCurve2D)
		return
	}

	stackofrotatedgrowthcurve2dTo = new(StackOfRotatedGrowthCurve2D)
	mapOrigCopy[stackofrotatedgrowthcurve2dFrom] = stackofrotatedgrowthcurve2dTo
	stackofrotatedgrowthcurve2dFrom.CopyBasicFields(stackofrotatedgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackOfRotatedGrowthCurve2DRibbon(mapOrigCopy map[any]any, stackofrotatedgrowthcurve2dribbonFrom *StackOfRotatedGrowthCurve2DRibbon) (stackofrotatedgrowthcurve2dribbonTo *StackOfRotatedGrowthCurve2DRibbon) {

	// stackofrotatedgrowthcurve2dribbonFrom has already been copied
	if _stackofrotatedgrowthcurve2dribbonTo, ok := mapOrigCopy[stackofrotatedgrowthcurve2dribbonFrom]; ok {
		stackofrotatedgrowthcurve2dribbonTo = _stackofrotatedgrowthcurve2dribbonTo.(*StackOfRotatedGrowthCurve2DRibbon)
		return
	}

	stackofrotatedgrowthcurve2dribbonTo = new(StackOfRotatedGrowthCurve2DRibbon)
	mapOrigCopy[stackofrotatedgrowthcurve2dribbonFrom] = stackofrotatedgrowthcurve2dribbonTo
	stackofrotatedgrowthcurve2dribbonFrom.CopyBasicFields(stackofrotatedgrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackRotatedGrowthCurve2DEndArcShape(mapOrigCopy map[any]any, stackrotatedgrowthcurve2dendarcshapeFrom *StackRotatedGrowthCurve2DEndArcShape) (stackrotatedgrowthcurve2dendarcshapeTo *StackRotatedGrowthCurve2DEndArcShape) {

	// stackrotatedgrowthcurve2dendarcshapeFrom has already been copied
	if _stackrotatedgrowthcurve2dendarcshapeTo, ok := mapOrigCopy[stackrotatedgrowthcurve2dendarcshapeFrom]; ok {
		stackrotatedgrowthcurve2dendarcshapeTo = _stackrotatedgrowthcurve2dendarcshapeTo.(*StackRotatedGrowthCurve2DEndArcShape)
		return
	}

	stackrotatedgrowthcurve2dendarcshapeTo = new(StackRotatedGrowthCurve2DEndArcShape)
	mapOrigCopy[stackrotatedgrowthcurve2dendarcshapeFrom] = stackrotatedgrowthcurve2dendarcshapeTo
	stackrotatedgrowthcurve2dendarcshapeFrom.CopyBasicFields(stackrotatedgrowthcurve2dendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackRotatedGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, stackrotatedgrowthcurve2dribbonendshapeFrom *StackRotatedGrowthCurve2DRibbonEndShape) (stackrotatedgrowthcurve2dribbonendshapeTo *StackRotatedGrowthCurve2DRibbonEndShape) {

	// stackrotatedgrowthcurve2dribbonendshapeFrom has already been copied
	if _stackrotatedgrowthcurve2dribbonendshapeTo, ok := mapOrigCopy[stackrotatedgrowthcurve2dribbonendshapeFrom]; ok {
		stackrotatedgrowthcurve2dribbonendshapeTo = _stackrotatedgrowthcurve2dribbonendshapeTo.(*StackRotatedGrowthCurve2DRibbonEndShape)
		return
	}

	stackrotatedgrowthcurve2dribbonendshapeTo = new(StackRotatedGrowthCurve2DRibbonEndShape)
	mapOrigCopy[stackrotatedgrowthcurve2dribbonendshapeFrom] = stackrotatedgrowthcurve2dribbonendshapeTo
	stackrotatedgrowthcurve2dribbonendshapeFrom.CopyBasicFields(stackrotatedgrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackRotatedGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, stackrotatedgrowthcurve2dribbonstartshapeFrom *StackRotatedGrowthCurve2DRibbonStartShape) (stackrotatedgrowthcurve2dribbonstartshapeTo *StackRotatedGrowthCurve2DRibbonStartShape) {

	// stackrotatedgrowthcurve2dribbonstartshapeFrom has already been copied
	if _stackrotatedgrowthcurve2dribbonstartshapeTo, ok := mapOrigCopy[stackrotatedgrowthcurve2dribbonstartshapeFrom]; ok {
		stackrotatedgrowthcurve2dribbonstartshapeTo = _stackrotatedgrowthcurve2dribbonstartshapeTo.(*StackRotatedGrowthCurve2DRibbonStartShape)
		return
	}

	stackrotatedgrowthcurve2dribbonstartshapeTo = new(StackRotatedGrowthCurve2DRibbonStartShape)
	mapOrigCopy[stackrotatedgrowthcurve2dribbonstartshapeFrom] = stackrotatedgrowthcurve2dribbonstartshapeTo
	stackrotatedgrowthcurve2dribbonstartshapeFrom.CopyBasicFields(stackrotatedgrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackRotatedGrowthCurve2DStartArcShape(mapOrigCopy map[any]any, stackrotatedgrowthcurve2dstartarcshapeFrom *StackRotatedGrowthCurve2DStartArcShape) (stackrotatedgrowthcurve2dstartarcshapeTo *StackRotatedGrowthCurve2DStartArcShape) {

	// stackrotatedgrowthcurve2dstartarcshapeFrom has already been copied
	if _stackrotatedgrowthcurve2dstartarcshapeTo, ok := mapOrigCopy[stackrotatedgrowthcurve2dstartarcshapeFrom]; ok {
		stackrotatedgrowthcurve2dstartarcshapeTo = _stackrotatedgrowthcurve2dstartarcshapeTo.(*StackRotatedGrowthCurve2DStartArcShape)
		return
	}

	stackrotatedgrowthcurve2dstartarcshapeTo = new(StackRotatedGrowthCurve2DStartArcShape)
	mapOrigCopy[stackrotatedgrowthcurve2dstartarcshapeFrom] = stackrotatedgrowthcurve2dstartarcshapeTo
	stackrotatedgrowthcurve2dstartarcshapeFrom.CopyBasicFields(stackrotatedgrowthcurve2dstartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStartArcShape(mapOrigCopy map[any]any, startarcshapeFrom *StartArcShape) (startarcshapeTo *StartArcShape) {

	// startarcshapeFrom has already been copied
	if _startarcshapeTo, ok := mapOrigCopy[startarcshapeFrom]; ok {
		startarcshapeTo = _startarcshapeTo.(*StartArcShape)
		return
	}

	startarcshapeTo = new(StartArcShape)
	mapOrigCopy[startarcshapeFrom] = startarcshapeTo
	startarcshapeFrom.CopyBasicFields(startarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStartArcShapeGrid(mapOrigCopy map[any]any, startarcshapegridFrom *StartArcShapeGrid) (startarcshapegridTo *StartArcShapeGrid) {

	// startarcshapegridFrom has already been copied
	if _startarcshapegridTo, ok := mapOrigCopy[startarcshapegridFrom]; ok {
		startarcshapegridTo = _startarcshapegridTo.(*StartArcShapeGrid)
		return
	}

	startarcshapegridTo = new(StartArcShapeGrid)
	mapOrigCopy[startarcshapegridFrom] = startarcshapegridTo
	startarcshapegridFrom.CopyBasicFields(startarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStartHalfwayArcShape(mapOrigCopy map[any]any, starthalfwayarcshapeFrom *StartHalfwayArcShape) (starthalfwayarcshapeTo *StartHalfwayArcShape) {

	// starthalfwayarcshapeFrom has already been copied
	if _starthalfwayarcshapeTo, ok := mapOrigCopy[starthalfwayarcshapeFrom]; ok {
		starthalfwayarcshapeTo = _starthalfwayarcshapeTo.(*StartHalfwayArcShape)
		return
	}

	starthalfwayarcshapeTo = new(StartHalfwayArcShape)
	mapOrigCopy[starthalfwayarcshapeFrom] = starthalfwayarcshapeTo
	starthalfwayarcshapeFrom.CopyBasicFields(starthalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStartHalfwayArcShapeGrid(mapOrigCopy map[any]any, starthalfwayarcshapegridFrom *StartHalfwayArcShapeGrid) (starthalfwayarcshapegridTo *StartHalfwayArcShapeGrid) {

	// starthalfwayarcshapegridFrom has already been copied
	if _starthalfwayarcshapegridTo, ok := mapOrigCopy[starthalfwayarcshapegridFrom]; ok {
		starthalfwayarcshapegridTo = _starthalfwayarcshapegridTo.(*StartHalfwayArcShapeGrid)
		return
	}

	starthalfwayarcshapegridTo = new(StartHalfwayArcShapeGrid)
	mapOrigCopy[starthalfwayarcshapegridFrom] = starthalfwayarcshapegridTo
	starthalfwayarcshapegridFrom.CopyBasicFields(starthalfwayarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStemCylinder3DShape(mapOrigCopy map[any]any, stemcylinder3dshapeFrom *StemCylinder3DShape) (stemcylinder3dshapeTo *StemCylinder3DShape) {

	// stemcylinder3dshapeFrom has already been copied
	if _stemcylinder3dshapeTo, ok := mapOrigCopy[stemcylinder3dshapeFrom]; ok {
		stemcylinder3dshapeTo = _stemcylinder3dshapeTo.(*StemCylinder3DShape)
		return
	}

	stemcylinder3dshapeTo = new(StemCylinder3DShape)
	mapOrigCopy[stemcylinder3dshapeFrom] = stemcylinder3dshapeTo
	stemcylinder3dshapeFrom.CopyBasicFields(stemcylinder3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStool2DDiagram(mapOrigCopy map[any]any, stool2ddiagramFrom *Stool2DDiagram) (stool2ddiagramTo *Stool2DDiagram) {

	// stool2ddiagramFrom has already been copied
	if _stool2ddiagramTo, ok := mapOrigCopy[stool2ddiagramFrom]; ok {
		stool2ddiagramTo = _stool2ddiagramTo.(*Stool2DDiagram)
		return
	}

	stool2ddiagramTo = new(Stool2DDiagram)
	mapOrigCopy[stool2ddiagramFrom] = stool2ddiagramTo
	stool2ddiagramFrom.CopyBasicFields(stool2ddiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStool3DDiagram(mapOrigCopy map[any]any, stool3ddiagramFrom *Stool3DDiagram) (stool3ddiagramTo *Stool3DDiagram) {

	// stool3ddiagramFrom has already been copied
	if _stool3ddiagramTo, ok := mapOrigCopy[stool3ddiagramFrom]; ok {
		stool3ddiagramTo = _stool3ddiagramTo.(*Stool3DDiagram)
		return
	}

	stool3ddiagramTo = new(Stool3DDiagram)
	mapOrigCopy[stool3ddiagramFrom] = stool3ddiagramTo
	stool3ddiagramFrom.CopyBasicFields(stool3ddiagramTo)

	//insertion point for the staging of instances referenced by pointers
	if stool3ddiagramFrom.SampledPoints3DShape != nil {
		stool3ddiagramTo.SampledPoints3DShape = CopyBranchSampledPoints3DShape(mapOrigCopy, stool3ddiagramFrom.SampledPoints3DShape)
	}
	if stool3ddiagramFrom.Rendered3DShape != nil {
		stool3ddiagramTo.Rendered3DShape = CopyBranchRendered3DShape(mapOrigCopy, stool3ddiagramFrom.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStoolAbstract(mapOrigCopy map[any]any, stoolabstractFrom *StoolAbstract) (stoolabstractTo *StoolAbstract) {

	// stoolabstractFrom has already been copied
	if _stoolabstractTo, ok := mapOrigCopy[stoolabstractFrom]; ok {
		stoolabstractTo = _stoolabstractTo.(*StoolAbstract)
		return
	}

	stoolabstractTo = new(StoolAbstract)
	mapOrigCopy[stoolabstractFrom] = stoolabstractTo
	stoolabstractFrom.CopyBasicFields(stoolabstractTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTiledFloor3DShape(mapOrigCopy map[any]any, tiledfloor3dshapeFrom *TiledFloor3DShape) (tiledfloor3dshapeTo *TiledFloor3DShape) {

	// tiledfloor3dshapeFrom has already been copied
	if _tiledfloor3dshapeTo, ok := mapOrigCopy[tiledfloor3dshapeFrom]; ok {
		tiledfloor3dshapeTo = _tiledfloor3dshapeTo.(*TiledFloor3DShape)
		return
	}

	tiledfloor3dshapeTo = new(TiledFloor3DShape)
	mapOrigCopy[tiledfloor3dshapeFrom] = tiledfloor3dshapeTo
	tiledfloor3dshapeFrom.CopyBasicFields(tiledfloor3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopEndArcShape(mapOrigCopy map[any]any, topendarcshapeFrom *TopEndArcShape) (topendarcshapeTo *TopEndArcShape) {

	// topendarcshapeFrom has already been copied
	if _topendarcshapeTo, ok := mapOrigCopy[topendarcshapeFrom]; ok {
		topendarcshapeTo = _topendarcshapeTo.(*TopEndArcShape)
		return
	}

	topendarcshapeTo = new(TopEndArcShape)
	mapOrigCopy[topendarcshapeFrom] = topendarcshapeTo
	topendarcshapeFrom.CopyBasicFields(topendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopEndArcShapeGrid(mapOrigCopy map[any]any, topendarcshapegridFrom *TopEndArcShapeGrid) (topendarcshapegridTo *TopEndArcShapeGrid) {

	// topendarcshapegridFrom has already been copied
	if _topendarcshapegridTo, ok := mapOrigCopy[topendarcshapegridFrom]; ok {
		topendarcshapegridTo = _topendarcshapegridTo.(*TopEndArcShapeGrid)
		return
	}

	topendarcshapegridTo = new(TopEndArcShapeGrid)
	mapOrigCopy[topendarcshapegridFrom] = topendarcshapegridTo
	topendarcshapegridFrom.CopyBasicFields(topendarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopEndHalfwayArcShape(mapOrigCopy map[any]any, topendhalfwayarcshapeFrom *TopEndHalfwayArcShape) (topendhalfwayarcshapeTo *TopEndHalfwayArcShape) {

	// topendhalfwayarcshapeFrom has already been copied
	if _topendhalfwayarcshapeTo, ok := mapOrigCopy[topendhalfwayarcshapeFrom]; ok {
		topendhalfwayarcshapeTo = _topendhalfwayarcshapeTo.(*TopEndHalfwayArcShape)
		return
	}

	topendhalfwayarcshapeTo = new(TopEndHalfwayArcShape)
	mapOrigCopy[topendhalfwayarcshapeFrom] = topendhalfwayarcshapeTo
	topendhalfwayarcshapeFrom.CopyBasicFields(topendhalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopEndHalfwayArcShapeGrid(mapOrigCopy map[any]any, topendhalfwayarcshapegridFrom *TopEndHalfwayArcShapeGrid) (topendhalfwayarcshapegridTo *TopEndHalfwayArcShapeGrid) {

	// topendhalfwayarcshapegridFrom has already been copied
	if _topendhalfwayarcshapegridTo, ok := mapOrigCopy[topendhalfwayarcshapegridFrom]; ok {
		topendhalfwayarcshapegridTo = _topendhalfwayarcshapegridTo.(*TopEndHalfwayArcShapeGrid)
		return
	}

	topendhalfwayarcshapegridTo = new(TopEndHalfwayArcShapeGrid)
	mapOrigCopy[topendhalfwayarcshapegridFrom] = topendhalfwayarcshapegridTo
	topendhalfwayarcshapegridFrom.CopyBasicFields(topendhalfwayarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopGrowthCurve2D(mapOrigCopy map[any]any, topgrowthcurve2dFrom *TopGrowthCurve2D) (topgrowthcurve2dTo *TopGrowthCurve2D) {

	// topgrowthcurve2dFrom has already been copied
	if _topgrowthcurve2dTo, ok := mapOrigCopy[topgrowthcurve2dFrom]; ok {
		topgrowthcurve2dTo = _topgrowthcurve2dTo.(*TopGrowthCurve2D)
		return
	}

	topgrowthcurve2dTo = new(TopGrowthCurve2D)
	mapOrigCopy[topgrowthcurve2dFrom] = topgrowthcurve2dTo
	topgrowthcurve2dFrom.CopyBasicFields(topgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopMidArcVectorShape(mapOrigCopy map[any]any, topmidarcvectorshapeFrom *TopMidArcVectorShape) (topmidarcvectorshapeTo *TopMidArcVectorShape) {

	// topmidarcvectorshapeFrom has already been copied
	if _topmidarcvectorshapeTo, ok := mapOrigCopy[topmidarcvectorshapeFrom]; ok {
		topmidarcvectorshapeTo = _topmidarcvectorshapeTo.(*TopMidArcVectorShape)
		return
	}

	topmidarcvectorshapeTo = new(TopMidArcVectorShape)
	mapOrigCopy[topmidarcvectorshapeFrom] = topmidarcvectorshapeTo
	topmidarcvectorshapeFrom.CopyBasicFields(topmidarcvectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopMidArcVectorShapeGrid(mapOrigCopy map[any]any, topmidarcvectorshapegridFrom *TopMidArcVectorShapeGrid) (topmidarcvectorshapegridTo *TopMidArcVectorShapeGrid) {

	// topmidarcvectorshapegridFrom has already been copied
	if _topmidarcvectorshapegridTo, ok := mapOrigCopy[topmidarcvectorshapegridFrom]; ok {
		topmidarcvectorshapegridTo = _topmidarcvectorshapegridTo.(*TopMidArcVectorShapeGrid)
		return
	}

	topmidarcvectorshapegridTo = new(TopMidArcVectorShapeGrid)
	mapOrigCopy[topmidarcvectorshapegridFrom] = topmidarcvectorshapegridTo
	topmidarcvectorshapegridFrom.CopyBasicFields(topmidarcvectorshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStackGrowthCurve2DEndHalfwayArcShape(mapOrigCopy map[any]any, topstackgrowthcurve2dendhalfwayarcshapeFrom *TopStackGrowthCurve2DEndHalfwayArcShape) (topstackgrowthcurve2dendhalfwayarcshapeTo *TopStackGrowthCurve2DEndHalfwayArcShape) {

	// topstackgrowthcurve2dendhalfwayarcshapeFrom has already been copied
	if _topstackgrowthcurve2dendhalfwayarcshapeTo, ok := mapOrigCopy[topstackgrowthcurve2dendhalfwayarcshapeFrom]; ok {
		topstackgrowthcurve2dendhalfwayarcshapeTo = _topstackgrowthcurve2dendhalfwayarcshapeTo.(*TopStackGrowthCurve2DEndHalfwayArcShape)
		return
	}

	topstackgrowthcurve2dendhalfwayarcshapeTo = new(TopStackGrowthCurve2DEndHalfwayArcShape)
	mapOrigCopy[topstackgrowthcurve2dendhalfwayarcshapeFrom] = topstackgrowthcurve2dendhalfwayarcshapeTo
	topstackgrowthcurve2dendhalfwayarcshapeFrom.CopyBasicFields(topstackgrowthcurve2dendhalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStackGrowthCurve2DStartHalfwayArcShape(mapOrigCopy map[any]any, topstackgrowthcurve2dstarthalfwayarcshapeFrom *TopStackGrowthCurve2DStartHalfwayArcShape) (topstackgrowthcurve2dstarthalfwayarcshapeTo *TopStackGrowthCurve2DStartHalfwayArcShape) {

	// topstackgrowthcurve2dstarthalfwayarcshapeFrom has already been copied
	if _topstackgrowthcurve2dstarthalfwayarcshapeTo, ok := mapOrigCopy[topstackgrowthcurve2dstarthalfwayarcshapeFrom]; ok {
		topstackgrowthcurve2dstarthalfwayarcshapeTo = _topstackgrowthcurve2dstarthalfwayarcshapeTo.(*TopStackGrowthCurve2DStartHalfwayArcShape)
		return
	}

	topstackgrowthcurve2dstarthalfwayarcshapeTo = new(TopStackGrowthCurve2DStartHalfwayArcShape)
	mapOrigCopy[topstackgrowthcurve2dstarthalfwayarcshapeFrom] = topstackgrowthcurve2dstarthalfwayarcshapeTo
	topstackgrowthcurve2dstarthalfwayarcshapeFrom.CopyBasicFields(topstackgrowthcurve2dstarthalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStackOfGrowthCurve2D(mapOrigCopy map[any]any, topstackofgrowthcurve2dFrom *TopStackOfGrowthCurve2D) (topstackofgrowthcurve2dTo *TopStackOfGrowthCurve2D) {

	// topstackofgrowthcurve2dFrom has already been copied
	if _topstackofgrowthcurve2dTo, ok := mapOrigCopy[topstackofgrowthcurve2dFrom]; ok {
		topstackofgrowthcurve2dTo = _topstackofgrowthcurve2dTo.(*TopStackOfGrowthCurve2D)
		return
	}

	topstackofgrowthcurve2dTo = new(TopStackOfGrowthCurve2D)
	mapOrigCopy[topstackofgrowthcurve2dFrom] = topstackofgrowthcurve2dTo
	topstackofgrowthcurve2dFrom.CopyBasicFields(topstackofgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStackOfRotatedGrowthCurve2D(mapOrigCopy map[any]any, topstackofrotatedgrowthcurve2dFrom *TopStackOfRotatedGrowthCurve2D) (topstackofrotatedgrowthcurve2dTo *TopStackOfRotatedGrowthCurve2D) {

	// topstackofrotatedgrowthcurve2dFrom has already been copied
	if _topstackofrotatedgrowthcurve2dTo, ok := mapOrigCopy[topstackofrotatedgrowthcurve2dFrom]; ok {
		topstackofrotatedgrowthcurve2dTo = _topstackofrotatedgrowthcurve2dTo.(*TopStackOfRotatedGrowthCurve2D)
		return
	}

	topstackofrotatedgrowthcurve2dTo = new(TopStackOfRotatedGrowthCurve2D)
	mapOrigCopy[topstackofrotatedgrowthcurve2dFrom] = topstackofrotatedgrowthcurve2dTo
	topstackofrotatedgrowthcurve2dFrom.CopyBasicFields(topstackofrotatedgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStackOfRotatedGrowthCurve2DEndArcShape(mapOrigCopy map[any]any, topstackofrotatedgrowthcurve2dendarcshapeFrom *TopStackOfRotatedGrowthCurve2DEndArcShape) (topstackofrotatedgrowthcurve2dendarcshapeTo *TopStackOfRotatedGrowthCurve2DEndArcShape) {

	// topstackofrotatedgrowthcurve2dendarcshapeFrom has already been copied
	if _topstackofrotatedgrowthcurve2dendarcshapeTo, ok := mapOrigCopy[topstackofrotatedgrowthcurve2dendarcshapeFrom]; ok {
		topstackofrotatedgrowthcurve2dendarcshapeTo = _topstackofrotatedgrowthcurve2dendarcshapeTo.(*TopStackOfRotatedGrowthCurve2DEndArcShape)
		return
	}

	topstackofrotatedgrowthcurve2dendarcshapeTo = new(TopStackOfRotatedGrowthCurve2DEndArcShape)
	mapOrigCopy[topstackofrotatedgrowthcurve2dendarcshapeFrom] = topstackofrotatedgrowthcurve2dendarcshapeTo
	topstackofrotatedgrowthcurve2dendarcshapeFrom.CopyBasicFields(topstackofrotatedgrowthcurve2dendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStackOfRotatedGrowthCurve2DStartArcShape(mapOrigCopy map[any]any, topstackofrotatedgrowthcurve2dstartarcshapeFrom *TopStackOfRotatedGrowthCurve2DStartArcShape) (topstackofrotatedgrowthcurve2dstartarcshapeTo *TopStackOfRotatedGrowthCurve2DStartArcShape) {

	// topstackofrotatedgrowthcurve2dstartarcshapeFrom has already been copied
	if _topstackofrotatedgrowthcurve2dstartarcshapeTo, ok := mapOrigCopy[topstackofrotatedgrowthcurve2dstartarcshapeFrom]; ok {
		topstackofrotatedgrowthcurve2dstartarcshapeTo = _topstackofrotatedgrowthcurve2dstartarcshapeTo.(*TopStackOfRotatedGrowthCurve2DStartArcShape)
		return
	}

	topstackofrotatedgrowthcurve2dstartarcshapeTo = new(TopStackOfRotatedGrowthCurve2DStartArcShape)
	mapOrigCopy[topstackofrotatedgrowthcurve2dstartarcshapeFrom] = topstackofrotatedgrowthcurve2dstartarcshapeTo
	topstackofrotatedgrowthcurve2dstartarcshapeFrom.CopyBasicFields(topstackofrotatedgrowthcurve2dstartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStartArcShape(mapOrigCopy map[any]any, topstartarcshapeFrom *TopStartArcShape) (topstartarcshapeTo *TopStartArcShape) {

	// topstartarcshapeFrom has already been copied
	if _topstartarcshapeTo, ok := mapOrigCopy[topstartarcshapeFrom]; ok {
		topstartarcshapeTo = _topstartarcshapeTo.(*TopStartArcShape)
		return
	}

	topstartarcshapeTo = new(TopStartArcShape)
	mapOrigCopy[topstartarcshapeFrom] = topstartarcshapeTo
	topstartarcshapeFrom.CopyBasicFields(topstartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStartArcShapeGrid(mapOrigCopy map[any]any, topstartarcshapegridFrom *TopStartArcShapeGrid) (topstartarcshapegridTo *TopStartArcShapeGrid) {

	// topstartarcshapegridFrom has already been copied
	if _topstartarcshapegridTo, ok := mapOrigCopy[topstartarcshapegridFrom]; ok {
		topstartarcshapegridTo = _topstartarcshapegridTo.(*TopStartArcShapeGrid)
		return
	}

	topstartarcshapegridTo = new(TopStartArcShapeGrid)
	mapOrigCopy[topstartarcshapegridFrom] = topstartarcshapegridTo
	topstartarcshapegridFrom.CopyBasicFields(topstartarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStartHalfwayArcShape(mapOrigCopy map[any]any, topstarthalfwayarcshapeFrom *TopStartHalfwayArcShape) (topstarthalfwayarcshapeTo *TopStartHalfwayArcShape) {

	// topstarthalfwayarcshapeFrom has already been copied
	if _topstarthalfwayarcshapeTo, ok := mapOrigCopy[topstarthalfwayarcshapeFrom]; ok {
		topstarthalfwayarcshapeTo = _topstarthalfwayarcshapeTo.(*TopStartHalfwayArcShape)
		return
	}

	topstarthalfwayarcshapeTo = new(TopStartHalfwayArcShape)
	mapOrigCopy[topstarthalfwayarcshapeFrom] = topstarthalfwayarcshapeTo
	topstarthalfwayarcshapeFrom.CopyBasicFields(topstarthalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStartHalfwayArcShapeGrid(mapOrigCopy map[any]any, topstarthalfwayarcshapegridFrom *TopStartHalfwayArcShapeGrid) (topstarthalfwayarcshapegridTo *TopStartHalfwayArcShapeGrid) {

	// topstarthalfwayarcshapegridFrom has already been copied
	if _topstarthalfwayarcshapegridTo, ok := mapOrigCopy[topstarthalfwayarcshapegridFrom]; ok {
		topstarthalfwayarcshapegridTo = _topstarthalfwayarcshapegridTo.(*TopStartHalfwayArcShapeGrid)
		return
	}

	topstarthalfwayarcshapegridTo = new(TopStartHalfwayArcShapeGrid)
	mapOrigCopy[topstarthalfwayarcshapegridFrom] = topstarthalfwayarcshapegridTo
	topstarthalfwayarcshapegridFrom.CopyBasicFields(topstarthalfwayarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTorus3DShape(mapOrigCopy map[any]any, torus3dshapeFrom *Torus3DShape) (torus3dshapeTo *Torus3DShape) {

	// torus3dshapeFrom has already been copied
	if _torus3dshapeTo, ok := mapOrigCopy[torus3dshapeFrom]; ok {
		torus3dshapeTo = _torus3dshapeTo.(*Torus3DShape)
		return
	}

	torus3dshapeTo = new(Torus3DShape)
	mapOrigCopy[torus3dshapeFrom] = torus3dshapeTo
	torus3dshapeFrom.CopyBasicFields(torus3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTorusEdge3DShape(mapOrigCopy map[any]any, torusedge3dshapeFrom *TorusEdge3DShape) (torusedge3dshapeTo *TorusEdge3DShape) {

	// torusedge3dshapeFrom has already been copied
	if _torusedge3dshapeTo, ok := mapOrigCopy[torusedge3dshapeFrom]; ok {
		torusedge3dshapeTo = _torusedge3dshapeTo.(*TorusEdge3DShape)
		return
	}

	torusedge3dshapeTo = new(TorusEdge3DShape)
	mapOrigCopy[torusedge3dshapeFrom] = torusedge3dshapeTo
	torusedge3dshapeFrom.CopyBasicFields(torusedge3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTorusStackShape(mapOrigCopy map[any]any, torusstackshapeFrom *TorusStackShape) (torusstackshapeTo *TorusStackShape) {

	// torusstackshapeFrom has already been copied
	if _torusstackshapeTo, ok := mapOrigCopy[torusstackshapeFrom]; ok {
		torusstackshapeTo = _torusstackshapeTo.(*TorusStackShape)
		return
	}

	torusstackshapeTo = new(TorusStackShape)
	mapOrigCopy[torusstackshapeFrom] = torusstackshapeTo
	torusstackshapeFrom.CopyBasicFields(torusstackshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchVase2DDiagram(mapOrigCopy map[any]any, vase2ddiagramFrom *Vase2DDiagram) (vase2ddiagramTo *Vase2DDiagram) {

	// vase2ddiagramFrom has already been copied
	if _vase2ddiagramTo, ok := mapOrigCopy[vase2ddiagramFrom]; ok {
		vase2ddiagramTo = _vase2ddiagramTo.(*Vase2DDiagram)
		return
	}

	vase2ddiagramTo = new(Vase2DDiagram)
	mapOrigCopy[vase2ddiagramFrom] = vase2ddiagramTo
	vase2ddiagramFrom.CopyBasicFields(vase2ddiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchVase3DDiagram(mapOrigCopy map[any]any, vase3ddiagramFrom *Vase3DDiagram) (vase3ddiagramTo *Vase3DDiagram) {

	// vase3ddiagramFrom has already been copied
	if _vase3ddiagramTo, ok := mapOrigCopy[vase3ddiagramFrom]; ok {
		vase3ddiagramTo = _vase3ddiagramTo.(*Vase3DDiagram)
		return
	}

	vase3ddiagramTo = new(Vase3DDiagram)
	mapOrigCopy[vase3ddiagramFrom] = vase3ddiagramTo
	vase3ddiagramFrom.CopyBasicFields(vase3ddiagramTo)

	//insertion point for the staging of instances referenced by pointers
	if vase3ddiagramFrom.Rendered3DShape != nil {
		vase3ddiagramTo.Rendered3DShape = CopyBranchRendered3DShape(mapOrigCopy, vase3ddiagramFrom.Rendered3DShape)
	}
	if vase3ddiagramFrom.SampledPoints3DShape != nil {
		vase3ddiagramTo.SampledPoints3DShape = CopyBranchSampledPoints3DShape(mapOrigCopy, vase3ddiagramFrom.SampledPoints3DShape)
	}
	if vase3ddiagramFrom.OriginalPoints3DShape != nil {
		vase3ddiagramTo.OriginalPoints3DShape = CopyBranchOriginalPoints3DShape(mapOrigCopy, vase3ddiagramFrom.OriginalPoints3DShape)
	}
	if vase3ddiagramFrom.Angle0Shape != nil {
		vase3ddiagramTo.Angle0Shape = CopyBranchAngle0Shape(mapOrigCopy, vase3ddiagramFrom.Angle0Shape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchVaseAbstract(mapOrigCopy map[any]any, vaseabstractFrom *VaseAbstract) (vaseabstractTo *VaseAbstract) {

	// vaseabstractFrom has already been copied
	if _vaseabstractTo, ok := mapOrigCopy[vaseabstractFrom]; ok {
		vaseabstractTo = _vaseabstractTo.(*VaseAbstract)
		return
	}

	vaseabstractTo = new(VaseAbstract)
	mapOrigCopy[vaseabstractFrom] = vaseabstractTo
	vaseabstractFrom.CopyBasicFields(vaseabstractTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchVerticalTorusStackShape(mapOrigCopy map[any]any, verticaltorusstackshapeFrom *VerticalTorusStackShape) (verticaltorusstackshapeTo *VerticalTorusStackShape) {

	// verticaltorusstackshapeFrom has already been copied
	if _verticaltorusstackshapeTo, ok := mapOrigCopy[verticaltorusstackshapeFrom]; ok {
		verticaltorusstackshapeTo = _verticaltorusstackshapeTo.(*VerticalTorusStackShape)
		return
	}

	verticaltorusstackshapeTo = new(VerticalTorusStackShape)
	mapOrigCopy[verticaltorusstackshapeFrom] = verticaltorusstackshapeTo
	verticaltorusstackshapeFrom.CopyBasicFields(verticaltorusstackshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchVolumeKey3DShape(mapOrigCopy map[any]any, volumekey3dshapeFrom *VolumeKey3DShape) (volumekey3dshapeTo *VolumeKey3DShape) {

	// volumekey3dshapeFrom has already been copied
	if _volumekey3dshapeTo, ok := mapOrigCopy[volumekey3dshapeFrom]; ok {
		volumekey3dshapeTo = _volumekey3dshapeTo.(*VolumeKey3DShape)
		return
	}

	volumekey3dshapeTo = new(VolumeKey3DShape)
	mapOrigCopy[volumekey3dshapeFrom] = volumekey3dshapeTo
	volumekey3dshapeFrom.CopyBasicFields(volumekey3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Angle0Shape:
		stage.UnstageBranchAngle0Shape(target)

	case *ArcNormalVectorShape:
		stage.UnstageBranchArcNormalVectorShape(target)

	case *ArcNormalVectorShapeGrid:
		stage.UnstageBranchArcNormalVectorShapeGrid(target)

	case *AxesShape:
		stage.UnstageBranchAxesShape(target)

	case *BaseVectorShape:
		stage.UnstageBranchBaseVectorShape(target)

	case *BaseVectorShapeGrid:
		stage.UnstageBranchBaseVectorShapeGrid(target)

	case *ChosenP1P2PairShape:
		stage.UnstageBranchChosenP1P2PairShape(target)

	case *CircleGridShape:
		stage.UnstageBranchCircleGridShape(target)

	case *Circumference3DShape:
		stage.UnstageBranchCircumference3DShape(target)

	case *Clock2DDiagram:
		stage.UnstageBranchClock2DDiagram(target)

	case *Clock3DDiagram:
		stage.UnstageBranchClock3DDiagram(target)

	case *ClockAbstract:
		stage.UnstageBranchClockAbstract(target)

	case *ClockTopCurveShape:
		stage.UnstageBranchClockTopCurveShape(target)

	case *CutLine3DShape:
		stage.UnstageBranchCutLine3DShape(target)

	case *EndArcShape:
		stage.UnstageBranchEndArcShape(target)

	case *EndArcShapeGrid:
		stage.UnstageBranchEndArcShapeGrid(target)

	case *EndHalfwayArcShape:
		stage.UnstageBranchEndHalfwayArcShape(target)

	case *EndHalfwayArcShapeGrid:
		stage.UnstageBranchEndHalfwayArcShapeGrid(target)

	case *ExplanationTextShape:
		stage.UnstageBranchExplanationTextShape(target)

	case *Eye3DShape:
		stage.UnstageBranchEye3DShape(target)

	case *EyeCornersSampledPoints3DShape:
		stage.UnstageBranchEyeCornersSampledPoints3DShape(target)

	case *EyeSampledPoints3DShape:
		stage.UnstageBranchEyeSampledPoints3DShape(target)

	case *EyeSeatBottomCurveShape:
		stage.UnstageBranchEyeSeatBottomCurveShape(target)

	case *EyeStoolBottomCurveShape:
		stage.UnstageBranchEyeStoolBottomCurveShape(target)

	case *EyeVolume3DShape:
		stage.UnstageBranchEyeVolume3DShape(target)

	case *GridPathShape:
		stage.UnstageBranchGridPathShape(target)

	case *GrowthCurve2D:
		stage.UnstageBranchGrowthCurve2D(target)

	case *GrowthCurve2DRibbon:
		stage.UnstageBranchGrowthCurve2DRibbon(target)

	case *GrowthCurve2DRibbonEndShape:
		stage.UnstageBranchGrowthCurve2DRibbonEndShape(target)

	case *GrowthCurve2DRibbonStartShape:
		stage.UnstageBranchGrowthCurve2DRibbonStartShape(target)

	case *GrowthCurveRhombusGridShape:
		stage.UnstageBranchGrowthCurveRhombusGridShape(target)

	case *GrowthCurveRhombusShape:
		stage.UnstageBranchGrowthCurveRhombusShape(target)

	case *GrowthVectorShape:
		stage.UnstageBranchGrowthVectorShape(target)

	case *InitialRhombusGridShape:
		stage.UnstageBranchInitialRhombusGridShape(target)

	case *InitialRhombusShape:
		stage.UnstageBranchInitialRhombusShape(target)

	case *Key3DShape:
		stage.UnstageBranchKey3DShape(target)

	case *KeyHole3DShape:
		stage.UnstageBranchKeyHole3DShape(target)

	case *KeyHoleShape:
		stage.UnstageBranchKeyHoleShape(target)

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *MidArcVectorShape:
		stage.UnstageBranchMidArcVectorShape(target)

	case *MidArcVectorShapeGrid:
		stage.UnstageBranchMidArcVectorShapeGrid(target)

	case *OriginalPoints3DShape:
		stage.UnstageBranchOriginalPoints3DShape(target)

	case *ParastichyMCurves3DShape:
		stage.UnstageBranchParastichyMCurves3DShape(target)

	case *ParastichyNCurves3DShape:
		stage.UnstageBranchParastichyNCurves3DShape(target)

	case *PartiallyGrowthCurve2DRibbon:
		stage.UnstageBranchPartiallyGrowthCurve2DRibbon(target)

	case *PartiallyGrowthCurve2DRibbonEndShape:
		stage.UnstageBranchPartiallyGrowthCurve2DRibbonEndShape(target)

	case *PartiallyGrowthCurve2DRibbonStartShape:
		stage.UnstageBranchPartiallyGrowthCurve2DRibbonStartShape(target)

	case *PartiallyGrowthCurve2DTrajectory:
		stage.UnstageBranchPartiallyGrowthCurve2DTrajectory(target)

	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryP1CurveShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryP1P2(target)

	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryP1PointShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryP2CurveShape(target)

	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryP2PointShape(target)

	case *PartiallyGrowthCurve2DTrajectoryShape:
		stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryShape(target)

	case *PartiallyRotatedSeatBottomCurveShape:
		stage.UnstageBranchPartiallyRotatedSeatBottomCurveShape(target)

	case *PartiallyRotatedSeatTopCurveShape:
		stage.UnstageBranchPartiallyRotatedSeatTopCurveShape(target)

	case *PartiallyRotatedTorusShape:
		stage.UnstageBranchPartiallyRotatedTorusShape(target)

	case *PerpendicularVector:
		stage.UnstageBranchPerpendicularVector(target)

	case *PerpendicularVectorGrid:
		stage.UnstageBranchPerpendicularVectorGrid(target)

	case *PerpendicularVectorGridHalfway:
		stage.UnstageBranchPerpendicularVectorGridHalfway(target)

	case *PerpendicularVectorHalfway:
		stage.UnstageBranchPerpendicularVectorHalfway(target)

	case *Plant2DDiagram:
		stage.UnstageBranchPlant2DDiagram(target)

	case *Plant3DDiagram:
		stage.UnstageBranchPlant3DDiagram(target)

	case *PlantAbstract:
		stage.UnstageBranchPlantAbstract(target)

	case *PlantCircumferenceShape:
		stage.UnstageBranchPlantCircumferenceShape(target)

	case *PointsAndLines3DShape:
		stage.UnstageBranchPointsAndLines3DShape(target)

	case *PxShape:
		stage.UnstageBranchPxShape(target)

	case *Rendered3DShape:
		stage.UnstageBranchRendered3DShape(target)

	case *RhombusShape:
		stage.UnstageBranchRhombusShape(target)

	case *RhombusStuff:
		stage.UnstageBranchRhombusStuff(target)

	case *RotatedRhombusGridShape:
		stage.UnstageBranchRotatedRhombusGridShape(target)

	case *RotatedRhombusShape:
		stage.UnstageBranchRotatedRhombusShape(target)

	case *RotatedSampledPoints3DShape:
		stage.UnstageBranchRotatedSampledPoints3DShape(target)

	case *RotatedSeatAndLegs3DShape:
		stage.UnstageBranchRotatedSeatAndLegs3DShape(target)

	case *SampledPoints3DShape:
		stage.UnstageBranchSampledPoints3DShape(target)

	case *Seat3DShape:
		stage.UnstageBranchSeat3DShape(target)

	case *SeatAndLegs3DShape:
		stage.UnstageBranchSeatAndLegs3DShape(target)

	case *SeatBottomCurveShape:
		stage.UnstageBranchSeatBottomCurveShape(target)

	case *SeatTopCurveShape:
		stage.UnstageBranchSeatTopCurveShape(target)

	case *ShiftedBottomTopStartArcShape:
		stage.UnstageBranchShiftedBottomTopStartArcShape(target)

	case *ShiftedBottomTopStartArcShapeGrid:
		stage.UnstageBranchShiftedBottomTopStartArcShapeGrid(target)

	case *ShiftedLeftGrowthCurve2DRibbon:
		stage.UnstageBranchShiftedLeftGrowthCurve2DRibbon(target)

	case *ShiftedLeftGrowthCurve2DRibbonEndShape:
		stage.UnstageBranchShiftedLeftGrowthCurve2DRibbonEndShape(target)

	case *ShiftedLeftGrowthCurve2DRibbonStartShape:
		stage.UnstageBranchShiftedLeftGrowthCurve2DRibbonStartShape(target)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbon:
		stage.UnstageBranchShiftedLeftPartiallyGrowthCurve2DRibbon(target)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape:
		stage.UnstageBranchShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(target)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape:
		stage.UnstageBranchShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(target)

	case *ShiftedLeftStackGrowthCurveEndArcShape:
		stage.UnstageBranchShiftedLeftStackGrowthCurveEndArcShape(target)

	case *ShiftedLeftStackGrowthCurveStartArcShape:
		stage.UnstageBranchShiftedLeftStackGrowthCurveStartArcShape(target)

	case *ShiftedLeftStackNormalVector:
		stage.UnstageBranchShiftedLeftStackNormalVector(target)

	case *ShiftedLeftStackOfGrowthCurve:
		stage.UnstageBranchShiftedLeftStackOfGrowthCurve(target)

	case *ShiftedLeftStackOfNormalVector:
		stage.UnstageBranchShiftedLeftStackOfNormalVector(target)

	case *ShiftedRightGrowthCurve2DRibbon:
		stage.UnstageBranchShiftedRightGrowthCurve2DRibbon(target)

	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		stage.UnstageBranchShiftedRightGrowthCurve2DRibbonEndShape(target)

	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		stage.UnstageBranchShiftedRightGrowthCurve2DRibbonStartShape(target)

	case *StackGrowthCurve2DEndHalfwayArcShape:
		stage.UnstageBranchStackGrowthCurve2DEndHalfwayArcShape(target)

	case *StackGrowthCurve2DRibbonEndShape:
		stage.UnstageBranchStackGrowthCurve2DRibbonEndShape(target)

	case *StackGrowthCurve2DRibbonStartShape:
		stage.UnstageBranchStackGrowthCurve2DRibbonStartShape(target)

	case *StackGrowthCurve2DStartHalfwayArcShape:
		stage.UnstageBranchStackGrowthCurve2DStartHalfwayArcShape(target)

	case *StackOfGrowthCurve2D:
		stage.UnstageBranchStackOfGrowthCurve2D(target)

	case *StackOfGrowthCurve2DByGrowthVector:
		stage.UnstageBranchStackOfGrowthCurve2DByGrowthVector(target)

	case *StackOfGrowthCurve2DRibbon:
		stage.UnstageBranchStackOfGrowthCurve2DRibbon(target)

	case *StackOfPartiallyRotatedTorusShape:
		stage.UnstageBranchStackOfPartiallyRotatedTorusShape(target)

	case *StackOfRotatedGrowthCurve2D:
		stage.UnstageBranchStackOfRotatedGrowthCurve2D(target)

	case *StackOfRotatedGrowthCurve2DRibbon:
		stage.UnstageBranchStackOfRotatedGrowthCurve2DRibbon(target)

	case *StackRotatedGrowthCurve2DEndArcShape:
		stage.UnstageBranchStackRotatedGrowthCurve2DEndArcShape(target)

	case *StackRotatedGrowthCurve2DRibbonEndShape:
		stage.UnstageBranchStackRotatedGrowthCurve2DRibbonEndShape(target)

	case *StackRotatedGrowthCurve2DRibbonStartShape:
		stage.UnstageBranchStackRotatedGrowthCurve2DRibbonStartShape(target)

	case *StackRotatedGrowthCurve2DStartArcShape:
		stage.UnstageBranchStackRotatedGrowthCurve2DStartArcShape(target)

	case *StartArcShape:
		stage.UnstageBranchStartArcShape(target)

	case *StartArcShapeGrid:
		stage.UnstageBranchStartArcShapeGrid(target)

	case *StartHalfwayArcShape:
		stage.UnstageBranchStartHalfwayArcShape(target)

	case *StartHalfwayArcShapeGrid:
		stage.UnstageBranchStartHalfwayArcShapeGrid(target)

	case *StemCylinder3DShape:
		stage.UnstageBranchStemCylinder3DShape(target)

	case *Stool2DDiagram:
		stage.UnstageBranchStool2DDiagram(target)

	case *Stool3DDiagram:
		stage.UnstageBranchStool3DDiagram(target)

	case *StoolAbstract:
		stage.UnstageBranchStoolAbstract(target)

	case *TiledFloor3DShape:
		stage.UnstageBranchTiledFloor3DShape(target)

	case *TopEndArcShape:
		stage.UnstageBranchTopEndArcShape(target)

	case *TopEndArcShapeGrid:
		stage.UnstageBranchTopEndArcShapeGrid(target)

	case *TopEndHalfwayArcShape:
		stage.UnstageBranchTopEndHalfwayArcShape(target)

	case *TopEndHalfwayArcShapeGrid:
		stage.UnstageBranchTopEndHalfwayArcShapeGrid(target)

	case *TopGrowthCurve2D:
		stage.UnstageBranchTopGrowthCurve2D(target)

	case *TopMidArcVectorShape:
		stage.UnstageBranchTopMidArcVectorShape(target)

	case *TopMidArcVectorShapeGrid:
		stage.UnstageBranchTopMidArcVectorShapeGrid(target)

	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		stage.UnstageBranchTopStackGrowthCurve2DEndHalfwayArcShape(target)

	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		stage.UnstageBranchTopStackGrowthCurve2DStartHalfwayArcShape(target)

	case *TopStackOfGrowthCurve2D:
		stage.UnstageBranchTopStackOfGrowthCurve2D(target)

	case *TopStackOfRotatedGrowthCurve2D:
		stage.UnstageBranchTopStackOfRotatedGrowthCurve2D(target)

	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		stage.UnstageBranchTopStackOfRotatedGrowthCurve2DEndArcShape(target)

	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		stage.UnstageBranchTopStackOfRotatedGrowthCurve2DStartArcShape(target)

	case *TopStartArcShape:
		stage.UnstageBranchTopStartArcShape(target)

	case *TopStartArcShapeGrid:
		stage.UnstageBranchTopStartArcShapeGrid(target)

	case *TopStartHalfwayArcShape:
		stage.UnstageBranchTopStartHalfwayArcShape(target)

	case *TopStartHalfwayArcShapeGrid:
		stage.UnstageBranchTopStartHalfwayArcShapeGrid(target)

	case *Torus3DShape:
		stage.UnstageBranchTorus3DShape(target)

	case *TorusEdge3DShape:
		stage.UnstageBranchTorusEdge3DShape(target)

	case *TorusStackShape:
		stage.UnstageBranchTorusStackShape(target)

	case *Vase2DDiagram:
		stage.UnstageBranchVase2DDiagram(target)

	case *Vase3DDiagram:
		stage.UnstageBranchVase3DDiagram(target)

	case *VaseAbstract:
		stage.UnstageBranchVaseAbstract(target)

	case *VerticalTorusStackShape:
		stage.UnstageBranchVerticalTorusStackShape(target)

	case *VolumeKey3DShape:
		stage.UnstageBranchVolumeKey3DShape(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchAngle0Shape(angle0shape *Angle0Shape) {

	// check if instance is already staged
	if !IsStaged(stage, angle0shape) {
		return
	}

	angle0shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchArcNormalVectorShape(arcnormalvectorshape *ArcNormalVectorShape) {

	// check if instance is already staged
	if !IsStaged(stage, arcnormalvectorshape) {
		return
	}

	arcnormalvectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchArcNormalVectorShapeGrid(arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, arcnormalvectorshapegrid) {
		return
	}

	arcnormalvectorshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchAxesShape(axesshape *AxesShape) {

	// check if instance is already staged
	if !IsStaged(stage, axesshape) {
		return
	}

	axesshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchBaseVectorShape(basevectorshape *BaseVectorShape) {

	// check if instance is already staged
	if !IsStaged(stage, basevectorshape) {
		return
	}

	basevectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchBaseVectorShapeGrid(basevectorshapegrid *BaseVectorShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, basevectorshapegrid) {
		return
	}

	basevectorshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchChosenP1P2PairShape(chosenp1p2pairshape *ChosenP1P2PairShape) {

	// check if instance is already staged
	if !IsStaged(stage, chosenp1p2pairshape) {
		return
	}

	chosenp1p2pairshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCircleGridShape(circlegridshape *CircleGridShape) {

	// check if instance is already staged
	if !IsStaged(stage, circlegridshape) {
		return
	}

	circlegridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCircumference3DShape(circumference3dshape *Circumference3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, circumference3dshape) {
		return
	}

	circumference3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchClock2DDiagram(clock2ddiagram *Clock2DDiagram) {

	// check if instance is already staged
	if !IsStaged(stage, clock2ddiagram) {
		return
	}

	clock2ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchClock3DDiagram(clock3ddiagram *Clock3DDiagram) {

	// check if instance is already staged
	if !IsStaged(stage, clock3ddiagram) {
		return
	}

	clock3ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if clock3ddiagram.SampledPoints3DShape != nil {
		UnstageBranch(stage, clock3ddiagram.SampledPoints3DShape)
	}
	if clock3ddiagram.Rendered3DShape != nil {
		UnstageBranch(stage, clock3ddiagram.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchClockAbstract(clockabstract *ClockAbstract) {

	// check if instance is already staged
	if !IsStaged(stage, clockabstract) {
		return
	}

	clockabstract.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchClockTopCurveShape(clocktopcurveshape *ClockTopCurveShape) {

	// check if instance is already staged
	if !IsStaged(stage, clocktopcurveshape) {
		return
	}

	clocktopcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCutLine3DShape(cutline3dshape *CutLine3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, cutline3dshape) {
		return
	}

	cutline3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEndArcShape(endarcshape *EndArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, endarcshape) {
		return
	}

	endarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEndArcShapeGrid(endarcshapegrid *EndArcShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, endarcshapegrid) {
		return
	}

	endarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEndHalfwayArcShape(endhalfwayarcshape *EndHalfwayArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, endhalfwayarcshape) {
		return
	}

	endhalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEndHalfwayArcShapeGrid(endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, endhalfwayarcshapegrid) {
		return
	}

	endhalfwayarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchExplanationTextShape(explanationtextshape *ExplanationTextShape) {

	// check if instance is already staged
	if !IsStaged(stage, explanationtextshape) {
		return
	}

	explanationtextshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEye3DShape(eye3dshape *Eye3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, eye3dshape) {
		return
	}

	eye3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEyeCornersSampledPoints3DShape(eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, eyecornerssampledpoints3dshape) {
		return
	}

	eyecornerssampledpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEyeSampledPoints3DShape(eyesampledpoints3dshape *EyeSampledPoints3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, eyesampledpoints3dshape) {
		return
	}

	eyesampledpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEyeSeatBottomCurveShape(eyeseatbottomcurveshape *EyeSeatBottomCurveShape) {

	// check if instance is already staged
	if !IsStaged(stage, eyeseatbottomcurveshape) {
		return
	}

	eyeseatbottomcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEyeStoolBottomCurveShape(eyestoolbottomcurveshape *EyeStoolBottomCurveShape) {

	// check if instance is already staged
	if !IsStaged(stage, eyestoolbottomcurveshape) {
		return
	}

	eyestoolbottomcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEyeVolume3DShape(eyevolume3dshape *EyeVolume3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, eyevolume3dshape) {
		return
	}

	eyevolume3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGridPathShape(gridpathshape *GridPathShape) {

	// check if instance is already staged
	if !IsStaged(stage, gridpathshape) {
		return
	}

	gridpathshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGrowthCurve2D(growthcurve2d *GrowthCurve2D) {

	// check if instance is already staged
	if !IsStaged(stage, growthcurve2d) {
		return
	}

	growthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGrowthCurve2DRibbon(growthcurve2dribbon *GrowthCurve2DRibbon) {

	// check if instance is already staged
	if !IsStaged(stage, growthcurve2dribbon) {
		return
	}

	growthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGrowthCurve2DRibbonEndShape(growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !IsStaged(stage, growthcurve2dribbonendshape) {
		return
	}

	growthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGrowthCurve2DRibbonStartShape(growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !IsStaged(stage, growthcurve2dribbonstartshape) {
		return
	}

	growthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGrowthCurveRhombusGridShape(growthcurverhombusgridshape *GrowthCurveRhombusGridShape) {

	// check if instance is already staged
	if !IsStaged(stage, growthcurverhombusgridshape) {
		return
	}

	growthcurverhombusgridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGrowthCurveRhombusShape(growthcurverhombusshape *GrowthCurveRhombusShape) {

	// check if instance is already staged
	if !IsStaged(stage, growthcurverhombusshape) {
		return
	}

	growthcurverhombusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGrowthVectorShape(growthvectorshape *GrowthVectorShape) {

	// check if instance is already staged
	if !IsStaged(stage, growthvectorshape) {
		return
	}

	growthvectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchInitialRhombusGridShape(initialrhombusgridshape *InitialRhombusGridShape) {

	// check if instance is already staged
	if !IsStaged(stage, initialrhombusgridshape) {
		return
	}

	initialrhombusgridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchInitialRhombusShape(initialrhombusshape *InitialRhombusShape) {

	// check if instance is already staged
	if !IsStaged(stage, initialrhombusshape) {
		return
	}

	initialrhombusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchKey3DShape(key3dshape *Key3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, key3dshape) {
		return
	}

	key3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchKeyHole3DShape(keyhole3dshape *KeyHole3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, keyhole3dshape) {
		return
	}

	keyhole3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchKeyHoleShape(keyholeshape *KeyHoleShape) {

	// check if instance is already staged
	if !IsStaged(stage, keyholeshape) {
		return
	}

	keyholeshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchLibrary(library *Library) {

	// check if instance is already staged
	if !IsStaged(stage, library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		UnstageBranch(stage, _library)
	}
	for _, _plantabstract := range library.Plants {
		UnstageBranch(stage, _plantabstract)
	}

}

func (stage *Stage) UnstageBranchMidArcVectorShape(midarcvectorshape *MidArcVectorShape) {

	// check if instance is already staged
	if !IsStaged(stage, midarcvectorshape) {
		return
	}

	midarcvectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMidArcVectorShapeGrid(midarcvectorshapegrid *MidArcVectorShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, midarcvectorshapegrid) {
		return
	}

	midarcvectorshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchOriginalPoints3DShape(originalpoints3dshape *OriginalPoints3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, originalpoints3dshape) {
		return
	}

	originalpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchParastichyMCurves3DShape(parastichymcurves3dshape *ParastichyMCurves3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, parastichymcurves3dshape) {
		return
	}

	parastichymcurves3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchParastichyNCurves3DShape(parastichyncurves3dshape *ParastichyNCurves3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, parastichyncurves3dshape) {
		return
	}

	parastichyncurves3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DRibbon(partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) {

	// check if instance is already staged
	if !IsStaged(stage, partiallygrowthcurve2dribbon) {
		return
	}

	partiallygrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DRibbonEndShape(partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !IsStaged(stage, partiallygrowthcurve2dribbonendshape) {
		return
	}

	partiallygrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DRibbonStartShape(partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !IsStaged(stage, partiallygrowthcurve2dribbonstartshape) {
		return
	}

	partiallygrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectory(partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) {

	// check if instance is already staged
	if !IsStaged(stage, partiallygrowthcurve2dtrajectory) {
		return
	}

	partiallygrowthcurve2dtrajectory.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryP1CurveShape(partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) {

	// check if instance is already staged
	if !IsStaged(stage, partiallygrowthcurve2dtrajectoryp1curveshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1curveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryP1P2(partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) {

	// check if instance is already staged
	if !IsStaged(stage, partiallygrowthcurve2dtrajectoryp1p2) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) {

	// check if instance is already staged
	if !IsStaged(stage, partiallygrowthcurve2dtrajectoryp1p2pairlineshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2pairlineshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryP1PointShape(partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) {

	// check if instance is already staged
	if !IsStaged(stage, partiallygrowthcurve2dtrajectoryp1pointshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1pointshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryP2CurveShape(partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) {

	// check if instance is already staged
	if !IsStaged(stage, partiallygrowthcurve2dtrajectoryp2curveshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp2curveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryP2PointShape(partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) {

	// check if instance is already staged
	if !IsStaged(stage, partiallygrowthcurve2dtrajectoryp2pointshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp2pointshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryShape(partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) {

	// check if instance is already staged
	if !IsStaged(stage, partiallygrowthcurve2dtrajectoryshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyRotatedSeatBottomCurveShape(partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) {

	// check if instance is already staged
	if !IsStaged(stage, partiallyrotatedseatbottomcurveshape) {
		return
	}

	partiallyrotatedseatbottomcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyRotatedSeatTopCurveShape(partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) {

	// check if instance is already staged
	if !IsStaged(stage, partiallyrotatedseattopcurveshape) {
		return
	}

	partiallyrotatedseattopcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPartiallyRotatedTorusShape(partiallyrotatedtorusshape *PartiallyRotatedTorusShape) {

	// check if instance is already staged
	if !IsStaged(stage, partiallyrotatedtorusshape) {
		return
	}

	partiallyrotatedtorusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPerpendicularVector(perpendicularvector *PerpendicularVector) {

	// check if instance is already staged
	if !IsStaged(stage, perpendicularvector) {
		return
	}

	perpendicularvector.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPerpendicularVectorGrid(perpendicularvectorgrid *PerpendicularVectorGrid) {

	// check if instance is already staged
	if !IsStaged(stage, perpendicularvectorgrid) {
		return
	}

	perpendicularvectorgrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) {

	// check if instance is already staged
	if !IsStaged(stage, perpendicularvectorgridhalfway) {
		return
	}

	perpendicularvectorgridhalfway.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPerpendicularVectorHalfway(perpendicularvectorhalfway *PerpendicularVectorHalfway) {

	// check if instance is already staged
	if !IsStaged(stage, perpendicularvectorhalfway) {
		return
	}

	perpendicularvectorhalfway.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPlant2DDiagram(plant2ddiagram *Plant2DDiagram) {

	// check if instance is already staged
	if !IsStaged(stage, plant2ddiagram) {
		return
	}

	plant2ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPlant3DDiagram(plant3ddiagram *Plant3DDiagram) {

	// check if instance is already staged
	if !IsStaged(stage, plant3ddiagram) {
		return
	}

	plant3ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plant3ddiagram.StemCylinder3DShape != nil {
		UnstageBranch(stage, plant3ddiagram.StemCylinder3DShape)
	}
	if plant3ddiagram.ParastichyNCurves3DShape != nil {
		UnstageBranch(stage, plant3ddiagram.ParastichyNCurves3DShape)
	}
	if plant3ddiagram.ParastichyMCurves3DShape != nil {
		UnstageBranch(stage, plant3ddiagram.ParastichyMCurves3DShape)
	}
	if plant3ddiagram.CutLine3DShape != nil {
		UnstageBranch(stage, plant3ddiagram.CutLine3DShape)
	}
	if plant3ddiagram.Circumference3DShape != nil {
		UnstageBranch(stage, plant3ddiagram.Circumference3DShape)
	}
	if plant3ddiagram.Rendered3DShape != nil {
		UnstageBranch(stage, plant3ddiagram.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPlantAbstract(plantabstract *PlantAbstract) {

	// check if instance is already staged
	if !IsStaged(stage, plantabstract) {
		return
	}

	plantabstract.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plantabstract.VaseAbstract != nil {
		UnstageBranch(stage, plantabstract.VaseAbstract)
	}
	if plantabstract.StoolAbstract != nil {
		UnstageBranch(stage, plantabstract.StoolAbstract)
	}
	if plantabstract.ClockAbstract != nil {
		UnstageBranch(stage, plantabstract.ClockAbstract)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plant2ddiagram := range plantabstract.Plant2DDiagrams {
		UnstageBranch(stage, _plant2ddiagram)
	}
	for _, _plant3ddiagram := range plantabstract.Plant3DDiagrams {
		UnstageBranch(stage, _plant3ddiagram)
	}
	for _, _vase2ddiagram := range plantabstract.Vase2DDiagrams {
		UnstageBranch(stage, _vase2ddiagram)
	}
	for _, _vase3ddiagram := range plantabstract.Vase3DDiagrams {
		UnstageBranch(stage, _vase3ddiagram)
	}
	for _, _stool2ddiagram := range plantabstract.Stool2DDiagrams {
		UnstageBranch(stage, _stool2ddiagram)
	}
	for _, _stool3ddiagram := range plantabstract.Stool3DDiagrams {
		UnstageBranch(stage, _stool3ddiagram)
	}
	for _, _clock2ddiagram := range plantabstract.Clock2DDiagrams {
		UnstageBranch(stage, _clock2ddiagram)
	}
	for _, _clock3ddiagram := range plantabstract.Clock3DDiagrams {
		UnstageBranch(stage, _clock3ddiagram)
	}

}

func (stage *Stage) UnstageBranchPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape) {

	// check if instance is already staged
	if !IsStaged(stage, plantcircumferenceshape) {
		return
	}

	plantcircumferenceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPointsAndLines3DShape(pointsandlines3dshape *PointsAndLines3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, pointsandlines3dshape) {
		return
	}

	pointsandlines3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPxShape(pxshape *PxShape) {

	// check if instance is already staged
	if !IsStaged(stage, pxshape) {
		return
	}

	pxshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRendered3DShape(rendered3dshape *Rendered3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, rendered3dshape) {
		return
	}

	rendered3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRhombusShape(rhombusshape *RhombusShape) {

	// check if instance is already staged
	if !IsStaged(stage, rhombusshape) {
		return
	}

	rhombusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRhombusStuff(rhombusstuff *RhombusStuff) {

	// check if instance is already staged
	if !IsStaged(stage, rhombusstuff) {
		return
	}

	rhombusstuff.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRotatedRhombusGridShape(rotatedrhombusgridshape *RotatedRhombusGridShape) {

	// check if instance is already staged
	if !IsStaged(stage, rotatedrhombusgridshape) {
		return
	}

	rotatedrhombusgridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRotatedRhombusShape(rotatedrhombusshape *RotatedRhombusShape) {

	// check if instance is already staged
	if !IsStaged(stage, rotatedrhombusshape) {
		return
	}

	rotatedrhombusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRotatedSampledPoints3DShape(rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, rotatedsampledpoints3dshape) {
		return
	}

	rotatedsampledpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRotatedSeatAndLegs3DShape(rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, rotatedseatandlegs3dshape) {
		return
	}

	rotatedseatandlegs3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSampledPoints3DShape(sampledpoints3dshape *SampledPoints3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, sampledpoints3dshape) {
		return
	}

	sampledpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSeat3DShape(seat3dshape *Seat3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, seat3dshape) {
		return
	}

	seat3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSeatAndLegs3DShape(seatandlegs3dshape *SeatAndLegs3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, seatandlegs3dshape) {
		return
	}

	seatandlegs3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSeatBottomCurveShape(seatbottomcurveshape *SeatBottomCurveShape) {

	// check if instance is already staged
	if !IsStaged(stage, seatbottomcurveshape) {
		return
	}

	seatbottomcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSeatTopCurveShape(seattopcurveshape *SeatTopCurveShape) {

	// check if instance is already staged
	if !IsStaged(stage, seattopcurveshape) {
		return
	}

	seattopcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedBottomTopStartArcShape(shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedbottomtopstartarcshape) {
		return
	}

	shiftedbottomtopstartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedBottomTopStartArcShapeGrid(shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedbottomtopstartarcshapegrid) {
		return
	}

	shiftedbottomtopstartarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedLeftGrowthCurve2DRibbon(shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedleftgrowthcurve2dribbon) {
		return
	}

	shiftedleftgrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedLeftGrowthCurve2DRibbonEndShape(shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedleftgrowthcurve2dribbonendshape) {
		return
	}

	shiftedleftgrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedLeftGrowthCurve2DRibbonStartShape(shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedleftgrowthcurve2dribbonstartshape) {
		return
	}

	shiftedleftgrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedLeftPartiallyGrowthCurve2DRibbon(shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedleftpartiallygrowthcurve2dribbon) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedleftpartiallygrowthcurve2dribbonendshape) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedleftpartiallygrowthcurve2dribbonstartshape) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedLeftStackGrowthCurveEndArcShape(shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedleftstackgrowthcurveendarcshape) {
		return
	}

	shiftedleftstackgrowthcurveendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedLeftStackGrowthCurveStartArcShape(shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedleftstackgrowthcurvestartarcshape) {
		return
	}

	shiftedleftstackgrowthcurvestartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedLeftStackNormalVector(shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedleftstacknormalvector) {
		return
	}

	shiftedleftstacknormalvector.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedLeftStackOfGrowthCurve(shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedleftstackofgrowthcurve) {
		return
	}

	shiftedleftstackofgrowthcurve.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedleftstackofnormalvector) {
		return
	}

	shiftedleftstackofnormalvector.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedRightGrowthCurve2DRibbon(shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedrightgrowthcurve2dribbon) {
		return
	}

	shiftedrightgrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedRightGrowthCurve2DRibbonEndShape(shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedrightgrowthcurve2dribbonendshape) {
		return
	}

	shiftedrightgrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShiftedRightGrowthCurve2DRibbonStartShape(shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedrightgrowthcurve2dribbonstartshape) {
		return
	}

	shiftedrightgrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackGrowthCurve2DEndHalfwayArcShape(stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, stackgrowthcurve2dendhalfwayarcshape) {
		return
	}

	stackgrowthcurve2dendhalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackGrowthCurve2DRibbonEndShape(stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !IsStaged(stage, stackgrowthcurve2dribbonendshape) {
		return
	}

	stackgrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackGrowthCurve2DRibbonStartShape(stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !IsStaged(stage, stackgrowthcurve2dribbonstartshape) {
		return
	}

	stackgrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackGrowthCurve2DStartHalfwayArcShape(stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, stackgrowthcurve2dstarthalfwayarcshape) {
		return
	}

	stackgrowthcurve2dstarthalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackOfGrowthCurve2D(stackofgrowthcurve2d *StackOfGrowthCurve2D) {

	// check if instance is already staged
	if !IsStaged(stage, stackofgrowthcurve2d) {
		return
	}

	stackofgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackOfGrowthCurve2DByGrowthVector(stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) {

	// check if instance is already staged
	if !IsStaged(stage, stackofgrowthcurve2dbygrowthvector) {
		return
	}

	stackofgrowthcurve2dbygrowthvector.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackOfGrowthCurve2DRibbon(stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) {

	// check if instance is already staged
	if !IsStaged(stage, stackofgrowthcurve2dribbon) {
		return
	}

	stackofgrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackOfPartiallyRotatedTorusShape(stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) {

	// check if instance is already staged
	if !IsStaged(stage, stackofpartiallyrotatedtorusshape) {
		return
	}

	stackofpartiallyrotatedtorusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackOfRotatedGrowthCurve2D(stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) {

	// check if instance is already staged
	if !IsStaged(stage, stackofrotatedgrowthcurve2d) {
		return
	}

	stackofrotatedgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackOfRotatedGrowthCurve2DRibbon(stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) {

	// check if instance is already staged
	if !IsStaged(stage, stackofrotatedgrowthcurve2dribbon) {
		return
	}

	stackofrotatedgrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackRotatedGrowthCurve2DEndArcShape(stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, stackrotatedgrowthcurve2dendarcshape) {
		return
	}

	stackrotatedgrowthcurve2dendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackRotatedGrowthCurve2DRibbonEndShape(stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !IsStaged(stage, stackrotatedgrowthcurve2dribbonendshape) {
		return
	}

	stackrotatedgrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackRotatedGrowthCurve2DRibbonStartShape(stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !IsStaged(stage, stackrotatedgrowthcurve2dribbonstartshape) {
		return
	}

	stackrotatedgrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackRotatedGrowthCurve2DStartArcShape(stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, stackrotatedgrowthcurve2dstartarcshape) {
		return
	}

	stackrotatedgrowthcurve2dstartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStartArcShape(startarcshape *StartArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, startarcshape) {
		return
	}

	startarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStartArcShapeGrid(startarcshapegrid *StartArcShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, startarcshapegrid) {
		return
	}

	startarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStartHalfwayArcShape(starthalfwayarcshape *StartHalfwayArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, starthalfwayarcshape) {
		return
	}

	starthalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStartHalfwayArcShapeGrid(starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, starthalfwayarcshapegrid) {
		return
	}

	starthalfwayarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStemCylinder3DShape(stemcylinder3dshape *StemCylinder3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, stemcylinder3dshape) {
		return
	}

	stemcylinder3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStool2DDiagram(stool2ddiagram *Stool2DDiagram) {

	// check if instance is already staged
	if !IsStaged(stage, stool2ddiagram) {
		return
	}

	stool2ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStool3DDiagram(stool3ddiagram *Stool3DDiagram) {

	// check if instance is already staged
	if !IsStaged(stage, stool3ddiagram) {
		return
	}

	stool3ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stool3ddiagram.SampledPoints3DShape != nil {
		UnstageBranch(stage, stool3ddiagram.SampledPoints3DShape)
	}
	if stool3ddiagram.Rendered3DShape != nil {
		UnstageBranch(stage, stool3ddiagram.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStoolAbstract(stoolabstract *StoolAbstract) {

	// check if instance is already staged
	if !IsStaged(stage, stoolabstract) {
		return
	}

	stoolabstract.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTiledFloor3DShape(tiledfloor3dshape *TiledFloor3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, tiledfloor3dshape) {
		return
	}

	tiledfloor3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopEndArcShape(topendarcshape *TopEndArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, topendarcshape) {
		return
	}

	topendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopEndArcShapeGrid(topendarcshapegrid *TopEndArcShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, topendarcshapegrid) {
		return
	}

	topendarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopEndHalfwayArcShape(topendhalfwayarcshape *TopEndHalfwayArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, topendhalfwayarcshape) {
		return
	}

	topendhalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopEndHalfwayArcShapeGrid(topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, topendhalfwayarcshapegrid) {
		return
	}

	topendhalfwayarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D) {

	// check if instance is already staged
	if !IsStaged(stage, topgrowthcurve2d) {
		return
	}

	topgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopMidArcVectorShape(topmidarcvectorshape *TopMidArcVectorShape) {

	// check if instance is already staged
	if !IsStaged(stage, topmidarcvectorshape) {
		return
	}

	topmidarcvectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopMidArcVectorShapeGrid(topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, topmidarcvectorshapegrid) {
		return
	}

	topmidarcvectorshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStackGrowthCurve2DEndHalfwayArcShape(topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, topstackgrowthcurve2dendhalfwayarcshape) {
		return
	}

	topstackgrowthcurve2dendhalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStackGrowthCurve2DStartHalfwayArcShape(topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, topstackgrowthcurve2dstarthalfwayarcshape) {
		return
	}

	topstackgrowthcurve2dstarthalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStackOfGrowthCurve2D(topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) {

	// check if instance is already staged
	if !IsStaged(stage, topstackofgrowthcurve2d) {
		return
	}

	topstackofgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStackOfRotatedGrowthCurve2D(topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) {

	// check if instance is already staged
	if !IsStaged(stage, topstackofrotatedgrowthcurve2d) {
		return
	}

	topstackofrotatedgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStackOfRotatedGrowthCurve2DEndArcShape(topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, topstackofrotatedgrowthcurve2dendarcshape) {
		return
	}

	topstackofrotatedgrowthcurve2dendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStackOfRotatedGrowthCurve2DStartArcShape(topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, topstackofrotatedgrowthcurve2dstartarcshape) {
		return
	}

	topstackofrotatedgrowthcurve2dstartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStartArcShape(topstartarcshape *TopStartArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, topstartarcshape) {
		return
	}

	topstartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStartArcShapeGrid(topstartarcshapegrid *TopStartArcShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, topstartarcshapegrid) {
		return
	}

	topstartarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStartHalfwayArcShape(topstarthalfwayarcshape *TopStartHalfwayArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, topstarthalfwayarcshape) {
		return
	}

	topstarthalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStartHalfwayArcShapeGrid(topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, topstarthalfwayarcshapegrid) {
		return
	}

	topstarthalfwayarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTorus3DShape(torus3dshape *Torus3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, torus3dshape) {
		return
	}

	torus3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTorusEdge3DShape(torusedge3dshape *TorusEdge3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, torusedge3dshape) {
		return
	}

	torusedge3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTorusStackShape(torusstackshape *TorusStackShape) {

	// check if instance is already staged
	if !IsStaged(stage, torusstackshape) {
		return
	}

	torusstackshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchVase2DDiagram(vase2ddiagram *Vase2DDiagram) {

	// check if instance is already staged
	if !IsStaged(stage, vase2ddiagram) {
		return
	}

	vase2ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchVase3DDiagram(vase3ddiagram *Vase3DDiagram) {

	// check if instance is already staged
	if !IsStaged(stage, vase3ddiagram) {
		return
	}

	vase3ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if vase3ddiagram.Rendered3DShape != nil {
		UnstageBranch(stage, vase3ddiagram.Rendered3DShape)
	}
	if vase3ddiagram.SampledPoints3DShape != nil {
		UnstageBranch(stage, vase3ddiagram.SampledPoints3DShape)
	}
	if vase3ddiagram.OriginalPoints3DShape != nil {
		UnstageBranch(stage, vase3ddiagram.OriginalPoints3DShape)
	}
	if vase3ddiagram.Angle0Shape != nil {
		UnstageBranch(stage, vase3ddiagram.Angle0Shape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchVaseAbstract(vaseabstract *VaseAbstract) {

	// check if instance is already staged
	if !IsStaged(stage, vaseabstract) {
		return
	}

	vaseabstract.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchVerticalTorusStackShape(verticaltorusstackshape *VerticalTorusStackShape) {

	// check if instance is already staged
	if !IsStaged(stage, verticaltorusstackshape) {
		return
	}

	verticaltorusstackshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchVolumeKey3DShape(volumekey3dshape *VolumeKey3DShape) {

	// check if instance is already staged
	if !IsStaged(stage, volumekey3dshape) {
		return
	}

	volumekey3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Angle0Shape) GongReconstructPointersFromReferences(stage *Stage, instance *Angle0Shape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ArcNormalVectorShape) GongReconstructPointersFromReferences(stage *Stage, instance *ArcNormalVectorShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ArcNormalVectorShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *ArcNormalVectorShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AxesShape) GongReconstructPointersFromReferences(stage *Stage, instance *AxesShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *BaseVectorShape) GongReconstructPointersFromReferences(stage *Stage, instance *BaseVectorShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *BaseVectorShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *BaseVectorShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ChosenP1P2PairShape) GongReconstructPointersFromReferences(stage *Stage, instance *ChosenP1P2PairShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CircleGridShape) GongReconstructPointersFromReferences(stage *Stage, instance *CircleGridShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Circumference3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *Circumference3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Clock2DDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *Clock2DDiagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Clock3DDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *Clock3DDiagram) {
	// insertion point for pointers field
	if instance.SampledPoints3DShape != nil {
		reference.SampledPoints3DShape = stage.SampledPoints3DShapes_reference[instance.SampledPoints3DShape]
	}
	if instance.Rendered3DShape != nil {
		reference.Rendered3DShape = stage.Rendered3DShapes_reference[instance.Rendered3DShape]
	}
	// insertion point for slice of pointers field
}

func (reference *ClockAbstract) GongReconstructPointersFromReferences(stage *Stage, instance *ClockAbstract) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ClockTopCurveShape) GongReconstructPointersFromReferences(stage *Stage, instance *ClockTopCurveShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CutLine3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *CutLine3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EndArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *EndArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EndArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *EndArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EndHalfwayArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *EndHalfwayArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EndHalfwayArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *EndHalfwayArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ExplanationTextShape) GongReconstructPointersFromReferences(stage *Stage, instance *ExplanationTextShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Eye3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *Eye3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EyeCornersSampledPoints3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *EyeCornersSampledPoints3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EyeSampledPoints3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *EyeSampledPoints3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EyeSeatBottomCurveShape) GongReconstructPointersFromReferences(stage *Stage, instance *EyeSeatBottomCurveShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EyeStoolBottomCurveShape) GongReconstructPointersFromReferences(stage *Stage, instance *EyeStoolBottomCurveShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EyeVolume3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *EyeVolume3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GridPathShape) GongReconstructPointersFromReferences(stage *Stage, instance *GridPathShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GrowthCurve2D) GongReconstructPointersFromReferences(stage *Stage, instance *GrowthCurve2D) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GrowthCurve2DRibbon) GongReconstructPointersFromReferences(stage *Stage, instance *GrowthCurve2DRibbon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GrowthCurve2DRibbonEndShape) GongReconstructPointersFromReferences(stage *Stage, instance *GrowthCurve2DRibbonEndShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GrowthCurve2DRibbonStartShape) GongReconstructPointersFromReferences(stage *Stage, instance *GrowthCurve2DRibbonStartShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GrowthCurveRhombusGridShape) GongReconstructPointersFromReferences(stage *Stage, instance *GrowthCurveRhombusGridShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GrowthCurveRhombusShape) GongReconstructPointersFromReferences(stage *Stage, instance *GrowthCurveRhombusShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GrowthVectorShape) GongReconstructPointersFromReferences(stage *Stage, instance *GrowthVectorShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *InitialRhombusGridShape) GongReconstructPointersFromReferences(stage *Stage, instance *InitialRhombusGridShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *InitialRhombusShape) GongReconstructPointersFromReferences(stage *Stage, instance *InitialRhombusShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Key3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *Key3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *KeyHole3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *KeyHole3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *KeyHoleShape) GongReconstructPointersFromReferences(stage *Stage, instance *KeyHoleShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SubLibraries = reference.SubLibraries[:0]
	for _, _b := range instance.SubLibraries {
		reference.SubLibraries = append(reference.SubLibraries, stage.Librarys_reference[_b])
	}
	reference.Plants = reference.Plants[:0]
	for _, _b := range instance.Plants {
		reference.Plants = append(reference.Plants, stage.PlantAbstracts_reference[_b])
	}
}

func (reference *MidArcVectorShape) GongReconstructPointersFromReferences(stage *Stage, instance *MidArcVectorShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MidArcVectorShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *MidArcVectorShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *OriginalPoints3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *OriginalPoints3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ParastichyMCurves3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *ParastichyMCurves3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ParastichyNCurves3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *ParastichyNCurves3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyGrowthCurve2DRibbon) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyGrowthCurve2DRibbon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyGrowthCurve2DRibbonEndShape) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyGrowthCurve2DRibbonEndShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyGrowthCurve2DRibbonStartShape) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyGrowthCurve2DRibbonStartShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyGrowthCurve2DTrajectory) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyGrowthCurve2DTrajectory) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyGrowthCurve2DTrajectoryP1CurveShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyGrowthCurve2DTrajectoryP1P2) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyGrowthCurve2DTrajectoryP1P2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyGrowthCurve2DTrajectoryP1PointShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyGrowthCurve2DTrajectoryP2CurveShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyGrowthCurve2DTrajectoryP2PointShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyGrowthCurve2DTrajectoryShape) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyGrowthCurve2DTrajectoryShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyRotatedSeatBottomCurveShape) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyRotatedSeatBottomCurveShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyRotatedSeatTopCurveShape) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyRotatedSeatTopCurveShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartiallyRotatedTorusShape) GongReconstructPointersFromReferences(stage *Stage, instance *PartiallyRotatedTorusShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PerpendicularVector) GongReconstructPointersFromReferences(stage *Stage, instance *PerpendicularVector) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PerpendicularVectorGrid) GongReconstructPointersFromReferences(stage *Stage, instance *PerpendicularVectorGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PerpendicularVectorGridHalfway) GongReconstructPointersFromReferences(stage *Stage, instance *PerpendicularVectorGridHalfway) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PerpendicularVectorHalfway) GongReconstructPointersFromReferences(stage *Stage, instance *PerpendicularVectorHalfway) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Plant2DDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *Plant2DDiagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Plant3DDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *Plant3DDiagram) {
	// insertion point for pointers field
	if instance.StemCylinder3DShape != nil {
		reference.StemCylinder3DShape = stage.StemCylinder3DShapes_reference[instance.StemCylinder3DShape]
	}
	if instance.ParastichyNCurves3DShape != nil {
		reference.ParastichyNCurves3DShape = stage.ParastichyNCurves3DShapes_reference[instance.ParastichyNCurves3DShape]
	}
	if instance.ParastichyMCurves3DShape != nil {
		reference.ParastichyMCurves3DShape = stage.ParastichyMCurves3DShapes_reference[instance.ParastichyMCurves3DShape]
	}
	if instance.CutLine3DShape != nil {
		reference.CutLine3DShape = stage.CutLine3DShapes_reference[instance.CutLine3DShape]
	}
	if instance.Circumference3DShape != nil {
		reference.Circumference3DShape = stage.Circumference3DShapes_reference[instance.Circumference3DShape]
	}
	if instance.Rendered3DShape != nil {
		reference.Rendered3DShape = stage.Rendered3DShapes_reference[instance.Rendered3DShape]
	}
	// insertion point for slice of pointers field
}

func (reference *PlantAbstract) GongReconstructPointersFromReferences(stage *Stage, instance *PlantAbstract) {
	// insertion point for pointers field
	if instance.VaseAbstract != nil {
		reference.VaseAbstract = stage.VaseAbstracts_reference[instance.VaseAbstract]
	}
	if instance.StoolAbstract != nil {
		reference.StoolAbstract = stage.StoolAbstracts_reference[instance.StoolAbstract]
	}
	if instance.ClockAbstract != nil {
		reference.ClockAbstract = stage.ClockAbstracts_reference[instance.ClockAbstract]
	}
	// insertion point for slice of pointers field
	reference.Plant2DDiagrams = reference.Plant2DDiagrams[:0]
	for _, _b := range instance.Plant2DDiagrams {
		reference.Plant2DDiagrams = append(reference.Plant2DDiagrams, stage.Plant2DDiagrams_reference[_b])
	}
	reference.Plant3DDiagrams = reference.Plant3DDiagrams[:0]
	for _, _b := range instance.Plant3DDiagrams {
		reference.Plant3DDiagrams = append(reference.Plant3DDiagrams, stage.Plant3DDiagrams_reference[_b])
	}
	reference.Vase2DDiagrams = reference.Vase2DDiagrams[:0]
	for _, _b := range instance.Vase2DDiagrams {
		reference.Vase2DDiagrams = append(reference.Vase2DDiagrams, stage.Vase2DDiagrams_reference[_b])
	}
	reference.Vase3DDiagrams = reference.Vase3DDiagrams[:0]
	for _, _b := range instance.Vase3DDiagrams {
		reference.Vase3DDiagrams = append(reference.Vase3DDiagrams, stage.Vase3DDiagrams_reference[_b])
	}
	reference.Stool2DDiagrams = reference.Stool2DDiagrams[:0]
	for _, _b := range instance.Stool2DDiagrams {
		reference.Stool2DDiagrams = append(reference.Stool2DDiagrams, stage.Stool2DDiagrams_reference[_b])
	}
	reference.Stool3DDiagrams = reference.Stool3DDiagrams[:0]
	for _, _b := range instance.Stool3DDiagrams {
		reference.Stool3DDiagrams = append(reference.Stool3DDiagrams, stage.Stool3DDiagrams_reference[_b])
	}
	reference.Clock2DDiagrams = reference.Clock2DDiagrams[:0]
	for _, _b := range instance.Clock2DDiagrams {
		reference.Clock2DDiagrams = append(reference.Clock2DDiagrams, stage.Clock2DDiagrams_reference[_b])
	}
	reference.Clock3DDiagrams = reference.Clock3DDiagrams[:0]
	for _, _b := range instance.Clock3DDiagrams {
		reference.Clock3DDiagrams = append(reference.Clock3DDiagrams, stage.Clock3DDiagrams_reference[_b])
	}
}

func (reference *PlantCircumferenceShape) GongReconstructPointersFromReferences(stage *Stage, instance *PlantCircumferenceShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PointsAndLines3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *PointsAndLines3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PxShape) GongReconstructPointersFromReferences(stage *Stage, instance *PxShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Rendered3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *Rendered3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *RhombusShape) GongReconstructPointersFromReferences(stage *Stage, instance *RhombusShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *RhombusStuff) GongReconstructPointersFromReferences(stage *Stage, instance *RhombusStuff) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *RotatedRhombusGridShape) GongReconstructPointersFromReferences(stage *Stage, instance *RotatedRhombusGridShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *RotatedRhombusShape) GongReconstructPointersFromReferences(stage *Stage, instance *RotatedRhombusShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *RotatedSampledPoints3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *RotatedSampledPoints3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *RotatedSeatAndLegs3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *RotatedSeatAndLegs3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *SampledPoints3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *SampledPoints3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Seat3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *Seat3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *SeatAndLegs3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *SeatAndLegs3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *SeatBottomCurveShape) GongReconstructPointersFromReferences(stage *Stage, instance *SeatBottomCurveShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *SeatTopCurveShape) GongReconstructPointersFromReferences(stage *Stage, instance *SeatTopCurveShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedBottomTopStartArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedBottomTopStartArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedBottomTopStartArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedBottomTopStartArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedLeftGrowthCurve2DRibbon) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedLeftGrowthCurve2DRibbon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedLeftGrowthCurve2DRibbonEndShape) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedLeftGrowthCurve2DRibbonEndShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedLeftGrowthCurve2DRibbonStartShape) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedLeftGrowthCurve2DRibbonStartShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedLeftPartiallyGrowthCurve2DRibbon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedLeftStackGrowthCurveEndArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedLeftStackGrowthCurveEndArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedLeftStackGrowthCurveStartArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedLeftStackGrowthCurveStartArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedLeftStackNormalVector) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedLeftStackNormalVector) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedLeftStackOfGrowthCurve) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedLeftStackOfGrowthCurve) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedLeftStackOfNormalVector) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedLeftStackOfNormalVector) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedRightGrowthCurve2DRibbon) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedRightGrowthCurve2DRibbon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedRightGrowthCurve2DRibbonEndShape) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedRightGrowthCurve2DRibbonEndShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ShiftedRightGrowthCurve2DRibbonStartShape) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedRightGrowthCurve2DRibbonStartShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackGrowthCurve2DEndHalfwayArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackGrowthCurve2DEndHalfwayArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackGrowthCurve2DRibbonEndShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackGrowthCurve2DRibbonEndShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackGrowthCurve2DRibbonStartShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackGrowthCurve2DRibbonStartShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackGrowthCurve2DStartHalfwayArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackGrowthCurve2DStartHalfwayArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackOfGrowthCurve2D) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfGrowthCurve2D) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackOfGrowthCurve2DByGrowthVector) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfGrowthCurve2DByGrowthVector) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackOfGrowthCurve2DRibbon) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfGrowthCurve2DRibbon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackOfPartiallyRotatedTorusShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfPartiallyRotatedTorusShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackOfRotatedGrowthCurve2D) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfRotatedGrowthCurve2D) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackOfRotatedGrowthCurve2DRibbon) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfRotatedGrowthCurve2DRibbon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackRotatedGrowthCurve2DEndArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackRotatedGrowthCurve2DEndArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackRotatedGrowthCurve2DRibbonEndShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackRotatedGrowthCurve2DRibbonEndShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackRotatedGrowthCurve2DRibbonStartShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackRotatedGrowthCurve2DRibbonStartShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackRotatedGrowthCurve2DStartArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackRotatedGrowthCurve2DStartArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StartArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *StartArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StartArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *StartArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StartHalfwayArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *StartHalfwayArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StartHalfwayArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *StartHalfwayArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StemCylinder3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *StemCylinder3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Stool2DDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *Stool2DDiagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Stool3DDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *Stool3DDiagram) {
	// insertion point for pointers field
	if instance.SampledPoints3DShape != nil {
		reference.SampledPoints3DShape = stage.SampledPoints3DShapes_reference[instance.SampledPoints3DShape]
	}
	if instance.Rendered3DShape != nil {
		reference.Rendered3DShape = stage.Rendered3DShapes_reference[instance.Rendered3DShape]
	}
	// insertion point for slice of pointers field
}

func (reference *StoolAbstract) GongReconstructPointersFromReferences(stage *Stage, instance *StoolAbstract) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TiledFloor3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *TiledFloor3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopEndArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopEndArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopEndArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *TopEndArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopEndHalfwayArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopEndHalfwayArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopEndHalfwayArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *TopEndHalfwayArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopGrowthCurve2D) GongReconstructPointersFromReferences(stage *Stage, instance *TopGrowthCurve2D) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopMidArcVectorShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopMidArcVectorShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopMidArcVectorShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *TopMidArcVectorShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStackGrowthCurve2DEndHalfwayArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopStackGrowthCurve2DEndHalfwayArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStackGrowthCurve2DStartHalfwayArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopStackGrowthCurve2DStartHalfwayArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStackOfGrowthCurve2D) GongReconstructPointersFromReferences(stage *Stage, instance *TopStackOfGrowthCurve2D) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStackOfRotatedGrowthCurve2D) GongReconstructPointersFromReferences(stage *Stage, instance *TopStackOfRotatedGrowthCurve2D) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStackOfRotatedGrowthCurve2DEndArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopStackOfRotatedGrowthCurve2DEndArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStackOfRotatedGrowthCurve2DStartArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopStackOfRotatedGrowthCurve2DStartArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStartArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopStartArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStartArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *TopStartArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStartHalfwayArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopStartHalfwayArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStartHalfwayArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *TopStartHalfwayArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Torus3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *Torus3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TorusEdge3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *TorusEdge3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TorusStackShape) GongReconstructPointersFromReferences(stage *Stage, instance *TorusStackShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Vase2DDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *Vase2DDiagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Vase3DDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *Vase3DDiagram) {
	// insertion point for pointers field
	if instance.Rendered3DShape != nil {
		reference.Rendered3DShape = stage.Rendered3DShapes_reference[instance.Rendered3DShape]
	}
	if instance.SampledPoints3DShape != nil {
		reference.SampledPoints3DShape = stage.SampledPoints3DShapes_reference[instance.SampledPoints3DShape]
	}
	if instance.OriginalPoints3DShape != nil {
		reference.OriginalPoints3DShape = stage.OriginalPoints3DShapes_reference[instance.OriginalPoints3DShape]
	}
	if instance.Angle0Shape != nil {
		reference.Angle0Shape = stage.Angle0Shapes_reference[instance.Angle0Shape]
	}
	// insertion point for slice of pointers field
}

func (reference *VaseAbstract) GongReconstructPointersFromReferences(stage *Stage, instance *VaseAbstract) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *VerticalTorusStackShape) GongReconstructPointersFromReferences(stage *Stage, instance *VerticalTorusStackShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *VolumeKey3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *VolumeKey3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Angle0Shape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ArcNormalVectorShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ArcNormalVectorShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AxesShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *BaseVectorShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *BaseVectorShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ChosenP1P2PairShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CircleGridShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Circumference3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Clock2DDiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Clock3DDiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.SampledPoints3DShape; _reference != nil {
		reference.SampledPoints3DShape = nil
		if _instance, ok := stage.SampledPoints3DShapes_instance[_reference]; ok {
			reference.SampledPoints3DShape = _instance
		}
	}
	if _reference := reference.Rendered3DShape; _reference != nil {
		reference.Rendered3DShape = nil
		if _instance, ok := stage.Rendered3DShapes_instance[_reference]; ok {
			reference.Rendered3DShape = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ClockAbstract) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ClockTopCurveShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CutLine3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EndArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EndArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EndHalfwayArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EndHalfwayArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ExplanationTextShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Eye3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EyeCornersSampledPoints3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EyeSampledPoints3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EyeSeatBottomCurveShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EyeStoolBottomCurveShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EyeVolume3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GridPathShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GrowthCurve2D) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GrowthCurve2DRibbon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GrowthCurve2DRibbonEndShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GrowthCurve2DRibbonStartShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GrowthCurveRhombusGridShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GrowthCurveRhombusShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GrowthVectorShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *InitialRhombusGridShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *InitialRhombusShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Key3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *KeyHole3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *KeyHoleShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SubLibraries []*Library
	for _, _reference := range reference.SubLibraries {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibraries = append(_SubLibraries, _instance)
		}
	}
	reference.SubLibraries = _SubLibraries
	var _Plants []*PlantAbstract
	for _, _reference := range reference.Plants {
		if _instance, ok := stage.PlantAbstracts_instance[_reference]; ok {
			_Plants = append(_Plants, _instance)
		}
	}
	reference.Plants = _Plants
}

func (reference *MidArcVectorShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MidArcVectorShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *OriginalPoints3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ParastichyMCurves3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ParastichyNCurves3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyGrowthCurve2DRibbon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyGrowthCurve2DRibbonEndShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyGrowthCurve2DRibbonStartShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyGrowthCurve2DTrajectory) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyGrowthCurve2DTrajectoryP1P2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyGrowthCurve2DTrajectoryShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyRotatedSeatBottomCurveShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyRotatedSeatTopCurveShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartiallyRotatedTorusShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PerpendicularVector) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PerpendicularVectorGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PerpendicularVectorGridHalfway) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PerpendicularVectorHalfway) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Plant2DDiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Plant3DDiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.StemCylinder3DShape; _reference != nil {
		reference.StemCylinder3DShape = nil
		if _instance, ok := stage.StemCylinder3DShapes_instance[_reference]; ok {
			reference.StemCylinder3DShape = _instance
		}
	}
	if _reference := reference.ParastichyNCurves3DShape; _reference != nil {
		reference.ParastichyNCurves3DShape = nil
		if _instance, ok := stage.ParastichyNCurves3DShapes_instance[_reference]; ok {
			reference.ParastichyNCurves3DShape = _instance
		}
	}
	if _reference := reference.ParastichyMCurves3DShape; _reference != nil {
		reference.ParastichyMCurves3DShape = nil
		if _instance, ok := stage.ParastichyMCurves3DShapes_instance[_reference]; ok {
			reference.ParastichyMCurves3DShape = _instance
		}
	}
	if _reference := reference.CutLine3DShape; _reference != nil {
		reference.CutLine3DShape = nil
		if _instance, ok := stage.CutLine3DShapes_instance[_reference]; ok {
			reference.CutLine3DShape = _instance
		}
	}
	if _reference := reference.Circumference3DShape; _reference != nil {
		reference.Circumference3DShape = nil
		if _instance, ok := stage.Circumference3DShapes_instance[_reference]; ok {
			reference.Circumference3DShape = _instance
		}
	}
	if _reference := reference.Rendered3DShape; _reference != nil {
		reference.Rendered3DShape = nil
		if _instance, ok := stage.Rendered3DShapes_instance[_reference]; ok {
			reference.Rendered3DShape = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *PlantAbstract) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.VaseAbstract; _reference != nil {
		reference.VaseAbstract = nil
		if _instance, ok := stage.VaseAbstracts_instance[_reference]; ok {
			reference.VaseAbstract = _instance
		}
	}
	if _reference := reference.StoolAbstract; _reference != nil {
		reference.StoolAbstract = nil
		if _instance, ok := stage.StoolAbstracts_instance[_reference]; ok {
			reference.StoolAbstract = _instance
		}
	}
	if _reference := reference.ClockAbstract; _reference != nil {
		reference.ClockAbstract = nil
		if _instance, ok := stage.ClockAbstracts_instance[_reference]; ok {
			reference.ClockAbstract = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Plant2DDiagrams []*Plant2DDiagram
	for _, _reference := range reference.Plant2DDiagrams {
		if _instance, ok := stage.Plant2DDiagrams_instance[_reference]; ok {
			_Plant2DDiagrams = append(_Plant2DDiagrams, _instance)
		}
	}
	reference.Plant2DDiagrams = _Plant2DDiagrams
	var _Plant3DDiagrams []*Plant3DDiagram
	for _, _reference := range reference.Plant3DDiagrams {
		if _instance, ok := stage.Plant3DDiagrams_instance[_reference]; ok {
			_Plant3DDiagrams = append(_Plant3DDiagrams, _instance)
		}
	}
	reference.Plant3DDiagrams = _Plant3DDiagrams
	var _Vase2DDiagrams []*Vase2DDiagram
	for _, _reference := range reference.Vase2DDiagrams {
		if _instance, ok := stage.Vase2DDiagrams_instance[_reference]; ok {
			_Vase2DDiagrams = append(_Vase2DDiagrams, _instance)
		}
	}
	reference.Vase2DDiagrams = _Vase2DDiagrams
	var _Vase3DDiagrams []*Vase3DDiagram
	for _, _reference := range reference.Vase3DDiagrams {
		if _instance, ok := stage.Vase3DDiagrams_instance[_reference]; ok {
			_Vase3DDiagrams = append(_Vase3DDiagrams, _instance)
		}
	}
	reference.Vase3DDiagrams = _Vase3DDiagrams
	var _Stool2DDiagrams []*Stool2DDiagram
	for _, _reference := range reference.Stool2DDiagrams {
		if _instance, ok := stage.Stool2DDiagrams_instance[_reference]; ok {
			_Stool2DDiagrams = append(_Stool2DDiagrams, _instance)
		}
	}
	reference.Stool2DDiagrams = _Stool2DDiagrams
	var _Stool3DDiagrams []*Stool3DDiagram
	for _, _reference := range reference.Stool3DDiagrams {
		if _instance, ok := stage.Stool3DDiagrams_instance[_reference]; ok {
			_Stool3DDiagrams = append(_Stool3DDiagrams, _instance)
		}
	}
	reference.Stool3DDiagrams = _Stool3DDiagrams
	var _Clock2DDiagrams []*Clock2DDiagram
	for _, _reference := range reference.Clock2DDiagrams {
		if _instance, ok := stage.Clock2DDiagrams_instance[_reference]; ok {
			_Clock2DDiagrams = append(_Clock2DDiagrams, _instance)
		}
	}
	reference.Clock2DDiagrams = _Clock2DDiagrams
	var _Clock3DDiagrams []*Clock3DDiagram
	for _, _reference := range reference.Clock3DDiagrams {
		if _instance, ok := stage.Clock3DDiagrams_instance[_reference]; ok {
			_Clock3DDiagrams = append(_Clock3DDiagrams, _instance)
		}
	}
	reference.Clock3DDiagrams = _Clock3DDiagrams
}

func (reference *PlantCircumferenceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PointsAndLines3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PxShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Rendered3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *RhombusShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *RhombusStuff) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *RotatedRhombusGridShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *RotatedRhombusShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *RotatedSampledPoints3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *RotatedSeatAndLegs3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *SampledPoints3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Seat3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *SeatAndLegs3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *SeatBottomCurveShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *SeatTopCurveShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedBottomTopStartArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedBottomTopStartArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedLeftGrowthCurve2DRibbon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedLeftGrowthCurve2DRibbonEndShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedLeftGrowthCurve2DRibbonStartShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedLeftStackGrowthCurveEndArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedLeftStackGrowthCurveStartArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedLeftStackNormalVector) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedLeftStackOfGrowthCurve) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedLeftStackOfNormalVector) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedRightGrowthCurve2DRibbon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedRightGrowthCurve2DRibbonEndShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ShiftedRightGrowthCurve2DRibbonStartShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackGrowthCurve2DEndHalfwayArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackGrowthCurve2DRibbonEndShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackGrowthCurve2DRibbonStartShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackGrowthCurve2DStartHalfwayArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackOfGrowthCurve2D) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackOfGrowthCurve2DByGrowthVector) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackOfGrowthCurve2DRibbon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackOfPartiallyRotatedTorusShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackOfRotatedGrowthCurve2D) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackOfRotatedGrowthCurve2DRibbon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackRotatedGrowthCurve2DEndArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackRotatedGrowthCurve2DRibbonEndShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackRotatedGrowthCurve2DRibbonStartShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackRotatedGrowthCurve2DStartArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StartArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StartArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StartHalfwayArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StartHalfwayArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StemCylinder3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Stool2DDiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Stool3DDiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.SampledPoints3DShape; _reference != nil {
		reference.SampledPoints3DShape = nil
		if _instance, ok := stage.SampledPoints3DShapes_instance[_reference]; ok {
			reference.SampledPoints3DShape = _instance
		}
	}
	if _reference := reference.Rendered3DShape; _reference != nil {
		reference.Rendered3DShape = nil
		if _instance, ok := stage.Rendered3DShapes_instance[_reference]; ok {
			reference.Rendered3DShape = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *StoolAbstract) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TiledFloor3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopEndArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopEndArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopEndHalfwayArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopEndHalfwayArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopGrowthCurve2D) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopMidArcVectorShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopMidArcVectorShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStackGrowthCurve2DEndHalfwayArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStackGrowthCurve2DStartHalfwayArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStackOfGrowthCurve2D) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStackOfRotatedGrowthCurve2D) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStackOfRotatedGrowthCurve2DEndArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStackOfRotatedGrowthCurve2DStartArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStartArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStartArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStartHalfwayArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStartHalfwayArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Torus3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TorusEdge3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TorusStackShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Vase2DDiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Vase3DDiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Rendered3DShape; _reference != nil {
		reference.Rendered3DShape = nil
		if _instance, ok := stage.Rendered3DShapes_instance[_reference]; ok {
			reference.Rendered3DShape = _instance
		}
	}
	if _reference := reference.SampledPoints3DShape; _reference != nil {
		reference.SampledPoints3DShape = nil
		if _instance, ok := stage.SampledPoints3DShapes_instance[_reference]; ok {
			reference.SampledPoints3DShape = _instance
		}
	}
	if _reference := reference.OriginalPoints3DShape; _reference != nil {
		reference.OriginalPoints3DShape = nil
		if _instance, ok := stage.OriginalPoints3DShapes_instance[_reference]; ok {
			reference.OriginalPoints3DShape = _instance
		}
	}
	if _reference := reference.Angle0Shape; _reference != nil {
		reference.Angle0Shape = nil
		if _instance, ok := stage.Angle0Shapes_instance[_reference]; ok {
			reference.Angle0Shape = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *VaseAbstract) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *VerticalTorusStackShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *VolumeKey3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (angle0shape *Angle0Shape) GongDiff(stage *Stage, angle0shapeOther *Angle0Shape) (diffs []string) {
	// insertion point for field diffs
	if angle0shape.Name != angle0shapeOther.Name {
		diffs = append(diffs, angle0shape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (arcnormalvectorshape *ArcNormalVectorShape) GongDiff(stage *Stage, arcnormalvectorshapeOther *ArcNormalVectorShape) (diffs []string) {
	// insertion point for field diffs
	if arcnormalvectorshape.Name != arcnormalvectorshapeOther.Name {
		diffs = append(diffs, arcnormalvectorshape.GongMarshallField(stage, "Name"))
	}
	if arcnormalvectorshape.StartX != arcnormalvectorshapeOther.StartX {
		diffs = append(diffs, arcnormalvectorshape.GongMarshallField(stage, "StartX"))
	}
	if arcnormalvectorshape.StartY != arcnormalvectorshapeOther.StartY {
		diffs = append(diffs, arcnormalvectorshape.GongMarshallField(stage, "StartY"))
	}
	if arcnormalvectorshape.EndX != arcnormalvectorshapeOther.EndX {
		diffs = append(diffs, arcnormalvectorshape.GongMarshallField(stage, "EndX"))
	}
	if arcnormalvectorshape.EndY != arcnormalvectorshapeOther.EndY {
		diffs = append(diffs, arcnormalvectorshape.GongMarshallField(stage, "EndY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongDiff(stage *Stage, arcnormalvectorshapegridOther *ArcNormalVectorShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if arcnormalvectorshapegrid.Name != arcnormalvectorshapegridOther.Name {
		diffs = append(diffs, arcnormalvectorshapegrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (axesshape *AxesShape) GongDiff(stage *Stage, axesshapeOther *AxesShape) (diffs []string) {
	// insertion point for field diffs
	if axesshape.Name != axesshapeOther.Name {
		diffs = append(diffs, axesshape.GongMarshallField(stage, "Name"))
	}
	if axesshape.LengthX != axesshapeOther.LengthX {
		diffs = append(diffs, axesshape.GongMarshallField(stage, "LengthX"))
	}
	if axesshape.LengthY != axesshapeOther.LengthY {
		diffs = append(diffs, axesshape.GongMarshallField(stage, "LengthY"))
	}
	if axesshape.IsWithHiddenHandle != axesshapeOther.IsWithHiddenHandle {
		diffs = append(diffs, axesshape.GongMarshallField(stage, "IsWithHiddenHandle"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (basevectorshape *BaseVectorShape) GongDiff(stage *Stage, basevectorshapeOther *BaseVectorShape) (diffs []string) {
	// insertion point for field diffs
	if basevectorshape.Name != basevectorshapeOther.Name {
		diffs = append(diffs, basevectorshape.GongMarshallField(stage, "Name"))
	}
	if basevectorshape.StartX != basevectorshapeOther.StartX {
		diffs = append(diffs, basevectorshape.GongMarshallField(stage, "StartX"))
	}
	if basevectorshape.StartY != basevectorshapeOther.StartY {
		diffs = append(diffs, basevectorshape.GongMarshallField(stage, "StartY"))
	}
	if basevectorshape.EndX != basevectorshapeOther.EndX {
		diffs = append(diffs, basevectorshape.GongMarshallField(stage, "EndX"))
	}
	if basevectorshape.EndY != basevectorshapeOther.EndY {
		diffs = append(diffs, basevectorshape.GongMarshallField(stage, "EndY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (basevectorshapegrid *BaseVectorShapeGrid) GongDiff(stage *Stage, basevectorshapegridOther *BaseVectorShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if basevectorshapegrid.Name != basevectorshapegridOther.Name {
		diffs = append(diffs, basevectorshapegrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (chosenp1p2pairshape *ChosenP1P2PairShape) GongDiff(stage *Stage, chosenp1p2pairshapeOther *ChosenP1P2PairShape) (diffs []string) {
	// insertion point for field diffs
	if chosenp1p2pairshape.Name != chosenp1p2pairshapeOther.Name {
		diffs = append(diffs, chosenp1p2pairshape.GongMarshallField(stage, "Name"))
	}
	if chosenp1p2pairshape.P1X != chosenp1p2pairshapeOther.P1X {
		diffs = append(diffs, chosenp1p2pairshape.GongMarshallField(stage, "P1X"))
	}
	if chosenp1p2pairshape.P1Y != chosenp1p2pairshapeOther.P1Y {
		diffs = append(diffs, chosenp1p2pairshape.GongMarshallField(stage, "P1Y"))
	}
	if chosenp1p2pairshape.P2X != chosenp1p2pairshapeOther.P2X {
		diffs = append(diffs, chosenp1p2pairshape.GongMarshallField(stage, "P2X"))
	}
	if chosenp1p2pairshape.P2Y != chosenp1p2pairshapeOther.P2Y {
		diffs = append(diffs, chosenp1p2pairshape.GongMarshallField(stage, "P2Y"))
	}
	if chosenp1p2pairshape.PxX != chosenp1p2pairshapeOther.PxX {
		diffs = append(diffs, chosenp1p2pairshape.GongMarshallField(stage, "PxX"))
	}
	if chosenp1p2pairshape.PxY != chosenp1p2pairshapeOther.PxY {
		diffs = append(diffs, chosenp1p2pairshape.GongMarshallField(stage, "PxY"))
	}
	if chosenp1p2pairshape.DistanceP1Px != chosenp1p2pairshapeOther.DistanceP1Px {
		diffs = append(diffs, chosenp1p2pairshape.GongMarshallField(stage, "DistanceP1Px"))
	}
	if chosenp1p2pairshape.DistanceP2Px != chosenp1p2pairshapeOther.DistanceP2Px {
		diffs = append(diffs, chosenp1p2pairshape.GongMarshallField(stage, "DistanceP2Px"))
	}
	if chosenp1p2pairshape.DistanceSum != chosenp1p2pairshapeOther.DistanceSum {
		diffs = append(diffs, chosenp1p2pairshape.GongMarshallField(stage, "DistanceSum"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (circlegridshape *CircleGridShape) GongDiff(stage *Stage, circlegridshapeOther *CircleGridShape) (diffs []string) {
	// insertion point for field diffs
	if circlegridshape.Name != circlegridshapeOther.Name {
		diffs = append(diffs, circlegridshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (circumference3dshape *Circumference3DShape) GongDiff(stage *Stage, circumference3dshapeOther *Circumference3DShape) (diffs []string) {
	// insertion point for field diffs
	if circumference3dshape.Name != circumference3dshapeOther.Name {
		diffs = append(diffs, circumference3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (clock2ddiagram *Clock2DDiagram) GongDiff(stage *Stage, clock2ddiagramOther *Clock2DDiagram) (diffs []string) {
	// insertion point for field diffs
	if clock2ddiagram.Name != clock2ddiagramOther.Name {
		diffs = append(diffs, clock2ddiagram.GongMarshallField(stage, "Name"))
	}
	if clock2ddiagram.Zoom != clock2ddiagramOther.Zoom {
		diffs = append(diffs, clock2ddiagram.GongMarshallField(stage, "Zoom"))
	}
	if clock2ddiagram.IsHiddenAxesShape != clock2ddiagramOther.IsHiddenAxesShape {
		diffs = append(diffs, clock2ddiagram.GongMarshallField(stage, "IsHiddenAxesShape"))
	}
	if clock2ddiagram.IsChecked != clock2ddiagramOther.IsChecked {
		diffs = append(diffs, clock2ddiagram.GongMarshallField(stage, "IsChecked"))
	}
	if clock2ddiagram.ComputedPrefix != clock2ddiagramOther.ComputedPrefix {
		diffs = append(diffs, clock2ddiagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if clock2ddiagram.IsExpanded != clock2ddiagramOther.IsExpanded {
		diffs = append(diffs, clock2ddiagram.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (clock3ddiagram *Clock3DDiagram) GongDiff(stage *Stage, clock3ddiagramOther *Clock3DDiagram) (diffs []string) {
	// insertion point for field diffs
	if clock3ddiagram.Name != clock3ddiagramOther.Name {
		diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "Name"))
	}
	if clock3ddiagram.IsHiddenClockTopCurveShape != clock3ddiagramOther.IsHiddenClockTopCurveShape {
		diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "IsHiddenClockTopCurveShape"))
	}
	if clock3ddiagram.IsHiddenTorus3DShape != clock3ddiagramOther.IsHiddenTorus3DShape {
		diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "IsHiddenTorus3DShape"))
	}
	if clock3ddiagram.IsHiddenSampledPoints3DShape != clock3ddiagramOther.IsHiddenSampledPoints3DShape {
		diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "IsHiddenSampledPoints3DShape"))
	}
	if (clock3ddiagram.SampledPoints3DShape == nil) != (clock3ddiagramOther.SampledPoints3DShape == nil) {
		diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
	} else if clock3ddiagram.SampledPoints3DShape != nil && clock3ddiagramOther.SampledPoints3DShape != nil {
		if clock3ddiagram.SampledPoints3DShape != clock3ddiagramOther.SampledPoints3DShape {
			diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
		}
	}
	if clock3ddiagram.IsHiddenTiledFloor3DShape != clock3ddiagramOther.IsHiddenTiledFloor3DShape {
		diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "IsHiddenTiledFloor3DShape"))
	}
	if (clock3ddiagram.Rendered3DShape == nil) != (clock3ddiagramOther.Rendered3DShape == nil) {
		diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
	} else if clock3ddiagram.Rendered3DShape != nil && clock3ddiagramOther.Rendered3DShape != nil {
		if clock3ddiagram.Rendered3DShape != clock3ddiagramOther.Rendered3DShape {
			diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
		}
	}
	if clock3ddiagram.IsChecked != clock3ddiagramOther.IsChecked {
		diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "IsChecked"))
	}
	if clock3ddiagram.ComputedPrefix != clock3ddiagramOther.ComputedPrefix {
		diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if clock3ddiagram.IsExpanded != clock3ddiagramOther.IsExpanded {
		diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (clockabstract *ClockAbstract) GongDiff(stage *Stage, clockabstractOther *ClockAbstract) (diffs []string) {
	// insertion point for field diffs
	if clockabstract.Name != clockabstractOther.Name {
		diffs = append(diffs, clockabstract.GongMarshallField(stage, "Name"))
	}
	if clockabstract.RadialRepetitions != clockabstractOther.RadialRepetitions {
		diffs = append(diffs, clockabstract.GongMarshallField(stage, "RadialRepetitions"))
	}
	if clockabstract.Transparency != clockabstractOther.Transparency {
		diffs = append(diffs, clockabstract.GongMarshallField(stage, "Transparency"))
	}
	if clockabstract.RelativeTubeDiameter != clockabstractOther.RelativeTubeDiameter {
		diffs = append(diffs, clockabstract.GongMarshallField(stage, "RelativeTubeDiameter"))
	}
	if clockabstract.RelativeHeight3DTorus != clockabstractOther.RelativeHeight3DTorus {
		diffs = append(diffs, clockabstract.GongMarshallField(stage, "RelativeHeight3DTorus"))
	}
	if clockabstract.ClockTorusVerticalScale != clockabstractOther.ClockTorusVerticalScale {
		diffs = append(diffs, clockabstract.GongMarshallField(stage, "ClockTorusVerticalScale"))
	}
	if clockabstract.RelativeHeight != clockabstractOther.RelativeHeight {
		diffs = append(diffs, clockabstract.GongMarshallField(stage, "RelativeHeight"))
	}
	if clockabstract.ProjectionAngle != clockabstractOther.ProjectionAngle {
		diffs = append(diffs, clockabstract.GongMarshallField(stage, "ProjectionAngle"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (clocktopcurveshape *ClockTopCurveShape) GongDiff(stage *Stage, clocktopcurveshapeOther *ClockTopCurveShape) (diffs []string) {
	// insertion point for field diffs
	if clocktopcurveshape.Name != clocktopcurveshapeOther.Name {
		diffs = append(diffs, clocktopcurveshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (cutline3dshape *CutLine3DShape) GongDiff(stage *Stage, cutline3dshapeOther *CutLine3DShape) (diffs []string) {
	// insertion point for field diffs
	if cutline3dshape.Name != cutline3dshapeOther.Name {
		diffs = append(diffs, cutline3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (endarcshape *EndArcShape) GongDiff(stage *Stage, endarcshapeOther *EndArcShape) (diffs []string) {
	// insertion point for field diffs
	if endarcshape.Name != endarcshapeOther.Name {
		diffs = append(diffs, endarcshape.GongMarshallField(stage, "Name"))
	}
	if endarcshape.StartX != endarcshapeOther.StartX {
		diffs = append(diffs, endarcshape.GongMarshallField(stage, "StartX"))
	}
	if endarcshape.StartY != endarcshapeOther.StartY {
		diffs = append(diffs, endarcshape.GongMarshallField(stage, "StartY"))
	}
	if endarcshape.EndX != endarcshapeOther.EndX {
		diffs = append(diffs, endarcshape.GongMarshallField(stage, "EndX"))
	}
	if endarcshape.EndY != endarcshapeOther.EndY {
		diffs = append(diffs, endarcshape.GongMarshallField(stage, "EndY"))
	}
	if endarcshape.XAxisRotation != endarcshapeOther.XAxisRotation {
		diffs = append(diffs, endarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if endarcshape.LargeArcFlag != endarcshapeOther.LargeArcFlag {
		diffs = append(diffs, endarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if endarcshape.SweepFlag != endarcshapeOther.SweepFlag {
		diffs = append(diffs, endarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if endarcshape.RadiusX != endarcshapeOther.RadiusX {
		diffs = append(diffs, endarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if endarcshape.RadiusY != endarcshapeOther.RadiusY {
		diffs = append(diffs, endarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (endarcshapegrid *EndArcShapeGrid) GongDiff(stage *Stage, endarcshapegridOther *EndArcShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if endarcshapegrid.Name != endarcshapegridOther.Name {
		diffs = append(diffs, endarcshapegrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (endhalfwayarcshape *EndHalfwayArcShape) GongDiff(stage *Stage, endhalfwayarcshapeOther *EndHalfwayArcShape) (diffs []string) {
	// insertion point for field diffs
	if endhalfwayarcshape.Name != endhalfwayarcshapeOther.Name {
		diffs = append(diffs, endhalfwayarcshape.GongMarshallField(stage, "Name"))
	}
	if endhalfwayarcshape.StartX != endhalfwayarcshapeOther.StartX {
		diffs = append(diffs, endhalfwayarcshape.GongMarshallField(stage, "StartX"))
	}
	if endhalfwayarcshape.StartY != endhalfwayarcshapeOther.StartY {
		diffs = append(diffs, endhalfwayarcshape.GongMarshallField(stage, "StartY"))
	}
	if endhalfwayarcshape.EndX != endhalfwayarcshapeOther.EndX {
		diffs = append(diffs, endhalfwayarcshape.GongMarshallField(stage, "EndX"))
	}
	if endhalfwayarcshape.EndY != endhalfwayarcshapeOther.EndY {
		diffs = append(diffs, endhalfwayarcshape.GongMarshallField(stage, "EndY"))
	}
	if endhalfwayarcshape.RadiusX != endhalfwayarcshapeOther.RadiusX {
		diffs = append(diffs, endhalfwayarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if endhalfwayarcshape.RadiusY != endhalfwayarcshapeOther.RadiusY {
		diffs = append(diffs, endhalfwayarcshape.GongMarshallField(stage, "RadiusY"))
	}
	if endhalfwayarcshape.XAxisRotation != endhalfwayarcshapeOther.XAxisRotation {
		diffs = append(diffs, endhalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if endhalfwayarcshape.LargeArcFlag != endhalfwayarcshapeOther.LargeArcFlag {
		diffs = append(diffs, endhalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if endhalfwayarcshape.SweepFlag != endhalfwayarcshapeOther.SweepFlag {
		diffs = append(diffs, endhalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongDiff(stage *Stage, endhalfwayarcshapegridOther *EndHalfwayArcShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if endhalfwayarcshapegrid.Name != endhalfwayarcshapegridOther.Name {
		diffs = append(diffs, endhalfwayarcshapegrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (explanationtextshape *ExplanationTextShape) GongDiff(stage *Stage, explanationtextshapeOther *ExplanationTextShape) (diffs []string) {
	// insertion point for field diffs
	if explanationtextshape.Name != explanationtextshapeOther.Name {
		diffs = append(diffs, explanationtextshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (eye3dshape *Eye3DShape) GongDiff(stage *Stage, eye3dshapeOther *Eye3DShape) (diffs []string) {
	// insertion point for field diffs
	if eye3dshape.Name != eye3dshapeOther.Name {
		diffs = append(diffs, eye3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongDiff(stage *Stage, eyecornerssampledpoints3dshapeOther *EyeCornersSampledPoints3DShape) (diffs []string) {
	// insertion point for field diffs
	if eyecornerssampledpoints3dshape.Name != eyecornerssampledpoints3dshapeOther.Name {
		diffs = append(diffs, eyecornerssampledpoints3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongDiff(stage *Stage, eyesampledpoints3dshapeOther *EyeSampledPoints3DShape) (diffs []string) {
	// insertion point for field diffs
	if eyesampledpoints3dshape.Name != eyesampledpoints3dshapeOther.Name {
		diffs = append(diffs, eyesampledpoints3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongDiff(stage *Stage, eyeseatbottomcurveshapeOther *EyeSeatBottomCurveShape) (diffs []string) {
	// insertion point for field diffs
	if eyeseatbottomcurveshape.Name != eyeseatbottomcurveshapeOther.Name {
		diffs = append(diffs, eyeseatbottomcurveshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongDiff(stage *Stage, eyestoolbottomcurveshapeOther *EyeStoolBottomCurveShape) (diffs []string) {
	// insertion point for field diffs
	if eyestoolbottomcurveshape.Name != eyestoolbottomcurveshapeOther.Name {
		diffs = append(diffs, eyestoolbottomcurveshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (eyevolume3dshape *EyeVolume3DShape) GongDiff(stage *Stage, eyevolume3dshapeOther *EyeVolume3DShape) (diffs []string) {
	// insertion point for field diffs
	if eyevolume3dshape.Name != eyevolume3dshapeOther.Name {
		diffs = append(diffs, eyevolume3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gridpathshape *GridPathShape) GongDiff(stage *Stage, gridpathshapeOther *GridPathShape) (diffs []string) {
	// insertion point for field diffs
	if gridpathshape.Name != gridpathshapeOther.Name {
		diffs = append(diffs, gridpathshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (growthcurve2d *GrowthCurve2D) GongDiff(stage *Stage, growthcurve2dOther *GrowthCurve2D) (diffs []string) {
	// insertion point for field diffs
	if growthcurve2d.Name != growthcurve2dOther.Name {
		diffs = append(diffs, growthcurve2d.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (growthcurve2dribbon *GrowthCurve2DRibbon) GongDiff(stage *Stage, growthcurve2dribbonOther *GrowthCurve2DRibbon) (diffs []string) {
	// insertion point for field diffs
	if growthcurve2dribbon.Name != growthcurve2dribbonOther.Name {
		diffs = append(diffs, growthcurve2dribbon.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongDiff(stage *Stage, growthcurve2dribbonendshapeOther *GrowthCurve2DRibbonEndShape) (diffs []string) {
	// insertion point for field diffs
	if growthcurve2dribbonendshape.Name != growthcurve2dribbonendshapeOther.Name {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
	}
	if growthcurve2dribbonendshape.BottomStartX != growthcurve2dribbonendshapeOther.BottomStartX {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
	}
	if growthcurve2dribbonendshape.BottomStartY != growthcurve2dribbonendshapeOther.BottomStartY {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
	}
	if growthcurve2dribbonendshape.BottomEndX != growthcurve2dribbonendshapeOther.BottomEndX {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
	}
	if growthcurve2dribbonendshape.BottomEndY != growthcurve2dribbonendshapeOther.BottomEndY {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
	}
	if growthcurve2dribbonendshape.BottomRadiusX != growthcurve2dribbonendshapeOther.BottomRadiusX {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if growthcurve2dribbonendshape.BottomRadiusY != growthcurve2dribbonendshapeOther.BottomRadiusY {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if growthcurve2dribbonendshape.BottomXAxisRotation != growthcurve2dribbonendshapeOther.BottomXAxisRotation {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if growthcurve2dribbonendshape.BottomLargeArcFlag != growthcurve2dribbonendshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if growthcurve2dribbonendshape.BottomSweepFlag != growthcurve2dribbonendshapeOther.BottomSweepFlag {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if growthcurve2dribbonendshape.TopStartX != growthcurve2dribbonendshapeOther.TopStartX {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
	}
	if growthcurve2dribbonendshape.TopStartY != growthcurve2dribbonendshapeOther.TopStartY {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
	}
	if growthcurve2dribbonendshape.TopEndX != growthcurve2dribbonendshapeOther.TopEndX {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
	}
	if growthcurve2dribbonendshape.TopEndY != growthcurve2dribbonendshapeOther.TopEndY {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
	}
	if growthcurve2dribbonendshape.TopRadiusX != growthcurve2dribbonendshapeOther.TopRadiusX {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if growthcurve2dribbonendshape.TopRadiusY != growthcurve2dribbonendshapeOther.TopRadiusY {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if growthcurve2dribbonendshape.TopXAxisRotation != growthcurve2dribbonendshapeOther.TopXAxisRotation {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if growthcurve2dribbonendshape.TopLargeArcFlag != growthcurve2dribbonendshapeOther.TopLargeArcFlag {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if growthcurve2dribbonendshape.TopSweepFlag != growthcurve2dribbonendshapeOther.TopSweepFlag {
		diffs = append(diffs, growthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongDiff(stage *Stage, growthcurve2dribbonstartshapeOther *GrowthCurve2DRibbonStartShape) (diffs []string) {
	// insertion point for field diffs
	if growthcurve2dribbonstartshape.Name != growthcurve2dribbonstartshapeOther.Name {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
	}
	if growthcurve2dribbonstartshape.BottomStartX != growthcurve2dribbonstartshapeOther.BottomStartX {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
	}
	if growthcurve2dribbonstartshape.BottomStartY != growthcurve2dribbonstartshapeOther.BottomStartY {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
	}
	if growthcurve2dribbonstartshape.BottomEndX != growthcurve2dribbonstartshapeOther.BottomEndX {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
	}
	if growthcurve2dribbonstartshape.BottomEndY != growthcurve2dribbonstartshapeOther.BottomEndY {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
	}
	if growthcurve2dribbonstartshape.BottomRadiusX != growthcurve2dribbonstartshapeOther.BottomRadiusX {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if growthcurve2dribbonstartshape.BottomRadiusY != growthcurve2dribbonstartshapeOther.BottomRadiusY {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if growthcurve2dribbonstartshape.BottomXAxisRotation != growthcurve2dribbonstartshapeOther.BottomXAxisRotation {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if growthcurve2dribbonstartshape.BottomLargeArcFlag != growthcurve2dribbonstartshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if growthcurve2dribbonstartshape.BottomSweepFlag != growthcurve2dribbonstartshapeOther.BottomSweepFlag {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if growthcurve2dribbonstartshape.TopStartX != growthcurve2dribbonstartshapeOther.TopStartX {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
	}
	if growthcurve2dribbonstartshape.TopStartY != growthcurve2dribbonstartshapeOther.TopStartY {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
	}
	if growthcurve2dribbonstartshape.TopEndX != growthcurve2dribbonstartshapeOther.TopEndX {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
	}
	if growthcurve2dribbonstartshape.TopEndY != growthcurve2dribbonstartshapeOther.TopEndY {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
	}
	if growthcurve2dribbonstartshape.TopRadiusX != growthcurve2dribbonstartshapeOther.TopRadiusX {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if growthcurve2dribbonstartshape.TopRadiusY != growthcurve2dribbonstartshapeOther.TopRadiusY {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if growthcurve2dribbonstartshape.TopXAxisRotation != growthcurve2dribbonstartshapeOther.TopXAxisRotation {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if growthcurve2dribbonstartshape.TopLargeArcFlag != growthcurve2dribbonstartshapeOther.TopLargeArcFlag {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if growthcurve2dribbonstartshape.TopSweepFlag != growthcurve2dribbonstartshapeOther.TopSweepFlag {
		diffs = append(diffs, growthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongDiff(stage *Stage, growthcurverhombusgridshapeOther *GrowthCurveRhombusGridShape) (diffs []string) {
	// insertion point for field diffs
	if growthcurverhombusgridshape.Name != growthcurverhombusgridshapeOther.Name {
		diffs = append(diffs, growthcurverhombusgridshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (growthcurverhombusshape *GrowthCurveRhombusShape) GongDiff(stage *Stage, growthcurverhombusshapeOther *GrowthCurveRhombusShape) (diffs []string) {
	// insertion point for field diffs
	if growthcurverhombusshape.Name != growthcurverhombusshapeOther.Name {
		diffs = append(diffs, growthcurverhombusshape.GongMarshallField(stage, "Name"))
	}
	if growthcurverhombusshape.X != growthcurverhombusshapeOther.X {
		diffs = append(diffs, growthcurverhombusshape.GongMarshallField(stage, "X"))
	}
	if growthcurverhombusshape.Y != growthcurverhombusshapeOther.Y {
		diffs = append(diffs, growthcurverhombusshape.GongMarshallField(stage, "Y"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (growthvectorshape *GrowthVectorShape) GongDiff(stage *Stage, growthvectorshapeOther *GrowthVectorShape) (diffs []string) {
	// insertion point for field diffs
	if growthvectorshape.Name != growthvectorshapeOther.Name {
		diffs = append(diffs, growthvectorshape.GongMarshallField(stage, "Name"))
	}
	if growthvectorshape.X != growthvectorshapeOther.X {
		diffs = append(diffs, growthvectorshape.GongMarshallField(stage, "X"))
	}
	if growthvectorshape.Y != growthvectorshapeOther.Y {
		diffs = append(diffs, growthvectorshape.GongMarshallField(stage, "Y"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (initialrhombusgridshape *InitialRhombusGridShape) GongDiff(stage *Stage, initialrhombusgridshapeOther *InitialRhombusGridShape) (diffs []string) {
	// insertion point for field diffs
	if initialrhombusgridshape.Name != initialrhombusgridshapeOther.Name {
		diffs = append(diffs, initialrhombusgridshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (initialrhombusshape *InitialRhombusShape) GongDiff(stage *Stage, initialrhombusshapeOther *InitialRhombusShape) (diffs []string) {
	// insertion point for field diffs
	if initialrhombusshape.Name != initialrhombusshapeOther.Name {
		diffs = append(diffs, initialrhombusshape.GongMarshallField(stage, "Name"))
	}
	if initialrhombusshape.X != initialrhombusshapeOther.X {
		diffs = append(diffs, initialrhombusshape.GongMarshallField(stage, "X"))
	}
	if initialrhombusshape.Y != initialrhombusshapeOther.Y {
		diffs = append(diffs, initialrhombusshape.GongMarshallField(stage, "Y"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (key3dshape *Key3DShape) GongDiff(stage *Stage, key3dshapeOther *Key3DShape) (diffs []string) {
	// insertion point for field diffs
	if key3dshape.Name != key3dshapeOther.Name {
		diffs = append(diffs, key3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (keyhole3dshape *KeyHole3DShape) GongDiff(stage *Stage, keyhole3dshapeOther *KeyHole3DShape) (diffs []string) {
	// insertion point for field diffs
	if keyhole3dshape.Name != keyhole3dshapeOther.Name {
		diffs = append(diffs, keyhole3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (keyholeshape *KeyHoleShape) GongDiff(stage *Stage, keyholeshapeOther *KeyHoleShape) (diffs []string) {
	// insertion point for field diffs
	if keyholeshape.Name != keyholeshapeOther.Name {
		diffs = append(diffs, keyholeshape.GongMarshallField(stage, "Name"))
	}
	if keyholeshape.X != keyholeshapeOther.X {
		diffs = append(diffs, keyholeshape.GongMarshallField(stage, "X"))
	}
	if keyholeshape.Y != keyholeshapeOther.Y {
		diffs = append(diffs, keyholeshape.GongMarshallField(stage, "Y"))
	}
	if keyholeshape.Width != keyholeshapeOther.Width {
		diffs = append(diffs, keyholeshape.GongMarshallField(stage, "Width"))
	}
	if keyholeshape.Height != keyholeshapeOther.Height {
		diffs = append(diffs, keyholeshape.GongMarshallField(stage, "Height"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (library *Library) GongDiff(stage *Stage, libraryOther *Library) (diffs []string) {
	// insertion point for field diffs
	if library.Name != libraryOther.Name {
		diffs = append(diffs, library.GongMarshallField(stage, "Name"))
	}
	SubLibrariesDifferent := false
	if len(library.SubLibraries) != len(libraryOther.SubLibraries) {
		SubLibrariesDifferent = true
	} else {
		for i := range library.SubLibraries {
			if (library.SubLibraries[i] == nil) != (libraryOther.SubLibraries[i] == nil) {
				SubLibrariesDifferent = true
				break
			} else if library.SubLibraries[i] != nil && libraryOther.SubLibraries[i] != nil {
				// this is a pointer comparaison
				if library.SubLibraries[i] != libraryOther.SubLibraries[i] {
					SubLibrariesDifferent = true
					break
				}
			}
		}
	}
	if SubLibrariesDifferent {
		ops := Diff(stage, library, libraryOther, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries)
		diffs = append(diffs, ops)
	}
	if library.NbPixPerCharacter != libraryOther.NbPixPerCharacter {
		diffs = append(diffs, library.GongMarshallField(stage, "NbPixPerCharacter"))
	}
	if library.LogoSVGFile != libraryOther.LogoSVGFile {
		diffs = append(diffs, library.GongMarshallField(stage, "LogoSVGFile"))
	}
	if library.ComputedPrefix != libraryOther.ComputedPrefix {
		diffs = append(diffs, library.GongMarshallField(stage, "ComputedPrefix"))
	}
	if library.IsExpanded != libraryOther.IsExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpanded"))
	}
	if library.IsRootLibrary != libraryOther.IsRootLibrary {
		diffs = append(diffs, library.GongMarshallField(stage, "IsRootLibrary"))
	}
	PlantsDifferent := false
	if len(library.Plants) != len(libraryOther.Plants) {
		PlantsDifferent = true
	} else {
		for i := range library.Plants {
			if (library.Plants[i] == nil) != (libraryOther.Plants[i] == nil) {
				PlantsDifferent = true
				break
			} else if library.Plants[i] != nil && libraryOther.Plants[i] != nil {
				// this is a pointer comparaison
				if library.Plants[i] != libraryOther.Plants[i] {
					PlantsDifferent = true
					break
				}
			}
		}
	}
	if PlantsDifferent {
		ops := Diff(stage, library, libraryOther, "Plants", libraryOther.Plants, library.Plants)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (midarcvectorshape *MidArcVectorShape) GongDiff(stage *Stage, midarcvectorshapeOther *MidArcVectorShape) (diffs []string) {
	// insertion point for field diffs
	if midarcvectorshape.Name != midarcvectorshapeOther.Name {
		diffs = append(diffs, midarcvectorshape.GongMarshallField(stage, "Name"))
	}
	if midarcvectorshape.StartX != midarcvectorshapeOther.StartX {
		diffs = append(diffs, midarcvectorshape.GongMarshallField(stage, "StartX"))
	}
	if midarcvectorshape.StartY != midarcvectorshapeOther.StartY {
		diffs = append(diffs, midarcvectorshape.GongMarshallField(stage, "StartY"))
	}
	if midarcvectorshape.EndX != midarcvectorshapeOther.EndX {
		diffs = append(diffs, midarcvectorshape.GongMarshallField(stage, "EndX"))
	}
	if midarcvectorshape.EndY != midarcvectorshapeOther.EndY {
		diffs = append(diffs, midarcvectorshape.GongMarshallField(stage, "EndY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongDiff(stage *Stage, midarcvectorshapegridOther *MidArcVectorShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if midarcvectorshapegrid.Name != midarcvectorshapegridOther.Name {
		diffs = append(diffs, midarcvectorshapegrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (originalpoints3dshape *OriginalPoints3DShape) GongDiff(stage *Stage, originalpoints3dshapeOther *OriginalPoints3DShape) (diffs []string) {
	// insertion point for field diffs
	if originalpoints3dshape.Name != originalpoints3dshapeOther.Name {
		diffs = append(diffs, originalpoints3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongDiff(stage *Stage, parastichymcurves3dshapeOther *ParastichyMCurves3DShape) (diffs []string) {
	// insertion point for field diffs
	if parastichymcurves3dshape.Name != parastichymcurves3dshapeOther.Name {
		diffs = append(diffs, parastichymcurves3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongDiff(stage *Stage, parastichyncurves3dshapeOther *ParastichyNCurves3DShape) (diffs []string) {
	// insertion point for field diffs
	if parastichyncurves3dshape.Name != parastichyncurves3dshapeOther.Name {
		diffs = append(diffs, parastichyncurves3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongDiff(stage *Stage, partiallygrowthcurve2dribbonOther *PartiallyGrowthCurve2DRibbon) (diffs []string) {
	// insertion point for field diffs
	if partiallygrowthcurve2dribbon.Name != partiallygrowthcurve2dribbonOther.Name {
		diffs = append(diffs, partiallygrowthcurve2dribbon.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongDiff(stage *Stage, partiallygrowthcurve2dribbonendshapeOther *PartiallyGrowthCurve2DRibbonEndShape) (diffs []string) {
	// insertion point for field diffs
	if partiallygrowthcurve2dribbonendshape.Name != partiallygrowthcurve2dribbonendshapeOther.Name {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
	}
	if partiallygrowthcurve2dribbonendshape.BottomStartX != partiallygrowthcurve2dribbonendshapeOther.BottomStartX {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
	}
	if partiallygrowthcurve2dribbonendshape.BottomStartY != partiallygrowthcurve2dribbonendshapeOther.BottomStartY {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
	}
	if partiallygrowthcurve2dribbonendshape.BottomEndX != partiallygrowthcurve2dribbonendshapeOther.BottomEndX {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
	}
	if partiallygrowthcurve2dribbonendshape.BottomEndY != partiallygrowthcurve2dribbonendshapeOther.BottomEndY {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
	}
	if partiallygrowthcurve2dribbonendshape.BottomRadiusX != partiallygrowthcurve2dribbonendshapeOther.BottomRadiusX {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if partiallygrowthcurve2dribbonendshape.BottomRadiusY != partiallygrowthcurve2dribbonendshapeOther.BottomRadiusY {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if partiallygrowthcurve2dribbonendshape.BottomXAxisRotation != partiallygrowthcurve2dribbonendshapeOther.BottomXAxisRotation {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if partiallygrowthcurve2dribbonendshape.BottomLargeArcFlag != partiallygrowthcurve2dribbonendshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if partiallygrowthcurve2dribbonendshape.BottomSweepFlag != partiallygrowthcurve2dribbonendshapeOther.BottomSweepFlag {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if partiallygrowthcurve2dribbonendshape.TopStartX != partiallygrowthcurve2dribbonendshapeOther.TopStartX {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
	}
	if partiallygrowthcurve2dribbonendshape.TopStartY != partiallygrowthcurve2dribbonendshapeOther.TopStartY {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
	}
	if partiallygrowthcurve2dribbonendshape.TopEndX != partiallygrowthcurve2dribbonendshapeOther.TopEndX {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
	}
	if partiallygrowthcurve2dribbonendshape.TopEndY != partiallygrowthcurve2dribbonendshapeOther.TopEndY {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
	}
	if partiallygrowthcurve2dribbonendshape.TopRadiusX != partiallygrowthcurve2dribbonendshapeOther.TopRadiusX {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if partiallygrowthcurve2dribbonendshape.TopRadiusY != partiallygrowthcurve2dribbonendshapeOther.TopRadiusY {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if partiallygrowthcurve2dribbonendshape.TopXAxisRotation != partiallygrowthcurve2dribbonendshapeOther.TopXAxisRotation {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if partiallygrowthcurve2dribbonendshape.TopLargeArcFlag != partiallygrowthcurve2dribbonendshapeOther.TopLargeArcFlag {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if partiallygrowthcurve2dribbonendshape.TopSweepFlag != partiallygrowthcurve2dribbonendshapeOther.TopSweepFlag {
		diffs = append(diffs, partiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongDiff(stage *Stage, partiallygrowthcurve2dribbonstartshapeOther *PartiallyGrowthCurve2DRibbonStartShape) (diffs []string) {
	// insertion point for field diffs
	if partiallygrowthcurve2dribbonstartshape.Name != partiallygrowthcurve2dribbonstartshapeOther.Name {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
	}
	if partiallygrowthcurve2dribbonstartshape.BottomStartX != partiallygrowthcurve2dribbonstartshapeOther.BottomStartX {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
	}
	if partiallygrowthcurve2dribbonstartshape.BottomStartY != partiallygrowthcurve2dribbonstartshapeOther.BottomStartY {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
	}
	if partiallygrowthcurve2dribbonstartshape.BottomEndX != partiallygrowthcurve2dribbonstartshapeOther.BottomEndX {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
	}
	if partiallygrowthcurve2dribbonstartshape.BottomEndY != partiallygrowthcurve2dribbonstartshapeOther.BottomEndY {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
	}
	if partiallygrowthcurve2dribbonstartshape.BottomRadiusX != partiallygrowthcurve2dribbonstartshapeOther.BottomRadiusX {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if partiallygrowthcurve2dribbonstartshape.BottomRadiusY != partiallygrowthcurve2dribbonstartshapeOther.BottomRadiusY {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if partiallygrowthcurve2dribbonstartshape.BottomXAxisRotation != partiallygrowthcurve2dribbonstartshapeOther.BottomXAxisRotation {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if partiallygrowthcurve2dribbonstartshape.BottomLargeArcFlag != partiallygrowthcurve2dribbonstartshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if partiallygrowthcurve2dribbonstartshape.BottomSweepFlag != partiallygrowthcurve2dribbonstartshapeOther.BottomSweepFlag {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if partiallygrowthcurve2dribbonstartshape.TopStartX != partiallygrowthcurve2dribbonstartshapeOther.TopStartX {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
	}
	if partiallygrowthcurve2dribbonstartshape.TopStartY != partiallygrowthcurve2dribbonstartshapeOther.TopStartY {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
	}
	if partiallygrowthcurve2dribbonstartshape.TopEndX != partiallygrowthcurve2dribbonstartshapeOther.TopEndX {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
	}
	if partiallygrowthcurve2dribbonstartshape.TopEndY != partiallygrowthcurve2dribbonstartshapeOther.TopEndY {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
	}
	if partiallygrowthcurve2dribbonstartshape.TopRadiusX != partiallygrowthcurve2dribbonstartshapeOther.TopRadiusX {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if partiallygrowthcurve2dribbonstartshape.TopRadiusY != partiallygrowthcurve2dribbonstartshapeOther.TopRadiusY {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if partiallygrowthcurve2dribbonstartshape.TopXAxisRotation != partiallygrowthcurve2dribbonstartshapeOther.TopXAxisRotation {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if partiallygrowthcurve2dribbonstartshape.TopLargeArcFlag != partiallygrowthcurve2dribbonstartshapeOther.TopLargeArcFlag {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if partiallygrowthcurve2dribbonstartshape.TopSweepFlag != partiallygrowthcurve2dribbonstartshapeOther.TopSweepFlag {
		diffs = append(diffs, partiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongDiff(stage *Stage, partiallygrowthcurve2dtrajectoryOther *PartiallyGrowthCurve2DTrajectory) (diffs []string) {
	// insertion point for field diffs
	if partiallygrowthcurve2dtrajectory.Name != partiallygrowthcurve2dtrajectoryOther.Name {
		diffs = append(diffs, partiallygrowthcurve2dtrajectory.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongDiff(stage *Stage, partiallygrowthcurve2dtrajectoryp1curveshapeOther *PartiallyGrowthCurve2DTrajectoryP1CurveShape) (diffs []string) {
	// insertion point for field diffs
	if partiallygrowthcurve2dtrajectoryp1curveshape.Name != partiallygrowthcurve2dtrajectoryp1curveshapeOther.Name {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1curveshape.GongMarshallField(stage, "Name"))
	}
	if partiallygrowthcurve2dtrajectoryp1curveshape.StartX != partiallygrowthcurve2dtrajectoryp1curveshapeOther.StartX {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1curveshape.GongMarshallField(stage, "StartX"))
	}
	if partiallygrowthcurve2dtrajectoryp1curveshape.StartY != partiallygrowthcurve2dtrajectoryp1curveshapeOther.StartY {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1curveshape.GongMarshallField(stage, "StartY"))
	}
	if partiallygrowthcurve2dtrajectoryp1curveshape.EndX != partiallygrowthcurve2dtrajectoryp1curveshapeOther.EndX {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1curveshape.GongMarshallField(stage, "EndX"))
	}
	if partiallygrowthcurve2dtrajectoryp1curveshape.EndY != partiallygrowthcurve2dtrajectoryp1curveshapeOther.EndY {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1curveshape.GongMarshallField(stage, "EndY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongDiff(stage *Stage, partiallygrowthcurve2dtrajectoryp1p2Other *PartiallyGrowthCurve2DTrajectoryP1P2) (diffs []string) {
	// insertion point for field diffs
	if partiallygrowthcurve2dtrajectoryp1p2.Name != partiallygrowthcurve2dtrajectoryp1p2Other.Name {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1p2.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongDiff(stage *Stage, partiallygrowthcurve2dtrajectoryp1p2pairlineshapeOther *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) (diffs []string) {
	// insertion point for field diffs
	if partiallygrowthcurve2dtrajectoryp1p2pairlineshape.Name != partiallygrowthcurve2dtrajectoryp1p2pairlineshapeOther.Name {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongMarshallField(stage, "Name"))
	}
	if partiallygrowthcurve2dtrajectoryp1p2pairlineshape.StartX != partiallygrowthcurve2dtrajectoryp1p2pairlineshapeOther.StartX {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongMarshallField(stage, "StartX"))
	}
	if partiallygrowthcurve2dtrajectoryp1p2pairlineshape.StartY != partiallygrowthcurve2dtrajectoryp1p2pairlineshapeOther.StartY {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongMarshallField(stage, "StartY"))
	}
	if partiallygrowthcurve2dtrajectoryp1p2pairlineshape.EndX != partiallygrowthcurve2dtrajectoryp1p2pairlineshapeOther.EndX {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongMarshallField(stage, "EndX"))
	}
	if partiallygrowthcurve2dtrajectoryp1p2pairlineshape.EndY != partiallygrowthcurve2dtrajectoryp1p2pairlineshapeOther.EndY {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongMarshallField(stage, "EndY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongDiff(stage *Stage, partiallygrowthcurve2dtrajectoryp1pointshapeOther *PartiallyGrowthCurve2DTrajectoryP1PointShape) (diffs []string) {
	// insertion point for field diffs
	if partiallygrowthcurve2dtrajectoryp1pointshape.Name != partiallygrowthcurve2dtrajectoryp1pointshapeOther.Name {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1pointshape.GongMarshallField(stage, "Name"))
	}
	if partiallygrowthcurve2dtrajectoryp1pointshape.X != partiallygrowthcurve2dtrajectoryp1pointshapeOther.X {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1pointshape.GongMarshallField(stage, "X"))
	}
	if partiallygrowthcurve2dtrajectoryp1pointshape.Y != partiallygrowthcurve2dtrajectoryp1pointshapeOther.Y {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp1pointshape.GongMarshallField(stage, "Y"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongDiff(stage *Stage, partiallygrowthcurve2dtrajectoryp2curveshapeOther *PartiallyGrowthCurve2DTrajectoryP2CurveShape) (diffs []string) {
	// insertion point for field diffs
	if partiallygrowthcurve2dtrajectoryp2curveshape.Name != partiallygrowthcurve2dtrajectoryp2curveshapeOther.Name {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp2curveshape.GongMarshallField(stage, "Name"))
	}
	if partiallygrowthcurve2dtrajectoryp2curveshape.StartX != partiallygrowthcurve2dtrajectoryp2curveshapeOther.StartX {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp2curveshape.GongMarshallField(stage, "StartX"))
	}
	if partiallygrowthcurve2dtrajectoryp2curveshape.StartY != partiallygrowthcurve2dtrajectoryp2curveshapeOther.StartY {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp2curveshape.GongMarshallField(stage, "StartY"))
	}
	if partiallygrowthcurve2dtrajectoryp2curveshape.EndX != partiallygrowthcurve2dtrajectoryp2curveshapeOther.EndX {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp2curveshape.GongMarshallField(stage, "EndX"))
	}
	if partiallygrowthcurve2dtrajectoryp2curveshape.EndY != partiallygrowthcurve2dtrajectoryp2curveshapeOther.EndY {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp2curveshape.GongMarshallField(stage, "EndY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongDiff(stage *Stage, partiallygrowthcurve2dtrajectoryp2pointshapeOther *PartiallyGrowthCurve2DTrajectoryP2PointShape) (diffs []string) {
	// insertion point for field diffs
	if partiallygrowthcurve2dtrajectoryp2pointshape.Name != partiallygrowthcurve2dtrajectoryp2pointshapeOther.Name {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp2pointshape.GongMarshallField(stage, "Name"))
	}
	if partiallygrowthcurve2dtrajectoryp2pointshape.X != partiallygrowthcurve2dtrajectoryp2pointshapeOther.X {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp2pointshape.GongMarshallField(stage, "X"))
	}
	if partiallygrowthcurve2dtrajectoryp2pointshape.Y != partiallygrowthcurve2dtrajectoryp2pointshapeOther.Y {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryp2pointshape.GongMarshallField(stage, "Y"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongDiff(stage *Stage, partiallygrowthcurve2dtrajectoryshapeOther *PartiallyGrowthCurve2DTrajectoryShape) (diffs []string) {
	// insertion point for field diffs
	if partiallygrowthcurve2dtrajectoryshape.Name != partiallygrowthcurve2dtrajectoryshapeOther.Name {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryshape.GongMarshallField(stage, "Name"))
	}
	if partiallygrowthcurve2dtrajectoryshape.StartX != partiallygrowthcurve2dtrajectoryshapeOther.StartX {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryshape.GongMarshallField(stage, "StartX"))
	}
	if partiallygrowthcurve2dtrajectoryshape.StartY != partiallygrowthcurve2dtrajectoryshapeOther.StartY {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryshape.GongMarshallField(stage, "StartY"))
	}
	if partiallygrowthcurve2dtrajectoryshape.EndX != partiallygrowthcurve2dtrajectoryshapeOther.EndX {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryshape.GongMarshallField(stage, "EndX"))
	}
	if partiallygrowthcurve2dtrajectoryshape.EndY != partiallygrowthcurve2dtrajectoryshapeOther.EndY {
		diffs = append(diffs, partiallygrowthcurve2dtrajectoryshape.GongMarshallField(stage, "EndY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongDiff(stage *Stage, partiallyrotatedseatbottomcurveshapeOther *PartiallyRotatedSeatBottomCurveShape) (diffs []string) {
	// insertion point for field diffs
	if partiallyrotatedseatbottomcurveshape.Name != partiallyrotatedseatbottomcurveshapeOther.Name {
		diffs = append(diffs, partiallyrotatedseatbottomcurveshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongDiff(stage *Stage, partiallyrotatedseattopcurveshapeOther *PartiallyRotatedSeatTopCurveShape) (diffs []string) {
	// insertion point for field diffs
	if partiallyrotatedseattopcurveshape.Name != partiallyrotatedseattopcurveshapeOther.Name {
		diffs = append(diffs, partiallyrotatedseattopcurveshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongDiff(stage *Stage, partiallyrotatedtorusshapeOther *PartiallyRotatedTorusShape) (diffs []string) {
	// insertion point for field diffs
	if partiallyrotatedtorusshape.Name != partiallyrotatedtorusshapeOther.Name {
		diffs = append(diffs, partiallyrotatedtorusshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (perpendicularvector *PerpendicularVector) GongDiff(stage *Stage, perpendicularvectorOther *PerpendicularVector) (diffs []string) {
	// insertion point for field diffs
	if perpendicularvector.Name != perpendicularvectorOther.Name {
		diffs = append(diffs, perpendicularvector.GongMarshallField(stage, "Name"))
	}
	if perpendicularvector.StartX != perpendicularvectorOther.StartX {
		diffs = append(diffs, perpendicularvector.GongMarshallField(stage, "StartX"))
	}
	if perpendicularvector.StartY != perpendicularvectorOther.StartY {
		diffs = append(diffs, perpendicularvector.GongMarshallField(stage, "StartY"))
	}
	if perpendicularvector.EndX != perpendicularvectorOther.EndX {
		diffs = append(diffs, perpendicularvector.GongMarshallField(stage, "EndX"))
	}
	if perpendicularvector.EndY != perpendicularvectorOther.EndY {
		diffs = append(diffs, perpendicularvector.GongMarshallField(stage, "EndY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (perpendicularvectorgrid *PerpendicularVectorGrid) GongDiff(stage *Stage, perpendicularvectorgridOther *PerpendicularVectorGrid) (diffs []string) {
	// insertion point for field diffs
	if perpendicularvectorgrid.Name != perpendicularvectorgridOther.Name {
		diffs = append(diffs, perpendicularvectorgrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongDiff(stage *Stage, perpendicularvectorgridhalfwayOther *PerpendicularVectorGridHalfway) (diffs []string) {
	// insertion point for field diffs
	if perpendicularvectorgridhalfway.Name != perpendicularvectorgridhalfwayOther.Name {
		diffs = append(diffs, perpendicularvectorgridhalfway.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongDiff(stage *Stage, perpendicularvectorhalfwayOther *PerpendicularVectorHalfway) (diffs []string) {
	// insertion point for field diffs
	if perpendicularvectorhalfway.Name != perpendicularvectorhalfwayOther.Name {
		diffs = append(diffs, perpendicularvectorhalfway.GongMarshallField(stage, "Name"))
	}
	if perpendicularvectorhalfway.StartX != perpendicularvectorhalfwayOther.StartX {
		diffs = append(diffs, perpendicularvectorhalfway.GongMarshallField(stage, "StartX"))
	}
	if perpendicularvectorhalfway.StartY != perpendicularvectorhalfwayOther.StartY {
		diffs = append(diffs, perpendicularvectorhalfway.GongMarshallField(stage, "StartY"))
	}
	if perpendicularvectorhalfway.EndX != perpendicularvectorhalfwayOther.EndX {
		diffs = append(diffs, perpendicularvectorhalfway.GongMarshallField(stage, "EndX"))
	}
	if perpendicularvectorhalfway.EndY != perpendicularvectorhalfwayOther.EndY {
		diffs = append(diffs, perpendicularvectorhalfway.GongMarshallField(stage, "EndY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (plant2ddiagram *Plant2DDiagram) GongDiff(stage *Stage, plant2ddiagramOther *Plant2DDiagram) (diffs []string) {
	// insertion point for field diffs
	if plant2ddiagram.Name != plant2ddiagramOther.Name {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "Name"))
	}
	if plant2ddiagram.OriginX != plant2ddiagramOther.OriginX {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "OriginX"))
	}
	if plant2ddiagram.OriginY != plant2ddiagramOther.OriginY {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "OriginY"))
	}
	if plant2ddiagram.Zoom != plant2ddiagramOther.Zoom {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "Zoom"))
	}
	if plant2ddiagram.IsRhombusNodesExpanded != plant2ddiagramOther.IsRhombusNodesExpanded {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsRhombusNodesExpanded"))
	}
	if plant2ddiagram.IsArcNodesExpanded != plant2ddiagramOther.IsArcNodesExpanded {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsArcNodesExpanded"))
	}
	if plant2ddiagram.IsHiddenAxesShape != plant2ddiagramOther.IsHiddenAxesShape {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenAxesShape"))
	}
	if plant2ddiagram.IsHiddenReferenceRhombus != plant2ddiagramOther.IsHiddenReferenceRhombus {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenReferenceRhombus"))
	}
	if plant2ddiagram.IsHiddenPlantCircumferenceShape != plant2ddiagramOther.IsHiddenPlantCircumferenceShape {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenPlantCircumferenceShape"))
	}
	if plant2ddiagram.IsHiddenGridPathShape != plant2ddiagramOther.IsHiddenGridPathShape {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenGridPathShape"))
	}
	if plant2ddiagram.IsHiddenRhombusGridShape != plant2ddiagramOther.IsHiddenRhombusGridShape {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenRhombusGridShape"))
	}
	if plant2ddiagram.IsHiddenExplanationTextShape != plant2ddiagramOther.IsHiddenExplanationTextShape {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenExplanationTextShape"))
	}
	if plant2ddiagram.IsHiddenRotatedReferenceRhombus != plant2ddiagramOther.IsHiddenRotatedReferenceRhombus {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenRotatedReferenceRhombus"))
	}
	if plant2ddiagram.IsHiddenRotatedPlantCircumferenceShape != plant2ddiagramOther.IsHiddenRotatedPlantCircumferenceShape {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenRotatedPlantCircumferenceShape"))
	}
	if plant2ddiagram.IsHiddenRotatedGridPathShape != plant2ddiagramOther.IsHiddenRotatedGridPathShape {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenRotatedGridPathShape"))
	}
	if plant2ddiagram.IsHiddenRotatedRhombusGridShape != plant2ddiagramOther.IsHiddenRotatedRhombusGridShape {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenRotatedRhombusGridShape"))
	}
	if plant2ddiagram.IsHiddenGrowthPathRhombusGridShape != plant2ddiagramOther.IsHiddenGrowthPathRhombusGridShape {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenGrowthPathRhombusGridShape"))
	}
	if plant2ddiagram.IsHiddenGrowthVectorShape != plant2ddiagramOther.IsHiddenGrowthVectorShape {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenGrowthVectorShape"))
	}
	if plant2ddiagram.IsHiddenPerpendicularVectorGrid != plant2ddiagramOther.IsHiddenPerpendicularVectorGrid {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenPerpendicularVectorGrid"))
	}
	if plant2ddiagram.IsHiddenBaseVectorShapeGrid != plant2ddiagramOther.IsHiddenBaseVectorShapeGrid {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenBaseVectorShapeGrid"))
	}
	if plant2ddiagram.IsHiddenArcNormalVectorShapeGrid != plant2ddiagramOther.IsHiddenArcNormalVectorShapeGrid {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenArcNormalVectorShapeGrid"))
	}
	if plant2ddiagram.IsHiddenStartArcShapeGrid != plant2ddiagramOther.IsHiddenStartArcShapeGrid {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenStartArcShapeGrid"))
	}
	if plant2ddiagram.IsHiddenMidArcVectorShapeGrid != plant2ddiagramOther.IsHiddenMidArcVectorShapeGrid {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenMidArcVectorShapeGrid"))
	}
	if plant2ddiagram.IsHiddenEndArcShapeGrid != plant2ddiagramOther.IsHiddenEndArcShapeGrid {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenEndArcShapeGrid"))
	}
	if plant2ddiagram.IsHiddenGrowthCurve2D != plant2ddiagramOther.IsHiddenGrowthCurve2D {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenGrowthCurve2D"))
	}
	if plant2ddiagram.IsHiddenStackOfGrowthCurve2DByGrowthVector != plant2ddiagramOther.IsHiddenStackOfGrowthCurve2DByGrowthVector {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurve2DByGrowthVector"))
	}
	if plant2ddiagram.IsChecked != plant2ddiagramOther.IsChecked {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsChecked"))
	}
	if plant2ddiagram.ComputedPrefix != plant2ddiagramOther.ComputedPrefix {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if plant2ddiagram.IsExpanded != plant2ddiagramOther.IsExpanded {
		diffs = append(diffs, plant2ddiagram.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (plant3ddiagram *Plant3DDiagram) GongDiff(stage *Stage, plant3ddiagramOther *Plant3DDiagram) (diffs []string) {
	// insertion point for field diffs
	if plant3ddiagram.Name != plant3ddiagramOther.Name {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "Name"))
	}
	if plant3ddiagram.IsHiddenStemCylinder3DShape != plant3ddiagramOther.IsHiddenStemCylinder3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsHiddenStemCylinder3DShape"))
	}
	if (plant3ddiagram.StemCylinder3DShape == nil) != (plant3ddiagramOther.StemCylinder3DShape == nil) {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "StemCylinder3DShape"))
	} else if plant3ddiagram.StemCylinder3DShape != nil && plant3ddiagramOther.StemCylinder3DShape != nil {
		if plant3ddiagram.StemCylinder3DShape != plant3ddiagramOther.StemCylinder3DShape {
			diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "StemCylinder3DShape"))
		}
	}
	if plant3ddiagram.IsHiddenParastichyNCurves3DShape != plant3ddiagramOther.IsHiddenParastichyNCurves3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsHiddenParastichyNCurves3DShape"))
	}
	if (plant3ddiagram.ParastichyNCurves3DShape == nil) != (plant3ddiagramOther.ParastichyNCurves3DShape == nil) {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "ParastichyNCurves3DShape"))
	} else if plant3ddiagram.ParastichyNCurves3DShape != nil && plant3ddiagramOther.ParastichyNCurves3DShape != nil {
		if plant3ddiagram.ParastichyNCurves3DShape != plant3ddiagramOther.ParastichyNCurves3DShape {
			diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "ParastichyNCurves3DShape"))
		}
	}
	if plant3ddiagram.IsHiddenParastichyMCurves3DShape != plant3ddiagramOther.IsHiddenParastichyMCurves3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsHiddenParastichyMCurves3DShape"))
	}
	if (plant3ddiagram.ParastichyMCurves3DShape == nil) != (plant3ddiagramOther.ParastichyMCurves3DShape == nil) {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "ParastichyMCurves3DShape"))
	} else if plant3ddiagram.ParastichyMCurves3DShape != nil && plant3ddiagramOther.ParastichyMCurves3DShape != nil {
		if plant3ddiagram.ParastichyMCurves3DShape != plant3ddiagramOther.ParastichyMCurves3DShape {
			diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "ParastichyMCurves3DShape"))
		}
	}
	if plant3ddiagram.IsHiddenCutLine3DShape != plant3ddiagramOther.IsHiddenCutLine3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsHiddenCutLine3DShape"))
	}
	if (plant3ddiagram.CutLine3DShape == nil) != (plant3ddiagramOther.CutLine3DShape == nil) {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "CutLine3DShape"))
	} else if plant3ddiagram.CutLine3DShape != nil && plant3ddiagramOther.CutLine3DShape != nil {
		if plant3ddiagram.CutLine3DShape != plant3ddiagramOther.CutLine3DShape {
			diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "CutLine3DShape"))
		}
	}
	if plant3ddiagram.IsHiddenCircumference3DShape != plant3ddiagramOther.IsHiddenCircumference3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsHiddenCircumference3DShape"))
	}
	if (plant3ddiagram.Circumference3DShape == nil) != (plant3ddiagramOther.Circumference3DShape == nil) {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "Circumference3DShape"))
	} else if plant3ddiagram.Circumference3DShape != nil && plant3ddiagramOther.Circumference3DShape != nil {
		if plant3ddiagram.Circumference3DShape != plant3ddiagramOther.Circumference3DShape {
			diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "Circumference3DShape"))
		}
	}
	if plant3ddiagram.IsHiddenTiledFloor3DShape != plant3ddiagramOther.IsHiddenTiledFloor3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsHiddenTiledFloor3DShape"))
	}
	if (plant3ddiagram.Rendered3DShape == nil) != (plant3ddiagramOther.Rendered3DShape == nil) {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
	} else if plant3ddiagram.Rendered3DShape != nil && plant3ddiagramOther.Rendered3DShape != nil {
		if plant3ddiagram.Rendered3DShape != plant3ddiagramOther.Rendered3DShape {
			diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
		}
	}
	if plant3ddiagram.IsChecked != plant3ddiagramOther.IsChecked {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsChecked"))
	}
	if plant3ddiagram.ComputedPrefix != plant3ddiagramOther.ComputedPrefix {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if plant3ddiagram.IsExpanded != plant3ddiagramOther.IsExpanded {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (plantabstract *PlantAbstract) GongDiff(stage *Stage, plantabstractOther *PlantAbstract) (diffs []string) {
	// insertion point for field diffs
	if plantabstract.Name != plantabstractOther.Name {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "Name"))
	}
	if plantabstract.N != plantabstractOther.N {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "N"))
	}
	if plantabstract.M != plantabstractOther.M {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "M"))
	}
	if plantabstract.StackHeight != plantabstractOther.StackHeight {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "StackHeight"))
	}
	if plantabstract.RhombusInsideAngle != plantabstractOther.RhombusInsideAngle {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "RhombusInsideAngle"))
	}
	if plantabstract.RhombusSideLength != plantabstractOther.RhombusSideLength {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "RhombusSideLength"))
	}
	if plantabstract.PlantType != plantabstractOther.PlantType {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "PlantType"))
	}
	if (plantabstract.VaseAbstract == nil) != (plantabstractOther.VaseAbstract == nil) {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "VaseAbstract"))
	} else if plantabstract.VaseAbstract != nil && plantabstractOther.VaseAbstract != nil {
		if plantabstract.VaseAbstract != plantabstractOther.VaseAbstract {
			diffs = append(diffs, plantabstract.GongMarshallField(stage, "VaseAbstract"))
		}
	}
	if (plantabstract.StoolAbstract == nil) != (plantabstractOther.StoolAbstract == nil) {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "StoolAbstract"))
	} else if plantabstract.StoolAbstract != nil && plantabstractOther.StoolAbstract != nil {
		if plantabstract.StoolAbstract != plantabstractOther.StoolAbstract {
			diffs = append(diffs, plantabstract.GongMarshallField(stage, "StoolAbstract"))
		}
	}
	if (plantabstract.ClockAbstract == nil) != (plantabstractOther.ClockAbstract == nil) {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "ClockAbstract"))
	} else if plantabstract.ClockAbstract != nil && plantabstractOther.ClockAbstract != nil {
		if plantabstract.ClockAbstract != plantabstractOther.ClockAbstract {
			diffs = append(diffs, plantabstract.GongMarshallField(stage, "ClockAbstract"))
		}
	}
	if plantabstract.CurrentView != plantabstractOther.CurrentView {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "CurrentView"))
	}
	if plantabstract.ComputedPrefix != plantabstractOther.ComputedPrefix {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "ComputedPrefix"))
	}
	if plantabstract.IsExpanded != plantabstractOther.IsExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsExpanded"))
	}
	if plantabstract.IsSelected != plantabstractOther.IsSelected {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsSelected"))
	}
	if plantabstract.IsPlant2DDiagramsNodeExpanded != plantabstractOther.IsPlant2DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsPlant2DDiagramsNodeExpanded"))
	}
	Plant2DDiagramsDifferent := false
	if len(plantabstract.Plant2DDiagrams) != len(plantabstractOther.Plant2DDiagrams) {
		Plant2DDiagramsDifferent = true
	} else {
		for i := range plantabstract.Plant2DDiagrams {
			if (plantabstract.Plant2DDiagrams[i] == nil) != (plantabstractOther.Plant2DDiagrams[i] == nil) {
				Plant2DDiagramsDifferent = true
				break
			} else if plantabstract.Plant2DDiagrams[i] != nil && plantabstractOther.Plant2DDiagrams[i] != nil {
				// this is a pointer comparaison
				if plantabstract.Plant2DDiagrams[i] != plantabstractOther.Plant2DDiagrams[i] {
					Plant2DDiagramsDifferent = true
					break
				}
			}
		}
	}
	if Plant2DDiagramsDifferent {
		ops := Diff(stage, plantabstract, plantabstractOther, "Plant2DDiagrams", plantabstractOther.Plant2DDiagrams, plantabstract.Plant2DDiagrams)
		diffs = append(diffs, ops)
	}
	if plantabstract.IsPlant3DDiagramsNodeExpanded != plantabstractOther.IsPlant3DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsPlant3DDiagramsNodeExpanded"))
	}
	Plant3DDiagramsDifferent := false
	if len(plantabstract.Plant3DDiagrams) != len(plantabstractOther.Plant3DDiagrams) {
		Plant3DDiagramsDifferent = true
	} else {
		for i := range plantabstract.Plant3DDiagrams {
			if (plantabstract.Plant3DDiagrams[i] == nil) != (plantabstractOther.Plant3DDiagrams[i] == nil) {
				Plant3DDiagramsDifferent = true
				break
			} else if plantabstract.Plant3DDiagrams[i] != nil && plantabstractOther.Plant3DDiagrams[i] != nil {
				// this is a pointer comparaison
				if plantabstract.Plant3DDiagrams[i] != plantabstractOther.Plant3DDiagrams[i] {
					Plant3DDiagramsDifferent = true
					break
				}
			}
		}
	}
	if Plant3DDiagramsDifferent {
		ops := Diff(stage, plantabstract, plantabstractOther, "Plant3DDiagrams", plantabstractOther.Plant3DDiagrams, plantabstract.Plant3DDiagrams)
		diffs = append(diffs, ops)
	}
	if plantabstract.IsVase2DDiagramsNodeExpanded != plantabstractOther.IsVase2DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsVase2DDiagramsNodeExpanded"))
	}
	Vase2DDiagramsDifferent := false
	if len(plantabstract.Vase2DDiagrams) != len(plantabstractOther.Vase2DDiagrams) {
		Vase2DDiagramsDifferent = true
	} else {
		for i := range plantabstract.Vase2DDiagrams {
			if (plantabstract.Vase2DDiagrams[i] == nil) != (plantabstractOther.Vase2DDiagrams[i] == nil) {
				Vase2DDiagramsDifferent = true
				break
			} else if plantabstract.Vase2DDiagrams[i] != nil && plantabstractOther.Vase2DDiagrams[i] != nil {
				// this is a pointer comparaison
				if plantabstract.Vase2DDiagrams[i] != plantabstractOther.Vase2DDiagrams[i] {
					Vase2DDiagramsDifferent = true
					break
				}
			}
		}
	}
	if Vase2DDiagramsDifferent {
		ops := Diff(stage, plantabstract, plantabstractOther, "Vase2DDiagrams", plantabstractOther.Vase2DDiagrams, plantabstract.Vase2DDiagrams)
		diffs = append(diffs, ops)
	}
	if plantabstract.IsVase3DDiagramsNodeExpanded != plantabstractOther.IsVase3DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsVase3DDiagramsNodeExpanded"))
	}
	Vase3DDiagramsDifferent := false
	if len(plantabstract.Vase3DDiagrams) != len(plantabstractOther.Vase3DDiagrams) {
		Vase3DDiagramsDifferent = true
	} else {
		for i := range plantabstract.Vase3DDiagrams {
			if (plantabstract.Vase3DDiagrams[i] == nil) != (plantabstractOther.Vase3DDiagrams[i] == nil) {
				Vase3DDiagramsDifferent = true
				break
			} else if plantabstract.Vase3DDiagrams[i] != nil && plantabstractOther.Vase3DDiagrams[i] != nil {
				// this is a pointer comparaison
				if plantabstract.Vase3DDiagrams[i] != plantabstractOther.Vase3DDiagrams[i] {
					Vase3DDiagramsDifferent = true
					break
				}
			}
		}
	}
	if Vase3DDiagramsDifferent {
		ops := Diff(stage, plantabstract, plantabstractOther, "Vase3DDiagrams", plantabstractOther.Vase3DDiagrams, plantabstract.Vase3DDiagrams)
		diffs = append(diffs, ops)
	}
	if plantabstract.IsStool2DDiagramsNodeExpanded != plantabstractOther.IsStool2DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsStool2DDiagramsNodeExpanded"))
	}
	Stool2DDiagramsDifferent := false
	if len(plantabstract.Stool2DDiagrams) != len(plantabstractOther.Stool2DDiagrams) {
		Stool2DDiagramsDifferent = true
	} else {
		for i := range plantabstract.Stool2DDiagrams {
			if (plantabstract.Stool2DDiagrams[i] == nil) != (plantabstractOther.Stool2DDiagrams[i] == nil) {
				Stool2DDiagramsDifferent = true
				break
			} else if plantabstract.Stool2DDiagrams[i] != nil && plantabstractOther.Stool2DDiagrams[i] != nil {
				// this is a pointer comparaison
				if plantabstract.Stool2DDiagrams[i] != plantabstractOther.Stool2DDiagrams[i] {
					Stool2DDiagramsDifferent = true
					break
				}
			}
		}
	}
	if Stool2DDiagramsDifferent {
		ops := Diff(stage, plantabstract, plantabstractOther, "Stool2DDiagrams", plantabstractOther.Stool2DDiagrams, plantabstract.Stool2DDiagrams)
		diffs = append(diffs, ops)
	}
	if plantabstract.IsStool3DDiagramsNodeExpanded != plantabstractOther.IsStool3DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsStool3DDiagramsNodeExpanded"))
	}
	Stool3DDiagramsDifferent := false
	if len(plantabstract.Stool3DDiagrams) != len(plantabstractOther.Stool3DDiagrams) {
		Stool3DDiagramsDifferent = true
	} else {
		for i := range plantabstract.Stool3DDiagrams {
			if (plantabstract.Stool3DDiagrams[i] == nil) != (plantabstractOther.Stool3DDiagrams[i] == nil) {
				Stool3DDiagramsDifferent = true
				break
			} else if plantabstract.Stool3DDiagrams[i] != nil && plantabstractOther.Stool3DDiagrams[i] != nil {
				// this is a pointer comparaison
				if plantabstract.Stool3DDiagrams[i] != plantabstractOther.Stool3DDiagrams[i] {
					Stool3DDiagramsDifferent = true
					break
				}
			}
		}
	}
	if Stool3DDiagramsDifferent {
		ops := Diff(stage, plantabstract, plantabstractOther, "Stool3DDiagrams", plantabstractOther.Stool3DDiagrams, plantabstract.Stool3DDiagrams)
		diffs = append(diffs, ops)
	}
	if plantabstract.IsClock2DDiagramsNodeExpanded != plantabstractOther.IsClock2DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsClock2DDiagramsNodeExpanded"))
	}
	Clock2DDiagramsDifferent := false
	if len(plantabstract.Clock2DDiagrams) != len(plantabstractOther.Clock2DDiagrams) {
		Clock2DDiagramsDifferent = true
	} else {
		for i := range plantabstract.Clock2DDiagrams {
			if (plantabstract.Clock2DDiagrams[i] == nil) != (plantabstractOther.Clock2DDiagrams[i] == nil) {
				Clock2DDiagramsDifferent = true
				break
			} else if plantabstract.Clock2DDiagrams[i] != nil && plantabstractOther.Clock2DDiagrams[i] != nil {
				// this is a pointer comparaison
				if plantabstract.Clock2DDiagrams[i] != plantabstractOther.Clock2DDiagrams[i] {
					Clock2DDiagramsDifferent = true
					break
				}
			}
		}
	}
	if Clock2DDiagramsDifferent {
		ops := Diff(stage, plantabstract, plantabstractOther, "Clock2DDiagrams", plantabstractOther.Clock2DDiagrams, plantabstract.Clock2DDiagrams)
		diffs = append(diffs, ops)
	}
	if plantabstract.IsClock3DDiagramsNodeExpanded != plantabstractOther.IsClock3DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsClock3DDiagramsNodeExpanded"))
	}
	Clock3DDiagramsDifferent := false
	if len(plantabstract.Clock3DDiagrams) != len(plantabstractOther.Clock3DDiagrams) {
		Clock3DDiagramsDifferent = true
	} else {
		for i := range plantabstract.Clock3DDiagrams {
			if (plantabstract.Clock3DDiagrams[i] == nil) != (plantabstractOther.Clock3DDiagrams[i] == nil) {
				Clock3DDiagramsDifferent = true
				break
			} else if plantabstract.Clock3DDiagrams[i] != nil && plantabstractOther.Clock3DDiagrams[i] != nil {
				// this is a pointer comparaison
				if plantabstract.Clock3DDiagrams[i] != plantabstractOther.Clock3DDiagrams[i] {
					Clock3DDiagramsDifferent = true
					break
				}
			}
		}
	}
	if Clock3DDiagramsDifferent {
		ops := Diff(stage, plantabstract, plantabstractOther, "Clock3DDiagrams", plantabstractOther.Clock3DDiagrams, plantabstract.Clock3DDiagrams)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (plantcircumferenceshape *PlantCircumferenceShape) GongDiff(stage *Stage, plantcircumferenceshapeOther *PlantCircumferenceShape) (diffs []string) {
	// insertion point for field diffs
	if plantcircumferenceshape.Name != plantcircumferenceshapeOther.Name {
		diffs = append(diffs, plantcircumferenceshape.GongMarshallField(stage, "Name"))
	}
	if plantcircumferenceshape.AngleDegree != plantcircumferenceshapeOther.AngleDegree {
		diffs = append(diffs, plantcircumferenceshape.GongMarshallField(stage, "AngleDegree"))
	}
	if plantcircumferenceshape.Length != plantcircumferenceshapeOther.Length {
		diffs = append(diffs, plantcircumferenceshape.GongMarshallField(stage, "Length"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pointsandlines3dshape *PointsAndLines3DShape) GongDiff(stage *Stage, pointsandlines3dshapeOther *PointsAndLines3DShape) (diffs []string) {
	// insertion point for field diffs
	if pointsandlines3dshape.Name != pointsandlines3dshapeOther.Name {
		diffs = append(diffs, pointsandlines3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pxshape *PxShape) GongDiff(stage *Stage, pxshapeOther *PxShape) (diffs []string) {
	// insertion point for field diffs
	if pxshape.Name != pxshapeOther.Name {
		diffs = append(diffs, pxshape.GongMarshallField(stage, "Name"))
	}
	if pxshape.X != pxshapeOther.X {
		diffs = append(diffs, pxshape.GongMarshallField(stage, "X"))
	}
	if pxshape.Y != pxshapeOther.Y {
		diffs = append(diffs, pxshape.GongMarshallField(stage, "Y"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rendered3dshape *Rendered3DShape) GongDiff(stage *Stage, rendered3dshapeOther *Rendered3DShape) (diffs []string) {
	// insertion point for field diffs
	if rendered3dshape.Name != rendered3dshapeOther.Name {
		diffs = append(diffs, rendered3dshape.GongMarshallField(stage, "Name"))
	}
	if rendered3dshape.ViewX != rendered3dshapeOther.ViewX {
		diffs = append(diffs, rendered3dshape.GongMarshallField(stage, "ViewX"))
	}
	if rendered3dshape.ViewY != rendered3dshapeOther.ViewY {
		diffs = append(diffs, rendered3dshape.GongMarshallField(stage, "ViewY"))
	}
	if rendered3dshape.ViewZ != rendered3dshapeOther.ViewZ {
		diffs = append(diffs, rendered3dshape.GongMarshallField(stage, "ViewZ"))
	}
	if rendered3dshape.TargetX != rendered3dshapeOther.TargetX {
		diffs = append(diffs, rendered3dshape.GongMarshallField(stage, "TargetX"))
	}
	if rendered3dshape.TargetY != rendered3dshapeOther.TargetY {
		diffs = append(diffs, rendered3dshape.GongMarshallField(stage, "TargetY"))
	}
	if rendered3dshape.TargetZ != rendered3dshapeOther.TargetZ {
		diffs = append(diffs, rendered3dshape.GongMarshallField(stage, "TargetZ"))
	}
	if rendered3dshape.Fov != rendered3dshapeOther.Fov {
		diffs = append(diffs, rendered3dshape.GongMarshallField(stage, "Fov"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rhombusshape *RhombusShape) GongDiff(stage *Stage, rhombusshapeOther *RhombusShape) (diffs []string) {
	// insertion point for field diffs
	if rhombusshape.Name != rhombusshapeOther.Name {
		diffs = append(diffs, rhombusshape.GongMarshallField(stage, "Name"))
	}
	if rhombusshape.X != rhombusshapeOther.X {
		diffs = append(diffs, rhombusshape.GongMarshallField(stage, "X"))
	}
	if rhombusshape.Y != rhombusshapeOther.Y {
		diffs = append(diffs, rhombusshape.GongMarshallField(stage, "Y"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rhombusstuff *RhombusStuff) GongDiff(stage *Stage, rhombusstuffOther *RhombusStuff) (diffs []string) {
	// insertion point for field diffs
	if rhombusstuff.Name != rhombusstuffOther.Name {
		diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongDiff(stage *Stage, rotatedrhombusgridshapeOther *RotatedRhombusGridShape) (diffs []string) {
	// insertion point for field diffs
	if rotatedrhombusgridshape.Name != rotatedrhombusgridshapeOther.Name {
		diffs = append(diffs, rotatedrhombusgridshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rotatedrhombusshape *RotatedRhombusShape) GongDiff(stage *Stage, rotatedrhombusshapeOther *RotatedRhombusShape) (diffs []string) {
	// insertion point for field diffs
	if rotatedrhombusshape.Name != rotatedrhombusshapeOther.Name {
		diffs = append(diffs, rotatedrhombusshape.GongMarshallField(stage, "Name"))
	}
	if rotatedrhombusshape.X != rotatedrhombusshapeOther.X {
		diffs = append(diffs, rotatedrhombusshape.GongMarshallField(stage, "X"))
	}
	if rotatedrhombusshape.Y != rotatedrhombusshapeOther.Y {
		diffs = append(diffs, rotatedrhombusshape.GongMarshallField(stage, "Y"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongDiff(stage *Stage, rotatedsampledpoints3dshapeOther *RotatedSampledPoints3DShape) (diffs []string) {
	// insertion point for field diffs
	if rotatedsampledpoints3dshape.Name != rotatedsampledpoints3dshapeOther.Name {
		diffs = append(diffs, rotatedsampledpoints3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongDiff(stage *Stage, rotatedseatandlegs3dshapeOther *RotatedSeatAndLegs3DShape) (diffs []string) {
	// insertion point for field diffs
	if rotatedseatandlegs3dshape.Name != rotatedseatandlegs3dshapeOther.Name {
		diffs = append(diffs, rotatedseatandlegs3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (sampledpoints3dshape *SampledPoints3DShape) GongDiff(stage *Stage, sampledpoints3dshapeOther *SampledPoints3DShape) (diffs []string) {
	// insertion point for field diffs
	if sampledpoints3dshape.Name != sampledpoints3dshapeOther.Name {
		diffs = append(diffs, sampledpoints3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (seat3dshape *Seat3DShape) GongDiff(stage *Stage, seat3dshapeOther *Seat3DShape) (diffs []string) {
	// insertion point for field diffs
	if seat3dshape.Name != seat3dshapeOther.Name {
		diffs = append(diffs, seat3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (seatandlegs3dshape *SeatAndLegs3DShape) GongDiff(stage *Stage, seatandlegs3dshapeOther *SeatAndLegs3DShape) (diffs []string) {
	// insertion point for field diffs
	if seatandlegs3dshape.Name != seatandlegs3dshapeOther.Name {
		diffs = append(diffs, seatandlegs3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (seatbottomcurveshape *SeatBottomCurveShape) GongDiff(stage *Stage, seatbottomcurveshapeOther *SeatBottomCurveShape) (diffs []string) {
	// insertion point for field diffs
	if seatbottomcurveshape.Name != seatbottomcurveshapeOther.Name {
		diffs = append(diffs, seatbottomcurveshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (seattopcurveshape *SeatTopCurveShape) GongDiff(stage *Stage, seattopcurveshapeOther *SeatTopCurveShape) (diffs []string) {
	// insertion point for field diffs
	if seattopcurveshape.Name != seattopcurveshapeOther.Name {
		diffs = append(diffs, seattopcurveshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongDiff(stage *Stage, shiftedbottomtopstartarcshapeOther *ShiftedBottomTopStartArcShape) (diffs []string) {
	// insertion point for field diffs
	if shiftedbottomtopstartarcshape.Name != shiftedbottomtopstartarcshapeOther.Name {
		diffs = append(diffs, shiftedbottomtopstartarcshape.GongMarshallField(stage, "Name"))
	}
	if shiftedbottomtopstartarcshape.StartX != shiftedbottomtopstartarcshapeOther.StartX {
		diffs = append(diffs, shiftedbottomtopstartarcshape.GongMarshallField(stage, "StartX"))
	}
	if shiftedbottomtopstartarcshape.StartY != shiftedbottomtopstartarcshapeOther.StartY {
		diffs = append(diffs, shiftedbottomtopstartarcshape.GongMarshallField(stage, "StartY"))
	}
	if shiftedbottomtopstartarcshape.EndX != shiftedbottomtopstartarcshapeOther.EndX {
		diffs = append(diffs, shiftedbottomtopstartarcshape.GongMarshallField(stage, "EndX"))
	}
	if shiftedbottomtopstartarcshape.EndY != shiftedbottomtopstartarcshapeOther.EndY {
		diffs = append(diffs, shiftedbottomtopstartarcshape.GongMarshallField(stage, "EndY"))
	}
	if shiftedbottomtopstartarcshape.XAxisRotation != shiftedbottomtopstartarcshapeOther.XAxisRotation {
		diffs = append(diffs, shiftedbottomtopstartarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if shiftedbottomtopstartarcshape.LargeArcFlag != shiftedbottomtopstartarcshapeOther.LargeArcFlag {
		diffs = append(diffs, shiftedbottomtopstartarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if shiftedbottomtopstartarcshape.SweepFlag != shiftedbottomtopstartarcshapeOther.SweepFlag {
		diffs = append(diffs, shiftedbottomtopstartarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if shiftedbottomtopstartarcshape.RadiusX != shiftedbottomtopstartarcshapeOther.RadiusX {
		diffs = append(diffs, shiftedbottomtopstartarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if shiftedbottomtopstartarcshape.RadiusY != shiftedbottomtopstartarcshapeOther.RadiusY {
		diffs = append(diffs, shiftedbottomtopstartarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongDiff(stage *Stage, shiftedbottomtopstartarcshapegridOther *ShiftedBottomTopStartArcShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if shiftedbottomtopstartarcshapegrid.Name != shiftedbottomtopstartarcshapegridOther.Name {
		diffs = append(diffs, shiftedbottomtopstartarcshapegrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongDiff(stage *Stage, shiftedleftgrowthcurve2dribbonOther *ShiftedLeftGrowthCurve2DRibbon) (diffs []string) {
	// insertion point for field diffs
	if shiftedleftgrowthcurve2dribbon.Name != shiftedleftgrowthcurve2dribbonOther.Name {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbon.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongDiff(stage *Stage, shiftedleftgrowthcurve2dribbonendshapeOther *ShiftedLeftGrowthCurve2DRibbonEndShape) (diffs []string) {
	// insertion point for field diffs
	if shiftedleftgrowthcurve2dribbonendshape.Name != shiftedleftgrowthcurve2dribbonendshapeOther.Name {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.BottomStartX != shiftedleftgrowthcurve2dribbonendshapeOther.BottomStartX {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.BottomStartY != shiftedleftgrowthcurve2dribbonendshapeOther.BottomStartY {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.BottomEndX != shiftedleftgrowthcurve2dribbonendshapeOther.BottomEndX {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.BottomEndY != shiftedleftgrowthcurve2dribbonendshapeOther.BottomEndY {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.BottomRadiusX != shiftedleftgrowthcurve2dribbonendshapeOther.BottomRadiusX {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.BottomRadiusY != shiftedleftgrowthcurve2dribbonendshapeOther.BottomRadiusY {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.BottomXAxisRotation != shiftedleftgrowthcurve2dribbonendshapeOther.BottomXAxisRotation {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.BottomLargeArcFlag != shiftedleftgrowthcurve2dribbonendshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.BottomSweepFlag != shiftedleftgrowthcurve2dribbonendshapeOther.BottomSweepFlag {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.TopStartX != shiftedleftgrowthcurve2dribbonendshapeOther.TopStartX {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.TopStartY != shiftedleftgrowthcurve2dribbonendshapeOther.TopStartY {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.TopEndX != shiftedleftgrowthcurve2dribbonendshapeOther.TopEndX {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.TopEndY != shiftedleftgrowthcurve2dribbonendshapeOther.TopEndY {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.TopRadiusX != shiftedleftgrowthcurve2dribbonendshapeOther.TopRadiusX {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.TopRadiusY != shiftedleftgrowthcurve2dribbonendshapeOther.TopRadiusY {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.TopXAxisRotation != shiftedleftgrowthcurve2dribbonendshapeOther.TopXAxisRotation {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.TopLargeArcFlag != shiftedleftgrowthcurve2dribbonendshapeOther.TopLargeArcFlag {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if shiftedleftgrowthcurve2dribbonendshape.TopSweepFlag != shiftedleftgrowthcurve2dribbonendshapeOther.TopSweepFlag {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongDiff(stage *Stage, shiftedleftgrowthcurve2dribbonstartshapeOther *ShiftedLeftGrowthCurve2DRibbonStartShape) (diffs []string) {
	// insertion point for field diffs
	if shiftedleftgrowthcurve2dribbonstartshape.Name != shiftedleftgrowthcurve2dribbonstartshapeOther.Name {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.BottomStartX != shiftedleftgrowthcurve2dribbonstartshapeOther.BottomStartX {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.BottomStartY != shiftedleftgrowthcurve2dribbonstartshapeOther.BottomStartY {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.BottomEndX != shiftedleftgrowthcurve2dribbonstartshapeOther.BottomEndX {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.BottomEndY != shiftedleftgrowthcurve2dribbonstartshapeOther.BottomEndY {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.BottomRadiusX != shiftedleftgrowthcurve2dribbonstartshapeOther.BottomRadiusX {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.BottomRadiusY != shiftedleftgrowthcurve2dribbonstartshapeOther.BottomRadiusY {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.BottomXAxisRotation != shiftedleftgrowthcurve2dribbonstartshapeOther.BottomXAxisRotation {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.BottomLargeArcFlag != shiftedleftgrowthcurve2dribbonstartshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.BottomSweepFlag != shiftedleftgrowthcurve2dribbonstartshapeOther.BottomSweepFlag {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.TopStartX != shiftedleftgrowthcurve2dribbonstartshapeOther.TopStartX {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.TopStartY != shiftedleftgrowthcurve2dribbonstartshapeOther.TopStartY {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.TopEndX != shiftedleftgrowthcurve2dribbonstartshapeOther.TopEndX {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.TopEndY != shiftedleftgrowthcurve2dribbonstartshapeOther.TopEndY {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.TopRadiusX != shiftedleftgrowthcurve2dribbonstartshapeOther.TopRadiusX {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.TopRadiusY != shiftedleftgrowthcurve2dribbonstartshapeOther.TopRadiusY {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.TopXAxisRotation != shiftedleftgrowthcurve2dribbonstartshapeOther.TopXAxisRotation {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.TopLargeArcFlag != shiftedleftgrowthcurve2dribbonstartshapeOther.TopLargeArcFlag {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if shiftedleftgrowthcurve2dribbonstartshape.TopSweepFlag != shiftedleftgrowthcurve2dribbonstartshapeOther.TopSweepFlag {
		diffs = append(diffs, shiftedleftgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongDiff(stage *Stage, shiftedleftpartiallygrowthcurve2dribbonOther *ShiftedLeftPartiallyGrowthCurve2DRibbon) (diffs []string) {
	// insertion point for field diffs
	if shiftedleftpartiallygrowthcurve2dribbon.Name != shiftedleftpartiallygrowthcurve2dribbonOther.Name {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbon.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongDiff(stage *Stage, shiftedleftpartiallygrowthcurve2dribbonendshapeOther *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) (diffs []string) {
	// insertion point for field diffs
	if shiftedleftpartiallygrowthcurve2dribbonendshape.Name != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.Name {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.BottomStartX != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.BottomStartX {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.BottomStartY != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.BottomStartY {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.BottomEndX != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.BottomEndX {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.BottomEndY != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.BottomEndY {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.BottomRadiusX != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.BottomRadiusX {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.BottomRadiusY != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.BottomRadiusY {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.BottomXAxisRotation != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.BottomXAxisRotation {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.BottomLargeArcFlag != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.BottomSweepFlag != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.BottomSweepFlag {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.TopStartX != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.TopStartX {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.TopStartY != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.TopStartY {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.TopEndX != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.TopEndX {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.TopEndY != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.TopEndY {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.TopRadiusX != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.TopRadiusX {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.TopRadiusY != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.TopRadiusY {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.TopXAxisRotation != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.TopXAxisRotation {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.TopLargeArcFlag != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.TopLargeArcFlag {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonendshape.TopSweepFlag != shiftedleftpartiallygrowthcurve2dribbonendshapeOther.TopSweepFlag {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongDiff(stage *Stage, shiftedleftpartiallygrowthcurve2dribbonstartshapeOther *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) (diffs []string) {
	// insertion point for field diffs
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.Name != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.Name {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomStartX != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.BottomStartX {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomStartY != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.BottomStartY {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomEndX != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.BottomEndX {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomEndY != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.BottomEndY {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomRadiusX != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.BottomRadiusX {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomRadiusY != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.BottomRadiusY {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomXAxisRotation != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.BottomXAxisRotation {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomLargeArcFlag != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.BottomSweepFlag != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.BottomSweepFlag {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.TopStartX != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.TopStartX {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.TopStartY != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.TopStartY {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.TopEndX != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.TopEndX {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.TopEndY != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.TopEndY {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.TopRadiusX != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.TopRadiusX {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.TopRadiusY != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.TopRadiusY {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.TopXAxisRotation != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.TopXAxisRotation {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.TopLargeArcFlag != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.TopLargeArcFlag {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if shiftedleftpartiallygrowthcurve2dribbonstartshape.TopSweepFlag != shiftedleftpartiallygrowthcurve2dribbonstartshapeOther.TopSweepFlag {
		diffs = append(diffs, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongDiff(stage *Stage, shiftedleftstackgrowthcurveendarcshapeOther *ShiftedLeftStackGrowthCurveEndArcShape) (diffs []string) {
	// insertion point for field diffs
	if shiftedleftstackgrowthcurveendarcshape.Name != shiftedleftstackgrowthcurveendarcshapeOther.Name {
		diffs = append(diffs, shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "Name"))
	}
	if shiftedleftstackgrowthcurveendarcshape.StartX != shiftedleftstackgrowthcurveendarcshapeOther.StartX {
		diffs = append(diffs, shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "StartX"))
	}
	if shiftedleftstackgrowthcurveendarcshape.StartY != shiftedleftstackgrowthcurveendarcshapeOther.StartY {
		diffs = append(diffs, shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "StartY"))
	}
	if shiftedleftstackgrowthcurveendarcshape.EndX != shiftedleftstackgrowthcurveendarcshapeOther.EndX {
		diffs = append(diffs, shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "EndX"))
	}
	if shiftedleftstackgrowthcurveendarcshape.EndY != shiftedleftstackgrowthcurveendarcshapeOther.EndY {
		diffs = append(diffs, shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "EndY"))
	}
	if shiftedleftstackgrowthcurveendarcshape.XAxisRotation != shiftedleftstackgrowthcurveendarcshapeOther.XAxisRotation {
		diffs = append(diffs, shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if shiftedleftstackgrowthcurveendarcshape.LargeArcFlag != shiftedleftstackgrowthcurveendarcshapeOther.LargeArcFlag {
		diffs = append(diffs, shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if shiftedleftstackgrowthcurveendarcshape.SweepFlag != shiftedleftstackgrowthcurveendarcshapeOther.SweepFlag {
		diffs = append(diffs, shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if shiftedleftstackgrowthcurveendarcshape.RadiusX != shiftedleftstackgrowthcurveendarcshapeOther.RadiusX {
		diffs = append(diffs, shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if shiftedleftstackgrowthcurveendarcshape.RadiusY != shiftedleftstackgrowthcurveendarcshapeOther.RadiusY {
		diffs = append(diffs, shiftedleftstackgrowthcurveendarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongDiff(stage *Stage, shiftedleftstackgrowthcurvestartarcshapeOther *ShiftedLeftStackGrowthCurveStartArcShape) (diffs []string) {
	// insertion point for field diffs
	if shiftedleftstackgrowthcurvestartarcshape.Name != shiftedleftstackgrowthcurvestartarcshapeOther.Name {
		diffs = append(diffs, shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "Name"))
	}
	if shiftedleftstackgrowthcurvestartarcshape.StartX != shiftedleftstackgrowthcurvestartarcshapeOther.StartX {
		diffs = append(diffs, shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "StartX"))
	}
	if shiftedleftstackgrowthcurvestartarcshape.StartY != shiftedleftstackgrowthcurvestartarcshapeOther.StartY {
		diffs = append(diffs, shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "StartY"))
	}
	if shiftedleftstackgrowthcurvestartarcshape.EndX != shiftedleftstackgrowthcurvestartarcshapeOther.EndX {
		diffs = append(diffs, shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "EndX"))
	}
	if shiftedleftstackgrowthcurvestartarcshape.EndY != shiftedleftstackgrowthcurvestartarcshapeOther.EndY {
		diffs = append(diffs, shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "EndY"))
	}
	if shiftedleftstackgrowthcurvestartarcshape.XAxisRotation != shiftedleftstackgrowthcurvestartarcshapeOther.XAxisRotation {
		diffs = append(diffs, shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if shiftedleftstackgrowthcurvestartarcshape.LargeArcFlag != shiftedleftstackgrowthcurvestartarcshapeOther.LargeArcFlag {
		diffs = append(diffs, shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if shiftedleftstackgrowthcurvestartarcshape.SweepFlag != shiftedleftstackgrowthcurvestartarcshapeOther.SweepFlag {
		diffs = append(diffs, shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if shiftedleftstackgrowthcurvestartarcshape.RadiusX != shiftedleftstackgrowthcurvestartarcshapeOther.RadiusX {
		diffs = append(diffs, shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if shiftedleftstackgrowthcurvestartarcshape.RadiusY != shiftedleftstackgrowthcurvestartarcshapeOther.RadiusY {
		diffs = append(diffs, shiftedleftstackgrowthcurvestartarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongDiff(stage *Stage, shiftedleftstacknormalvectorOther *ShiftedLeftStackNormalVector) (diffs []string) {
	// insertion point for field diffs
	if shiftedleftstacknormalvector.Name != shiftedleftstacknormalvectorOther.Name {
		diffs = append(diffs, shiftedleftstacknormalvector.GongMarshallField(stage, "Name"))
	}
	if shiftedleftstacknormalvector.StartX != shiftedleftstacknormalvectorOther.StartX {
		diffs = append(diffs, shiftedleftstacknormalvector.GongMarshallField(stage, "StartX"))
	}
	if shiftedleftstacknormalvector.StartY != shiftedleftstacknormalvectorOther.StartY {
		diffs = append(diffs, shiftedleftstacknormalvector.GongMarshallField(stage, "StartY"))
	}
	if shiftedleftstacknormalvector.EndX != shiftedleftstacknormalvectorOther.EndX {
		diffs = append(diffs, shiftedleftstacknormalvector.GongMarshallField(stage, "EndX"))
	}
	if shiftedleftstacknormalvector.EndY != shiftedleftstacknormalvectorOther.EndY {
		diffs = append(diffs, shiftedleftstacknormalvector.GongMarshallField(stage, "EndY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongDiff(stage *Stage, shiftedleftstackofgrowthcurveOther *ShiftedLeftStackOfGrowthCurve) (diffs []string) {
	// insertion point for field diffs
	if shiftedleftstackofgrowthcurve.Name != shiftedleftstackofgrowthcurveOther.Name {
		diffs = append(diffs, shiftedleftstackofgrowthcurve.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongDiff(stage *Stage, shiftedleftstackofnormalvectorOther *ShiftedLeftStackOfNormalVector) (diffs []string) {
	// insertion point for field diffs
	if shiftedleftstackofnormalvector.Name != shiftedleftstackofnormalvectorOther.Name {
		diffs = append(diffs, shiftedleftstackofnormalvector.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongDiff(stage *Stage, shiftedrightgrowthcurve2dribbonOther *ShiftedRightGrowthCurve2DRibbon) (diffs []string) {
	// insertion point for field diffs
	if shiftedrightgrowthcurve2dribbon.Name != shiftedrightgrowthcurve2dribbonOther.Name {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbon.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongDiff(stage *Stage, shiftedrightgrowthcurve2dribbonendshapeOther *ShiftedRightGrowthCurve2DRibbonEndShape) (diffs []string) {
	// insertion point for field diffs
	if shiftedrightgrowthcurve2dribbonendshape.Name != shiftedrightgrowthcurve2dribbonendshapeOther.Name {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.BottomStartX != shiftedrightgrowthcurve2dribbonendshapeOther.BottomStartX {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.BottomStartY != shiftedrightgrowthcurve2dribbonendshapeOther.BottomStartY {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.BottomEndX != shiftedrightgrowthcurve2dribbonendshapeOther.BottomEndX {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.BottomEndY != shiftedrightgrowthcurve2dribbonendshapeOther.BottomEndY {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.BottomRadiusX != shiftedrightgrowthcurve2dribbonendshapeOther.BottomRadiusX {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.BottomRadiusY != shiftedrightgrowthcurve2dribbonendshapeOther.BottomRadiusY {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.BottomXAxisRotation != shiftedrightgrowthcurve2dribbonendshapeOther.BottomXAxisRotation {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.BottomLargeArcFlag != shiftedrightgrowthcurve2dribbonendshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.BottomSweepFlag != shiftedrightgrowthcurve2dribbonendshapeOther.BottomSweepFlag {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.TopStartX != shiftedrightgrowthcurve2dribbonendshapeOther.TopStartX {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.TopStartY != shiftedrightgrowthcurve2dribbonendshapeOther.TopStartY {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.TopEndX != shiftedrightgrowthcurve2dribbonendshapeOther.TopEndX {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.TopEndY != shiftedrightgrowthcurve2dribbonendshapeOther.TopEndY {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.TopRadiusX != shiftedrightgrowthcurve2dribbonendshapeOther.TopRadiusX {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.TopRadiusY != shiftedrightgrowthcurve2dribbonendshapeOther.TopRadiusY {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.TopXAxisRotation != shiftedrightgrowthcurve2dribbonendshapeOther.TopXAxisRotation {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.TopLargeArcFlag != shiftedrightgrowthcurve2dribbonendshapeOther.TopLargeArcFlag {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if shiftedrightgrowthcurve2dribbonendshape.TopSweepFlag != shiftedrightgrowthcurve2dribbonendshapeOther.TopSweepFlag {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongDiff(stage *Stage, shiftedrightgrowthcurve2dribbonstartshapeOther *ShiftedRightGrowthCurve2DRibbonStartShape) (diffs []string) {
	// insertion point for field diffs
	if shiftedrightgrowthcurve2dribbonstartshape.Name != shiftedrightgrowthcurve2dribbonstartshapeOther.Name {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.BottomStartX != shiftedrightgrowthcurve2dribbonstartshapeOther.BottomStartX {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.BottomStartY != shiftedrightgrowthcurve2dribbonstartshapeOther.BottomStartY {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.BottomEndX != shiftedrightgrowthcurve2dribbonstartshapeOther.BottomEndX {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.BottomEndY != shiftedrightgrowthcurve2dribbonstartshapeOther.BottomEndY {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.BottomRadiusX != shiftedrightgrowthcurve2dribbonstartshapeOther.BottomRadiusX {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.BottomRadiusY != shiftedrightgrowthcurve2dribbonstartshapeOther.BottomRadiusY {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.BottomXAxisRotation != shiftedrightgrowthcurve2dribbonstartshapeOther.BottomXAxisRotation {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.BottomLargeArcFlag != shiftedrightgrowthcurve2dribbonstartshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.BottomSweepFlag != shiftedrightgrowthcurve2dribbonstartshapeOther.BottomSweepFlag {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.TopStartX != shiftedrightgrowthcurve2dribbonstartshapeOther.TopStartX {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.TopStartY != shiftedrightgrowthcurve2dribbonstartshapeOther.TopStartY {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.TopEndX != shiftedrightgrowthcurve2dribbonstartshapeOther.TopEndX {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.TopEndY != shiftedrightgrowthcurve2dribbonstartshapeOther.TopEndY {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.TopRadiusX != shiftedrightgrowthcurve2dribbonstartshapeOther.TopRadiusX {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.TopRadiusY != shiftedrightgrowthcurve2dribbonstartshapeOther.TopRadiusY {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.TopXAxisRotation != shiftedrightgrowthcurve2dribbonstartshapeOther.TopXAxisRotation {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.TopLargeArcFlag != shiftedrightgrowthcurve2dribbonstartshapeOther.TopLargeArcFlag {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if shiftedrightgrowthcurve2dribbonstartshape.TopSweepFlag != shiftedrightgrowthcurve2dribbonstartshapeOther.TopSweepFlag {
		diffs = append(diffs, shiftedrightgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongDiff(stage *Stage, stackgrowthcurve2dendhalfwayarcshapeOther *StackGrowthCurve2DEndHalfwayArcShape) (diffs []string) {
	// insertion point for field diffs
	if stackgrowthcurve2dendhalfwayarcshape.Name != stackgrowthcurve2dendhalfwayarcshapeOther.Name {
		diffs = append(diffs, stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "Name"))
	}
	if stackgrowthcurve2dendhalfwayarcshape.StartX != stackgrowthcurve2dendhalfwayarcshapeOther.StartX {
		diffs = append(diffs, stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "StartX"))
	}
	if stackgrowthcurve2dendhalfwayarcshape.StartY != stackgrowthcurve2dendhalfwayarcshapeOther.StartY {
		diffs = append(diffs, stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "StartY"))
	}
	if stackgrowthcurve2dendhalfwayarcshape.EndX != stackgrowthcurve2dendhalfwayarcshapeOther.EndX {
		diffs = append(diffs, stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "EndX"))
	}
	if stackgrowthcurve2dendhalfwayarcshape.EndY != stackgrowthcurve2dendhalfwayarcshapeOther.EndY {
		diffs = append(diffs, stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "EndY"))
	}
	if stackgrowthcurve2dendhalfwayarcshape.RadiusX != stackgrowthcurve2dendhalfwayarcshapeOther.RadiusX {
		diffs = append(diffs, stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if stackgrowthcurve2dendhalfwayarcshape.RadiusY != stackgrowthcurve2dendhalfwayarcshapeOther.RadiusY {
		diffs = append(diffs, stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "RadiusY"))
	}
	if stackgrowthcurve2dendhalfwayarcshape.XAxisRotation != stackgrowthcurve2dendhalfwayarcshapeOther.XAxisRotation {
		diffs = append(diffs, stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if stackgrowthcurve2dendhalfwayarcshape.LargeArcFlag != stackgrowthcurve2dendhalfwayarcshapeOther.LargeArcFlag {
		diffs = append(diffs, stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if stackgrowthcurve2dendhalfwayarcshape.SweepFlag != stackgrowthcurve2dendhalfwayarcshapeOther.SweepFlag {
		diffs = append(diffs, stackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongDiff(stage *Stage, stackgrowthcurve2dribbonendshapeOther *StackGrowthCurve2DRibbonEndShape) (diffs []string) {
	// insertion point for field diffs
	if stackgrowthcurve2dribbonendshape.Name != stackgrowthcurve2dribbonendshapeOther.Name {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
	}
	if stackgrowthcurve2dribbonendshape.BottomStartX != stackgrowthcurve2dribbonendshapeOther.BottomStartX {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
	}
	if stackgrowthcurve2dribbonendshape.BottomStartY != stackgrowthcurve2dribbonendshapeOther.BottomStartY {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
	}
	if stackgrowthcurve2dribbonendshape.BottomEndX != stackgrowthcurve2dribbonendshapeOther.BottomEndX {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
	}
	if stackgrowthcurve2dribbonendshape.BottomEndY != stackgrowthcurve2dribbonendshapeOther.BottomEndY {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
	}
	if stackgrowthcurve2dribbonendshape.BottomRadiusX != stackgrowthcurve2dribbonendshapeOther.BottomRadiusX {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if stackgrowthcurve2dribbonendshape.BottomRadiusY != stackgrowthcurve2dribbonendshapeOther.BottomRadiusY {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if stackgrowthcurve2dribbonendshape.BottomXAxisRotation != stackgrowthcurve2dribbonendshapeOther.BottomXAxisRotation {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if stackgrowthcurve2dribbonendshape.BottomLargeArcFlag != stackgrowthcurve2dribbonendshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if stackgrowthcurve2dribbonendshape.BottomSweepFlag != stackgrowthcurve2dribbonendshapeOther.BottomSweepFlag {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if stackgrowthcurve2dribbonendshape.TopStartX != stackgrowthcurve2dribbonendshapeOther.TopStartX {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
	}
	if stackgrowthcurve2dribbonendshape.TopStartY != stackgrowthcurve2dribbonendshapeOther.TopStartY {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
	}
	if stackgrowthcurve2dribbonendshape.TopEndX != stackgrowthcurve2dribbonendshapeOther.TopEndX {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
	}
	if stackgrowthcurve2dribbonendshape.TopEndY != stackgrowthcurve2dribbonendshapeOther.TopEndY {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
	}
	if stackgrowthcurve2dribbonendshape.TopRadiusX != stackgrowthcurve2dribbonendshapeOther.TopRadiusX {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if stackgrowthcurve2dribbonendshape.TopRadiusY != stackgrowthcurve2dribbonendshapeOther.TopRadiusY {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if stackgrowthcurve2dribbonendshape.TopXAxisRotation != stackgrowthcurve2dribbonendshapeOther.TopXAxisRotation {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if stackgrowthcurve2dribbonendshape.TopLargeArcFlag != stackgrowthcurve2dribbonendshapeOther.TopLargeArcFlag {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if stackgrowthcurve2dribbonendshape.TopSweepFlag != stackgrowthcurve2dribbonendshapeOther.TopSweepFlag {
		diffs = append(diffs, stackgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongDiff(stage *Stage, stackgrowthcurve2dribbonstartshapeOther *StackGrowthCurve2DRibbonStartShape) (diffs []string) {
	// insertion point for field diffs
	if stackgrowthcurve2dribbonstartshape.Name != stackgrowthcurve2dribbonstartshapeOther.Name {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
	}
	if stackgrowthcurve2dribbonstartshape.BottomStartX != stackgrowthcurve2dribbonstartshapeOther.BottomStartX {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
	}
	if stackgrowthcurve2dribbonstartshape.BottomStartY != stackgrowthcurve2dribbonstartshapeOther.BottomStartY {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
	}
	if stackgrowthcurve2dribbonstartshape.BottomEndX != stackgrowthcurve2dribbonstartshapeOther.BottomEndX {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
	}
	if stackgrowthcurve2dribbonstartshape.BottomEndY != stackgrowthcurve2dribbonstartshapeOther.BottomEndY {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
	}
	if stackgrowthcurve2dribbonstartshape.BottomRadiusX != stackgrowthcurve2dribbonstartshapeOther.BottomRadiusX {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if stackgrowthcurve2dribbonstartshape.BottomRadiusY != stackgrowthcurve2dribbonstartshapeOther.BottomRadiusY {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if stackgrowthcurve2dribbonstartshape.BottomXAxisRotation != stackgrowthcurve2dribbonstartshapeOther.BottomXAxisRotation {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if stackgrowthcurve2dribbonstartshape.BottomLargeArcFlag != stackgrowthcurve2dribbonstartshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if stackgrowthcurve2dribbonstartshape.BottomSweepFlag != stackgrowthcurve2dribbonstartshapeOther.BottomSweepFlag {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if stackgrowthcurve2dribbonstartshape.TopStartX != stackgrowthcurve2dribbonstartshapeOther.TopStartX {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
	}
	if stackgrowthcurve2dribbonstartshape.TopStartY != stackgrowthcurve2dribbonstartshapeOther.TopStartY {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
	}
	if stackgrowthcurve2dribbonstartshape.TopEndX != stackgrowthcurve2dribbonstartshapeOther.TopEndX {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
	}
	if stackgrowthcurve2dribbonstartshape.TopEndY != stackgrowthcurve2dribbonstartshapeOther.TopEndY {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
	}
	if stackgrowthcurve2dribbonstartshape.TopRadiusX != stackgrowthcurve2dribbonstartshapeOther.TopRadiusX {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if stackgrowthcurve2dribbonstartshape.TopRadiusY != stackgrowthcurve2dribbonstartshapeOther.TopRadiusY {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if stackgrowthcurve2dribbonstartshape.TopXAxisRotation != stackgrowthcurve2dribbonstartshapeOther.TopXAxisRotation {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if stackgrowthcurve2dribbonstartshape.TopLargeArcFlag != stackgrowthcurve2dribbonstartshapeOther.TopLargeArcFlag {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if stackgrowthcurve2dribbonstartshape.TopSweepFlag != stackgrowthcurve2dribbonstartshapeOther.TopSweepFlag {
		diffs = append(diffs, stackgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongDiff(stage *Stage, stackgrowthcurve2dstarthalfwayarcshapeOther *StackGrowthCurve2DStartHalfwayArcShape) (diffs []string) {
	// insertion point for field diffs
	if stackgrowthcurve2dstarthalfwayarcshape.Name != stackgrowthcurve2dstarthalfwayarcshapeOther.Name {
		diffs = append(diffs, stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "Name"))
	}
	if stackgrowthcurve2dstarthalfwayarcshape.StartX != stackgrowthcurve2dstarthalfwayarcshapeOther.StartX {
		diffs = append(diffs, stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "StartX"))
	}
	if stackgrowthcurve2dstarthalfwayarcshape.StartY != stackgrowthcurve2dstarthalfwayarcshapeOther.StartY {
		diffs = append(diffs, stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "StartY"))
	}
	if stackgrowthcurve2dstarthalfwayarcshape.EndX != stackgrowthcurve2dstarthalfwayarcshapeOther.EndX {
		diffs = append(diffs, stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "EndX"))
	}
	if stackgrowthcurve2dstarthalfwayarcshape.EndY != stackgrowthcurve2dstarthalfwayarcshapeOther.EndY {
		diffs = append(diffs, stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "EndY"))
	}
	if stackgrowthcurve2dstarthalfwayarcshape.RadiusX != stackgrowthcurve2dstarthalfwayarcshapeOther.RadiusX {
		diffs = append(diffs, stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if stackgrowthcurve2dstarthalfwayarcshape.RadiusY != stackgrowthcurve2dstarthalfwayarcshapeOther.RadiusY {
		diffs = append(diffs, stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "RadiusY"))
	}
	if stackgrowthcurve2dstarthalfwayarcshape.XAxisRotation != stackgrowthcurve2dstarthalfwayarcshapeOther.XAxisRotation {
		diffs = append(diffs, stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if stackgrowthcurve2dstarthalfwayarcshape.LargeArcFlag != stackgrowthcurve2dstarthalfwayarcshapeOther.LargeArcFlag {
		diffs = append(diffs, stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if stackgrowthcurve2dstarthalfwayarcshape.SweepFlag != stackgrowthcurve2dstarthalfwayarcshapeOther.SweepFlag {
		diffs = append(diffs, stackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongDiff(stage *Stage, stackofgrowthcurve2dOther *StackOfGrowthCurve2D) (diffs []string) {
	// insertion point for field diffs
	if stackofgrowthcurve2d.Name != stackofgrowthcurve2dOther.Name {
		diffs = append(diffs, stackofgrowthcurve2d.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongDiff(stage *Stage, stackofgrowthcurve2dbygrowthvectorOther *StackOfGrowthCurve2DByGrowthVector) (diffs []string) {
	// insertion point for field diffs
	if stackofgrowthcurve2dbygrowthvector.Name != stackofgrowthcurve2dbygrowthvectorOther.Name {
		diffs = append(diffs, stackofgrowthcurve2dbygrowthvector.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongDiff(stage *Stage, stackofgrowthcurve2dribbonOther *StackOfGrowthCurve2DRibbon) (diffs []string) {
	// insertion point for field diffs
	if stackofgrowthcurve2dribbon.Name != stackofgrowthcurve2dribbonOther.Name {
		diffs = append(diffs, stackofgrowthcurve2dribbon.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongDiff(stage *Stage, stackofpartiallyrotatedtorusshapeOther *StackOfPartiallyRotatedTorusShape) (diffs []string) {
	// insertion point for field diffs
	if stackofpartiallyrotatedtorusshape.Name != stackofpartiallyrotatedtorusshapeOther.Name {
		diffs = append(diffs, stackofpartiallyrotatedtorusshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongDiff(stage *Stage, stackofrotatedgrowthcurve2dOther *StackOfRotatedGrowthCurve2D) (diffs []string) {
	// insertion point for field diffs
	if stackofrotatedgrowthcurve2d.Name != stackofrotatedgrowthcurve2dOther.Name {
		diffs = append(diffs, stackofrotatedgrowthcurve2d.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongDiff(stage *Stage, stackofrotatedgrowthcurve2dribbonOther *StackOfRotatedGrowthCurve2DRibbon) (diffs []string) {
	// insertion point for field diffs
	if stackofrotatedgrowthcurve2dribbon.Name != stackofrotatedgrowthcurve2dribbonOther.Name {
		diffs = append(diffs, stackofrotatedgrowthcurve2dribbon.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongDiff(stage *Stage, stackrotatedgrowthcurve2dendarcshapeOther *StackRotatedGrowthCurve2DEndArcShape) (diffs []string) {
	// insertion point for field diffs
	if stackrotatedgrowthcurve2dendarcshape.Name != stackrotatedgrowthcurve2dendarcshapeOther.Name {
		diffs = append(diffs, stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "Name"))
	}
	if stackrotatedgrowthcurve2dendarcshape.StartX != stackrotatedgrowthcurve2dendarcshapeOther.StartX {
		diffs = append(diffs, stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "StartX"))
	}
	if stackrotatedgrowthcurve2dendarcshape.StartY != stackrotatedgrowthcurve2dendarcshapeOther.StartY {
		diffs = append(diffs, stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "StartY"))
	}
	if stackrotatedgrowthcurve2dendarcshape.EndX != stackrotatedgrowthcurve2dendarcshapeOther.EndX {
		diffs = append(diffs, stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "EndX"))
	}
	if stackrotatedgrowthcurve2dendarcshape.EndY != stackrotatedgrowthcurve2dendarcshapeOther.EndY {
		diffs = append(diffs, stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "EndY"))
	}
	if stackrotatedgrowthcurve2dendarcshape.XAxisRotation != stackrotatedgrowthcurve2dendarcshapeOther.XAxisRotation {
		diffs = append(diffs, stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if stackrotatedgrowthcurve2dendarcshape.LargeArcFlag != stackrotatedgrowthcurve2dendarcshapeOther.LargeArcFlag {
		diffs = append(diffs, stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if stackrotatedgrowthcurve2dendarcshape.SweepFlag != stackrotatedgrowthcurve2dendarcshapeOther.SweepFlag {
		diffs = append(diffs, stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if stackrotatedgrowthcurve2dendarcshape.RadiusX != stackrotatedgrowthcurve2dendarcshapeOther.RadiusX {
		diffs = append(diffs, stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if stackrotatedgrowthcurve2dendarcshape.RadiusY != stackrotatedgrowthcurve2dendarcshapeOther.RadiusY {
		diffs = append(diffs, stackrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongDiff(stage *Stage, stackrotatedgrowthcurve2dribbonendshapeOther *StackRotatedGrowthCurve2DRibbonEndShape) (diffs []string) {
	// insertion point for field diffs
	if stackrotatedgrowthcurve2dribbonendshape.Name != stackrotatedgrowthcurve2dribbonendshapeOther.Name {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "Name"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.BottomStartX != stackrotatedgrowthcurve2dribbonendshapeOther.BottomStartX {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartX"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.BottomStartY != stackrotatedgrowthcurve2dribbonendshapeOther.BottomStartY {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomStartY"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.BottomEndX != stackrotatedgrowthcurve2dribbonendshapeOther.BottomEndX {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndX"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.BottomEndY != stackrotatedgrowthcurve2dribbonendshapeOther.BottomEndY {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomEndY"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.BottomRadiusX != stackrotatedgrowthcurve2dribbonendshapeOther.BottomRadiusX {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.BottomRadiusY != stackrotatedgrowthcurve2dribbonendshapeOther.BottomRadiusY {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.BottomXAxisRotation != stackrotatedgrowthcurve2dribbonendshapeOther.BottomXAxisRotation {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.BottomLargeArcFlag != stackrotatedgrowthcurve2dribbonendshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.BottomSweepFlag != stackrotatedgrowthcurve2dribbonendshapeOther.BottomSweepFlag {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.TopStartX != stackrotatedgrowthcurve2dribbonendshapeOther.TopStartX {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartX"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.TopStartY != stackrotatedgrowthcurve2dribbonendshapeOther.TopStartY {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopStartY"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.TopEndX != stackrotatedgrowthcurve2dribbonendshapeOther.TopEndX {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndX"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.TopEndY != stackrotatedgrowthcurve2dribbonendshapeOther.TopEndY {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopEndY"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.TopRadiusX != stackrotatedgrowthcurve2dribbonendshapeOther.TopRadiusX {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.TopRadiusY != stackrotatedgrowthcurve2dribbonendshapeOther.TopRadiusY {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.TopXAxisRotation != stackrotatedgrowthcurve2dribbonendshapeOther.TopXAxisRotation {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.TopLargeArcFlag != stackrotatedgrowthcurve2dribbonendshapeOther.TopLargeArcFlag {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if stackrotatedgrowthcurve2dribbonendshape.TopSweepFlag != stackrotatedgrowthcurve2dribbonendshapeOther.TopSweepFlag {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonendshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongDiff(stage *Stage, stackrotatedgrowthcurve2dribbonstartshapeOther *StackRotatedGrowthCurve2DRibbonStartShape) (diffs []string) {
	// insertion point for field diffs
	if stackrotatedgrowthcurve2dribbonstartshape.Name != stackrotatedgrowthcurve2dribbonstartshapeOther.Name {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "Name"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.BottomStartX != stackrotatedgrowthcurve2dribbonstartshapeOther.BottomStartX {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartX"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.BottomStartY != stackrotatedgrowthcurve2dribbonstartshapeOther.BottomStartY {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomStartY"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.BottomEndX != stackrotatedgrowthcurve2dribbonstartshapeOther.BottomEndX {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndX"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.BottomEndY != stackrotatedgrowthcurve2dribbonstartshapeOther.BottomEndY {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomEndY"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.BottomRadiusX != stackrotatedgrowthcurve2dribbonstartshapeOther.BottomRadiusX {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusX"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.BottomRadiusY != stackrotatedgrowthcurve2dribbonstartshapeOther.BottomRadiusY {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomRadiusY"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.BottomXAxisRotation != stackrotatedgrowthcurve2dribbonstartshapeOther.BottomXAxisRotation {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomXAxisRotation"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.BottomLargeArcFlag != stackrotatedgrowthcurve2dribbonstartshapeOther.BottomLargeArcFlag {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomLargeArcFlag"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.BottomSweepFlag != stackrotatedgrowthcurve2dribbonstartshapeOther.BottomSweepFlag {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "BottomSweepFlag"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.TopStartX != stackrotatedgrowthcurve2dribbonstartshapeOther.TopStartX {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartX"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.TopStartY != stackrotatedgrowthcurve2dribbonstartshapeOther.TopStartY {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopStartY"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.TopEndX != stackrotatedgrowthcurve2dribbonstartshapeOther.TopEndX {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndX"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.TopEndY != stackrotatedgrowthcurve2dribbonstartshapeOther.TopEndY {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopEndY"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.TopRadiusX != stackrotatedgrowthcurve2dribbonstartshapeOther.TopRadiusX {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusX"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.TopRadiusY != stackrotatedgrowthcurve2dribbonstartshapeOther.TopRadiusY {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopRadiusY"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.TopXAxisRotation != stackrotatedgrowthcurve2dribbonstartshapeOther.TopXAxisRotation {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopXAxisRotation"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.TopLargeArcFlag != stackrotatedgrowthcurve2dribbonstartshapeOther.TopLargeArcFlag {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopLargeArcFlag"))
	}
	if stackrotatedgrowthcurve2dribbonstartshape.TopSweepFlag != stackrotatedgrowthcurve2dribbonstartshapeOther.TopSweepFlag {
		diffs = append(diffs, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallField(stage, "TopSweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongDiff(stage *Stage, stackrotatedgrowthcurve2dstartarcshapeOther *StackRotatedGrowthCurve2DStartArcShape) (diffs []string) {
	// insertion point for field diffs
	if stackrotatedgrowthcurve2dstartarcshape.Name != stackrotatedgrowthcurve2dstartarcshapeOther.Name {
		diffs = append(diffs, stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "Name"))
	}
	if stackrotatedgrowthcurve2dstartarcshape.StartX != stackrotatedgrowthcurve2dstartarcshapeOther.StartX {
		diffs = append(diffs, stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "StartX"))
	}
	if stackrotatedgrowthcurve2dstartarcshape.StartY != stackrotatedgrowthcurve2dstartarcshapeOther.StartY {
		diffs = append(diffs, stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "StartY"))
	}
	if stackrotatedgrowthcurve2dstartarcshape.EndX != stackrotatedgrowthcurve2dstartarcshapeOther.EndX {
		diffs = append(diffs, stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "EndX"))
	}
	if stackrotatedgrowthcurve2dstartarcshape.EndY != stackrotatedgrowthcurve2dstartarcshapeOther.EndY {
		diffs = append(diffs, stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "EndY"))
	}
	if stackrotatedgrowthcurve2dstartarcshape.XAxisRotation != stackrotatedgrowthcurve2dstartarcshapeOther.XAxisRotation {
		diffs = append(diffs, stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if stackrotatedgrowthcurve2dstartarcshape.LargeArcFlag != stackrotatedgrowthcurve2dstartarcshapeOther.LargeArcFlag {
		diffs = append(diffs, stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if stackrotatedgrowthcurve2dstartarcshape.SweepFlag != stackrotatedgrowthcurve2dstartarcshapeOther.SweepFlag {
		diffs = append(diffs, stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if stackrotatedgrowthcurve2dstartarcshape.RadiusX != stackrotatedgrowthcurve2dstartarcshapeOther.RadiusX {
		diffs = append(diffs, stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if stackrotatedgrowthcurve2dstartarcshape.RadiusY != stackrotatedgrowthcurve2dstartarcshapeOther.RadiusY {
		diffs = append(diffs, stackrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (startarcshape *StartArcShape) GongDiff(stage *Stage, startarcshapeOther *StartArcShape) (diffs []string) {
	// insertion point for field diffs
	if startarcshape.Name != startarcshapeOther.Name {
		diffs = append(diffs, startarcshape.GongMarshallField(stage, "Name"))
	}
	if startarcshape.StartX != startarcshapeOther.StartX {
		diffs = append(diffs, startarcshape.GongMarshallField(stage, "StartX"))
	}
	if startarcshape.StartY != startarcshapeOther.StartY {
		diffs = append(diffs, startarcshape.GongMarshallField(stage, "StartY"))
	}
	if startarcshape.EndX != startarcshapeOther.EndX {
		diffs = append(diffs, startarcshape.GongMarshallField(stage, "EndX"))
	}
	if startarcshape.EndY != startarcshapeOther.EndY {
		diffs = append(diffs, startarcshape.GongMarshallField(stage, "EndY"))
	}
	if startarcshape.XAxisRotation != startarcshapeOther.XAxisRotation {
		diffs = append(diffs, startarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if startarcshape.LargeArcFlag != startarcshapeOther.LargeArcFlag {
		diffs = append(diffs, startarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if startarcshape.SweepFlag != startarcshapeOther.SweepFlag {
		diffs = append(diffs, startarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if startarcshape.RadiusX != startarcshapeOther.RadiusX {
		diffs = append(diffs, startarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if startarcshape.RadiusY != startarcshapeOther.RadiusY {
		diffs = append(diffs, startarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (startarcshapegrid *StartArcShapeGrid) GongDiff(stage *Stage, startarcshapegridOther *StartArcShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if startarcshapegrid.Name != startarcshapegridOther.Name {
		diffs = append(diffs, startarcshapegrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (starthalfwayarcshape *StartHalfwayArcShape) GongDiff(stage *Stage, starthalfwayarcshapeOther *StartHalfwayArcShape) (diffs []string) {
	// insertion point for field diffs
	if starthalfwayarcshape.Name != starthalfwayarcshapeOther.Name {
		diffs = append(diffs, starthalfwayarcshape.GongMarshallField(stage, "Name"))
	}
	if starthalfwayarcshape.StartX != starthalfwayarcshapeOther.StartX {
		diffs = append(diffs, starthalfwayarcshape.GongMarshallField(stage, "StartX"))
	}
	if starthalfwayarcshape.StartY != starthalfwayarcshapeOther.StartY {
		diffs = append(diffs, starthalfwayarcshape.GongMarshallField(stage, "StartY"))
	}
	if starthalfwayarcshape.EndX != starthalfwayarcshapeOther.EndX {
		diffs = append(diffs, starthalfwayarcshape.GongMarshallField(stage, "EndX"))
	}
	if starthalfwayarcshape.EndY != starthalfwayarcshapeOther.EndY {
		diffs = append(diffs, starthalfwayarcshape.GongMarshallField(stage, "EndY"))
	}
	if starthalfwayarcshape.RadiusX != starthalfwayarcshapeOther.RadiusX {
		diffs = append(diffs, starthalfwayarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if starthalfwayarcshape.RadiusY != starthalfwayarcshapeOther.RadiusY {
		diffs = append(diffs, starthalfwayarcshape.GongMarshallField(stage, "RadiusY"))
	}
	if starthalfwayarcshape.XAxisRotation != starthalfwayarcshapeOther.XAxisRotation {
		diffs = append(diffs, starthalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if starthalfwayarcshape.LargeArcFlag != starthalfwayarcshapeOther.LargeArcFlag {
		diffs = append(diffs, starthalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if starthalfwayarcshape.SweepFlag != starthalfwayarcshapeOther.SweepFlag {
		diffs = append(diffs, starthalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongDiff(stage *Stage, starthalfwayarcshapegridOther *StartHalfwayArcShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if starthalfwayarcshapegrid.Name != starthalfwayarcshapegridOther.Name {
		diffs = append(diffs, starthalfwayarcshapegrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stemcylinder3dshape *StemCylinder3DShape) GongDiff(stage *Stage, stemcylinder3dshapeOther *StemCylinder3DShape) (diffs []string) {
	// insertion point for field diffs
	if stemcylinder3dshape.Name != stemcylinder3dshapeOther.Name {
		diffs = append(diffs, stemcylinder3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stool2ddiagram *Stool2DDiagram) GongDiff(stage *Stage, stool2ddiagramOther *Stool2DDiagram) (diffs []string) {
	// insertion point for field diffs
	if stool2ddiagram.Name != stool2ddiagramOther.Name {
		diffs = append(diffs, stool2ddiagram.GongMarshallField(stage, "Name"))
	}
	if stool2ddiagram.Zoom != stool2ddiagramOther.Zoom {
		diffs = append(diffs, stool2ddiagram.GongMarshallField(stage, "Zoom"))
	}
	if stool2ddiagram.IsHiddenAxesShape != stool2ddiagramOther.IsHiddenAxesShape {
		diffs = append(diffs, stool2ddiagram.GongMarshallField(stage, "IsHiddenAxesShape"))
	}
	if stool2ddiagram.IsChecked != stool2ddiagramOther.IsChecked {
		diffs = append(diffs, stool2ddiagram.GongMarshallField(stage, "IsChecked"))
	}
	if stool2ddiagram.ComputedPrefix != stool2ddiagramOther.ComputedPrefix {
		diffs = append(diffs, stool2ddiagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if stool2ddiagram.IsExpanded != stool2ddiagramOther.IsExpanded {
		diffs = append(diffs, stool2ddiagram.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stool3ddiagram *Stool3DDiagram) GongDiff(stage *Stage, stool3ddiagramOther *Stool3DDiagram) (diffs []string) {
	// insertion point for field diffs
	if stool3ddiagram.Name != stool3ddiagramOther.Name {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "Name"))
	}
	if stool3ddiagram.IsHiddenSeatTopCurveShape != stool3ddiagramOther.IsHiddenSeatTopCurveShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenSeatTopCurveShape"))
	}
	if stool3ddiagram.IsHiddenRotatedSeatTopCurveShape != stool3ddiagramOther.IsHiddenRotatedSeatTopCurveShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenRotatedSeatTopCurveShape"))
	}
	if stool3ddiagram.IsHiddenSeatBottomCurveShape != stool3ddiagramOther.IsHiddenSeatBottomCurveShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenSeatBottomCurveShape"))
	}
	if stool3ddiagram.IsHiddenRotatedSeatBottomCurveShape != stool3ddiagramOther.IsHiddenRotatedSeatBottomCurveShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenRotatedSeatBottomCurveShape"))
	}
	if stool3ddiagram.IsHiddenTorus3DShape != stool3ddiagramOther.IsHiddenTorus3DShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenTorus3DShape"))
	}
	if stool3ddiagram.IsHiddenRotatedTorusShape != stool3ddiagramOther.IsHiddenRotatedTorusShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenRotatedTorusShape"))
	}
	if stool3ddiagram.IsHiddenSampledPoints3DShape != stool3ddiagramOther.IsHiddenSampledPoints3DShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenSampledPoints3DShape"))
	}
	if (stool3ddiagram.SampledPoints3DShape == nil) != (stool3ddiagramOther.SampledPoints3DShape == nil) {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
	} else if stool3ddiagram.SampledPoints3DShape != nil && stool3ddiagramOther.SampledPoints3DShape != nil {
		if stool3ddiagram.SampledPoints3DShape != stool3ddiagramOther.SampledPoints3DShape {
			diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
		}
	}
	if stool3ddiagram.IsHiddenRotatedSampledPoints3DShape != stool3ddiagramOther.IsHiddenRotatedSampledPoints3DShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenRotatedSampledPoints3DShape"))
	}
	if stool3ddiagram.IsHiddenEyeSampledPoints3DShape != stool3ddiagramOther.IsHiddenEyeSampledPoints3DShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenEyeSampledPoints3DShape"))
	}
	if stool3ddiagram.IsHiddenEyeCornersSampledPoints3DShape != stool3ddiagramOther.IsHiddenEyeCornersSampledPoints3DShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenEyeCornersSampledPoints3DShape"))
	}
	if stool3ddiagram.IsHiddenEye3DShape != stool3ddiagramOther.IsHiddenEye3DShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenEye3DShape"))
	}
	if stool3ddiagram.IsHiddenEyeSeatBottomCurveShape != stool3ddiagramOther.IsHiddenEyeSeatBottomCurveShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenEyeSeatBottomCurveShape"))
	}
	if stool3ddiagram.IsHiddenEyeStoolBottomCurveShape != stool3ddiagramOther.IsHiddenEyeStoolBottomCurveShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenEyeStoolBottomCurveShape"))
	}
	if stool3ddiagram.IsHiddenSeat3DShape != stool3ddiagramOther.IsHiddenSeat3DShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenSeat3DShape"))
	}
	if stool3ddiagram.IsHiddenEyeVolume3DShape != stool3ddiagramOther.IsHiddenEyeVolume3DShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenEyeVolume3DShape"))
	}
	if stool3ddiagram.IsHiddenSeatAndLegs3DShape != stool3ddiagramOther.IsHiddenSeatAndLegs3DShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenSeatAndLegs3DShape"))
	}
	if stool3ddiagram.IsHiddenRotatedSeatAndLegs3DShape != stool3ddiagramOther.IsHiddenRotatedSeatAndLegs3DShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenRotatedSeatAndLegs3DShape"))
	}
	if stool3ddiagram.IsHiddenTiledFloor3DShape != stool3ddiagramOther.IsHiddenTiledFloor3DShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsHiddenTiledFloor3DShape"))
	}
	if (stool3ddiagram.Rendered3DShape == nil) != (stool3ddiagramOther.Rendered3DShape == nil) {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
	} else if stool3ddiagram.Rendered3DShape != nil && stool3ddiagramOther.Rendered3DShape != nil {
		if stool3ddiagram.Rendered3DShape != stool3ddiagramOther.Rendered3DShape {
			diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
		}
	}
	if stool3ddiagram.IsChecked != stool3ddiagramOther.IsChecked {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsChecked"))
	}
	if stool3ddiagram.ComputedPrefix != stool3ddiagramOther.ComputedPrefix {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if stool3ddiagram.IsExpanded != stool3ddiagramOther.IsExpanded {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stoolabstract *StoolAbstract) GongDiff(stage *Stage, stoolabstractOther *StoolAbstract) (diffs []string) {
	// insertion point for field diffs
	if stoolabstract.Name != stoolabstractOther.Name {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "Name"))
	}
	if stoolabstract.RadialRepetitions != stoolabstractOther.RadialRepetitions {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RadialRepetitions"))
	}
	if stoolabstract.Transparency != stoolabstractOther.Transparency {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "Transparency"))
	}
	if stoolabstract.RelativeTubeDiameter != stoolabstractOther.RelativeTubeDiameter {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RelativeTubeDiameter"))
	}
	if stoolabstract.RelativeHeight3DTorus != stoolabstractOther.RelativeHeight3DTorus {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RelativeHeight3DTorus"))
	}
	if stoolabstract.StoolTorusVerticalScale != stoolabstractOther.StoolTorusVerticalScale {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "StoolTorusVerticalScale"))
	}
	if stoolabstract.RelativeHeight != stoolabstractOther.RelativeHeight {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RelativeHeight"))
	}
	if stoolabstract.RelativeSeatThickness != stoolabstractOther.RelativeSeatThickness {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RelativeSeatThickness"))
	}
	if stoolabstract.ProjectionAngle != stoolabstractOther.ProjectionAngle {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "ProjectionAngle"))
	}
	if stoolabstract.RelativeEyeSeparationCriteria != stoolabstractOther.RelativeEyeSeparationCriteria {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RelativeEyeSeparationCriteria"))
	}
	if stoolabstract.RelativeEyeCornerControlVectorStrength != stoolabstractOther.RelativeEyeCornerControlVectorStrength {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RelativeEyeCornerControlVectorStrength"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tiledfloor3dshape *TiledFloor3DShape) GongDiff(stage *Stage, tiledfloor3dshapeOther *TiledFloor3DShape) (diffs []string) {
	// insertion point for field diffs
	if tiledfloor3dshape.Name != tiledfloor3dshapeOther.Name {
		diffs = append(diffs, tiledfloor3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topendarcshape *TopEndArcShape) GongDiff(stage *Stage, topendarcshapeOther *TopEndArcShape) (diffs []string) {
	// insertion point for field diffs
	if topendarcshape.Name != topendarcshapeOther.Name {
		diffs = append(diffs, topendarcshape.GongMarshallField(stage, "Name"))
	}
	if topendarcshape.StartX != topendarcshapeOther.StartX {
		diffs = append(diffs, topendarcshape.GongMarshallField(stage, "StartX"))
	}
	if topendarcshape.StartY != topendarcshapeOther.StartY {
		diffs = append(diffs, topendarcshape.GongMarshallField(stage, "StartY"))
	}
	if topendarcshape.EndX != topendarcshapeOther.EndX {
		diffs = append(diffs, topendarcshape.GongMarshallField(stage, "EndX"))
	}
	if topendarcshape.EndY != topendarcshapeOther.EndY {
		diffs = append(diffs, topendarcshape.GongMarshallField(stage, "EndY"))
	}
	if topendarcshape.XAxisRotation != topendarcshapeOther.XAxisRotation {
		diffs = append(diffs, topendarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if topendarcshape.LargeArcFlag != topendarcshapeOther.LargeArcFlag {
		diffs = append(diffs, topendarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topendarcshape.SweepFlag != topendarcshapeOther.SweepFlag {
		diffs = append(diffs, topendarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if topendarcshape.RadiusX != topendarcshapeOther.RadiusX {
		diffs = append(diffs, topendarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if topendarcshape.RadiusY != topendarcshapeOther.RadiusY {
		diffs = append(diffs, topendarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topendarcshapegrid *TopEndArcShapeGrid) GongDiff(stage *Stage, topendarcshapegridOther *TopEndArcShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if topendarcshapegrid.Name != topendarcshapegridOther.Name {
		diffs = append(diffs, topendarcshapegrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongDiff(stage *Stage, topendhalfwayarcshapeOther *TopEndHalfwayArcShape) (diffs []string) {
	// insertion point for field diffs
	if topendhalfwayarcshape.Name != topendhalfwayarcshapeOther.Name {
		diffs = append(diffs, topendhalfwayarcshape.GongMarshallField(stage, "Name"))
	}
	if topendhalfwayarcshape.StartX != topendhalfwayarcshapeOther.StartX {
		diffs = append(diffs, topendhalfwayarcshape.GongMarshallField(stage, "StartX"))
	}
	if topendhalfwayarcshape.StartY != topendhalfwayarcshapeOther.StartY {
		diffs = append(diffs, topendhalfwayarcshape.GongMarshallField(stage, "StartY"))
	}
	if topendhalfwayarcshape.EndX != topendhalfwayarcshapeOther.EndX {
		diffs = append(diffs, topendhalfwayarcshape.GongMarshallField(stage, "EndX"))
	}
	if topendhalfwayarcshape.EndY != topendhalfwayarcshapeOther.EndY {
		diffs = append(diffs, topendhalfwayarcshape.GongMarshallField(stage, "EndY"))
	}
	if topendhalfwayarcshape.RadiusX != topendhalfwayarcshapeOther.RadiusX {
		diffs = append(diffs, topendhalfwayarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if topendhalfwayarcshape.RadiusY != topendhalfwayarcshapeOther.RadiusY {
		diffs = append(diffs, topendhalfwayarcshape.GongMarshallField(stage, "RadiusY"))
	}
	if topendhalfwayarcshape.XAxisRotation != topendhalfwayarcshapeOther.XAxisRotation {
		diffs = append(diffs, topendhalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if topendhalfwayarcshape.LargeArcFlag != topendhalfwayarcshapeOther.LargeArcFlag {
		diffs = append(diffs, topendhalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topendhalfwayarcshape.SweepFlag != topendhalfwayarcshapeOther.SweepFlag {
		diffs = append(diffs, topendhalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongDiff(stage *Stage, topendhalfwayarcshapegridOther *TopEndHalfwayArcShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if topendhalfwayarcshapegrid.Name != topendhalfwayarcshapegridOther.Name {
		diffs = append(diffs, topendhalfwayarcshapegrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topgrowthcurve2d *TopGrowthCurve2D) GongDiff(stage *Stage, topgrowthcurve2dOther *TopGrowthCurve2D) (diffs []string) {
	// insertion point for field diffs
	if topgrowthcurve2d.Name != topgrowthcurve2dOther.Name {
		diffs = append(diffs, topgrowthcurve2d.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topmidarcvectorshape *TopMidArcVectorShape) GongDiff(stage *Stage, topmidarcvectorshapeOther *TopMidArcVectorShape) (diffs []string) {
	// insertion point for field diffs
	if topmidarcvectorshape.Name != topmidarcvectorshapeOther.Name {
		diffs = append(diffs, topmidarcvectorshape.GongMarshallField(stage, "Name"))
	}
	if topmidarcvectorshape.StartX != topmidarcvectorshapeOther.StartX {
		diffs = append(diffs, topmidarcvectorshape.GongMarshallField(stage, "StartX"))
	}
	if topmidarcvectorshape.StartY != topmidarcvectorshapeOther.StartY {
		diffs = append(diffs, topmidarcvectorshape.GongMarshallField(stage, "StartY"))
	}
	if topmidarcvectorshape.EndX != topmidarcvectorshapeOther.EndX {
		diffs = append(diffs, topmidarcvectorshape.GongMarshallField(stage, "EndX"))
	}
	if topmidarcvectorshape.EndY != topmidarcvectorshapeOther.EndY {
		diffs = append(diffs, topmidarcvectorshape.GongMarshallField(stage, "EndY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongDiff(stage *Stage, topmidarcvectorshapegridOther *TopMidArcVectorShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if topmidarcvectorshapegrid.Name != topmidarcvectorshapegridOther.Name {
		diffs = append(diffs, topmidarcvectorshapegrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongDiff(stage *Stage, topstackgrowthcurve2dendhalfwayarcshapeOther *TopStackGrowthCurve2DEndHalfwayArcShape) (diffs []string) {
	// insertion point for field diffs
	if topstackgrowthcurve2dendhalfwayarcshape.Name != topstackgrowthcurve2dendhalfwayarcshapeOther.Name {
		diffs = append(diffs, topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "Name"))
	}
	if topstackgrowthcurve2dendhalfwayarcshape.StartX != topstackgrowthcurve2dendhalfwayarcshapeOther.StartX {
		diffs = append(diffs, topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "StartX"))
	}
	if topstackgrowthcurve2dendhalfwayarcshape.StartY != topstackgrowthcurve2dendhalfwayarcshapeOther.StartY {
		diffs = append(diffs, topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "StartY"))
	}
	if topstackgrowthcurve2dendhalfwayarcshape.EndX != topstackgrowthcurve2dendhalfwayarcshapeOther.EndX {
		diffs = append(diffs, topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "EndX"))
	}
	if topstackgrowthcurve2dendhalfwayarcshape.EndY != topstackgrowthcurve2dendhalfwayarcshapeOther.EndY {
		diffs = append(diffs, topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "EndY"))
	}
	if topstackgrowthcurve2dendhalfwayarcshape.RadiusX != topstackgrowthcurve2dendhalfwayarcshapeOther.RadiusX {
		diffs = append(diffs, topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if topstackgrowthcurve2dendhalfwayarcshape.RadiusY != topstackgrowthcurve2dendhalfwayarcshapeOther.RadiusY {
		diffs = append(diffs, topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "RadiusY"))
	}
	if topstackgrowthcurve2dendhalfwayarcshape.XAxisRotation != topstackgrowthcurve2dendhalfwayarcshapeOther.XAxisRotation {
		diffs = append(diffs, topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if topstackgrowthcurve2dendhalfwayarcshape.LargeArcFlag != topstackgrowthcurve2dendhalfwayarcshapeOther.LargeArcFlag {
		diffs = append(diffs, topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topstackgrowthcurve2dendhalfwayarcshape.SweepFlag != topstackgrowthcurve2dendhalfwayarcshapeOther.SweepFlag {
		diffs = append(diffs, topstackgrowthcurve2dendhalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongDiff(stage *Stage, topstackgrowthcurve2dstarthalfwayarcshapeOther *TopStackGrowthCurve2DStartHalfwayArcShape) (diffs []string) {
	// insertion point for field diffs
	if topstackgrowthcurve2dstarthalfwayarcshape.Name != topstackgrowthcurve2dstarthalfwayarcshapeOther.Name {
		diffs = append(diffs, topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "Name"))
	}
	if topstackgrowthcurve2dstarthalfwayarcshape.StartX != topstackgrowthcurve2dstarthalfwayarcshapeOther.StartX {
		diffs = append(diffs, topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "StartX"))
	}
	if topstackgrowthcurve2dstarthalfwayarcshape.StartY != topstackgrowthcurve2dstarthalfwayarcshapeOther.StartY {
		diffs = append(diffs, topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "StartY"))
	}
	if topstackgrowthcurve2dstarthalfwayarcshape.EndX != topstackgrowthcurve2dstarthalfwayarcshapeOther.EndX {
		diffs = append(diffs, topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "EndX"))
	}
	if topstackgrowthcurve2dstarthalfwayarcshape.EndY != topstackgrowthcurve2dstarthalfwayarcshapeOther.EndY {
		diffs = append(diffs, topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "EndY"))
	}
	if topstackgrowthcurve2dstarthalfwayarcshape.RadiusX != topstackgrowthcurve2dstarthalfwayarcshapeOther.RadiusX {
		diffs = append(diffs, topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if topstackgrowthcurve2dstarthalfwayarcshape.RadiusY != topstackgrowthcurve2dstarthalfwayarcshapeOther.RadiusY {
		diffs = append(diffs, topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "RadiusY"))
	}
	if topstackgrowthcurve2dstarthalfwayarcshape.XAxisRotation != topstackgrowthcurve2dstarthalfwayarcshapeOther.XAxisRotation {
		diffs = append(diffs, topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if topstackgrowthcurve2dstarthalfwayarcshape.LargeArcFlag != topstackgrowthcurve2dstarthalfwayarcshapeOther.LargeArcFlag {
		diffs = append(diffs, topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topstackgrowthcurve2dstarthalfwayarcshape.SweepFlag != topstackgrowthcurve2dstarthalfwayarcshapeOther.SweepFlag {
		diffs = append(diffs, topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongDiff(stage *Stage, topstackofgrowthcurve2dOther *TopStackOfGrowthCurve2D) (diffs []string) {
	// insertion point for field diffs
	if topstackofgrowthcurve2d.Name != topstackofgrowthcurve2dOther.Name {
		diffs = append(diffs, topstackofgrowthcurve2d.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongDiff(stage *Stage, topstackofrotatedgrowthcurve2dOther *TopStackOfRotatedGrowthCurve2D) (diffs []string) {
	// insertion point for field diffs
	if topstackofrotatedgrowthcurve2d.Name != topstackofrotatedgrowthcurve2dOther.Name {
		diffs = append(diffs, topstackofrotatedgrowthcurve2d.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongDiff(stage *Stage, topstackofrotatedgrowthcurve2dendarcshapeOther *TopStackOfRotatedGrowthCurve2DEndArcShape) (diffs []string) {
	// insertion point for field diffs
	if topstackofrotatedgrowthcurve2dendarcshape.Name != topstackofrotatedgrowthcurve2dendarcshapeOther.Name {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "Name"))
	}
	if topstackofrotatedgrowthcurve2dendarcshape.StartX != topstackofrotatedgrowthcurve2dendarcshapeOther.StartX {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "StartX"))
	}
	if topstackofrotatedgrowthcurve2dendarcshape.StartY != topstackofrotatedgrowthcurve2dendarcshapeOther.StartY {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "StartY"))
	}
	if topstackofrotatedgrowthcurve2dendarcshape.EndX != topstackofrotatedgrowthcurve2dendarcshapeOther.EndX {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "EndX"))
	}
	if topstackofrotatedgrowthcurve2dendarcshape.EndY != topstackofrotatedgrowthcurve2dendarcshapeOther.EndY {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "EndY"))
	}
	if topstackofrotatedgrowthcurve2dendarcshape.XAxisRotation != topstackofrotatedgrowthcurve2dendarcshapeOther.XAxisRotation {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if topstackofrotatedgrowthcurve2dendarcshape.LargeArcFlag != topstackofrotatedgrowthcurve2dendarcshapeOther.LargeArcFlag {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topstackofrotatedgrowthcurve2dendarcshape.SweepFlag != topstackofrotatedgrowthcurve2dendarcshapeOther.SweepFlag {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if topstackofrotatedgrowthcurve2dendarcshape.RadiusX != topstackofrotatedgrowthcurve2dendarcshapeOther.RadiusX {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if topstackofrotatedgrowthcurve2dendarcshape.RadiusY != topstackofrotatedgrowthcurve2dendarcshapeOther.RadiusY {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dendarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongDiff(stage *Stage, topstackofrotatedgrowthcurve2dstartarcshapeOther *TopStackOfRotatedGrowthCurve2DStartArcShape) (diffs []string) {
	// insertion point for field diffs
	if topstackofrotatedgrowthcurve2dstartarcshape.Name != topstackofrotatedgrowthcurve2dstartarcshapeOther.Name {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "Name"))
	}
	if topstackofrotatedgrowthcurve2dstartarcshape.StartX != topstackofrotatedgrowthcurve2dstartarcshapeOther.StartX {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "StartX"))
	}
	if topstackofrotatedgrowthcurve2dstartarcshape.StartY != topstackofrotatedgrowthcurve2dstartarcshapeOther.StartY {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "StartY"))
	}
	if topstackofrotatedgrowthcurve2dstartarcshape.EndX != topstackofrotatedgrowthcurve2dstartarcshapeOther.EndX {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "EndX"))
	}
	if topstackofrotatedgrowthcurve2dstartarcshape.EndY != topstackofrotatedgrowthcurve2dstartarcshapeOther.EndY {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "EndY"))
	}
	if topstackofrotatedgrowthcurve2dstartarcshape.XAxisRotation != topstackofrotatedgrowthcurve2dstartarcshapeOther.XAxisRotation {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if topstackofrotatedgrowthcurve2dstartarcshape.LargeArcFlag != topstackofrotatedgrowthcurve2dstartarcshapeOther.LargeArcFlag {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topstackofrotatedgrowthcurve2dstartarcshape.SweepFlag != topstackofrotatedgrowthcurve2dstartarcshapeOther.SweepFlag {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if topstackofrotatedgrowthcurve2dstartarcshape.RadiusX != topstackofrotatedgrowthcurve2dstartarcshapeOther.RadiusX {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if topstackofrotatedgrowthcurve2dstartarcshape.RadiusY != topstackofrotatedgrowthcurve2dstartarcshapeOther.RadiusY {
		diffs = append(diffs, topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstartarcshape *TopStartArcShape) GongDiff(stage *Stage, topstartarcshapeOther *TopStartArcShape) (diffs []string) {
	// insertion point for field diffs
	if topstartarcshape.Name != topstartarcshapeOther.Name {
		diffs = append(diffs, topstartarcshape.GongMarshallField(stage, "Name"))
	}
	if topstartarcshape.StartX != topstartarcshapeOther.StartX {
		diffs = append(diffs, topstartarcshape.GongMarshallField(stage, "StartX"))
	}
	if topstartarcshape.StartY != topstartarcshapeOther.StartY {
		diffs = append(diffs, topstartarcshape.GongMarshallField(stage, "StartY"))
	}
	if topstartarcshape.EndX != topstartarcshapeOther.EndX {
		diffs = append(diffs, topstartarcshape.GongMarshallField(stage, "EndX"))
	}
	if topstartarcshape.EndY != topstartarcshapeOther.EndY {
		diffs = append(diffs, topstartarcshape.GongMarshallField(stage, "EndY"))
	}
	if topstartarcshape.XAxisRotation != topstartarcshapeOther.XAxisRotation {
		diffs = append(diffs, topstartarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if topstartarcshape.LargeArcFlag != topstartarcshapeOther.LargeArcFlag {
		diffs = append(diffs, topstartarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topstartarcshape.SweepFlag != topstartarcshapeOther.SweepFlag {
		diffs = append(diffs, topstartarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if topstartarcshape.RadiusX != topstartarcshapeOther.RadiusX {
		diffs = append(diffs, topstartarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if topstartarcshape.RadiusY != topstartarcshapeOther.RadiusY {
		diffs = append(diffs, topstartarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstartarcshapegrid *TopStartArcShapeGrid) GongDiff(stage *Stage, topstartarcshapegridOther *TopStartArcShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if topstartarcshapegrid.Name != topstartarcshapegridOther.Name {
		diffs = append(diffs, topstartarcshapegrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongDiff(stage *Stage, topstarthalfwayarcshapeOther *TopStartHalfwayArcShape) (diffs []string) {
	// insertion point for field diffs
	if topstarthalfwayarcshape.Name != topstarthalfwayarcshapeOther.Name {
		diffs = append(diffs, topstarthalfwayarcshape.GongMarshallField(stage, "Name"))
	}
	if topstarthalfwayarcshape.StartX != topstarthalfwayarcshapeOther.StartX {
		diffs = append(diffs, topstarthalfwayarcshape.GongMarshallField(stage, "StartX"))
	}
	if topstarthalfwayarcshape.StartY != topstarthalfwayarcshapeOther.StartY {
		diffs = append(diffs, topstarthalfwayarcshape.GongMarshallField(stage, "StartY"))
	}
	if topstarthalfwayarcshape.EndX != topstarthalfwayarcshapeOther.EndX {
		diffs = append(diffs, topstarthalfwayarcshape.GongMarshallField(stage, "EndX"))
	}
	if topstarthalfwayarcshape.EndY != topstarthalfwayarcshapeOther.EndY {
		diffs = append(diffs, topstarthalfwayarcshape.GongMarshallField(stage, "EndY"))
	}
	if topstarthalfwayarcshape.RadiusX != topstarthalfwayarcshapeOther.RadiusX {
		diffs = append(diffs, topstarthalfwayarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if topstarthalfwayarcshape.RadiusY != topstarthalfwayarcshapeOther.RadiusY {
		diffs = append(diffs, topstarthalfwayarcshape.GongMarshallField(stage, "RadiusY"))
	}
	if topstarthalfwayarcshape.XAxisRotation != topstarthalfwayarcshapeOther.XAxisRotation {
		diffs = append(diffs, topstarthalfwayarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if topstarthalfwayarcshape.LargeArcFlag != topstarthalfwayarcshapeOther.LargeArcFlag {
		diffs = append(diffs, topstarthalfwayarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topstarthalfwayarcshape.SweepFlag != topstarthalfwayarcshapeOther.SweepFlag {
		diffs = append(diffs, topstarthalfwayarcshape.GongMarshallField(stage, "SweepFlag"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongDiff(stage *Stage, topstarthalfwayarcshapegridOther *TopStartHalfwayArcShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if topstarthalfwayarcshapegrid.Name != topstarthalfwayarcshapegridOther.Name {
		diffs = append(diffs, topstarthalfwayarcshapegrid.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (torus3dshape *Torus3DShape) GongDiff(stage *Stage, torus3dshapeOther *Torus3DShape) (diffs []string) {
	// insertion point for field diffs
	if torus3dshape.Name != torus3dshapeOther.Name {
		diffs = append(diffs, torus3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (torusedge3dshape *TorusEdge3DShape) GongDiff(stage *Stage, torusedge3dshapeOther *TorusEdge3DShape) (diffs []string) {
	// insertion point for field diffs
	if torusedge3dshape.Name != torusedge3dshapeOther.Name {
		diffs = append(diffs, torusedge3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (torusstackshape *TorusStackShape) GongDiff(stage *Stage, torusstackshapeOther *TorusStackShape) (diffs []string) {
	// insertion point for field diffs
	if torusstackshape.Name != torusstackshapeOther.Name {
		diffs = append(diffs, torusstackshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (vase2ddiagram *Vase2DDiagram) GongDiff(stage *Stage, vase2ddiagramOther *Vase2DDiagram) (diffs []string) {
	// insertion point for field diffs
	if vase2ddiagram.Name != vase2ddiagramOther.Name {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "Name"))
	}
	if vase2ddiagram.Zoom != vase2ddiagramOther.Zoom {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "Zoom"))
	}
	if vase2ddiagram.IsVaseArcNodesExpanded != vase2ddiagramOther.IsVaseArcNodesExpanded {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsVaseArcNodesExpanded"))
	}
	if vase2ddiagram.IsVaseClampingNodesExpanded != vase2ddiagramOther.IsVaseClampingNodesExpanded {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsVaseClampingNodesExpanded"))
	}
	if vase2ddiagram.IsHiddenAxesShape != vase2ddiagramOther.IsHiddenAxesShape {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenAxesShape"))
	}
	if vase2ddiagram.IsHiddenBottomStartArcShapeGrid != vase2ddiagramOther.IsHiddenBottomStartArcShapeGrid {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenBottomStartArcShapeGrid"))
	}
	if vase2ddiagram.IsHiddenBottomEndArcShapeGrid != vase2ddiagramOther.IsHiddenBottomEndArcShapeGrid {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenBottomEndArcShapeGrid"))
	}
	if vase2ddiagram.IsHiddenBottomStackOfGrowthCurve != vase2ddiagramOther.IsHiddenBottomStackOfGrowthCurve {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenBottomStackOfGrowthCurve"))
	}
	if vase2ddiagram.IsHiddenShiftedLeftStackOfGrowthCurve != vase2ddiagramOther.IsHiddenShiftedLeftStackOfGrowthCurve {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenShiftedLeftStackOfGrowthCurve"))
	}
	if vase2ddiagram.IsHiddenShiftedLeftStackOfNormalVector != vase2ddiagramOther.IsHiddenShiftedLeftStackOfNormalVector {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenShiftedLeftStackOfNormalVector"))
	}
	if vase2ddiagram.IsHiddenPerpendicularVectorGridHalfway != vase2ddiagramOther.IsHiddenPerpendicularVectorGridHalfway {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenPerpendicularVectorGridHalfway"))
	}
	if vase2ddiagram.IsHiddenTopStartArcShapeGrid != vase2ddiagramOther.IsHiddenTopStartArcShapeGrid {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenTopStartArcShapeGrid"))
	}
	if vase2ddiagram.IsHiddenShiftedBottomTopStartArcShapeGrid != vase2ddiagramOther.IsHiddenShiftedBottomTopStartArcShapeGrid {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenShiftedBottomTopStartArcShapeGrid"))
	}
	if vase2ddiagram.IsHiddenTopMidArcVectorShapeGrid != vase2ddiagramOther.IsHiddenTopMidArcVectorShapeGrid {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenTopMidArcVectorShapeGrid"))
	}
	if vase2ddiagram.IsHiddenStartHalfwayArcShapeGrid != vase2ddiagramOther.IsHiddenStartHalfwayArcShapeGrid {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenStartHalfwayArcShapeGrid"))
	}
	if vase2ddiagram.IsHiddenTopStartHalfwayArcShapeGrid != vase2ddiagramOther.IsHiddenTopStartHalfwayArcShapeGrid {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenTopStartHalfwayArcShapeGrid"))
	}
	if vase2ddiagram.IsHiddenEndHalfwayArcShapeGrid != vase2ddiagramOther.IsHiddenEndHalfwayArcShapeGrid {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenEndHalfwayArcShapeGrid"))
	}
	if vase2ddiagram.IsHiddenTopEndHalfwayArcShapeGrid != vase2ddiagramOther.IsHiddenTopEndHalfwayArcShapeGrid {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenTopEndHalfwayArcShapeGrid"))
	}
	if vase2ddiagram.IsHiddenTopEndArcShapeGrid != vase2ddiagramOther.IsHiddenTopEndArcShapeGrid {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenTopEndArcShapeGrid"))
	}
	if vase2ddiagram.IsHiddenStackOfGrowthCurve != vase2ddiagramOther.IsHiddenStackOfGrowthCurve {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurve"))
	}
	if vase2ddiagram.IsHiddenTopStackOfGrowthCurve != vase2ddiagramOther.IsHiddenTopStackOfGrowthCurve {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenTopStackOfGrowthCurve"))
	}
	if vase2ddiagram.IsHiddenTopGrowthCurve2D != vase2ddiagramOther.IsHiddenTopGrowthCurve2D {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenTopGrowthCurve2D"))
	}
	if vase2ddiagram.IsHiddenStackOfGrowthCurve2D != vase2ddiagramOther.IsHiddenStackOfGrowthCurve2D {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurve2D"))
	}
	if vase2ddiagram.IsHiddenTopStackOfGrowthCurve2D != vase2ddiagramOther.IsHiddenTopStackOfGrowthCurve2D {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenTopStackOfGrowthCurve2D"))
	}
	if vase2ddiagram.IsHiddenGrowthCurve2DRibbon != vase2ddiagramOther.IsHiddenGrowthCurve2DRibbon {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenGrowthCurve2DRibbon"))
	}
	if vase2ddiagram.IsHiddenShiftedRightGrowthCurve2DRibbon != vase2ddiagramOther.IsHiddenShiftedRightGrowthCurve2DRibbon {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenShiftedRightGrowthCurve2DRibbon"))
	}
	if vase2ddiagram.IsHiddenShiftedLeftGrowthCurve2DRibbon != vase2ddiagramOther.IsHiddenShiftedLeftGrowthCurve2DRibbon {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenShiftedLeftGrowthCurve2DRibbon"))
	}
	if vase2ddiagram.IsHiddenStackOfGrowthCurve2DRibbon != vase2ddiagramOther.IsHiddenStackOfGrowthCurve2DRibbon {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurve2DRibbon"))
	}
	if vase2ddiagram.IsHiddenStackOfRotatedGrowthCurve2DRibbon != vase2ddiagramOther.IsHiddenStackOfRotatedGrowthCurve2DRibbon {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenStackOfRotatedGrowthCurve2DRibbon"))
	}
	if vase2ddiagram.IsHiddenPartiallyGrowthCurve2DRibbon != vase2ddiagramOther.IsHiddenPartiallyGrowthCurve2DRibbon {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenPartiallyGrowthCurve2DRibbon"))
	}
	if vase2ddiagram.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon != vase2ddiagramOther.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon"))
	}
	if vase2ddiagram.IsHiddenPartiallyGrowthCurve2DTrajectory != vase2ddiagramOther.IsHiddenPartiallyGrowthCurve2DTrajectory {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenPartiallyGrowthCurve2DTrajectory"))
	}
	if vase2ddiagram.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2 != vase2ddiagramOther.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2 {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2"))
	}
	if vase2ddiagram.IsHiddenPxShape != vase2ddiagramOther.IsHiddenPxShape {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenPxShape"))
	}
	if vase2ddiagram.IsHiddenChosenP1P2PairShape != vase2ddiagramOther.IsHiddenChosenP1P2PairShape {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenChosenP1P2PairShape"))
	}
	if vase2ddiagram.IsHiddenKeyHoleShape != vase2ddiagramOther.IsHiddenKeyHoleShape {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsHiddenKeyHoleShape"))
	}
	if vase2ddiagram.IsChecked != vase2ddiagramOther.IsChecked {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsChecked"))
	}
	if vase2ddiagram.ComputedPrefix != vase2ddiagramOther.ComputedPrefix {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if vase2ddiagram.IsExpanded != vase2ddiagramOther.IsExpanded {
		diffs = append(diffs, vase2ddiagram.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (vase3ddiagram *Vase3DDiagram) GongDiff(stage *Stage, vase3ddiagramOther *Vase3DDiagram) (diffs []string) {
	// insertion point for field diffs
	if vase3ddiagram.Name != vase3ddiagramOther.Name {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "Name"))
	}
	if vase3ddiagram.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon != vase3ddiagramOther.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon"))
	}
	if vase3ddiagram.IsHiddenTorusStackShape != vase3ddiagramOther.IsHiddenTorusStackShape {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenTorusStackShape"))
	}
	if vase3ddiagram.IsHiddenVerticalTorusStackShape != vase3ddiagramOther.IsHiddenVerticalTorusStackShape {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenVerticalTorusStackShape"))
	}
	if vase3ddiagram.IsHiddenPartiallyRotatedTorusShape != vase3ddiagramOther.IsHiddenPartiallyRotatedTorusShape {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenPartiallyRotatedTorusShape"))
	}
	if vase3ddiagram.IsHiddenStackOfPartiallyRotatedTorusShape != vase3ddiagramOther.IsHiddenStackOfPartiallyRotatedTorusShape {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenStackOfPartiallyRotatedTorusShape"))
	}
	if vase3ddiagram.IsHiddenPointsAndLines3DShape != vase3ddiagramOther.IsHiddenPointsAndLines3DShape {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenPointsAndLines3DShape"))
	}
	if vase3ddiagram.IsHiddenKeyHole3DShape != vase3ddiagramOther.IsHiddenKeyHole3DShape {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenKeyHole3DShape"))
	}
	if vase3ddiagram.IsHiddenKey3DShape != vase3ddiagramOther.IsHiddenKey3DShape {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenKey3DShape"))
	}
	if vase3ddiagram.IsHiddenVolumeKey3DShape != vase3ddiagramOther.IsHiddenVolumeKey3DShape {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenVolumeKey3DShape"))
	}
	if vase3ddiagram.IsHiddenTorusEdge3DShape != vase3ddiagramOther.IsHiddenTorusEdge3DShape {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenTorusEdge3DShape"))
	}
	if vase3ddiagram.IsHiddenSampledPoints3DShape != vase3ddiagramOther.IsHiddenSampledPoints3DShape {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenSampledPoints3DShape"))
	}
	if vase3ddiagram.IsHiddenOriginalPoints3DShape != vase3ddiagramOther.IsHiddenOriginalPoints3DShape {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenOriginalPoints3DShape"))
	}
	if vase3ddiagram.IsHiddenAngle0Shape != vase3ddiagramOther.IsHiddenAngle0Shape {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenAngle0Shape"))
	}
	if vase3ddiagram.IsHiddenTiledFloor3DShape != vase3ddiagramOther.IsHiddenTiledFloor3DShape {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsHiddenTiledFloor3DShape"))
	}
	if (vase3ddiagram.Rendered3DShape == nil) != (vase3ddiagramOther.Rendered3DShape == nil) {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
	} else if vase3ddiagram.Rendered3DShape != nil && vase3ddiagramOther.Rendered3DShape != nil {
		if vase3ddiagram.Rendered3DShape != vase3ddiagramOther.Rendered3DShape {
			diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
		}
	}
	if (vase3ddiagram.SampledPoints3DShape == nil) != (vase3ddiagramOther.SampledPoints3DShape == nil) {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
	} else if vase3ddiagram.SampledPoints3DShape != nil && vase3ddiagramOther.SampledPoints3DShape != nil {
		if vase3ddiagram.SampledPoints3DShape != vase3ddiagramOther.SampledPoints3DShape {
			diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
		}
	}
	if (vase3ddiagram.OriginalPoints3DShape == nil) != (vase3ddiagramOther.OriginalPoints3DShape == nil) {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "OriginalPoints3DShape"))
	} else if vase3ddiagram.OriginalPoints3DShape != nil && vase3ddiagramOther.OriginalPoints3DShape != nil {
		if vase3ddiagram.OriginalPoints3DShape != vase3ddiagramOther.OriginalPoints3DShape {
			diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "OriginalPoints3DShape"))
		}
	}
	if (vase3ddiagram.Angle0Shape == nil) != (vase3ddiagramOther.Angle0Shape == nil) {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "Angle0Shape"))
	} else if vase3ddiagram.Angle0Shape != nil && vase3ddiagramOther.Angle0Shape != nil {
		if vase3ddiagram.Angle0Shape != vase3ddiagramOther.Angle0Shape {
			diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "Angle0Shape"))
		}
	}
	if vase3ddiagram.IsChecked != vase3ddiagramOther.IsChecked {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsChecked"))
	}
	if vase3ddiagram.ComputedPrefix != vase3ddiagramOther.ComputedPrefix {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if vase3ddiagram.IsExpanded != vase3ddiagramOther.IsExpanded {
		diffs = append(diffs, vase3ddiagram.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (vaseabstract *VaseAbstract) GongDiff(stage *Stage, vaseabstractOther *VaseAbstract) (diffs []string) {
	// insertion point for field diffs
	if vaseabstract.Name != vaseabstractOther.Name {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "Name"))
	}
	if vaseabstract.RelativeVerticalThickness != vaseabstractOther.RelativeVerticalThickness {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "RelativeVerticalThickness"))
	}
	if vaseabstract.RelativeRadialThickness != vaseabstractOther.RelativeRadialThickness {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "RelativeRadialThickness"))
	}
	if vaseabstract.RelativeCuttedStackFloorHeight != vaseabstractOther.RelativeCuttedStackFloorHeight {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "RelativeCuttedStackFloorHeight"))
	}
	if vaseabstract.RelativeRotatedTorusSeparation != vaseabstractOther.RelativeRotatedTorusSeparation {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "RelativeRotatedTorusSeparation"))
	}
	if vaseabstract.RotationRatio != vaseabstractOther.RotationRatio {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "RotationRatio"))
	}
	if vaseabstract.RadialRepetitions != vaseabstractOther.RadialRepetitions {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "RadialRepetitions"))
	}
	if vaseabstract.Transparency != vaseabstractOther.Transparency {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "Transparency"))
	}
	if vaseabstract.HasAlternatingRingColors != vaseabstractOther.HasAlternatingRingColors {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "HasAlternatingRingColors"))
	}
	if vaseabstract.RelativeTrajectoryOffsetX != vaseabstractOther.RelativeTrajectoryOffsetX {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "RelativeTrajectoryOffsetX"))
	}
	if vaseabstract.RelativeTrajectoryOffsetY != vaseabstractOther.RelativeTrajectoryOffsetY {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "RelativeTrajectoryOffsetY"))
	}
	if vaseabstract.NbStepP1P2 != vaseabstractOther.NbStepP1P2 {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "NbStepP1P2"))
	}
	if vaseabstract.ChosenStep != vaseabstractOther.ChosenStep {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "ChosenStep"))
	}
	if vaseabstract.RelativeHorizontalRingsHeight != vaseabstractOther.RelativeHorizontalRingsHeight {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "RelativeHorizontalRingsHeight"))
	}
	if vaseabstract.OffsetKeyX != vaseabstractOther.OffsetKeyX {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "OffsetKeyX"))
	}
	if vaseabstract.OffsetKeyY != vaseabstractOther.OffsetKeyY {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "OffsetKeyY"))
	}
	if vaseabstract.HeightKey != vaseabstractOther.HeightKey {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "HeightKey"))
	}
	if vaseabstract.WidthKey != vaseabstractOther.WidthKey {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "WidthKey"))
	}
	if vaseabstract.RelativeKeySize != vaseabstractOther.RelativeKeySize {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "RelativeKeySize"))
	}
	if vaseabstract.MovieNbFrames != vaseabstractOther.MovieNbFrames {
		diffs = append(diffs, vaseabstract.GongMarshallField(stage, "MovieNbFrames"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (verticaltorusstackshape *VerticalTorusStackShape) GongDiff(stage *Stage, verticaltorusstackshapeOther *VerticalTorusStackShape) (diffs []string) {
	// insertion point for field diffs
	if verticaltorusstackshape.Name != verticaltorusstackshapeOther.Name {
		diffs = append(diffs, verticaltorusstackshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (volumekey3dshape *VolumeKey3DShape) GongDiff(stage *Stage, volumekey3dshapeOther *VolumeKey3DShape) (diffs []string) {
	// insertion point for field diffs
	if volumekey3dshape.Name != volumekey3dshapeOther.Name {
		diffs = append(diffs, volumekey3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
