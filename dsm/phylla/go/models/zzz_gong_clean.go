// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

type GongCleaner interface {
	GongClean(stage *Stage) (modified bool)
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by ArcNormalVectorShapeGrid
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&arcnormalvectorshapegrid.ArcNormalVectorShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by BaseVectorShapeGrid
func (basevectorshapegrid *BaseVectorShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&basevectorshapegrid.BaseVectorShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Clock3DDiagram
func (clock3ddiagram *Clock3DDiagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&clock3ddiagram.ClockTopCurveShape) || modified
	modified = stage.CleanPointer(&clock3ddiagram.Torus3DShape) || modified
	modified = stage.CleanPointer(&clock3ddiagram.SampledPoints3DShape) || modified
	modified = stage.CleanPointer(&clock3ddiagram.TiledFloor3DShape) || modified
	modified = stage.CleanPointer(&clock3ddiagram.Rendered3DShape) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EndArcShapeGrid
func (endarcshapegrid *EndArcShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&endarcshapegrid.EndArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by EndHalfwayArcShapeGrid
func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&endhalfwayarcshapegrid.EndHalfwayArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GrowthCurve2D
func (growthcurve2d *GrowthCurve2D) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&growthcurve2d.StartHalfwayArcShapeGrid) || modified
	modified = stage.CleanPointer(&growthcurve2d.EndHalfwayArcShapeGrid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GrowthCurve2DRibbon
func (growthcurve2dribbon *GrowthCurve2DRibbon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&growthcurve2dribbon.GrowthCurve2DRibbonStartShapes) || modified
	modified = stage.CleanSlice(&growthcurve2dribbon.GrowthCurve2DRibbonEndShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GrowthCurveRhombusGridShape
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&growthcurverhombusgridshape.GrowthCurveRhombusShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by InitialRhombusGridShape
func (initialrhombusgridshape *InitialRhombusGridShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&initialrhombusgridshape.InitialRhombusShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&library.Plants) || modified
	modified = stage.CleanSlice(&library.SubLibraries) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MidArcVectorShapeGrid
func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&midarcvectorshapegrid.MidArcVectorShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartiallyGrowthCurve2DRibbon
func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonStartShapes) || modified
	modified = stage.CleanSlice(&partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonEndShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartiallyGrowthCurve2DTrajectory
func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&partiallygrowthcurve2dtrajectory.PartiallyGrowthCurve2DTrajectoryShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartiallyGrowthCurve2DTrajectoryP1P2
func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&partiallygrowthcurve2dtrajectoryp1p2.P1PointShapes) || modified
	modified = stage.CleanSlice(&partiallygrowthcurve2dtrajectoryp1p2.P2PointShapes) || modified
	modified = stage.CleanSlice(&partiallygrowthcurve2dtrajectoryp1p2.P1CurveShapes) || modified
	modified = stage.CleanSlice(&partiallygrowthcurve2dtrajectoryp1p2.P2CurveShapes) || modified
	modified = stage.CleanSlice(&partiallygrowthcurve2dtrajectoryp1p2.P1P2PairLineShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PerpendicularVectorGrid
func (perpendicularvectorgrid *PerpendicularVectorGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&perpendicularvectorgrid.PerpendicularVectors) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PerpendicularVectorGridHalfway
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&perpendicularvectorgridhalfway.PerpendicularVectorHalfways) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Plant3DDiagram
func (plant3ddiagram *Plant3DDiagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&plant3ddiagram.StemCylinder3DShape) || modified
	modified = stage.CleanPointer(&plant3ddiagram.ParastichyNCurves3DShape) || modified
	modified = stage.CleanPointer(&plant3ddiagram.ParastichyMCurves3DShape) || modified
	modified = stage.CleanPointer(&plant3ddiagram.CutLine3DShape) || modified
	modified = stage.CleanPointer(&plant3ddiagram.Circumference3DShape) || modified
	modified = stage.CleanPointer(&plant3ddiagram.TiledFloor3DShape) || modified
	modified = stage.CleanPointer(&plant3ddiagram.Leaves3DShape) || modified
	modified = stage.CleanPointer(&plant3ddiagram.Rendered3DShape) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PlantAbstract
func (plantabstract *PlantAbstract) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&plantabstract.Plant2DDiagrams) || modified
	modified = stage.CleanSlice(&plantabstract.Plant3DDiagrams) || modified
	modified = stage.CleanSlice(&plantabstract.Vase2DDiagrams) || modified
	modified = stage.CleanSlice(&plantabstract.TubeVase3DDiagrams) || modified
	modified = stage.CleanSlice(&plantabstract.Stool2DDiagrams) || modified
	modified = stage.CleanSlice(&plantabstract.Stool3DDiagrams) || modified
	modified = stage.CleanSlice(&plantabstract.Clock2DDiagrams) || modified
	modified = stage.CleanSlice(&plantabstract.Clock3DDiagrams) || modified
	// insertion point per field
	modified = stage.CleanPointer(&plantabstract.TubeVaseAbstract) || modified
	modified = stage.CleanPointer(&plantabstract.AxesShape) || modified
	modified = stage.CleanPointer(&plantabstract.RhombusStuff) || modified
	modified = stage.CleanPointer(&plantabstract.GrowthVectorShape) || modified
	modified = stage.CleanPointer(&plantabstract.PerpendicularVectorGrid) || modified
	modified = stage.CleanPointer(&plantabstract.BaseVectorShapeGrid) || modified
	modified = stage.CleanPointer(&plantabstract.ArcNormalVectorShapeGrid) || modified
	modified = stage.CleanPointer(&plantabstract.StartArcShapeGrid) || modified
	modified = stage.CleanPointer(&plantabstract.MidArcVectorShapeGrid) || modified
	modified = stage.CleanPointer(&plantabstract.EndArcShapeGrid) || modified
	modified = stage.CleanPointer(&plantabstract.GrowthCurve2D) || modified
	modified = stage.CleanPointer(&plantabstract.StackOfGrowthCurve2DByGrowthVector) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RhombusStuff
func (rhombusstuff *RhombusStuff) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&rhombusstuff.ReferenceRhombus) || modified
	modified = stage.CleanPointer(&rhombusstuff.PlantCircumferenceShape) || modified
	modified = stage.CleanPointer(&rhombusstuff.GridPathShape) || modified
	modified = stage.CleanPointer(&rhombusstuff.InitialRhombusGridShape) || modified
	modified = stage.CleanPointer(&rhombusstuff.ExplanationTextShape) || modified
	modified = stage.CleanPointer(&rhombusstuff.RotatedReferenceRhombus) || modified
	modified = stage.CleanPointer(&rhombusstuff.RotatedPlantCircumferenceShape) || modified
	modified = stage.CleanPointer(&rhombusstuff.RotatedGridPathShape) || modified
	modified = stage.CleanPointer(&rhombusstuff.RotatedRhombusGridShape2) || modified
	modified = stage.CleanPointer(&rhombusstuff.GrowthCurveRhombusGridShape) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RotatedRhombusGridShape
func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&rotatedrhombusgridshape.RotatedRhombusShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedBottomTopStartArcShapeGrid
func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedLeftGrowthCurve2DRibbon
func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&shiftedleftgrowthcurve2dribbon.ShiftedLeftGrowthCurve2DRibbonStartShapes) || modified
	modified = stage.CleanSlice(&shiftedleftgrowthcurve2dribbon.ShiftedLeftGrowthCurve2DRibbonEndShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedLeftPartiallyGrowthCurve2DRibbon
func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&shiftedleftpartiallygrowthcurve2dribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes) || modified
	modified = stage.CleanSlice(&shiftedleftpartiallygrowthcurve2dribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedLeftStackOfGrowthCurve
func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes) || modified
	modified = stage.CleanSlice(&shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedLeftStackOfNormalVector
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedRightGrowthCurve2DRibbon
func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonStartShapes) || modified
	modified = stage.CleanSlice(&shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonEndShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackOfGrowthCurve2D
func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes) || modified
	modified = stage.CleanSlice(&stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackOfGrowthCurve2DRibbon
func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes) || modified
	modified = stage.CleanSlice(&stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackOfRotatedGrowthCurve2D
func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes) || modified
	modified = stage.CleanSlice(&stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackOfRotatedGrowthCurve2DRibbon
func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonStartShapes) || modified
	modified = stage.CleanSlice(&stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonEndShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StartArcShapeGrid
func (startarcshapegrid *StartArcShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&startarcshapegrid.StartArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StartHalfwayArcShapeGrid
func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&starthalfwayarcshapegrid.StartHalfwayArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Stool3DDiagram
func (stool3ddiagram *Stool3DDiagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&stool3ddiagram.SeatTopCurveShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.RotatedSeatTopCurveShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.SeatBottomCurveShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.RotatedSeatBottomCurveShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.Torus3DShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.RotatedTorusShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.SampledPoints3DShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.RotatedSampledPoints3DShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.EyeSampledPoints3DShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.EyeCornersSampledPoints3DShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.Eye3DShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.EyeSeatBottomCurveShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.EyeStoolBottomCurveShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.Seat3DShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.EyeVolume3DShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.SeatAndLegs3DShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.RotatedSeatAndLegs3DShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.TiledFloor3DShape) || modified
	modified = stage.CleanPointer(&stool3ddiagram.Rendered3DShape) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopEndArcShapeGrid
func (topendarcshapegrid *TopEndArcShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topendarcshapegrid.TopEndArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopEndHalfwayArcShapeGrid
func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topendhalfwayarcshapegrid.TopEndHalfwayArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopGrowthCurve2D
func (topgrowthcurve2d *TopGrowthCurve2D) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&topgrowthcurve2d.TopStartHalfwayArcShapeGrid) || modified
	modified = stage.CleanPointer(&topgrowthcurve2d.TopEndHalfwayArcShapeGrid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopMidArcVectorShapeGrid
func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topmidarcvectorshapegrid.TopMidArcVectorShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopStackOfGrowthCurve2D
func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes) || modified
	modified = stage.CleanSlice(&topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopStackOfRotatedGrowthCurve2D
func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes) || modified
	modified = stage.CleanSlice(&topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopStartArcShapeGrid
func (topstartarcshapegrid *TopStartArcShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topstartarcshapegrid.TopStartArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopStartHalfwayArcShapeGrid
func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TubeVase3DDiagram
func (tubevase3ddiagram *TubeVase3DDiagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&tubevase3ddiagram.Rendered3DShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.TorusStackShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.VerticalTorusStackShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.PartiallyRotatedTorusShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.StackOfPartiallyRotatedTorusShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.PointsAndLines3DShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.SampledPoints3DShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.OriginalPoints3DShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.Angle0Shape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.KeyHole3DShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.Key3DShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.VolumeKey3DShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.TorusEdge3DShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.TiledFloor3DShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.TopCurvePlane1Shape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.BottomCurvePlane1Shape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.TopCurvePlane2Shape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.BottomCurvePlane2Shape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.VaseTrapezeRingShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.StackOfVaseTrapezeRingsShape) || modified
	modified = stage.CleanPointer(&tubevase3ddiagram.StackOfRotatedVaseTrapezeRingsShape) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TubeVaseAbstract
func (tubevaseabstract *TubeVaseAbstract) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&tubevaseabstract.PerpendicularVectorGridHalfway) || modified
	modified = stage.CleanPointer(&tubevaseabstract.TopStartArcShapeGrid) || modified
	modified = stage.CleanPointer(&tubevaseabstract.TopEndArcShapeGrid) || modified
	modified = stage.CleanPointer(&tubevaseabstract.ShiftedBottomTopStartArcShapeGrid) || modified
	modified = stage.CleanPointer(&tubevaseabstract.TopMidArcVectorShapeGrid) || modified
	modified = stage.CleanPointer(&tubevaseabstract.StartHalfwayArcShapeGrid) || modified
	modified = stage.CleanPointer(&tubevaseabstract.TopStartHalfwayArcShapeGrid) || modified
	modified = stage.CleanPointer(&tubevaseabstract.EndHalfwayArcShapeGrid) || modified
	modified = stage.CleanPointer(&tubevaseabstract.TopEndHalfwayArcShapeGrid) || modified
	modified = stage.CleanPointer(&tubevaseabstract.StackOfRotatedGrowthCurve2D) || modified
	modified = stage.CleanPointer(&tubevaseabstract.TopStackOfRotatedGrowthCurve2D) || modified
	modified = stage.CleanPointer(&tubevaseabstract.TopGrowthCurve2D) || modified
	modified = stage.CleanPointer(&tubevaseabstract.StackOfGrowthCurve2D) || modified
	modified = stage.CleanPointer(&tubevaseabstract.TopStackOfGrowthCurve2D) || modified
	modified = stage.CleanPointer(&tubevaseabstract.StackOfGrowthCurve2DRibbon) || modified
	modified = stage.CleanPointer(&tubevaseabstract.StackOfRotatedGrowthCurve2DRibbon) || modified
	modified = stage.CleanPointer(&tubevaseabstract.GrowthCurve2DRibbon) || modified
	modified = stage.CleanPointer(&tubevaseabstract.ShiftedRightGrowthCurve2DRibbon) || modified
	modified = stage.CleanPointer(&tubevaseabstract.ShiftedLeftGrowthCurve2DRibbon) || modified
	modified = stage.CleanPointer(&tubevaseabstract.PartiallyGrowthCurve2DRibbon) || modified
	modified = stage.CleanPointer(&tubevaseabstract.ShiftedLeftPartiallyGrowthCurve2DRibbon) || modified
	modified = stage.CleanPointer(&tubevaseabstract.PartiallyGrowthCurve2DTrajectory) || modified
	modified = stage.CleanPointer(&tubevaseabstract.PartiallyGrowthCurve2DTrajectoryP1P2) || modified
	modified = stage.CleanPointer(&tubevaseabstract.PxShape) || modified
	modified = stage.CleanPointer(&tubevaseabstract.ChosenP1P2PairShape) || modified
	modified = stage.CleanPointer(&tubevaseabstract.KeyHoleShape) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		if cleaner, ok := any(instance).(GongCleaner); ok {
			modified = cleaner.GongClean(stage) || modified
		}
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
