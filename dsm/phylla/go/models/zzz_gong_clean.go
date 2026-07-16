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
	modified = GongCleanPointer(stage, &growthcurve2d.StartArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &growthcurve2d.EndArcShapeGrid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GrowthCurveBezierShape
func (growthcurvebeziershape *GrowthCurveBezierShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GrowthCurveBezierShapeGrid
func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &growthcurvebeziershapegrid.GrowthCurveBezierShapes) || modified
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

// Clean garbage collect unstaged instances that are referenced by NextCircleShape
func (nextcircleshape *NextCircleShape) GongClean(stage *Stage) (modified bool) {
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
	modified = GongCleanPointer(stage, &plant.ReferenceRhombus) || modified
	modified = GongCleanPointer(stage, &plant.PlantCircumferenceShape) || modified
	modified = GongCleanPointer(stage, &plant.GridPathShape) || modified
	modified = GongCleanPointer(stage, &plant.InitialRhombusGridShape) || modified
	modified = GongCleanPointer(stage, &plant.ExplanationTextShape) || modified
	modified = GongCleanPointer(stage, &plant.RotatedReferenceRhombus) || modified
	modified = GongCleanPointer(stage, &plant.RotatedPlantCircumferenceShape) || modified
	modified = GongCleanPointer(stage, &plant.RotatedGridPathShape) || modified
	modified = GongCleanPointer(stage, &plant.RotatedRhombusGridShape2) || modified
	modified = GongCleanPointer(stage, &plant.GrowthCurveRhombusGridShape) || modified
	modified = GongCleanPointer(stage, &plant.GrowthVectorShape) || modified
	modified = GongCleanPointer(stage, &plant.PerpendicularVectorGrid) || modified
	modified = GongCleanPointer(stage, &plant.PerpendicularVectorGridHalfway) || modified
	modified = GongCleanPointer(stage, &plant.BaseVectorShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.ArcNormalVectorShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.StartArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.TopStartArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.ShiftedBottomTopStartArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.MidArcVectorShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.TopMidArcVectorShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.EndArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.TopEndArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.GrowthCurveBezierShapeGrid) || modified
	modified = GongCleanPointer(stage, &plant.StackOfGrowthCurve) || modified
	modified = GongCleanPointer(stage, &plant.TopStackOfGrowthCurve) || modified
	modified = GongCleanPointer(stage, &plant.ShiftedLeftStackOfGrowthCurve) || modified
	modified = GongCleanPointer(stage, &plant.ShiftedLeftStackOfNormalVector) || modified
	modified = GongCleanPointer(stage, &plant.GrowthCurve2D) || modified
	modified = GongCleanPointer(stage, &plant.TopGrowthCurve2D) || modified
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

// Clean garbage collect unstaged instances that are referenced by StackGrowthCurveEndArcShape
func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackGrowthCurveStartArcShape
func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by StackOfGrowthCurve
func (stackofgrowthcurve *StackOfGrowthCurve) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &stackofgrowthcurve.StackGrowthCurveStartArcShapes) || modified
	modified = GongCleanSlice(stage, &stackofgrowthcurve.StackGrowthCurveEndArcShapes) || modified
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

// Clean garbage collect unstaged instances that are referenced by TopGrowthCurve2D
func (topgrowthcurve2d *TopGrowthCurve2D) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &topgrowthcurve2d.TopStartArcShapeGrid) || modified
	modified = GongCleanPointer(stage, &topgrowthcurve2d.TopEndArcShapeGrid) || modified
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

// Clean garbage collect unstaged instances that are referenced by TopStackGrowthCurveEndArcShape
func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopStackGrowthCurveStartArcShape
func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopStackOfGrowthCurve
func (topstackofgrowthcurve *TopStackOfGrowthCurve) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes) || modified
	modified = GongCleanSlice(stage, &topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes) || modified
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
