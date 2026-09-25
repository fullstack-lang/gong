// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (angle0shape *Angle0Shape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Angle0Shapes[angle0shape]
	return ok
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ArcNormalVectorShapes[arcnormalvectorshape]
	return ok
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ArcNormalVectorShapeGrids[arcnormalvectorshapegrid]
	return ok
}

func (axesshape *AxesShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AxesShapes[axesshape]
	return ok
}

func (basevectorshape *BaseVectorShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BaseVectorShapes[basevectorshape]
	return ok
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BaseVectorShapeGrids[basevectorshapegrid]
	return ok
}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BottomCurvePlane1Shapes[bottomcurveplane1shape]
	return ok
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BottomCurvePlane2Shapes[bottomcurveplane2shape]
	return ok
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ChosenP1P2PairShapes[chosenp1p2pairshape]
	return ok
}

func (circlegridshape *CircleGridShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CircleGridShapes[circlegridshape]
	return ok
}

func (circumference3dshape *Circumference3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Circumference3DShapes[circumference3dshape]
	return ok
}

func (clock2ddiagram *Clock2DDiagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Clock2DDiagrams[clock2ddiagram]
	return ok
}

func (clock3ddiagram *Clock3DDiagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Clock3DDiagrams[clock3ddiagram]
	return ok
}

func (clocktopcurveshape *ClockTopCurveShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ClockTopCurveShapes[clocktopcurveshape]
	return ok
}

func (cutline3dshape *CutLine3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CutLine3DShapes[cutline3dshape]
	return ok
}

func (endarcshape *EndArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EndArcShapes[endarcshape]
	return ok
}

func (endarcshapegrid *EndArcShapeGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EndArcShapeGrids[endarcshapegrid]
	return ok
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EndHalfwayArcShapes[endhalfwayarcshape]
	return ok
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EndHalfwayArcShapeGrids[endhalfwayarcshapegrid]
	return ok
}

func (explanationtextshape *ExplanationTextShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ExplanationTextShapes[explanationtextshape]
	return ok
}

func (eye3dshape *Eye3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Eye3DShapes[eye3dshape]
	return ok
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EyeCornersSampledPoints3DShapes[eyecornerssampledpoints3dshape]
	return ok
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EyeSampledPoints3DShapes[eyesampledpoints3dshape]
	return ok
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EyeSeatBottomCurveShapes[eyeseatbottomcurveshape]
	return ok
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EyeStoolBottomCurveShapes[eyestoolbottomcurveshape]
	return ok
}

func (eyevolume3dshape *EyeVolume3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EyeVolume3DShapes[eyevolume3dshape]
	return ok
}

func (gridpathshape *GridPathShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GridPathShapes[gridpathshape]
	return ok
}

func (growthcurve2d *GrowthCurve2D) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GrowthCurve2Ds[growthcurve2d]
	return ok
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GrowthCurve2DRibbons[growthcurve2dribbon]
	return ok
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GrowthCurve2DRibbonEndShapes[growthcurve2dribbonendshape]
	return ok
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GrowthCurve2DRibbonStartShapes[growthcurve2dribbonstartshape]
	return ok
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GrowthCurveRhombusGridShapes[growthcurverhombusgridshape]
	return ok
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GrowthCurveRhombusShapes[growthcurverhombusshape]
	return ok
}

func (growthvectorshape *GrowthVectorShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GrowthVectorShapes[growthvectorshape]
	return ok
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.InitialRhombusGridShapes[initialrhombusgridshape]
	return ok
}

func (initialrhombusshape *InitialRhombusShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.InitialRhombusShapes[initialrhombusshape]
	return ok
}

func (key3dshape *Key3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Key3DShapes[key3dshape]
	return ok
}

func (keyhole3dshape *KeyHole3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.KeyHole3DShapes[keyhole3dshape]
	return ok
}

func (keyholeshape *KeyHoleShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.KeyHoleShapes[keyholeshape]
	return ok
}

func (leaves3dshape *Leaves3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Leaves3DShapes[leaves3dshape]
	return ok
}

func (library *Library) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Librarys[library]
	return ok
}

func (midarcvectorshape *MidArcVectorShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MidArcVectorShapes[midarcvectorshape]
	return ok
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MidArcVectorShapeGrids[midarcvectorshapegrid]
	return ok
}

func (originalpoints3dshape *OriginalPoints3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.OriginalPoints3DShapes[originalpoints3dshape]
	return ok
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ParastichyMCurves3DShapes[parastichymcurves3dshape]
	return ok
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ParastichyNCurves3DShapes[parastichyncurves3dshape]
	return ok
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyGrowthCurve2DRibbons[partiallygrowthcurve2dribbon]
	return ok
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyGrowthCurve2DRibbonEndShapes[partiallygrowthcurve2dribbonendshape]
	return ok
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyGrowthCurve2DRibbonStartShapes[partiallygrowthcurve2dribbonstartshape]
	return ok
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyGrowthCurve2DTrajectorys[partiallygrowthcurve2dtrajectory]
	return ok
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes[partiallygrowthcurve2dtrajectoryp1curveshape]
	return ok
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyGrowthCurve2DTrajectoryP1P2s[partiallygrowthcurve2dtrajectoryp1p2]
	return ok
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes[partiallygrowthcurve2dtrajectoryp1p2pairlineshape]
	return ok
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes[partiallygrowthcurve2dtrajectoryp1pointshape]
	return ok
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes[partiallygrowthcurve2dtrajectoryp2curveshape]
	return ok
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes[partiallygrowthcurve2dtrajectoryp2pointshape]
	return ok
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyGrowthCurve2DTrajectoryShapes[partiallygrowthcurve2dtrajectoryshape]
	return ok
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyRotatedSeatBottomCurveShapes[partiallyrotatedseatbottomcurveshape]
	return ok
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyRotatedSeatTopCurveShapes[partiallyrotatedseattopcurveshape]
	return ok
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartiallyRotatedTorusShapes[partiallyrotatedtorusshape]
	return ok
}

func (perpendicularvector *PerpendicularVector) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PerpendicularVectors[perpendicularvector]
	return ok
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PerpendicularVectorGrids[perpendicularvectorgrid]
	return ok
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PerpendicularVectorGridHalfways[perpendicularvectorgridhalfway]
	return ok
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PerpendicularVectorHalfways[perpendicularvectorhalfway]
	return ok
}

func (plant2ddiagram *Plant2DDiagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Plant2DDiagrams[plant2ddiagram]
	return ok
}

func (plant3ddiagram *Plant3DDiagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Plant3DDiagrams[plant3ddiagram]
	return ok
}

func (plantabstract *PlantAbstract) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PlantAbstracts[plantabstract]
	return ok
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PlantCircumferenceShapes[plantcircumferenceshape]
	return ok
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PointsAndLines3DShapes[pointsandlines3dshape]
	return ok
}

func (pxshape *PxShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PxShapes[pxshape]
	return ok
}

func (rendered3dshape *Rendered3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Rendered3DShapes[rendered3dshape]
	return ok
}

func (rhombusshape *RhombusShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RhombusShapes[rhombusshape]
	return ok
}

func (rhombusstuff *RhombusStuff) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RhombusStuffs[rhombusstuff]
	return ok
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RotatedRhombusGridShapes[rotatedrhombusgridshape]
	return ok
}

func (rotatedrhombusshape *RotatedRhombusShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RotatedRhombusShapes[rotatedrhombusshape]
	return ok
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RotatedSampledPoints3DShapes[rotatedsampledpoints3dshape]
	return ok
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RotatedSeatAndLegs3DShapes[rotatedseatandlegs3dshape]
	return ok
}

func (sampledpoints3dshape *SampledPoints3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SampledPoints3DShapes[sampledpoints3dshape]
	return ok
}

func (seat3dshape *Seat3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Seat3DShapes[seat3dshape]
	return ok
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SeatAndLegs3DShapes[seatandlegs3dshape]
	return ok
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SeatBottomCurveShapes[seatbottomcurveshape]
	return ok
}

func (seattopcurveshape *SeatTopCurveShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SeatTopCurveShapes[seattopcurveshape]
	return ok
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedBottomTopStartArcShapes[shiftedbottomtopstartarcshape]
	return ok
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedBottomTopStartArcShapeGrids[shiftedbottomtopstartarcshapegrid]
	return ok
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedLeftGrowthCurve2DRibbons[shiftedleftgrowthcurve2dribbon]
	return ok
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedLeftGrowthCurve2DRibbonEndShapes[shiftedleftgrowthcurve2dribbonendshape]
	return ok
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedLeftGrowthCurve2DRibbonStartShapes[shiftedleftgrowthcurve2dribbonstartshape]
	return ok
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedLeftPartiallyGrowthCurve2DRibbons[shiftedleftpartiallygrowthcurve2dribbon]
	return ok
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes[shiftedleftpartiallygrowthcurve2dribbonendshape]
	return ok
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes[shiftedleftpartiallygrowthcurve2dribbonstartshape]
	return ok
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedLeftStackGrowthCurveEndArcShapes[shiftedleftstackgrowthcurveendarcshape]
	return ok
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedLeftStackGrowthCurveStartArcShapes[shiftedleftstackgrowthcurvestartarcshape]
	return ok
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedLeftStackNormalVectors[shiftedleftstacknormalvector]
	return ok
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedLeftStackOfGrowthCurves[shiftedleftstackofgrowthcurve]
	return ok
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedLeftStackOfNormalVectors[shiftedleftstackofnormalvector]
	return ok
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedRightGrowthCurve2DRibbons[shiftedrightgrowthcurve2dribbon]
	return ok
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedRightGrowthCurve2DRibbonEndShapes[shiftedrightgrowthcurve2dribbonendshape]
	return ok
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ShiftedRightGrowthCurve2DRibbonStartShapes[shiftedrightgrowthcurve2dribbonstartshape]
	return ok
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackGrowthCurve2DEndHalfwayArcShapes[stackgrowthcurve2dendhalfwayarcshape]
	return ok
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackGrowthCurve2DRibbonEndShapes[stackgrowthcurve2dribbonendshape]
	return ok
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackGrowthCurve2DRibbonStartShapes[stackgrowthcurve2dribbonstartshape]
	return ok
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackGrowthCurve2DStartHalfwayArcShapes[stackgrowthcurve2dstarthalfwayarcshape]
	return ok
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackOfGrowthCurve2Ds[stackofgrowthcurve2d]
	return ok
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackOfGrowthCurve2DByGrowthVectors[stackofgrowthcurve2dbygrowthvector]
	return ok
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackOfGrowthCurve2DRibbons[stackofgrowthcurve2dribbon]
	return ok
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackOfPartiallyRotatedTorusShapes[stackofpartiallyrotatedtorusshape]
	return ok
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackOfRotatedGrowthCurve2Ds[stackofrotatedgrowthcurve2d]
	return ok
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackOfRotatedGrowthCurve2DRibbons[stackofrotatedgrowthcurve2dribbon]
	return ok
}

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackOfRotatedVaseTrapezeRingsShapes[stackofrotatedvasetrapezeringsshape]
	return ok
}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackOfVaseTrapezeRingsShapes[stackofvasetrapezeringsshape]
	return ok
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackRotatedGrowthCurve2DEndArcShapes[stackrotatedgrowthcurve2dendarcshape]
	return ok
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackRotatedGrowthCurve2DRibbonEndShapes[stackrotatedgrowthcurve2dribbonendshape]
	return ok
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackRotatedGrowthCurve2DRibbonStartShapes[stackrotatedgrowthcurve2dribbonstartshape]
	return ok
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StackRotatedGrowthCurve2DStartArcShapes[stackrotatedgrowthcurve2dstartarcshape]
	return ok
}

func (startarcshape *StartArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StartArcShapes[startarcshape]
	return ok
}

func (startarcshapegrid *StartArcShapeGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StartArcShapeGrids[startarcshapegrid]
	return ok
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StartHalfwayArcShapes[starthalfwayarcshape]
	return ok
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StartHalfwayArcShapeGrids[starthalfwayarcshapegrid]
	return ok
}

func (stemcylinder3dshape *StemCylinder3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StemCylinder3DShapes[stemcylinder3dshape]
	return ok
}

func (stool2ddiagram *Stool2DDiagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Stool2DDiagrams[stool2ddiagram]
	return ok
}

func (stool3ddiagram *Stool3DDiagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Stool3DDiagrams[stool3ddiagram]
	return ok
}

func (tiledfloor3dshape *TiledFloor3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TiledFloor3DShapes[tiledfloor3dshape]
	return ok
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopCurvePlane1Shapes[topcurveplane1shape]
	return ok
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopCurvePlane2Shapes[topcurveplane2shape]
	return ok
}

func (topendarcshape *TopEndArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopEndArcShapes[topendarcshape]
	return ok
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopEndArcShapeGrids[topendarcshapegrid]
	return ok
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopEndHalfwayArcShapes[topendhalfwayarcshape]
	return ok
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopEndHalfwayArcShapeGrids[topendhalfwayarcshapegrid]
	return ok
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopGrowthCurve2Ds[topgrowthcurve2d]
	return ok
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopMidArcVectorShapes[topmidarcvectorshape]
	return ok
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopMidArcVectorShapeGrids[topmidarcvectorshapegrid]
	return ok
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopStackGrowthCurve2DEndHalfwayArcShapes[topstackgrowthcurve2dendhalfwayarcshape]
	return ok
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopStackGrowthCurve2DStartHalfwayArcShapes[topstackgrowthcurve2dstarthalfwayarcshape]
	return ok
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopStackOfGrowthCurve2Ds[topstackofgrowthcurve2d]
	return ok
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopStackOfRotatedGrowthCurve2Ds[topstackofrotatedgrowthcurve2d]
	return ok
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopStackOfRotatedGrowthCurve2DEndArcShapes[topstackofrotatedgrowthcurve2dendarcshape]
	return ok
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopStackOfRotatedGrowthCurve2DStartArcShapes[topstackofrotatedgrowthcurve2dstartarcshape]
	return ok
}

func (topstartarcshape *TopStartArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopStartArcShapes[topstartarcshape]
	return ok
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopStartArcShapeGrids[topstartarcshapegrid]
	return ok
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopStartHalfwayArcShapes[topstarthalfwayarcshape]
	return ok
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TopStartHalfwayArcShapeGrids[topstarthalfwayarcshapegrid]
	return ok
}

func (torus3dshape *Torus3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Torus3DShapes[torus3dshape]
	return ok
}

func (torusedge3dshape *TorusEdge3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TorusEdge3DShapes[torusedge3dshape]
	return ok
}

func (torusstackshape *TorusStackShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TorusStackShapes[torusstackshape]
	return ok
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TubeVase3DDiagrams[tubevase3ddiagram]
	return ok
}

func (tubevaseabstract *TubeVaseAbstract) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TubeVaseAbstracts[tubevaseabstract]
	return ok
}

func (vase2ddiagram *Vase2DDiagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Vase2DDiagrams[vase2ddiagram]
	return ok
}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.VaseTrapezeRingShapes[vasetrapezeringshape]
	return ok
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.VerticalTorusStackShapes[verticaltorusstackshape]
	return ok
}

