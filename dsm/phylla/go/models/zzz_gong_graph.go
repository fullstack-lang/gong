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

	case *BottomEndArcShapeV2:
		ok = stage.IsStagedBottomEndArcShapeV2(target)

	case *BottomEndArcShapeV2Grid:
		ok = stage.IsStagedBottomEndArcShapeV2Grid(target)

	case *BottomStackGrowthCurveEndArcShapeV2:
		ok = stage.IsStagedBottomStackGrowthCurveEndArcShapeV2(target)

	case *BottomStackGrowthCurveStartArcShapeV2:
		ok = stage.IsStagedBottomStackGrowthCurveStartArcShapeV2(target)

	case *BottomStackOfGrowthCurveV2:
		ok = stage.IsStagedBottomStackOfGrowthCurveV2(target)

	case *BottomStartArcShapeV2:
		ok = stage.IsStagedBottomStartArcShapeV2(target)

	case *BottomStartArcShapeV2Grid:
		ok = stage.IsStagedBottomStartArcShapeV2Grid(target)

	case *CircleGridShape:
		ok = stage.IsStagedCircleGridShape(target)

	case *EndArcShapeV2:
		ok = stage.IsStagedEndArcShapeV2(target)

	case *EndArcShapeV2Grid:
		ok = stage.IsStagedEndArcShapeV2Grid(target)

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

	case *StackGrowthCurveBezierShape:
		ok = stage.IsStagedStackGrowthCurveBezierShape(target)

	case *StackGrowthCurveEndArcShapeV2:
		ok = stage.IsStagedStackGrowthCurveEndArcShapeV2(target)

	case *StackGrowthCurveStartArcShapeV2:
		ok = stage.IsStagedStackGrowthCurveStartArcShapeV2(target)

	case *StackOfGrowthCurve:
		ok = stage.IsStagedStackOfGrowthCurve(target)

	case *StackOfGrowthCurveV2:
		ok = stage.IsStagedStackOfGrowthCurveV2(target)

	case *StartArcShapeV2:
		ok = stage.IsStagedStartArcShapeV2(target)

	case *StartArcShapeV2Grid:
		ok = stage.IsStagedStartArcShapeV2Grid(target)

	case *TopEndArcShapeV2:
		ok = stage.IsStagedTopEndArcShapeV2(target)

	case *TopEndArcShapeV2Grid:
		ok = stage.IsStagedTopEndArcShapeV2Grid(target)

	case *TopGrowthCurve2D:
		ok = stage.IsStagedTopGrowthCurve2D(target)

	case *TopStackGrowthCurveEndArcShapeV2:
		ok = stage.IsStagedTopStackGrowthCurveEndArcShapeV2(target)

	case *TopStackGrowthCurveStartArcShapeV2:
		ok = stage.IsStagedTopStackGrowthCurveStartArcShapeV2(target)

	case *TopStackOfGrowthCurveV2:
		ok = stage.IsStagedTopStackOfGrowthCurveV2(target)

	case *TopStartArcShapeV2:
		ok = stage.IsStagedTopStartArcShapeV2(target)

	case *TopStartArcShapeV2Grid:
		ok = stage.IsStagedTopStartArcShapeV2Grid(target)

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

	case *BottomEndArcShapeV2:
		ok = stage.IsStagedBottomEndArcShapeV2(target)

	case *BottomEndArcShapeV2Grid:
		ok = stage.IsStagedBottomEndArcShapeV2Grid(target)

	case *BottomStackGrowthCurveEndArcShapeV2:
		ok = stage.IsStagedBottomStackGrowthCurveEndArcShapeV2(target)

	case *BottomStackGrowthCurveStartArcShapeV2:
		ok = stage.IsStagedBottomStackGrowthCurveStartArcShapeV2(target)

	case *BottomStackOfGrowthCurveV2:
		ok = stage.IsStagedBottomStackOfGrowthCurveV2(target)

	case *BottomStartArcShapeV2:
		ok = stage.IsStagedBottomStartArcShapeV2(target)

	case *BottomStartArcShapeV2Grid:
		ok = stage.IsStagedBottomStartArcShapeV2Grid(target)

	case *CircleGridShape:
		ok = stage.IsStagedCircleGridShape(target)

	case *EndArcShapeV2:
		ok = stage.IsStagedEndArcShapeV2(target)

	case *EndArcShapeV2Grid:
		ok = stage.IsStagedEndArcShapeV2Grid(target)

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

	case *StackGrowthCurveBezierShape:
		ok = stage.IsStagedStackGrowthCurveBezierShape(target)

	case *StackGrowthCurveEndArcShapeV2:
		ok = stage.IsStagedStackGrowthCurveEndArcShapeV2(target)

	case *StackGrowthCurveStartArcShapeV2:
		ok = stage.IsStagedStackGrowthCurveStartArcShapeV2(target)

	case *StackOfGrowthCurve:
		ok = stage.IsStagedStackOfGrowthCurve(target)

	case *StackOfGrowthCurveV2:
		ok = stage.IsStagedStackOfGrowthCurveV2(target)

	case *StartArcShapeV2:
		ok = stage.IsStagedStartArcShapeV2(target)

	case *StartArcShapeV2Grid:
		ok = stage.IsStagedStartArcShapeV2Grid(target)

	case *TopEndArcShapeV2:
		ok = stage.IsStagedTopEndArcShapeV2(target)

	case *TopEndArcShapeV2Grid:
		ok = stage.IsStagedTopEndArcShapeV2Grid(target)

	case *TopGrowthCurve2D:
		ok = stage.IsStagedTopGrowthCurve2D(target)

	case *TopStackGrowthCurveEndArcShapeV2:
		ok = stage.IsStagedTopStackGrowthCurveEndArcShapeV2(target)

	case *TopStackGrowthCurveStartArcShapeV2:
		ok = stage.IsStagedTopStackGrowthCurveStartArcShapeV2(target)

	case *TopStackOfGrowthCurveV2:
		ok = stage.IsStagedTopStackOfGrowthCurveV2(target)

	case *TopStartArcShapeV2:
		ok = stage.IsStagedTopStartArcShapeV2(target)

	case *TopStartArcShapeV2Grid:
		ok = stage.IsStagedTopStartArcShapeV2Grid(target)

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

func (stage *Stage) IsStagedBottomEndArcShapeV2(bottomendarcshapev2 *BottomEndArcShapeV2) (ok bool) {

	_, ok = stage.BottomEndArcShapeV2s[bottomendarcshapev2]

	return
}

func (stage *Stage) IsStagedBottomEndArcShapeV2Grid(bottomendarcshapev2grid *BottomEndArcShapeV2Grid) (ok bool) {

	_, ok = stage.BottomEndArcShapeV2Grids[bottomendarcshapev2grid]

	return
}

func (stage *Stage) IsStagedBottomStackGrowthCurveEndArcShapeV2(bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) (ok bool) {

	_, ok = stage.BottomStackGrowthCurveEndArcShapeV2s[bottomstackgrowthcurveendarcshapev2]

	return
}

func (stage *Stage) IsStagedBottomStackGrowthCurveStartArcShapeV2(bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) (ok bool) {

	_, ok = stage.BottomStackGrowthCurveStartArcShapeV2s[bottomstackgrowthcurvestartarcshapev2]

	return
}

func (stage *Stage) IsStagedBottomStackOfGrowthCurveV2(bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) (ok bool) {

	_, ok = stage.BottomStackOfGrowthCurveV2s[bottomstackofgrowthcurvev2]

	return
}

func (stage *Stage) IsStagedBottomStartArcShapeV2(bottomstartarcshapev2 *BottomStartArcShapeV2) (ok bool) {

	_, ok = stage.BottomStartArcShapeV2s[bottomstartarcshapev2]

	return
}

func (stage *Stage) IsStagedBottomStartArcShapeV2Grid(bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) (ok bool) {

	_, ok = stage.BottomStartArcShapeV2Grids[bottomstartarcshapev2grid]

	return
}

func (stage *Stage) IsStagedCircleGridShape(circlegridshape *CircleGridShape) (ok bool) {

	_, ok = stage.CircleGridShapes[circlegridshape]

	return
}

func (stage *Stage) IsStagedEndArcShapeV2(endarcshapev2 *EndArcShapeV2) (ok bool) {

	_, ok = stage.EndArcShapeV2s[endarcshapev2]

	return
}

func (stage *Stage) IsStagedEndArcShapeV2Grid(endarcshapev2grid *EndArcShapeV2Grid) (ok bool) {

	_, ok = stage.EndArcShapeV2Grids[endarcshapev2grid]

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

func (stage *Stage) IsStagedStackGrowthCurveBezierShape(stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) (ok bool) {

	_, ok = stage.StackGrowthCurveBezierShapes[stackgrowthcurvebeziershape]

	return
}

func (stage *Stage) IsStagedStackGrowthCurveEndArcShapeV2(stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) (ok bool) {

	_, ok = stage.StackGrowthCurveEndArcShapeV2s[stackgrowthcurveendarcshapev2]

	return
}

func (stage *Stage) IsStagedStackGrowthCurveStartArcShapeV2(stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) (ok bool) {

	_, ok = stage.StackGrowthCurveStartArcShapeV2s[stackgrowthcurvestartarcshapev2]

	return
}

func (stage *Stage) IsStagedStackOfGrowthCurve(stackofgrowthcurve *StackOfGrowthCurve) (ok bool) {

	_, ok = stage.StackOfGrowthCurves[stackofgrowthcurve]

	return
}

func (stage *Stage) IsStagedStackOfGrowthCurveV2(stackofgrowthcurvev2 *StackOfGrowthCurveV2) (ok bool) {

	_, ok = stage.StackOfGrowthCurveV2s[stackofgrowthcurvev2]

	return
}

func (stage *Stage) IsStagedStartArcShapeV2(startarcshapev2 *StartArcShapeV2) (ok bool) {

	_, ok = stage.StartArcShapeV2s[startarcshapev2]

	return
}

func (stage *Stage) IsStagedStartArcShapeV2Grid(startarcshapev2grid *StartArcShapeV2Grid) (ok bool) {

	_, ok = stage.StartArcShapeV2Grids[startarcshapev2grid]

	return
}

func (stage *Stage) IsStagedTopEndArcShapeV2(topendarcshapev2 *TopEndArcShapeV2) (ok bool) {

	_, ok = stage.TopEndArcShapeV2s[topendarcshapev2]

	return
}

func (stage *Stage) IsStagedTopEndArcShapeV2Grid(topendarcshapev2grid *TopEndArcShapeV2Grid) (ok bool) {

	_, ok = stage.TopEndArcShapeV2Grids[topendarcshapev2grid]

	return
}

func (stage *Stage) IsStagedTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D) (ok bool) {

	_, ok = stage.TopGrowthCurve2Ds[topgrowthcurve2d]

	return
}

func (stage *Stage) IsStagedTopStackGrowthCurveEndArcShapeV2(topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) (ok bool) {

	_, ok = stage.TopStackGrowthCurveEndArcShapeV2s[topstackgrowthcurveendarcshapev2]

	return
}

func (stage *Stage) IsStagedTopStackGrowthCurveStartArcShapeV2(topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) (ok bool) {

	_, ok = stage.TopStackGrowthCurveStartArcShapeV2s[topstackgrowthcurvestartarcshapev2]

	return
}

func (stage *Stage) IsStagedTopStackOfGrowthCurveV2(topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) (ok bool) {

	_, ok = stage.TopStackOfGrowthCurveV2s[topstackofgrowthcurvev2]

	return
}

func (stage *Stage) IsStagedTopStartArcShapeV2(topstartarcshapev2 *TopStartArcShapeV2) (ok bool) {

	_, ok = stage.TopStartArcShapeV2s[topstartarcshapev2]

	return
}

