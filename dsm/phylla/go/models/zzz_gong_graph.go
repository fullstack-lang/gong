// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (angle0shape *Angle0Shape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Angle0Shapes[angle0shape]

	return
}

func (stage *Stage) IsStagedAngle0Shape(angle0shape *Angle0Shape) (ok bool) {

	return angle0shape.GongIsStaged(stage)
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ArcNormalVectorShapes[arcnormalvectorshape]

	return
}

func (stage *Stage) IsStagedArcNormalVectorShape(arcnormalvectorshape *ArcNormalVectorShape) (ok bool) {

	return arcnormalvectorshape.GongIsStaged(stage)
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ArcNormalVectorShapeGrids[arcnormalvectorshapegrid]

	return
}

func (stage *Stage) IsStagedArcNormalVectorShapeGrid(arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) (ok bool) {

	return arcnormalvectorshapegrid.GongIsStaged(stage)
}

func (axesshape *AxesShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.AxesShapes[axesshape]

	return
}

func (stage *Stage) IsStagedAxesShape(axesshape *AxesShape) (ok bool) {

	return axesshape.GongIsStaged(stage)
}

func (basevectorshape *BaseVectorShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.BaseVectorShapes[basevectorshape]

	return
}

func (stage *Stage) IsStagedBaseVectorShape(basevectorshape *BaseVectorShape) (ok bool) {

	return basevectorshape.GongIsStaged(stage)
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.BaseVectorShapeGrids[basevectorshapegrid]

	return
}

func (stage *Stage) IsStagedBaseVectorShapeGrid(basevectorshapegrid *BaseVectorShapeGrid) (ok bool) {

	return basevectorshapegrid.GongIsStaged(stage)
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ChosenP1P2PairShapes[chosenp1p2pairshape]

	return
}

func (stage *Stage) IsStagedChosenP1P2PairShape(chosenp1p2pairshape *ChosenP1P2PairShape) (ok bool) {

	return chosenp1p2pairshape.GongIsStaged(stage)
}

func (circlegridshape *CircleGridShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.CircleGridShapes[circlegridshape]

	return
}

func (stage *Stage) IsStagedCircleGridShape(circlegridshape *CircleGridShape) (ok bool) {

	return circlegridshape.GongIsStaged(stage)
}

func (circumference3dshape *Circumference3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Circumference3DShapes[circumference3dshape]

	return
}

func (stage *Stage) IsStagedCircumference3DShape(circumference3dshape *Circumference3DShape) (ok bool) {

	return circumference3dshape.GongIsStaged(stage)
}

func (clock2ddiagram *Clock2DDiagram) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Clock2DDiagrams[clock2ddiagram]

	return
}

func (stage *Stage) IsStagedClock2DDiagram(clock2ddiagram *Clock2DDiagram) (ok bool) {

	return clock2ddiagram.GongIsStaged(stage)
}

func (clock3ddiagram *Clock3DDiagram) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Clock3DDiagrams[clock3ddiagram]

	return
}

func (stage *Stage) IsStagedClock3DDiagram(clock3ddiagram *Clock3DDiagram) (ok bool) {

	return clock3ddiagram.GongIsStaged(stage)
}

func (clockabstract *ClockAbstract) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ClockAbstracts[clockabstract]

	return
}

func (stage *Stage) IsStagedClockAbstract(clockabstract *ClockAbstract) (ok bool) {

	return clockabstract.GongIsStaged(stage)
}

func (clocktopcurveshape *ClockTopCurveShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ClockTopCurveShapes[clocktopcurveshape]

	return
}

func (stage *Stage) IsStagedClockTopCurveShape(clocktopcurveshape *ClockTopCurveShape) (ok bool) {

	return clocktopcurveshape.GongIsStaged(stage)
}

func (cutline3dshape *CutLine3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.CutLine3DShapes[cutline3dshape]

	return
}

func (stage *Stage) IsStagedCutLine3DShape(cutline3dshape *CutLine3DShape) (ok bool) {

	return cutline3dshape.GongIsStaged(stage)
}

func (endarcshape *EndArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.EndArcShapes[endarcshape]

	return
}

func (stage *Stage) IsStagedEndArcShape(endarcshape *EndArcShape) (ok bool) {

	return endarcshape.GongIsStaged(stage)
}

func (endarcshapegrid *EndArcShapeGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.EndArcShapeGrids[endarcshapegrid]

	return
}

func (stage *Stage) IsStagedEndArcShapeGrid(endarcshapegrid *EndArcShapeGrid) (ok bool) {

	return endarcshapegrid.GongIsStaged(stage)
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.EndHalfwayArcShapes[endhalfwayarcshape]

	return
}

func (stage *Stage) IsStagedEndHalfwayArcShape(endhalfwayarcshape *EndHalfwayArcShape) (ok bool) {

	return endhalfwayarcshape.GongIsStaged(stage)
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.EndHalfwayArcShapeGrids[endhalfwayarcshapegrid]

	return
}

func (stage *Stage) IsStagedEndHalfwayArcShapeGrid(endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) (ok bool) {

	return endhalfwayarcshapegrid.GongIsStaged(stage)
}

func (explanationtextshape *ExplanationTextShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ExplanationTextShapes[explanationtextshape]

	return
}

func (stage *Stage) IsStagedExplanationTextShape(explanationtextshape *ExplanationTextShape) (ok bool) {

	return explanationtextshape.GongIsStaged(stage)
}

func (eye3dshape *Eye3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Eye3DShapes[eye3dshape]

	return
}

func (stage *Stage) IsStagedEye3DShape(eye3dshape *Eye3DShape) (ok bool) {

	return eye3dshape.GongIsStaged(stage)
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.EyeCornersSampledPoints3DShapes[eyecornerssampledpoints3dshape]

	return
}

func (stage *Stage) IsStagedEyeCornersSampledPoints3DShape(eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) (ok bool) {

	return eyecornerssampledpoints3dshape.GongIsStaged(stage)
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.EyeSampledPoints3DShapes[eyesampledpoints3dshape]

	return
}

func (stage *Stage) IsStagedEyeSampledPoints3DShape(eyesampledpoints3dshape *EyeSampledPoints3DShape) (ok bool) {

	return eyesampledpoints3dshape.GongIsStaged(stage)
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.EyeSeatBottomCurveShapes[eyeseatbottomcurveshape]

	return
}

func (stage *Stage) IsStagedEyeSeatBottomCurveShape(eyeseatbottomcurveshape *EyeSeatBottomCurveShape) (ok bool) {

	return eyeseatbottomcurveshape.GongIsStaged(stage)
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.EyeStoolBottomCurveShapes[eyestoolbottomcurveshape]

	return
}

func (stage *Stage) IsStagedEyeStoolBottomCurveShape(eyestoolbottomcurveshape *EyeStoolBottomCurveShape) (ok bool) {

	return eyestoolbottomcurveshape.GongIsStaged(stage)
}

func (eyevolume3dshape *EyeVolume3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.EyeVolume3DShapes[eyevolume3dshape]

	return
}

func (stage *Stage) IsStagedEyeVolume3DShape(eyevolume3dshape *EyeVolume3DShape) (ok bool) {

	return eyevolume3dshape.GongIsStaged(stage)
}

func (gridpathshape *GridPathShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GridPathShapes[gridpathshape]

	return
}

func (stage *Stage) IsStagedGridPathShape(gridpathshape *GridPathShape) (ok bool) {

	return gridpathshape.GongIsStaged(stage)
}

func (growthcurve2d *GrowthCurve2D) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GrowthCurve2Ds[growthcurve2d]

	return
}

