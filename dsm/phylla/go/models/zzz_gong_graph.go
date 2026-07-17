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

	case *GrowthCurveBezierShape:
		ok = stage.IsStagedGrowthCurveBezierShape(target)

	case *GrowthCurveBezierShapeGrid:
		ok = stage.IsStagedGrowthCurveBezierShapeGrid(target)

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

	case *NextCircleShape:
		ok = stage.IsStagedNextCircleShape(target)

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

	case *StackGrowthCurveEndArcShape:
		ok = stage.IsStagedStackGrowthCurveEndArcShape(target)

	case *StackGrowthCurveStartArcShape:
		ok = stage.IsStagedStackGrowthCurveStartArcShape(target)

	case *StackOfGrowthCurve:
		ok = stage.IsStagedStackOfGrowthCurve(target)

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

	case *TopStackGrowthCurveEndArcShape:
		ok = stage.IsStagedTopStackGrowthCurveEndArcShape(target)

	case *TopStackGrowthCurveStartArcShape:
		ok = stage.IsStagedTopStackGrowthCurveStartArcShape(target)

	case *TopStackOfGrowthCurve:
		ok = stage.IsStagedTopStackOfGrowthCurve(target)

	case *TopStartArcShape:
		ok = stage.IsStagedTopStartArcShape(target)

	case *TopStartArcShapeGrid:
		ok = stage.IsStagedTopStartArcShapeGrid(target)

	case *TopStartHalfwayArcShape:
		ok = stage.IsStagedTopStartHalfwayArcShape(target)

	case *TopStartHalfwayArcShapeGrid:
		ok = stage.IsStagedTopStartHalfwayArcShapeGrid(target)

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

	case *GrowthCurveBezierShape:
		ok = stage.IsStagedGrowthCurveBezierShape(target)

	case *GrowthCurveBezierShapeGrid:
		ok = stage.IsStagedGrowthCurveBezierShapeGrid(target)

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

	case *NextCircleShape:
		ok = stage.IsStagedNextCircleShape(target)

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

	case *StackGrowthCurveEndArcShape:
		ok = stage.IsStagedStackGrowthCurveEndArcShape(target)

	case *StackGrowthCurveStartArcShape:
		ok = stage.IsStagedStackGrowthCurveStartArcShape(target)

	case *StackOfGrowthCurve:
		ok = stage.IsStagedStackOfGrowthCurve(target)

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

	case *TopStackGrowthCurveEndArcShape:
		ok = stage.IsStagedTopStackGrowthCurveEndArcShape(target)

	case *TopStackGrowthCurveStartArcShape:
		ok = stage.IsStagedTopStackGrowthCurveStartArcShape(target)

	case *TopStackOfGrowthCurve:
		ok = stage.IsStagedTopStackOfGrowthCurve(target)

	case *TopStartArcShape:
		ok = stage.IsStagedTopStartArcShape(target)

	case *TopStartArcShapeGrid:
		ok = stage.IsStagedTopStartArcShapeGrid(target)

	case *TopStartHalfwayArcShape:
		ok = stage.IsStagedTopStartHalfwayArcShape(target)

	case *TopStartHalfwayArcShapeGrid:
		ok = stage.IsStagedTopStartHalfwayArcShapeGrid(target)

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

func (stage *Stage) IsStagedGrowthCurveBezierShape(growthcurvebeziershape *GrowthCurveBezierShape) (ok bool) {

	_, ok = stage.GrowthCurveBezierShapes[growthcurvebeziershape]

	return
}

