// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
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

	case *CircleGridShape:
		ok = stage.IsStagedCircleGridShape(target)

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

	case *GridPathShape:
		ok = stage.IsStagedGridPathShape(target)

	case *GrowthCurve2D:
		ok = stage.IsStagedGrowthCurve2D(target)

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

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *MidArcVectorShape:
		ok = stage.IsStagedMidArcVectorShape(target)

	case *MidArcVectorShapeGrid:
		ok = stage.IsStagedMidArcVectorShapeGrid(target)

	case *PerpendicularVector:
		ok = stage.IsStagedPerpendicularVector(target)

	case *PerpendicularVectorGrid:
		ok = stage.IsStagedPerpendicularVectorGrid(target)

	case *PerpendicularVectorGridHalfway:
		ok = stage.IsStagedPerpendicularVectorGridHalfway(target)

	case *PerpendicularVectorHalfway:
		ok = stage.IsStagedPerpendicularVectorHalfway(target)

	case *Plant:
		ok = stage.IsStagedPlant(target)

	case *PlantCircumferenceShape:
		ok = stage.IsStagedPlantCircumferenceShape(target)

	case *PlantDiagram:
		ok = stage.IsStagedPlantDiagram(target)

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

	case *ShiftedBottomTopStartArcShape:
		ok = stage.IsStagedShiftedBottomTopStartArcShape(target)

	case *ShiftedBottomTopStartArcShapeGrid:
		ok = stage.IsStagedShiftedBottomTopStartArcShapeGrid(target)

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

	case *StackOfGrowthCurve2DRibbon:
		ok = stage.IsStagedStackOfGrowthCurve2DRibbon(target)

	case *StackOfRotatedGrowthCurve2D:
		ok = stage.IsStagedStackOfRotatedGrowthCurve2D(target)

	case *StackRotatedGrowthCurve2DEndArcShape:
		ok = stage.IsStagedStackRotatedGrowthCurve2DEndArcShape(target)

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

	case *TorusStackShape:
		ok = stage.IsStagedTorusStackShape(target)

	case *VerticalTorusStackShape:
		ok = stage.IsStagedVerticalTorusStackShape(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
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

	case *CircleGridShape:
		ok = stage.IsStagedCircleGridShape(target)

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

	case *GridPathShape:
		ok = stage.IsStagedGridPathShape(target)

	case *GrowthCurve2D:
		ok = stage.IsStagedGrowthCurve2D(target)

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

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *MidArcVectorShape:
		ok = stage.IsStagedMidArcVectorShape(target)

	case *MidArcVectorShapeGrid:
		ok = stage.IsStagedMidArcVectorShapeGrid(target)

	case *PerpendicularVector:
		ok = stage.IsStagedPerpendicularVector(target)

	case *PerpendicularVectorGrid:
		ok = stage.IsStagedPerpendicularVectorGrid(target)

	case *PerpendicularVectorGridHalfway:
		ok = stage.IsStagedPerpendicularVectorGridHalfway(target)

	case *PerpendicularVectorHalfway:
		ok = stage.IsStagedPerpendicularVectorHalfway(target)

	case *Plant:
		ok = stage.IsStagedPlant(target)

	case *PlantCircumferenceShape:
		ok = stage.IsStagedPlantCircumferenceShape(target)

	case *PlantDiagram:
		ok = stage.IsStagedPlantDiagram(target)

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

	case *ShiftedBottomTopStartArcShape:
		ok = stage.IsStagedShiftedBottomTopStartArcShape(target)

	case *ShiftedBottomTopStartArcShapeGrid:
		ok = stage.IsStagedShiftedBottomTopStartArcShapeGrid(target)

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

	case *StackOfGrowthCurve2DRibbon:
		ok = stage.IsStagedStackOfGrowthCurve2DRibbon(target)

	case *StackOfRotatedGrowthCurve2D:
		ok = stage.IsStagedStackOfRotatedGrowthCurve2D(target)

	case *StackRotatedGrowthCurve2DEndArcShape:
		ok = stage.IsStagedStackRotatedGrowthCurve2DEndArcShape(target)

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

	case *TorusStackShape:
		ok = stage.IsStagedTorusStackShape(target)

	case *VerticalTorusStackShape:
		ok = stage.IsStagedVerticalTorusStackShape(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
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

func (stage *Stage) IsStagedCircleGridShape(circlegridshape *CircleGridShape) (ok bool) {

	_, ok = stage.CircleGridShapes[circlegridshape]

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

func (stage *Stage) IsStagedGridPathShape(gridpathshape *GridPathShape) (ok bool) {

	_, ok = stage.GridPathShapes[gridpathshape]

	return
}

func (stage *Stage) IsStagedGrowthCurve2D(growthcurve2d *GrowthCurve2D) (ok bool) {

	_, ok = stage.GrowthCurve2Ds[growthcurve2d]

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

func (stage *Stage) IsStagedPlant(plant *Plant) (ok bool) {

	_, ok = stage.Plants[plant]

	return
}

func (stage *Stage) IsStagedPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape) (ok bool) {

	_, ok = stage.PlantCircumferenceShapes[plantcircumferenceshape]

	return
}

func (stage *Stage) IsStagedPlantDiagram(plantdiagram *PlantDiagram) (ok bool) {

	_, ok = stage.PlantDiagrams[plantdiagram]

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

func (stage *Stage) IsStagedShiftedBottomTopStartArcShape(shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) (ok bool) {

	_, ok = stage.ShiftedBottomTopStartArcShapes[shiftedbottomtopstartarcshape]

	return
}

func (stage *Stage) IsStagedShiftedBottomTopStartArcShapeGrid(shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) (ok bool) {

	_, ok = stage.ShiftedBottomTopStartArcShapeGrids[shiftedbottomtopstartarcshapegrid]

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

func (stage *Stage) IsStagedStackOfGrowthCurve2DRibbon(stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) (ok bool) {

	_, ok = stage.StackOfGrowthCurve2DRibbons[stackofgrowthcurve2dribbon]

	return
}

func (stage *Stage) IsStagedStackOfRotatedGrowthCurve2D(stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) (ok bool) {

	_, ok = stage.StackOfRotatedGrowthCurve2Ds[stackofrotatedgrowthcurve2d]

	return
}

func (stage *Stage) IsStagedStackRotatedGrowthCurve2DEndArcShape(stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) (ok bool) {

	_, ok = stage.StackRotatedGrowthCurve2DEndArcShapes[stackrotatedgrowthcurve2dendarcshape]

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

func (stage *Stage) IsStagedTorusStackShape(torusstackshape *TorusStackShape) (ok bool) {

	_, ok = stage.TorusStackShapes[torusstackshape]

	return
}

func (stage *Stage) IsStagedVerticalTorusStackShape(verticaltorusstackshape *VerticalTorusStackShape) (ok bool) {

	_, ok = stage.VerticalTorusStackShapes[verticaltorusstackshape]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
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

	case *CircleGridShape:
		stage.StageBranchCircleGridShape(target)

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

	case *GridPathShape:
		stage.StageBranchGridPathShape(target)

	case *GrowthCurve2D:
		stage.StageBranchGrowthCurve2D(target)

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

	case *Library:
		stage.StageBranchLibrary(target)

	case *MidArcVectorShape:
		stage.StageBranchMidArcVectorShape(target)

	case *MidArcVectorShapeGrid:
		stage.StageBranchMidArcVectorShapeGrid(target)

	case *PerpendicularVector:
		stage.StageBranchPerpendicularVector(target)

	case *PerpendicularVectorGrid:
		stage.StageBranchPerpendicularVectorGrid(target)

	case *PerpendicularVectorGridHalfway:
		stage.StageBranchPerpendicularVectorGridHalfway(target)

	case *PerpendicularVectorHalfway:
		stage.StageBranchPerpendicularVectorHalfway(target)

	case *Plant:
		stage.StageBranchPlant(target)

	case *PlantCircumferenceShape:
		stage.StageBranchPlantCircumferenceShape(target)

	case *PlantDiagram:
		stage.StageBranchPlantDiagram(target)

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

	case *ShiftedBottomTopStartArcShape:
		stage.StageBranchShiftedBottomTopStartArcShape(target)

	case *ShiftedBottomTopStartArcShapeGrid:
		stage.StageBranchShiftedBottomTopStartArcShapeGrid(target)

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

	case *StackOfGrowthCurve2DRibbon:
		stage.StageBranchStackOfGrowthCurve2DRibbon(target)

	case *StackOfRotatedGrowthCurve2D:
		stage.StageBranchStackOfRotatedGrowthCurve2D(target)

	case *StackRotatedGrowthCurve2DEndArcShape:
		stage.StageBranchStackRotatedGrowthCurve2DEndArcShape(target)

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

	case *TorusStackShape:
		stage.StageBranchTorusStackShape(target)

	case *VerticalTorusStackShape:
		stage.StageBranchVerticalTorusStackShape(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
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
	for _, _arcnormalvectorshape := range arcnormalvectorshapegrid.ArcNormalVectorShapes {
		StageBranch(stage, _arcnormalvectorshape)
	}

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
	for _, _basevectorshape := range basevectorshapegrid.BaseVectorShapes {
		StageBranch(stage, _basevectorshape)
	}

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
	for _, _endarcshape := range endarcshapegrid.EndArcShapes {
		StageBranch(stage, _endarcshape)
	}

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
	for _, _endhalfwayarcshape := range endhalfwayarcshapegrid.EndHalfwayArcShapes {
		StageBranch(stage, _endhalfwayarcshape)
	}

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
	if growthcurve2d.StartHalfwayArcShapeGrid != nil {
		StageBranch(stage, growthcurve2d.StartHalfwayArcShapeGrid)
	}
	if growthcurve2d.EndHalfwayArcShapeGrid != nil {
		StageBranch(stage, growthcurve2d.EndHalfwayArcShapeGrid)
	}

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
	for _, _growthcurverhombusshape := range growthcurverhombusgridshape.GrowthCurveRhombusShapes {
		StageBranch(stage, _growthcurverhombusshape)
	}

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
	for _, _initialrhombusshape := range initialrhombusgridshape.InitialRhombusShapes {
		StageBranch(stage, _initialrhombusshape)
	}

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
	for _, _plant := range library.Plants {
		StageBranch(stage, _plant)
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
	for _, _midarcvectorshape := range midarcvectorshapegrid.MidArcVectorShapes {
		StageBranch(stage, _midarcvectorshape)
	}

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
	for _, _perpendicularvector := range perpendicularvectorgrid.PerpendicularVectors {
		StageBranch(stage, _perpendicularvector)
	}

}

func (stage *Stage) StageBranchPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) {

	// check if instance is already staged
	if IsStaged(stage, perpendicularvectorgridhalfway) {
		return
	}

	perpendicularvectorgridhalfway.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _perpendicularvectorhalfway := range perpendicularvectorgridhalfway.PerpendicularVectorHalfways {
		StageBranch(stage, _perpendicularvectorhalfway)
	}

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

func (stage *Stage) StageBranchPlant(plant *Plant) {

	// check if instance is already staged
	if IsStaged(stage, plant) {
		return
	}

	plant.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plant.AxesShape != nil {
		StageBranch(stage, plant.AxesShape)
	}
	if plant.RhombusStuff != nil {
		StageBranch(stage, plant.RhombusStuff)
	}
	if plant.GrowthVectorShape != nil {
		StageBranch(stage, plant.GrowthVectorShape)
	}
	if plant.PerpendicularVectorGrid != nil {
		StageBranch(stage, plant.PerpendicularVectorGrid)
	}
	if plant.PerpendicularVectorGridHalfway != nil {
		StageBranch(stage, plant.PerpendicularVectorGridHalfway)
	}
	if plant.BaseVectorShapeGrid != nil {
		StageBranch(stage, plant.BaseVectorShapeGrid)
	}
	if plant.ArcNormalVectorShapeGrid != nil {
		StageBranch(stage, plant.ArcNormalVectorShapeGrid)
	}
	if plant.StartArcShapeGrid != nil {
		StageBranch(stage, plant.StartArcShapeGrid)
	}
	if plant.TopStartArcShapeGrid != nil {
		StageBranch(stage, plant.TopStartArcShapeGrid)
	}
	if plant.EndArcShapeGrid != nil {
		StageBranch(stage, plant.EndArcShapeGrid)
	}
	if plant.TopEndArcShapeGrid != nil {
		StageBranch(stage, plant.TopEndArcShapeGrid)
	}
	if plant.ShiftedBottomTopStartArcShapeGrid != nil {
		StageBranch(stage, plant.ShiftedBottomTopStartArcShapeGrid)
	}
	if plant.MidArcVectorShapeGrid != nil {
		StageBranch(stage, plant.MidArcVectorShapeGrid)
	}
	if plant.TopMidArcVectorShapeGrid != nil {
		StageBranch(stage, plant.TopMidArcVectorShapeGrid)
	}
	if plant.StartHalfwayArcShapeGrid != nil {
		StageBranch(stage, plant.StartHalfwayArcShapeGrid)
	}
	if plant.TopStartHalfwayArcShapeGrid != nil {
		StageBranch(stage, plant.TopStartHalfwayArcShapeGrid)
	}
	if plant.EndHalfwayArcShapeGrid != nil {
		StageBranch(stage, plant.EndHalfwayArcShapeGrid)
	}
	if plant.TopEndHalfwayArcShapeGrid != nil {
		StageBranch(stage, plant.TopEndHalfwayArcShapeGrid)
	}
	if plant.StackOfRotatedGrowthCurve2D != nil {
		StageBranch(stage, plant.StackOfRotatedGrowthCurve2D)
	}
	if plant.TopStackOfRotatedGrowthCurve2D != nil {
		StageBranch(stage, plant.TopStackOfRotatedGrowthCurve2D)
	}
	if plant.GrowthCurve2D != nil {
		StageBranch(stage, plant.GrowthCurve2D)
	}
	if plant.TopGrowthCurve2D != nil {
		StageBranch(stage, plant.TopGrowthCurve2D)
	}
	if plant.StackOfGrowthCurve2D != nil {
		StageBranch(stage, plant.StackOfGrowthCurve2D)
	}
	if plant.TopStackOfGrowthCurve2D != nil {
		StageBranch(stage, plant.TopStackOfGrowthCurve2D)
	}
	if plant.StackOfGrowthCurve2DRibbon != nil {
		StageBranch(stage, plant.StackOfGrowthCurve2DRibbon)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plantdiagram := range plant.PlantDiagrams {
		StageBranch(stage, _plantdiagram)
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

func (stage *Stage) StageBranchPlantDiagram(plantdiagram *PlantDiagram) {

	// check if instance is already staged
	if IsStaged(stage, plantdiagram) {
		return
	}

	plantdiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plantdiagram.Rendered3DShape != nil {
		StageBranch(stage, plantdiagram.Rendered3DShape)
	}
	if plantdiagram.TorusStackShape != nil {
		StageBranch(stage, plantdiagram.TorusStackShape)
	}
	if plantdiagram.VerticalTorusStackShape != nil {
		StageBranch(stage, plantdiagram.VerticalTorusStackShape)
	}

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
	if rhombusstuff.ReferenceRhombus != nil {
		StageBranch(stage, rhombusstuff.ReferenceRhombus)
	}
	if rhombusstuff.PlantCircumferenceShape != nil {
		StageBranch(stage, rhombusstuff.PlantCircumferenceShape)
	}
	if rhombusstuff.GridPathShape != nil {
		StageBranch(stage, rhombusstuff.GridPathShape)
	}
	if rhombusstuff.InitialRhombusGridShape != nil {
		StageBranch(stage, rhombusstuff.InitialRhombusGridShape)
	}
	if rhombusstuff.ExplanationTextShape != nil {
		StageBranch(stage, rhombusstuff.ExplanationTextShape)
	}
	if rhombusstuff.RotatedReferenceRhombus != nil {
		StageBranch(stage, rhombusstuff.RotatedReferenceRhombus)
	}
	if rhombusstuff.RotatedPlantCircumferenceShape != nil {
		StageBranch(stage, rhombusstuff.RotatedPlantCircumferenceShape)
	}
	if rhombusstuff.RotatedGridPathShape != nil {
		StageBranch(stage, rhombusstuff.RotatedGridPathShape)
	}
	if rhombusstuff.RotatedRhombusGridShape2 != nil {
		StageBranch(stage, rhombusstuff.RotatedRhombusGridShape2)
	}
	if rhombusstuff.GrowthCurveRhombusGridShape != nil {
		StageBranch(stage, rhombusstuff.GrowthCurveRhombusGridShape)
	}

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
	for _, _rotatedrhombusshape := range rotatedrhombusgridshape.RotatedRhombusShapes {
		StageBranch(stage, _rotatedrhombusshape)
	}

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
	for _, _shiftedbottomtopstartarcshape := range shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes {
		StageBranch(stage, _shiftedbottomtopstartarcshape)
	}

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
	for _, _shiftedleftstackgrowthcurvestartarcshape := range shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes {
		StageBranch(stage, _shiftedleftstackgrowthcurvestartarcshape)
	}
	for _, _shiftedleftstackgrowthcurveendarcshape := range shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes {
		StageBranch(stage, _shiftedleftstackgrowthcurveendarcshape)
	}

}

func (stage *Stage) StageBranchShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) {

	// check if instance is already staged
	if IsStaged(stage, shiftedleftstackofnormalvector) {
		return
	}

	shiftedleftstackofnormalvector.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _shiftedleftstacknormalvector := range shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors {
		StageBranch(stage, _shiftedleftstacknormalvector)
	}

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
	for _, _stackgrowthcurve2dstarthalfwayarcshape := range stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes {
		StageBranch(stage, _stackgrowthcurve2dstarthalfwayarcshape)
	}
	for _, _stackgrowthcurve2dendhalfwayarcshape := range stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes {
		StageBranch(stage, _stackgrowthcurve2dendhalfwayarcshape)
	}

}

func (stage *Stage) StageBranchStackOfGrowthCurve2DRibbon(stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) {

	// check if instance is already staged
	if IsStaged(stage, stackofgrowthcurve2dribbon) {
		return
	}

	stackofgrowthcurve2dribbon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stackgrowthcurve2dribbonstartshape := range stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes {
		StageBranch(stage, _stackgrowthcurve2dribbonstartshape)
	}
	for _, _stackgrowthcurve2dribbonendshape := range stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes {
		StageBranch(stage, _stackgrowthcurve2dribbonendshape)
	}

}

func (stage *Stage) StageBranchStackOfRotatedGrowthCurve2D(stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) {

	// check if instance is already staged
	if IsStaged(stage, stackofrotatedgrowthcurve2d) {
		return
	}

	stackofrotatedgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stackrotatedgrowthcurve2dstartarcshape := range stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes {
		StageBranch(stage, _stackrotatedgrowthcurve2dstartarcshape)
	}
	for _, _stackrotatedgrowthcurve2dendarcshape := range stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes {
		StageBranch(stage, _stackrotatedgrowthcurve2dendarcshape)
	}

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
	for _, _startarcshape := range startarcshapegrid.StartArcShapes {
		StageBranch(stage, _startarcshape)
	}

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
	for _, _starthalfwayarcshape := range starthalfwayarcshapegrid.StartHalfwayArcShapes {
		StageBranch(stage, _starthalfwayarcshape)
	}

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
	for _, _topendarcshape := range topendarcshapegrid.TopEndArcShapes {
		StageBranch(stage, _topendarcshape)
	}

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
	for _, _topendhalfwayarcshape := range topendhalfwayarcshapegrid.TopEndHalfwayArcShapes {
		StageBranch(stage, _topendhalfwayarcshape)
	}

}

func (stage *Stage) StageBranchTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D) {

	// check if instance is already staged
	if IsStaged(stage, topgrowthcurve2d) {
		return
	}

	topgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if topgrowthcurve2d.TopStartHalfwayArcShapeGrid != nil {
		StageBranch(stage, topgrowthcurve2d.TopStartHalfwayArcShapeGrid)
	}
	if topgrowthcurve2d.TopEndHalfwayArcShapeGrid != nil {
		StageBranch(stage, topgrowthcurve2d.TopEndHalfwayArcShapeGrid)
	}

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
	for _, _topmidarcvectorshape := range topmidarcvectorshapegrid.TopMidArcVectorShapes {
		StageBranch(stage, _topmidarcvectorshape)
	}

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
	for _, _topstackgrowthcurve2dstarthalfwayarcshape := range topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes {
		StageBranch(stage, _topstackgrowthcurve2dstarthalfwayarcshape)
	}
	for _, _topstackgrowthcurve2dendhalfwayarcshape := range topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes {
		StageBranch(stage, _topstackgrowthcurve2dendhalfwayarcshape)
	}

}

func (stage *Stage) StageBranchTopStackOfRotatedGrowthCurve2D(topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) {

	// check if instance is already staged
	if IsStaged(stage, topstackofrotatedgrowthcurve2d) {
		return
	}

	topstackofrotatedgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topstackofrotatedgrowthcurve2dstartarcshape := range topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes {
		StageBranch(stage, _topstackofrotatedgrowthcurve2dstartarcshape)
	}
	for _, _topstackofrotatedgrowthcurve2dendarcshape := range topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes {
		StageBranch(stage, _topstackofrotatedgrowthcurve2dendarcshape)
	}

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
	for _, _topstartarcshape := range topstartarcshapegrid.TopStartArcShapes {
		StageBranch(stage, _topstartarcshape)
	}

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
	for _, _topstarthalfwayarcshape := range topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes {
		StageBranch(stage, _topstarthalfwayarcshape)
	}

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

func (stage *Stage) StageBranchVerticalTorusStackShape(verticaltorusstackshape *VerticalTorusStackShape) {

	// check if instance is already staged
	if IsStaged(stage, verticaltorusstackshape) {
		return
	}

	verticaltorusstackshape.Stage(stage)

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

	case *CircleGridShape:
		toT := CopyBranchCircleGridShape(mapOrigCopy, fromT)
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

	case *GridPathShape:
		toT := CopyBranchGridPathShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurve2D:
		toT := CopyBranchGrowthCurve2D(mapOrigCopy, fromT)
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

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MidArcVectorShape:
		toT := CopyBranchMidArcVectorShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MidArcVectorShapeGrid:
		toT := CopyBranchMidArcVectorShapeGrid(mapOrigCopy, fromT)
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

	case *Plant:
		toT := CopyBranchPlant(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PlantCircumferenceShape:
		toT := CopyBranchPlantCircumferenceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PlantDiagram:
		toT := CopyBranchPlantDiagram(mapOrigCopy, fromT)
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

	case *ShiftedBottomTopStartArcShape:
		toT := CopyBranchShiftedBottomTopStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ShiftedBottomTopStartArcShapeGrid:
		toT := CopyBranchShiftedBottomTopStartArcShapeGrid(mapOrigCopy, fromT)
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

	case *StackOfGrowthCurve2DRibbon:
		toT := CopyBranchStackOfGrowthCurve2DRibbon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfRotatedGrowthCurve2D:
		toT := CopyBranchStackOfRotatedGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackRotatedGrowthCurve2DEndArcShape:
		toT := CopyBranchStackRotatedGrowthCurve2DEndArcShape(mapOrigCopy, fromT)
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

	case *TorusStackShape:
		toT := CopyBranchTorusStackShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *VerticalTorusStackShape:
		toT := CopyBranchVerticalTorusStackShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
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
	for _, _arcnormalvectorshape := range arcnormalvectorshapegridFrom.ArcNormalVectorShapes {
		arcnormalvectorshapegridTo.ArcNormalVectorShapes = append(arcnormalvectorshapegridTo.ArcNormalVectorShapes, CopyBranchArcNormalVectorShape(mapOrigCopy, _arcnormalvectorshape))
	}

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
	for _, _basevectorshape := range basevectorshapegridFrom.BaseVectorShapes {
		basevectorshapegridTo.BaseVectorShapes = append(basevectorshapegridTo.BaseVectorShapes, CopyBranchBaseVectorShape(mapOrigCopy, _basevectorshape))
	}

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
	for _, _endarcshape := range endarcshapegridFrom.EndArcShapes {
		endarcshapegridTo.EndArcShapes = append(endarcshapegridTo.EndArcShapes, CopyBranchEndArcShape(mapOrigCopy, _endarcshape))
	}

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
	for _, _endhalfwayarcshape := range endhalfwayarcshapegridFrom.EndHalfwayArcShapes {
		endhalfwayarcshapegridTo.EndHalfwayArcShapes = append(endhalfwayarcshapegridTo.EndHalfwayArcShapes, CopyBranchEndHalfwayArcShape(mapOrigCopy, _endhalfwayarcshape))
	}

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
	if growthcurve2dFrom.StartHalfwayArcShapeGrid != nil {
		growthcurve2dTo.StartHalfwayArcShapeGrid = CopyBranchStartHalfwayArcShapeGrid(mapOrigCopy, growthcurve2dFrom.StartHalfwayArcShapeGrid)
	}
	if growthcurve2dFrom.EndHalfwayArcShapeGrid != nil {
		growthcurve2dTo.EndHalfwayArcShapeGrid = CopyBranchEndHalfwayArcShapeGrid(mapOrigCopy, growthcurve2dFrom.EndHalfwayArcShapeGrid)
	}

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
	for _, _growthcurverhombusshape := range growthcurverhombusgridshapeFrom.GrowthCurveRhombusShapes {
		growthcurverhombusgridshapeTo.GrowthCurveRhombusShapes = append(growthcurverhombusgridshapeTo.GrowthCurveRhombusShapes, CopyBranchGrowthCurveRhombusShape(mapOrigCopy, _growthcurverhombusshape))
	}

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
	for _, _initialrhombusshape := range initialrhombusgridshapeFrom.InitialRhombusShapes {
		initialrhombusgridshapeTo.InitialRhombusShapes = append(initialrhombusgridshapeTo.InitialRhombusShapes, CopyBranchInitialRhombusShape(mapOrigCopy, _initialrhombusshape))
	}

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
	for _, _plant := range libraryFrom.Plants {
		libraryTo.Plants = append(libraryTo.Plants, CopyBranchPlant(mapOrigCopy, _plant))
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
	for _, _midarcvectorshape := range midarcvectorshapegridFrom.MidArcVectorShapes {
		midarcvectorshapegridTo.MidArcVectorShapes = append(midarcvectorshapegridTo.MidArcVectorShapes, CopyBranchMidArcVectorShape(mapOrigCopy, _midarcvectorshape))
	}

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
	for _, _perpendicularvector := range perpendicularvectorgridFrom.PerpendicularVectors {
		perpendicularvectorgridTo.PerpendicularVectors = append(perpendicularvectorgridTo.PerpendicularVectors, CopyBranchPerpendicularVector(mapOrigCopy, _perpendicularvector))
	}

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
	for _, _perpendicularvectorhalfway := range perpendicularvectorgridhalfwayFrom.PerpendicularVectorHalfways {
		perpendicularvectorgridhalfwayTo.PerpendicularVectorHalfways = append(perpendicularvectorgridhalfwayTo.PerpendicularVectorHalfways, CopyBranchPerpendicularVectorHalfway(mapOrigCopy, _perpendicularvectorhalfway))
	}

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

func CopyBranchPlant(mapOrigCopy map[any]any, plantFrom *Plant) (plantTo *Plant) {

	// plantFrom has already been copied
	if _plantTo, ok := mapOrigCopy[plantFrom]; ok {
		plantTo = _plantTo.(*Plant)
		return
	}

	plantTo = new(Plant)
	mapOrigCopy[plantFrom] = plantTo
	plantFrom.CopyBasicFields(plantTo)

	//insertion point for the staging of instances referenced by pointers
	if plantFrom.AxesShape != nil {
		plantTo.AxesShape = CopyBranchAxesShape(mapOrigCopy, plantFrom.AxesShape)
	}
	if plantFrom.RhombusStuff != nil {
		plantTo.RhombusStuff = CopyBranchRhombusStuff(mapOrigCopy, plantFrom.RhombusStuff)
	}
	if plantFrom.GrowthVectorShape != nil {
		plantTo.GrowthVectorShape = CopyBranchGrowthVectorShape(mapOrigCopy, plantFrom.GrowthVectorShape)
	}
	if plantFrom.PerpendicularVectorGrid != nil {
		plantTo.PerpendicularVectorGrid = CopyBranchPerpendicularVectorGrid(mapOrigCopy, plantFrom.PerpendicularVectorGrid)
	}
	if plantFrom.PerpendicularVectorGridHalfway != nil {
		plantTo.PerpendicularVectorGridHalfway = CopyBranchPerpendicularVectorGridHalfway(mapOrigCopy, plantFrom.PerpendicularVectorGridHalfway)
	}
	if plantFrom.BaseVectorShapeGrid != nil {
		plantTo.BaseVectorShapeGrid = CopyBranchBaseVectorShapeGrid(mapOrigCopy, plantFrom.BaseVectorShapeGrid)
	}
	if plantFrom.ArcNormalVectorShapeGrid != nil {
		plantTo.ArcNormalVectorShapeGrid = CopyBranchArcNormalVectorShapeGrid(mapOrigCopy, plantFrom.ArcNormalVectorShapeGrid)
	}
	if plantFrom.StartArcShapeGrid != nil {
		plantTo.StartArcShapeGrid = CopyBranchStartArcShapeGrid(mapOrigCopy, plantFrom.StartArcShapeGrid)
	}
	if plantFrom.TopStartArcShapeGrid != nil {
		plantTo.TopStartArcShapeGrid = CopyBranchTopStartArcShapeGrid(mapOrigCopy, plantFrom.TopStartArcShapeGrid)
	}
	if plantFrom.EndArcShapeGrid != nil {
		plantTo.EndArcShapeGrid = CopyBranchEndArcShapeGrid(mapOrigCopy, plantFrom.EndArcShapeGrid)
	}
	if plantFrom.TopEndArcShapeGrid != nil {
		plantTo.TopEndArcShapeGrid = CopyBranchTopEndArcShapeGrid(mapOrigCopy, plantFrom.TopEndArcShapeGrid)
	}
	if plantFrom.ShiftedBottomTopStartArcShapeGrid != nil {
		plantTo.ShiftedBottomTopStartArcShapeGrid = CopyBranchShiftedBottomTopStartArcShapeGrid(mapOrigCopy, plantFrom.ShiftedBottomTopStartArcShapeGrid)
	}
	if plantFrom.MidArcVectorShapeGrid != nil {
		plantTo.MidArcVectorShapeGrid = CopyBranchMidArcVectorShapeGrid(mapOrigCopy, plantFrom.MidArcVectorShapeGrid)
	}
	if plantFrom.TopMidArcVectorShapeGrid != nil {
		plantTo.TopMidArcVectorShapeGrid = CopyBranchTopMidArcVectorShapeGrid(mapOrigCopy, plantFrom.TopMidArcVectorShapeGrid)
	}
	if plantFrom.StartHalfwayArcShapeGrid != nil {
		plantTo.StartHalfwayArcShapeGrid = CopyBranchStartHalfwayArcShapeGrid(mapOrigCopy, plantFrom.StartHalfwayArcShapeGrid)
	}
	if plantFrom.TopStartHalfwayArcShapeGrid != nil {
		plantTo.TopStartHalfwayArcShapeGrid = CopyBranchTopStartHalfwayArcShapeGrid(mapOrigCopy, plantFrom.TopStartHalfwayArcShapeGrid)
	}
	if plantFrom.EndHalfwayArcShapeGrid != nil {
		plantTo.EndHalfwayArcShapeGrid = CopyBranchEndHalfwayArcShapeGrid(mapOrigCopy, plantFrom.EndHalfwayArcShapeGrid)
	}
	if plantFrom.TopEndHalfwayArcShapeGrid != nil {
		plantTo.TopEndHalfwayArcShapeGrid = CopyBranchTopEndHalfwayArcShapeGrid(mapOrigCopy, plantFrom.TopEndHalfwayArcShapeGrid)
	}
	if plantFrom.StackOfRotatedGrowthCurve2D != nil {
		plantTo.StackOfRotatedGrowthCurve2D = CopyBranchStackOfRotatedGrowthCurve2D(mapOrigCopy, plantFrom.StackOfRotatedGrowthCurve2D)
	}
	if plantFrom.TopStackOfRotatedGrowthCurve2D != nil {
		plantTo.TopStackOfRotatedGrowthCurve2D = CopyBranchTopStackOfRotatedGrowthCurve2D(mapOrigCopy, plantFrom.TopStackOfRotatedGrowthCurve2D)
	}
	if plantFrom.GrowthCurve2D != nil {
		plantTo.GrowthCurve2D = CopyBranchGrowthCurve2D(mapOrigCopy, plantFrom.GrowthCurve2D)
	}
	if plantFrom.TopGrowthCurve2D != nil {
		plantTo.TopGrowthCurve2D = CopyBranchTopGrowthCurve2D(mapOrigCopy, plantFrom.TopGrowthCurve2D)
	}
	if plantFrom.StackOfGrowthCurve2D != nil {
		plantTo.StackOfGrowthCurve2D = CopyBranchStackOfGrowthCurve2D(mapOrigCopy, plantFrom.StackOfGrowthCurve2D)
	}
	if plantFrom.TopStackOfGrowthCurve2D != nil {
		plantTo.TopStackOfGrowthCurve2D = CopyBranchTopStackOfGrowthCurve2D(mapOrigCopy, plantFrom.TopStackOfGrowthCurve2D)
	}
	if plantFrom.StackOfGrowthCurve2DRibbon != nil {
		plantTo.StackOfGrowthCurve2DRibbon = CopyBranchStackOfGrowthCurve2DRibbon(mapOrigCopy, plantFrom.StackOfGrowthCurve2DRibbon)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plantdiagram := range plantFrom.PlantDiagrams {
		plantTo.PlantDiagrams = append(plantTo.PlantDiagrams, CopyBranchPlantDiagram(mapOrigCopy, _plantdiagram))
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

func CopyBranchPlantDiagram(mapOrigCopy map[any]any, plantdiagramFrom *PlantDiagram) (plantdiagramTo *PlantDiagram) {

	// plantdiagramFrom has already been copied
	if _plantdiagramTo, ok := mapOrigCopy[plantdiagramFrom]; ok {
		plantdiagramTo = _plantdiagramTo.(*PlantDiagram)
		return
	}

	plantdiagramTo = new(PlantDiagram)
	mapOrigCopy[plantdiagramFrom] = plantdiagramTo
	plantdiagramFrom.CopyBasicFields(plantdiagramTo)

	//insertion point for the staging of instances referenced by pointers
	if plantdiagramFrom.Rendered3DShape != nil {
		plantdiagramTo.Rendered3DShape = CopyBranchRendered3DShape(mapOrigCopy, plantdiagramFrom.Rendered3DShape)
	}
	if plantdiagramFrom.TorusStackShape != nil {
		plantdiagramTo.TorusStackShape = CopyBranchTorusStackShape(mapOrigCopy, plantdiagramFrom.TorusStackShape)
	}
	if plantdiagramFrom.VerticalTorusStackShape != nil {
		plantdiagramTo.VerticalTorusStackShape = CopyBranchVerticalTorusStackShape(mapOrigCopy, plantdiagramFrom.VerticalTorusStackShape)
	}

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
	if rhombusstuffFrom.ReferenceRhombus != nil {
		rhombusstuffTo.ReferenceRhombus = CopyBranchRhombusShape(mapOrigCopy, rhombusstuffFrom.ReferenceRhombus)
	}
	if rhombusstuffFrom.PlantCircumferenceShape != nil {
		rhombusstuffTo.PlantCircumferenceShape = CopyBranchPlantCircumferenceShape(mapOrigCopy, rhombusstuffFrom.PlantCircumferenceShape)
	}
	if rhombusstuffFrom.GridPathShape != nil {
		rhombusstuffTo.GridPathShape = CopyBranchGridPathShape(mapOrigCopy, rhombusstuffFrom.GridPathShape)
	}
	if rhombusstuffFrom.InitialRhombusGridShape != nil {
		rhombusstuffTo.InitialRhombusGridShape = CopyBranchInitialRhombusGridShape(mapOrigCopy, rhombusstuffFrom.InitialRhombusGridShape)
	}
	if rhombusstuffFrom.ExplanationTextShape != nil {
		rhombusstuffTo.ExplanationTextShape = CopyBranchExplanationTextShape(mapOrigCopy, rhombusstuffFrom.ExplanationTextShape)
	}
	if rhombusstuffFrom.RotatedReferenceRhombus != nil {
		rhombusstuffTo.RotatedReferenceRhombus = CopyBranchRhombusShape(mapOrigCopy, rhombusstuffFrom.RotatedReferenceRhombus)
	}
	if rhombusstuffFrom.RotatedPlantCircumferenceShape != nil {
		rhombusstuffTo.RotatedPlantCircumferenceShape = CopyBranchPlantCircumferenceShape(mapOrigCopy, rhombusstuffFrom.RotatedPlantCircumferenceShape)
	}
	if rhombusstuffFrom.RotatedGridPathShape != nil {
		rhombusstuffTo.RotatedGridPathShape = CopyBranchGridPathShape(mapOrigCopy, rhombusstuffFrom.RotatedGridPathShape)
	}
	if rhombusstuffFrom.RotatedRhombusGridShape2 != nil {
		rhombusstuffTo.RotatedRhombusGridShape2 = CopyBranchRotatedRhombusGridShape(mapOrigCopy, rhombusstuffFrom.RotatedRhombusGridShape2)
	}
	if rhombusstuffFrom.GrowthCurveRhombusGridShape != nil {
		rhombusstuffTo.GrowthCurveRhombusGridShape = CopyBranchGrowthCurveRhombusGridShape(mapOrigCopy, rhombusstuffFrom.GrowthCurveRhombusGridShape)
	}

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
	for _, _rotatedrhombusshape := range rotatedrhombusgridshapeFrom.RotatedRhombusShapes {
		rotatedrhombusgridshapeTo.RotatedRhombusShapes = append(rotatedrhombusgridshapeTo.RotatedRhombusShapes, CopyBranchRotatedRhombusShape(mapOrigCopy, _rotatedrhombusshape))
	}

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
	for _, _shiftedbottomtopstartarcshape := range shiftedbottomtopstartarcshapegridFrom.ShiftedBottomTopStartArcShapes {
		shiftedbottomtopstartarcshapegridTo.ShiftedBottomTopStartArcShapes = append(shiftedbottomtopstartarcshapegridTo.ShiftedBottomTopStartArcShapes, CopyBranchShiftedBottomTopStartArcShape(mapOrigCopy, _shiftedbottomtopstartarcshape))
	}

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
	for _, _shiftedleftstackgrowthcurvestartarcshape := range shiftedleftstackofgrowthcurveFrom.ShiftedLeftStackGrowthCurveStartArcShapes {
		shiftedleftstackofgrowthcurveTo.ShiftedLeftStackGrowthCurveStartArcShapes = append(shiftedleftstackofgrowthcurveTo.ShiftedLeftStackGrowthCurveStartArcShapes, CopyBranchShiftedLeftStackGrowthCurveStartArcShape(mapOrigCopy, _shiftedleftstackgrowthcurvestartarcshape))
	}
	for _, _shiftedleftstackgrowthcurveendarcshape := range shiftedleftstackofgrowthcurveFrom.ShiftedLeftStackGrowthCurveEndArcShapes {
		shiftedleftstackofgrowthcurveTo.ShiftedLeftStackGrowthCurveEndArcShapes = append(shiftedleftstackofgrowthcurveTo.ShiftedLeftStackGrowthCurveEndArcShapes, CopyBranchShiftedLeftStackGrowthCurveEndArcShape(mapOrigCopy, _shiftedleftstackgrowthcurveendarcshape))
	}

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
	for _, _shiftedleftstacknormalvector := range shiftedleftstackofnormalvectorFrom.ShiftedLeftStackNormalVectors {
		shiftedleftstackofnormalvectorTo.ShiftedLeftStackNormalVectors = append(shiftedleftstackofnormalvectorTo.ShiftedLeftStackNormalVectors, CopyBranchShiftedLeftStackNormalVector(mapOrigCopy, _shiftedleftstacknormalvector))
	}

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
	for _, _stackgrowthcurve2dstarthalfwayarcshape := range stackofgrowthcurve2dFrom.StackGrowthCurve2DStartHalfwayArcShapes {
		stackofgrowthcurve2dTo.StackGrowthCurve2DStartHalfwayArcShapes = append(stackofgrowthcurve2dTo.StackGrowthCurve2DStartHalfwayArcShapes, CopyBranchStackGrowthCurve2DStartHalfwayArcShape(mapOrigCopy, _stackgrowthcurve2dstarthalfwayarcshape))
	}
	for _, _stackgrowthcurve2dendhalfwayarcshape := range stackofgrowthcurve2dFrom.StackGrowthCurve2DEndHalfwayArcShapes {
		stackofgrowthcurve2dTo.StackGrowthCurve2DEndHalfwayArcShapes = append(stackofgrowthcurve2dTo.StackGrowthCurve2DEndHalfwayArcShapes, CopyBranchStackGrowthCurve2DEndHalfwayArcShape(mapOrigCopy, _stackgrowthcurve2dendhalfwayarcshape))
	}

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
	for _, _stackgrowthcurve2dribbonstartshape := range stackofgrowthcurve2dribbonFrom.StackGrowthCurve2DRibbonStartShapes {
		stackofgrowthcurve2dribbonTo.StackGrowthCurve2DRibbonStartShapes = append(stackofgrowthcurve2dribbonTo.StackGrowthCurve2DRibbonStartShapes, CopyBranchStackGrowthCurve2DRibbonStartShape(mapOrigCopy, _stackgrowthcurve2dribbonstartshape))
	}
	for _, _stackgrowthcurve2dribbonendshape := range stackofgrowthcurve2dribbonFrom.StackGrowthCurve2DRibbonEndShapes {
		stackofgrowthcurve2dribbonTo.StackGrowthCurve2DRibbonEndShapes = append(stackofgrowthcurve2dribbonTo.StackGrowthCurve2DRibbonEndShapes, CopyBranchStackGrowthCurve2DRibbonEndShape(mapOrigCopy, _stackgrowthcurve2dribbonendshape))
	}

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
	for _, _stackrotatedgrowthcurve2dstartarcshape := range stackofrotatedgrowthcurve2dFrom.StackRotatedGrowthCurve2DStartArcShapes {
		stackofrotatedgrowthcurve2dTo.StackRotatedGrowthCurve2DStartArcShapes = append(stackofrotatedgrowthcurve2dTo.StackRotatedGrowthCurve2DStartArcShapes, CopyBranchStackRotatedGrowthCurve2DStartArcShape(mapOrigCopy, _stackrotatedgrowthcurve2dstartarcshape))
	}
	for _, _stackrotatedgrowthcurve2dendarcshape := range stackofrotatedgrowthcurve2dFrom.StackRotatedGrowthCurve2DEndArcShapes {
		stackofrotatedgrowthcurve2dTo.StackRotatedGrowthCurve2DEndArcShapes = append(stackofrotatedgrowthcurve2dTo.StackRotatedGrowthCurve2DEndArcShapes, CopyBranchStackRotatedGrowthCurve2DEndArcShape(mapOrigCopy, _stackrotatedgrowthcurve2dendarcshape))
	}

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
	for _, _startarcshape := range startarcshapegridFrom.StartArcShapes {
		startarcshapegridTo.StartArcShapes = append(startarcshapegridTo.StartArcShapes, CopyBranchStartArcShape(mapOrigCopy, _startarcshape))
	}

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
	for _, _starthalfwayarcshape := range starthalfwayarcshapegridFrom.StartHalfwayArcShapes {
		starthalfwayarcshapegridTo.StartHalfwayArcShapes = append(starthalfwayarcshapegridTo.StartHalfwayArcShapes, CopyBranchStartHalfwayArcShape(mapOrigCopy, _starthalfwayarcshape))
	}

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
	for _, _topendarcshape := range topendarcshapegridFrom.TopEndArcShapes {
		topendarcshapegridTo.TopEndArcShapes = append(topendarcshapegridTo.TopEndArcShapes, CopyBranchTopEndArcShape(mapOrigCopy, _topendarcshape))
	}

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
	for _, _topendhalfwayarcshape := range topendhalfwayarcshapegridFrom.TopEndHalfwayArcShapes {
		topendhalfwayarcshapegridTo.TopEndHalfwayArcShapes = append(topendhalfwayarcshapegridTo.TopEndHalfwayArcShapes, CopyBranchTopEndHalfwayArcShape(mapOrigCopy, _topendhalfwayarcshape))
	}

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
	if topgrowthcurve2dFrom.TopStartHalfwayArcShapeGrid != nil {
		topgrowthcurve2dTo.TopStartHalfwayArcShapeGrid = CopyBranchTopStartHalfwayArcShapeGrid(mapOrigCopy, topgrowthcurve2dFrom.TopStartHalfwayArcShapeGrid)
	}
	if topgrowthcurve2dFrom.TopEndHalfwayArcShapeGrid != nil {
		topgrowthcurve2dTo.TopEndHalfwayArcShapeGrid = CopyBranchTopEndHalfwayArcShapeGrid(mapOrigCopy, topgrowthcurve2dFrom.TopEndHalfwayArcShapeGrid)
	}

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
	for _, _topmidarcvectorshape := range topmidarcvectorshapegridFrom.TopMidArcVectorShapes {
		topmidarcvectorshapegridTo.TopMidArcVectorShapes = append(topmidarcvectorshapegridTo.TopMidArcVectorShapes, CopyBranchTopMidArcVectorShape(mapOrigCopy, _topmidarcvectorshape))
	}

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
	for _, _topstackgrowthcurve2dstarthalfwayarcshape := range topstackofgrowthcurve2dFrom.TopStackGrowthCurve2DStartHalfwayArcShapes {
		topstackofgrowthcurve2dTo.TopStackGrowthCurve2DStartHalfwayArcShapes = append(topstackofgrowthcurve2dTo.TopStackGrowthCurve2DStartHalfwayArcShapes, CopyBranchTopStackGrowthCurve2DStartHalfwayArcShape(mapOrigCopy, _topstackgrowthcurve2dstarthalfwayarcshape))
	}
	for _, _topstackgrowthcurve2dendhalfwayarcshape := range topstackofgrowthcurve2dFrom.TopStackGrowthCurve2DEndHalfwayArcShapes {
		topstackofgrowthcurve2dTo.TopStackGrowthCurve2DEndHalfwayArcShapes = append(topstackofgrowthcurve2dTo.TopStackGrowthCurve2DEndHalfwayArcShapes, CopyBranchTopStackGrowthCurve2DEndHalfwayArcShape(mapOrigCopy, _topstackgrowthcurve2dendhalfwayarcshape))
	}

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
	for _, _topstackofrotatedgrowthcurve2dstartarcshape := range topstackofrotatedgrowthcurve2dFrom.TopStackOfRotatedGrowthCurve2DStartArcShapes {
		topstackofrotatedgrowthcurve2dTo.TopStackOfRotatedGrowthCurve2DStartArcShapes = append(topstackofrotatedgrowthcurve2dTo.TopStackOfRotatedGrowthCurve2DStartArcShapes, CopyBranchTopStackOfRotatedGrowthCurve2DStartArcShape(mapOrigCopy, _topstackofrotatedgrowthcurve2dstartarcshape))
	}
	for _, _topstackofrotatedgrowthcurve2dendarcshape := range topstackofrotatedgrowthcurve2dFrom.TopStackOfRotatedGrowthCurve2DEndArcShapes {
		topstackofrotatedgrowthcurve2dTo.TopStackOfRotatedGrowthCurve2DEndArcShapes = append(topstackofrotatedgrowthcurve2dTo.TopStackOfRotatedGrowthCurve2DEndArcShapes, CopyBranchTopStackOfRotatedGrowthCurve2DEndArcShape(mapOrigCopy, _topstackofrotatedgrowthcurve2dendarcshape))
	}

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
	for _, _topstartarcshape := range topstartarcshapegridFrom.TopStartArcShapes {
		topstartarcshapegridTo.TopStartArcShapes = append(topstartarcshapegridTo.TopStartArcShapes, CopyBranchTopStartArcShape(mapOrigCopy, _topstartarcshape))
	}

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
	for _, _topstarthalfwayarcshape := range topstarthalfwayarcshapegridFrom.TopStartHalfwayArcShapes {
		topstarthalfwayarcshapegridTo.TopStartHalfwayArcShapes = append(topstarthalfwayarcshapegridTo.TopStartHalfwayArcShapes, CopyBranchTopStartHalfwayArcShape(mapOrigCopy, _topstarthalfwayarcshape))
	}

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

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
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

	case *CircleGridShape:
		stage.UnstageBranchCircleGridShape(target)

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

	case *GridPathShape:
		stage.UnstageBranchGridPathShape(target)

	case *GrowthCurve2D:
		stage.UnstageBranchGrowthCurve2D(target)

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

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *MidArcVectorShape:
		stage.UnstageBranchMidArcVectorShape(target)

	case *MidArcVectorShapeGrid:
		stage.UnstageBranchMidArcVectorShapeGrid(target)

	case *PerpendicularVector:
		stage.UnstageBranchPerpendicularVector(target)

	case *PerpendicularVectorGrid:
		stage.UnstageBranchPerpendicularVectorGrid(target)

	case *PerpendicularVectorGridHalfway:
		stage.UnstageBranchPerpendicularVectorGridHalfway(target)

	case *PerpendicularVectorHalfway:
		stage.UnstageBranchPerpendicularVectorHalfway(target)

	case *Plant:
		stage.UnstageBranchPlant(target)

	case *PlantCircumferenceShape:
		stage.UnstageBranchPlantCircumferenceShape(target)

	case *PlantDiagram:
		stage.UnstageBranchPlantDiagram(target)

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

	case *ShiftedBottomTopStartArcShape:
		stage.UnstageBranchShiftedBottomTopStartArcShape(target)

	case *ShiftedBottomTopStartArcShapeGrid:
		stage.UnstageBranchShiftedBottomTopStartArcShapeGrid(target)

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

	case *StackOfGrowthCurve2DRibbon:
		stage.UnstageBranchStackOfGrowthCurve2DRibbon(target)

	case *StackOfRotatedGrowthCurve2D:
		stage.UnstageBranchStackOfRotatedGrowthCurve2D(target)

	case *StackRotatedGrowthCurve2DEndArcShape:
		stage.UnstageBranchStackRotatedGrowthCurve2DEndArcShape(target)

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

	case *TorusStackShape:
		stage.UnstageBranchTorusStackShape(target)

	case *VerticalTorusStackShape:
		stage.UnstageBranchVerticalTorusStackShape(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
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
	for _, _arcnormalvectorshape := range arcnormalvectorshapegrid.ArcNormalVectorShapes {
		UnstageBranch(stage, _arcnormalvectorshape)
	}

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
	for _, _basevectorshape := range basevectorshapegrid.BaseVectorShapes {
		UnstageBranch(stage, _basevectorshape)
	}

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
	for _, _endarcshape := range endarcshapegrid.EndArcShapes {
		UnstageBranch(stage, _endarcshape)
	}

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
	for _, _endhalfwayarcshape := range endhalfwayarcshapegrid.EndHalfwayArcShapes {
		UnstageBranch(stage, _endhalfwayarcshape)
	}

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
	if growthcurve2d.StartHalfwayArcShapeGrid != nil {
		UnstageBranch(stage, growthcurve2d.StartHalfwayArcShapeGrid)
	}
	if growthcurve2d.EndHalfwayArcShapeGrid != nil {
		UnstageBranch(stage, growthcurve2d.EndHalfwayArcShapeGrid)
	}

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
	for _, _growthcurverhombusshape := range growthcurverhombusgridshape.GrowthCurveRhombusShapes {
		UnstageBranch(stage, _growthcurverhombusshape)
	}

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
	for _, _initialrhombusshape := range initialrhombusgridshape.InitialRhombusShapes {
		UnstageBranch(stage, _initialrhombusshape)
	}

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
	for _, _plant := range library.Plants {
		UnstageBranch(stage, _plant)
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
	for _, _midarcvectorshape := range midarcvectorshapegrid.MidArcVectorShapes {
		UnstageBranch(stage, _midarcvectorshape)
	}

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
	for _, _perpendicularvector := range perpendicularvectorgrid.PerpendicularVectors {
		UnstageBranch(stage, _perpendicularvector)
	}

}

func (stage *Stage) UnstageBranchPerpendicularVectorGridHalfway(perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) {

	// check if instance is already staged
	if !IsStaged(stage, perpendicularvectorgridhalfway) {
		return
	}

	perpendicularvectorgridhalfway.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _perpendicularvectorhalfway := range perpendicularvectorgridhalfway.PerpendicularVectorHalfways {
		UnstageBranch(stage, _perpendicularvectorhalfway)
	}

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

func (stage *Stage) UnstageBranchPlant(plant *Plant) {

	// check if instance is already staged
	if !IsStaged(stage, plant) {
		return
	}

	plant.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plant.AxesShape != nil {
		UnstageBranch(stage, plant.AxesShape)
	}
	if plant.RhombusStuff != nil {
		UnstageBranch(stage, plant.RhombusStuff)
	}
	if plant.GrowthVectorShape != nil {
		UnstageBranch(stage, plant.GrowthVectorShape)
	}
	if plant.PerpendicularVectorGrid != nil {
		UnstageBranch(stage, plant.PerpendicularVectorGrid)
	}
	if plant.PerpendicularVectorGridHalfway != nil {
		UnstageBranch(stage, plant.PerpendicularVectorGridHalfway)
	}
	if plant.BaseVectorShapeGrid != nil {
		UnstageBranch(stage, plant.BaseVectorShapeGrid)
	}
	if plant.ArcNormalVectorShapeGrid != nil {
		UnstageBranch(stage, plant.ArcNormalVectorShapeGrid)
	}
	if plant.StartArcShapeGrid != nil {
		UnstageBranch(stage, plant.StartArcShapeGrid)
	}
	if plant.TopStartArcShapeGrid != nil {
		UnstageBranch(stage, plant.TopStartArcShapeGrid)
	}
	if plant.EndArcShapeGrid != nil {
		UnstageBranch(stage, plant.EndArcShapeGrid)
	}
	if plant.TopEndArcShapeGrid != nil {
		UnstageBranch(stage, plant.TopEndArcShapeGrid)
	}
	if plant.ShiftedBottomTopStartArcShapeGrid != nil {
		UnstageBranch(stage, plant.ShiftedBottomTopStartArcShapeGrid)
	}
	if plant.MidArcVectorShapeGrid != nil {
		UnstageBranch(stage, plant.MidArcVectorShapeGrid)
	}
	if plant.TopMidArcVectorShapeGrid != nil {
		UnstageBranch(stage, plant.TopMidArcVectorShapeGrid)
	}
	if plant.StartHalfwayArcShapeGrid != nil {
		UnstageBranch(stage, plant.StartHalfwayArcShapeGrid)
	}
	if plant.TopStartHalfwayArcShapeGrid != nil {
		UnstageBranch(stage, plant.TopStartHalfwayArcShapeGrid)
	}
	if plant.EndHalfwayArcShapeGrid != nil {
		UnstageBranch(stage, plant.EndHalfwayArcShapeGrid)
	}
	if plant.TopEndHalfwayArcShapeGrid != nil {
		UnstageBranch(stage, plant.TopEndHalfwayArcShapeGrid)
	}
	if plant.StackOfRotatedGrowthCurve2D != nil {
		UnstageBranch(stage, plant.StackOfRotatedGrowthCurve2D)
	}
	if plant.TopStackOfRotatedGrowthCurve2D != nil {
		UnstageBranch(stage, plant.TopStackOfRotatedGrowthCurve2D)
	}
	if plant.GrowthCurve2D != nil {
		UnstageBranch(stage, plant.GrowthCurve2D)
	}
	if plant.TopGrowthCurve2D != nil {
		UnstageBranch(stage, plant.TopGrowthCurve2D)
	}
	if plant.StackOfGrowthCurve2D != nil {
		UnstageBranch(stage, plant.StackOfGrowthCurve2D)
	}
	if plant.TopStackOfGrowthCurve2D != nil {
		UnstageBranch(stage, plant.TopStackOfGrowthCurve2D)
	}
	if plant.StackOfGrowthCurve2DRibbon != nil {
		UnstageBranch(stage, plant.StackOfGrowthCurve2DRibbon)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plantdiagram := range plant.PlantDiagrams {
		UnstageBranch(stage, _plantdiagram)
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

func (stage *Stage) UnstageBranchPlantDiagram(plantdiagram *PlantDiagram) {

	// check if instance is already staged
	if !IsStaged(stage, plantdiagram) {
		return
	}

	plantdiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plantdiagram.Rendered3DShape != nil {
		UnstageBranch(stage, plantdiagram.Rendered3DShape)
	}
	if plantdiagram.TorusStackShape != nil {
		UnstageBranch(stage, plantdiagram.TorusStackShape)
	}
	if plantdiagram.VerticalTorusStackShape != nil {
		UnstageBranch(stage, plantdiagram.VerticalTorusStackShape)
	}

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
	if rhombusstuff.ReferenceRhombus != nil {
		UnstageBranch(stage, rhombusstuff.ReferenceRhombus)
	}
	if rhombusstuff.PlantCircumferenceShape != nil {
		UnstageBranch(stage, rhombusstuff.PlantCircumferenceShape)
	}
	if rhombusstuff.GridPathShape != nil {
		UnstageBranch(stage, rhombusstuff.GridPathShape)
	}
	if rhombusstuff.InitialRhombusGridShape != nil {
		UnstageBranch(stage, rhombusstuff.InitialRhombusGridShape)
	}
	if rhombusstuff.ExplanationTextShape != nil {
		UnstageBranch(stage, rhombusstuff.ExplanationTextShape)
	}
	if rhombusstuff.RotatedReferenceRhombus != nil {
		UnstageBranch(stage, rhombusstuff.RotatedReferenceRhombus)
	}
	if rhombusstuff.RotatedPlantCircumferenceShape != nil {
		UnstageBranch(stage, rhombusstuff.RotatedPlantCircumferenceShape)
	}
	if rhombusstuff.RotatedGridPathShape != nil {
		UnstageBranch(stage, rhombusstuff.RotatedGridPathShape)
	}
	if rhombusstuff.RotatedRhombusGridShape2 != nil {
		UnstageBranch(stage, rhombusstuff.RotatedRhombusGridShape2)
	}
	if rhombusstuff.GrowthCurveRhombusGridShape != nil {
		UnstageBranch(stage, rhombusstuff.GrowthCurveRhombusGridShape)
	}

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
	for _, _rotatedrhombusshape := range rotatedrhombusgridshape.RotatedRhombusShapes {
		UnstageBranch(stage, _rotatedrhombusshape)
	}

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
	for _, _shiftedbottomtopstartarcshape := range shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes {
		UnstageBranch(stage, _shiftedbottomtopstartarcshape)
	}

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
	for _, _shiftedleftstackgrowthcurvestartarcshape := range shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes {
		UnstageBranch(stage, _shiftedleftstackgrowthcurvestartarcshape)
	}
	for _, _shiftedleftstackgrowthcurveendarcshape := range shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes {
		UnstageBranch(stage, _shiftedleftstackgrowthcurveendarcshape)
	}

}

func (stage *Stage) UnstageBranchShiftedLeftStackOfNormalVector(shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) {

	// check if instance is already staged
	if !IsStaged(stage, shiftedleftstackofnormalvector) {
		return
	}

	shiftedleftstackofnormalvector.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _shiftedleftstacknormalvector := range shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors {
		UnstageBranch(stage, _shiftedleftstacknormalvector)
	}

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
	for _, _stackgrowthcurve2dstarthalfwayarcshape := range stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes {
		UnstageBranch(stage, _stackgrowthcurve2dstarthalfwayarcshape)
	}
	for _, _stackgrowthcurve2dendhalfwayarcshape := range stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes {
		UnstageBranch(stage, _stackgrowthcurve2dendhalfwayarcshape)
	}

}

func (stage *Stage) UnstageBranchStackOfGrowthCurve2DRibbon(stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) {

	// check if instance is already staged
	if !IsStaged(stage, stackofgrowthcurve2dribbon) {
		return
	}

	stackofgrowthcurve2dribbon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stackgrowthcurve2dribbonstartshape := range stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes {
		UnstageBranch(stage, _stackgrowthcurve2dribbonstartshape)
	}
	for _, _stackgrowthcurve2dribbonendshape := range stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes {
		UnstageBranch(stage, _stackgrowthcurve2dribbonendshape)
	}

}

func (stage *Stage) UnstageBranchStackOfRotatedGrowthCurve2D(stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) {

	// check if instance is already staged
	if !IsStaged(stage, stackofrotatedgrowthcurve2d) {
		return
	}

	stackofrotatedgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stackrotatedgrowthcurve2dstartarcshape := range stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes {
		UnstageBranch(stage, _stackrotatedgrowthcurve2dstartarcshape)
	}
	for _, _stackrotatedgrowthcurve2dendarcshape := range stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes {
		UnstageBranch(stage, _stackrotatedgrowthcurve2dendarcshape)
	}

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
	for _, _startarcshape := range startarcshapegrid.StartArcShapes {
		UnstageBranch(stage, _startarcshape)
	}

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
	for _, _starthalfwayarcshape := range starthalfwayarcshapegrid.StartHalfwayArcShapes {
		UnstageBranch(stage, _starthalfwayarcshape)
	}

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
	for _, _topendarcshape := range topendarcshapegrid.TopEndArcShapes {
		UnstageBranch(stage, _topendarcshape)
	}

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
	for _, _topendhalfwayarcshape := range topendhalfwayarcshapegrid.TopEndHalfwayArcShapes {
		UnstageBranch(stage, _topendhalfwayarcshape)
	}

}

func (stage *Stage) UnstageBranchTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D) {

	// check if instance is already staged
	if !IsStaged(stage, topgrowthcurve2d) {
		return
	}

	topgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if topgrowthcurve2d.TopStartHalfwayArcShapeGrid != nil {
		UnstageBranch(stage, topgrowthcurve2d.TopStartHalfwayArcShapeGrid)
	}
	if topgrowthcurve2d.TopEndHalfwayArcShapeGrid != nil {
		UnstageBranch(stage, topgrowthcurve2d.TopEndHalfwayArcShapeGrid)
	}

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
	for _, _topmidarcvectorshape := range topmidarcvectorshapegrid.TopMidArcVectorShapes {
		UnstageBranch(stage, _topmidarcvectorshape)
	}

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
	for _, _topstackgrowthcurve2dstarthalfwayarcshape := range topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes {
		UnstageBranch(stage, _topstackgrowthcurve2dstarthalfwayarcshape)
	}
	for _, _topstackgrowthcurve2dendhalfwayarcshape := range topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes {
		UnstageBranch(stage, _topstackgrowthcurve2dendhalfwayarcshape)
	}

}

func (stage *Stage) UnstageBranchTopStackOfRotatedGrowthCurve2D(topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) {

	// check if instance is already staged
	if !IsStaged(stage, topstackofrotatedgrowthcurve2d) {
		return
	}

	topstackofrotatedgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topstackofrotatedgrowthcurve2dstartarcshape := range topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes {
		UnstageBranch(stage, _topstackofrotatedgrowthcurve2dstartarcshape)
	}
	for _, _topstackofrotatedgrowthcurve2dendarcshape := range topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes {
		UnstageBranch(stage, _topstackofrotatedgrowthcurve2dendarcshape)
	}

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
	for _, _topstartarcshape := range topstartarcshapegrid.TopStartArcShapes {
		UnstageBranch(stage, _topstartarcshape)
	}

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
	for _, _topstarthalfwayarcshape := range topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes {
		UnstageBranch(stage, _topstarthalfwayarcshape)
	}

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

func (stage *Stage) UnstageBranchVerticalTorusStackShape(verticaltorusstackshape *VerticalTorusStackShape) {

	// check if instance is already staged
	if !IsStaged(stage, verticaltorusstackshape) {
		return
	}

	verticaltorusstackshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *ArcNormalVectorShape) GongReconstructPointersFromReferences(stage *Stage, instance *ArcNormalVectorShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ArcNormalVectorShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *ArcNormalVectorShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ArcNormalVectorShapes = reference.ArcNormalVectorShapes[:0]
	for _, _b := range instance.ArcNormalVectorShapes {
		reference.ArcNormalVectorShapes = append(reference.ArcNormalVectorShapes, stage.ArcNormalVectorShapes_reference[_b])
	}
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
	reference.BaseVectorShapes = reference.BaseVectorShapes[:0]
	for _, _b := range instance.BaseVectorShapes {
		reference.BaseVectorShapes = append(reference.BaseVectorShapes, stage.BaseVectorShapes_reference[_b])
	}
}

func (reference *CircleGridShape) GongReconstructPointersFromReferences(stage *Stage, instance *CircleGridShape) {
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
	reference.EndArcShapes = reference.EndArcShapes[:0]
	for _, _b := range instance.EndArcShapes {
		reference.EndArcShapes = append(reference.EndArcShapes, stage.EndArcShapes_reference[_b])
	}
}

func (reference *EndHalfwayArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *EndHalfwayArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EndHalfwayArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *EndHalfwayArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.EndHalfwayArcShapes = reference.EndHalfwayArcShapes[:0]
	for _, _b := range instance.EndHalfwayArcShapes {
		reference.EndHalfwayArcShapes = append(reference.EndHalfwayArcShapes, stage.EndHalfwayArcShapes_reference[_b])
	}
}

func (reference *ExplanationTextShape) GongReconstructPointersFromReferences(stage *Stage, instance *ExplanationTextShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GridPathShape) GongReconstructPointersFromReferences(stage *Stage, instance *GridPathShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GrowthCurve2D) GongReconstructPointersFromReferences(stage *Stage, instance *GrowthCurve2D) {
	// insertion point for pointers field
	if instance.StartHalfwayArcShapeGrid != nil {
		reference.StartHalfwayArcShapeGrid = stage.StartHalfwayArcShapeGrids_reference[instance.StartHalfwayArcShapeGrid]
	}
	if instance.EndHalfwayArcShapeGrid != nil {
		reference.EndHalfwayArcShapeGrid = stage.EndHalfwayArcShapeGrids_reference[instance.EndHalfwayArcShapeGrid]
	}
	// insertion point for slice of pointers field
}

func (reference *GrowthCurveRhombusGridShape) GongReconstructPointersFromReferences(stage *Stage, instance *GrowthCurveRhombusGridShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.GrowthCurveRhombusShapes = reference.GrowthCurveRhombusShapes[:0]
	for _, _b := range instance.GrowthCurveRhombusShapes {
		reference.GrowthCurveRhombusShapes = append(reference.GrowthCurveRhombusShapes, stage.GrowthCurveRhombusShapes_reference[_b])
	}
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
	reference.InitialRhombusShapes = reference.InitialRhombusShapes[:0]
	for _, _b := range instance.InitialRhombusShapes {
		reference.InitialRhombusShapes = append(reference.InitialRhombusShapes, stage.InitialRhombusShapes_reference[_b])
	}
}

func (reference *InitialRhombusShape) GongReconstructPointersFromReferences(stage *Stage, instance *InitialRhombusShape) {
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
		reference.Plants = append(reference.Plants, stage.Plants_reference[_b])
	}
}

func (reference *MidArcVectorShape) GongReconstructPointersFromReferences(stage *Stage, instance *MidArcVectorShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MidArcVectorShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *MidArcVectorShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.MidArcVectorShapes = reference.MidArcVectorShapes[:0]
	for _, _b := range instance.MidArcVectorShapes {
		reference.MidArcVectorShapes = append(reference.MidArcVectorShapes, stage.MidArcVectorShapes_reference[_b])
	}
}

func (reference *PerpendicularVector) GongReconstructPointersFromReferences(stage *Stage, instance *PerpendicularVector) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PerpendicularVectorGrid) GongReconstructPointersFromReferences(stage *Stage, instance *PerpendicularVectorGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.PerpendicularVectors = reference.PerpendicularVectors[:0]
	for _, _b := range instance.PerpendicularVectors {
		reference.PerpendicularVectors = append(reference.PerpendicularVectors, stage.PerpendicularVectors_reference[_b])
	}
}

func (reference *PerpendicularVectorGridHalfway) GongReconstructPointersFromReferences(stage *Stage, instance *PerpendicularVectorGridHalfway) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.PerpendicularVectorHalfways = reference.PerpendicularVectorHalfways[:0]
	for _, _b := range instance.PerpendicularVectorHalfways {
		reference.PerpendicularVectorHalfways = append(reference.PerpendicularVectorHalfways, stage.PerpendicularVectorHalfways_reference[_b])
	}
}

func (reference *PerpendicularVectorHalfway) GongReconstructPointersFromReferences(stage *Stage, instance *PerpendicularVectorHalfway) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Plant) GongReconstructPointersFromReferences(stage *Stage, instance *Plant) {
	// insertion point for pointers field
	if instance.AxesShape != nil {
		reference.AxesShape = stage.AxesShapes_reference[instance.AxesShape]
	}
	if instance.RhombusStuff != nil {
		reference.RhombusStuff = stage.RhombusStuffs_reference[instance.RhombusStuff]
	}
	if instance.GrowthVectorShape != nil {
		reference.GrowthVectorShape = stage.GrowthVectorShapes_reference[instance.GrowthVectorShape]
	}
	if instance.PerpendicularVectorGrid != nil {
		reference.PerpendicularVectorGrid = stage.PerpendicularVectorGrids_reference[instance.PerpendicularVectorGrid]
	}
	if instance.PerpendicularVectorGridHalfway != nil {
		reference.PerpendicularVectorGridHalfway = stage.PerpendicularVectorGridHalfways_reference[instance.PerpendicularVectorGridHalfway]
	}
	if instance.BaseVectorShapeGrid != nil {
		reference.BaseVectorShapeGrid = stage.BaseVectorShapeGrids_reference[instance.BaseVectorShapeGrid]
	}
	if instance.ArcNormalVectorShapeGrid != nil {
		reference.ArcNormalVectorShapeGrid = stage.ArcNormalVectorShapeGrids_reference[instance.ArcNormalVectorShapeGrid]
	}
	if instance.StartArcShapeGrid != nil {
		reference.StartArcShapeGrid = stage.StartArcShapeGrids_reference[instance.StartArcShapeGrid]
	}
	if instance.TopStartArcShapeGrid != nil {
		reference.TopStartArcShapeGrid = stage.TopStartArcShapeGrids_reference[instance.TopStartArcShapeGrid]
	}
	if instance.EndArcShapeGrid != nil {
		reference.EndArcShapeGrid = stage.EndArcShapeGrids_reference[instance.EndArcShapeGrid]
	}
	if instance.TopEndArcShapeGrid != nil {
		reference.TopEndArcShapeGrid = stage.TopEndArcShapeGrids_reference[instance.TopEndArcShapeGrid]
	}
	if instance.ShiftedBottomTopStartArcShapeGrid != nil {
		reference.ShiftedBottomTopStartArcShapeGrid = stage.ShiftedBottomTopStartArcShapeGrids_reference[instance.ShiftedBottomTopStartArcShapeGrid]
	}
	if instance.MidArcVectorShapeGrid != nil {
		reference.MidArcVectorShapeGrid = stage.MidArcVectorShapeGrids_reference[instance.MidArcVectorShapeGrid]
	}
	if instance.TopMidArcVectorShapeGrid != nil {
		reference.TopMidArcVectorShapeGrid = stage.TopMidArcVectorShapeGrids_reference[instance.TopMidArcVectorShapeGrid]
	}
	if instance.StartHalfwayArcShapeGrid != nil {
		reference.StartHalfwayArcShapeGrid = stage.StartHalfwayArcShapeGrids_reference[instance.StartHalfwayArcShapeGrid]
	}
	if instance.TopStartHalfwayArcShapeGrid != nil {
		reference.TopStartHalfwayArcShapeGrid = stage.TopStartHalfwayArcShapeGrids_reference[instance.TopStartHalfwayArcShapeGrid]
	}
	if instance.EndHalfwayArcShapeGrid != nil {
		reference.EndHalfwayArcShapeGrid = stage.EndHalfwayArcShapeGrids_reference[instance.EndHalfwayArcShapeGrid]
	}
	if instance.TopEndHalfwayArcShapeGrid != nil {
		reference.TopEndHalfwayArcShapeGrid = stage.TopEndHalfwayArcShapeGrids_reference[instance.TopEndHalfwayArcShapeGrid]
	}
	if instance.StackOfRotatedGrowthCurve2D != nil {
		reference.StackOfRotatedGrowthCurve2D = stage.StackOfRotatedGrowthCurve2Ds_reference[instance.StackOfRotatedGrowthCurve2D]
	}
	if instance.TopStackOfRotatedGrowthCurve2D != nil {
		reference.TopStackOfRotatedGrowthCurve2D = stage.TopStackOfRotatedGrowthCurve2Ds_reference[instance.TopStackOfRotatedGrowthCurve2D]
	}
	if instance.GrowthCurve2D != nil {
		reference.GrowthCurve2D = stage.GrowthCurve2Ds_reference[instance.GrowthCurve2D]
	}
	if instance.TopGrowthCurve2D != nil {
		reference.TopGrowthCurve2D = stage.TopGrowthCurve2Ds_reference[instance.TopGrowthCurve2D]
	}
	if instance.StackOfGrowthCurve2D != nil {
		reference.StackOfGrowthCurve2D = stage.StackOfGrowthCurve2Ds_reference[instance.StackOfGrowthCurve2D]
	}
	if instance.TopStackOfGrowthCurve2D != nil {
		reference.TopStackOfGrowthCurve2D = stage.TopStackOfGrowthCurve2Ds_reference[instance.TopStackOfGrowthCurve2D]
	}
	if instance.StackOfGrowthCurve2DRibbon != nil {
		reference.StackOfGrowthCurve2DRibbon = stage.StackOfGrowthCurve2DRibbons_reference[instance.StackOfGrowthCurve2DRibbon]
	}
	// insertion point for slice of pointers field
	reference.PlantDiagrams = reference.PlantDiagrams[:0]
	for _, _b := range instance.PlantDiagrams {
		reference.PlantDiagrams = append(reference.PlantDiagrams, stage.PlantDiagrams_reference[_b])
	}
}

func (reference *PlantCircumferenceShape) GongReconstructPointersFromReferences(stage *Stage, instance *PlantCircumferenceShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PlantDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *PlantDiagram) {
	// insertion point for pointers field
	if instance.Rendered3DShape != nil {
		reference.Rendered3DShape = stage.Rendered3DShapes_reference[instance.Rendered3DShape]
	}
	if instance.TorusStackShape != nil {
		reference.TorusStackShape = stage.TorusStackShapes_reference[instance.TorusStackShape]
	}
	if instance.VerticalTorusStackShape != nil {
		reference.VerticalTorusStackShape = stage.VerticalTorusStackShapes_reference[instance.VerticalTorusStackShape]
	}
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
	if instance.ReferenceRhombus != nil {
		reference.ReferenceRhombus = stage.RhombusShapes_reference[instance.ReferenceRhombus]
	}
	if instance.PlantCircumferenceShape != nil {
		reference.PlantCircumferenceShape = stage.PlantCircumferenceShapes_reference[instance.PlantCircumferenceShape]
	}
	if instance.GridPathShape != nil {
		reference.GridPathShape = stage.GridPathShapes_reference[instance.GridPathShape]
	}
	if instance.InitialRhombusGridShape != nil {
		reference.InitialRhombusGridShape = stage.InitialRhombusGridShapes_reference[instance.InitialRhombusGridShape]
	}
	if instance.ExplanationTextShape != nil {
		reference.ExplanationTextShape = stage.ExplanationTextShapes_reference[instance.ExplanationTextShape]
	}
	if instance.RotatedReferenceRhombus != nil {
		reference.RotatedReferenceRhombus = stage.RhombusShapes_reference[instance.RotatedReferenceRhombus]
	}
	if instance.RotatedPlantCircumferenceShape != nil {
		reference.RotatedPlantCircumferenceShape = stage.PlantCircumferenceShapes_reference[instance.RotatedPlantCircumferenceShape]
	}
	if instance.RotatedGridPathShape != nil {
		reference.RotatedGridPathShape = stage.GridPathShapes_reference[instance.RotatedGridPathShape]
	}
	if instance.RotatedRhombusGridShape2 != nil {
		reference.RotatedRhombusGridShape2 = stage.RotatedRhombusGridShapes_reference[instance.RotatedRhombusGridShape2]
	}
	if instance.GrowthCurveRhombusGridShape != nil {
		reference.GrowthCurveRhombusGridShape = stage.GrowthCurveRhombusGridShapes_reference[instance.GrowthCurveRhombusGridShape]
	}
	// insertion point for slice of pointers field
}

func (reference *RotatedRhombusGridShape) GongReconstructPointersFromReferences(stage *Stage, instance *RotatedRhombusGridShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.RotatedRhombusShapes = reference.RotatedRhombusShapes[:0]
	for _, _b := range instance.RotatedRhombusShapes {
		reference.RotatedRhombusShapes = append(reference.RotatedRhombusShapes, stage.RotatedRhombusShapes_reference[_b])
	}
}

func (reference *RotatedRhombusShape) GongReconstructPointersFromReferences(stage *Stage, instance *RotatedRhombusShape) {
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
	reference.ShiftedBottomTopStartArcShapes = reference.ShiftedBottomTopStartArcShapes[:0]
	for _, _b := range instance.ShiftedBottomTopStartArcShapes {
		reference.ShiftedBottomTopStartArcShapes = append(reference.ShiftedBottomTopStartArcShapes, stage.ShiftedBottomTopStartArcShapes_reference[_b])
	}
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
	reference.ShiftedLeftStackGrowthCurveStartArcShapes = reference.ShiftedLeftStackGrowthCurveStartArcShapes[:0]
	for _, _b := range instance.ShiftedLeftStackGrowthCurveStartArcShapes {
		reference.ShiftedLeftStackGrowthCurveStartArcShapes = append(reference.ShiftedLeftStackGrowthCurveStartArcShapes, stage.ShiftedLeftStackGrowthCurveStartArcShapes_reference[_b])
	}
	reference.ShiftedLeftStackGrowthCurveEndArcShapes = reference.ShiftedLeftStackGrowthCurveEndArcShapes[:0]
	for _, _b := range instance.ShiftedLeftStackGrowthCurveEndArcShapes {
		reference.ShiftedLeftStackGrowthCurveEndArcShapes = append(reference.ShiftedLeftStackGrowthCurveEndArcShapes, stage.ShiftedLeftStackGrowthCurveEndArcShapes_reference[_b])
	}
}

func (reference *ShiftedLeftStackOfNormalVector) GongReconstructPointersFromReferences(stage *Stage, instance *ShiftedLeftStackOfNormalVector) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ShiftedLeftStackNormalVectors = reference.ShiftedLeftStackNormalVectors[:0]
	for _, _b := range instance.ShiftedLeftStackNormalVectors {
		reference.ShiftedLeftStackNormalVectors = append(reference.ShiftedLeftStackNormalVectors, stage.ShiftedLeftStackNormalVectors_reference[_b])
	}
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
	reference.StackGrowthCurve2DStartHalfwayArcShapes = reference.StackGrowthCurve2DStartHalfwayArcShapes[:0]
	for _, _b := range instance.StackGrowthCurve2DStartHalfwayArcShapes {
		reference.StackGrowthCurve2DStartHalfwayArcShapes = append(reference.StackGrowthCurve2DStartHalfwayArcShapes, stage.StackGrowthCurve2DStartHalfwayArcShapes_reference[_b])
	}
	reference.StackGrowthCurve2DEndHalfwayArcShapes = reference.StackGrowthCurve2DEndHalfwayArcShapes[:0]
	for _, _b := range instance.StackGrowthCurve2DEndHalfwayArcShapes {
		reference.StackGrowthCurve2DEndHalfwayArcShapes = append(reference.StackGrowthCurve2DEndHalfwayArcShapes, stage.StackGrowthCurve2DEndHalfwayArcShapes_reference[_b])
	}
}

func (reference *StackOfGrowthCurve2DRibbon) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfGrowthCurve2DRibbon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.StackGrowthCurve2DRibbonStartShapes = reference.StackGrowthCurve2DRibbonStartShapes[:0]
	for _, _b := range instance.StackGrowthCurve2DRibbonStartShapes {
		reference.StackGrowthCurve2DRibbonStartShapes = append(reference.StackGrowthCurve2DRibbonStartShapes, stage.StackGrowthCurve2DRibbonStartShapes_reference[_b])
	}
	reference.StackGrowthCurve2DRibbonEndShapes = reference.StackGrowthCurve2DRibbonEndShapes[:0]
	for _, _b := range instance.StackGrowthCurve2DRibbonEndShapes {
		reference.StackGrowthCurve2DRibbonEndShapes = append(reference.StackGrowthCurve2DRibbonEndShapes, stage.StackGrowthCurve2DRibbonEndShapes_reference[_b])
	}
}

func (reference *StackOfRotatedGrowthCurve2D) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfRotatedGrowthCurve2D) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.StackRotatedGrowthCurve2DStartArcShapes = reference.StackRotatedGrowthCurve2DStartArcShapes[:0]
	for _, _b := range instance.StackRotatedGrowthCurve2DStartArcShapes {
		reference.StackRotatedGrowthCurve2DStartArcShapes = append(reference.StackRotatedGrowthCurve2DStartArcShapes, stage.StackRotatedGrowthCurve2DStartArcShapes_reference[_b])
	}
	reference.StackRotatedGrowthCurve2DEndArcShapes = reference.StackRotatedGrowthCurve2DEndArcShapes[:0]
	for _, _b := range instance.StackRotatedGrowthCurve2DEndArcShapes {
		reference.StackRotatedGrowthCurve2DEndArcShapes = append(reference.StackRotatedGrowthCurve2DEndArcShapes, stage.StackRotatedGrowthCurve2DEndArcShapes_reference[_b])
	}
}

func (reference *StackRotatedGrowthCurve2DEndArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackRotatedGrowthCurve2DEndArcShape) {
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
	reference.StartArcShapes = reference.StartArcShapes[:0]
	for _, _b := range instance.StartArcShapes {
		reference.StartArcShapes = append(reference.StartArcShapes, stage.StartArcShapes_reference[_b])
	}
}

func (reference *StartHalfwayArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *StartHalfwayArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StartHalfwayArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *StartHalfwayArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.StartHalfwayArcShapes = reference.StartHalfwayArcShapes[:0]
	for _, _b := range instance.StartHalfwayArcShapes {
		reference.StartHalfwayArcShapes = append(reference.StartHalfwayArcShapes, stage.StartHalfwayArcShapes_reference[_b])
	}
}

func (reference *TopEndArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopEndArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopEndArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *TopEndArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.TopEndArcShapes = reference.TopEndArcShapes[:0]
	for _, _b := range instance.TopEndArcShapes {
		reference.TopEndArcShapes = append(reference.TopEndArcShapes, stage.TopEndArcShapes_reference[_b])
	}
}

func (reference *TopEndHalfwayArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopEndHalfwayArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopEndHalfwayArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *TopEndHalfwayArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.TopEndHalfwayArcShapes = reference.TopEndHalfwayArcShapes[:0]
	for _, _b := range instance.TopEndHalfwayArcShapes {
		reference.TopEndHalfwayArcShapes = append(reference.TopEndHalfwayArcShapes, stage.TopEndHalfwayArcShapes_reference[_b])
	}
}

func (reference *TopGrowthCurve2D) GongReconstructPointersFromReferences(stage *Stage, instance *TopGrowthCurve2D) {
	// insertion point for pointers field
	if instance.TopStartHalfwayArcShapeGrid != nil {
		reference.TopStartHalfwayArcShapeGrid = stage.TopStartHalfwayArcShapeGrids_reference[instance.TopStartHalfwayArcShapeGrid]
	}
	if instance.TopEndHalfwayArcShapeGrid != nil {
		reference.TopEndHalfwayArcShapeGrid = stage.TopEndHalfwayArcShapeGrids_reference[instance.TopEndHalfwayArcShapeGrid]
	}
	// insertion point for slice of pointers field
}

func (reference *TopMidArcVectorShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopMidArcVectorShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopMidArcVectorShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *TopMidArcVectorShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.TopMidArcVectorShapes = reference.TopMidArcVectorShapes[:0]
	for _, _b := range instance.TopMidArcVectorShapes {
		reference.TopMidArcVectorShapes = append(reference.TopMidArcVectorShapes, stage.TopMidArcVectorShapes_reference[_b])
	}
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
	reference.TopStackGrowthCurve2DStartHalfwayArcShapes = reference.TopStackGrowthCurve2DStartHalfwayArcShapes[:0]
	for _, _b := range instance.TopStackGrowthCurve2DStartHalfwayArcShapes {
		reference.TopStackGrowthCurve2DStartHalfwayArcShapes = append(reference.TopStackGrowthCurve2DStartHalfwayArcShapes, stage.TopStackGrowthCurve2DStartHalfwayArcShapes_reference[_b])
	}
	reference.TopStackGrowthCurve2DEndHalfwayArcShapes = reference.TopStackGrowthCurve2DEndHalfwayArcShapes[:0]
	for _, _b := range instance.TopStackGrowthCurve2DEndHalfwayArcShapes {
		reference.TopStackGrowthCurve2DEndHalfwayArcShapes = append(reference.TopStackGrowthCurve2DEndHalfwayArcShapes, stage.TopStackGrowthCurve2DEndHalfwayArcShapes_reference[_b])
	}
}

func (reference *TopStackOfRotatedGrowthCurve2D) GongReconstructPointersFromReferences(stage *Stage, instance *TopStackOfRotatedGrowthCurve2D) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.TopStackOfRotatedGrowthCurve2DStartArcShapes = reference.TopStackOfRotatedGrowthCurve2DStartArcShapes[:0]
	for _, _b := range instance.TopStackOfRotatedGrowthCurve2DStartArcShapes {
		reference.TopStackOfRotatedGrowthCurve2DStartArcShapes = append(reference.TopStackOfRotatedGrowthCurve2DStartArcShapes, stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_reference[_b])
	}
	reference.TopStackOfRotatedGrowthCurve2DEndArcShapes = reference.TopStackOfRotatedGrowthCurve2DEndArcShapes[:0]
	for _, _b := range instance.TopStackOfRotatedGrowthCurve2DEndArcShapes {
		reference.TopStackOfRotatedGrowthCurve2DEndArcShapes = append(reference.TopStackOfRotatedGrowthCurve2DEndArcShapes, stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_reference[_b])
	}
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
	reference.TopStartArcShapes = reference.TopStartArcShapes[:0]
	for _, _b := range instance.TopStartArcShapes {
		reference.TopStartArcShapes = append(reference.TopStartArcShapes, stage.TopStartArcShapes_reference[_b])
	}
}

func (reference *TopStartHalfwayArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopStartHalfwayArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStartHalfwayArcShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *TopStartHalfwayArcShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.TopStartHalfwayArcShapes = reference.TopStartHalfwayArcShapes[:0]
	for _, _b := range instance.TopStartHalfwayArcShapes {
		reference.TopStartHalfwayArcShapes = append(reference.TopStartHalfwayArcShapes, stage.TopStartHalfwayArcShapes_reference[_b])
	}
}

func (reference *TorusStackShape) GongReconstructPointersFromReferences(stage *Stage, instance *TorusStackShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *VerticalTorusStackShape) GongReconstructPointersFromReferences(stage *Stage, instance *VerticalTorusStackShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *ArcNormalVectorShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ArcNormalVectorShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ArcNormalVectorShapes []*ArcNormalVectorShape
	for _, _reference := range reference.ArcNormalVectorShapes {
		if _instance, ok := stage.ArcNormalVectorShapes_instance[_reference]; ok {
			_ArcNormalVectorShapes = append(_ArcNormalVectorShapes, _instance)
		}
	}
	reference.ArcNormalVectorShapes = _ArcNormalVectorShapes
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
	var _BaseVectorShapes []*BaseVectorShape
	for _, _reference := range reference.BaseVectorShapes {
		if _instance, ok := stage.BaseVectorShapes_instance[_reference]; ok {
			_BaseVectorShapes = append(_BaseVectorShapes, _instance)
		}
	}
	reference.BaseVectorShapes = _BaseVectorShapes
}

func (reference *CircleGridShape) GongReconstructPointersFromInstances(stage *Stage) {
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
	var _EndArcShapes []*EndArcShape
	for _, _reference := range reference.EndArcShapes {
		if _instance, ok := stage.EndArcShapes_instance[_reference]; ok {
			_EndArcShapes = append(_EndArcShapes, _instance)
		}
	}
	reference.EndArcShapes = _EndArcShapes
}

func (reference *EndHalfwayArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EndHalfwayArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _EndHalfwayArcShapes []*EndHalfwayArcShape
	for _, _reference := range reference.EndHalfwayArcShapes {
		if _instance, ok := stage.EndHalfwayArcShapes_instance[_reference]; ok {
			_EndHalfwayArcShapes = append(_EndHalfwayArcShapes, _instance)
		}
	}
	reference.EndHalfwayArcShapes = _EndHalfwayArcShapes
}

func (reference *ExplanationTextShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GridPathShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GrowthCurve2D) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.StartHalfwayArcShapeGrid; _reference != nil {
		reference.StartHalfwayArcShapeGrid = nil
		if _instance, ok := stage.StartHalfwayArcShapeGrids_instance[_reference]; ok {
			reference.StartHalfwayArcShapeGrid = _instance
		}
	}
	if _reference := reference.EndHalfwayArcShapeGrid; _reference != nil {
		reference.EndHalfwayArcShapeGrid = nil
		if _instance, ok := stage.EndHalfwayArcShapeGrids_instance[_reference]; ok {
			reference.EndHalfwayArcShapeGrid = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *GrowthCurveRhombusGridShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _GrowthCurveRhombusShapes []*GrowthCurveRhombusShape
	for _, _reference := range reference.GrowthCurveRhombusShapes {
		if _instance, ok := stage.GrowthCurveRhombusShapes_instance[_reference]; ok {
			_GrowthCurveRhombusShapes = append(_GrowthCurveRhombusShapes, _instance)
		}
	}
	reference.GrowthCurveRhombusShapes = _GrowthCurveRhombusShapes
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
	var _InitialRhombusShapes []*InitialRhombusShape
	for _, _reference := range reference.InitialRhombusShapes {
		if _instance, ok := stage.InitialRhombusShapes_instance[_reference]; ok {
			_InitialRhombusShapes = append(_InitialRhombusShapes, _instance)
		}
	}
	reference.InitialRhombusShapes = _InitialRhombusShapes
}

func (reference *InitialRhombusShape) GongReconstructPointersFromInstances(stage *Stage) {
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
	var _Plants []*Plant
	for _, _reference := range reference.Plants {
		if _instance, ok := stage.Plants_instance[_reference]; ok {
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
	var _MidArcVectorShapes []*MidArcVectorShape
	for _, _reference := range reference.MidArcVectorShapes {
		if _instance, ok := stage.MidArcVectorShapes_instance[_reference]; ok {
			_MidArcVectorShapes = append(_MidArcVectorShapes, _instance)
		}
	}
	reference.MidArcVectorShapes = _MidArcVectorShapes
}

func (reference *PerpendicularVector) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PerpendicularVectorGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _PerpendicularVectors []*PerpendicularVector
	for _, _reference := range reference.PerpendicularVectors {
		if _instance, ok := stage.PerpendicularVectors_instance[_reference]; ok {
			_PerpendicularVectors = append(_PerpendicularVectors, _instance)
		}
	}
	reference.PerpendicularVectors = _PerpendicularVectors
}

func (reference *PerpendicularVectorGridHalfway) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _PerpendicularVectorHalfways []*PerpendicularVectorHalfway
	for _, _reference := range reference.PerpendicularVectorHalfways {
		if _instance, ok := stage.PerpendicularVectorHalfways_instance[_reference]; ok {
			_PerpendicularVectorHalfways = append(_PerpendicularVectorHalfways, _instance)
		}
	}
	reference.PerpendicularVectorHalfways = _PerpendicularVectorHalfways
}

func (reference *PerpendicularVectorHalfway) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Plant) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.AxesShape; _reference != nil {
		reference.AxesShape = nil
		if _instance, ok := stage.AxesShapes_instance[_reference]; ok {
			reference.AxesShape = _instance
		}
	}
	if _reference := reference.RhombusStuff; _reference != nil {
		reference.RhombusStuff = nil
		if _instance, ok := stage.RhombusStuffs_instance[_reference]; ok {
			reference.RhombusStuff = _instance
		}
	}
	if _reference := reference.GrowthVectorShape; _reference != nil {
		reference.GrowthVectorShape = nil
		if _instance, ok := stage.GrowthVectorShapes_instance[_reference]; ok {
			reference.GrowthVectorShape = _instance
		}
	}
	if _reference := reference.PerpendicularVectorGrid; _reference != nil {
		reference.PerpendicularVectorGrid = nil
		if _instance, ok := stage.PerpendicularVectorGrids_instance[_reference]; ok {
			reference.PerpendicularVectorGrid = _instance
		}
	}
	if _reference := reference.PerpendicularVectorGridHalfway; _reference != nil {
		reference.PerpendicularVectorGridHalfway = nil
		if _instance, ok := stage.PerpendicularVectorGridHalfways_instance[_reference]; ok {
			reference.PerpendicularVectorGridHalfway = _instance
		}
	}
	if _reference := reference.BaseVectorShapeGrid; _reference != nil {
		reference.BaseVectorShapeGrid = nil
		if _instance, ok := stage.BaseVectorShapeGrids_instance[_reference]; ok {
			reference.BaseVectorShapeGrid = _instance
		}
	}
	if _reference := reference.ArcNormalVectorShapeGrid; _reference != nil {
		reference.ArcNormalVectorShapeGrid = nil
		if _instance, ok := stage.ArcNormalVectorShapeGrids_instance[_reference]; ok {
			reference.ArcNormalVectorShapeGrid = _instance
		}
	}
	if _reference := reference.StartArcShapeGrid; _reference != nil {
		reference.StartArcShapeGrid = nil
		if _instance, ok := stage.StartArcShapeGrids_instance[_reference]; ok {
			reference.StartArcShapeGrid = _instance
		}
	}
	if _reference := reference.TopStartArcShapeGrid; _reference != nil {
		reference.TopStartArcShapeGrid = nil
		if _instance, ok := stage.TopStartArcShapeGrids_instance[_reference]; ok {
			reference.TopStartArcShapeGrid = _instance
		}
	}
	if _reference := reference.EndArcShapeGrid; _reference != nil {
		reference.EndArcShapeGrid = nil
		if _instance, ok := stage.EndArcShapeGrids_instance[_reference]; ok {
			reference.EndArcShapeGrid = _instance
		}
	}
	if _reference := reference.TopEndArcShapeGrid; _reference != nil {
		reference.TopEndArcShapeGrid = nil
		if _instance, ok := stage.TopEndArcShapeGrids_instance[_reference]; ok {
			reference.TopEndArcShapeGrid = _instance
		}
	}
	if _reference := reference.ShiftedBottomTopStartArcShapeGrid; _reference != nil {
		reference.ShiftedBottomTopStartArcShapeGrid = nil
		if _instance, ok := stage.ShiftedBottomTopStartArcShapeGrids_instance[_reference]; ok {
			reference.ShiftedBottomTopStartArcShapeGrid = _instance
		}
	}
	if _reference := reference.MidArcVectorShapeGrid; _reference != nil {
		reference.MidArcVectorShapeGrid = nil
		if _instance, ok := stage.MidArcVectorShapeGrids_instance[_reference]; ok {
			reference.MidArcVectorShapeGrid = _instance
		}
	}
	if _reference := reference.TopMidArcVectorShapeGrid; _reference != nil {
		reference.TopMidArcVectorShapeGrid = nil
		if _instance, ok := stage.TopMidArcVectorShapeGrids_instance[_reference]; ok {
			reference.TopMidArcVectorShapeGrid = _instance
		}
	}
	if _reference := reference.StartHalfwayArcShapeGrid; _reference != nil {
		reference.StartHalfwayArcShapeGrid = nil
		if _instance, ok := stage.StartHalfwayArcShapeGrids_instance[_reference]; ok {
			reference.StartHalfwayArcShapeGrid = _instance
		}
	}
	if _reference := reference.TopStartHalfwayArcShapeGrid; _reference != nil {
		reference.TopStartHalfwayArcShapeGrid = nil
		if _instance, ok := stage.TopStartHalfwayArcShapeGrids_instance[_reference]; ok {
			reference.TopStartHalfwayArcShapeGrid = _instance
		}
	}
	if _reference := reference.EndHalfwayArcShapeGrid; _reference != nil {
		reference.EndHalfwayArcShapeGrid = nil
		if _instance, ok := stage.EndHalfwayArcShapeGrids_instance[_reference]; ok {
			reference.EndHalfwayArcShapeGrid = _instance
		}
	}
	if _reference := reference.TopEndHalfwayArcShapeGrid; _reference != nil {
		reference.TopEndHalfwayArcShapeGrid = nil
		if _instance, ok := stage.TopEndHalfwayArcShapeGrids_instance[_reference]; ok {
			reference.TopEndHalfwayArcShapeGrid = _instance
		}
	}
	if _reference := reference.StackOfRotatedGrowthCurve2D; _reference != nil {
		reference.StackOfRotatedGrowthCurve2D = nil
		if _instance, ok := stage.StackOfRotatedGrowthCurve2Ds_instance[_reference]; ok {
			reference.StackOfRotatedGrowthCurve2D = _instance
		}
	}
	if _reference := reference.TopStackOfRotatedGrowthCurve2D; _reference != nil {
		reference.TopStackOfRotatedGrowthCurve2D = nil
		if _instance, ok := stage.TopStackOfRotatedGrowthCurve2Ds_instance[_reference]; ok {
			reference.TopStackOfRotatedGrowthCurve2D = _instance
		}
	}
	if _reference := reference.GrowthCurve2D; _reference != nil {
		reference.GrowthCurve2D = nil
		if _instance, ok := stage.GrowthCurve2Ds_instance[_reference]; ok {
			reference.GrowthCurve2D = _instance
		}
	}
	if _reference := reference.TopGrowthCurve2D; _reference != nil {
		reference.TopGrowthCurve2D = nil
		if _instance, ok := stage.TopGrowthCurve2Ds_instance[_reference]; ok {
			reference.TopGrowthCurve2D = _instance
		}
	}
	if _reference := reference.StackOfGrowthCurve2D; _reference != nil {
		reference.StackOfGrowthCurve2D = nil
		if _instance, ok := stage.StackOfGrowthCurve2Ds_instance[_reference]; ok {
			reference.StackOfGrowthCurve2D = _instance
		}
	}
	if _reference := reference.TopStackOfGrowthCurve2D; _reference != nil {
		reference.TopStackOfGrowthCurve2D = nil
		if _instance, ok := stage.TopStackOfGrowthCurve2Ds_instance[_reference]; ok {
			reference.TopStackOfGrowthCurve2D = _instance
		}
	}
	if _reference := reference.StackOfGrowthCurve2DRibbon; _reference != nil {
		reference.StackOfGrowthCurve2DRibbon = nil
		if _instance, ok := stage.StackOfGrowthCurve2DRibbons_instance[_reference]; ok {
			reference.StackOfGrowthCurve2DRibbon = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _PlantDiagrams []*PlantDiagram
	for _, _reference := range reference.PlantDiagrams {
		if _instance, ok := stage.PlantDiagrams_instance[_reference]; ok {
			_PlantDiagrams = append(_PlantDiagrams, _instance)
		}
	}
	reference.PlantDiagrams = _PlantDiagrams
}

func (reference *PlantCircumferenceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PlantDiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Rendered3DShape; _reference != nil {
		reference.Rendered3DShape = nil
		if _instance, ok := stage.Rendered3DShapes_instance[_reference]; ok {
			reference.Rendered3DShape = _instance
		}
	}
	if _reference := reference.TorusStackShape; _reference != nil {
		reference.TorusStackShape = nil
		if _instance, ok := stage.TorusStackShapes_instance[_reference]; ok {
			reference.TorusStackShape = _instance
		}
	}
	if _reference := reference.VerticalTorusStackShape; _reference != nil {
		reference.VerticalTorusStackShape = nil
		if _instance, ok := stage.VerticalTorusStackShapes_instance[_reference]; ok {
			reference.VerticalTorusStackShape = _instance
		}
	}
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
	if _reference := reference.ReferenceRhombus; _reference != nil {
		reference.ReferenceRhombus = nil
		if _instance, ok := stage.RhombusShapes_instance[_reference]; ok {
			reference.ReferenceRhombus = _instance
		}
	}
	if _reference := reference.PlantCircumferenceShape; _reference != nil {
		reference.PlantCircumferenceShape = nil
		if _instance, ok := stage.PlantCircumferenceShapes_instance[_reference]; ok {
			reference.PlantCircumferenceShape = _instance
		}
	}
	if _reference := reference.GridPathShape; _reference != nil {
		reference.GridPathShape = nil
		if _instance, ok := stage.GridPathShapes_instance[_reference]; ok {
			reference.GridPathShape = _instance
		}
	}
	if _reference := reference.InitialRhombusGridShape; _reference != nil {
		reference.InitialRhombusGridShape = nil
		if _instance, ok := stage.InitialRhombusGridShapes_instance[_reference]; ok {
			reference.InitialRhombusGridShape = _instance
		}
	}
	if _reference := reference.ExplanationTextShape; _reference != nil {
		reference.ExplanationTextShape = nil
		if _instance, ok := stage.ExplanationTextShapes_instance[_reference]; ok {
			reference.ExplanationTextShape = _instance
		}
	}
	if _reference := reference.RotatedReferenceRhombus; _reference != nil {
		reference.RotatedReferenceRhombus = nil
		if _instance, ok := stage.RhombusShapes_instance[_reference]; ok {
			reference.RotatedReferenceRhombus = _instance
		}
	}
	if _reference := reference.RotatedPlantCircumferenceShape; _reference != nil {
		reference.RotatedPlantCircumferenceShape = nil
		if _instance, ok := stage.PlantCircumferenceShapes_instance[_reference]; ok {
			reference.RotatedPlantCircumferenceShape = _instance
		}
	}
	if _reference := reference.RotatedGridPathShape; _reference != nil {
		reference.RotatedGridPathShape = nil
		if _instance, ok := stage.GridPathShapes_instance[_reference]; ok {
			reference.RotatedGridPathShape = _instance
		}
	}
	if _reference := reference.RotatedRhombusGridShape2; _reference != nil {
		reference.RotatedRhombusGridShape2 = nil
		if _instance, ok := stage.RotatedRhombusGridShapes_instance[_reference]; ok {
			reference.RotatedRhombusGridShape2 = _instance
		}
	}
	if _reference := reference.GrowthCurveRhombusGridShape; _reference != nil {
		reference.GrowthCurveRhombusGridShape = nil
		if _instance, ok := stage.GrowthCurveRhombusGridShapes_instance[_reference]; ok {
			reference.GrowthCurveRhombusGridShape = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *RotatedRhombusGridShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _RotatedRhombusShapes []*RotatedRhombusShape
	for _, _reference := range reference.RotatedRhombusShapes {
		if _instance, ok := stage.RotatedRhombusShapes_instance[_reference]; ok {
			_RotatedRhombusShapes = append(_RotatedRhombusShapes, _instance)
		}
	}
	reference.RotatedRhombusShapes = _RotatedRhombusShapes
}

func (reference *RotatedRhombusShape) GongReconstructPointersFromInstances(stage *Stage) {
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
	var _ShiftedBottomTopStartArcShapes []*ShiftedBottomTopStartArcShape
	for _, _reference := range reference.ShiftedBottomTopStartArcShapes {
		if _instance, ok := stage.ShiftedBottomTopStartArcShapes_instance[_reference]; ok {
			_ShiftedBottomTopStartArcShapes = append(_ShiftedBottomTopStartArcShapes, _instance)
		}
	}
	reference.ShiftedBottomTopStartArcShapes = _ShiftedBottomTopStartArcShapes
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
	var _ShiftedLeftStackGrowthCurveStartArcShapes []*ShiftedLeftStackGrowthCurveStartArcShape
	for _, _reference := range reference.ShiftedLeftStackGrowthCurveStartArcShapes {
		if _instance, ok := stage.ShiftedLeftStackGrowthCurveStartArcShapes_instance[_reference]; ok {
			_ShiftedLeftStackGrowthCurveStartArcShapes = append(_ShiftedLeftStackGrowthCurveStartArcShapes, _instance)
		}
	}
	reference.ShiftedLeftStackGrowthCurveStartArcShapes = _ShiftedLeftStackGrowthCurveStartArcShapes
	var _ShiftedLeftStackGrowthCurveEndArcShapes []*ShiftedLeftStackGrowthCurveEndArcShape
	for _, _reference := range reference.ShiftedLeftStackGrowthCurveEndArcShapes {
		if _instance, ok := stage.ShiftedLeftStackGrowthCurveEndArcShapes_instance[_reference]; ok {
			_ShiftedLeftStackGrowthCurveEndArcShapes = append(_ShiftedLeftStackGrowthCurveEndArcShapes, _instance)
		}
	}
	reference.ShiftedLeftStackGrowthCurveEndArcShapes = _ShiftedLeftStackGrowthCurveEndArcShapes
}

func (reference *ShiftedLeftStackOfNormalVector) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ShiftedLeftStackNormalVectors []*ShiftedLeftStackNormalVector
	for _, _reference := range reference.ShiftedLeftStackNormalVectors {
		if _instance, ok := stage.ShiftedLeftStackNormalVectors_instance[_reference]; ok {
			_ShiftedLeftStackNormalVectors = append(_ShiftedLeftStackNormalVectors, _instance)
		}
	}
	reference.ShiftedLeftStackNormalVectors = _ShiftedLeftStackNormalVectors
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
	var _StackGrowthCurve2DStartHalfwayArcShapes []*StackGrowthCurve2DStartHalfwayArcShape
	for _, _reference := range reference.StackGrowthCurve2DStartHalfwayArcShapes {
		if _instance, ok := stage.StackGrowthCurve2DStartHalfwayArcShapes_instance[_reference]; ok {
			_StackGrowthCurve2DStartHalfwayArcShapes = append(_StackGrowthCurve2DStartHalfwayArcShapes, _instance)
		}
	}
	reference.StackGrowthCurve2DStartHalfwayArcShapes = _StackGrowthCurve2DStartHalfwayArcShapes
	var _StackGrowthCurve2DEndHalfwayArcShapes []*StackGrowthCurve2DEndHalfwayArcShape
	for _, _reference := range reference.StackGrowthCurve2DEndHalfwayArcShapes {
		if _instance, ok := stage.StackGrowthCurve2DEndHalfwayArcShapes_instance[_reference]; ok {
			_StackGrowthCurve2DEndHalfwayArcShapes = append(_StackGrowthCurve2DEndHalfwayArcShapes, _instance)
		}
	}
	reference.StackGrowthCurve2DEndHalfwayArcShapes = _StackGrowthCurve2DEndHalfwayArcShapes
}

func (reference *StackOfGrowthCurve2DRibbon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _StackGrowthCurve2DRibbonStartShapes []*StackGrowthCurve2DRibbonStartShape
	for _, _reference := range reference.StackGrowthCurve2DRibbonStartShapes {
		if _instance, ok := stage.StackGrowthCurve2DRibbonStartShapes_instance[_reference]; ok {
			_StackGrowthCurve2DRibbonStartShapes = append(_StackGrowthCurve2DRibbonStartShapes, _instance)
		}
	}
	reference.StackGrowthCurve2DRibbonStartShapes = _StackGrowthCurve2DRibbonStartShapes
	var _StackGrowthCurve2DRibbonEndShapes []*StackGrowthCurve2DRibbonEndShape
	for _, _reference := range reference.StackGrowthCurve2DRibbonEndShapes {
		if _instance, ok := stage.StackGrowthCurve2DRibbonEndShapes_instance[_reference]; ok {
			_StackGrowthCurve2DRibbonEndShapes = append(_StackGrowthCurve2DRibbonEndShapes, _instance)
		}
	}
	reference.StackGrowthCurve2DRibbonEndShapes = _StackGrowthCurve2DRibbonEndShapes
}

func (reference *StackOfRotatedGrowthCurve2D) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _StackRotatedGrowthCurve2DStartArcShapes []*StackRotatedGrowthCurve2DStartArcShape
	for _, _reference := range reference.StackRotatedGrowthCurve2DStartArcShapes {
		if _instance, ok := stage.StackRotatedGrowthCurve2DStartArcShapes_instance[_reference]; ok {
			_StackRotatedGrowthCurve2DStartArcShapes = append(_StackRotatedGrowthCurve2DStartArcShapes, _instance)
		}
	}
	reference.StackRotatedGrowthCurve2DStartArcShapes = _StackRotatedGrowthCurve2DStartArcShapes
	var _StackRotatedGrowthCurve2DEndArcShapes []*StackRotatedGrowthCurve2DEndArcShape
	for _, _reference := range reference.StackRotatedGrowthCurve2DEndArcShapes {
		if _instance, ok := stage.StackRotatedGrowthCurve2DEndArcShapes_instance[_reference]; ok {
			_StackRotatedGrowthCurve2DEndArcShapes = append(_StackRotatedGrowthCurve2DEndArcShapes, _instance)
		}
	}
	reference.StackRotatedGrowthCurve2DEndArcShapes = _StackRotatedGrowthCurve2DEndArcShapes
}

func (reference *StackRotatedGrowthCurve2DEndArcShape) GongReconstructPointersFromInstances(stage *Stage) {
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
	var _StartArcShapes []*StartArcShape
	for _, _reference := range reference.StartArcShapes {
		if _instance, ok := stage.StartArcShapes_instance[_reference]; ok {
			_StartArcShapes = append(_StartArcShapes, _instance)
		}
	}
	reference.StartArcShapes = _StartArcShapes
}

func (reference *StartHalfwayArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StartHalfwayArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _StartHalfwayArcShapes []*StartHalfwayArcShape
	for _, _reference := range reference.StartHalfwayArcShapes {
		if _instance, ok := stage.StartHalfwayArcShapes_instance[_reference]; ok {
			_StartHalfwayArcShapes = append(_StartHalfwayArcShapes, _instance)
		}
	}
	reference.StartHalfwayArcShapes = _StartHalfwayArcShapes
}

func (reference *TopEndArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopEndArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _TopEndArcShapes []*TopEndArcShape
	for _, _reference := range reference.TopEndArcShapes {
		if _instance, ok := stage.TopEndArcShapes_instance[_reference]; ok {
			_TopEndArcShapes = append(_TopEndArcShapes, _instance)
		}
	}
	reference.TopEndArcShapes = _TopEndArcShapes
}

func (reference *TopEndHalfwayArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopEndHalfwayArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _TopEndHalfwayArcShapes []*TopEndHalfwayArcShape
	for _, _reference := range reference.TopEndHalfwayArcShapes {
		if _instance, ok := stage.TopEndHalfwayArcShapes_instance[_reference]; ok {
			_TopEndHalfwayArcShapes = append(_TopEndHalfwayArcShapes, _instance)
		}
	}
	reference.TopEndHalfwayArcShapes = _TopEndHalfwayArcShapes
}

func (reference *TopGrowthCurve2D) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.TopStartHalfwayArcShapeGrid; _reference != nil {
		reference.TopStartHalfwayArcShapeGrid = nil
		if _instance, ok := stage.TopStartHalfwayArcShapeGrids_instance[_reference]; ok {
			reference.TopStartHalfwayArcShapeGrid = _instance
		}
	}
	if _reference := reference.TopEndHalfwayArcShapeGrid; _reference != nil {
		reference.TopEndHalfwayArcShapeGrid = nil
		if _instance, ok := stage.TopEndHalfwayArcShapeGrids_instance[_reference]; ok {
			reference.TopEndHalfwayArcShapeGrid = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *TopMidArcVectorShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopMidArcVectorShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _TopMidArcVectorShapes []*TopMidArcVectorShape
	for _, _reference := range reference.TopMidArcVectorShapes {
		if _instance, ok := stage.TopMidArcVectorShapes_instance[_reference]; ok {
			_TopMidArcVectorShapes = append(_TopMidArcVectorShapes, _instance)
		}
	}
	reference.TopMidArcVectorShapes = _TopMidArcVectorShapes
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
	var _TopStackGrowthCurve2DStartHalfwayArcShapes []*TopStackGrowthCurve2DStartHalfwayArcShape
	for _, _reference := range reference.TopStackGrowthCurve2DStartHalfwayArcShapes {
		if _instance, ok := stage.TopStackGrowthCurve2DStartHalfwayArcShapes_instance[_reference]; ok {
			_TopStackGrowthCurve2DStartHalfwayArcShapes = append(_TopStackGrowthCurve2DStartHalfwayArcShapes, _instance)
		}
	}
	reference.TopStackGrowthCurve2DStartHalfwayArcShapes = _TopStackGrowthCurve2DStartHalfwayArcShapes
	var _TopStackGrowthCurve2DEndHalfwayArcShapes []*TopStackGrowthCurve2DEndHalfwayArcShape
	for _, _reference := range reference.TopStackGrowthCurve2DEndHalfwayArcShapes {
		if _instance, ok := stage.TopStackGrowthCurve2DEndHalfwayArcShapes_instance[_reference]; ok {
			_TopStackGrowthCurve2DEndHalfwayArcShapes = append(_TopStackGrowthCurve2DEndHalfwayArcShapes, _instance)
		}
	}
	reference.TopStackGrowthCurve2DEndHalfwayArcShapes = _TopStackGrowthCurve2DEndHalfwayArcShapes
}

func (reference *TopStackOfRotatedGrowthCurve2D) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _TopStackOfRotatedGrowthCurve2DStartArcShapes []*TopStackOfRotatedGrowthCurve2DStartArcShape
	for _, _reference := range reference.TopStackOfRotatedGrowthCurve2DStartArcShapes {
		if _instance, ok := stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_instance[_reference]; ok {
			_TopStackOfRotatedGrowthCurve2DStartArcShapes = append(_TopStackOfRotatedGrowthCurve2DStartArcShapes, _instance)
		}
	}
	reference.TopStackOfRotatedGrowthCurve2DStartArcShapes = _TopStackOfRotatedGrowthCurve2DStartArcShapes
	var _TopStackOfRotatedGrowthCurve2DEndArcShapes []*TopStackOfRotatedGrowthCurve2DEndArcShape
	for _, _reference := range reference.TopStackOfRotatedGrowthCurve2DEndArcShapes {
		if _instance, ok := stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_instance[_reference]; ok {
			_TopStackOfRotatedGrowthCurve2DEndArcShapes = append(_TopStackOfRotatedGrowthCurve2DEndArcShapes, _instance)
		}
	}
	reference.TopStackOfRotatedGrowthCurve2DEndArcShapes = _TopStackOfRotatedGrowthCurve2DEndArcShapes
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
	var _TopStartArcShapes []*TopStartArcShape
	for _, _reference := range reference.TopStartArcShapes {
		if _instance, ok := stage.TopStartArcShapes_instance[_reference]; ok {
			_TopStartArcShapes = append(_TopStartArcShapes, _instance)
		}
	}
	reference.TopStartArcShapes = _TopStartArcShapes
}

func (reference *TopStartHalfwayArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStartHalfwayArcShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _TopStartHalfwayArcShapes []*TopStartHalfwayArcShape
	for _, _reference := range reference.TopStartHalfwayArcShapes {
		if _instance, ok := stage.TopStartHalfwayArcShapes_instance[_reference]; ok {
			_TopStartHalfwayArcShapes = append(_TopStartHalfwayArcShapes, _instance)
		}
	}
	reference.TopStartHalfwayArcShapes = _TopStartHalfwayArcShapes
}

func (reference *TorusStackShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *VerticalTorusStackShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
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
	ArcNormalVectorShapesDifferent := false
	if len(arcnormalvectorshapegrid.ArcNormalVectorShapes) != len(arcnormalvectorshapegridOther.ArcNormalVectorShapes) {
		ArcNormalVectorShapesDifferent = true
	} else {
		for i := range arcnormalvectorshapegrid.ArcNormalVectorShapes {
			if (arcnormalvectorshapegrid.ArcNormalVectorShapes[i] == nil) != (arcnormalvectorshapegridOther.ArcNormalVectorShapes[i] == nil) {
				ArcNormalVectorShapesDifferent = true
				break
			} else if arcnormalvectorshapegrid.ArcNormalVectorShapes[i] != nil && arcnormalvectorshapegridOther.ArcNormalVectorShapes[i] != nil {
				// this is a pointer comparaison
				if arcnormalvectorshapegrid.ArcNormalVectorShapes[i] != arcnormalvectorshapegridOther.ArcNormalVectorShapes[i] {
					ArcNormalVectorShapesDifferent = true
					break
				}
			}
		}
	}
	if ArcNormalVectorShapesDifferent {
		ops := Diff(stage, arcnormalvectorshapegrid, arcnormalvectorshapegridOther, "ArcNormalVectorShapes", arcnormalvectorshapegridOther.ArcNormalVectorShapes, arcnormalvectorshapegrid.ArcNormalVectorShapes)
		diffs = append(diffs, ops)
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
	BaseVectorShapesDifferent := false
	if len(basevectorshapegrid.BaseVectorShapes) != len(basevectorshapegridOther.BaseVectorShapes) {
		BaseVectorShapesDifferent = true
	} else {
		for i := range basevectorshapegrid.BaseVectorShapes {
			if (basevectorshapegrid.BaseVectorShapes[i] == nil) != (basevectorshapegridOther.BaseVectorShapes[i] == nil) {
				BaseVectorShapesDifferent = true
				break
			} else if basevectorshapegrid.BaseVectorShapes[i] != nil && basevectorshapegridOther.BaseVectorShapes[i] != nil {
				// this is a pointer comparaison
				if basevectorshapegrid.BaseVectorShapes[i] != basevectorshapegridOther.BaseVectorShapes[i] {
					BaseVectorShapesDifferent = true
					break
				}
			}
		}
	}
	if BaseVectorShapesDifferent {
		ops := Diff(stage, basevectorshapegrid, basevectorshapegridOther, "BaseVectorShapes", basevectorshapegridOther.BaseVectorShapes, basevectorshapegrid.BaseVectorShapes)
		diffs = append(diffs, ops)
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
	EndArcShapesDifferent := false
	if len(endarcshapegrid.EndArcShapes) != len(endarcshapegridOther.EndArcShapes) {
		EndArcShapesDifferent = true
	} else {
		for i := range endarcshapegrid.EndArcShapes {
			if (endarcshapegrid.EndArcShapes[i] == nil) != (endarcshapegridOther.EndArcShapes[i] == nil) {
				EndArcShapesDifferent = true
				break
			} else if endarcshapegrid.EndArcShapes[i] != nil && endarcshapegridOther.EndArcShapes[i] != nil {
				// this is a pointer comparaison
				if endarcshapegrid.EndArcShapes[i] != endarcshapegridOther.EndArcShapes[i] {
					EndArcShapesDifferent = true
					break
				}
			}
		}
	}
	if EndArcShapesDifferent {
		ops := Diff(stage, endarcshapegrid, endarcshapegridOther, "EndArcShapes", endarcshapegridOther.EndArcShapes, endarcshapegrid.EndArcShapes)
		diffs = append(diffs, ops)
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
	EndHalfwayArcShapesDifferent := false
	if len(endhalfwayarcshapegrid.EndHalfwayArcShapes) != len(endhalfwayarcshapegridOther.EndHalfwayArcShapes) {
		EndHalfwayArcShapesDifferent = true
	} else {
		for i := range endhalfwayarcshapegrid.EndHalfwayArcShapes {
			if (endhalfwayarcshapegrid.EndHalfwayArcShapes[i] == nil) != (endhalfwayarcshapegridOther.EndHalfwayArcShapes[i] == nil) {
				EndHalfwayArcShapesDifferent = true
				break
			} else if endhalfwayarcshapegrid.EndHalfwayArcShapes[i] != nil && endhalfwayarcshapegridOther.EndHalfwayArcShapes[i] != nil {
				// this is a pointer comparaison
				if endhalfwayarcshapegrid.EndHalfwayArcShapes[i] != endhalfwayarcshapegridOther.EndHalfwayArcShapes[i] {
					EndHalfwayArcShapesDifferent = true
					break
				}
			}
		}
	}
	if EndHalfwayArcShapesDifferent {
		ops := Diff(stage, endhalfwayarcshapegrid, endhalfwayarcshapegridOther, "EndHalfwayArcShapes", endhalfwayarcshapegridOther.EndHalfwayArcShapes, endhalfwayarcshapegrid.EndHalfwayArcShapes)
		diffs = append(diffs, ops)
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
	if (growthcurve2d.StartHalfwayArcShapeGrid == nil) != (growthcurve2dOther.StartHalfwayArcShapeGrid == nil) {
		diffs = append(diffs, growthcurve2d.GongMarshallField(stage, "StartHalfwayArcShapeGrid"))
	} else if growthcurve2d.StartHalfwayArcShapeGrid != nil && growthcurve2dOther.StartHalfwayArcShapeGrid != nil {
		if growthcurve2d.StartHalfwayArcShapeGrid != growthcurve2dOther.StartHalfwayArcShapeGrid {
			diffs = append(diffs, growthcurve2d.GongMarshallField(stage, "StartHalfwayArcShapeGrid"))
		}
	}
	if (growthcurve2d.EndHalfwayArcShapeGrid == nil) != (growthcurve2dOther.EndHalfwayArcShapeGrid == nil) {
		diffs = append(diffs, growthcurve2d.GongMarshallField(stage, "EndHalfwayArcShapeGrid"))
	} else if growthcurve2d.EndHalfwayArcShapeGrid != nil && growthcurve2dOther.EndHalfwayArcShapeGrid != nil {
		if growthcurve2d.EndHalfwayArcShapeGrid != growthcurve2dOther.EndHalfwayArcShapeGrid {
			diffs = append(diffs, growthcurve2d.GongMarshallField(stage, "EndHalfwayArcShapeGrid"))
		}
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
	GrowthCurveRhombusShapesDifferent := false
	if len(growthcurverhombusgridshape.GrowthCurveRhombusShapes) != len(growthcurverhombusgridshapeOther.GrowthCurveRhombusShapes) {
		GrowthCurveRhombusShapesDifferent = true
	} else {
		for i := range growthcurverhombusgridshape.GrowthCurveRhombusShapes {
			if (growthcurverhombusgridshape.GrowthCurveRhombusShapes[i] == nil) != (growthcurverhombusgridshapeOther.GrowthCurveRhombusShapes[i] == nil) {
				GrowthCurveRhombusShapesDifferent = true
				break
			} else if growthcurverhombusgridshape.GrowthCurveRhombusShapes[i] != nil && growthcurverhombusgridshapeOther.GrowthCurveRhombusShapes[i] != nil {
				// this is a pointer comparaison
				if growthcurverhombusgridshape.GrowthCurveRhombusShapes[i] != growthcurverhombusgridshapeOther.GrowthCurveRhombusShapes[i] {
					GrowthCurveRhombusShapesDifferent = true
					break
				}
			}
		}
	}
	if GrowthCurveRhombusShapesDifferent {
		ops := Diff(stage, growthcurverhombusgridshape, growthcurverhombusgridshapeOther, "GrowthCurveRhombusShapes", growthcurverhombusgridshapeOther.GrowthCurveRhombusShapes, growthcurverhombusgridshape.GrowthCurveRhombusShapes)
		diffs = append(diffs, ops)
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
	InitialRhombusShapesDifferent := false
	if len(initialrhombusgridshape.InitialRhombusShapes) != len(initialrhombusgridshapeOther.InitialRhombusShapes) {
		InitialRhombusShapesDifferent = true
	} else {
		for i := range initialrhombusgridshape.InitialRhombusShapes {
			if (initialrhombusgridshape.InitialRhombusShapes[i] == nil) != (initialrhombusgridshapeOther.InitialRhombusShapes[i] == nil) {
				InitialRhombusShapesDifferent = true
				break
			} else if initialrhombusgridshape.InitialRhombusShapes[i] != nil && initialrhombusgridshapeOther.InitialRhombusShapes[i] != nil {
				// this is a pointer comparaison
				if initialrhombusgridshape.InitialRhombusShapes[i] != initialrhombusgridshapeOther.InitialRhombusShapes[i] {
					InitialRhombusShapesDifferent = true
					break
				}
			}
		}
	}
	if InitialRhombusShapesDifferent {
		ops := Diff(stage, initialrhombusgridshape, initialrhombusgridshapeOther, "InitialRhombusShapes", initialrhombusgridshapeOther.InitialRhombusShapes, initialrhombusgridshape.InitialRhombusShapes)
		diffs = append(diffs, ops)
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
	MidArcVectorShapesDifferent := false
	if len(midarcvectorshapegrid.MidArcVectorShapes) != len(midarcvectorshapegridOther.MidArcVectorShapes) {
		MidArcVectorShapesDifferent = true
	} else {
		for i := range midarcvectorshapegrid.MidArcVectorShapes {
			if (midarcvectorshapegrid.MidArcVectorShapes[i] == nil) != (midarcvectorshapegridOther.MidArcVectorShapes[i] == nil) {
				MidArcVectorShapesDifferent = true
				break
			} else if midarcvectorshapegrid.MidArcVectorShapes[i] != nil && midarcvectorshapegridOther.MidArcVectorShapes[i] != nil {
				// this is a pointer comparaison
				if midarcvectorshapegrid.MidArcVectorShapes[i] != midarcvectorshapegridOther.MidArcVectorShapes[i] {
					MidArcVectorShapesDifferent = true
					break
				}
			}
		}
	}
	if MidArcVectorShapesDifferent {
		ops := Diff(stage, midarcvectorshapegrid, midarcvectorshapegridOther, "MidArcVectorShapes", midarcvectorshapegridOther.MidArcVectorShapes, midarcvectorshapegrid.MidArcVectorShapes)
		diffs = append(diffs, ops)
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
	PerpendicularVectorsDifferent := false
	if len(perpendicularvectorgrid.PerpendicularVectors) != len(perpendicularvectorgridOther.PerpendicularVectors) {
		PerpendicularVectorsDifferent = true
	} else {
		for i := range perpendicularvectorgrid.PerpendicularVectors {
			if (perpendicularvectorgrid.PerpendicularVectors[i] == nil) != (perpendicularvectorgridOther.PerpendicularVectors[i] == nil) {
				PerpendicularVectorsDifferent = true
				break
			} else if perpendicularvectorgrid.PerpendicularVectors[i] != nil && perpendicularvectorgridOther.PerpendicularVectors[i] != nil {
				// this is a pointer comparaison
				if perpendicularvectorgrid.PerpendicularVectors[i] != perpendicularvectorgridOther.PerpendicularVectors[i] {
					PerpendicularVectorsDifferent = true
					break
				}
			}
		}
	}
	if PerpendicularVectorsDifferent {
		ops := Diff(stage, perpendicularvectorgrid, perpendicularvectorgridOther, "PerpendicularVectors", perpendicularvectorgridOther.PerpendicularVectors, perpendicularvectorgrid.PerpendicularVectors)
		diffs = append(diffs, ops)
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
	PerpendicularVectorHalfwaysDifferent := false
	if len(perpendicularvectorgridhalfway.PerpendicularVectorHalfways) != len(perpendicularvectorgridhalfwayOther.PerpendicularVectorHalfways) {
		PerpendicularVectorHalfwaysDifferent = true
	} else {
		for i := range perpendicularvectorgridhalfway.PerpendicularVectorHalfways {
			if (perpendicularvectorgridhalfway.PerpendicularVectorHalfways[i] == nil) != (perpendicularvectorgridhalfwayOther.PerpendicularVectorHalfways[i] == nil) {
				PerpendicularVectorHalfwaysDifferent = true
				break
			} else if perpendicularvectorgridhalfway.PerpendicularVectorHalfways[i] != nil && perpendicularvectorgridhalfwayOther.PerpendicularVectorHalfways[i] != nil {
				// this is a pointer comparaison
				if perpendicularvectorgridhalfway.PerpendicularVectorHalfways[i] != perpendicularvectorgridhalfwayOther.PerpendicularVectorHalfways[i] {
					PerpendicularVectorHalfwaysDifferent = true
					break
				}
			}
		}
	}
	if PerpendicularVectorHalfwaysDifferent {
		ops := Diff(stage, perpendicularvectorgridhalfway, perpendicularvectorgridhalfwayOther, "PerpendicularVectorHalfways", perpendicularvectorgridhalfwayOther.PerpendicularVectorHalfways, perpendicularvectorgridhalfway.PerpendicularVectorHalfways)
		diffs = append(diffs, ops)
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
func (plant *Plant) GongDiff(stage *Stage, plantOther *Plant) (diffs []string) {
	// insertion point for field diffs
	if plant.Name != plantOther.Name {
		diffs = append(diffs, plant.GongMarshallField(stage, "Name"))
	}
	if plant.N != plantOther.N {
		diffs = append(diffs, plant.GongMarshallField(stage, "N"))
	}
	if plant.M != plantOther.M {
		diffs = append(diffs, plant.GongMarshallField(stage, "M"))
	}
	if plant.StackHeight != plantOther.StackHeight {
		diffs = append(diffs, plant.GongMarshallField(stage, "StackHeight"))
	}
	if plant.RhombusInsideAngle != plantOther.RhombusInsideAngle {
		diffs = append(diffs, plant.GongMarshallField(stage, "RhombusInsideAngle"))
	}
	if plant.RelativeVerticalThickness != plantOther.RelativeVerticalThickness {
		diffs = append(diffs, plant.GongMarshallField(stage, "RelativeVerticalThickness"))
	}
	if plant.RelativeRadialThickness != plantOther.RelativeRadialThickness {
		diffs = append(diffs, plant.GongMarshallField(stage, "RelativeRadialThickness"))
	}
	if plant.RhombusSideLength != plantOther.RhombusSideLength {
		diffs = append(diffs, plant.GongMarshallField(stage, "RhombusSideLength"))
	}
	if plant.RelativeCuttedStackFloorHeight != plantOther.RelativeCuttedStackFloorHeight {
		diffs = append(diffs, plant.GongMarshallField(stage, "RelativeCuttedStackFloorHeight"))
	}
	if plant.RelativeRotatedTorusSeparation != plantOther.RelativeRotatedTorusSeparation {
		diffs = append(diffs, plant.GongMarshallField(stage, "RelativeRotatedTorusSeparation"))
	}
	if plant.RotationRatio != plantOther.RotationRatio {
		diffs = append(diffs, plant.GongMarshallField(stage, "RotationRatio"))
	}
	if plant.ComputedPrefix != plantOther.ComputedPrefix {
		diffs = append(diffs, plant.GongMarshallField(stage, "ComputedPrefix"))
	}
	if plant.IsExpanded != plantOther.IsExpanded {
		diffs = append(diffs, plant.GongMarshallField(stage, "IsExpanded"))
	}
	if plant.IsSelected != plantOther.IsSelected {
		diffs = append(diffs, plant.GongMarshallField(stage, "IsSelected"))
	}
	if plant.IsPlantDiagramsNodeExpanded != plantOther.IsPlantDiagramsNodeExpanded {
		diffs = append(diffs, plant.GongMarshallField(stage, "IsPlantDiagramsNodeExpanded"))
	}
	PlantDiagramsDifferent := false
	if len(plant.PlantDiagrams) != len(plantOther.PlantDiagrams) {
		PlantDiagramsDifferent = true
	} else {
		for i := range plant.PlantDiagrams {
			if (plant.PlantDiagrams[i] == nil) != (plantOther.PlantDiagrams[i] == nil) {
				PlantDiagramsDifferent = true
				break
			} else if plant.PlantDiagrams[i] != nil && plantOther.PlantDiagrams[i] != nil {
				// this is a pointer comparaison
				if plant.PlantDiagrams[i] != plantOther.PlantDiagrams[i] {
					PlantDiagramsDifferent = true
					break
				}
			}
		}
	}
	if PlantDiagramsDifferent {
		ops := Diff(stage, plant, plantOther, "PlantDiagrams", plantOther.PlantDiagrams, plant.PlantDiagrams)
		diffs = append(diffs, ops)
	}
	if (plant.AxesShape == nil) != (plantOther.AxesShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "AxesShape"))
	} else if plant.AxesShape != nil && plantOther.AxesShape != nil {
		if plant.AxesShape != plantOther.AxesShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "AxesShape"))
		}
	}
	if (plant.RhombusStuff == nil) != (plantOther.RhombusStuff == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "RhombusStuff"))
	} else if plant.RhombusStuff != nil && plantOther.RhombusStuff != nil {
		if plant.RhombusStuff != plantOther.RhombusStuff {
			diffs = append(diffs, plant.GongMarshallField(stage, "RhombusStuff"))
		}
	}
	if (plant.GrowthVectorShape == nil) != (plantOther.GrowthVectorShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "GrowthVectorShape"))
	} else if plant.GrowthVectorShape != nil && plantOther.GrowthVectorShape != nil {
		if plant.GrowthVectorShape != plantOther.GrowthVectorShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "GrowthVectorShape"))
		}
	}
	if (plant.PerpendicularVectorGrid == nil) != (plantOther.PerpendicularVectorGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "PerpendicularVectorGrid"))
	} else if plant.PerpendicularVectorGrid != nil && plantOther.PerpendicularVectorGrid != nil {
		if plant.PerpendicularVectorGrid != plantOther.PerpendicularVectorGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "PerpendicularVectorGrid"))
		}
	}
	if (plant.PerpendicularVectorGridHalfway == nil) != (plantOther.PerpendicularVectorGridHalfway == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "PerpendicularVectorGridHalfway"))
	} else if plant.PerpendicularVectorGridHalfway != nil && plantOther.PerpendicularVectorGridHalfway != nil {
		if plant.PerpendicularVectorGridHalfway != plantOther.PerpendicularVectorGridHalfway {
			diffs = append(diffs, plant.GongMarshallField(stage, "PerpendicularVectorGridHalfway"))
		}
	}
	if (plant.BaseVectorShapeGrid == nil) != (plantOther.BaseVectorShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "BaseVectorShapeGrid"))
	} else if plant.BaseVectorShapeGrid != nil && plantOther.BaseVectorShapeGrid != nil {
		if plant.BaseVectorShapeGrid != plantOther.BaseVectorShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "BaseVectorShapeGrid"))
		}
	}
	if (plant.ArcNormalVectorShapeGrid == nil) != (plantOther.ArcNormalVectorShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "ArcNormalVectorShapeGrid"))
	} else if plant.ArcNormalVectorShapeGrid != nil && plantOther.ArcNormalVectorShapeGrid != nil {
		if plant.ArcNormalVectorShapeGrid != plantOther.ArcNormalVectorShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "ArcNormalVectorShapeGrid"))
		}
	}
	if (plant.StartArcShapeGrid == nil) != (plantOther.StartArcShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "StartArcShapeGrid"))
	} else if plant.StartArcShapeGrid != nil && plantOther.StartArcShapeGrid != nil {
		if plant.StartArcShapeGrid != plantOther.StartArcShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "StartArcShapeGrid"))
		}
	}
	if (plant.TopStartArcShapeGrid == nil) != (plantOther.TopStartArcShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "TopStartArcShapeGrid"))
	} else if plant.TopStartArcShapeGrid != nil && plantOther.TopStartArcShapeGrid != nil {
		if plant.TopStartArcShapeGrid != plantOther.TopStartArcShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "TopStartArcShapeGrid"))
		}
	}
	if (plant.EndArcShapeGrid == nil) != (plantOther.EndArcShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "EndArcShapeGrid"))
	} else if plant.EndArcShapeGrid != nil && plantOther.EndArcShapeGrid != nil {
		if plant.EndArcShapeGrid != plantOther.EndArcShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "EndArcShapeGrid"))
		}
	}
	if (plant.TopEndArcShapeGrid == nil) != (plantOther.TopEndArcShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "TopEndArcShapeGrid"))
	} else if plant.TopEndArcShapeGrid != nil && plantOther.TopEndArcShapeGrid != nil {
		if plant.TopEndArcShapeGrid != plantOther.TopEndArcShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "TopEndArcShapeGrid"))
		}
	}
	if (plant.ShiftedBottomTopStartArcShapeGrid == nil) != (plantOther.ShiftedBottomTopStartArcShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "ShiftedBottomTopStartArcShapeGrid"))
	} else if plant.ShiftedBottomTopStartArcShapeGrid != nil && plantOther.ShiftedBottomTopStartArcShapeGrid != nil {
		if plant.ShiftedBottomTopStartArcShapeGrid != plantOther.ShiftedBottomTopStartArcShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "ShiftedBottomTopStartArcShapeGrid"))
		}
	}
	if (plant.MidArcVectorShapeGrid == nil) != (plantOther.MidArcVectorShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "MidArcVectorShapeGrid"))
	} else if plant.MidArcVectorShapeGrid != nil && plantOther.MidArcVectorShapeGrid != nil {
		if plant.MidArcVectorShapeGrid != plantOther.MidArcVectorShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "MidArcVectorShapeGrid"))
		}
	}
	if (plant.TopMidArcVectorShapeGrid == nil) != (plantOther.TopMidArcVectorShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "TopMidArcVectorShapeGrid"))
	} else if plant.TopMidArcVectorShapeGrid != nil && plantOther.TopMidArcVectorShapeGrid != nil {
		if plant.TopMidArcVectorShapeGrid != plantOther.TopMidArcVectorShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "TopMidArcVectorShapeGrid"))
		}
	}
	if (plant.StartHalfwayArcShapeGrid == nil) != (plantOther.StartHalfwayArcShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "StartHalfwayArcShapeGrid"))
	} else if plant.StartHalfwayArcShapeGrid != nil && plantOther.StartHalfwayArcShapeGrid != nil {
		if plant.StartHalfwayArcShapeGrid != plantOther.StartHalfwayArcShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "StartHalfwayArcShapeGrid"))
		}
	}
	if (plant.TopStartHalfwayArcShapeGrid == nil) != (plantOther.TopStartHalfwayArcShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "TopStartHalfwayArcShapeGrid"))
	} else if plant.TopStartHalfwayArcShapeGrid != nil && plantOther.TopStartHalfwayArcShapeGrid != nil {
		if plant.TopStartHalfwayArcShapeGrid != plantOther.TopStartHalfwayArcShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "TopStartHalfwayArcShapeGrid"))
		}
	}
	if (plant.EndHalfwayArcShapeGrid == nil) != (plantOther.EndHalfwayArcShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "EndHalfwayArcShapeGrid"))
	} else if plant.EndHalfwayArcShapeGrid != nil && plantOther.EndHalfwayArcShapeGrid != nil {
		if plant.EndHalfwayArcShapeGrid != plantOther.EndHalfwayArcShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "EndHalfwayArcShapeGrid"))
		}
	}
	if (plant.TopEndHalfwayArcShapeGrid == nil) != (plantOther.TopEndHalfwayArcShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "TopEndHalfwayArcShapeGrid"))
	} else if plant.TopEndHalfwayArcShapeGrid != nil && plantOther.TopEndHalfwayArcShapeGrid != nil {
		if plant.TopEndHalfwayArcShapeGrid != plantOther.TopEndHalfwayArcShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "TopEndHalfwayArcShapeGrid"))
		}
	}
	if (plant.StackOfRotatedGrowthCurve2D == nil) != (plantOther.StackOfRotatedGrowthCurve2D == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "StackOfRotatedGrowthCurve2D"))
	} else if plant.StackOfRotatedGrowthCurve2D != nil && plantOther.StackOfRotatedGrowthCurve2D != nil {
		if plant.StackOfRotatedGrowthCurve2D != plantOther.StackOfRotatedGrowthCurve2D {
			diffs = append(diffs, plant.GongMarshallField(stage, "StackOfRotatedGrowthCurve2D"))
		}
	}
	if (plant.TopStackOfRotatedGrowthCurve2D == nil) != (plantOther.TopStackOfRotatedGrowthCurve2D == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "TopStackOfRotatedGrowthCurve2D"))
	} else if plant.TopStackOfRotatedGrowthCurve2D != nil && plantOther.TopStackOfRotatedGrowthCurve2D != nil {
		if plant.TopStackOfRotatedGrowthCurve2D != plantOther.TopStackOfRotatedGrowthCurve2D {
			diffs = append(diffs, plant.GongMarshallField(stage, "TopStackOfRotatedGrowthCurve2D"))
		}
	}
	if (plant.GrowthCurve2D == nil) != (plantOther.GrowthCurve2D == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "GrowthCurve2D"))
	} else if plant.GrowthCurve2D != nil && plantOther.GrowthCurve2D != nil {
		if plant.GrowthCurve2D != plantOther.GrowthCurve2D {
			diffs = append(diffs, plant.GongMarshallField(stage, "GrowthCurve2D"))
		}
	}
	if (plant.TopGrowthCurve2D == nil) != (plantOther.TopGrowthCurve2D == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "TopGrowthCurve2D"))
	} else if plant.TopGrowthCurve2D != nil && plantOther.TopGrowthCurve2D != nil {
		if plant.TopGrowthCurve2D != plantOther.TopGrowthCurve2D {
			diffs = append(diffs, plant.GongMarshallField(stage, "TopGrowthCurve2D"))
		}
	}
	if (plant.StackOfGrowthCurve2D == nil) != (plantOther.StackOfGrowthCurve2D == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "StackOfGrowthCurve2D"))
	} else if plant.StackOfGrowthCurve2D != nil && plantOther.StackOfGrowthCurve2D != nil {
		if plant.StackOfGrowthCurve2D != plantOther.StackOfGrowthCurve2D {
			diffs = append(diffs, plant.GongMarshallField(stage, "StackOfGrowthCurve2D"))
		}
	}
	if (plant.TopStackOfGrowthCurve2D == nil) != (plantOther.TopStackOfGrowthCurve2D == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "TopStackOfGrowthCurve2D"))
	} else if plant.TopStackOfGrowthCurve2D != nil && plantOther.TopStackOfGrowthCurve2D != nil {
		if plant.TopStackOfGrowthCurve2D != plantOther.TopStackOfGrowthCurve2D {
			diffs = append(diffs, plant.GongMarshallField(stage, "TopStackOfGrowthCurve2D"))
		}
	}
	if (plant.StackOfGrowthCurve2DRibbon == nil) != (plantOther.StackOfGrowthCurve2DRibbon == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "StackOfGrowthCurve2DRibbon"))
	} else if plant.StackOfGrowthCurve2DRibbon != nil && plantOther.StackOfGrowthCurve2DRibbon != nil {
		if plant.StackOfGrowthCurve2DRibbon != plantOther.StackOfGrowthCurve2DRibbon {
			diffs = append(diffs, plant.GongMarshallField(stage, "StackOfGrowthCurve2DRibbon"))
		}
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
func (plantdiagram *PlantDiagram) GongDiff(stage *Stage, plantdiagramOther *PlantDiagram) (diffs []string) {
	// insertion point for field diffs
	if plantdiagram.Name != plantdiagramOther.Name {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "Name"))
	}
	if plantdiagram.OriginX != plantdiagramOther.OriginX {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "OriginX"))
	}
	if plantdiagram.OriginY != plantdiagramOther.OriginY {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "OriginY"))
	}
	if plantdiagram.IsRhombusNodesExpanded != plantdiagramOther.IsRhombusNodesExpanded {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsRhombusNodesExpanded"))
	}
	if plantdiagram.IsArcNodesExpanded != plantdiagramOther.IsArcNodesExpanded {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsArcNodesExpanded"))
	}
	if plantdiagram.IsHiddenAxesShape != plantdiagramOther.IsHiddenAxesShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenAxesShape"))
	}
	if plantdiagram.IsHiddenReferenceRhombus != plantdiagramOther.IsHiddenReferenceRhombus {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenReferenceRhombus"))
	}
	if plantdiagram.IsHiddenPlantCircumferenceShape != plantdiagramOther.IsHiddenPlantCircumferenceShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenPlantCircumferenceShape"))
	}
	if plantdiagram.IsHiddenGridPathShape != plantdiagramOther.IsHiddenGridPathShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenGridPathShape"))
	}
	if plantdiagram.IsHiddenRhombusGridShape != plantdiagramOther.IsHiddenRhombusGridShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenRhombusGridShape"))
	}
	if plantdiagram.IsHiddenExplanationTextShape != plantdiagramOther.IsHiddenExplanationTextShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenExplanationTextShape"))
	}
	if plantdiagram.IsHiddenRotatedReferenceRhombus != plantdiagramOther.IsHiddenRotatedReferenceRhombus {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenRotatedReferenceRhombus"))
	}
	if plantdiagram.IsHiddenRotatedPlantCircumferenceShape != plantdiagramOther.IsHiddenRotatedPlantCircumferenceShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenRotatedPlantCircumferenceShape"))
	}
	if plantdiagram.IsHiddenRotatedGridPathShape != plantdiagramOther.IsHiddenRotatedGridPathShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenRotatedGridPathShape"))
	}
	if plantdiagram.IsHiddenRotatedRhombusGridShape != plantdiagramOther.IsHiddenRotatedRhombusGridShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenRotatedRhombusGridShape"))
	}
	if plantdiagram.IsHiddenGrowthPathRhombusGridShape != plantdiagramOther.IsHiddenGrowthPathRhombusGridShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenGrowthPathRhombusGridShape"))
	}
	if plantdiagram.IsHiddenGrowthVectorShape != plantdiagramOther.IsHiddenGrowthVectorShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenGrowthVectorShape"))
	}
	if plantdiagram.IsHiddenPerpendicularVectorGrid != plantdiagramOther.IsHiddenPerpendicularVectorGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenPerpendicularVectorGrid"))
	}
	if plantdiagram.IsHiddenPerpendicularVectorGridHalfway != plantdiagramOther.IsHiddenPerpendicularVectorGridHalfway {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenPerpendicularVectorGridHalfway"))
	}
	if plantdiagram.IsHiddenBaseVectorShapeGrid != plantdiagramOther.IsHiddenBaseVectorShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenBaseVectorShapeGrid"))
	}
	if plantdiagram.IsHiddenArcNormalVectorShapeGrid != plantdiagramOther.IsHiddenArcNormalVectorShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenArcNormalVectorShapeGrid"))
	}
	if plantdiagram.IsHiddenStartArcShapeGrid != plantdiagramOther.IsHiddenStartArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenStartArcShapeGrid"))
	}
	if plantdiagram.IsHiddenTopStartArcShapeGrid != plantdiagramOther.IsHiddenTopStartArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenTopStartArcShapeGrid"))
	}
	if plantdiagram.IsHiddenShiftedBottomTopStartArcShapeGrid != plantdiagramOther.IsHiddenShiftedBottomTopStartArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenShiftedBottomTopStartArcShapeGrid"))
	}
	if plantdiagram.IsHiddenMidArcVectorShapeGrid != plantdiagramOther.IsHiddenMidArcVectorShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenMidArcVectorShapeGrid"))
	}
	if plantdiagram.IsHiddenTopMidArcVectorShapeGrid != plantdiagramOther.IsHiddenTopMidArcVectorShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenTopMidArcVectorShapeGrid"))
	}
	if plantdiagram.IsHiddenStartHalfwayArcShapeGrid != plantdiagramOther.IsHiddenStartHalfwayArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenStartHalfwayArcShapeGrid"))
	}
	if plantdiagram.IsHiddenTopStartHalfwayArcShapeGrid != plantdiagramOther.IsHiddenTopStartHalfwayArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenTopStartHalfwayArcShapeGrid"))
	}
	if plantdiagram.IsHiddenEndHalfwayArcShapeGrid != plantdiagramOther.IsHiddenEndHalfwayArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenEndHalfwayArcShapeGrid"))
	}
	if plantdiagram.IsHiddenTopEndHalfwayArcShapeGrid != plantdiagramOther.IsHiddenTopEndHalfwayArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenTopEndHalfwayArcShapeGrid"))
	}
	if plantdiagram.IsHiddenEndArcShapeGrid != plantdiagramOther.IsHiddenEndArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenEndArcShapeGrid"))
	}
	if plantdiagram.IsHiddenTopEndArcShapeGrid != plantdiagramOther.IsHiddenTopEndArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenTopEndArcShapeGrid"))
	}
	if plantdiagram.IsHiddenBottomStartArcShapeGrid != plantdiagramOther.IsHiddenBottomStartArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenBottomStartArcShapeGrid"))
	}
	if plantdiagram.IsHiddenBottomEndArcShapeGrid != plantdiagramOther.IsHiddenBottomEndArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenBottomEndArcShapeGrid"))
	}
	if plantdiagram.IsHiddenStackOfGrowthCurve != plantdiagramOther.IsHiddenStackOfGrowthCurve {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurve"))
	}
	if plantdiagram.IsHiddenTopStackOfGrowthCurve != plantdiagramOther.IsHiddenTopStackOfGrowthCurve {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenTopStackOfGrowthCurve"))
	}
	if plantdiagram.IsHiddenBottomStackOfGrowthCurve != plantdiagramOther.IsHiddenBottomStackOfGrowthCurve {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenBottomStackOfGrowthCurve"))
	}
	if plantdiagram.IsHiddenShiftedLeftStackOfGrowthCurve != plantdiagramOther.IsHiddenShiftedLeftStackOfGrowthCurve {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenShiftedLeftStackOfGrowthCurve"))
	}
	if plantdiagram.IsHiddenShiftedLeftStackOfNormalVector != plantdiagramOther.IsHiddenShiftedLeftStackOfNormalVector {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenShiftedLeftStackOfNormalVector"))
	}
	if plantdiagram.IsHiddenGrowthCurve2D != plantdiagramOther.IsHiddenGrowthCurve2D {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenGrowthCurve2D"))
	}
	if plantdiagram.IsHiddenTopGrowthCurve2D != plantdiagramOther.IsHiddenTopGrowthCurve2D {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenTopGrowthCurve2D"))
	}
	if plantdiagram.IsHiddenStackOfGrowthCurve2D != plantdiagramOther.IsHiddenStackOfGrowthCurve2D {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurve2D"))
	}
	if plantdiagram.IsHiddenTopStackOfGrowthCurve2D != plantdiagramOther.IsHiddenTopStackOfGrowthCurve2D {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenTopStackOfGrowthCurve2D"))
	}
	if plantdiagram.IsHiddenStackOfGrowthCurve2DRibbon != plantdiagramOther.IsHiddenStackOfGrowthCurve2DRibbon {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurve2DRibbon"))
	}
	if plantdiagram.IsHiddenTorusStackShape != plantdiagramOther.IsHiddenTorusStackShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenTorusStackShape"))
	}
	if plantdiagram.IsHiddenVerticalTorusStackShape != plantdiagramOther.IsHiddenVerticalTorusStackShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenVerticalTorusStackShape"))
	}
	if plantdiagram.IsChecked != plantdiagramOther.IsChecked {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsChecked"))
	}
	if plantdiagram.ComputedPrefix != plantdiagramOther.ComputedPrefix {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if plantdiagram.IsExpanded != plantdiagramOther.IsExpanded {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsExpanded"))
	}
	if (plantdiagram.Rendered3DShape == nil) != (plantdiagramOther.Rendered3DShape == nil) {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "Rendered3DShape"))
	} else if plantdiagram.Rendered3DShape != nil && plantdiagramOther.Rendered3DShape != nil {
		if plantdiagram.Rendered3DShape != plantdiagramOther.Rendered3DShape {
			diffs = append(diffs, plantdiagram.GongMarshallField(stage, "Rendered3DShape"))
		}
	}
	if (plantdiagram.TorusStackShape == nil) != (plantdiagramOther.TorusStackShape == nil) {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "TorusStackShape"))
	} else if plantdiagram.TorusStackShape != nil && plantdiagramOther.TorusStackShape != nil {
		if plantdiagram.TorusStackShape != plantdiagramOther.TorusStackShape {
			diffs = append(diffs, plantdiagram.GongMarshallField(stage, "TorusStackShape"))
		}
	}
	if (plantdiagram.VerticalTorusStackShape == nil) != (plantdiagramOther.VerticalTorusStackShape == nil) {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "VerticalTorusStackShape"))
	} else if plantdiagram.VerticalTorusStackShape != nil && plantdiagramOther.VerticalTorusStackShape != nil {
		if plantdiagram.VerticalTorusStackShape != plantdiagramOther.VerticalTorusStackShape {
			diffs = append(diffs, plantdiagram.GongMarshallField(stage, "VerticalTorusStackShape"))
		}
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
	if (rhombusstuff.ReferenceRhombus == nil) != (rhombusstuffOther.ReferenceRhombus == nil) {
		diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "ReferenceRhombus"))
	} else if rhombusstuff.ReferenceRhombus != nil && rhombusstuffOther.ReferenceRhombus != nil {
		if rhombusstuff.ReferenceRhombus != rhombusstuffOther.ReferenceRhombus {
			diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "ReferenceRhombus"))
		}
	}
	if (rhombusstuff.PlantCircumferenceShape == nil) != (rhombusstuffOther.PlantCircumferenceShape == nil) {
		diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "PlantCircumferenceShape"))
	} else if rhombusstuff.PlantCircumferenceShape != nil && rhombusstuffOther.PlantCircumferenceShape != nil {
		if rhombusstuff.PlantCircumferenceShape != rhombusstuffOther.PlantCircumferenceShape {
			diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "PlantCircumferenceShape"))
		}
	}
	if (rhombusstuff.GridPathShape == nil) != (rhombusstuffOther.GridPathShape == nil) {
		diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "GridPathShape"))
	} else if rhombusstuff.GridPathShape != nil && rhombusstuffOther.GridPathShape != nil {
		if rhombusstuff.GridPathShape != rhombusstuffOther.GridPathShape {
			diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "GridPathShape"))
		}
	}
	if (rhombusstuff.InitialRhombusGridShape == nil) != (rhombusstuffOther.InitialRhombusGridShape == nil) {
		diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "InitialRhombusGridShape"))
	} else if rhombusstuff.InitialRhombusGridShape != nil && rhombusstuffOther.InitialRhombusGridShape != nil {
		if rhombusstuff.InitialRhombusGridShape != rhombusstuffOther.InitialRhombusGridShape {
			diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "InitialRhombusGridShape"))
		}
	}
	if (rhombusstuff.ExplanationTextShape == nil) != (rhombusstuffOther.ExplanationTextShape == nil) {
		diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "ExplanationTextShape"))
	} else if rhombusstuff.ExplanationTextShape != nil && rhombusstuffOther.ExplanationTextShape != nil {
		if rhombusstuff.ExplanationTextShape != rhombusstuffOther.ExplanationTextShape {
			diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "ExplanationTextShape"))
		}
	}
	if (rhombusstuff.RotatedReferenceRhombus == nil) != (rhombusstuffOther.RotatedReferenceRhombus == nil) {
		diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "RotatedReferenceRhombus"))
	} else if rhombusstuff.RotatedReferenceRhombus != nil && rhombusstuffOther.RotatedReferenceRhombus != nil {
		if rhombusstuff.RotatedReferenceRhombus != rhombusstuffOther.RotatedReferenceRhombus {
			diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "RotatedReferenceRhombus"))
		}
	}
	if (rhombusstuff.RotatedPlantCircumferenceShape == nil) != (rhombusstuffOther.RotatedPlantCircumferenceShape == nil) {
		diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "RotatedPlantCircumferenceShape"))
	} else if rhombusstuff.RotatedPlantCircumferenceShape != nil && rhombusstuffOther.RotatedPlantCircumferenceShape != nil {
		if rhombusstuff.RotatedPlantCircumferenceShape != rhombusstuffOther.RotatedPlantCircumferenceShape {
			diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "RotatedPlantCircumferenceShape"))
		}
	}
	if (rhombusstuff.RotatedGridPathShape == nil) != (rhombusstuffOther.RotatedGridPathShape == nil) {
		diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "RotatedGridPathShape"))
	} else if rhombusstuff.RotatedGridPathShape != nil && rhombusstuffOther.RotatedGridPathShape != nil {
		if rhombusstuff.RotatedGridPathShape != rhombusstuffOther.RotatedGridPathShape {
			diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "RotatedGridPathShape"))
		}
	}
	if (rhombusstuff.RotatedRhombusGridShape2 == nil) != (rhombusstuffOther.RotatedRhombusGridShape2 == nil) {
		diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "RotatedRhombusGridShape2"))
	} else if rhombusstuff.RotatedRhombusGridShape2 != nil && rhombusstuffOther.RotatedRhombusGridShape2 != nil {
		if rhombusstuff.RotatedRhombusGridShape2 != rhombusstuffOther.RotatedRhombusGridShape2 {
			diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "RotatedRhombusGridShape2"))
		}
	}
	if (rhombusstuff.GrowthCurveRhombusGridShape == nil) != (rhombusstuffOther.GrowthCurveRhombusGridShape == nil) {
		diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "GrowthCurveRhombusGridShape"))
	} else if rhombusstuff.GrowthCurveRhombusGridShape != nil && rhombusstuffOther.GrowthCurveRhombusGridShape != nil {
		if rhombusstuff.GrowthCurveRhombusGridShape != rhombusstuffOther.GrowthCurveRhombusGridShape {
			diffs = append(diffs, rhombusstuff.GongMarshallField(stage, "GrowthCurveRhombusGridShape"))
		}
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
	RotatedRhombusShapesDifferent := false
	if len(rotatedrhombusgridshape.RotatedRhombusShapes) != len(rotatedrhombusgridshapeOther.RotatedRhombusShapes) {
		RotatedRhombusShapesDifferent = true
	} else {
		for i := range rotatedrhombusgridshape.RotatedRhombusShapes {
			if (rotatedrhombusgridshape.RotatedRhombusShapes[i] == nil) != (rotatedrhombusgridshapeOther.RotatedRhombusShapes[i] == nil) {
				RotatedRhombusShapesDifferent = true
				break
			} else if rotatedrhombusgridshape.RotatedRhombusShapes[i] != nil && rotatedrhombusgridshapeOther.RotatedRhombusShapes[i] != nil {
				// this is a pointer comparaison
				if rotatedrhombusgridshape.RotatedRhombusShapes[i] != rotatedrhombusgridshapeOther.RotatedRhombusShapes[i] {
					RotatedRhombusShapesDifferent = true
					break
				}
			}
		}
	}
	if RotatedRhombusShapesDifferent {
		ops := Diff(stage, rotatedrhombusgridshape, rotatedrhombusgridshapeOther, "RotatedRhombusShapes", rotatedrhombusgridshapeOther.RotatedRhombusShapes, rotatedrhombusgridshape.RotatedRhombusShapes)
		diffs = append(diffs, ops)
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
	ShiftedBottomTopStartArcShapesDifferent := false
	if len(shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes) != len(shiftedbottomtopstartarcshapegridOther.ShiftedBottomTopStartArcShapes) {
		ShiftedBottomTopStartArcShapesDifferent = true
	} else {
		for i := range shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes {
			if (shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes[i] == nil) != (shiftedbottomtopstartarcshapegridOther.ShiftedBottomTopStartArcShapes[i] == nil) {
				ShiftedBottomTopStartArcShapesDifferent = true
				break
			} else if shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes[i] != nil && shiftedbottomtopstartarcshapegridOther.ShiftedBottomTopStartArcShapes[i] != nil {
				// this is a pointer comparaison
				if shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes[i] != shiftedbottomtopstartarcshapegridOther.ShiftedBottomTopStartArcShapes[i] {
					ShiftedBottomTopStartArcShapesDifferent = true
					break
				}
			}
		}
	}
	if ShiftedBottomTopStartArcShapesDifferent {
		ops := Diff(stage, shiftedbottomtopstartarcshapegrid, shiftedbottomtopstartarcshapegridOther, "ShiftedBottomTopStartArcShapes", shiftedbottomtopstartarcshapegridOther.ShiftedBottomTopStartArcShapes, shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes)
		diffs = append(diffs, ops)
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
	ShiftedLeftStackGrowthCurveStartArcShapesDifferent := false
	if len(shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes) != len(shiftedleftstackofgrowthcurveOther.ShiftedLeftStackGrowthCurveStartArcShapes) {
		ShiftedLeftStackGrowthCurveStartArcShapesDifferent = true
	} else {
		for i := range shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes {
			if (shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes[i] == nil) != (shiftedleftstackofgrowthcurveOther.ShiftedLeftStackGrowthCurveStartArcShapes[i] == nil) {
				ShiftedLeftStackGrowthCurveStartArcShapesDifferent = true
				break
			} else if shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes[i] != nil && shiftedleftstackofgrowthcurveOther.ShiftedLeftStackGrowthCurveStartArcShapes[i] != nil {
				// this is a pointer comparaison
				if shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes[i] != shiftedleftstackofgrowthcurveOther.ShiftedLeftStackGrowthCurveStartArcShapes[i] {
					ShiftedLeftStackGrowthCurveStartArcShapesDifferent = true
					break
				}
			}
		}
	}
	if ShiftedLeftStackGrowthCurveStartArcShapesDifferent {
		ops := Diff(stage, shiftedleftstackofgrowthcurve, shiftedleftstackofgrowthcurveOther, "ShiftedLeftStackGrowthCurveStartArcShapes", shiftedleftstackofgrowthcurveOther.ShiftedLeftStackGrowthCurveStartArcShapes, shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes)
		diffs = append(diffs, ops)
	}
	ShiftedLeftStackGrowthCurveEndArcShapesDifferent := false
	if len(shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes) != len(shiftedleftstackofgrowthcurveOther.ShiftedLeftStackGrowthCurveEndArcShapes) {
		ShiftedLeftStackGrowthCurveEndArcShapesDifferent = true
	} else {
		for i := range shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes {
			if (shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes[i] == nil) != (shiftedleftstackofgrowthcurveOther.ShiftedLeftStackGrowthCurveEndArcShapes[i] == nil) {
				ShiftedLeftStackGrowthCurveEndArcShapesDifferent = true
				break
			} else if shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes[i] != nil && shiftedleftstackofgrowthcurveOther.ShiftedLeftStackGrowthCurveEndArcShapes[i] != nil {
				// this is a pointer comparaison
				if shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes[i] != shiftedleftstackofgrowthcurveOther.ShiftedLeftStackGrowthCurveEndArcShapes[i] {
					ShiftedLeftStackGrowthCurveEndArcShapesDifferent = true
					break
				}
			}
		}
	}
	if ShiftedLeftStackGrowthCurveEndArcShapesDifferent {
		ops := Diff(stage, shiftedleftstackofgrowthcurve, shiftedleftstackofgrowthcurveOther, "ShiftedLeftStackGrowthCurveEndArcShapes", shiftedleftstackofgrowthcurveOther.ShiftedLeftStackGrowthCurveEndArcShapes, shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes)
		diffs = append(diffs, ops)
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
	ShiftedLeftStackNormalVectorsDifferent := false
	if len(shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors) != len(shiftedleftstackofnormalvectorOther.ShiftedLeftStackNormalVectors) {
		ShiftedLeftStackNormalVectorsDifferent = true
	} else {
		for i := range shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors {
			if (shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors[i] == nil) != (shiftedleftstackofnormalvectorOther.ShiftedLeftStackNormalVectors[i] == nil) {
				ShiftedLeftStackNormalVectorsDifferent = true
				break
			} else if shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors[i] != nil && shiftedleftstackofnormalvectorOther.ShiftedLeftStackNormalVectors[i] != nil {
				// this is a pointer comparaison
				if shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors[i] != shiftedleftstackofnormalvectorOther.ShiftedLeftStackNormalVectors[i] {
					ShiftedLeftStackNormalVectorsDifferent = true
					break
				}
			}
		}
	}
	if ShiftedLeftStackNormalVectorsDifferent {
		ops := Diff(stage, shiftedleftstackofnormalvector, shiftedleftstackofnormalvectorOther, "ShiftedLeftStackNormalVectors", shiftedleftstackofnormalvectorOther.ShiftedLeftStackNormalVectors, shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors)
		diffs = append(diffs, ops)
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
	StackGrowthCurve2DStartHalfwayArcShapesDifferent := false
	if len(stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes) != len(stackofgrowthcurve2dOther.StackGrowthCurve2DStartHalfwayArcShapes) {
		StackGrowthCurve2DStartHalfwayArcShapesDifferent = true
	} else {
		for i := range stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes {
			if (stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes[i] == nil) != (stackofgrowthcurve2dOther.StackGrowthCurve2DStartHalfwayArcShapes[i] == nil) {
				StackGrowthCurve2DStartHalfwayArcShapesDifferent = true
				break
			} else if stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes[i] != nil && stackofgrowthcurve2dOther.StackGrowthCurve2DStartHalfwayArcShapes[i] != nil {
				// this is a pointer comparaison
				if stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes[i] != stackofgrowthcurve2dOther.StackGrowthCurve2DStartHalfwayArcShapes[i] {
					StackGrowthCurve2DStartHalfwayArcShapesDifferent = true
					break
				}
			}
		}
	}
	if StackGrowthCurve2DStartHalfwayArcShapesDifferent {
		ops := Diff(stage, stackofgrowthcurve2d, stackofgrowthcurve2dOther, "StackGrowthCurve2DStartHalfwayArcShapes", stackofgrowthcurve2dOther.StackGrowthCurve2DStartHalfwayArcShapes, stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes)
		diffs = append(diffs, ops)
	}
	StackGrowthCurve2DEndHalfwayArcShapesDifferent := false
	if len(stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes) != len(stackofgrowthcurve2dOther.StackGrowthCurve2DEndHalfwayArcShapes) {
		StackGrowthCurve2DEndHalfwayArcShapesDifferent = true
	} else {
		for i := range stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes {
			if (stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes[i] == nil) != (stackofgrowthcurve2dOther.StackGrowthCurve2DEndHalfwayArcShapes[i] == nil) {
				StackGrowthCurve2DEndHalfwayArcShapesDifferent = true
				break
			} else if stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes[i] != nil && stackofgrowthcurve2dOther.StackGrowthCurve2DEndHalfwayArcShapes[i] != nil {
				// this is a pointer comparaison
				if stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes[i] != stackofgrowthcurve2dOther.StackGrowthCurve2DEndHalfwayArcShapes[i] {
					StackGrowthCurve2DEndHalfwayArcShapesDifferent = true
					break
				}
			}
		}
	}
	if StackGrowthCurve2DEndHalfwayArcShapesDifferent {
		ops := Diff(stage, stackofgrowthcurve2d, stackofgrowthcurve2dOther, "StackGrowthCurve2DEndHalfwayArcShapes", stackofgrowthcurve2dOther.StackGrowthCurve2DEndHalfwayArcShapes, stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes)
		diffs = append(diffs, ops)
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
	StackGrowthCurve2DRibbonStartShapesDifferent := false
	if len(stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes) != len(stackofgrowthcurve2dribbonOther.StackGrowthCurve2DRibbonStartShapes) {
		StackGrowthCurve2DRibbonStartShapesDifferent = true
	} else {
		for i := range stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes {
			if (stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes[i] == nil) != (stackofgrowthcurve2dribbonOther.StackGrowthCurve2DRibbonStartShapes[i] == nil) {
				StackGrowthCurve2DRibbonStartShapesDifferent = true
				break
			} else if stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes[i] != nil && stackofgrowthcurve2dribbonOther.StackGrowthCurve2DRibbonStartShapes[i] != nil {
				// this is a pointer comparaison
				if stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes[i] != stackofgrowthcurve2dribbonOther.StackGrowthCurve2DRibbonStartShapes[i] {
					StackGrowthCurve2DRibbonStartShapesDifferent = true
					break
				}
			}
		}
	}
	if StackGrowthCurve2DRibbonStartShapesDifferent {
		ops := Diff(stage, stackofgrowthcurve2dribbon, stackofgrowthcurve2dribbonOther, "StackGrowthCurve2DRibbonStartShapes", stackofgrowthcurve2dribbonOther.StackGrowthCurve2DRibbonStartShapes, stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes)
		diffs = append(diffs, ops)
	}
	StackGrowthCurve2DRibbonEndShapesDifferent := false
	if len(stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes) != len(stackofgrowthcurve2dribbonOther.StackGrowthCurve2DRibbonEndShapes) {
		StackGrowthCurve2DRibbonEndShapesDifferent = true
	} else {
		for i := range stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes {
			if (stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes[i] == nil) != (stackofgrowthcurve2dribbonOther.StackGrowthCurve2DRibbonEndShapes[i] == nil) {
				StackGrowthCurve2DRibbonEndShapesDifferent = true
				break
			} else if stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes[i] != nil && stackofgrowthcurve2dribbonOther.StackGrowthCurve2DRibbonEndShapes[i] != nil {
				// this is a pointer comparaison
				if stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes[i] != stackofgrowthcurve2dribbonOther.StackGrowthCurve2DRibbonEndShapes[i] {
					StackGrowthCurve2DRibbonEndShapesDifferent = true
					break
				}
			}
		}
	}
	if StackGrowthCurve2DRibbonEndShapesDifferent {
		ops := Diff(stage, stackofgrowthcurve2dribbon, stackofgrowthcurve2dribbonOther, "StackGrowthCurve2DRibbonEndShapes", stackofgrowthcurve2dribbonOther.StackGrowthCurve2DRibbonEndShapes, stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes)
		diffs = append(diffs, ops)
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
	StackRotatedGrowthCurve2DStartArcShapesDifferent := false
	if len(stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes) != len(stackofrotatedgrowthcurve2dOther.StackRotatedGrowthCurve2DStartArcShapes) {
		StackRotatedGrowthCurve2DStartArcShapesDifferent = true
	} else {
		for i := range stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes {
			if (stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes[i] == nil) != (stackofrotatedgrowthcurve2dOther.StackRotatedGrowthCurve2DStartArcShapes[i] == nil) {
				StackRotatedGrowthCurve2DStartArcShapesDifferent = true
				break
			} else if stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes[i] != nil && stackofrotatedgrowthcurve2dOther.StackRotatedGrowthCurve2DStartArcShapes[i] != nil {
				// this is a pointer comparaison
				if stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes[i] != stackofrotatedgrowthcurve2dOther.StackRotatedGrowthCurve2DStartArcShapes[i] {
					StackRotatedGrowthCurve2DStartArcShapesDifferent = true
					break
				}
			}
		}
	}
	if StackRotatedGrowthCurve2DStartArcShapesDifferent {
		ops := Diff(stage, stackofrotatedgrowthcurve2d, stackofrotatedgrowthcurve2dOther, "StackRotatedGrowthCurve2DStartArcShapes", stackofrotatedgrowthcurve2dOther.StackRotatedGrowthCurve2DStartArcShapes, stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes)
		diffs = append(diffs, ops)
	}
	StackRotatedGrowthCurve2DEndArcShapesDifferent := false
	if len(stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes) != len(stackofrotatedgrowthcurve2dOther.StackRotatedGrowthCurve2DEndArcShapes) {
		StackRotatedGrowthCurve2DEndArcShapesDifferent = true
	} else {
		for i := range stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes {
			if (stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes[i] == nil) != (stackofrotatedgrowthcurve2dOther.StackRotatedGrowthCurve2DEndArcShapes[i] == nil) {
				StackRotatedGrowthCurve2DEndArcShapesDifferent = true
				break
			} else if stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes[i] != nil && stackofrotatedgrowthcurve2dOther.StackRotatedGrowthCurve2DEndArcShapes[i] != nil {
				// this is a pointer comparaison
				if stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes[i] != stackofrotatedgrowthcurve2dOther.StackRotatedGrowthCurve2DEndArcShapes[i] {
					StackRotatedGrowthCurve2DEndArcShapesDifferent = true
					break
				}
			}
		}
	}
	if StackRotatedGrowthCurve2DEndArcShapesDifferent {
		ops := Diff(stage, stackofrotatedgrowthcurve2d, stackofrotatedgrowthcurve2dOther, "StackRotatedGrowthCurve2DEndArcShapes", stackofrotatedgrowthcurve2dOther.StackRotatedGrowthCurve2DEndArcShapes, stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes)
		diffs = append(diffs, ops)
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
	StartArcShapesDifferent := false
	if len(startarcshapegrid.StartArcShapes) != len(startarcshapegridOther.StartArcShapes) {
		StartArcShapesDifferent = true
	} else {
		for i := range startarcshapegrid.StartArcShapes {
			if (startarcshapegrid.StartArcShapes[i] == nil) != (startarcshapegridOther.StartArcShapes[i] == nil) {
				StartArcShapesDifferent = true
				break
			} else if startarcshapegrid.StartArcShapes[i] != nil && startarcshapegridOther.StartArcShapes[i] != nil {
				// this is a pointer comparaison
				if startarcshapegrid.StartArcShapes[i] != startarcshapegridOther.StartArcShapes[i] {
					StartArcShapesDifferent = true
					break
				}
			}
		}
	}
	if StartArcShapesDifferent {
		ops := Diff(stage, startarcshapegrid, startarcshapegridOther, "StartArcShapes", startarcshapegridOther.StartArcShapes, startarcshapegrid.StartArcShapes)
		diffs = append(diffs, ops)
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
	StartHalfwayArcShapesDifferent := false
	if len(starthalfwayarcshapegrid.StartHalfwayArcShapes) != len(starthalfwayarcshapegridOther.StartHalfwayArcShapes) {
		StartHalfwayArcShapesDifferent = true
	} else {
		for i := range starthalfwayarcshapegrid.StartHalfwayArcShapes {
			if (starthalfwayarcshapegrid.StartHalfwayArcShapes[i] == nil) != (starthalfwayarcshapegridOther.StartHalfwayArcShapes[i] == nil) {
				StartHalfwayArcShapesDifferent = true
				break
			} else if starthalfwayarcshapegrid.StartHalfwayArcShapes[i] != nil && starthalfwayarcshapegridOther.StartHalfwayArcShapes[i] != nil {
				// this is a pointer comparaison
				if starthalfwayarcshapegrid.StartHalfwayArcShapes[i] != starthalfwayarcshapegridOther.StartHalfwayArcShapes[i] {
					StartHalfwayArcShapesDifferent = true
					break
				}
			}
		}
	}
	if StartHalfwayArcShapesDifferent {
		ops := Diff(stage, starthalfwayarcshapegrid, starthalfwayarcshapegridOther, "StartHalfwayArcShapes", starthalfwayarcshapegridOther.StartHalfwayArcShapes, starthalfwayarcshapegrid.StartHalfwayArcShapes)
		diffs = append(diffs, ops)
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
	TopEndArcShapesDifferent := false
	if len(topendarcshapegrid.TopEndArcShapes) != len(topendarcshapegridOther.TopEndArcShapes) {
		TopEndArcShapesDifferent = true
	} else {
		for i := range topendarcshapegrid.TopEndArcShapes {
			if (topendarcshapegrid.TopEndArcShapes[i] == nil) != (topendarcshapegridOther.TopEndArcShapes[i] == nil) {
				TopEndArcShapesDifferent = true
				break
			} else if topendarcshapegrid.TopEndArcShapes[i] != nil && topendarcshapegridOther.TopEndArcShapes[i] != nil {
				// this is a pointer comparaison
				if topendarcshapegrid.TopEndArcShapes[i] != topendarcshapegridOther.TopEndArcShapes[i] {
					TopEndArcShapesDifferent = true
					break
				}
			}
		}
	}
	if TopEndArcShapesDifferent {
		ops := Diff(stage, topendarcshapegrid, topendarcshapegridOther, "TopEndArcShapes", topendarcshapegridOther.TopEndArcShapes, topendarcshapegrid.TopEndArcShapes)
		diffs = append(diffs, ops)
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
	TopEndHalfwayArcShapesDifferent := false
	if len(topendhalfwayarcshapegrid.TopEndHalfwayArcShapes) != len(topendhalfwayarcshapegridOther.TopEndHalfwayArcShapes) {
		TopEndHalfwayArcShapesDifferent = true
	} else {
		for i := range topendhalfwayarcshapegrid.TopEndHalfwayArcShapes {
			if (topendhalfwayarcshapegrid.TopEndHalfwayArcShapes[i] == nil) != (topendhalfwayarcshapegridOther.TopEndHalfwayArcShapes[i] == nil) {
				TopEndHalfwayArcShapesDifferent = true
				break
			} else if topendhalfwayarcshapegrid.TopEndHalfwayArcShapes[i] != nil && topendhalfwayarcshapegridOther.TopEndHalfwayArcShapes[i] != nil {
				// this is a pointer comparaison
				if topendhalfwayarcshapegrid.TopEndHalfwayArcShapes[i] != topendhalfwayarcshapegridOther.TopEndHalfwayArcShapes[i] {
					TopEndHalfwayArcShapesDifferent = true
					break
				}
			}
		}
	}
	if TopEndHalfwayArcShapesDifferent {
		ops := Diff(stage, topendhalfwayarcshapegrid, topendhalfwayarcshapegridOther, "TopEndHalfwayArcShapes", topendhalfwayarcshapegridOther.TopEndHalfwayArcShapes, topendhalfwayarcshapegrid.TopEndHalfwayArcShapes)
		diffs = append(diffs, ops)
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
	if (topgrowthcurve2d.TopStartHalfwayArcShapeGrid == nil) != (topgrowthcurve2dOther.TopStartHalfwayArcShapeGrid == nil) {
		diffs = append(diffs, topgrowthcurve2d.GongMarshallField(stage, "TopStartHalfwayArcShapeGrid"))
	} else if topgrowthcurve2d.TopStartHalfwayArcShapeGrid != nil && topgrowthcurve2dOther.TopStartHalfwayArcShapeGrid != nil {
		if topgrowthcurve2d.TopStartHalfwayArcShapeGrid != topgrowthcurve2dOther.TopStartHalfwayArcShapeGrid {
			diffs = append(diffs, topgrowthcurve2d.GongMarshallField(stage, "TopStartHalfwayArcShapeGrid"))
		}
	}
	if (topgrowthcurve2d.TopEndHalfwayArcShapeGrid == nil) != (topgrowthcurve2dOther.TopEndHalfwayArcShapeGrid == nil) {
		diffs = append(diffs, topgrowthcurve2d.GongMarshallField(stage, "TopEndHalfwayArcShapeGrid"))
	} else if topgrowthcurve2d.TopEndHalfwayArcShapeGrid != nil && topgrowthcurve2dOther.TopEndHalfwayArcShapeGrid != nil {
		if topgrowthcurve2d.TopEndHalfwayArcShapeGrid != topgrowthcurve2dOther.TopEndHalfwayArcShapeGrid {
			diffs = append(diffs, topgrowthcurve2d.GongMarshallField(stage, "TopEndHalfwayArcShapeGrid"))
		}
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
	TopMidArcVectorShapesDifferent := false
	if len(topmidarcvectorshapegrid.TopMidArcVectorShapes) != len(topmidarcvectorshapegridOther.TopMidArcVectorShapes) {
		TopMidArcVectorShapesDifferent = true
	} else {
		for i := range topmidarcvectorshapegrid.TopMidArcVectorShapes {
			if (topmidarcvectorshapegrid.TopMidArcVectorShapes[i] == nil) != (topmidarcvectorshapegridOther.TopMidArcVectorShapes[i] == nil) {
				TopMidArcVectorShapesDifferent = true
				break
			} else if topmidarcvectorshapegrid.TopMidArcVectorShapes[i] != nil && topmidarcvectorshapegridOther.TopMidArcVectorShapes[i] != nil {
				// this is a pointer comparaison
				if topmidarcvectorshapegrid.TopMidArcVectorShapes[i] != topmidarcvectorshapegridOther.TopMidArcVectorShapes[i] {
					TopMidArcVectorShapesDifferent = true
					break
				}
			}
		}
	}
	if TopMidArcVectorShapesDifferent {
		ops := Diff(stage, topmidarcvectorshapegrid, topmidarcvectorshapegridOther, "TopMidArcVectorShapes", topmidarcvectorshapegridOther.TopMidArcVectorShapes, topmidarcvectorshapegrid.TopMidArcVectorShapes)
		diffs = append(diffs, ops)
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
	TopStackGrowthCurve2DStartHalfwayArcShapesDifferent := false
	if len(topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes) != len(topstackofgrowthcurve2dOther.TopStackGrowthCurve2DStartHalfwayArcShapes) {
		TopStackGrowthCurve2DStartHalfwayArcShapesDifferent = true
	} else {
		for i := range topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes {
			if (topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes[i] == nil) != (topstackofgrowthcurve2dOther.TopStackGrowthCurve2DStartHalfwayArcShapes[i] == nil) {
				TopStackGrowthCurve2DStartHalfwayArcShapesDifferent = true
				break
			} else if topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes[i] != nil && topstackofgrowthcurve2dOther.TopStackGrowthCurve2DStartHalfwayArcShapes[i] != nil {
				// this is a pointer comparaison
				if topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes[i] != topstackofgrowthcurve2dOther.TopStackGrowthCurve2DStartHalfwayArcShapes[i] {
					TopStackGrowthCurve2DStartHalfwayArcShapesDifferent = true
					break
				}
			}
		}
	}
	if TopStackGrowthCurve2DStartHalfwayArcShapesDifferent {
		ops := Diff(stage, topstackofgrowthcurve2d, topstackofgrowthcurve2dOther, "TopStackGrowthCurve2DStartHalfwayArcShapes", topstackofgrowthcurve2dOther.TopStackGrowthCurve2DStartHalfwayArcShapes, topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes)
		diffs = append(diffs, ops)
	}
	TopStackGrowthCurve2DEndHalfwayArcShapesDifferent := false
	if len(topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes) != len(topstackofgrowthcurve2dOther.TopStackGrowthCurve2DEndHalfwayArcShapes) {
		TopStackGrowthCurve2DEndHalfwayArcShapesDifferent = true
	} else {
		for i := range topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes {
			if (topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes[i] == nil) != (topstackofgrowthcurve2dOther.TopStackGrowthCurve2DEndHalfwayArcShapes[i] == nil) {
				TopStackGrowthCurve2DEndHalfwayArcShapesDifferent = true
				break
			} else if topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes[i] != nil && topstackofgrowthcurve2dOther.TopStackGrowthCurve2DEndHalfwayArcShapes[i] != nil {
				// this is a pointer comparaison
				if topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes[i] != topstackofgrowthcurve2dOther.TopStackGrowthCurve2DEndHalfwayArcShapes[i] {
					TopStackGrowthCurve2DEndHalfwayArcShapesDifferent = true
					break
				}
			}
		}
	}
	if TopStackGrowthCurve2DEndHalfwayArcShapesDifferent {
		ops := Diff(stage, topstackofgrowthcurve2d, topstackofgrowthcurve2dOther, "TopStackGrowthCurve2DEndHalfwayArcShapes", topstackofgrowthcurve2dOther.TopStackGrowthCurve2DEndHalfwayArcShapes, topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes)
		diffs = append(diffs, ops)
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
	TopStackOfRotatedGrowthCurve2DStartArcShapesDifferent := false
	if len(topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes) != len(topstackofrotatedgrowthcurve2dOther.TopStackOfRotatedGrowthCurve2DStartArcShapes) {
		TopStackOfRotatedGrowthCurve2DStartArcShapesDifferent = true
	} else {
		for i := range topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes {
			if (topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes[i] == nil) != (topstackofrotatedgrowthcurve2dOther.TopStackOfRotatedGrowthCurve2DStartArcShapes[i] == nil) {
				TopStackOfRotatedGrowthCurve2DStartArcShapesDifferent = true
				break
			} else if topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes[i] != nil && topstackofrotatedgrowthcurve2dOther.TopStackOfRotatedGrowthCurve2DStartArcShapes[i] != nil {
				// this is a pointer comparaison
				if topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes[i] != topstackofrotatedgrowthcurve2dOther.TopStackOfRotatedGrowthCurve2DStartArcShapes[i] {
					TopStackOfRotatedGrowthCurve2DStartArcShapesDifferent = true
					break
				}
			}
		}
	}
	if TopStackOfRotatedGrowthCurve2DStartArcShapesDifferent {
		ops := Diff(stage, topstackofrotatedgrowthcurve2d, topstackofrotatedgrowthcurve2dOther, "TopStackOfRotatedGrowthCurve2DStartArcShapes", topstackofrotatedgrowthcurve2dOther.TopStackOfRotatedGrowthCurve2DStartArcShapes, topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes)
		diffs = append(diffs, ops)
	}
	TopStackOfRotatedGrowthCurve2DEndArcShapesDifferent := false
	if len(topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes) != len(topstackofrotatedgrowthcurve2dOther.TopStackOfRotatedGrowthCurve2DEndArcShapes) {
		TopStackOfRotatedGrowthCurve2DEndArcShapesDifferent = true
	} else {
		for i := range topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes {
			if (topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes[i] == nil) != (topstackofrotatedgrowthcurve2dOther.TopStackOfRotatedGrowthCurve2DEndArcShapes[i] == nil) {
				TopStackOfRotatedGrowthCurve2DEndArcShapesDifferent = true
				break
			} else if topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes[i] != nil && topstackofrotatedgrowthcurve2dOther.TopStackOfRotatedGrowthCurve2DEndArcShapes[i] != nil {
				// this is a pointer comparaison
				if topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes[i] != topstackofrotatedgrowthcurve2dOther.TopStackOfRotatedGrowthCurve2DEndArcShapes[i] {
					TopStackOfRotatedGrowthCurve2DEndArcShapesDifferent = true
					break
				}
			}
		}
	}
	if TopStackOfRotatedGrowthCurve2DEndArcShapesDifferent {
		ops := Diff(stage, topstackofrotatedgrowthcurve2d, topstackofrotatedgrowthcurve2dOther, "TopStackOfRotatedGrowthCurve2DEndArcShapes", topstackofrotatedgrowthcurve2dOther.TopStackOfRotatedGrowthCurve2DEndArcShapes, topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes)
		diffs = append(diffs, ops)
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
	TopStartArcShapesDifferent := false
	if len(topstartarcshapegrid.TopStartArcShapes) != len(topstartarcshapegridOther.TopStartArcShapes) {
		TopStartArcShapesDifferent = true
	} else {
		for i := range topstartarcshapegrid.TopStartArcShapes {
			if (topstartarcshapegrid.TopStartArcShapes[i] == nil) != (topstartarcshapegridOther.TopStartArcShapes[i] == nil) {
				TopStartArcShapesDifferent = true
				break
			} else if topstartarcshapegrid.TopStartArcShapes[i] != nil && topstartarcshapegridOther.TopStartArcShapes[i] != nil {
				// this is a pointer comparaison
				if topstartarcshapegrid.TopStartArcShapes[i] != topstartarcshapegridOther.TopStartArcShapes[i] {
					TopStartArcShapesDifferent = true
					break
				}
			}
		}
	}
	if TopStartArcShapesDifferent {
		ops := Diff(stage, topstartarcshapegrid, topstartarcshapegridOther, "TopStartArcShapes", topstartarcshapegridOther.TopStartArcShapes, topstartarcshapegrid.TopStartArcShapes)
		diffs = append(diffs, ops)
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
	TopStartHalfwayArcShapesDifferent := false
	if len(topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes) != len(topstarthalfwayarcshapegridOther.TopStartHalfwayArcShapes) {
		TopStartHalfwayArcShapesDifferent = true
	} else {
		for i := range topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes {
			if (topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes[i] == nil) != (topstarthalfwayarcshapegridOther.TopStartHalfwayArcShapes[i] == nil) {
				TopStartHalfwayArcShapesDifferent = true
				break
			} else if topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes[i] != nil && topstarthalfwayarcshapegridOther.TopStartHalfwayArcShapes[i] != nil {
				// this is a pointer comparaison
				if topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes[i] != topstarthalfwayarcshapegridOther.TopStartHalfwayArcShapes[i] {
					TopStartHalfwayArcShapesDifferent = true
					break
				}
			}
		}
	}
	if TopStartHalfwayArcShapesDifferent {
		ops := Diff(stage, topstarthalfwayarcshapegrid, topstarthalfwayarcshapegridOther, "TopStartHalfwayArcShapes", topstarthalfwayarcshapegridOther.TopStartHalfwayArcShapes, topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes)
		diffs = append(diffs, ops)
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
func (verticaltorusstackshape *VerticalTorusStackShape) GongDiff(stage *Stage, verticaltorusstackshapeOther *VerticalTorusStackShape) (diffs []string) {
	// insertion point for field diffs
	if verticaltorusstackshape.Name != verticaltorusstackshapeOther.Name {
		diffs = append(diffs, verticaltorusstackshape.GongMarshallField(stage, "Name"))
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