func (stage *Stage) IsStagedGrowthCurve2D(growthcurve2d *GrowthCurve2D) (ok bool) {

	return growthcurve2d.GongIsStaged(stage)
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GrowthCurve2DRibbons[growthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedGrowthCurve2DRibbon(growthcurve2dribbon *GrowthCurve2DRibbon) (ok bool) {

	return growthcurve2dribbon.GongIsStaged(stage)
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GrowthCurve2DRibbonEndShapes[growthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedGrowthCurve2DRibbonEndShape(growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) (ok bool) {

	return growthcurve2dribbonendshape.GongIsStaged(stage)
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GrowthCurve2DRibbonStartShapes[growthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedGrowthCurve2DRibbonStartShape(growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) (ok bool) {

	return growthcurve2dribbonstartshape.GongIsStaged(stage)
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GrowthCurveRhombusGridShapes[growthcurverhombusgridshape]

	return
}

func (stage *Stage) IsStagedGrowthCurveRhombusGridShape(growthcurverhombusgridshape *GrowthCurveRhombusGridShape) (ok bool) {

	return growthcurverhombusgridshape.GongIsStaged(stage)
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GrowthCurveRhombusShapes[growthcurverhombusshape]

	return
}

func (stage *Stage) IsStagedGrowthCurveRhombusShape(growthcurverhombusshape *GrowthCurveRhombusShape) (ok bool) {

	return growthcurverhombusshape.GongIsStaged(stage)
}

func (growthvectorshape *GrowthVectorShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GrowthVectorShapes[growthvectorshape]

	return
}

func (stage *Stage) IsStagedGrowthVectorShape(growthvectorshape *GrowthVectorShape) (ok bool) {

	return growthvectorshape.GongIsStaged(stage)
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.InitialRhombusGridShapes[initialrhombusgridshape]

	return
}

func (stage *Stage) IsStagedInitialRhombusGridShape(initialrhombusgridshape *InitialRhombusGridShape) (ok bool) {

	return initialrhombusgridshape.GongIsStaged(stage)
}

func (initialrhombusshape *InitialRhombusShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.InitialRhombusShapes[initialrhombusshape]

	return
}

func (stage *Stage) IsStagedInitialRhombusShape(initialrhombusshape *InitialRhombusShape) (ok bool) {

	return initialrhombusshape.GongIsStaged(stage)
}

func (key3dshape *Key3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Key3DShapes[key3dshape]

	return
}

func (stage *Stage) IsStagedKey3DShape(key3dshape *Key3DShape) (ok bool) {

	return key3dshape.GongIsStaged(stage)
}

func (keyhole3dshape *KeyHole3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.KeyHole3DShapes[keyhole3dshape]

	return
}

func (stage *Stage) IsStagedKeyHole3DShape(keyhole3dshape *KeyHole3DShape) (ok bool) {

	return keyhole3dshape.GongIsStaged(stage)
}

func (keyholeshape *KeyHoleShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.KeyHoleShapes[keyholeshape]

	return
}

func (stage *Stage) IsStagedKeyHoleShape(keyholeshape *KeyHoleShape) (ok bool) {

	return keyholeshape.GongIsStaged(stage)
}

func (leaves3dshape *Leaves3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Leaves3DShapes[leaves3dshape]

	return
}

func (stage *Stage) IsStagedLeaves3DShape(leaves3dshape *Leaves3DShape) (ok bool) {

	return leaves3dshape.GongIsStaged(stage)
}

func (library *Library) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	return library.GongIsStaged(stage)
}

func (midarcvectorshape *MidArcVectorShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.MidArcVectorShapes[midarcvectorshape]

	return
}

func (stage *Stage) IsStagedMidArcVectorShape(midarcvectorshape *MidArcVectorShape) (ok bool) {

	return midarcvectorshape.GongIsStaged(stage)
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.MidArcVectorShapeGrids[midarcvectorshapegrid]

	return
}

func (stage *Stage) IsStagedMidArcVectorShapeGrid(midarcvectorshapegrid *MidArcVectorShapeGrid) (ok bool) {

	return midarcvectorshapegrid.GongIsStaged(stage)
}

func (musicabstract *MusicAbstract) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.MusicAbstracts[musicabstract]

	return
}

func (stage *Stage) IsStagedMusicAbstract(musicabstract *MusicAbstract) (ok bool) {

	return musicabstract.GongIsStaged(stage)
}

func (originalpoints3dshape *OriginalPoints3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.OriginalPoints3DShapes[originalpoints3dshape]

	return
}

func (stage *Stage) IsStagedOriginalPoints3DShape(originalpoints3dshape *OriginalPoints3DShape) (ok bool) {

	return originalpoints3dshape.GongIsStaged(stage)
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ParastichyMCurves3DShapes[parastichymcurves3dshape]

	return
}

func (stage *Stage) IsStagedParastichyMCurves3DShape(parastichymcurves3dshape *ParastichyMCurves3DShape) (ok bool) {

	return parastichymcurves3dshape.GongIsStaged(stage)
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ParastichyNCurves3DShapes[parastichyncurves3dshape]

	return
}

func (stage *Stage) IsStagedParastichyNCurves3DShape(parastichyncurves3dshape *ParastichyNCurves3DShape) (ok bool) {

	return parastichyncurves3dshape.GongIsStaged(stage)
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DRibbons[partiallygrowthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DRibbon(partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) (ok bool) {

	return partiallygrowthcurve2dribbon.GongIsStaged(stage)
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DRibbonEndShapes[partiallygrowthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DRibbonEndShape(partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) (ok bool) {

	return partiallygrowthcurve2dribbonendshape.GongIsStaged(stage)
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DRibbonStartShapes[partiallygrowthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DRibbonStartShape(partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) (ok bool) {

	return partiallygrowthcurve2dribbonstartshape.GongIsStaged(stage)
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectorys[partiallygrowthcurve2dtrajectory]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectory(partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) (ok bool) {

	return partiallygrowthcurve2dtrajectory.GongIsStaged(stage)
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes[partiallygrowthcurve2dtrajectoryp1curveshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryP1CurveShape(partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) (ok bool) {

	return partiallygrowthcurve2dtrajectoryp1curveshape.GongIsStaged(stage)
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryP1P2s[partiallygrowthcurve2dtrajectoryp1p2]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryP1P2(partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) (ok bool) {

	return partiallygrowthcurve2dtrajectoryp1p2.GongIsStaged(stage)
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes[partiallygrowthcurve2dtrajectoryp1p2pairlineshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) (ok bool) {

	return partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongIsStaged(stage)
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes[partiallygrowthcurve2dtrajectoryp1pointshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryP1PointShape(partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) (ok bool) {

	return partiallygrowthcurve2dtrajectoryp1pointshape.GongIsStaged(stage)
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes[partiallygrowthcurve2dtrajectoryp2curveshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryP2CurveShape(partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) (ok bool) {

	return partiallygrowthcurve2dtrajectoryp2curveshape.GongIsStaged(stage)
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes[partiallygrowthcurve2dtrajectoryp2pointshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryP2PointShape(partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) (ok bool) {

	return partiallygrowthcurve2dtrajectoryp2pointshape.GongIsStaged(stage)
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyGrowthCurve2DTrajectoryShapes[partiallygrowthcurve2dtrajectoryshape]

	return
}

func (stage *Stage) IsStagedPartiallyGrowthCurve2DTrajectoryShape(partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) (ok bool) {

	return partiallygrowthcurve2dtrajectoryshape.GongIsStaged(stage)
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyRotatedSeatBottomCurveShapes[partiallyrotatedseatbottomcurveshape]

	return
}

func (stage *Stage) IsStagedPartiallyRotatedSeatBottomCurveShape(partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) (ok bool) {

	return partiallyrotatedseatbottomcurveshape.GongIsStaged(stage)
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyRotatedSeatTopCurveShapes[partiallyrotatedseattopcurveshape]

	return
}

func (stage *Stage) IsStagedPartiallyRotatedSeatTopCurveShape(partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) (ok bool) {

	return partiallyrotatedseattopcurveshape.GongIsStaged(stage)
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartiallyRotatedTorusShapes[partiallyrotatedtorusshape]

	return
}

func (stage *Stage) IsStagedPartiallyRotatedTorusShape(partiallyrotatedtorusshape *PartiallyRotatedTorusShape) (ok bool) {

	return partiallyrotatedtorusshape.GongIsStaged(stage)
}

func (perpendicularvector *PerpendicularVector) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PerpendicularVectors[perpendicularvector]

	return
}

func (stage *Stage) IsStagedPerpendicularVector(perpendicularvector *PerpendicularVector) (ok bool) {

	return perpendicularvector.GongIsStaged(stage)
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PerpendicularVectorGrids[perpendicularvectorgrid]

	return
}

func (stage *Stage) IsStagedPerpendicularVectorGrid(perpendicularvectorgrid *PerpendicularVectorGrid) (ok bool) {

	return perpendicularvectorgrid.GongIsStaged(stage)
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PerpendicularVectorGridHalfways[perpendicularvectorgridhalfway]

	return
}

func (stage *Stage) IsStagedPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) (ok bool) {

	return perpendicularvectorgridhalfway.GongIsStaged(stage)
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PerpendicularVectorHalfways[perpendicularvectorhalfway]

	return
}

func (stage *Stage) IsStagedPerpendicularVectorHalfway(perpendicularvectorhalfway *PerpendicularVectorHalfway) (ok bool) {

	return perpendicularvectorhalfway.GongIsStaged(stage)
}

func (plant2ddiagram *Plant2DDiagram) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Plant2DDiagrams[plant2ddiagram]

	return
}

func (stage *Stage) IsStagedPlant2DDiagram(plant2ddiagram *Plant2DDiagram) (ok bool) {

	return plant2ddiagram.GongIsStaged(stage)
}

func (plant3ddiagram *Plant3DDiagram) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Plant3DDiagrams[plant3ddiagram]

	return
}

func (stage *Stage) IsStagedPlant3DDiagram(plant3ddiagram *Plant3DDiagram) (ok bool) {

	return plant3ddiagram.GongIsStaged(stage)
}

func (plantabstract *PlantAbstract) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PlantAbstracts[plantabstract]

	return
}

func (stage *Stage) IsStagedPlantAbstract(plantabstract *PlantAbstract) (ok bool) {

	return plantabstract.GongIsStaged(stage)
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PlantCircumferenceShapes[plantcircumferenceshape]

	return
}

func (stage *Stage) IsStagedPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape) (ok bool) {

	return plantcircumferenceshape.GongIsStaged(stage)
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PointsAndLines3DShapes[pointsandlines3dshape]

	return
}

func (stage *Stage) IsStagedPointsAndLines3DShape(pointsandlines3dshape *PointsAndLines3DShape) (ok bool) {

	return pointsandlines3dshape.GongIsStaged(stage)
}

func (pxshape *PxShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PxShapes[pxshape]

	return
}

func (stage *Stage) IsStagedPxShape(pxshape *PxShape) (ok bool) {

	return pxshape.GongIsStaged(stage)
}

func (rendered3dshape *Rendered3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Rendered3DShapes[rendered3dshape]

	return
}

func (stage *Stage) IsStagedRendered3DShape(rendered3dshape *Rendered3DShape) (ok bool) {

	return rendered3dshape.GongIsStaged(stage)
}

func (rhombusshape *RhombusShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RhombusShapes[rhombusshape]

	return
}

func (stage *Stage) IsStagedRhombusShape(rhombusshape *RhombusShape) (ok bool) {

	return rhombusshape.GongIsStaged(stage)
}

func (rhombusstuff *RhombusStuff) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RhombusStuffs[rhombusstuff]

	return
}

func (stage *Stage) IsStagedRhombusStuff(rhombusstuff *RhombusStuff) (ok bool) {

	return rhombusstuff.GongIsStaged(stage)
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RotatedRhombusGridShapes[rotatedrhombusgridshape]

	return
}

func (stage *Stage) IsStagedRotatedRhombusGridShape(rotatedrhombusgridshape *RotatedRhombusGridShape) (ok bool) {

	return rotatedrhombusgridshape.GongIsStaged(stage)
}

func (rotatedrhombusshape *RotatedRhombusShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RotatedRhombusShapes[rotatedrhombusshape]

	return
}

func (stage *Stage) IsStagedRotatedRhombusShape(rotatedrhombusshape *RotatedRhombusShape) (ok bool) {

	return rotatedrhombusshape.GongIsStaged(stage)
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RotatedSampledPoints3DShapes[rotatedsampledpoints3dshape]

	return
}

func (stage *Stage) IsStagedRotatedSampledPoints3DShape(rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) (ok bool) {

	return rotatedsampledpoints3dshape.GongIsStaged(stage)
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RotatedSeatAndLegs3DShapes[rotatedseatandlegs3dshape]

	return
}

func (stage *Stage) IsStagedRotatedSeatAndLegs3DShape(rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) (ok bool) {

	return rotatedseatandlegs3dshape.GongIsStaged(stage)
}

func (sampledpoints3dshape *SampledPoints3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SampledPoints3DShapes[sampledpoints3dshape]

	return
}

func (stage *Stage) IsStagedSampledPoints3DShape(sampledpoints3dshape *SampledPoints3DShape) (ok bool) {

	return sampledpoints3dshape.GongIsStaged(stage)
}

func (seat3dshape *Seat3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Seat3DShapes[seat3dshape]

	return
}

func (stage *Stage) IsStagedSeat3DShape(seat3dshape *Seat3DShape) (ok bool) {

	return seat3dshape.GongIsStaged(stage)
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SeatAndLegs3DShapes[seatandlegs3dshape]

	return
}

func (stage *Stage) IsStagedSeatAndLegs3DShape(seatandlegs3dshape *SeatAndLegs3DShape) (ok bool) {

	return seatandlegs3dshape.GongIsStaged(stage)
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SeatBottomCurveShapes[seatbottomcurveshape]

	return
}

func (stage *Stage) IsStagedSeatBottomCurveShape(seatbottomcurveshape *SeatBottomCurveShape) (ok bool) {

	return seatbottomcurveshape.GongIsStaged(stage)
}

func (seattopcurveshape *SeatTopCurveShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SeatTopCurveShapes[seattopcurveshape]

	return
}

func (stage *Stage) IsStagedSeatTopCurveShape(seattopcurveshape *SeatTopCurveShape) (ok bool) {

	return seattopcurveshape.GongIsStaged(stage)
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedBottomTopStartArcShapes[shiftedbottomtopstartarcshape]

	return
}

func (stage *Stage) IsStagedShiftedBottomTopStartArcShape(shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) (ok bool) {

	return shiftedbottomtopstartarcshape.GongIsStaged(stage)
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedBottomTopStartArcShapeGrids[shiftedbottomtopstartarcshapegrid]

	return
}

func (stage *Stage) IsStagedShiftedBottomTopStartArcShapeGrid(shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) (ok bool) {

	return shiftedbottomtopstartarcshapegrid.GongIsStaged(stage)
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedLeftGrowthCurve2DRibbons[shiftedleftgrowthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedShiftedLeftGrowthCurve2DRibbon(shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) (ok bool) {

	return shiftedleftgrowthcurve2dribbon.GongIsStaged(stage)
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedLeftGrowthCurve2DRibbonEndShapes[shiftedleftgrowthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedShiftedLeftGrowthCurve2DRibbonEndShape(shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) (ok bool) {

	return shiftedleftgrowthcurve2dribbonendshape.GongIsStaged(stage)
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedLeftGrowthCurve2DRibbonStartShapes[shiftedleftgrowthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedShiftedLeftGrowthCurve2DRibbonStartShape(shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) (ok bool) {

	return shiftedleftgrowthcurve2dribbonstartshape.GongIsStaged(stage)
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedLeftPartiallyGrowthCurve2DRibbons[shiftedleftpartiallygrowthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedShiftedLeftPartiallyGrowthCurve2DRibbon(shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) (ok bool) {

	return shiftedleftpartiallygrowthcurve2dribbon.GongIsStaged(stage)
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes[shiftedleftpartiallygrowthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) (ok bool) {

	return shiftedleftpartiallygrowthcurve2dribbonendshape.GongIsStaged(stage)
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes[shiftedleftpartiallygrowthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) (ok bool) {

	return shiftedleftpartiallygrowthcurve2dribbonstartshape.GongIsStaged(stage)
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedLeftStackGrowthCurveEndArcShapes[shiftedleftstackgrowthcurveendarcshape]

	return
}

func (stage *Stage) IsStagedShiftedLeftStackGrowthCurveEndArcShape(shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) (ok bool) {

	return shiftedleftstackgrowthcurveendarcshape.GongIsStaged(stage)
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedLeftStackGrowthCurveStartArcShapes[shiftedleftstackgrowthcurvestartarcshape]

	return
}

func (stage *Stage) IsStagedShiftedLeftStackGrowthCurveStartArcShape(shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) (ok bool) {

	return shiftedleftstackgrowthcurvestartarcshape.GongIsStaged(stage)
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedLeftStackNormalVectors[shiftedleftstacknormalvector]

	return
}

func (stage *Stage) IsStagedShiftedLeftStackNormalVector(shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) (ok bool) {

	return shiftedleftstacknormalvector.GongIsStaged(stage)
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedLeftStackOfGrowthCurves[shiftedleftstackofgrowthcurve]

	return
}

func (stage *Stage) IsStagedShiftedLeftStackOfGrowthCurve(shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) (ok bool) {

	return shiftedleftstackofgrowthcurve.GongIsStaged(stage)
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedLeftStackOfNormalVectors[shiftedleftstackofnormalvector]

	return
}

func (stage *Stage) IsStagedShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) (ok bool) {

	return shiftedleftstackofnormalvector.GongIsStaged(stage)
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedRightGrowthCurve2DRibbons[shiftedrightgrowthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedShiftedRightGrowthCurve2DRibbon(shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) (ok bool) {

	return shiftedrightgrowthcurve2dribbon.GongIsStaged(stage)
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedRightGrowthCurve2DRibbonEndShapes[shiftedrightgrowthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedShiftedRightGrowthCurve2DRibbonEndShape(shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) (ok bool) {

	return shiftedrightgrowthcurve2dribbonendshape.GongIsStaged(stage)
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ShiftedRightGrowthCurve2DRibbonStartShapes[shiftedrightgrowthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedShiftedRightGrowthCurve2DRibbonStartShape(shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) (ok bool) {

	return shiftedrightgrowthcurve2dribbonstartshape.GongIsStaged(stage)
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackGrowthCurve2DEndHalfwayArcShapes[stackgrowthcurve2dendhalfwayarcshape]

	return
}

func (stage *Stage) IsStagedStackGrowthCurve2DEndHalfwayArcShape(stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) (ok bool) {

	return stackgrowthcurve2dendhalfwayarcshape.GongIsStaged(stage)
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackGrowthCurve2DRibbonEndShapes[stackgrowthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedStackGrowthCurve2DRibbonEndShape(stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) (ok bool) {

	return stackgrowthcurve2dribbonendshape.GongIsStaged(stage)
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackGrowthCurve2DRibbonStartShapes[stackgrowthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedStackGrowthCurve2DRibbonStartShape(stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) (ok bool) {

	return stackgrowthcurve2dribbonstartshape.GongIsStaged(stage)
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackGrowthCurve2DStartHalfwayArcShapes[stackgrowthcurve2dstarthalfwayarcshape]

	return
}

func (stage *Stage) IsStagedStackGrowthCurve2DStartHalfwayArcShape(stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) (ok bool) {

	return stackgrowthcurve2dstarthalfwayarcshape.GongIsStaged(stage)
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackOfGrowthCurve2Ds[stackofgrowthcurve2d]

	return
}

func (stage *Stage) IsStagedStackOfGrowthCurve2D(stackofgrowthcurve2d *StackOfGrowthCurve2D) (ok bool) {

	return stackofgrowthcurve2d.GongIsStaged(stage)
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackOfGrowthCurve2DByGrowthVectors[stackofgrowthcurve2dbygrowthvector]

	return
}

func (stage *Stage) IsStagedStackOfGrowthCurve2DByGrowthVector(stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) (ok bool) {

	return stackofgrowthcurve2dbygrowthvector.GongIsStaged(stage)
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackOfGrowthCurve2DRibbons[stackofgrowthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedStackOfGrowthCurve2DRibbon(stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) (ok bool) {

	return stackofgrowthcurve2dribbon.GongIsStaged(stage)
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackOfPartiallyRotatedTorusShapes[stackofpartiallyrotatedtorusshape]

	return
}

func (stage *Stage) IsStagedStackOfPartiallyRotatedTorusShape(stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) (ok bool) {

	return stackofpartiallyrotatedtorusshape.GongIsStaged(stage)
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackOfRotatedGrowthCurve2Ds[stackofrotatedgrowthcurve2d]

	return
}

func (stage *Stage) IsStagedStackOfRotatedGrowthCurve2D(stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) (ok bool) {

	return stackofrotatedgrowthcurve2d.GongIsStaged(stage)
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackOfRotatedGrowthCurve2DRibbons[stackofrotatedgrowthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedStackOfRotatedGrowthCurve2DRibbon(stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) (ok bool) {

	return stackofrotatedgrowthcurve2dribbon.GongIsStaged(stage)
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackRotatedGrowthCurve2DEndArcShapes[stackrotatedgrowthcurve2dendarcshape]

	return
}

func (stage *Stage) IsStagedStackRotatedGrowthCurve2DEndArcShape(stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) (ok bool) {

	return stackrotatedgrowthcurve2dendarcshape.GongIsStaged(stage)
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackRotatedGrowthCurve2DRibbonEndShapes[stackrotatedgrowthcurve2dribbonendshape]

	return
}

func (stage *Stage) IsStagedStackRotatedGrowthCurve2DRibbonEndShape(stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) (ok bool) {

	return stackrotatedgrowthcurve2dribbonendshape.GongIsStaged(stage)
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackRotatedGrowthCurve2DRibbonStartShapes[stackrotatedgrowthcurve2dribbonstartshape]

	return
}

func (stage *Stage) IsStagedStackRotatedGrowthCurve2DRibbonStartShape(stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) (ok bool) {

	return stackrotatedgrowthcurve2dribbonstartshape.GongIsStaged(stage)
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StackRotatedGrowthCurve2DStartArcShapes[stackrotatedgrowthcurve2dstartarcshape]

	return
}

func (stage *Stage) IsStagedStackRotatedGrowthCurve2DStartArcShape(stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) (ok bool) {

	return stackrotatedgrowthcurve2dstartarcshape.GongIsStaged(stage)
}

func (startarcshape *StartArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StartArcShapes[startarcshape]

	return
}

func (stage *Stage) IsStagedStartArcShape(startarcshape *StartArcShape) (ok bool) {

	return startarcshape.GongIsStaged(stage)
}

func (startarcshapegrid *StartArcShapeGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StartArcShapeGrids[startarcshapegrid]

	return
}

func (stage *Stage) IsStagedStartArcShapeGrid(startarcshapegrid *StartArcShapeGrid) (ok bool) {

	return startarcshapegrid.GongIsStaged(stage)
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StartHalfwayArcShapes[starthalfwayarcshape]

	return
}

func (stage *Stage) IsStagedStartHalfwayArcShape(starthalfwayarcshape *StartHalfwayArcShape) (ok bool) {

	return starthalfwayarcshape.GongIsStaged(stage)
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StartHalfwayArcShapeGrids[starthalfwayarcshapegrid]

	return
}

func (stage *Stage) IsStagedStartHalfwayArcShapeGrid(starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) (ok bool) {

	return starthalfwayarcshapegrid.GongIsStaged(stage)
}

func (stemcylinder3dshape *StemCylinder3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StemCylinder3DShapes[stemcylinder3dshape]

	return
}

func (stage *Stage) IsStagedStemCylinder3DShape(stemcylinder3dshape *StemCylinder3DShape) (ok bool) {

	return stemcylinder3dshape.GongIsStaged(stage)
}

func (stool2ddiagram *Stool2DDiagram) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Stool2DDiagrams[stool2ddiagram]

	return
}

func (stage *Stage) IsStagedStool2DDiagram(stool2ddiagram *Stool2DDiagram) (ok bool) {

	return stool2ddiagram.GongIsStaged(stage)
}

func (stool3ddiagram *Stool3DDiagram) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Stool3DDiagrams[stool3ddiagram]

	return
}

func (stage *Stage) IsStagedStool3DDiagram(stool3ddiagram *Stool3DDiagram) (ok bool) {

	return stool3ddiagram.GongIsStaged(stage)
}

func (stoolabstract *StoolAbstract) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StoolAbstracts[stoolabstract]

	return
}

func (stage *Stage) IsStagedStoolAbstract(stoolabstract *StoolAbstract) (ok bool) {

	return stoolabstract.GongIsStaged(stage)
}

func (tiledfloor3dshape *TiledFloor3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TiledFloor3DShapes[tiledfloor3dshape]

	return
}

func (stage *Stage) IsStagedTiledFloor3DShape(tiledfloor3dshape *TiledFloor3DShape) (ok bool) {

	return tiledfloor3dshape.GongIsStaged(stage)
}

func (topendarcshape *TopEndArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopEndArcShapes[topendarcshape]

	return
}

func (stage *Stage) IsStagedTopEndArcShape(topendarcshape *TopEndArcShape) (ok bool) {

	return topendarcshape.GongIsStaged(stage)
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopEndArcShapeGrids[topendarcshapegrid]

	return
}

func (stage *Stage) IsStagedTopEndArcShapeGrid(topendarcshapegrid *TopEndArcShapeGrid) (ok bool) {

	return topendarcshapegrid.GongIsStaged(stage)
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopEndHalfwayArcShapes[topendhalfwayarcshape]

	return
}

func (stage *Stage) IsStagedTopEndHalfwayArcShape(topendhalfwayarcshape *TopEndHalfwayArcShape) (ok bool) {

	return topendhalfwayarcshape.GongIsStaged(stage)
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopEndHalfwayArcShapeGrids[topendhalfwayarcshapegrid]

	return
}

func (stage *Stage) IsStagedTopEndHalfwayArcShapeGrid(topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) (ok bool) {

	return topendhalfwayarcshapegrid.GongIsStaged(stage)
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopGrowthCurve2Ds[topgrowthcurve2d]

	return
}

func (stage *Stage) IsStagedTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D) (ok bool) {

	return topgrowthcurve2d.GongIsStaged(stage)
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopMidArcVectorShapes[topmidarcvectorshape]

	return
}

func (stage *Stage) IsStagedTopMidArcVectorShape(topmidarcvectorshape *TopMidArcVectorShape) (ok bool) {

	return topmidarcvectorshape.GongIsStaged(stage)
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopMidArcVectorShapeGrids[topmidarcvectorshapegrid]

	return
}

func (stage *Stage) IsStagedTopMidArcVectorShapeGrid(topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) (ok bool) {

	return topmidarcvectorshapegrid.GongIsStaged(stage)
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopStackGrowthCurve2DEndHalfwayArcShapes[topstackgrowthcurve2dendhalfwayarcshape]

	return
}

func (stage *Stage) IsStagedTopStackGrowthCurve2DEndHalfwayArcShape(topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) (ok bool) {

	return topstackgrowthcurve2dendhalfwayarcshape.GongIsStaged(stage)
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopStackGrowthCurve2DStartHalfwayArcShapes[topstackgrowthcurve2dstarthalfwayarcshape]

	return
}

func (stage *Stage) IsStagedTopStackGrowthCurve2DStartHalfwayArcShape(topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) (ok bool) {

	return topstackgrowthcurve2dstarthalfwayarcshape.GongIsStaged(stage)
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopStackOfGrowthCurve2Ds[topstackofgrowthcurve2d]

	return
}

func (stage *Stage) IsStagedTopStackOfGrowthCurve2D(topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) (ok bool) {

	return topstackofgrowthcurve2d.GongIsStaged(stage)
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopStackOfRotatedGrowthCurve2Ds[topstackofrotatedgrowthcurve2d]

	return
}

func (stage *Stage) IsStagedTopStackOfRotatedGrowthCurve2D(topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) (ok bool) {

	return topstackofrotatedgrowthcurve2d.GongIsStaged(stage)
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopStackOfRotatedGrowthCurve2DEndArcShapes[topstackofrotatedgrowthcurve2dendarcshape]

	return
}

func (stage *Stage) IsStagedTopStackOfRotatedGrowthCurve2DEndArcShape(topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) (ok bool) {

	return topstackofrotatedgrowthcurve2dendarcshape.GongIsStaged(stage)
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopStackOfRotatedGrowthCurve2DStartArcShapes[topstackofrotatedgrowthcurve2dstartarcshape]

	return
}

func (stage *Stage) IsStagedTopStackOfRotatedGrowthCurve2DStartArcShape(topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) (ok bool) {

	return topstackofrotatedgrowthcurve2dstartarcshape.GongIsStaged(stage)
}

func (topstartarcshape *TopStartArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopStartArcShapes[topstartarcshape]

	return
}

func (stage *Stage) IsStagedTopStartArcShape(topstartarcshape *TopStartArcShape) (ok bool) {

	return topstartarcshape.GongIsStaged(stage)
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopStartArcShapeGrids[topstartarcshapegrid]

	return
}

func (stage *Stage) IsStagedTopStartArcShapeGrid(topstartarcshapegrid *TopStartArcShapeGrid) (ok bool) {

	return topstartarcshapegrid.GongIsStaged(stage)
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopStartHalfwayArcShapes[topstarthalfwayarcshape]

	return
}

func (stage *Stage) IsStagedTopStartHalfwayArcShape(topstarthalfwayarcshape *TopStartHalfwayArcShape) (ok bool) {

	return topstarthalfwayarcshape.GongIsStaged(stage)
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TopStartHalfwayArcShapeGrids[topstarthalfwayarcshapegrid]

	return
}

func (stage *Stage) IsStagedTopStartHalfwayArcShapeGrid(topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) (ok bool) {

	return topstarthalfwayarcshapegrid.GongIsStaged(stage)
}

func (torus3dshape *Torus3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Torus3DShapes[torus3dshape]

	return
}

func (stage *Stage) IsStagedTorus3DShape(torus3dshape *Torus3DShape) (ok bool) {

	return torus3dshape.GongIsStaged(stage)
}

func (torusedge3dshape *TorusEdge3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TorusEdge3DShapes[torusedge3dshape]

	return
}

func (stage *Stage) IsStagedTorusEdge3DShape(torusedge3dshape *TorusEdge3DShape) (ok bool) {

	return torusedge3dshape.GongIsStaged(stage)
}

func (torusstackshape *TorusStackShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TorusStackShapes[torusstackshape]

	return
}

func (stage *Stage) IsStagedTorusStackShape(torusstackshape *TorusStackShape) (ok bool) {

	return torusstackshape.GongIsStaged(stage)
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TubeVase3DDiagrams[tubevase3ddiagram]

	return
}

func (stage *Stage) IsStagedTubeVase3DDiagram(tubevase3ddiagram *TubeVase3DDiagram) (ok bool) {

	return tubevase3ddiagram.GongIsStaged(stage)
}

func (tubevaseabstract *TubeVaseAbstract) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TubeVaseAbstracts[tubevaseabstract]

	return
}

func (stage *Stage) IsStagedTubeVaseAbstract(tubevaseabstract *TubeVaseAbstract) (ok bool) {

	return tubevaseabstract.GongIsStaged(stage)
}

func (vase2ddiagram *Vase2DDiagram) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Vase2DDiagrams[vase2ddiagram]

	return
}

func (stage *Stage) IsStagedVase2DDiagram(vase2ddiagram *Vase2DDiagram) (ok bool) {

	return vase2ddiagram.GongIsStaged(stage)
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.VerticalTorusStackShapes[verticaltorusstackshape]

	return
}

func (stage *Stage) IsStagedVerticalTorusStackShape(verticaltorusstackshape *VerticalTorusStackShape) (ok bool) {

	return verticaltorusstackshape.GongIsStaged(stage)
}

func (volumekey3dshape *VolumeKey3DShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.VolumeKey3DShapes[volumekey3dshape]

	return
}

func (stage *Stage) IsStagedVolumeKey3DShape(volumekey3dshape *VolumeKey3DShape) (ok bool) {

	return volumekey3dshape.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (angle0shape *Angle0Shape) GongStageBranch(stage *Stage) {
	stage.StageBranchAngle0Shape(angle0shape)
}

func (stage *Stage) StageBranchAngle0Shape(angle0shape *Angle0Shape) {

	// check if instance is already staged
	if stage.IsStaged(angle0shape) {
		return
	}

	angle0shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arcnormalvectorshape *ArcNormalVectorShape) GongStageBranch(stage *Stage) {
	stage.StageBranchArcNormalVectorShape(arcnormalvectorshape)
}

func (stage *Stage) StageBranchArcNormalVectorShape(arcnormalvectorshape *ArcNormalVectorShape) {

	// check if instance is already staged
	if stage.IsStaged(arcnormalvectorshape) {
		return
	}

	arcnormalvectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchArcNormalVectorShapeGrid(arcnormalvectorshapegrid)
}

func (stage *Stage) StageBranchArcNormalVectorShapeGrid(arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) {

	// check if instance is already staged
	if stage.IsStaged(arcnormalvectorshapegrid) {
		return
	}

	arcnormalvectorshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (axesshape *AxesShape) GongStageBranch(stage *Stage) {
	stage.StageBranchAxesShape(axesshape)
}

func (stage *Stage) StageBranchAxesShape(axesshape *AxesShape) {

	// check if instance is already staged
	if stage.IsStaged(axesshape) {
		return
	}

	axesshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (basevectorshape *BaseVectorShape) GongStageBranch(stage *Stage) {
	stage.StageBranchBaseVectorShape(basevectorshape)
}

func (stage *Stage) StageBranchBaseVectorShape(basevectorshape *BaseVectorShape) {

	// check if instance is already staged
	if stage.IsStaged(basevectorshape) {
		return
	}

	basevectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (basevectorshapegrid *BaseVectorShapeGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchBaseVectorShapeGrid(basevectorshapegrid)
}

func (stage *Stage) StageBranchBaseVectorShapeGrid(basevectorshapegrid *BaseVectorShapeGrid) {

	// check if instance is already staged
	if stage.IsStaged(basevectorshapegrid) {
		return
	}

	basevectorshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongStageBranch(stage *Stage) {
	stage.StageBranchChosenP1P2PairShape(chosenp1p2pairshape)
}

func (stage *Stage) StageBranchChosenP1P2PairShape(chosenp1p2pairshape *ChosenP1P2PairShape) {

	// check if instance is already staged
	if stage.IsStaged(chosenp1p2pairshape) {
		return
	}

	chosenp1p2pairshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circlegridshape *CircleGridShape) GongStageBranch(stage *Stage) {
	stage.StageBranchCircleGridShape(circlegridshape)
}

func (stage *Stage) StageBranchCircleGridShape(circlegridshape *CircleGridShape) {

	// check if instance is already staged
	if stage.IsStaged(circlegridshape) {
		return
	}

	circlegridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circumference3dshape *Circumference3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchCircumference3DShape(circumference3dshape)
}

func (stage *Stage) StageBranchCircumference3DShape(circumference3dshape *Circumference3DShape) {

	// check if instance is already staged
	if stage.IsStaged(circumference3dshape) {
		return
	}

	circumference3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clock2ddiagram *Clock2DDiagram) GongStageBranch(stage *Stage) {
	stage.StageBranchClock2DDiagram(clock2ddiagram)
}

func (stage *Stage) StageBranchClock2DDiagram(clock2ddiagram *Clock2DDiagram) {

	// check if instance is already staged
	if stage.IsStaged(clock2ddiagram) {
		return
	}

	clock2ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clock3ddiagram *Clock3DDiagram) GongStageBranch(stage *Stage) {
	stage.StageBranchClock3DDiagram(clock3ddiagram)
}

func (stage *Stage) StageBranchClock3DDiagram(clock3ddiagram *Clock3DDiagram) {

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

func (clockabstract *ClockAbstract) GongStageBranch(stage *Stage) {
	stage.StageBranchClockAbstract(clockabstract)
}

func (stage *Stage) StageBranchClockAbstract(clockabstract *ClockAbstract) {

	// check if instance is already staged
	if stage.IsStaged(clockabstract) {
		return
	}

	clockabstract.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clocktopcurveshape *ClockTopCurveShape) GongStageBranch(stage *Stage) {
	stage.StageBranchClockTopCurveShape(clocktopcurveshape)
}

func (stage *Stage) StageBranchClockTopCurveShape(clocktopcurveshape *ClockTopCurveShape) {

	// check if instance is already staged
	if stage.IsStaged(clocktopcurveshape) {
		return
	}

	clocktopcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cutline3dshape *CutLine3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchCutLine3DShape(cutline3dshape)
}

func (stage *Stage) StageBranchCutLine3DShape(cutline3dshape *CutLine3DShape) {

	// check if instance is already staged
	if stage.IsStaged(cutline3dshape) {
		return
	}

	cutline3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endarcshape *EndArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchEndArcShape(endarcshape)
}

func (stage *Stage) StageBranchEndArcShape(endarcshape *EndArcShape) {

	// check if instance is already staged
	if stage.IsStaged(endarcshape) {
		return
	}

	endarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endarcshapegrid *EndArcShapeGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchEndArcShapeGrid(endarcshapegrid)
}

func (stage *Stage) StageBranchEndArcShapeGrid(endarcshapegrid *EndArcShapeGrid) {

	// check if instance is already staged
	if stage.IsStaged(endarcshapegrid) {
		return
	}

	endarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endhalfwayarcshape *EndHalfwayArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchEndHalfwayArcShape(endhalfwayarcshape)
}

func (stage *Stage) StageBranchEndHalfwayArcShape(endhalfwayarcshape *EndHalfwayArcShape) {

	// check if instance is already staged
	if stage.IsStaged(endhalfwayarcshape) {
		return
	}

	endhalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchEndHalfwayArcShapeGrid(endhalfwayarcshapegrid)
}

func (stage *Stage) StageBranchEndHalfwayArcShapeGrid(endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) {

	// check if instance is already staged
	if stage.IsStaged(endhalfwayarcshapegrid) {
		return
	}

	endhalfwayarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (explanationtextshape *ExplanationTextShape) GongStageBranch(stage *Stage) {
	stage.StageBranchExplanationTextShape(explanationtextshape)
}

func (stage *Stage) StageBranchExplanationTextShape(explanationtextshape *ExplanationTextShape) {

	// check if instance is already staged
	if stage.IsStaged(explanationtextshape) {
		return
	}

	explanationtextshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eye3dshape *Eye3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchEye3DShape(eye3dshape)
}

func (stage *Stage) StageBranchEye3DShape(eye3dshape *Eye3DShape) {

	// check if instance is already staged
	if stage.IsStaged(eye3dshape) {
		return
	}

	eye3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchEyeCornersSampledPoints3DShape(eyecornerssampledpoints3dshape)
}

func (stage *Stage) StageBranchEyeCornersSampledPoints3DShape(eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) {

	// check if instance is already staged
	if stage.IsStaged(eyecornerssampledpoints3dshape) {
		return
	}

	eyecornerssampledpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchEyeSampledPoints3DShape(eyesampledpoints3dshape)
}

func (stage *Stage) StageBranchEyeSampledPoints3DShape(eyesampledpoints3dshape *EyeSampledPoints3DShape) {

	// check if instance is already staged
	if stage.IsStaged(eyesampledpoints3dshape) {
		return
	}

	eyesampledpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongStageBranch(stage *Stage) {
	stage.StageBranchEyeSeatBottomCurveShape(eyeseatbottomcurveshape)
}

func (stage *Stage) StageBranchEyeSeatBottomCurveShape(eyeseatbottomcurveshape *EyeSeatBottomCurveShape) {

	// check if instance is already staged
	if stage.IsStaged(eyeseatbottomcurveshape) {
		return
	}

	eyeseatbottomcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongStageBranch(stage *Stage) {
	stage.StageBranchEyeStoolBottomCurveShape(eyestoolbottomcurveshape)
}

func (stage *Stage) StageBranchEyeStoolBottomCurveShape(eyestoolbottomcurveshape *EyeStoolBottomCurveShape) {

	// check if instance is already staged
	if stage.IsStaged(eyestoolbottomcurveshape) {
		return
	}

	eyestoolbottomcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyevolume3dshape *EyeVolume3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchEyeVolume3DShape(eyevolume3dshape)
}

func (stage *Stage) StageBranchEyeVolume3DShape(eyevolume3dshape *EyeVolume3DShape) {

	// check if instance is already staged
	if stage.IsStaged(eyevolume3dshape) {
		return
	}

	eyevolume3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gridpathshape *GridPathShape) GongStageBranch(stage *Stage) {
	stage.StageBranchGridPathShape(gridpathshape)
}

func (stage *Stage) StageBranchGridPathShape(gridpathshape *GridPathShape) {

	// check if instance is already staged
	if stage.IsStaged(gridpathshape) {
		return
	}

	gridpathshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2d *GrowthCurve2D) GongStageBranch(stage *Stage) {
	stage.StageBranchGrowthCurve2D(growthcurve2d)
}

func (stage *Stage) StageBranchGrowthCurve2D(growthcurve2d *GrowthCurve2D) {

	// check if instance is already staged
	if stage.IsStaged(growthcurve2d) {
		return
	}

	growthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongStageBranch(stage *Stage) {
	stage.StageBranchGrowthCurve2DRibbon(growthcurve2dribbon)
}

func (stage *Stage) StageBranchGrowthCurve2DRibbon(growthcurve2dribbon *GrowthCurve2DRibbon) {

	// check if instance is already staged
	if stage.IsStaged(growthcurve2dribbon) {
		return
	}

	growthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {
	stage.StageBranchGrowthCurve2DRibbonEndShape(growthcurve2dribbonendshape)
}

func (stage *Stage) StageBranchGrowthCurve2DRibbonEndShape(growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if stage.IsStaged(growthcurve2dribbonendshape) {
		return
	}

	growthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {
	stage.StageBranchGrowthCurve2DRibbonStartShape(growthcurve2dribbonstartshape)
}

func (stage *Stage) StageBranchGrowthCurve2DRibbonStartShape(growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if stage.IsStaged(growthcurve2dribbonstartshape) {
		return
	}

	growthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongStageBranch(stage *Stage) {
	stage.StageBranchGrowthCurveRhombusGridShape(growthcurverhombusgridshape)
}

func (stage *Stage) StageBranchGrowthCurveRhombusGridShape(growthcurverhombusgridshape *GrowthCurveRhombusGridShape) {

	// check if instance is already staged
	if stage.IsStaged(growthcurverhombusgridshape) {
		return
	}

	growthcurverhombusgridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongStageBranch(stage *Stage) {
	stage.StageBranchGrowthCurveRhombusShape(growthcurverhombusshape)
}

func (stage *Stage) StageBranchGrowthCurveRhombusShape(growthcurverhombusshape *GrowthCurveRhombusShape) {

	// check if instance is already staged
	if stage.IsStaged(growthcurverhombusshape) {
		return
	}

	growthcurverhombusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthvectorshape *GrowthVectorShape) GongStageBranch(stage *Stage) {
	stage.StageBranchGrowthVectorShape(growthvectorshape)
}

func (stage *Stage) StageBranchGrowthVectorShape(growthvectorshape *GrowthVectorShape) {

	// check if instance is already staged
	if stage.IsStaged(growthvectorshape) {
		return
	}

	growthvectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (initialrhombusgridshape *InitialRhombusGridShape) GongStageBranch(stage *Stage) {
	stage.StageBranchInitialRhombusGridShape(initialrhombusgridshape)
}

func (stage *Stage) StageBranchInitialRhombusGridShape(initialrhombusgridshape *InitialRhombusGridShape) {

	// check if instance is already staged
	if stage.IsStaged(initialrhombusgridshape) {
		return
	}

	initialrhombusgridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (initialrhombusshape *InitialRhombusShape) GongStageBranch(stage *Stage) {
	stage.StageBranchInitialRhombusShape(initialrhombusshape)
}

func (stage *Stage) StageBranchInitialRhombusShape(initialrhombusshape *InitialRhombusShape) {

	// check if instance is already staged
	if stage.IsStaged(initialrhombusshape) {
		return
	}

	initialrhombusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (key3dshape *Key3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchKey3DShape(key3dshape)
}

func (stage *Stage) StageBranchKey3DShape(key3dshape *Key3DShape) {

	// check if instance is already staged
	if stage.IsStaged(key3dshape) {
		return
	}

	key3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (keyhole3dshape *KeyHole3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchKeyHole3DShape(keyhole3dshape)
}

func (stage *Stage) StageBranchKeyHole3DShape(keyhole3dshape *KeyHole3DShape) {

	// check if instance is already staged
	if stage.IsStaged(keyhole3dshape) {
		return
	}

	keyhole3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (keyholeshape *KeyHoleShape) GongStageBranch(stage *Stage) {
	stage.StageBranchKeyHoleShape(keyholeshape)
}

func (stage *Stage) StageBranchKeyHoleShape(keyholeshape *KeyHoleShape) {

	// check if instance is already staged
	if stage.IsStaged(keyholeshape) {
		return
	}

	keyholeshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (leaves3dshape *Leaves3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchLeaves3DShape(leaves3dshape)
}

func (stage *Stage) StageBranchLeaves3DShape(leaves3dshape *Leaves3DShape) {

	// check if instance is already staged
	if stage.IsStaged(leaves3dshape) {
		return
	}

	leaves3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongStageBranch(stage *Stage) {
	stage.StageBranchLibrary(library)
}

func (stage *Stage) StageBranchLibrary(library *Library) {

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
	stage.StageBranchMidArcVectorShape(midarcvectorshape)
}

func (stage *Stage) StageBranchMidArcVectorShape(midarcvectorshape *MidArcVectorShape) {

	// check if instance is already staged
	if stage.IsStaged(midarcvectorshape) {
		return
	}

	midarcvectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchMidArcVectorShapeGrid(midarcvectorshapegrid)
}

func (stage *Stage) StageBranchMidArcVectorShapeGrid(midarcvectorshapegrid *MidArcVectorShapeGrid) {

	// check if instance is already staged
	if stage.IsStaged(midarcvectorshapegrid) {
		return
	}

	midarcvectorshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (musicabstract *MusicAbstract) GongStageBranch(stage *Stage) {
	stage.StageBranchMusicAbstract(musicabstract)
}

func (stage *Stage) StageBranchMusicAbstract(musicabstract *MusicAbstract) {

	// check if instance is already staged
	if stage.IsStaged(musicabstract) {
		return
	}

	musicabstract.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (originalpoints3dshape *OriginalPoints3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchOriginalPoints3DShape(originalpoints3dshape)
}

func (stage *Stage) StageBranchOriginalPoints3DShape(originalpoints3dshape *OriginalPoints3DShape) {

	// check if instance is already staged
	if stage.IsStaged(originalpoints3dshape) {
		return
	}

	originalpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchParastichyMCurves3DShape(parastichymcurves3dshape)
}

func (stage *Stage) StageBranchParastichyMCurves3DShape(parastichymcurves3dshape *ParastichyMCurves3DShape) {

	// check if instance is already staged
	if stage.IsStaged(parastichymcurves3dshape) {
		return
	}

	parastichymcurves3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchParastichyNCurves3DShape(parastichyncurves3dshape)
}

func (stage *Stage) StageBranchParastichyNCurves3DShape(parastichyncurves3dshape *ParastichyNCurves3DShape) {

	// check if instance is already staged
	if stage.IsStaged(parastichyncurves3dshape) {
		return
	}

	parastichyncurves3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyGrowthCurve2DRibbon(partiallygrowthcurve2dribbon)
}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DRibbon(partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dribbon) {
		return
	}

	partiallygrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyGrowthCurve2DRibbonEndShape(partiallygrowthcurve2dribbonendshape)
}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DRibbonEndShape(partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dribbonendshape) {
		return
	}

	partiallygrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyGrowthCurve2DRibbonStartShape(partiallygrowthcurve2dribbonstartshape)
}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DRibbonStartShape(partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dribbonstartshape) {
		return
	}

	partiallygrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyGrowthCurve2DTrajectory(partiallygrowthcurve2dtrajectory)
}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectory(partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectory) {
		return
	}

	partiallygrowthcurve2dtrajectory.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyGrowthCurve2DTrajectoryP1CurveShape(partiallygrowthcurve2dtrajectoryp1curveshape)
}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryP1CurveShape(partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryp1curveshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1curveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyGrowthCurve2DTrajectoryP1P2(partiallygrowthcurve2dtrajectoryp1p2)
}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryP1P2(partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryp1p2) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(partiallygrowthcurve2dtrajectoryp1p2pairlineshape)
}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryp1p2pairlineshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2pairlineshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyGrowthCurve2DTrajectoryP1PointShape(partiallygrowthcurve2dtrajectoryp1pointshape)
}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryP1PointShape(partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryp1pointshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1pointshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyGrowthCurve2DTrajectoryP2CurveShape(partiallygrowthcurve2dtrajectoryp2curveshape)
}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryP2CurveShape(partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryp2curveshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp2curveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyGrowthCurve2DTrajectoryP2PointShape(partiallygrowthcurve2dtrajectoryp2pointshape)
}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryP2PointShape(partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryp2pointshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp2pointshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyGrowthCurve2DTrajectoryShape(partiallygrowthcurve2dtrajectoryshape)
}

func (stage *Stage) StageBranchPartiallyGrowthCurve2DTrajectoryShape(partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) {

	// check if instance is already staged
	if stage.IsStaged(partiallygrowthcurve2dtrajectoryshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyRotatedSeatBottomCurveShape(partiallyrotatedseatbottomcurveshape)
}

func (stage *Stage) StageBranchPartiallyRotatedSeatBottomCurveShape(partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) {

	// check if instance is already staged
	if stage.IsStaged(partiallyrotatedseatbottomcurveshape) {
		return
	}

	partiallyrotatedseatbottomcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyRotatedSeatTopCurveShape(partiallyrotatedseattopcurveshape)
}

func (stage *Stage) StageBranchPartiallyRotatedSeatTopCurveShape(partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) {

	// check if instance is already staged
	if stage.IsStaged(partiallyrotatedseattopcurveshape) {
		return
	}

	partiallyrotatedseattopcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPartiallyRotatedTorusShape(partiallyrotatedtorusshape)
}

func (stage *Stage) StageBranchPartiallyRotatedTorusShape(partiallyrotatedtorusshape *PartiallyRotatedTorusShape) {

	// check if instance is already staged
	if stage.IsStaged(partiallyrotatedtorusshape) {
		return
	}

	partiallyrotatedtorusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvector *PerpendicularVector) GongStageBranch(stage *Stage) {
	stage.StageBranchPerpendicularVector(perpendicularvector)
}

func (stage *Stage) StageBranchPerpendicularVector(perpendicularvector *PerpendicularVector) {

	// check if instance is already staged
	if stage.IsStaged(perpendicularvector) {
		return
	}

	perpendicularvector.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchPerpendicularVectorGrid(perpendicularvectorgrid)
}

func (stage *Stage) StageBranchPerpendicularVectorGrid(perpendicularvectorgrid *PerpendicularVectorGrid) {

	// check if instance is already staged
	if stage.IsStaged(perpendicularvectorgrid) {
		return
	}

	perpendicularvectorgrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongStageBranch(stage *Stage) {
	stage.StageBranchPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway)
}

func (stage *Stage) StageBranchPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) {

	// check if instance is already staged
	if stage.IsStaged(perpendicularvectorgridhalfway) {
		return
	}

	perpendicularvectorgridhalfway.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongStageBranch(stage *Stage) {
	stage.StageBranchPerpendicularVectorHalfway(perpendicularvectorhalfway)
}

func (stage *Stage) StageBranchPerpendicularVectorHalfway(perpendicularvectorhalfway *PerpendicularVectorHalfway) {

	// check if instance is already staged
	if stage.IsStaged(perpendicularvectorhalfway) {
		return
	}

	perpendicularvectorhalfway.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (plant2ddiagram *Plant2DDiagram) GongStageBranch(stage *Stage) {
	stage.StageBranchPlant2DDiagram(plant2ddiagram)
}

func (stage *Stage) StageBranchPlant2DDiagram(plant2ddiagram *Plant2DDiagram) {

	// check if instance is already staged
	if stage.IsStaged(plant2ddiagram) {
		return
	}

	plant2ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (plant3ddiagram *Plant3DDiagram) GongStageBranch(stage *Stage) {
	stage.StageBranchPlant3DDiagram(plant3ddiagram)
}

func (stage *Stage) StageBranchPlant3DDiagram(plant3ddiagram *Plant3DDiagram) {

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
	stage.StageBranchPlantAbstract(plantabstract)
}

func (stage *Stage) StageBranchPlantAbstract(plantabstract *PlantAbstract) {

	// check if instance is already staged
	if stage.IsStaged(plantabstract) {
		return
	}

	plantabstract.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plantabstract.TubeVaseAbstract != nil {
		stage.StageBranch(plantabstract.TubeVaseAbstract)
	}
	if plantabstract.StoolAbstract != nil {
		stage.StageBranch(plantabstract.StoolAbstract)
	}
	if plantabstract.ClockAbstract != nil {
		stage.StageBranch(plantabstract.ClockAbstract)
	}
	if plantabstract.MusicAbstract != nil {
		stage.StageBranch(plantabstract.MusicAbstract)
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
	stage.StageBranchPlantCircumferenceShape(plantcircumferenceshape)
}

func (stage *Stage) StageBranchPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape) {

	// check if instance is already staged
	if stage.IsStaged(plantcircumferenceshape) {
		return
	}

	plantcircumferenceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointsandlines3dshape *PointsAndLines3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPointsAndLines3DShape(pointsandlines3dshape)
}

func (stage *Stage) StageBranchPointsAndLines3DShape(pointsandlines3dshape *PointsAndLines3DShape) {

	// check if instance is already staged
	if stage.IsStaged(pointsandlines3dshape) {
		return
	}

	pointsandlines3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pxshape *PxShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPxShape(pxshape)
}

func (stage *Stage) StageBranchPxShape(pxshape *PxShape) {

	// check if instance is already staged
	if stage.IsStaged(pxshape) {
		return
	}

	pxshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rendered3dshape *Rendered3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchRendered3DShape(rendered3dshape)
}

func (stage *Stage) StageBranchRendered3DShape(rendered3dshape *Rendered3DShape) {

	// check if instance is already staged
	if stage.IsStaged(rendered3dshape) {
		return
	}

	rendered3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rhombusshape *RhombusShape) GongStageBranch(stage *Stage) {
	stage.StageBranchRhombusShape(rhombusshape)
}

func (stage *Stage) StageBranchRhombusShape(rhombusshape *RhombusShape) {

	// check if instance is already staged
	if stage.IsStaged(rhombusshape) {
		return
	}

	rhombusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rhombusstuff *RhombusStuff) GongStageBranch(stage *Stage) {
	stage.StageBranchRhombusStuff(rhombusstuff)
}

func (stage *Stage) StageBranchRhombusStuff(rhombusstuff *RhombusStuff) {

	// check if instance is already staged
	if stage.IsStaged(rhombusstuff) {
		return
	}

	rhombusstuff.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongStageBranch(stage *Stage) {
	stage.StageBranchRotatedRhombusGridShape(rotatedrhombusgridshape)
}

func (stage *Stage) StageBranchRotatedRhombusGridShape(rotatedrhombusgridshape *RotatedRhombusGridShape) {

	// check if instance is already staged
	if stage.IsStaged(rotatedrhombusgridshape) {
		return
	}

	rotatedrhombusgridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedrhombusshape *RotatedRhombusShape) GongStageBranch(stage *Stage) {
	stage.StageBranchRotatedRhombusShape(rotatedrhombusshape)
}

func (stage *Stage) StageBranchRotatedRhombusShape(rotatedrhombusshape *RotatedRhombusShape) {

	// check if instance is already staged
	if stage.IsStaged(rotatedrhombusshape) {
		return
	}

	rotatedrhombusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchRotatedSampledPoints3DShape(rotatedsampledpoints3dshape)
}

func (stage *Stage) StageBranchRotatedSampledPoints3DShape(rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) {

	// check if instance is already staged
	if stage.IsStaged(rotatedsampledpoints3dshape) {
		return
	}

	rotatedsampledpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchRotatedSeatAndLegs3DShape(rotatedseatandlegs3dshape)
}

func (stage *Stage) StageBranchRotatedSeatAndLegs3DShape(rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) {

	// check if instance is already staged
	if stage.IsStaged(rotatedseatandlegs3dshape) {
		return
	}

	rotatedseatandlegs3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (sampledpoints3dshape *SampledPoints3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchSampledPoints3DShape(sampledpoints3dshape)
}

func (stage *Stage) StageBranchSampledPoints3DShape(sampledpoints3dshape *SampledPoints3DShape) {

	// check if instance is already staged
	if stage.IsStaged(sampledpoints3dshape) {
		return
	}

	sampledpoints3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seat3dshape *Seat3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchSeat3DShape(seat3dshape)
}

func (stage *Stage) StageBranchSeat3DShape(seat3dshape *Seat3DShape) {

	// check if instance is already staged
	if stage.IsStaged(seat3dshape) {
		return
	}

	seat3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchSeatAndLegs3DShape(seatandlegs3dshape)
}

func (stage *Stage) StageBranchSeatAndLegs3DShape(seatandlegs3dshape *SeatAndLegs3DShape) {

	// check if instance is already staged
	if stage.IsStaged(seatandlegs3dshape) {
		return
	}

	seatandlegs3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seatbottomcurveshape *SeatBottomCurveShape) GongStageBranch(stage *Stage) {
	stage.StageBranchSeatBottomCurveShape(seatbottomcurveshape)
}

func (stage *Stage) StageBranchSeatBottomCurveShape(seatbottomcurveshape *SeatBottomCurveShape) {

	// check if instance is already staged
	if stage.IsStaged(seatbottomcurveshape) {
		return
	}

	seatbottomcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seattopcurveshape *SeatTopCurveShape) GongStageBranch(stage *Stage) {
	stage.StageBranchSeatTopCurveShape(seattopcurveshape)
}

func (stage *Stage) StageBranchSeatTopCurveShape(seattopcurveshape *SeatTopCurveShape) {

	// check if instance is already staged
	if stage.IsStaged(seattopcurveshape) {
		return
	}

	seattopcurveshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedBottomTopStartArcShape(shiftedbottomtopstartarcshape)
}

func (stage *Stage) StageBranchShiftedBottomTopStartArcShape(shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) {

	// check if instance is already staged
	if stage.IsStaged(shiftedbottomtopstartarcshape) {
		return
	}

	shiftedbottomtopstartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedBottomTopStartArcShapeGrid(shiftedbottomtopstartarcshapegrid)
}

func (stage *Stage) StageBranchShiftedBottomTopStartArcShapeGrid(shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) {

	// check if instance is already staged
	if stage.IsStaged(shiftedbottomtopstartarcshapegrid) {
		return
	}

	shiftedbottomtopstartarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedLeftGrowthCurve2DRibbon(shiftedleftgrowthcurve2dribbon)
}

func (stage *Stage) StageBranchShiftedLeftGrowthCurve2DRibbon(shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftgrowthcurve2dribbon) {
		return
	}

	shiftedleftgrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedLeftGrowthCurve2DRibbonEndShape(shiftedleftgrowthcurve2dribbonendshape)
}

func (stage *Stage) StageBranchShiftedLeftGrowthCurve2DRibbonEndShape(shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftgrowthcurve2dribbonendshape) {
		return
	}

	shiftedleftgrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedLeftGrowthCurve2DRibbonStartShape(shiftedleftgrowthcurve2dribbonstartshape)
}

func (stage *Stage) StageBranchShiftedLeftGrowthCurve2DRibbonStartShape(shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftgrowthcurve2dribbonstartshape) {
		return
	}

	shiftedleftgrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedLeftPartiallyGrowthCurve2DRibbon(shiftedleftpartiallygrowthcurve2dribbon)
}

func (stage *Stage) StageBranchShiftedLeftPartiallyGrowthCurve2DRibbon(shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftpartiallygrowthcurve2dribbon) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(shiftedleftpartiallygrowthcurve2dribbonendshape)
}

func (stage *Stage) StageBranchShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftpartiallygrowthcurve2dribbonendshape) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(shiftedleftpartiallygrowthcurve2dribbonstartshape)
}

func (stage *Stage) StageBranchShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftpartiallygrowthcurve2dribbonstartshape) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedLeftStackGrowthCurveEndArcShape(shiftedleftstackgrowthcurveendarcshape)
}

func (stage *Stage) StageBranchShiftedLeftStackGrowthCurveEndArcShape(shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftstackgrowthcurveendarcshape) {
		return
	}

	shiftedleftstackgrowthcurveendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedLeftStackGrowthCurveStartArcShape(shiftedleftstackgrowthcurvestartarcshape)
}

func (stage *Stage) StageBranchShiftedLeftStackGrowthCurveStartArcShape(shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftstackgrowthcurvestartarcshape) {
		return
	}

	shiftedleftstackgrowthcurvestartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedLeftStackNormalVector(shiftedleftstacknormalvector)
}

func (stage *Stage) StageBranchShiftedLeftStackNormalVector(shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftstacknormalvector) {
		return
	}

	shiftedleftstacknormalvector.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedLeftStackOfGrowthCurve(shiftedleftstackofgrowthcurve)
}

func (stage *Stage) StageBranchShiftedLeftStackOfGrowthCurve(shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftstackofgrowthcurve) {
		return
	}

	shiftedleftstackofgrowthcurve.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector)
}

func (stage *Stage) StageBranchShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) {

	// check if instance is already staged
	if stage.IsStaged(shiftedleftstackofnormalvector) {
		return
	}

	shiftedleftstackofnormalvector.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedRightGrowthCurve2DRibbon(shiftedrightgrowthcurve2dribbon)
}

func (stage *Stage) StageBranchShiftedRightGrowthCurve2DRibbon(shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) {

	// check if instance is already staged
	if stage.IsStaged(shiftedrightgrowthcurve2dribbon) {
		return
	}

	shiftedrightgrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedRightGrowthCurve2DRibbonEndShape(shiftedrightgrowthcurve2dribbonendshape)
}

func (stage *Stage) StageBranchShiftedRightGrowthCurve2DRibbonEndShape(shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if stage.IsStaged(shiftedrightgrowthcurve2dribbonendshape) {
		return
	}

	shiftedrightgrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {
	stage.StageBranchShiftedRightGrowthCurve2DRibbonStartShape(shiftedrightgrowthcurve2dribbonstartshape)
}

func (stage *Stage) StageBranchShiftedRightGrowthCurve2DRibbonStartShape(shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if stage.IsStaged(shiftedrightgrowthcurve2dribbonstartshape) {
		return
	}

	shiftedrightgrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchStackGrowthCurve2DEndHalfwayArcShape(stackgrowthcurve2dendhalfwayarcshape)
}

func (stage *Stage) StageBranchStackGrowthCurve2DEndHalfwayArcShape(stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) {

	// check if instance is already staged
	if stage.IsStaged(stackgrowthcurve2dendhalfwayarcshape) {
		return
	}

	stackgrowthcurve2dendhalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {
	stage.StageBranchStackGrowthCurve2DRibbonEndShape(stackgrowthcurve2dribbonendshape)
}

func (stage *Stage) StageBranchStackGrowthCurve2DRibbonEndShape(stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if stage.IsStaged(stackgrowthcurve2dribbonendshape) {
		return
	}

	stackgrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {
	stage.StageBranchStackGrowthCurve2DRibbonStartShape(stackgrowthcurve2dribbonstartshape)
}

func (stage *Stage) StageBranchStackGrowthCurve2DRibbonStartShape(stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if stage.IsStaged(stackgrowthcurve2dribbonstartshape) {
		return
	}

	stackgrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchStackGrowthCurve2DStartHalfwayArcShape(stackgrowthcurve2dstarthalfwayarcshape)
}

func (stage *Stage) StageBranchStackGrowthCurve2DStartHalfwayArcShape(stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) {

	// check if instance is already staged
	if stage.IsStaged(stackgrowthcurve2dstarthalfwayarcshape) {
		return
	}

	stackgrowthcurve2dstarthalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongStageBranch(stage *Stage) {
	stage.StageBranchStackOfGrowthCurve2D(stackofgrowthcurve2d)
}

func (stage *Stage) StageBranchStackOfGrowthCurve2D(stackofgrowthcurve2d *StackOfGrowthCurve2D) {

	// check if instance is already staged
	if stage.IsStaged(stackofgrowthcurve2d) {
		return
	}

	stackofgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongStageBranch(stage *Stage) {
	stage.StageBranchStackOfGrowthCurve2DByGrowthVector(stackofgrowthcurve2dbygrowthvector)
}

func (stage *Stage) StageBranchStackOfGrowthCurve2DByGrowthVector(stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) {

	// check if instance is already staged
	if stage.IsStaged(stackofgrowthcurve2dbygrowthvector) {
		return
	}

	stackofgrowthcurve2dbygrowthvector.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongStageBranch(stage *Stage) {
	stage.StageBranchStackOfGrowthCurve2DRibbon(stackofgrowthcurve2dribbon)
}

func (stage *Stage) StageBranchStackOfGrowthCurve2DRibbon(stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) {

	// check if instance is already staged
	if stage.IsStaged(stackofgrowthcurve2dribbon) {
		return
	}

	stackofgrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongStageBranch(stage *Stage) {
	stage.StageBranchStackOfPartiallyRotatedTorusShape(stackofpartiallyrotatedtorusshape)
}

func (stage *Stage) StageBranchStackOfPartiallyRotatedTorusShape(stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) {

	// check if instance is already staged
	if stage.IsStaged(stackofpartiallyrotatedtorusshape) {
		return
	}

	stackofpartiallyrotatedtorusshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongStageBranch(stage *Stage) {
	stage.StageBranchStackOfRotatedGrowthCurve2D(stackofrotatedgrowthcurve2d)
}

func (stage *Stage) StageBranchStackOfRotatedGrowthCurve2D(stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) {

	// check if instance is already staged
	if stage.IsStaged(stackofrotatedgrowthcurve2d) {
		return
	}

	stackofrotatedgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongStageBranch(stage *Stage) {
	stage.StageBranchStackOfRotatedGrowthCurve2DRibbon(stackofrotatedgrowthcurve2dribbon)
}

func (stage *Stage) StageBranchStackOfRotatedGrowthCurve2DRibbon(stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) {

	// check if instance is already staged
	if stage.IsStaged(stackofrotatedgrowthcurve2dribbon) {
		return
	}

	stackofrotatedgrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchStackRotatedGrowthCurve2DEndArcShape(stackrotatedgrowthcurve2dendarcshape)
}

func (stage *Stage) StageBranchStackRotatedGrowthCurve2DEndArcShape(stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) {

	// check if instance is already staged
	if stage.IsStaged(stackrotatedgrowthcurve2dendarcshape) {
		return
	}

	stackrotatedgrowthcurve2dendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongStageBranch(stage *Stage) {
	stage.StageBranchStackRotatedGrowthCurve2DRibbonEndShape(stackrotatedgrowthcurve2dribbonendshape)
}

func (stage *Stage) StageBranchStackRotatedGrowthCurve2DRibbonEndShape(stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if stage.IsStaged(stackrotatedgrowthcurve2dribbonendshape) {
		return
	}

	stackrotatedgrowthcurve2dribbonendshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongStageBranch(stage *Stage) {
	stage.StageBranchStackRotatedGrowthCurve2DRibbonStartShape(stackrotatedgrowthcurve2dribbonstartshape)
}

func (stage *Stage) StageBranchStackRotatedGrowthCurve2DRibbonStartShape(stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if stage.IsStaged(stackrotatedgrowthcurve2dribbonstartshape) {
		return
	}

	stackrotatedgrowthcurve2dribbonstartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchStackRotatedGrowthCurve2DStartArcShape(stackrotatedgrowthcurve2dstartarcshape)
}

func (stage *Stage) StageBranchStackRotatedGrowthCurve2DStartArcShape(stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) {

	// check if instance is already staged
	if stage.IsStaged(stackrotatedgrowthcurve2dstartarcshape) {
		return
	}

	stackrotatedgrowthcurve2dstartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (startarcshape *StartArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchStartArcShape(startarcshape)
}

func (stage *Stage) StageBranchStartArcShape(startarcshape *StartArcShape) {

	// check if instance is already staged
	if stage.IsStaged(startarcshape) {
		return
	}

	startarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (startarcshapegrid *StartArcShapeGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchStartArcShapeGrid(startarcshapegrid)
}

func (stage *Stage) StageBranchStartArcShapeGrid(startarcshapegrid *StartArcShapeGrid) {

	// check if instance is already staged
	if stage.IsStaged(startarcshapegrid) {
		return
	}

	startarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (starthalfwayarcshape *StartHalfwayArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchStartHalfwayArcShape(starthalfwayarcshape)
}

func (stage *Stage) StageBranchStartHalfwayArcShape(starthalfwayarcshape *StartHalfwayArcShape) {

	// check if instance is already staged
	if stage.IsStaged(starthalfwayarcshape) {
		return
	}

	starthalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchStartHalfwayArcShapeGrid(starthalfwayarcshapegrid)
}

func (stage *Stage) StageBranchStartHalfwayArcShapeGrid(starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) {

	// check if instance is already staged
	if stage.IsStaged(starthalfwayarcshapegrid) {
		return
	}

	starthalfwayarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stemcylinder3dshape *StemCylinder3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchStemCylinder3DShape(stemcylinder3dshape)
}

func (stage *Stage) StageBranchStemCylinder3DShape(stemcylinder3dshape *StemCylinder3DShape) {

	// check if instance is already staged
	if stage.IsStaged(stemcylinder3dshape) {
		return
	}

	stemcylinder3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stool2ddiagram *Stool2DDiagram) GongStageBranch(stage *Stage) {
	stage.StageBranchStool2DDiagram(stool2ddiagram)
}

func (stage *Stage) StageBranchStool2DDiagram(stool2ddiagram *Stool2DDiagram) {

	// check if instance is already staged
	if stage.IsStaged(stool2ddiagram) {
		return
	}

	stool2ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stool3ddiagram *Stool3DDiagram) GongStageBranch(stage *Stage) {
	stage.StageBranchStool3DDiagram(stool3ddiagram)
}

func (stage *Stage) StageBranchStool3DDiagram(stool3ddiagram *Stool3DDiagram) {

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

func (stoolabstract *StoolAbstract) GongStageBranch(stage *Stage) {
	stage.StageBranchStoolAbstract(stoolabstract)
}

func (stage *Stage) StageBranchStoolAbstract(stoolabstract *StoolAbstract) {

	// check if instance is already staged
	if stage.IsStaged(stoolabstract) {
		return
	}

	stoolabstract.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tiledfloor3dshape *TiledFloor3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTiledFloor3DShape(tiledfloor3dshape)
}

func (stage *Stage) StageBranchTiledFloor3DShape(tiledfloor3dshape *TiledFloor3DShape) {

	// check if instance is already staged
	if stage.IsStaged(tiledfloor3dshape) {
		return
	}

	tiledfloor3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendarcshape *TopEndArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTopEndArcShape(topendarcshape)
}

func (stage *Stage) StageBranchTopEndArcShape(topendarcshape *TopEndArcShape) {

	// check if instance is already staged
	if stage.IsStaged(topendarcshape) {
		return
	}

	topendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendarcshapegrid *TopEndArcShapeGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchTopEndArcShapeGrid(topendarcshapegrid)
}

func (stage *Stage) StageBranchTopEndArcShapeGrid(topendarcshapegrid *TopEndArcShapeGrid) {

	// check if instance is already staged
	if stage.IsStaged(topendarcshapegrid) {
		return
	}

	topendarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTopEndHalfwayArcShape(topendhalfwayarcshape)
}

func (stage *Stage) StageBranchTopEndHalfwayArcShape(topendhalfwayarcshape *TopEndHalfwayArcShape) {

	// check if instance is already staged
	if stage.IsStaged(topendhalfwayarcshape) {
		return
	}

	topendhalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchTopEndHalfwayArcShapeGrid(topendhalfwayarcshapegrid)
}

func (stage *Stage) StageBranchTopEndHalfwayArcShapeGrid(topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) {

	// check if instance is already staged
	if stage.IsStaged(topendhalfwayarcshapegrid) {
		return
	}

	topendhalfwayarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topgrowthcurve2d *TopGrowthCurve2D) GongStageBranch(stage *Stage) {
	stage.StageBranchTopGrowthCurve2D(topgrowthcurve2d)
}

func (stage *Stage) StageBranchTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D) {

	// check if instance is already staged
	if stage.IsStaged(topgrowthcurve2d) {
		return
	}

	topgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topmidarcvectorshape *TopMidArcVectorShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTopMidArcVectorShape(topmidarcvectorshape)
}

func (stage *Stage) StageBranchTopMidArcVectorShape(topmidarcvectorshape *TopMidArcVectorShape) {

	// check if instance is already staged
	if stage.IsStaged(topmidarcvectorshape) {
		return
	}

	topmidarcvectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchTopMidArcVectorShapeGrid(topmidarcvectorshapegrid)
}

func (stage *Stage) StageBranchTopMidArcVectorShapeGrid(topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) {

	// check if instance is already staged
	if stage.IsStaged(topmidarcvectorshapegrid) {
		return
	}

	topmidarcvectorshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTopStackGrowthCurve2DEndHalfwayArcShape(topstackgrowthcurve2dendhalfwayarcshape)
}

func (stage *Stage) StageBranchTopStackGrowthCurve2DEndHalfwayArcShape(topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) {

	// check if instance is already staged
	if stage.IsStaged(topstackgrowthcurve2dendhalfwayarcshape) {
		return
	}

	topstackgrowthcurve2dendhalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTopStackGrowthCurve2DStartHalfwayArcShape(topstackgrowthcurve2dstarthalfwayarcshape)
}

func (stage *Stage) StageBranchTopStackGrowthCurve2DStartHalfwayArcShape(topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) {

	// check if instance is already staged
	if stage.IsStaged(topstackgrowthcurve2dstarthalfwayarcshape) {
		return
	}

	topstackgrowthcurve2dstarthalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongStageBranch(stage *Stage) {
	stage.StageBranchTopStackOfGrowthCurve2D(topstackofgrowthcurve2d)
}

func (stage *Stage) StageBranchTopStackOfGrowthCurve2D(topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) {

	// check if instance is already staged
	if stage.IsStaged(topstackofgrowthcurve2d) {
		return
	}

	topstackofgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongStageBranch(stage *Stage) {
	stage.StageBranchTopStackOfRotatedGrowthCurve2D(topstackofrotatedgrowthcurve2d)
}

func (stage *Stage) StageBranchTopStackOfRotatedGrowthCurve2D(topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) {

	// check if instance is already staged
	if stage.IsStaged(topstackofrotatedgrowthcurve2d) {
		return
	}

	topstackofrotatedgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTopStackOfRotatedGrowthCurve2DEndArcShape(topstackofrotatedgrowthcurve2dendarcshape)
}

func (stage *Stage) StageBranchTopStackOfRotatedGrowthCurve2DEndArcShape(topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) {

	// check if instance is already staged
	if stage.IsStaged(topstackofrotatedgrowthcurve2dendarcshape) {
		return
	}

	topstackofrotatedgrowthcurve2dendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTopStackOfRotatedGrowthCurve2DStartArcShape(topstackofrotatedgrowthcurve2dstartarcshape)
}

func (stage *Stage) StageBranchTopStackOfRotatedGrowthCurve2DStartArcShape(topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) {

	// check if instance is already staged
	if stage.IsStaged(topstackofrotatedgrowthcurve2dstartarcshape) {
		return
	}

	topstackofrotatedgrowthcurve2dstartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstartarcshape *TopStartArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTopStartArcShape(topstartarcshape)
}

func (stage *Stage) StageBranchTopStartArcShape(topstartarcshape *TopStartArcShape) {

	// check if instance is already staged
	if stage.IsStaged(topstartarcshape) {
		return
	}

	topstartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchTopStartArcShapeGrid(topstartarcshapegrid)
}

func (stage *Stage) StageBranchTopStartArcShapeGrid(topstartarcshapegrid *TopStartArcShapeGrid) {

	// check if instance is already staged
	if stage.IsStaged(topstartarcshapegrid) {
		return
	}

	topstartarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTopStartHalfwayArcShape(topstarthalfwayarcshape)
}

func (stage *Stage) StageBranchTopStartHalfwayArcShape(topstarthalfwayarcshape *TopStartHalfwayArcShape) {

	// check if instance is already staged
	if stage.IsStaged(topstarthalfwayarcshape) {
		return
	}

	topstarthalfwayarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongStageBranch(stage *Stage) {
	stage.StageBranchTopStartHalfwayArcShapeGrid(topstarthalfwayarcshapegrid)
}

func (stage *Stage) StageBranchTopStartHalfwayArcShapeGrid(topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) {

	// check if instance is already staged
	if stage.IsStaged(topstarthalfwayarcshapegrid) {
		return
	}

	topstarthalfwayarcshapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torus3dshape *Torus3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTorus3DShape(torus3dshape)
}

func (stage *Stage) StageBranchTorus3DShape(torus3dshape *Torus3DShape) {

	// check if instance is already staged
	if stage.IsStaged(torus3dshape) {
		return
	}

	torus3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torusedge3dshape *TorusEdge3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTorusEdge3DShape(torusedge3dshape)
}

func (stage *Stage) StageBranchTorusEdge3DShape(torusedge3dshape *TorusEdge3DShape) {

	// check if instance is already staged
	if stage.IsStaged(torusedge3dshape) {
		return
	}

	torusedge3dshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torusstackshape *TorusStackShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTorusStackShape(torusstackshape)
}

func (stage *Stage) StageBranchTorusStackShape(torusstackshape *TorusStackShape) {

	// check if instance is already staged
	if stage.IsStaged(torusstackshape) {
		return
	}

	torusstackshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tubevase3ddiagram *TubeVase3DDiagram) GongStageBranch(stage *Stage) {
	stage.StageBranchTubeVase3DDiagram(tubevase3ddiagram)
}

func (stage *Stage) StageBranchTubeVase3DDiagram(tubevase3ddiagram *TubeVase3DDiagram) {

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

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tubevaseabstract *TubeVaseAbstract) GongStageBranch(stage *Stage) {
	stage.StageBranchTubeVaseAbstract(tubevaseabstract)
}

func (stage *Stage) StageBranchTubeVaseAbstract(tubevaseabstract *TubeVaseAbstract) {

	// check if instance is already staged
	if stage.IsStaged(tubevaseabstract) {
		return
	}

	tubevaseabstract.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vase2ddiagram *Vase2DDiagram) GongStageBranch(stage *Stage) {
	stage.StageBranchVase2DDiagram(vase2ddiagram)
}

func (stage *Stage) StageBranchVase2DDiagram(vase2ddiagram *Vase2DDiagram) {

	// check if instance is already staged
	if stage.IsStaged(vase2ddiagram) {
		return
	}

	vase2ddiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticaltorusstackshape *VerticalTorusStackShape) GongStageBranch(stage *Stage) {
	stage.StageBranchVerticalTorusStackShape(verticaltorusstackshape)
}

func (stage *Stage) StageBranchVerticalTorusStackShape(verticaltorusstackshape *VerticalTorusStackShape) {

	// check if instance is already staged
	if stage.IsStaged(verticaltorusstackshape) {
		return
	}

	verticaltorusstackshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (volumekey3dshape *VolumeKey3DShape) GongStageBranch(stage *Stage) {
	stage.StageBranchVolumeKey3DShape(volumekey3dshape)
}

func (stage *Stage) StageBranchVolumeKey3DShape(volumekey3dshape *VolumeKey3DShape) {

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

	case *ClockAbstract:
		toT := GongCopyBranchClockAbstract(mapOrigCopy, fromT)
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

	case *MusicAbstract:
		toT := GongCopyBranchMusicAbstract(mapOrigCopy, fromT)
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

	case *StoolAbstract:
		toT := GongCopyBranchStoolAbstract(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TiledFloor3DShape:
		toT := GongCopyBranchTiledFloor3DShape(mapOrigCopy, fromT)
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

	// angle0shapeFrom has already been copied
	if _angle0shapeTo, ok := mapOrigCopy[angle0shapeFrom]; ok {
		angle0shapeTo = _angle0shapeTo.(*Angle0Shape)
		return
	}

	angle0shapeTo = new(Angle0Shape)
	mapOrigCopy[angle0shapeFrom] = angle0shapeTo
	angle0shapeFrom.GongCopyBasicFields(angle0shapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArcNormalVectorShape(mapOrigCopy map[any]any, arcnormalvectorshapeFrom *ArcNormalVectorShape) (arcnormalvectorshapeTo *ArcNormalVectorShape) {

	// arcnormalvectorshapeFrom has already been copied
	if _arcnormalvectorshapeTo, ok := mapOrigCopy[arcnormalvectorshapeFrom]; ok {
		arcnormalvectorshapeTo = _arcnormalvectorshapeTo.(*ArcNormalVectorShape)
		return
	}

	arcnormalvectorshapeTo = new(ArcNormalVectorShape)
	mapOrigCopy[arcnormalvectorshapeFrom] = arcnormalvectorshapeTo
	arcnormalvectorshapeFrom.GongCopyBasicFields(arcnormalvectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArcNormalVectorShapeGrid(mapOrigCopy map[any]any, arcnormalvectorshapegridFrom *ArcNormalVectorShapeGrid) (arcnormalvectorshapegridTo *ArcNormalVectorShapeGrid) {

	// arcnormalvectorshapegridFrom has already been copied
	if _arcnormalvectorshapegridTo, ok := mapOrigCopy[arcnormalvectorshapegridFrom]; ok {
		arcnormalvectorshapegridTo = _arcnormalvectorshapegridTo.(*ArcNormalVectorShapeGrid)
		return
	}

	arcnormalvectorshapegridTo = new(ArcNormalVectorShapeGrid)
	mapOrigCopy[arcnormalvectorshapegridFrom] = arcnormalvectorshapegridTo
	arcnormalvectorshapegridFrom.GongCopyBasicFields(arcnormalvectorshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAxesShape(mapOrigCopy map[any]any, axesshapeFrom *AxesShape) (axesshapeTo *AxesShape) {

	// axesshapeFrom has already been copied
	if _axesshapeTo, ok := mapOrigCopy[axesshapeFrom]; ok {
		axesshapeTo = _axesshapeTo.(*AxesShape)
		return
	}

	axesshapeTo = new(AxesShape)
	mapOrigCopy[axesshapeFrom] = axesshapeTo
	axesshapeFrom.GongCopyBasicFields(axesshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBaseVectorShape(mapOrigCopy map[any]any, basevectorshapeFrom *BaseVectorShape) (basevectorshapeTo *BaseVectorShape) {

	// basevectorshapeFrom has already been copied
	if _basevectorshapeTo, ok := mapOrigCopy[basevectorshapeFrom]; ok {
		basevectorshapeTo = _basevectorshapeTo.(*BaseVectorShape)
		return
	}

	basevectorshapeTo = new(BaseVectorShape)
	mapOrigCopy[basevectorshapeFrom] = basevectorshapeTo
	basevectorshapeFrom.GongCopyBasicFields(basevectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBaseVectorShapeGrid(mapOrigCopy map[any]any, basevectorshapegridFrom *BaseVectorShapeGrid) (basevectorshapegridTo *BaseVectorShapeGrid) {

	// basevectorshapegridFrom has already been copied
	if _basevectorshapegridTo, ok := mapOrigCopy[basevectorshapegridFrom]; ok {
		basevectorshapegridTo = _basevectorshapegridTo.(*BaseVectorShapeGrid)
		return
	}

	basevectorshapegridTo = new(BaseVectorShapeGrid)
	mapOrigCopy[basevectorshapegridFrom] = basevectorshapegridTo
	basevectorshapegridFrom.GongCopyBasicFields(basevectorshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchChosenP1P2PairShape(mapOrigCopy map[any]any, chosenp1p2pairshapeFrom *ChosenP1P2PairShape) (chosenp1p2pairshapeTo *ChosenP1P2PairShape) {

	// chosenp1p2pairshapeFrom has already been copied
	if _chosenp1p2pairshapeTo, ok := mapOrigCopy[chosenp1p2pairshapeFrom]; ok {
		chosenp1p2pairshapeTo = _chosenp1p2pairshapeTo.(*ChosenP1P2PairShape)
		return
	}

	chosenp1p2pairshapeTo = new(ChosenP1P2PairShape)
	mapOrigCopy[chosenp1p2pairshapeFrom] = chosenp1p2pairshapeTo
	chosenp1p2pairshapeFrom.GongCopyBasicFields(chosenp1p2pairshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCircleGridShape(mapOrigCopy map[any]any, circlegridshapeFrom *CircleGridShape) (circlegridshapeTo *CircleGridShape) {

	// circlegridshapeFrom has already been copied
	if _circlegridshapeTo, ok := mapOrigCopy[circlegridshapeFrom]; ok {
		circlegridshapeTo = _circlegridshapeTo.(*CircleGridShape)
		return
	}

	circlegridshapeTo = new(CircleGridShape)
	mapOrigCopy[circlegridshapeFrom] = circlegridshapeTo
	circlegridshapeFrom.GongCopyBasicFields(circlegridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCircumference3DShape(mapOrigCopy map[any]any, circumference3dshapeFrom *Circumference3DShape) (circumference3dshapeTo *Circumference3DShape) {

	// circumference3dshapeFrom has already been copied
	if _circumference3dshapeTo, ok := mapOrigCopy[circumference3dshapeFrom]; ok {
		circumference3dshapeTo = _circumference3dshapeTo.(*Circumference3DShape)
		return
	}

	circumference3dshapeTo = new(Circumference3DShape)
	mapOrigCopy[circumference3dshapeFrom] = circumference3dshapeTo
	circumference3dshapeFrom.GongCopyBasicFields(circumference3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchClock2DDiagram(mapOrigCopy map[any]any, clock2ddiagramFrom *Clock2DDiagram) (clock2ddiagramTo *Clock2DDiagram) {

	// clock2ddiagramFrom has already been copied
	if _clock2ddiagramTo, ok := mapOrigCopy[clock2ddiagramFrom]; ok {
		clock2ddiagramTo = _clock2ddiagramTo.(*Clock2DDiagram)
		return
	}

	clock2ddiagramTo = new(Clock2DDiagram)
	mapOrigCopy[clock2ddiagramFrom] = clock2ddiagramTo
	clock2ddiagramFrom.GongCopyBasicFields(clock2ddiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchClock3DDiagram(mapOrigCopy map[any]any, clock3ddiagramFrom *Clock3DDiagram) (clock3ddiagramTo *Clock3DDiagram) {

	// clock3ddiagramFrom has already been copied
	if _clock3ddiagramTo, ok := mapOrigCopy[clock3ddiagramFrom]; ok {
		clock3ddiagramTo = _clock3ddiagramTo.(*Clock3DDiagram)
		return
	}

	clock3ddiagramTo = new(Clock3DDiagram)
	mapOrigCopy[clock3ddiagramFrom] = clock3ddiagramTo
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

func GongCopyBranchClockAbstract(mapOrigCopy map[any]any, clockabstractFrom *ClockAbstract) (clockabstractTo *ClockAbstract) {

	// clockabstractFrom has already been copied
	if _clockabstractTo, ok := mapOrigCopy[clockabstractFrom]; ok {
		clockabstractTo = _clockabstractTo.(*ClockAbstract)
		return
	}

	clockabstractTo = new(ClockAbstract)
	mapOrigCopy[clockabstractFrom] = clockabstractTo
	clockabstractFrom.GongCopyBasicFields(clockabstractTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchClockTopCurveShape(mapOrigCopy map[any]any, clocktopcurveshapeFrom *ClockTopCurveShape) (clocktopcurveshapeTo *ClockTopCurveShape) {

	// clocktopcurveshapeFrom has already been copied
	if _clocktopcurveshapeTo, ok := mapOrigCopy[clocktopcurveshapeFrom]; ok {
		clocktopcurveshapeTo = _clocktopcurveshapeTo.(*ClockTopCurveShape)
		return
	}

	clocktopcurveshapeTo = new(ClockTopCurveShape)
	mapOrigCopy[clocktopcurveshapeFrom] = clocktopcurveshapeTo
	clocktopcurveshapeFrom.GongCopyBasicFields(clocktopcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCutLine3DShape(mapOrigCopy map[any]any, cutline3dshapeFrom *CutLine3DShape) (cutline3dshapeTo *CutLine3DShape) {

	// cutline3dshapeFrom has already been copied
	if _cutline3dshapeTo, ok := mapOrigCopy[cutline3dshapeFrom]; ok {
		cutline3dshapeTo = _cutline3dshapeTo.(*CutLine3DShape)
		return
	}

	cutline3dshapeTo = new(CutLine3DShape)
	mapOrigCopy[cutline3dshapeFrom] = cutline3dshapeTo
	cutline3dshapeFrom.GongCopyBasicFields(cutline3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEndArcShape(mapOrigCopy map[any]any, endarcshapeFrom *EndArcShape) (endarcshapeTo *EndArcShape) {

	// endarcshapeFrom has already been copied
	if _endarcshapeTo, ok := mapOrigCopy[endarcshapeFrom]; ok {
		endarcshapeTo = _endarcshapeTo.(*EndArcShape)
		return
	}

	endarcshapeTo = new(EndArcShape)
	mapOrigCopy[endarcshapeFrom] = endarcshapeTo
	endarcshapeFrom.GongCopyBasicFields(endarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEndArcShapeGrid(mapOrigCopy map[any]any, endarcshapegridFrom *EndArcShapeGrid) (endarcshapegridTo *EndArcShapeGrid) {

	// endarcshapegridFrom has already been copied
	if _endarcshapegridTo, ok := mapOrigCopy[endarcshapegridFrom]; ok {
		endarcshapegridTo = _endarcshapegridTo.(*EndArcShapeGrid)
		return
	}

	endarcshapegridTo = new(EndArcShapeGrid)
	mapOrigCopy[endarcshapegridFrom] = endarcshapegridTo
	endarcshapegridFrom.GongCopyBasicFields(endarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEndHalfwayArcShape(mapOrigCopy map[any]any, endhalfwayarcshapeFrom *EndHalfwayArcShape) (endhalfwayarcshapeTo *EndHalfwayArcShape) {

	// endhalfwayarcshapeFrom has already been copied
	if _endhalfwayarcshapeTo, ok := mapOrigCopy[endhalfwayarcshapeFrom]; ok {
		endhalfwayarcshapeTo = _endhalfwayarcshapeTo.(*EndHalfwayArcShape)
		return
	}

	endhalfwayarcshapeTo = new(EndHalfwayArcShape)
	mapOrigCopy[endhalfwayarcshapeFrom] = endhalfwayarcshapeTo
	endhalfwayarcshapeFrom.GongCopyBasicFields(endhalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEndHalfwayArcShapeGrid(mapOrigCopy map[any]any, endhalfwayarcshapegridFrom *EndHalfwayArcShapeGrid) (endhalfwayarcshapegridTo *EndHalfwayArcShapeGrid) {

	// endhalfwayarcshapegridFrom has already been copied
	if _endhalfwayarcshapegridTo, ok := mapOrigCopy[endhalfwayarcshapegridFrom]; ok {
		endhalfwayarcshapegridTo = _endhalfwayarcshapegridTo.(*EndHalfwayArcShapeGrid)
		return
	}

	endhalfwayarcshapegridTo = new(EndHalfwayArcShapeGrid)
	mapOrigCopy[endhalfwayarcshapegridFrom] = endhalfwayarcshapegridTo
	endhalfwayarcshapegridFrom.GongCopyBasicFields(endhalfwayarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchExplanationTextShape(mapOrigCopy map[any]any, explanationtextshapeFrom *ExplanationTextShape) (explanationtextshapeTo *ExplanationTextShape) {

	// explanationtextshapeFrom has already been copied
	if _explanationtextshapeTo, ok := mapOrigCopy[explanationtextshapeFrom]; ok {
		explanationtextshapeTo = _explanationtextshapeTo.(*ExplanationTextShape)
		return
	}

	explanationtextshapeTo = new(ExplanationTextShape)
	mapOrigCopy[explanationtextshapeFrom] = explanationtextshapeTo
	explanationtextshapeFrom.GongCopyBasicFields(explanationtextshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEye3DShape(mapOrigCopy map[any]any, eye3dshapeFrom *Eye3DShape) (eye3dshapeTo *Eye3DShape) {

	// eye3dshapeFrom has already been copied
	if _eye3dshapeTo, ok := mapOrigCopy[eye3dshapeFrom]; ok {
		eye3dshapeTo = _eye3dshapeTo.(*Eye3DShape)
		return
	}

	eye3dshapeTo = new(Eye3DShape)
	mapOrigCopy[eye3dshapeFrom] = eye3dshapeTo
	eye3dshapeFrom.GongCopyBasicFields(eye3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEyeCornersSampledPoints3DShape(mapOrigCopy map[any]any, eyecornerssampledpoints3dshapeFrom *EyeCornersSampledPoints3DShape) (eyecornerssampledpoints3dshapeTo *EyeCornersSampledPoints3DShape) {

	// eyecornerssampledpoints3dshapeFrom has already been copied
	if _eyecornerssampledpoints3dshapeTo, ok := mapOrigCopy[eyecornerssampledpoints3dshapeFrom]; ok {
		eyecornerssampledpoints3dshapeTo = _eyecornerssampledpoints3dshapeTo.(*EyeCornersSampledPoints3DShape)
		return
	}

	eyecornerssampledpoints3dshapeTo = new(EyeCornersSampledPoints3DShape)
	mapOrigCopy[eyecornerssampledpoints3dshapeFrom] = eyecornerssampledpoints3dshapeTo
	eyecornerssampledpoints3dshapeFrom.GongCopyBasicFields(eyecornerssampledpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEyeSampledPoints3DShape(mapOrigCopy map[any]any, eyesampledpoints3dshapeFrom *EyeSampledPoints3DShape) (eyesampledpoints3dshapeTo *EyeSampledPoints3DShape) {

	// eyesampledpoints3dshapeFrom has already been copied
	if _eyesampledpoints3dshapeTo, ok := mapOrigCopy[eyesampledpoints3dshapeFrom]; ok {
		eyesampledpoints3dshapeTo = _eyesampledpoints3dshapeTo.(*EyeSampledPoints3DShape)
		return
	}

	eyesampledpoints3dshapeTo = new(EyeSampledPoints3DShape)
	mapOrigCopy[eyesampledpoints3dshapeFrom] = eyesampledpoints3dshapeTo
	eyesampledpoints3dshapeFrom.GongCopyBasicFields(eyesampledpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEyeSeatBottomCurveShape(mapOrigCopy map[any]any, eyeseatbottomcurveshapeFrom *EyeSeatBottomCurveShape) (eyeseatbottomcurveshapeTo *EyeSeatBottomCurveShape) {

	// eyeseatbottomcurveshapeFrom has already been copied
	if _eyeseatbottomcurveshapeTo, ok := mapOrigCopy[eyeseatbottomcurveshapeFrom]; ok {
		eyeseatbottomcurveshapeTo = _eyeseatbottomcurveshapeTo.(*EyeSeatBottomCurveShape)
		return
	}

	eyeseatbottomcurveshapeTo = new(EyeSeatBottomCurveShape)
	mapOrigCopy[eyeseatbottomcurveshapeFrom] = eyeseatbottomcurveshapeTo
	eyeseatbottomcurveshapeFrom.GongCopyBasicFields(eyeseatbottomcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEyeStoolBottomCurveShape(mapOrigCopy map[any]any, eyestoolbottomcurveshapeFrom *EyeStoolBottomCurveShape) (eyestoolbottomcurveshapeTo *EyeStoolBottomCurveShape) {

	// eyestoolbottomcurveshapeFrom has already been copied
	if _eyestoolbottomcurveshapeTo, ok := mapOrigCopy[eyestoolbottomcurveshapeFrom]; ok {
		eyestoolbottomcurveshapeTo = _eyestoolbottomcurveshapeTo.(*EyeStoolBottomCurveShape)
		return
	}

	eyestoolbottomcurveshapeTo = new(EyeStoolBottomCurveShape)
	mapOrigCopy[eyestoolbottomcurveshapeFrom] = eyestoolbottomcurveshapeTo
	eyestoolbottomcurveshapeFrom.GongCopyBasicFields(eyestoolbottomcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEyeVolume3DShape(mapOrigCopy map[any]any, eyevolume3dshapeFrom *EyeVolume3DShape) (eyevolume3dshapeTo *EyeVolume3DShape) {

	// eyevolume3dshapeFrom has already been copied
	if _eyevolume3dshapeTo, ok := mapOrigCopy[eyevolume3dshapeFrom]; ok {
		eyevolume3dshapeTo = _eyevolume3dshapeTo.(*EyeVolume3DShape)
		return
	}

	eyevolume3dshapeTo = new(EyeVolume3DShape)
	mapOrigCopy[eyevolume3dshapeFrom] = eyevolume3dshapeTo
	eyevolume3dshapeFrom.GongCopyBasicFields(eyevolume3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGridPathShape(mapOrigCopy map[any]any, gridpathshapeFrom *GridPathShape) (gridpathshapeTo *GridPathShape) {

	// gridpathshapeFrom has already been copied
	if _gridpathshapeTo, ok := mapOrigCopy[gridpathshapeFrom]; ok {
		gridpathshapeTo = _gridpathshapeTo.(*GridPathShape)
		return
	}

	gridpathshapeTo = new(GridPathShape)
	mapOrigCopy[gridpathshapeFrom] = gridpathshapeTo
	gridpathshapeFrom.GongCopyBasicFields(gridpathshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthCurve2D(mapOrigCopy map[any]any, growthcurve2dFrom *GrowthCurve2D) (growthcurve2dTo *GrowthCurve2D) {

	// growthcurve2dFrom has already been copied
	if _growthcurve2dTo, ok := mapOrigCopy[growthcurve2dFrom]; ok {
		growthcurve2dTo = _growthcurve2dTo.(*GrowthCurve2D)
		return
	}

	growthcurve2dTo = new(GrowthCurve2D)
	mapOrigCopy[growthcurve2dFrom] = growthcurve2dTo
	growthcurve2dFrom.GongCopyBasicFields(growthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthCurve2DRibbon(mapOrigCopy map[any]any, growthcurve2dribbonFrom *GrowthCurve2DRibbon) (growthcurve2dribbonTo *GrowthCurve2DRibbon) {

	// growthcurve2dribbonFrom has already been copied
	if _growthcurve2dribbonTo, ok := mapOrigCopy[growthcurve2dribbonFrom]; ok {
		growthcurve2dribbonTo = _growthcurve2dribbonTo.(*GrowthCurve2DRibbon)
		return
	}

	growthcurve2dribbonTo = new(GrowthCurve2DRibbon)
	mapOrigCopy[growthcurve2dribbonFrom] = growthcurve2dribbonTo
	growthcurve2dribbonFrom.GongCopyBasicFields(growthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, growthcurve2dribbonendshapeFrom *GrowthCurve2DRibbonEndShape) (growthcurve2dribbonendshapeTo *GrowthCurve2DRibbonEndShape) {

	// growthcurve2dribbonendshapeFrom has already been copied
	if _growthcurve2dribbonendshapeTo, ok := mapOrigCopy[growthcurve2dribbonendshapeFrom]; ok {
		growthcurve2dribbonendshapeTo = _growthcurve2dribbonendshapeTo.(*GrowthCurve2DRibbonEndShape)
		return
	}

	growthcurve2dribbonendshapeTo = new(GrowthCurve2DRibbonEndShape)
	mapOrigCopy[growthcurve2dribbonendshapeFrom] = growthcurve2dribbonendshapeTo
	growthcurve2dribbonendshapeFrom.GongCopyBasicFields(growthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, growthcurve2dribbonstartshapeFrom *GrowthCurve2DRibbonStartShape) (growthcurve2dribbonstartshapeTo *GrowthCurve2DRibbonStartShape) {

	// growthcurve2dribbonstartshapeFrom has already been copied
	if _growthcurve2dribbonstartshapeTo, ok := mapOrigCopy[growthcurve2dribbonstartshapeFrom]; ok {
		growthcurve2dribbonstartshapeTo = _growthcurve2dribbonstartshapeTo.(*GrowthCurve2DRibbonStartShape)
		return
	}

	growthcurve2dribbonstartshapeTo = new(GrowthCurve2DRibbonStartShape)
	mapOrigCopy[growthcurve2dribbonstartshapeFrom] = growthcurve2dribbonstartshapeTo
	growthcurve2dribbonstartshapeFrom.GongCopyBasicFields(growthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthCurveRhombusGridShape(mapOrigCopy map[any]any, growthcurverhombusgridshapeFrom *GrowthCurveRhombusGridShape) (growthcurverhombusgridshapeTo *GrowthCurveRhombusGridShape) {

	// growthcurverhombusgridshapeFrom has already been copied
	if _growthcurverhombusgridshapeTo, ok := mapOrigCopy[growthcurverhombusgridshapeFrom]; ok {
		growthcurverhombusgridshapeTo = _growthcurverhombusgridshapeTo.(*GrowthCurveRhombusGridShape)
		return
	}

	growthcurverhombusgridshapeTo = new(GrowthCurveRhombusGridShape)
	mapOrigCopy[growthcurverhombusgridshapeFrom] = growthcurverhombusgridshapeTo
	growthcurverhombusgridshapeFrom.GongCopyBasicFields(growthcurverhombusgridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthCurveRhombusShape(mapOrigCopy map[any]any, growthcurverhombusshapeFrom *GrowthCurveRhombusShape) (growthcurverhombusshapeTo *GrowthCurveRhombusShape) {

	// growthcurverhombusshapeFrom has already been copied
	if _growthcurverhombusshapeTo, ok := mapOrigCopy[growthcurverhombusshapeFrom]; ok {
		growthcurverhombusshapeTo = _growthcurverhombusshapeTo.(*GrowthCurveRhombusShape)
		return
	}

	growthcurverhombusshapeTo = new(GrowthCurveRhombusShape)
	mapOrigCopy[growthcurverhombusshapeFrom] = growthcurverhombusshapeTo
	growthcurverhombusshapeFrom.GongCopyBasicFields(growthcurverhombusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrowthVectorShape(mapOrigCopy map[any]any, growthvectorshapeFrom *GrowthVectorShape) (growthvectorshapeTo *GrowthVectorShape) {

	// growthvectorshapeFrom has already been copied
	if _growthvectorshapeTo, ok := mapOrigCopy[growthvectorshapeFrom]; ok {
		growthvectorshapeTo = _growthvectorshapeTo.(*GrowthVectorShape)
		return
	}

	growthvectorshapeTo = new(GrowthVectorShape)
	mapOrigCopy[growthvectorshapeFrom] = growthvectorshapeTo
	growthvectorshapeFrom.GongCopyBasicFields(growthvectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInitialRhombusGridShape(mapOrigCopy map[any]any, initialrhombusgridshapeFrom *InitialRhombusGridShape) (initialrhombusgridshapeTo *InitialRhombusGridShape) {

	// initialrhombusgridshapeFrom has already been copied
	if _initialrhombusgridshapeTo, ok := mapOrigCopy[initialrhombusgridshapeFrom]; ok {
		initialrhombusgridshapeTo = _initialrhombusgridshapeTo.(*InitialRhombusGridShape)
		return
	}

	initialrhombusgridshapeTo = new(InitialRhombusGridShape)
	mapOrigCopy[initialrhombusgridshapeFrom] = initialrhombusgridshapeTo
	initialrhombusgridshapeFrom.GongCopyBasicFields(initialrhombusgridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInitialRhombusShape(mapOrigCopy map[any]any, initialrhombusshapeFrom *InitialRhombusShape) (initialrhombusshapeTo *InitialRhombusShape) {

	// initialrhombusshapeFrom has already been copied
	if _initialrhombusshapeTo, ok := mapOrigCopy[initialrhombusshapeFrom]; ok {
		initialrhombusshapeTo = _initialrhombusshapeTo.(*InitialRhombusShape)
		return
	}

	initialrhombusshapeTo = new(InitialRhombusShape)
	mapOrigCopy[initialrhombusshapeFrom] = initialrhombusshapeTo
	initialrhombusshapeFrom.GongCopyBasicFields(initialrhombusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKey3DShape(mapOrigCopy map[any]any, key3dshapeFrom *Key3DShape) (key3dshapeTo *Key3DShape) {

	// key3dshapeFrom has already been copied
	if _key3dshapeTo, ok := mapOrigCopy[key3dshapeFrom]; ok {
		key3dshapeTo = _key3dshapeTo.(*Key3DShape)
		return
	}

	key3dshapeTo = new(Key3DShape)
	mapOrigCopy[key3dshapeFrom] = key3dshapeTo
	key3dshapeFrom.GongCopyBasicFields(key3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKeyHole3DShape(mapOrigCopy map[any]any, keyhole3dshapeFrom *KeyHole3DShape) (keyhole3dshapeTo *KeyHole3DShape) {

	// keyhole3dshapeFrom has already been copied
	if _keyhole3dshapeTo, ok := mapOrigCopy[keyhole3dshapeFrom]; ok {
		keyhole3dshapeTo = _keyhole3dshapeTo.(*KeyHole3DShape)
		return
	}

	keyhole3dshapeTo = new(KeyHole3DShape)
	mapOrigCopy[keyhole3dshapeFrom] = keyhole3dshapeTo
	keyhole3dshapeFrom.GongCopyBasicFields(keyhole3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKeyHoleShape(mapOrigCopy map[any]any, keyholeshapeFrom *KeyHoleShape) (keyholeshapeTo *KeyHoleShape) {

	// keyholeshapeFrom has already been copied
	if _keyholeshapeTo, ok := mapOrigCopy[keyholeshapeFrom]; ok {
		keyholeshapeTo = _keyholeshapeTo.(*KeyHoleShape)
		return
	}

	keyholeshapeTo = new(KeyHoleShape)
	mapOrigCopy[keyholeshapeFrom] = keyholeshapeTo
	keyholeshapeFrom.GongCopyBasicFields(keyholeshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLeaves3DShape(mapOrigCopy map[any]any, leaves3dshapeFrom *Leaves3DShape) (leaves3dshapeTo *Leaves3DShape) {

	// leaves3dshapeFrom has already been copied
	if _leaves3dshapeTo, ok := mapOrigCopy[leaves3dshapeFrom]; ok {
		leaves3dshapeTo = _leaves3dshapeTo.(*Leaves3DShape)
		return
	}

	leaves3dshapeTo = new(Leaves3DShape)
	mapOrigCopy[leaves3dshapeFrom] = leaves3dshapeTo
	leaves3dshapeFrom.GongCopyBasicFields(leaves3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {

	// libraryFrom has already been copied
	if _libraryTo, ok := mapOrigCopy[libraryFrom]; ok {
		libraryTo = _libraryTo.(*Library)
		return
	}

	libraryTo = new(Library)
	mapOrigCopy[libraryFrom] = libraryTo
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

	// midarcvectorshapeFrom has already been copied
	if _midarcvectorshapeTo, ok := mapOrigCopy[midarcvectorshapeFrom]; ok {
		midarcvectorshapeTo = _midarcvectorshapeTo.(*MidArcVectorShape)
		return
	}

	midarcvectorshapeTo = new(MidArcVectorShape)
	mapOrigCopy[midarcvectorshapeFrom] = midarcvectorshapeTo
	midarcvectorshapeFrom.GongCopyBasicFields(midarcvectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMidArcVectorShapeGrid(mapOrigCopy map[any]any, midarcvectorshapegridFrom *MidArcVectorShapeGrid) (midarcvectorshapegridTo *MidArcVectorShapeGrid) {

	// midarcvectorshapegridFrom has already been copied
	if _midarcvectorshapegridTo, ok := mapOrigCopy[midarcvectorshapegridFrom]; ok {
		midarcvectorshapegridTo = _midarcvectorshapegridTo.(*MidArcVectorShapeGrid)
		return
	}

	midarcvectorshapegridTo = new(MidArcVectorShapeGrid)
	mapOrigCopy[midarcvectorshapegridFrom] = midarcvectorshapegridTo
	midarcvectorshapegridFrom.GongCopyBasicFields(midarcvectorshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMusicAbstract(mapOrigCopy map[any]any, musicabstractFrom *MusicAbstract) (musicabstractTo *MusicAbstract) {

	// musicabstractFrom has already been copied
	if _musicabstractTo, ok := mapOrigCopy[musicabstractFrom]; ok {
		musicabstractTo = _musicabstractTo.(*MusicAbstract)
		return
	}

	musicabstractTo = new(MusicAbstract)
	mapOrigCopy[musicabstractFrom] = musicabstractTo
	musicabstractFrom.GongCopyBasicFields(musicabstractTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOriginalPoints3DShape(mapOrigCopy map[any]any, originalpoints3dshapeFrom *OriginalPoints3DShape) (originalpoints3dshapeTo *OriginalPoints3DShape) {

	// originalpoints3dshapeFrom has already been copied
	if _originalpoints3dshapeTo, ok := mapOrigCopy[originalpoints3dshapeFrom]; ok {
		originalpoints3dshapeTo = _originalpoints3dshapeTo.(*OriginalPoints3DShape)
		return
	}

	originalpoints3dshapeTo = new(OriginalPoints3DShape)
	mapOrigCopy[originalpoints3dshapeFrom] = originalpoints3dshapeTo
	originalpoints3dshapeFrom.GongCopyBasicFields(originalpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchParastichyMCurves3DShape(mapOrigCopy map[any]any, parastichymcurves3dshapeFrom *ParastichyMCurves3DShape) (parastichymcurves3dshapeTo *ParastichyMCurves3DShape) {

	// parastichymcurves3dshapeFrom has already been copied
	if _parastichymcurves3dshapeTo, ok := mapOrigCopy[parastichymcurves3dshapeFrom]; ok {
		parastichymcurves3dshapeTo = _parastichymcurves3dshapeTo.(*ParastichyMCurves3DShape)
		return
	}

	parastichymcurves3dshapeTo = new(ParastichyMCurves3DShape)
	mapOrigCopy[parastichymcurves3dshapeFrom] = parastichymcurves3dshapeTo
	parastichymcurves3dshapeFrom.GongCopyBasicFields(parastichymcurves3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchParastichyNCurves3DShape(mapOrigCopy map[any]any, parastichyncurves3dshapeFrom *ParastichyNCurves3DShape) (parastichyncurves3dshapeTo *ParastichyNCurves3DShape) {

	// parastichyncurves3dshapeFrom has already been copied
	if _parastichyncurves3dshapeTo, ok := mapOrigCopy[parastichyncurves3dshapeFrom]; ok {
		parastichyncurves3dshapeTo = _parastichyncurves3dshapeTo.(*ParastichyNCurves3DShape)
		return
	}

	parastichyncurves3dshapeTo = new(ParastichyNCurves3DShape)
	mapOrigCopy[parastichyncurves3dshapeFrom] = parastichyncurves3dshapeTo
	parastichyncurves3dshapeFrom.GongCopyBasicFields(parastichyncurves3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DRibbon(mapOrigCopy map[any]any, partiallygrowthcurve2dribbonFrom *PartiallyGrowthCurve2DRibbon) (partiallygrowthcurve2dribbonTo *PartiallyGrowthCurve2DRibbon) {

	// partiallygrowthcurve2dribbonFrom has already been copied
	if _partiallygrowthcurve2dribbonTo, ok := mapOrigCopy[partiallygrowthcurve2dribbonFrom]; ok {
		partiallygrowthcurve2dribbonTo = _partiallygrowthcurve2dribbonTo.(*PartiallyGrowthCurve2DRibbon)
		return
	}

	partiallygrowthcurve2dribbonTo = new(PartiallyGrowthCurve2DRibbon)
	mapOrigCopy[partiallygrowthcurve2dribbonFrom] = partiallygrowthcurve2dribbonTo
	partiallygrowthcurve2dribbonFrom.GongCopyBasicFields(partiallygrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, partiallygrowthcurve2dribbonendshapeFrom *PartiallyGrowthCurve2DRibbonEndShape) (partiallygrowthcurve2dribbonendshapeTo *PartiallyGrowthCurve2DRibbonEndShape) {

	// partiallygrowthcurve2dribbonendshapeFrom has already been copied
	if _partiallygrowthcurve2dribbonendshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dribbonendshapeFrom]; ok {
		partiallygrowthcurve2dribbonendshapeTo = _partiallygrowthcurve2dribbonendshapeTo.(*PartiallyGrowthCurve2DRibbonEndShape)
		return
	}

	partiallygrowthcurve2dribbonendshapeTo = new(PartiallyGrowthCurve2DRibbonEndShape)
	mapOrigCopy[partiallygrowthcurve2dribbonendshapeFrom] = partiallygrowthcurve2dribbonendshapeTo
	partiallygrowthcurve2dribbonendshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, partiallygrowthcurve2dribbonstartshapeFrom *PartiallyGrowthCurve2DRibbonStartShape) (partiallygrowthcurve2dribbonstartshapeTo *PartiallyGrowthCurve2DRibbonStartShape) {

	// partiallygrowthcurve2dribbonstartshapeFrom has already been copied
	if _partiallygrowthcurve2dribbonstartshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dribbonstartshapeFrom]; ok {
		partiallygrowthcurve2dribbonstartshapeTo = _partiallygrowthcurve2dribbonstartshapeTo.(*PartiallyGrowthCurve2DRibbonStartShape)
		return
	}

	partiallygrowthcurve2dribbonstartshapeTo = new(PartiallyGrowthCurve2DRibbonStartShape)
	mapOrigCopy[partiallygrowthcurve2dribbonstartshapeFrom] = partiallygrowthcurve2dribbonstartshapeTo
	partiallygrowthcurve2dribbonstartshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectory(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryFrom *PartiallyGrowthCurve2DTrajectory) (partiallygrowthcurve2dtrajectoryTo *PartiallyGrowthCurve2DTrajectory) {

	// partiallygrowthcurve2dtrajectoryFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryFrom]; ok {
		partiallygrowthcurve2dtrajectoryTo = _partiallygrowthcurve2dtrajectoryTo.(*PartiallyGrowthCurve2DTrajectory)
		return
	}

	partiallygrowthcurve2dtrajectoryTo = new(PartiallyGrowthCurve2DTrajectory)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryFrom] = partiallygrowthcurve2dtrajectoryTo
	partiallygrowthcurve2dtrajectoryFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryP1CurveShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp1curveshapeFrom *PartiallyGrowthCurve2DTrajectoryP1CurveShape) (partiallygrowthcurve2dtrajectoryp1curveshapeTo *PartiallyGrowthCurve2DTrajectoryP1CurveShape) {

	// partiallygrowthcurve2dtrajectoryp1curveshapeFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryp1curveshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryp1curveshapeFrom]; ok {
		partiallygrowthcurve2dtrajectoryp1curveshapeTo = _partiallygrowthcurve2dtrajectoryp1curveshapeTo.(*PartiallyGrowthCurve2DTrajectoryP1CurveShape)
		return
	}

	partiallygrowthcurve2dtrajectoryp1curveshapeTo = new(PartiallyGrowthCurve2DTrajectoryP1CurveShape)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryp1curveshapeFrom] = partiallygrowthcurve2dtrajectoryp1curveshapeTo
	partiallygrowthcurve2dtrajectoryp1curveshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryp1curveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryP1P2(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp1p2From *PartiallyGrowthCurve2DTrajectoryP1P2) (partiallygrowthcurve2dtrajectoryp1p2To *PartiallyGrowthCurve2DTrajectoryP1P2) {

	// partiallygrowthcurve2dtrajectoryp1p2From has already been copied
	if _partiallygrowthcurve2dtrajectoryp1p2To, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryp1p2From]; ok {
		partiallygrowthcurve2dtrajectoryp1p2To = _partiallygrowthcurve2dtrajectoryp1p2To.(*PartiallyGrowthCurve2DTrajectoryP1P2)
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2To = new(PartiallyGrowthCurve2DTrajectoryP1P2)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryp1p2From] = partiallygrowthcurve2dtrajectoryp1p2To
	partiallygrowthcurve2dtrajectoryp1p2From.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryp1p2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFrom *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) (partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) {

	// partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFrom]; ok {
		partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo = _partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo.(*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo = new(PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFrom] = partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo
	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryp1p2pairlineshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryP1PointShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp1pointshapeFrom *PartiallyGrowthCurve2DTrajectoryP1PointShape) (partiallygrowthcurve2dtrajectoryp1pointshapeTo *PartiallyGrowthCurve2DTrajectoryP1PointShape) {

	// partiallygrowthcurve2dtrajectoryp1pointshapeFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryp1pointshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryp1pointshapeFrom]; ok {
		partiallygrowthcurve2dtrajectoryp1pointshapeTo = _partiallygrowthcurve2dtrajectoryp1pointshapeTo.(*PartiallyGrowthCurve2DTrajectoryP1PointShape)
		return
	}

	partiallygrowthcurve2dtrajectoryp1pointshapeTo = new(PartiallyGrowthCurve2DTrajectoryP1PointShape)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryp1pointshapeFrom] = partiallygrowthcurve2dtrajectoryp1pointshapeTo
	partiallygrowthcurve2dtrajectoryp1pointshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryp1pointshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryP2CurveShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp2curveshapeFrom *PartiallyGrowthCurve2DTrajectoryP2CurveShape) (partiallygrowthcurve2dtrajectoryp2curveshapeTo *PartiallyGrowthCurve2DTrajectoryP2CurveShape) {

	// partiallygrowthcurve2dtrajectoryp2curveshapeFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryp2curveshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryp2curveshapeFrom]; ok {
		partiallygrowthcurve2dtrajectoryp2curveshapeTo = _partiallygrowthcurve2dtrajectoryp2curveshapeTo.(*PartiallyGrowthCurve2DTrajectoryP2CurveShape)
		return
	}

	partiallygrowthcurve2dtrajectoryp2curveshapeTo = new(PartiallyGrowthCurve2DTrajectoryP2CurveShape)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryp2curveshapeFrom] = partiallygrowthcurve2dtrajectoryp2curveshapeTo
	partiallygrowthcurve2dtrajectoryp2curveshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryp2curveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryP2PointShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryp2pointshapeFrom *PartiallyGrowthCurve2DTrajectoryP2PointShape) (partiallygrowthcurve2dtrajectoryp2pointshapeTo *PartiallyGrowthCurve2DTrajectoryP2PointShape) {

	// partiallygrowthcurve2dtrajectoryp2pointshapeFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryp2pointshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryp2pointshapeFrom]; ok {
		partiallygrowthcurve2dtrajectoryp2pointshapeTo = _partiallygrowthcurve2dtrajectoryp2pointshapeTo.(*PartiallyGrowthCurve2DTrajectoryP2PointShape)
		return
	}

	partiallygrowthcurve2dtrajectoryp2pointshapeTo = new(PartiallyGrowthCurve2DTrajectoryP2PointShape)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryp2pointshapeFrom] = partiallygrowthcurve2dtrajectoryp2pointshapeTo
	partiallygrowthcurve2dtrajectoryp2pointshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryp2pointshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyGrowthCurve2DTrajectoryShape(mapOrigCopy map[any]any, partiallygrowthcurve2dtrajectoryshapeFrom *PartiallyGrowthCurve2DTrajectoryShape) (partiallygrowthcurve2dtrajectoryshapeTo *PartiallyGrowthCurve2DTrajectoryShape) {

	// partiallygrowthcurve2dtrajectoryshapeFrom has already been copied
	if _partiallygrowthcurve2dtrajectoryshapeTo, ok := mapOrigCopy[partiallygrowthcurve2dtrajectoryshapeFrom]; ok {
		partiallygrowthcurve2dtrajectoryshapeTo = _partiallygrowthcurve2dtrajectoryshapeTo.(*PartiallyGrowthCurve2DTrajectoryShape)
		return
	}

	partiallygrowthcurve2dtrajectoryshapeTo = new(PartiallyGrowthCurve2DTrajectoryShape)
	mapOrigCopy[partiallygrowthcurve2dtrajectoryshapeFrom] = partiallygrowthcurve2dtrajectoryshapeTo
	partiallygrowthcurve2dtrajectoryshapeFrom.GongCopyBasicFields(partiallygrowthcurve2dtrajectoryshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyRotatedSeatBottomCurveShape(mapOrigCopy map[any]any, partiallyrotatedseatbottomcurveshapeFrom *PartiallyRotatedSeatBottomCurveShape) (partiallyrotatedseatbottomcurveshapeTo *PartiallyRotatedSeatBottomCurveShape) {

	// partiallyrotatedseatbottomcurveshapeFrom has already been copied
	if _partiallyrotatedseatbottomcurveshapeTo, ok := mapOrigCopy[partiallyrotatedseatbottomcurveshapeFrom]; ok {
		partiallyrotatedseatbottomcurveshapeTo = _partiallyrotatedseatbottomcurveshapeTo.(*PartiallyRotatedSeatBottomCurveShape)
		return
	}

	partiallyrotatedseatbottomcurveshapeTo = new(PartiallyRotatedSeatBottomCurveShape)
	mapOrigCopy[partiallyrotatedseatbottomcurveshapeFrom] = partiallyrotatedseatbottomcurveshapeTo
	partiallyrotatedseatbottomcurveshapeFrom.GongCopyBasicFields(partiallyrotatedseatbottomcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyRotatedSeatTopCurveShape(mapOrigCopy map[any]any, partiallyrotatedseattopcurveshapeFrom *PartiallyRotatedSeatTopCurveShape) (partiallyrotatedseattopcurveshapeTo *PartiallyRotatedSeatTopCurveShape) {

	// partiallyrotatedseattopcurveshapeFrom has already been copied
	if _partiallyrotatedseattopcurveshapeTo, ok := mapOrigCopy[partiallyrotatedseattopcurveshapeFrom]; ok {
		partiallyrotatedseattopcurveshapeTo = _partiallyrotatedseattopcurveshapeTo.(*PartiallyRotatedSeatTopCurveShape)
		return
	}

	partiallyrotatedseattopcurveshapeTo = new(PartiallyRotatedSeatTopCurveShape)
	mapOrigCopy[partiallyrotatedseattopcurveshapeFrom] = partiallyrotatedseattopcurveshapeTo
	partiallyrotatedseattopcurveshapeFrom.GongCopyBasicFields(partiallyrotatedseattopcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartiallyRotatedTorusShape(mapOrigCopy map[any]any, partiallyrotatedtorusshapeFrom *PartiallyRotatedTorusShape) (partiallyrotatedtorusshapeTo *PartiallyRotatedTorusShape) {

	// partiallyrotatedtorusshapeFrom has already been copied
	if _partiallyrotatedtorusshapeTo, ok := mapOrigCopy[partiallyrotatedtorusshapeFrom]; ok {
		partiallyrotatedtorusshapeTo = _partiallyrotatedtorusshapeTo.(*PartiallyRotatedTorusShape)
		return
	}

	partiallyrotatedtorusshapeTo = new(PartiallyRotatedTorusShape)
	mapOrigCopy[partiallyrotatedtorusshapeFrom] = partiallyrotatedtorusshapeTo
	partiallyrotatedtorusshapeFrom.GongCopyBasicFields(partiallyrotatedtorusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPerpendicularVector(mapOrigCopy map[any]any, perpendicularvectorFrom *PerpendicularVector) (perpendicularvectorTo *PerpendicularVector) {

	// perpendicularvectorFrom has already been copied
	if _perpendicularvectorTo, ok := mapOrigCopy[perpendicularvectorFrom]; ok {
		perpendicularvectorTo = _perpendicularvectorTo.(*PerpendicularVector)
		return
	}

	perpendicularvectorTo = new(PerpendicularVector)
	mapOrigCopy[perpendicularvectorFrom] = perpendicularvectorTo
	perpendicularvectorFrom.GongCopyBasicFields(perpendicularvectorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPerpendicularVectorGrid(mapOrigCopy map[any]any, perpendicularvectorgridFrom *PerpendicularVectorGrid) (perpendicularvectorgridTo *PerpendicularVectorGrid) {

	// perpendicularvectorgridFrom has already been copied
	if _perpendicularvectorgridTo, ok := mapOrigCopy[perpendicularvectorgridFrom]; ok {
		perpendicularvectorgridTo = _perpendicularvectorgridTo.(*PerpendicularVectorGrid)
		return
	}

	perpendicularvectorgridTo = new(PerpendicularVectorGrid)
	mapOrigCopy[perpendicularvectorgridFrom] = perpendicularvectorgridTo
	perpendicularvectorgridFrom.GongCopyBasicFields(perpendicularvectorgridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPerpendicularVectorGridHalfway(mapOrigCopy map[any]any, perpendicularvectorgridhalfwayFrom *PerpendicularVectorGridHalfway) (perpendicularvectorgridhalfwayTo *PerpendicularVectorGridHalfway) {

	// perpendicularvectorgridhalfwayFrom has already been copied
	if _perpendicularvectorgridhalfwayTo, ok := mapOrigCopy[perpendicularvectorgridhalfwayFrom]; ok {
		perpendicularvectorgridhalfwayTo = _perpendicularvectorgridhalfwayTo.(*PerpendicularVectorGridHalfway)
		return
	}

	perpendicularvectorgridhalfwayTo = new(PerpendicularVectorGridHalfway)
	mapOrigCopy[perpendicularvectorgridhalfwayFrom] = perpendicularvectorgridhalfwayTo
	perpendicularvectorgridhalfwayFrom.GongCopyBasicFields(perpendicularvectorgridhalfwayTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPerpendicularVectorHalfway(mapOrigCopy map[any]any, perpendicularvectorhalfwayFrom *PerpendicularVectorHalfway) (perpendicularvectorhalfwayTo *PerpendicularVectorHalfway) {

	// perpendicularvectorhalfwayFrom has already been copied
	if _perpendicularvectorhalfwayTo, ok := mapOrigCopy[perpendicularvectorhalfwayFrom]; ok {
		perpendicularvectorhalfwayTo = _perpendicularvectorhalfwayTo.(*PerpendicularVectorHalfway)
		return
	}

	perpendicularvectorhalfwayTo = new(PerpendicularVectorHalfway)
	mapOrigCopy[perpendicularvectorhalfwayFrom] = perpendicularvectorhalfwayTo
	perpendicularvectorhalfwayFrom.GongCopyBasicFields(perpendicularvectorhalfwayTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPlant2DDiagram(mapOrigCopy map[any]any, plant2ddiagramFrom *Plant2DDiagram) (plant2ddiagramTo *Plant2DDiagram) {

	// plant2ddiagramFrom has already been copied
	if _plant2ddiagramTo, ok := mapOrigCopy[plant2ddiagramFrom]; ok {
		plant2ddiagramTo = _plant2ddiagramTo.(*Plant2DDiagram)
		return
	}

	plant2ddiagramTo = new(Plant2DDiagram)
	mapOrigCopy[plant2ddiagramFrom] = plant2ddiagramTo
	plant2ddiagramFrom.GongCopyBasicFields(plant2ddiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPlant3DDiagram(mapOrigCopy map[any]any, plant3ddiagramFrom *Plant3DDiagram) (plant3ddiagramTo *Plant3DDiagram) {

	// plant3ddiagramFrom has already been copied
	if _plant3ddiagramTo, ok := mapOrigCopy[plant3ddiagramFrom]; ok {
		plant3ddiagramTo = _plant3ddiagramTo.(*Plant3DDiagram)
		return
	}

	plant3ddiagramTo = new(Plant3DDiagram)
	mapOrigCopy[plant3ddiagramFrom] = plant3ddiagramTo
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

	// plantabstractFrom has already been copied
	if _plantabstractTo, ok := mapOrigCopy[plantabstractFrom]; ok {
		plantabstractTo = _plantabstractTo.(*PlantAbstract)
		return
	}

	plantabstractTo = new(PlantAbstract)
	mapOrigCopy[plantabstractFrom] = plantabstractTo
	plantabstractFrom.GongCopyBasicFields(plantabstractTo)

	//insertion point for the staging of instances referenced by pointers
	if plantabstractFrom.TubeVaseAbstract != nil {
		plantabstractTo.TubeVaseAbstract = GongCopyBranchTubeVaseAbstract(mapOrigCopy, plantabstractFrom.TubeVaseAbstract)
	}
	if plantabstractFrom.StoolAbstract != nil {
		plantabstractTo.StoolAbstract = GongCopyBranchStoolAbstract(mapOrigCopy, plantabstractFrom.StoolAbstract)
	}
	if plantabstractFrom.ClockAbstract != nil {
		plantabstractTo.ClockAbstract = GongCopyBranchClockAbstract(mapOrigCopy, plantabstractFrom.ClockAbstract)
	}
	if plantabstractFrom.MusicAbstract != nil {
		plantabstractTo.MusicAbstract = GongCopyBranchMusicAbstract(mapOrigCopy, plantabstractFrom.MusicAbstract)
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

	// plantcircumferenceshapeFrom has already been copied
	if _plantcircumferenceshapeTo, ok := mapOrigCopy[plantcircumferenceshapeFrom]; ok {
		plantcircumferenceshapeTo = _plantcircumferenceshapeTo.(*PlantCircumferenceShape)
		return
	}

	plantcircumferenceshapeTo = new(PlantCircumferenceShape)
	mapOrigCopy[plantcircumferenceshapeFrom] = plantcircumferenceshapeTo
	plantcircumferenceshapeFrom.GongCopyBasicFields(plantcircumferenceshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPointsAndLines3DShape(mapOrigCopy map[any]any, pointsandlines3dshapeFrom *PointsAndLines3DShape) (pointsandlines3dshapeTo *PointsAndLines3DShape) {

	// pointsandlines3dshapeFrom has already been copied
	if _pointsandlines3dshapeTo, ok := mapOrigCopy[pointsandlines3dshapeFrom]; ok {
		pointsandlines3dshapeTo = _pointsandlines3dshapeTo.(*PointsAndLines3DShape)
		return
	}

	pointsandlines3dshapeTo = new(PointsAndLines3DShape)
	mapOrigCopy[pointsandlines3dshapeFrom] = pointsandlines3dshapeTo
	pointsandlines3dshapeFrom.GongCopyBasicFields(pointsandlines3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPxShape(mapOrigCopy map[any]any, pxshapeFrom *PxShape) (pxshapeTo *PxShape) {

	// pxshapeFrom has already been copied
	if _pxshapeTo, ok := mapOrigCopy[pxshapeFrom]; ok {
		pxshapeTo = _pxshapeTo.(*PxShape)
		return
	}

	pxshapeTo = new(PxShape)
	mapOrigCopy[pxshapeFrom] = pxshapeTo
	pxshapeFrom.GongCopyBasicFields(pxshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRendered3DShape(mapOrigCopy map[any]any, rendered3dshapeFrom *Rendered3DShape) (rendered3dshapeTo *Rendered3DShape) {

	// rendered3dshapeFrom has already been copied
	if _rendered3dshapeTo, ok := mapOrigCopy[rendered3dshapeFrom]; ok {
		rendered3dshapeTo = _rendered3dshapeTo.(*Rendered3DShape)
		return
	}

	rendered3dshapeTo = new(Rendered3DShape)
	mapOrigCopy[rendered3dshapeFrom] = rendered3dshapeTo
	rendered3dshapeFrom.GongCopyBasicFields(rendered3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRhombusShape(mapOrigCopy map[any]any, rhombusshapeFrom *RhombusShape) (rhombusshapeTo *RhombusShape) {

	// rhombusshapeFrom has already been copied
	if _rhombusshapeTo, ok := mapOrigCopy[rhombusshapeFrom]; ok {
		rhombusshapeTo = _rhombusshapeTo.(*RhombusShape)
		return
	}

	rhombusshapeTo = new(RhombusShape)
	mapOrigCopy[rhombusshapeFrom] = rhombusshapeTo
	rhombusshapeFrom.GongCopyBasicFields(rhombusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRhombusStuff(mapOrigCopy map[any]any, rhombusstuffFrom *RhombusStuff) (rhombusstuffTo *RhombusStuff) {

	// rhombusstuffFrom has already been copied
	if _rhombusstuffTo, ok := mapOrigCopy[rhombusstuffFrom]; ok {
		rhombusstuffTo = _rhombusstuffTo.(*RhombusStuff)
		return
	}

	rhombusstuffTo = new(RhombusStuff)
	mapOrigCopy[rhombusstuffFrom] = rhombusstuffTo
	rhombusstuffFrom.GongCopyBasicFields(rhombusstuffTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRotatedRhombusGridShape(mapOrigCopy map[any]any, rotatedrhombusgridshapeFrom *RotatedRhombusGridShape) (rotatedrhombusgridshapeTo *RotatedRhombusGridShape) {

	// rotatedrhombusgridshapeFrom has already been copied
	if _rotatedrhombusgridshapeTo, ok := mapOrigCopy[rotatedrhombusgridshapeFrom]; ok {
		rotatedrhombusgridshapeTo = _rotatedrhombusgridshapeTo.(*RotatedRhombusGridShape)
		return
	}

	rotatedrhombusgridshapeTo = new(RotatedRhombusGridShape)
	mapOrigCopy[rotatedrhombusgridshapeFrom] = rotatedrhombusgridshapeTo
	rotatedrhombusgridshapeFrom.GongCopyBasicFields(rotatedrhombusgridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRotatedRhombusShape(mapOrigCopy map[any]any, rotatedrhombusshapeFrom *RotatedRhombusShape) (rotatedrhombusshapeTo *RotatedRhombusShape) {

	// rotatedrhombusshapeFrom has already been copied
	if _rotatedrhombusshapeTo, ok := mapOrigCopy[rotatedrhombusshapeFrom]; ok {
		rotatedrhombusshapeTo = _rotatedrhombusshapeTo.(*RotatedRhombusShape)
		return
	}

	rotatedrhombusshapeTo = new(RotatedRhombusShape)
	mapOrigCopy[rotatedrhombusshapeFrom] = rotatedrhombusshapeTo
	rotatedrhombusshapeFrom.GongCopyBasicFields(rotatedrhombusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRotatedSampledPoints3DShape(mapOrigCopy map[any]any, rotatedsampledpoints3dshapeFrom *RotatedSampledPoints3DShape) (rotatedsampledpoints3dshapeTo *RotatedSampledPoints3DShape) {

	// rotatedsampledpoints3dshapeFrom has already been copied
	if _rotatedsampledpoints3dshapeTo, ok := mapOrigCopy[rotatedsampledpoints3dshapeFrom]; ok {
		rotatedsampledpoints3dshapeTo = _rotatedsampledpoints3dshapeTo.(*RotatedSampledPoints3DShape)
		return
	}

	rotatedsampledpoints3dshapeTo = new(RotatedSampledPoints3DShape)
	mapOrigCopy[rotatedsampledpoints3dshapeFrom] = rotatedsampledpoints3dshapeTo
	rotatedsampledpoints3dshapeFrom.GongCopyBasicFields(rotatedsampledpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRotatedSeatAndLegs3DShape(mapOrigCopy map[any]any, rotatedseatandlegs3dshapeFrom *RotatedSeatAndLegs3DShape) (rotatedseatandlegs3dshapeTo *RotatedSeatAndLegs3DShape) {

	// rotatedseatandlegs3dshapeFrom has already been copied
	if _rotatedseatandlegs3dshapeTo, ok := mapOrigCopy[rotatedseatandlegs3dshapeFrom]; ok {
		rotatedseatandlegs3dshapeTo = _rotatedseatandlegs3dshapeTo.(*RotatedSeatAndLegs3DShape)
		return
	}

	rotatedseatandlegs3dshapeTo = new(RotatedSeatAndLegs3DShape)
	mapOrigCopy[rotatedseatandlegs3dshapeFrom] = rotatedseatandlegs3dshapeTo
	rotatedseatandlegs3dshapeFrom.GongCopyBasicFields(rotatedseatandlegs3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSampledPoints3DShape(mapOrigCopy map[any]any, sampledpoints3dshapeFrom *SampledPoints3DShape) (sampledpoints3dshapeTo *SampledPoints3DShape) {

	// sampledpoints3dshapeFrom has already been copied
	if _sampledpoints3dshapeTo, ok := mapOrigCopy[sampledpoints3dshapeFrom]; ok {
		sampledpoints3dshapeTo = _sampledpoints3dshapeTo.(*SampledPoints3DShape)
		return
	}

	sampledpoints3dshapeTo = new(SampledPoints3DShape)
	mapOrigCopy[sampledpoints3dshapeFrom] = sampledpoints3dshapeTo
	sampledpoints3dshapeFrom.GongCopyBasicFields(sampledpoints3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSeat3DShape(mapOrigCopy map[any]any, seat3dshapeFrom *Seat3DShape) (seat3dshapeTo *Seat3DShape) {

	// seat3dshapeFrom has already been copied
	if _seat3dshapeTo, ok := mapOrigCopy[seat3dshapeFrom]; ok {
		seat3dshapeTo = _seat3dshapeTo.(*Seat3DShape)
		return
	}

	seat3dshapeTo = new(Seat3DShape)
	mapOrigCopy[seat3dshapeFrom] = seat3dshapeTo
	seat3dshapeFrom.GongCopyBasicFields(seat3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSeatAndLegs3DShape(mapOrigCopy map[any]any, seatandlegs3dshapeFrom *SeatAndLegs3DShape) (seatandlegs3dshapeTo *SeatAndLegs3DShape) {

	// seatandlegs3dshapeFrom has already been copied
	if _seatandlegs3dshapeTo, ok := mapOrigCopy[seatandlegs3dshapeFrom]; ok {
		seatandlegs3dshapeTo = _seatandlegs3dshapeTo.(*SeatAndLegs3DShape)
		return
	}

	seatandlegs3dshapeTo = new(SeatAndLegs3DShape)
	mapOrigCopy[seatandlegs3dshapeFrom] = seatandlegs3dshapeTo
	seatandlegs3dshapeFrom.GongCopyBasicFields(seatandlegs3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSeatBottomCurveShape(mapOrigCopy map[any]any, seatbottomcurveshapeFrom *SeatBottomCurveShape) (seatbottomcurveshapeTo *SeatBottomCurveShape) {

	// seatbottomcurveshapeFrom has already been copied
	if _seatbottomcurveshapeTo, ok := mapOrigCopy[seatbottomcurveshapeFrom]; ok {
		seatbottomcurveshapeTo = _seatbottomcurveshapeTo.(*SeatBottomCurveShape)
		return
	}

	seatbottomcurveshapeTo = new(SeatBottomCurveShape)
	mapOrigCopy[seatbottomcurveshapeFrom] = seatbottomcurveshapeTo
	seatbottomcurveshapeFrom.GongCopyBasicFields(seatbottomcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSeatTopCurveShape(mapOrigCopy map[any]any, seattopcurveshapeFrom *SeatTopCurveShape) (seattopcurveshapeTo *SeatTopCurveShape) {

	// seattopcurveshapeFrom has already been copied
	if _seattopcurveshapeTo, ok := mapOrigCopy[seattopcurveshapeFrom]; ok {
		seattopcurveshapeTo = _seattopcurveshapeTo.(*SeatTopCurveShape)
		return
	}

	seattopcurveshapeTo = new(SeatTopCurveShape)
	mapOrigCopy[seattopcurveshapeFrom] = seattopcurveshapeTo
	seattopcurveshapeFrom.GongCopyBasicFields(seattopcurveshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedBottomTopStartArcShape(mapOrigCopy map[any]any, shiftedbottomtopstartarcshapeFrom *ShiftedBottomTopStartArcShape) (shiftedbottomtopstartarcshapeTo *ShiftedBottomTopStartArcShape) {

	// shiftedbottomtopstartarcshapeFrom has already been copied
	if _shiftedbottomtopstartarcshapeTo, ok := mapOrigCopy[shiftedbottomtopstartarcshapeFrom]; ok {
		shiftedbottomtopstartarcshapeTo = _shiftedbottomtopstartarcshapeTo.(*ShiftedBottomTopStartArcShape)
		return
	}

	shiftedbottomtopstartarcshapeTo = new(ShiftedBottomTopStartArcShape)
	mapOrigCopy[shiftedbottomtopstartarcshapeFrom] = shiftedbottomtopstartarcshapeTo
	shiftedbottomtopstartarcshapeFrom.GongCopyBasicFields(shiftedbottomtopstartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedBottomTopStartArcShapeGrid(mapOrigCopy map[any]any, shiftedbottomtopstartarcshapegridFrom *ShiftedBottomTopStartArcShapeGrid) (shiftedbottomtopstartarcshapegridTo *ShiftedBottomTopStartArcShapeGrid) {

	// shiftedbottomtopstartarcshapegridFrom has already been copied
	if _shiftedbottomtopstartarcshapegridTo, ok := mapOrigCopy[shiftedbottomtopstartarcshapegridFrom]; ok {
		shiftedbottomtopstartarcshapegridTo = _shiftedbottomtopstartarcshapegridTo.(*ShiftedBottomTopStartArcShapeGrid)
		return
	}

	shiftedbottomtopstartarcshapegridTo = new(ShiftedBottomTopStartArcShapeGrid)
	mapOrigCopy[shiftedbottomtopstartarcshapegridFrom] = shiftedbottomtopstartarcshapegridTo
	shiftedbottomtopstartarcshapegridFrom.GongCopyBasicFields(shiftedbottomtopstartarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftGrowthCurve2DRibbon(mapOrigCopy map[any]any, shiftedleftgrowthcurve2dribbonFrom *ShiftedLeftGrowthCurve2DRibbon) (shiftedleftgrowthcurve2dribbonTo *ShiftedLeftGrowthCurve2DRibbon) {

	// shiftedleftgrowthcurve2dribbonFrom has already been copied
	if _shiftedleftgrowthcurve2dribbonTo, ok := mapOrigCopy[shiftedleftgrowthcurve2dribbonFrom]; ok {
		shiftedleftgrowthcurve2dribbonTo = _shiftedleftgrowthcurve2dribbonTo.(*ShiftedLeftGrowthCurve2DRibbon)
		return
	}

	shiftedleftgrowthcurve2dribbonTo = new(ShiftedLeftGrowthCurve2DRibbon)
	mapOrigCopy[shiftedleftgrowthcurve2dribbonFrom] = shiftedleftgrowthcurve2dribbonTo
	shiftedleftgrowthcurve2dribbonFrom.GongCopyBasicFields(shiftedleftgrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, shiftedleftgrowthcurve2dribbonendshapeFrom *ShiftedLeftGrowthCurve2DRibbonEndShape) (shiftedleftgrowthcurve2dribbonendshapeTo *ShiftedLeftGrowthCurve2DRibbonEndShape) {

	// shiftedleftgrowthcurve2dribbonendshapeFrom has already been copied
	if _shiftedleftgrowthcurve2dribbonendshapeTo, ok := mapOrigCopy[shiftedleftgrowthcurve2dribbonendshapeFrom]; ok {
		shiftedleftgrowthcurve2dribbonendshapeTo = _shiftedleftgrowthcurve2dribbonendshapeTo.(*ShiftedLeftGrowthCurve2DRibbonEndShape)
		return
	}

	shiftedleftgrowthcurve2dribbonendshapeTo = new(ShiftedLeftGrowthCurve2DRibbonEndShape)
	mapOrigCopy[shiftedleftgrowthcurve2dribbonendshapeFrom] = shiftedleftgrowthcurve2dribbonendshapeTo
	shiftedleftgrowthcurve2dribbonendshapeFrom.GongCopyBasicFields(shiftedleftgrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, shiftedleftgrowthcurve2dribbonstartshapeFrom *ShiftedLeftGrowthCurve2DRibbonStartShape) (shiftedleftgrowthcurve2dribbonstartshapeTo *ShiftedLeftGrowthCurve2DRibbonStartShape) {

	// shiftedleftgrowthcurve2dribbonstartshapeFrom has already been copied
	if _shiftedleftgrowthcurve2dribbonstartshapeTo, ok := mapOrigCopy[shiftedleftgrowthcurve2dribbonstartshapeFrom]; ok {
		shiftedleftgrowthcurve2dribbonstartshapeTo = _shiftedleftgrowthcurve2dribbonstartshapeTo.(*ShiftedLeftGrowthCurve2DRibbonStartShape)
		return
	}

	shiftedleftgrowthcurve2dribbonstartshapeTo = new(ShiftedLeftGrowthCurve2DRibbonStartShape)
	mapOrigCopy[shiftedleftgrowthcurve2dribbonstartshapeFrom] = shiftedleftgrowthcurve2dribbonstartshapeTo
	shiftedleftgrowthcurve2dribbonstartshapeFrom.GongCopyBasicFields(shiftedleftgrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftPartiallyGrowthCurve2DRibbon(mapOrigCopy map[any]any, shiftedleftpartiallygrowthcurve2dribbonFrom *ShiftedLeftPartiallyGrowthCurve2DRibbon) (shiftedleftpartiallygrowthcurve2dribbonTo *ShiftedLeftPartiallyGrowthCurve2DRibbon) {

	// shiftedleftpartiallygrowthcurve2dribbonFrom has already been copied
	if _shiftedleftpartiallygrowthcurve2dribbonTo, ok := mapOrigCopy[shiftedleftpartiallygrowthcurve2dribbonFrom]; ok {
		shiftedleftpartiallygrowthcurve2dribbonTo = _shiftedleftpartiallygrowthcurve2dribbonTo.(*ShiftedLeftPartiallyGrowthCurve2DRibbon)
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonTo = new(ShiftedLeftPartiallyGrowthCurve2DRibbon)
	mapOrigCopy[shiftedleftpartiallygrowthcurve2dribbonFrom] = shiftedleftpartiallygrowthcurve2dribbonTo
	shiftedleftpartiallygrowthcurve2dribbonFrom.GongCopyBasicFields(shiftedleftpartiallygrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, shiftedleftpartiallygrowthcurve2dribbonendshapeFrom *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) (shiftedleftpartiallygrowthcurve2dribbonendshapeTo *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) {

	// shiftedleftpartiallygrowthcurve2dribbonendshapeFrom has already been copied
	if _shiftedleftpartiallygrowthcurve2dribbonendshapeTo, ok := mapOrigCopy[shiftedleftpartiallygrowthcurve2dribbonendshapeFrom]; ok {
		shiftedleftpartiallygrowthcurve2dribbonendshapeTo = _shiftedleftpartiallygrowthcurve2dribbonendshapeTo.(*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonendshapeTo = new(ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
	mapOrigCopy[shiftedleftpartiallygrowthcurve2dribbonendshapeFrom] = shiftedleftpartiallygrowthcurve2dribbonendshapeTo
	shiftedleftpartiallygrowthcurve2dribbonendshapeFrom.GongCopyBasicFields(shiftedleftpartiallygrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, shiftedleftpartiallygrowthcurve2dribbonstartshapeFrom *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) (shiftedleftpartiallygrowthcurve2dribbonstartshapeTo *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) {

	// shiftedleftpartiallygrowthcurve2dribbonstartshapeFrom has already been copied
	if _shiftedleftpartiallygrowthcurve2dribbonstartshapeTo, ok := mapOrigCopy[shiftedleftpartiallygrowthcurve2dribbonstartshapeFrom]; ok {
		shiftedleftpartiallygrowthcurve2dribbonstartshapeTo = _shiftedleftpartiallygrowthcurve2dribbonstartshapeTo.(*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonstartshapeTo = new(ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
	mapOrigCopy[shiftedleftpartiallygrowthcurve2dribbonstartshapeFrom] = shiftedleftpartiallygrowthcurve2dribbonstartshapeTo
	shiftedleftpartiallygrowthcurve2dribbonstartshapeFrom.GongCopyBasicFields(shiftedleftpartiallygrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftStackGrowthCurveEndArcShape(mapOrigCopy map[any]any, shiftedleftstackgrowthcurveendarcshapeFrom *ShiftedLeftStackGrowthCurveEndArcShape) (shiftedleftstackgrowthcurveendarcshapeTo *ShiftedLeftStackGrowthCurveEndArcShape) {

	// shiftedleftstackgrowthcurveendarcshapeFrom has already been copied
	if _shiftedleftstackgrowthcurveendarcshapeTo, ok := mapOrigCopy[shiftedleftstackgrowthcurveendarcshapeFrom]; ok {
		shiftedleftstackgrowthcurveendarcshapeTo = _shiftedleftstackgrowthcurveendarcshapeTo.(*ShiftedLeftStackGrowthCurveEndArcShape)
		return
	}

	shiftedleftstackgrowthcurveendarcshapeTo = new(ShiftedLeftStackGrowthCurveEndArcShape)
	mapOrigCopy[shiftedleftstackgrowthcurveendarcshapeFrom] = shiftedleftstackgrowthcurveendarcshapeTo
	shiftedleftstackgrowthcurveendarcshapeFrom.GongCopyBasicFields(shiftedleftstackgrowthcurveendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftStackGrowthCurveStartArcShape(mapOrigCopy map[any]any, shiftedleftstackgrowthcurvestartarcshapeFrom *ShiftedLeftStackGrowthCurveStartArcShape) (shiftedleftstackgrowthcurvestartarcshapeTo *ShiftedLeftStackGrowthCurveStartArcShape) {

	// shiftedleftstackgrowthcurvestartarcshapeFrom has already been copied
	if _shiftedleftstackgrowthcurvestartarcshapeTo, ok := mapOrigCopy[shiftedleftstackgrowthcurvestartarcshapeFrom]; ok {
		shiftedleftstackgrowthcurvestartarcshapeTo = _shiftedleftstackgrowthcurvestartarcshapeTo.(*ShiftedLeftStackGrowthCurveStartArcShape)
		return
	}

	shiftedleftstackgrowthcurvestartarcshapeTo = new(ShiftedLeftStackGrowthCurveStartArcShape)
	mapOrigCopy[shiftedleftstackgrowthcurvestartarcshapeFrom] = shiftedleftstackgrowthcurvestartarcshapeTo
	shiftedleftstackgrowthcurvestartarcshapeFrom.GongCopyBasicFields(shiftedleftstackgrowthcurvestartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftStackNormalVector(mapOrigCopy map[any]any, shiftedleftstacknormalvectorFrom *ShiftedLeftStackNormalVector) (shiftedleftstacknormalvectorTo *ShiftedLeftStackNormalVector) {

	// shiftedleftstacknormalvectorFrom has already been copied
	if _shiftedleftstacknormalvectorTo, ok := mapOrigCopy[shiftedleftstacknormalvectorFrom]; ok {
		shiftedleftstacknormalvectorTo = _shiftedleftstacknormalvectorTo.(*ShiftedLeftStackNormalVector)
		return
	}

	shiftedleftstacknormalvectorTo = new(ShiftedLeftStackNormalVector)
	mapOrigCopy[shiftedleftstacknormalvectorFrom] = shiftedleftstacknormalvectorTo
	shiftedleftstacknormalvectorFrom.GongCopyBasicFields(shiftedleftstacknormalvectorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftStackOfGrowthCurve(mapOrigCopy map[any]any, shiftedleftstackofgrowthcurveFrom *ShiftedLeftStackOfGrowthCurve) (shiftedleftstackofgrowthcurveTo *ShiftedLeftStackOfGrowthCurve) {

	// shiftedleftstackofgrowthcurveFrom has already been copied
	if _shiftedleftstackofgrowthcurveTo, ok := mapOrigCopy[shiftedleftstackofgrowthcurveFrom]; ok {
		shiftedleftstackofgrowthcurveTo = _shiftedleftstackofgrowthcurveTo.(*ShiftedLeftStackOfGrowthCurve)
		return
	}

	shiftedleftstackofgrowthcurveTo = new(ShiftedLeftStackOfGrowthCurve)
	mapOrigCopy[shiftedleftstackofgrowthcurveFrom] = shiftedleftstackofgrowthcurveTo
	shiftedleftstackofgrowthcurveFrom.GongCopyBasicFields(shiftedleftstackofgrowthcurveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedLeftStackOfNormalVector(mapOrigCopy map[any]any, shiftedleftstackofnormalvectorFrom *ShiftedLeftStackOfNormalVector) (shiftedleftstackofnormalvectorTo *ShiftedLeftStackOfNormalVector) {

	// shiftedleftstackofnormalvectorFrom has already been copied
	if _shiftedleftstackofnormalvectorTo, ok := mapOrigCopy[shiftedleftstackofnormalvectorFrom]; ok {
		shiftedleftstackofnormalvectorTo = _shiftedleftstackofnormalvectorTo.(*ShiftedLeftStackOfNormalVector)
		return
	}

	shiftedleftstackofnormalvectorTo = new(ShiftedLeftStackOfNormalVector)
	mapOrigCopy[shiftedleftstackofnormalvectorFrom] = shiftedleftstackofnormalvectorTo
	shiftedleftstackofnormalvectorFrom.GongCopyBasicFields(shiftedleftstackofnormalvectorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedRightGrowthCurve2DRibbon(mapOrigCopy map[any]any, shiftedrightgrowthcurve2dribbonFrom *ShiftedRightGrowthCurve2DRibbon) (shiftedrightgrowthcurve2dribbonTo *ShiftedRightGrowthCurve2DRibbon) {

	// shiftedrightgrowthcurve2dribbonFrom has already been copied
	if _shiftedrightgrowthcurve2dribbonTo, ok := mapOrigCopy[shiftedrightgrowthcurve2dribbonFrom]; ok {
		shiftedrightgrowthcurve2dribbonTo = _shiftedrightgrowthcurve2dribbonTo.(*ShiftedRightGrowthCurve2DRibbon)
		return
	}

	shiftedrightgrowthcurve2dribbonTo = new(ShiftedRightGrowthCurve2DRibbon)
	mapOrigCopy[shiftedrightgrowthcurve2dribbonFrom] = shiftedrightgrowthcurve2dribbonTo
	shiftedrightgrowthcurve2dribbonFrom.GongCopyBasicFields(shiftedrightgrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedRightGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, shiftedrightgrowthcurve2dribbonendshapeFrom *ShiftedRightGrowthCurve2DRibbonEndShape) (shiftedrightgrowthcurve2dribbonendshapeTo *ShiftedRightGrowthCurve2DRibbonEndShape) {

	// shiftedrightgrowthcurve2dribbonendshapeFrom has already been copied
	if _shiftedrightgrowthcurve2dribbonendshapeTo, ok := mapOrigCopy[shiftedrightgrowthcurve2dribbonendshapeFrom]; ok {
		shiftedrightgrowthcurve2dribbonendshapeTo = _shiftedrightgrowthcurve2dribbonendshapeTo.(*ShiftedRightGrowthCurve2DRibbonEndShape)
		return
	}

	shiftedrightgrowthcurve2dribbonendshapeTo = new(ShiftedRightGrowthCurve2DRibbonEndShape)
	mapOrigCopy[shiftedrightgrowthcurve2dribbonendshapeFrom] = shiftedrightgrowthcurve2dribbonendshapeTo
	shiftedrightgrowthcurve2dribbonendshapeFrom.GongCopyBasicFields(shiftedrightgrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShiftedRightGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, shiftedrightgrowthcurve2dribbonstartshapeFrom *ShiftedRightGrowthCurve2DRibbonStartShape) (shiftedrightgrowthcurve2dribbonstartshapeTo *ShiftedRightGrowthCurve2DRibbonStartShape) {

	// shiftedrightgrowthcurve2dribbonstartshapeFrom has already been copied
	if _shiftedrightgrowthcurve2dribbonstartshapeTo, ok := mapOrigCopy[shiftedrightgrowthcurve2dribbonstartshapeFrom]; ok {
		shiftedrightgrowthcurve2dribbonstartshapeTo = _shiftedrightgrowthcurve2dribbonstartshapeTo.(*ShiftedRightGrowthCurve2DRibbonStartShape)
		return
	}

	shiftedrightgrowthcurve2dribbonstartshapeTo = new(ShiftedRightGrowthCurve2DRibbonStartShape)
	mapOrigCopy[shiftedrightgrowthcurve2dribbonstartshapeFrom] = shiftedrightgrowthcurve2dribbonstartshapeTo
	shiftedrightgrowthcurve2dribbonstartshapeFrom.GongCopyBasicFields(shiftedrightgrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackGrowthCurve2DEndHalfwayArcShape(mapOrigCopy map[any]any, stackgrowthcurve2dendhalfwayarcshapeFrom *StackGrowthCurve2DEndHalfwayArcShape) (stackgrowthcurve2dendhalfwayarcshapeTo *StackGrowthCurve2DEndHalfwayArcShape) {

	// stackgrowthcurve2dendhalfwayarcshapeFrom has already been copied
	if _stackgrowthcurve2dendhalfwayarcshapeTo, ok := mapOrigCopy[stackgrowthcurve2dendhalfwayarcshapeFrom]; ok {
		stackgrowthcurve2dendhalfwayarcshapeTo = _stackgrowthcurve2dendhalfwayarcshapeTo.(*StackGrowthCurve2DEndHalfwayArcShape)
		return
	}

	stackgrowthcurve2dendhalfwayarcshapeTo = new(StackGrowthCurve2DEndHalfwayArcShape)
	mapOrigCopy[stackgrowthcurve2dendhalfwayarcshapeFrom] = stackgrowthcurve2dendhalfwayarcshapeTo
	stackgrowthcurve2dendhalfwayarcshapeFrom.GongCopyBasicFields(stackgrowthcurve2dendhalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, stackgrowthcurve2dribbonendshapeFrom *StackGrowthCurve2DRibbonEndShape) (stackgrowthcurve2dribbonendshapeTo *StackGrowthCurve2DRibbonEndShape) {

	// stackgrowthcurve2dribbonendshapeFrom has already been copied
	if _stackgrowthcurve2dribbonendshapeTo, ok := mapOrigCopy[stackgrowthcurve2dribbonendshapeFrom]; ok {
		stackgrowthcurve2dribbonendshapeTo = _stackgrowthcurve2dribbonendshapeTo.(*StackGrowthCurve2DRibbonEndShape)
		return
	}

	stackgrowthcurve2dribbonendshapeTo = new(StackGrowthCurve2DRibbonEndShape)
	mapOrigCopy[stackgrowthcurve2dribbonendshapeFrom] = stackgrowthcurve2dribbonendshapeTo
	stackgrowthcurve2dribbonendshapeFrom.GongCopyBasicFields(stackgrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, stackgrowthcurve2dribbonstartshapeFrom *StackGrowthCurve2DRibbonStartShape) (stackgrowthcurve2dribbonstartshapeTo *StackGrowthCurve2DRibbonStartShape) {

	// stackgrowthcurve2dribbonstartshapeFrom has already been copied
	if _stackgrowthcurve2dribbonstartshapeTo, ok := mapOrigCopy[stackgrowthcurve2dribbonstartshapeFrom]; ok {
		stackgrowthcurve2dribbonstartshapeTo = _stackgrowthcurve2dribbonstartshapeTo.(*StackGrowthCurve2DRibbonStartShape)
		return
	}

	stackgrowthcurve2dribbonstartshapeTo = new(StackGrowthCurve2DRibbonStartShape)
	mapOrigCopy[stackgrowthcurve2dribbonstartshapeFrom] = stackgrowthcurve2dribbonstartshapeTo
	stackgrowthcurve2dribbonstartshapeFrom.GongCopyBasicFields(stackgrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackGrowthCurve2DStartHalfwayArcShape(mapOrigCopy map[any]any, stackgrowthcurve2dstarthalfwayarcshapeFrom *StackGrowthCurve2DStartHalfwayArcShape) (stackgrowthcurve2dstarthalfwayarcshapeTo *StackGrowthCurve2DStartHalfwayArcShape) {

	// stackgrowthcurve2dstarthalfwayarcshapeFrom has already been copied
	if _stackgrowthcurve2dstarthalfwayarcshapeTo, ok := mapOrigCopy[stackgrowthcurve2dstarthalfwayarcshapeFrom]; ok {
		stackgrowthcurve2dstarthalfwayarcshapeTo = _stackgrowthcurve2dstarthalfwayarcshapeTo.(*StackGrowthCurve2DStartHalfwayArcShape)
		return
	}

	stackgrowthcurve2dstarthalfwayarcshapeTo = new(StackGrowthCurve2DStartHalfwayArcShape)
	mapOrigCopy[stackgrowthcurve2dstarthalfwayarcshapeFrom] = stackgrowthcurve2dstarthalfwayarcshapeTo
	stackgrowthcurve2dstarthalfwayarcshapeFrom.GongCopyBasicFields(stackgrowthcurve2dstarthalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfGrowthCurve2D(mapOrigCopy map[any]any, stackofgrowthcurve2dFrom *StackOfGrowthCurve2D) (stackofgrowthcurve2dTo *StackOfGrowthCurve2D) {

	// stackofgrowthcurve2dFrom has already been copied
	if _stackofgrowthcurve2dTo, ok := mapOrigCopy[stackofgrowthcurve2dFrom]; ok {
		stackofgrowthcurve2dTo = _stackofgrowthcurve2dTo.(*StackOfGrowthCurve2D)
		return
	}

	stackofgrowthcurve2dTo = new(StackOfGrowthCurve2D)
	mapOrigCopy[stackofgrowthcurve2dFrom] = stackofgrowthcurve2dTo
	stackofgrowthcurve2dFrom.GongCopyBasicFields(stackofgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfGrowthCurve2DByGrowthVector(mapOrigCopy map[any]any, stackofgrowthcurve2dbygrowthvectorFrom *StackOfGrowthCurve2DByGrowthVector) (stackofgrowthcurve2dbygrowthvectorTo *StackOfGrowthCurve2DByGrowthVector) {

	// stackofgrowthcurve2dbygrowthvectorFrom has already been copied
	if _stackofgrowthcurve2dbygrowthvectorTo, ok := mapOrigCopy[stackofgrowthcurve2dbygrowthvectorFrom]; ok {
		stackofgrowthcurve2dbygrowthvectorTo = _stackofgrowthcurve2dbygrowthvectorTo.(*StackOfGrowthCurve2DByGrowthVector)
		return
	}

	stackofgrowthcurve2dbygrowthvectorTo = new(StackOfGrowthCurve2DByGrowthVector)
	mapOrigCopy[stackofgrowthcurve2dbygrowthvectorFrom] = stackofgrowthcurve2dbygrowthvectorTo
	stackofgrowthcurve2dbygrowthvectorFrom.GongCopyBasicFields(stackofgrowthcurve2dbygrowthvectorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfGrowthCurve2DRibbon(mapOrigCopy map[any]any, stackofgrowthcurve2dribbonFrom *StackOfGrowthCurve2DRibbon) (stackofgrowthcurve2dribbonTo *StackOfGrowthCurve2DRibbon) {

	// stackofgrowthcurve2dribbonFrom has already been copied
	if _stackofgrowthcurve2dribbonTo, ok := mapOrigCopy[stackofgrowthcurve2dribbonFrom]; ok {
		stackofgrowthcurve2dribbonTo = _stackofgrowthcurve2dribbonTo.(*StackOfGrowthCurve2DRibbon)
		return
	}

	stackofgrowthcurve2dribbonTo = new(StackOfGrowthCurve2DRibbon)
	mapOrigCopy[stackofgrowthcurve2dribbonFrom] = stackofgrowthcurve2dribbonTo
	stackofgrowthcurve2dribbonFrom.GongCopyBasicFields(stackofgrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfPartiallyRotatedTorusShape(mapOrigCopy map[any]any, stackofpartiallyrotatedtorusshapeFrom *StackOfPartiallyRotatedTorusShape) (stackofpartiallyrotatedtorusshapeTo *StackOfPartiallyRotatedTorusShape) {

	// stackofpartiallyrotatedtorusshapeFrom has already been copied
	if _stackofpartiallyrotatedtorusshapeTo, ok := mapOrigCopy[stackofpartiallyrotatedtorusshapeFrom]; ok {
		stackofpartiallyrotatedtorusshapeTo = _stackofpartiallyrotatedtorusshapeTo.(*StackOfPartiallyRotatedTorusShape)
		return
	}

	stackofpartiallyrotatedtorusshapeTo = new(StackOfPartiallyRotatedTorusShape)
	mapOrigCopy[stackofpartiallyrotatedtorusshapeFrom] = stackofpartiallyrotatedtorusshapeTo
	stackofpartiallyrotatedtorusshapeFrom.GongCopyBasicFields(stackofpartiallyrotatedtorusshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfRotatedGrowthCurve2D(mapOrigCopy map[any]any, stackofrotatedgrowthcurve2dFrom *StackOfRotatedGrowthCurve2D) (stackofrotatedgrowthcurve2dTo *StackOfRotatedGrowthCurve2D) {

	// stackofrotatedgrowthcurve2dFrom has already been copied
	if _stackofrotatedgrowthcurve2dTo, ok := mapOrigCopy[stackofrotatedgrowthcurve2dFrom]; ok {
		stackofrotatedgrowthcurve2dTo = _stackofrotatedgrowthcurve2dTo.(*StackOfRotatedGrowthCurve2D)
		return
	}

	stackofrotatedgrowthcurve2dTo = new(StackOfRotatedGrowthCurve2D)
	mapOrigCopy[stackofrotatedgrowthcurve2dFrom] = stackofrotatedgrowthcurve2dTo
	stackofrotatedgrowthcurve2dFrom.GongCopyBasicFields(stackofrotatedgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackOfRotatedGrowthCurve2DRibbon(mapOrigCopy map[any]any, stackofrotatedgrowthcurve2dribbonFrom *StackOfRotatedGrowthCurve2DRibbon) (stackofrotatedgrowthcurve2dribbonTo *StackOfRotatedGrowthCurve2DRibbon) {

	// stackofrotatedgrowthcurve2dribbonFrom has already been copied
	if _stackofrotatedgrowthcurve2dribbonTo, ok := mapOrigCopy[stackofrotatedgrowthcurve2dribbonFrom]; ok {
		stackofrotatedgrowthcurve2dribbonTo = _stackofrotatedgrowthcurve2dribbonTo.(*StackOfRotatedGrowthCurve2DRibbon)
		return
	}

	stackofrotatedgrowthcurve2dribbonTo = new(StackOfRotatedGrowthCurve2DRibbon)
	mapOrigCopy[stackofrotatedgrowthcurve2dribbonFrom] = stackofrotatedgrowthcurve2dribbonTo
	stackofrotatedgrowthcurve2dribbonFrom.GongCopyBasicFields(stackofrotatedgrowthcurve2dribbonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackRotatedGrowthCurve2DEndArcShape(mapOrigCopy map[any]any, stackrotatedgrowthcurve2dendarcshapeFrom *StackRotatedGrowthCurve2DEndArcShape) (stackrotatedgrowthcurve2dendarcshapeTo *StackRotatedGrowthCurve2DEndArcShape) {

	// stackrotatedgrowthcurve2dendarcshapeFrom has already been copied
	if _stackrotatedgrowthcurve2dendarcshapeTo, ok := mapOrigCopy[stackrotatedgrowthcurve2dendarcshapeFrom]; ok {
		stackrotatedgrowthcurve2dendarcshapeTo = _stackrotatedgrowthcurve2dendarcshapeTo.(*StackRotatedGrowthCurve2DEndArcShape)
		return
	}

	stackrotatedgrowthcurve2dendarcshapeTo = new(StackRotatedGrowthCurve2DEndArcShape)
	mapOrigCopy[stackrotatedgrowthcurve2dendarcshapeFrom] = stackrotatedgrowthcurve2dendarcshapeTo
	stackrotatedgrowthcurve2dendarcshapeFrom.GongCopyBasicFields(stackrotatedgrowthcurve2dendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackRotatedGrowthCurve2DRibbonEndShape(mapOrigCopy map[any]any, stackrotatedgrowthcurve2dribbonendshapeFrom *StackRotatedGrowthCurve2DRibbonEndShape) (stackrotatedgrowthcurve2dribbonendshapeTo *StackRotatedGrowthCurve2DRibbonEndShape) {

	// stackrotatedgrowthcurve2dribbonendshapeFrom has already been copied
	if _stackrotatedgrowthcurve2dribbonendshapeTo, ok := mapOrigCopy[stackrotatedgrowthcurve2dribbonendshapeFrom]; ok {
		stackrotatedgrowthcurve2dribbonendshapeTo = _stackrotatedgrowthcurve2dribbonendshapeTo.(*StackRotatedGrowthCurve2DRibbonEndShape)
		return
	}

	stackrotatedgrowthcurve2dribbonendshapeTo = new(StackRotatedGrowthCurve2DRibbonEndShape)
	mapOrigCopy[stackrotatedgrowthcurve2dribbonendshapeFrom] = stackrotatedgrowthcurve2dribbonendshapeTo
	stackrotatedgrowthcurve2dribbonendshapeFrom.GongCopyBasicFields(stackrotatedgrowthcurve2dribbonendshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackRotatedGrowthCurve2DRibbonStartShape(mapOrigCopy map[any]any, stackrotatedgrowthcurve2dribbonstartshapeFrom *StackRotatedGrowthCurve2DRibbonStartShape) (stackrotatedgrowthcurve2dribbonstartshapeTo *StackRotatedGrowthCurve2DRibbonStartShape) {

	// stackrotatedgrowthcurve2dribbonstartshapeFrom has already been copied
	if _stackrotatedgrowthcurve2dribbonstartshapeTo, ok := mapOrigCopy[stackrotatedgrowthcurve2dribbonstartshapeFrom]; ok {
		stackrotatedgrowthcurve2dribbonstartshapeTo = _stackrotatedgrowthcurve2dribbonstartshapeTo.(*StackRotatedGrowthCurve2DRibbonStartShape)
		return
	}

	stackrotatedgrowthcurve2dribbonstartshapeTo = new(StackRotatedGrowthCurve2DRibbonStartShape)
	mapOrigCopy[stackrotatedgrowthcurve2dribbonstartshapeFrom] = stackrotatedgrowthcurve2dribbonstartshapeTo
	stackrotatedgrowthcurve2dribbonstartshapeFrom.GongCopyBasicFields(stackrotatedgrowthcurve2dribbonstartshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStackRotatedGrowthCurve2DStartArcShape(mapOrigCopy map[any]any, stackrotatedgrowthcurve2dstartarcshapeFrom *StackRotatedGrowthCurve2DStartArcShape) (stackrotatedgrowthcurve2dstartarcshapeTo *StackRotatedGrowthCurve2DStartArcShape) {

	// stackrotatedgrowthcurve2dstartarcshapeFrom has already been copied
	if _stackrotatedgrowthcurve2dstartarcshapeTo, ok := mapOrigCopy[stackrotatedgrowthcurve2dstartarcshapeFrom]; ok {
		stackrotatedgrowthcurve2dstartarcshapeTo = _stackrotatedgrowthcurve2dstartarcshapeTo.(*StackRotatedGrowthCurve2DStartArcShape)
		return
	}

	stackrotatedgrowthcurve2dstartarcshapeTo = new(StackRotatedGrowthCurve2DStartArcShape)
	mapOrigCopy[stackrotatedgrowthcurve2dstartarcshapeFrom] = stackrotatedgrowthcurve2dstartarcshapeTo
	stackrotatedgrowthcurve2dstartarcshapeFrom.GongCopyBasicFields(stackrotatedgrowthcurve2dstartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStartArcShape(mapOrigCopy map[any]any, startarcshapeFrom *StartArcShape) (startarcshapeTo *StartArcShape) {

	// startarcshapeFrom has already been copied
	if _startarcshapeTo, ok := mapOrigCopy[startarcshapeFrom]; ok {
		startarcshapeTo = _startarcshapeTo.(*StartArcShape)
		return
	}

	startarcshapeTo = new(StartArcShape)
	mapOrigCopy[startarcshapeFrom] = startarcshapeTo
	startarcshapeFrom.GongCopyBasicFields(startarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStartArcShapeGrid(mapOrigCopy map[any]any, startarcshapegridFrom *StartArcShapeGrid) (startarcshapegridTo *StartArcShapeGrid) {

	// startarcshapegridFrom has already been copied
	if _startarcshapegridTo, ok := mapOrigCopy[startarcshapegridFrom]; ok {
		startarcshapegridTo = _startarcshapegridTo.(*StartArcShapeGrid)
		return
	}

	startarcshapegridTo = new(StartArcShapeGrid)
	mapOrigCopy[startarcshapegridFrom] = startarcshapegridTo
	startarcshapegridFrom.GongCopyBasicFields(startarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStartHalfwayArcShape(mapOrigCopy map[any]any, starthalfwayarcshapeFrom *StartHalfwayArcShape) (starthalfwayarcshapeTo *StartHalfwayArcShape) {

	// starthalfwayarcshapeFrom has already been copied
	if _starthalfwayarcshapeTo, ok := mapOrigCopy[starthalfwayarcshapeFrom]; ok {
		starthalfwayarcshapeTo = _starthalfwayarcshapeTo.(*StartHalfwayArcShape)
		return
	}

	starthalfwayarcshapeTo = new(StartHalfwayArcShape)
	mapOrigCopy[starthalfwayarcshapeFrom] = starthalfwayarcshapeTo
	starthalfwayarcshapeFrom.GongCopyBasicFields(starthalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStartHalfwayArcShapeGrid(mapOrigCopy map[any]any, starthalfwayarcshapegridFrom *StartHalfwayArcShapeGrid) (starthalfwayarcshapegridTo *StartHalfwayArcShapeGrid) {

	// starthalfwayarcshapegridFrom has already been copied
	if _starthalfwayarcshapegridTo, ok := mapOrigCopy[starthalfwayarcshapegridFrom]; ok {
		starthalfwayarcshapegridTo = _starthalfwayarcshapegridTo.(*StartHalfwayArcShapeGrid)
		return
	}

	starthalfwayarcshapegridTo = new(StartHalfwayArcShapeGrid)
	mapOrigCopy[starthalfwayarcshapegridFrom] = starthalfwayarcshapegridTo
	starthalfwayarcshapegridFrom.GongCopyBasicFields(starthalfwayarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStemCylinder3DShape(mapOrigCopy map[any]any, stemcylinder3dshapeFrom *StemCylinder3DShape) (stemcylinder3dshapeTo *StemCylinder3DShape) {

	// stemcylinder3dshapeFrom has already been copied
	if _stemcylinder3dshapeTo, ok := mapOrigCopy[stemcylinder3dshapeFrom]; ok {
		stemcylinder3dshapeTo = _stemcylinder3dshapeTo.(*StemCylinder3DShape)
		return
	}

	stemcylinder3dshapeTo = new(StemCylinder3DShape)
	mapOrigCopy[stemcylinder3dshapeFrom] = stemcylinder3dshapeTo
	stemcylinder3dshapeFrom.GongCopyBasicFields(stemcylinder3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStool2DDiagram(mapOrigCopy map[any]any, stool2ddiagramFrom *Stool2DDiagram) (stool2ddiagramTo *Stool2DDiagram) {

	// stool2ddiagramFrom has already been copied
	if _stool2ddiagramTo, ok := mapOrigCopy[stool2ddiagramFrom]; ok {
		stool2ddiagramTo = _stool2ddiagramTo.(*Stool2DDiagram)
		return
	}

	stool2ddiagramTo = new(Stool2DDiagram)
	mapOrigCopy[stool2ddiagramFrom] = stool2ddiagramTo
	stool2ddiagramFrom.GongCopyBasicFields(stool2ddiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStool3DDiagram(mapOrigCopy map[any]any, stool3ddiagramFrom *Stool3DDiagram) (stool3ddiagramTo *Stool3DDiagram) {

	// stool3ddiagramFrom has already been copied
	if _stool3ddiagramTo, ok := mapOrigCopy[stool3ddiagramFrom]; ok {
		stool3ddiagramTo = _stool3ddiagramTo.(*Stool3DDiagram)
		return
	}

	stool3ddiagramTo = new(Stool3DDiagram)
	mapOrigCopy[stool3ddiagramFrom] = stool3ddiagramTo
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

func GongCopyBranchStoolAbstract(mapOrigCopy map[any]any, stoolabstractFrom *StoolAbstract) (stoolabstractTo *StoolAbstract) {

	// stoolabstractFrom has already been copied
	if _stoolabstractTo, ok := mapOrigCopy[stoolabstractFrom]; ok {
		stoolabstractTo = _stoolabstractTo.(*StoolAbstract)
		return
	}

	stoolabstractTo = new(StoolAbstract)
	mapOrigCopy[stoolabstractFrom] = stoolabstractTo
	stoolabstractFrom.GongCopyBasicFields(stoolabstractTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTiledFloor3DShape(mapOrigCopy map[any]any, tiledfloor3dshapeFrom *TiledFloor3DShape) (tiledfloor3dshapeTo *TiledFloor3DShape) {

	// tiledfloor3dshapeFrom has already been copied
	if _tiledfloor3dshapeTo, ok := mapOrigCopy[tiledfloor3dshapeFrom]; ok {
		tiledfloor3dshapeTo = _tiledfloor3dshapeTo.(*TiledFloor3DShape)
		return
	}

	tiledfloor3dshapeTo = new(TiledFloor3DShape)
	mapOrigCopy[tiledfloor3dshapeFrom] = tiledfloor3dshapeTo
	tiledfloor3dshapeFrom.GongCopyBasicFields(tiledfloor3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopEndArcShape(mapOrigCopy map[any]any, topendarcshapeFrom *TopEndArcShape) (topendarcshapeTo *TopEndArcShape) {

	// topendarcshapeFrom has already been copied
	if _topendarcshapeTo, ok := mapOrigCopy[topendarcshapeFrom]; ok {
		topendarcshapeTo = _topendarcshapeTo.(*TopEndArcShape)
		return
	}

	topendarcshapeTo = new(TopEndArcShape)
	mapOrigCopy[topendarcshapeFrom] = topendarcshapeTo
	topendarcshapeFrom.GongCopyBasicFields(topendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopEndArcShapeGrid(mapOrigCopy map[any]any, topendarcshapegridFrom *TopEndArcShapeGrid) (topendarcshapegridTo *TopEndArcShapeGrid) {

	// topendarcshapegridFrom has already been copied
	if _topendarcshapegridTo, ok := mapOrigCopy[topendarcshapegridFrom]; ok {
		topendarcshapegridTo = _topendarcshapegridTo.(*TopEndArcShapeGrid)
		return
	}

	topendarcshapegridTo = new(TopEndArcShapeGrid)
	mapOrigCopy[topendarcshapegridFrom] = topendarcshapegridTo
	topendarcshapegridFrom.GongCopyBasicFields(topendarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopEndHalfwayArcShape(mapOrigCopy map[any]any, topendhalfwayarcshapeFrom *TopEndHalfwayArcShape) (topendhalfwayarcshapeTo *TopEndHalfwayArcShape) {

	// topendhalfwayarcshapeFrom has already been copied
	if _topendhalfwayarcshapeTo, ok := mapOrigCopy[topendhalfwayarcshapeFrom]; ok {
		topendhalfwayarcshapeTo = _topendhalfwayarcshapeTo.(*TopEndHalfwayArcShape)
		return
	}

	topendhalfwayarcshapeTo = new(TopEndHalfwayArcShape)
	mapOrigCopy[topendhalfwayarcshapeFrom] = topendhalfwayarcshapeTo
	topendhalfwayarcshapeFrom.GongCopyBasicFields(topendhalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopEndHalfwayArcShapeGrid(mapOrigCopy map[any]any, topendhalfwayarcshapegridFrom *TopEndHalfwayArcShapeGrid) (topendhalfwayarcshapegridTo *TopEndHalfwayArcShapeGrid) {

	// topendhalfwayarcshapegridFrom has already been copied
	if _topendhalfwayarcshapegridTo, ok := mapOrigCopy[topendhalfwayarcshapegridFrom]; ok {
		topendhalfwayarcshapegridTo = _topendhalfwayarcshapegridTo.(*TopEndHalfwayArcShapeGrid)
		return
	}

	topendhalfwayarcshapegridTo = new(TopEndHalfwayArcShapeGrid)
	mapOrigCopy[topendhalfwayarcshapegridFrom] = topendhalfwayarcshapegridTo
	topendhalfwayarcshapegridFrom.GongCopyBasicFields(topendhalfwayarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopGrowthCurve2D(mapOrigCopy map[any]any, topgrowthcurve2dFrom *TopGrowthCurve2D) (topgrowthcurve2dTo *TopGrowthCurve2D) {

	// topgrowthcurve2dFrom has already been copied
	if _topgrowthcurve2dTo, ok := mapOrigCopy[topgrowthcurve2dFrom]; ok {
		topgrowthcurve2dTo = _topgrowthcurve2dTo.(*TopGrowthCurve2D)
		return
	}

	topgrowthcurve2dTo = new(TopGrowthCurve2D)
	mapOrigCopy[topgrowthcurve2dFrom] = topgrowthcurve2dTo
	topgrowthcurve2dFrom.GongCopyBasicFields(topgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopMidArcVectorShape(mapOrigCopy map[any]any, topmidarcvectorshapeFrom *TopMidArcVectorShape) (topmidarcvectorshapeTo *TopMidArcVectorShape) {

	// topmidarcvectorshapeFrom has already been copied
	if _topmidarcvectorshapeTo, ok := mapOrigCopy[topmidarcvectorshapeFrom]; ok {
		topmidarcvectorshapeTo = _topmidarcvectorshapeTo.(*TopMidArcVectorShape)
		return
	}

	topmidarcvectorshapeTo = new(TopMidArcVectorShape)
	mapOrigCopy[topmidarcvectorshapeFrom] = topmidarcvectorshapeTo
	topmidarcvectorshapeFrom.GongCopyBasicFields(topmidarcvectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopMidArcVectorShapeGrid(mapOrigCopy map[any]any, topmidarcvectorshapegridFrom *TopMidArcVectorShapeGrid) (topmidarcvectorshapegridTo *TopMidArcVectorShapeGrid) {

	// topmidarcvectorshapegridFrom has already been copied
	if _topmidarcvectorshapegridTo, ok := mapOrigCopy[topmidarcvectorshapegridFrom]; ok {
		topmidarcvectorshapegridTo = _topmidarcvectorshapegridTo.(*TopMidArcVectorShapeGrid)
		return
	}

	topmidarcvectorshapegridTo = new(TopMidArcVectorShapeGrid)
	mapOrigCopy[topmidarcvectorshapegridFrom] = topmidarcvectorshapegridTo
	topmidarcvectorshapegridFrom.GongCopyBasicFields(topmidarcvectorshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStackGrowthCurve2DEndHalfwayArcShape(mapOrigCopy map[any]any, topstackgrowthcurve2dendhalfwayarcshapeFrom *TopStackGrowthCurve2DEndHalfwayArcShape) (topstackgrowthcurve2dendhalfwayarcshapeTo *TopStackGrowthCurve2DEndHalfwayArcShape) {

	// topstackgrowthcurve2dendhalfwayarcshapeFrom has already been copied
	if _topstackgrowthcurve2dendhalfwayarcshapeTo, ok := mapOrigCopy[topstackgrowthcurve2dendhalfwayarcshapeFrom]; ok {
		topstackgrowthcurve2dendhalfwayarcshapeTo = _topstackgrowthcurve2dendhalfwayarcshapeTo.(*TopStackGrowthCurve2DEndHalfwayArcShape)
		return
	}

	topstackgrowthcurve2dendhalfwayarcshapeTo = new(TopStackGrowthCurve2DEndHalfwayArcShape)
	mapOrigCopy[topstackgrowthcurve2dendhalfwayarcshapeFrom] = topstackgrowthcurve2dendhalfwayarcshapeTo
	topstackgrowthcurve2dendhalfwayarcshapeFrom.GongCopyBasicFields(topstackgrowthcurve2dendhalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStackGrowthCurve2DStartHalfwayArcShape(mapOrigCopy map[any]any, topstackgrowthcurve2dstarthalfwayarcshapeFrom *TopStackGrowthCurve2DStartHalfwayArcShape) (topstackgrowthcurve2dstarthalfwayarcshapeTo *TopStackGrowthCurve2DStartHalfwayArcShape) {

	// topstackgrowthcurve2dstarthalfwayarcshapeFrom has already been copied
	if _topstackgrowthcurve2dstarthalfwayarcshapeTo, ok := mapOrigCopy[topstackgrowthcurve2dstarthalfwayarcshapeFrom]; ok {
		topstackgrowthcurve2dstarthalfwayarcshapeTo = _topstackgrowthcurve2dstarthalfwayarcshapeTo.(*TopStackGrowthCurve2DStartHalfwayArcShape)
		return
	}

	topstackgrowthcurve2dstarthalfwayarcshapeTo = new(TopStackGrowthCurve2DStartHalfwayArcShape)
	mapOrigCopy[topstackgrowthcurve2dstarthalfwayarcshapeFrom] = topstackgrowthcurve2dstarthalfwayarcshapeTo
	topstackgrowthcurve2dstarthalfwayarcshapeFrom.GongCopyBasicFields(topstackgrowthcurve2dstarthalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStackOfGrowthCurve2D(mapOrigCopy map[any]any, topstackofgrowthcurve2dFrom *TopStackOfGrowthCurve2D) (topstackofgrowthcurve2dTo *TopStackOfGrowthCurve2D) {

	// topstackofgrowthcurve2dFrom has already been copied
	if _topstackofgrowthcurve2dTo, ok := mapOrigCopy[topstackofgrowthcurve2dFrom]; ok {
		topstackofgrowthcurve2dTo = _topstackofgrowthcurve2dTo.(*TopStackOfGrowthCurve2D)
		return
	}

	topstackofgrowthcurve2dTo = new(TopStackOfGrowthCurve2D)
	mapOrigCopy[topstackofgrowthcurve2dFrom] = topstackofgrowthcurve2dTo
	topstackofgrowthcurve2dFrom.GongCopyBasicFields(topstackofgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStackOfRotatedGrowthCurve2D(mapOrigCopy map[any]any, topstackofrotatedgrowthcurve2dFrom *TopStackOfRotatedGrowthCurve2D) (topstackofrotatedgrowthcurve2dTo *TopStackOfRotatedGrowthCurve2D) {

	// topstackofrotatedgrowthcurve2dFrom has already been copied
	if _topstackofrotatedgrowthcurve2dTo, ok := mapOrigCopy[topstackofrotatedgrowthcurve2dFrom]; ok {
		topstackofrotatedgrowthcurve2dTo = _topstackofrotatedgrowthcurve2dTo.(*TopStackOfRotatedGrowthCurve2D)
		return
	}

	topstackofrotatedgrowthcurve2dTo = new(TopStackOfRotatedGrowthCurve2D)
	mapOrigCopy[topstackofrotatedgrowthcurve2dFrom] = topstackofrotatedgrowthcurve2dTo
	topstackofrotatedgrowthcurve2dFrom.GongCopyBasicFields(topstackofrotatedgrowthcurve2dTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStackOfRotatedGrowthCurve2DEndArcShape(mapOrigCopy map[any]any, topstackofrotatedgrowthcurve2dendarcshapeFrom *TopStackOfRotatedGrowthCurve2DEndArcShape) (topstackofrotatedgrowthcurve2dendarcshapeTo *TopStackOfRotatedGrowthCurve2DEndArcShape) {

	// topstackofrotatedgrowthcurve2dendarcshapeFrom has already been copied
	if _topstackofrotatedgrowthcurve2dendarcshapeTo, ok := mapOrigCopy[topstackofrotatedgrowthcurve2dendarcshapeFrom]; ok {
		topstackofrotatedgrowthcurve2dendarcshapeTo = _topstackofrotatedgrowthcurve2dendarcshapeTo.(*TopStackOfRotatedGrowthCurve2DEndArcShape)
		return
	}

	topstackofrotatedgrowthcurve2dendarcshapeTo = new(TopStackOfRotatedGrowthCurve2DEndArcShape)
	mapOrigCopy[topstackofrotatedgrowthcurve2dendarcshapeFrom] = topstackofrotatedgrowthcurve2dendarcshapeTo
	topstackofrotatedgrowthcurve2dendarcshapeFrom.GongCopyBasicFields(topstackofrotatedgrowthcurve2dendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStackOfRotatedGrowthCurve2DStartArcShape(mapOrigCopy map[any]any, topstackofrotatedgrowthcurve2dstartarcshapeFrom *TopStackOfRotatedGrowthCurve2DStartArcShape) (topstackofrotatedgrowthcurve2dstartarcshapeTo *TopStackOfRotatedGrowthCurve2DStartArcShape) {

	// topstackofrotatedgrowthcurve2dstartarcshapeFrom has already been copied
	if _topstackofrotatedgrowthcurve2dstartarcshapeTo, ok := mapOrigCopy[topstackofrotatedgrowthcurve2dstartarcshapeFrom]; ok {
		topstackofrotatedgrowthcurve2dstartarcshapeTo = _topstackofrotatedgrowthcurve2dstartarcshapeTo.(*TopStackOfRotatedGrowthCurve2DStartArcShape)
		return
	}

	topstackofrotatedgrowthcurve2dstartarcshapeTo = new(TopStackOfRotatedGrowthCurve2DStartArcShape)
	mapOrigCopy[topstackofrotatedgrowthcurve2dstartarcshapeFrom] = topstackofrotatedgrowthcurve2dstartarcshapeTo
	topstackofrotatedgrowthcurve2dstartarcshapeFrom.GongCopyBasicFields(topstackofrotatedgrowthcurve2dstartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStartArcShape(mapOrigCopy map[any]any, topstartarcshapeFrom *TopStartArcShape) (topstartarcshapeTo *TopStartArcShape) {

	// topstartarcshapeFrom has already been copied
	if _topstartarcshapeTo, ok := mapOrigCopy[topstartarcshapeFrom]; ok {
		topstartarcshapeTo = _topstartarcshapeTo.(*TopStartArcShape)
		return
	}

	topstartarcshapeTo = new(TopStartArcShape)
	mapOrigCopy[topstartarcshapeFrom] = topstartarcshapeTo
	topstartarcshapeFrom.GongCopyBasicFields(topstartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStartArcShapeGrid(mapOrigCopy map[any]any, topstartarcshapegridFrom *TopStartArcShapeGrid) (topstartarcshapegridTo *TopStartArcShapeGrid) {

	// topstartarcshapegridFrom has already been copied
	if _topstartarcshapegridTo, ok := mapOrigCopy[topstartarcshapegridFrom]; ok {
		topstartarcshapegridTo = _topstartarcshapegridTo.(*TopStartArcShapeGrid)
		return
	}

	topstartarcshapegridTo = new(TopStartArcShapeGrid)
	mapOrigCopy[topstartarcshapegridFrom] = topstartarcshapegridTo
	topstartarcshapegridFrom.GongCopyBasicFields(topstartarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStartHalfwayArcShape(mapOrigCopy map[any]any, topstarthalfwayarcshapeFrom *TopStartHalfwayArcShape) (topstarthalfwayarcshapeTo *TopStartHalfwayArcShape) {

	// topstarthalfwayarcshapeFrom has already been copied
	if _topstarthalfwayarcshapeTo, ok := mapOrigCopy[topstarthalfwayarcshapeFrom]; ok {
		topstarthalfwayarcshapeTo = _topstarthalfwayarcshapeTo.(*TopStartHalfwayArcShape)
		return
	}

	topstarthalfwayarcshapeTo = new(TopStartHalfwayArcShape)
	mapOrigCopy[topstarthalfwayarcshapeFrom] = topstarthalfwayarcshapeTo
	topstarthalfwayarcshapeFrom.GongCopyBasicFields(topstarthalfwayarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTopStartHalfwayArcShapeGrid(mapOrigCopy map[any]any, topstarthalfwayarcshapegridFrom *TopStartHalfwayArcShapeGrid) (topstarthalfwayarcshapegridTo *TopStartHalfwayArcShapeGrid) {

	// topstarthalfwayarcshapegridFrom has already been copied
	if _topstarthalfwayarcshapegridTo, ok := mapOrigCopy[topstarthalfwayarcshapegridFrom]; ok {
		topstarthalfwayarcshapegridTo = _topstarthalfwayarcshapegridTo.(*TopStartHalfwayArcShapeGrid)
		return
	}

	topstarthalfwayarcshapegridTo = new(TopStartHalfwayArcShapeGrid)
	mapOrigCopy[topstarthalfwayarcshapegridFrom] = topstarthalfwayarcshapegridTo
	topstarthalfwayarcshapegridFrom.GongCopyBasicFields(topstarthalfwayarcshapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTorus3DShape(mapOrigCopy map[any]any, torus3dshapeFrom *Torus3DShape) (torus3dshapeTo *Torus3DShape) {

	// torus3dshapeFrom has already been copied
	if _torus3dshapeTo, ok := mapOrigCopy[torus3dshapeFrom]; ok {
		torus3dshapeTo = _torus3dshapeTo.(*Torus3DShape)
		return
	}

	torus3dshapeTo = new(Torus3DShape)
	mapOrigCopy[torus3dshapeFrom] = torus3dshapeTo
	torus3dshapeFrom.GongCopyBasicFields(torus3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTorusEdge3DShape(mapOrigCopy map[any]any, torusedge3dshapeFrom *TorusEdge3DShape) (torusedge3dshapeTo *TorusEdge3DShape) {

	// torusedge3dshapeFrom has already been copied
	if _torusedge3dshapeTo, ok := mapOrigCopy[torusedge3dshapeFrom]; ok {
		torusedge3dshapeTo = _torusedge3dshapeTo.(*TorusEdge3DShape)
		return
	}

	torusedge3dshapeTo = new(TorusEdge3DShape)
	mapOrigCopy[torusedge3dshapeFrom] = torusedge3dshapeTo
	torusedge3dshapeFrom.GongCopyBasicFields(torusedge3dshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTorusStackShape(mapOrigCopy map[any]any, torusstackshapeFrom *TorusStackShape) (torusstackshapeTo *TorusStackShape) {

	// torusstackshapeFrom has already been copied
	if _torusstackshapeTo, ok := mapOrigCopy[torusstackshapeFrom]; ok {
		torusstackshapeTo = _torusstackshapeTo.(*TorusStackShape)
		return
	}

	torusstackshapeTo = new(TorusStackShape)
	mapOrigCopy[torusstackshapeFrom] = torusstackshapeTo
	torusstackshapeFrom.GongCopyBasicFields(torusstackshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTubeVase3DDiagram(mapOrigCopy map[any]any, tubevase3ddiagramFrom *TubeVase3DDiagram) (tubevase3ddiagramTo *TubeVase3DDiagram) {

	// tubevase3ddiagramFrom has already been copied
	if _tubevase3ddiagramTo, ok := mapOrigCopy[tubevase3ddiagramFrom]; ok {
		tubevase3ddiagramTo = _tubevase3ddiagramTo.(*TubeVase3DDiagram)
		return
	}

	tubevase3ddiagramTo = new(TubeVase3DDiagram)
	mapOrigCopy[tubevase3ddiagramFrom] = tubevase3ddiagramTo
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

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTubeVaseAbstract(mapOrigCopy map[any]any, tubevaseabstractFrom *TubeVaseAbstract) (tubevaseabstractTo *TubeVaseAbstract) {

	// tubevaseabstractFrom has already been copied
	if _tubevaseabstractTo, ok := mapOrigCopy[tubevaseabstractFrom]; ok {
		tubevaseabstractTo = _tubevaseabstractTo.(*TubeVaseAbstract)
		return
	}

	tubevaseabstractTo = new(TubeVaseAbstract)
	mapOrigCopy[tubevaseabstractFrom] = tubevaseabstractTo
	tubevaseabstractFrom.GongCopyBasicFields(tubevaseabstractTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVase2DDiagram(mapOrigCopy map[any]any, vase2ddiagramFrom *Vase2DDiagram) (vase2ddiagramTo *Vase2DDiagram) {

	// vase2ddiagramFrom has already been copied
	if _vase2ddiagramTo, ok := mapOrigCopy[vase2ddiagramFrom]; ok {
		vase2ddiagramTo = _vase2ddiagramTo.(*Vase2DDiagram)
		return
	}

	vase2ddiagramTo = new(Vase2DDiagram)
	mapOrigCopy[vase2ddiagramFrom] = vase2ddiagramTo
	vase2ddiagramFrom.GongCopyBasicFields(vase2ddiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVerticalTorusStackShape(mapOrigCopy map[any]any, verticaltorusstackshapeFrom *VerticalTorusStackShape) (verticaltorusstackshapeTo *VerticalTorusStackShape) {

	// verticaltorusstackshapeFrom has already been copied
	if _verticaltorusstackshapeTo, ok := mapOrigCopy[verticaltorusstackshapeFrom]; ok {
		verticaltorusstackshapeTo = _verticaltorusstackshapeTo.(*VerticalTorusStackShape)
		return
	}

	verticaltorusstackshapeTo = new(VerticalTorusStackShape)
	mapOrigCopy[verticaltorusstackshapeFrom] = verticaltorusstackshapeTo
	verticaltorusstackshapeFrom.GongCopyBasicFields(verticaltorusstackshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVolumeKey3DShape(mapOrigCopy map[any]any, volumekey3dshapeFrom *VolumeKey3DShape) (volumekey3dshapeTo *VolumeKey3DShape) {

	// volumekey3dshapeFrom has already been copied
	if _volumekey3dshapeTo, ok := mapOrigCopy[volumekey3dshapeFrom]; ok {
		volumekey3dshapeTo = _volumekey3dshapeTo.(*VolumeKey3DShape)
		return
	}

	volumekey3dshapeTo = new(VolumeKey3DShape)
	mapOrigCopy[volumekey3dshapeFrom] = volumekey3dshapeTo
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

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (angle0shape *Angle0Shape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchAngle0Shape(angle0shape)
}

func (stage *Stage) UnstageBranchAngle0Shape(angle0shape *Angle0Shape) {

	// check if instance is already staged
	if !stage.IsStaged(angle0shape) {
		return
	}

	angle0shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arcnormalvectorshape *ArcNormalVectorShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchArcNormalVectorShape(arcnormalvectorshape)
}

func (stage *Stage) UnstageBranchArcNormalVectorShape(arcnormalvectorshape *ArcNormalVectorShape) {

	// check if instance is already staged
	if !stage.IsStaged(arcnormalvectorshape) {
		return
	}

	arcnormalvectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchArcNormalVectorShapeGrid(arcnormalvectorshapegrid)
}

func (stage *Stage) UnstageBranchArcNormalVectorShapeGrid(arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) {

	// check if instance is already staged
	if !stage.IsStaged(arcnormalvectorshapegrid) {
		return
	}

	arcnormalvectorshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (axesshape *AxesShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchAxesShape(axesshape)
}

func (stage *Stage) UnstageBranchAxesShape(axesshape *AxesShape) {

	// check if instance is already staged
	if !stage.IsStaged(axesshape) {
		return
	}

	axesshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (basevectorshape *BaseVectorShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchBaseVectorShape(basevectorshape)
}

func (stage *Stage) UnstageBranchBaseVectorShape(basevectorshape *BaseVectorShape) {

	// check if instance is already staged
	if !stage.IsStaged(basevectorshape) {
		return
	}

	basevectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (basevectorshapegrid *BaseVectorShapeGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchBaseVectorShapeGrid(basevectorshapegrid)
}

func (stage *Stage) UnstageBranchBaseVectorShapeGrid(basevectorshapegrid *BaseVectorShapeGrid) {

	// check if instance is already staged
	if !stage.IsStaged(basevectorshapegrid) {
		return
	}

	basevectorshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchChosenP1P2PairShape(chosenp1p2pairshape)
}

func (stage *Stage) UnstageBranchChosenP1P2PairShape(chosenp1p2pairshape *ChosenP1P2PairShape) {

	// check if instance is already staged
	if !stage.IsStaged(chosenp1p2pairshape) {
		return
	}

	chosenp1p2pairshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circlegridshape *CircleGridShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCircleGridShape(circlegridshape)
}

func (stage *Stage) UnstageBranchCircleGridShape(circlegridshape *CircleGridShape) {

	// check if instance is already staged
	if !stage.IsStaged(circlegridshape) {
		return
	}

	circlegridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circumference3dshape *Circumference3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCircumference3DShape(circumference3dshape)
}

func (stage *Stage) UnstageBranchCircumference3DShape(circumference3dshape *Circumference3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(circumference3dshape) {
		return
	}

	circumference3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clock2ddiagram *Clock2DDiagram) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchClock2DDiagram(clock2ddiagram)
}

func (stage *Stage) UnstageBranchClock2DDiagram(clock2ddiagram *Clock2DDiagram) {

	// check if instance is already staged
	if !stage.IsStaged(clock2ddiagram) {
		return
	}

	clock2ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clock3ddiagram *Clock3DDiagram) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchClock3DDiagram(clock3ddiagram)
}

func (stage *Stage) UnstageBranchClock3DDiagram(clock3ddiagram *Clock3DDiagram) {

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

func (clockabstract *ClockAbstract) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchClockAbstract(clockabstract)
}

func (stage *Stage) UnstageBranchClockAbstract(clockabstract *ClockAbstract) {

	// check if instance is already staged
	if !stage.IsStaged(clockabstract) {
		return
	}

	clockabstract.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clocktopcurveshape *ClockTopCurveShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchClockTopCurveShape(clocktopcurveshape)
}

func (stage *Stage) UnstageBranchClockTopCurveShape(clocktopcurveshape *ClockTopCurveShape) {

	// check if instance is already staged
	if !stage.IsStaged(clocktopcurveshape) {
		return
	}

	clocktopcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cutline3dshape *CutLine3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCutLine3DShape(cutline3dshape)
}

func (stage *Stage) UnstageBranchCutLine3DShape(cutline3dshape *CutLine3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(cutline3dshape) {
		return
	}

	cutline3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endarcshape *EndArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchEndArcShape(endarcshape)
}

func (stage *Stage) UnstageBranchEndArcShape(endarcshape *EndArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(endarcshape) {
		return
	}

	endarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endarcshapegrid *EndArcShapeGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchEndArcShapeGrid(endarcshapegrid)
}

func (stage *Stage) UnstageBranchEndArcShapeGrid(endarcshapegrid *EndArcShapeGrid) {

	// check if instance is already staged
	if !stage.IsStaged(endarcshapegrid) {
		return
	}

	endarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endhalfwayarcshape *EndHalfwayArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchEndHalfwayArcShape(endhalfwayarcshape)
}

func (stage *Stage) UnstageBranchEndHalfwayArcShape(endhalfwayarcshape *EndHalfwayArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(endhalfwayarcshape) {
		return
	}

	endhalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchEndHalfwayArcShapeGrid(endhalfwayarcshapegrid)
}

func (stage *Stage) UnstageBranchEndHalfwayArcShapeGrid(endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) {

	// check if instance is already staged
	if !stage.IsStaged(endhalfwayarcshapegrid) {
		return
	}

	endhalfwayarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (explanationtextshape *ExplanationTextShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchExplanationTextShape(explanationtextshape)
}

func (stage *Stage) UnstageBranchExplanationTextShape(explanationtextshape *ExplanationTextShape) {

	// check if instance is already staged
	if !stage.IsStaged(explanationtextshape) {
		return
	}

	explanationtextshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eye3dshape *Eye3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchEye3DShape(eye3dshape)
}

func (stage *Stage) UnstageBranchEye3DShape(eye3dshape *Eye3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(eye3dshape) {
		return
	}

	eye3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchEyeCornersSampledPoints3DShape(eyecornerssampledpoints3dshape)
}

func (stage *Stage) UnstageBranchEyeCornersSampledPoints3DShape(eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(eyecornerssampledpoints3dshape) {
		return
	}

	eyecornerssampledpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchEyeSampledPoints3DShape(eyesampledpoints3dshape)
}

func (stage *Stage) UnstageBranchEyeSampledPoints3DShape(eyesampledpoints3dshape *EyeSampledPoints3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(eyesampledpoints3dshape) {
		return
	}

	eyesampledpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchEyeSeatBottomCurveShape(eyeseatbottomcurveshape)
}

func (stage *Stage) UnstageBranchEyeSeatBottomCurveShape(eyeseatbottomcurveshape *EyeSeatBottomCurveShape) {

	// check if instance is already staged
	if !stage.IsStaged(eyeseatbottomcurveshape) {
		return
	}

	eyeseatbottomcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchEyeStoolBottomCurveShape(eyestoolbottomcurveshape)
}

func (stage *Stage) UnstageBranchEyeStoolBottomCurveShape(eyestoolbottomcurveshape *EyeStoolBottomCurveShape) {

	// check if instance is already staged
	if !stage.IsStaged(eyestoolbottomcurveshape) {
		return
	}

	eyestoolbottomcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (eyevolume3dshape *EyeVolume3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchEyeVolume3DShape(eyevolume3dshape)
}

func (stage *Stage) UnstageBranchEyeVolume3DShape(eyevolume3dshape *EyeVolume3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(eyevolume3dshape) {
		return
	}

	eyevolume3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gridpathshape *GridPathShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGridPathShape(gridpathshape)
}

func (stage *Stage) UnstageBranchGridPathShape(gridpathshape *GridPathShape) {

	// check if instance is already staged
	if !stage.IsStaged(gridpathshape) {
		return
	}

	gridpathshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2d *GrowthCurve2D) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGrowthCurve2D(growthcurve2d)
}

func (stage *Stage) UnstageBranchGrowthCurve2D(growthcurve2d *GrowthCurve2D) {

	// check if instance is already staged
	if !stage.IsStaged(growthcurve2d) {
		return
	}

	growthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGrowthCurve2DRibbon(growthcurve2dribbon)
}

func (stage *Stage) UnstageBranchGrowthCurve2DRibbon(growthcurve2dribbon *GrowthCurve2DRibbon) {

	// check if instance is already staged
	if !stage.IsStaged(growthcurve2dribbon) {
		return
	}

	growthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGrowthCurve2DRibbonEndShape(growthcurve2dribbonendshape)
}

func (stage *Stage) UnstageBranchGrowthCurve2DRibbonEndShape(growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !stage.IsStaged(growthcurve2dribbonendshape) {
		return
	}

	growthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGrowthCurve2DRibbonStartShape(growthcurve2dribbonstartshape)
}

func (stage *Stage) UnstageBranchGrowthCurve2DRibbonStartShape(growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !stage.IsStaged(growthcurve2dribbonstartshape) {
		return
	}

	growthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGrowthCurveRhombusGridShape(growthcurverhombusgridshape)
}

func (stage *Stage) UnstageBranchGrowthCurveRhombusGridShape(growthcurverhombusgridshape *GrowthCurveRhombusGridShape) {

	// check if instance is already staged
	if !stage.IsStaged(growthcurverhombusgridshape) {
		return
	}

	growthcurverhombusgridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGrowthCurveRhombusShape(growthcurverhombusshape)
}

func (stage *Stage) UnstageBranchGrowthCurveRhombusShape(growthcurverhombusshape *GrowthCurveRhombusShape) {

	// check if instance is already staged
	if !stage.IsStaged(growthcurverhombusshape) {
		return
	}

	growthcurverhombusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (growthvectorshape *GrowthVectorShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGrowthVectorShape(growthvectorshape)
}

func (stage *Stage) UnstageBranchGrowthVectorShape(growthvectorshape *GrowthVectorShape) {

	// check if instance is already staged
	if !stage.IsStaged(growthvectorshape) {
		return
	}

	growthvectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (initialrhombusgridshape *InitialRhombusGridShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchInitialRhombusGridShape(initialrhombusgridshape)
}

func (stage *Stage) UnstageBranchInitialRhombusGridShape(initialrhombusgridshape *InitialRhombusGridShape) {

	// check if instance is already staged
	if !stage.IsStaged(initialrhombusgridshape) {
		return
	}

	initialrhombusgridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (initialrhombusshape *InitialRhombusShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchInitialRhombusShape(initialrhombusshape)
}

func (stage *Stage) UnstageBranchInitialRhombusShape(initialrhombusshape *InitialRhombusShape) {

	// check if instance is already staged
	if !stage.IsStaged(initialrhombusshape) {
		return
	}

	initialrhombusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (key3dshape *Key3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchKey3DShape(key3dshape)
}

func (stage *Stage) UnstageBranchKey3DShape(key3dshape *Key3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(key3dshape) {
		return
	}

	key3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (keyhole3dshape *KeyHole3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchKeyHole3DShape(keyhole3dshape)
}

func (stage *Stage) UnstageBranchKeyHole3DShape(keyhole3dshape *KeyHole3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(keyhole3dshape) {
		return
	}

	keyhole3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (keyholeshape *KeyHoleShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchKeyHoleShape(keyholeshape)
}

func (stage *Stage) UnstageBranchKeyHoleShape(keyholeshape *KeyHoleShape) {

	// check if instance is already staged
	if !stage.IsStaged(keyholeshape) {
		return
	}

	keyholeshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (leaves3dshape *Leaves3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLeaves3DShape(leaves3dshape)
}

func (stage *Stage) UnstageBranchLeaves3DShape(leaves3dshape *Leaves3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(leaves3dshape) {
		return
	}

	leaves3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLibrary(library)
}

func (stage *Stage) UnstageBranchLibrary(library *Library) {

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
	stage.UnstageBranchMidArcVectorShape(midarcvectorshape)
}

func (stage *Stage) UnstageBranchMidArcVectorShape(midarcvectorshape *MidArcVectorShape) {

	// check if instance is already staged
	if !stage.IsStaged(midarcvectorshape) {
		return
	}

	midarcvectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchMidArcVectorShapeGrid(midarcvectorshapegrid)
}

func (stage *Stage) UnstageBranchMidArcVectorShapeGrid(midarcvectorshapegrid *MidArcVectorShapeGrid) {

	// check if instance is already staged
	if !stage.IsStaged(midarcvectorshapegrid) {
		return
	}

	midarcvectorshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (musicabstract *MusicAbstract) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchMusicAbstract(musicabstract)
}

func (stage *Stage) UnstageBranchMusicAbstract(musicabstract *MusicAbstract) {

	// check if instance is already staged
	if !stage.IsStaged(musicabstract) {
		return
	}

	musicabstract.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (originalpoints3dshape *OriginalPoints3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchOriginalPoints3DShape(originalpoints3dshape)
}

func (stage *Stage) UnstageBranchOriginalPoints3DShape(originalpoints3dshape *OriginalPoints3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(originalpoints3dshape) {
		return
	}

	originalpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchParastichyMCurves3DShape(parastichymcurves3dshape)
}

func (stage *Stage) UnstageBranchParastichyMCurves3DShape(parastichymcurves3dshape *ParastichyMCurves3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(parastichymcurves3dshape) {
		return
	}

	parastichymcurves3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchParastichyNCurves3DShape(parastichyncurves3dshape)
}

func (stage *Stage) UnstageBranchParastichyNCurves3DShape(parastichyncurves3dshape *ParastichyNCurves3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(parastichyncurves3dshape) {
		return
	}

	parastichyncurves3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyGrowthCurve2DRibbon(partiallygrowthcurve2dribbon)
}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DRibbon(partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dribbon) {
		return
	}

	partiallygrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyGrowthCurve2DRibbonEndShape(partiallygrowthcurve2dribbonendshape)
}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DRibbonEndShape(partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dribbonendshape) {
		return
	}

	partiallygrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyGrowthCurve2DRibbonStartShape(partiallygrowthcurve2dribbonstartshape)
}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DRibbonStartShape(partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dribbonstartshape) {
		return
	}

	partiallygrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyGrowthCurve2DTrajectory(partiallygrowthcurve2dtrajectory)
}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectory(partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectory) {
		return
	}

	partiallygrowthcurve2dtrajectory.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryP1CurveShape(partiallygrowthcurve2dtrajectoryp1curveshape)
}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryP1CurveShape(partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryp1curveshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1curveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryP1P2(partiallygrowthcurve2dtrajectoryp1p2)
}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryP1P2(partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryp1p2) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(partiallygrowthcurve2dtrajectoryp1p2pairlineshape)
}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape(partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryp1p2pairlineshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1p2pairlineshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryP1PointShape(partiallygrowthcurve2dtrajectoryp1pointshape)
}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryP1PointShape(partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryp1pointshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp1pointshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryP2CurveShape(partiallygrowthcurve2dtrajectoryp2curveshape)
}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryP2CurveShape(partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryp2curveshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp2curveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryP2PointShape(partiallygrowthcurve2dtrajectoryp2pointshape)
}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryP2PointShape(partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryp2pointshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryp2pointshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyGrowthCurve2DTrajectoryShape(partiallygrowthcurve2dtrajectoryshape)
}

func (stage *Stage) UnstageBranchPartiallyGrowthCurve2DTrajectoryShape(partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) {

	// check if instance is already staged
	if !stage.IsStaged(partiallygrowthcurve2dtrajectoryshape) {
		return
	}

	partiallygrowthcurve2dtrajectoryshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyRotatedSeatBottomCurveShape(partiallyrotatedseatbottomcurveshape)
}

func (stage *Stage) UnstageBranchPartiallyRotatedSeatBottomCurveShape(partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) {

	// check if instance is already staged
	if !stage.IsStaged(partiallyrotatedseatbottomcurveshape) {
		return
	}

	partiallyrotatedseatbottomcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyRotatedSeatTopCurveShape(partiallyrotatedseattopcurveshape)
}

func (stage *Stage) UnstageBranchPartiallyRotatedSeatTopCurveShape(partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) {

	// check if instance is already staged
	if !stage.IsStaged(partiallyrotatedseattopcurveshape) {
		return
	}

	partiallyrotatedseattopcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartiallyRotatedTorusShape(partiallyrotatedtorusshape)
}

func (stage *Stage) UnstageBranchPartiallyRotatedTorusShape(partiallyrotatedtorusshape *PartiallyRotatedTorusShape) {

	// check if instance is already staged
	if !stage.IsStaged(partiallyrotatedtorusshape) {
		return
	}

	partiallyrotatedtorusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvector *PerpendicularVector) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPerpendicularVector(perpendicularvector)
}

func (stage *Stage) UnstageBranchPerpendicularVector(perpendicularvector *PerpendicularVector) {

	// check if instance is already staged
	if !stage.IsStaged(perpendicularvector) {
		return
	}

	perpendicularvector.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPerpendicularVectorGrid(perpendicularvectorgrid)
}

func (stage *Stage) UnstageBranchPerpendicularVectorGrid(perpendicularvectorgrid *PerpendicularVectorGrid) {

	// check if instance is already staged
	if !stage.IsStaged(perpendicularvectorgrid) {
		return
	}

	perpendicularvectorgrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway)
}

func (stage *Stage) UnstageBranchPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) {

	// check if instance is already staged
	if !stage.IsStaged(perpendicularvectorgridhalfway) {
		return
	}

	perpendicularvectorgridhalfway.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPerpendicularVectorHalfway(perpendicularvectorhalfway)
}

func (stage *Stage) UnstageBranchPerpendicularVectorHalfway(perpendicularvectorhalfway *PerpendicularVectorHalfway) {

	// check if instance is already staged
	if !stage.IsStaged(perpendicularvectorhalfway) {
		return
	}

	perpendicularvectorhalfway.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (plant2ddiagram *Plant2DDiagram) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPlant2DDiagram(plant2ddiagram)
}

func (stage *Stage) UnstageBranchPlant2DDiagram(plant2ddiagram *Plant2DDiagram) {

	// check if instance is already staged
	if !stage.IsStaged(plant2ddiagram) {
		return
	}

	plant2ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (plant3ddiagram *Plant3DDiagram) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPlant3DDiagram(plant3ddiagram)
}

func (stage *Stage) UnstageBranchPlant3DDiagram(plant3ddiagram *Plant3DDiagram) {

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
	stage.UnstageBranchPlantAbstract(plantabstract)
}

func (stage *Stage) UnstageBranchPlantAbstract(plantabstract *PlantAbstract) {

	// check if instance is already staged
	if !stage.IsStaged(plantabstract) {
		return
	}

	plantabstract.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plantabstract.TubeVaseAbstract != nil {
		stage.UnstageBranch(plantabstract.TubeVaseAbstract)
	}
	if plantabstract.StoolAbstract != nil {
		stage.UnstageBranch(plantabstract.StoolAbstract)
	}
	if plantabstract.ClockAbstract != nil {
		stage.UnstageBranch(plantabstract.ClockAbstract)
	}
	if plantabstract.MusicAbstract != nil {
		stage.UnstageBranch(plantabstract.MusicAbstract)
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
	stage.UnstageBranchPlantCircumferenceShape(plantcircumferenceshape)
}

func (stage *Stage) UnstageBranchPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape) {

	// check if instance is already staged
	if !stage.IsStaged(plantcircumferenceshape) {
		return
	}

	plantcircumferenceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointsandlines3dshape *PointsAndLines3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPointsAndLines3DShape(pointsandlines3dshape)
}

func (stage *Stage) UnstageBranchPointsAndLines3DShape(pointsandlines3dshape *PointsAndLines3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(pointsandlines3dshape) {
		return
	}

	pointsandlines3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pxshape *PxShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPxShape(pxshape)
}

func (stage *Stage) UnstageBranchPxShape(pxshape *PxShape) {

	// check if instance is already staged
	if !stage.IsStaged(pxshape) {
		return
	}

	pxshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rendered3dshape *Rendered3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRendered3DShape(rendered3dshape)
}

func (stage *Stage) UnstageBranchRendered3DShape(rendered3dshape *Rendered3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(rendered3dshape) {
		return
	}

	rendered3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rhombusshape *RhombusShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRhombusShape(rhombusshape)
}

func (stage *Stage) UnstageBranchRhombusShape(rhombusshape *RhombusShape) {

	// check if instance is already staged
	if !stage.IsStaged(rhombusshape) {
		return
	}

	rhombusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rhombusstuff *RhombusStuff) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRhombusStuff(rhombusstuff)
}

func (stage *Stage) UnstageBranchRhombusStuff(rhombusstuff *RhombusStuff) {

	// check if instance is already staged
	if !stage.IsStaged(rhombusstuff) {
		return
	}

	rhombusstuff.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRotatedRhombusGridShape(rotatedrhombusgridshape)
}

func (stage *Stage) UnstageBranchRotatedRhombusGridShape(rotatedrhombusgridshape *RotatedRhombusGridShape) {

	// check if instance is already staged
	if !stage.IsStaged(rotatedrhombusgridshape) {
		return
	}

	rotatedrhombusgridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedrhombusshape *RotatedRhombusShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRotatedRhombusShape(rotatedrhombusshape)
}

func (stage *Stage) UnstageBranchRotatedRhombusShape(rotatedrhombusshape *RotatedRhombusShape) {

	// check if instance is already staged
	if !stage.IsStaged(rotatedrhombusshape) {
		return
	}

	rotatedrhombusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRotatedSampledPoints3DShape(rotatedsampledpoints3dshape)
}

func (stage *Stage) UnstageBranchRotatedSampledPoints3DShape(rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(rotatedsampledpoints3dshape) {
		return
	}

	rotatedsampledpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRotatedSeatAndLegs3DShape(rotatedseatandlegs3dshape)
}

func (stage *Stage) UnstageBranchRotatedSeatAndLegs3DShape(rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(rotatedseatandlegs3dshape) {
		return
	}

	rotatedseatandlegs3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (sampledpoints3dshape *SampledPoints3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSampledPoints3DShape(sampledpoints3dshape)
}

func (stage *Stage) UnstageBranchSampledPoints3DShape(sampledpoints3dshape *SampledPoints3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(sampledpoints3dshape) {
		return
	}

	sampledpoints3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seat3dshape *Seat3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSeat3DShape(seat3dshape)
}

func (stage *Stage) UnstageBranchSeat3DShape(seat3dshape *Seat3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(seat3dshape) {
		return
	}

	seat3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSeatAndLegs3DShape(seatandlegs3dshape)
}

func (stage *Stage) UnstageBranchSeatAndLegs3DShape(seatandlegs3dshape *SeatAndLegs3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(seatandlegs3dshape) {
		return
	}

	seatandlegs3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seatbottomcurveshape *SeatBottomCurveShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSeatBottomCurveShape(seatbottomcurveshape)
}

func (stage *Stage) UnstageBranchSeatBottomCurveShape(seatbottomcurveshape *SeatBottomCurveShape) {

	// check if instance is already staged
	if !stage.IsStaged(seatbottomcurveshape) {
		return
	}

	seatbottomcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (seattopcurveshape *SeatTopCurveShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSeatTopCurveShape(seattopcurveshape)
}

func (stage *Stage) UnstageBranchSeatTopCurveShape(seattopcurveshape *SeatTopCurveShape) {

	// check if instance is already staged
	if !stage.IsStaged(seattopcurveshape) {
		return
	}

	seattopcurveshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedBottomTopStartArcShape(shiftedbottomtopstartarcshape)
}

func (stage *Stage) UnstageBranchShiftedBottomTopStartArcShape(shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedbottomtopstartarcshape) {
		return
	}

	shiftedbottomtopstartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedBottomTopStartArcShapeGrid(shiftedbottomtopstartarcshapegrid)
}

func (stage *Stage) UnstageBranchShiftedBottomTopStartArcShapeGrid(shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedbottomtopstartarcshapegrid) {
		return
	}

	shiftedbottomtopstartarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedLeftGrowthCurve2DRibbon(shiftedleftgrowthcurve2dribbon)
}

func (stage *Stage) UnstageBranchShiftedLeftGrowthCurve2DRibbon(shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftgrowthcurve2dribbon) {
		return
	}

	shiftedleftgrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedLeftGrowthCurve2DRibbonEndShape(shiftedleftgrowthcurve2dribbonendshape)
}

func (stage *Stage) UnstageBranchShiftedLeftGrowthCurve2DRibbonEndShape(shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftgrowthcurve2dribbonendshape) {
		return
	}

	shiftedleftgrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedLeftGrowthCurve2DRibbonStartShape(shiftedleftgrowthcurve2dribbonstartshape)
}

func (stage *Stage) UnstageBranchShiftedLeftGrowthCurve2DRibbonStartShape(shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftgrowthcurve2dribbonstartshape) {
		return
	}

	shiftedleftgrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedLeftPartiallyGrowthCurve2DRibbon(shiftedleftpartiallygrowthcurve2dribbon)
}

func (stage *Stage) UnstageBranchShiftedLeftPartiallyGrowthCurve2DRibbon(shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftpartiallygrowthcurve2dribbon) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(shiftedleftpartiallygrowthcurve2dribbonendshape)
}

func (stage *Stage) UnstageBranchShiftedLeftPartiallyGrowthCurve2DRibbonEndShape(shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftpartiallygrowthcurve2dribbonendshape) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(shiftedleftpartiallygrowthcurve2dribbonstartshape)
}

func (stage *Stage) UnstageBranchShiftedLeftPartiallyGrowthCurve2DRibbonStartShape(shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftpartiallygrowthcurve2dribbonstartshape) {
		return
	}

	shiftedleftpartiallygrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedLeftStackGrowthCurveEndArcShape(shiftedleftstackgrowthcurveendarcshape)
}

func (stage *Stage) UnstageBranchShiftedLeftStackGrowthCurveEndArcShape(shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftstackgrowthcurveendarcshape) {
		return
	}

	shiftedleftstackgrowthcurveendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedLeftStackGrowthCurveStartArcShape(shiftedleftstackgrowthcurvestartarcshape)
}

func (stage *Stage) UnstageBranchShiftedLeftStackGrowthCurveStartArcShape(shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftstackgrowthcurvestartarcshape) {
		return
	}

	shiftedleftstackgrowthcurvestartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedLeftStackNormalVector(shiftedleftstacknormalvector)
}

func (stage *Stage) UnstageBranchShiftedLeftStackNormalVector(shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftstacknormalvector) {
		return
	}

	shiftedleftstacknormalvector.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedLeftStackOfGrowthCurve(shiftedleftstackofgrowthcurve)
}

func (stage *Stage) UnstageBranchShiftedLeftStackOfGrowthCurve(shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftstackofgrowthcurve) {
		return
	}

	shiftedleftstackofgrowthcurve.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector)
}

func (stage *Stage) UnstageBranchShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedleftstackofnormalvector) {
		return
	}

	shiftedleftstackofnormalvector.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedRightGrowthCurve2DRibbon(shiftedrightgrowthcurve2dribbon)
}

func (stage *Stage) UnstageBranchShiftedRightGrowthCurve2DRibbon(shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedrightgrowthcurve2dribbon) {
		return
	}

	shiftedrightgrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedRightGrowthCurve2DRibbonEndShape(shiftedrightgrowthcurve2dribbonendshape)
}

func (stage *Stage) UnstageBranchShiftedRightGrowthCurve2DRibbonEndShape(shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedrightgrowthcurve2dribbonendshape) {
		return
	}

	shiftedrightgrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchShiftedRightGrowthCurve2DRibbonStartShape(shiftedrightgrowthcurve2dribbonstartshape)
}

func (stage *Stage) UnstageBranchShiftedRightGrowthCurve2DRibbonStartShape(shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !stage.IsStaged(shiftedrightgrowthcurve2dribbonstartshape) {
		return
	}

	shiftedrightgrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackGrowthCurve2DEndHalfwayArcShape(stackgrowthcurve2dendhalfwayarcshape)
}

func (stage *Stage) UnstageBranchStackGrowthCurve2DEndHalfwayArcShape(stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(stackgrowthcurve2dendhalfwayarcshape) {
		return
	}

	stackgrowthcurve2dendhalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackGrowthCurve2DRibbonEndShape(stackgrowthcurve2dribbonendshape)
}

func (stage *Stage) UnstageBranchStackGrowthCurve2DRibbonEndShape(stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !stage.IsStaged(stackgrowthcurve2dribbonendshape) {
		return
	}

	stackgrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackGrowthCurve2DRibbonStartShape(stackgrowthcurve2dribbonstartshape)
}

func (stage *Stage) UnstageBranchStackGrowthCurve2DRibbonStartShape(stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !stage.IsStaged(stackgrowthcurve2dribbonstartshape) {
		return
	}

	stackgrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackGrowthCurve2DStartHalfwayArcShape(stackgrowthcurve2dstarthalfwayarcshape)
}

func (stage *Stage) UnstageBranchStackGrowthCurve2DStartHalfwayArcShape(stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(stackgrowthcurve2dstarthalfwayarcshape) {
		return
	}

	stackgrowthcurve2dstarthalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackOfGrowthCurve2D(stackofgrowthcurve2d)
}

func (stage *Stage) UnstageBranchStackOfGrowthCurve2D(stackofgrowthcurve2d *StackOfGrowthCurve2D) {

	// check if instance is already staged
	if !stage.IsStaged(stackofgrowthcurve2d) {
		return
	}

	stackofgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackOfGrowthCurve2DByGrowthVector(stackofgrowthcurve2dbygrowthvector)
}

func (stage *Stage) UnstageBranchStackOfGrowthCurve2DByGrowthVector(stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) {

	// check if instance is already staged
	if !stage.IsStaged(stackofgrowthcurve2dbygrowthvector) {
		return
	}

	stackofgrowthcurve2dbygrowthvector.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackOfGrowthCurve2DRibbon(stackofgrowthcurve2dribbon)
}

func (stage *Stage) UnstageBranchStackOfGrowthCurve2DRibbon(stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) {

	// check if instance is already staged
	if !stage.IsStaged(stackofgrowthcurve2dribbon) {
		return
	}

	stackofgrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackOfPartiallyRotatedTorusShape(stackofpartiallyrotatedtorusshape)
}

func (stage *Stage) UnstageBranchStackOfPartiallyRotatedTorusShape(stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) {

	// check if instance is already staged
	if !stage.IsStaged(stackofpartiallyrotatedtorusshape) {
		return
	}

	stackofpartiallyrotatedtorusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackOfRotatedGrowthCurve2D(stackofrotatedgrowthcurve2d)
}

func (stage *Stage) UnstageBranchStackOfRotatedGrowthCurve2D(stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) {

	// check if instance is already staged
	if !stage.IsStaged(stackofrotatedgrowthcurve2d) {
		return
	}

	stackofrotatedgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackOfRotatedGrowthCurve2DRibbon(stackofrotatedgrowthcurve2dribbon)
}

func (stage *Stage) UnstageBranchStackOfRotatedGrowthCurve2DRibbon(stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) {

	// check if instance is already staged
	if !stage.IsStaged(stackofrotatedgrowthcurve2dribbon) {
		return
	}

	stackofrotatedgrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackRotatedGrowthCurve2DEndArcShape(stackrotatedgrowthcurve2dendarcshape)
}

func (stage *Stage) UnstageBranchStackRotatedGrowthCurve2DEndArcShape(stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(stackrotatedgrowthcurve2dendarcshape) {
		return
	}

	stackrotatedgrowthcurve2dendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackRotatedGrowthCurve2DRibbonEndShape(stackrotatedgrowthcurve2dribbonendshape)
}

func (stage *Stage) UnstageBranchStackRotatedGrowthCurve2DRibbonEndShape(stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) {

	// check if instance is already staged
	if !stage.IsStaged(stackrotatedgrowthcurve2dribbonendshape) {
		return
	}

	stackrotatedgrowthcurve2dribbonendshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackRotatedGrowthCurve2DRibbonStartShape(stackrotatedgrowthcurve2dribbonstartshape)
}

func (stage *Stage) UnstageBranchStackRotatedGrowthCurve2DRibbonStartShape(stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) {

	// check if instance is already staged
	if !stage.IsStaged(stackrotatedgrowthcurve2dribbonstartshape) {
		return
	}

	stackrotatedgrowthcurve2dribbonstartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStackRotatedGrowthCurve2DStartArcShape(stackrotatedgrowthcurve2dstartarcshape)
}

func (stage *Stage) UnstageBranchStackRotatedGrowthCurve2DStartArcShape(stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(stackrotatedgrowthcurve2dstartarcshape) {
		return
	}

	stackrotatedgrowthcurve2dstartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (startarcshape *StartArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStartArcShape(startarcshape)
}

func (stage *Stage) UnstageBranchStartArcShape(startarcshape *StartArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(startarcshape) {
		return
	}

	startarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (startarcshapegrid *StartArcShapeGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStartArcShapeGrid(startarcshapegrid)
}

func (stage *Stage) UnstageBranchStartArcShapeGrid(startarcshapegrid *StartArcShapeGrid) {

	// check if instance is already staged
	if !stage.IsStaged(startarcshapegrid) {
		return
	}

	startarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (starthalfwayarcshape *StartHalfwayArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStartHalfwayArcShape(starthalfwayarcshape)
}

func (stage *Stage) UnstageBranchStartHalfwayArcShape(starthalfwayarcshape *StartHalfwayArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(starthalfwayarcshape) {
		return
	}

	starthalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStartHalfwayArcShapeGrid(starthalfwayarcshapegrid)
}

func (stage *Stage) UnstageBranchStartHalfwayArcShapeGrid(starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) {

	// check if instance is already staged
	if !stage.IsStaged(starthalfwayarcshapegrid) {
		return
	}

	starthalfwayarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stemcylinder3dshape *StemCylinder3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStemCylinder3DShape(stemcylinder3dshape)
}

func (stage *Stage) UnstageBranchStemCylinder3DShape(stemcylinder3dshape *StemCylinder3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(stemcylinder3dshape) {
		return
	}

	stemcylinder3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stool2ddiagram *Stool2DDiagram) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStool2DDiagram(stool2ddiagram)
}

func (stage *Stage) UnstageBranchStool2DDiagram(stool2ddiagram *Stool2DDiagram) {

	// check if instance is already staged
	if !stage.IsStaged(stool2ddiagram) {
		return
	}

	stool2ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stool3ddiagram *Stool3DDiagram) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStool3DDiagram(stool3ddiagram)
}

func (stage *Stage) UnstageBranchStool3DDiagram(stool3ddiagram *Stool3DDiagram) {

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

func (stoolabstract *StoolAbstract) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStoolAbstract(stoolabstract)
}

func (stage *Stage) UnstageBranchStoolAbstract(stoolabstract *StoolAbstract) {

	// check if instance is already staged
	if !stage.IsStaged(stoolabstract) {
		return
	}

	stoolabstract.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tiledfloor3dshape *TiledFloor3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTiledFloor3DShape(tiledfloor3dshape)
}

func (stage *Stage) UnstageBranchTiledFloor3DShape(tiledfloor3dshape *TiledFloor3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(tiledfloor3dshape) {
		return
	}

	tiledfloor3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendarcshape *TopEndArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopEndArcShape(topendarcshape)
}

func (stage *Stage) UnstageBranchTopEndArcShape(topendarcshape *TopEndArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(topendarcshape) {
		return
	}

	topendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendarcshapegrid *TopEndArcShapeGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopEndArcShapeGrid(topendarcshapegrid)
}

func (stage *Stage) UnstageBranchTopEndArcShapeGrid(topendarcshapegrid *TopEndArcShapeGrid) {

	// check if instance is already staged
	if !stage.IsStaged(topendarcshapegrid) {
		return
	}

	topendarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopEndHalfwayArcShape(topendhalfwayarcshape)
}

func (stage *Stage) UnstageBranchTopEndHalfwayArcShape(topendhalfwayarcshape *TopEndHalfwayArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(topendhalfwayarcshape) {
		return
	}

	topendhalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopEndHalfwayArcShapeGrid(topendhalfwayarcshapegrid)
}

func (stage *Stage) UnstageBranchTopEndHalfwayArcShapeGrid(topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) {

	// check if instance is already staged
	if !stage.IsStaged(topendhalfwayarcshapegrid) {
		return
	}

	topendhalfwayarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topgrowthcurve2d *TopGrowthCurve2D) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopGrowthCurve2D(topgrowthcurve2d)
}

func (stage *Stage) UnstageBranchTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D) {

	// check if instance is already staged
	if !stage.IsStaged(topgrowthcurve2d) {
		return
	}

	topgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topmidarcvectorshape *TopMidArcVectorShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopMidArcVectorShape(topmidarcvectorshape)
}

func (stage *Stage) UnstageBranchTopMidArcVectorShape(topmidarcvectorshape *TopMidArcVectorShape) {

	// check if instance is already staged
	if !stage.IsStaged(topmidarcvectorshape) {
		return
	}

	topmidarcvectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopMidArcVectorShapeGrid(topmidarcvectorshapegrid)
}

func (stage *Stage) UnstageBranchTopMidArcVectorShapeGrid(topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) {

	// check if instance is already staged
	if !stage.IsStaged(topmidarcvectorshapegrid) {
		return
	}

	topmidarcvectorshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopStackGrowthCurve2DEndHalfwayArcShape(topstackgrowthcurve2dendhalfwayarcshape)
}

func (stage *Stage) UnstageBranchTopStackGrowthCurve2DEndHalfwayArcShape(topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(topstackgrowthcurve2dendhalfwayarcshape) {
		return
	}

	topstackgrowthcurve2dendhalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopStackGrowthCurve2DStartHalfwayArcShape(topstackgrowthcurve2dstarthalfwayarcshape)
}

func (stage *Stage) UnstageBranchTopStackGrowthCurve2DStartHalfwayArcShape(topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(topstackgrowthcurve2dstarthalfwayarcshape) {
		return
	}

	topstackgrowthcurve2dstarthalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopStackOfGrowthCurve2D(topstackofgrowthcurve2d)
}

func (stage *Stage) UnstageBranchTopStackOfGrowthCurve2D(topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) {

	// check if instance is already staged
	if !stage.IsStaged(topstackofgrowthcurve2d) {
		return
	}

	topstackofgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopStackOfRotatedGrowthCurve2D(topstackofrotatedgrowthcurve2d)
}

func (stage *Stage) UnstageBranchTopStackOfRotatedGrowthCurve2D(topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) {

	// check if instance is already staged
	if !stage.IsStaged(topstackofrotatedgrowthcurve2d) {
		return
	}

	topstackofrotatedgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopStackOfRotatedGrowthCurve2DEndArcShape(topstackofrotatedgrowthcurve2dendarcshape)
}

func (stage *Stage) UnstageBranchTopStackOfRotatedGrowthCurve2DEndArcShape(topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(topstackofrotatedgrowthcurve2dendarcshape) {
		return
	}

	topstackofrotatedgrowthcurve2dendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopStackOfRotatedGrowthCurve2DStartArcShape(topstackofrotatedgrowthcurve2dstartarcshape)
}

func (stage *Stage) UnstageBranchTopStackOfRotatedGrowthCurve2DStartArcShape(topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(topstackofrotatedgrowthcurve2dstartarcshape) {
		return
	}

	topstackofrotatedgrowthcurve2dstartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstartarcshape *TopStartArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopStartArcShape(topstartarcshape)
}

func (stage *Stage) UnstageBranchTopStartArcShape(topstartarcshape *TopStartArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(topstartarcshape) {
		return
	}

	topstartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopStartArcShapeGrid(topstartarcshapegrid)
}

func (stage *Stage) UnstageBranchTopStartArcShapeGrid(topstartarcshapegrid *TopStartArcShapeGrid) {

	// check if instance is already staged
	if !stage.IsStaged(topstartarcshapegrid) {
		return
	}

	topstartarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopStartHalfwayArcShape(topstarthalfwayarcshape)
}

func (stage *Stage) UnstageBranchTopStartHalfwayArcShape(topstarthalfwayarcshape *TopStartHalfwayArcShape) {

	// check if instance is already staged
	if !stage.IsStaged(topstarthalfwayarcshape) {
		return
	}

	topstarthalfwayarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTopStartHalfwayArcShapeGrid(topstarthalfwayarcshapegrid)
}

func (stage *Stage) UnstageBranchTopStartHalfwayArcShapeGrid(topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) {

	// check if instance is already staged
	if !stage.IsStaged(topstarthalfwayarcshapegrid) {
		return
	}

	topstarthalfwayarcshapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torus3dshape *Torus3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTorus3DShape(torus3dshape)
}

func (stage *Stage) UnstageBranchTorus3DShape(torus3dshape *Torus3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(torus3dshape) {
		return
	}

	torus3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torusedge3dshape *TorusEdge3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTorusEdge3DShape(torusedge3dshape)
}

func (stage *Stage) UnstageBranchTorusEdge3DShape(torusedge3dshape *TorusEdge3DShape) {

	// check if instance is already staged
	if !stage.IsStaged(torusedge3dshape) {
		return
	}

	torusedge3dshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torusstackshape *TorusStackShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTorusStackShape(torusstackshape)
}

func (stage *Stage) UnstageBranchTorusStackShape(torusstackshape *TorusStackShape) {

	// check if instance is already staged
	if !stage.IsStaged(torusstackshape) {
		return
	}

	torusstackshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tubevase3ddiagram *TubeVase3DDiagram) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTubeVase3DDiagram(tubevase3ddiagram)
}

func (stage *Stage) UnstageBranchTubeVase3DDiagram(tubevase3ddiagram *TubeVase3DDiagram) {

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

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tubevaseabstract *TubeVaseAbstract) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTubeVaseAbstract(tubevaseabstract)
}

func (stage *Stage) UnstageBranchTubeVaseAbstract(tubevaseabstract *TubeVaseAbstract) {

	// check if instance is already staged
	if !stage.IsStaged(tubevaseabstract) {
		return
	}

	tubevaseabstract.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vase2ddiagram *Vase2DDiagram) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchVase2DDiagram(vase2ddiagram)
}

func (stage *Stage) UnstageBranchVase2DDiagram(vase2ddiagram *Vase2DDiagram) {

	// check if instance is already staged
	if !stage.IsStaged(vase2ddiagram) {
		return
	}

	vase2ddiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (verticaltorusstackshape *VerticalTorusStackShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchVerticalTorusStackShape(verticaltorusstackshape)
}

func (stage *Stage) UnstageBranchVerticalTorusStackShape(verticaltorusstackshape *VerticalTorusStackShape) {

	// check if instance is already staged
	if !stage.IsStaged(verticaltorusstackshape) {
		return
	}

	verticaltorusstackshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (volumekey3dshape *VolumeKey3DShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchVolumeKey3DShape(volumekey3dshape)
}

func (stage *Stage) UnstageBranchVolumeKey3DShape(volumekey3dshape *VolumeKey3DShape) {

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

func (reference *Leaves3DShape) GongReconstructPointersFromReferences(stage *Stage, instance *Leaves3DShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Plants = reference.Plants[:0]
	for _, _b := range instance.Plants {
		reference.Plants = append(reference.Plants, stage.PlantAbstracts_reference[_b])
	}
	reference.SubLibraries = reference.SubLibraries[:0]
	for _, _b := range instance.SubLibraries {
		reference.SubLibraries = append(reference.SubLibraries, stage.Librarys_reference[_b])
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

func (reference *MusicAbstract) GongReconstructPointersFromReferences(stage *Stage, instance *MusicAbstract) {
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
	if instance.Leaves3DShape != nil {
		reference.Leaves3DShape = stage.Leaves3DShapes_reference[instance.Leaves3DShape]
	}
	if instance.Rendered3DShape != nil {
		reference.Rendered3DShape = stage.Rendered3DShapes_reference[instance.Rendered3DShape]
	}
	// insertion point for slice of pointers field
}

func (reference *PlantAbstract) GongReconstructPointersFromReferences(stage *Stage, instance *PlantAbstract) {
	// insertion point for pointers field
	if instance.TubeVaseAbstract != nil {
		reference.TubeVaseAbstract = stage.TubeVaseAbstracts_reference[instance.TubeVaseAbstract]
	}
	if instance.StoolAbstract != nil {
		reference.StoolAbstract = stage.StoolAbstracts_reference[instance.StoolAbstract]
	}
	if instance.ClockAbstract != nil {
		reference.ClockAbstract = stage.ClockAbstracts_reference[instance.ClockAbstract]
	}
	if instance.MusicAbstract != nil {
		reference.MusicAbstract = stage.MusicAbstracts_reference[instance.MusicAbstract]
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
	reference.TubeVase3DDiagrams = reference.TubeVase3DDiagrams[:0]
	for _, _b := range instance.TubeVase3DDiagrams {
		reference.TubeVase3DDiagrams = append(reference.TubeVase3DDiagrams, stage.TubeVase3DDiagrams_reference[_b])
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

func (reference *TubeVase3DDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *TubeVase3DDiagram) {
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

func (reference *TubeVaseAbstract) GongReconstructPointersFromReferences(stage *Stage, instance *TubeVaseAbstract) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Vase2DDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *Vase2DDiagram) {
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

func (reference *Leaves3DShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Plants []*PlantAbstract
	for _, _reference := range reference.Plants {
		if _instance, ok := stage.PlantAbstracts_instance[_reference]; ok {
			_Plants = append(_Plants, _instance)
		}
	}
	reference.Plants = _Plants
	var _SubLibraries []*Library
	for _, _reference := range reference.SubLibraries {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibraries = append(_SubLibraries, _instance)
		}
	}
	reference.SubLibraries = _SubLibraries
}

func (reference *MidArcVectorShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MidArcVectorShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MusicAbstract) GongReconstructPointersFromInstances(stage *Stage) {
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
	if _reference := reference.Leaves3DShape; _reference != nil {
		reference.Leaves3DShape = nil
		if _instance, ok := stage.Leaves3DShapes_instance[_reference]; ok {
			reference.Leaves3DShape = _instance
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
	if _reference := reference.TubeVaseAbstract; _reference != nil {
		reference.TubeVaseAbstract = nil
		if _instance, ok := stage.TubeVaseAbstracts_instance[_reference]; ok {
			reference.TubeVaseAbstract = _instance
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
	if _reference := reference.MusicAbstract; _reference != nil {
		reference.MusicAbstract = nil
		if _instance, ok := stage.MusicAbstracts_instance[_reference]; ok {
			reference.MusicAbstract = _instance
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
	var _TubeVase3DDiagrams []*TubeVase3DDiagram
	for _, _reference := range reference.TubeVase3DDiagrams {
		if _instance, ok := stage.TubeVase3DDiagrams_instance[_reference]; ok {
			_TubeVase3DDiagrams = append(_TubeVase3DDiagrams, _instance)
		}
	}
	reference.TubeVase3DDiagrams = _TubeVase3DDiagrams
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

func (reference *TubeVase3DDiagram) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *TubeVaseAbstract) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Vase2DDiagram) GongReconstructPointersFromInstances(stage *Stage) {
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
		ops := stage.Diff(
			library,
			"Plants",
			len(libraryOther.Plants),
			len(library.Plants),
			func(i, j int) bool {
				return libraryOther.Plants[i] == library.Plants[j]
			},
			func(j int) string {
				return library.Plants[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
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
		ops := stage.Diff(
			library,
			"SubLibraries",
			len(libraryOther.SubLibraries),
			len(library.SubLibraries),
			func(i, j int) bool {
				return libraryOther.SubLibraries[i] == library.SubLibraries[j]
			},
			func(j int) string {
				return library.SubLibraries[j].GongGetIdentifier(stage)
			},
		)
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
func (musicabstract *MusicAbstract) GongDiff(stage *Stage, musicabstractOther *MusicAbstract) (diffs []string) {
	// insertion point for field diffs
	if musicabstract.Name != musicabstractOther.Name {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "Name"))
	}
	if musicabstract.IsChecked != musicabstractOther.IsChecked {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "IsChecked"))
	}
	if musicabstract.PitchHeight != musicabstractOther.PitchHeight {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "PitchHeight"))
	}
	if musicabstract.NbOfBeatsInTheme != musicabstractOther.NbOfBeatsInTheme {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "NbOfBeatsInTheme"))
	}
	if musicabstract.BeatsPerSecond != musicabstractOther.BeatsPerSecond {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "BeatsPerSecond"))
	}
	if musicabstract.FirstVoiceShiftX != musicabstractOther.FirstVoiceShiftX {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "FirstVoiceShiftX"))
	}
	if musicabstract.FirstVoiceShiftY != musicabstractOther.FirstVoiceShiftY {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "FirstVoiceShiftY"))
	}
	if musicabstract.PitchDifference != musicabstractOther.PitchDifference {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "PitchDifference"))
	}
	if musicabstract.Level != musicabstractOther.Level {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "Level"))
	}
	if musicabstract.ActualBeatsTemporalShift != musicabstractOther.ActualBeatsTemporalShift {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ActualBeatsTemporalShift"))
	}
	if musicabstract.IsMinor != musicabstractOther.IsMinor {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "IsMinor"))
	}
	if musicabstract.ThemeBinaryEncoding != musicabstractOther.ThemeBinaryEncoding {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ThemeBinaryEncoding"))
	}
	if musicabstract.BezierControlLengthRatio != musicabstractOther.BezierControlLengthRatio {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "BezierControlLengthRatio"))
	}
	if musicabstract.NbPitchLines != musicabstractOther.NbPitchLines {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "NbPitchLines"))
	}
	if musicabstract.NbBeatLines != musicabstractOther.NbBeatLines {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "NbBeatLines"))
	}
	if musicabstract.OriginX != musicabstractOther.OriginX {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "OriginX"))
	}
	if musicabstract.OriginY != musicabstractOther.OriginY {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "OriginY"))
	}
	if musicabstract.ScoreScale != musicabstractOther.ScoreScale {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ScoreScale"))
	}
	if musicabstract.ShowFirstVoice != musicabstractOther.ShowFirstVoice {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowFirstVoice"))
	}
	if musicabstract.ShowFirstVoiceShiftRight != musicabstractOther.ShowFirstVoiceShiftRight {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowFirstVoiceShiftRight"))
	}
	if musicabstract.ShowSecondVoice != musicabstractOther.ShowSecondVoice {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowSecondVoice"))
	}
	if musicabstract.ShowSecondVoiceShiftRight != musicabstractOther.ShowSecondVoiceShiftRight {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowSecondVoiceShiftRight"))
	}
	if musicabstract.ShowFirstVoiceNotes != musicabstractOther.ShowFirstVoiceNotes {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowFirstVoiceNotes"))
	}
	if musicabstract.ShowFirstVoiceNotesShiftRight != musicabstractOther.ShowFirstVoiceNotesShiftRight {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowFirstVoiceNotesShiftRight"))
	}
	if musicabstract.ShowSecondVoiceNotes != musicabstractOther.ShowSecondVoiceNotes {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowSecondVoiceNotes"))
	}
	if musicabstract.ShowSecondVoiceNotesShiftRight != musicabstractOther.ShowSecondVoiceNotesShiftRight {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowSecondVoiceNotesShiftRight"))
	}
	if musicabstract.IsComposerNodeExpanded != musicabstractOther.IsComposerNodeExpanded {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "IsComposerNodeExpanded"))
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
	if plant3ddiagram.IsHiddenLeaves3DShape != plant3ddiagramOther.IsHiddenLeaves3DShape {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "IsHiddenLeaves3DShape"))
	}
	if (plant3ddiagram.Leaves3DShape == nil) != (plant3ddiagramOther.Leaves3DShape == nil) {
		diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "Leaves3DShape"))
	} else if plant3ddiagram.Leaves3DShape != nil && plant3ddiagramOther.Leaves3DShape != nil {
		if plant3ddiagram.Leaves3DShape != plant3ddiagramOther.Leaves3DShape {
			diffs = append(diffs, plant3ddiagram.GongMarshallField(stage, "Leaves3DShape"))
		}
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
	if (plantabstract.TubeVaseAbstract == nil) != (plantabstractOther.TubeVaseAbstract == nil) {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "TubeVaseAbstract"))
	} else if plantabstract.TubeVaseAbstract != nil && plantabstractOther.TubeVaseAbstract != nil {
		if plantabstract.TubeVaseAbstract != plantabstractOther.TubeVaseAbstract {
			diffs = append(diffs, plantabstract.GongMarshallField(stage, "TubeVaseAbstract"))
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
	if (plantabstract.MusicAbstract == nil) != (plantabstractOther.MusicAbstract == nil) {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "MusicAbstract"))
	} else if plantabstract.MusicAbstract != nil && plantabstractOther.MusicAbstract != nil {
		if plantabstract.MusicAbstract != plantabstractOther.MusicAbstract {
			diffs = append(diffs, plantabstract.GongMarshallField(stage, "MusicAbstract"))
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
		ops := stage.Diff(
			plantabstract,
			"Plant2DDiagrams",
			len(plantabstractOther.Plant2DDiagrams),
			len(plantabstract.Plant2DDiagrams),
			func(i, j int) bool {
				return plantabstractOther.Plant2DDiagrams[i] == plantabstract.Plant2DDiagrams[j]
			},
			func(j int) string {
				return plantabstract.Plant2DDiagrams[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			plantabstract,
			"Plant3DDiagrams",
			len(plantabstractOther.Plant3DDiagrams),
			len(plantabstract.Plant3DDiagrams),
			func(i, j int) bool {
				return plantabstractOther.Plant3DDiagrams[i] == plantabstract.Plant3DDiagrams[j]
			},
			func(j int) string {
				return plantabstract.Plant3DDiagrams[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			plantabstract,
			"Vase2DDiagrams",
			len(plantabstractOther.Vase2DDiagrams),
			len(plantabstract.Vase2DDiagrams),
			func(i, j int) bool {
				return plantabstractOther.Vase2DDiagrams[i] == plantabstract.Vase2DDiagrams[j]
			},
			func(j int) string {
				return plantabstract.Vase2DDiagrams[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if plantabstract.IsTubeVase3DDiagramsNodeExpanded != plantabstractOther.IsTubeVase3DDiagramsNodeExpanded {
		diffs = append(diffs, plantabstract.GongMarshallField(stage, "IsTubeVase3DDiagramsNodeExpanded"))
	}
	TubeVase3DDiagramsDifferent := false
	if len(plantabstract.TubeVase3DDiagrams) != len(plantabstractOther.TubeVase3DDiagrams) {
		TubeVase3DDiagramsDifferent = true
	} else {
		for i := range plantabstract.TubeVase3DDiagrams {
			if (plantabstract.TubeVase3DDiagrams[i] == nil) != (plantabstractOther.TubeVase3DDiagrams[i] == nil) {
				TubeVase3DDiagramsDifferent = true
				break
			} else if plantabstract.TubeVase3DDiagrams[i] != nil && plantabstractOther.TubeVase3DDiagrams[i] != nil {
				// this is a pointer comparaison
				if plantabstract.TubeVase3DDiagrams[i] != plantabstractOther.TubeVase3DDiagrams[i] {
					TubeVase3DDiagramsDifferent = true
					break
				}
			}
		}
	}
	if TubeVase3DDiagramsDifferent {
		ops := stage.Diff(
			plantabstract,
			"TubeVase3DDiagrams",
			len(plantabstractOther.TubeVase3DDiagrams),
			len(plantabstract.TubeVase3DDiagrams),
			func(i, j int) bool {
				return plantabstractOther.TubeVase3DDiagrams[i] == plantabstract.TubeVase3DDiagrams[j]
			},
			func(j int) string {
				return plantabstract.TubeVase3DDiagrams[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			plantabstract,
			"Stool2DDiagrams",
			len(plantabstractOther.Stool2DDiagrams),
			len(plantabstract.Stool2DDiagrams),
			func(i, j int) bool {
				return plantabstractOther.Stool2DDiagrams[i] == plantabstract.Stool2DDiagrams[j]
			},
			func(j int) string {
				return plantabstract.Stool2DDiagrams[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			plantabstract,
			"Stool3DDiagrams",
			len(plantabstractOther.Stool3DDiagrams),
			len(plantabstract.Stool3DDiagrams),
			func(i, j int) bool {
				return plantabstractOther.Stool3DDiagrams[i] == plantabstract.Stool3DDiagrams[j]
			},
			func(j int) string {
				return plantabstract.Stool3DDiagrams[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			plantabstract,
			"Clock2DDiagrams",
			len(plantabstractOther.Clock2DDiagrams),
			len(plantabstract.Clock2DDiagrams),
			func(i, j int) bool {
				return plantabstractOther.Clock2DDiagrams[i] == plantabstract.Clock2DDiagrams[j]
			},
			func(j int) string {
				return plantabstract.Clock2DDiagrams[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			plantabstract,
			"Clock3DDiagrams",
			len(plantabstractOther.Clock3DDiagrams),
			len(plantabstract.Clock3DDiagrams),
			func(i, j int) bool {
				return plantabstractOther.Clock3DDiagrams[i] == plantabstract.Clock3DDiagrams[j]
			},
			func(j int) string {
				return plantabstract.Clock3DDiagrams[j].GongGetIdentifier(stage)
			},
		)
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
	if (tubevase3ddiagram.Rendered3DShape == nil) != (tubevase3ddiagramOther.Rendered3DShape == nil) {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
	} else if tubevase3ddiagram.Rendered3DShape != nil && tubevase3ddiagramOther.Rendered3DShape != nil {
		if tubevase3ddiagram.Rendered3DShape != tubevase3ddiagramOther.Rendered3DShape {
			diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "Rendered3DShape"))
		}
	}
	if (tubevase3ddiagram.SampledPoints3DShape == nil) != (tubevase3ddiagramOther.SampledPoints3DShape == nil) {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
	} else if tubevase3ddiagram.SampledPoints3DShape != nil && tubevase3ddiagramOther.SampledPoints3DShape != nil {
		if tubevase3ddiagram.SampledPoints3DShape != tubevase3ddiagramOther.SampledPoints3DShape {
			diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "SampledPoints3DShape"))
		}
	}
	if (tubevase3ddiagram.OriginalPoints3DShape == nil) != (tubevase3ddiagramOther.OriginalPoints3DShape == nil) {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "OriginalPoints3DShape"))
	} else if tubevase3ddiagram.OriginalPoints3DShape != nil && tubevase3ddiagramOther.OriginalPoints3DShape != nil {
		if tubevase3ddiagram.OriginalPoints3DShape != tubevase3ddiagramOther.OriginalPoints3DShape {
			diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "OriginalPoints3DShape"))
		}
	}
	if (tubevase3ddiagram.Angle0Shape == nil) != (tubevase3ddiagramOther.Angle0Shape == nil) {
		diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "Angle0Shape"))
	} else if tubevase3ddiagram.Angle0Shape != nil && tubevase3ddiagramOther.Angle0Shape != nil {
		if tubevase3ddiagram.Angle0Shape != tubevase3ddiagramOther.Angle0Shape {
			diffs = append(diffs, tubevase3ddiagram.GongMarshallField(stage, "Angle0Shape"))
		}
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