func (volumekey3dshape *VolumeKey3DShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.VolumeKey3DShapes[volumekey3dshape]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (angle0shape *Angle0Shape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(angle0shape) {
		return
	}

	angle0shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arcnormalvectorshape *ArcNormalVectorShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(arcnormalvectorshape) {
		return
	}

	arcnormalvectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(arcnormalvectorshapegrid) {
		return
	}

	arcnormalvectorshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (axesshape *AxesShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(axesshape) {
		return
	}

	axesshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (basevectorshape *BaseVectorShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(basevectorshape) {
		return
	}

	basevectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (basevectorshapegrid *BaseVectorShapeGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(basevectorshapegrid) {
		return
	}

	basevectorshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bottomcurveplane1shape) {
		return
	}

	bottomcurveplane1shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bottomcurveplane2shape) {
		return
	}

	bottomcurveplane2shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(chosenp1p2pairshape) {
		return
	}

	chosenp1p2pairshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circlegridshape *CircleGridShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(circlegridshape) {
		return
	}

	circlegridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circumference3dshape *Circumference3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(circumference3dshape) {
		return
	}

	circumference3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clock2ddiagram *Clock2DDiagram) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(clock2ddiagram) {
		return
	}

	clock2ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clock3ddiagram *Clock3DDiagram) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(clock3ddiagram) {
		return
	}

	clock3ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if clock3ddiagram.SampledPoints3DShape != nil {
		stage.StageBranch(clock3ddiagram.SampledPoints3DShape)
	}
	if clock3ddiagram.Rendered3DShape != nil {
		stage.StageBranch(clock3ddiagram.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clocktopcurveshape *ClockTopCurveShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(clocktopcurveshape) {
		return
	}

	clocktopcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cutline3dshape *CutLine3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(cutline3dshape) {
		return
	}

	cutline3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endarcshape *EndArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(endarcshape) {
		return
	}

	endarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endarcshapegrid *EndArcShapeGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(endarcshapegrid) {
		return
	}

	endarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endhalfwayarcshape *EndHalfwayArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(endhalfwayarcshape) {
		return
	}

	endhalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(endhalfwayarcshapegrid) {
		return
	}

	endhalfwayarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (explanationtextshape *ExplanationTextShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(explanationtextshape) {
		return
	}

	explanationtextshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eye3dshape *Eye3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(eye3dshape) {
		return
	}

	eye3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(eyecornerssampledpoints3dshape) {
		return
	}

	eyecornerssampledpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(eyesampledpoints3dshape) {
		return
	}

	eyesampledpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(eyeseatbottomcurveshape) {
		return
	}

	eyeseatbottomcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(eyestoolbottomcurveshape) {
		return
	}

	eyestoolbottomcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyevolume3dshape *EyeVolume3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(eyevolume3dshape) {
		return
	}

	eyevolume3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gridpathshape *GridPathShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(gridpathshape) {
		return
	}

	gridpathshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2d *GrowthCurve2D) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(growthcurve2d) {
		return
	}

	growthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(growthcurve2dribbon) {
		return
	}

	growthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(growthcurve2dribbonendshape) {
		return
	}

	growthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(growthcurve2dribbonstartshape) {
		return
	}

	growthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(growthcurverhombusgridshape) {
		return
	}

	growthcurverhombusgridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(growthcurverhombusshape) {
		return
	}

	growthcurverhombusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthvectorshape *GrowthVectorShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(growthvectorshape) {
		return
	}

	growthvectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (initialrhombusgridshape *InitialRhombusGridShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(initialrhombusgridshape) {
		return
	}

	initialrhombusgridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (initialrhombusshape *InitialRhombusShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(initialrhombusshape) {
		return
	}

	initialrhombusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (key3dshape *Key3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(key3dshape) {
		return
	}

	key3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (keyhole3dshape *KeyHole3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(keyhole3dshape) {
		return
	}

	keyhole3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (keyholeshape *KeyHoleShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(keyholeshape) {
		return
	}

	keyholeshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (leaves3dshape *Leaves3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(leaves3dshape) {
		return
	}

	leaves3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plantabstract := range library.Plants {
		stage.StageBranch(_plantabstract)
	}
	for _, _library := range library.SubLibraries {
		stage.StageBranch(_library)
	}

}

func (midarcvectorshape *MidArcVectorShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(midarcvectorshape) {
		return
	}

	midarcvectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(midarcvectorshapegrid) {
		return
	}

	midarcvectorshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (originalpoints3dshape *OriginalPoints3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(originalpoints3dshape) {
		return
	}

	originalpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(parastichymcurves3dshape) {
		return
	}

	parastichymcurves3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(parastichyncurves3dshape) {
		return
	}

	parastichyncurves3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dribbon) {
		return
	}

	partiallygrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dribbonendshape) {
		return
	}

	partiallygrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dribbonstartshape) {
		return
	}

	partiallygrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectory) {
		return
	}

	partiallygrowthcurve2dtrajectory.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryp1curveshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1curveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryp1p2) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryp1p2pairlineshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2pairlineshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryp1pointshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1pointshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryp2curveshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp2curveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryp2pointshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp2pointshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallyrotatedseatbottomcurveshape) {
		return
	}

	partiallyrotatedseatbottomcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallyrotatedseattopcurveshape) {
		return
	}

	partiallyrotatedseattopcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partiallyrotatedtorusshape) {
		return
	}

	partiallyrotatedtorusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvector *PerpendicularVector) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(perpendicularvector) {
		return
	}

	perpendicularvector.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(perpendicularvectorgrid) {
		return
	}

	perpendicularvectorgrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(perpendicularvectorgridhalfway) {
		return
	}

	perpendicularvectorgridhalfway.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(perpendicularvectorhalfway) {
		return
	}

	perpendicularvectorhalfway.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (plant2ddiagram *Plant2DDiagram) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(plant2ddiagram) {
		return
	}

	plant2ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (plant3ddiagram *Plant3DDiagram) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(plant3ddiagram) {
		return
	}

	plant3ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plant3ddiagram.StemCylinder3DShape != nil {
		stage.StageBranch(plant3ddiagram.StemCylinder3DShape)
	}
	if plant3ddiagram.ParastichyNCurves3DShape != nil {
		stage.StageBranch(plant3ddiagram.ParastichyNCurves3DShape)
	}
	if plant3ddiagram.ParastichyMCurves3DShape != nil {
		stage.StageBranch(plant3ddiagram.ParastichyMCurves3DShape)
	}
	if plant3ddiagram.CutLine3DShape != nil {
		stage.StageBranch(plant3ddiagram.CutLine3DShape)
	}
	if plant3ddiagram.Circumference3DShape != nil {
		stage.StageBranch(plant3ddiagram.Circumference3DShape)
	}
	if plant3ddiagram.Leaves3DShape != nil {
		stage.StageBranch(plant3ddiagram.Leaves3DShape)
	}
	if plant3ddiagram.Rendered3DShape != nil {
		stage.StageBranch(plant3ddiagram.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (plantabstract *PlantAbstract) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(plantabstract) {
		return
	}

	plantabstract.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plantabstract.TubeVaseAbstract != nil {
		stage.StageBranch(plantabstract.TubeVaseAbstract)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plant2ddiagram := range plantabstract.Plant2DDiagrams {
		stage.StageBranch(_plant2ddiagram)
	}
	for _, _plant3ddiagram := range plantabstract.Plant3DDiagrams {
		stage.StageBranch(_plant3ddiagram)
	}
	for _, _vase2ddiagram := range plantabstract.Vase2DDiagrams {
		stage.StageBranch(_vase2ddiagram)
	}
	for _, _tubevase3ddiagram := range plantabstract.TubeVase3DDiagrams {
		stage.StageBranch(_tubevase3ddiagram)
	}
	for _, _stool2ddiagram := range plantabstract.Stool2DDiagrams {
		stage.StageBranch(_stool2ddiagram)
	}
	for _, _stool3ddiagram := range plantabstract.Stool3DDiagrams {
		stage.StageBranch(_stool3ddiagram)
	}
	for _, _clock2ddiagram := range plantabstract.Clock2DDiagrams {
		stage.StageBranch(_clock2ddiagram)
	}
	for _, _clock3ddiagram := range plantabstract.Clock3DDiagrams {
		stage.StageBranch(_clock3ddiagram)
	}

}

func (plantcircumferenceshape *PlantCircumferenceShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(plantcircumferenceshape) {
		return
	}

	plantcircumferenceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointsandlines3dshape *PointsAndLines3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pointsandlines3dshape) {
		return
	}

	pointsandlines3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pxshape *PxShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pxshape) {
		return
	}

	pxshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rendered3dshape *Rendered3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rendered3dshape) {
		return
	}

	rendered3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rhombusshape *RhombusShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rhombusshape) {
		return
	}

	rhombusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rhombusstuff *RhombusStuff) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rhombusstuff) {
		return
	}

	rhombusstuff.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rotatedrhombusgridshape) {
		return
	}

	rotatedrhombusgridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedrhombusshape *RotatedRhombusShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rotatedrhombusshape) {
		return
	}

	rotatedrhombusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rotatedsampledpoints3dshape) {
		return
	}

	rotatedsampledpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rotatedseatandlegs3dshape) {
		return
	}

	rotatedseatandlegs3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (sampledpoints3dshape *SampledPoints3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(sampledpoints3dshape) {
		return
	}

	sampledpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seat3dshape *Seat3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(seat3dshape) {
		return
	}

	seat3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(seatandlegs3dshape) {
		return
	}

	seatandlegs3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seatbottomcurveshape *SeatBottomCurveShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(seatbottomcurveshape) {
		return
	}

	seatbottomcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seattopcurveshape *SeatTopCurveShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(seattopcurveshape) {
		return
	}

	seattopcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedbottomtopstartarcshape) {
		return
	}

	shiftedbottomtopstartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedbottomtopstartarcshapegrid) {
		return
	}

	shiftedbottomtopstartarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftgrowthcurve2dribbon) {
		return
	}

	shiftedleftgrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftgrowthcurve2dribbonendshape) {
		return
	}

	shiftedleftgrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftgrowthcurve2dribbonstartshape) {
		return
	}

	shiftedleftgrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftpartiallygrowthcurve2dribbon) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftpartiallygrowthcurve2dribbonendshape) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftpartiallygrowthcurve2dribbonstartshape) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftstackgrowthcurveendarcshape) {
		return
	}

	shiftedleftstackgrowthcurveendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftstackgrowthcurvestartarcshape) {
		return
	}

	shiftedleftstackgrowthcurvestartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftstacknormalvector) {
		return
	}

	shiftedleftstacknormalvector.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftstackofgrowthcurve) {
		return
	}

	shiftedleftstackofgrowthcurve.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftstackofnormalvector) {
		return
	}

	shiftedleftstackofnormalvector.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedrightgrowthcurve2dribbon) {
		return
	}

	shiftedrightgrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedrightgrowthcurve2dribbonendshape) {
		return
	}

	shiftedrightgrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shiftedrightgrowthcurve2dribbonstartshape) {
		return
	}

	shiftedrightgrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackgrowthcurve2dendhalfwayarcshape) {
		return
	}

	stackgrowthcurve2dendhalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackgrowthcurve2dribbonendshape) {
		return
	}

	stackgrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackgrowthcurve2dribbonstartshape) {
		return
	}

	stackgrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackgrowthcurve2dstarthalfwayarcshape) {
		return
	}

	stackgrowthcurve2dstarthalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackofgrowthcurve2d) {
		return
	}

	stackofgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackofgrowthcurve2dbygrowthvector) {
		return
	}

	stackofgrowthcurve2dbygrowthvector.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackofgrowthcurve2dribbon) {
		return
	}

	stackofgrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackofpartiallyrotatedtorusshape) {
		return
	}

	stackofpartiallyrotatedtorusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackofrotatedgrowthcurve2d) {
		return
	}

	stackofrotatedgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackofrotatedgrowthcurve2dribbon) {
		return
	}

	stackofrotatedgrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackofrotatedvasetrapezeringsshape) {
		return
	}

	stackofrotatedvasetrapezeringsshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackofvasetrapezeringsshape) {
		return
	}

	stackofvasetrapezeringsshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackrotatedgrowthcurve2dendarcshape) {
		return
	}

	stackrotatedgrowthcurve2dendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackrotatedgrowthcurve2dribbonendshape) {
		return
	}

	stackrotatedgrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackrotatedgrowthcurve2dribbonstartshape) {
		return
	}

	stackrotatedgrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stackrotatedgrowthcurve2dstartarcshape) {
		return
	}

	stackrotatedgrowthcurve2dstartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (startarcshape *StartArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(startarcshape) {
		return
	}

	startarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (startarcshapegrid *StartArcShapeGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(startarcshapegrid) {
		return
	}

	startarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (starthalfwayarcshape *StartHalfwayArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(starthalfwayarcshape) {
		return
	}

	starthalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(starthalfwayarcshapegrid) {
		return
	}

	starthalfwayarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stemcylinder3dshape *StemCylinder3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stemcylinder3dshape) {
		return
	}

	stemcylinder3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stool2ddiagram *Stool2DDiagram) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stool2ddiagram) {
		return
	}

	stool2ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stool3ddiagram *Stool3DDiagram) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stool3ddiagram) {
		return
	}

	stool3ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stool3ddiagram.SampledPoints3DShape != nil {
		stage.StageBranch(stool3ddiagram.SampledPoints3DShape)
	}
	if stool3ddiagram.Rendered3DShape != nil {
		stage.StageBranch(stool3ddiagram.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tiledfloor3dshape *TiledFloor3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tiledfloor3dshape) {
		return
	}

	tiledfloor3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topcurveplane1shape *TopCurvePlane1Shape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topcurveplane1shape) {
		return
	}

	topcurveplane1shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topcurveplane2shape *TopCurvePlane2Shape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topcurveplane2shape) {
		return
	}

	topcurveplane2shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendarcshape *TopEndArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topendarcshape) {
		return
	}

	topendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendarcshapegrid *TopEndArcShapeGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topendarcshapegrid) {
		return
	}

	topendarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topendhalfwayarcshape) {
		return
	}

	topendhalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topendhalfwayarcshapegrid) {
		return
	}

	topendhalfwayarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topgrowthcurve2d *TopGrowthCurve2D) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topgrowthcurve2d) {
		return
	}

	topgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topmidarcvectorshape *TopMidArcVectorShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topmidarcvectorshape) {
		return
	}

	topmidarcvectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topmidarcvectorshapegrid) {
		return
	}

	topmidarcvectorshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topstackgrowthcurve2dendhalfwayarcshape) {
		return
	}

	topstackgrowthcurve2dendhalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topstackgrowthcurve2dstarthalfwayarcshape) {
		return
	}

	topstackgrowthcurve2dstarthalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topstackofgrowthcurve2d) {
		return
	}

	topstackofgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topstackofrotatedgrowthcurve2d) {
		return
	}

	topstackofrotatedgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topstackofrotatedgrowthcurve2dendarcshape) {
		return
	}

	topstackofrotatedgrowthcurve2dendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topstackofrotatedgrowthcurve2dstartarcshape) {
		return
	}

	topstackofrotatedgrowthcurve2dstartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstartarcshape *TopStartArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topstartarcshape) {
		return
	}

	topstartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topstartarcshapegrid) {
		return
	}

	topstartarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topstarthalfwayarcshape) {
		return
	}

	topstarthalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(topstarthalfwayarcshapegrid) {
		return
	}

	topstarthalfwayarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torus3dshape *Torus3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(torus3dshape) {
		return
	}

	torus3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torusedge3dshape *TorusEdge3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(torusedge3dshape) {
		return
	}

	torusedge3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torusstackshape *TorusStackShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(torusstackshape) {
		return
	}

	torusstackshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tubevase3ddiagram *TubeVase3DDiagram) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tubevase3ddiagram) {
		return
	}

	tubevase3ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tubevase3ddiagram.Rendered3DShape != nil {
		stage.StageBranch(tubevase3ddiagram.Rendered3DShape)
	}
	if tubevase3ddiagram.SampledPoints3DShape != nil {
		stage.StageBranch(tubevase3ddiagram.SampledPoints3DShape)
	}
	if tubevase3ddiagram.OriginalPoints3DShape != nil {
		stage.StageBranch(tubevase3ddiagram.OriginalPoints3DShape)
	}
	if tubevase3ddiagram.Angle0Shape != nil {
		stage.StageBranch(tubevase3ddiagram.Angle0Shape)
	}
	if tubevase3ddiagram.TopCurvePlane1Shape != nil {
		stage.StageBranch(tubevase3ddiagram.TopCurvePlane1Shape)
	}
	if tubevase3ddiagram.BottomCurvePlane1Shape != nil {
		stage.StageBranch(tubevase3ddiagram.BottomCurvePlane1Shape)
	}
	if tubevase3ddiagram.TopCurvePlane2Shape != nil {
		stage.StageBranch(tubevase3ddiagram.TopCurvePlane2Shape)
	}
	if tubevase3ddiagram.BottomCurvePlane2Shape != nil {
		stage.StageBranch(tubevase3ddiagram.BottomCurvePlane2Shape)
	}
	if tubevase3ddiagram.VaseTrapezeRingShape != nil {
		stage.StageBranch(tubevase3ddiagram.VaseTrapezeRingShape)
	}
	if tubevase3ddiagram.StackOfVaseTrapezeRingsShape != nil {
		stage.StageBranch(tubevase3ddiagram.StackOfVaseTrapezeRingsShape)
	}
	if tubevase3ddiagram.StackOfRotatedVaseTrapezeRingsShape != nil {
		stage.StageBranch(tubevase3ddiagram.StackOfRotatedVaseTrapezeRingsShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tubevaseabstract *TubeVaseAbstract) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tubevaseabstract) {
		return
	}

	tubevaseabstract.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vase2ddiagram *Vase2DDiagram) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(vase2ddiagram) {
		return
	}

	vase2ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(vasetrapezeringshape) {
		return
	}

	vasetrapezeringshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticaltorusstackshape *VerticalTorusStackShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(verticaltorusstackshape) {
		return
	}

	verticaltorusstackshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (volumekey3dshape *VolumeKey3DShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(volumekey3dshape) {
		return
	}

	volumekey3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Angle0Shape:
		toT := GongCopyBranchAngle0Shape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ArcNormalVectorShape:
		toT := GongCopyBranchArcNormalVectorShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ArcNormalVectorShapeGrid:
		toT := GongCopyBranchArcNormalVectorShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AxesShape:
		toT := GongCopyBranchAxesShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BaseVectorShape:
		toT := GongCopyBranchBaseVectorShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BaseVectorShapeGrid:
		toT := GongCopyBranchBaseVectorShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BottomCurvePlane1Shape:
		toT := GongCopyBranchBottomCurvePlane1Shape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BottomCurvePlane2Shape:
		toT := GongCopyBranchBottomCurvePlane2Shape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ChosenP1P2PairShape:
		toT := GongCopyBranchChosenP1P2PairShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CircleGridShape:
		toT := GongCopyBranchCircleGridShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Circumference3DShape:
		toT := GongCopyBranchCircumference3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Clock2DDiagram:
		toT := GongCopyBranchClock2DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Clock3DDiagram:
		toT := GongCopyBranchClock3DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ClockTopCurveShape:
		toT := GongCopyBranchClockTopCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CutLine3DShape:
		toT := GongCopyBranchCutLine3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EndArcShape:
		toT := GongCopyBranchEndArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EndArcShapeGrid:
		toT := GongCopyBranchEndArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EndHalfwayArcShape:
		toT := GongCopyBranchEndHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EndHalfwayArcShapeGrid:
		toT := GongCopyBranchEndHalfwayArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ExplanationTextShape:
		toT := GongCopyBranchExplanationTextShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Eye3DShape:
		toT := GongCopyBranchEye3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EyeCornersSampledPoints3DShape:
		toT := GongCopyBranchEyeCornersSampledPoints3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EyeSampledPoints3DShape:
		toT := GongCopyBranchEyeSampledPoints3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EyeSeatBottomCurveShape:
		toT := GongCopyBranchEyeSeatBottomCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EyeStoolBottomCurveShape:
		toT := GongCopyBranchEyeStoolBottomCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EyeVolume3DShape:
		toT := GongCopyBranchEyeVolume3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GridPathShape:
		toT := GongCopyBranchGridPathShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurve2D:
		toT := GongCopyBranchGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurve2DRibbon:
		toT := GongCopyBranchGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurve2DRibbonEndShape:
		toT := GongCopyBranchGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurve2DRibbonStartShape:
		toT := GongCopyBranchGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurveRhombusGridShape:
		toT := GongCopyBranchGrowthCurveRhombusGridShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurveRhombusShape:
		toT := GongCopyBranchGrowthCurveRhombusShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthVectorShape:
		toT := GongCopyBranchGrowthVectorShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *InitialRhombusGridShape:
		toT := GongCopyBranchInitialRhombusGridShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *InitialRhombusShape:
		toT := GongCopyBranchInitialRhombusShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Key3DShape:
		toT := GongCopyBranchKey3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *KeyHole3DShape:
		toT := GongCopyBranchKeyHole3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *KeyHoleShape:
		toT := GongCopyBranchKeyHoleShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Leaves3DShape:
		toT := GongCopyBranchLeaves3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := GongCopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MidArcVectorShape:
		toT := GongCopyBranchMidArcVectorShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MidArcVectorShapeGrid:
		toT := GongCopyBranchMidArcVectorShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *OriginalPoints3DShape:
		toT := GongCopyBranchOriginalPoints3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParastichyMCurves3DShape:
		toT := GongCopyBranchParastichyMCurves3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParastichyNCurves3DShape:
		toT := GongCopyBranchParastichyNCurves3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DRibbon:
		toT := GongCopyBranchPartiallyGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DRibbonEndShape:
		toT := GongCopyBranchPartiallyGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DRibbonStartShape:
		toT := GongCopyBranchPartiallyGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectory:
		toT := GongCopyBranchPartiallyGrowthCurve2DTrajectory(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		toT := GongCopyBranchPartiallyGrowthCurve2DTrajectoryP1CurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		toT := GongCopyBranchPartiallyGrowthCurve2DTrajectoryP1P2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		toT := GongCopyBranchPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		toT := GongCopyBranchPartiallyGrowthCurve2DTrajectoryP1PointShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		toT := GongCopyBranchPartiallyGrowthCurve2DTrajectoryP2CurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		toT := GongCopyBranchPartiallyGrowthCurve2DTrajectoryP2PointShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyGrowthCurve2DTrajectoryShape:
		toT := GongCopyBranchPartiallyGrowthCurve2DTrajectoryShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyRotatedSeatBottomCurveShape:
		toT := GongCopyBranchPartiallyRotatedSeatBottomCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyRotatedSeatTopCurveShape:
		toT := GongCopyBranchPartiallyRotatedSeatTopCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartiallyRotatedTorusShape:
		toT := GongCopyBranchPartiallyRotatedTorusShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PerpendicularVector:
		toT := GongCopyBranchPerpendicularVector(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PerpendicularVectorGrid:
		toT := GongCopyBranchPerpendicularVectorGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PerpendicularVectorGridHalfway:
		toT := GongCopyBranchPerpendicularVectorGridHalfway(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PerpendicularVectorHalfway:
		toT := GongCopyBranchPerpendicularVectorHalfway(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Plant2DDiagram:
		toT := GongCopyBranchPlant2DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Plant3DDiagram:
		toT := GongCopyBranchPlant3DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PlantAbstract:
		toT := GongCopyBranchPlantAbstract(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PlantCircumferenceShape:
		toT := GongCopyBranchPlantCircumferenceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PointsAndLines3DShape:
		toT := GongCopyBranchPointsAndLines3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PxShape:
		toT := GongCopyBranchPxShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Rendered3DShape:
		toT := GongCopyBranchRendered3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RhombusShape:
		toT := GongCopyBranchRhombusShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RhombusStuff:
		toT := GongCopyBranchRhombusStuff(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RotatedRhombusGridShape:
		toT := GongCopyBranchRotatedRhombusGridShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RotatedRhombusShape:
		toT := GongCopyBranchRotatedRhombusShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RotatedSampledPoints3DShape:
		toT := GongCopyBranchRotatedSampledPoints3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RotatedSeatAndLegs3DShape:
		toT := GongCopyBranchRotatedSeatAndLegs3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SampledPoints3DShape:
		toT := GongCopyBranchSampledPoints3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Seat3DShape:
		toT := GongCopyBranchSeat3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SeatAndLegs3DShape:
		toT := GongCopyBranchSeatAndLegs3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SeatBottomCurveShape:
		toT := GongCopyBranchSeatBottomCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SeatTopCurveShape:
		toT := GongCopyBranchSeatTopCurveShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedBottomTopStartArcShape:
		toT := GongCopyBranchShiftedBottomTopStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedBottomTopStartArcShapeGrid:
		toT := GongCopyBranchShiftedBottomTopStartArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftGrowthCurve2DRibbon:
		toT := GongCopyBranchShiftedLeftGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftGrowthCurve2DRibbonEndShape:
		toT := GongCopyBranchShiftedLeftGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftGrowthCurve2DRibbonStartShape:
		toT := GongCopyBranchShiftedLeftGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbon:
		toT := GongCopyBranchShiftedLeftPartiallyGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape:
		toT := GongCopyBranchShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape:
		toT := GongCopyBranchShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftStackGrowthCurveEndArcShape:
		toT := GongCopyBranchShiftedLeftStackGrowthCurveEndArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftStackGrowthCurveStartArcShape:
		toT := GongCopyBranchShiftedLeftStackGrowthCurveStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftStackNormalVector:
		toT := GongCopyBranchShiftedLeftStackNormalVector(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftStackOfGrowthCurve:
		toT := GongCopyBranchShiftedLeftStackOfGrowthCurve(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedLeftStackOfNormalVector:
		toT := GongCopyBranchShiftedLeftStackOfNormalVector(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedRightGrowthCurve2DRibbon:
		toT := GongCopyBranchShiftedRightGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		toT := GongCopyBranchShiftedRightGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		toT := GongCopyBranchShiftedRightGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackGrowthCurve2DEndHalfwayArcShape:
		toT := GongCopyBranchStackGrowthCurve2DEndHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackGrowthCurve2DRibbonEndShape:
		toT := GongCopyBranchStackGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackGrowthCurve2DRibbonStartShape:
		toT := GongCopyBranchStackGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackGrowthCurve2DStartHalfwayArcShape:
		toT := GongCopyBranchStackGrowthCurve2DStartHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfGrowthCurve2D:
		toT := GongCopyBranchStackOfGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfGrowthCurve2DByGrowthVector:
		toT := GongCopyBranchStackOfGrowthCurve2DByGrowthVector(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfGrowthCurve2DRibbon:
		toT := GongCopyBranchStackOfGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfPartiallyRotatedTorusShape:
		toT := GongCopyBranchStackOfPartiallyRotatedTorusShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfRotatedGrowthCurve2D:
		toT := GongCopyBranchStackOfRotatedGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfRotatedGrowthCurve2DRibbon:
		toT := GongCopyBranchStackOfRotatedGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfRotatedVaseTrapezeRingsShape:
		toT := GongCopyBranchStackOfRotatedVaseTrapezeRingsShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfVaseTrapezeRingsShape:
		toT := GongCopyBranchStackOfVaseTrapezeRingsShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackRotatedGrowthCurve2DEndArcShape:
		toT := GongCopyBranchStackRotatedGrowthCurve2DEndArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackRotatedGrowthCurve2DRibbonEndShape:
		toT := GongCopyBranchStackRotatedGrowthCurve2DRibbonEndShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackRotatedGrowthCurve2DRibbonStartShape:
		toT := GongCopyBranchStackRotatedGrowthCurve2DRibbonStartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackRotatedGrowthCurve2DStartArcShape:
		toT := GongCopyBranchStackRotatedGrowthCurve2DStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartArcShape:
		toT := GongCopyBranchStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartArcShapeGrid:
		toT := GongCopyBranchStartArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartHalfwayArcShape:
		toT := GongCopyBranchStartHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartHalfwayArcShapeGrid:
		toT := GongCopyBranchStartHalfwayArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StemCylinder3DShape:
		toT := GongCopyBranchStemCylinder3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Stool2DDiagram:
		toT := GongCopyBranchStool2DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Stool3DDiagram:
		toT := GongCopyBranchStool3DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TiledFloor3DShape:
		toT := GongCopyBranchTiledFloor3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopCurvePlane1Shape:
		toT := GongCopyBranchTopCurvePlane1Shape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopCurvePlane2Shape:
		toT := GongCopyBranchTopCurvePlane2Shape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopEndArcShape:
		toT := GongCopyBranchTopEndArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopEndArcShapeGrid:
		toT := GongCopyBranchTopEndArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopEndHalfwayArcShape:
		toT := GongCopyBranchTopEndHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopEndHalfwayArcShapeGrid:
		toT := GongCopyBranchTopEndHalfwayArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopGrowthCurve2D:
		toT := GongCopyBranchTopGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopMidArcVectorShape:
		toT := GongCopyBranchTopMidArcVectorShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopMidArcVectorShapeGrid:
		toT := GongCopyBranchTopMidArcVectorShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		toT := GongCopyBranchTopStackGrowthCurve2DEndHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		toT := GongCopyBranchTopStackGrowthCurve2DStartHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackOfGrowthCurve2D:
		toT := GongCopyBranchTopStackOfGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackOfRotatedGrowthCurve2D:
		toT := GongCopyBranchTopStackOfRotatedGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		toT := GongCopyBranchTopStackOfRotatedGrowthCurve2DEndArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		toT := GongCopyBranchTopStackOfRotatedGrowthCurve2DStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStartArcShape:
		toT := GongCopyBranchTopStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStartArcShapeGrid:
		toT := GongCopyBranchTopStartArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStartHalfwayArcShape:
		toT := GongCopyBranchTopStartHalfwayArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStartHalfwayArcShapeGrid:
		toT := GongCopyBranchTopStartHalfwayArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Torus3DShape:
		toT := GongCopyBranchTorus3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TorusEdge3DShape:
		toT := GongCopyBranchTorusEdge3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TorusStackShape:
		toT := GongCopyBranchTorusStackShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TubeVase3DDiagram:
		toT := GongCopyBranchTubeVase3DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TubeVaseAbstract:
		toT := GongCopyBranchTubeVaseAbstract(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Vase2DDiagram:
		toT := GongCopyBranchVase2DDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VaseTrapezeRingShape:
		toT := GongCopyBranchVaseTrapezeRingShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VerticalTorusStackShape:
		toT := GongCopyBranchVerticalTorusStackShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VolumeKey3DShape:
		toT := GongCopyBranchVolumeKey3DShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAngle0Shape(mapOrigCopy map[any]any, angle0shapeFrom *Angle0Shape) (angle0shapeTo *Angle0Shape) {
	var alreadyCopied bool
	angle0shapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, angle0shapeFrom)
	if alreadyCopied {
		return
	}
	angle0shapeFrom.GongCopyBasicFields(angle0shapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArcNormalVectorShape(mapOrigCopy map[any]any, arcnormalvectorshapeFrom *ArcNormalVectorShape) (arcnormalvectorshapeTo *ArcNormalVectorShape) {
	var alreadyCopied bool
	arcnormalvectorshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, arcnormalvectorshapeFrom)
	if alreadyCopied {
		return
	}
	arcnormalvectorshapeFrom.GongCopyBasicFields(arcnormalvectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArcNormalVectorShapeGrid(mapOrigCopy map[any]any, arcnormalvectorshapegridFrom *ArcNormalVectorShapeGrid) (arcnormalvectorshapegridTo *ArcNormalVectorShapeGrid) {
	var alreadyCopied bool
	arcnormalvectorshapegridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, arcnormalvectorshapegridFrom)
	if alreadyCopied {
		return
	}
	arcnormalvectorshapegridFrom.GongCopyBasicFields(arcnormalvectorshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAxesShape(mapOrigCopy map[any]any, axesshapeFrom *AxesShape) (axesshapeTo *AxesShape) {
	var alreadyCopied bool
	axesshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, axesshapeFrom)
	if alreadyCopied {
		return
	}
	axesshapeFrom.GongCopyBasicFields(axesshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBaseVectorShape(mapOrigCopy map[any]any, basevectorshapeFrom *BaseVectorShape) (basevectorshapeTo *BaseVectorShape) {
	var alreadyCopied bool
	basevectorshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, basevectorshapeFrom)
	if alreadyCopied {
		return
	}
	basevectorshapeFrom.GongCopyBasicFields(basevectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBaseVectorShapeGrid(mapOrigCopy map[any]any, basevectorshapegridFrom *BaseVectorShapeGrid) (basevectorshapegridTo *BaseVectorShapeGrid) {
	var alreadyCopied bool
	basevectorshapegridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, basevectorshapegridFrom)
	if alreadyCopied {
		return
	}
	basevectorshapegridFrom.GongCopyBasicFields(basevectorshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBottomCurvePlane1Shape(mapOrigCopy map[any]any, bottomcurveplane1shapeFrom *BottomCurvePlane1Shape) (bottomcurveplane1shapeTo *BottomCurvePlane1Shape) {
	var alreadyCopied bool
	bottomcurveplane1shapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bottomcurveplane1shapeFrom)
	if alreadyCopied {
		return
	}
	bottomcurveplane1shapeFrom.GongCopyBasicFields(bottomcurveplane1shapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBottomCurvePlane2Shape(mapOrigCopy map[any]any, bottomcurveplane2shapeFrom *BottomCurvePlane2Shape) (bottomcurveplane2shapeTo *BottomCurvePlane2Shape) {
	var alreadyCopied bool
	bottomcurveplane2shapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bottomcurveplane2shapeFrom)
	if alreadyCopied {
		return
	}
	bottomcurveplane2shapeFrom.GongCopyBasicFields(bottomcurveplane2shapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchChosenP1P2PairShape(mapOrigCopy map[any]any, chosenp1p2pairshapeFrom *ChosenP1P2PairShape) (chosenp1p2pairshapeTo *ChosenP1P2PairShape) {
	var alreadyCopied bool
	chosenp1p2pairshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, chosenp1p2pairshapeFrom)
	if alreadyCopied {
		return
	}
	chosenp1p2pairshapeFrom.GongCopyBasicFields(chosenp1p2pairshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCircleGridShape(mapOrigCopy map[any]any, circlegridshapeFrom *CircleGridShape) (circlegridshapeTo *CircleGridShape) {
	var alreadyCopied bool
	circlegridshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, circlegridshapeFrom)
	if alreadyCopied {
		return
	}
	circlegridshapeFrom.GongCopyBasicFields(circlegridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCircumference3DShape(mapOrigCopy map[any]any, circumference3dshapeFrom *Circumference3DShape) (circumference3dshapeTo *Circumference3DShape) {
	var alreadyCopied bool
	circumference3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, circumference3dshapeFrom)
	if alreadyCopied {
		return
	}
	circumference3dshapeFrom.GongCopyBasicFields(circumference3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchClock2DDiagram(mapOrigCopy map[any]any, clock2ddiagramFrom *Clock2DDiagram) (clock2ddiagramTo *Clock2DDiagram) {
	var alreadyCopied bool
	clock2ddiagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, clock2ddiagramFrom)
	if alreadyCopied {
		return
	}
	clock2ddiagramFrom.GongCopyBasicFields(clock2ddiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchClock3DDiagram(mapOrigCopy map[any]any, clock3ddiagramFrom *Clock3DDiagram) (clock3ddiagramTo *Clock3DDiagram) {
	var alreadyCopied bool
	clock3ddiagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, clock3ddiagramFrom)
	if alreadyCopied {
		return
	}
	clock3ddiagramFrom.GongCopyBasicFields(clock3ddiagramTo)

	//insertion point for the staging of instances referenced by pointers
	if clock3ddiagramFrom.SampledPoints3DShape != nil {
		clock3ddiagramTo.SampledPoints3DShape = GongCopyBranchSampledPoints3DShape(mapOrigCopy, clock3ddiagramFrom.SampledPoints3DShape)
	}
	if clock3ddiagramFrom.Rendered3DShape != nil {
		clock3ddiagramTo.Rendered3DShape = GongCopyBranchRendered3DShape(mapOrigCopy, clock3ddiagramFrom.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchClockTopCurveShape(mapOrigCopy map[any]any, clocktopcurveshapeFrom *ClockTopCurveShape) (clocktopcurveshapeTo *ClockTopCurveShape) {
	var alreadyCopied bool
	clocktopcurveshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, clocktopcurveshapeFrom)
	if alreadyCopied {
		return
	}
	clocktopcurveshapeFrom.GongCopyBasicFields(clocktopcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCutLine3DShape(mapOrigCopy map[any]any, cutline3dshapeFrom *CutLine3DShape) (cutline3dshapeTo *CutLine3DShape) {
	var alreadyCopied bool
	cutline3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, cutline3dshapeFrom)
	if alreadyCopied {
		return
	}
	cutline3dshapeFrom.GongCopyBasicFields(cutline3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEndArcShape(mapOrigCopy map[any]any, endarcshapeFrom *EndArcShape) (endarcshapeTo *EndArcShape) {
	var alreadyCopied bool
	endarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, endarcshapeFrom)
	if alreadyCopied {
		return
	}
	endarcshapeFrom.GongCopyBasicFields(endarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEndArcShapeGrid(mapOrigCopy map[any]any, endarcshapegridFrom *EndArcShapeGrid) (endarcshapegridTo *EndArcShapeGrid) {
	var alreadyCopied bool
	endarcshapegridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, endarcshapegridFrom)
	if alreadyCopied {
		return
	}
	endarcshapegridFrom.GongCopyBasicFields(endarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEndHalfwayArcShape(mapOrigCopy map[any]any, endhalfwayarcshapeFrom *EndHalfwayArcShape) (endhalfwayarcshapeTo *EndHalfwayArcShape) {
	var alreadyCopied bool
	endhalfwayarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, endhalfwayarcshapeFrom)
	if alreadyCopied {
		return
	}
	endhalfwayarcshapeFrom.GongCopyBasicFields(endhalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEndHalfwayArcShapeGrid(mapOrigCopy map[any]any, endhalfwayarcshapegridFrom *EndHalfwayArcShapeGrid) (endhalfwayarcshapegridTo *EndHalfwayArcShapeGrid) {
	var alreadyCopied bool
	endhalfwayarcshapegridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, endhalfwayarcshapegridFrom)
	if alreadyCopied {
		return
	}
	endhalfwayarcshapegridFrom.GongCopyBasicFields(endhalfwayarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchExplanationTextShape(mapOrigCopy map[any]any, explanationtextshapeFrom *ExplanationTextShape) (explanationtextshapeTo *ExplanationTextShape) {
	var alreadyCopied bool
	explanationtextshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, explanationtextshapeFrom)
	if alreadyCopied {
		return
	}
	explanationtextshapeFrom.GongCopyBasicFields(explanationtextshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEye3DShape(mapOrigCopy map[any]any, eye3dshapeFrom *Eye3DShape) (eye3dshapeTo *Eye3DShape) {
	var alreadyCopied bool
	eye3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, eye3dshapeFrom)
	if alreadyCopied {
		return
	}
	eye3dshapeFrom.GongCopyBasicFields(eye3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEyeCornersSampledPoints3DShape(mapOrigCopy map[any]any, eyecornerssampledpoints3dshapeFrom *EyeCornersSampledPoints3DShape) (eyecornerssampledpoints3dshapeTo *EyeCornersSampledPoints3DShape) {
	var alreadyCopied bool
	eyecornerssampledpoints3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, eyecornerssampledpoints3dshapeFrom)
	if alreadyCopied {
		return
	}
	eyecornerssampledpoints3dshapeFrom.GongCopyBasicFields(eyecornerssampledpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEyeSampledPoints3DShape(mapOrigCopy map[any]any, eyesampledpoints3dshapeFrom *EyeSampledPoints3DShape) (eyesampledpoints3dshapeTo *EyeSampledPoints3DShape) {
	var alreadyCopied bool
	eyesampledpoints3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, eyesampledpoints3dshapeFrom)
	if alreadyCopied {
		return
	}
	eyesampledpoints3dshapeFrom.GongCopyBasicFields(eyesampledpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEyeSeatBottomCurveShape(mapOrigCopy map[any]any, eyeseatbottomcurveshapeFrom *EyeSeatBottomCurveShape) (eyeseatbottomcurveshapeTo *EyeSeatBottomCurveShape) {
	var alreadyCopied bool
	eyeseatbottomcurveshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, eyeseatbottomcurveshapeFrom)
	if alreadyCopied {
		return
	}
	eyeseatbottomcurveshapeFrom.GongCopyBasicFields(eyeseatbottomcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEyeStoolBottomCurveShape(mapOrigCopy map[any]any, eyestoolbottomcurveshapeFrom *EyeStoolBottomCurveShape) (eyestoolbottomcurveshapeTo *EyeStoolBottomCurveShape) {
	var alreadyCopied bool
	eyestoolbottomcurveshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, eyestoolbottomcurveshapeFrom)
	if alreadyCopied {
		return
	}
	eyestoolbottomcurveshapeFrom.GongCopyBasicFields(eyestoolbottomcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEyeVolume3DShape(mapOrigCopy map[any]any, eyevolume3dshapeFrom *EyeVolume3DShape) (eyevolume3dshapeTo *EyeVolume3DShape) {
	var alreadyCopied bool
	eyevolume3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, eyevolume3dshapeFrom)
	if alreadyCopied {
		return
	}
	eyevolume3dshapeFrom.GongCopyBasicFields(eyevolume3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGridPathShape(mapOrigCopy map[any]any, gridpathshapeFrom *GridPathShape) (gridpathshapeTo *GridPathShape) {
	var alreadyCopied bool
	gridpathshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gridpathshapeFrom)
	if alreadyCopied {
		return
	}
	gridpathshapeFrom.GongCopyBasicFields(gridpathshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthCurve2D(mapOrigCopy map[any]any, growthcurve2dFrom *GrowthCurve2D) (growthcurve2dTo *GrowthCurve2D) {
	var alreadyCopied bool
	growthcurve2dTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, growthcurve2dFrom)
	if alreadyCopied {
		return
	}
	growthcurve2dFrom.GongCopyBasicFields(growthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthCurve2DRibbon(mapOrigCopy map[any]any, growthcurve2dribbonFrom *GrowthCurve2DRibbon) (growthcurve2dribbonTo *GrowthCurve2DRibbon) {
	var alreadyCopied bool
	growthcurve2dribbonTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, growthcurve2dribbonFrom)
	if alreadyCopied {
		return
	}
	growthcurve2dribbonFrom.GongCopyBasicFields(growthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, growthcurve2dribbonendshapeFrom *GrowthCurve2DRibbonEndShape) (growthcurve2dribbonendshapeTo *GrowthCurve2DRibbonEndShape) {
	var alreadyCopied bool
	growthcurve2dribbonendshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, growthcurve2dribbonendshapeFrom)
	if alreadyCopied {
		return
	}
	growthcurve2dribbonendshapeFrom.GongCopyBasicFields(growthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, growthcurve2dribbonstartshapeFrom *GrowthCurve2DRibbonStartShape) (growthcurve2dribbonstartshapeTo *GrowthCurve2DRibbonStartShape) {
	var alreadyCopied bool
	growthcurve2dribbonstartshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, growthcurve2dribbonstartshapeFrom)
	if alreadyCopied {
		return
	}
	growthcurve2dribbonstartshapeFrom.GongCopyBasicFields(growthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthCurveRhombusGridShape(mapOrigCopy map[any]any, growthcurverhombusgridshapeFrom *GrowthCurveRhombusGridShape) (growthcurverhombusgridshapeTo *GrowthCurveRhombusGridShape) {
	var alreadyCopied bool
	growthcurverhombusgridshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, growthcurverhombusgridshapeFrom)
	if alreadyCopied {
		return
	}
	growthcurverhombusgridshapeFrom.GongCopyBasicFields(growthcurverhombusgridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthCurveRhombusShape(mapOrigCopy map[any]any, growthcurverhombusshapeFrom *GrowthCurveRhombusShape) (growthcurverhombusshapeTo *GrowthCurveRhombusShape) {
	var alreadyCopied bool
	growthcurverhombusshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, growthcurverhombusshapeFrom)
	if alreadyCopied {
		return
	}
	growthcurverhombusshapeFrom.GongCopyBasicFields(growthcurverhombusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthVectorShape(mapOrigCopy map[any]any, growthvectorshapeFrom *GrowthVectorShape) (growthvectorshapeTo *GrowthVectorShape) {
	var alreadyCopied bool
	growthvectorshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, growthvectorshapeFrom)
	if alreadyCopied {
		return
	}
	growthvectorshapeFrom.GongCopyBasicFields(growthvectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInitialRhombusGridShape(mapOrigCopy map[any]any, initialrhombusgridshapeFrom *InitialRhombusGridShape) (initialrhombusgridshapeTo *InitialRhombusGridShape) {
	var alreadyCopied bool
	initialrhombusgridshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, initialrhombusgridshapeFrom)
	if alreadyCopied {
		return
	}
	initialrhombusgridshapeFrom.GongCopyBasicFields(initialrhombusgridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInitialRhombusShape(mapOrigCopy map[any]any, initialrhombusshapeFrom *InitialRhombusShape) (initialrhombusshapeTo *InitialRhombusShape) {
	var alreadyCopied bool
	initialrhombusshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, initialrhombusshapeFrom)
	if alreadyCopied {
		return
	}
	initialrhombusshapeFrom.GongCopyBasicFields(initialrhombusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKey3DShape(mapOrigCopy map[any]any, key3dshapeFrom *Key3DShape) (key3dshapeTo *Key3DShape) {
	var alreadyCopied bool
	key3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, key3dshapeFrom)
	if alreadyCopied {
		return
	}
	key3dshapeFrom.GongCopyBasicFields(key3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKeyHole3DShape(mapOrigCopy map[any]any, keyhole3dshapeFrom *KeyHole3DShape) (keyhole3dshapeTo *KeyHole3DShape) {
	var alreadyCopied bool
	keyhole3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, keyhole3dshapeFrom)
	if alreadyCopied {
		return
	}
	keyhole3dshapeFrom.GongCopyBasicFields(keyhole3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKeyHoleShape(mapOrigCopy map[any]any, keyholeshapeFrom *KeyHoleShape) (keyholeshapeTo *KeyHoleShape) {
	var alreadyCopied bool
	keyholeshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, keyholeshapeFrom)
	if alreadyCopied {
		return
	}
	keyholeshapeFrom.GongCopyBasicFields(keyholeshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLeaves3DShape(mapOrigCopy map[any]any, leaves3dshapeFrom *Leaves3DShape) (leaves3dshapeTo *Leaves3DShape) {
	var alreadyCopied bool
	leaves3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, leaves3dshapeFrom)
	if alreadyCopied {
		return
	}
	leaves3dshapeFrom.GongCopyBasicFields(leaves3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {
	var alreadyCopied bool
	libraryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, libraryFrom)
	if alreadyCopied {
		return
	}
	libraryFrom.GongCopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plantabstract := range libraryFrom.Plants {
		libraryTo.Plants = append(libraryTo.Plants, GongCopyBranchPlantAbstract(mapOrigCopy, _plantabstract))
	}
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, GongCopyBranchLibrary(mapOrigCopy, _library))
	}

	return
}

func GongCopyBranchMidArcVectorShape(mapOrigCopy map[any]any, midarcvectorshapeFrom *MidArcVectorShape) (midarcvectorshapeTo *MidArcVectorShape) {
	var alreadyCopied bool
	midarcvectorshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, midarcvectorshapeFrom)
	if alreadyCopied {
		return
	}
	midarcvectorshapeFrom.GongCopyBasicFields(midarcvectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMidArcVectorShapeGrid(mapOrigCopy map[any]any, midarcvectorshapegridFrom *MidArcVectorShapeGrid) (midarcvectorshapegridTo *MidArcVectorShapeGrid) {
	var alreadyCopied bool
	midarcvectorshapegridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, midarcvectorshapegridFrom)
	if alreadyCopied {
		return
	}
	midarcvectorshapegridFrom.GongCopyBasicFields(midarcvectorshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOriginalPoints3DShape(mapOrigCopy map[any]any, originalpoints3dshapeFrom *OriginalPoints3DShape) (originalpoints3dshapeTo *OriginalPoints3DShape) {
	var alreadyCopied bool
	originalpoints3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, originalpoints3dshapeFrom)
	if alreadyCopied {
		return
	}
	originalpoints3dshapeFrom.GongCopyBasicFields(originalpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchParastichyMCurves3DShape(mapOrigCopy map[any]any, parastichymcurves3dshapeFrom *ParastichyMCurves3DShape) (parastichymcurves3dshapeTo *ParastichyMCurves3DShape) {
	var alreadyCopied bool
	parastichymcurves3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, parastichymcurves3dshapeFrom)
	if alreadyCopied {
		return
	}
	parastichymcurves3dshapeFrom.GongCopyBasicFields(parastichymcurves3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchParastichyNCurves3DShape(mapOrigCopy map[any]any, parastichyncurves3dshapeFrom *ParastichyNCurves3DShape) (parastichyncurves3dshapeTo *ParastichyNCurves3DShape) {
	var alreadyCopied bool
	parastichyncurves3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, parastichyncurves3dshapeFrom)
	if alreadyCopied {
		return
	}
	parastichyncurves3dshapeFrom.GongCopyBasicFields(parastichyncurves3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DRibbon(mapOrigCopy map[any]any, partiallygrowthcurve2dribbonFrom *PartiallyGrowthCurve2DRibbon) (partiallygrowthcurve2dribbonTo *PartiallyGrowthCurve2DRibbon) {
	var alreadyCopied bool
	partiallygrowthcurve2dribbonTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallygrowthcurve2dribbonFrom)
	if alreadyCopied {
		return
	}
	partiallygrowthcurve2dribbonFrom.GongCopyBasicFields(partiallygrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, partiallygrowthcurve2dribbonendshapeFrom *PartiallyGrowthCurve2DRibbonEndShape) (partiallygrowthcurve2dribbonendshapeTo *PartiallyGrowthCurve2DRibbonEndShape) {
	var alreadyCopied bool
	partiallygrowthcurve2dribbonendshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallygrowthcurve2dribbonendshapeFrom)
	if alreadyCopied {
		return
	}
	partiallygrowthcurve2dribbonendshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, partiallygrowthcurve2dribbonstartshapeFrom *PartiallyGrowthCurve2DRibbonStartShape) (partiallygrowthcurve2dribbonstartshapeTo *PartiallyGrowthCurve2DRibbonStartShape) {
	var alreadyCopied bool
	partiallygrowthcurve2dribbonstartshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallygrowthcurve2dribbonstartshapeFrom)
	if alreadyCopied {
		return
	}
	partiallygrowthcurve2dribbonstartshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectory(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryFrom *PartiallyGrowthCurve2DTrajectory) (partiallygrowthcurve2dtrajectoryTo *PartiallyGrowthCurve2DTrajectory) {
	var alreadyCopied bool
	partiallygrowthcurve2dtrajectoryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallygrowthcurve2dtrajectoryFrom)
	if alreadyCopied {
		return
	}
	partiallygrowthcurve2dtrajectoryFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryP1CurveShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp1curveshapeFrom *PartiallyGrowthCurve2DTrajectoryP1CurveShape) (partiallygrowthcurve2dtrajectoryp1curveshapeTo *PartiallyGrowthCurve2DTrajectoryP1CurveShape) {
	var alreadyCopied bool
	partiallygrowthcurve2dtrajectoryp1curveshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallygrowthcurve2dtrajectoryp1curveshapeFrom)
	if alreadyCopied {
		return
	}
	partiallygrowthcurve2dtrajectoryp1curveshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryp1curveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryP1P2(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp1p2From *PartiallyGrowthCurve2DTrajectoryP1P2) (partiallygrowthcurve2dtrajectoryp1p2To *PartiallyGrowthCurve2DTrajectoryP1P2) {
	var alreadyCopied bool
	partiallygrowthcurve2dtrajectoryp1p2To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallygrowthcurve2dtrajectoryp1p2From)
	if alreadyCopied {
		return
	}
	partiallygrowthcurve2dtrajectoryp1p2From.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryp1p2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFrom *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) (partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) {
	var alreadyCopied bool
	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFrom)
	if alreadyCopied {
		return
	}
	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryP1PointShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp1pointshapeFrom *PartiallyGrowthCurve2DTrajectoryP1PointShape) (partiallygrowthcurve2dtrajectoryp1pointshapeTo *PartiallyGrowthCurve2DTrajectoryP1PointShape) {
	var alreadyCopied bool
	partiallygrowthcurve2dtrajectoryp1pointshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallygrowthcurve2dtrajectoryp1pointshapeFrom)
	if alreadyCopied {
		return
	}
	partiallygrowthcurve2dtrajectoryp1pointshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryp1pointshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryP2CurveShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp2curveshapeFrom *PartiallyGrowthCurve2DTrajectoryP2CurveShape) (partiallygrowthcurve2dtrajectoryp2curveshapeTo *PartiallyGrowthCurve2DTrajectoryP2CurveShape) {
	var alreadyCopied bool
	partiallygrowthcurve2dtrajectoryp2curveshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallygrowthcurve2dtrajectoryp2curveshapeFrom)
	if alreadyCopied {
		return
	}
	partiallygrowthcurve2dtrajectoryp2curveshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryp2curveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryP2PointShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp2pointshapeFrom *PartiallyGrowthCurve2DTrajectoryP2PointShape) (partiallygrowthcurve2dtrajectoryp2pointshapeTo *PartiallyGrowthCurve2DTrajectoryP2PointShape) {
	var alreadyCopied bool
	partiallygrowthcurve2dtrajectoryp2pointshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallygrowthcurve2dtrajectoryp2pointshapeFrom)
	if alreadyCopied {
		return
	}
	partiallygrowthcurve2dtrajectoryp2pointshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryp2pointshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryshapeFrom *PartiallyGrowthCurve2DTrajectoryShape) (partiallygrowthcurve2dtrajectoryshapeTo *PartiallyGrowthCurve2DTrajectoryShape) {
	var alreadyCopied bool
	partiallygrowthcurve2dtrajectoryshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallygrowthcurve2dtrajectoryshapeFrom)
	if alreadyCopied {
		return
	}
	partiallygrowthcurve2dtrajectoryshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyRotatedSeatBottomCurveShape(mapOrigCopy map[any]any, partiallyrotatedseatbottomcurveshapeFrom *PartiallyRotatedSeatBottomCurveShape) (partiallyrotatedseatbottomcurveshapeTo *PartiallyRotatedSeatBottomCurveShape) {
	var alreadyCopied bool
	partiallyrotatedseatbottomcurveshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallyrotatedseatbottomcurveshapeFrom)
	if alreadyCopied {
		return
	}
	partiallyrotatedseatbottomcurveshapeFrom.GongCopyBasicFields(partiallyrotatedseatbottomcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyRotatedSeatTopCurveShape(mapOrigCopy map[any]any, partiallyrotatedseattopcurveshapeFrom *PartiallyRotatedSeatTopCurveShape) (partiallyrotatedseattopcurveshapeTo *PartiallyRotatedSeatTopCurveShape) {
	var alreadyCopied bool
	partiallyrotatedseattopcurveshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallyrotatedseattopcurveshapeFrom)
	if alreadyCopied {
		return
	}
	partiallyrotatedseattopcurveshapeFrom.GongCopyBasicFields(partiallyrotatedseattopcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyRotatedTorusShape(mapOrigCopy map[any]any, partiallyrotatedtorusshapeFrom *PartiallyRotatedTorusShape) (partiallyrotatedtorusshapeTo *PartiallyRotatedTorusShape) {
	var alreadyCopied bool
	partiallyrotatedtorusshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partiallyrotatedtorusshapeFrom)
	if alreadyCopied {
		return
	}
	partiallyrotatedtorusshapeFrom.GongCopyBasicFields(partiallyrotatedtorusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPerpendicularVector(mapOrigCopy map[any]any, perpendicularvectorFrom *PerpendicularVector) (perpendicularvectorTo *PerpendicularVector) {
	var alreadyCopied bool
	perpendicularvectorTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, perpendicularvectorFrom)
	if alreadyCopied {
		return
	}
	perpendicularvectorFrom.GongCopyBasicFields(perpendicularvectorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPerpendicularVectorGrid(mapOrigCopy map[any]any, perpendicularvectorgridFrom *PerpendicularVectorGrid) (perpendicularvectorgridTo *PerpendicularVectorGrid) {
	var alreadyCopied bool
	perpendicularvectorgridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, perpendicularvectorgridFrom)
	if alreadyCopied {
		return
	}
	perpendicularvectorgridFrom.GongCopyBasicFields(perpendicularvectorgridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPerpendicularVectorGridHalfway(mapOrigCopy map[any]any, perpendicularvectorgridhalfwayFrom *PerpendicularVectorGridHalfway) (perpendicularvectorgridhalfwayTo *PerpendicularVectorGridHalfway) {
	var alreadyCopied bool
	perpendicularvectorgridhalfwayTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, perpendicularvectorgridhalfwayFrom)
	if alreadyCopied {
		return
	}
	perpendicularvectorgridhalfwayFrom.GongCopyBasicFields(perpendicularvectorgridhalfwayTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPerpendicularVectorHalfway(mapOrigCopy map[any]any, perpendicularvectorhalfwayFrom *PerpendicularVectorHalfway) (perpendicularvectorhalfwayTo *PerpendicularVectorHalfway) {
	var alreadyCopied bool
	perpendicularvectorhalfwayTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, perpendicularvectorhalfwayFrom)
	if alreadyCopied {
		return
	}
	perpendicularvectorhalfwayFrom.GongCopyBasicFields(perpendicularvectorhalfwayTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPlant2DDiagram(mapOrigCopy map[any]any, plant2ddiagramFrom *Plant2DDiagram) (plant2ddiagramTo *Plant2DDiagram) {
	var alreadyCopied bool
	plant2ddiagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, plant2ddiagramFrom)
	if alreadyCopied {
		return
	}
	plant2ddiagramFrom.GongCopyBasicFields(plant2ddiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPlant3DDiagram(mapOrigCopy map[any]any, plant3ddiagramFrom *Plant3DDiagram) (plant3ddiagramTo *Plant3DDiagram) {
	var alreadyCopied bool
	plant3ddiagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, plant3ddiagramFrom)
	if alreadyCopied {
		return
	}
	plant3ddiagramFrom.GongCopyBasicFields(plant3ddiagramTo)

	//insertion point for the staging of instances referenced by pointers
	if plant3ddiagramFrom.StemCylinder3DShape != nil {
		plant3ddiagramTo.StemCylinder3DShape = GongCopyBranchStemCylinder3DShape(mapOrigCopy, plant3ddiagramFrom.StemCylinder3DShape)
	}
	if plant3ddiagramFrom.ParastichyNCurves3DShape != nil {
		plant3ddiagramTo.ParastichyNCurves3DShape = GongCopyBranchParastichyNCurves3DShape(mapOrigCopy, plant3ddiagramFrom.ParastichyNCurves3DShape)
	}
	if plant3ddiagramFrom.ParastichyMCurves3DShape != nil {
		plant3ddiagramTo.ParastichyMCurves3DShape = GongCopyBranchParastichyMCurves3DShape(mapOrigCopy, plant3ddiagramFrom.ParastichyMCurves3DShape)
	}
	if plant3ddiagramFrom.CutLine3DShape != nil {
		plant3ddiagramTo.CutLine3DShape = GongCopyBranchCutLine3DShape(mapOrigCopy, plant3ddiagramFrom.CutLine3DShape)
	}
	if plant3ddiagramFrom.Circumference3DShape != nil {
		plant3ddiagramTo.Circumference3DShape = GongCopyBranchCircumference3DShape(mapOrigCopy, plant3ddiagramFrom.Circumference3DShape)
	}
	if plant3ddiagramFrom.Leaves3DShape != nil {
		plant3ddiagramTo.Leaves3DShape = GongCopyBranchLeaves3DShape(mapOrigCopy, plant3ddiagramFrom.Leaves3DShape)
	}
	if plant3ddiagramFrom.Rendered3DShape != nil {
		plant3ddiagramTo.Rendered3DShape = GongCopyBranchRendered3DShape(mapOrigCopy, plant3ddiagramFrom.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPlantAbstract(mapOrigCopy map[any]any, plantabstractFrom *PlantAbstract) (plantabstractTo *PlantAbstract) {
	var alreadyCopied bool
	plantabstractTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, plantabstractFrom)
	if alreadyCopied {
		return
	}
	plantabstractFrom.GongCopyBasicFields(plantabstractTo)

	//insertion point for the staging of instances referenced by pointers
	if plantabstractFrom.TubeVaseAbstract != nil {
		plantabstractTo.TubeVaseAbstract = GongCopyBranchTubeVaseAbstract(mapOrigCopy, plantabstractFrom.TubeVaseAbstract)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plant2ddiagram := range plantabstractFrom.Plant2DDiagrams {
		plantabstractTo.Plant2DDiagrams = append(plantabstractTo.Plant2DDiagrams, GongCopyBranchPlant2DDiagram(mapOrigCopy, _plant2ddiagram))
	}
	for _, _plant3ddiagram := range plantabstractFrom.Plant3DDiagrams {
		plantabstractTo.Plant3DDiagrams = append(plantabstractTo.Plant3DDiagrams, GongCopyBranchPlant3DDiagram(mapOrigCopy, _plant3ddiagram))
	}
	for _, _vase2ddiagram := range plantabstractFrom.Vase2DDiagrams {
		plantabstractTo.Vase2DDiagrams = append(plantabstractTo.Vase2DDiagrams, GongCopyBranchVase2DDiagram(mapOrigCopy, _vase2ddiagram))
	}
	for _, _tubevase3ddiagram := range plantabstractFrom.TubeVase3DDiagrams {
		plantabstractTo.TubeVase3DDiagrams = append(plantabstractTo.TubeVase3DDiagrams, GongCopyBranchTubeVase3DDiagram(mapOrigCopy, _tubevase3ddiagram))
	}
	for _, _stool2ddiagram := range plantabstractFrom.Stool2DDiagrams {
		plantabstractTo.Stool2DDiagrams = append(plantabstractTo.Stool2DDiagrams, GongCopyBranchStool2DDiagram(mapOrigCopy, _stool2ddiagram))
	}
	for _, _stool3ddiagram := range plantabstractFrom.Stool3DDiagrams {
		plantabstractTo.Stool3DDiagrams = append(plantabstractTo.Stool3DDiagrams, GongCopyBranchStool3DDiagram(mapOrigCopy, _stool3ddiagram))
	}
	for _, _clock2ddiagram := range plantabstractFrom.Clock2DDiagrams {
		plantabstractTo.Clock2DDiagrams = append(plantabstractTo.Clock2DDiagrams, GongCopyBranchClock2DDiagram(mapOrigCopy, _clock2ddiagram))
	}
	for _, _clock3ddiagram := range plantabstractFrom.Clock3DDiagrams {
		plantabstractTo.Clock3DDiagrams = append(plantabstractTo.Clock3DDiagrams, GongCopyBranchClock3DDiagram(mapOrigCopy, _clock3ddiagram))
	}

	return
}

func GongCopyBranchPlantCircumferenceShape(mapOrigCopy map[any]any, plantcircumferenceshapeFrom *PlantCircumferenceShape) (plantcircumferenceshapeTo *PlantCircumferenceShape) {
	var alreadyCopied bool
	plantcircumferenceshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, plantcircumferenceshapeFrom)
	if alreadyCopied {
		return
	}
	plantcircumferenceshapeFrom.GongCopyBasicFields(plantcircumferenceshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPointsAndLines3DShape(mapOrigCopy map[any]any, pointsandlines3dshapeFrom *PointsAndLines3DShape) (pointsandlines3dshapeTo *PointsAndLines3DShape) {
	var alreadyCopied bool
	pointsandlines3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pointsandlines3dshapeFrom)
	if alreadyCopied {
		return
	}
	pointsandlines3dshapeFrom.GongCopyBasicFields(pointsandlines3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPxShape(mapOrigCopy map[any]any, pxshapeFrom *PxShape) (pxshapeTo *PxShape) {
	var alreadyCopied bool
	pxshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pxshapeFrom)
	if alreadyCopied {
		return
	}
	pxshapeFrom.GongCopyBasicFields(pxshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRendered3DShape(mapOrigCopy map[any]any, rendered3dshapeFrom *Rendered3DShape) (rendered3dshapeTo *Rendered3DShape) {
	var alreadyCopied bool
	rendered3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rendered3dshapeFrom)
	if alreadyCopied {
		return
	}
	rendered3dshapeFrom.GongCopyBasicFields(rendered3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRhombusShape(mapOrigCopy map[any]any, rhombusshapeFrom *RhombusShape) (rhombusshapeTo *RhombusShape) {
	var alreadyCopied bool
	rhombusshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rhombusshapeFrom)
	if alreadyCopied {
		return
	}
	rhombusshapeFrom.GongCopyBasicFields(rhombusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRhombusStuff(mapOrigCopy map[any]any, rhombusstuffFrom *RhombusStuff) (rhombusstuffTo *RhombusStuff) {
	var alreadyCopied bool
	rhombusstuffTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rhombusstuffFrom)
	if alreadyCopied {
		return
	}
	rhombusstuffFrom.GongCopyBasicFields(rhombusstuffTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRotatedRhombusGridShape(mapOrigCopy map[any]any, rotatedrhombusgridshapeFrom *RotatedRhombusGridShape) (rotatedrhombusgridshapeTo *RotatedRhombusGridShape) {
	var alreadyCopied bool
	rotatedrhombusgridshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rotatedrhombusgridshapeFrom)
	if alreadyCopied {
		return
	}
	rotatedrhombusgridshapeFrom.GongCopyBasicFields(rotatedrhombusgridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRotatedRhombusShape(mapOrigCopy map[any]any, rotatedrhombusshapeFrom *RotatedRhombusShape) (rotatedrhombusshapeTo *RotatedRhombusShape) {
	var alreadyCopied bool
	rotatedrhombusshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rotatedrhombusshapeFrom)
	if alreadyCopied {
		return
	}
	rotatedrhombusshapeFrom.GongCopyBasicFields(rotatedrhombusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRotatedSampledPoints3DShape(mapOrigCopy map[any]any, rotatedsampledpoints3dshapeFrom *RotatedSampledPoints3DShape) (rotatedsampledpoints3dshapeTo *RotatedSampledPoints3DShape) {
	var alreadyCopied bool
	rotatedsampledpoints3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rotatedsampledpoints3dshapeFrom)
	if alreadyCopied {
		return
	}
	rotatedsampledpoints3dshapeFrom.GongCopyBasicFields(rotatedsampledpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRotatedSeatAndLegs3DShape(mapOrigCopy map[any]any, rotatedseatandlegs3dshapeFrom *RotatedSeatAndLegs3DShape) (rotatedseatandlegs3dshapeTo *RotatedSeatAndLegs3DShape) {
	var alreadyCopied bool
	rotatedseatandlegs3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rotatedseatandlegs3dshapeFrom)
	if alreadyCopied {
		return
	}
	rotatedseatandlegs3dshapeFrom.GongCopyBasicFields(rotatedseatandlegs3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSampledPoints3DShape(mapOrigCopy map[any]any, sampledpoints3dshapeFrom *SampledPoints3DShape) (sampledpoints3dshapeTo *SampledPoints3DShape) {
	var alreadyCopied bool
	sampledpoints3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, sampledpoints3dshapeFrom)
	if alreadyCopied {
		return
	}
	sampledpoints3dshapeFrom.GongCopyBasicFields(sampledpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSeat3DShape(mapOrigCopy map[any]any, seat3dshapeFrom *Seat3DShape) (seat3dshapeTo *Seat3DShape) {
	var alreadyCopied bool
	seat3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, seat3dshapeFrom)
	if alreadyCopied {
		return
	}
	seat3dshapeFrom.GongCopyBasicFields(seat3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSeatAndLegs3DShape(mapOrigCopy map[any]any, seatandlegs3dshapeFrom *SeatAndLegs3DShape) (seatandlegs3dshapeTo *SeatAndLegs3DShape) {
	var alreadyCopied bool
	seatandlegs3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, seatandlegs3dshapeFrom)
	if alreadyCopied {
		return
	}
	seatandlegs3dshapeFrom.GongCopyBasicFields(seatandlegs3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSeatBottomCurveShape(mapOrigCopy map[any]any, seatbottomcurveshapeFrom *SeatBottomCurveShape) (seatbottomcurveshapeTo *SeatBottomCurveShape) {
	var alreadyCopied bool
	seatbottomcurveshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, seatbottomcurveshapeFrom)
	if alreadyCopied {
		return
	}
	seatbottomcurveshapeFrom.GongCopyBasicFields(seatbottomcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSeatTopCurveShape(mapOrigCopy map[any]any, seattopcurveshapeFrom *SeatTopCurveShape) (seattopcurveshapeTo *SeatTopCurveShape) {
	var alreadyCopied bool
	seattopcurveshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, seattopcurveshapeFrom)
	if alreadyCopied {
		return
	}
	seattopcurveshapeFrom.GongCopyBasicFields(seattopcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedBottomTopStartArcShape(mapOrigCopy map[any]any, shiftedbottomtopstartarcshapeFrom *ShiftedBottomTopStartArcShape) (shiftedbottomtopstartarcshapeTo *ShiftedBottomTopStartArcShape) {
	var alreadyCopied bool
	shiftedbottomtopstartarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedbottomtopstartarcshapeFrom)
	if alreadyCopied {
		return
	}
	shiftedbottomtopstartarcshapeFrom.GongCopyBasicFields(shiftedbottomtopstartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedBottomTopStartArcShapeGrid(mapOrigCopy map[any]any, shiftedbottomtopstartarcshapegridFrom *ShiftedBottomTopStartArcShapeGrid) (shiftedbottomtopstartarcshapegridTo *ShiftedBottomTopStartArcShapeGrid) {
	var alreadyCopied bool
	shiftedbottomtopstartarcshapegridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedbottomtopstartarcshapegridFrom)
	if alreadyCopied {
		return
	}
	shiftedbottomtopstartarcshapegridFrom.GongCopyBasicFields(shiftedbottomtopstartarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftGrowthCurve2DRibbon(mapOrigCopy map[any]any, shiftedleftgrowthcurve2dribbonFrom *ShiftedLeftGrowthCurve2DRibbon) (shiftedleftgrowthcurve2dribbonTo *ShiftedLeftGrowthCurve2DRibbon) {
	var alreadyCopied bool
	shiftedleftgrowthcurve2dribbonTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedleftgrowthcurve2dribbonFrom)
	if alreadyCopied {
		return
	}
	shiftedleftgrowthcurve2dribbonFrom.GongCopyBasicFields(shiftedleftgrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, shiftedleftgrowthcurve2dribbonendshapeFrom *ShiftedLeftGrowthCurve2DRibbonEndShape) (shiftedleftgrowthcurve2dribbonendshapeTo *ShiftedLeftGrowthCurve2DRibbonEndShape) {
	var alreadyCopied bool
	shiftedleftgrowthcurve2dribbonendshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedleftgrowthcurve2dribbonendshapeFrom)
	if alreadyCopied {
		return
	}
	shiftedleftgrowthcurve2dribbonendshapeFrom.GongCopyBasicFields(shiftedleftgrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, shiftedleftgrowthcurve2dribbonstartshapeFrom *ShiftedLeftGrowthCurve2DRibbonStartShape) (shiftedleftgrowthcurve2dribbonstartshapeTo *ShiftedLeftGrowthCurve2DRibbonStartShape) {
	var alreadyCopied bool
	shiftedleftgrowthcurve2dribbonstartshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedleftgrowthcurve2dribbonstartshapeFrom)
	if alreadyCopied {
		return
	}
	shiftedleftgrowthcurve2dribbonstartshapeFrom.GongCopyBasicFields(shiftedleftgrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftPartiallyGrowthCurve2DRibbon(mapOrigCopy map[any]any, shiftedleftpartiallygrowthcurve2dribbonFrom *ShiftedLeftPartiallyGrowthCurve2DRibbon) (shiftedleftpartiallygrowthcurve2dribbonTo *ShiftedLeftPartiallyGrowthCurve2DRibbon) {
	var alreadyCopied bool
	shiftedleftpartiallygrowthcurve2dribbonTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedleftpartiallygrowthcurve2dribbonFrom)
	if alreadyCopied {
		return
	}
	shiftedleftpartiallygrowthcurve2dribbonFrom.GongCopyBasicFields(shiftedleftpartiallygrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, shiftedleftpartiallygrowthcurve2dribbonendshapeFrom *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) (shiftedleftpartiallygrowthcurve2dribbonendshapeTo *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) {
	var alreadyCopied bool
	shiftedleftpartiallygrowthcurve2dribbonendshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedleftpartiallygrowthcurve2dribbonendshapeFrom)
	if alreadyCopied {
		return
	}
	shiftedleftpartiallygrowthcurve2dribbonendshapeFrom.GongCopyBasicFields(shiftedleftpartiallygrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, shiftedleftpartiallygrowthcurve2dribbonstartshapeFrom *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) (shiftedleftpartiallygrowthcurve2dribbonstartshapeTo *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) {
	var alreadyCopied bool
	shiftedleftpartiallygrowthcurve2dribbonstartshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedleftpartiallygrowthcurve2dribbonstartshapeFrom)
	if alreadyCopied {
		return
	}
	shiftedleftpartiallygrowthcurve2dribbonstartshapeFrom.GongCopyBasicFields(shiftedleftpartiallygrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftStackGrowthCurveEndArcShape(mapOrigCopy map[any]any, shiftedleftstackgrowthcurveendarcshapeFrom *ShiftedLeftStackGrowthCurveEndArcShape) (shiftedleftstackgrowthcurveendarcshapeTo *ShiftedLeftStackGrowthCurveEndArcShape) {
	var alreadyCopied bool
	shiftedleftstackgrowthcurveendarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedleftstackgrowthcurveendarcshapeFrom)
	if alreadyCopied {
		return
	}
	shiftedleftstackgrowthcurveendarcshapeFrom.GongCopyBasicFields(shiftedleftstackgrowthcurveendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftStackGrowthCurveStartArcShape(mapOrigCopy map[any]any, shiftedleftstackgrowthcurvestartarcshapeFrom *ShiftedLeftStackGrowthCurveStartArcShape) (shiftedleftstackgrowthcurvestartarcshapeTo *ShiftedLeftStackGrowthCurveStartArcShape) {
	var alreadyCopied bool
	shiftedleftstackgrowthcurvestartarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedleftstackgrowthcurvestartarcshapeFrom)
	if alreadyCopied {
		return
	}
	shiftedleftstackgrowthcurvestartarcshapeFrom.GongCopyBasicFields(shiftedleftstackgrowthcurvestartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftStackNormalVector(mapOrigCopy map[any]any, shiftedleftstacknormalvectorFrom *ShiftedLeftStackNormalVector) (shiftedleftstacknormalvectorTo *ShiftedLeftStackNormalVector) {
	var alreadyCopied bool
	shiftedleftstacknormalvectorTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedleftstacknormalvectorFrom)
	if alreadyCopied {
		return
	}
	shiftedleftstacknormalvectorFrom.GongCopyBasicFields(shiftedleftstacknormalvectorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftStackOfGrowthCurve(mapOrigCopy map[any]any, shiftedleftstackofgrowthcurveFrom *ShiftedLeftStackOfGrowthCurve) (shiftedleftstackofgrowthcurveTo *ShiftedLeftStackOfGrowthCurve) {
	var alreadyCopied bool
	shiftedleftstackofgrowthcurveTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedleftstackofgrowthcurveFrom)
	if alreadyCopied {
		return
	}
	shiftedleftstackofgrowthcurveFrom.GongCopyBasicFields(shiftedleftstackofgrowthcurveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftStackOfNormalVector(mapOrigCopy map[any]any, shiftedleftstackofnormalvectorFrom *ShiftedLeftStackOfNormalVector) (shiftedleftstackofnormalvectorTo *ShiftedLeftStackOfNormalVector) {
	var alreadyCopied bool
	shiftedleftstackofnormalvectorTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedleftstackofnormalvectorFrom)
	if alreadyCopied {
		return
	}
	shiftedleftstackofnormalvectorFrom.GongCopyBasicFields(shiftedleftstackofnormalvectorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedRightGrowthCurve2DRibbon(mapOrigCopy map[any]any, shiftedrightgrowthcurve2dribbonFrom *ShiftedRightGrowthCurve2DRibbon) (shiftedrightgrowthcurve2dribbonTo *ShiftedRightGrowthCurve2DRibbon) {
	var alreadyCopied bool
	shiftedrightgrowthcurve2dribbonTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedrightgrowthcurve2dribbonFrom)
	if alreadyCopied {
		return
	}
	shiftedrightgrowthcurve2dribbonFrom.GongCopyBasicFields(shiftedrightgrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedRightGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, shiftedrightgrowthcurve2dribbonendshapeFrom *ShiftedRightGrowthCurve2DRibbonEndShape) (shiftedrightgrowthcurve2dribbonendshapeTo *ShiftedRightGrowthCurve2DRibbonEndShape) {
	var alreadyCopied bool
	shiftedrightgrowthcurve2dribbonendshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedrightgrowthcurve2dribbonendshapeFrom)
	if alreadyCopied {
		return
	}
	shiftedrightgrowthcurve2dribbonendshapeFrom.GongCopyBasicFields(shiftedrightgrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedRightGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, shiftedrightgrowthcurve2dribbonstartshapeFrom *ShiftedRightGrowthCurve2DRibbonStartShape) (shiftedrightgrowthcurve2dribbonstartshapeTo *ShiftedRightGrowthCurve2DRibbonStartShape) {
	var alreadyCopied bool
	shiftedrightgrowthcurve2dribbonstartshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shiftedrightgrowthcurve2dribbonstartshapeFrom)
	if alreadyCopied {
		return
	}
	shiftedrightgrowthcurve2dribbonstartshapeFrom.GongCopyBasicFields(shiftedrightgrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackGrowthCurve2DEndHalfwayArcShape(mapOrigCopy map[any]any, stackgrowthcurve2dendhalfwayarcshapeFrom *StackGrowthCurve2DEndHalfwayArcShape) (stackgrowthcurve2dendhalfwayarcshapeTo *StackGrowthCurve2DEndHalfwayArcShape) {
	var alreadyCopied bool
	stackgrowthcurve2dendhalfwayarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackgrowthcurve2dendhalfwayarcshapeFrom)
	if alreadyCopied {
		return
	}
	stackgrowthcurve2dendhalfwayarcshapeFrom.GongCopyBasicFields(stackgrowthcurve2dendhalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, stackgrowthcurve2dribbonendshapeFrom *StackGrowthCurve2DRibbonEndShape) (stackgrowthcurve2dribbonendshapeTo *StackGrowthCurve2DRibbonEndShape) {
	var alreadyCopied bool
	stackgrowthcurve2dribbonendshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackgrowthcurve2dribbonendshapeFrom)
	if alreadyCopied {
		return
	}
	stackgrowthcurve2dribbonendshapeFrom.GongCopyBasicFields(stackgrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, stackgrowthcurve2dribbonstartshapeFrom *StackGrowthCurve2DRibbonStartShape) (stackgrowthcurve2dribbonstartshapeTo *StackGrowthCurve2DRibbonStartShape) {
	var alreadyCopied bool
	stackgrowthcurve2dribbonstartshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackgrowthcurve2dribbonstartshapeFrom)
	if alreadyCopied {
		return
	}
	stackgrowthcurve2dribbonstartshapeFrom.GongCopyBasicFields(stackgrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackGrowthCurve2DStartHalfwayArcShape(mapOrigCopy map[any]any, stackgrowthcurve2dstarthalfwayarcshapeFrom *StackGrowthCurve2DStartHalfwayArcShape) (stackgrowthcurve2dstarthalfwayarcshapeTo *StackGrowthCurve2DStartHalfwayArcShape) {
	var alreadyCopied bool
	stackgrowthcurve2dstarthalfwayarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackgrowthcurve2dstarthalfwayarcshapeFrom)
	if alreadyCopied {
		return
	}
	stackgrowthcurve2dstarthalfwayarcshapeFrom.GongCopyBasicFields(stackgrowthcurve2dstarthalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfGrowthCurve2D(mapOrigCopy map[any]any, stackofgrowthcurve2dFrom *StackOfGrowthCurve2D) (stackofgrowthcurve2dTo *StackOfGrowthCurve2D) {
	var alreadyCopied bool
	stackofgrowthcurve2dTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackofgrowthcurve2dFrom)
	if alreadyCopied {
		return
	}
	stackofgrowthcurve2dFrom.GongCopyBasicFields(stackofgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfGrowthCurve2DByGrowthVector(mapOrigCopy map[any]any, stackofgrowthcurve2dbygrowthvectorFrom *StackOfGrowthCurve2DByGrowthVector) (stackofgrowthcurve2dbygrowthvectorTo *StackOfGrowthCurve2DByGrowthVector) {
	var alreadyCopied bool
	stackofgrowthcurve2dbygrowthvectorTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackofgrowthcurve2dbygrowthvectorFrom)
	if alreadyCopied {
		return
	}
	stackofgrowthcurve2dbygrowthvectorFrom.GongCopyBasicFields(stackofgrowthcurve2dbygrowthvectorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfGrowthCurve2DRibbon(mapOrigCopy map[any]any, stackofgrowthcurve2dribbonFrom *StackOfGrowthCurve2DRibbon) (stackofgrowthcurve2dribbonTo *StackOfGrowthCurve2DRibbon) {
	var alreadyCopied bool
	stackofgrowthcurve2dribbonTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackofgrowthcurve2dribbonFrom)
	if alreadyCopied {
		return
	}
	stackofgrowthcurve2dribbonFrom.GongCopyBasicFields(stackofgrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfPartiallyRotatedTorusShape(mapOrigCopy map[any]any, stackofpartiallyrotatedtorusshapeFrom *StackOfPartiallyRotatedTorusShape) (stackofpartiallyrotatedtorusshapeTo *StackOfPartiallyRotatedTorusShape) {
	var alreadyCopied bool
	stackofpartiallyrotatedtorusshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackofpartiallyrotatedtorusshapeFrom)
	if alreadyCopied {
		return
	}
	stackofpartiallyrotatedtorusshapeFrom.GongCopyBasicFields(stackofpartiallyrotatedtorusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfRotatedGrowthCurve2D(mapOrigCopy map[any]any, stackofrotatedgrowthcurve2dFrom *StackOfRotatedGrowthCurve2D) (stackofrotatedgrowthcurve2dTo *StackOfRotatedGrowthCurve2D) {
	var alreadyCopied bool
	stackofrotatedgrowthcurve2dTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackofrotatedgrowthcurve2dFrom)
	if alreadyCopied {
		return
	}
	stackofrotatedgrowthcurve2dFrom.GongCopyBasicFields(stackofrotatedgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfRotatedGrowthCurve2DRibbon(mapOrigCopy map[any]any, stackofrotatedgrowthcurve2dribbonFrom *StackOfRotatedGrowthCurve2DRibbon) (stackofrotatedgrowthcurve2dribbonTo *StackOfRotatedGrowthCurve2DRibbon) {
	var alreadyCopied bool
	stackofrotatedgrowthcurve2dribbonTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackofrotatedgrowthcurve2dribbonFrom)
	if alreadyCopied {
		return
	}
	stackofrotatedgrowthcurve2dribbonFrom.GongCopyBasicFields(stackofrotatedgrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfRotatedVaseTrapezeRingsShape(mapOrigCopy map[any]any, stackofrotatedvasetrapezeringsshapeFrom *StackOfRotatedVaseTrapezeRingsShape) (stackofrotatedvasetrapezeringsshapeTo *StackOfRotatedVaseTrapezeRingsShape) {
	var alreadyCopied bool
	stackofrotatedvasetrapezeringsshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackofrotatedvasetrapezeringsshapeFrom)
	if alreadyCopied {
		return
	}
	stackofrotatedvasetrapezeringsshapeFrom.GongCopyBasicFields(stackofrotatedvasetrapezeringsshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfVaseTrapezeRingsShape(mapOrigCopy map[any]any, stackofvasetrapezeringsshapeFrom *StackOfVaseTrapezeRingsShape) (stackofvasetrapezeringsshapeTo *StackOfVaseTrapezeRingsShape) {
	var alreadyCopied bool
	stackofvasetrapezeringsshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackofvasetrapezeringsshapeFrom)
	if alreadyCopied {
		return
	}
	stackofvasetrapezeringsshapeFrom.GongCopyBasicFields(stackofvasetrapezeringsshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackRotatedGrowthCurve2DEndArcShape(mapOrigCopy map[any]any, stackrotatedgrowthcurve2dendarcshapeFrom *StackRotatedGrowthCurve2DEndArcShape) (stackrotatedgrowthcurve2dendarcshapeTo *StackRotatedGrowthCurve2DEndArcShape) {
	var alreadyCopied bool
	stackrotatedgrowthcurve2dendarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackrotatedgrowthcurve2dendarcshapeFrom)
	if alreadyCopied {
		return
	}
	stackrotatedgrowthcurve2dendarcshapeFrom.GongCopyBasicFields(stackrotatedgrowthcurve2dendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackRotatedGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, stackrotatedgrowthcurve2dribbonendshapeFrom *StackRotatedGrowthCurve2DRibbonEndShape) (stackrotatedgrowthcurve2dribbonendshapeTo *StackRotatedGrowthCurve2DRibbonEndShape) {
	var alreadyCopied bool
	stackrotatedgrowthcurve2dribbonendshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackrotatedgrowthcurve2dribbonendshapeFrom)
	if alreadyCopied {
		return
	}
	stackrotatedgrowthcurve2dribbonendshapeFrom.GongCopyBasicFields(stackrotatedgrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackRotatedGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, stackrotatedgrowthcurve2dribbonstartshapeFrom *StackRotatedGrowthCurve2DRibbonStartShape) (stackrotatedgrowthcurve2dribbonstartshapeTo *StackRotatedGrowthCurve2DRibbonStartShape) {
	var alreadyCopied bool
	stackrotatedgrowthcurve2dribbonstartshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackrotatedgrowthcurve2dribbonstartshapeFrom)
	if alreadyCopied {
		return
	}
	stackrotatedgrowthcurve2dribbonstartshapeFrom.GongCopyBasicFields(stackrotatedgrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackRotatedGrowthCurve2DStartArcShape(mapOrigCopy map[any]any, stackrotatedgrowthcurve2dstartarcshapeFrom *StackRotatedGrowthCurve2DStartArcShape) (stackrotatedgrowthcurve2dstartarcshapeTo *StackRotatedGrowthCurve2DStartArcShape) {
	var alreadyCopied bool
	stackrotatedgrowthcurve2dstartarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stackrotatedgrowthcurve2dstartarcshapeFrom)
	if alreadyCopied {
		return
	}
	stackrotatedgrowthcurve2dstartarcshapeFrom.GongCopyBasicFields(stackrotatedgrowthcurve2dstartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStartArcShape(mapOrigCopy map[any]any, startarcshapeFrom *StartArcShape) (startarcshapeTo *StartArcShape) {
	var alreadyCopied bool
	startarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, startarcshapeFrom)
	if alreadyCopied {
		return
	}
	startarcshapeFrom.GongCopyBasicFields(startarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStartArcShapeGrid(mapOrigCopy map[any]any, startarcshapegridFrom *StartArcShapeGrid) (startarcshapegridTo *StartArcShapeGrid) {
	var alreadyCopied bool
	startarcshapegridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, startarcshapegridFrom)
	if alreadyCopied {
		return
	}
	startarcshapegridFrom.GongCopyBasicFields(startarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStartHalfwayArcShape(mapOrigCopy map[any]any, starthalfwayarcshapeFrom *StartHalfwayArcShape) (starthalfwayarcshapeTo *StartHalfwayArcShape) {
	var alreadyCopied bool
	starthalfwayarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, starthalfwayarcshapeFrom)
	if alreadyCopied {
		return
	}
	starthalfwayarcshapeFrom.GongCopyBasicFields(starthalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStartHalfwayArcShapeGrid(mapOrigCopy map[any]any, starthalfwayarcshapegridFrom *StartHalfwayArcShapeGrid) (starthalfwayarcshapegridTo *StartHalfwayArcShapeGrid) {
	var alreadyCopied bool
	starthalfwayarcshapegridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, starthalfwayarcshapegridFrom)
	if alreadyCopied {
		return
	}
	starthalfwayarcshapegridFrom.GongCopyBasicFields(starthalfwayarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStemCylinder3DShape(mapOrigCopy map[any]any, stemcylinder3dshapeFrom *StemCylinder3DShape) (stemcylinder3dshapeTo *StemCylinder3DShape) {
	var alreadyCopied bool
	stemcylinder3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stemcylinder3dshapeFrom)
	if alreadyCopied {
		return
	}
	stemcylinder3dshapeFrom.GongCopyBasicFields(stemcylinder3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStool2DDiagram(mapOrigCopy map[any]any, stool2ddiagramFrom *Stool2DDiagram) (stool2ddiagramTo *Stool2DDiagram) {
	var alreadyCopied bool
	stool2ddiagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stool2ddiagramFrom)
	if alreadyCopied {
		return
	}
	stool2ddiagramFrom.GongCopyBasicFields(stool2ddiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStool3DDiagram(mapOrigCopy map[any]any, stool3ddiagramFrom *Stool3DDiagram) (stool3ddiagramTo *Stool3DDiagram) {
	var alreadyCopied bool
	stool3ddiagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stool3ddiagramFrom)
	if alreadyCopied {
		return
	}
	stool3ddiagramFrom.GongCopyBasicFields(stool3ddiagramTo)

	//insertion point for the staging of instances referenced by pointers
	if stool3ddiagramFrom.SampledPoints3DShape != nil {
		stool3ddiagramTo.SampledPoints3DShape = GongCopyBranchSampledPoints3DShape(mapOrigCopy, stool3ddiagramFrom.SampledPoints3DShape)
	}
	if stool3ddiagramFrom.Rendered3DShape != nil {
		stool3ddiagramTo.Rendered3DShape = GongCopyBranchRendered3DShape(mapOrigCopy, stool3ddiagramFrom.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTiledFloor3DShape(mapOrigCopy map[any]any, tiledfloor3dshapeFrom *TiledFloor3DShape) (tiledfloor3dshapeTo *TiledFloor3DShape) {
	var alreadyCopied bool
	tiledfloor3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tiledfloor3dshapeFrom)
	if alreadyCopied {
		return
	}
	tiledfloor3dshapeFrom.GongCopyBasicFields(tiledfloor3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopCurvePlane1Shape(mapOrigCopy map[any]any, topcurveplane1shapeFrom *TopCurvePlane1Shape) (topcurveplane1shapeTo *TopCurvePlane1Shape) {
	var alreadyCopied bool
	topcurveplane1shapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topcurveplane1shapeFrom)
	if alreadyCopied {
		return
	}
	topcurveplane1shapeFrom.GongCopyBasicFields(topcurveplane1shapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopCurvePlane2Shape(mapOrigCopy map[any]any, topcurveplane2shapeFrom *TopCurvePlane2Shape) (topcurveplane2shapeTo *TopCurvePlane2Shape) {
	var alreadyCopied bool
	topcurveplane2shapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topcurveplane2shapeFrom)
	if alreadyCopied {
		return
	}
	topcurveplane2shapeFrom.GongCopyBasicFields(topcurveplane2shapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopEndArcShape(mapOrigCopy map[any]any, topendarcshapeFrom *TopEndArcShape) (topendarcshapeTo *TopEndArcShape) {
	var alreadyCopied bool
	topendarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topendarcshapeFrom)
	if alreadyCopied {
		return
	}
	topendarcshapeFrom.GongCopyBasicFields(topendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopEndArcShapeGrid(mapOrigCopy map[any]any, topendarcshapegridFrom *TopEndArcShapeGrid) (topendarcshapegridTo *TopEndArcShapeGrid) {
	var alreadyCopied bool
	topendarcshapegridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topendarcshapegridFrom)
	if alreadyCopied {
		return
	}
	topendarcshapegridFrom.GongCopyBasicFields(topendarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopEndHalfwayArcShape(mapOrigCopy map[any]any, topendhalfwayarcshapeFrom *TopEndHalfwayArcShape) (topendhalfwayarcshapeTo *TopEndHalfwayArcShape) {
	var alreadyCopied bool
	topendhalfwayarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topendhalfwayarcshapeFrom)
	if alreadyCopied {
		return
	}
	topendhalfwayarcshapeFrom.GongCopyBasicFields(topendhalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopEndHalfwayArcShapeGrid(mapOrigCopy map[any]any, topendhalfwayarcshapegridFrom *TopEndHalfwayArcShapeGrid) (topendhalfwayarcshapegridTo *TopEndHalfwayArcShapeGrid) {
	var alreadyCopied bool
	topendhalfwayarcshapegridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topendhalfwayarcshapegridFrom)
	if alreadyCopied {
		return
	}
	topendhalfwayarcshapegridFrom.GongCopyBasicFields(topendhalfwayarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopGrowthCurve2D(mapOrigCopy map[any]any, topgrowthcurve2dFrom *TopGrowthCurve2D) (topgrowthcurve2dTo *TopGrowthCurve2D) {
	var alreadyCopied bool
	topgrowthcurve2dTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topgrowthcurve2dFrom)
	if alreadyCopied {
		return
	}
	topgrowthcurve2dFrom.GongCopyBasicFields(topgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopMidArcVectorShape(mapOrigCopy map[any]any, topmidarcvectorshapeFrom *TopMidArcVectorShape) (topmidarcvectorshapeTo *TopMidArcVectorShape) {
	var alreadyCopied bool
	topmidarcvectorshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topmidarcvectorshapeFrom)
	if alreadyCopied {
		return
	}
	topmidarcvectorshapeFrom.GongCopyBasicFields(topmidarcvectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopMidArcVectorShapeGrid(mapOrigCopy map[any]any, topmidarcvectorshapegridFrom *TopMidArcVectorShapeGrid) (topmidarcvectorshapegridTo *TopMidArcVectorShapeGrid) {
	var alreadyCopied bool
	topmidarcvectorshapegridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topmidarcvectorshapegridFrom)
	if alreadyCopied {
		return
	}
	topmidarcvectorshapegridFrom.GongCopyBasicFields(topmidarcvectorshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStackGrowthCurve2DEndHalfwayArcShape(mapOrigCopy map[any]any, topstackgrowthcurve2dendhalfwayarcshapeFrom *TopStackGrowthCurve2DEndHalfwayArcShape) (topstackgrowthcurve2dendhalfwayarcshapeTo *TopStackGrowthCurve2DEndHalfwayArcShape) {
	var alreadyCopied bool
	topstackgrowthcurve2dendhalfwayarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topstackgrowthcurve2dendhalfwayarcshapeFrom)
	if alreadyCopied {
		return
	}
	topstackgrowthcurve2dendhalfwayarcshapeFrom.GongCopyBasicFields(topstackgrowthcurve2dendhalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStackGrowthCurve2DStartHalfwayArcShape(mapOrigCopy map[any]any, topstackgrowthcurve2dstarthalfwayarcshapeFrom *TopStackGrowthCurve2DStartHalfwayArcShape) (topstackgrowthcurve2dstarthalfwayarcshapeTo *TopStackGrowthCurve2DStartHalfwayArcShape) {
	var alreadyCopied bool
	topstackgrowthcurve2dstarthalfwayarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topstackgrowthcurve2dstarthalfwayarcshapeFrom)
	if alreadyCopied {
		return
	}
	topstackgrowthcurve2dstarthalfwayarcshapeFrom.GongCopyBasicFields(topstackgrowthcurve2dstarthalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStackOfGrowthCurve2D(mapOrigCopy map[any]any, topstackofgrowthcurve2dFrom *TopStackOfGrowthCurve2D) (topstackofgrowthcurve2dTo *TopStackOfGrowthCurve2D) {
	var alreadyCopied bool
	topstackofgrowthcurve2dTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topstackofgrowthcurve2dFrom)
	if alreadyCopied {
		return
	}
	topstackofgrowthcurve2dFrom.GongCopyBasicFields(topstackofgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStackOfRotatedGrowthCurve2D(mapOrigCopy map[any]any, topstackofrotatedgrowthcurve2dFrom *TopStackOfRotatedGrowthCurve2D) (topstackofrotatedgrowthcurve2dTo *TopStackOfRotatedGrowthCurve2D) {
	var alreadyCopied bool
	topstackofrotatedgrowthcurve2dTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topstackofrotatedgrowthcurve2dFrom)
	if alreadyCopied {
		return
	}
	topstackofrotatedgrowthcurve2dFrom.GongCopyBasicFields(topstackofrotatedgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStackOfRotatedGrowthCurve2DEndArcShape(mapOrigCopy map[any]any, topstackofrotatedgrowthcurve2dendarcshapeFrom *TopStackOfRotatedGrowthCurve2DEndArcShape) (topstackofrotatedgrowthcurve2dendarcshapeTo *TopStackOfRotatedGrowthCurve2DEndArcShape) {
	var alreadyCopied bool
	topstackofrotatedgrowthcurve2dendarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topstackofrotatedgrowthcurve2dendarcshapeFrom)
	if alreadyCopied {
		return
	}
	topstackofrotatedgrowthcurve2dendarcshapeFrom.GongCopyBasicFields(topstackofrotatedgrowthcurve2dendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStackOfRotatedGrowthCurve2DStartArcShape(mapOrigCopy map[any]any, topstackofrotatedgrowthcurve2dstartarcshapeFrom *TopStackOfRotatedGrowthCurve2DStartArcShape) (topstackofrotatedgrowthcurve2dstartarcshapeTo *TopStackOfRotatedGrowthCurve2DStartArcShape) {
	var alreadyCopied bool
	topstackofrotatedgrowthcurve2dstartarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topstackofrotatedgrowthcurve2dstartarcshapeFrom)
	if alreadyCopied {
		return
	}
	topstackofrotatedgrowthcurve2dstartarcshapeFrom.GongCopyBasicFields(topstackofrotatedgrowthcurve2dstartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStartArcShape(mapOrigCopy map[any]any, topstartarcshapeFrom *TopStartArcShape) (topstartarcshapeTo *TopStartArcShape) {
	var alreadyCopied bool
	topstartarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topstartarcshapeFrom)
	if alreadyCopied {
		return
	}
	topstartarcshapeFrom.GongCopyBasicFields(topstartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStartArcShapeGrid(mapOrigCopy map[any]any, topstartarcshapegridFrom *TopStartArcShapeGrid) (topstartarcshapegridTo *TopStartArcShapeGrid) {
	var alreadyCopied bool
	topstartarcshapegridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topstartarcshapegridFrom)
	if alreadyCopied {
		return
	}
	topstartarcshapegridFrom.GongCopyBasicFields(topstartarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStartHalfwayArcShape(mapOrigCopy map[any]any, topstarthalfwayarcshapeFrom *TopStartHalfwayArcShape) (topstarthalfwayarcshapeTo *TopStartHalfwayArcShape) {
	var alreadyCopied bool
	topstarthalfwayarcshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topstarthalfwayarcshapeFrom)
	if alreadyCopied {
		return
	}
	topstarthalfwayarcshapeFrom.GongCopyBasicFields(topstarthalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStartHalfwayArcShapeGrid(mapOrigCopy map[any]any, topstarthalfwayarcshapegridFrom *TopStartHalfwayArcShapeGrid) (topstarthalfwayarcshapegridTo *TopStartHalfwayArcShapeGrid) {
	var alreadyCopied bool
	topstarthalfwayarcshapegridTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, topstarthalfwayarcshapegridFrom)
	if alreadyCopied {
		return
	}
	topstarthalfwayarcshapegridFrom.GongCopyBasicFields(topstarthalfwayarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTorus3DShape(mapOrigCopy map[any]any, torus3dshapeFrom *Torus3DShape) (torus3dshapeTo *Torus3DShape) {
	var alreadyCopied bool
	torus3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, torus3dshapeFrom)
	if alreadyCopied {
		return
	}
	torus3dshapeFrom.GongCopyBasicFields(torus3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTorusEdge3DShape(mapOrigCopy map[any]any, torusedge3dshapeFrom *TorusEdge3DShape) (torusedge3dshapeTo *TorusEdge3DShape) {
	var alreadyCopied bool
	torusedge3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, torusedge3dshapeFrom)
	if alreadyCopied {
		return
	}
	torusedge3dshapeFrom.GongCopyBasicFields(torusedge3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTorusStackShape(mapOrigCopy map[any]any, torusstackshapeFrom *TorusStackShape) (torusstackshapeTo *TorusStackShape) {
	var alreadyCopied bool
	torusstackshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, torusstackshapeFrom)
	if alreadyCopied {
		return
	}
	torusstackshapeFrom.GongCopyBasicFields(torusstackshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTubeVase3DDiagram(mapOrigCopy map[any]any, tubevase3ddiagramFrom *TubeVase3DDiagram) (tubevase3ddiagramTo *TubeVase3DDiagram) {
	var alreadyCopied bool
	tubevase3ddiagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tubevase3ddiagramFrom)
	if alreadyCopied {
		return
	}
	tubevase3ddiagramFrom.GongCopyBasicFields(tubevase3ddiagramTo)

	//insertion point for the staging of instances referenced by pointers
	if tubevase3ddiagramFrom.Rendered3DShape != nil {
		tubevase3ddiagramTo.Rendered3DShape = GongCopyBranchRendered3DShape(mapOrigCopy, tubevase3ddiagramFrom.Rendered3DShape)
	}
	if tubevase3ddiagramFrom.SampledPoints3DShape != nil {
		tubevase3ddiagramTo.SampledPoints3DShape = GongCopyBranchSampledPoints3DShape(mapOrigCopy, tubevase3ddiagramFrom.SampledPoints3DShape)
	}
	if tubevase3ddiagramFrom.OriginalPoints3DShape != nil {
		tubevase3ddiagramTo.OriginalPoints3DShape = GongCopyBranchOriginalPoints3DShape(mapOrigCopy, tubevase3ddiagramFrom.OriginalPoints3DShape)
	}
	if tubevase3ddiagramFrom.Angle0Shape != nil {
		tubevase3ddiagramTo.Angle0Shape = GongCopyBranchAngle0Shape(mapOrigCopy, tubevase3ddiagramFrom.Angle0Shape)
	}
	if tubevase3ddiagramFrom.TopCurvePlane1Shape != nil {
		tubevase3ddiagramTo.TopCurvePlane1Shape = GongCopyBranchTopCurvePlane1Shape(mapOrigCopy, tubevase3ddiagramFrom.TopCurvePlane1Shape)
	}
	if tubevase3ddiagramFrom.BottomCurvePlane1Shape != nil {
		tubevase3ddiagramTo.BottomCurvePlane1Shape = GongCopyBranchBottomCurvePlane1Shape(mapOrigCopy, tubevase3ddiagramFrom.BottomCurvePlane1Shape)
	}
	if tubevase3ddiagramFrom.TopCurvePlane2Shape != nil {
		tubevase3ddiagramTo.TopCurvePlane2Shape = GongCopyBranchTopCurvePlane2Shape(mapOrigCopy, tubevase3ddiagramFrom.TopCurvePlane2Shape)
	}
	if tubevase3ddiagramFrom.BottomCurvePlane2Shape != nil {
		tubevase3ddiagramTo.BottomCurvePlane2Shape = GongCopyBranchBottomCurvePlane2Shape(mapOrigCopy, tubevase3ddiagramFrom.BottomCurvePlane2Shape)
	}
	if tubevase3ddiagramFrom.VaseTrapezeRingShape != nil {
		tubevase3ddiagramTo.VaseTrapezeRingShape = GongCopyBranchVaseTrapezeRingShape(mapOrigCopy, tubevase3ddiagramFrom.VaseTrapezeRingShape)
	}
	if tubevase3ddiagramFrom.StackOfVaseTrapezeRingsShape != nil {
		tubevase3ddiagramTo.StackOfVaseTrapezeRingsShape = GongCopyBranchStackOfVaseTrapezeRingsShape(mapOrigCopy, tubevase3ddiagramFrom.StackOfVaseTrapezeRingsShape)
	}
	if tubevase3ddiagramFrom.StackOfRotatedVaseTrapezeRingsShape != nil {
		tubevase3ddiagramTo.StackOfRotatedVaseTrapezeRingsShape = GongCopyBranchStackOfRotatedVaseTrapezeRingsShape(mapOrigCopy, tubevase3ddiagramFrom.StackOfRotatedVaseTrapezeRingsShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTubeVaseAbstract(mapOrigCopy map[any]any, tubevaseabstractFrom *TubeVaseAbstract) (tubevaseabstractTo *TubeVaseAbstract) {
	var alreadyCopied bool
	tubevaseabstractTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tubevaseabstractFrom)
	if alreadyCopied {
		return
	}
	tubevaseabstractFrom.GongCopyBasicFields(tubevaseabstractTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVase2DDiagram(mapOrigCopy map[any]any, vase2ddiagramFrom *Vase2DDiagram) (vase2ddiagramTo *Vase2DDiagram) {
	var alreadyCopied bool
	vase2ddiagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, vase2ddiagramFrom)
	if alreadyCopied {
		return
	}
	vase2ddiagramFrom.GongCopyBasicFields(vase2ddiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVaseTrapezeRingShape(mapOrigCopy map[any]any, vasetrapezeringshapeFrom *VaseTrapezeRingShape) (vasetrapezeringshapeTo *VaseTrapezeRingShape) {
	var alreadyCopied bool
	vasetrapezeringshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, vasetrapezeringshapeFrom)
	if alreadyCopied {
		return
	}
	vasetrapezeringshapeFrom.GongCopyBasicFields(vasetrapezeringshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVerticalTorusStackShape(mapOrigCopy map[any]any, verticaltorusstackshapeFrom *VerticalTorusStackShape) (verticaltorusstackshapeTo *VerticalTorusStackShape) {
	var alreadyCopied bool
	verticaltorusstackshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, verticaltorusstackshapeFrom)
	if alreadyCopied {
		return
	}
	verticaltorusstackshapeFrom.GongCopyBasicFields(verticaltorusstackshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVolumeKey3DShape(mapOrigCopy map[any]any, volumekey3dshapeFrom *VolumeKey3DShape) (volumekey3dshapeTo *VolumeKey3DShape) {
	var alreadyCopied bool
	volumekey3dshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, volumekey3dshapeFrom)
	if alreadyCopied {
		return
	}
	volumekey3dshapeFrom.GongCopyBasicFields(volumekey3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (angle0shape *Angle0Shape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(angle0shape) {
		return
	}

	angle0shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arcnormalvectorshape *ArcNormalVectorShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(arcnormalvectorshape) {
		return
	}

	arcnormalvectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(arcnormalvectorshapegrid) {
		return
	}

	arcnormalvectorshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (axesshape *AxesShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(axesshape) {
		return
	}

	axesshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (basevectorshape *BaseVectorShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(basevectorshape) {
		return
	}

	basevectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (basevectorshapegrid *BaseVectorShapeGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(basevectorshapegrid) {
		return
	}

	basevectorshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bottomcurveplane1shape) {
		return
	}

	bottomcurveplane1shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bottomcurveplane2shape) {
		return
	}

	bottomcurveplane2shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(chosenp1p2pairshape) {
		return
	}

	chosenp1p2pairshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circlegridshape *CircleGridShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(circlegridshape) {
		return
	}

	circlegridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circumference3dshape *Circumference3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(circumference3dshape) {
		return
	}

	circumference3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clock2ddiagram *Clock2DDiagram) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(clock2ddiagram) {
		return
	}

	clock2ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clock3ddiagram *Clock3DDiagram) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(clock3ddiagram) {
		return
	}

	clock3ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if clock3ddiagram.SampledPoints3DShape != nil {
		stage.UnstageBranch(clock3ddiagram.SampledPoints3DShape)
	}
	if clock3ddiagram.Rendered3DShape != nil {
		stage.UnstageBranch(clock3ddiagram.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clocktopcurveshape *ClockTopCurveShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(clocktopcurveshape) {
		return
	}

	clocktopcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cutline3dshape *CutLine3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(cutline3dshape) {
		return
	}

	cutline3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endarcshape *EndArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(endarcshape) {
		return
	}

	endarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endarcshapegrid *EndArcShapeGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(endarcshapegrid) {
		return
	}

	endarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endhalfwayarcshape *EndHalfwayArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(endhalfwayarcshape) {
		return
	}

	endhalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(endhalfwayarcshapegrid) {
		return
	}

	endhalfwayarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (explanationtextshape *ExplanationTextShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(explanationtextshape) {
		return
	}

	explanationtextshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eye3dshape *Eye3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(eye3dshape) {
		return
	}

	eye3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(eyecornerssampledpoints3dshape) {
		return
	}

	eyecornerssampledpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(eyesampledpoints3dshape) {
		return
	}

	eyesampledpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(eyeseatbottomcurveshape) {
		return
	}

	eyeseatbottomcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(eyestoolbottomcurveshape) {
		return
	}

	eyestoolbottomcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyevolume3dshape *EyeVolume3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(eyevolume3dshape) {
		return
	}

	eyevolume3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gridpathshape *GridPathShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(gridpathshape) {
		return
	}

	gridpathshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2d *GrowthCurve2D) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(growthcurve2d) {
		return
	}

	growthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(growthcurve2dribbon) {
		return
	}

	growthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(growthcurve2dribbonendshape) {
		return
	}

	growthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(growthcurve2dribbonstartshape) {
		return
	}

	growthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(growthcurverhombusgridshape) {
		return
	}

	growthcurverhombusgridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(growthcurverhombusshape) {
		return
	}

	growthcurverhombusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthvectorshape *GrowthVectorShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(growthvectorshape) {
		return
	}

	growthvectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (initialrhombusgridshape *InitialRhombusGridShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(initialrhombusgridshape) {
		return
	}

	initialrhombusgridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (initialrhombusshape *InitialRhombusShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(initialrhombusshape) {
		return
	}

	initialrhombusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (key3dshape *Key3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(key3dshape) {
		return
	}

	key3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (keyhole3dshape *KeyHole3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(keyhole3dshape) {
		return
	}

	keyhole3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (keyholeshape *KeyHoleShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(keyholeshape) {
		return
	}

	keyholeshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (leaves3dshape *Leaves3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(leaves3dshape) {
		return
	}

	leaves3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plantabstract := range library.Plants {
		stage.UnstageBranch(_plantabstract)
	}
	for _, _library := range library.SubLibraries {
		stage.UnstageBranch(_library)
	}

}

func (midarcvectorshape *MidArcVectorShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(midarcvectorshape) {
		return
	}

	midarcvectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(midarcvectorshapegrid) {
		return
	}

	midarcvectorshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (originalpoints3dshape *OriginalPoints3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(originalpoints3dshape) {
		return
	}

	originalpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(parastichymcurves3dshape) {
		return
	}

	parastichymcurves3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(parastichyncurves3dshape) {
		return
	}

	parastichyncurves3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dribbon) {
		return
	}

	partiallygrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dribbonendshape) {
		return
	}

	partiallygrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dribbonstartshape) {
		return
	}

	partiallygrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectory) {
		return
	}

	partiallygrowthcurve2dtrajectory.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryp1curveshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1curveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryp1p2) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryp1p2pairlineshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2pairlineshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryp1pointshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1pointshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryp2curveshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp2curveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryp2pointshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp2pointshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallyrotatedseatbottomcurveshape) {
		return
	}

	partiallyrotatedseatbottomcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallyrotatedseattopcurveshape) {
		return
	}

	partiallyrotatedseattopcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partiallyrotatedtorusshape) {
		return
	}

	partiallyrotatedtorusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvector *PerpendicularVector) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(perpendicularvector) {
		return
	}

	perpendicularvector.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(perpendicularvectorgrid) {
		return
	}

	perpendicularvectorgrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(perpendicularvectorgridhalfway) {
		return
	}

	perpendicularvectorgridhalfway.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(perpendicularvectorhalfway) {
		return
	}

	perpendicularvectorhalfway.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (plant2ddiagram *Plant2DDiagram) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(plant2ddiagram) {
		return
	}

	plant2ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (plant3ddiagram *Plant3DDiagram) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(plant3ddiagram) {
		return
	}

	plant3ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plant3ddiagram.StemCylinder3DShape != nil {
		stage.UnstageBranch(plant3ddiagram.StemCylinder3DShape)
	}
	if plant3ddiagram.ParastichyNCurves3DShape != nil {
		stage.UnstageBranch(plant3ddiagram.ParastichyNCurves3DShape)
	}
	if plant3ddiagram.ParastichyMCurves3DShape != nil {
		stage.UnstageBranch(plant3ddiagram.ParastichyMCurves3DShape)
	}
	if plant3ddiagram.CutLine3DShape != nil {
		stage.UnstageBranch(plant3ddiagram.CutLine3DShape)
	}
	if plant3ddiagram.Circumference3DShape != nil {
		stage.UnstageBranch(plant3ddiagram.Circumference3DShape)
	}
	if plant3ddiagram.Leaves3DShape != nil {
		stage.UnstageBranch(plant3ddiagram.Leaves3DShape)
	}
	if plant3ddiagram.Rendered3DShape != nil {
		stage.UnstageBranch(plant3ddiagram.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (plantabstract *PlantAbstract) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(plantabstract) {
		return
	}

	plantabstract.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plantabstract.TubeVaseAbstract != nil {
		stage.UnstageBranch(plantabstract.TubeVaseAbstract)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plant2ddiagram := range plantabstract.Plant2DDiagrams {
		stage.UnstageBranch(_plant2ddiagram)
	}
	for _, _plant3ddiagram := range plantabstract.Plant3DDiagrams {
		stage.UnstageBranch(_plant3ddiagram)
	}
	for _, _vase2ddiagram := range plantabstract.Vase2DDiagrams {
		stage.UnstageBranch(_vase2ddiagram)
	}
	for _, _tubevase3ddiagram := range plantabstract.TubeVase3DDiagrams {
		stage.UnstageBranch(_tubevase3ddiagram)
	}
	for _, _stool2ddiagram := range plantabstract.Stool2DDiagrams {
		stage.UnstageBranch(_stool2ddiagram)
	}
	for _, _stool3ddiagram := range plantabstract.Stool3DDiagrams {
		stage.UnstageBranch(_stool3ddiagram)
	}
	for _, _clock2ddiagram := range plantabstract.Clock2DDiagrams {
		stage.UnstageBranch(_clock2ddiagram)
	}
	for _, _clock3ddiagram := range plantabstract.Clock3DDiagrams {
		stage.UnstageBranch(_clock3ddiagram)
	}

}

func (plantcircumferenceshape *PlantCircumferenceShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(plantcircumferenceshape) {
		return
	}

	plantcircumferenceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointsandlines3dshape *PointsAndLines3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pointsandlines3dshape) {
		return
	}

	pointsandlines3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pxshape *PxShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pxshape) {
		return
	}

	pxshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rendered3dshape *Rendered3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rendered3dshape) {
		return
	}

	rendered3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rhombusshape *RhombusShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rhombusshape) {
		return
	}

	rhombusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rhombusstuff *RhombusStuff) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rhombusstuff) {
		return
	}

	rhombusstuff.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rotatedrhombusgridshape) {
		return
	}

	rotatedrhombusgridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedrhombusshape *RotatedRhombusShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rotatedrhombusshape) {
		return
	}

	rotatedrhombusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rotatedsampledpoints3dshape) {
		return
	}

	rotatedsampledpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rotatedseatandlegs3dshape) {
		return
	}

	rotatedseatandlegs3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (sampledpoints3dshape *SampledPoints3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(sampledpoints3dshape) {
		return
	}

	sampledpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seat3dshape *Seat3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(seat3dshape) {
		return
	}

	seat3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(seatandlegs3dshape) {
		return
	}

	seatandlegs3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seatbottomcurveshape *SeatBottomCurveShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(seatbottomcurveshape) {
		return
	}

	seatbottomcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seattopcurveshape *SeatTopCurveShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(seattopcurveshape) {
		return
	}

	seattopcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedbottomtopstartarcshape) {
		return
	}

	shiftedbottomtopstartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedbottomtopstartarcshapegrid) {
		return
	}

	shiftedbottomtopstartarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftgrowthcurve2dribbon) {
		return
	}

	shiftedleftgrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftgrowthcurve2dribbonendshape) {
		return
	}

	shiftedleftgrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftgrowthcurve2dribbonstartshape) {
		return
	}

	shiftedleftgrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftpartiallygrowthcurve2dribbon) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftpartiallygrowthcurve2dribbonendshape) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftpartiallygrowthcurve2dribbonstartshape) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftstackgrowthcurveendarcshape) {
		return
	}

	shiftedleftstackgrowthcurveendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftstackgrowthcurvestartarcshape) {
		return
	}

	shiftedleftstackgrowthcurvestartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftstacknormalvector) {
		return
	}

	shiftedleftstacknormalvector.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftstackofgrowthcurve) {
		return
	}

	shiftedleftstackofgrowthcurve.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftstackofnormalvector) {
		return
	}

	shiftedleftstackofnormalvector.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedrightgrowthcurve2dribbon) {
		return
	}

	shiftedrightgrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedrightgrowthcurve2dribbonendshape) {
		return
	}

	shiftedrightgrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedrightgrowthcurve2dribbonstartshape) {
		return
	}

	shiftedrightgrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackgrowthcurve2dendhalfwayarcshape) {
		return
	}

	stackgrowthcurve2dendhalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackgrowthcurve2dribbonendshape) {
		return
	}

	stackgrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackgrowthcurve2dribbonstartshape) {
		return
	}

	stackgrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackgrowthcurve2dstarthalfwayarcshape) {
		return
	}

	stackgrowthcurve2dstarthalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackofgrowthcurve2d) {
		return
	}

	stackofgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackofgrowthcurve2dbygrowthvector) {
		return
	}

	stackofgrowthcurve2dbygrowthvector.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackofgrowthcurve2dribbon) {
		return
	}

	stackofgrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackofpartiallyrotatedtorusshape) {
		return
	}

	stackofpartiallyrotatedtorusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackofrotatedgrowthcurve2d) {
		return
	}

	stackofrotatedgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackofrotatedgrowthcurve2dribbon) {
		return
	}

	stackofrotatedgrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackofrotatedvasetrapezeringsshape) {
		return
	}

	stackofrotatedvasetrapezeringsshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackofvasetrapezeringsshape) {
		return
	}

	stackofvasetrapezeringsshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackrotatedgrowthcurve2dendarcshape) {
		return
	}

	stackrotatedgrowthcurve2dendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackrotatedgrowthcurve2dribbonendshape) {
		return
	}

	stackrotatedgrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackrotatedgrowthcurve2dribbonstartshape) {
		return
	}

	stackrotatedgrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stackrotatedgrowthcurve2dstartarcshape) {
		return
	}

	stackrotatedgrowthcurve2dstartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (startarcshape *StartArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(startarcshape) {
		return
	}

	startarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (startarcshapegrid *StartArcShapeGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(startarcshapegrid) {
		return
	}

	startarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (starthalfwayarcshape *StartHalfwayArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(starthalfwayarcshape) {
		return
	}

	starthalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(starthalfwayarcshapegrid) {
		return
	}

	starthalfwayarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stemcylinder3dshape *StemCylinder3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stemcylinder3dshape) {
		return
	}

	stemcylinder3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stool2ddiagram *Stool2DDiagram) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stool2ddiagram) {
		return
	}

	stool2ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stool3ddiagram *Stool3DDiagram) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stool3ddiagram) {
		return
	}

	stool3ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stool3ddiagram.SampledPoints3DShape != nil {
		stage.UnstageBranch(stool3ddiagram.SampledPoints3DShape)
	}
	if stool3ddiagram.Rendered3DShape != nil {
		stage.UnstageBranch(stool3ddiagram.Rendered3DShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tiledfloor3dshape *TiledFloor3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tiledfloor3dshape) {
		return
	}

	tiledfloor3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topcurveplane1shape *TopCurvePlane1Shape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topcurveplane1shape) {
		return
	}

	topcurveplane1shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topcurveplane2shape *TopCurvePlane2Shape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topcurveplane2shape) {
		return
	}

	topcurveplane2shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendarcshape *TopEndArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topendarcshape) {
		return
	}

	topendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendarcshapegrid *TopEndArcShapeGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topendarcshapegrid) {
		return
	}

	topendarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topendhalfwayarcshape) {
		return
	}

	topendhalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topendhalfwayarcshapegrid) {
		return
	}

	topendhalfwayarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topgrowthcurve2d *TopGrowthCurve2D) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topgrowthcurve2d) {
		return
	}

	topgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topmidarcvectorshape *TopMidArcVectorShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topmidarcvectorshape) {
		return
	}

	topmidarcvectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topmidarcvectorshapegrid) {
		return
	}

	topmidarcvectorshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topstackgrowthcurve2dendhalfwayarcshape) {
		return
	}

	topstackgrowthcurve2dendhalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topstackgrowthcurve2dstarthalfwayarcshape) {
		return
	}

	topstackgrowthcurve2dstarthalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topstackofgrowthcurve2d) {
		return
	}

	topstackofgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topstackofrotatedgrowthcurve2d) {
		return
	}

	topstackofrotatedgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topstackofrotatedgrowthcurve2dendarcshape) {
		return
	}

	topstackofrotatedgrowthcurve2dendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topstackofrotatedgrowthcurve2dstartarcshape) {
		return
	}

	topstackofrotatedgrowthcurve2dstartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstartarcshape *TopStartArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topstartarcshape) {
		return
	}

	topstartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topstartarcshapegrid) {
		return
	}

	topstartarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topstarthalfwayarcshape) {
		return
	}

	topstarthalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(topstarthalfwayarcshapegrid) {
		return
	}

	topstarthalfwayarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torus3dshape *Torus3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(torus3dshape) {
		return
	}

	torus3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torusedge3dshape *TorusEdge3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(torusedge3dshape) {
		return
	}

	torusedge3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torusstackshape *TorusStackShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(torusstackshape) {
		return
	}

	torusstackshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tubevase3ddiagram *TubeVase3DDiagram) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tubevase3ddiagram) {
		return
	}

	tubevase3ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tubevase3ddiagram.Rendered3DShape != nil {
		stage.UnstageBranch(tubevase3ddiagram.Rendered3DShape)
	}
	if tubevase3ddiagram.SampledPoints3DShape != nil {
		stage.UnstageBranch(tubevase3ddiagram.SampledPoints3DShape)
	}
	if tubevase3ddiagram.OriginalPoints3DShape != nil {
		stage.UnstageBranch(tubevase3ddiagram.OriginalPoints3DShape)
	}
	if tubevase3ddiagram.Angle0Shape != nil {
		stage.UnstageBranch(tubevase3ddiagram.Angle0Shape)
	}
	if tubevase3ddiagram.TopCurvePlane1Shape != nil {
		stage.UnstageBranch(tubevase3ddiagram.TopCurvePlane1Shape)
	}
	if tubevase3ddiagram.BottomCurvePlane1Shape != nil {
		stage.UnstageBranch(tubevase3ddiagram.BottomCurvePlane1Shape)
	}
	if tubevase3ddiagram.TopCurvePlane2Shape != nil {
		stage.UnstageBranch(tubevase3ddiagram.TopCurvePlane2Shape)
	}
	if tubevase3ddiagram.BottomCurvePlane2Shape != nil {
		stage.UnstageBranch(tubevase3ddiagram.BottomCurvePlane2Shape)
	}
	if tubevase3ddiagram.VaseTrapezeRingShape != nil {
		stage.UnstageBranch(tubevase3ddiagram.VaseTrapezeRingShape)
	}
	if tubevase3ddiagram.StackOfVaseTrapezeRingsShape != nil {
		stage.UnstageBranch(tubevase3ddiagram.StackOfVaseTrapezeRingsShape)
	}
	if tubevase3ddiagram.StackOfRotatedVaseTrapezeRingsShape != nil {
		stage.UnstageBranch(tubevase3ddiagram.StackOfRotatedVaseTrapezeRingsShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tubevaseabstract *TubeVaseAbstract) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tubevaseabstract) {
		return
	}

	tubevaseabstract.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vase2ddiagram *Vase2DDiagram) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(vase2ddiagram) {
		return
	}

	vase2ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(vasetrapezeringshape) {
		return
	}

	vasetrapezeringshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticaltorusstackshape *VerticalTorusStackShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(verticaltorusstackshape) {
		return
	}

	verticaltorusstackshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (volumekey3dshape *VolumeKey3DShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(volumekey3dshape) {
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

func (reference *BottomCurvePlane1Shape) GongReconstructPointersFromReferences(stage *Stage, instance *BottomCurvePlane1Shape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *BottomCurvePlane2Shape) GongReconstructPointersFromReferences(stage *Stage, instance *BottomCurvePlane2Shape) {
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
	__gong__reconstructPointer(&reference.SampledPoints3DShape, stage.SampledPoints3DShapes_reference, instance.SampledPoints3DShape)
	__gong__reconstructPointer(&reference.Rendered3DShape, stage.Rendered3DShapes_reference, instance.Rendered3DShape)
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

func (reference *Leaves3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *Leaves3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Plants, stage.PlantAbstracts_reference, instance.Plants)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibraries, stage.Librarys_reference, instance.SubLibraries)
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
	__gong__reconstructPointer(&reference.StemCylinder3DShape, stage.StemCylinder3DShapes_reference, instance.StemCylinder3DShape)
	__gong__reconstructPointer(&reference.ParastichyNCurves3DShape, stage.ParastichyNCurves3DShapes_reference, instance.ParastichyNCurves3DShape)
	__gong__reconstructPointer(&reference.ParastichyMCurves3DShape, stage.ParastichyMCurves3DShapes_reference, instance.ParastichyMCurves3DShape)
	__gong__reconstructPointer(&reference.CutLine3DShape, stage.CutLine3DShapes_reference, instance.CutLine3DShape)
	__gong__reconstructPointer(&reference.Circumference3DShape, stage.Circumference3DShapes_reference, instance.Circumference3DShape)
	__gong__reconstructPointer(&reference.Leaves3DShape, stage.Leaves3DShapes_reference, instance.Leaves3DShape)
	__gong__reconstructPointer(&reference.Rendered3DShape, stage.Rendered3DShapes_reference, instance.Rendered3DShape)
	// insertion point for slice of pointers field
}

func (reference *PlantAbstract) GongReconstructPointersFromReferences(stage *Stage, instance *PlantAbstract) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.TubeVaseAbstract, stage.TubeVaseAbstracts_reference, instance.TubeVaseAbstract)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Plant2DDiagrams, stage.Plant2DDiagrams_reference, instance.Plant2DDiagrams)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Plant3DDiagrams, stage.Plant3DDiagrams_reference, instance.Plant3DDiagrams)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Vase2DDiagrams, stage.Vase2DDiagrams_reference, instance.Vase2DDiagrams)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TubeVase3DDiagrams, stage.TubeVase3DDiagrams_reference, instance.TubeVase3DDiagrams)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Stool2DDiagrams, stage.Stool2DDiagrams_reference, instance.Stool2DDiagrams)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Stool3DDiagrams, stage.Stool3DDiagrams_reference, instance.Stool3DDiagrams)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Clock2DDiagrams, stage.Clock2DDiagrams_reference, instance.Clock2DDiagrams)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Clock3DDiagrams, stage.Clock3DDiagrams_reference, instance.Clock3DDiagrams)
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

func (reference *StackOfRotatedVaseTrapezeRingsShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfRotatedVaseTrapezeRingsShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackOfVaseTrapezeRingsShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfVaseTrapezeRingsShape) {
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
	__gong__reconstructPointer(&reference.SampledPoints3DShape, stage.SampledPoints3DShapes_reference, instance.SampledPoints3DShape)
	__gong__reconstructPointer(&reference.Rendered3DShape, stage.Rendered3DShapes_reference, instance.Rendered3DShape)
	// insertion point for slice of pointers field
}

func (reference *TiledFloor3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *TiledFloor3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopCurvePlane1Shape) GongReconstructPointersFromReferences(stage *Stage, instance *TopCurvePlane1Shape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopCurvePlane2Shape) GongReconstructPointersFromReferences(stage *Stage, instance *TopCurvePlane2Shape) {
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

func (reference *TubeVase3DDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *TubeVase3DDiagram) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Rendered3DShape, stage.Rendered3DShapes_reference, instance.Rendered3DShape)
	__gong__reconstructPointer(&reference.SampledPoints3DShape, stage.SampledPoints3DShapes_reference, instance.SampledPoints3DShape)
	__gong__reconstructPointer(&reference.OriginalPoints3DShape, stage.OriginalPoints3DShapes_reference, instance.OriginalPoints3DShape)
	__gong__reconstructPointer(&reference.Angle0Shape, stage.Angle0Shapes_reference, instance.Angle0Shape)
	__gong__reconstructPointer(&reference.TopCurvePlane1Shape, stage.TopCurvePlane1Shapes_reference, instance.TopCurvePlane1Shape)
	__gong__reconstructPointer(&reference.BottomCurvePlane1Shape, stage.BottomCurvePlane1Shapes_reference, instance.BottomCurvePlane1Shape)
	__gong__reconstructPointer(&reference.TopCurvePlane2Shape, stage.TopCurvePlane2Shapes_reference, instance.TopCurvePlane2Shape)
	__gong__reconstructPointer(&reference.BottomCurvePlane2Shape, stage.BottomCurvePlane2Shapes_reference, instance.BottomCurvePlane2Shape)
	__gong__reconstructPointer(&reference.VaseTrapezeRingShape, stage.VaseTrapezeRingShapes_reference, instance.VaseTrapezeRingShape)
	__gong__reconstructPointer(&reference.StackOfVaseTrapezeRingsShape, stage.StackOfVaseTrapezeRingsShapes_reference, instance.StackOfVaseTrapezeRingsShape)
	__gong__reconstructPointer(&reference.StackOfRotatedVaseTrapezeRingsShape, stage.StackOfRotatedVaseTrapezeRingsShapes_reference, instance.StackOfRotatedVaseTrapezeRingsShape)
	// insertion point for slice of pointers field
}

func (reference *TubeVaseAbstract) GongReconstructPointersFromReferences(stage *Stage, instance *TubeVaseAbstract) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Vase2DDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *Vase2DDiagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *VaseTrapezeRingShape) GongReconstructPointersFromReferences(stage *Stage, instance *VaseTrapezeRingShape) {
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

func (reference *BottomCurvePlane1Shape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *BottomCurvePlane2Shape) GongReconstructPointersFromInstances(stage *Stage) {
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
	__gong__reconstructPointerFromInstance(&reference.SampledPoints3DShape, stage.SampledPoints3DShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.Rendered3DShape, stage.Rendered3DShapes_instance)
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

func (reference *Leaves3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Plants, stage.PlantAbstracts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibraries, stage.Librarys_instance)
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
	__gong__reconstructPointerFromInstance(&reference.StemCylinder3DShape, stage.StemCylinder3DShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.ParastichyNCurves3DShape, stage.ParastichyNCurves3DShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.ParastichyMCurves3DShape, stage.ParastichyMCurves3DShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.CutLine3DShape, stage.CutLine3DShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.Circumference3DShape, stage.Circumference3DShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.Leaves3DShape, stage.Leaves3DShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.Rendered3DShape, stage.Rendered3DShapes_instance)
	// insertion point for slice of pointers fields
}

func (reference *PlantAbstract) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.TubeVaseAbstract, stage.TubeVaseAbstracts_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Plant2DDiagrams, stage.Plant2DDiagrams_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Plant3DDiagrams, stage.Plant3DDiagrams_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Vase2DDiagrams, stage.Vase2DDiagrams_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TubeVase3DDiagrams, stage.TubeVase3DDiagrams_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Stool2DDiagrams, stage.Stool2DDiagrams_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Stool3DDiagrams, stage.Stool3DDiagrams_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Clock2DDiagrams, stage.Clock2DDiagrams_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Clock3DDiagrams, stage.Clock3DDiagrams_instance)
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

func (reference *StackOfRotatedVaseTrapezeRingsShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackOfVaseTrapezeRingsShape) GongReconstructPointersFromInstances(stage *Stage) {
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
	__gong__reconstructPointerFromInstance(&reference.SampledPoints3DShape, stage.SampledPoints3DShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.Rendered3DShape, stage.Rendered3DShapes_instance)
	// insertion point for slice of pointers fields
}

func (reference *TiledFloor3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopCurvePlane1Shape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopCurvePlane2Shape) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *TubeVase3DDiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Rendered3DShape, stage.Rendered3DShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.SampledPoints3DShape, stage.SampledPoints3DShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.OriginalPoints3DShape, stage.OriginalPoints3DShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.Angle0Shape, stage.Angle0Shapes_instance)
	__gong__reconstructPointerFromInstance(&reference.TopCurvePlane1Shape, stage.TopCurvePlane1Shapes_instance)
	__gong__reconstructPointerFromInstance(&reference.BottomCurvePlane1Shape, stage.BottomCurvePlane1Shapes_instance)
	__gong__reconstructPointerFromInstance(&reference.TopCurvePlane2Shape, stage.TopCurvePlane2Shapes_instance)
	__gong__reconstructPointerFromInstance(&reference.BottomCurvePlane2Shape, stage.BottomCurvePlane2Shapes_instance)
	__gong__reconstructPointerFromInstance(&reference.VaseTrapezeRingShape, stage.VaseTrapezeRingShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.StackOfVaseTrapezeRingsShape, stage.StackOfVaseTrapezeRingsShapes_instance)
	__gong__reconstructPointerFromInstance(&reference.StackOfRotatedVaseTrapezeRingsShape, stage.StackOfRotatedVaseTrapezeRingsShapes_instance)
	// insertion point for slice of pointers fields
}

func (reference *TubeVaseAbstract) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Vase2DDiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *VaseTrapezeRingShape) GongReconstructPointersFromInstances(stage *Stage) {
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
func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongDiff(stage *Stage, bottomcurveplane1shapeOther *BottomCurvePlane1Shape) (diffs []string) {
	// insertion point for field diffs
	if bottomcurveplane1shape.Name != bottomcurveplane1shapeOther.Name {
		diffs = append(diffs, bottomcurveplane1shape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongDiff(stage *Stage, bottomcurveplane2shapeOther *BottomCurvePlane2Shape) (diffs []string) {
	// insertion point for field diffs
	if bottomcurveplane2shape.Name != bottomcurveplane2shapeOther.Name {
		diffs = append(diffs, bottomcurveplane2shape.GongMarshallField(stage, "Name"))
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
	if clock3ddiagram.SampledPoints3DShape != clock3ddiagramOther.SampledPoints3DShape {
		diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
	}
	if clock3ddiagram.IsHiddenTiledFloor3DShape != clock3ddiagramOther.IsHiddenTiledFloor3DShape {
		diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "IsHiddenTiledFloor3DShape"))
	}
	if clock3ddiagram.Rendered3DShape != clock3ddiagramOther.Rendered3DShape {
		diffs = append(diffs, clock3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
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
func (leaves3dshape *Leaves3DShape) GongDiff(stage *Stage, leaves3dshapeOther *Leaves3DShape) (diffs []string) {
	// insertion point for field diffs
	if leaves3dshape.Name != leaves3dshapeOther.Name {
		diffs = append(diffs, leaves3dshape.GongMarshallField(stage, "Name"))
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
	if ops := __gong__diffSliceOfPointers(stage, library, "Plants", libraryOther.Plants, library.Plants); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries); ops != "" {
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
	if plant3ddiagram.StemCylinder3DShape != plant3ddiagramOther.StemCylinder3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "StemCylinder3DShape"))
	}
	if plant3ddiagram.IsHiddenParastichyNCurves3DShape != plant3ddiagramOther.IsHiddenParastichyNCurves3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsHiddenParastichyNCurves3DShape"))
	}
	if plant3ddiagram.ParastichyNCurves3DShape != plant3ddiagramOther.ParastichyNCurves3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "ParastichyNCurves3DShape"))
	}
	if plant3ddiagram.IsHiddenParastichyMCurves3DShape != plant3ddiagramOther.IsHiddenParastichyMCurves3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsHiddenParastichyMCurves3DShape"))
	}
	if plant3ddiagram.ParastichyMCurves3DShape != plant3ddiagramOther.ParastichyMCurves3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "ParastichyMCurves3DShape"))
	}
	if plant3ddiagram.IsHiddenCutLine3DShape != plant3ddiagramOther.IsHiddenCutLine3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsHiddenCutLine3DShape"))
	}
	if plant3ddiagram.CutLine3DShape != plant3ddiagramOther.CutLine3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "CutLine3DShape"))
	}
	if plant3ddiagram.IsHiddenCircumference3DShape != plant3ddiagramOther.IsHiddenCircumference3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsHiddenCircumference3DShape"))
	}
	if plant3ddiagram.Circumference3DShape != plant3ddiagramOther.Circumference3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "Circumference3DShape"))
	}
	if plant3ddiagram.IsHiddenTiledFloor3DShape != plant3ddiagramOther.IsHiddenTiledFloor3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsHiddenTiledFloor3DShape"))
	}
	if plant3ddiagram.IsHiddenLeaves3DShape != plant3ddiagramOther.IsHiddenLeaves3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsHiddenLeaves3DShape"))
	}
	if plant3ddiagram.Leaves3DShape != plant3ddiagramOther.Leaves3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "Leaves3DShape"))
	}
	if plant3ddiagram.Rendered3DShape != plant3ddiagramOther.Rendered3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
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
	if plantabstract.TubeVaseAbstract != plantabstractOther.TubeVaseAbstract {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "TubeVaseAbstract"))
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
	if ops := __gong__diffSliceOfPointers(stage, plantabstract, "Plant2DDiagrams", plantabstractOther.Plant2DDiagrams, plantabstract.Plant2DDiagrams); ops != "" {
		diffs = append(diffs, ops)
	}
	if plantabstract.IsPlant3DDiagramsNodeExpanded != plantabstractOther.IsPlant3DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsPlant3DDiagramsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, plantabstract, "Plant3DDiagrams", plantabstractOther.Plant3DDiagrams, plantabstract.Plant3DDiagrams); ops != "" {
		diffs = append(diffs, ops)
	}
	if plantabstract.IsVase2DDiagramsNodeExpanded != plantabstractOther.IsVase2DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsVase2DDiagramsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, plantabstract, "Vase2DDiagrams", plantabstractOther.Vase2DDiagrams, plantabstract.Vase2DDiagrams); ops != "" {
		diffs = append(diffs, ops)
	}
	if plantabstract.IsTubeVase3DDiagramsNodeExpanded != plantabstractOther.IsTubeVase3DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsTubeVase3DDiagramsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, plantabstract, "TubeVase3DDiagrams", plantabstractOther.TubeVase3DDiagrams, plantabstract.TubeVase3DDiagrams); ops != "" {
		diffs = append(diffs, ops)
	}
	if plantabstract.IsStool2DDiagramsNodeExpanded != plantabstractOther.IsStool2DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsStool2DDiagramsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, plantabstract, "Stool2DDiagrams", plantabstractOther.Stool2DDiagrams, plantabstract.Stool2DDiagrams); ops != "" {
		diffs = append(diffs, ops)
	}
	if plantabstract.IsStool3DDiagramsNodeExpanded != plantabstractOther.IsStool3DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsStool3DDiagramsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, plantabstract, "Stool3DDiagrams", plantabstractOther.Stool3DDiagrams, plantabstract.Stool3DDiagrams); ops != "" {
		diffs = append(diffs, ops)
	}
	if plantabstract.IsClock2DDiagramsNodeExpanded != plantabstractOther.IsClock2DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsClock2DDiagramsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, plantabstract, "Clock2DDiagrams", plantabstractOther.Clock2DDiagrams, plantabstract.Clock2DDiagrams); ops != "" {
		diffs = append(diffs, ops)
	}
	if plantabstract.IsClock3DDiagramsNodeExpanded != plantabstractOther.IsClock3DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsClock3DDiagramsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, plantabstract, "Clock3DDiagrams", plantabstractOther.Clock3DDiagrams, plantabstract.Clock3DDiagrams); ops != "" {
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
func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongDiff(stage *Stage, stackofrotatedvasetrapezeringsshapeOther *StackOfRotatedVaseTrapezeRingsShape) (diffs []string) {
	// insertion point for field diffs
	if stackofrotatedvasetrapezeringsshape.Name != stackofrotatedvasetrapezeringsshapeOther.Name {
		diffs = append(diffs, stackofrotatedvasetrapezeringsshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongDiff(stage *Stage, stackofvasetrapezeringsshapeOther *StackOfVaseTrapezeRingsShape) (diffs []string) {
	// insertion point for field diffs
	if stackofvasetrapezeringsshape.Name != stackofvasetrapezeringsshapeOther.Name {
		diffs = append(diffs, stackofvasetrapezeringsshape.GongMarshallField(stage, "Name"))
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
	if stemcylinder3dshape.Transparency != stemcylinder3dshapeOther.Transparency {
		diffs = append(diffs, stemcylinder3dshape.GongMarshallField(stage, "Transparency"))
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
	if stool3ddiagram.SampledPoints3DShape != stool3ddiagramOther.SampledPoints3DShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
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
	if stool3ddiagram.Rendered3DShape != stool3ddiagramOther.Rendered3DShape {
		diffs = append(diffs, stool3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
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
func (tiledfloor3dshape *TiledFloor3DShape) GongDiff(stage *Stage, tiledfloor3dshapeOther *TiledFloor3DShape) (diffs []string) {
	// insertion point for field diffs
	if tiledfloor3dshape.Name != tiledfloor3dshapeOther.Name {
		diffs = append(diffs, tiledfloor3dshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topcurveplane1shape *TopCurvePlane1Shape) GongDiff(stage *Stage, topcurveplane1shapeOther *TopCurvePlane1Shape) (diffs []string) {
	// insertion point for field diffs
	if topcurveplane1shape.Name != topcurveplane1shapeOther.Name {
		diffs = append(diffs, topcurveplane1shape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topcurveplane2shape *TopCurvePlane2Shape) GongDiff(stage *Stage, topcurveplane2shapeOther *TopCurvePlane2Shape) (diffs []string) {
	// insertion point for field diffs
	if topcurveplane2shape.Name != topcurveplane2shapeOther.Name {
		diffs = append(diffs, topcurveplane2shape.GongMarshallField(stage, "Name"))
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
func (tubevase3ddiagram *TubeVase3DDiagram) GongDiff(stage *Stage, tubevase3ddiagramOther *TubeVase3DDiagram) (diffs []string) {
	// insertion point for field diffs
	if tubevase3ddiagram.Name != tubevase3ddiagramOther.Name {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "Name"))
	}
	if tubevase3ddiagram.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon != tubevase3ddiagramOther.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon"))
	}
	if tubevase3ddiagram.IsHiddenTorusStackShape != tubevase3ddiagramOther.IsHiddenTorusStackShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenTorusStackShape"))
	}
	if tubevase3ddiagram.IsHiddenVerticalTorusStackShape != tubevase3ddiagramOther.IsHiddenVerticalTorusStackShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenVerticalTorusStackShape"))
	}
	if tubevase3ddiagram.IsHiddenPartiallyRotatedTorusShape != tubevase3ddiagramOther.IsHiddenPartiallyRotatedTorusShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenPartiallyRotatedTorusShape"))
	}
	if tubevase3ddiagram.IsHiddenStackOfPartiallyRotatedTorusShape != tubevase3ddiagramOther.IsHiddenStackOfPartiallyRotatedTorusShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenStackOfPartiallyRotatedTorusShape"))
	}
	if tubevase3ddiagram.IsHiddenPointsAndLines3DShape != tubevase3ddiagramOther.IsHiddenPointsAndLines3DShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenPointsAndLines3DShape"))
	}
	if tubevase3ddiagram.IsHiddenKeyHole3DShape != tubevase3ddiagramOther.IsHiddenKeyHole3DShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenKeyHole3DShape"))
	}
	if tubevase3ddiagram.IsHiddenKey3DShape != tubevase3ddiagramOther.IsHiddenKey3DShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenKey3DShape"))
	}
	if tubevase3ddiagram.IsHiddenVolumeKey3DShape != tubevase3ddiagramOther.IsHiddenVolumeKey3DShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenVolumeKey3DShape"))
	}
	if tubevase3ddiagram.IsHiddenTorusEdge3DShape != tubevase3ddiagramOther.IsHiddenTorusEdge3DShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenTorusEdge3DShape"))
	}
	if tubevase3ddiagram.IsHiddenSampledPoints3DShape != tubevase3ddiagramOther.IsHiddenSampledPoints3DShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenSampledPoints3DShape"))
	}
	if tubevase3ddiagram.IsHiddenOriginalPoints3DShape != tubevase3ddiagramOther.IsHiddenOriginalPoints3DShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenOriginalPoints3DShape"))
	}
	if tubevase3ddiagram.IsHiddenAngle0Shape != tubevase3ddiagramOther.IsHiddenAngle0Shape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenAngle0Shape"))
	}
	if tubevase3ddiagram.IsHiddenTiledFloor3DShape != tubevase3ddiagramOther.IsHiddenTiledFloor3DShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenTiledFloor3DShape"))
	}
	if tubevase3ddiagram.IsHiddenTopCurvePlane1Shape != tubevase3ddiagramOther.IsHiddenTopCurvePlane1Shape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenTopCurvePlane1Shape"))
	}
	if tubevase3ddiagram.IsHiddenBottomCurvePlane1Shape != tubevase3ddiagramOther.IsHiddenBottomCurvePlane1Shape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenBottomCurvePlane1Shape"))
	}
	if tubevase3ddiagram.IsHiddenTopCurvePlane2Shape != tubevase3ddiagramOther.IsHiddenTopCurvePlane2Shape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenTopCurvePlane2Shape"))
	}
	if tubevase3ddiagram.IsHiddenBottomCurvePlane2Shape != tubevase3ddiagramOther.IsHiddenBottomCurvePlane2Shape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenBottomCurvePlane2Shape"))
	}
	if tubevase3ddiagram.IsHiddenVaseTrapezeRingShape != tubevase3ddiagramOther.IsHiddenVaseTrapezeRingShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenVaseTrapezeRingShape"))
	}
	if tubevase3ddiagram.IsHiddenStackOfVaseTrapezeRingsShape != tubevase3ddiagramOther.IsHiddenStackOfVaseTrapezeRingsShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenStackOfVaseTrapezeRingsShape"))
	}
	if tubevase3ddiagram.IsHiddenStackOfRotatedVaseTrapezeRingsShape != tubevase3ddiagramOther.IsHiddenStackOfRotatedVaseTrapezeRingsShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsHiddenStackOfRotatedVaseTrapezeRingsShape"))
	}
	if tubevase3ddiagram.Rendered3DShape != tubevase3ddiagramOther.Rendered3DShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
	}
	if tubevase3ddiagram.SampledPoints3DShape != tubevase3ddiagramOther.SampledPoints3DShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
	}
	if tubevase3ddiagram.OriginalPoints3DShape != tubevase3ddiagramOther.OriginalPoints3DShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "OriginalPoints3DShape"))
	}
	if tubevase3ddiagram.Angle0Shape != tubevase3ddiagramOther.Angle0Shape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "Angle0Shape"))
	}
	if tubevase3ddiagram.TopCurvePlane1Shape != tubevase3ddiagramOther.TopCurvePlane1Shape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "TopCurvePlane1Shape"))
	}
	if tubevase3ddiagram.BottomCurvePlane1Shape != tubevase3ddiagramOther.BottomCurvePlane1Shape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "BottomCurvePlane1Shape"))
	}
	if tubevase3ddiagram.TopCurvePlane2Shape != tubevase3ddiagramOther.TopCurvePlane2Shape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "TopCurvePlane2Shape"))
	}
	if tubevase3ddiagram.BottomCurvePlane2Shape != tubevase3ddiagramOther.BottomCurvePlane2Shape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "BottomCurvePlane2Shape"))
	}
	if tubevase3ddiagram.VaseTrapezeRingShape != tubevase3ddiagramOther.VaseTrapezeRingShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "VaseTrapezeRingShape"))
	}
	if tubevase3ddiagram.StackOfVaseTrapezeRingsShape != tubevase3ddiagramOther.StackOfVaseTrapezeRingsShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "StackOfVaseTrapezeRingsShape"))
	}
	if tubevase3ddiagram.StackOfRotatedVaseTrapezeRingsShape != tubevase3ddiagramOther.StackOfRotatedVaseTrapezeRingsShape {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "StackOfRotatedVaseTrapezeRingsShape"))
	}
	if tubevase3ddiagram.IsChecked != tubevase3ddiagramOther.IsChecked {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsChecked"))
	}
	if tubevase3ddiagram.ComputedPrefix != tubevase3ddiagramOther.ComputedPrefix {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if tubevase3ddiagram.IsExpanded != tubevase3ddiagramOther.IsExpanded {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tubevaseabstract *TubeVaseAbstract) GongDiff(stage *Stage, tubevaseabstractOther *TubeVaseAbstract) (diffs []string) {
	// insertion point for field diffs
	if tubevaseabstract.Name != tubevaseabstractOther.Name {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "Name"))
	}
	if tubevaseabstract.Z_Ribbon != tubevaseabstractOther.Z_Ribbon {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "Z_Ribbon"))
	}
	if tubevaseabstract.RibbonVerticalScale != tubevaseabstractOther.RibbonVerticalScale {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "RibbonVerticalScale"))
	}
	if tubevaseabstract.Plane1Height != tubevaseabstractOther.Plane1Height {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "Plane1Height"))
	}
	if tubevaseabstract.Plane2Height != tubevaseabstractOther.Plane2Height {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "Plane2Height"))
	}
	if tubevaseabstract.ProjectionAngle != tubevaseabstractOther.ProjectionAngle {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "ProjectionAngle"))
	}
	if tubevaseabstract.RelativeVerticalThickness != tubevaseabstractOther.RelativeVerticalThickness {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "RelativeVerticalThickness"))
	}
	if tubevaseabstract.RelativeRadialThickness != tubevaseabstractOther.RelativeRadialThickness {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "RelativeRadialThickness"))
	}
	if tubevaseabstract.RelativeCuttedStackFloorHeight != tubevaseabstractOther.RelativeCuttedStackFloorHeight {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "RelativeCuttedStackFloorHeight"))
	}
	if tubevaseabstract.RelativeRotatedTorusSeparation != tubevaseabstractOther.RelativeRotatedTorusSeparation {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "RelativeRotatedTorusSeparation"))
	}
	if tubevaseabstract.RotationRatio != tubevaseabstractOther.RotationRatio {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "RotationRatio"))
	}
	if tubevaseabstract.RadialRepetitions != tubevaseabstractOther.RadialRepetitions {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "RadialRepetitions"))
	}
	if tubevaseabstract.Transparency != tubevaseabstractOther.Transparency {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "Transparency"))
	}
	if tubevaseabstract.HasAlternatingRingColors != tubevaseabstractOther.HasAlternatingRingColors {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "HasAlternatingRingColors"))
	}
	if tubevaseabstract.RelativeTrajectoryOffsetX != tubevaseabstractOther.RelativeTrajectoryOffsetX {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "RelativeTrajectoryOffsetX"))
	}
	if tubevaseabstract.RelativeTrajectoryOffsetY != tubevaseabstractOther.RelativeTrajectoryOffsetY {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "RelativeTrajectoryOffsetY"))
	}
	if tubevaseabstract.NbStepP1P2 != tubevaseabstractOther.NbStepP1P2 {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "NbStepP1P2"))
	}
	if tubevaseabstract.ChosenStep != tubevaseabstractOther.ChosenStep {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "ChosenStep"))
	}
	if tubevaseabstract.RelativeHorizontalRingsHeight != tubevaseabstractOther.RelativeHorizontalRingsHeight {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "RelativeHorizontalRingsHeight"))
	}
	if tubevaseabstract.OffsetKeyX != tubevaseabstractOther.OffsetKeyX {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "OffsetKeyX"))
	}
	if tubevaseabstract.OffsetKeyY != tubevaseabstractOther.OffsetKeyY {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "OffsetKeyY"))
	}
	if tubevaseabstract.HeightKey != tubevaseabstractOther.HeightKey {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "HeightKey"))
	}
	if tubevaseabstract.WidthKey != tubevaseabstractOther.WidthKey {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "WidthKey"))
	}
	if tubevaseabstract.RelativeKeySize != tubevaseabstractOther.RelativeKeySize {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "RelativeKeySize"))
	}
	if tubevaseabstract.MovieNbFrames != tubevaseabstractOther.MovieNbFrames {
		diffs = append(diffs, tubevaseabstract.GongMarshallField(stage, "MovieNbFrames"))
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
func (vasetrapezeringshape *VaseTrapezeRingShape) GongDiff(stage *Stage, vasetrapezeringshapeOther *VaseTrapezeRingShape) (diffs []string) {
	// insertion point for field diffs
	if vasetrapezeringshape.Name != vasetrapezeringshapeOther.Name {
		diffs = append(diffs, vasetrapezeringshape.GongMarshallField(stage, "Name"))
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

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
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

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
