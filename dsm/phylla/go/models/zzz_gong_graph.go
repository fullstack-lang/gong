// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
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

	case *ExplanationTextShape:
		ok = stage.IsStagedExplanationTextShape(target)

	case *GridPathShape:
		ok = stage.IsStagedGridPathShape(target)

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

	case *StackOfGrowthCurve:
		ok = stage.IsStagedStackOfGrowthCurve(target)

	case *StartArcShape:
		ok = stage.IsStagedStartArcShape(target)

	case *StartArcShapeGrid:
		ok = stage.IsStagedStartArcShapeGrid(target)

	case *StartArcShapeV2:
		ok = stage.IsStagedStartArcShapeV2(target)

	case *StartArcShapeV2Grid:
		ok = stage.IsStagedStartArcShapeV2Grid(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
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

	case *ExplanationTextShape:
		ok = stage.IsStagedExplanationTextShape(target)

	case *GridPathShape:
		ok = stage.IsStagedGridPathShape(target)

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

	case *StackOfGrowthCurve:
		ok = stage.IsStagedStackOfGrowthCurve(target)

	case *StartArcShape:
		ok = stage.IsStagedStartArcShape(target)

	case *StartArcShapeGrid:
		ok = stage.IsStagedStartArcShapeGrid(target)

	case *StartArcShapeV2:
		ok = stage.IsStagedStartArcShapeV2(target)

	case *StartArcShapeV2Grid:
		ok = stage.IsStagedStartArcShapeV2Grid(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
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

func (stage *Stage) IsStagedExplanationTextShape(explanationtextshape *ExplanationTextShape) (ok bool) {

	_, ok = stage.ExplanationTextShapes[explanationtextshape]

	return
}

func (stage *Stage) IsStagedGridPathShape(gridpathshape *GridPathShape) (ok bool) {

	_, ok = stage.GridPathShapes[gridpathshape]

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

func (stage *Stage) IsStagedStartArcShapeV2(startarcshapev2 *StartArcShapeV2) (ok bool) {

	_, ok = stage.StartArcShapeV2s[startarcshapev2]

	return
}

func (stage *Stage) IsStagedStartArcShapeV2Grid(startarcshapev2grid *StartArcShapeV2Grid) (ok bool) {

	_, ok = stage.StartArcShapeV2Grids[startarcshapev2grid]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
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

	case *ExplanationTextShape:
		stage.StageBranchExplanationTextShape(target)

	case *GridPathShape:
		stage.StageBranchGridPathShape(target)

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

	case *StackOfGrowthCurve:
		stage.StageBranchStackOfGrowthCurve(target)

	case *StartArcShape:
		stage.StageBranchStartArcShape(target)

	case *StartArcShapeGrid:
		stage.StageBranchStartArcShapeGrid(target)

	case *StartArcShapeV2:
		stage.StageBranchStartArcShapeV2(target)

	case *StartArcShapeV2Grid:
		stage.StageBranchStartArcShapeV2Grid(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
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
	if plant.StartArcShapeGrid != nil {
		StageBranch(stage, plant.StartArcShapeGrid)
	}
	if plant.StartArcShapeV2Grid != nil {
		StageBranch(stage, plant.StartArcShapeV2Grid)
	}
	if plant.EndArcShapeGrid != nil {
		StageBranch(stage, plant.EndArcShapeGrid)
	}
	if plant.GrowthCurveBezierShapeGrid != nil {
		StageBranch(stage, plant.GrowthCurveBezierShapeGrid)
	}
	if plant.StackOfGrowthCurve != nil {
		StageBranch(stage, plant.StackOfGrowthCurve)
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

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
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

	case *ExplanationTextShape:
		toT := CopyBranchExplanationTextShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GridPathShape:
		toT := CopyBranchGridPathShape(mapOrigCopy, fromT)
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

	case *StackOfGrowthCurve:
		toT := CopyBranchStackOfGrowthCurve(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartArcShape:
		toT := CopyBranchStartArcShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartArcShapeGrid:
		toT := CopyBranchStartArcShapeGrid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartArcShapeV2:
		toT := CopyBranchStartArcShapeV2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StartArcShapeV2Grid:
		toT := CopyBranchStartArcShapeV2Grid(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
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
	if plantFrom.StartArcShapeGrid != nil {
		plantTo.StartArcShapeGrid = CopyBranchStartArcShapeGrid(mapOrigCopy, plantFrom.StartArcShapeGrid)
	}
	if plantFrom.StartArcShapeV2Grid != nil {
		plantTo.StartArcShapeV2Grid = CopyBranchStartArcShapeV2Grid(mapOrigCopy, plantFrom.StartArcShapeV2Grid)
	}
	if plantFrom.EndArcShapeGrid != nil {
		plantTo.EndArcShapeGrid = CopyBranchEndArcShapeGrid(mapOrigCopy, plantFrom.EndArcShapeGrid)
	}
	if plantFrom.GrowthCurveBezierShapeGrid != nil {
		plantTo.GrowthCurveBezierShapeGrid = CopyBranchGrowthCurveBezierShapeGrid(mapOrigCopy, plantFrom.GrowthCurveBezierShapeGrid)
	}
	if plantFrom.StackOfGrowthCurve != nil {
		plantTo.StackOfGrowthCurve = CopyBranchStackOfGrowthCurve(mapOrigCopy, plantFrom.StackOfGrowthCurve)
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

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
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

	case *ExplanationTextShape:
		stage.UnstageBranchExplanationTextShape(target)

	case *GridPathShape:
		stage.UnstageBranchGridPathShape(target)

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

	case *StackOfGrowthCurve:
		stage.UnstageBranchStackOfGrowthCurve(target)

	case *StartArcShape:
		stage.UnstageBranchStartArcShape(target)

	case *StartArcShapeGrid:
		stage.UnstageBranchStartArcShapeGrid(target)

	case *StartArcShapeV2:
		stage.UnstageBranchStartArcShapeV2(target)

	case *StartArcShapeV2Grid:
		stage.UnstageBranchStartArcShapeV2Grid(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
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
	if plant.StartArcShapeGrid != nil {
		UnstageBranch(stage, plant.StartArcShapeGrid)
	}
	if plant.StartArcShapeV2Grid != nil {
		UnstageBranch(stage, plant.StartArcShapeV2Grid)
	}
	if plant.EndArcShapeGrid != nil {
		UnstageBranch(stage, plant.EndArcShapeGrid)
	}
	if plant.GrowthCurveBezierShapeGrid != nil {
		UnstageBranch(stage, plant.GrowthCurveBezierShapeGrid)
	}
	if plant.StackOfGrowthCurve != nil {
		UnstageBranch(stage, plant.StackOfGrowthCurve)
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

// insertion point for pointer reconstruction from references
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

func (reference *ExplanationTextShape) GongReconstructPointersFromReferences(stage *Stage, instance *ExplanationTextShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GridPathShape) GongReconstructPointersFromReferences(stage *Stage, instance *GridPathShape) {
	// insertion point for pointers field
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
	if instance.StartArcShapeGrid != nil {
		reference.StartArcShapeGrid = stage.StartArcShapeGrids_reference[instance.StartArcShapeGrid]
	}
	if instance.StartArcShapeV2Grid != nil {
		reference.StartArcShapeV2Grid = stage.StartArcShapeV2Grids_reference[instance.StartArcShapeV2Grid]
	}
	if instance.EndArcShapeGrid != nil {
		reference.EndArcShapeGrid = stage.EndArcShapeGrids_reference[instance.EndArcShapeGrid]
	}
	if instance.GrowthCurveBezierShapeGrid != nil {
		reference.GrowthCurveBezierShapeGrid = stage.GrowthCurveBezierShapeGrids_reference[instance.GrowthCurveBezierShapeGrid]
	}
	if instance.StackOfGrowthCurve != nil {
		reference.StackOfGrowthCurve = stage.StackOfGrowthCurves_reference[instance.StackOfGrowthCurve]
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

func (reference *StackOfGrowthCurve) GongReconstructPointersFromReferences(stage *Stage, instance *StackOfGrowthCurve) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.StackGrowthCurveBezierShapes = reference.StackGrowthCurveBezierShapes[:0]
	for _, _b := range instance.StackGrowthCurveBezierShapes {
		reference.StackGrowthCurveBezierShapes = append(reference.StackGrowthCurveBezierShapes, stage.StackGrowthCurveBezierShapes_reference[_b])
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

// insertion point for pointer reconstruction from instances
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

func (reference *ExplanationTextShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GridPathShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
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
	if _reference := reference.StartArcShapeGrid; _reference != nil {
		reference.StartArcShapeGrid = nil
		if _instance, ok := stage.StartArcShapeGrids_instance[_reference]; ok {
			reference.StartArcShapeGrid = _instance
		}
	}
	if _reference := reference.StartArcShapeV2Grid; _reference != nil {
		reference.StartArcShapeV2Grid = nil
		if _instance, ok := stage.StartArcShapeV2Grids_instance[_reference]; ok {
			reference.StartArcShapeV2Grid = _instance
		}
	}
	if _reference := reference.EndArcShapeGrid; _reference != nil {
		reference.EndArcShapeGrid = nil
		if _instance, ok := stage.EndArcShapeGrids_instance[_reference]; ok {
			reference.EndArcShapeGrid = _instance
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

// insertion point for diff per struct
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
	if (plant.StartArcShapeGrid == nil) != (plantOther.StartArcShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "StartArcShapeGrid"))
	} else if plant.StartArcShapeGrid != nil && plantOther.StartArcShapeGrid != nil {
		if plant.StartArcShapeGrid != plantOther.StartArcShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "StartArcShapeGrid"))
		}
	}
	if (plant.StartArcShapeV2Grid == nil) != (plantOther.StartArcShapeV2Grid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "StartArcShapeV2Grid"))
	} else if plant.StartArcShapeV2Grid != nil && plantOther.StartArcShapeV2Grid != nil {
		if plant.StartArcShapeV2Grid != plantOther.StartArcShapeV2Grid {
			diffs = append(diffs, plant.GongMarshallField(stage, "StartArcShapeV2Grid"))
		}
	}
	if (plant.EndArcShapeGrid == nil) != (plantOther.EndArcShapeGrid == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "EndArcShapeGrid"))
	} else if plant.EndArcShapeGrid != nil && plantOther.EndArcShapeGrid != nil {
		if plant.EndArcShapeGrid != plantOther.EndArcShapeGrid {
			diffs = append(diffs, plant.GongMarshallField(stage, "EndArcShapeGrid"))
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
	if plantdiagram.IsHiddenStartArcShapeGrid != plantdiagramOther.IsHiddenStartArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenStartArcShapeGrid"))
	}
	if plantdiagram.IsHiddenStartArcShapeV2Grid != plantdiagramOther.IsHiddenStartArcShapeV2Grid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenStartArcShapeV2Grid"))
	}
	if plantdiagram.IsHiddenEndArcShapeGrid != plantdiagramOther.IsHiddenEndArcShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenEndArcShapeGrid"))
	}
	if plantdiagram.IsHiddenGrowthCurveBezierShapeGrid != plantdiagramOther.IsHiddenGrowthCurveBezierShapeGrid {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenGrowthCurveBezierShapeGrid"))
	}
	if plantdiagram.IsHiddenStackOfGrowthCurve != plantdiagramOther.IsHiddenStackOfGrowthCurve {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenStackOfGrowthCurve"))
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