func (stage *Stage) IsStagedTopStartArcShapeV2Grid(topstartarcshapev2grid *TopStartArcShapeV2Grid) (ok bool) {

	_, ok = stage.TopStartArcShapeV2Grids[topstartarcshapev2grid]

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

	case *BottomEndArcShapeV2:
		stage.StageBranchBottomEndArcShapeV2(target)

	case *BottomEndArcShapeV2Grid:
		stage.StageBranchBottomEndArcShapeV2Grid(target)

	case *BottomStackGrowthCurveEndArcShapeV2:
		stage.StageBranchBottomStackGrowthCurveEndArcShapeV2(target)

	case *BottomStackGrowthCurveStartArcShapeV2:
		stage.StageBranchBottomStackGrowthCurveStartArcShapeV2(target)

	case *BottomStackOfGrowthCurveV2:
		stage.StageBranchBottomStackOfGrowthCurveV2(target)

	case *BottomStartArcShapeV2:
		stage.StageBranchBottomStartArcShapeV2(target)

	case *BottomStartArcShapeV2Grid:
		stage.StageBranchBottomStartArcShapeV2Grid(target)

	case *CircleGridShape:
		stage.StageBranchCircleGridShape(target)

	case *EndArcShapeV2:
		stage.StageBranchEndArcShapeV2(target)

	case *EndArcShapeV2Grid:
		stage.StageBranchEndArcShapeV2Grid(target)

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

	case *StackGrowthCurveBezierShape:
		stage.StageBranchStackGrowthCurveBezierShape(target)

	case *StackGrowthCurveEndArcShapeV2:
		stage.StageBranchStackGrowthCurveEndArcShapeV2(target)

	case *StackGrowthCurveStartArcShapeV2:
		stage.StageBranchStackGrowthCurveStartArcShapeV2(target)

	case *StackOfGrowthCurve:
		stage.StageBranchStackOfGrowthCurve(target)

	case *StackOfGrowthCurveV2:
		stage.StageBranchStackOfGrowthCurveV2(target)

	case *StartArcShapeV2:
		stage.StageBranchStartArcShapeV2(target)

	case *StartArcShapeV2Grid:
		stage.StageBranchStartArcShapeV2Grid(target)

	case *TopEndArcShapeV2:
		stage.StageBranchTopEndArcShapeV2(target)

	case *TopEndArcShapeV2Grid:
		stage.StageBranchTopEndArcShapeV2Grid(target)

	case *TopGrowthCurve2D:
		stage.StageBranchTopGrowthCurve2D(target)

	case *TopStackGrowthCurveEndArcShapeV2:
		stage.StageBranchTopStackGrowthCurveEndArcShapeV2(target)

	case *TopStackGrowthCurveStartArcShapeV2:
		stage.StageBranchTopStackGrowthCurveStartArcShapeV2(target)

	case *TopStackOfGrowthCurveV2:
		stage.StageBranchTopStackOfGrowthCurveV2(target)

	case *TopStartArcShapeV2:
		stage.StageBranchTopStartArcShapeV2(target)

	case *TopStartArcShapeV2Grid:
		stage.StageBranchTopStartArcShapeV2Grid(target)

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

func (stage *Stage) StageBranchBottomEndArcShapeV2(bottomendarcshapev2 *BottomEndArcShapeV2) {

	// check if instance is already staged
	if IsStaged(stage, bottomendarcshapev2) {
		return
	}

	bottomendarcshapev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchBottomEndArcShapeV2Grid(bottomendarcshapev2grid *BottomEndArcShapeV2Grid) {

	// check if instance is already staged
	if IsStaged(stage, bottomendarcshapev2grid) {
		return
	}

	bottomendarcshapev2grid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bottomendarcshapev2 := range bottomendarcshapev2grid.BottomEndArcShapesV2 {
		StageBranch(stage, _bottomendarcshapev2)
	}

}

func (stage *Stage) StageBranchBottomStackGrowthCurveEndArcShapeV2(bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) {

	// check if instance is already staged
	if IsStaged(stage, bottomstackgrowthcurveendarcshapev2) {
		return
	}

	bottomstackgrowthcurveendarcshapev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchBottomStackGrowthCurveStartArcShapeV2(bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) {

	// check if instance is already staged
	if IsStaged(stage, bottomstackgrowthcurvestartarcshapev2) {
		return
	}

	bottomstackgrowthcurvestartarcshapev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchBottomStackOfGrowthCurveV2(bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) {

	// check if instance is already staged
	if IsStaged(stage, bottomstackofgrowthcurvev2) {
		return
	}

	bottomstackofgrowthcurvev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bottomstackgrowthcurvestartarcshapev2 := range bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s {
		StageBranch(stage, _bottomstackgrowthcurvestartarcshapev2)
	}
	for _, _bottomstackgrowthcurveendarcshapev2 := range bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s {
		StageBranch(stage, _bottomstackgrowthcurveendarcshapev2)
	}

}

func (stage *Stage) StageBranchBottomStartArcShapeV2(bottomstartarcshapev2 *BottomStartArcShapeV2) {

	// check if instance is already staged
	if IsStaged(stage, bottomstartarcshapev2) {
		return
	}

	bottomstartarcshapev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchBottomStartArcShapeV2Grid(bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) {

	// check if instance is already staged
	if IsStaged(stage, bottomstartarcshapev2grid) {
		return
	}

	bottomstartarcshapev2grid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bottomstartarcshapev2 := range bottomstartarcshapev2grid.BottomStartArcShapesV2 {
		StageBranch(stage, _bottomstartarcshapev2)
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

func (stage *Stage) StageBranchEndArcShapeV2(endarcshapev2 *EndArcShapeV2) {

	// check if instance is already staged
	if IsStaged(stage, endarcshapev2) {
		return
	}

	endarcshapev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEndArcShapeV2Grid(endarcshapev2grid *EndArcShapeV2Grid) {

	// check if instance is already staged
	if IsStaged(stage, endarcshapev2grid) {
		return
	}

	endarcshapev2grid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _endarcshapev2 := range endarcshapev2grid.EndArcShapesV2 {
		StageBranch(stage, _endarcshapev2)
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
	if growthcurve2d.StartArcShapeV2Grid != nil {
		StageBranch(stage, growthcurve2d.StartArcShapeV2Grid)
	}
	if growthcurve2d.EndArcShapeV2Grid != nil {
		StageBranch(stage, growthcurve2d.EndArcShapeV2Grid)
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
	if plant.StartArcShapeV2Grid != nil {
		StageBranch(stage, plant.StartArcShapeV2Grid)
	}
	if plant.TopStartArcShapeV2Grid != nil {
		StageBranch(stage, plant.TopStartArcShapeV2Grid)
	}
	if plant.EndArcShapeV2Grid != nil {
		StageBranch(stage, plant.EndArcShapeV2Grid)
	}
	if plant.TopEndArcShapeV2Grid != nil {
		StageBranch(stage, plant.TopEndArcShapeV2Grid)
	}
	if plant.BottomStartArcShapeV2Grid != nil {
		StageBranch(stage, plant.BottomStartArcShapeV2Grid)
	}
	if plant.BottomEndArcShapeV2Grid != nil {
		StageBranch(stage, plant.BottomEndArcShapeV2Grid)
	}
	if plant.GrowthCurveBezierShapeGrid != nil {
		StageBranch(stage, plant.GrowthCurveBezierShapeGrid)
	}
	if plant.StackOfGrowthCurve != nil {
		StageBranch(stage, plant.StackOfGrowthCurve)
	}
	if plant.StackOfGrowthCurveV2 != nil {
		StageBranch(stage, plant.StackOfGrowthCurveV2)
	}
	if plant.TopStackOfGrowthCurveV2 != nil {
		StageBranch(stage, plant.TopStackOfGrowthCurveV2)
	}
	if plant.BottomStackOfGrowthCurveV2 != nil {
		StageBranch(stage, plant.BottomStackOfGrowthCurveV2)
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

func (stage *Stage) StageBranchStackGrowthCurveBezierShape(stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) {

	// check if instance is already staged
	if IsStaged(stage, stackgrowthcurvebeziershape) {
		return
	}

	stackgrowthcurvebeziershape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackGrowthCurveEndArcShapeV2(stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) {

	// check if instance is already staged
	if IsStaged(stage, stackgrowthcurveendarcshapev2) {
		return
	}

	stackgrowthcurveendarcshapev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStackGrowthCurveStartArcShapeV2(stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) {

	// check if instance is already staged
	if IsStaged(stage, stackgrowthcurvestartarcshapev2) {
		return
	}

	stackgrowthcurvestartarcshapev2.Stage(stage)

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
	for _, _stackgrowthcurvebeziershape := range stackofgrowthcurve.StackGrowthCurveBezierShapes {
		StageBranch(stage, _stackgrowthcurvebeziershape)
	}

}

func (stage *Stage) StageBranchStackOfGrowthCurveV2(stackofgrowthcurvev2 *StackOfGrowthCurveV2) {

	// check if instance is already staged
	if IsStaged(stage, stackofgrowthcurvev2) {
		return
	}

	stackofgrowthcurvev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stackgrowthcurvestartarcshapev2 := range stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s {
		StageBranch(stage, _stackgrowthcurvestartarcshapev2)
	}
	for _, _stackgrowthcurveendarcshapev2 := range stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s {
		StageBranch(stage, _stackgrowthcurveendarcshapev2)
	}

}

func (stage *Stage) StageBranchStartArcShapeV2(startarcshapev2 *StartArcShapeV2) {

	// check if instance is already staged
	if IsStaged(stage, startarcshapev2) {
		return
	}

	startarcshapev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchStartArcShapeV2Grid(startarcshapev2grid *StartArcShapeV2Grid) {

	// check if instance is already staged
	if IsStaged(stage, startarcshapev2grid) {
		return
	}

	startarcshapev2grid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _startarcshapev2 := range startarcshapev2grid.StartArcShapesV2 {
		StageBranch(stage, _startarcshapev2)
	}

}

func (stage *Stage) StageBranchTopEndArcShapeV2(topendarcshapev2 *TopEndArcShapeV2) {

	// check if instance is already staged
	if IsStaged(stage, topendarcshapev2) {
		return
	}

	topendarcshapev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopEndArcShapeV2Grid(topendarcshapev2grid *TopEndArcShapeV2Grid) {

	// check if instance is already staged
	if IsStaged(stage, topendarcshapev2grid) {
		return
	}

	topendarcshapev2grid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topendarcshapev2 := range topendarcshapev2grid.TopEndArcShapesV2 {
		StageBranch(stage, _topendarcshapev2)
	}

}

func (stage *Stage) StageBranchTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D) {

	// check if instance is already staged
	if IsStaged(stage, topgrowthcurve2d) {
		return
	}

	topgrowthcurve2d.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if topgrowthcurve2d.TopStartArcShapeV2Grid != nil {
		StageBranch(stage, topgrowthcurve2d.TopStartArcShapeV2Grid)
	}
	if topgrowthcurve2d.TopEndArcShapeV2Grid != nil {
		StageBranch(stage, topgrowthcurve2d.TopEndArcShapeV2Grid)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStackGrowthCurveEndArcShapeV2(topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) {

	// check if instance is already staged
	if IsStaged(stage, topstackgrowthcurveendarcshapev2) {
		return
	}

	topstackgrowthcurveendarcshapev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStackGrowthCurveStartArcShapeV2(topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) {

	// check if instance is already staged
	if IsStaged(stage, topstackgrowthcurvestartarcshapev2) {
		return
	}

	topstackgrowthcurvestartarcshapev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStackOfGrowthCurveV2(topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) {

	// check if instance is already staged
	if IsStaged(stage, topstackofgrowthcurvev2) {
		return
	}

	topstackofgrowthcurvev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topstackgrowthcurvestartarcshapev2 := range topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s {
		StageBranch(stage, _topstackgrowthcurvestartarcshapev2)
	}
	for _, _topstackgrowthcurveendarcshapev2 := range topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s {
		StageBranch(stage, _topstackgrowthcurveendarcshapev2)
	}

}

func (stage *Stage) StageBranchTopStartArcShapeV2(topstartarcshapev2 *TopStartArcShapeV2) {

	// check if instance is already staged
	if IsStaged(stage, topstartarcshapev2) {
		return
	}

	topstartarcshapev2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTopStartArcShapeV2Grid(topstartarcshapev2grid *TopStartArcShapeV2Grid) {

	// check if instance is already staged
	if IsStaged(stage, topstartarcshapev2grid) {
		return
	}

	topstartarcshapev2grid.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topstartarcshapev2 := range topstartarcshapev2grid.TopStartArcShapesV2 {
		StageBranch(stage, _topstartarcshapev2)
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

	case *BottomEndArcShapeV2:
		toT := CopyBranchBottomEndArcShapeV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BottomEndArcShapeV2Grid:
		toT := CopyBranchBottomEndArcShapeV2Grid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BottomStackGrowthCurveEndArcShapeV2:
		toT := CopyBranchBottomStackGrowthCurveEndArcShapeV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BottomStackGrowthCurveStartArcShapeV2:
		toT := CopyBranchBottomStackGrowthCurveStartArcShapeV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BottomStackOfGrowthCurveV2:
		toT := CopyBranchBottomStackOfGrowthCurveV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BottomStartArcShapeV2:
		toT := CopyBranchBottomStartArcShapeV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BottomStartArcShapeV2Grid:
		toT := CopyBranchBottomStartArcShapeV2Grid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CircleGridShape:
		toT := CopyBranchCircleGridShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EndArcShapeV2:
		toT := CopyBranchEndArcShapeV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EndArcShapeV2Grid:
		toT := CopyBranchEndArcShapeV2Grid(mapOrigCopy, fromT)
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

	case *StackGrowthCurveBezierShape:
		toT := CopyBranchStackGrowthCurveBezierShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackGrowthCurveEndArcShapeV2:
		toT := CopyBranchStackGrowthCurveEndArcShapeV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackGrowthCurveStartArcShapeV2:
		toT := CopyBranchStackGrowthCurveStartArcShapeV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfGrowthCurve:
		toT := CopyBranchStackOfGrowthCurve(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StackOfGrowthCurveV2:
		toT := CopyBranchStackOfGrowthCurveV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartArcShapeV2:
		toT := CopyBranchStartArcShapeV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartArcShapeV2Grid:
		toT := CopyBranchStartArcShapeV2Grid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopEndArcShapeV2:
		toT := CopyBranchTopEndArcShapeV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopEndArcShapeV2Grid:
		toT := CopyBranchTopEndArcShapeV2Grid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopGrowthCurve2D:
		toT := CopyBranchTopGrowthCurve2D(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackGrowthCurveEndArcShapeV2:
		toT := CopyBranchTopStackGrowthCurveEndArcShapeV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackGrowthCurveStartArcShapeV2:
		toT := CopyBranchTopStackGrowthCurveStartArcShapeV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStackOfGrowthCurveV2:
		toT := CopyBranchTopStackOfGrowthCurveV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStartArcShapeV2:
		toT := CopyBranchTopStartArcShapeV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TopStartArcShapeV2Grid:
		toT := CopyBranchTopStartArcShapeV2Grid(mapOrigCopy, fromT)
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

func CopyBranchBottomEndArcShapeV2(mapOrigCopy map[any]any, bottomendarcshapev2From *BottomEndArcShapeV2) (bottomendarcshapev2To *BottomEndArcShapeV2) {

	// bottomendarcshapev2From has already been copied
	if _bottomendarcshapev2To, ok := mapOrigCopy[bottomendarcshapev2From]; ok {
		bottomendarcshapev2To = _bottomendarcshapev2To.(*BottomEndArcShapeV2)
		return
	}

	bottomendarcshapev2To = new(BottomEndArcShapeV2)
	mapOrigCopy[bottomendarcshapev2From] = bottomendarcshapev2To
	bottomendarcshapev2From.CopyBasicFields(bottomendarcshapev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBottomEndArcShapeV2Grid(mapOrigCopy map[any]any, bottomendarcshapev2gridFrom *BottomEndArcShapeV2Grid) (bottomendarcshapev2gridTo *BottomEndArcShapeV2Grid) {

	// bottomendarcshapev2gridFrom has already been copied
	if _bottomendarcshapev2gridTo, ok := mapOrigCopy[bottomendarcshapev2gridFrom]; ok {
		bottomendarcshapev2gridTo = _bottomendarcshapev2gridTo.(*BottomEndArcShapeV2Grid)
		return
	}

	bottomendarcshapev2gridTo = new(BottomEndArcShapeV2Grid)
	mapOrigCopy[bottomendarcshapev2gridFrom] = bottomendarcshapev2gridTo
	bottomendarcshapev2gridFrom.CopyBasicFields(bottomendarcshapev2gridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bottomendarcshapev2 := range bottomendarcshapev2gridFrom.BottomEndArcShapesV2 {
		bottomendarcshapev2gridTo.BottomEndArcShapesV2 = append(bottomendarcshapev2gridTo.BottomEndArcShapesV2, CopyBranchBottomEndArcShapeV2(mapOrigCopy, _bottomendarcshapev2))
	}

	return
}

func CopyBranchBottomStackGrowthCurveEndArcShapeV2(mapOrigCopy map[any]any, bottomstackgrowthcurveendarcshapev2From *BottomStackGrowthCurveEndArcShapeV2) (bottomstackgrowthcurveendarcshapev2To *BottomStackGrowthCurveEndArcShapeV2) {

	// bottomstackgrowthcurveendarcshapev2From has already been copied
	if _bottomstackgrowthcurveendarcshapev2To, ok := mapOrigCopy[bottomstackgrowthcurveendarcshapev2From]; ok {
		bottomstackgrowthcurveendarcshapev2To = _bottomstackgrowthcurveendarcshapev2To.(*BottomStackGrowthCurveEndArcShapeV2)
		return
	}

	bottomstackgrowthcurveendarcshapev2To = new(BottomStackGrowthCurveEndArcShapeV2)
	mapOrigCopy[bottomstackgrowthcurveendarcshapev2From] = bottomstackgrowthcurveendarcshapev2To
	bottomstackgrowthcurveendarcshapev2From.CopyBasicFields(bottomstackgrowthcurveendarcshapev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBottomStackGrowthCurveStartArcShapeV2(mapOrigCopy map[any]any, bottomstackgrowthcurvestartarcshapev2From *BottomStackGrowthCurveStartArcShapeV2) (bottomstackgrowthcurvestartarcshapev2To *BottomStackGrowthCurveStartArcShapeV2) {

	// bottomstackgrowthcurvestartarcshapev2From has already been copied
	if _bottomstackgrowthcurvestartarcshapev2To, ok := mapOrigCopy[bottomstackgrowthcurvestartarcshapev2From]; ok {
		bottomstackgrowthcurvestartarcshapev2To = _bottomstackgrowthcurvestartarcshapev2To.(*BottomStackGrowthCurveStartArcShapeV2)
		return
	}

	bottomstackgrowthcurvestartarcshapev2To = new(BottomStackGrowthCurveStartArcShapeV2)
	mapOrigCopy[bottomstackgrowthcurvestartarcshapev2From] = bottomstackgrowthcurvestartarcshapev2To
	bottomstackgrowthcurvestartarcshapev2From.CopyBasicFields(bottomstackgrowthcurvestartarcshapev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBottomStackOfGrowthCurveV2(mapOrigCopy map[any]any, bottomstackofgrowthcurvev2From *BottomStackOfGrowthCurveV2) (bottomstackofgrowthcurvev2To *BottomStackOfGrowthCurveV2) {

	// bottomstackofgrowthcurvev2From has already been copied
	if _bottomstackofgrowthcurvev2To, ok := mapOrigCopy[bottomstackofgrowthcurvev2From]; ok {
		bottomstackofgrowthcurvev2To = _bottomstackofgrowthcurvev2To.(*BottomStackOfGrowthCurveV2)
		return
	}

	bottomstackofgrowthcurvev2To = new(BottomStackOfGrowthCurveV2)
	mapOrigCopy[bottomstackofgrowthcurvev2From] = bottomstackofgrowthcurvev2To
	bottomstackofgrowthcurvev2From.CopyBasicFields(bottomstackofgrowthcurvev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bottomstackgrowthcurvestartarcshapev2 := range bottomstackofgrowthcurvev2From.BottomStackGrowthCurveStartArcShapeV2s {
		bottomstackofgrowthcurvev2To.BottomStackGrowthCurveStartArcShapeV2s = append(bottomstackofgrowthcurvev2To.BottomStackGrowthCurveStartArcShapeV2s, CopyBranchBottomStackGrowthCurveStartArcShapeV2(mapOrigCopy, _bottomstackgrowthcurvestartarcshapev2))
	}
	for _, _bottomstackgrowthcurveendarcshapev2 := range bottomstackofgrowthcurvev2From.BottomStackGrowthCurveEndArcShapeV2s {
		bottomstackofgrowthcurvev2To.BottomStackGrowthCurveEndArcShapeV2s = append(bottomstackofgrowthcurvev2To.BottomStackGrowthCurveEndArcShapeV2s, CopyBranchBottomStackGrowthCurveEndArcShapeV2(mapOrigCopy, _bottomstackgrowthcurveendarcshapev2))
	}

	return
}

func CopyBranchBottomStartArcShapeV2(mapOrigCopy map[any]any, bottomstartarcshapev2From *BottomStartArcShapeV2) (bottomstartarcshapev2To *BottomStartArcShapeV2) {

	// bottomstartarcshapev2From has already been copied
	if _bottomstartarcshapev2To, ok := mapOrigCopy[bottomstartarcshapev2From]; ok {
		bottomstartarcshapev2To = _bottomstartarcshapev2To.(*BottomStartArcShapeV2)
		return
	}

	bottomstartarcshapev2To = new(BottomStartArcShapeV2)
	mapOrigCopy[bottomstartarcshapev2From] = bottomstartarcshapev2To
	bottomstartarcshapev2From.CopyBasicFields(bottomstartarcshapev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBottomStartArcShapeV2Grid(mapOrigCopy map[any]any, bottomstartarcshapev2gridFrom *BottomStartArcShapeV2Grid) (bottomstartarcshapev2gridTo *BottomStartArcShapeV2Grid) {

	// bottomstartarcshapev2gridFrom has already been copied
	if _bottomstartarcshapev2gridTo, ok := mapOrigCopy[bottomstartarcshapev2gridFrom]; ok {
		bottomstartarcshapev2gridTo = _bottomstartarcshapev2gridTo.(*BottomStartArcShapeV2Grid)
		return
	}

	bottomstartarcshapev2gridTo = new(BottomStartArcShapeV2Grid)
	mapOrigCopy[bottomstartarcshapev2gridFrom] = bottomstartarcshapev2gridTo
	bottomstartarcshapev2gridFrom.CopyBasicFields(bottomstartarcshapev2gridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bottomstartarcshapev2 := range bottomstartarcshapev2gridFrom.BottomStartArcShapesV2 {
		bottomstartarcshapev2gridTo.BottomStartArcShapesV2 = append(bottomstartarcshapev2gridTo.BottomStartArcShapesV2, CopyBranchBottomStartArcShapeV2(mapOrigCopy, _bottomstartarcshapev2))
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

func CopyBranchEndArcShapeV2(mapOrigCopy map[any]any, endarcshapev2From *EndArcShapeV2) (endarcshapev2To *EndArcShapeV2) {

	// endarcshapev2From has already been copied
	if _endarcshapev2To, ok := mapOrigCopy[endarcshapev2From]; ok {
		endarcshapev2To = _endarcshapev2To.(*EndArcShapeV2)
		return
	}

	endarcshapev2To = new(EndArcShapeV2)
	mapOrigCopy[endarcshapev2From] = endarcshapev2To
	endarcshapev2From.CopyBasicFields(endarcshapev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEndArcShapeV2Grid(mapOrigCopy map[any]any, endarcshapev2gridFrom *EndArcShapeV2Grid) (endarcshapev2gridTo *EndArcShapeV2Grid) {

	// endarcshapev2gridFrom has already been copied
	if _endarcshapev2gridTo, ok := mapOrigCopy[endarcshapev2gridFrom]; ok {
		endarcshapev2gridTo = _endarcshapev2gridTo.(*EndArcShapeV2Grid)
		return
	}

	endarcshapev2gridTo = new(EndArcShapeV2Grid)
	mapOrigCopy[endarcshapev2gridFrom] = endarcshapev2gridTo
	endarcshapev2gridFrom.CopyBasicFields(endarcshapev2gridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _endarcshapev2 := range endarcshapev2gridFrom.EndArcShapesV2 {
		endarcshapev2gridTo.EndArcShapesV2 = append(endarcshapev2gridTo.EndArcShapesV2, CopyBranchEndArcShapeV2(mapOrigCopy, _endarcshapev2))
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
	if growthcurve2dFrom.StartArcShapeV2Grid != nil {
		growthcurve2dTo.StartArcShapeV2Grid = CopyBranchStartArcShapeV2Grid(mapOrigCopy, growthcurve2dFrom.StartArcShapeV2Grid)
	}
	if growthcurve2dFrom.EndArcShapeV2Grid != nil {
		growthcurve2dTo.EndArcShapeV2Grid = CopyBranchEndArcShapeV2Grid(mapOrigCopy, growthcurve2dFrom.EndArcShapeV2Grid)
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
	if plantFrom.StartArcShapeV2Grid != nil {
		plantTo.StartArcShapeV2Grid = CopyBranchStartArcShapeV2Grid(mapOrigCopy, plantFrom.StartArcShapeV2Grid)
	}
	if plantFrom.TopStartArcShapeV2Grid != nil {
		plantTo.TopStartArcShapeV2Grid = CopyBranchTopStartArcShapeV2Grid(mapOrigCopy, plantFrom.TopStartArcShapeV2Grid)
	}
	if plantFrom.EndArcShapeV2Grid != nil {
		plantTo.EndArcShapeV2Grid = CopyBranchEndArcShapeV2Grid(mapOrigCopy, plantFrom.EndArcShapeV2Grid)
	}
	if plantFrom.TopEndArcShapeV2Grid != nil {
		plantTo.TopEndArcShapeV2Grid = CopyBranchTopEndArcShapeV2Grid(mapOrigCopy, plantFrom.TopEndArcShapeV2Grid)
	}
	if plantFrom.BottomStartArcShapeV2Grid != nil {
		plantTo.BottomStartArcShapeV2Grid = CopyBranchBottomStartArcShapeV2Grid(mapOrigCopy, plantFrom.BottomStartArcShapeV2Grid)
	}
	if plantFrom.BottomEndArcShapeV2Grid != nil {
		plantTo.BottomEndArcShapeV2Grid = CopyBranchBottomEndArcShapeV2Grid(mapOrigCopy, plantFrom.BottomEndArcShapeV2Grid)
	}
	if plantFrom.GrowthCurveBezierShapeGrid != nil {
		plantTo.GrowthCurveBezierShapeGrid = CopyBranchGrowthCurveBezierShapeGrid(mapOrigCopy, plantFrom.GrowthCurveBezierShapeGrid)
	}
	if plantFrom.StackOfGrowthCurve != nil {
		plantTo.StackOfGrowthCurve = CopyBranchStackOfGrowthCurve(mapOrigCopy, plantFrom.StackOfGrowthCurve)
	}
	if plantFrom.StackOfGrowthCurveV2 != nil {
		plantTo.StackOfGrowthCurveV2 = CopyBranchStackOfGrowthCurveV2(mapOrigCopy, plantFrom.StackOfGrowthCurveV2)
	}
	if plantFrom.TopStackOfGrowthCurveV2 != nil {
		plantTo.TopStackOfGrowthCurveV2 = CopyBranchTopStackOfGrowthCurveV2(mapOrigCopy, plantFrom.TopStackOfGrowthCurveV2)
	}
	if plantFrom.BottomStackOfGrowthCurveV2 != nil {
		plantTo.BottomStackOfGrowthCurveV2 = CopyBranchBottomStackOfGrowthCurveV2(mapOrigCopy, plantFrom.BottomStackOfGrowthCurveV2)
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

func CopyBranchStackGrowthCurveBezierShape(mapOrigCopy map[any]any, stackgrowthcurvebeziershapeFrom *StackGrowthCurveBezierShape) (stackgrowthcurvebeziershapeTo *StackGrowthCurveBezierShape) {

	// stackgrowthcurvebeziershapeFrom has already been copied
	if _stackgrowthcurvebeziershapeTo, ok := mapOrigCopy[stackgrowthcurvebeziershapeFrom]; ok {
		stackgrowthcurvebeziershapeTo = _stackgrowthcurvebeziershapeTo.(*StackGrowthCurveBezierShape)
		return
	}

	stackgrowthcurvebeziershapeTo = new(StackGrowthCurveBezierShape)
	mapOrigCopy[stackgrowthcurvebeziershapeFrom] = stackgrowthcurvebeziershapeTo
	stackgrowthcurvebeziershapeFrom.CopyBasicFields(stackgrowthcurvebeziershapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackGrowthCurveEndArcShapeV2(mapOrigCopy map[any]any, stackgrowthcurveendarcshapev2From *StackGrowthCurveEndArcShapeV2) (stackgrowthcurveendarcshapev2To *StackGrowthCurveEndArcShapeV2) {

	// stackgrowthcurveendarcshapev2From has already been copied
	if _stackgrowthcurveendarcshapev2To, ok := mapOrigCopy[stackgrowthcurveendarcshapev2From]; ok {
		stackgrowthcurveendarcshapev2To = _stackgrowthcurveendarcshapev2To.(*StackGrowthCurveEndArcShapeV2)
		return
	}

	stackgrowthcurveendarcshapev2To = new(StackGrowthCurveEndArcShapeV2)
	mapOrigCopy[stackgrowthcurveendarcshapev2From] = stackgrowthcurveendarcshapev2To
	stackgrowthcurveendarcshapev2From.CopyBasicFields(stackgrowthcurveendarcshapev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStackGrowthCurveStartArcShapeV2(mapOrigCopy map[any]any, stackgrowthcurvestartarcshapev2From *StackGrowthCurveStartArcShapeV2) (stackgrowthcurvestartarcshapev2To *StackGrowthCurveStartArcShapeV2) {

	// stackgrowthcurvestartarcshapev2From has already been copied
	if _stackgrowthcurvestartarcshapev2To, ok := mapOrigCopy[stackgrowthcurvestartarcshapev2From]; ok {
		stackgrowthcurvestartarcshapev2To = _stackgrowthcurvestartarcshapev2To.(*StackGrowthCurveStartArcShapeV2)
		return
	}

	stackgrowthcurvestartarcshapev2To = new(StackGrowthCurveStartArcShapeV2)
	mapOrigCopy[stackgrowthcurvestartarcshapev2From] = stackgrowthcurvestartarcshapev2To
	stackgrowthcurvestartarcshapev2From.CopyBasicFields(stackgrowthcurvestartarcshapev2To)

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
	for _, _stackgrowthcurvebeziershape := range stackofgrowthcurveFrom.StackGrowthCurveBezierShapes {
		stackofgrowthcurveTo.StackGrowthCurveBezierShapes = append(stackofgrowthcurveTo.StackGrowthCurveBezierShapes, CopyBranchStackGrowthCurveBezierShape(mapOrigCopy, _stackgrowthcurvebeziershape))
	}

	return
}

func CopyBranchStackOfGrowthCurveV2(mapOrigCopy map[any]any, stackofgrowthcurvev2From *StackOfGrowthCurveV2) (stackofgrowthcurvev2To *StackOfGrowthCurveV2) {

	// stackofgrowthcurvev2From has already been copied
	if _stackofgrowthcurvev2To, ok := mapOrigCopy[stackofgrowthcurvev2From]; ok {
		stackofgrowthcurvev2To = _stackofgrowthcurvev2To.(*StackOfGrowthCurveV2)
		return
	}

	stackofgrowthcurvev2To = new(StackOfGrowthCurveV2)
	mapOrigCopy[stackofgrowthcurvev2From] = stackofgrowthcurvev2To
	stackofgrowthcurvev2From.CopyBasicFields(stackofgrowthcurvev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stackgrowthcurvestartarcshapev2 := range stackofgrowthcurvev2From.StackGrowthCurveStartArcShapeV2s {
		stackofgrowthcurvev2To.StackGrowthCurveStartArcShapeV2s = append(stackofgrowthcurvev2To.StackGrowthCurveStartArcShapeV2s, CopyBranchStackGrowthCurveStartArcShapeV2(mapOrigCopy, _stackgrowthcurvestartarcshapev2))
	}
	for _, _stackgrowthcurveendarcshapev2 := range stackofgrowthcurvev2From.StackGrowthCurveEndArcShapeV2s {
		stackofgrowthcurvev2To.StackGrowthCurveEndArcShapeV2s = append(stackofgrowthcurvev2To.StackGrowthCurveEndArcShapeV2s, CopyBranchStackGrowthCurveEndArcShapeV2(mapOrigCopy, _stackgrowthcurveendarcshapev2))
	}

	return
}

func CopyBranchStartArcShapeV2(mapOrigCopy map[any]any, startarcshapev2From *StartArcShapeV2) (startarcshapev2To *StartArcShapeV2) {

	// startarcshapev2From has already been copied
	if _startarcshapev2To, ok := mapOrigCopy[startarcshapev2From]; ok {
		startarcshapev2To = _startarcshapev2To.(*StartArcShapeV2)
		return
	}

	startarcshapev2To = new(StartArcShapeV2)
	mapOrigCopy[startarcshapev2From] = startarcshapev2To
	startarcshapev2From.CopyBasicFields(startarcshapev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStartArcShapeV2Grid(mapOrigCopy map[any]any, startarcshapev2gridFrom *StartArcShapeV2Grid) (startarcshapev2gridTo *StartArcShapeV2Grid) {

	// startarcshapev2gridFrom has already been copied
	if _startarcshapev2gridTo, ok := mapOrigCopy[startarcshapev2gridFrom]; ok {
		startarcshapev2gridTo = _startarcshapev2gridTo.(*StartArcShapeV2Grid)
		return
	}

	startarcshapev2gridTo = new(StartArcShapeV2Grid)
	mapOrigCopy[startarcshapev2gridFrom] = startarcshapev2gridTo
	startarcshapev2gridFrom.CopyBasicFields(startarcshapev2gridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _startarcshapev2 := range startarcshapev2gridFrom.StartArcShapesV2 {
		startarcshapev2gridTo.StartArcShapesV2 = append(startarcshapev2gridTo.StartArcShapesV2, CopyBranchStartArcShapeV2(mapOrigCopy, _startarcshapev2))
	}

	return
}

func CopyBranchTopEndArcShapeV2(mapOrigCopy map[any]any, topendarcshapev2From *TopEndArcShapeV2) (topendarcshapev2To *TopEndArcShapeV2) {

	// topendarcshapev2From has already been copied
	if _topendarcshapev2To, ok := mapOrigCopy[topendarcshapev2From]; ok {
		topendarcshapev2To = _topendarcshapev2To.(*TopEndArcShapeV2)
		return
	}

	topendarcshapev2To = new(TopEndArcShapeV2)
	mapOrigCopy[topendarcshapev2From] = topendarcshapev2To
	topendarcshapev2From.CopyBasicFields(topendarcshapev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopEndArcShapeV2Grid(mapOrigCopy map[any]any, topendarcshapev2gridFrom *TopEndArcShapeV2Grid) (topendarcshapev2gridTo *TopEndArcShapeV2Grid) {

	// topendarcshapev2gridFrom has already been copied
	if _topendarcshapev2gridTo, ok := mapOrigCopy[topendarcshapev2gridFrom]; ok {
		topendarcshapev2gridTo = _topendarcshapev2gridTo.(*TopEndArcShapeV2Grid)
		return
	}

	topendarcshapev2gridTo = new(TopEndArcShapeV2Grid)
	mapOrigCopy[topendarcshapev2gridFrom] = topendarcshapev2gridTo
	topendarcshapev2gridFrom.CopyBasicFields(topendarcshapev2gridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topendarcshapev2 := range topendarcshapev2gridFrom.TopEndArcShapesV2 {
		topendarcshapev2gridTo.TopEndArcShapesV2 = append(topendarcshapev2gridTo.TopEndArcShapesV2, CopyBranchTopEndArcShapeV2(mapOrigCopy, _topendarcshapev2))
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
	if topgrowthcurve2dFrom.TopStartArcShapeV2Grid != nil {
		topgrowthcurve2dTo.TopStartArcShapeV2Grid = CopyBranchTopStartArcShapeV2Grid(mapOrigCopy, topgrowthcurve2dFrom.TopStartArcShapeV2Grid)
	}
	if topgrowthcurve2dFrom.TopEndArcShapeV2Grid != nil {
		topgrowthcurve2dTo.TopEndArcShapeV2Grid = CopyBranchTopEndArcShapeV2Grid(mapOrigCopy, topgrowthcurve2dFrom.TopEndArcShapeV2Grid)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStackGrowthCurveEndArcShapeV2(mapOrigCopy map[any]any, topstackgrowthcurveendarcshapev2From *TopStackGrowthCurveEndArcShapeV2) (topstackgrowthcurveendarcshapev2To *TopStackGrowthCurveEndArcShapeV2) {

	// topstackgrowthcurveendarcshapev2From has already been copied
	if _topstackgrowthcurveendarcshapev2To, ok := mapOrigCopy[topstackgrowthcurveendarcshapev2From]; ok {
		topstackgrowthcurveendarcshapev2To = _topstackgrowthcurveendarcshapev2To.(*TopStackGrowthCurveEndArcShapeV2)
		return
	}

	topstackgrowthcurveendarcshapev2To = new(TopStackGrowthCurveEndArcShapeV2)
	mapOrigCopy[topstackgrowthcurveendarcshapev2From] = topstackgrowthcurveendarcshapev2To
	topstackgrowthcurveendarcshapev2From.CopyBasicFields(topstackgrowthcurveendarcshapev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStackGrowthCurveStartArcShapeV2(mapOrigCopy map[any]any, topstackgrowthcurvestartarcshapev2From *TopStackGrowthCurveStartArcShapeV2) (topstackgrowthcurvestartarcshapev2To *TopStackGrowthCurveStartArcShapeV2) {

	// topstackgrowthcurvestartarcshapev2From has already been copied
	if _topstackgrowthcurvestartarcshapev2To, ok := mapOrigCopy[topstackgrowthcurvestartarcshapev2From]; ok {
		topstackgrowthcurvestartarcshapev2To = _topstackgrowthcurvestartarcshapev2To.(*TopStackGrowthCurveStartArcShapeV2)
		return
	}

	topstackgrowthcurvestartarcshapev2To = new(TopStackGrowthCurveStartArcShapeV2)
	mapOrigCopy[topstackgrowthcurvestartarcshapev2From] = topstackgrowthcurvestartarcshapev2To
	topstackgrowthcurvestartarcshapev2From.CopyBasicFields(topstackgrowthcurvestartarcshapev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStackOfGrowthCurveV2(mapOrigCopy map[any]any, topstackofgrowthcurvev2From *TopStackOfGrowthCurveV2) (topstackofgrowthcurvev2To *TopStackOfGrowthCurveV2) {

	// topstackofgrowthcurvev2From has already been copied
	if _topstackofgrowthcurvev2To, ok := mapOrigCopy[topstackofgrowthcurvev2From]; ok {
		topstackofgrowthcurvev2To = _topstackofgrowthcurvev2To.(*TopStackOfGrowthCurveV2)
		return
	}

	topstackofgrowthcurvev2To = new(TopStackOfGrowthCurveV2)
	mapOrigCopy[topstackofgrowthcurvev2From] = topstackofgrowthcurvev2To
	topstackofgrowthcurvev2From.CopyBasicFields(topstackofgrowthcurvev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topstackgrowthcurvestartarcshapev2 := range topstackofgrowthcurvev2From.TopStackGrowthCurveStartArcShapeV2s {
		topstackofgrowthcurvev2To.TopStackGrowthCurveStartArcShapeV2s = append(topstackofgrowthcurvev2To.TopStackGrowthCurveStartArcShapeV2s, CopyBranchTopStackGrowthCurveStartArcShapeV2(mapOrigCopy, _topstackgrowthcurvestartarcshapev2))
	}
	for _, _topstackgrowthcurveendarcshapev2 := range topstackofgrowthcurvev2From.TopStackGrowthCurveEndArcShapeV2s {
		topstackofgrowthcurvev2To.TopStackGrowthCurveEndArcShapeV2s = append(topstackofgrowthcurvev2To.TopStackGrowthCurveEndArcShapeV2s, CopyBranchTopStackGrowthCurveEndArcShapeV2(mapOrigCopy, _topstackgrowthcurveendarcshapev2))
	}

	return
}

func CopyBranchTopStartArcShapeV2(mapOrigCopy map[any]any, topstartarcshapev2From *TopStartArcShapeV2) (topstartarcshapev2To *TopStartArcShapeV2) {

	// topstartarcshapev2From has already been copied
	if _topstartarcshapev2To, ok := mapOrigCopy[topstartarcshapev2From]; ok {
		topstartarcshapev2To = _topstartarcshapev2To.(*TopStartArcShapeV2)
		return
	}

	topstartarcshapev2To = new(TopStartArcShapeV2)
	mapOrigCopy[topstartarcshapev2From] = topstartarcshapev2To
	topstartarcshapev2From.CopyBasicFields(topstartarcshapev2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTopStartArcShapeV2Grid(mapOrigCopy map[any]any, topstartarcshapev2gridFrom *TopStartArcShapeV2Grid) (topstartarcshapev2gridTo *TopStartArcShapeV2Grid) {

	// topstartarcshapev2gridFrom has already been copied
	if _topstartarcshapev2gridTo, ok := mapOrigCopy[topstartarcshapev2gridFrom]; ok {
		topstartarcshapev2gridTo = _topstartarcshapev2gridTo.(*TopStartArcShapeV2Grid)
		return
	}

	topstartarcshapev2gridTo = new(TopStartArcShapeV2Grid)
	mapOrigCopy[topstartarcshapev2gridFrom] = topstartarcshapev2gridTo
	topstartarcshapev2gridFrom.CopyBasicFields(topstartarcshapev2gridTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topstartarcshapev2 := range topstartarcshapev2gridFrom.TopStartArcShapesV2 {
		topstartarcshapev2gridTo.TopStartArcShapesV2 = append(topstartarcshapev2gridTo.TopStartArcShapesV2, CopyBranchTopStartArcShapeV2(mapOrigCopy, _topstartarcshapev2))
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

	case *BottomEndArcShapeV2:
		stage.UnstageBranchBottomEndArcShapeV2(target)

	case *BottomEndArcShapeV2Grid:
		stage.UnstageBranchBottomEndArcShapeV2Grid(target)

	case *BottomStackGrowthCurveEndArcShapeV2:
		stage.UnstageBranchBottomStackGrowthCurveEndArcShapeV2(target)

	case *BottomStackGrowthCurveStartArcShapeV2:
		stage.UnstageBranchBottomStackGrowthCurveStartArcShapeV2(target)

	case *BottomStackOfGrowthCurveV2:
		stage.UnstageBranchBottomStackOfGrowthCurveV2(target)

	case *BottomStartArcShapeV2:
		stage.UnstageBranchBottomStartArcShapeV2(target)

	case *BottomStartArcShapeV2Grid:
		stage.UnstageBranchBottomStartArcShapeV2Grid(target)

	case *CircleGridShape:
		stage.UnstageBranchCircleGridShape(target)

	case *EndArcShapeV2:
		stage.UnstageBranchEndArcShapeV2(target)

	case *EndArcShapeV2Grid:
		stage.UnstageBranchEndArcShapeV2Grid(target)

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

	case *StackGrowthCurveBezierShape:
		stage.UnstageBranchStackGrowthCurveBezierShape(target)

	case *StackGrowthCurveEndArcShapeV2:
		stage.UnstageBranchStackGrowthCurveEndArcShapeV2(target)

	case *StackGrowthCurveStartArcShapeV2:
		stage.UnstageBranchStackGrowthCurveStartArcShapeV2(target)

	case *StackOfGrowthCurve:
		stage.UnstageBranchStackOfGrowthCurve(target)

	case *StackOfGrowthCurveV2:
		stage.UnstageBranchStackOfGrowthCurveV2(target)

	case *StartArcShapeV2:
		stage.UnstageBranchStartArcShapeV2(target)

	case *StartArcShapeV2Grid:
		stage.UnstageBranchStartArcShapeV2Grid(target)

	case *TopEndArcShapeV2:
		stage.UnstageBranchTopEndArcShapeV2(target)

	case *TopEndArcShapeV2Grid:
		stage.UnstageBranchTopEndArcShapeV2Grid(target)

	case *TopGrowthCurve2D:
		stage.UnstageBranchTopGrowthCurve2D(target)

	case *TopStackGrowthCurveEndArcShapeV2:
		stage.UnstageBranchTopStackGrowthCurveEndArcShapeV2(target)

	case *TopStackGrowthCurveStartArcShapeV2:
		stage.UnstageBranchTopStackGrowthCurveStartArcShapeV2(target)

	case *TopStackOfGrowthCurveV2:
		stage.UnstageBranchTopStackOfGrowthCurveV2(target)

	case *TopStartArcShapeV2:
		stage.UnstageBranchTopStartArcShapeV2(target)

	case *TopStartArcShapeV2Grid:
		stage.UnstageBranchTopStartArcShapeV2Grid(target)

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

func (stage *Stage) UnstageBranchBottomEndArcShapeV2(bottomendarcshapev2 *BottomEndArcShapeV2) {

	// check if instance is already staged
	if !IsStaged(stage, bottomendarcshapev2) {
		return
	}

	bottomendarcshapev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchBottomEndArcShapeV2Grid(bottomendarcshapev2grid *BottomEndArcShapeV2Grid) {

	// check if instance is already staged
	if !IsStaged(stage, bottomendarcshapev2grid) {
		return
	}

	bottomendarcshapev2grid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bottomendarcshapev2 := range bottomendarcshapev2grid.BottomEndArcShapesV2 {
		UnstageBranch(stage, _bottomendarcshapev2)
	}

}

func (stage *Stage) UnstageBranchBottomStackGrowthCurveEndArcShapeV2(bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) {

	// check if instance is already staged
	if !IsStaged(stage, bottomstackgrowthcurveendarcshapev2) {
		return
	}

	bottomstackgrowthcurveendarcshapev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchBottomStackGrowthCurveStartArcShapeV2(bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) {

	// check if instance is already staged
	if !IsStaged(stage, bottomstackgrowthcurvestartarcshapev2) {
		return
	}

	bottomstackgrowthcurvestartarcshapev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchBottomStackOfGrowthCurveV2(bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) {

	// check if instance is already staged
	if !IsStaged(stage, bottomstackofgrowthcurvev2) {
		return
	}

	bottomstackofgrowthcurvev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bottomstackgrowthcurvestartarcshapev2 := range bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s {
		UnstageBranch(stage, _bottomstackgrowthcurvestartarcshapev2)
	}
	for _, _bottomstackgrowthcurveendarcshapev2 := range bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s {
		UnstageBranch(stage, _bottomstackgrowthcurveendarcshapev2)
	}

}

func (stage *Stage) UnstageBranchBottomStartArcShapeV2(bottomstartarcshapev2 *BottomStartArcShapeV2) {

	// check if instance is already staged
	if !IsStaged(stage, bottomstartarcshapev2) {
		return
	}

	bottomstartarcshapev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchBottomStartArcShapeV2Grid(bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) {

	// check if instance is already staged
	if !IsStaged(stage, bottomstartarcshapev2grid) {
		return
	}

	bottomstartarcshapev2grid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bottomstartarcshapev2 := range bottomstartarcshapev2grid.BottomStartArcShapesV2 {
		UnstageBranch(stage, _bottomstartarcshapev2)
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

func (stage *Stage) UnstageBranchEndArcShapeV2(endarcshapev2 *EndArcShapeV2) {

	// check if instance is already staged
	if !IsStaged(stage, endarcshapev2) {
		return
	}

	endarcshapev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEndArcShapeV2Grid(endarcshapev2grid *EndArcShapeV2Grid) {

	// check if instance is already staged
	if !IsStaged(stage, endarcshapev2grid) {
		return
	}

	endarcshapev2grid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _endarcshapev2 := range endarcshapev2grid.EndArcShapesV2 {
		UnstageBranch(stage, _endarcshapev2)
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
	if growthcurve2d.StartArcShapeV2Grid != nil {
		UnstageBranch(stage, growthcurve2d.StartArcShapeV2Grid)
	}
	if growthcurve2d.EndArcShapeV2Grid != nil {
		UnstageBranch(stage, growthcurve2d.EndArcShapeV2Grid)
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
	if plant.StartArcShapeV2Grid != nil {
		UnstageBranch(stage, plant.StartArcShapeV2Grid)
	}
	if plant.TopStartArcShapeV2Grid != nil {
		UnstageBranch(stage, plant.TopStartArcShapeV2Grid)
	}
	if plant.EndArcShapeV2Grid != nil {
		UnstageBranch(stage, plant.EndArcShapeV2Grid)
	}
	if plant.TopEndArcShapeV2Grid != nil {
		UnstageBranch(stage, plant.TopEndArcShapeV2Grid)
	}
	if plant.BottomStartArcShapeV2Grid != nil {
		UnstageBranch(stage, plant.BottomStartArcShapeV2Grid)
	}
	if plant.BottomEndArcShapeV2Grid != nil {
		UnstageBranch(stage, plant.BottomEndArcShapeV2Grid)
	}
	if plant.GrowthCurveBezierShapeGrid != nil {
		UnstageBranch(stage, plant.GrowthCurveBezierShapeGrid)
	}
	if plant.StackOfGrowthCurve != nil {
		UnstageBranch(stage, plant.StackOfGrowthCurve)
	}
	if plant.StackOfGrowthCurveV2 != nil {
		UnstageBranch(stage, plant.StackOfGrowthCurveV2)
	}
	if plant.TopStackOfGrowthCurveV2 != nil {
		UnstageBranch(stage, plant.TopStackOfGrowthCurveV2)
	}
	if plant.BottomStackOfGrowthCurveV2 != nil {
		UnstageBranch(stage, plant.BottomStackOfGrowthCurveV2)
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

func (stage *Stage) UnstageBranchStackGrowthCurveBezierShape(stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) {

	// check if instance is already staged
	if !IsStaged(stage, stackgrowthcurvebeziershape) {
		return
	}

	stackgrowthcurvebeziershape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackGrowthCurveEndArcShapeV2(stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) {

	// check if instance is already staged
	if !IsStaged(stage, stackgrowthcurveendarcshapev2) {
		return
	}

	stackgrowthcurveendarcshapev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStackGrowthCurveStartArcShapeV2(stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) {

	// check if instance is already staged
	if !IsStaged(stage, stackgrowthcurvestartarcshapev2) {
		return
	}

	stackgrowthcurvestartarcshapev2.Unstage(stage)

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
	for _, _stackgrowthcurvebeziershape := range stackofgrowthcurve.StackGrowthCurveBezierShapes {
		UnstageBranch(stage, _stackgrowthcurvebeziershape)
	}

}

func (stage *Stage) UnstageBranchStackOfGrowthCurveV2(stackofgrowthcurvev2 *StackOfGrowthCurveV2) {

	// check if instance is already staged
	if !IsStaged(stage, stackofgrowthcurvev2) {
		return
	}

	stackofgrowthcurvev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stackgrowthcurvestartarcshapev2 := range stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s {
		UnstageBranch(stage, _stackgrowthcurvestartarcshapev2)
	}
	for _, _stackgrowthcurveendarcshapev2 := range stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s {
		UnstageBranch(stage, _stackgrowthcurveendarcshapev2)
	}

}

func (stage *Stage) UnstageBranchStartArcShapeV2(startarcshapev2 *StartArcShapeV2) {

	// check if instance is already staged
	if !IsStaged(stage, startarcshapev2) {
		return
	}

	startarcshapev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchStartArcShapeV2Grid(startarcshapev2grid *StartArcShapeV2Grid) {

	// check if instance is already staged
	if !IsStaged(stage, startarcshapev2grid) {
		return
	}

	startarcshapev2grid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _startarcshapev2 := range startarcshapev2grid.StartArcShapesV2 {
		UnstageBranch(stage, _startarcshapev2)
	}

}

func (stage *Stage) UnstageBranchTopEndArcShapeV2(topendarcshapev2 *TopEndArcShapeV2) {

	// check if instance is already staged
	if !IsStaged(stage, topendarcshapev2) {
		return
	}

	topendarcshapev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopEndArcShapeV2Grid(topendarcshapev2grid *TopEndArcShapeV2Grid) {

	// check if instance is already staged
	if !IsStaged(stage, topendarcshapev2grid) {
		return
	}

	topendarcshapev2grid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topendarcshapev2 := range topendarcshapev2grid.TopEndArcShapesV2 {
		UnstageBranch(stage, _topendarcshapev2)
	}

}

func (stage *Stage) UnstageBranchTopGrowthCurve2D(topgrowthcurve2d *TopGrowthCurve2D) {

	// check if instance is already staged
	if !IsStaged(stage, topgrowthcurve2d) {
		return
	}

	topgrowthcurve2d.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if topgrowthcurve2d.TopStartArcShapeV2Grid != nil {
		UnstageBranch(stage, topgrowthcurve2d.TopStartArcShapeV2Grid)
	}
	if topgrowthcurve2d.TopEndArcShapeV2Grid != nil {
		UnstageBranch(stage, topgrowthcurve2d.TopEndArcShapeV2Grid)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStackGrowthCurveEndArcShapeV2(topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) {

	// check if instance is already staged
	if !IsStaged(stage, topstackgrowthcurveendarcshapev2) {
		return
	}

	topstackgrowthcurveendarcshapev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStackGrowthCurveStartArcShapeV2(topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) {

	// check if instance is already staged
	if !IsStaged(stage, topstackgrowthcurvestartarcshapev2) {
		return
	}

	topstackgrowthcurvestartarcshapev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStackOfGrowthCurveV2(topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) {

	// check if instance is already staged
	if !IsStaged(stage, topstackofgrowthcurvev2) {
		return
	}

	topstackofgrowthcurvev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topstackgrowthcurvestartarcshapev2 := range topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s {
		UnstageBranch(stage, _topstackgrowthcurvestartarcshapev2)
	}
	for _, _topstackgrowthcurveendarcshapev2 := range topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s {
		UnstageBranch(stage, _topstackgrowthcurveendarcshapev2)
	}

}

func (stage *Stage) UnstageBranchTopStartArcShapeV2(topstartarcshapev2 *TopStartArcShapeV2) {

	// check if instance is already staged
	if !IsStaged(stage, topstartarcshapev2) {
		return
	}

	topstartarcshapev2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTopStartArcShapeV2Grid(topstartarcshapev2grid *TopStartArcShapeV2Grid) {

	// check if instance is already staged
	if !IsStaged(stage, topstartarcshapev2grid) {
		return
	}

	topstartarcshapev2grid.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _topstartarcshapev2 := range topstartarcshapev2grid.TopStartArcShapesV2 {
		UnstageBranch(stage, _topstartarcshapev2)
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

func (reference *BottomEndArcShapeV2) GongReconstructPointersFromReferences(stage *Stage, instance *BottomEndArcShapeV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *BottomEndArcShapeV2Grid) GongReconstructPointersFromReferences(stage *Stage, instance *BottomEndArcShapeV2Grid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.BottomEndArcShapesV2 = reference.BottomEndArcShapesV2[:0]
	for _, _b := range instance.BottomEndArcShapesV2 {
		reference.BottomEndArcShapesV2 = append(reference.BottomEndArcShapesV2, stage.BottomEndArcShapeV2s_reference[_b])
	}
}

func (reference *BottomStackGrowthCurveEndArcShapeV2) GongReconstructPointersFromReferences(stage *Stage, instance *BottomStackGrowthCurveEndArcShapeV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *BottomStackGrowthCurveStartArcShapeV2) GongReconstructPointersFromReferences(stage *Stage, instance *BottomStackGrowthCurveStartArcShapeV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *BottomStackOfGrowthCurveV2) GongReconstructPointersFromReferences(stage *Stage, instance *BottomStackOfGrowthCurveV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.BottomStackGrowthCurveStartArcShapeV2s = reference.BottomStackGrowthCurveStartArcShapeV2s[:0]
	for _, _b := range instance.BottomStackGrowthCurveStartArcShapeV2s {
		reference.BottomStackGrowthCurveStartArcShapeV2s = append(reference.BottomStackGrowthCurveStartArcShapeV2s, stage.BottomStackGrowthCurveStartArcShapeV2s_reference[_b])
	}
	reference.BottomStackGrowthCurveEndArcShapeV2s = reference.BottomStackGrowthCurveEndArcShapeV2s[:0]
	for _, _b := range instance.BottomStackGrowthCurveEndArcShapeV2s {
		reference.BottomStackGrowthCurveEndArcShapeV2s = append(reference.BottomStackGrowthCurveEndArcShapeV2s, stage.BottomStackGrowthCurveEndArcShapeV2s_reference[_b])
	}
}

func (reference *BottomStartArcShapeV2) GongReconstructPointersFromReferences(stage *Stage, instance *BottomStartArcShapeV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *BottomStartArcShapeV2Grid) GongReconstructPointersFromReferences(stage *Stage, instance *BottomStartArcShapeV2Grid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.BottomStartArcShapesV2 = reference.BottomStartArcShapesV2[:0]
	for _, _b := range instance.BottomStartArcShapesV2 {
		reference.BottomStartArcShapesV2 = append(reference.BottomStartArcShapesV2, stage.BottomStartArcShapeV2s_reference[_b])
	}
}

func (reference *CircleGridShape) GongReconstructPointersFromReferences(stage *Stage, instance *CircleGridShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EndArcShapeV2) GongReconstructPointersFromReferences(stage *Stage, instance *EndArcShapeV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *EndArcShapeV2Grid) GongReconstructPointersFromReferences(stage *Stage, instance *EndArcShapeV2Grid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.EndArcShapesV2 = reference.EndArcShapesV2[:0]
	for _, _b := range instance.EndArcShapesV2 {
		reference.EndArcShapesV2 = append(reference.EndArcShapesV2, stage.EndArcShapeV2s_reference[_b])
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
	if instance.StartArcShapeV2Grid != nil {
		reference.StartArcShapeV2Grid = stage.StartArcShapeV2Grids_reference[instance.StartArcShapeV2Grid]
	}
	if instance.EndArcShapeV2Grid != nil {
		reference.EndArcShapeV2Grid = stage.EndArcShapeV2Grids_reference[instance.EndArcShapeV2Grid]
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
	if instance.StartArcShapeV2Grid != nil {
		reference.StartArcShapeV2Grid = stage.StartArcShapeV2Grids_reference[instance.StartArcShapeV2Grid]
	}
	if instance.TopStartArcShapeV2Grid != nil {
		reference.TopStartArcShapeV2Grid = stage.TopStartArcShapeV2Grids_reference[instance.TopStartArcShapeV2Grid]
	}
	if instance.EndArcShapeV2Grid != nil {
		reference.EndArcShapeV2Grid = stage.EndArcShapeV2Grids_reference[instance.EndArcShapeV2Grid]
	}
	if instance.TopEndArcShapeV2Grid != nil {
		reference.TopEndArcShapeV2Grid = stage.TopEndArcShapeV2Grids_reference[instance.TopEndArcShapeV2Grid]
	}
	if instance.BottomStartArcShapeV2Grid != nil {
		reference.BottomStartArcShapeV2Grid = stage.BottomStartArcShapeV2Grids_reference[instance.BottomStartArcShapeV2Grid]
	}
	if instance.BottomEndArcShapeV2Grid != nil {
		reference.BottomEndArcShapeV2Grid = stage.BottomEndArcShapeV2Grids_reference[instance.BottomEndArcShapeV2Grid]
	}
	if instance.GrowthCurveBezierShapeGrid != nil {
		reference.GrowthCurveBezierShapeGrid = stage.GrowthCurveBezierShapeGrids_reference[instance.GrowthCurveBezierShapeGrid]
	}
	if instance.StackOfGrowthCurve != nil {
		reference.StackOfGrowthCurve = stage.StackOfGrowthCurves_reference[instance.StackOfGrowthCurve]
	}
	if instance.StackOfGrowthCurveV2 != nil {
		reference.StackOfGrowthCurveV2 = stage.StackOfGrowthCurveV2s_reference[instance.StackOfGrowthCurveV2]
	}
	if instance.TopStackOfGrowthCurveV2 != nil {
		reference.TopStackOfGrowthCurveV2 = stage.TopStackOfGrowthCurveV2s_reference[instance.TopStackOfGrowthCurveV2]
	}
	if instance.BottomStackOfGrowthCurveV2 != nil {
		reference.BottomStackOfGrowthCurveV2 = stage.BottomStackOfGrowthCurveV2s_reference[instance.BottomStackOfGrowthCurveV2]
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

func (reference *StackGrowthCurveBezierShape) GongReconstructPointersFromReferences(stage *Stage, instance *StackGrowthCurveBezierShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackGrowthCurveEndArcShapeV2) GongReconstructPointersFromReferences(stage *Stage, instance *StackGrowthCurveEndArcShapeV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackGrowthCurveStartArcShapeV2) GongReconstructPointersFromReferences(stage *Stage, instance *StackGrowthCurveStartArcShapeV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StackOfGrowthCurve) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfGrowthCurve) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.StackGrowthCurveBezierShapes = reference.StackGrowthCurveBezierShapes[:0]
	for _, _b := range instance.StackGrowthCurveBezierShapes {
		reference.StackGrowthCurveBezierShapes = append(reference.StackGrowthCurveBezierShapes, stage.StackGrowthCurveBezierShapes_reference[_b])
	}
}

func (reference *StackOfGrowthCurveV2) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfGrowthCurveV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.StackGrowthCurveStartArcShapeV2s = reference.StackGrowthCurveStartArcShapeV2s[:0]
	for _, _b := range instance.StackGrowthCurveStartArcShapeV2s {
		reference.StackGrowthCurveStartArcShapeV2s = append(reference.StackGrowthCurveStartArcShapeV2s, stage.StackGrowthCurveStartArcShapeV2s_reference[_b])
	}
	reference.StackGrowthCurveEndArcShapeV2s = reference.StackGrowthCurveEndArcShapeV2s[:0]
	for _, _b := range instance.StackGrowthCurveEndArcShapeV2s {
		reference.StackGrowthCurveEndArcShapeV2s = append(reference.StackGrowthCurveEndArcShapeV2s, stage.StackGrowthCurveEndArcShapeV2s_reference[_b])
	}
}

func (reference *StartArcShapeV2) GongReconstructPointersFromReferences(stage *Stage, instance *StartArcShapeV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StartArcShapeV2Grid) GongReconstructPointersFromReferences(stage *Stage, instance *StartArcShapeV2Grid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.StartArcShapesV2 = reference.StartArcShapesV2[:0]
	for _, _b := range instance.StartArcShapesV2 {
		reference.StartArcShapesV2 = append(reference.StartArcShapesV2, stage.StartArcShapeV2s_reference[_b])
	}
}

func (reference *TopEndArcShapeV2) GongReconstructPointersFromReferences(stage *Stage, instance *TopEndArcShapeV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopEndArcShapeV2Grid) GongReconstructPointersFromReferences(stage *Stage, instance *TopEndArcShapeV2Grid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.TopEndArcShapesV2 = reference.TopEndArcShapesV2[:0]
	for _, _b := range instance.TopEndArcShapesV2 {
		reference.TopEndArcShapesV2 = append(reference.TopEndArcShapesV2, stage.TopEndArcShapeV2s_reference[_b])
	}
}

func (reference *TopGrowthCurve2D) GongReconstructPointersFromReferences(stage *Stage, instance *TopGrowthCurve2D) {
	// insertion point for pointers field
	if instance.TopStartArcShapeV2Grid != nil {
		reference.TopStartArcShapeV2Grid = stage.TopStartArcShapeV2Grids_reference[instance.TopStartArcShapeV2Grid]
	}
	if instance.TopEndArcShapeV2Grid != nil {
		reference.TopEndArcShapeV2Grid = stage.TopEndArcShapeV2Grids_reference[instance.TopEndArcShapeV2Grid]
	}
	// insertion point for slice of pointers field
}

func (reference *TopStackGrowthCurveEndArcShapeV2) GongReconstructPointersFromReferences(stage *Stage, instance *TopStackGrowthCurveEndArcShapeV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStackGrowthCurveStartArcShapeV2) GongReconstructPointersFromReferences(stage *Stage, instance *TopStackGrowthCurveStartArcShapeV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStackOfGrowthCurveV2) GongReconstructPointersFromReferences(stage *Stage, instance *TopStackOfGrowthCurveV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.TopStackGrowthCurveStartArcShapeV2s = reference.TopStackGrowthCurveStartArcShapeV2s[:0]
	for _, _b := range instance.TopStackGrowthCurveStartArcShapeV2s {
		reference.TopStackGrowthCurveStartArcShapeV2s = append(reference.TopStackGrowthCurveStartArcShapeV2s, stage.TopStackGrowthCurveStartArcShapeV2s_reference[_b])
	}
	reference.TopStackGrowthCurveEndArcShapeV2s = reference.TopStackGrowthCurveEndArcShapeV2s[:0]
	for _, _b := range instance.TopStackGrowthCurveEndArcShapeV2s {
		reference.TopStackGrowthCurveEndArcShapeV2s = append(reference.TopStackGrowthCurveEndArcShapeV2s, stage.TopStackGrowthCurveEndArcShapeV2s_reference[_b])
	}
}

func (reference *TopStartArcShapeV2) GongReconstructPointersFromReferences(stage *Stage, instance *TopStartArcShapeV2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TopStartArcShapeV2Grid) GongReconstructPointersFromReferences(stage *Stage, instance *TopStartArcShapeV2Grid) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.TopStartArcShapesV2 = reference.TopStartArcShapesV2[:0]
	for _, _b := range instance.TopStartArcShapesV2 {
		reference.TopStartArcShapesV2 = append(reference.TopStartArcShapesV2, stage.TopStartArcShapeV2s_reference[_b])
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

func (reference *BottomEndArcShapeV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *BottomEndArcShapeV2Grid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _BottomEndArcShapesV2 []*BottomEndArcShapeV2
	for _, _reference := range reference.BottomEndArcShapesV2 {
		if _instance, ok := stage.BottomEndArcShapeV2s_instance[_reference]; ok {
			_BottomEndArcShapesV2 = append(_BottomEndArcShapesV2, _instance)
		}
	}
	reference.BottomEndArcShapesV2 = _BottomEndArcShapesV2
}

func (reference *BottomStackGrowthCurveEndArcShapeV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *BottomStackGrowthCurveStartArcShapeV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *BottomStackOfGrowthCurveV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _BottomStackGrowthCurveStartArcShapeV2s []*BottomStackGrowthCurveStartArcShapeV2
	for _, _reference := range reference.BottomStackGrowthCurveStartArcShapeV2s {
		if _instance, ok := stage.BottomStackGrowthCurveStartArcShapeV2s_instance[_reference]; ok {
			_BottomStackGrowthCurveStartArcShapeV2s = append(_BottomStackGrowthCurveStartArcShapeV2s, _instance)
		}
	}
	reference.BottomStackGrowthCurveStartArcShapeV2s = _BottomStackGrowthCurveStartArcShapeV2s
	var _BottomStackGrowthCurveEndArcShapeV2s []*BottomStackGrowthCurveEndArcShapeV2
	for _, _reference := range reference.BottomStackGrowthCurveEndArcShapeV2s {
		if _instance, ok := stage.BottomStackGrowthCurveEndArcShapeV2s_instance[_reference]; ok {
			_BottomStackGrowthCurveEndArcShapeV2s = append(_BottomStackGrowthCurveEndArcShapeV2s, _instance)
		}
	}
	reference.BottomStackGrowthCurveEndArcShapeV2s = _BottomStackGrowthCurveEndArcShapeV2s
}

func (reference *BottomStartArcShapeV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *BottomStartArcShapeV2Grid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _BottomStartArcShapesV2 []*BottomStartArcShapeV2
	for _, _reference := range reference.BottomStartArcShapesV2 {
		if _instance, ok := stage.BottomStartArcShapeV2s_instance[_reference]; ok {
			_BottomStartArcShapesV2 = append(_BottomStartArcShapesV2, _instance)
		}
	}
	reference.BottomStartArcShapesV2 = _BottomStartArcShapesV2
}

func (reference *CircleGridShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EndArcShapeV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *EndArcShapeV2Grid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _EndArcShapesV2 []*EndArcShapeV2
	for _, _reference := range reference.EndArcShapesV2 {
		if _instance, ok := stage.EndArcShapeV2s_instance[_reference]; ok {
			_EndArcShapesV2 = append(_EndArcShapesV2, _instance)
		}
	}
	reference.EndArcShapesV2 = _EndArcShapesV2
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
	if _reference := reference.StartArcShapeV2Grid; _reference != nil {
		reference.StartArcShapeV2Grid = nil
		if _instance, ok := stage.StartArcShapeV2Grids_instance[_reference]; ok {
			reference.StartArcShapeV2Grid = _instance
		}
	}
	if _reference := reference.EndArcShapeV2Grid; _reference != nil {
		reference.EndArcShapeV2Grid = nil
		if _instance, ok := stage.EndArcShapeV2Grids_instance[_reference]; ok {
			reference.EndArcShapeV2Grid = _instance
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
	if _reference := reference.StartArcShapeV2Grid; _reference != nil {
		reference.StartArcShapeV2Grid = nil
		if _instance, ok := stage.StartArcShapeV2Grids_instance[_reference]; ok {
			reference.StartArcShapeV2Grid = _instance
		}
	}
	if _reference := reference.TopStartArcShapeV2Grid; _reference != nil {
		reference.TopStartArcShapeV2Grid = nil
		if _instance, ok := stage.TopStartArcShapeV2Grids_instance[_reference]; ok {
			reference.TopStartArcShapeV2Grid = _instance
		}
	}
	if _reference := reference.EndArcShapeV2Grid; _reference != nil {
		reference.EndArcShapeV2Grid = nil
		if _instance, ok := stage.EndArcShapeV2Grids_instance[_reference]; ok {
			reference.EndArcShapeV2Grid = _instance
		}
	}
	if _reference := reference.TopEndArcShapeV2Grid; _reference != nil {
		reference.TopEndArcShapeV2Grid = nil
		if _instance, ok := stage.TopEndArcShapeV2Grids_instance[_reference]; ok {
			reference.TopEndArcShapeV2Grid = _instance
		}
	}
	if _reference := reference.BottomStartArcShapeV2Grid; _reference != nil {
		reference.BottomStartArcShapeV2Grid = nil
		if _instance, ok := stage.BottomStartArcShapeV2Grids_instance[_reference]; ok {
			reference.BottomStartArcShapeV2Grid = _instance
		}
	}
	if _reference := reference.BottomEndArcShapeV2Grid; _reference != nil {
		reference.BottomEndArcShapeV2Grid = nil
		if _instance, ok := stage.BottomEndArcShapeV2Grids_instance[_reference]; ok {
			reference.BottomEndArcShapeV2Grid = _instance
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
	if _reference := reference.StackOfGrowthCurveV2; _reference != nil {
		reference.StackOfGrowthCurveV2 = nil
		if _instance, ok := stage.StackOfGrowthCurveV2s_instance[_reference]; ok {
			reference.StackOfGrowthCurveV2 = _instance
		}
	}
	if _reference := reference.TopStackOfGrowthCurveV2; _reference != nil {
		reference.TopStackOfGrowthCurveV2 = nil
		if _instance, ok := stage.TopStackOfGrowthCurveV2s_instance[_reference]; ok {
			reference.TopStackOfGrowthCurveV2 = _instance
		}
	}
	if _reference := reference.BottomStackOfGrowthCurveV2; _reference != nil {
		reference.BottomStackOfGrowthCurveV2 = nil
		if _instance, ok := stage.BottomStackOfGrowthCurveV2s_instance[_reference]; ok {
			reference.BottomStackOfGrowthCurveV2 = _instance
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

func (reference *StackGrowthCurveBezierShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackGrowthCurveEndArcShapeV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackGrowthCurveStartArcShapeV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StackOfGrowthCurve) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _StackGrowthCurveBezierShapes []*StackGrowthCurveBezierShape
	for _, _reference := range reference.StackGrowthCurveBezierShapes {
		if _instance, ok := stage.StackGrowthCurveBezierShapes_instance[_reference]; ok {
			_StackGrowthCurveBezierShapes = append(_StackGrowthCurveBezierShapes, _instance)
		}
	}
	reference.StackGrowthCurveBezierShapes = _StackGrowthCurveBezierShapes
}

func (reference *StackOfGrowthCurveV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _StackGrowthCurveStartArcShapeV2s []*StackGrowthCurveStartArcShapeV2
	for _, _reference := range reference.StackGrowthCurveStartArcShapeV2s {
		if _instance, ok := stage.StackGrowthCurveStartArcShapeV2s_instance[_reference]; ok {
			_StackGrowthCurveStartArcShapeV2s = append(_StackGrowthCurveStartArcShapeV2s, _instance)
		}
	}
	reference.StackGrowthCurveStartArcShapeV2s = _StackGrowthCurveStartArcShapeV2s
	var _StackGrowthCurveEndArcShapeV2s []*StackGrowthCurveEndArcShapeV2
	for _, _reference := range reference.StackGrowthCurveEndArcShapeV2s {
		if _instance, ok := stage.StackGrowthCurveEndArcShapeV2s_instance[_reference]; ok {
			_StackGrowthCurveEndArcShapeV2s = append(_StackGrowthCurveEndArcShapeV2s, _instance)
		}
	}
	reference.StackGrowthCurveEndArcShapeV2s = _StackGrowthCurveEndArcShapeV2s
}

func (reference *StartArcShapeV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StartArcShapeV2Grid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _StartArcShapesV2 []*StartArcShapeV2
	for _, _reference := range reference.StartArcShapesV2 {
		if _instance, ok := stage.StartArcShapeV2s_instance[_reference]; ok {
			_StartArcShapesV2 = append(_StartArcShapesV2, _instance)
		}
	}
	reference.StartArcShapesV2 = _StartArcShapesV2
}

func (reference *TopEndArcShapeV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopEndArcShapeV2Grid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _TopEndArcShapesV2 []*TopEndArcShapeV2
	for _, _reference := range reference.TopEndArcShapesV2 {
		if _instance, ok := stage.TopEndArcShapeV2s_instance[_reference]; ok {
			_TopEndArcShapesV2 = append(_TopEndArcShapesV2, _instance)
		}
	}
	reference.TopEndArcShapesV2 = _TopEndArcShapesV2
}

func (reference *TopGrowthCurve2D) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.TopStartArcShapeV2Grid; _reference != nil {
		reference.TopStartArcShapeV2Grid = nil
		if _instance, ok := stage.TopStartArcShapeV2Grids_instance[_reference]; ok {
			reference.TopStartArcShapeV2Grid = _instance
		}
	}
	if _reference := reference.TopEndArcShapeV2Grid; _reference != nil {
		reference.TopEndArcShapeV2Grid = nil
		if _instance, ok := stage.TopEndArcShapeV2Grids_instance[_reference]; ok {
			reference.TopEndArcShapeV2Grid = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *TopStackGrowthCurveEndArcShapeV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStackGrowthCurveStartArcShapeV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStackOfGrowthCurveV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _TopStackGrowthCurveStartArcShapeV2s []*TopStackGrowthCurveStartArcShapeV2
	for _, _reference := range reference.TopStackGrowthCurveStartArcShapeV2s {
		if _instance, ok := stage.TopStackGrowthCurveStartArcShapeV2s_instance[_reference]; ok {
			_TopStackGrowthCurveStartArcShapeV2s = append(_TopStackGrowthCurveStartArcShapeV2s, _instance)
		}
	}
	reference.TopStackGrowthCurveStartArcShapeV2s = _TopStackGrowthCurveStartArcShapeV2s
	var _TopStackGrowthCurveEndArcShapeV2s []*TopStackGrowthCurveEndArcShapeV2
	for _, _reference := range reference.TopStackGrowthCurveEndArcShapeV2s {
		if _instance, ok := stage.TopStackGrowthCurveEndArcShapeV2s_instance[_reference]; ok {
			_TopStackGrowthCurveEndArcShapeV2s = append(_TopStackGrowthCurveEndArcShapeV2s, _instance)
		}
	}
	reference.TopStackGrowthCurveEndArcShapeV2s = _TopStackGrowthCurveEndArcShapeV2s
}

func (reference *TopStartArcShapeV2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TopStartArcShapeV2Grid) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _TopStartArcShapesV2 []*TopStartArcShapeV2
	for _, _reference := range reference.TopStartArcShapesV2 {
		if _instance, ok := stage.TopStartArcShapeV2s_instance[_reference]; ok {
			_TopStartArcShapesV2 = append(_TopStartArcShapesV2, _instance)
		}
	}
	reference.TopStartArcShapesV2 = _TopStartArcShapesV2
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
func (bottomendarcshapev2 *BottomEndArcShapeV2) GongDiff(stage *Stage, bottomendarcshapev2Other *BottomEndArcShapeV2) (diffs []string) {
	// insertion point for field diffs
	if bottomendarcshapev2.Name != bottomendarcshapev2Other.Name {
		diffs = append(diffs, bottomendarcshapev2.GongMarshallField(stage, "Name"))
	}
	if bottomendarcshapev2.StartX != bottomendarcshapev2Other.StartX {
		diffs = append(diffs, bottomendarcshapev2.GongMarshallField(stage, "StartX"))
	}
	if bottomendarcshapev2.StartY != bottomendarcshapev2Other.StartY {
		diffs = append(diffs, bottomendarcshapev2.GongMarshallField(stage, "StartY"))
	}
	if bottomendarcshapev2.EndX != bottomendarcshapev2Other.EndX {
		diffs = append(diffs, bottomendarcshapev2.GongMarshallField(stage, "EndX"))
	}
	if bottomendarcshapev2.EndY != bottomendarcshapev2Other.EndY {
		diffs = append(diffs, bottomendarcshapev2.GongMarshallField(stage, "EndY"))
	}
	if bottomendarcshapev2.XAxisRotation != bottomendarcshapev2Other.XAxisRotation {
		diffs = append(diffs, bottomendarcshapev2.GongMarshallField(stage, "XAxisRotation"))
	}
	if bottomendarcshapev2.LargeArcFlag != bottomendarcshapev2Other.LargeArcFlag {
		diffs = append(diffs, bottomendarcshapev2.GongMarshallField(stage, "LargeArcFlag"))
	}
	if bottomendarcshapev2.SweepFlag != bottomendarcshapev2Other.SweepFlag {
		diffs = append(diffs, bottomendarcshapev2.GongMarshallField(stage, "SweepFlag"))
	}
	if bottomendarcshapev2.RadiusX != bottomendarcshapev2Other.RadiusX {
		diffs = append(diffs, bottomendarcshapev2.GongMarshallField(stage, "RadiusX"))
	}
	if bottomendarcshapev2.RadiusY != bottomendarcshapev2Other.RadiusY {
		diffs = append(diffs, bottomendarcshapev2.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) GongDiff(stage *Stage, bottomendarcshapev2gridOther *BottomEndArcShapeV2Grid) (diffs []string) {
	// insertion point for field diffs
	if bottomendarcshapev2grid.Name != bottomendarcshapev2gridOther.Name {
		diffs = append(diffs, bottomendarcshapev2grid.GongMarshallField(stage, "Name"))
	}
	BottomEndArcShapesV2Different := false
	if len(bottomendarcshapev2grid.BottomEndArcShapesV2) != len(bottomendarcshapev2gridOther.BottomEndArcShapesV2) {
		BottomEndArcShapesV2Different = true
	} else {
		for i := range bottomendarcshapev2grid.BottomEndArcShapesV2 {
			if (bottomendarcshapev2grid.BottomEndArcShapesV2[i] == nil) != (bottomendarcshapev2gridOther.BottomEndArcShapesV2[i] == nil) {
				BottomEndArcShapesV2Different = true
				break
			} else if bottomendarcshapev2grid.BottomEndArcShapesV2[i] != nil && bottomendarcshapev2gridOther.BottomEndArcShapesV2[i] != nil {
				// this is a pointer comparaison
				if bottomendarcshapev2grid.BottomEndArcShapesV2[i] != bottomendarcshapev2gridOther.BottomEndArcShapesV2[i] {
					BottomEndArcShapesV2Different = true
					break
				}
			}
		}
	}
	if BottomEndArcShapesV2Different {
		ops := Diff(stage, bottomendarcshapev2grid, bottomendarcshapev2gridOther, "BottomEndArcShapesV2", bottomendarcshapev2gridOther.BottomEndArcShapesV2, bottomendarcshapev2grid.BottomEndArcShapesV2)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) GongDiff(stage *Stage, bottomstackgrowthcurveendarcshapev2Other *BottomStackGrowthCurveEndArcShapeV2) (diffs []string) {
	// insertion point for field diffs
	if bottomstackgrowthcurveendarcshapev2.Name != bottomstackgrowthcurveendarcshapev2Other.Name {
		diffs = append(diffs, bottomstackgrowthcurveendarcshapev2.GongMarshallField(stage, "Name"))
	}
	if bottomstackgrowthcurveendarcshapev2.StartX != bottomstackgrowthcurveendarcshapev2Other.StartX {
		diffs = append(diffs, bottomstackgrowthcurveendarcshapev2.GongMarshallField(stage, "StartX"))
	}
	if bottomstackgrowthcurveendarcshapev2.StartY != bottomstackgrowthcurveendarcshapev2Other.StartY {
		diffs = append(diffs, bottomstackgrowthcurveendarcshapev2.GongMarshallField(stage, "StartY"))
	}
	if bottomstackgrowthcurveendarcshapev2.EndX != bottomstackgrowthcurveendarcshapev2Other.EndX {
		diffs = append(diffs, bottomstackgrowthcurveendarcshapev2.GongMarshallField(stage, "EndX"))
	}
	if bottomstackgrowthcurveendarcshapev2.EndY != bottomstackgrowthcurveendarcshapev2Other.EndY {
		diffs = append(diffs, bottomstackgrowthcurveendarcshapev2.GongMarshallField(stage, "EndY"))
	}
	if bottomstackgrowthcurveendarcshapev2.XAxisRotation != bottomstackgrowthcurveendarcshapev2Other.XAxisRotation {
		diffs = append(diffs, bottomstackgrowthcurveendarcshapev2.GongMarshallField(stage, "XAxisRotation"))
	}
	if bottomstackgrowthcurveendarcshapev2.LargeArcFlag != bottomstackgrowthcurveendarcshapev2Other.LargeArcFlag {
		diffs = append(diffs, bottomstackgrowthcurveendarcshapev2.GongMarshallField(stage, "LargeArcFlag"))
	}
	if bottomstackgrowthcurveendarcshapev2.SweepFlag != bottomstackgrowthcurveendarcshapev2Other.SweepFlag {
		diffs = append(diffs, bottomstackgrowthcurveendarcshapev2.GongMarshallField(stage, "SweepFlag"))
	}
	if bottomstackgrowthcurveendarcshapev2.RadiusX != bottomstackgrowthcurveendarcshapev2Other.RadiusX {
		diffs = append(diffs, bottomstackgrowthcurveendarcshapev2.GongMarshallField(stage, "RadiusX"))
	}
	if bottomstackgrowthcurveendarcshapev2.RadiusY != bottomstackgrowthcurveendarcshapev2Other.RadiusY {
		diffs = append(diffs, bottomstackgrowthcurveendarcshapev2.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) GongDiff(stage *Stage, bottomstackgrowthcurvestartarcshapev2Other *BottomStackGrowthCurveStartArcShapeV2) (diffs []string) {
	// insertion point for field diffs
	if bottomstackgrowthcurvestartarcshapev2.Name != bottomstackgrowthcurvestartarcshapev2Other.Name {
		diffs = append(diffs, bottomstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "Name"))
	}
	if bottomstackgrowthcurvestartarcshapev2.StartX != bottomstackgrowthcurvestartarcshapev2Other.StartX {
		diffs = append(diffs, bottomstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "StartX"))
	}
	if bottomstackgrowthcurvestartarcshapev2.StartY != bottomstackgrowthcurvestartarcshapev2Other.StartY {
		diffs = append(diffs, bottomstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "StartY"))
	}
	if bottomstackgrowthcurvestartarcshapev2.EndX != bottomstackgrowthcurvestartarcshapev2Other.EndX {
		diffs = append(diffs, bottomstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "EndX"))
	}
	if bottomstackgrowthcurvestartarcshapev2.EndY != bottomstackgrowthcurvestartarcshapev2Other.EndY {
		diffs = append(diffs, bottomstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "EndY"))
	}
	if bottomstackgrowthcurvestartarcshapev2.XAxisRotation != bottomstackgrowthcurvestartarcshapev2Other.XAxisRotation {
		diffs = append(diffs, bottomstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "XAxisRotation"))
	}
	if bottomstackgrowthcurvestartarcshapev2.LargeArcFlag != bottomstackgrowthcurvestartarcshapev2Other.LargeArcFlag {
		diffs = append(diffs, bottomstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "LargeArcFlag"))
	}
	if bottomstackgrowthcurvestartarcshapev2.SweepFlag != bottomstackgrowthcurvestartarcshapev2Other.SweepFlag {
		diffs = append(diffs, bottomstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "SweepFlag"))
	}
	if bottomstackgrowthcurvestartarcshapev2.RadiusX != bottomstackgrowthcurvestartarcshapev2Other.RadiusX {
		diffs = append(diffs, bottomstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "RadiusX"))
	}
	if bottomstackgrowthcurvestartarcshapev2.RadiusY != bottomstackgrowthcurvestartarcshapev2Other.RadiusY {
		diffs = append(diffs, bottomstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) GongDiff(stage *Stage, bottomstackofgrowthcurvev2Other *BottomStackOfGrowthCurveV2) (diffs []string) {
	// insertion point for field diffs
	if bottomstackofgrowthcurvev2.Name != bottomstackofgrowthcurvev2Other.Name {
		diffs = append(diffs, bottomstackofgrowthcurvev2.GongMarshallField(stage, "Name"))
	}
	BottomStackGrowthCurveStartArcShapeV2sDifferent := false
	if len(bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s) != len(bottomstackofgrowthcurvev2Other.BottomStackGrowthCurveStartArcShapeV2s) {
		BottomStackGrowthCurveStartArcShapeV2sDifferent = true
	} else {
		for i := range bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s {
			if (bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s[i] == nil) != (bottomstackofgrowthcurvev2Other.BottomStackGrowthCurveStartArcShapeV2s[i] == nil) {
				BottomStackGrowthCurveStartArcShapeV2sDifferent = true
				break
			} else if bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s[i] != nil && bottomstackofgrowthcurvev2Other.BottomStackGrowthCurveStartArcShapeV2s[i] != nil {
				// this is a pointer comparaison
				if bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s[i] != bottomstackofgrowthcurvev2Other.BottomStackGrowthCurveStartArcShapeV2s[i] {
					BottomStackGrowthCurveStartArcShapeV2sDifferent = true
					break
				}
			}
		}
	}
	if BottomStackGrowthCurveStartArcShapeV2sDifferent {
		ops := Diff(stage, bottomstackofgrowthcurvev2, bottomstackofgrowthcurvev2Other, "BottomStackGrowthCurveStartArcShapeV2s", bottomstackofgrowthcurvev2Other.BottomStackGrowthCurveStartArcShapeV2s, bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s)
		diffs = append(diffs, ops)
	}
	BottomStackGrowthCurveEndArcShapeV2sDifferent := false
	if len(bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s) != len(bottomstackofgrowthcurvev2Other.BottomStackGrowthCurveEndArcShapeV2s) {
		BottomStackGrowthCurveEndArcShapeV2sDifferent = true
	} else {
		for i := range bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s {
			if (bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s[i] == nil) != (bottomstackofgrowthcurvev2Other.BottomStackGrowthCurveEndArcShapeV2s[i] == nil) {
				BottomStackGrowthCurveEndArcShapeV2sDifferent = true
				break
			} else if bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s[i] != nil && bottomstackofgrowthcurvev2Other.BottomStackGrowthCurveEndArcShapeV2s[i] != nil {
				// this is a pointer comparaison
				if bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s[i] != bottomstackofgrowthcurvev2Other.BottomStackGrowthCurveEndArcShapeV2s[i] {
					BottomStackGrowthCurveEndArcShapeV2sDifferent = true
					break
				}
			}
		}
	}
	if BottomStackGrowthCurveEndArcShapeV2sDifferent {
		ops := Diff(stage, bottomstackofgrowthcurvev2, bottomstackofgrowthcurvev2Other, "BottomStackGrowthCurveEndArcShapeV2s", bottomstackofgrowthcurvev2Other.BottomStackGrowthCurveEndArcShapeV2s, bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bottomstartarcshapev2 *BottomStartArcShapeV2) GongDiff(stage *Stage, bottomstartarcshapev2Other *BottomStartArcShapeV2) (diffs []string) {
	// insertion point for field diffs
	if bottomstartarcshapev2.Name != bottomstartarcshapev2Other.Name {
		diffs = append(diffs, bottomstartarcshapev2.GongMarshallField(stage, "Name"))
	}
	if bottomstartarcshapev2.StartX != bottomstartarcshapev2Other.StartX {
		diffs = append(diffs, bottomstartarcshapev2.GongMarshallField(stage, "StartX"))
	}
	if bottomstartarcshapev2.StartY != bottomstartarcshapev2Other.StartY {
		diffs = append(diffs, bottomstartarcshapev2.GongMarshallField(stage, "StartY"))
	}
	if bottomstartarcshapev2.EndX != bottomstartarcshapev2Other.EndX {
		diffs = append(diffs, bottomstartarcshapev2.GongMarshallField(stage, "EndX"))
	}
	if bottomstartarcshapev2.EndY != bottomstartarcshapev2Other.EndY {
		diffs = append(diffs, bottomstartarcshapev2.GongMarshallField(stage, "EndY"))
	}
	if bottomstartarcshapev2.XAxisRotation != bottomstartarcshapev2Other.XAxisRotation {
		diffs = append(diffs, bottomstartarcshapev2.GongMarshallField(stage, "XAxisRotation"))
	}
	if bottomstartarcshapev2.LargeArcFlag != bottomstartarcshapev2Other.LargeArcFlag {
		diffs = append(diffs, bottomstartarcshapev2.GongMarshallField(stage, "LargeArcFlag"))
	}
	if bottomstartarcshapev2.SweepFlag != bottomstartarcshapev2Other.SweepFlag {
		diffs = append(diffs, bottomstartarcshapev2.GongMarshallField(stage, "SweepFlag"))
	}
	if bottomstartarcshapev2.RadiusX != bottomstartarcshapev2Other.RadiusX {
		diffs = append(diffs, bottomstartarcshapev2.GongMarshallField(stage, "RadiusX"))
	}
	if bottomstartarcshapev2.RadiusY != bottomstartarcshapev2Other.RadiusY {
		diffs = append(diffs, bottomstartarcshapev2.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) GongDiff(stage *Stage, bottomstartarcshapev2gridOther *BottomStartArcShapeV2Grid) (diffs []string) {
	// insertion point for field diffs
	if bottomstartarcshapev2grid.Name != bottomstartarcshapev2gridOther.Name {
		diffs = append(diffs, bottomstartarcshapev2grid.GongMarshallField(stage, "Name"))
	}
	BottomStartArcShapesV2Different := false
	if len(bottomstartarcshapev2grid.BottomStartArcShapesV2) != len(bottomstartarcshapev2gridOther.BottomStartArcShapesV2) {
		BottomStartArcShapesV2Different = true
	} else {
		for i := range bottomstartarcshapev2grid.BottomStartArcShapesV2 {
			if (bottomstartarcshapev2grid.BottomStartArcShapesV2[i] == nil) != (bottomstartarcshapev2gridOther.BottomStartArcShapesV2[i] == nil) {
				BottomStartArcShapesV2Different = true
				break
			} else if bottomstartarcshapev2grid.BottomStartArcShapesV2[i] != nil && bottomstartarcshapev2gridOther.BottomStartArcShapesV2[i] != nil {
				// this is a pointer comparaison
				if bottomstartarcshapev2grid.BottomStartArcShapesV2[i] != bottomstartarcshapev2gridOther.BottomStartArcShapesV2[i] {
					BottomStartArcShapesV2Different = true
					break
				}
			}
		}
	}
	if BottomStartArcShapesV2Different {
		ops := Diff(stage, bottomstartarcshapev2grid, bottomstartarcshapev2gridOther, "BottomStartArcShapesV2", bottomstartarcshapev2gridOther.BottomStartArcShapesV2, bottomstartarcshapev2grid.BottomStartArcShapesV2)
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
func (endarcshapev2 *EndArcShapeV2) GongDiff(stage *Stage, endarcshapev2Other *EndArcShapeV2) (diffs []string) {
	// insertion point for field diffs
	if endarcshapev2.Name != endarcshapev2Other.Name {
		diffs = append(diffs, endarcshapev2.GongMarshallField(stage, "Name"))
	}
	if endarcshapev2.StartX != endarcshapev2Other.StartX {
		diffs = append(diffs, endarcshapev2.GongMarshallField(stage, "StartX"))
	}
	if endarcshapev2.StartY != endarcshapev2Other.StartY {
		diffs = append(diffs, endarcshapev2.GongMarshallField(stage, "StartY"))
	}
	if endarcshapev2.EndX != endarcshapev2Other.EndX {
		diffs = append(diffs, endarcshapev2.GongMarshallField(stage, "EndX"))
	}
	if endarcshapev2.EndY != endarcshapev2Other.EndY {
		diffs = append(diffs, endarcshapev2.GongMarshallField(stage, "EndY"))
	}
	if endarcshapev2.XAxisRotation != endarcshapev2Other.XAxisRotation {
		diffs = append(diffs, endarcshapev2.GongMarshallField(stage, "XAxisRotation"))
	}
	if endarcshapev2.LargeArcFlag != endarcshapev2Other.LargeArcFlag {
		diffs = append(diffs, endarcshapev2.GongMarshallField(stage, "LargeArcFlag"))
	}
	if endarcshapev2.SweepFlag != endarcshapev2Other.SweepFlag {
		diffs = append(diffs, endarcshapev2.GongMarshallField(stage, "SweepFlag"))
	}
	if endarcshapev2.RadiusX != endarcshapev2Other.RadiusX {
		diffs = append(diffs, endarcshapev2.GongMarshallField(stage, "RadiusX"))
	}
	if endarcshapev2.RadiusY != endarcshapev2Other.RadiusY {
		diffs = append(diffs, endarcshapev2.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (endarcshapev2grid *EndArcShapeV2Grid) GongDiff(stage *Stage, endarcshapev2gridOther *EndArcShapeV2Grid) (diffs []string) {
	// insertion point for field diffs
	if endarcshapev2grid.Name != endarcshapev2gridOther.Name {
		diffs = append(diffs, endarcshapev2grid.GongMarshallField(stage, "Name"))
	}
	EndArcShapesV2Different := false
	if len(endarcshapev2grid.EndArcShapesV2) != len(endarcshapev2gridOther.EndArcShapesV2) {
		EndArcShapesV2Different = true
	} else {
		for i := range endarcshapev2grid.EndArcShapesV2 {
			if (endarcshapev2grid.EndArcShapesV2[i] == nil) != (endarcshapev2gridOther.EndArcShapesV2[i] == nil) {
				EndArcShapesV2Different = true
				break
			} else if endarcshapev2grid.EndArcShapesV2[i] != nil && endarcshapev2gridOther.EndArcShapesV2[i] != nil {
				// this is a pointer comparaison
				if endarcshapev2grid.EndArcShapesV2[i] != endarcshapev2gridOther.EndArcShapesV2[i] {
					EndArcShapesV2Different = true
					break
				}
			}
		}
	}
	if EndArcShapesV2Different {
		ops := Diff(stage, endarcshapev2grid, endarcshapev2gridOther, "EndArcShapesV2", endarcshapev2gridOther.EndArcShapesV2, endarcshapev2grid.EndArcShapesV2)
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
	if (growthcurve2d.StartArcShapeV2Grid == nil) != (growthcurve2dOther.StartArcShapeV2Grid == nil) {
		diffs = append(diffs, growthcurve2d.GongMarshallField(stage, "StartArcShapeV2Grid"))
	} else if growthcurve2d.StartArcShapeV2Grid != nil && growthcurve2dOther.StartArcShapeV2Grid != nil {
		if growthcurve2d.StartArcShapeV2Grid != growthcurve2dOther.StartArcShapeV2Grid {
			diffs = append(diffs, growthcurve2d.GongMarshallField(stage, "StartArcShapeV2Grid"))
		}
	}
	if (growthcurve2d.EndArcShapeV2Grid == nil) != (growthcurve2dOther.EndArcShapeV2Grid == nil) {
		diffs = append(diffs, growthcurve2d.GongMarshallField(stage, "EndArcShapeV2Grid"))
	} else if growthcurve2d.EndArcShapeV2Grid != nil && growthcurve2dOther.EndArcShapeV2Grid != nil {
		if growthcurve2d.EndArcShapeV2Grid != growthcurve2dOther.EndArcShapeV2Grid {
			diffs = append(diffs, growthcurve2d.GongMarshallField(stage, "EndArcShapeV2Grid"))
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
	if plant.Thickness != plantOther.Thickness {
		diffs = append(diffs, plant.GongMarshallField(stage, "Thickness"))
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
	if (plant.StartArcShapeV2Grid == nil) != (plantOther.StartArcShapeV2Grid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "StartArcShapeV2Grid"))
	} else if plant.StartArcShapeV2Grid != nil && plantOther.StartArcShapeV2Grid != nil {
		if plant.StartArcShapeV2Grid != plantOther.StartArcShapeV2Grid {
			diffs = append(diffs, plant.GongMarshallField(stage, "StartArcShapeV2Grid"))
		}
	}
	if (plant.TopStartArcShapeV2Grid == nil) != (plantOther.TopStartArcShapeV2Grid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "TopStartArcShapeV2Grid"))
	} else if plant.TopStartArcShapeV2Grid != nil && plantOther.TopStartArcShapeV2Grid != nil {
		if plant.TopStartArcShapeV2Grid != plantOther.TopStartArcShapeV2Grid {
			diffs = append(diffs, plant.GongMarshallField(stage, "TopStartArcShapeV2Grid"))
		}
	}
	if (plant.EndArcShapeV2Grid == nil) != (plantOther.EndArcShapeV2Grid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "EndArcShapeV2Grid"))
	} else if plant.EndArcShapeV2Grid != nil && plantOther.EndArcShapeV2Grid != nil {
		if plant.EndArcShapeV2Grid != plantOther.EndArcShapeV2Grid {
			diffs = append(diffs, plant.GongMarshallField(stage, "EndArcShapeV2Grid"))
		}
	}
	if (plant.TopEndArcShapeV2Grid == nil) != (plantOther.TopEndArcShapeV2Grid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "TopEndArcShapeV2Grid"))
	} else if plant.TopEndArcShapeV2Grid != nil && plantOther.TopEndArcShapeV2Grid != nil {
		if plant.TopEndArcShapeV2Grid != plantOther.TopEndArcShapeV2Grid {
			diffs = append(diffs, plant.GongMarshallField(stage, "TopEndArcShapeV2Grid"))
		}
	}
	if (plant.BottomStartArcShapeV2Grid == nil) != (plantOther.BottomStartArcShapeV2Grid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "BottomStartArcShapeV2Grid"))
	} else if plant.BottomStartArcShapeV2Grid != nil && plantOther.BottomStartArcShapeV2Grid != nil {
		if plant.BottomStartArcShapeV2Grid != plantOther.BottomStartArcShapeV2Grid {
			diffs = append(diffs, plant.GongMarshallField(stage, "BottomStartArcShapeV2Grid"))
		}
	}
	if (plant.BottomEndArcShapeV2Grid == nil) != (plantOther.BottomEndArcShapeV2Grid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "BottomEndArcShapeV2Grid"))
	} else if plant.BottomEndArcShapeV2Grid != nil && plantOther.BottomEndArcShapeV2Grid != nil {
		if plant.BottomEndArcShapeV2Grid != plantOther.BottomEndArcShapeV2Grid {
			diffs = append(diffs, plant.GongMarshallField(stage, "BottomEndArcShapeV2Grid"))
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
	if (plant.StackOfGrowthCurveV2 == nil) != (plantOther.StackOfGrowthCurveV2 == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "StackOfGrowthCurveV2"))
	} else if plant.StackOfGrowthCurveV2 != nil && plantOther.StackOfGrowthCurveV2 != nil {
		if plant.StackOfGrowthCurveV2 != plantOther.StackOfGrowthCurveV2 {
			diffs = append(diffs, plant.GongMarshallField(stage, "StackOfGrowthCurveV2"))
		}
	}
	if (plant.TopStackOfGrowthCurveV2 == nil) != (plantOther.TopStackOfGrowthCurveV2 == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "TopStackOfGrowthCurveV2"))
	} else if plant.TopStackOfGrowthCurveV2 != nil && plantOther.TopStackOfGrowthCurveV2 != nil {
		if plant.TopStackOfGrowthCurveV2 != plantOther.TopStackOfGrowthCurveV2 {
			diffs = append(diffs, plant.GongMarshallField(stage, "TopStackOfGrowthCurveV2"))
		}
	}
	if (plant.BottomStackOfGrowthCurveV2 == nil) != (plantOther.BottomStackOfGrowthCurveV2 == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "BottomStackOfGrowthCurveV2"))
	} else if plant.BottomStackOfGrowthCurveV2 != nil && plantOther.BottomStackOfGrowthCurveV2 != nil {
		if plant.BottomStackOfGrowthCurveV2 != plantOther.BottomStackOfGrowthCurveV2 {
			diffs = append(diffs, plant.GongMarshallField(stage, "BottomStackOfGrowthCurveV2"))
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
	if plantdiagram.IsHiddenStartArcShapeV2Grid != plantdiagramOther.IsHiddenStartArcShapeV2Grid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenStartArcShapeV2Grid"))
	}
	if plantdiagram.IsHiddenTopStartArcShapeV2Grid != plantdiagramOther.IsHiddenTopStartArcShapeV2Grid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenTopStartArcShapeV2Grid"))
	}
	if plantdiagram.IsHiddenEndArcShapeGrid != plantdiagramOther.IsHiddenEndArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenEndArcShapeGrid"))
	}
	if plantdiagram.IsHiddenEndArcShapeV2Grid != plantdiagramOther.IsHiddenEndArcShapeV2Grid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenEndArcShapeV2Grid"))
	}
	if plantdiagram.IsHiddenTopEndArcShapeV2Grid != plantdiagramOther.IsHiddenTopEndArcShapeV2Grid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenTopEndArcShapeV2Grid"))
	}
	if plantdiagram.IsHiddenBottomStartArcShapeV2Grid != plantdiagramOther.IsHiddenBottomStartArcShapeV2Grid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenBottomStartArcShapeV2Grid"))
	}
	if plantdiagram.IsHiddenBottomEndArcShapeV2Grid != plantdiagramOther.IsHiddenBottomEndArcShapeV2Grid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenBottomEndArcShapeV2Grid"))
	}
	if plantdiagram.IsHiddenGrowthCurveBezierShapeGrid != plantdiagramOther.IsHiddenGrowthCurveBezierShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenGrowthCurveBezierShapeGrid"))
	}
	if plantdiagram.IsHiddenStackOfGrowthCurve != plantdiagramOther.IsHiddenStackOfGrowthCurve {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurve"))
	}
	if plantdiagram.IsHiddenStackOfGrowthCurveV2 != plantdiagramOther.IsHiddenStackOfGrowthCurveV2 {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurveV2"))
	}
	if plantdiagram.IsHiddenTopStackOfGrowthCurveV2 != plantdiagramOther.IsHiddenTopStackOfGrowthCurveV2 {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenTopStackOfGrowthCurveV2"))
	}
	if plantdiagram.IsHiddenBottomStackOfGrowthCurveV2 != plantdiagramOther.IsHiddenBottomStackOfGrowthCurveV2 {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenBottomStackOfGrowthCurveV2"))
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
func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GongDiff(stage *Stage, stackgrowthcurvebeziershapeOther *StackGrowthCurveBezierShape) (diffs []string) {
	// insertion point for field diffs
	if stackgrowthcurvebeziershape.Name != stackgrowthcurvebeziershapeOther.Name {
		diffs = append(diffs, stackgrowthcurvebeziershape.GongMarshallField(stage, "Name"))
	}
	if stackgrowthcurvebeziershape.StartX != stackgrowthcurvebeziershapeOther.StartX {
		diffs = append(diffs, stackgrowthcurvebeziershape.GongMarshallField(stage, "StartX"))
	}
	if stackgrowthcurvebeziershape.StartY != stackgrowthcurvebeziershapeOther.StartY {
		diffs = append(diffs, stackgrowthcurvebeziershape.GongMarshallField(stage, "StartY"))
	}
	if stackgrowthcurvebeziershape.ControlPointStartX != stackgrowthcurvebeziershapeOther.ControlPointStartX {
		diffs = append(diffs, stackgrowthcurvebeziershape.GongMarshallField(stage, "ControlPointStartX"))
	}
	if stackgrowthcurvebeziershape.ControlPointStartY != stackgrowthcurvebeziershapeOther.ControlPointStartY {
		diffs = append(diffs, stackgrowthcurvebeziershape.GongMarshallField(stage, "ControlPointStartY"))
	}
	if stackgrowthcurvebeziershape.EndX != stackgrowthcurvebeziershapeOther.EndX {
		diffs = append(diffs, stackgrowthcurvebeziershape.GongMarshallField(stage, "EndX"))
	}
	if stackgrowthcurvebeziershape.EndY != stackgrowthcurvebeziershapeOther.EndY {
		diffs = append(diffs, stackgrowthcurvebeziershape.GongMarshallField(stage, "EndY"))
	}
	if stackgrowthcurvebeziershape.ControlPointEndX != stackgrowthcurvebeziershapeOther.ControlPointEndX {
		diffs = append(diffs, stackgrowthcurvebeziershape.GongMarshallField(stage, "ControlPointEndX"))
	}
	if stackgrowthcurvebeziershape.ControlPointEndY != stackgrowthcurvebeziershapeOther.ControlPointEndY {
		diffs = append(diffs, stackgrowthcurvebeziershape.GongMarshallField(stage, "ControlPointEndY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) GongDiff(stage *Stage, stackgrowthcurveendarcshapev2Other *StackGrowthCurveEndArcShapeV2) (diffs []string) {
	// insertion point for field diffs
	if stackgrowthcurveendarcshapev2.Name != stackgrowthcurveendarcshapev2Other.Name {
		diffs = append(diffs, stackgrowthcurveendarcshapev2.GongMarshallField(stage, "Name"))
	}
	if stackgrowthcurveendarcshapev2.StartX != stackgrowthcurveendarcshapev2Other.StartX {
		diffs = append(diffs, stackgrowthcurveendarcshapev2.GongMarshallField(stage, "StartX"))
	}
	if stackgrowthcurveendarcshapev2.StartY != stackgrowthcurveendarcshapev2Other.StartY {
		diffs = append(diffs, stackgrowthcurveendarcshapev2.GongMarshallField(stage, "StartY"))
	}
	if stackgrowthcurveendarcshapev2.EndX != stackgrowthcurveendarcshapev2Other.EndX {
		diffs = append(diffs, stackgrowthcurveendarcshapev2.GongMarshallField(stage, "EndX"))
	}
	if stackgrowthcurveendarcshapev2.EndY != stackgrowthcurveendarcshapev2Other.EndY {
		diffs = append(diffs, stackgrowthcurveendarcshapev2.GongMarshallField(stage, "EndY"))
	}
	if stackgrowthcurveendarcshapev2.XAxisRotation != stackgrowthcurveendarcshapev2Other.XAxisRotation {
		diffs = append(diffs, stackgrowthcurveendarcshapev2.GongMarshallField(stage, "XAxisRotation"))
	}
	if stackgrowthcurveendarcshapev2.LargeArcFlag != stackgrowthcurveendarcshapev2Other.LargeArcFlag {
		diffs = append(diffs, stackgrowthcurveendarcshapev2.GongMarshallField(stage, "LargeArcFlag"))
	}
	if stackgrowthcurveendarcshapev2.SweepFlag != stackgrowthcurveendarcshapev2Other.SweepFlag {
		diffs = append(diffs, stackgrowthcurveendarcshapev2.GongMarshallField(stage, "SweepFlag"))
	}
	if stackgrowthcurveendarcshapev2.RadiusX != stackgrowthcurveendarcshapev2Other.RadiusX {
		diffs = append(diffs, stackgrowthcurveendarcshapev2.GongMarshallField(stage, "RadiusX"))
	}
	if stackgrowthcurveendarcshapev2.RadiusY != stackgrowthcurveendarcshapev2Other.RadiusY {
		diffs = append(diffs, stackgrowthcurveendarcshapev2.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) GongDiff(stage *Stage, stackgrowthcurvestartarcshapev2Other *StackGrowthCurveStartArcShapeV2) (diffs []string) {
	// insertion point for field diffs
	if stackgrowthcurvestartarcshapev2.Name != stackgrowthcurvestartarcshapev2Other.Name {
		diffs = append(diffs, stackgrowthcurvestartarcshapev2.GongMarshallField(stage, "Name"))
	}
	if stackgrowthcurvestartarcshapev2.StartX != stackgrowthcurvestartarcshapev2Other.StartX {
		diffs = append(diffs, stackgrowthcurvestartarcshapev2.GongMarshallField(stage, "StartX"))
	}
	if stackgrowthcurvestartarcshapev2.StartY != stackgrowthcurvestartarcshapev2Other.StartY {
		diffs = append(diffs, stackgrowthcurvestartarcshapev2.GongMarshallField(stage, "StartY"))
	}
	if stackgrowthcurvestartarcshapev2.EndX != stackgrowthcurvestartarcshapev2Other.EndX {
		diffs = append(diffs, stackgrowthcurvestartarcshapev2.GongMarshallField(stage, "EndX"))
	}
	if stackgrowthcurvestartarcshapev2.EndY != stackgrowthcurvestartarcshapev2Other.EndY {
		diffs = append(diffs, stackgrowthcurvestartarcshapev2.GongMarshallField(stage, "EndY"))
	}
	if stackgrowthcurvestartarcshapev2.XAxisRotation != stackgrowthcurvestartarcshapev2Other.XAxisRotation {
		diffs = append(diffs, stackgrowthcurvestartarcshapev2.GongMarshallField(stage, "XAxisRotation"))
	}
	if stackgrowthcurvestartarcshapev2.LargeArcFlag != stackgrowthcurvestartarcshapev2Other.LargeArcFlag {
		diffs = append(diffs, stackgrowthcurvestartarcshapev2.GongMarshallField(stage, "LargeArcFlag"))
	}
	if stackgrowthcurvestartarcshapev2.SweepFlag != stackgrowthcurvestartarcshapev2Other.SweepFlag {
		diffs = append(diffs, stackgrowthcurvestartarcshapev2.GongMarshallField(stage, "SweepFlag"))
	}
	if stackgrowthcurvestartarcshapev2.RadiusX != stackgrowthcurvestartarcshapev2Other.RadiusX {
		diffs = append(diffs, stackgrowthcurvestartarcshapev2.GongMarshallField(stage, "RadiusX"))
	}
	if stackgrowthcurvestartarcshapev2.RadiusY != stackgrowthcurvestartarcshapev2Other.RadiusY {
		diffs = append(diffs, stackgrowthcurvestartarcshapev2.GongMarshallField(stage, "RadiusY"))
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
	StackGrowthCurveBezierShapesDifferent := false
	if len(stackofgrowthcurve.StackGrowthCurveBezierShapes) != len(stackofgrowthcurveOther.StackGrowthCurveBezierShapes) {
		StackGrowthCurveBezierShapesDifferent = true
	} else {
		for i := range stackofgrowthcurve.StackGrowthCurveBezierShapes {
			if (stackofgrowthcurve.StackGrowthCurveBezierShapes[i] == nil) != (stackofgrowthcurveOther.StackGrowthCurveBezierShapes[i] == nil) {
				StackGrowthCurveBezierShapesDifferent = true
				break
			} else if stackofgrowthcurve.StackGrowthCurveBezierShapes[i] != nil && stackofgrowthcurveOther.StackGrowthCurveBezierShapes[i] != nil {
				// this is a pointer comparaison
				if stackofgrowthcurve.StackGrowthCurveBezierShapes[i] != stackofgrowthcurveOther.StackGrowthCurveBezierShapes[i] {
					StackGrowthCurveBezierShapesDifferent = true
					break
				}
			}
		}
	}
	if StackGrowthCurveBezierShapesDifferent {
		ops := Diff(stage, stackofgrowthcurve, stackofgrowthcurveOther, "StackGrowthCurveBezierShapes", stackofgrowthcurveOther.StackGrowthCurveBezierShapes, stackofgrowthcurve.StackGrowthCurveBezierShapes)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) GongDiff(stage *Stage, stackofgrowthcurvev2Other *StackOfGrowthCurveV2) (diffs []string) {
	// insertion point for field diffs
	if stackofgrowthcurvev2.Name != stackofgrowthcurvev2Other.Name {
		diffs = append(diffs, stackofgrowthcurvev2.GongMarshallField(stage, "Name"))
	}
	StackGrowthCurveStartArcShapeV2sDifferent := false
	if len(stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s) != len(stackofgrowthcurvev2Other.StackGrowthCurveStartArcShapeV2s) {
		StackGrowthCurveStartArcShapeV2sDifferent = true
	} else {
		for i := range stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s {
			if (stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s[i] == nil) != (stackofgrowthcurvev2Other.StackGrowthCurveStartArcShapeV2s[i] == nil) {
				StackGrowthCurveStartArcShapeV2sDifferent = true
				break
			} else if stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s[i] != nil && stackofgrowthcurvev2Other.StackGrowthCurveStartArcShapeV2s[i] != nil {
				// this is a pointer comparaison
				if stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s[i] != stackofgrowthcurvev2Other.StackGrowthCurveStartArcShapeV2s[i] {
					StackGrowthCurveStartArcShapeV2sDifferent = true
					break
				}
			}
		}
	}
	if StackGrowthCurveStartArcShapeV2sDifferent {
		ops := Diff(stage, stackofgrowthcurvev2, stackofgrowthcurvev2Other, "StackGrowthCurveStartArcShapeV2s", stackofgrowthcurvev2Other.StackGrowthCurveStartArcShapeV2s, stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s)
		diffs = append(diffs, ops)
	}
	StackGrowthCurveEndArcShapeV2sDifferent := false
	if len(stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s) != len(stackofgrowthcurvev2Other.StackGrowthCurveEndArcShapeV2s) {
		StackGrowthCurveEndArcShapeV2sDifferent = true
	} else {
		for i := range stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s {
			if (stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s[i] == nil) != (stackofgrowthcurvev2Other.StackGrowthCurveEndArcShapeV2s[i] == nil) {
				StackGrowthCurveEndArcShapeV2sDifferent = true
				break
			} else if stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s[i] != nil && stackofgrowthcurvev2Other.StackGrowthCurveEndArcShapeV2s[i] != nil {
				// this is a pointer comparaison
				if stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s[i] != stackofgrowthcurvev2Other.StackGrowthCurveEndArcShapeV2s[i] {
					StackGrowthCurveEndArcShapeV2sDifferent = true
					break
				}
			}
		}
	}
	if StackGrowthCurveEndArcShapeV2sDifferent {
		ops := Diff(stage, stackofgrowthcurvev2, stackofgrowthcurvev2Other, "StackGrowthCurveEndArcShapeV2s", stackofgrowthcurvev2Other.StackGrowthCurveEndArcShapeV2s, stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (startarcshapev2 *StartArcShapeV2) GongDiff(stage *Stage, startarcshapev2Other *StartArcShapeV2) (diffs []string) {
	// insertion point for field diffs
	if startarcshapev2.Name != startarcshapev2Other.Name {
		diffs = append(diffs, startarcshapev2.GongMarshallField(stage, "Name"))
	}
	if startarcshapev2.StartX != startarcshapev2Other.StartX {
		diffs = append(diffs, startarcshapev2.GongMarshallField(stage, "StartX"))
	}
	if startarcshapev2.StartY != startarcshapev2Other.StartY {
		diffs = append(diffs, startarcshapev2.GongMarshallField(stage, "StartY"))
	}
	if startarcshapev2.EndX != startarcshapev2Other.EndX {
		diffs = append(diffs, startarcshapev2.GongMarshallField(stage, "EndX"))
	}
	if startarcshapev2.EndY != startarcshapev2Other.EndY {
		diffs = append(diffs, startarcshapev2.GongMarshallField(stage, "EndY"))
	}
	if startarcshapev2.XAxisRotation != startarcshapev2Other.XAxisRotation {
		diffs = append(diffs, startarcshapev2.GongMarshallField(stage, "XAxisRotation"))
	}
	if startarcshapev2.LargeArcFlag != startarcshapev2Other.LargeArcFlag {
		diffs = append(diffs, startarcshapev2.GongMarshallField(stage, "LargeArcFlag"))
	}
	if startarcshapev2.SweepFlag != startarcshapev2Other.SweepFlag {
		diffs = append(diffs, startarcshapev2.GongMarshallField(stage, "SweepFlag"))
	}
	if startarcshapev2.RadiusX != startarcshapev2Other.RadiusX {
		diffs = append(diffs, startarcshapev2.GongMarshallField(stage, "RadiusX"))
	}
	if startarcshapev2.RadiusY != startarcshapev2Other.RadiusY {
		diffs = append(diffs, startarcshapev2.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (startarcshapev2grid *StartArcShapeV2Grid) GongDiff(stage *Stage, startarcshapev2gridOther *StartArcShapeV2Grid) (diffs []string) {
	// insertion point for field diffs
	if startarcshapev2grid.Name != startarcshapev2gridOther.Name {
		diffs = append(diffs, startarcshapev2grid.GongMarshallField(stage, "Name"))
	}
	StartArcShapesV2Different := false
	if len(startarcshapev2grid.StartArcShapesV2) != len(startarcshapev2gridOther.StartArcShapesV2) {
		StartArcShapesV2Different = true
	} else {
		for i := range startarcshapev2grid.StartArcShapesV2 {
			if (startarcshapev2grid.StartArcShapesV2[i] == nil) != (startarcshapev2gridOther.StartArcShapesV2[i] == nil) {
				StartArcShapesV2Different = true
				break
			} else if startarcshapev2grid.StartArcShapesV2[i] != nil && startarcshapev2gridOther.StartArcShapesV2[i] != nil {
				// this is a pointer comparaison
				if startarcshapev2grid.StartArcShapesV2[i] != startarcshapev2gridOther.StartArcShapesV2[i] {
					StartArcShapesV2Different = true
					break
				}
			}
		}
	}
	if StartArcShapesV2Different {
		ops := Diff(stage, startarcshapev2grid, startarcshapev2gridOther, "StartArcShapesV2", startarcshapev2gridOther.StartArcShapesV2, startarcshapev2grid.StartArcShapesV2)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topendarcshapev2 *TopEndArcShapeV2) GongDiff(stage *Stage, topendarcshapev2Other *TopEndArcShapeV2) (diffs []string) {
	// insertion point for field diffs
	if topendarcshapev2.Name != topendarcshapev2Other.Name {
		diffs = append(diffs, topendarcshapev2.GongMarshallField(stage, "Name"))
	}
	if topendarcshapev2.StartX != topendarcshapev2Other.StartX {
		diffs = append(diffs, topendarcshapev2.GongMarshallField(stage, "StartX"))
	}
	if topendarcshapev2.StartY != topendarcshapev2Other.StartY {
		diffs = append(diffs, topendarcshapev2.GongMarshallField(stage, "StartY"))
	}
	if topendarcshapev2.EndX != topendarcshapev2Other.EndX {
		diffs = append(diffs, topendarcshapev2.GongMarshallField(stage, "EndX"))
	}
	if topendarcshapev2.EndY != topendarcshapev2Other.EndY {
		diffs = append(diffs, topendarcshapev2.GongMarshallField(stage, "EndY"))
	}
	if topendarcshapev2.XAxisRotation != topendarcshapev2Other.XAxisRotation {
		diffs = append(diffs, topendarcshapev2.GongMarshallField(stage, "XAxisRotation"))
	}
	if topendarcshapev2.LargeArcFlag != topendarcshapev2Other.LargeArcFlag {
		diffs = append(diffs, topendarcshapev2.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topendarcshapev2.SweepFlag != topendarcshapev2Other.SweepFlag {
		diffs = append(diffs, topendarcshapev2.GongMarshallField(stage, "SweepFlag"))
	}
	if topendarcshapev2.RadiusX != topendarcshapev2Other.RadiusX {
		diffs = append(diffs, topendarcshapev2.GongMarshallField(stage, "RadiusX"))
	}
	if topendarcshapev2.RadiusY != topendarcshapev2Other.RadiusY {
		diffs = append(diffs, topendarcshapev2.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topendarcshapev2grid *TopEndArcShapeV2Grid) GongDiff(stage *Stage, topendarcshapev2gridOther *TopEndArcShapeV2Grid) (diffs []string) {
	// insertion point for field diffs
	if topendarcshapev2grid.Name != topendarcshapev2gridOther.Name {
		diffs = append(diffs, topendarcshapev2grid.GongMarshallField(stage, "Name"))
	}
	TopEndArcShapesV2Different := false
	if len(topendarcshapev2grid.TopEndArcShapesV2) != len(topendarcshapev2gridOther.TopEndArcShapesV2) {
		TopEndArcShapesV2Different = true
	} else {
		for i := range topendarcshapev2grid.TopEndArcShapesV2 {
			if (topendarcshapev2grid.TopEndArcShapesV2[i] == nil) != (topendarcshapev2gridOther.TopEndArcShapesV2[i] == nil) {
				TopEndArcShapesV2Different = true
				break
			} else if topendarcshapev2grid.TopEndArcShapesV2[i] != nil && topendarcshapev2gridOther.TopEndArcShapesV2[i] != nil {
				// this is a pointer comparaison
				if topendarcshapev2grid.TopEndArcShapesV2[i] != topendarcshapev2gridOther.TopEndArcShapesV2[i] {
					TopEndArcShapesV2Different = true
					break
				}
			}
		}
	}
	if TopEndArcShapesV2Different {
		ops := Diff(stage, topendarcshapev2grid, topendarcshapev2gridOther, "TopEndArcShapesV2", topendarcshapev2gridOther.TopEndArcShapesV2, topendarcshapev2grid.TopEndArcShapesV2)
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
	if (topgrowthcurve2d.TopStartArcShapeV2Grid == nil) != (topgrowthcurve2dOther.TopStartArcShapeV2Grid == nil) {
		diffs = append(diffs, topgrowthcurve2d.GongMarshallField(stage, "TopStartArcShapeV2Grid"))
	} else if topgrowthcurve2d.TopStartArcShapeV2Grid != nil && topgrowthcurve2dOther.TopStartArcShapeV2Grid != nil {
		if topgrowthcurve2d.TopStartArcShapeV2Grid != topgrowthcurve2dOther.TopStartArcShapeV2Grid {
			diffs = append(diffs, topgrowthcurve2d.GongMarshallField(stage, "TopStartArcShapeV2Grid"))
		}
	}
	if (topgrowthcurve2d.TopEndArcShapeV2Grid == nil) != (topgrowthcurve2dOther.TopEndArcShapeV2Grid == nil) {
		diffs = append(diffs, topgrowthcurve2d.GongMarshallField(stage, "TopEndArcShapeV2Grid"))
	} else if topgrowthcurve2d.TopEndArcShapeV2Grid != nil && topgrowthcurve2dOther.TopEndArcShapeV2Grid != nil {
		if topgrowthcurve2d.TopEndArcShapeV2Grid != topgrowthcurve2dOther.TopEndArcShapeV2Grid {
			diffs = append(diffs, topgrowthcurve2d.GongMarshallField(stage, "TopEndArcShapeV2Grid"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) GongDiff(stage *Stage, topstackgrowthcurveendarcshapev2Other *TopStackGrowthCurveEndArcShapeV2) (diffs []string) {
	// insertion point for field diffs
	if topstackgrowthcurveendarcshapev2.Name != topstackgrowthcurveendarcshapev2Other.Name {
		diffs = append(diffs, topstackgrowthcurveendarcshapev2.GongMarshallField(stage, "Name"))
	}
	if topstackgrowthcurveendarcshapev2.StartX != topstackgrowthcurveendarcshapev2Other.StartX {
		diffs = append(diffs, topstackgrowthcurveendarcshapev2.GongMarshallField(stage, "StartX"))
	}
	if topstackgrowthcurveendarcshapev2.StartY != topstackgrowthcurveendarcshapev2Other.StartY {
		diffs = append(diffs, topstackgrowthcurveendarcshapev2.GongMarshallField(stage, "StartY"))
	}
	if topstackgrowthcurveendarcshapev2.EndX != topstackgrowthcurveendarcshapev2Other.EndX {
		diffs = append(diffs, topstackgrowthcurveendarcshapev2.GongMarshallField(stage, "EndX"))
	}
	if topstackgrowthcurveendarcshapev2.EndY != topstackgrowthcurveendarcshapev2Other.EndY {
		diffs = append(diffs, topstackgrowthcurveendarcshapev2.GongMarshallField(stage, "EndY"))
	}
	if topstackgrowthcurveendarcshapev2.XAxisRotation != topstackgrowthcurveendarcshapev2Other.XAxisRotation {
		diffs = append(diffs, topstackgrowthcurveendarcshapev2.GongMarshallField(stage, "XAxisRotation"))
	}
	if topstackgrowthcurveendarcshapev2.LargeArcFlag != topstackgrowthcurveendarcshapev2Other.LargeArcFlag {
		diffs = append(diffs, topstackgrowthcurveendarcshapev2.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topstackgrowthcurveendarcshapev2.SweepFlag != topstackgrowthcurveendarcshapev2Other.SweepFlag {
		diffs = append(diffs, topstackgrowthcurveendarcshapev2.GongMarshallField(stage, "SweepFlag"))
	}
	if topstackgrowthcurveendarcshapev2.RadiusX != topstackgrowthcurveendarcshapev2Other.RadiusX {
		diffs = append(diffs, topstackgrowthcurveendarcshapev2.GongMarshallField(stage, "RadiusX"))
	}
	if topstackgrowthcurveendarcshapev2.RadiusY != topstackgrowthcurveendarcshapev2Other.RadiusY {
		diffs = append(diffs, topstackgrowthcurveendarcshapev2.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) GongDiff(stage *Stage, topstackgrowthcurvestartarcshapev2Other *TopStackGrowthCurveStartArcShapeV2) (diffs []string) {
	// insertion point for field diffs
	if topstackgrowthcurvestartarcshapev2.Name != topstackgrowthcurvestartarcshapev2Other.Name {
		diffs = append(diffs, topstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "Name"))
	}
	if topstackgrowthcurvestartarcshapev2.StartX != topstackgrowthcurvestartarcshapev2Other.StartX {
		diffs = append(diffs, topstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "StartX"))
	}
	if topstackgrowthcurvestartarcshapev2.StartY != topstackgrowthcurvestartarcshapev2Other.StartY {
		diffs = append(diffs, topstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "StartY"))
	}
	if topstackgrowthcurvestartarcshapev2.EndX != topstackgrowthcurvestartarcshapev2Other.EndX {
		diffs = append(diffs, topstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "EndX"))
	}
	if topstackgrowthcurvestartarcshapev2.EndY != topstackgrowthcurvestartarcshapev2Other.EndY {
		diffs = append(diffs, topstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "EndY"))
	}
	if topstackgrowthcurvestartarcshapev2.XAxisRotation != topstackgrowthcurvestartarcshapev2Other.XAxisRotation {
		diffs = append(diffs, topstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "XAxisRotation"))
	}
	if topstackgrowthcurvestartarcshapev2.LargeArcFlag != topstackgrowthcurvestartarcshapev2Other.LargeArcFlag {
		diffs = append(diffs, topstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topstackgrowthcurvestartarcshapev2.SweepFlag != topstackgrowthcurvestartarcshapev2Other.SweepFlag {
		diffs = append(diffs, topstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "SweepFlag"))
	}
	if topstackgrowthcurvestartarcshapev2.RadiusX != topstackgrowthcurvestartarcshapev2Other.RadiusX {
		diffs = append(diffs, topstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "RadiusX"))
	}
	if topstackgrowthcurvestartarcshapev2.RadiusY != topstackgrowthcurvestartarcshapev2Other.RadiusY {
		diffs = append(diffs, topstackgrowthcurvestartarcshapev2.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) GongDiff(stage *Stage, topstackofgrowthcurvev2Other *TopStackOfGrowthCurveV2) (diffs []string) {
	// insertion point for field diffs
	if topstackofgrowthcurvev2.Name != topstackofgrowthcurvev2Other.Name {
		diffs = append(diffs, topstackofgrowthcurvev2.GongMarshallField(stage, "Name"))
	}
	TopStackGrowthCurveStartArcShapeV2sDifferent := false
	if len(topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s) != len(topstackofgrowthcurvev2Other.TopStackGrowthCurveStartArcShapeV2s) {
		TopStackGrowthCurveStartArcShapeV2sDifferent = true
	} else {
		for i := range topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s {
			if (topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s[i] == nil) != (topstackofgrowthcurvev2Other.TopStackGrowthCurveStartArcShapeV2s[i] == nil) {
				TopStackGrowthCurveStartArcShapeV2sDifferent = true
				break
			} else if topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s[i] != nil && topstackofgrowthcurvev2Other.TopStackGrowthCurveStartArcShapeV2s[i] != nil {
				// this is a pointer comparaison
				if topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s[i] != topstackofgrowthcurvev2Other.TopStackGrowthCurveStartArcShapeV2s[i] {
					TopStackGrowthCurveStartArcShapeV2sDifferent = true
					break
				}
			}
		}
	}
	if TopStackGrowthCurveStartArcShapeV2sDifferent {
		ops := Diff(stage, topstackofgrowthcurvev2, topstackofgrowthcurvev2Other, "TopStackGrowthCurveStartArcShapeV2s", topstackofgrowthcurvev2Other.TopStackGrowthCurveStartArcShapeV2s, topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s)
		diffs = append(diffs, ops)
	}
	TopStackGrowthCurveEndArcShapeV2sDifferent := false
	if len(topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s) != len(topstackofgrowthcurvev2Other.TopStackGrowthCurveEndArcShapeV2s) {
		TopStackGrowthCurveEndArcShapeV2sDifferent = true
	} else {
		for i := range topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s {
			if (topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s[i] == nil) != (topstackofgrowthcurvev2Other.TopStackGrowthCurveEndArcShapeV2s[i] == nil) {
				TopStackGrowthCurveEndArcShapeV2sDifferent = true
				break
			} else if topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s[i] != nil && topstackofgrowthcurvev2Other.TopStackGrowthCurveEndArcShapeV2s[i] != nil {
				// this is a pointer comparaison
				if topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s[i] != topstackofgrowthcurvev2Other.TopStackGrowthCurveEndArcShapeV2s[i] {
					TopStackGrowthCurveEndArcShapeV2sDifferent = true
					break
				}
			}
		}
	}
	if TopStackGrowthCurveEndArcShapeV2sDifferent {
		ops := Diff(stage, topstackofgrowthcurvev2, topstackofgrowthcurvev2Other, "TopStackGrowthCurveEndArcShapeV2s", topstackofgrowthcurvev2Other.TopStackGrowthCurveEndArcShapeV2s, topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstartarcshapev2 *TopStartArcShapeV2) GongDiff(stage *Stage, topstartarcshapev2Other *TopStartArcShapeV2) (diffs []string) {
	// insertion point for field diffs
	if topstartarcshapev2.Name != topstartarcshapev2Other.Name {
		diffs = append(diffs, topstartarcshapev2.GongMarshallField(stage, "Name"))
	}
	if topstartarcshapev2.StartX != topstartarcshapev2Other.StartX {
		diffs = append(diffs, topstartarcshapev2.GongMarshallField(stage, "StartX"))
	}
	if topstartarcshapev2.StartY != topstartarcshapev2Other.StartY {
		diffs = append(diffs, topstartarcshapev2.GongMarshallField(stage, "StartY"))
	}
	if topstartarcshapev2.EndX != topstartarcshapev2Other.EndX {
		diffs = append(diffs, topstartarcshapev2.GongMarshallField(stage, "EndX"))
	}
	if topstartarcshapev2.EndY != topstartarcshapev2Other.EndY {
		diffs = append(diffs, topstartarcshapev2.GongMarshallField(stage, "EndY"))
	}
	if topstartarcshapev2.XAxisRotation != topstartarcshapev2Other.XAxisRotation {
		diffs = append(diffs, topstartarcshapev2.GongMarshallField(stage, "XAxisRotation"))
	}
	if topstartarcshapev2.LargeArcFlag != topstartarcshapev2Other.LargeArcFlag {
		diffs = append(diffs, topstartarcshapev2.GongMarshallField(stage, "LargeArcFlag"))
	}
	if topstartarcshapev2.SweepFlag != topstartarcshapev2Other.SweepFlag {
		diffs = append(diffs, topstartarcshapev2.GongMarshallField(stage, "SweepFlag"))
	}
	if topstartarcshapev2.RadiusX != topstartarcshapev2Other.RadiusX {
		diffs = append(diffs, topstartarcshapev2.GongMarshallField(stage, "RadiusX"))
	}
	if topstartarcshapev2.RadiusY != topstartarcshapev2Other.RadiusY {
		diffs = append(diffs, topstartarcshapev2.GongMarshallField(stage, "RadiusY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (topstartarcshapev2grid *TopStartArcShapeV2Grid) GongDiff(stage *Stage, topstartarcshapev2gridOther *TopStartArcShapeV2Grid) (diffs []string) {
	// insertion point for field diffs
	if topstartarcshapev2grid.Name != topstartarcshapev2gridOther.Name {
		diffs = append(diffs, topstartarcshapev2grid.GongMarshallField(stage, "Name"))
	}
	TopStartArcShapesV2Different := false
	if len(topstartarcshapev2grid.TopStartArcShapesV2) != len(topstartarcshapev2gridOther.TopStartArcShapesV2) {
		TopStartArcShapesV2Different = true
	} else {
		for i := range topstartarcshapev2grid.TopStartArcShapesV2 {
			if (topstartarcshapev2grid.TopStartArcShapesV2[i] == nil) != (topstartarcshapev2gridOther.TopStartArcShapesV2[i] == nil) {
				TopStartArcShapesV2Different = true
				break
			} else if topstartarcshapev2grid.TopStartArcShapesV2[i] != nil && topstartarcshapev2gridOther.TopStartArcShapesV2[i] != nil {
				// this is a pointer comparaison
				if topstartarcshapev2grid.TopStartArcShapesV2[i] != topstartarcshapev2gridOther.TopStartArcShapesV2[i] {
					TopStartArcShapesV2Different = true
					break
				}
			}
		}
	}
	if TopStartArcShapesV2Different {
		ops := Diff(stage, topstartarcshapev2grid, topstartarcshapev2gridOther, "TopStartArcShapesV2", topstartarcshapev2gridOther.TopStartArcShapesV2, topstartarcshapev2grid.TopStartArcShapesV2)
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
