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

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Angle0Shape
func (angle0shape *Angle0Shape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ArcNormalVectorShape
func (arcnormalvectorshape *ArcNormalVectorShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ArcNormalVectorShapeGrid
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&arcnormalvectorshapegrid.ArcNormalVectorShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AxesShape
func (axesshape *AxesShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by BaseVectorShape
func (basevectorshape *BaseVectorShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by BottomCurvePlane1Shape
func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by BottomCurvePlane2Shape
func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ChosenP1P2PairShape
func (chosenp1p2pairshape *ChosenP1P2PairShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CircleGridShape
func (circlegridshape *CircleGridShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Circumference3DShape
func (circumference3dshape *Circumference3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Clock2DDiagram
func (clock2ddiagram *Clock2DDiagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by ClockAbstract
func (clockabstract *ClockAbstract) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ClockTopCurveShape
func (clocktopcurveshape *ClockTopCurveShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CutLine3DShape
func (cutline3dshape *CutLine3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by EndArcShape
func (endarcshape *EndArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by EndArcShapeGrid
func (endarcshapegrid *EndArcShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&endarcshapegrid.EndArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by EndHalfwayArcShape
func (endhalfwayarcshape *EndHalfwayArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by ExplanationTextShape
func (explanationtextshape *ExplanationTextShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Eye3DShape
func (eye3dshape *Eye3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by EyeCornersSampledPoints3DShape
func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by EyeSampledPoints3DShape
func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by EyeSeatBottomCurveShape
func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by EyeStoolBottomCurveShape
func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by EyeVolume3DShape
func (eyevolume3dshape *EyeVolume3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GridPathShape
func (gridpathshape *GridPathShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by GrowthCurve2DRibbonEndShape
func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GrowthCurve2DRibbonStartShape
func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by GrowthCurveRhombusShape
func (growthcurverhombusshape *GrowthCurveRhombusShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GrowthVectorShape
func (growthvectorshape *GrowthVectorShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by InitialRhombusShape
func (initialrhombusshape *InitialRhombusShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Key3DShape
func (key3dshape *Key3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by KeyHole3DShape
func (keyhole3dshape *KeyHole3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by KeyHoleShape
func (keyholeshape *KeyHoleShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Leaves3DShape
func (leaves3dshape *Leaves3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by MidArcVectorShape
func (midarcvectorshape *MidArcVectorShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by MusicAbstract
func (musicabstract *MusicAbstract) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by OriginalPoints3DShape
func (originalpoints3dshape *OriginalPoints3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ParastichyMCurves3DShape
func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ParastichyNCurves3DShape
func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by PartiallyGrowthCurve2DRibbonEndShape
func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartiallyGrowthCurve2DRibbonStartShape
func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by PartiallyGrowthCurve2DTrajectoryP1CurveShape
func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape
func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartiallyGrowthCurve2DTrajectoryP1PointShape
func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartiallyGrowthCurve2DTrajectoryP2CurveShape
func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartiallyGrowthCurve2DTrajectoryP2PointShape
func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartiallyGrowthCurve2DTrajectoryShape
func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartiallyRotatedSeatBottomCurveShape
func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartiallyRotatedSeatTopCurveShape
func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartiallyRotatedTorusShape
func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PerpendicularVector
func (perpendicularvector *PerpendicularVector) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by PerpendicularVectorHalfway
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Plant2DDiagram
func (plant2ddiagram *Plant2DDiagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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
	modified = stage.CleanPointer(&plantabstract.StoolAbstract) || modified
	modified = stage.CleanPointer(&plantabstract.ClockAbstract) || modified
	modified = stage.CleanPointer(&plantabstract.MusicAbstract) || modified
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

// Clean garbage collect unstaged instances that are referenced by PlantCircumferenceShape
func (plantcircumferenceshape *PlantCircumferenceShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PointsAndLines3DShape
func (pointsandlines3dshape *PointsAndLines3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PxShape
func (pxshape *PxShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Rendered3DShape
func (rendered3dshape *Rendered3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RhombusShape
func (rhombusshape *RhombusShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by RotatedRhombusShape
func (rotatedrhombusshape *RotatedRhombusShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RotatedSampledPoints3DShape
func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RotatedSeatAndLegs3DShape
func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SampledPoints3DShape
func (sampledpoints3dshape *SampledPoints3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Seat3DShape
func (seat3dshape *Seat3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SeatAndLegs3DShape
func (seatandlegs3dshape *SeatAndLegs3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SeatBottomCurveShape
func (seatbottomcurveshape *SeatBottomCurveShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SeatTopCurveShape
func (seattopcurveshape *SeatTopCurveShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedBottomTopStartArcShape
func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by ShiftedLeftGrowthCurve2DRibbonEndShape
func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedLeftGrowthCurve2DRibbonStartShape
func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape
func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape
func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedLeftStackGrowthCurveEndArcShape
func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedLeftStackGrowthCurveStartArcShape
func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedLeftStackNormalVector
func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by ShiftedRightGrowthCurve2DRibbonEndShape
func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedRightGrowthCurve2DRibbonStartShape
func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackGrowthCurve2DEndHalfwayArcShape
func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackGrowthCurve2DRibbonEndShape
func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackGrowthCurve2DRibbonStartShape
func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackGrowthCurve2DStartHalfwayArcShape
func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by StackOfGrowthCurve2DByGrowthVector
func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by StackOfPartiallyRotatedTorusShape
func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by StackRotatedGrowthCurve2DEndArcShape
func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackRotatedGrowthCurve2DRibbonEndShape
func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackRotatedGrowthCurve2DRibbonStartShape
func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackRotatedGrowthCurve2DStartArcShape
func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StartArcShape
func (startarcshape *StartArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by StartHalfwayArcShape
func (starthalfwayarcshape *StartHalfwayArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by StemCylinder3DShape
func (stemcylinder3dshape *StemCylinder3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Stool2DDiagram
func (stool2ddiagram *Stool2DDiagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by StoolAbstract
func (stoolabstract *StoolAbstract) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TiledFloor3DShape
func (tiledfloor3dshape *TiledFloor3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopCurvePlane1Shape
func (topcurveplane1shape *TopCurvePlane1Shape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopCurvePlane2Shape
func (topcurveplane2shape *TopCurvePlane2Shape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopEndArcShape
func (topendarcshape *TopEndArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopEndArcShapeGrid
func (topendarcshapegrid *TopEndArcShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topendarcshapegrid.TopEndArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopEndHalfwayArcShape
func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by TopMidArcVectorShape
func (topmidarcvectorshape *TopMidArcVectorShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopMidArcVectorShapeGrid
func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topmidarcvectorshapegrid.TopMidArcVectorShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopStackGrowthCurve2DEndHalfwayArcShape
func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopStackGrowthCurve2DStartHalfwayArcShape
func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by TopStackOfRotatedGrowthCurve2DEndArcShape
func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopStackOfRotatedGrowthCurve2DStartArcShape
func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopStartArcShape
func (topstartarcshape *TopStartArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by TopStartHalfwayArcShape
func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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

// Clean garbage collect unstaged instances that are referenced by Torus3DShape
func (torus3dshape *Torus3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TorusEdge3DShape
func (torusedge3dshape *TorusEdge3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TorusStackShape
func (torusstackshape *TorusStackShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TrapezeVolume3DShape
func (trapezevolume3dshape *TrapezeVolume3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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
	modified = stage.CleanPointer(&tubevase3ddiagram.TrapezeVolume3DShape) || modified
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

// Clean garbage collect unstaged instances that are referenced by Vase2DDiagram
func (vase2ddiagram *Vase2DDiagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalTorusStackShape
func (verticaltorusstackshape *VerticalTorusStackShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by VolumeKey3DShape
func (volumekey3dshape *VolumeKey3DShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
