// generated code - do not edit
package models

import "time"

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !IsStagedPointerToGongstruct(stage, *element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by ArcNormalVectorShape
func (arcnormalvectorshape *ArcNormalVectorShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ArcNormalVectorShapeGrid
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &arcnormalvectorshapegrid.ArcNormalVectorShapes) || modified
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
	modified = GongCleanSlice(stage, &basevectorshapegrid.BaseVectorShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CircleGridShape
func (circlegridshape *CircleGridShape) GongClean(stage *Stage) (modified bool) {
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
	modified = GongCleanSlice(stage, &endarcshapegrid.EndArcShapes) || modified
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
	modified = GongCleanSlice(stage, &endhalfwayarcshapegrid.EndHalfwayArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ExplanationTextShape
func (explanationtextshape *ExplanationTextShape) GongClean(stage *Stage) (modified bool) {
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
	modified = GongCleanPointer(stage, &growthcurve2d.StartHalfwayArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &growthcurve2d.EndHalfwayArcShapeGrid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GrowthCurve2DRibbon
func (growthcurve2dribbon *GrowthCurve2DRibbon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &growthcurve2dribbon.GrowthCurve2DRibbonStartShapes) || modified
	modified = GongCleanSlice(stage, &growthcurve2dribbon.GrowthCurve2DRibbonEndShapes) || modified
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
	modified = GongCleanSlice(stage, &growthcurverhombusgridshape.GrowthCurveRhombusShapes) || modified
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
	modified = GongCleanSlice(stage, &initialrhombusgridshape.InitialRhombusShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by InitialRhombusShape
func (initialrhombusshape *InitialRhombusShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &library.SubLibraries) || modified
	modified = GongCleanSlice(stage, &library.Plants) || modified
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
	modified = GongCleanSlice(stage, &midarcvectorshapegrid.MidArcVectorShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartiallyGrowthCurve2DRibbon
func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonStartShapes) || modified
	modified = GongCleanSlice(stage, &partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonEndShapes) || modified
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
	modified = GongCleanSlice(stage, &partiallygrowthcurve2dtrajectory.PartiallyGrowthCurve2DTrajectoryShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartiallyGrowthCurve2DTrajectoryShape
func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongClean(stage *Stage) (modified bool) {
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
	modified = GongCleanSlice(stage, &perpendicularvectorgrid.PerpendicularVectors) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PerpendicularVectorGridHalfway
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &perpendicularvectorgridhalfway.PerpendicularVectorHalfways) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PerpendicularVectorHalfway
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Plant
func (plant *Plant) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &plant.PlantDiagrams) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &plant.AxesShape) || modified
	modified = GongCleanPointer(stage, &plant.RhombusStuff) || modified
	modified = GongCleanPointer(stage, &plant.GrowthVectorShape) || modified
	modified = GongCleanPointer(stage, &plant.PerpendicularVectorGrid) || modified
	modified = GongCleanPointer(stage, &plant.PerpendicularVectorGridHalfway) || modified
	modified = GongCleanPointer(stage, &plant.BaseVectorShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.ArcNormalVectorShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.StartArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.TopStartArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.EndArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.TopEndArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.ShiftedBottomTopStartArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.MidArcVectorShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.TopMidArcVectorShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.StartHalfwayArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.TopStartHalfwayArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.EndHalfwayArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.TopEndHalfwayArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.StackOfRotatedGrowthCurve2D) || modified
	modified = GongCleanPointer(stage, &plant.TopStackOfRotatedGrowthCurve2D) || modified
	modified = GongCleanPointer(stage, &plant.GrowthCurve2D) || modified
	modified = GongCleanPointer(stage, &plant.TopGrowthCurve2D) || modified
	modified = GongCleanPointer(stage, &plant.StackOfGrowthCurve2D) || modified
	modified = GongCleanPointer(stage, &plant.TopStackOfGrowthCurve2D) || modified
	modified = GongCleanPointer(stage, &plant.StackOfGrowthCurve2DRibbon) || modified
	modified = GongCleanPointer(stage, &plant.StackOfRotatedGrowthCurve2DRibbon) || modified
	modified = GongCleanPointer(stage, &plant.GrowthCurve2DRibbon) || modified
	modified = GongCleanPointer(stage, &plant.ShiftedRightGrowthCurve2DRibbon) || modified
	modified = GongCleanPointer(stage, &plant.PartiallyGrowthCurve2DRibbon) || modified
	modified = GongCleanPointer(stage, &plant.PartiallyGrowthCurve2DTrajectory) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PlantCircumferenceShape
func (plantcircumferenceshape *PlantCircumferenceShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PlantDiagram
func (plantdiagram *PlantDiagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &plantdiagram.Rendered3DShape) || modified
	modified = GongCleanPointer(stage, &plantdiagram.GrowthCurve2DRibbon) || modified
	modified = GongCleanPointer(stage, &plantdiagram.ShiftedRightGrowthCurve2DRibbon) || modified
	modified = GongCleanPointer(stage, &plantdiagram.TorusStackShape) || modified
	modified = GongCleanPointer(stage, &plantdiagram.VerticalTorusStackShape) || modified
	modified = GongCleanPointer(stage, &plantdiagram.PartiallyRotatedTorusShape) || modified
	modified = GongCleanPointer(stage, &plantdiagram.StackOfPartiallyRotatedTorusShape) || modified
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
	modified = GongCleanPointer(stage, &rhombusstuff.ReferenceRhombus) || modified
	modified = GongCleanPointer(stage, &rhombusstuff.PlantCircumferenceShape) || modified
	modified = GongCleanPointer(stage, &rhombusstuff.GridPathShape) || modified
	modified = GongCleanPointer(stage, &rhombusstuff.InitialRhombusGridShape) || modified
	modified = GongCleanPointer(stage, &rhombusstuff.ExplanationTextShape) || modified
	modified = GongCleanPointer(stage, &rhombusstuff.RotatedReferenceRhombus) || modified
	modified = GongCleanPointer(stage, &rhombusstuff.RotatedPlantCircumferenceShape) || modified
	modified = GongCleanPointer(stage, &rhombusstuff.RotatedGridPathShape) || modified
	modified = GongCleanPointer(stage, &rhombusstuff.RotatedRhombusGridShape2) || modified
	modified = GongCleanPointer(stage, &rhombusstuff.GrowthCurveRhombusGridShape) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RotatedRhombusGridShape
func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &rotatedrhombusgridshape.RotatedRhombusShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RotatedRhombusShape
func (rotatedrhombusshape *RotatedRhombusShape) GongClean(stage *Stage) (modified bool) {
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
	modified = GongCleanSlice(stage, &shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes) || modified
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
	modified = GongCleanSlice(stage, &shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes) || modified
	modified = GongCleanSlice(stage, &shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedLeftStackOfNormalVector
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ShiftedRightGrowthCurve2DRibbon
func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonStartShapes) || modified
	modified = GongCleanSlice(stage, &shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonEndShapes) || modified
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
	modified = GongCleanSlice(stage, &stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes) || modified
	modified = GongCleanSlice(stage, &stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackOfGrowthCurve2DRibbon
func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes) || modified
	modified = GongCleanSlice(stage, &stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes) || modified
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
	modified = GongCleanSlice(stage, &stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes) || modified
	modified = GongCleanSlice(stage, &stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackOfRotatedGrowthCurve2DRibbon
func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonStartShapes) || modified
	modified = GongCleanSlice(stage, &stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonEndShapes) || modified
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
	modified = GongCleanSlice(stage, &startarcshapegrid.StartArcShapes) || modified
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
	modified = GongCleanSlice(stage, &starthalfwayarcshapegrid.StartHalfwayArcShapes) || modified
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
	modified = GongCleanSlice(stage, &topendarcshapegrid.TopEndArcShapes) || modified
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
	modified = GongCleanSlice(stage, &topendhalfwayarcshapegrid.TopEndHalfwayArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopGrowthCurve2D
func (topgrowthcurve2d *TopGrowthCurve2D) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &topgrowthcurve2d.TopStartHalfwayArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &topgrowthcurve2d.TopEndHalfwayArcShapeGrid) || modified
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
	modified = GongCleanSlice(stage, &topmidarcvectorshapegrid.TopMidArcVectorShapes) || modified
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
	modified = GongCleanSlice(stage, &topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes) || modified
	modified = GongCleanSlice(stage, &topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopStackOfRotatedGrowthCurve2D
func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes) || modified
	modified = GongCleanSlice(stage, &topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes) || modified
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
	modified = GongCleanSlice(stage, &topstartarcshapegrid.TopStartArcShapes) || modified
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
	modified = GongCleanSlice(stage, &topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TorusStackShape
func (torusstackshape *TorusStackShape) GongClean(stage *Stage) (modified bool) {
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