func (stage *Stage) IsStagedGrowthCurveBezierShapeGrid(growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) (ok bool) {

	_, ok = stage.GrowthCurveBezierShapeGrids[growthcurvebeziershapegrid]

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

func (stage *Stage) IsStagedNextCircleShape(nextcircleshape *NextCircleShape) (ok bool) {

	_, ok = stage.NextCircleShapes[nextcircleshape]

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

func (stage *Stage) IsStagedStackGrowthCurveEndArcShape(stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) (ok bool) {

	_, ok = stage.StackGrowthCurveEndArcShapes[stackgrowthcurveendarcshape]

	return
}

func (stage *Stage) IsStagedStackGrowthCurveStartArcShape(stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) (ok bool) {

	_, ok = stage.StackGrowthCurveStartArcShapes[stackgrowthcurvestartarcshape]

	return
}

func (stage *Stage) IsStagedStackOfGrowthCurve(stackofgrowthcurve *StackOfGrowthCurve) (ok bool) {

	_, ok = stage.StackOfGrowthCurves[stackofgrowthcurve]

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

func (stage *Stage) IsStagedTopStackGrowthCurveEndArcShape(topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) (ok bool) {

	_, ok = stage.TopStackGrowthCurveEndArcShapes[topstackgrowthcurveendarcshape]

	return
}

func (stage *Stage) IsStagedTopStackGrowthCurveStartArcShape(topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) (ok bool) {

	_, ok = stage.TopStackGrowthCurveStartArcShapes[topstackgrowthcurvestartarcshape]

	return
}

func (stage *Stage) IsStagedTopStackOfGrowthCurve(topstackofgrowthcurve *TopStackOfGrowthCurve) (ok bool) {

	_, ok = stage.TopStackOfGrowthCurves[topstackofgrowthcurve]

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

	case *GrowthCurveBezierShape:
		stage.StageBranchGrowthCurveBezierShape(target)

	case *GrowthCurveBezierShapeGrid:
		stage.StageBranchGrowthCurveBezierShapeGrid(target)

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

	case *NextCircleShape:
		stage.StageBranchNextCircleShape(target)

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

	case *StackGrowthCurveEndArcShape:
		stage.StageBranchStackGrowthCurveEndArcShape(target)

	case *StackGrowthCurveStartArcShape:
		stage.StageBranchStackGrowthCurveStartArcShape(target)

	case *StackOfGrowthCurve:
		stage.StageBranchStackOfGrowthCurve(target)

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

	case *TopStackGrowthCurveEndArcShape:
		stage.StageBranchTopStackGrowthCurveEndArcShape(target)

	case *TopStackGrowthCurveStartArcShape:
		stage.StageBranchTopStackGrowthCurveStartArcShape(target)

	case *TopStackOfGrowthCurve:
		stage.StageBranchTopStackOfGrowthCurve(target)

	case *TopStartArcShape:
		stage.StageBranchTopStartArcShape(target)

	case *TopStartArcShapeGrid:
		stage.StageBranchTopStartArcShapeGrid(target)

	case *TopStartHalfwayArcShape:
		stage.StageBranchTopStartHalfwayArcShape(target)

	case *TopStartHalfwayArcShapeGrid:
		stage.StageBranchTopStartHalfwayArcShapeGrid(target)

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
	if growthcurve2d.StartArcShapeGrid != nil {
		StageBranch(stage, growthcurve2d.StartArcShapeGrid)
	}
	if growthcurve2d.EndArcShapeGrid != nil {
		StageBranch(stage, growthcurve2d.EndArcShapeGrid)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGrowthCurveBezierShape(growthcurvebeziershape *GrowthCurveBezierShape) {

	// check if instance is already staged
	if IsStaged(stage, growthcurvebeziershape) {
		return
	}

	growthcurvebeziershape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGrowthCurveBezierShapeGrid(growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) {

	// check if instance is already staged
	if IsStaged(stage, growthcurvebeziershapegrid) {
		return
	}

	growthcurvebeziershapegrid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _growthcurvebeziershape := range growthcurvebeziershapegrid.GrowthCurveBezierShapes {
		StageBranch(stage, _growthcurvebeziershape)
	}

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

func (stage *Stage) StageBranchNextCircleShape(nextcircleshape *NextCircleShape) {

	// check if instance is already staged
	if IsStaged(stage, nextcircleshape) {
		return
	}

	nextcircleshape.Stage(stage)

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
	if plant.ReferenceRhombus != nil {
		StageBranch(stage, plant.ReferenceRhombus)
	}
	if plant.PlantCircumferenceShape != nil {
		StageBranch(stage, plant.PlantCircumferenceShape)
	}
	if plant.GridPathShape != nil {
		StageBranch(stage, plant.GridPathShape)
	}
	if plant.InitialRhombusGridShape != nil {
		StageBranch(stage, plant.InitialRhombusGridShape)
	}
	if plant.ExplanationTextShape != nil {
		StageBranch(stage, plant.ExplanationTextShape)
	}
	if plant.RotatedReferenceRhombus != nil {
		StageBranch(stage, plant.RotatedReferenceRhombus)
	}
	if plant.RotatedPlantCircumferenceShape != nil {
		StageBranch(stage, plant.RotatedPlantCircumferenceShape)
	}
	if plant.RotatedGridPathShape != nil {
		StageBranch(stage, plant.RotatedGridPathShape)
	}
	if plant.RotatedRhombusGridShape2 != nil {
		StageBranch(stage, plant.RotatedRhombusGridShape2)
	}
	if plant.GrowthCurveRhombusGridShape != nil {
		StageBranch(stage, plant.GrowthCurveRhombusGridShape)
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
	if plant.EndArcShapeGrid != nil {
		StageBranch(stage, plant.EndArcShapeGrid)
	}
	if plant.TopEndArcShapeGrid != nil {
		StageBranch(stage, plant.TopEndArcShapeGrid)
	}
	if plant.GrowthCurveBezierShapeGrid != nil {
		StageBranch(stage, plant.GrowthCurveBezierShapeGrid)
	}
	if plant.StackOfGrowthCurve != nil {
		StageBranch(stage, plant.StackOfGrowthCurve)
	}
	if plant.TopStackOfGrowthCurve != nil {
		StageBranch(stage, plant.TopStackOfGrowthCurve)
	}
	if plant.ShiftedLeftStackOfGrowthCurve != nil {
		StageBranch(stage, plant.ShiftedLeftStackOfGrowthCurve)
	}
	if plant.ShiftedLeftStackOfNormalVector != nil {
		StageBranch(stage, plant.ShiftedLeftStackOfNormalVector)
	}
	if plant.GrowthCurve2D != nil {
		StageBranch(stage, plant.GrowthCurve2D)
	}
	if plant.TopGrowthCurve2D != nil {
		StageBranch(stage, plant.TopGrowthCurve2D)
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

func (stage *Stage) StageBranchStackGrowthCurveEndArcShape(stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) {

	// check if instance is already staged
	if IsStaged(stage, stackgrowthcurveendarcshape) {
		return
	}

	stackgrowthcurveendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackGrowthCurveStartArcShape(stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) {

	// check if instance is already staged
	if IsStaged(stage, stackgrowthcurvestartarcshape) {
		return
	}

	stackgrowthcurvestartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackOfGrowthCurve(stackofgrowthcurve *StackOfGrowthCurve) {

	// check if instance is already staged
	if IsStaged(stage, stackofgrowthcurve) {
		return
	}

	stackofgrowthcurve.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stackgrowthcurvestartarcshape := range stackofgrowthcurve.StackGrowthCurveStartArcShapes {
		StageBranch(stage, _stackgrowthcurvestartarcshape)
	}
	for _, _stackgrowthcurveendarcshape := range stackofgrowthcurve.StackGrowthCurveEndArcShapes {
		StageBranch(stage, _stackgrowthcurveendarcshape)
	}

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
	if topgrowthcurve2d.TopStartArcShapeGrid != nil {
		StageBranch(stage, topgrowthcurve2d.TopStartArcShapeGrid)
	}
	if topgrowthcurve2d.TopEndArcShapeGrid != nil {
		StageBranch(stage, topgrowthcurve2d.TopEndArcShapeGrid)
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

func (stage *Stage) StageBranchTopStackGrowthCurveEndArcShape(topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) {

	// check if instance is already staged
	if IsStaged(stage, topstackgrowthcurveendarcshape) {
		return
	}

	topstackgrowthcurveendarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStackGrowthCurveStartArcShape(topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) {

	// check if instance is already staged
	if IsStaged(stage, topstackgrowthcurvestartarcshape) {
		return
	}

	topstackgrowthcurvestartarcshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStackOfGrowthCurve(topstackofgrowthcurve *TopStackOfGrowthCurve) {

	// check if instance is already staged
	if IsStaged(stage, topstackofgrowthcurve) {
		return
	}

	topstackofgrowthcurve.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topstackgrowthcurvestartarcshape := range topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes {
		StageBranch(stage, _topstackgrowthcurvestartarcshape)
	}
	for _, _topstackgrowthcurveendarcshape := range topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes {
		StageBranch(stage, _topstackgrowthcurveendarcshape)
	}

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

	case *GrowthCurveBezierShape:
		toT := CopyBranchGrowthCurveBezierShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthCurveBezierShapeGrid:
		toT := CopyBranchGrowthCurveBezierShapeGrid(mapOrigCopy, fromT)
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

	case *NextCircleShape:
		toT := CopyBranchNextCircleShape(mapOrigCopy, fromT)
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

	case *StackGrowthCurveEndArcShape:
		toT := CopyBranchStackGrowthCurveEndArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackGrowthCurveStartArcShape:
		toT := CopyBranchStackGrowthCurveStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfGrowthCurve:
		toT := CopyBranchStackOfGrowthCurve(mapOrigCopy, fromT)
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

	case *TopStackGrowthCurveEndArcShape:
		toT := CopyBranchTopStackGrowthCurveEndArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackGrowthCurveStartArcShape:
		toT := CopyBranchTopStackGrowthCurveStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackOfGrowthCurve:
		toT := CopyBranchTopStackOfGrowthCurve(mapOrigCopy, fromT)
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
	if growthcurve2dFrom.StartArcShapeGrid != nil {
		growthcurve2dTo.StartArcShapeGrid = CopyBranchStartArcShapeGrid(mapOrigCopy, growthcurve2dFrom.StartArcShapeGrid)
	}
	if growthcurve2dFrom.EndArcShapeGrid != nil {
		growthcurve2dTo.EndArcShapeGrid = CopyBranchEndArcShapeGrid(mapOrigCopy, growthcurve2dFrom.EndArcShapeGrid)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGrowthCurveBezierShape(mapOrigCopy map[any]any, growthcurvebeziershapeFrom *GrowthCurveBezierShape) (growthcurvebeziershapeTo *GrowthCurveBezierShape) {

	// growthcurvebeziershapeFrom has already been copied
	if _growthcurvebeziershapeTo, ok := mapOrigCopy[growthcurvebeziershapeFrom]; ok {
		growthcurvebeziershapeTo = _growthcurvebeziershapeTo.(*GrowthCurveBezierShape)
		return
	}

	growthcurvebeziershapeTo = new(GrowthCurveBezierShape)
	mapOrigCopy[growthcurvebeziershapeFrom] = growthcurvebeziershapeTo
	growthcurvebeziershapeFrom.CopyBasicFields(growthcurvebeziershapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGrowthCurveBezierShapeGrid(mapOrigCopy map[any]any, growthcurvebeziershapegridFrom *GrowthCurveBezierShapeGrid) (growthcurvebeziershapegridTo *GrowthCurveBezierShapeGrid) {

	// growthcurvebeziershapegridFrom has already been copied
	if _growthcurvebeziershapegridTo, ok := mapOrigCopy[growthcurvebeziershapegridFrom]; ok {
		growthcurvebeziershapegridTo = _growthcurvebeziershapegridTo.(*GrowthCurveBezierShapeGrid)
		return
	}

	growthcurvebeziershapegridTo = new(GrowthCurveBezierShapeGrid)
	mapOrigCopy[growthcurvebeziershapegridFrom] = growthcurvebeziershapegridTo
	growthcurvebeziershapegridFrom.CopyBasicFields(growthcurvebeziershapegridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _growthcurvebeziershape := range growthcurvebeziershapegridFrom.GrowthCurveBezierShapes {
		growthcurvebeziershapegridTo.GrowthCurveBezierShapes = append(growthcurvebeziershapegridTo.GrowthCurveBezierShapes, CopyBranchGrowthCurveBezierShape(mapOrigCopy, _growthcurvebeziershape))
	}

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

func CopyBranchNextCircleShape(mapOrigCopy map[any]any, nextcircleshapeFrom *NextCircleShape) (nextcircleshapeTo *NextCircleShape) {

	// nextcircleshapeFrom has already been copied
	if _nextcircleshapeTo, ok := mapOrigCopy[nextcircleshapeFrom]; ok {
		nextcircleshapeTo = _nextcircleshapeTo.(*NextCircleShape)
		return
	}

	nextcircleshapeTo = new(NextCircleShape)
	mapOrigCopy[nextcircleshapeFrom] = nextcircleshapeTo
	nextcircleshapeFrom.CopyBasicFields(nextcircleshapeTo)

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
	if plantFrom.ReferenceRhombus != nil {
		plantTo.ReferenceRhombus = CopyBranchRhombusShape(mapOrigCopy, plantFrom.ReferenceRhombus)
	}
	if plantFrom.PlantCircumferenceShape != nil {
		plantTo.PlantCircumferenceShape = CopyBranchPlantCircumferenceShape(mapOrigCopy, plantFrom.PlantCircumferenceShape)
	}
	if plantFrom.GridPathShape != nil {
		plantTo.GridPathShape = CopyBranchGridPathShape(mapOrigCopy, plantFrom.GridPathShape)
	}
	if plantFrom.InitialRhombusGridShape != nil {
		plantTo.InitialRhombusGridShape = CopyBranchInitialRhombusGridShape(mapOrigCopy, plantFrom.InitialRhombusGridShape)
	}
	if plantFrom.ExplanationTextShape != nil {
		plantTo.ExplanationTextShape = CopyBranchExplanationTextShape(mapOrigCopy, plantFrom.ExplanationTextShape)
	}
	if plantFrom.RotatedReferenceRhombus != nil {
		plantTo.RotatedReferenceRhombus = CopyBranchRhombusShape(mapOrigCopy, plantFrom.RotatedReferenceRhombus)
	}
	if plantFrom.RotatedPlantCircumferenceShape != nil {
		plantTo.RotatedPlantCircumferenceShape = CopyBranchPlantCircumferenceShape(mapOrigCopy, plantFrom.RotatedPlantCircumferenceShape)
	}
	if plantFrom.RotatedGridPathShape != nil {
		plantTo.RotatedGridPathShape = CopyBranchGridPathShape(mapOrigCopy, plantFrom.RotatedGridPathShape)
	}
	if plantFrom.RotatedRhombusGridShape2 != nil {
		plantTo.RotatedRhombusGridShape2 = CopyBranchRotatedRhombusGridShape(mapOrigCopy, plantFrom.RotatedRhombusGridShape2)
	}
	if plantFrom.GrowthCurveRhombusGridShape != nil {
		plantTo.GrowthCurveRhombusGridShape = CopyBranchGrowthCurveRhombusGridShape(mapOrigCopy, plantFrom.GrowthCurveRhombusGridShape)
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
	if plantFrom.EndArcShapeGrid != nil {
		plantTo.EndArcShapeGrid = CopyBranchEndArcShapeGrid(mapOrigCopy, plantFrom.EndArcShapeGrid)
	}
	if plantFrom.TopEndArcShapeGrid != nil {
		plantTo.TopEndArcShapeGrid = CopyBranchTopEndArcShapeGrid(mapOrigCopy, plantFrom.TopEndArcShapeGrid)
	}
	if plantFrom.GrowthCurveBezierShapeGrid != nil {
		plantTo.GrowthCurveBezierShapeGrid = CopyBranchGrowthCurveBezierShapeGrid(mapOrigCopy, plantFrom.GrowthCurveBezierShapeGrid)
	}
	if plantFrom.StackOfGrowthCurve != nil {
		plantTo.StackOfGrowthCurve = CopyBranchStackOfGrowthCurve(mapOrigCopy, plantFrom.StackOfGrowthCurve)
	}
	if plantFrom.TopStackOfGrowthCurve != nil {
		plantTo.TopStackOfGrowthCurve = CopyBranchTopStackOfGrowthCurve(mapOrigCopy, plantFrom.TopStackOfGrowthCurve)
	}
	if plantFrom.ShiftedLeftStackOfGrowthCurve != nil {
		plantTo.ShiftedLeftStackOfGrowthCurve = CopyBranchShiftedLeftStackOfGrowthCurve(mapOrigCopy, plantFrom.ShiftedLeftStackOfGrowthCurve)
	}
	if plantFrom.ShiftedLeftStackOfNormalVector != nil {
		plantTo.ShiftedLeftStackOfNormalVector = CopyBranchShiftedLeftStackOfNormalVector(mapOrigCopy, plantFrom.ShiftedLeftStackOfNormalVector)
	}
	if plantFrom.GrowthCurve2D != nil {
		plantTo.GrowthCurve2D = CopyBranchGrowthCurve2D(mapOrigCopy, plantFrom.GrowthCurve2D)
	}
	if plantFrom.TopGrowthCurve2D != nil {
		plantTo.TopGrowthCurve2D = CopyBranchTopGrowthCurve2D(mapOrigCopy, plantFrom.TopGrowthCurve2D)
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

func CopyBranchStackGrowthCurveEndArcShape(mapOrigCopy map[any]any, stackgrowthcurveendarcshapeFrom *StackGrowthCurveEndArcShape) (stackgrowthcurveendarcshapeTo *StackGrowthCurveEndArcShape) {

	// stackgrowthcurveendarcshapeFrom has already been copied
	if _stackgrowthcurveendarcshapeTo, ok := mapOrigCopy[stackgrowthcurveendarcshapeFrom]; ok {
		stackgrowthcurveendarcshapeTo = _stackgrowthcurveendarcshapeTo.(*StackGrowthCurveEndArcShape)
		return
	}

	stackgrowthcurveendarcshapeTo = new(StackGrowthCurveEndArcShape)
	mapOrigCopy[stackgrowthcurveendarcshapeFrom] = stackgrowthcurveendarcshapeTo
	stackgrowthcurveendarcshapeFrom.CopyBasicFields(stackgrowthcurveendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackGrowthCurveStartArcShape(mapOrigCopy map[any]any, stackgrowthcurvestartarcshapeFrom *StackGrowthCurveStartArcShape) (stackgrowthcurvestartarcshapeTo *StackGrowthCurveStartArcShape) {

	// stackgrowthcurvestartarcshapeFrom has already been copied
	if _stackgrowthcurvestartarcshapeTo, ok := mapOrigCopy[stackgrowthcurvestartarcshapeFrom]; ok {
		stackgrowthcurvestartarcshapeTo = _stackgrowthcurvestartarcshapeTo.(*StackGrowthCurveStartArcShape)
		return
	}

	stackgrowthcurvestartarcshapeTo = new(StackGrowthCurveStartArcShape)
	mapOrigCopy[stackgrowthcurvestartarcshapeFrom] = stackgrowthcurvestartarcshapeTo
	stackgrowthcurvestartarcshapeFrom.CopyBasicFields(stackgrowthcurvestartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackOfGrowthCurve(mapOrigCopy map[any]any, stackofgrowthcurveFrom *StackOfGrowthCurve) (stackofgrowthcurveTo *StackOfGrowthCurve) {

	// stackofgrowthcurveFrom has already been copied
	if _stackofgrowthcurveTo, ok := mapOrigCopy[stackofgrowthcurveFrom]; ok {
		stackofgrowthcurveTo = _stackofgrowthcurveTo.(*StackOfGrowthCurve)
		return
	}

	stackofgrowthcurveTo = new(StackOfGrowthCurve)
	mapOrigCopy[stackofgrowthcurveFrom] = stackofgrowthcurveTo
	stackofgrowthcurveFrom.CopyBasicFields(stackofgrowthcurveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stackgrowthcurvestartarcshape := range stackofgrowthcurveFrom.StackGrowthCurveStartArcShapes {
		stackofgrowthcurveTo.StackGrowthCurveStartArcShapes = append(stackofgrowthcurveTo.StackGrowthCurveStartArcShapes, CopyBranchStackGrowthCurveStartArcShape(mapOrigCopy, _stackgrowthcurvestartarcshape))
	}
	for _, _stackgrowthcurveendarcshape := range stackofgrowthcurveFrom.StackGrowthCurveEndArcShapes {
		stackofgrowthcurveTo.StackGrowthCurveEndArcShapes = append(stackofgrowthcurveTo.StackGrowthCurveEndArcShapes, CopyBranchStackGrowthCurveEndArcShape(mapOrigCopy, _stackgrowthcurveendarcshape))
	}

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
	if topgrowthcurve2dFrom.TopStartArcShapeGrid != nil {
		topgrowthcurve2dTo.TopStartArcShapeGrid = CopyBranchTopStartArcShapeGrid(mapOrigCopy, topgrowthcurve2dFrom.TopStartArcShapeGrid)
	}
	if topgrowthcurve2dFrom.TopEndArcShapeGrid != nil {
		topgrowthcurve2dTo.TopEndArcShapeGrid = CopyBranchTopEndArcShapeGrid(mapOrigCopy, topgrowthcurve2dFrom.TopEndArcShapeGrid)
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

func CopyBranchTopStackGrowthCurveEndArcShape(mapOrigCopy map[any]any, topstackgrowthcurveendarcshapeFrom *TopStackGrowthCurveEndArcShape) (topstackgrowthcurveendarcshapeTo *TopStackGrowthCurveEndArcShape) {

	// topstackgrowthcurveendarcshapeFrom has already been copied
	if _topstackgrowthcurveendarcshapeTo, ok := mapOrigCopy[topstackgrowthcurveendarcshapeFrom]; ok {
		topstackgrowthcurveendarcshapeTo = _topstackgrowthcurveendarcshapeTo.(*TopStackGrowthCurveEndArcShape)
		return
	}

	topstackgrowthcurveendarcshapeTo = new(TopStackGrowthCurveEndArcShape)
	mapOrigCopy[topstackgrowthcurveendarcshapeFrom] = topstackgrowthcurveendarcshapeTo
	topstackgrowthcurveendarcshapeFrom.CopyBasicFields(topstackgrowthcurveendarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStackGrowthCurveStartArcShape(mapOrigCopy map[any]any, topstackgrowthcurvestartarcshapeFrom *TopStackGrowthCurveStartArcShape) (topstackgrowthcurvestartarcshapeTo *TopStackGrowthCurveStartArcShape) {

	// topstackgrowthcurvestartarcshapeFrom has already been copied
	if _topstackgrowthcurvestartarcshapeTo, ok := mapOrigCopy[topstackgrowthcurvestartarcshapeFrom]; ok {
		topstackgrowthcurvestartarcshapeTo = _topstackgrowthcurvestartarcshapeTo.(*TopStackGrowthCurveStartArcShape)
		return
	}

	topstackgrowthcurvestartarcshapeTo = new(TopStackGrowthCurveStartArcShape)
	mapOrigCopy[topstackgrowthcurvestartarcshapeFrom] = topstackgrowthcurvestartarcshapeTo
	topstackgrowthcurvestartarcshapeFrom.CopyBasicFields(topstackgrowthcurvestartarcshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStackOfGrowthCurve(mapOrigCopy map[any]any, topstackofgrowthcurveFrom *TopStackOfGrowthCurve) (topstackofgrowthcurveTo *TopStackOfGrowthCurve) {

	// topstackofgrowthcurveFrom has already been copied
	if _topstackofgrowthcurveTo, ok := mapOrigCopy[topstackofgrowthcurveFrom]; ok {
		topstackofgrowthcurveTo = _topstackofgrowthcurveTo.(*TopStackOfGrowthCurve)
		return
	}

	topstackofgrowthcurveTo = new(TopStackOfGrowthCurve)
	mapOrigCopy[topstackofgrowthcurveFrom] = topstackofgrowthcurveTo
	topstackofgrowthcurveFrom.CopyBasicFields(topstackofgrowthcurveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topstackgrowthcurvestartarcshape := range topstackofgrowthcurveFrom.TopStackGrowthCurveStartArcShapes {
		topstackofgrowthcurveTo.TopStackGrowthCurveStartArcShapes = append(topstackofgrowthcurveTo.TopStackGrowthCurveStartArcShapes, CopyBranchTopStackGrowthCurveStartArcShape(mapOrigCopy, _topstackgrowthcurvestartarcshape))
	}
	for _, _topstackgrowthcurveendarcshape := range topstackofgrowthcurveFrom.TopStackGrowthCurveEndArcShapes {
		topstackofgrowthcurveTo.TopStackGrowthCurveEndArcShapes = append(topstackofgrowthcurveTo.TopStackGrowthCurveEndArcShapes, CopyBranchTopStackGrowthCurveEndArcShape(mapOrigCopy, _topstackgrowthcurveendarcshape))
	}

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

	case *GrowthCurveBezierShape:
		stage.UnstageBranchGrowthCurveBezierShape(target)

	case *GrowthCurveBezierShapeGrid:
		stage.UnstageBranchGrowthCurveBezierShapeGrid(target)

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

	case *NextCircleShape:
		stage.UnstageBranchNextCircleShape(target)

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

	case *StackGrowthCurveEndArcShape:
		stage.UnstageBranchStackGrowthCurveEndArcShape(target)

	case *StackGrowthCurveStartArcShape:
		stage.UnstageBranchStackGrowthCurveStartArcShape(target)

	case *StackOfGrowthCurve:
		stage.UnstageBranchStackOfGrowthCurve(target)

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

	case *TopStackGrowthCurveEndArcShape:
		stage.UnstageBranchTopStackGrowthCurveEndArcShape(target)

	case *TopStackGrowthCurveStartArcShape:
		stage.UnstageBranchTopStackGrowthCurveStartArcShape(target)

	case *TopStackOfGrowthCurve:
		stage.UnstageBranchTopStackOfGrowthCurve(target)

	case *TopStartArcShape:
		stage.UnstageBranchTopStartArcShape(target)

	case *TopStartArcShapeGrid:
		stage.UnstageBranchTopStartArcShapeGrid(target)

	case *TopStartHalfwayArcShape:
		stage.UnstageBranchTopStartHalfwayArcShape(target)

	case *TopStartHalfwayArcShapeGrid:
		stage.UnstageBranchTopStartHalfwayArcShapeGrid(target)

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
	if growthcurve2d.StartArcShapeGrid != nil {
		UnstageBranch(stage, growthcurve2d.StartArcShapeGrid)
	}
	if growthcurve2d.EndArcShapeGrid != nil {
		UnstageBranch(stage, growthcurve2d.EndArcShapeGrid)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGrowthCurveBezierShape(growthcurvebeziershape *GrowthCurveBezierShape) {

	// check if instance is already staged
	if !IsStaged(stage, growthcurvebeziershape) {
		return
	}

	growthcurvebeziershape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGrowthCurveBezierShapeGrid(growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) {

	// check if instance is already staged
	if !IsStaged(stage, growthcurvebeziershapegrid) {
		return
	}

	growthcurvebeziershapegrid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _growthcurvebeziershape := range growthcurvebeziershapegrid.GrowthCurveBezierShapes {
		UnstageBranch(stage, _growthcurvebeziershape)
	}

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

func (stage *Stage) UnstageBranchNextCircleShape(nextcircleshape *NextCircleShape) {

	// check if instance is already staged
	if !IsStaged(stage, nextcircleshape) {
		return
	}

	nextcircleshape.Unstage(stage)

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
	if plant.ReferenceRhombus != nil {
		UnstageBranch(stage, plant.ReferenceRhombus)
	}
	if plant.PlantCircumferenceShape != nil {
		UnstageBranch(stage, plant.PlantCircumferenceShape)
	}
	if plant.GridPathShape != nil {
		UnstageBranch(stage, plant.GridPathShape)
	}
	if plant.InitialRhombusGridShape != nil {
		UnstageBranch(stage, plant.InitialRhombusGridShape)
	}
	if plant.ExplanationTextShape != nil {
		UnstageBranch(stage, plant.ExplanationTextShape)
	}
	if plant.RotatedReferenceRhombus != nil {
		UnstageBranch(stage, plant.RotatedReferenceRhombus)
	}
	if plant.RotatedPlantCircumferenceShape != nil {
		UnstageBranch(stage, plant.RotatedPlantCircumferenceShape)
	}
	if plant.RotatedGridPathShape != nil {
		UnstageBranch(stage, plant.RotatedGridPathShape)
	}
	if plant.RotatedRhombusGridShape2 != nil {
		UnstageBranch(stage, plant.RotatedRhombusGridShape2)
	}
	if plant.GrowthCurveRhombusGridShape != nil {
		UnstageBranch(stage, plant.GrowthCurveRhombusGridShape)
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
	if plant.EndArcShapeGrid != nil {
		UnstageBranch(stage, plant.EndArcShapeGrid)
	}
	if plant.TopEndArcShapeGrid != nil {
		UnstageBranch(stage, plant.TopEndArcShapeGrid)
	}
	if plant.GrowthCurveBezierShapeGrid != nil {
		UnstageBranch(stage, plant.GrowthCurveBezierShapeGrid)
	}
	if plant.StackOfGrowthCurve != nil {
		UnstageBranch(stage, plant.StackOfGrowthCurve)
	}
	if plant.TopStackOfGrowthCurve != nil {
		UnstageBranch(stage, plant.TopStackOfGrowthCurve)
	}
	if plant.ShiftedLeftStackOfGrowthCurve != nil {
		UnstageBranch(stage, plant.ShiftedLeftStackOfGrowthCurve)
	}
	if plant.ShiftedLeftStackOfNormalVector != nil {
		UnstageBranch(stage, plant.ShiftedLeftStackOfNormalVector)
	}
	if plant.GrowthCurve2D != nil {
		UnstageBranch(stage, plant.GrowthCurve2D)
	}
	if plant.TopGrowthCurve2D != nil {
		UnstageBranch(stage, plant.TopGrowthCurve2D)
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

func (stage *Stage) UnstageBranchStackGrowthCurveEndArcShape(stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, stackgrowthcurveendarcshape) {
		return
	}

	stackgrowthcurveendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackGrowthCurveStartArcShape(stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, stackgrowthcurvestartarcshape) {
		return
	}

	stackgrowthcurvestartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackOfGrowthCurve(stackofgrowthcurve *StackOfGrowthCurve) {

	// check if instance is already staged
	if !IsStaged(stage, stackofgrowthcurve) {
		return
	}

	stackofgrowthcurve.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stackgrowthcurvestartarcshape := range stackofgrowthcurve.StackGrowthCurveStartArcShapes {
		UnstageBranch(stage, _stackgrowthcurvestartarcshape)
	}
	for _, _stackgrowthcurveendarcshape := range stackofgrowthcurve.StackGrowthCurveEndArcShapes {
		UnstageBranch(stage, _stackgrowthcurveendarcshape)
	}

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
	if topgrowthcurve2d.TopStartArcShapeGrid != nil {
		UnstageBranch(stage, topgrowthcurve2d.TopStartArcShapeGrid)
	}
	if topgrowthcurve2d.TopEndArcShapeGrid != nil {
		UnstageBranch(stage, topgrowthcurve2d.TopEndArcShapeGrid)
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

func (stage *Stage) UnstageBranchTopStackGrowthCurveEndArcShape(topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, topstackgrowthcurveendarcshape) {
		return
	}

	topstackgrowthcurveendarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStackGrowthCurveStartArcShape(topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) {

	// check if instance is already staged
	if !IsStaged(stage, topstackgrowthcurvestartarcshape) {
		return
	}

	topstackgrowthcurvestartarcshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStackOfGrowthCurve(topstackofgrowthcurve *TopStackOfGrowthCurve) {

	// check if instance is already staged
	if !IsStaged(stage, topstackofgrowthcurve) {
		return
	}

	topstackofgrowthcurve.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topstackgrowthcurvestartarcshape := range topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes {
		UnstageBranch(stage, _topstackgrowthcurvestartarcshape)
	}
	for _, _topstackgrowthcurveendarcshape := range topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes {
		UnstageBranch(stage, _topstackgrowthcurveendarcshape)
	}

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
	if instance.StartArcShapeGrid != nil {
		reference.StartArcShapeGrid = stage.StartArcShapeGrids_reference[instance.StartArcShapeGrid]
	}
	if instance.EndArcShapeGrid != nil {
		reference.EndArcShapeGrid = stage.EndArcShapeGrids_reference[instance.EndArcShapeGrid]
	}
	// insertion point for slice of pointers field
}

func (reference *GrowthCurveBezierShape) GongReconstructPointersFromReferences(stage *Stage, instance *GrowthCurveBezierShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GrowthCurveBezierShapeGrid) GongReconstructPointersFromReferences(stage *Stage, instance *GrowthCurveBezierShapeGrid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.GrowthCurveBezierShapes = reference.GrowthCurveBezierShapes[:0]
	for _, _b := range instance.GrowthCurveBezierShapes {
		reference.GrowthCurveBezierShapes = append(reference.GrowthCurveBezierShapes, stage.GrowthCurveBezierShapes_reference[_b])
	}
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

func (reference *NextCircleShape) GongReconstructPointersFromReferences(stage *Stage, instance *NextCircleShape) {
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
	if instance.EndArcShapeGrid != nil {
		reference.EndArcShapeGrid = stage.EndArcShapeGrids_reference[instance.EndArcShapeGrid]
	}
	if instance.TopEndArcShapeGrid != nil {
		reference.TopEndArcShapeGrid = stage.TopEndArcShapeGrids_reference[instance.TopEndArcShapeGrid]
	}
	if instance.GrowthCurveBezierShapeGrid != nil {
		reference.GrowthCurveBezierShapeGrid = stage.GrowthCurveBezierShapeGrids_reference[instance.GrowthCurveBezierShapeGrid]
	}
	if instance.StackOfGrowthCurve != nil {
		reference.StackOfGrowthCurve = stage.StackOfGrowthCurves_reference[instance.StackOfGrowthCurve]
	}
	if instance.TopStackOfGrowthCurve != nil {
		reference.TopStackOfGrowthCurve = stage.TopStackOfGrowthCurves_reference[instance.TopStackOfGrowthCurve]
	}
	if instance.ShiftedLeftStackOfGrowthCurve != nil {
		reference.ShiftedLeftStackOfGrowthCurve = stage.ShiftedLeftStackOfGrowthCurves_reference[instance.ShiftedLeftStackOfGrowthCurve]
	}
	if instance.ShiftedLeftStackOfNormalVector != nil {
		reference.ShiftedLeftStackOfNormalVector = stage.ShiftedLeftStackOfNormalVectors_reference[instance.ShiftedLeftStackOfNormalVector]
	}
	if instance.GrowthCurve2D != nil {
		reference.GrowthCurve2D = stage.GrowthCurve2Ds_reference[instance.GrowthCurve2D]
	}
	if instance.TopGrowthCurve2D != nil {
		reference.TopGrowthCurve2D = stage.TopGrowthCurve2Ds_reference[instance.TopGrowthCurve2D]
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

func (reference *StackGrowthCurveEndArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackGrowthCurveEndArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackGrowthCurveStartArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackGrowthCurveStartArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackOfGrowthCurve) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfGrowthCurve) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.StackGrowthCurveStartArcShapes = reference.StackGrowthCurveStartArcShapes[:0]
	for _, _b := range instance.StackGrowthCurveStartArcShapes {
		reference.StackGrowthCurveStartArcShapes = append(reference.StackGrowthCurveStartArcShapes, stage.StackGrowthCurveStartArcShapes_reference[_b])
	}
	reference.StackGrowthCurveEndArcShapes = reference.StackGrowthCurveEndArcShapes[:0]
	for _, _b := range instance.StackGrowthCurveEndArcShapes {
		reference.StackGrowthCurveEndArcShapes = append(reference.StackGrowthCurveEndArcShapes, stage.StackGrowthCurveEndArcShapes_reference[_b])
	}
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
	if instance.TopStartArcShapeGrid != nil {
		reference.TopStartArcShapeGrid = stage.TopStartArcShapeGrids_reference[instance.TopStartArcShapeGrid]
	}
	if instance.TopEndArcShapeGrid != nil {
		reference.TopEndArcShapeGrid = stage.TopEndArcShapeGrids_reference[instance.TopEndArcShapeGrid]
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

func (reference *TopStackGrowthCurveEndArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopStackGrowthCurveEndArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStackGrowthCurveStartArcShape) GongReconstructPointersFromReferences(stage *Stage, instance *TopStackGrowthCurveStartArcShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStackOfGrowthCurve) GongReconstructPointersFromReferences(stage *Stage, instance *TopStackOfGrowthCurve) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.TopStackGrowthCurveStartArcShapes = reference.TopStackGrowthCurveStartArcShapes[:0]
	for _, _b := range instance.TopStackGrowthCurveStartArcShapes {
		reference.TopStackGrowthCurveStartArcShapes = append(reference.TopStackGrowthCurveStartArcShapes, stage.TopStackGrowthCurveStartArcShapes_reference[_b])
	}
	reference.TopStackGrowthCurveEndArcShapes = reference.TopStackGrowthCurveEndArcShapes[:0]
	for _, _b := range instance.TopStackGrowthCurveEndArcShapes {
		reference.TopStackGrowthCurveEndArcShapes = append(reference.TopStackGrowthCurveEndArcShapes, stage.TopStackGrowthCurveEndArcShapes_reference[_b])
	}
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
	if _reference := reference.StartArcShapeGrid; _reference != nil {
		reference.StartArcShapeGrid = nil
		if _instance, ok := stage.StartArcShapeGrids_instance[_reference]; ok {
			reference.StartArcShapeGrid = _instance
		}
	}
	if _reference := reference.EndArcShapeGrid; _reference != nil {
		reference.EndArcShapeGrid = nil
		if _instance, ok := stage.EndArcShapeGrids_instance[_reference]; ok {
			reference.EndArcShapeGrid = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *GrowthCurveBezierShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GrowthCurveBezierShapeGrid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _GrowthCurveBezierShapes []*GrowthCurveBezierShape
	for _, _reference := range reference.GrowthCurveBezierShapes {
		if _instance, ok := stage.GrowthCurveBezierShapes_instance[_reference]; ok {
			_GrowthCurveBezierShapes = append(_GrowthCurveBezierShapes, _instance)
		}
	}
	reference.GrowthCurveBezierShapes = _GrowthCurveBezierShapes
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

func (reference *NextCircleShape) GongReconstructPointersFromInstances(stage *Stage) {
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
	if _reference := reference.GrowthCurveBezierShapeGrid; _reference != nil {
		reference.GrowthCurveBezierShapeGrid = nil
		if _instance, ok := stage.GrowthCurveBezierShapeGrids_instance[_reference]; ok {
			reference.GrowthCurveBezierShapeGrid = _instance
		}
	}
	if _reference := reference.StackOfGrowthCurve; _reference != nil {
		reference.StackOfGrowthCurve = nil
		if _instance, ok := stage.StackOfGrowthCurves_instance[_reference]; ok {
			reference.StackOfGrowthCurve = _instance
		}
	}
	if _reference := reference.TopStackOfGrowthCurve; _reference != nil {
		reference.TopStackOfGrowthCurve = nil
		if _instance, ok := stage.TopStackOfGrowthCurves_instance[_reference]; ok {
			reference.TopStackOfGrowthCurve = _instance
		}
	}
	if _reference := reference.ShiftedLeftStackOfGrowthCurve; _reference != nil {
		reference.ShiftedLeftStackOfGrowthCurve = nil
		if _instance, ok := stage.ShiftedLeftStackOfGrowthCurves_instance[_reference]; ok {
			reference.ShiftedLeftStackOfGrowthCurve = _instance
		}
	}
	if _reference := reference.ShiftedLeftStackOfNormalVector; _reference != nil {
		reference.ShiftedLeftStackOfNormalVector = nil
		if _instance, ok := stage.ShiftedLeftStackOfNormalVectors_instance[_reference]; ok {
			reference.ShiftedLeftStackOfNormalVector = _instance
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

func (reference *StackGrowthCurveEndArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackGrowthCurveStartArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackOfGrowthCurve) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _StackGrowthCurveStartArcShapes []*StackGrowthCurveStartArcShape
	for _, _reference := range reference.StackGrowthCurveStartArcShapes {
		if _instance, ok := stage.StackGrowthCurveStartArcShapes_instance[_reference]; ok {
			_StackGrowthCurveStartArcShapes = append(_StackGrowthCurveStartArcShapes, _instance)
		}
	}
	reference.StackGrowthCurveStartArcShapes = _StackGrowthCurveStartArcShapes
	var _StackGrowthCurveEndArcShapes []*StackGrowthCurveEndArcShape
	for _, _reference := range reference.StackGrowthCurveEndArcShapes {
		if _instance, ok := stage.StackGrowthCurveEndArcShapes_instance[_reference]; ok {
			_StackGrowthCurveEndArcShapes = append(_StackGrowthCurveEndArcShapes, _instance)
		}
	}
	reference.StackGrowthCurveEndArcShapes = _StackGrowthCurveEndArcShapes
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
	if _reference := reference.TopStartArcShapeGrid; _reference != nil {
		reference.TopStartArcShapeGrid = nil
		if _instance, ok := stage.TopStartArcShapeGrids_instance[_reference]; ok {
			reference.TopStartArcShapeGrid = _instance
		}
	}
	if _reference := reference.TopEndArcShapeGrid; _reference != nil {
		reference.TopEndArcShapeGrid = nil
		if _instance, ok := stage.TopEndArcShapeGrids_instance[_reference]; ok {
			reference.TopEndArcShapeGrid = _instance
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

func (reference *TopStackGrowthCurveEndArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStackGrowthCurveStartArcShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStackOfGrowthCurve) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _TopStackGrowthCurveStartArcShapes []*TopStackGrowthCurveStartArcShape
	for _, _reference := range reference.TopStackGrowthCurveStartArcShapes {
		if _instance, ok := stage.TopStackGrowthCurveStartArcShapes_instance[_reference]; ok {
			_TopStackGrowthCurveStartArcShapes = append(_TopStackGrowthCurveStartArcShapes, _instance)
		}
	}
	reference.TopStackGrowthCurveStartArcShapes = _TopStackGrowthCurveStartArcShapes
	var _TopStackGrowthCurveEndArcShapes []*TopStackGrowthCurveEndArcShape
	for _, _reference := range reference.TopStackGrowthCurveEndArcShapes {
		if _instance, ok := stage.TopStackGrowthCurveEndArcShapes_instance[_reference]; ok {
			_TopStackGrowthCurveEndArcShapes = append(_TopStackGrowthCurveEndArcShapes, _instance)
		}
	}
	reference.TopStackGrowthCurveEndArcShapes = _TopStackGrowthCurveEndArcShapes
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
	if (growthcurve2d.StartArcShapeGrid == nil) != (growthcurve2dOther.StartArcShapeGrid == nil) {
		diffs = append(diffs, growthcurve2d.GongMarshallField(stage, "StartArcShapeGrid"))
	} else if growthcurve2d.StartArcShapeGrid != nil && growthcurve2dOther.StartArcShapeGrid != nil {
		if growthcurve2d.StartArcShapeGrid != growthcurve2dOther.StartArcShapeGrid {
			diffs = append(diffs, growthcurve2d.GongMarshallField(stage, "StartArcShapeGrid"))
		}
	}
	if (growthcurve2d.EndArcShapeGrid == nil) != (growthcurve2dOther.EndArcShapeGrid == nil) {
		diffs = append(diffs, growthcurve2d.GongMarshallField(stage, "EndArcShapeGrid"))
	} else if growthcurve2d.EndArcShapeGrid != nil && growthcurve2dOther.EndArcShapeGrid != nil {
		if growthcurve2d.EndArcShapeGrid != growthcurve2dOther.EndArcShapeGrid {
			diffs = append(diffs, growthcurve2d.GongMarshallField(stage, "EndArcShapeGrid"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (growthcurvebeziershape *GrowthCurveBezierShape) GongDiff(stage *Stage, growthcurvebeziershapeOther *GrowthCurveBezierShape) (diffs []string) {
	// insertion point for field diffs
	if growthcurvebeziershape.Name != growthcurvebeziershapeOther.Name {
		diffs = append(diffs, growthcurvebeziershape.GongMarshallField(stage, "Name"))
	}
	if growthcurvebeziershape.StartX != growthcurvebeziershapeOther.StartX {
		diffs = append(diffs, growthcurvebeziershape.GongMarshallField(stage, "StartX"))
	}
	if growthcurvebeziershape.StartY != growthcurvebeziershapeOther.StartY {
		diffs = append(diffs, growthcurvebeziershape.GongMarshallField(stage, "StartY"))
	}
	if growthcurvebeziershape.ControlPointStartX != growthcurvebeziershapeOther.ControlPointStartX {
		diffs = append(diffs, growthcurvebeziershape.GongMarshallField(stage, "ControlPointStartX"))
	}
	if growthcurvebeziershape.ControlPointStartY != growthcurvebeziershapeOther.ControlPointStartY {
		diffs = append(diffs, growthcurvebeziershape.GongMarshallField(stage, "ControlPointStartY"))
	}
	if growthcurvebeziershape.EndX != growthcurvebeziershapeOther.EndX {
		diffs = append(diffs, growthcurvebeziershape.GongMarshallField(stage, "EndX"))
	}
	if growthcurvebeziershape.EndY != growthcurvebeziershapeOther.EndY {
		diffs = append(diffs, growthcurvebeziershape.GongMarshallField(stage, "EndY"))
	}
	if growthcurvebeziershape.ControlPointEndX != growthcurvebeziershapeOther.ControlPointEndX {
		diffs = append(diffs, growthcurvebeziershape.GongMarshallField(stage, "ControlPointEndX"))
	}
	if growthcurvebeziershape.ControlPointEndY != growthcurvebeziershapeOther.ControlPointEndY {
		diffs = append(diffs, growthcurvebeziershape.GongMarshallField(stage, "ControlPointEndY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GongDiff(stage *Stage, growthcurvebeziershapegridOther *GrowthCurveBezierShapeGrid) (diffs []string) {
	// insertion point for field diffs
	if growthcurvebeziershapegrid.Name != growthcurvebeziershapegridOther.Name {
		diffs = append(diffs, growthcurvebeziershapegrid.GongMarshallField(stage, "Name"))
	}
	GrowthCurveBezierShapesDifferent := false
	if len(growthcurvebeziershapegrid.GrowthCurveBezierShapes) != len(growthcurvebeziershapegridOther.GrowthCurveBezierShapes) {
		GrowthCurveBezierShapesDifferent = true
	} else {
		for i := range growthcurvebeziershapegrid.GrowthCurveBezierShapes {
			if (growthcurvebeziershapegrid.GrowthCurveBezierShapes[i] == nil) != (growthcurvebeziershapegridOther.GrowthCurveBezierShapes[i] == nil) {
				GrowthCurveBezierShapesDifferent = true
				break
			} else if growthcurvebeziershapegrid.GrowthCurveBezierShapes[i] != nil && growthcurvebeziershapegridOther.GrowthCurveBezierShapes[i] != nil {
				// this is a pointer comparaison
				if growthcurvebeziershapegrid.GrowthCurveBezierShapes[i] != growthcurvebeziershapegridOther.GrowthCurveBezierShapes[i] {
					GrowthCurveBezierShapesDifferent = true
					break
				}
			}
		}
	}
	if GrowthCurveBezierShapesDifferent {
		ops := Diff(stage, growthcurvebeziershapegrid, growthcurvebeziershapegridOther, "GrowthCurveBezierShapes", growthcurvebeziershapegridOther.GrowthCurveBezierShapes, growthcurvebeziershapegrid.GrowthCurveBezierShapes)
		diffs = append(diffs, ops)
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
func (nextcircleshape *NextCircleShape) GongDiff(stage *Stage, nextcircleshapeOther *NextCircleShape) (diffs []string) {
	// insertion point for field diffs
	if nextcircleshape.Name != nextcircleshapeOther.Name {
		diffs = append(diffs, nextcircleshape.GongMarshallField(stage, "Name"))
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
	if plant.RadialThickness != plantOther.RadialThickness {
		diffs = append(diffs, plant.GongMarshallField(stage, "RadialThickness"))
	}
	if plant.RhombusSideLength != plantOther.RhombusSideLength {
		diffs = append(diffs, plant.GongMarshallField(stage, "RhombusSideLength"))
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
	if (plant.ReferenceRhombus == nil) != (plantOther.ReferenceRhombus == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "ReferenceRhombus"))
	} else if plant.ReferenceRhombus != nil && plantOther.ReferenceRhombus != nil {
		if plant.ReferenceRhombus != plantOther.ReferenceRhombus {
			diffs = append(diffs, plant.GongMarshallField(stage, "ReferenceRhombus"))
		}
	}
	if (plant.PlantCircumferenceShape == nil) != (plantOther.PlantCircumferenceShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "PlantCircumferenceShape"))
	} else if plant.PlantCircumferenceShape != nil && plantOther.PlantCircumferenceShape != nil {
		if plant.PlantCircumferenceShape != plantOther.PlantCircumferenceShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "PlantCircumferenceShape"))
		}
	}
	if (plant.GridPathShape == nil) != (plantOther.GridPathShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "GridPathShape"))
	} else if plant.GridPathShape != nil && plantOther.GridPathShape != nil {
		if plant.GridPathShape != plantOther.GridPathShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "GridPathShape"))
		}
	}
	if (plant.InitialRhombusGridShape == nil) != (plantOther.InitialRhombusGridShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "InitialRhombusGridShape"))
	} else if plant.InitialRhombusGridShape != nil && plantOther.InitialRhombusGridShape != nil {
		if plant.InitialRhombusGridShape != plantOther.InitialRhombusGridShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "InitialRhombusGridShape"))
		}
	}
	if (plant.ExplanationTextShape == nil) != (plantOther.ExplanationTextShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "ExplanationTextShape"))
	} else if plant.ExplanationTextShape != nil && plantOther.ExplanationTextShape != nil {
		if plant.ExplanationTextShape != plantOther.ExplanationTextShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "ExplanationTextShape"))
		}
	}
	if (plant.RotatedReferenceRhombus == nil) != (plantOther.RotatedReferenceRhombus == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "RotatedReferenceRhombus"))
	} else if plant.RotatedReferenceRhombus != nil && plantOther.RotatedReferenceRhombus != nil {
		if plant.RotatedReferenceRhombus != plantOther.RotatedReferenceRhombus {
			diffs = append(diffs, plant.GongMarshallField(stage, "RotatedReferenceRhombus"))
		}
	}
	if (plant.RotatedPlantCircumferenceShape == nil) != (plantOther.RotatedPlantCircumferenceShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "RotatedPlantCircumferenceShape"))
	} else if plant.RotatedPlantCircumferenceShape != nil && plantOther.RotatedPlantCircumferenceShape != nil {
		if plant.RotatedPlantCircumferenceShape != plantOther.RotatedPlantCircumferenceShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "RotatedPlantCircumferenceShape"))
		}
	}
	if (plant.RotatedGridPathShape == nil) != (plantOther.RotatedGridPathShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "RotatedGridPathShape"))
	} else if plant.RotatedGridPathShape != nil && plantOther.RotatedGridPathShape != nil {
		if plant.RotatedGridPathShape != plantOther.RotatedGridPathShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "RotatedGridPathShape"))
		}
	}
	if (plant.RotatedRhombusGridShape2 == nil) != (plantOther.RotatedRhombusGridShape2 == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "RotatedRhombusGridShape2"))
	} else if plant.RotatedRhombusGridShape2 != nil && plantOther.RotatedRhombusGridShape2 != nil {
		if plant.RotatedRhombusGridShape2 != plantOther.RotatedRhombusGridShape2 {
			diffs = append(diffs, plant.GongMarshallField(stage, "RotatedRhombusGridShape2"))
		}
	}
	if (plant.GrowthCurveRhombusGridShape == nil) != (plantOther.GrowthCurveRhombusGridShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "GrowthCurveRhombusGridShape"))
	} else if plant.GrowthCurveRhombusGridShape != nil && plantOther.GrowthCurveRhombusGridShape != nil {
		if plant.GrowthCurveRhombusGridShape != plantOther.GrowthCurveRhombusGridShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "GrowthCurveRhombusGridShape"))
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
	if (plant.GrowthCurveBezierShapeGrid == nil) != (plantOther.GrowthCurveBezierShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "GrowthCurveBezierShapeGrid"))
	} else if plant.GrowthCurveBezierShapeGrid != nil && plantOther.GrowthCurveBezierShapeGrid != nil {
		if plant.GrowthCurveBezierShapeGrid != plantOther.GrowthCurveBezierShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "GrowthCurveBezierShapeGrid"))
		}
	}
	if (plant.StackOfGrowthCurve == nil) != (plantOther.StackOfGrowthCurve == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "StackOfGrowthCurve"))
	} else if plant.StackOfGrowthCurve != nil && plantOther.StackOfGrowthCurve != nil {
		if plant.StackOfGrowthCurve != plantOther.StackOfGrowthCurve {
			diffs = append(diffs, plant.GongMarshallField(stage, "StackOfGrowthCurve"))
		}
	}
	if (plant.TopStackOfGrowthCurve == nil) != (plantOther.TopStackOfGrowthCurve == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "TopStackOfGrowthCurve"))
	} else if plant.TopStackOfGrowthCurve != nil && plantOther.TopStackOfGrowthCurve != nil {
		if plant.TopStackOfGrowthCurve != plantOther.TopStackOfGrowthCurve {
			diffs = append(diffs, plant.GongMarshallField(stage, "TopStackOfGrowthCurve"))
		}
	}
	if (plant.ShiftedLeftStackOfGrowthCurve == nil) != (plantOther.ShiftedLeftStackOfGrowthCurve == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "ShiftedLeftStackOfGrowthCurve"))
	} else if plant.ShiftedLeftStackOfGrowthCurve != nil && plantOther.ShiftedLeftStackOfGrowthCurve != nil {
		if plant.ShiftedLeftStackOfGrowthCurve != plantOther.ShiftedLeftStackOfGrowthCurve {
			diffs = append(diffs, plant.GongMarshallField(stage, "ShiftedLeftStackOfGrowthCurve"))
		}
	}
	if (plant.ShiftedLeftStackOfNormalVector == nil) != (plantOther.ShiftedLeftStackOfNormalVector == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "ShiftedLeftStackOfNormalVector"))
	} else if plant.ShiftedLeftStackOfNormalVector != nil && plantOther.ShiftedLeftStackOfNormalVector != nil {
		if plant.ShiftedLeftStackOfNormalVector != plantOther.ShiftedLeftStackOfNormalVector {
			diffs = append(diffs, plant.GongMarshallField(stage, "ShiftedLeftStackOfNormalVector"))
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
	if plantdiagram.IsHiddenGrowthCurveBezierShapeGrid != plantdiagramOther.IsHiddenGrowthCurveBezierShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenGrowthCurveBezierShapeGrid"))
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
func (stackgrowthcurveendarcshape *StackGrowthCurveEndArcShape) GongDiff(stage *Stage, stackgrowthcurveendarcshapeOther *StackGrowthCurveEndArcShape) (diffs []string) {
	// insertion point for field diffs
	if stackgrowthcurveendarcshape.Name != stackgrowthcurveendarcshapeOther.Name {
		diffs = append(diffs, stackgrowthcurveendarcshape.GongMarshallField(stage, "Name"))
	}
	if stackgrowthcurveendarcshape.StartX != stackgrowthcurveendarcshapeOther.StartX {
		diffs = append(diffs, stackgrowthcurveendarcshape.GongMarshallField(stage, "StartX"))
	}
	if stackgrowthcurveendarcshape.StartY != stackgrowthcurveendarcshapeOther.StartY {
		diffs = append(diffs, stackgrowthcurveendarcshape.GongMarshallField(stage, "StartY"))
	}
	if stackgrowthcurveendarcshape.EndX != stackgrowthcurveendarcshapeOther.EndX {
		diffs = append(diffs, stackgrowthcurveendarcshape.GongMarshallField(stage, "EndX"))
	}
	if stackgrowthcurveendarcshape.EndY != stackgrowthcurveendarcshapeOther.EndY {
		diffs = append(diffs, stackgrowthcurveendarcshape.GongMarshallField(stage, "EndY"))
	}
	if stackgrowthcurveendarcshape.XAxisRotation != stackgrowthcurveendarcshapeOther.XAxisRotation {
		diffs = append(diffs, stackgrowthcurveendarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if stackgrowthcurveendarcshape.LargeArcFlag != stackgrowthcurveendarcshapeOther.LargeArcFlag {
		diffs = append(diffs, stackgrowthcurveendarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if stackgrowthcurveendarcshape.SweepFlag != stackgrowthcurveendarcshapeOther.SweepFlag {
		diffs = append(diffs, stackgrowthcurveendarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if stackgrowthcurveendarcshape.RadiusX != stackgrowthcurveendarcshapeOther.RadiusX {
		diffs = append(diffs, stackgrowthcurveendarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if stackgrowthcurveendarcshape.RadiusY != stackgrowthcurveendarcshapeOther.RadiusY {
		diffs = append(diffs, stackgrowthcurveendarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackgrowthcurvestartarcshape *StackGrowthCurveStartArcShape) GongDiff(stage *Stage, stackgrowthcurvestartarcshapeOther *StackGrowthCurveStartArcShape) (diffs []string) {
	// insertion point for field diffs
	if stackgrowthcurvestartarcshape.Name != stackgrowthcurvestartarcshapeOther.Name {
		diffs = append(diffs, stackgrowthcurvestartarcshape.GongMarshallField(stage, "Name"))
	}
	if stackgrowthcurvestartarcshape.StartX != stackgrowthcurvestartarcshapeOther.StartX {
		diffs = append(diffs, stackgrowthcurvestartarcshape.GongMarshallField(stage, "StartX"))
	}
	if stackgrowthcurvestartarcshape.StartY != stackgrowthcurvestartarcshapeOther.StartY {
		diffs = append(diffs, stackgrowthcurvestartarcshape.GongMarshallField(stage, "StartY"))
	}
	if stackgrowthcurvestartarcshape.EndX != stackgrowthcurvestartarcshapeOther.EndX {
		diffs = append(diffs, stackgrowthcurvestartarcshape.GongMarshallField(stage, "EndX"))
	}
	if stackgrowthcurvestartarcshape.EndY != stackgrowthcurvestartarcshapeOther.EndY {
		diffs = append(diffs, stackgrowthcurvestartarcshape.GongMarshallField(stage, "EndY"))
	}
	if stackgrowthcurvestartarcshape.XAxisRotation != stackgrowthcurvestartarcshapeOther.XAxisRotation {
		diffs = append(diffs, stackgrowthcurvestartarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if stackgrowthcurvestartarcshape.LargeArcFlag != stackgrowthcurvestartarcshapeOther.LargeArcFlag {
		diffs = append(diffs, stackgrowthcurvestartarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if stackgrowthcurvestartarcshape.SweepFlag != stackgrowthcurvestartarcshapeOther.SweepFlag {
		diffs = append(diffs, stackgrowthcurvestartarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if stackgrowthcurvestartarcshape.RadiusX != stackgrowthcurvestartarcshapeOther.RadiusX {
		diffs = append(diffs, stackgrowthcurvestartarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if stackgrowthcurvestartarcshape.RadiusY != stackgrowthcurvestartarcshapeOther.RadiusY {
		diffs = append(diffs, stackgrowthcurvestartarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackofgrowthcurve *StackOfGrowthCurve) GongDiff(stage *Stage, stackofgrowthcurveOther *StackOfGrowthCurve) (diffs []string) {
	// insertion point for field diffs
	if stackofgrowthcurve.Name != stackofgrowthcurveOther.Name {
		diffs = append(diffs, stackofgrowthcurve.GongMarshallField(stage, "Name"))
	}
	StackGrowthCurveStartArcShapesDifferent := false
	if len(stackofgrowthcurve.StackGrowthCurveStartArcShapes) != len(stackofgrowthcurveOther.StackGrowthCurveStartArcShapes) {
		StackGrowthCurveStartArcShapesDifferent = true
	} else {
		for i := range stackofgrowthcurve.StackGrowthCurveStartArcShapes {
			if (stackofgrowthcurve.StackGrowthCurveStartArcShapes[i] == nil) != (stackofgrowthcurveOther.StackGrowthCurveStartArcShapes[i] == nil) {
				StackGrowthCurveStartArcShapesDifferent = true
				break
			} else if stackofgrowthcurve.StackGrowthCurveStartArcShapes[i] != nil && stackofgrowthcurveOther.StackGrowthCurveStartArcShapes[i] != nil {
				// this is a pointer comparaison
				if stackofgrowthcurve.StackGrowthCurveStartArcShapes[i] != stackofgrowthcurveOther.StackGrowthCurveStartArcShapes[i] {
					StackGrowthCurveStartArcShapesDifferent = true
					break
				}
			}
		}
	}
	if StackGrowthCurveStartArcShapesDifferent {
		ops := Diff(stage, stackofgrowthcurve, stackofgrowthcurveOther, "StackGrowthCurveStartArcShapes", stackofgrowthcurveOther.StackGrowthCurveStartArcShapes, stackofgrowthcurve.StackGrowthCurveStartArcShapes)
		diffs = append(diffs, ops)
	}
	StackGrowthCurveEndArcShapesDifferent := false
	if len(stackofgrowthcurve.StackGrowthCurveEndArcShapes) != len(stackofgrowthcurveOther.StackGrowthCurveEndArcShapes) {
		StackGrowthCurveEndArcShapesDifferent = true
	} else {
		for i := range stackofgrowthcurve.StackGrowthCurveEndArcShapes {
			if (stackofgrowthcurve.StackGrowthCurveEndArcShapes[i] == nil) != (stackofgrowthcurveOther.StackGrowthCurveEndArcShapes[i] == nil) {
				StackGrowthCurveEndArcShapesDifferent = true
				break
			} else if stackofgrowthcurve.StackGrowthCurveEndArcShapes[i] != nil && stackofgrowthcurveOther.StackGrowthCurveEndArcShapes[i] != nil {
				// this is a pointer comparaison
				if stackofgrowthcurve.StackGrowthCurveEndArcShapes[i] != stackofgrowthcurveOther.StackGrowthCurveEndArcShapes[i] {
					StackGrowthCurveEndArcShapesDifferent = true
					break
				}
			}
		}
	}
	if StackGrowthCurveEndArcShapesDifferent {
		ops := Diff(stage, stackofgrowthcurve, stackofgrowthcurveOther, "StackGrowthCurveEndArcShapes", stackofgrowthcurveOther.StackGrowthCurveEndArcShapes, stackofgrowthcurve.StackGrowthCurveEndArcShapes)
		diffs = append(diffs, ops)
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
	if (topgrowthcurve2d.TopStartArcShapeGrid == nil) != (topgrowthcurve2dOther.TopStartArcShapeGrid == nil) {
		diffs = append(diffs, topgrowthcurve2d.GongMarshallField(stage, "TopStartArcShapeGrid"))
	} else if topgrowthcurve2d.TopStartArcShapeGrid != nil && topgrowthcurve2dOther.TopStartArcShapeGrid != nil {
		if topgrowthcurve2d.TopStartArcShapeGrid != topgrowthcurve2dOther.TopStartArcShapeGrid {
			diffs = append(diffs, topgrowthcurve2d.GongMarshallField(stage, "TopStartArcShapeGrid"))
		}
	}
	if (topgrowthcurve2d.TopEndArcShapeGrid == nil) != (topgrowthcurve2dOther.TopEndArcShapeGrid == nil) {
		diffs = append(diffs, topgrowthcurve2d.GongMarshallField(stage, "TopEndArcShapeGrid"))
	} else if topgrowthcurve2d.TopEndArcShapeGrid != nil && topgrowthcurve2dOther.TopEndArcShapeGrid != nil {
		if topgrowthcurve2d.TopEndArcShapeGrid != topgrowthcurve2dOther.TopEndArcShapeGrid {
			diffs = append(diffs, topgrowthcurve2d.GongMarshallField(stage, "TopEndArcShapeGrid"))
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
func (topstackgrowthcurveendarcshape *TopStackGrowthCurveEndArcShape) GongDiff(stage *Stage, topstackgrowthcurveendarcshapeOther *TopStackGrowthCurveEndArcShape) (diffs []string) {
	// insertion point for field diffs
	if topstackgrowthcurveendarcshape.Name != topstackgrowthcurveendarcshapeOther.Name {
		diffs = append(diffs, topstackgrowthcurveendarcshape.GongMarshallField(stage, "Name"))
	}
	if topstackgrowthcurveendarcshape.StartX != topstackgrowthcurveendarcshapeOther.StartX {
		diffs = append(diffs, topstackgrowthcurveendarcshape.GongMarshallField(stage, "StartX"))
	}
	if topstackgrowthcurveendarcshape.StartY != topstackgrowthcurveendarcshapeOther.StartY {
		diffs = append(diffs, topstackgrowthcurveendarcshape.GongMarshallField(stage, "StartY"))
	}
	if topstackgrowthcurveendarcshape.EndX != topstackgrowthcurveendarcshapeOther.EndX {
		diffs = append(diffs, topstackgrowthcurveendarcshape.GongMarshallField(stage, "EndX"))
	}
	if topstackgrowthcurveendarcshape.EndY != topstackgrowthcurveendarcshapeOther.EndY {
		diffs = append(diffs, topstackgrowthcurveendarcshape.GongMarshallField(stage, "EndY"))
	}
	if topstackgrowthcurveendarcshape.XAxisRotation != topstackgrowthcurveendarcshapeOther.XAxisRotation {
		diffs = append(diffs, topstackgrowthcurveendarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if topstackgrowthcurveendarcshape.LargeArcFlag != topstackgrowthcurveendarcshapeOther.LargeArcFlag {
		diffs = append(diffs, topstackgrowthcurveendarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topstackgrowthcurveendarcshape.SweepFlag != topstackgrowthcurveendarcshapeOther.SweepFlag {
		diffs = append(diffs, topstackgrowthcurveendarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if topstackgrowthcurveendarcshape.RadiusX != topstackgrowthcurveendarcshapeOther.RadiusX {
		diffs = append(diffs, topstackgrowthcurveendarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if topstackgrowthcurveendarcshape.RadiusY != topstackgrowthcurveendarcshapeOther.RadiusY {
		diffs = append(diffs, topstackgrowthcurveendarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstackgrowthcurvestartarcshape *TopStackGrowthCurveStartArcShape) GongDiff(stage *Stage, topstackgrowthcurvestartarcshapeOther *TopStackGrowthCurveStartArcShape) (diffs []string) {
	// insertion point for field diffs
	if topstackgrowthcurvestartarcshape.Name != topstackgrowthcurvestartarcshapeOther.Name {
		diffs = append(diffs, topstackgrowthcurvestartarcshape.GongMarshallField(stage, "Name"))
	}
	if topstackgrowthcurvestartarcshape.StartX != topstackgrowthcurvestartarcshapeOther.StartX {
		diffs = append(diffs, topstackgrowthcurvestartarcshape.GongMarshallField(stage, "StartX"))
	}
	if topstackgrowthcurvestartarcshape.StartY != topstackgrowthcurvestartarcshapeOther.StartY {
		diffs = append(diffs, topstackgrowthcurvestartarcshape.GongMarshallField(stage, "StartY"))
	}
	if topstackgrowthcurvestartarcshape.EndX != topstackgrowthcurvestartarcshapeOther.EndX {
		diffs = append(diffs, topstackgrowthcurvestartarcshape.GongMarshallField(stage, "EndX"))
	}
	if topstackgrowthcurvestartarcshape.EndY != topstackgrowthcurvestartarcshapeOther.EndY {
		diffs = append(diffs, topstackgrowthcurvestartarcshape.GongMarshallField(stage, "EndY"))
	}
	if topstackgrowthcurvestartarcshape.XAxisRotation != topstackgrowthcurvestartarcshapeOther.XAxisRotation {
		diffs = append(diffs, topstackgrowthcurvestartarcshape.GongMarshallField(stage, "XAxisRotation"))
	}
	if topstackgrowthcurvestartarcshape.LargeArcFlag != topstackgrowthcurvestartarcshapeOther.LargeArcFlag {
		diffs = append(diffs, topstackgrowthcurvestartarcshape.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topstackgrowthcurvestartarcshape.SweepFlag != topstackgrowthcurvestartarcshapeOther.SweepFlag {
		diffs = append(diffs, topstackgrowthcurvestartarcshape.GongMarshallField(stage, "SweepFlag"))
	}
	if topstackgrowthcurvestartarcshape.RadiusX != topstackgrowthcurvestartarcshapeOther.RadiusX {
		diffs = append(diffs, topstackgrowthcurvestartarcshape.GongMarshallField(stage, "RadiusX"))
	}
	if topstackgrowthcurvestartarcshape.RadiusY != topstackgrowthcurvestartarcshapeOther.RadiusY {
		diffs = append(diffs, topstackgrowthcurvestartarcshape.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstackofgrowthcurve *TopStackOfGrowthCurve) GongDiff(stage *Stage, topstackofgrowthcurveOther *TopStackOfGrowthCurve) (diffs []string) {
	// insertion point for field diffs
	if topstackofgrowthcurve.Name != topstackofgrowthcurveOther.Name {
		diffs = append(diffs, topstackofgrowthcurve.GongMarshallField(stage, "Name"))
	}
	TopStackGrowthCurveStartArcShapesDifferent := false
	if len(topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes) != len(topstackofgrowthcurveOther.TopStackGrowthCurveStartArcShapes) {
		TopStackGrowthCurveStartArcShapesDifferent = true
	} else {
		for i := range topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes {
			if (topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes[i] == nil) != (topstackofgrowthcurveOther.TopStackGrowthCurveStartArcShapes[i] == nil) {
				TopStackGrowthCurveStartArcShapesDifferent = true
				break
			} else if topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes[i] != nil && topstackofgrowthcurveOther.TopStackGrowthCurveStartArcShapes[i] != nil {
				// this is a pointer comparaison
				if topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes[i] != topstackofgrowthcurveOther.TopStackGrowthCurveStartArcShapes[i] {
					TopStackGrowthCurveStartArcShapesDifferent = true
					break
				}
			}
		}
	}
	if TopStackGrowthCurveStartArcShapesDifferent {
		ops := Diff(stage, topstackofgrowthcurve, topstackofgrowthcurveOther, "TopStackGrowthCurveStartArcShapes", topstackofgrowthcurveOther.TopStackGrowthCurveStartArcShapes, topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes)
		diffs = append(diffs, ops)
	}
	TopStackGrowthCurveEndArcShapesDifferent := false
	if len(topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes) != len(topstackofgrowthcurveOther.TopStackGrowthCurveEndArcShapes) {
		TopStackGrowthCurveEndArcShapesDifferent = true
	} else {
		for i := range topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes {
			if (topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes[i] == nil) != (topstackofgrowthcurveOther.TopStackGrowthCurveEndArcShapes[i] == nil) {
				TopStackGrowthCurveEndArcShapesDifferent = true
				break
			} else if topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes[i] != nil && topstackofgrowthcurveOther.TopStackGrowthCurveEndArcShapes[i] != nil {
				// this is a pointer comparaison
				if topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes[i] != topstackofgrowthcurveOther.TopStackGrowthCurveEndArcShapes[i] {
					TopStackGrowthCurveEndArcShapesDifferent = true
					break
				}
			}
		}
	}
	if TopStackGrowthCurveEndArcShapesDifferent {
		ops := Diff(stage, topstackofgrowthcurve, topstackofgrowthcurveOther, "TopStackGrowthCurveEndArcShapes", topstackofgrowthcurveOther.TopStackGrowthCurveEndArcShapes, topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes)
		diffs = append(diffs, ops)
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
