// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct ArcNormalVectorShape
	// insertion point per field

	// Compute reverse map for named struct ArcNormalVectorShapeGrid
	// insertion point per field
	stage.ArcNormalVectorShapeGrid_ArcNormalVectorShapes_reverseMap = make(map[*ArcNormalVectorShape]*ArcNormalVectorShapeGrid)
	for arcnormalvectorshapegrid := range stage.ArcNormalVectorShapeGrids {
		_ = arcnormalvectorshapegrid
		for _, _arcnormalvectorshape := range arcnormalvectorshapegrid.ArcNormalVectorShapes {
			stage.ArcNormalVectorShapeGrid_ArcNormalVectorShapes_reverseMap[_arcnormalvectorshape] = arcnormalvectorshapegrid
		}
	}

	// Compute reverse map for named struct AxesShape
	// insertion point per field

	// Compute reverse map for named struct BaseVectorShape
	// insertion point per field

	// Compute reverse map for named struct BaseVectorShapeGrid
	// insertion point per field
	stage.BaseVectorShapeGrid_BaseVectorShapes_reverseMap = make(map[*BaseVectorShape]*BaseVectorShapeGrid)
	for basevectorshapegrid := range stage.BaseVectorShapeGrids {
		_ = basevectorshapegrid
		for _, _basevectorshape := range basevectorshapegrid.BaseVectorShapes {
			stage.BaseVectorShapeGrid_BaseVectorShapes_reverseMap[_basevectorshape] = basevectorshapegrid
		}
	}

	// Compute reverse map for named struct BottomEndArcShapeV2
	// insertion point per field

	// Compute reverse map for named struct BottomEndArcShapeV2Grid
	// insertion point per field
	stage.BottomEndArcShapeV2Grid_BottomEndArcShapesV2_reverseMap = make(map[*BottomEndArcShapeV2]*BottomEndArcShapeV2Grid)
	for bottomendarcshapev2grid := range stage.BottomEndArcShapeV2Grids {
		_ = bottomendarcshapev2grid
		for _, _bottomendarcshapev2 := range bottomendarcshapev2grid.BottomEndArcShapesV2 {
			stage.BottomEndArcShapeV2Grid_BottomEndArcShapesV2_reverseMap[_bottomendarcshapev2] = bottomendarcshapev2grid
		}
	}

	// Compute reverse map for named struct BottomStackGrowthCurveEndArcShapeV2
	// insertion point per field

	// Compute reverse map for named struct BottomStackGrowthCurveStartArcShapeV2
	// insertion point per field

	// Compute reverse map for named struct BottomStackOfGrowthCurveV2
	// insertion point per field
	stage.BottomStackOfGrowthCurveV2_BottomStackGrowthCurveStartArcShapeV2s_reverseMap = make(map[*BottomStackGrowthCurveStartArcShapeV2]*BottomStackOfGrowthCurveV2)
	for bottomstackofgrowthcurvev2 := range stage.BottomStackOfGrowthCurveV2s {
		_ = bottomstackofgrowthcurvev2
		for _, _bottomstackgrowthcurvestartarcshapev2 := range bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s {
			stage.BottomStackOfGrowthCurveV2_BottomStackGrowthCurveStartArcShapeV2s_reverseMap[_bottomstackgrowthcurvestartarcshapev2] = bottomstackofgrowthcurvev2
		}
	}
	stage.BottomStackOfGrowthCurveV2_BottomStackGrowthCurveEndArcShapeV2s_reverseMap = make(map[*BottomStackGrowthCurveEndArcShapeV2]*BottomStackOfGrowthCurveV2)
	for bottomstackofgrowthcurvev2 := range stage.BottomStackOfGrowthCurveV2s {
		_ = bottomstackofgrowthcurvev2
		for _, _bottomstackgrowthcurveendarcshapev2 := range bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s {
			stage.BottomStackOfGrowthCurveV2_BottomStackGrowthCurveEndArcShapeV2s_reverseMap[_bottomstackgrowthcurveendarcshapev2] = bottomstackofgrowthcurvev2
		}
	}

	// Compute reverse map for named struct BottomStartArcShapeV2
	// insertion point per field

	// Compute reverse map for named struct BottomStartArcShapeV2Grid
	// insertion point per field
	stage.BottomStartArcShapeV2Grid_BottomStartArcShapesV2_reverseMap = make(map[*BottomStartArcShapeV2]*BottomStartArcShapeV2Grid)
	for bottomstartarcshapev2grid := range stage.BottomStartArcShapeV2Grids {
		_ = bottomstartarcshapev2grid
		for _, _bottomstartarcshapev2 := range bottomstartarcshapev2grid.BottomStartArcShapesV2 {
			stage.BottomStartArcShapeV2Grid_BottomStartArcShapesV2_reverseMap[_bottomstartarcshapev2] = bottomstartarcshapev2grid
		}
	}

	// Compute reverse map for named struct CircleGridShape
	// insertion point per field

	// Compute reverse map for named struct EndArcShape
	// insertion point per field

	// Compute reverse map for named struct EndArcShapeGrid
	// insertion point per field
	stage.EndArcShapeGrid_EndArcShapes_reverseMap = make(map[*EndArcShape]*EndArcShapeGrid)
	for endarcshapegrid := range stage.EndArcShapeGrids {
		_ = endarcshapegrid
		for _, _endarcshape := range endarcshapegrid.EndArcShapes {
			stage.EndArcShapeGrid_EndArcShapes_reverseMap[_endarcshape] = endarcshapegrid
		}
	}

	// Compute reverse map for named struct EndArcShapeV2
	// insertion point per field

	// Compute reverse map for named struct EndArcShapeV2Grid
	// insertion point per field
	stage.EndArcShapeV2Grid_EndArcShapesV2_reverseMap = make(map[*EndArcShapeV2]*EndArcShapeV2Grid)
	for endarcshapev2grid := range stage.EndArcShapeV2Grids {
		_ = endarcshapev2grid
		for _, _endarcshapev2 := range endarcshapev2grid.EndArcShapesV2 {
			stage.EndArcShapeV2Grid_EndArcShapesV2_reverseMap[_endarcshapev2] = endarcshapev2grid
		}
	}

	// Compute reverse map for named struct ExplanationTextShape
	// insertion point per field

	// Compute reverse map for named struct GridPathShape
	// insertion point per field

	// Compute reverse map for named struct GrowthCurveBezierShape
	// insertion point per field

	// Compute reverse map for named struct GrowthCurveBezierShapeGrid
	// insertion point per field
	stage.GrowthCurveBezierShapeGrid_GrowthCurveBezierShapes_reverseMap = make(map[*GrowthCurveBezierShape]*GrowthCurveBezierShapeGrid)
	for growthcurvebeziershapegrid := range stage.GrowthCurveBezierShapeGrids {
		_ = growthcurvebeziershapegrid
		for _, _growthcurvebeziershape := range growthcurvebeziershapegrid.GrowthCurveBezierShapes {
			stage.GrowthCurveBezierShapeGrid_GrowthCurveBezierShapes_reverseMap[_growthcurvebeziershape] = growthcurvebeziershapegrid
		}
	}

	// Compute reverse map for named struct GrowthCurveRhombusGridShape
	// insertion point per field
	stage.GrowthCurveRhombusGridShape_GrowthCurveRhombusShapes_reverseMap = make(map[*GrowthCurveRhombusShape]*GrowthCurveRhombusGridShape)
	for growthcurverhombusgridshape := range stage.GrowthCurveRhombusGridShapes {
		_ = growthcurverhombusgridshape
		for _, _growthcurverhombusshape := range growthcurverhombusgridshape.GrowthCurveRhombusShapes {
			stage.GrowthCurveRhombusGridShape_GrowthCurveRhombusShapes_reverseMap[_growthcurverhombusshape] = growthcurverhombusgridshape
		}
	}

	// Compute reverse map for named struct GrowthCurveRhombusShape
	// insertion point per field

	// Compute reverse map for named struct GrowthVectorShape
	// insertion point per field

	// Compute reverse map for named struct InitialRhombusGridShape
	// insertion point per field
	stage.InitialRhombusGridShape_InitialRhombusShapes_reverseMap = make(map[*InitialRhombusShape]*InitialRhombusGridShape)
	for initialrhombusgridshape := range stage.InitialRhombusGridShapes {
		_ = initialrhombusgridshape
		for _, _initialrhombusshape := range initialrhombusgridshape.InitialRhombusShapes {
			stage.InitialRhombusGridShape_InitialRhombusShapes_reverseMap[_initialrhombusshape] = initialrhombusgridshape
		}
	}

	// Compute reverse map for named struct InitialRhombusShape
	// insertion point per field

	// Compute reverse map for named struct Library
	// insertion point per field
	stage.Library_SubLibraries_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibraries {
			stage.Library_SubLibraries_reverseMap[_library] = library
		}
	}
	stage.Library_Plants_reverseMap = make(map[*Plant]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _plant := range library.Plants {
			stage.Library_Plants_reverseMap[_plant] = library
		}
	}

	// Compute reverse map for named struct NextCircleShape
	// insertion point per field

	// Compute reverse map for named struct PerpendicularVector
	// insertion point per field

	// Compute reverse map for named struct PerpendicularVectorGrid
	// insertion point per field
	stage.PerpendicularVectorGrid_PerpendicularVectors_reverseMap = make(map[*PerpendicularVector]*PerpendicularVectorGrid)
	for perpendicularvectorgrid := range stage.PerpendicularVectorGrids {
		_ = perpendicularvectorgrid
		for _, _perpendicularvector := range perpendicularvectorgrid.PerpendicularVectors {
			stage.PerpendicularVectorGrid_PerpendicularVectors_reverseMap[_perpendicularvector] = perpendicularvectorgrid
		}
	}

	// Compute reverse map for named struct PerpendicularVectorGridHalfway
	// insertion point per field
	stage.PerpendicularVectorGridHalfway_PerpendicularVectorHalfways_reverseMap = make(map[*PerpendicularVectorHalfway]*PerpendicularVectorGridHalfway)
	for perpendicularvectorgridhalfway := range stage.PerpendicularVectorGridHalfways {
		_ = perpendicularvectorgridhalfway
		for _, _perpendicularvectorhalfway := range perpendicularvectorgridhalfway.PerpendicularVectorHalfways {
			stage.PerpendicularVectorGridHalfway_PerpendicularVectorHalfways_reverseMap[_perpendicularvectorhalfway] = perpendicularvectorgridhalfway
		}
	}

	// Compute reverse map for named struct PerpendicularVectorHalfway
	// insertion point per field

	// Compute reverse map for named struct Plant
	// insertion point per field
	stage.Plant_PlantDiagrams_reverseMap = make(map[*PlantDiagram]*Plant)
	for plant := range stage.Plants {
		_ = plant
		for _, _plantdiagram := range plant.PlantDiagrams {
			stage.Plant_PlantDiagrams_reverseMap[_plantdiagram] = plant
		}
	}

	// Compute reverse map for named struct PlantCircumferenceShape
	// insertion point per field

	// Compute reverse map for named struct PlantDiagram
	// insertion point per field

	// Compute reverse map for named struct Rendered3DShape
	// insertion point per field

	// Compute reverse map for named struct RhombusShape
	// insertion point per field

	// Compute reverse map for named struct RotatedRhombusGridShape
	// insertion point per field
	stage.RotatedRhombusGridShape_RotatedRhombusShapes_reverseMap = make(map[*RotatedRhombusShape]*RotatedRhombusGridShape)
	for rotatedrhombusgridshape := range stage.RotatedRhombusGridShapes {
		_ = rotatedrhombusgridshape
		for _, _rotatedrhombusshape := range rotatedrhombusgridshape.RotatedRhombusShapes {
			stage.RotatedRhombusGridShape_RotatedRhombusShapes_reverseMap[_rotatedrhombusshape] = rotatedrhombusgridshape
		}
	}

	// Compute reverse map for named struct RotatedRhombusShape
	// insertion point per field

	// Compute reverse map for named struct StackGrowthCurveBezierShape
	// insertion point per field

	// Compute reverse map for named struct StackGrowthCurveEndArcShapeV2
	// insertion point per field

	// Compute reverse map for named struct StackGrowthCurveStartArcShapeV2
	// insertion point per field

	// Compute reverse map for named struct StackOfGrowthCurve
	// insertion point per field
	stage.StackOfGrowthCurve_StackGrowthCurveBezierShapes_reverseMap = make(map[*StackGrowthCurveBezierShape]*StackOfGrowthCurve)
	for stackofgrowthcurve := range stage.StackOfGrowthCurves {
		_ = stackofgrowthcurve
		for _, _stackgrowthcurvebeziershape := range stackofgrowthcurve.StackGrowthCurveBezierShapes {
			stage.StackOfGrowthCurve_StackGrowthCurveBezierShapes_reverseMap[_stackgrowthcurvebeziershape] = stackofgrowthcurve
		}
	}

	// Compute reverse map for named struct StackOfGrowthCurveV2
	// insertion point per field
	stage.StackOfGrowthCurveV2_StackGrowthCurveStartArcShapeV2s_reverseMap = make(map[*StackGrowthCurveStartArcShapeV2]*StackOfGrowthCurveV2)
	for stackofgrowthcurvev2 := range stage.StackOfGrowthCurveV2s {
		_ = stackofgrowthcurvev2
		for _, _stackgrowthcurvestartarcshapev2 := range stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s {
			stage.StackOfGrowthCurveV2_StackGrowthCurveStartArcShapeV2s_reverseMap[_stackgrowthcurvestartarcshapev2] = stackofgrowthcurvev2
		}
	}
	stage.StackOfGrowthCurveV2_StackGrowthCurveEndArcShapeV2s_reverseMap = make(map[*StackGrowthCurveEndArcShapeV2]*StackOfGrowthCurveV2)
	for stackofgrowthcurvev2 := range stage.StackOfGrowthCurveV2s {
		_ = stackofgrowthcurvev2
		for _, _stackgrowthcurveendarcshapev2 := range stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s {
			stage.StackOfGrowthCurveV2_StackGrowthCurveEndArcShapeV2s_reverseMap[_stackgrowthcurveendarcshapev2] = stackofgrowthcurvev2
		}
	}

	// Compute reverse map for named struct StartArcShape
	// insertion point per field

	// Compute reverse map for named struct StartArcShapeGrid
	// insertion point per field
	stage.StartArcShapeGrid_StartArcShapes_reverseMap = make(map[*StartArcShape]*StartArcShapeGrid)
	for startarcshapegrid := range stage.StartArcShapeGrids {
		_ = startarcshapegrid
		for _, _startarcshape := range startarcshapegrid.StartArcShapes {
			stage.StartArcShapeGrid_StartArcShapes_reverseMap[_startarcshape] = startarcshapegrid
		}
	}

	// Compute reverse map for named struct StartArcShapeV2
	// insertion point per field

	// Compute reverse map for named struct StartArcShapeV2Grid
	// insertion point per field
	stage.StartArcShapeV2Grid_StartArcShapesV2_reverseMap = make(map[*StartArcShapeV2]*StartArcShapeV2Grid)
	for startarcshapev2grid := range stage.StartArcShapeV2Grids {
		_ = startarcshapev2grid
		for _, _startarcshapev2 := range startarcshapev2grid.StartArcShapesV2 {
			stage.StartArcShapeV2Grid_StartArcShapesV2_reverseMap[_startarcshapev2] = startarcshapev2grid
		}
	}

	// Compute reverse map for named struct TopEndArcShapeV2
	// insertion point per field

	// Compute reverse map for named struct TopEndArcShapeV2Grid
	// insertion point per field
	stage.TopEndArcShapeV2Grid_TopEndArcShapesV2_reverseMap = make(map[*TopEndArcShapeV2]*TopEndArcShapeV2Grid)
	for topendarcshapev2grid := range stage.TopEndArcShapeV2Grids {
		_ = topendarcshapev2grid
		for _, _topendarcshapev2 := range topendarcshapev2grid.TopEndArcShapesV2 {
			stage.TopEndArcShapeV2Grid_TopEndArcShapesV2_reverseMap[_topendarcshapev2] = topendarcshapev2grid
		}
	}

	// Compute reverse map for named struct TopStackGrowthCurveEndArcShapeV2
	// insertion point per field

	// Compute reverse map for named struct TopStackGrowthCurveStartArcShapeV2
	// insertion point per field

	// Compute reverse map for named struct TopStackOfGrowthCurveV2
	// insertion point per field
	stage.TopStackOfGrowthCurveV2_TopStackGrowthCurveStartArcShapeV2s_reverseMap = make(map[*TopStackGrowthCurveStartArcShapeV2]*TopStackOfGrowthCurveV2)
	for topstackofgrowthcurvev2 := range stage.TopStackOfGrowthCurveV2s {
		_ = topstackofgrowthcurvev2
		for _, _topstackgrowthcurvestartarcshapev2 := range topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s {
			stage.TopStackOfGrowthCurveV2_TopStackGrowthCurveStartArcShapeV2s_reverseMap[_topstackgrowthcurvestartarcshapev2] = topstackofgrowthcurvev2
		}
	}
	stage.TopStackOfGrowthCurveV2_TopStackGrowthCurveEndArcShapeV2s_reverseMap = make(map[*TopStackGrowthCurveEndArcShapeV2]*TopStackOfGrowthCurveV2)
	for topstackofgrowthcurvev2 := range stage.TopStackOfGrowthCurveV2s {
		_ = topstackofgrowthcurvev2
		for _, _topstackgrowthcurveendarcshapev2 := range topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s {
			stage.TopStackOfGrowthCurveV2_TopStackGrowthCurveEndArcShapeV2s_reverseMap[_topstackgrowthcurveendarcshapev2] = topstackofgrowthcurvev2
		}
	}

	// Compute reverse map for named struct TopStartArcShapeV2
	// insertion point per field

	// Compute reverse map for named struct TopStartArcShapeV2Grid
	// insertion point per field
	stage.TopStartArcShapeV2Grid_TopStartArcShapesV2_reverseMap = make(map[*TopStartArcShapeV2]*TopStartArcShapeV2Grid)
	for topstartarcshapev2grid := range stage.TopStartArcShapeV2Grids {
		_ = topstartarcshapev2grid
		for _, _topstartarcshapev2 := range topstartarcshapev2grid.TopStartArcShapesV2 {
			stage.TopStartArcShapeV2Grid_TopStartArcShapesV2_reverseMap[_topstartarcshapev2] = topstartarcshapev2grid
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.ArcNormalVectorShapes {
		res = append(res, instance)
	}

	for instance := range stage.ArcNormalVectorShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.AxesShapes {
		res = append(res, instance)
	}

	for instance := range stage.BaseVectorShapes {
		res = append(res, instance)
	}

	for instance := range stage.BaseVectorShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.BottomEndArcShapeV2s {
		res = append(res, instance)
	}

	for instance := range stage.BottomEndArcShapeV2Grids {
		res = append(res, instance)
	}

	for instance := range stage.BottomStackGrowthCurveEndArcShapeV2s {
		res = append(res, instance)
	}

	for instance := range stage.BottomStackGrowthCurveStartArcShapeV2s {
		res = append(res, instance)
	}

	for instance := range stage.BottomStackOfGrowthCurveV2s {
		res = append(res, instance)
	}

	for instance := range stage.BottomStartArcShapeV2s {
		res = append(res, instance)
	}

	for instance := range stage.BottomStartArcShapeV2Grids {
		res = append(res, instance)
	}

	for instance := range stage.CircleGridShapes {
		res = append(res, instance)
	}

	for instance := range stage.EndArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.EndArcShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.EndArcShapeV2s {
		res = append(res, instance)
	}

	for instance := range stage.EndArcShapeV2Grids {
		res = append(res, instance)
	}

	for instance := range stage.ExplanationTextShapes {
		res = append(res, instance)
	}

	for instance := range stage.GridPathShapes {
		res = append(res, instance)
	}

	for instance := range stage.GrowthCurveBezierShapes {
		res = append(res, instance)
	}

	for instance := range stage.GrowthCurveBezierShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.GrowthCurveRhombusGridShapes {
		res = append(res, instance)
	}

	for instance := range stage.GrowthCurveRhombusShapes {
		res = append(res, instance)
	}

	for instance := range stage.GrowthVectorShapes {
		res = append(res, instance)
	}

	for instance := range stage.InitialRhombusGridShapes {
		res = append(res, instance)
	}

	for instance := range stage.InitialRhombusShapes {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.NextCircleShapes {
		res = append(res, instance)
	}

	for instance := range stage.PerpendicularVectors {
		res = append(res, instance)
	}

	for instance := range stage.PerpendicularVectorGrids {
		res = append(res, instance)
	}

	for instance := range stage.PerpendicularVectorGridHalfways {
		res = append(res, instance)
	}

	for instance := range stage.PerpendicularVectorHalfways {
		res = append(res, instance)
	}

	for instance := range stage.Plants {
		res = append(res, instance)
	}

	for instance := range stage.PlantCircumferenceShapes {
		res = append(res, instance)
	}

	for instance := range stage.PlantDiagrams {
		res = append(res, instance)
	}

	for instance := range stage.Rendered3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.RhombusShapes {
		res = append(res, instance)
	}

	for instance := range stage.RotatedRhombusGridShapes {
		res = append(res, instance)
	}

	for instance := range stage.RotatedRhombusShapes {
		res = append(res, instance)
	}

	for instance := range stage.StackGrowthCurveBezierShapes {
		res = append(res, instance)
	}

	for instance := range stage.StackGrowthCurveEndArcShapeV2s {
		res = append(res, instance)
	}

	for instance := range stage.StackGrowthCurveStartArcShapeV2s {
		res = append(res, instance)
	}

	for instance := range stage.StackOfGrowthCurves {
		res = append(res, instance)
	}

	for instance := range stage.StackOfGrowthCurveV2s {
		res = append(res, instance)
	}

	for instance := range stage.StartArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.StartArcShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.StartArcShapeV2s {
		res = append(res, instance)
	}

	for instance := range stage.StartArcShapeV2Grids {
		res = append(res, instance)
	}

	for instance := range stage.TopEndArcShapeV2s {
		res = append(res, instance)
	}

	for instance := range stage.TopEndArcShapeV2Grids {
		res = append(res, instance)
	}

	for instance := range stage.TopStackGrowthCurveEndArcShapeV2s {
		res = append(res, instance)
	}

	for instance := range stage.TopStackGrowthCurveStartArcShapeV2s {
		res = append(res, instance)
	}

	for instance := range stage.TopStackOfGrowthCurveV2s {
		res = append(res, instance)
	}

	for instance := range stage.TopStartArcShapeV2s {
		res = append(res, instance)
	}

	for instance := range stage.TopStartArcShapeV2Grids {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (arcnormalvectorshape *ArcNormalVectorShape) GongCopy() GongstructIF {
	newInstance := new(ArcNormalVectorShape)
	arcnormalvectorshape.CopyBasicFields(newInstance)
	return newInstance
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongCopy() GongstructIF {
	newInstance := new(ArcNormalVectorShapeGrid)
	arcnormalvectorshapegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (axesshape *AxesShape) GongCopy() GongstructIF {
	newInstance := new(AxesShape)
	axesshape.CopyBasicFields(newInstance)
	return newInstance
}

func (basevectorshape *BaseVectorShape) GongCopy() GongstructIF {
	newInstance := new(BaseVectorShape)
	basevectorshape.CopyBasicFields(newInstance)
	return newInstance
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongCopy() GongstructIF {
	newInstance := new(BaseVectorShapeGrid)
	basevectorshapegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (bottomendarcshapev2 *BottomEndArcShapeV2) GongCopy() GongstructIF {
	newInstance := new(BottomEndArcShapeV2)
	bottomendarcshapev2.CopyBasicFields(newInstance)
	return newInstance
}

func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) GongCopy() GongstructIF {
	newInstance := new(BottomEndArcShapeV2Grid)
	bottomendarcshapev2grid.CopyBasicFields(newInstance)
	return newInstance
}

func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) GongCopy() GongstructIF {
	newInstance := new(BottomStackGrowthCurveEndArcShapeV2)
	bottomstackgrowthcurveendarcshapev2.CopyBasicFields(newInstance)
	return newInstance
}

func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) GongCopy() GongstructIF {
	newInstance := new(BottomStackGrowthCurveStartArcShapeV2)
	bottomstackgrowthcurvestartarcshapev2.CopyBasicFields(newInstance)
	return newInstance
}

func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) GongCopy() GongstructIF {
	newInstance := new(BottomStackOfGrowthCurveV2)
	bottomstackofgrowthcurvev2.CopyBasicFields(newInstance)
	return newInstance
}

func (bottomstartarcshapev2 *BottomStartArcShapeV2) GongCopy() GongstructIF {
	newInstance := new(BottomStartArcShapeV2)
	bottomstartarcshapev2.CopyBasicFields(newInstance)
	return newInstance
}

func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) GongCopy() GongstructIF {
	newInstance := new(BottomStartArcShapeV2Grid)
	bottomstartarcshapev2grid.CopyBasicFields(newInstance)
	return newInstance
}

func (circlegridshape *CircleGridShape) GongCopy() GongstructIF {
	newInstance := new(CircleGridShape)
	circlegridshape.CopyBasicFields(newInstance)
	return newInstance
}

func (endarcshape *EndArcShape) GongCopy() GongstructIF {
	newInstance := new(EndArcShape)
	endarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (endarcshapegrid *EndArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(EndArcShapeGrid)
	endarcshapegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (endarcshapev2 *EndArcShapeV2) GongCopy() GongstructIF {
	newInstance := new(EndArcShapeV2)
	endarcshapev2.CopyBasicFields(newInstance)
	return newInstance
}

func (endarcshapev2grid *EndArcShapeV2Grid) GongCopy() GongstructIF {
	newInstance := new(EndArcShapeV2Grid)
	endarcshapev2grid.CopyBasicFields(newInstance)
	return newInstance
}

func (explanationtextshape *ExplanationTextShape) GongCopy() GongstructIF {
	newInstance := new(ExplanationTextShape)
	explanationtextshape.CopyBasicFields(newInstance)
	return newInstance
}

func (gridpathshape *GridPathShape) GongCopy() GongstructIF {
	newInstance := new(GridPathShape)
	gridpathshape.CopyBasicFields(newInstance)
	return newInstance
}

func (growthcurvebeziershape *GrowthCurveBezierShape) GongCopy() GongstructIF {
	newInstance := new(GrowthCurveBezierShape)
	growthcurvebeziershape.CopyBasicFields(newInstance)
	return newInstance
}

func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GongCopy() GongstructIF {
	newInstance := new(GrowthCurveBezierShapeGrid)
	growthcurvebeziershapegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongCopy() GongstructIF {
	newInstance := new(GrowthCurveRhombusGridShape)
	growthcurverhombusgridshape.CopyBasicFields(newInstance)
	return newInstance
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongCopy() GongstructIF {
	newInstance := new(GrowthCurveRhombusShape)
	growthcurverhombusshape.CopyBasicFields(newInstance)
	return newInstance
}

func (growthvectorshape *GrowthVectorShape) GongCopy() GongstructIF {
	newInstance := new(GrowthVectorShape)
	growthvectorshape.CopyBasicFields(newInstance)
	return newInstance
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongCopy() GongstructIF {
	newInstance := new(InitialRhombusGridShape)
	initialrhombusgridshape.CopyBasicFields(newInstance)
	return newInstance
}

func (initialrhombusshape *InitialRhombusShape) GongCopy() GongstructIF {
	newInstance := new(InitialRhombusShape)
	initialrhombusshape.CopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.CopyBasicFields(newInstance)
	return newInstance
}

func (nextcircleshape *NextCircleShape) GongCopy() GongstructIF {
	newInstance := new(NextCircleShape)
	nextcircleshape.CopyBasicFields(newInstance)
	return newInstance
}

func (perpendicularvector *PerpendicularVector) GongCopy() GongstructIF {
	newInstance := new(PerpendicularVector)
	perpendicularvector.CopyBasicFields(newInstance)
	return newInstance
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongCopy() GongstructIF {
	newInstance := new(PerpendicularVectorGrid)
	perpendicularvectorgrid.CopyBasicFields(newInstance)
	return newInstance
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongCopy() GongstructIF {
	newInstance := new(PerpendicularVectorGridHalfway)
	perpendicularvectorgridhalfway.CopyBasicFields(newInstance)
	return newInstance
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongCopy() GongstructIF {
	newInstance := new(PerpendicularVectorHalfway)
	perpendicularvectorhalfway.CopyBasicFields(newInstance)
	return newInstance
}

func (plant *Plant) GongCopy() GongstructIF {
	newInstance := new(Plant)
	plant.CopyBasicFields(newInstance)
	return newInstance
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongCopy() GongstructIF {
	newInstance := new(PlantCircumferenceShape)
	plantcircumferenceshape.CopyBasicFields(newInstance)
	return newInstance
}

func (plantdiagram *PlantDiagram) GongCopy() GongstructIF {
	newInstance := new(PlantDiagram)
	plantdiagram.CopyBasicFields(newInstance)
	return newInstance
}

func (rendered3dshape *Rendered3DShape) GongCopy() GongstructIF {
	newInstance := new(Rendered3DShape)
	rendered3dshape.CopyBasicFields(newInstance)
	return newInstance
}

func (rhombusshape *RhombusShape) GongCopy() GongstructIF {
	newInstance := new(RhombusShape)
	rhombusshape.CopyBasicFields(newInstance)
	return newInstance
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongCopy() GongstructIF {
	newInstance := new(RotatedRhombusGridShape)
	rotatedrhombusgridshape.CopyBasicFields(newInstance)
	return newInstance
}

func (rotatedrhombusshape *RotatedRhombusShape) GongCopy() GongstructIF {
	newInstance := new(RotatedRhombusShape)
	rotatedrhombusshape.CopyBasicFields(newInstance)
	return newInstance
}

func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GongCopy() GongstructIF {
	newInstance := new(StackGrowthCurveBezierShape)
	stackgrowthcurvebeziershape.CopyBasicFields(newInstance)
	return newInstance
}

func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) GongCopy() GongstructIF {
	newInstance := new(StackGrowthCurveEndArcShapeV2)
	stackgrowthcurveendarcshapev2.CopyBasicFields(newInstance)
	return newInstance
}

func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) GongCopy() GongstructIF {
	newInstance := new(StackGrowthCurveStartArcShapeV2)
	stackgrowthcurvestartarcshapev2.CopyBasicFields(newInstance)
	return newInstance
}

func (stackofgrowthcurve *StackOfGrowthCurve) GongCopy() GongstructIF {
	newInstance := new(StackOfGrowthCurve)
	stackofgrowthcurve.CopyBasicFields(newInstance)
	return newInstance
}

func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) GongCopy() GongstructIF {
	newInstance := new(StackOfGrowthCurveV2)
	stackofgrowthcurvev2.CopyBasicFields(newInstance)
	return newInstance
}

func (startarcshape *StartArcShape) GongCopy() GongstructIF {
	newInstance := new(StartArcShape)
	startarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (startarcshapegrid *StartArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(StartArcShapeGrid)
	startarcshapegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (startarcshapev2 *StartArcShapeV2) GongCopy() GongstructIF {
	newInstance := new(StartArcShapeV2)
	startarcshapev2.CopyBasicFields(newInstance)
	return newInstance
}

func (startarcshapev2grid *StartArcShapeV2Grid) GongCopy() GongstructIF {
	newInstance := new(StartArcShapeV2Grid)
	startarcshapev2grid.CopyBasicFields(newInstance)
	return newInstance
}

func (topendarcshapev2 *TopEndArcShapeV2) GongCopy() GongstructIF {
	newInstance := new(TopEndArcShapeV2)
	topendarcshapev2.CopyBasicFields(newInstance)
	return newInstance
}

func (topendarcshapev2grid *TopEndArcShapeV2Grid) GongCopy() GongstructIF {
	newInstance := new(TopEndArcShapeV2Grid)
	topendarcshapev2grid.CopyBasicFields(newInstance)
	return newInstance
}

func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) GongCopy() GongstructIF {
	newInstance := new(TopStackGrowthCurveEndArcShapeV2)
	topstackgrowthcurveendarcshapev2.CopyBasicFields(newInstance)
	return newInstance
}

func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) GongCopy() GongstructIF {
	newInstance := new(TopStackGrowthCurveStartArcShapeV2)
	topstackgrowthcurvestartarcshapev2.CopyBasicFields(newInstance)
	return newInstance
}

func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) GongCopy() GongstructIF {
	newInstance := new(TopStackOfGrowthCurveV2)
	topstackofgrowthcurvev2.CopyBasicFields(newInstance)
	return newInstance
}

func (topstartarcshapev2 *TopStartArcShapeV2) GongCopy() GongstructIF {
	newInstance := new(TopStartArcShapeV2)
	topstartarcshapev2.CopyBasicFields(newInstance)
	return newInstance
}

func (topstartarcshapev2grid *TopStartArcShapeV2Grid) GongCopy() GongstructIF {
	newInstance := new(TopStartArcShapeV2Grid)
	topstartarcshapev2grid.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (arcnormalvectorshape *ArcNormalVectorShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(arcnormalvectorshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(arcnormalvectorshape), uint64(GetOrderPointerGongstruct(stage, arcnormalvectorshape)))
	return
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(arcnormalvectorshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(arcnormalvectorshapegrid), uint64(GetOrderPointerGongstruct(stage, arcnormalvectorshapegrid)))
	return
}

func (axesshape *AxesShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(axesshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(axesshape), uint64(GetOrderPointerGongstruct(stage, axesshape)))
	return
}

func (basevectorshape *BaseVectorShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(basevectorshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(basevectorshape), uint64(GetOrderPointerGongstruct(stage, basevectorshape)))
	return
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(basevectorshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(basevectorshapegrid), uint64(GetOrderPointerGongstruct(stage, basevectorshapegrid)))
	return
}

func (bottomendarcshapev2 *BottomEndArcShapeV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bottomendarcshapev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(bottomendarcshapev2), uint64(GetOrderPointerGongstruct(stage, bottomendarcshapev2)))
	return
}

func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bottomendarcshapev2grid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(bottomendarcshapev2grid), uint64(GetOrderPointerGongstruct(stage, bottomendarcshapev2grid)))
	return
}

func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bottomstackgrowthcurveendarcshapev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(bottomstackgrowthcurveendarcshapev2), uint64(GetOrderPointerGongstruct(stage, bottomstackgrowthcurveendarcshapev2)))
	return
}

func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bottomstackgrowthcurvestartarcshapev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(bottomstackgrowthcurvestartarcshapev2), uint64(GetOrderPointerGongstruct(stage, bottomstackgrowthcurvestartarcshapev2)))
	return
}

func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bottomstackofgrowthcurvev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(bottomstackofgrowthcurvev2), uint64(GetOrderPointerGongstruct(stage, bottomstackofgrowthcurvev2)))
	return
}

func (bottomstartarcshapev2 *BottomStartArcShapeV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bottomstartarcshapev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(bottomstartarcshapev2), uint64(GetOrderPointerGongstruct(stage, bottomstartarcshapev2)))
	return
}

func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bottomstartarcshapev2grid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(bottomstartarcshapev2grid), uint64(GetOrderPointerGongstruct(stage, bottomstartarcshapev2grid)))
	return
}

func (circlegridshape *CircleGridShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(circlegridshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(circlegridshape), uint64(GetOrderPointerGongstruct(stage, circlegridshape)))
	return
}

func (endarcshape *EndArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(endarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(endarcshape), uint64(GetOrderPointerGongstruct(stage, endarcshape)))
	return
}

func (endarcshapegrid *EndArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(endarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(endarcshapegrid), uint64(GetOrderPointerGongstruct(stage, endarcshapegrid)))
	return
}

func (endarcshapev2 *EndArcShapeV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(endarcshapev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(endarcshapev2), uint64(GetOrderPointerGongstruct(stage, endarcshapev2)))
	return
}

func (endarcshapev2grid *EndArcShapeV2Grid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(endarcshapev2grid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(endarcshapev2grid), uint64(GetOrderPointerGongstruct(stage, endarcshapev2grid)))
	return
}

func (explanationtextshape *ExplanationTextShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(explanationtextshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(explanationtextshape), uint64(GetOrderPointerGongstruct(stage, explanationtextshape)))
	return
}

func (gridpathshape *GridPathShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(gridpathshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(gridpathshape), uint64(GetOrderPointerGongstruct(stage, gridpathshape)))
	return
}

func (growthcurvebeziershape *GrowthCurveBezierShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(growthcurvebeziershape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(growthcurvebeziershape), uint64(GetOrderPointerGongstruct(stage, growthcurvebeziershape)))
	return
}

func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(growthcurvebeziershapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(growthcurvebeziershapegrid), uint64(GetOrderPointerGongstruct(stage, growthcurvebeziershapegrid)))
	return
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(growthcurverhombusgridshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(growthcurverhombusgridshape), uint64(GetOrderPointerGongstruct(stage, growthcurverhombusgridshape)))
	return
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(growthcurverhombusshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(growthcurverhombusshape), uint64(GetOrderPointerGongstruct(stage, growthcurverhombusshape)))
	return
}

func (growthvectorshape *GrowthVectorShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(growthvectorshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(growthvectorshape), uint64(GetOrderPointerGongstruct(stage, growthvectorshape)))
	return
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(initialrhombusgridshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(initialrhombusgridshape), uint64(GetOrderPointerGongstruct(stage, initialrhombusgridshape)))
	return
}

func (initialrhombusshape *InitialRhombusShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(initialrhombusshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(initialrhombusshape), uint64(GetOrderPointerGongstruct(stage, initialrhombusshape)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(library), uint64(GetOrderPointerGongstruct(stage, library)))
	return
}

func (nextcircleshape *NextCircleShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(nextcircleshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(nextcircleshape), uint64(GetOrderPointerGongstruct(stage, nextcircleshape)))
	return
}

func (perpendicularvector *PerpendicularVector) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(perpendicularvector).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(perpendicularvector), uint64(GetOrderPointerGongstruct(stage, perpendicularvector)))
	return
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(perpendicularvectorgrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(perpendicularvectorgrid), uint64(GetOrderPointerGongstruct(stage, perpendicularvectorgrid)))
	return
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(perpendicularvectorgridhalfway).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(perpendicularvectorgridhalfway), uint64(GetOrderPointerGongstruct(stage, perpendicularvectorgridhalfway)))
	return
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(perpendicularvectorhalfway).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(perpendicularvectorhalfway), uint64(GetOrderPointerGongstruct(stage, perpendicularvectorhalfway)))
	return
}

func (plant *Plant) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(plant).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(plant), uint64(GetOrderPointerGongstruct(stage, plant)))
	return
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(plantcircumferenceshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(plantcircumferenceshape), uint64(GetOrderPointerGongstruct(stage, plantcircumferenceshape)))
	return
}

func (plantdiagram *PlantDiagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(plantdiagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(plantdiagram), uint64(GetOrderPointerGongstruct(stage, plantdiagram)))
	return
}

func (rendered3dshape *Rendered3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rendered3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(rendered3dshape), uint64(GetOrderPointerGongstruct(stage, rendered3dshape)))
	return
}

func (rhombusshape *RhombusShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rhombusshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(rhombusshape), uint64(GetOrderPointerGongstruct(stage, rhombusshape)))
	return
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rotatedrhombusgridshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(rotatedrhombusgridshape), uint64(GetOrderPointerGongstruct(stage, rotatedrhombusgridshape)))
	return
}

func (rotatedrhombusshape *RotatedRhombusShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rotatedrhombusshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(rotatedrhombusshape), uint64(GetOrderPointerGongstruct(stage, rotatedrhombusshape)))
	return
}

func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackgrowthcurvebeziershape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackgrowthcurvebeziershape), uint64(GetOrderPointerGongstruct(stage, stackgrowthcurvebeziershape)))
	return
}

func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackgrowthcurveendarcshapev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackgrowthcurveendarcshapev2), uint64(GetOrderPointerGongstruct(stage, stackgrowthcurveendarcshapev2)))
	return
}

func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackgrowthcurvestartarcshapev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackgrowthcurvestartarcshapev2), uint64(GetOrderPointerGongstruct(stage, stackgrowthcurvestartarcshapev2)))
	return
}

func (stackofgrowthcurve *StackOfGrowthCurve) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofgrowthcurve).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackofgrowthcurve), uint64(GetOrderPointerGongstruct(stage, stackofgrowthcurve)))
	return
}

func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofgrowthcurvev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackofgrowthcurvev2), uint64(GetOrderPointerGongstruct(stage, stackofgrowthcurvev2)))
	return
}

func (startarcshape *StartArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(startarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(startarcshape), uint64(GetOrderPointerGongstruct(stage, startarcshape)))
	return
}

func (startarcshapegrid *StartArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(startarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(startarcshapegrid), uint64(GetOrderPointerGongstruct(stage, startarcshapegrid)))
	return
}

func (startarcshapev2 *StartArcShapeV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(startarcshapev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(startarcshapev2), uint64(GetOrderPointerGongstruct(stage, startarcshapev2)))
	return
}

func (startarcshapev2grid *StartArcShapeV2Grid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(startarcshapev2grid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(startarcshapev2grid), uint64(GetOrderPointerGongstruct(stage, startarcshapev2grid)))
	return
}

func (topendarcshapev2 *TopEndArcShapeV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topendarcshapev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topendarcshapev2), uint64(GetOrderPointerGongstruct(stage, topendarcshapev2)))
	return
}

func (topendarcshapev2grid *TopEndArcShapeV2Grid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topendarcshapev2grid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topendarcshapev2grid), uint64(GetOrderPointerGongstruct(stage, topendarcshapev2grid)))
	return
}

func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackgrowthcurveendarcshapev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstackgrowthcurveendarcshapev2), uint64(GetOrderPointerGongstruct(stage, topstackgrowthcurveendarcshapev2)))
	return
}

func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackgrowthcurvestartarcshapev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstackgrowthcurvestartarcshapev2), uint64(GetOrderPointerGongstruct(stage, topstackgrowthcurvestartarcshapev2)))
	return
}

func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackofgrowthcurvev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstackofgrowthcurvev2), uint64(GetOrderPointerGongstruct(stage, topstackofgrowthcurvev2)))
	return
}

func (topstartarcshapev2 *TopStartArcShapeV2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstartarcshapev2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstartarcshapev2), uint64(GetOrderPointerGongstruct(stage, topstartarcshapev2)))
	return
}

func (topstartarcshapev2grid *TopStartArcShapeV2Grid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstartarcshapev2grid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstartarcshapev2grid), uint64(GetOrderPointerGongstruct(stage, topstartarcshapev2grid)))
	return
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var arcnormalvectorshapes_newInstances []*ArcNormalVectorShape
	var arcnormalvectorshapes_deletedInstances []*ArcNormalVectorShape

	// parse all staged instances and check if they have a reference
	for arcnormalvectorshape := range stage.ArcNormalVectorShapes {
		if ref, ok := stage.ArcNormalVectorShapes_reference[arcnormalvectorshape]; !ok {
			arcnormalvectorshapes_newInstances = append(arcnormalvectorshapes_newInstances, arcnormalvectorshape)
			newInstancesSlice = append(newInstancesSlice, arcnormalvectorshape.GongMarshallIdentifier(stage))
			if stage.ArcNormalVectorShapes_referenceOrder == nil {
				stage.ArcNormalVectorShapes_referenceOrder = make(map[*ArcNormalVectorShape]uint)
			}
			stage.ArcNormalVectorShapes_referenceOrder[arcnormalvectorshape] = stage.ArcNormalVectorShape_stagedOrder[arcnormalvectorshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, arcnormalvectorshape.GongMarshallUnstaging(stage))
			// delete(stage.ArcNormalVectorShapes_referenceOrder, arcnormalvectorshape)
			fieldInitializers, pointersInitializations := arcnormalvectorshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ArcNormalVectorShape_stagedOrder[ref] = stage.ArcNormalVectorShape_stagedOrder[arcnormalvectorshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := arcnormalvectorshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, arcnormalvectorshape)
			// delete(stage.ArcNormalVectorShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if arcnormalvectorshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", arcnormalvectorshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ArcNormalVectorShapes_reference {
		instance := stage.ArcNormalVectorShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ArcNormalVectorShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			arcnormalvectorshapes_deletedInstances = append(arcnormalvectorshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(arcnormalvectorshapes_newInstances)
	lenDeletedInstances += len(arcnormalvectorshapes_deletedInstances)
	var arcnormalvectorshapegrids_newInstances []*ArcNormalVectorShapeGrid
	var arcnormalvectorshapegrids_deletedInstances []*ArcNormalVectorShapeGrid

	// parse all staged instances and check if they have a reference
	for arcnormalvectorshapegrid := range stage.ArcNormalVectorShapeGrids {
		if ref, ok := stage.ArcNormalVectorShapeGrids_reference[arcnormalvectorshapegrid]; !ok {
			arcnormalvectorshapegrids_newInstances = append(arcnormalvectorshapegrids_newInstances, arcnormalvectorshapegrid)
			newInstancesSlice = append(newInstancesSlice, arcnormalvectorshapegrid.GongMarshallIdentifier(stage))
			if stage.ArcNormalVectorShapeGrids_referenceOrder == nil {
				stage.ArcNormalVectorShapeGrids_referenceOrder = make(map[*ArcNormalVectorShapeGrid]uint)
			}
			stage.ArcNormalVectorShapeGrids_referenceOrder[arcnormalvectorshapegrid] = stage.ArcNormalVectorShapeGrid_stagedOrder[arcnormalvectorshapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, arcnormalvectorshapegrid.GongMarshallUnstaging(stage))
			// delete(stage.ArcNormalVectorShapeGrids_referenceOrder, arcnormalvectorshapegrid)
			fieldInitializers, pointersInitializations := arcnormalvectorshapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ArcNormalVectorShapeGrid_stagedOrder[ref] = stage.ArcNormalVectorShapeGrid_stagedOrder[arcnormalvectorshapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := arcnormalvectorshapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, arcnormalvectorshapegrid)
			// delete(stage.ArcNormalVectorShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if arcnormalvectorshapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", arcnormalvectorshapegrid.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ArcNormalVectorShapeGrids_reference {
		instance := stage.ArcNormalVectorShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ArcNormalVectorShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			arcnormalvectorshapegrids_deletedInstances = append(arcnormalvectorshapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(arcnormalvectorshapegrids_newInstances)
	lenDeletedInstances += len(arcnormalvectorshapegrids_deletedInstances)
	var axesshapes_newInstances []*AxesShape
	var axesshapes_deletedInstances []*AxesShape

	// parse all staged instances and check if they have a reference
	for axesshape := range stage.AxesShapes {
		if ref, ok := stage.AxesShapes_reference[axesshape]; !ok {
			axesshapes_newInstances = append(axesshapes_newInstances, axesshape)
			newInstancesSlice = append(newInstancesSlice, axesshape.GongMarshallIdentifier(stage))
			if stage.AxesShapes_referenceOrder == nil {
				stage.AxesShapes_referenceOrder = make(map[*AxesShape]uint)
			}
			stage.AxesShapes_referenceOrder[axesshape] = stage.AxesShape_stagedOrder[axesshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, axesshape.GongMarshallUnstaging(stage))
			// delete(stage.AxesShapes_referenceOrder, axesshape)
			fieldInitializers, pointersInitializations := axesshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AxesShape_stagedOrder[ref] = stage.AxesShape_stagedOrder[axesshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := axesshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, axesshape)
			// delete(stage.AxesShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if axesshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", axesshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.AxesShapes_reference {
		instance := stage.AxesShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.AxesShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			axesshapes_deletedInstances = append(axesshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(axesshapes_newInstances)
	lenDeletedInstances += len(axesshapes_deletedInstances)
	var basevectorshapes_newInstances []*BaseVectorShape
	var basevectorshapes_deletedInstances []*BaseVectorShape

	// parse all staged instances and check if they have a reference
	for basevectorshape := range stage.BaseVectorShapes {
		if ref, ok := stage.BaseVectorShapes_reference[basevectorshape]; !ok {
			basevectorshapes_newInstances = append(basevectorshapes_newInstances, basevectorshape)
			newInstancesSlice = append(newInstancesSlice, basevectorshape.GongMarshallIdentifier(stage))
			if stage.BaseVectorShapes_referenceOrder == nil {
				stage.BaseVectorShapes_referenceOrder = make(map[*BaseVectorShape]uint)
			}
			stage.BaseVectorShapes_referenceOrder[basevectorshape] = stage.BaseVectorShape_stagedOrder[basevectorshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, basevectorshape.GongMarshallUnstaging(stage))
			// delete(stage.BaseVectorShapes_referenceOrder, basevectorshape)
			fieldInitializers, pointersInitializations := basevectorshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BaseVectorShape_stagedOrder[ref] = stage.BaseVectorShape_stagedOrder[basevectorshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := basevectorshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, basevectorshape)
			// delete(stage.BaseVectorShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if basevectorshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", basevectorshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.BaseVectorShapes_reference {
		instance := stage.BaseVectorShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.BaseVectorShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			basevectorshapes_deletedInstances = append(basevectorshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(basevectorshapes_newInstances)
	lenDeletedInstances += len(basevectorshapes_deletedInstances)
	var basevectorshapegrids_newInstances []*BaseVectorShapeGrid
	var basevectorshapegrids_deletedInstances []*BaseVectorShapeGrid

	// parse all staged instances and check if they have a reference
	for basevectorshapegrid := range stage.BaseVectorShapeGrids {
		if ref, ok := stage.BaseVectorShapeGrids_reference[basevectorshapegrid]; !ok {
			basevectorshapegrids_newInstances = append(basevectorshapegrids_newInstances, basevectorshapegrid)
			newInstancesSlice = append(newInstancesSlice, basevectorshapegrid.GongMarshallIdentifier(stage))
			if stage.BaseVectorShapeGrids_referenceOrder == nil {
				stage.BaseVectorShapeGrids_referenceOrder = make(map[*BaseVectorShapeGrid]uint)
			}
			stage.BaseVectorShapeGrids_referenceOrder[basevectorshapegrid] = stage.BaseVectorShapeGrid_stagedOrder[basevectorshapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, basevectorshapegrid.GongMarshallUnstaging(stage))
			// delete(stage.BaseVectorShapeGrids_referenceOrder, basevectorshapegrid)
			fieldInitializers, pointersInitializations := basevectorshapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BaseVectorShapeGrid_stagedOrder[ref] = stage.BaseVectorShapeGrid_stagedOrder[basevectorshapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := basevectorshapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, basevectorshapegrid)
			// delete(stage.BaseVectorShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if basevectorshapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", basevectorshapegrid.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.BaseVectorShapeGrids_reference {
		instance := stage.BaseVectorShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.BaseVectorShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			basevectorshapegrids_deletedInstances = append(basevectorshapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(basevectorshapegrids_newInstances)
	lenDeletedInstances += len(basevectorshapegrids_deletedInstances)
	var bottomendarcshapev2s_newInstances []*BottomEndArcShapeV2
	var bottomendarcshapev2s_deletedInstances []*BottomEndArcShapeV2

	// parse all staged instances and check if they have a reference
	for bottomendarcshapev2 := range stage.BottomEndArcShapeV2s {
		if ref, ok := stage.BottomEndArcShapeV2s_reference[bottomendarcshapev2]; !ok {
			bottomendarcshapev2s_newInstances = append(bottomendarcshapev2s_newInstances, bottomendarcshapev2)
			newInstancesSlice = append(newInstancesSlice, bottomendarcshapev2.GongMarshallIdentifier(stage))
			if stage.BottomEndArcShapeV2s_referenceOrder == nil {
				stage.BottomEndArcShapeV2s_referenceOrder = make(map[*BottomEndArcShapeV2]uint)
			}
			stage.BottomEndArcShapeV2s_referenceOrder[bottomendarcshapev2] = stage.BottomEndArcShapeV2_stagedOrder[bottomendarcshapev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, bottomendarcshapev2.GongMarshallUnstaging(stage))
			// delete(stage.BottomEndArcShapeV2s_referenceOrder, bottomendarcshapev2)
			fieldInitializers, pointersInitializations := bottomendarcshapev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BottomEndArcShapeV2_stagedOrder[ref] = stage.BottomEndArcShapeV2_stagedOrder[bottomendarcshapev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := bottomendarcshapev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, bottomendarcshapev2)
			// delete(stage.BottomEndArcShapeV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if bottomendarcshapev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", bottomendarcshapev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.BottomEndArcShapeV2s_reference {
		instance := stage.BottomEndArcShapeV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.BottomEndArcShapeV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			bottomendarcshapev2s_deletedInstances = append(bottomendarcshapev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(bottomendarcshapev2s_newInstances)
	lenDeletedInstances += len(bottomendarcshapev2s_deletedInstances)
	var bottomendarcshapev2grids_newInstances []*BottomEndArcShapeV2Grid
	var bottomendarcshapev2grids_deletedInstances []*BottomEndArcShapeV2Grid

	// parse all staged instances and check if they have a reference
	for bottomendarcshapev2grid := range stage.BottomEndArcShapeV2Grids {
		if ref, ok := stage.BottomEndArcShapeV2Grids_reference[bottomendarcshapev2grid]; !ok {
			bottomendarcshapev2grids_newInstances = append(bottomendarcshapev2grids_newInstances, bottomendarcshapev2grid)
			newInstancesSlice = append(newInstancesSlice, bottomendarcshapev2grid.GongMarshallIdentifier(stage))
			if stage.BottomEndArcShapeV2Grids_referenceOrder == nil {
				stage.BottomEndArcShapeV2Grids_referenceOrder = make(map[*BottomEndArcShapeV2Grid]uint)
			}
			stage.BottomEndArcShapeV2Grids_referenceOrder[bottomendarcshapev2grid] = stage.BottomEndArcShapeV2Grid_stagedOrder[bottomendarcshapev2grid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, bottomendarcshapev2grid.GongMarshallUnstaging(stage))
			// delete(stage.BottomEndArcShapeV2Grids_referenceOrder, bottomendarcshapev2grid)
			fieldInitializers, pointersInitializations := bottomendarcshapev2grid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BottomEndArcShapeV2Grid_stagedOrder[ref] = stage.BottomEndArcShapeV2Grid_stagedOrder[bottomendarcshapev2grid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := bottomendarcshapev2grid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, bottomendarcshapev2grid)
			// delete(stage.BottomEndArcShapeV2Grid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if bottomendarcshapev2grid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", bottomendarcshapev2grid.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.BottomEndArcShapeV2Grids_reference {
		instance := stage.BottomEndArcShapeV2Grids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.BottomEndArcShapeV2Grids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			bottomendarcshapev2grids_deletedInstances = append(bottomendarcshapev2grids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(bottomendarcshapev2grids_newInstances)
	lenDeletedInstances += len(bottomendarcshapev2grids_deletedInstances)
	var bottomstackgrowthcurveendarcshapev2s_newInstances []*BottomStackGrowthCurveEndArcShapeV2
	var bottomstackgrowthcurveendarcshapev2s_deletedInstances []*BottomStackGrowthCurveEndArcShapeV2

	// parse all staged instances and check if they have a reference
	for bottomstackgrowthcurveendarcshapev2 := range stage.BottomStackGrowthCurveEndArcShapeV2s {
		if ref, ok := stage.BottomStackGrowthCurveEndArcShapeV2s_reference[bottomstackgrowthcurveendarcshapev2]; !ok {
			bottomstackgrowthcurveendarcshapev2s_newInstances = append(bottomstackgrowthcurveendarcshapev2s_newInstances, bottomstackgrowthcurveendarcshapev2)
			newInstancesSlice = append(newInstancesSlice, bottomstackgrowthcurveendarcshapev2.GongMarshallIdentifier(stage))
			if stage.BottomStackGrowthCurveEndArcShapeV2s_referenceOrder == nil {
				stage.BottomStackGrowthCurveEndArcShapeV2s_referenceOrder = make(map[*BottomStackGrowthCurveEndArcShapeV2]uint)
			}
			stage.BottomStackGrowthCurveEndArcShapeV2s_referenceOrder[bottomstackgrowthcurveendarcshapev2] = stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder[bottomstackgrowthcurveendarcshapev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, bottomstackgrowthcurveendarcshapev2.GongMarshallUnstaging(stage))
			// delete(stage.BottomStackGrowthCurveEndArcShapeV2s_referenceOrder, bottomstackgrowthcurveendarcshapev2)
			fieldInitializers, pointersInitializations := bottomstackgrowthcurveendarcshapev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder[ref] = stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder[bottomstackgrowthcurveendarcshapev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := bottomstackgrowthcurveendarcshapev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, bottomstackgrowthcurveendarcshapev2)
			// delete(stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if bottomstackgrowthcurveendarcshapev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", bottomstackgrowthcurveendarcshapev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.BottomStackGrowthCurveEndArcShapeV2s_reference {
		instance := stage.BottomStackGrowthCurveEndArcShapeV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.BottomStackGrowthCurveEndArcShapeV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			bottomstackgrowthcurveendarcshapev2s_deletedInstances = append(bottomstackgrowthcurveendarcshapev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(bottomstackgrowthcurveendarcshapev2s_newInstances)
	lenDeletedInstances += len(bottomstackgrowthcurveendarcshapev2s_deletedInstances)
	var bottomstackgrowthcurvestartarcshapev2s_newInstances []*BottomStackGrowthCurveStartArcShapeV2
	var bottomstackgrowthcurvestartarcshapev2s_deletedInstances []*BottomStackGrowthCurveStartArcShapeV2

	// parse all staged instances and check if they have a reference
	for bottomstackgrowthcurvestartarcshapev2 := range stage.BottomStackGrowthCurveStartArcShapeV2s {
		if ref, ok := stage.BottomStackGrowthCurveStartArcShapeV2s_reference[bottomstackgrowthcurvestartarcshapev2]; !ok {
			bottomstackgrowthcurvestartarcshapev2s_newInstances = append(bottomstackgrowthcurvestartarcshapev2s_newInstances, bottomstackgrowthcurvestartarcshapev2)
			newInstancesSlice = append(newInstancesSlice, bottomstackgrowthcurvestartarcshapev2.GongMarshallIdentifier(stage))
			if stage.BottomStackGrowthCurveStartArcShapeV2s_referenceOrder == nil {
				stage.BottomStackGrowthCurveStartArcShapeV2s_referenceOrder = make(map[*BottomStackGrowthCurveStartArcShapeV2]uint)
			}
			stage.BottomStackGrowthCurveStartArcShapeV2s_referenceOrder[bottomstackgrowthcurvestartarcshapev2] = stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder[bottomstackgrowthcurvestartarcshapev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, bottomstackgrowthcurvestartarcshapev2.GongMarshallUnstaging(stage))
			// delete(stage.BottomStackGrowthCurveStartArcShapeV2s_referenceOrder, bottomstackgrowthcurvestartarcshapev2)
			fieldInitializers, pointersInitializations := bottomstackgrowthcurvestartarcshapev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder[ref] = stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder[bottomstackgrowthcurvestartarcshapev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := bottomstackgrowthcurvestartarcshapev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, bottomstackgrowthcurvestartarcshapev2)
			// delete(stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if bottomstackgrowthcurvestartarcshapev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", bottomstackgrowthcurvestartarcshapev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.BottomStackGrowthCurveStartArcShapeV2s_reference {
		instance := stage.BottomStackGrowthCurveStartArcShapeV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.BottomStackGrowthCurveStartArcShapeV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			bottomstackgrowthcurvestartarcshapev2s_deletedInstances = append(bottomstackgrowthcurvestartarcshapev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(bottomstackgrowthcurvestartarcshapev2s_newInstances)
	lenDeletedInstances += len(bottomstackgrowthcurvestartarcshapev2s_deletedInstances)
	var bottomstackofgrowthcurvev2s_newInstances []*BottomStackOfGrowthCurveV2
	var bottomstackofgrowthcurvev2s_deletedInstances []*BottomStackOfGrowthCurveV2

	// parse all staged instances and check if they have a reference
	for bottomstackofgrowthcurvev2 := range stage.BottomStackOfGrowthCurveV2s {
		if ref, ok := stage.BottomStackOfGrowthCurveV2s_reference[bottomstackofgrowthcurvev2]; !ok {
			bottomstackofgrowthcurvev2s_newInstances = append(bottomstackofgrowthcurvev2s_newInstances, bottomstackofgrowthcurvev2)
			newInstancesSlice = append(newInstancesSlice, bottomstackofgrowthcurvev2.GongMarshallIdentifier(stage))
			if stage.BottomStackOfGrowthCurveV2s_referenceOrder == nil {
				stage.BottomStackOfGrowthCurveV2s_referenceOrder = make(map[*BottomStackOfGrowthCurveV2]uint)
			}
			stage.BottomStackOfGrowthCurveV2s_referenceOrder[bottomstackofgrowthcurvev2] = stage.BottomStackOfGrowthCurveV2_stagedOrder[bottomstackofgrowthcurvev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, bottomstackofgrowthcurvev2.GongMarshallUnstaging(stage))
			// delete(stage.BottomStackOfGrowthCurveV2s_referenceOrder, bottomstackofgrowthcurvev2)
			fieldInitializers, pointersInitializations := bottomstackofgrowthcurvev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BottomStackOfGrowthCurveV2_stagedOrder[ref] = stage.BottomStackOfGrowthCurveV2_stagedOrder[bottomstackofgrowthcurvev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := bottomstackofgrowthcurvev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, bottomstackofgrowthcurvev2)
			// delete(stage.BottomStackOfGrowthCurveV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if bottomstackofgrowthcurvev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", bottomstackofgrowthcurvev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.BottomStackOfGrowthCurveV2s_reference {
		instance := stage.BottomStackOfGrowthCurveV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.BottomStackOfGrowthCurveV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			bottomstackofgrowthcurvev2s_deletedInstances = append(bottomstackofgrowthcurvev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(bottomstackofgrowthcurvev2s_newInstances)
	lenDeletedInstances += len(bottomstackofgrowthcurvev2s_deletedInstances)
	var bottomstartarcshapev2s_newInstances []*BottomStartArcShapeV2
	var bottomstartarcshapev2s_deletedInstances []*BottomStartArcShapeV2

	// parse all staged instances and check if they have a reference
	for bottomstartarcshapev2 := range stage.BottomStartArcShapeV2s {
		if ref, ok := stage.BottomStartArcShapeV2s_reference[bottomstartarcshapev2]; !ok {
			bottomstartarcshapev2s_newInstances = append(bottomstartarcshapev2s_newInstances, bottomstartarcshapev2)
			newInstancesSlice = append(newInstancesSlice, bottomstartarcshapev2.GongMarshallIdentifier(stage))
			if stage.BottomStartArcShapeV2s_referenceOrder == nil {
				stage.BottomStartArcShapeV2s_referenceOrder = make(map[*BottomStartArcShapeV2]uint)
			}
			stage.BottomStartArcShapeV2s_referenceOrder[bottomstartarcshapev2] = stage.BottomStartArcShapeV2_stagedOrder[bottomstartarcshapev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, bottomstartarcshapev2.GongMarshallUnstaging(stage))
			// delete(stage.BottomStartArcShapeV2s_referenceOrder, bottomstartarcshapev2)
			fieldInitializers, pointersInitializations := bottomstartarcshapev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BottomStartArcShapeV2_stagedOrder[ref] = stage.BottomStartArcShapeV2_stagedOrder[bottomstartarcshapev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := bottomstartarcshapev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, bottomstartarcshapev2)
			// delete(stage.BottomStartArcShapeV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if bottomstartarcshapev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", bottomstartarcshapev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.BottomStartArcShapeV2s_reference {
		instance := stage.BottomStartArcShapeV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.BottomStartArcShapeV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			bottomstartarcshapev2s_deletedInstances = append(bottomstartarcshapev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(bottomstartarcshapev2s_newInstances)
	lenDeletedInstances += len(bottomstartarcshapev2s_deletedInstances)
	var bottomstartarcshapev2grids_newInstances []*BottomStartArcShapeV2Grid
	var bottomstartarcshapev2grids_deletedInstances []*BottomStartArcShapeV2Grid

	// parse all staged instances and check if they have a reference
	for bottomstartarcshapev2grid := range stage.BottomStartArcShapeV2Grids {
		if ref, ok := stage.BottomStartArcShapeV2Grids_reference[bottomstartarcshapev2grid]; !ok {
			bottomstartarcshapev2grids_newInstances = append(bottomstartarcshapev2grids_newInstances, bottomstartarcshapev2grid)
			newInstancesSlice = append(newInstancesSlice, bottomstartarcshapev2grid.GongMarshallIdentifier(stage))
			if stage.BottomStartArcShapeV2Grids_referenceOrder == nil {
				stage.BottomStartArcShapeV2Grids_referenceOrder = make(map[*BottomStartArcShapeV2Grid]uint)
			}
			stage.BottomStartArcShapeV2Grids_referenceOrder[bottomstartarcshapev2grid] = stage.BottomStartArcShapeV2Grid_stagedOrder[bottomstartarcshapev2grid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, bottomstartarcshapev2grid.GongMarshallUnstaging(stage))
			// delete(stage.BottomStartArcShapeV2Grids_referenceOrder, bottomstartarcshapev2grid)
			fieldInitializers, pointersInitializations := bottomstartarcshapev2grid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BottomStartArcShapeV2Grid_stagedOrder[ref] = stage.BottomStartArcShapeV2Grid_stagedOrder[bottomstartarcshapev2grid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := bottomstartarcshapev2grid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, bottomstartarcshapev2grid)
			// delete(stage.BottomStartArcShapeV2Grid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if bottomstartarcshapev2grid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", bottomstartarcshapev2grid.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.BottomStartArcShapeV2Grids_reference {
		instance := stage.BottomStartArcShapeV2Grids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.BottomStartArcShapeV2Grids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			bottomstartarcshapev2grids_deletedInstances = append(bottomstartarcshapev2grids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(bottomstartarcshapev2grids_newInstances)
	lenDeletedInstances += len(bottomstartarcshapev2grids_deletedInstances)
	var circlegridshapes_newInstances []*CircleGridShape
	var circlegridshapes_deletedInstances []*CircleGridShape

	// parse all staged instances and check if they have a reference
	for circlegridshape := range stage.CircleGridShapes {
		if ref, ok := stage.CircleGridShapes_reference[circlegridshape]; !ok {
			circlegridshapes_newInstances = append(circlegridshapes_newInstances, circlegridshape)
			newInstancesSlice = append(newInstancesSlice, circlegridshape.GongMarshallIdentifier(stage))
			if stage.CircleGridShapes_referenceOrder == nil {
				stage.CircleGridShapes_referenceOrder = make(map[*CircleGridShape]uint)
			}
			stage.CircleGridShapes_referenceOrder[circlegridshape] = stage.CircleGridShape_stagedOrder[circlegridshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, circlegridshape.GongMarshallUnstaging(stage))
			// delete(stage.CircleGridShapes_referenceOrder, circlegridshape)
			fieldInitializers, pointersInitializations := circlegridshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CircleGridShape_stagedOrder[ref] = stage.CircleGridShape_stagedOrder[circlegridshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := circlegridshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, circlegridshape)
			// delete(stage.CircleGridShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if circlegridshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", circlegridshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.CircleGridShapes_reference {
		instance := stage.CircleGridShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.CircleGridShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			circlegridshapes_deletedInstances = append(circlegridshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(circlegridshapes_newInstances)
	lenDeletedInstances += len(circlegridshapes_deletedInstances)
	var endarcshapes_newInstances []*EndArcShape
	var endarcshapes_deletedInstances []*EndArcShape

	// parse all staged instances and check if they have a reference
	for endarcshape := range stage.EndArcShapes {
		if ref, ok := stage.EndArcShapes_reference[endarcshape]; !ok {
			endarcshapes_newInstances = append(endarcshapes_newInstances, endarcshape)
			newInstancesSlice = append(newInstancesSlice, endarcshape.GongMarshallIdentifier(stage))
			if stage.EndArcShapes_referenceOrder == nil {
				stage.EndArcShapes_referenceOrder = make(map[*EndArcShape]uint)
			}
			stage.EndArcShapes_referenceOrder[endarcshape] = stage.EndArcShape_stagedOrder[endarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, endarcshape.GongMarshallUnstaging(stage))
			// delete(stage.EndArcShapes_referenceOrder, endarcshape)
			fieldInitializers, pointersInitializations := endarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EndArcShape_stagedOrder[ref] = stage.EndArcShape_stagedOrder[endarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := endarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, endarcshape)
			// delete(stage.EndArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if endarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", endarcshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.EndArcShapes_reference {
		instance := stage.EndArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.EndArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			endarcshapes_deletedInstances = append(endarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(endarcshapes_newInstances)
	lenDeletedInstances += len(endarcshapes_deletedInstances)
	var endarcshapegrids_newInstances []*EndArcShapeGrid
	var endarcshapegrids_deletedInstances []*EndArcShapeGrid

	// parse all staged instances and check if they have a reference
	for endarcshapegrid := range stage.EndArcShapeGrids {
		if ref, ok := stage.EndArcShapeGrids_reference[endarcshapegrid]; !ok {
			endarcshapegrids_newInstances = append(endarcshapegrids_newInstances, endarcshapegrid)
			newInstancesSlice = append(newInstancesSlice, endarcshapegrid.GongMarshallIdentifier(stage))
			if stage.EndArcShapeGrids_referenceOrder == nil {
				stage.EndArcShapeGrids_referenceOrder = make(map[*EndArcShapeGrid]uint)
			}
			stage.EndArcShapeGrids_referenceOrder[endarcshapegrid] = stage.EndArcShapeGrid_stagedOrder[endarcshapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, endarcshapegrid.GongMarshallUnstaging(stage))
			// delete(stage.EndArcShapeGrids_referenceOrder, endarcshapegrid)
			fieldInitializers, pointersInitializations := endarcshapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EndArcShapeGrid_stagedOrder[ref] = stage.EndArcShapeGrid_stagedOrder[endarcshapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := endarcshapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, endarcshapegrid)
			// delete(stage.EndArcShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if endarcshapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", endarcshapegrid.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.EndArcShapeGrids_reference {
		instance := stage.EndArcShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.EndArcShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			endarcshapegrids_deletedInstances = append(endarcshapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(endarcshapegrids_newInstances)
	lenDeletedInstances += len(endarcshapegrids_deletedInstances)
	var endarcshapev2s_newInstances []*EndArcShapeV2
	var endarcshapev2s_deletedInstances []*EndArcShapeV2

	// parse all staged instances and check if they have a reference
	for endarcshapev2 := range stage.EndArcShapeV2s {
		if ref, ok := stage.EndArcShapeV2s_reference[endarcshapev2]; !ok {
			endarcshapev2s_newInstances = append(endarcshapev2s_newInstances, endarcshapev2)
			newInstancesSlice = append(newInstancesSlice, endarcshapev2.GongMarshallIdentifier(stage))
			if stage.EndArcShapeV2s_referenceOrder == nil {
				stage.EndArcShapeV2s_referenceOrder = make(map[*EndArcShapeV2]uint)
			}
			stage.EndArcShapeV2s_referenceOrder[endarcshapev2] = stage.EndArcShapeV2_stagedOrder[endarcshapev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, endarcshapev2.GongMarshallUnstaging(stage))
			// delete(stage.EndArcShapeV2s_referenceOrder, endarcshapev2)
			fieldInitializers, pointersInitializations := endarcshapev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EndArcShapeV2_stagedOrder[ref] = stage.EndArcShapeV2_stagedOrder[endarcshapev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := endarcshapev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, endarcshapev2)
			// delete(stage.EndArcShapeV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if endarcshapev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", endarcshapev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.EndArcShapeV2s_reference {
		instance := stage.EndArcShapeV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.EndArcShapeV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			endarcshapev2s_deletedInstances = append(endarcshapev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(endarcshapev2s_newInstances)
	lenDeletedInstances += len(endarcshapev2s_deletedInstances)
	var endarcshapev2grids_newInstances []*EndArcShapeV2Grid
	var endarcshapev2grids_deletedInstances []*EndArcShapeV2Grid

	// parse all staged instances and check if they have a reference
	for endarcshapev2grid := range stage.EndArcShapeV2Grids {
		if ref, ok := stage.EndArcShapeV2Grids_reference[endarcshapev2grid]; !ok {
			endarcshapev2grids_newInstances = append(endarcshapev2grids_newInstances, endarcshapev2grid)
			newInstancesSlice = append(newInstancesSlice, endarcshapev2grid.GongMarshallIdentifier(stage))
			if stage.EndArcShapeV2Grids_referenceOrder == nil {
				stage.EndArcShapeV2Grids_referenceOrder = make(map[*EndArcShapeV2Grid]uint)
			}
			stage.EndArcShapeV2Grids_referenceOrder[endarcshapev2grid] = stage.EndArcShapeV2Grid_stagedOrder[endarcshapev2grid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, endarcshapev2grid.GongMarshallUnstaging(stage))
			// delete(stage.EndArcShapeV2Grids_referenceOrder, endarcshapev2grid)
			fieldInitializers, pointersInitializations := endarcshapev2grid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EndArcShapeV2Grid_stagedOrder[ref] = stage.EndArcShapeV2Grid_stagedOrder[endarcshapev2grid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := endarcshapev2grid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, endarcshapev2grid)
			// delete(stage.EndArcShapeV2Grid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if endarcshapev2grid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", endarcshapev2grid.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.EndArcShapeV2Grids_reference {
		instance := stage.EndArcShapeV2Grids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.EndArcShapeV2Grids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			endarcshapev2grids_deletedInstances = append(endarcshapev2grids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(endarcshapev2grids_newInstances)
	lenDeletedInstances += len(endarcshapev2grids_deletedInstances)
	var explanationtextshapes_newInstances []*ExplanationTextShape
	var explanationtextshapes_deletedInstances []*ExplanationTextShape

	// parse all staged instances and check if they have a reference
	for explanationtextshape := range stage.ExplanationTextShapes {
		if ref, ok := stage.ExplanationTextShapes_reference[explanationtextshape]; !ok {
			explanationtextshapes_newInstances = append(explanationtextshapes_newInstances, explanationtextshape)
			newInstancesSlice = append(newInstancesSlice, explanationtextshape.GongMarshallIdentifier(stage))
			if stage.ExplanationTextShapes_referenceOrder == nil {
				stage.ExplanationTextShapes_referenceOrder = make(map[*ExplanationTextShape]uint)
			}
			stage.ExplanationTextShapes_referenceOrder[explanationtextshape] = stage.ExplanationTextShape_stagedOrder[explanationtextshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, explanationtextshape.GongMarshallUnstaging(stage))
			// delete(stage.ExplanationTextShapes_referenceOrder, explanationtextshape)
			fieldInitializers, pointersInitializations := explanationtextshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ExplanationTextShape_stagedOrder[ref] = stage.ExplanationTextShape_stagedOrder[explanationtextshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := explanationtextshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, explanationtextshape)
			// delete(stage.ExplanationTextShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if explanationtextshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", explanationtextshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ExplanationTextShapes_reference {
		instance := stage.ExplanationTextShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ExplanationTextShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			explanationtextshapes_deletedInstances = append(explanationtextshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(explanationtextshapes_newInstances)
	lenDeletedInstances += len(explanationtextshapes_deletedInstances)
	var gridpathshapes_newInstances []*GridPathShape
	var gridpathshapes_deletedInstances []*GridPathShape

	// parse all staged instances and check if they have a reference
	for gridpathshape := range stage.GridPathShapes {
		if ref, ok := stage.GridPathShapes_reference[gridpathshape]; !ok {
			gridpathshapes_newInstances = append(gridpathshapes_newInstances, gridpathshape)
			newInstancesSlice = append(newInstancesSlice, gridpathshape.GongMarshallIdentifier(stage))
			if stage.GridPathShapes_referenceOrder == nil {
				stage.GridPathShapes_referenceOrder = make(map[*GridPathShape]uint)
			}
			stage.GridPathShapes_referenceOrder[gridpathshape] = stage.GridPathShape_stagedOrder[gridpathshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gridpathshape.GongMarshallUnstaging(stage))
			// delete(stage.GridPathShapes_referenceOrder, gridpathshape)
			fieldInitializers, pointersInitializations := gridpathshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GridPathShape_stagedOrder[ref] = stage.GridPathShape_stagedOrder[gridpathshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := gridpathshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gridpathshape)
			// delete(stage.GridPathShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if gridpathshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", gridpathshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.GridPathShapes_reference {
		instance := stage.GridPathShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GridPathShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			gridpathshapes_deletedInstances = append(gridpathshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gridpathshapes_newInstances)
	lenDeletedInstances += len(gridpathshapes_deletedInstances)
	var growthcurvebeziershapes_newInstances []*GrowthCurveBezierShape
	var growthcurvebeziershapes_deletedInstances []*GrowthCurveBezierShape

	// parse all staged instances and check if they have a reference
	for growthcurvebeziershape := range stage.GrowthCurveBezierShapes {
		if ref, ok := stage.GrowthCurveBezierShapes_reference[growthcurvebeziershape]; !ok {
			growthcurvebeziershapes_newInstances = append(growthcurvebeziershapes_newInstances, growthcurvebeziershape)
			newInstancesSlice = append(newInstancesSlice, growthcurvebeziershape.GongMarshallIdentifier(stage))
			if stage.GrowthCurveBezierShapes_referenceOrder == nil {
				stage.GrowthCurveBezierShapes_referenceOrder = make(map[*GrowthCurveBezierShape]uint)
			}
			stage.GrowthCurveBezierShapes_referenceOrder[growthcurvebeziershape] = stage.GrowthCurveBezierShape_stagedOrder[growthcurvebeziershape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, growthcurvebeziershape.GongMarshallUnstaging(stage))
			// delete(stage.GrowthCurveBezierShapes_referenceOrder, growthcurvebeziershape)
			fieldInitializers, pointersInitializations := growthcurvebeziershape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GrowthCurveBezierShape_stagedOrder[ref] = stage.GrowthCurveBezierShape_stagedOrder[growthcurvebeziershape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := growthcurvebeziershape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, growthcurvebeziershape)
			// delete(stage.GrowthCurveBezierShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if growthcurvebeziershape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", growthcurvebeziershape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.GrowthCurveBezierShapes_reference {
		instance := stage.GrowthCurveBezierShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GrowthCurveBezierShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			growthcurvebeziershapes_deletedInstances = append(growthcurvebeziershapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(growthcurvebeziershapes_newInstances)
	lenDeletedInstances += len(growthcurvebeziershapes_deletedInstances)
	var growthcurvebeziershapegrids_newInstances []*GrowthCurveBezierShapeGrid
	var growthcurvebeziershapegrids_deletedInstances []*GrowthCurveBezierShapeGrid

	// parse all staged instances and check if they have a reference
	for growthcurvebeziershapegrid := range stage.GrowthCurveBezierShapeGrids {
		if ref, ok := stage.GrowthCurveBezierShapeGrids_reference[growthcurvebeziershapegrid]; !ok {
			growthcurvebeziershapegrids_newInstances = append(growthcurvebeziershapegrids_newInstances, growthcurvebeziershapegrid)
			newInstancesSlice = append(newInstancesSlice, growthcurvebeziershapegrid.GongMarshallIdentifier(stage))
			if stage.GrowthCurveBezierShapeGrids_referenceOrder == nil {
				stage.GrowthCurveBezierShapeGrids_referenceOrder = make(map[*GrowthCurveBezierShapeGrid]uint)
			}
			stage.GrowthCurveBezierShapeGrids_referenceOrder[growthcurvebeziershapegrid] = stage.GrowthCurveBezierShapeGrid_stagedOrder[growthcurvebeziershapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, growthcurvebeziershapegrid.GongMarshallUnstaging(stage))
			// delete(stage.GrowthCurveBezierShapeGrids_referenceOrder, growthcurvebeziershapegrid)
			fieldInitializers, pointersInitializations := growthcurvebeziershapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GrowthCurveBezierShapeGrid_stagedOrder[ref] = stage.GrowthCurveBezierShapeGrid_stagedOrder[growthcurvebeziershapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := growthcurvebeziershapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, growthcurvebeziershapegrid)
			// delete(stage.GrowthCurveBezierShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if growthcurvebeziershapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", growthcurvebeziershapegrid.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.GrowthCurveBezierShapeGrids_reference {
		instance := stage.GrowthCurveBezierShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GrowthCurveBezierShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			growthcurvebeziershapegrids_deletedInstances = append(growthcurvebeziershapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(growthcurvebeziershapegrids_newInstances)
	lenDeletedInstances += len(growthcurvebeziershapegrids_deletedInstances)
	var growthcurverhombusgridshapes_newInstances []*GrowthCurveRhombusGridShape
	var growthcurverhombusgridshapes_deletedInstances []*GrowthCurveRhombusGridShape

	// parse all staged instances and check if they have a reference
	for growthcurverhombusgridshape := range stage.GrowthCurveRhombusGridShapes {
		if ref, ok := stage.GrowthCurveRhombusGridShapes_reference[growthcurverhombusgridshape]; !ok {
			growthcurverhombusgridshapes_newInstances = append(growthcurverhombusgridshapes_newInstances, growthcurverhombusgridshape)
			newInstancesSlice = append(newInstancesSlice, growthcurverhombusgridshape.GongMarshallIdentifier(stage))
			if stage.GrowthCurveRhombusGridShapes_referenceOrder == nil {
				stage.GrowthCurveRhombusGridShapes_referenceOrder = make(map[*GrowthCurveRhombusGridShape]uint)
			}
			stage.GrowthCurveRhombusGridShapes_referenceOrder[growthcurverhombusgridshape] = stage.GrowthCurveRhombusGridShape_stagedOrder[growthcurverhombusgridshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, growthcurverhombusgridshape.GongMarshallUnstaging(stage))
			// delete(stage.GrowthCurveRhombusGridShapes_referenceOrder, growthcurverhombusgridshape)
			fieldInitializers, pointersInitializations := growthcurverhombusgridshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GrowthCurveRhombusGridShape_stagedOrder[ref] = stage.GrowthCurveRhombusGridShape_stagedOrder[growthcurverhombusgridshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := growthcurverhombusgridshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, growthcurverhombusgridshape)
			// delete(stage.GrowthCurveRhombusGridShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if growthcurverhombusgridshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", growthcurverhombusgridshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.GrowthCurveRhombusGridShapes_reference {
		instance := stage.GrowthCurveRhombusGridShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GrowthCurveRhombusGridShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			growthcurverhombusgridshapes_deletedInstances = append(growthcurverhombusgridshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(growthcurverhombusgridshapes_newInstances)
	lenDeletedInstances += len(growthcurverhombusgridshapes_deletedInstances)
	var growthcurverhombusshapes_newInstances []*GrowthCurveRhombusShape
	var growthcurverhombusshapes_deletedInstances []*GrowthCurveRhombusShape

	// parse all staged instances and check if they have a reference
	for growthcurverhombusshape := range stage.GrowthCurveRhombusShapes {
		if ref, ok := stage.GrowthCurveRhombusShapes_reference[growthcurverhombusshape]; !ok {
			growthcurverhombusshapes_newInstances = append(growthcurverhombusshapes_newInstances, growthcurverhombusshape)
			newInstancesSlice = append(newInstancesSlice, growthcurverhombusshape.GongMarshallIdentifier(stage))
			if stage.GrowthCurveRhombusShapes_referenceOrder == nil {
				stage.GrowthCurveRhombusShapes_referenceOrder = make(map[*GrowthCurveRhombusShape]uint)
			}
			stage.GrowthCurveRhombusShapes_referenceOrder[growthcurverhombusshape] = stage.GrowthCurveRhombusShape_stagedOrder[growthcurverhombusshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, growthcurverhombusshape.GongMarshallUnstaging(stage))
			// delete(stage.GrowthCurveRhombusShapes_referenceOrder, growthcurverhombusshape)
			fieldInitializers, pointersInitializations := growthcurverhombusshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GrowthCurveRhombusShape_stagedOrder[ref] = stage.GrowthCurveRhombusShape_stagedOrder[growthcurverhombusshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := growthcurverhombusshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, growthcurverhombusshape)
			// delete(stage.GrowthCurveRhombusShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if growthcurverhombusshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", growthcurverhombusshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.GrowthCurveRhombusShapes_reference {
		instance := stage.GrowthCurveRhombusShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GrowthCurveRhombusShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			growthcurverhombusshapes_deletedInstances = append(growthcurverhombusshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(growthcurverhombusshapes_newInstances)
	lenDeletedInstances += len(growthcurverhombusshapes_deletedInstances)
	var growthvectorshapes_newInstances []*GrowthVectorShape
	var growthvectorshapes_deletedInstances []*GrowthVectorShape

	// parse all staged instances and check if they have a reference
	for growthvectorshape := range stage.GrowthVectorShapes {
		if ref, ok := stage.GrowthVectorShapes_reference[growthvectorshape]; !ok {
			growthvectorshapes_newInstances = append(growthvectorshapes_newInstances, growthvectorshape)
			newInstancesSlice = append(newInstancesSlice, growthvectorshape.GongMarshallIdentifier(stage))
			if stage.GrowthVectorShapes_referenceOrder == nil {
				stage.GrowthVectorShapes_referenceOrder = make(map[*GrowthVectorShape]uint)
			}
			stage.GrowthVectorShapes_referenceOrder[growthvectorshape] = stage.GrowthVectorShape_stagedOrder[growthvectorshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, growthvectorshape.GongMarshallUnstaging(stage))
			// delete(stage.GrowthVectorShapes_referenceOrder, growthvectorshape)
			fieldInitializers, pointersInitializations := growthvectorshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GrowthVectorShape_stagedOrder[ref] = stage.GrowthVectorShape_stagedOrder[growthvectorshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := growthvectorshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, growthvectorshape)
			// delete(stage.GrowthVectorShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if growthvectorshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", growthvectorshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.GrowthVectorShapes_reference {
		instance := stage.GrowthVectorShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GrowthVectorShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			growthvectorshapes_deletedInstances = append(growthvectorshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(growthvectorshapes_newInstances)
	lenDeletedInstances += len(growthvectorshapes_deletedInstances)
	var initialrhombusgridshapes_newInstances []*InitialRhombusGridShape
	var initialrhombusgridshapes_deletedInstances []*InitialRhombusGridShape

	// parse all staged instances and check if they have a reference
	for initialrhombusgridshape := range stage.InitialRhombusGridShapes {
		if ref, ok := stage.InitialRhombusGridShapes_reference[initialrhombusgridshape]; !ok {
			initialrhombusgridshapes_newInstances = append(initialrhombusgridshapes_newInstances, initialrhombusgridshape)
			newInstancesSlice = append(newInstancesSlice, initialrhombusgridshape.GongMarshallIdentifier(stage))
			if stage.InitialRhombusGridShapes_referenceOrder == nil {
				stage.InitialRhombusGridShapes_referenceOrder = make(map[*InitialRhombusGridShape]uint)
			}
			stage.InitialRhombusGridShapes_referenceOrder[initialrhombusgridshape] = stage.InitialRhombusGridShape_stagedOrder[initialrhombusgridshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, initialrhombusgridshape.GongMarshallUnstaging(stage))
			// delete(stage.InitialRhombusGridShapes_referenceOrder, initialrhombusgridshape)
			fieldInitializers, pointersInitializations := initialrhombusgridshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.InitialRhombusGridShape_stagedOrder[ref] = stage.InitialRhombusGridShape_stagedOrder[initialrhombusgridshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := initialrhombusgridshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, initialrhombusgridshape)
			// delete(stage.InitialRhombusGridShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if initialrhombusgridshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", initialrhombusgridshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.InitialRhombusGridShapes_reference {
		instance := stage.InitialRhombusGridShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.InitialRhombusGridShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			initialrhombusgridshapes_deletedInstances = append(initialrhombusgridshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(initialrhombusgridshapes_newInstances)
	lenDeletedInstances += len(initialrhombusgridshapes_deletedInstances)
	var initialrhombusshapes_newInstances []*InitialRhombusShape
	var initialrhombusshapes_deletedInstances []*InitialRhombusShape

	// parse all staged instances and check if they have a reference
	for initialrhombusshape := range stage.InitialRhombusShapes {
		if ref, ok := stage.InitialRhombusShapes_reference[initialrhombusshape]; !ok {
			initialrhombusshapes_newInstances = append(initialrhombusshapes_newInstances, initialrhombusshape)
			newInstancesSlice = append(newInstancesSlice, initialrhombusshape.GongMarshallIdentifier(stage))
			if stage.InitialRhombusShapes_referenceOrder == nil {
				stage.InitialRhombusShapes_referenceOrder = make(map[*InitialRhombusShape]uint)
			}
			stage.InitialRhombusShapes_referenceOrder[initialrhombusshape] = stage.InitialRhombusShape_stagedOrder[initialrhombusshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, initialrhombusshape.GongMarshallUnstaging(stage))
			// delete(stage.InitialRhombusShapes_referenceOrder, initialrhombusshape)
			fieldInitializers, pointersInitializations := initialrhombusshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.InitialRhombusShape_stagedOrder[ref] = stage.InitialRhombusShape_stagedOrder[initialrhombusshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := initialrhombusshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, initialrhombusshape)
			// delete(stage.InitialRhombusShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if initialrhombusshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", initialrhombusshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.InitialRhombusShapes_reference {
		instance := stage.InitialRhombusShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.InitialRhombusShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			initialrhombusshapes_deletedInstances = append(initialrhombusshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(initialrhombusshapes_newInstances)
	lenDeletedInstances += len(initialrhombusshapes_deletedInstances)
	var librarys_newInstances []*Library
	var librarys_deletedInstances []*Library

	// parse all staged instances and check if they have a reference
	for library := range stage.Librarys {
		if ref, ok := stage.Librarys_reference[library]; !ok {
			librarys_newInstances = append(librarys_newInstances, library)
			newInstancesSlice = append(newInstancesSlice, library.GongMarshallIdentifier(stage))
			if stage.Librarys_referenceOrder == nil {
				stage.Librarys_referenceOrder = make(map[*Library]uint)
			}
			stage.Librarys_referenceOrder[library] = stage.Library_stagedOrder[library]
			newInstancesReverseSlice = append(newInstancesReverseSlice, library.GongMarshallUnstaging(stage))
			// delete(stage.Librarys_referenceOrder, library)
			fieldInitializers, pointersInitializations := library.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Library_stagedOrder[ref] = stage.Library_stagedOrder[library]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := library.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, library)
			// delete(stage.Library_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if library.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", library.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Librarys_reference {
		instance := stage.Librarys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Librarys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			librarys_deletedInstances = append(librarys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(librarys_newInstances)
	lenDeletedInstances += len(librarys_deletedInstances)
	var nextcircleshapes_newInstances []*NextCircleShape
	var nextcircleshapes_deletedInstances []*NextCircleShape

	// parse all staged instances and check if they have a reference
	for nextcircleshape := range stage.NextCircleShapes {
		if ref, ok := stage.NextCircleShapes_reference[nextcircleshape]; !ok {
			nextcircleshapes_newInstances = append(nextcircleshapes_newInstances, nextcircleshape)
			newInstancesSlice = append(newInstancesSlice, nextcircleshape.GongMarshallIdentifier(stage))
			if stage.NextCircleShapes_referenceOrder == nil {
				stage.NextCircleShapes_referenceOrder = make(map[*NextCircleShape]uint)
			}
			stage.NextCircleShapes_referenceOrder[nextcircleshape] = stage.NextCircleShape_stagedOrder[nextcircleshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, nextcircleshape.GongMarshallUnstaging(stage))
			// delete(stage.NextCircleShapes_referenceOrder, nextcircleshape)
			fieldInitializers, pointersInitializations := nextcircleshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NextCircleShape_stagedOrder[ref] = stage.NextCircleShape_stagedOrder[nextcircleshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := nextcircleshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, nextcircleshape)
			// delete(stage.NextCircleShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if nextcircleshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", nextcircleshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.NextCircleShapes_reference {
		instance := stage.NextCircleShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NextCircleShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			nextcircleshapes_deletedInstances = append(nextcircleshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(nextcircleshapes_newInstances)
	lenDeletedInstances += len(nextcircleshapes_deletedInstances)
	var perpendicularvectors_newInstances []*PerpendicularVector
	var perpendicularvectors_deletedInstances []*PerpendicularVector

	// parse all staged instances and check if they have a reference
	for perpendicularvector := range stage.PerpendicularVectors {
		if ref, ok := stage.PerpendicularVectors_reference[perpendicularvector]; !ok {
			perpendicularvectors_newInstances = append(perpendicularvectors_newInstances, perpendicularvector)
			newInstancesSlice = append(newInstancesSlice, perpendicularvector.GongMarshallIdentifier(stage))
			if stage.PerpendicularVectors_referenceOrder == nil {
				stage.PerpendicularVectors_referenceOrder = make(map[*PerpendicularVector]uint)
			}
			stage.PerpendicularVectors_referenceOrder[perpendicularvector] = stage.PerpendicularVector_stagedOrder[perpendicularvector]
			newInstancesReverseSlice = append(newInstancesReverseSlice, perpendicularvector.GongMarshallUnstaging(stage))
			// delete(stage.PerpendicularVectors_referenceOrder, perpendicularvector)
			fieldInitializers, pointersInitializations := perpendicularvector.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PerpendicularVector_stagedOrder[ref] = stage.PerpendicularVector_stagedOrder[perpendicularvector]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := perpendicularvector.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, perpendicularvector)
			// delete(stage.PerpendicularVector_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if perpendicularvector.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", perpendicularvector.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.PerpendicularVectors_reference {
		instance := stage.PerpendicularVectors_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PerpendicularVectors[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			perpendicularvectors_deletedInstances = append(perpendicularvectors_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(perpendicularvectors_newInstances)
	lenDeletedInstances += len(perpendicularvectors_deletedInstances)
	var perpendicularvectorgrids_newInstances []*PerpendicularVectorGrid
	var perpendicularvectorgrids_deletedInstances []*PerpendicularVectorGrid

	// parse all staged instances and check if they have a reference
	for perpendicularvectorgrid := range stage.PerpendicularVectorGrids {
		if ref, ok := stage.PerpendicularVectorGrids_reference[perpendicularvectorgrid]; !ok {
			perpendicularvectorgrids_newInstances = append(perpendicularvectorgrids_newInstances, perpendicularvectorgrid)
			newInstancesSlice = append(newInstancesSlice, perpendicularvectorgrid.GongMarshallIdentifier(stage))
			if stage.PerpendicularVectorGrids_referenceOrder == nil {
				stage.PerpendicularVectorGrids_referenceOrder = make(map[*PerpendicularVectorGrid]uint)
			}
			stage.PerpendicularVectorGrids_referenceOrder[perpendicularvectorgrid] = stage.PerpendicularVectorGrid_stagedOrder[perpendicularvectorgrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, perpendicularvectorgrid.GongMarshallUnstaging(stage))
			// delete(stage.PerpendicularVectorGrids_referenceOrder, perpendicularvectorgrid)
			fieldInitializers, pointersInitializations := perpendicularvectorgrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PerpendicularVectorGrid_stagedOrder[ref] = stage.PerpendicularVectorGrid_stagedOrder[perpendicularvectorgrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := perpendicularvectorgrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, perpendicularvectorgrid)
			// delete(stage.PerpendicularVectorGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if perpendicularvectorgrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", perpendicularvectorgrid.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.PerpendicularVectorGrids_reference {
		instance := stage.PerpendicularVectorGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PerpendicularVectorGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			perpendicularvectorgrids_deletedInstances = append(perpendicularvectorgrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(perpendicularvectorgrids_newInstances)
	lenDeletedInstances += len(perpendicularvectorgrids_deletedInstances)
	var perpendicularvectorgridhalfways_newInstances []*PerpendicularVectorGridHalfway
	var perpendicularvectorgridhalfways_deletedInstances []*PerpendicularVectorGridHalfway

	// parse all staged instances and check if they have a reference
	for perpendicularvectorgridhalfway := range stage.PerpendicularVectorGridHalfways {
		if ref, ok := stage.PerpendicularVectorGridHalfways_reference[perpendicularvectorgridhalfway]; !ok {
			perpendicularvectorgridhalfways_newInstances = append(perpendicularvectorgridhalfways_newInstances, perpendicularvectorgridhalfway)
			newInstancesSlice = append(newInstancesSlice, perpendicularvectorgridhalfway.GongMarshallIdentifier(stage))
			if stage.PerpendicularVectorGridHalfways_referenceOrder == nil {
				stage.PerpendicularVectorGridHalfways_referenceOrder = make(map[*PerpendicularVectorGridHalfway]uint)
			}
			stage.PerpendicularVectorGridHalfways_referenceOrder[perpendicularvectorgridhalfway] = stage.PerpendicularVectorGridHalfway_stagedOrder[perpendicularvectorgridhalfway]
			newInstancesReverseSlice = append(newInstancesReverseSlice, perpendicularvectorgridhalfway.GongMarshallUnstaging(stage))
			// delete(stage.PerpendicularVectorGridHalfways_referenceOrder, perpendicularvectorgridhalfway)
			fieldInitializers, pointersInitializations := perpendicularvectorgridhalfway.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PerpendicularVectorGridHalfway_stagedOrder[ref] = stage.PerpendicularVectorGridHalfway_stagedOrder[perpendicularvectorgridhalfway]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := perpendicularvectorgridhalfway.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, perpendicularvectorgridhalfway)
			// delete(stage.PerpendicularVectorGridHalfway_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if perpendicularvectorgridhalfway.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", perpendicularvectorgridhalfway.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.PerpendicularVectorGridHalfways_reference {
		instance := stage.PerpendicularVectorGridHalfways_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PerpendicularVectorGridHalfways[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			perpendicularvectorgridhalfways_deletedInstances = append(perpendicularvectorgridhalfways_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(perpendicularvectorgridhalfways_newInstances)
	lenDeletedInstances += len(perpendicularvectorgridhalfways_deletedInstances)
	var perpendicularvectorhalfways_newInstances []*PerpendicularVectorHalfway
	var perpendicularvectorhalfways_deletedInstances []*PerpendicularVectorHalfway

	// parse all staged instances and check if they have a reference
	for perpendicularvectorhalfway := range stage.PerpendicularVectorHalfways {
		if ref, ok := stage.PerpendicularVectorHalfways_reference[perpendicularvectorhalfway]; !ok {
			perpendicularvectorhalfways_newInstances = append(perpendicularvectorhalfways_newInstances, perpendicularvectorhalfway)
			newInstancesSlice = append(newInstancesSlice, perpendicularvectorhalfway.GongMarshallIdentifier(stage))
			if stage.PerpendicularVectorHalfways_referenceOrder == nil {
				stage.PerpendicularVectorHalfways_referenceOrder = make(map[*PerpendicularVectorHalfway]uint)
			}
			stage.PerpendicularVectorHalfways_referenceOrder[perpendicularvectorhalfway] = stage.PerpendicularVectorHalfway_stagedOrder[perpendicularvectorhalfway]
			newInstancesReverseSlice = append(newInstancesReverseSlice, perpendicularvectorhalfway.GongMarshallUnstaging(stage))
			// delete(stage.PerpendicularVectorHalfways_referenceOrder, perpendicularvectorhalfway)
			fieldInitializers, pointersInitializations := perpendicularvectorhalfway.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PerpendicularVectorHalfway_stagedOrder[ref] = stage.PerpendicularVectorHalfway_stagedOrder[perpendicularvectorhalfway]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := perpendicularvectorhalfway.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, perpendicularvectorhalfway)
			// delete(stage.PerpendicularVectorHalfway_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if perpendicularvectorhalfway.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", perpendicularvectorhalfway.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.PerpendicularVectorHalfways_reference {
		instance := stage.PerpendicularVectorHalfways_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PerpendicularVectorHalfways[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			perpendicularvectorhalfways_deletedInstances = append(perpendicularvectorhalfways_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(perpendicularvectorhalfways_newInstances)
	lenDeletedInstances += len(perpendicularvectorhalfways_deletedInstances)
	var plants_newInstances []*Plant
	var plants_deletedInstances []*Plant

	// parse all staged instances and check if they have a reference
	for plant := range stage.Plants {
		if ref, ok := stage.Plants_reference[plant]; !ok {
			plants_newInstances = append(plants_newInstances, plant)
			newInstancesSlice = append(newInstancesSlice, plant.GongMarshallIdentifier(stage))
			if stage.Plants_referenceOrder == nil {
				stage.Plants_referenceOrder = make(map[*Plant]uint)
			}
			stage.Plants_referenceOrder[plant] = stage.Plant_stagedOrder[plant]
			newInstancesReverseSlice = append(newInstancesReverseSlice, plant.GongMarshallUnstaging(stage))
			// delete(stage.Plants_referenceOrder, plant)
			fieldInitializers, pointersInitializations := plant.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Plant_stagedOrder[ref] = stage.Plant_stagedOrder[plant]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := plant.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, plant)
			// delete(stage.Plant_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if plant.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", plant.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Plants_reference {
		instance := stage.Plants_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Plants[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			plants_deletedInstances = append(plants_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(plants_newInstances)
	lenDeletedInstances += len(plants_deletedInstances)
	var plantcircumferenceshapes_newInstances []*PlantCircumferenceShape
	var plantcircumferenceshapes_deletedInstances []*PlantCircumferenceShape

	// parse all staged instances and check if they have a reference
	for plantcircumferenceshape := range stage.PlantCircumferenceShapes {
		if ref, ok := stage.PlantCircumferenceShapes_reference[plantcircumferenceshape]; !ok {
			plantcircumferenceshapes_newInstances = append(plantcircumferenceshapes_newInstances, plantcircumferenceshape)
			newInstancesSlice = append(newInstancesSlice, plantcircumferenceshape.GongMarshallIdentifier(stage))
			if stage.PlantCircumferenceShapes_referenceOrder == nil {
				stage.PlantCircumferenceShapes_referenceOrder = make(map[*PlantCircumferenceShape]uint)
			}
			stage.PlantCircumferenceShapes_referenceOrder[plantcircumferenceshape] = stage.PlantCircumferenceShape_stagedOrder[plantcircumferenceshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, plantcircumferenceshape.GongMarshallUnstaging(stage))
			// delete(stage.PlantCircumferenceShapes_referenceOrder, plantcircumferenceshape)
			fieldInitializers, pointersInitializations := plantcircumferenceshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PlantCircumferenceShape_stagedOrder[ref] = stage.PlantCircumferenceShape_stagedOrder[plantcircumferenceshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := plantcircumferenceshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, plantcircumferenceshape)
			// delete(stage.PlantCircumferenceShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if plantcircumferenceshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", plantcircumferenceshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.PlantCircumferenceShapes_reference {
		instance := stage.PlantCircumferenceShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PlantCircumferenceShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			plantcircumferenceshapes_deletedInstances = append(plantcircumferenceshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(plantcircumferenceshapes_newInstances)
	lenDeletedInstances += len(plantcircumferenceshapes_deletedInstances)
	var plantdiagrams_newInstances []*PlantDiagram
	var plantdiagrams_deletedInstances []*PlantDiagram

	// parse all staged instances and check if they have a reference
	for plantdiagram := range stage.PlantDiagrams {
		if ref, ok := stage.PlantDiagrams_reference[plantdiagram]; !ok {
			plantdiagrams_newInstances = append(plantdiagrams_newInstances, plantdiagram)
			newInstancesSlice = append(newInstancesSlice, plantdiagram.GongMarshallIdentifier(stage))
			if stage.PlantDiagrams_referenceOrder == nil {
				stage.PlantDiagrams_referenceOrder = make(map[*PlantDiagram]uint)
			}
			stage.PlantDiagrams_referenceOrder[plantdiagram] = stage.PlantDiagram_stagedOrder[plantdiagram]
			newInstancesReverseSlice = append(newInstancesReverseSlice, plantdiagram.GongMarshallUnstaging(stage))
			// delete(stage.PlantDiagrams_referenceOrder, plantdiagram)
			fieldInitializers, pointersInitializations := plantdiagram.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PlantDiagram_stagedOrder[ref] = stage.PlantDiagram_stagedOrder[plantdiagram]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := plantdiagram.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, plantdiagram)
			// delete(stage.PlantDiagram_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if plantdiagram.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", plantdiagram.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.PlantDiagrams_reference {
		instance := stage.PlantDiagrams_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PlantDiagrams[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			plantdiagrams_deletedInstances = append(plantdiagrams_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(plantdiagrams_newInstances)
	lenDeletedInstances += len(plantdiagrams_deletedInstances)
	var rendered3dshapes_newInstances []*Rendered3DShape
	var rendered3dshapes_deletedInstances []*Rendered3DShape

	// parse all staged instances and check if they have a reference
	for rendered3dshape := range stage.Rendered3DShapes {
		if ref, ok := stage.Rendered3DShapes_reference[rendered3dshape]; !ok {
			rendered3dshapes_newInstances = append(rendered3dshapes_newInstances, rendered3dshape)
			newInstancesSlice = append(newInstancesSlice, rendered3dshape.GongMarshallIdentifier(stage))
			if stage.Rendered3DShapes_referenceOrder == nil {
				stage.Rendered3DShapes_referenceOrder = make(map[*Rendered3DShape]uint)
			}
			stage.Rendered3DShapes_referenceOrder[rendered3dshape] = stage.Rendered3DShape_stagedOrder[rendered3dshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, rendered3dshape.GongMarshallUnstaging(stage))
			// delete(stage.Rendered3DShapes_referenceOrder, rendered3dshape)
			fieldInitializers, pointersInitializations := rendered3dshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Rendered3DShape_stagedOrder[ref] = stage.Rendered3DShape_stagedOrder[rendered3dshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := rendered3dshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, rendered3dshape)
			// delete(stage.Rendered3DShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if rendered3dshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", rendered3dshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Rendered3DShapes_reference {
		instance := stage.Rendered3DShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Rendered3DShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			rendered3dshapes_deletedInstances = append(rendered3dshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(rendered3dshapes_newInstances)
	lenDeletedInstances += len(rendered3dshapes_deletedInstances)
	var rhombusshapes_newInstances []*RhombusShape
	var rhombusshapes_deletedInstances []*RhombusShape

	// parse all staged instances and check if they have a reference
	for rhombusshape := range stage.RhombusShapes {
		if ref, ok := stage.RhombusShapes_reference[rhombusshape]; !ok {
			rhombusshapes_newInstances = append(rhombusshapes_newInstances, rhombusshape)
			newInstancesSlice = append(newInstancesSlice, rhombusshape.GongMarshallIdentifier(stage))
			if stage.RhombusShapes_referenceOrder == nil {
				stage.RhombusShapes_referenceOrder = make(map[*RhombusShape]uint)
			}
			stage.RhombusShapes_referenceOrder[rhombusshape] = stage.RhombusShape_stagedOrder[rhombusshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, rhombusshape.GongMarshallUnstaging(stage))
			// delete(stage.RhombusShapes_referenceOrder, rhombusshape)
			fieldInitializers, pointersInitializations := rhombusshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RhombusShape_stagedOrder[ref] = stage.RhombusShape_stagedOrder[rhombusshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := rhombusshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, rhombusshape)
			// delete(stage.RhombusShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if rhombusshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", rhombusshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.RhombusShapes_reference {
		instance := stage.RhombusShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.RhombusShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			rhombusshapes_deletedInstances = append(rhombusshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(rhombusshapes_newInstances)
	lenDeletedInstances += len(rhombusshapes_deletedInstances)
	var rotatedrhombusgridshapes_newInstances []*RotatedRhombusGridShape
	var rotatedrhombusgridshapes_deletedInstances []*RotatedRhombusGridShape

	// parse all staged instances and check if they have a reference
	for rotatedrhombusgridshape := range stage.RotatedRhombusGridShapes {
		if ref, ok := stage.RotatedRhombusGridShapes_reference[rotatedrhombusgridshape]; !ok {
			rotatedrhombusgridshapes_newInstances = append(rotatedrhombusgridshapes_newInstances, rotatedrhombusgridshape)
			newInstancesSlice = append(newInstancesSlice, rotatedrhombusgridshape.GongMarshallIdentifier(stage))
			if stage.RotatedRhombusGridShapes_referenceOrder == nil {
				stage.RotatedRhombusGridShapes_referenceOrder = make(map[*RotatedRhombusGridShape]uint)
			}
			stage.RotatedRhombusGridShapes_referenceOrder[rotatedrhombusgridshape] = stage.RotatedRhombusGridShape_stagedOrder[rotatedrhombusgridshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, rotatedrhombusgridshape.GongMarshallUnstaging(stage))
			// delete(stage.RotatedRhombusGridShapes_referenceOrder, rotatedrhombusgridshape)
			fieldInitializers, pointersInitializations := rotatedrhombusgridshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RotatedRhombusGridShape_stagedOrder[ref] = stage.RotatedRhombusGridShape_stagedOrder[rotatedrhombusgridshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := rotatedrhombusgridshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, rotatedrhombusgridshape)
			// delete(stage.RotatedRhombusGridShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if rotatedrhombusgridshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", rotatedrhombusgridshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.RotatedRhombusGridShapes_reference {
		instance := stage.RotatedRhombusGridShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.RotatedRhombusGridShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			rotatedrhombusgridshapes_deletedInstances = append(rotatedrhombusgridshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(rotatedrhombusgridshapes_newInstances)
	lenDeletedInstances += len(rotatedrhombusgridshapes_deletedInstances)
	var rotatedrhombusshapes_newInstances []*RotatedRhombusShape
	var rotatedrhombusshapes_deletedInstances []*RotatedRhombusShape

	// parse all staged instances and check if they have a reference
	for rotatedrhombusshape := range stage.RotatedRhombusShapes {
		if ref, ok := stage.RotatedRhombusShapes_reference[rotatedrhombusshape]; !ok {
			rotatedrhombusshapes_newInstances = append(rotatedrhombusshapes_newInstances, rotatedrhombusshape)
			newInstancesSlice = append(newInstancesSlice, rotatedrhombusshape.GongMarshallIdentifier(stage))
			if stage.RotatedRhombusShapes_referenceOrder == nil {
				stage.RotatedRhombusShapes_referenceOrder = make(map[*RotatedRhombusShape]uint)
			}
			stage.RotatedRhombusShapes_referenceOrder[rotatedrhombusshape] = stage.RotatedRhombusShape_stagedOrder[rotatedrhombusshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, rotatedrhombusshape.GongMarshallUnstaging(stage))
			// delete(stage.RotatedRhombusShapes_referenceOrder, rotatedrhombusshape)
			fieldInitializers, pointersInitializations := rotatedrhombusshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RotatedRhombusShape_stagedOrder[ref] = stage.RotatedRhombusShape_stagedOrder[rotatedrhombusshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := rotatedrhombusshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, rotatedrhombusshape)
			// delete(stage.RotatedRhombusShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if rotatedrhombusshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", rotatedrhombusshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.RotatedRhombusShapes_reference {
		instance := stage.RotatedRhombusShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.RotatedRhombusShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			rotatedrhombusshapes_deletedInstances = append(rotatedrhombusshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(rotatedrhombusshapes_newInstances)
	lenDeletedInstances += len(rotatedrhombusshapes_deletedInstances)
	var stackgrowthcurvebeziershapes_newInstances []*StackGrowthCurveBezierShape
	var stackgrowthcurvebeziershapes_deletedInstances []*StackGrowthCurveBezierShape

	// parse all staged instances and check if they have a reference
	for stackgrowthcurvebeziershape := range stage.StackGrowthCurveBezierShapes {
		if ref, ok := stage.StackGrowthCurveBezierShapes_reference[stackgrowthcurvebeziershape]; !ok {
			stackgrowthcurvebeziershapes_newInstances = append(stackgrowthcurvebeziershapes_newInstances, stackgrowthcurvebeziershape)
			newInstancesSlice = append(newInstancesSlice, stackgrowthcurvebeziershape.GongMarshallIdentifier(stage))
			if stage.StackGrowthCurveBezierShapes_referenceOrder == nil {
				stage.StackGrowthCurveBezierShapes_referenceOrder = make(map[*StackGrowthCurveBezierShape]uint)
			}
			stage.StackGrowthCurveBezierShapes_referenceOrder[stackgrowthcurvebeziershape] = stage.StackGrowthCurveBezierShape_stagedOrder[stackgrowthcurvebeziershape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackgrowthcurvebeziershape.GongMarshallUnstaging(stage))
			// delete(stage.StackGrowthCurveBezierShapes_referenceOrder, stackgrowthcurvebeziershape)
			fieldInitializers, pointersInitializations := stackgrowthcurvebeziershape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackGrowthCurveBezierShape_stagedOrder[ref] = stage.StackGrowthCurveBezierShape_stagedOrder[stackgrowthcurvebeziershape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackgrowthcurvebeziershape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackgrowthcurvebeziershape)
			// delete(stage.StackGrowthCurveBezierShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackgrowthcurvebeziershape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackgrowthcurvebeziershape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StackGrowthCurveBezierShapes_reference {
		instance := stage.StackGrowthCurveBezierShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackGrowthCurveBezierShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackgrowthcurvebeziershapes_deletedInstances = append(stackgrowthcurvebeziershapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackgrowthcurvebeziershapes_newInstances)
	lenDeletedInstances += len(stackgrowthcurvebeziershapes_deletedInstances)
	var stackgrowthcurveendarcshapev2s_newInstances []*StackGrowthCurveEndArcShapeV2
	var stackgrowthcurveendarcshapev2s_deletedInstances []*StackGrowthCurveEndArcShapeV2

	// parse all staged instances and check if they have a reference
	for stackgrowthcurveendarcshapev2 := range stage.StackGrowthCurveEndArcShapeV2s {
		if ref, ok := stage.StackGrowthCurveEndArcShapeV2s_reference[stackgrowthcurveendarcshapev2]; !ok {
			stackgrowthcurveendarcshapev2s_newInstances = append(stackgrowthcurveendarcshapev2s_newInstances, stackgrowthcurveendarcshapev2)
			newInstancesSlice = append(newInstancesSlice, stackgrowthcurveendarcshapev2.GongMarshallIdentifier(stage))
			if stage.StackGrowthCurveEndArcShapeV2s_referenceOrder == nil {
				stage.StackGrowthCurveEndArcShapeV2s_referenceOrder = make(map[*StackGrowthCurveEndArcShapeV2]uint)
			}
			stage.StackGrowthCurveEndArcShapeV2s_referenceOrder[stackgrowthcurveendarcshapev2] = stage.StackGrowthCurveEndArcShapeV2_stagedOrder[stackgrowthcurveendarcshapev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackgrowthcurveendarcshapev2.GongMarshallUnstaging(stage))
			// delete(stage.StackGrowthCurveEndArcShapeV2s_referenceOrder, stackgrowthcurveendarcshapev2)
			fieldInitializers, pointersInitializations := stackgrowthcurveendarcshapev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackGrowthCurveEndArcShapeV2_stagedOrder[ref] = stage.StackGrowthCurveEndArcShapeV2_stagedOrder[stackgrowthcurveendarcshapev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackgrowthcurveendarcshapev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackgrowthcurveendarcshapev2)
			// delete(stage.StackGrowthCurveEndArcShapeV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackgrowthcurveendarcshapev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackgrowthcurveendarcshapev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StackGrowthCurveEndArcShapeV2s_reference {
		instance := stage.StackGrowthCurveEndArcShapeV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackGrowthCurveEndArcShapeV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackgrowthcurveendarcshapev2s_deletedInstances = append(stackgrowthcurveendarcshapev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackgrowthcurveendarcshapev2s_newInstances)
	lenDeletedInstances += len(stackgrowthcurveendarcshapev2s_deletedInstances)
	var stackgrowthcurvestartarcshapev2s_newInstances []*StackGrowthCurveStartArcShapeV2
	var stackgrowthcurvestartarcshapev2s_deletedInstances []*StackGrowthCurveStartArcShapeV2

	// parse all staged instances and check if they have a reference
	for stackgrowthcurvestartarcshapev2 := range stage.StackGrowthCurveStartArcShapeV2s {
		if ref, ok := stage.StackGrowthCurveStartArcShapeV2s_reference[stackgrowthcurvestartarcshapev2]; !ok {
			stackgrowthcurvestartarcshapev2s_newInstances = append(stackgrowthcurvestartarcshapev2s_newInstances, stackgrowthcurvestartarcshapev2)
			newInstancesSlice = append(newInstancesSlice, stackgrowthcurvestartarcshapev2.GongMarshallIdentifier(stage))
			if stage.StackGrowthCurveStartArcShapeV2s_referenceOrder == nil {
				stage.StackGrowthCurveStartArcShapeV2s_referenceOrder = make(map[*StackGrowthCurveStartArcShapeV2]uint)
			}
			stage.StackGrowthCurveStartArcShapeV2s_referenceOrder[stackgrowthcurvestartarcshapev2] = stage.StackGrowthCurveStartArcShapeV2_stagedOrder[stackgrowthcurvestartarcshapev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackgrowthcurvestartarcshapev2.GongMarshallUnstaging(stage))
			// delete(stage.StackGrowthCurveStartArcShapeV2s_referenceOrder, stackgrowthcurvestartarcshapev2)
			fieldInitializers, pointersInitializations := stackgrowthcurvestartarcshapev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackGrowthCurveStartArcShapeV2_stagedOrder[ref] = stage.StackGrowthCurveStartArcShapeV2_stagedOrder[stackgrowthcurvestartarcshapev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackgrowthcurvestartarcshapev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackgrowthcurvestartarcshapev2)
			// delete(stage.StackGrowthCurveStartArcShapeV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackgrowthcurvestartarcshapev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackgrowthcurvestartarcshapev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StackGrowthCurveStartArcShapeV2s_reference {
		instance := stage.StackGrowthCurveStartArcShapeV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackGrowthCurveStartArcShapeV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackgrowthcurvestartarcshapev2s_deletedInstances = append(stackgrowthcurvestartarcshapev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackgrowthcurvestartarcshapev2s_newInstances)
	lenDeletedInstances += len(stackgrowthcurvestartarcshapev2s_deletedInstances)
	var stackofgrowthcurves_newInstances []*StackOfGrowthCurve
	var stackofgrowthcurves_deletedInstances []*StackOfGrowthCurve

	// parse all staged instances and check if they have a reference
	for stackofgrowthcurve := range stage.StackOfGrowthCurves {
		if ref, ok := stage.StackOfGrowthCurves_reference[stackofgrowthcurve]; !ok {
			stackofgrowthcurves_newInstances = append(stackofgrowthcurves_newInstances, stackofgrowthcurve)
			newInstancesSlice = append(newInstancesSlice, stackofgrowthcurve.GongMarshallIdentifier(stage))
			if stage.StackOfGrowthCurves_referenceOrder == nil {
				stage.StackOfGrowthCurves_referenceOrder = make(map[*StackOfGrowthCurve]uint)
			}
			stage.StackOfGrowthCurves_referenceOrder[stackofgrowthcurve] = stage.StackOfGrowthCurve_stagedOrder[stackofgrowthcurve]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackofgrowthcurve.GongMarshallUnstaging(stage))
			// delete(stage.StackOfGrowthCurves_referenceOrder, stackofgrowthcurve)
			fieldInitializers, pointersInitializations := stackofgrowthcurve.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackOfGrowthCurve_stagedOrder[ref] = stage.StackOfGrowthCurve_stagedOrder[stackofgrowthcurve]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackofgrowthcurve.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackofgrowthcurve)
			// delete(stage.StackOfGrowthCurve_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackofgrowthcurve.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackofgrowthcurve.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StackOfGrowthCurves_reference {
		instance := stage.StackOfGrowthCurves_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackOfGrowthCurves[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackofgrowthcurves_deletedInstances = append(stackofgrowthcurves_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackofgrowthcurves_newInstances)
	lenDeletedInstances += len(stackofgrowthcurves_deletedInstances)
	var stackofgrowthcurvev2s_newInstances []*StackOfGrowthCurveV2
	var stackofgrowthcurvev2s_deletedInstances []*StackOfGrowthCurveV2

	// parse all staged instances and check if they have a reference
	for stackofgrowthcurvev2 := range stage.StackOfGrowthCurveV2s {
		if ref, ok := stage.StackOfGrowthCurveV2s_reference[stackofgrowthcurvev2]; !ok {
			stackofgrowthcurvev2s_newInstances = append(stackofgrowthcurvev2s_newInstances, stackofgrowthcurvev2)
			newInstancesSlice = append(newInstancesSlice, stackofgrowthcurvev2.GongMarshallIdentifier(stage))
			if stage.StackOfGrowthCurveV2s_referenceOrder == nil {
				stage.StackOfGrowthCurveV2s_referenceOrder = make(map[*StackOfGrowthCurveV2]uint)
			}
			stage.StackOfGrowthCurveV2s_referenceOrder[stackofgrowthcurvev2] = stage.StackOfGrowthCurveV2_stagedOrder[stackofgrowthcurvev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackofgrowthcurvev2.GongMarshallUnstaging(stage))
			// delete(stage.StackOfGrowthCurveV2s_referenceOrder, stackofgrowthcurvev2)
			fieldInitializers, pointersInitializations := stackofgrowthcurvev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackOfGrowthCurveV2_stagedOrder[ref] = stage.StackOfGrowthCurveV2_stagedOrder[stackofgrowthcurvev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackofgrowthcurvev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackofgrowthcurvev2)
			// delete(stage.StackOfGrowthCurveV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackofgrowthcurvev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackofgrowthcurvev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StackOfGrowthCurveV2s_reference {
		instance := stage.StackOfGrowthCurveV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackOfGrowthCurveV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackofgrowthcurvev2s_deletedInstances = append(stackofgrowthcurvev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackofgrowthcurvev2s_newInstances)
	lenDeletedInstances += len(stackofgrowthcurvev2s_deletedInstances)
	var startarcshapes_newInstances []*StartArcShape
	var startarcshapes_deletedInstances []*StartArcShape

	// parse all staged instances and check if they have a reference
	for startarcshape := range stage.StartArcShapes {
		if ref, ok := stage.StartArcShapes_reference[startarcshape]; !ok {
			startarcshapes_newInstances = append(startarcshapes_newInstances, startarcshape)
			newInstancesSlice = append(newInstancesSlice, startarcshape.GongMarshallIdentifier(stage))
			if stage.StartArcShapes_referenceOrder == nil {
				stage.StartArcShapes_referenceOrder = make(map[*StartArcShape]uint)
			}
			stage.StartArcShapes_referenceOrder[startarcshape] = stage.StartArcShape_stagedOrder[startarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, startarcshape.GongMarshallUnstaging(stage))
			// delete(stage.StartArcShapes_referenceOrder, startarcshape)
			fieldInitializers, pointersInitializations := startarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StartArcShape_stagedOrder[ref] = stage.StartArcShape_stagedOrder[startarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := startarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, startarcshape)
			// delete(stage.StartArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if startarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", startarcshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StartArcShapes_reference {
		instance := stage.StartArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StartArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			startarcshapes_deletedInstances = append(startarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(startarcshapes_newInstances)
	lenDeletedInstances += len(startarcshapes_deletedInstances)
	var startarcshapegrids_newInstances []*StartArcShapeGrid
	var startarcshapegrids_deletedInstances []*StartArcShapeGrid

	// parse all staged instances and check if they have a reference
	for startarcshapegrid := range stage.StartArcShapeGrids {
		if ref, ok := stage.StartArcShapeGrids_reference[startarcshapegrid]; !ok {
			startarcshapegrids_newInstances = append(startarcshapegrids_newInstances, startarcshapegrid)
			newInstancesSlice = append(newInstancesSlice, startarcshapegrid.GongMarshallIdentifier(stage))
			if stage.StartArcShapeGrids_referenceOrder == nil {
				stage.StartArcShapeGrids_referenceOrder = make(map[*StartArcShapeGrid]uint)
			}
			stage.StartArcShapeGrids_referenceOrder[startarcshapegrid] = stage.StartArcShapeGrid_stagedOrder[startarcshapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, startarcshapegrid.GongMarshallUnstaging(stage))
			// delete(stage.StartArcShapeGrids_referenceOrder, startarcshapegrid)
			fieldInitializers, pointersInitializations := startarcshapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StartArcShapeGrid_stagedOrder[ref] = stage.StartArcShapeGrid_stagedOrder[startarcshapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := startarcshapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, startarcshapegrid)
			// delete(stage.StartArcShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if startarcshapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", startarcshapegrid.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StartArcShapeGrids_reference {
		instance := stage.StartArcShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StartArcShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			startarcshapegrids_deletedInstances = append(startarcshapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(startarcshapegrids_newInstances)
	lenDeletedInstances += len(startarcshapegrids_deletedInstances)
	var startarcshapev2s_newInstances []*StartArcShapeV2
	var startarcshapev2s_deletedInstances []*StartArcShapeV2

	// parse all staged instances and check if they have a reference
	for startarcshapev2 := range stage.StartArcShapeV2s {
		if ref, ok := stage.StartArcShapeV2s_reference[startarcshapev2]; !ok {
			startarcshapev2s_newInstances = append(startarcshapev2s_newInstances, startarcshapev2)
			newInstancesSlice = append(newInstancesSlice, startarcshapev2.GongMarshallIdentifier(stage))
			if stage.StartArcShapeV2s_referenceOrder == nil {
				stage.StartArcShapeV2s_referenceOrder = make(map[*StartArcShapeV2]uint)
			}
			stage.StartArcShapeV2s_referenceOrder[startarcshapev2] = stage.StartArcShapeV2_stagedOrder[startarcshapev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, startarcshapev2.GongMarshallUnstaging(stage))
			// delete(stage.StartArcShapeV2s_referenceOrder, startarcshapev2)
			fieldInitializers, pointersInitializations := startarcshapev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StartArcShapeV2_stagedOrder[ref] = stage.StartArcShapeV2_stagedOrder[startarcshapev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := startarcshapev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, startarcshapev2)
			// delete(stage.StartArcShapeV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if startarcshapev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", startarcshapev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StartArcShapeV2s_reference {
		instance := stage.StartArcShapeV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StartArcShapeV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			startarcshapev2s_deletedInstances = append(startarcshapev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(startarcshapev2s_newInstances)
	lenDeletedInstances += len(startarcshapev2s_deletedInstances)
	var startarcshapev2grids_newInstances []*StartArcShapeV2Grid
	var startarcshapev2grids_deletedInstances []*StartArcShapeV2Grid

	// parse all staged instances and check if they have a reference
	for startarcshapev2grid := range stage.StartArcShapeV2Grids {
		if ref, ok := stage.StartArcShapeV2Grids_reference[startarcshapev2grid]; !ok {
			startarcshapev2grids_newInstances = append(startarcshapev2grids_newInstances, startarcshapev2grid)
			newInstancesSlice = append(newInstancesSlice, startarcshapev2grid.GongMarshallIdentifier(stage))
			if stage.StartArcShapeV2Grids_referenceOrder == nil {
				stage.StartArcShapeV2Grids_referenceOrder = make(map[*StartArcShapeV2Grid]uint)
			}
			stage.StartArcShapeV2Grids_referenceOrder[startarcshapev2grid] = stage.StartArcShapeV2Grid_stagedOrder[startarcshapev2grid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, startarcshapev2grid.GongMarshallUnstaging(stage))
			// delete(stage.StartArcShapeV2Grids_referenceOrder, startarcshapev2grid)
			fieldInitializers, pointersInitializations := startarcshapev2grid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StartArcShapeV2Grid_stagedOrder[ref] = stage.StartArcShapeV2Grid_stagedOrder[startarcshapev2grid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := startarcshapev2grid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, startarcshapev2grid)
			// delete(stage.StartArcShapeV2Grid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if startarcshapev2grid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", startarcshapev2grid.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StartArcShapeV2Grids_reference {
		instance := stage.StartArcShapeV2Grids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StartArcShapeV2Grids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			startarcshapev2grids_deletedInstances = append(startarcshapev2grids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(startarcshapev2grids_newInstances)
	lenDeletedInstances += len(startarcshapev2grids_deletedInstances)
	var topendarcshapev2s_newInstances []*TopEndArcShapeV2
	var topendarcshapev2s_deletedInstances []*TopEndArcShapeV2

	// parse all staged instances and check if they have a reference
	for topendarcshapev2 := range stage.TopEndArcShapeV2s {
		if ref, ok := stage.TopEndArcShapeV2s_reference[topendarcshapev2]; !ok {
			topendarcshapev2s_newInstances = append(topendarcshapev2s_newInstances, topendarcshapev2)
			newInstancesSlice = append(newInstancesSlice, topendarcshapev2.GongMarshallIdentifier(stage))
			if stage.TopEndArcShapeV2s_referenceOrder == nil {
				stage.TopEndArcShapeV2s_referenceOrder = make(map[*TopEndArcShapeV2]uint)
			}
			stage.TopEndArcShapeV2s_referenceOrder[topendarcshapev2] = stage.TopEndArcShapeV2_stagedOrder[topendarcshapev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topendarcshapev2.GongMarshallUnstaging(stage))
			// delete(stage.TopEndArcShapeV2s_referenceOrder, topendarcshapev2)
			fieldInitializers, pointersInitializations := topendarcshapev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopEndArcShapeV2_stagedOrder[ref] = stage.TopEndArcShapeV2_stagedOrder[topendarcshapev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topendarcshapev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topendarcshapev2)
			// delete(stage.TopEndArcShapeV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topendarcshapev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topendarcshapev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.TopEndArcShapeV2s_reference {
		instance := stage.TopEndArcShapeV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopEndArcShapeV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topendarcshapev2s_deletedInstances = append(topendarcshapev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topendarcshapev2s_newInstances)
	lenDeletedInstances += len(topendarcshapev2s_deletedInstances)
	var topendarcshapev2grids_newInstances []*TopEndArcShapeV2Grid
	var topendarcshapev2grids_deletedInstances []*TopEndArcShapeV2Grid

	// parse all staged instances and check if they have a reference
	for topendarcshapev2grid := range stage.TopEndArcShapeV2Grids {
		if ref, ok := stage.TopEndArcShapeV2Grids_reference[topendarcshapev2grid]; !ok {
			topendarcshapev2grids_newInstances = append(topendarcshapev2grids_newInstances, topendarcshapev2grid)
			newInstancesSlice = append(newInstancesSlice, topendarcshapev2grid.GongMarshallIdentifier(stage))
			if stage.TopEndArcShapeV2Grids_referenceOrder == nil {
				stage.TopEndArcShapeV2Grids_referenceOrder = make(map[*TopEndArcShapeV2Grid]uint)
			}
			stage.TopEndArcShapeV2Grids_referenceOrder[topendarcshapev2grid] = stage.TopEndArcShapeV2Grid_stagedOrder[topendarcshapev2grid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topendarcshapev2grid.GongMarshallUnstaging(stage))
			// delete(stage.TopEndArcShapeV2Grids_referenceOrder, topendarcshapev2grid)
			fieldInitializers, pointersInitializations := topendarcshapev2grid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopEndArcShapeV2Grid_stagedOrder[ref] = stage.TopEndArcShapeV2Grid_stagedOrder[topendarcshapev2grid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topendarcshapev2grid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topendarcshapev2grid)
			// delete(stage.TopEndArcShapeV2Grid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topendarcshapev2grid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topendarcshapev2grid.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.TopEndArcShapeV2Grids_reference {
		instance := stage.TopEndArcShapeV2Grids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopEndArcShapeV2Grids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topendarcshapev2grids_deletedInstances = append(topendarcshapev2grids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topendarcshapev2grids_newInstances)
	lenDeletedInstances += len(topendarcshapev2grids_deletedInstances)
	var topstackgrowthcurveendarcshapev2s_newInstances []*TopStackGrowthCurveEndArcShapeV2
	var topstackgrowthcurveendarcshapev2s_deletedInstances []*TopStackGrowthCurveEndArcShapeV2

	// parse all staged instances and check if they have a reference
	for topstackgrowthcurveendarcshapev2 := range stage.TopStackGrowthCurveEndArcShapeV2s {
		if ref, ok := stage.TopStackGrowthCurveEndArcShapeV2s_reference[topstackgrowthcurveendarcshapev2]; !ok {
			topstackgrowthcurveendarcshapev2s_newInstances = append(topstackgrowthcurveendarcshapev2s_newInstances, topstackgrowthcurveendarcshapev2)
			newInstancesSlice = append(newInstancesSlice, topstackgrowthcurveendarcshapev2.GongMarshallIdentifier(stage))
			if stage.TopStackGrowthCurveEndArcShapeV2s_referenceOrder == nil {
				stage.TopStackGrowthCurveEndArcShapeV2s_referenceOrder = make(map[*TopStackGrowthCurveEndArcShapeV2]uint)
			}
			stage.TopStackGrowthCurveEndArcShapeV2s_referenceOrder[topstackgrowthcurveendarcshapev2] = stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder[topstackgrowthcurveendarcshapev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstackgrowthcurveendarcshapev2.GongMarshallUnstaging(stage))
			// delete(stage.TopStackGrowthCurveEndArcShapeV2s_referenceOrder, topstackgrowthcurveendarcshapev2)
			fieldInitializers, pointersInitializations := topstackgrowthcurveendarcshapev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder[ref] = stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder[topstackgrowthcurveendarcshapev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstackgrowthcurveendarcshapev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstackgrowthcurveendarcshapev2)
			// delete(stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstackgrowthcurveendarcshapev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstackgrowthcurveendarcshapev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.TopStackGrowthCurveEndArcShapeV2s_reference {
		instance := stage.TopStackGrowthCurveEndArcShapeV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStackGrowthCurveEndArcShapeV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstackgrowthcurveendarcshapev2s_deletedInstances = append(topstackgrowthcurveendarcshapev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstackgrowthcurveendarcshapev2s_newInstances)
	lenDeletedInstances += len(topstackgrowthcurveendarcshapev2s_deletedInstances)
	var topstackgrowthcurvestartarcshapev2s_newInstances []*TopStackGrowthCurveStartArcShapeV2
	var topstackgrowthcurvestartarcshapev2s_deletedInstances []*TopStackGrowthCurveStartArcShapeV2

	// parse all staged instances and check if they have a reference
	for topstackgrowthcurvestartarcshapev2 := range stage.TopStackGrowthCurveStartArcShapeV2s {
		if ref, ok := stage.TopStackGrowthCurveStartArcShapeV2s_reference[topstackgrowthcurvestartarcshapev2]; !ok {
			topstackgrowthcurvestartarcshapev2s_newInstances = append(topstackgrowthcurvestartarcshapev2s_newInstances, topstackgrowthcurvestartarcshapev2)
			newInstancesSlice = append(newInstancesSlice, topstackgrowthcurvestartarcshapev2.GongMarshallIdentifier(stage))
			if stage.TopStackGrowthCurveStartArcShapeV2s_referenceOrder == nil {
				stage.TopStackGrowthCurveStartArcShapeV2s_referenceOrder = make(map[*TopStackGrowthCurveStartArcShapeV2]uint)
			}
			stage.TopStackGrowthCurveStartArcShapeV2s_referenceOrder[topstackgrowthcurvestartarcshapev2] = stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder[topstackgrowthcurvestartarcshapev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstackgrowthcurvestartarcshapev2.GongMarshallUnstaging(stage))
			// delete(stage.TopStackGrowthCurveStartArcShapeV2s_referenceOrder, topstackgrowthcurvestartarcshapev2)
			fieldInitializers, pointersInitializations := topstackgrowthcurvestartarcshapev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder[ref] = stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder[topstackgrowthcurvestartarcshapev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstackgrowthcurvestartarcshapev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstackgrowthcurvestartarcshapev2)
			// delete(stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstackgrowthcurvestartarcshapev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstackgrowthcurvestartarcshapev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.TopStackGrowthCurveStartArcShapeV2s_reference {
		instance := stage.TopStackGrowthCurveStartArcShapeV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStackGrowthCurveStartArcShapeV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstackgrowthcurvestartarcshapev2s_deletedInstances = append(topstackgrowthcurvestartarcshapev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstackgrowthcurvestartarcshapev2s_newInstances)
	lenDeletedInstances += len(topstackgrowthcurvestartarcshapev2s_deletedInstances)
	var topstackofgrowthcurvev2s_newInstances []*TopStackOfGrowthCurveV2
	var topstackofgrowthcurvev2s_deletedInstances []*TopStackOfGrowthCurveV2

	// parse all staged instances and check if they have a reference
	for topstackofgrowthcurvev2 := range stage.TopStackOfGrowthCurveV2s {
		if ref, ok := stage.TopStackOfGrowthCurveV2s_reference[topstackofgrowthcurvev2]; !ok {
			topstackofgrowthcurvev2s_newInstances = append(topstackofgrowthcurvev2s_newInstances, topstackofgrowthcurvev2)
			newInstancesSlice = append(newInstancesSlice, topstackofgrowthcurvev2.GongMarshallIdentifier(stage))
			if stage.TopStackOfGrowthCurveV2s_referenceOrder == nil {
				stage.TopStackOfGrowthCurveV2s_referenceOrder = make(map[*TopStackOfGrowthCurveV2]uint)
			}
			stage.TopStackOfGrowthCurveV2s_referenceOrder[topstackofgrowthcurvev2] = stage.TopStackOfGrowthCurveV2_stagedOrder[topstackofgrowthcurvev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstackofgrowthcurvev2.GongMarshallUnstaging(stage))
			// delete(stage.TopStackOfGrowthCurveV2s_referenceOrder, topstackofgrowthcurvev2)
			fieldInitializers, pointersInitializations := topstackofgrowthcurvev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStackOfGrowthCurveV2_stagedOrder[ref] = stage.TopStackOfGrowthCurveV2_stagedOrder[topstackofgrowthcurvev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstackofgrowthcurvev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstackofgrowthcurvev2)
			// delete(stage.TopStackOfGrowthCurveV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstackofgrowthcurvev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstackofgrowthcurvev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.TopStackOfGrowthCurveV2s_reference {
		instance := stage.TopStackOfGrowthCurveV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStackOfGrowthCurveV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstackofgrowthcurvev2s_deletedInstances = append(topstackofgrowthcurvev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstackofgrowthcurvev2s_newInstances)
	lenDeletedInstances += len(topstackofgrowthcurvev2s_deletedInstances)
	var topstartarcshapev2s_newInstances []*TopStartArcShapeV2
	var topstartarcshapev2s_deletedInstances []*TopStartArcShapeV2

	// parse all staged instances and check if they have a reference
	for topstartarcshapev2 := range stage.TopStartArcShapeV2s {
		if ref, ok := stage.TopStartArcShapeV2s_reference[topstartarcshapev2]; !ok {
			topstartarcshapev2s_newInstances = append(topstartarcshapev2s_newInstances, topstartarcshapev2)
			newInstancesSlice = append(newInstancesSlice, topstartarcshapev2.GongMarshallIdentifier(stage))
			if stage.TopStartArcShapeV2s_referenceOrder == nil {
				stage.TopStartArcShapeV2s_referenceOrder = make(map[*TopStartArcShapeV2]uint)
			}
			stage.TopStartArcShapeV2s_referenceOrder[topstartarcshapev2] = stage.TopStartArcShapeV2_stagedOrder[topstartarcshapev2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstartarcshapev2.GongMarshallUnstaging(stage))
			// delete(stage.TopStartArcShapeV2s_referenceOrder, topstartarcshapev2)
			fieldInitializers, pointersInitializations := topstartarcshapev2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStartArcShapeV2_stagedOrder[ref] = stage.TopStartArcShapeV2_stagedOrder[topstartarcshapev2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstartarcshapev2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstartarcshapev2)
			// delete(stage.TopStartArcShapeV2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstartarcshapev2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstartarcshapev2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.TopStartArcShapeV2s_reference {
		instance := stage.TopStartArcShapeV2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStartArcShapeV2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstartarcshapev2s_deletedInstances = append(topstartarcshapev2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstartarcshapev2s_newInstances)
	lenDeletedInstances += len(topstartarcshapev2s_deletedInstances)
	var topstartarcshapev2grids_newInstances []*TopStartArcShapeV2Grid
	var topstartarcshapev2grids_deletedInstances []*TopStartArcShapeV2Grid

	// parse all staged instances and check if they have a reference
	for topstartarcshapev2grid := range stage.TopStartArcShapeV2Grids {
		if ref, ok := stage.TopStartArcShapeV2Grids_reference[topstartarcshapev2grid]; !ok {
			topstartarcshapev2grids_newInstances = append(topstartarcshapev2grids_newInstances, topstartarcshapev2grid)
			newInstancesSlice = append(newInstancesSlice, topstartarcshapev2grid.GongMarshallIdentifier(stage))
			if stage.TopStartArcShapeV2Grids_referenceOrder == nil {
				stage.TopStartArcShapeV2Grids_referenceOrder = make(map[*TopStartArcShapeV2Grid]uint)
			}
			stage.TopStartArcShapeV2Grids_referenceOrder[topstartarcshapev2grid] = stage.TopStartArcShapeV2Grid_stagedOrder[topstartarcshapev2grid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstartarcshapev2grid.GongMarshallUnstaging(stage))
			// delete(stage.TopStartArcShapeV2Grids_referenceOrder, topstartarcshapev2grid)
			fieldInitializers, pointersInitializations := topstartarcshapev2grid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStartArcShapeV2Grid_stagedOrder[ref] = stage.TopStartArcShapeV2Grid_stagedOrder[topstartarcshapev2grid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstartarcshapev2grid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstartarcshapev2grid)
			// delete(stage.TopStartArcShapeV2Grid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstartarcshapev2grid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstartarcshapev2grid.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.TopStartArcShapeV2Grids_reference {
		instance := stage.TopStartArcShapeV2Grids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStartArcShapeV2Grids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstartarcshapev2grids_deletedInstances = append(topstartarcshapev2grids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstartarcshapev2grids_newInstances)
	lenDeletedInstances += len(topstartarcshapev2grids_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.ArcNormalVectorShapes_reference = make(map[*ArcNormalVectorShape]*ArcNormalVectorShape)
	stage.ArcNormalVectorShapes_referenceOrder = make(map[*ArcNormalVectorShape]uint) // diff Unstage needs the reference order
	stage.ArcNormalVectorShapes_instance = make(map[*ArcNormalVectorShape]*ArcNormalVectorShape)
	for instance := range stage.ArcNormalVectorShapes {
		_copy := instance.GongCopy().(*ArcNormalVectorShape)
		stage.ArcNormalVectorShapes_reference[instance] = _copy
		stage.ArcNormalVectorShapes_instance[_copy] = instance
		stage.ArcNormalVectorShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ArcNormalVectorShapeGrids_reference = make(map[*ArcNormalVectorShapeGrid]*ArcNormalVectorShapeGrid)
	stage.ArcNormalVectorShapeGrids_referenceOrder = make(map[*ArcNormalVectorShapeGrid]uint) // diff Unstage needs the reference order
	stage.ArcNormalVectorShapeGrids_instance = make(map[*ArcNormalVectorShapeGrid]*ArcNormalVectorShapeGrid)
	for instance := range stage.ArcNormalVectorShapeGrids {
		_copy := instance.GongCopy().(*ArcNormalVectorShapeGrid)
		stage.ArcNormalVectorShapeGrids_reference[instance] = _copy
		stage.ArcNormalVectorShapeGrids_instance[_copy] = instance
		stage.ArcNormalVectorShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.AxesShapes_reference = make(map[*AxesShape]*AxesShape)
	stage.AxesShapes_referenceOrder = make(map[*AxesShape]uint) // diff Unstage needs the reference order
	stage.AxesShapes_instance = make(map[*AxesShape]*AxesShape)
	for instance := range stage.AxesShapes {
		_copy := instance.GongCopy().(*AxesShape)
		stage.AxesShapes_reference[instance] = _copy
		stage.AxesShapes_instance[_copy] = instance
		stage.AxesShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BaseVectorShapes_reference = make(map[*BaseVectorShape]*BaseVectorShape)
	stage.BaseVectorShapes_referenceOrder = make(map[*BaseVectorShape]uint) // diff Unstage needs the reference order
	stage.BaseVectorShapes_instance = make(map[*BaseVectorShape]*BaseVectorShape)
	for instance := range stage.BaseVectorShapes {
		_copy := instance.GongCopy().(*BaseVectorShape)
		stage.BaseVectorShapes_reference[instance] = _copy
		stage.BaseVectorShapes_instance[_copy] = instance
		stage.BaseVectorShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BaseVectorShapeGrids_reference = make(map[*BaseVectorShapeGrid]*BaseVectorShapeGrid)
	stage.BaseVectorShapeGrids_referenceOrder = make(map[*BaseVectorShapeGrid]uint) // diff Unstage needs the reference order
	stage.BaseVectorShapeGrids_instance = make(map[*BaseVectorShapeGrid]*BaseVectorShapeGrid)
	for instance := range stage.BaseVectorShapeGrids {
		_copy := instance.GongCopy().(*BaseVectorShapeGrid)
		stage.BaseVectorShapeGrids_reference[instance] = _copy
		stage.BaseVectorShapeGrids_instance[_copy] = instance
		stage.BaseVectorShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BottomEndArcShapeV2s_reference = make(map[*BottomEndArcShapeV2]*BottomEndArcShapeV2)
	stage.BottomEndArcShapeV2s_referenceOrder = make(map[*BottomEndArcShapeV2]uint) // diff Unstage needs the reference order
	stage.BottomEndArcShapeV2s_instance = make(map[*BottomEndArcShapeV2]*BottomEndArcShapeV2)
	for instance := range stage.BottomEndArcShapeV2s {
		_copy := instance.GongCopy().(*BottomEndArcShapeV2)
		stage.BottomEndArcShapeV2s_reference[instance] = _copy
		stage.BottomEndArcShapeV2s_instance[_copy] = instance
		stage.BottomEndArcShapeV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BottomEndArcShapeV2Grids_reference = make(map[*BottomEndArcShapeV2Grid]*BottomEndArcShapeV2Grid)
	stage.BottomEndArcShapeV2Grids_referenceOrder = make(map[*BottomEndArcShapeV2Grid]uint) // diff Unstage needs the reference order
	stage.BottomEndArcShapeV2Grids_instance = make(map[*BottomEndArcShapeV2Grid]*BottomEndArcShapeV2Grid)
	for instance := range stage.BottomEndArcShapeV2Grids {
		_copy := instance.GongCopy().(*BottomEndArcShapeV2Grid)
		stage.BottomEndArcShapeV2Grids_reference[instance] = _copy
		stage.BottomEndArcShapeV2Grids_instance[_copy] = instance
		stage.BottomEndArcShapeV2Grids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BottomStackGrowthCurveEndArcShapeV2s_reference = make(map[*BottomStackGrowthCurveEndArcShapeV2]*BottomStackGrowthCurveEndArcShapeV2)
	stage.BottomStackGrowthCurveEndArcShapeV2s_referenceOrder = make(map[*BottomStackGrowthCurveEndArcShapeV2]uint) // diff Unstage needs the reference order
	stage.BottomStackGrowthCurveEndArcShapeV2s_instance = make(map[*BottomStackGrowthCurveEndArcShapeV2]*BottomStackGrowthCurveEndArcShapeV2)
	for instance := range stage.BottomStackGrowthCurveEndArcShapeV2s {
		_copy := instance.GongCopy().(*BottomStackGrowthCurveEndArcShapeV2)
		stage.BottomStackGrowthCurveEndArcShapeV2s_reference[instance] = _copy
		stage.BottomStackGrowthCurveEndArcShapeV2s_instance[_copy] = instance
		stage.BottomStackGrowthCurveEndArcShapeV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BottomStackGrowthCurveStartArcShapeV2s_reference = make(map[*BottomStackGrowthCurveStartArcShapeV2]*BottomStackGrowthCurveStartArcShapeV2)
	stage.BottomStackGrowthCurveStartArcShapeV2s_referenceOrder = make(map[*BottomStackGrowthCurveStartArcShapeV2]uint) // diff Unstage needs the reference order
	stage.BottomStackGrowthCurveStartArcShapeV2s_instance = make(map[*BottomStackGrowthCurveStartArcShapeV2]*BottomStackGrowthCurveStartArcShapeV2)
	for instance := range stage.BottomStackGrowthCurveStartArcShapeV2s {
		_copy := instance.GongCopy().(*BottomStackGrowthCurveStartArcShapeV2)
		stage.BottomStackGrowthCurveStartArcShapeV2s_reference[instance] = _copy
		stage.BottomStackGrowthCurveStartArcShapeV2s_instance[_copy] = instance
		stage.BottomStackGrowthCurveStartArcShapeV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BottomStackOfGrowthCurveV2s_reference = make(map[*BottomStackOfGrowthCurveV2]*BottomStackOfGrowthCurveV2)
	stage.BottomStackOfGrowthCurveV2s_referenceOrder = make(map[*BottomStackOfGrowthCurveV2]uint) // diff Unstage needs the reference order
	stage.BottomStackOfGrowthCurveV2s_instance = make(map[*BottomStackOfGrowthCurveV2]*BottomStackOfGrowthCurveV2)
	for instance := range stage.BottomStackOfGrowthCurveV2s {
		_copy := instance.GongCopy().(*BottomStackOfGrowthCurveV2)
		stage.BottomStackOfGrowthCurveV2s_reference[instance] = _copy
		stage.BottomStackOfGrowthCurveV2s_instance[_copy] = instance
		stage.BottomStackOfGrowthCurveV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BottomStartArcShapeV2s_reference = make(map[*BottomStartArcShapeV2]*BottomStartArcShapeV2)
	stage.BottomStartArcShapeV2s_referenceOrder = make(map[*BottomStartArcShapeV2]uint) // diff Unstage needs the reference order
	stage.BottomStartArcShapeV2s_instance = make(map[*BottomStartArcShapeV2]*BottomStartArcShapeV2)
	for instance := range stage.BottomStartArcShapeV2s {
		_copy := instance.GongCopy().(*BottomStartArcShapeV2)
		stage.BottomStartArcShapeV2s_reference[instance] = _copy
		stage.BottomStartArcShapeV2s_instance[_copy] = instance
		stage.BottomStartArcShapeV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BottomStartArcShapeV2Grids_reference = make(map[*BottomStartArcShapeV2Grid]*BottomStartArcShapeV2Grid)
	stage.BottomStartArcShapeV2Grids_referenceOrder = make(map[*BottomStartArcShapeV2Grid]uint) // diff Unstage needs the reference order
	stage.BottomStartArcShapeV2Grids_instance = make(map[*BottomStartArcShapeV2Grid]*BottomStartArcShapeV2Grid)
	for instance := range stage.BottomStartArcShapeV2Grids {
		_copy := instance.GongCopy().(*BottomStartArcShapeV2Grid)
		stage.BottomStartArcShapeV2Grids_reference[instance] = _copy
		stage.BottomStartArcShapeV2Grids_instance[_copy] = instance
		stage.BottomStartArcShapeV2Grids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.CircleGridShapes_reference = make(map[*CircleGridShape]*CircleGridShape)
	stage.CircleGridShapes_referenceOrder = make(map[*CircleGridShape]uint) // diff Unstage needs the reference order
	stage.CircleGridShapes_instance = make(map[*CircleGridShape]*CircleGridShape)
	for instance := range stage.CircleGridShapes {
		_copy := instance.GongCopy().(*CircleGridShape)
		stage.CircleGridShapes_reference[instance] = _copy
		stage.CircleGridShapes_instance[_copy] = instance
		stage.CircleGridShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EndArcShapes_reference = make(map[*EndArcShape]*EndArcShape)
	stage.EndArcShapes_referenceOrder = make(map[*EndArcShape]uint) // diff Unstage needs the reference order
	stage.EndArcShapes_instance = make(map[*EndArcShape]*EndArcShape)
	for instance := range stage.EndArcShapes {
		_copy := instance.GongCopy().(*EndArcShape)
		stage.EndArcShapes_reference[instance] = _copy
		stage.EndArcShapes_instance[_copy] = instance
		stage.EndArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EndArcShapeGrids_reference = make(map[*EndArcShapeGrid]*EndArcShapeGrid)
	stage.EndArcShapeGrids_referenceOrder = make(map[*EndArcShapeGrid]uint) // diff Unstage needs the reference order
	stage.EndArcShapeGrids_instance = make(map[*EndArcShapeGrid]*EndArcShapeGrid)
	for instance := range stage.EndArcShapeGrids {
		_copy := instance.GongCopy().(*EndArcShapeGrid)
		stage.EndArcShapeGrids_reference[instance] = _copy
		stage.EndArcShapeGrids_instance[_copy] = instance
		stage.EndArcShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EndArcShapeV2s_reference = make(map[*EndArcShapeV2]*EndArcShapeV2)
	stage.EndArcShapeV2s_referenceOrder = make(map[*EndArcShapeV2]uint) // diff Unstage needs the reference order
	stage.EndArcShapeV2s_instance = make(map[*EndArcShapeV2]*EndArcShapeV2)
	for instance := range stage.EndArcShapeV2s {
		_copy := instance.GongCopy().(*EndArcShapeV2)
		stage.EndArcShapeV2s_reference[instance] = _copy
		stage.EndArcShapeV2s_instance[_copy] = instance
		stage.EndArcShapeV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EndArcShapeV2Grids_reference = make(map[*EndArcShapeV2Grid]*EndArcShapeV2Grid)
	stage.EndArcShapeV2Grids_referenceOrder = make(map[*EndArcShapeV2Grid]uint) // diff Unstage needs the reference order
	stage.EndArcShapeV2Grids_instance = make(map[*EndArcShapeV2Grid]*EndArcShapeV2Grid)
	for instance := range stage.EndArcShapeV2Grids {
		_copy := instance.GongCopy().(*EndArcShapeV2Grid)
		stage.EndArcShapeV2Grids_reference[instance] = _copy
		stage.EndArcShapeV2Grids_instance[_copy] = instance
		stage.EndArcShapeV2Grids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ExplanationTextShapes_reference = make(map[*ExplanationTextShape]*ExplanationTextShape)
	stage.ExplanationTextShapes_referenceOrder = make(map[*ExplanationTextShape]uint) // diff Unstage needs the reference order
	stage.ExplanationTextShapes_instance = make(map[*ExplanationTextShape]*ExplanationTextShape)
	for instance := range stage.ExplanationTextShapes {
		_copy := instance.GongCopy().(*ExplanationTextShape)
		stage.ExplanationTextShapes_reference[instance] = _copy
		stage.ExplanationTextShapes_instance[_copy] = instance
		stage.ExplanationTextShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GridPathShapes_reference = make(map[*GridPathShape]*GridPathShape)
	stage.GridPathShapes_referenceOrder = make(map[*GridPathShape]uint) // diff Unstage needs the reference order
	stage.GridPathShapes_instance = make(map[*GridPathShape]*GridPathShape)
	for instance := range stage.GridPathShapes {
		_copy := instance.GongCopy().(*GridPathShape)
		stage.GridPathShapes_reference[instance] = _copy
		stage.GridPathShapes_instance[_copy] = instance
		stage.GridPathShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GrowthCurveBezierShapes_reference = make(map[*GrowthCurveBezierShape]*GrowthCurveBezierShape)
	stage.GrowthCurveBezierShapes_referenceOrder = make(map[*GrowthCurveBezierShape]uint) // diff Unstage needs the reference order
	stage.GrowthCurveBezierShapes_instance = make(map[*GrowthCurveBezierShape]*GrowthCurveBezierShape)
	for instance := range stage.GrowthCurveBezierShapes {
		_copy := instance.GongCopy().(*GrowthCurveBezierShape)
		stage.GrowthCurveBezierShapes_reference[instance] = _copy
		stage.GrowthCurveBezierShapes_instance[_copy] = instance
		stage.GrowthCurveBezierShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GrowthCurveBezierShapeGrids_reference = make(map[*GrowthCurveBezierShapeGrid]*GrowthCurveBezierShapeGrid)
	stage.GrowthCurveBezierShapeGrids_referenceOrder = make(map[*GrowthCurveBezierShapeGrid]uint) // diff Unstage needs the reference order
	stage.GrowthCurveBezierShapeGrids_instance = make(map[*GrowthCurveBezierShapeGrid]*GrowthCurveBezierShapeGrid)
	for instance := range stage.GrowthCurveBezierShapeGrids {
		_copy := instance.GongCopy().(*GrowthCurveBezierShapeGrid)
		stage.GrowthCurveBezierShapeGrids_reference[instance] = _copy
		stage.GrowthCurveBezierShapeGrids_instance[_copy] = instance
		stage.GrowthCurveBezierShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GrowthCurveRhombusGridShapes_reference = make(map[*GrowthCurveRhombusGridShape]*GrowthCurveRhombusGridShape)
	stage.GrowthCurveRhombusGridShapes_referenceOrder = make(map[*GrowthCurveRhombusGridShape]uint) // diff Unstage needs the reference order
	stage.GrowthCurveRhombusGridShapes_instance = make(map[*GrowthCurveRhombusGridShape]*GrowthCurveRhombusGridShape)
	for instance := range stage.GrowthCurveRhombusGridShapes {
		_copy := instance.GongCopy().(*GrowthCurveRhombusGridShape)
		stage.GrowthCurveRhombusGridShapes_reference[instance] = _copy
		stage.GrowthCurveRhombusGridShapes_instance[_copy] = instance
		stage.GrowthCurveRhombusGridShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GrowthCurveRhombusShapes_reference = make(map[*GrowthCurveRhombusShape]*GrowthCurveRhombusShape)
	stage.GrowthCurveRhombusShapes_referenceOrder = make(map[*GrowthCurveRhombusShape]uint) // diff Unstage needs the reference order
	stage.GrowthCurveRhombusShapes_instance = make(map[*GrowthCurveRhombusShape]*GrowthCurveRhombusShape)
	for instance := range stage.GrowthCurveRhombusShapes {
		_copy := instance.GongCopy().(*GrowthCurveRhombusShape)
		stage.GrowthCurveRhombusShapes_reference[instance] = _copy
		stage.GrowthCurveRhombusShapes_instance[_copy] = instance
		stage.GrowthCurveRhombusShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GrowthVectorShapes_reference = make(map[*GrowthVectorShape]*GrowthVectorShape)
	stage.GrowthVectorShapes_referenceOrder = make(map[*GrowthVectorShape]uint) // diff Unstage needs the reference order
	stage.GrowthVectorShapes_instance = make(map[*GrowthVectorShape]*GrowthVectorShape)
	for instance := range stage.GrowthVectorShapes {
		_copy := instance.GongCopy().(*GrowthVectorShape)
		stage.GrowthVectorShapes_reference[instance] = _copy
		stage.GrowthVectorShapes_instance[_copy] = instance
		stage.GrowthVectorShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.InitialRhombusGridShapes_reference = make(map[*InitialRhombusGridShape]*InitialRhombusGridShape)
	stage.InitialRhombusGridShapes_referenceOrder = make(map[*InitialRhombusGridShape]uint) // diff Unstage needs the reference order
	stage.InitialRhombusGridShapes_instance = make(map[*InitialRhombusGridShape]*InitialRhombusGridShape)
	for instance := range stage.InitialRhombusGridShapes {
		_copy := instance.GongCopy().(*InitialRhombusGridShape)
		stage.InitialRhombusGridShapes_reference[instance] = _copy
		stage.InitialRhombusGridShapes_instance[_copy] = instance
		stage.InitialRhombusGridShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.InitialRhombusShapes_reference = make(map[*InitialRhombusShape]*InitialRhombusShape)
	stage.InitialRhombusShapes_referenceOrder = make(map[*InitialRhombusShape]uint) // diff Unstage needs the reference order
	stage.InitialRhombusShapes_instance = make(map[*InitialRhombusShape]*InitialRhombusShape)
	for instance := range stage.InitialRhombusShapes {
		_copy := instance.GongCopy().(*InitialRhombusShape)
		stage.InitialRhombusShapes_reference[instance] = _copy
		stage.InitialRhombusShapes_instance[_copy] = instance
		stage.InitialRhombusShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint) // diff Unstage needs the reference order
	stage.Librarys_instance = make(map[*Library]*Library)
	for instance := range stage.Librarys {
		_copy := instance.GongCopy().(*Library)
		stage.Librarys_reference[instance] = _copy
		stage.Librarys_instance[_copy] = instance
		stage.Librarys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NextCircleShapes_reference = make(map[*NextCircleShape]*NextCircleShape)
	stage.NextCircleShapes_referenceOrder = make(map[*NextCircleShape]uint) // diff Unstage needs the reference order
	stage.NextCircleShapes_instance = make(map[*NextCircleShape]*NextCircleShape)
	for instance := range stage.NextCircleShapes {
		_copy := instance.GongCopy().(*NextCircleShape)
		stage.NextCircleShapes_reference[instance] = _copy
		stage.NextCircleShapes_instance[_copy] = instance
		stage.NextCircleShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PerpendicularVectors_reference = make(map[*PerpendicularVector]*PerpendicularVector)
	stage.PerpendicularVectors_referenceOrder = make(map[*PerpendicularVector]uint) // diff Unstage needs the reference order
	stage.PerpendicularVectors_instance = make(map[*PerpendicularVector]*PerpendicularVector)
	for instance := range stage.PerpendicularVectors {
		_copy := instance.GongCopy().(*PerpendicularVector)
		stage.PerpendicularVectors_reference[instance] = _copy
		stage.PerpendicularVectors_instance[_copy] = instance
		stage.PerpendicularVectors_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PerpendicularVectorGrids_reference = make(map[*PerpendicularVectorGrid]*PerpendicularVectorGrid)
	stage.PerpendicularVectorGrids_referenceOrder = make(map[*PerpendicularVectorGrid]uint) // diff Unstage needs the reference order
	stage.PerpendicularVectorGrids_instance = make(map[*PerpendicularVectorGrid]*PerpendicularVectorGrid)
	for instance := range stage.PerpendicularVectorGrids {
		_copy := instance.GongCopy().(*PerpendicularVectorGrid)
		stage.PerpendicularVectorGrids_reference[instance] = _copy
		stage.PerpendicularVectorGrids_instance[_copy] = instance
		stage.PerpendicularVectorGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PerpendicularVectorGridHalfways_reference = make(map[*PerpendicularVectorGridHalfway]*PerpendicularVectorGridHalfway)
	stage.PerpendicularVectorGridHalfways_referenceOrder = make(map[*PerpendicularVectorGridHalfway]uint) // diff Unstage needs the reference order
	stage.PerpendicularVectorGridHalfways_instance = make(map[*PerpendicularVectorGridHalfway]*PerpendicularVectorGridHalfway)
	for instance := range stage.PerpendicularVectorGridHalfways {
		_copy := instance.GongCopy().(*PerpendicularVectorGridHalfway)
		stage.PerpendicularVectorGridHalfways_reference[instance] = _copy
		stage.PerpendicularVectorGridHalfways_instance[_copy] = instance
		stage.PerpendicularVectorGridHalfways_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PerpendicularVectorHalfways_reference = make(map[*PerpendicularVectorHalfway]*PerpendicularVectorHalfway)
	stage.PerpendicularVectorHalfways_referenceOrder = make(map[*PerpendicularVectorHalfway]uint) // diff Unstage needs the reference order
	stage.PerpendicularVectorHalfways_instance = make(map[*PerpendicularVectorHalfway]*PerpendicularVectorHalfway)
	for instance := range stage.PerpendicularVectorHalfways {
		_copy := instance.GongCopy().(*PerpendicularVectorHalfway)
		stage.PerpendicularVectorHalfways_reference[instance] = _copy
		stage.PerpendicularVectorHalfways_instance[_copy] = instance
		stage.PerpendicularVectorHalfways_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Plants_reference = make(map[*Plant]*Plant)
	stage.Plants_referenceOrder = make(map[*Plant]uint) // diff Unstage needs the reference order
	stage.Plants_instance = make(map[*Plant]*Plant)
	for instance := range stage.Plants {
		_copy := instance.GongCopy().(*Plant)
		stage.Plants_reference[instance] = _copy
		stage.Plants_instance[_copy] = instance
		stage.Plants_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PlantCircumferenceShapes_reference = make(map[*PlantCircumferenceShape]*PlantCircumferenceShape)
	stage.PlantCircumferenceShapes_referenceOrder = make(map[*PlantCircumferenceShape]uint) // diff Unstage needs the reference order
	stage.PlantCircumferenceShapes_instance = make(map[*PlantCircumferenceShape]*PlantCircumferenceShape)
	for instance := range stage.PlantCircumferenceShapes {
		_copy := instance.GongCopy().(*PlantCircumferenceShape)
		stage.PlantCircumferenceShapes_reference[instance] = _copy
		stage.PlantCircumferenceShapes_instance[_copy] = instance
		stage.PlantCircumferenceShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PlantDiagrams_reference = make(map[*PlantDiagram]*PlantDiagram)
	stage.PlantDiagrams_referenceOrder = make(map[*PlantDiagram]uint) // diff Unstage needs the reference order
	stage.PlantDiagrams_instance = make(map[*PlantDiagram]*PlantDiagram)
	for instance := range stage.PlantDiagrams {
		_copy := instance.GongCopy().(*PlantDiagram)
		stage.PlantDiagrams_reference[instance] = _copy
		stage.PlantDiagrams_instance[_copy] = instance
		stage.PlantDiagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Rendered3DShapes_reference = make(map[*Rendered3DShape]*Rendered3DShape)
	stage.Rendered3DShapes_referenceOrder = make(map[*Rendered3DShape]uint) // diff Unstage needs the reference order
	stage.Rendered3DShapes_instance = make(map[*Rendered3DShape]*Rendered3DShape)
	for instance := range stage.Rendered3DShapes {
		_copy := instance.GongCopy().(*Rendered3DShape)
		stage.Rendered3DShapes_reference[instance] = _copy
		stage.Rendered3DShapes_instance[_copy] = instance
		stage.Rendered3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RhombusShapes_reference = make(map[*RhombusShape]*RhombusShape)
	stage.RhombusShapes_referenceOrder = make(map[*RhombusShape]uint) // diff Unstage needs the reference order
	stage.RhombusShapes_instance = make(map[*RhombusShape]*RhombusShape)
	for instance := range stage.RhombusShapes {
		_copy := instance.GongCopy().(*RhombusShape)
		stage.RhombusShapes_reference[instance] = _copy
		stage.RhombusShapes_instance[_copy] = instance
		stage.RhombusShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RotatedRhombusGridShapes_reference = make(map[*RotatedRhombusGridShape]*RotatedRhombusGridShape)
	stage.RotatedRhombusGridShapes_referenceOrder = make(map[*RotatedRhombusGridShape]uint) // diff Unstage needs the reference order
	stage.RotatedRhombusGridShapes_instance = make(map[*RotatedRhombusGridShape]*RotatedRhombusGridShape)
	for instance := range stage.RotatedRhombusGridShapes {
		_copy := instance.GongCopy().(*RotatedRhombusGridShape)
		stage.RotatedRhombusGridShapes_reference[instance] = _copy
		stage.RotatedRhombusGridShapes_instance[_copy] = instance
		stage.RotatedRhombusGridShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RotatedRhombusShapes_reference = make(map[*RotatedRhombusShape]*RotatedRhombusShape)
	stage.RotatedRhombusShapes_referenceOrder = make(map[*RotatedRhombusShape]uint) // diff Unstage needs the reference order
	stage.RotatedRhombusShapes_instance = make(map[*RotatedRhombusShape]*RotatedRhombusShape)
	for instance := range stage.RotatedRhombusShapes {
		_copy := instance.GongCopy().(*RotatedRhombusShape)
		stage.RotatedRhombusShapes_reference[instance] = _copy
		stage.RotatedRhombusShapes_instance[_copy] = instance
		stage.RotatedRhombusShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackGrowthCurveBezierShapes_reference = make(map[*StackGrowthCurveBezierShape]*StackGrowthCurveBezierShape)
	stage.StackGrowthCurveBezierShapes_referenceOrder = make(map[*StackGrowthCurveBezierShape]uint) // diff Unstage needs the reference order
	stage.StackGrowthCurveBezierShapes_instance = make(map[*StackGrowthCurveBezierShape]*StackGrowthCurveBezierShape)
	for instance := range stage.StackGrowthCurveBezierShapes {
		_copy := instance.GongCopy().(*StackGrowthCurveBezierShape)
		stage.StackGrowthCurveBezierShapes_reference[instance] = _copy
		stage.StackGrowthCurveBezierShapes_instance[_copy] = instance
		stage.StackGrowthCurveBezierShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackGrowthCurveEndArcShapeV2s_reference = make(map[*StackGrowthCurveEndArcShapeV2]*StackGrowthCurveEndArcShapeV2)
	stage.StackGrowthCurveEndArcShapeV2s_referenceOrder = make(map[*StackGrowthCurveEndArcShapeV2]uint) // diff Unstage needs the reference order
	stage.StackGrowthCurveEndArcShapeV2s_instance = make(map[*StackGrowthCurveEndArcShapeV2]*StackGrowthCurveEndArcShapeV2)
	for instance := range stage.StackGrowthCurveEndArcShapeV2s {
		_copy := instance.GongCopy().(*StackGrowthCurveEndArcShapeV2)
		stage.StackGrowthCurveEndArcShapeV2s_reference[instance] = _copy
		stage.StackGrowthCurveEndArcShapeV2s_instance[_copy] = instance
		stage.StackGrowthCurveEndArcShapeV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackGrowthCurveStartArcShapeV2s_reference = make(map[*StackGrowthCurveStartArcShapeV2]*StackGrowthCurveStartArcShapeV2)
	stage.StackGrowthCurveStartArcShapeV2s_referenceOrder = make(map[*StackGrowthCurveStartArcShapeV2]uint) // diff Unstage needs the reference order
	stage.StackGrowthCurveStartArcShapeV2s_instance = make(map[*StackGrowthCurveStartArcShapeV2]*StackGrowthCurveStartArcShapeV2)
	for instance := range stage.StackGrowthCurveStartArcShapeV2s {
		_copy := instance.GongCopy().(*StackGrowthCurveStartArcShapeV2)
		stage.StackGrowthCurveStartArcShapeV2s_reference[instance] = _copy
		stage.StackGrowthCurveStartArcShapeV2s_instance[_copy] = instance
		stage.StackGrowthCurveStartArcShapeV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackOfGrowthCurves_reference = make(map[*StackOfGrowthCurve]*StackOfGrowthCurve)
	stage.StackOfGrowthCurves_referenceOrder = make(map[*StackOfGrowthCurve]uint) // diff Unstage needs the reference order
	stage.StackOfGrowthCurves_instance = make(map[*StackOfGrowthCurve]*StackOfGrowthCurve)
	for instance := range stage.StackOfGrowthCurves {
		_copy := instance.GongCopy().(*StackOfGrowthCurve)
		stage.StackOfGrowthCurves_reference[instance] = _copy
		stage.StackOfGrowthCurves_instance[_copy] = instance
		stage.StackOfGrowthCurves_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackOfGrowthCurveV2s_reference = make(map[*StackOfGrowthCurveV2]*StackOfGrowthCurveV2)
	stage.StackOfGrowthCurveV2s_referenceOrder = make(map[*StackOfGrowthCurveV2]uint) // diff Unstage needs the reference order
	stage.StackOfGrowthCurveV2s_instance = make(map[*StackOfGrowthCurveV2]*StackOfGrowthCurveV2)
	for instance := range stage.StackOfGrowthCurveV2s {
		_copy := instance.GongCopy().(*StackOfGrowthCurveV2)
		stage.StackOfGrowthCurveV2s_reference[instance] = _copy
		stage.StackOfGrowthCurveV2s_instance[_copy] = instance
		stage.StackOfGrowthCurveV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StartArcShapes_reference = make(map[*StartArcShape]*StartArcShape)
	stage.StartArcShapes_referenceOrder = make(map[*StartArcShape]uint) // diff Unstage needs the reference order
	stage.StartArcShapes_instance = make(map[*StartArcShape]*StartArcShape)
	for instance := range stage.StartArcShapes {
		_copy := instance.GongCopy().(*StartArcShape)
		stage.StartArcShapes_reference[instance] = _copy
		stage.StartArcShapes_instance[_copy] = instance
		stage.StartArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StartArcShapeGrids_reference = make(map[*StartArcShapeGrid]*StartArcShapeGrid)
	stage.StartArcShapeGrids_referenceOrder = make(map[*StartArcShapeGrid]uint) // diff Unstage needs the reference order
	stage.StartArcShapeGrids_instance = make(map[*StartArcShapeGrid]*StartArcShapeGrid)
	for instance := range stage.StartArcShapeGrids {
		_copy := instance.GongCopy().(*StartArcShapeGrid)
		stage.StartArcShapeGrids_reference[instance] = _copy
		stage.StartArcShapeGrids_instance[_copy] = instance
		stage.StartArcShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StartArcShapeV2s_reference = make(map[*StartArcShapeV2]*StartArcShapeV2)
	stage.StartArcShapeV2s_referenceOrder = make(map[*StartArcShapeV2]uint) // diff Unstage needs the reference order
	stage.StartArcShapeV2s_instance = make(map[*StartArcShapeV2]*StartArcShapeV2)
	for instance := range stage.StartArcShapeV2s {
		_copy := instance.GongCopy().(*StartArcShapeV2)
		stage.StartArcShapeV2s_reference[instance] = _copy
		stage.StartArcShapeV2s_instance[_copy] = instance
		stage.StartArcShapeV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StartArcShapeV2Grids_reference = make(map[*StartArcShapeV2Grid]*StartArcShapeV2Grid)
	stage.StartArcShapeV2Grids_referenceOrder = make(map[*StartArcShapeV2Grid]uint) // diff Unstage needs the reference order
	stage.StartArcShapeV2Grids_instance = make(map[*StartArcShapeV2Grid]*StartArcShapeV2Grid)
	for instance := range stage.StartArcShapeV2Grids {
		_copy := instance.GongCopy().(*StartArcShapeV2Grid)
		stage.StartArcShapeV2Grids_reference[instance] = _copy
		stage.StartArcShapeV2Grids_instance[_copy] = instance
		stage.StartArcShapeV2Grids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopEndArcShapeV2s_reference = make(map[*TopEndArcShapeV2]*TopEndArcShapeV2)
	stage.TopEndArcShapeV2s_referenceOrder = make(map[*TopEndArcShapeV2]uint) // diff Unstage needs the reference order
	stage.TopEndArcShapeV2s_instance = make(map[*TopEndArcShapeV2]*TopEndArcShapeV2)
	for instance := range stage.TopEndArcShapeV2s {
		_copy := instance.GongCopy().(*TopEndArcShapeV2)
		stage.TopEndArcShapeV2s_reference[instance] = _copy
		stage.TopEndArcShapeV2s_instance[_copy] = instance
		stage.TopEndArcShapeV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopEndArcShapeV2Grids_reference = make(map[*TopEndArcShapeV2Grid]*TopEndArcShapeV2Grid)
	stage.TopEndArcShapeV2Grids_referenceOrder = make(map[*TopEndArcShapeV2Grid]uint) // diff Unstage needs the reference order
	stage.TopEndArcShapeV2Grids_instance = make(map[*TopEndArcShapeV2Grid]*TopEndArcShapeV2Grid)
	for instance := range stage.TopEndArcShapeV2Grids {
		_copy := instance.GongCopy().(*TopEndArcShapeV2Grid)
		stage.TopEndArcShapeV2Grids_reference[instance] = _copy
		stage.TopEndArcShapeV2Grids_instance[_copy] = instance
		stage.TopEndArcShapeV2Grids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStackGrowthCurveEndArcShapeV2s_reference = make(map[*TopStackGrowthCurveEndArcShapeV2]*TopStackGrowthCurveEndArcShapeV2)
	stage.TopStackGrowthCurveEndArcShapeV2s_referenceOrder = make(map[*TopStackGrowthCurveEndArcShapeV2]uint) // diff Unstage needs the reference order
	stage.TopStackGrowthCurveEndArcShapeV2s_instance = make(map[*TopStackGrowthCurveEndArcShapeV2]*TopStackGrowthCurveEndArcShapeV2)
	for instance := range stage.TopStackGrowthCurveEndArcShapeV2s {
		_copy := instance.GongCopy().(*TopStackGrowthCurveEndArcShapeV2)
		stage.TopStackGrowthCurveEndArcShapeV2s_reference[instance] = _copy
		stage.TopStackGrowthCurveEndArcShapeV2s_instance[_copy] = instance
		stage.TopStackGrowthCurveEndArcShapeV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStackGrowthCurveStartArcShapeV2s_reference = make(map[*TopStackGrowthCurveStartArcShapeV2]*TopStackGrowthCurveStartArcShapeV2)
	stage.TopStackGrowthCurveStartArcShapeV2s_referenceOrder = make(map[*TopStackGrowthCurveStartArcShapeV2]uint) // diff Unstage needs the reference order
	stage.TopStackGrowthCurveStartArcShapeV2s_instance = make(map[*TopStackGrowthCurveStartArcShapeV2]*TopStackGrowthCurveStartArcShapeV2)
	for instance := range stage.TopStackGrowthCurveStartArcShapeV2s {
		_copy := instance.GongCopy().(*TopStackGrowthCurveStartArcShapeV2)
		stage.TopStackGrowthCurveStartArcShapeV2s_reference[instance] = _copy
		stage.TopStackGrowthCurveStartArcShapeV2s_instance[_copy] = instance
		stage.TopStackGrowthCurveStartArcShapeV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStackOfGrowthCurveV2s_reference = make(map[*TopStackOfGrowthCurveV2]*TopStackOfGrowthCurveV2)
	stage.TopStackOfGrowthCurveV2s_referenceOrder = make(map[*TopStackOfGrowthCurveV2]uint) // diff Unstage needs the reference order
	stage.TopStackOfGrowthCurveV2s_instance = make(map[*TopStackOfGrowthCurveV2]*TopStackOfGrowthCurveV2)
	for instance := range stage.TopStackOfGrowthCurveV2s {
		_copy := instance.GongCopy().(*TopStackOfGrowthCurveV2)
		stage.TopStackOfGrowthCurveV2s_reference[instance] = _copy
		stage.TopStackOfGrowthCurveV2s_instance[_copy] = instance
		stage.TopStackOfGrowthCurveV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStartArcShapeV2s_reference = make(map[*TopStartArcShapeV2]*TopStartArcShapeV2)
	stage.TopStartArcShapeV2s_referenceOrder = make(map[*TopStartArcShapeV2]uint) // diff Unstage needs the reference order
	stage.TopStartArcShapeV2s_instance = make(map[*TopStartArcShapeV2]*TopStartArcShapeV2)
	for instance := range stage.TopStartArcShapeV2s {
		_copy := instance.GongCopy().(*TopStartArcShapeV2)
		stage.TopStartArcShapeV2s_reference[instance] = _copy
		stage.TopStartArcShapeV2s_instance[_copy] = instance
		stage.TopStartArcShapeV2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStartArcShapeV2Grids_reference = make(map[*TopStartArcShapeV2Grid]*TopStartArcShapeV2Grid)
	stage.TopStartArcShapeV2Grids_referenceOrder = make(map[*TopStartArcShapeV2Grid]uint) // diff Unstage needs the reference order
	stage.TopStartArcShapeV2Grids_instance = make(map[*TopStartArcShapeV2Grid]*TopStartArcShapeV2Grid)
	for instance := range stage.TopStartArcShapeV2Grids {
		_copy := instance.GongCopy().(*TopStartArcShapeV2Grid)
		stage.TopStartArcShapeV2Grids_reference[instance] = _copy
		stage.TopStartArcShapeV2Grids_instance[_copy] = instance
		stage.TopStartArcShapeV2Grids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.ArcNormalVectorShapes {
		reference := stage.ArcNormalVectorShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ArcNormalVectorShapeGrids {
		reference := stage.ArcNormalVectorShapeGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.AxesShapes {
		reference := stage.AxesShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BaseVectorShapes {
		reference := stage.BaseVectorShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BaseVectorShapeGrids {
		reference := stage.BaseVectorShapeGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BottomEndArcShapeV2s {
		reference := stage.BottomEndArcShapeV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BottomEndArcShapeV2Grids {
		reference := stage.BottomEndArcShapeV2Grids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BottomStackGrowthCurveEndArcShapeV2s {
		reference := stage.BottomStackGrowthCurveEndArcShapeV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BottomStackGrowthCurveStartArcShapeV2s {
		reference := stage.BottomStackGrowthCurveStartArcShapeV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BottomStackOfGrowthCurveV2s {
		reference := stage.BottomStackOfGrowthCurveV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BottomStartArcShapeV2s {
		reference := stage.BottomStartArcShapeV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BottomStartArcShapeV2Grids {
		reference := stage.BottomStartArcShapeV2Grids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.CircleGridShapes {
		reference := stage.CircleGridShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EndArcShapes {
		reference := stage.EndArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EndArcShapeGrids {
		reference := stage.EndArcShapeGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EndArcShapeV2s {
		reference := stage.EndArcShapeV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EndArcShapeV2Grids {
		reference := stage.EndArcShapeV2Grids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ExplanationTextShapes {
		reference := stage.ExplanationTextShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GridPathShapes {
		reference := stage.GridPathShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GrowthCurveBezierShapes {
		reference := stage.GrowthCurveBezierShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GrowthCurveBezierShapeGrids {
		reference := stage.GrowthCurveBezierShapeGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GrowthCurveRhombusGridShapes {
		reference := stage.GrowthCurveRhombusGridShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GrowthCurveRhombusShapes {
		reference := stage.GrowthCurveRhombusShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GrowthVectorShapes {
		reference := stage.GrowthVectorShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.InitialRhombusGridShapes {
		reference := stage.InitialRhombusGridShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.InitialRhombusShapes {
		reference := stage.InitialRhombusShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Librarys {
		reference := stage.Librarys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NextCircleShapes {
		reference := stage.NextCircleShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PerpendicularVectors {
		reference := stage.PerpendicularVectors_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PerpendicularVectorGrids {
		reference := stage.PerpendicularVectorGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PerpendicularVectorGridHalfways {
		reference := stage.PerpendicularVectorGridHalfways_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PerpendicularVectorHalfways {
		reference := stage.PerpendicularVectorHalfways_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Plants {
		reference := stage.Plants_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PlantCircumferenceShapes {
		reference := stage.PlantCircumferenceShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PlantDiagrams {
		reference := stage.PlantDiagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Rendered3DShapes {
		reference := stage.Rendered3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RhombusShapes {
		reference := stage.RhombusShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RotatedRhombusGridShapes {
		reference := stage.RotatedRhombusGridShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RotatedRhombusShapes {
		reference := stage.RotatedRhombusShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackGrowthCurveBezierShapes {
		reference := stage.StackGrowthCurveBezierShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackGrowthCurveEndArcShapeV2s {
		reference := stage.StackGrowthCurveEndArcShapeV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackGrowthCurveStartArcShapeV2s {
		reference := stage.StackGrowthCurveStartArcShapeV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackOfGrowthCurves {
		reference := stage.StackOfGrowthCurves_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackOfGrowthCurveV2s {
		reference := stage.StackOfGrowthCurveV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StartArcShapes {
		reference := stage.StartArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StartArcShapeGrids {
		reference := stage.StartArcShapeGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StartArcShapeV2s {
		reference := stage.StartArcShapeV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StartArcShapeV2Grids {
		reference := stage.StartArcShapeV2Grids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopEndArcShapeV2s {
		reference := stage.TopEndArcShapeV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopEndArcShapeV2Grids {
		reference := stage.TopEndArcShapeV2Grids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStackGrowthCurveEndArcShapeV2s {
		reference := stage.TopStackGrowthCurveEndArcShapeV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStackGrowthCurveStartArcShapeV2s {
		reference := stage.TopStackGrowthCurveStartArcShapeV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStackOfGrowthCurveV2s {
		reference := stage.TopStackOfGrowthCurveV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStartArcShapeV2s {
		reference := stage.TopStartArcShapeV2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStartArcShapeV2Grids {
		reference := stage.TopStartArcShapeV2Grids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (arcnormalvectorshape *ArcNormalVectorShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ArcNormalVectorShape_stagedOrder[arcnormalvectorshape]; ok {
		return order
	}
	if order, ok := stage.ArcNormalVectorShapes_referenceOrder[arcnormalvectorshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ArcNormalVectorShape was not staged and does not have a reference order", arcnormalvectorshape)
		return 0
	}
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ArcNormalVectorShapeGrid_stagedOrder[arcnormalvectorshapegrid]; ok {
		return order
	}
	if order, ok := stage.ArcNormalVectorShapeGrids_referenceOrder[arcnormalvectorshapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type ArcNormalVectorShapeGrid was not staged and does not have a reference order", arcnormalvectorshapegrid)
		return 0
	}
}

func (axesshape *AxesShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AxesShape_stagedOrder[axesshape]; ok {
		return order
	}
	if order, ok := stage.AxesShapes_referenceOrder[axesshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type AxesShape was not staged and does not have a reference order", axesshape)
		return 0
	}
}

func (basevectorshape *BaseVectorShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BaseVectorShape_stagedOrder[basevectorshape]; ok {
		return order
	}
	if order, ok := stage.BaseVectorShapes_referenceOrder[basevectorshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type BaseVectorShape was not staged and does not have a reference order", basevectorshape)
		return 0
	}
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BaseVectorShapeGrid_stagedOrder[basevectorshapegrid]; ok {
		return order
	}
	if order, ok := stage.BaseVectorShapeGrids_referenceOrder[basevectorshapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type BaseVectorShapeGrid was not staged and does not have a reference order", basevectorshapegrid)
		return 0
	}
}

func (bottomendarcshapev2 *BottomEndArcShapeV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BottomEndArcShapeV2_stagedOrder[bottomendarcshapev2]; ok {
		return order
	}
	if order, ok := stage.BottomEndArcShapeV2s_referenceOrder[bottomendarcshapev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type BottomEndArcShapeV2 was not staged and does not have a reference order", bottomendarcshapev2)
		return 0
	}
}

func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BottomEndArcShapeV2Grid_stagedOrder[bottomendarcshapev2grid]; ok {
		return order
	}
	if order, ok := stage.BottomEndArcShapeV2Grids_referenceOrder[bottomendarcshapev2grid]; ok {
		return order
	} else {
		log.Printf("instance %p of type BottomEndArcShapeV2Grid was not staged and does not have a reference order", bottomendarcshapev2grid)
		return 0
	}
}

func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BottomStackGrowthCurveEndArcShapeV2_stagedOrder[bottomstackgrowthcurveendarcshapev2]; ok {
		return order
	}
	if order, ok := stage.BottomStackGrowthCurveEndArcShapeV2s_referenceOrder[bottomstackgrowthcurveendarcshapev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type BottomStackGrowthCurveEndArcShapeV2 was not staged and does not have a reference order", bottomstackgrowthcurveendarcshapev2)
		return 0
	}
}

func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BottomStackGrowthCurveStartArcShapeV2_stagedOrder[bottomstackgrowthcurvestartarcshapev2]; ok {
		return order
	}
	if order, ok := stage.BottomStackGrowthCurveStartArcShapeV2s_referenceOrder[bottomstackgrowthcurvestartarcshapev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type BottomStackGrowthCurveStartArcShapeV2 was not staged and does not have a reference order", bottomstackgrowthcurvestartarcshapev2)
		return 0
	}
}

func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BottomStackOfGrowthCurveV2_stagedOrder[bottomstackofgrowthcurvev2]; ok {
		return order
	}
	if order, ok := stage.BottomStackOfGrowthCurveV2s_referenceOrder[bottomstackofgrowthcurvev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type BottomStackOfGrowthCurveV2 was not staged and does not have a reference order", bottomstackofgrowthcurvev2)
		return 0
	}
}

func (bottomstartarcshapev2 *BottomStartArcShapeV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BottomStartArcShapeV2_stagedOrder[bottomstartarcshapev2]; ok {
		return order
	}
	if order, ok := stage.BottomStartArcShapeV2s_referenceOrder[bottomstartarcshapev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type BottomStartArcShapeV2 was not staged and does not have a reference order", bottomstartarcshapev2)
		return 0
	}
}

func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BottomStartArcShapeV2Grid_stagedOrder[bottomstartarcshapev2grid]; ok {
		return order
	}
	if order, ok := stage.BottomStartArcShapeV2Grids_referenceOrder[bottomstartarcshapev2grid]; ok {
		return order
	} else {
		log.Printf("instance %p of type BottomStartArcShapeV2Grid was not staged and does not have a reference order", bottomstartarcshapev2grid)
		return 0
	}
}

func (circlegridshape *CircleGridShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CircleGridShape_stagedOrder[circlegridshape]; ok {
		return order
	}
	if order, ok := stage.CircleGridShapes_referenceOrder[circlegridshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type CircleGridShape was not staged and does not have a reference order", circlegridshape)
		return 0
	}
}

func (endarcshape *EndArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EndArcShape_stagedOrder[endarcshape]; ok {
		return order
	}
	if order, ok := stage.EndArcShapes_referenceOrder[endarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type EndArcShape was not staged and does not have a reference order", endarcshape)
		return 0
	}
}

func (endarcshapegrid *EndArcShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EndArcShapeGrid_stagedOrder[endarcshapegrid]; ok {
		return order
	}
	if order, ok := stage.EndArcShapeGrids_referenceOrder[endarcshapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type EndArcShapeGrid was not staged and does not have a reference order", endarcshapegrid)
		return 0
	}
}

func (endarcshapev2 *EndArcShapeV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EndArcShapeV2_stagedOrder[endarcshapev2]; ok {
		return order
	}
	if order, ok := stage.EndArcShapeV2s_referenceOrder[endarcshapev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type EndArcShapeV2 was not staged and does not have a reference order", endarcshapev2)
		return 0
	}
}

func (endarcshapev2grid *EndArcShapeV2Grid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EndArcShapeV2Grid_stagedOrder[endarcshapev2grid]; ok {
		return order
	}
	if order, ok := stage.EndArcShapeV2Grids_referenceOrder[endarcshapev2grid]; ok {
		return order
	} else {
		log.Printf("instance %p of type EndArcShapeV2Grid was not staged and does not have a reference order", endarcshapev2grid)
		return 0
	}
}

func (explanationtextshape *ExplanationTextShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ExplanationTextShape_stagedOrder[explanationtextshape]; ok {
		return order
	}
	if order, ok := stage.ExplanationTextShapes_referenceOrder[explanationtextshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ExplanationTextShape was not staged and does not have a reference order", explanationtextshape)
		return 0
	}
}

func (gridpathshape *GridPathShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GridPathShape_stagedOrder[gridpathshape]; ok {
		return order
	}
	if order, ok := stage.GridPathShapes_referenceOrder[gridpathshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type GridPathShape was not staged and does not have a reference order", gridpathshape)
		return 0
	}
}

func (growthcurvebeziershape *GrowthCurveBezierShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GrowthCurveBezierShape_stagedOrder[growthcurvebeziershape]; ok {
		return order
	}
	if order, ok := stage.GrowthCurveBezierShapes_referenceOrder[growthcurvebeziershape]; ok {
		return order
	} else {
		log.Printf("instance %p of type GrowthCurveBezierShape was not staged and does not have a reference order", growthcurvebeziershape)
		return 0
	}
}

func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GrowthCurveBezierShapeGrid_stagedOrder[growthcurvebeziershapegrid]; ok {
		return order
	}
	if order, ok := stage.GrowthCurveBezierShapeGrids_referenceOrder[growthcurvebeziershapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type GrowthCurveBezierShapeGrid was not staged and does not have a reference order", growthcurvebeziershapegrid)
		return 0
	}
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GrowthCurveRhombusGridShape_stagedOrder[growthcurverhombusgridshape]; ok {
		return order
	}
	if order, ok := stage.GrowthCurveRhombusGridShapes_referenceOrder[growthcurverhombusgridshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type GrowthCurveRhombusGridShape was not staged and does not have a reference order", growthcurverhombusgridshape)
		return 0
	}
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GrowthCurveRhombusShape_stagedOrder[growthcurverhombusshape]; ok {
		return order
	}
	if order, ok := stage.GrowthCurveRhombusShapes_referenceOrder[growthcurverhombusshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type GrowthCurveRhombusShape was not staged and does not have a reference order", growthcurverhombusshape)
		return 0
	}
}

func (growthvectorshape *GrowthVectorShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GrowthVectorShape_stagedOrder[growthvectorshape]; ok {
		return order
	}
	if order, ok := stage.GrowthVectorShapes_referenceOrder[growthvectorshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type GrowthVectorShape was not staged and does not have a reference order", growthvectorshape)
		return 0
	}
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.InitialRhombusGridShape_stagedOrder[initialrhombusgridshape]; ok {
		return order
	}
	if order, ok := stage.InitialRhombusGridShapes_referenceOrder[initialrhombusgridshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type InitialRhombusGridShape was not staged and does not have a reference order", initialrhombusgridshape)
		return 0
	}
}

func (initialrhombusshape *InitialRhombusShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.InitialRhombusShape_stagedOrder[initialrhombusshape]; ok {
		return order
	}
	if order, ok := stage.InitialRhombusShapes_referenceOrder[initialrhombusshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type InitialRhombusShape was not staged and does not have a reference order", initialrhombusshape)
		return 0
	}
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Library_stagedOrder[library]; ok {
		return order
	}
	if order, ok := stage.Librarys_referenceOrder[library]; ok {
		return order
	} else {
		log.Printf("instance %p of type Library was not staged and does not have a reference order", library)
		return 0
	}
}

func (nextcircleshape *NextCircleShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NextCircleShape_stagedOrder[nextcircleshape]; ok {
		return order
	}
	if order, ok := stage.NextCircleShapes_referenceOrder[nextcircleshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NextCircleShape was not staged and does not have a reference order", nextcircleshape)
		return 0
	}
}

func (perpendicularvector *PerpendicularVector) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PerpendicularVector_stagedOrder[perpendicularvector]; ok {
		return order
	}
	if order, ok := stage.PerpendicularVectors_referenceOrder[perpendicularvector]; ok {
		return order
	} else {
		log.Printf("instance %p of type PerpendicularVector was not staged and does not have a reference order", perpendicularvector)
		return 0
	}
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PerpendicularVectorGrid_stagedOrder[perpendicularvectorgrid]; ok {
		return order
	}
	if order, ok := stage.PerpendicularVectorGrids_referenceOrder[perpendicularvectorgrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type PerpendicularVectorGrid was not staged and does not have a reference order", perpendicularvectorgrid)
		return 0
	}
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PerpendicularVectorGridHalfway_stagedOrder[perpendicularvectorgridhalfway]; ok {
		return order
	}
	if order, ok := stage.PerpendicularVectorGridHalfways_referenceOrder[perpendicularvectorgridhalfway]; ok {
		return order
	} else {
		log.Printf("instance %p of type PerpendicularVectorGridHalfway was not staged and does not have a reference order", perpendicularvectorgridhalfway)
		return 0
	}
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PerpendicularVectorHalfway_stagedOrder[perpendicularvectorhalfway]; ok {
		return order
	}
	if order, ok := stage.PerpendicularVectorHalfways_referenceOrder[perpendicularvectorhalfway]; ok {
		return order
	} else {
		log.Printf("instance %p of type PerpendicularVectorHalfway was not staged and does not have a reference order", perpendicularvectorhalfway)
		return 0
	}
}

func (plant *Plant) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Plant_stagedOrder[plant]; ok {
		return order
	}
	if order, ok := stage.Plants_referenceOrder[plant]; ok {
		return order
	} else {
		log.Printf("instance %p of type Plant was not staged and does not have a reference order", plant)
		return 0
	}
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PlantCircumferenceShape_stagedOrder[plantcircumferenceshape]; ok {
		return order
	}
	if order, ok := stage.PlantCircumferenceShapes_referenceOrder[plantcircumferenceshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PlantCircumferenceShape was not staged and does not have a reference order", plantcircumferenceshape)
		return 0
	}
}

func (plantdiagram *PlantDiagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PlantDiagram_stagedOrder[plantdiagram]; ok {
		return order
	}
	if order, ok := stage.PlantDiagrams_referenceOrder[plantdiagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type PlantDiagram was not staged and does not have a reference order", plantdiagram)
		return 0
	}
}

func (rendered3dshape *Rendered3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Rendered3DShape_stagedOrder[rendered3dshape]; ok {
		return order
	}
	if order, ok := stage.Rendered3DShapes_referenceOrder[rendered3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type Rendered3DShape was not staged and does not have a reference order", rendered3dshape)
		return 0
	}
}

func (rhombusshape *RhombusShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RhombusShape_stagedOrder[rhombusshape]; ok {
		return order
	}
	if order, ok := stage.RhombusShapes_referenceOrder[rhombusshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type RhombusShape was not staged and does not have a reference order", rhombusshape)
		return 0
	}
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RotatedRhombusGridShape_stagedOrder[rotatedrhombusgridshape]; ok {
		return order
	}
	if order, ok := stage.RotatedRhombusGridShapes_referenceOrder[rotatedrhombusgridshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type RotatedRhombusGridShape was not staged and does not have a reference order", rotatedrhombusgridshape)
		return 0
	}
}

func (rotatedrhombusshape *RotatedRhombusShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RotatedRhombusShape_stagedOrder[rotatedrhombusshape]; ok {
		return order
	}
	if order, ok := stage.RotatedRhombusShapes_referenceOrder[rotatedrhombusshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type RotatedRhombusShape was not staged and does not have a reference order", rotatedrhombusshape)
		return 0
	}
}

func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackGrowthCurveBezierShape_stagedOrder[stackgrowthcurvebeziershape]; ok {
		return order
	}
	if order, ok := stage.StackGrowthCurveBezierShapes_referenceOrder[stackgrowthcurvebeziershape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackGrowthCurveBezierShape was not staged and does not have a reference order", stackgrowthcurvebeziershape)
		return 0
	}
}

func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackGrowthCurveEndArcShapeV2_stagedOrder[stackgrowthcurveendarcshapev2]; ok {
		return order
	}
	if order, ok := stage.StackGrowthCurveEndArcShapeV2s_referenceOrder[stackgrowthcurveendarcshapev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackGrowthCurveEndArcShapeV2 was not staged and does not have a reference order", stackgrowthcurveendarcshapev2)
		return 0
	}
}

func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackGrowthCurveStartArcShapeV2_stagedOrder[stackgrowthcurvestartarcshapev2]; ok {
		return order
	}
	if order, ok := stage.StackGrowthCurveStartArcShapeV2s_referenceOrder[stackgrowthcurvestartarcshapev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackGrowthCurveStartArcShapeV2 was not staged and does not have a reference order", stackgrowthcurvestartarcshapev2)
		return 0
	}
}

func (stackofgrowthcurve *StackOfGrowthCurve) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackOfGrowthCurve_stagedOrder[stackofgrowthcurve]; ok {
		return order
	}
	if order, ok := stage.StackOfGrowthCurves_referenceOrder[stackofgrowthcurve]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackOfGrowthCurve was not staged and does not have a reference order", stackofgrowthcurve)
		return 0
	}
}

func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackOfGrowthCurveV2_stagedOrder[stackofgrowthcurvev2]; ok {
		return order
	}
	if order, ok := stage.StackOfGrowthCurveV2s_referenceOrder[stackofgrowthcurvev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackOfGrowthCurveV2 was not staged and does not have a reference order", stackofgrowthcurvev2)
		return 0
	}
}

func (startarcshape *StartArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StartArcShape_stagedOrder[startarcshape]; ok {
		return order
	}
	if order, ok := stage.StartArcShapes_referenceOrder[startarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StartArcShape was not staged and does not have a reference order", startarcshape)
		return 0
	}
}

func (startarcshapegrid *StartArcShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StartArcShapeGrid_stagedOrder[startarcshapegrid]; ok {
		return order
	}
	if order, ok := stage.StartArcShapeGrids_referenceOrder[startarcshapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type StartArcShapeGrid was not staged and does not have a reference order", startarcshapegrid)
		return 0
	}
}

func (startarcshapev2 *StartArcShapeV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StartArcShapeV2_stagedOrder[startarcshapev2]; ok {
		return order
	}
	if order, ok := stage.StartArcShapeV2s_referenceOrder[startarcshapev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type StartArcShapeV2 was not staged and does not have a reference order", startarcshapev2)
		return 0
	}
}

func (startarcshapev2grid *StartArcShapeV2Grid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StartArcShapeV2Grid_stagedOrder[startarcshapev2grid]; ok {
		return order
	}
	if order, ok := stage.StartArcShapeV2Grids_referenceOrder[startarcshapev2grid]; ok {
		return order
	} else {
		log.Printf("instance %p of type StartArcShapeV2Grid was not staged and does not have a reference order", startarcshapev2grid)
		return 0
	}
}

func (topendarcshapev2 *TopEndArcShapeV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopEndArcShapeV2_stagedOrder[topendarcshapev2]; ok {
		return order
	}
	if order, ok := stage.TopEndArcShapeV2s_referenceOrder[topendarcshapev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopEndArcShapeV2 was not staged and does not have a reference order", topendarcshapev2)
		return 0
	}
}

func (topendarcshapev2grid *TopEndArcShapeV2Grid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopEndArcShapeV2Grid_stagedOrder[topendarcshapev2grid]; ok {
		return order
	}
	if order, ok := stage.TopEndArcShapeV2Grids_referenceOrder[topendarcshapev2grid]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopEndArcShapeV2Grid was not staged and does not have a reference order", topendarcshapev2grid)
		return 0
	}
}

func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStackGrowthCurveEndArcShapeV2_stagedOrder[topstackgrowthcurveendarcshapev2]; ok {
		return order
	}
	if order, ok := stage.TopStackGrowthCurveEndArcShapeV2s_referenceOrder[topstackgrowthcurveendarcshapev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStackGrowthCurveEndArcShapeV2 was not staged and does not have a reference order", topstackgrowthcurveendarcshapev2)
		return 0
	}
}

func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStackGrowthCurveStartArcShapeV2_stagedOrder[topstackgrowthcurvestartarcshapev2]; ok {
		return order
	}
	if order, ok := stage.TopStackGrowthCurveStartArcShapeV2s_referenceOrder[topstackgrowthcurvestartarcshapev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStackGrowthCurveStartArcShapeV2 was not staged and does not have a reference order", topstackgrowthcurvestartarcshapev2)
		return 0
	}
}

func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStackOfGrowthCurveV2_stagedOrder[topstackofgrowthcurvev2]; ok {
		return order
	}
	if order, ok := stage.TopStackOfGrowthCurveV2s_referenceOrder[topstackofgrowthcurvev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStackOfGrowthCurveV2 was not staged and does not have a reference order", topstackofgrowthcurvev2)
		return 0
	}
}

func (topstartarcshapev2 *TopStartArcShapeV2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStartArcShapeV2_stagedOrder[topstartarcshapev2]; ok {
		return order
	}
	if order, ok := stage.TopStartArcShapeV2s_referenceOrder[topstartarcshapev2]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStartArcShapeV2 was not staged and does not have a reference order", topstartarcshapev2)
		return 0
	}
}

func (topstartarcshapev2grid *TopStartArcShapeV2Grid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStartArcShapeV2Grid_stagedOrder[topstartarcshapev2grid]; ok {
		return order
	}
	if order, ok := stage.TopStartArcShapeV2Grids_referenceOrder[topstartarcshapev2grid]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStartArcShapeV2Grid was not staged and does not have a reference order", topstartarcshapev2grid)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (arcnormalvectorshape *ArcNormalVectorShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", arcnormalvectorshape.GongGetGongstructName(), arcnormalvectorshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arcnormalvectorshape *ArcNormalVectorShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", arcnormalvectorshape.GongGetGongstructName(), arcnormalvectorshape.GongGetOrder(stage))
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", arcnormalvectorshapegrid.GongGetGongstructName(), arcnormalvectorshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", arcnormalvectorshapegrid.GongGetGongstructName(), arcnormalvectorshapegrid.GongGetOrder(stage))
}

func (axesshape *AxesShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", axesshape.GongGetGongstructName(), axesshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (axesshape *AxesShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", axesshape.GongGetGongstructName(), axesshape.GongGetOrder(stage))
}

func (basevectorshape *BaseVectorShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", basevectorshape.GongGetGongstructName(), basevectorshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (basevectorshape *BaseVectorShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", basevectorshape.GongGetGongstructName(), basevectorshape.GongGetOrder(stage))
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", basevectorshapegrid.GongGetGongstructName(), basevectorshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (basevectorshapegrid *BaseVectorShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", basevectorshapegrid.GongGetGongstructName(), basevectorshapegrid.GongGetOrder(stage))
}

func (bottomendarcshapev2 *BottomEndArcShapeV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomendarcshapev2.GongGetGongstructName(), bottomendarcshapev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bottomendarcshapev2 *BottomEndArcShapeV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomendarcshapev2.GongGetGongstructName(), bottomendarcshapev2.GongGetOrder(stage))
}

func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomendarcshapev2grid.GongGetGongstructName(), bottomendarcshapev2grid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomendarcshapev2grid.GongGetGongstructName(), bottomendarcshapev2grid.GongGetOrder(stage))
}

func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomstackgrowthcurveendarcshapev2.GongGetGongstructName(), bottomstackgrowthcurveendarcshapev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomstackgrowthcurveendarcshapev2.GongGetGongstructName(), bottomstackgrowthcurveendarcshapev2.GongGetOrder(stage))
}

func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomstackgrowthcurvestartarcshapev2.GongGetGongstructName(), bottomstackgrowthcurvestartarcshapev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomstackgrowthcurvestartarcshapev2.GongGetGongstructName(), bottomstackgrowthcurvestartarcshapev2.GongGetOrder(stage))
}

func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomstackofgrowthcurvev2.GongGetGongstructName(), bottomstackofgrowthcurvev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomstackofgrowthcurvev2.GongGetGongstructName(), bottomstackofgrowthcurvev2.GongGetOrder(stage))
}

func (bottomstartarcshapev2 *BottomStartArcShapeV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomstartarcshapev2.GongGetGongstructName(), bottomstartarcshapev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bottomstartarcshapev2 *BottomStartArcShapeV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomstartarcshapev2.GongGetGongstructName(), bottomstartarcshapev2.GongGetOrder(stage))
}

func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomstartarcshapev2grid.GongGetGongstructName(), bottomstartarcshapev2grid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomstartarcshapev2grid.GongGetGongstructName(), bottomstartarcshapev2grid.GongGetOrder(stage))
}

func (circlegridshape *CircleGridShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circlegridshape.GongGetGongstructName(), circlegridshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circlegridshape *CircleGridShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circlegridshape.GongGetGongstructName(), circlegridshape.GongGetOrder(stage))
}

func (endarcshape *EndArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", endarcshape.GongGetGongstructName(), endarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (endarcshape *EndArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", endarcshape.GongGetGongstructName(), endarcshape.GongGetOrder(stage))
}

func (endarcshapegrid *EndArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", endarcshapegrid.GongGetGongstructName(), endarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (endarcshapegrid *EndArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", endarcshapegrid.GongGetGongstructName(), endarcshapegrid.GongGetOrder(stage))
}

func (endarcshapev2 *EndArcShapeV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", endarcshapev2.GongGetGongstructName(), endarcshapev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (endarcshapev2 *EndArcShapeV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", endarcshapev2.GongGetGongstructName(), endarcshapev2.GongGetOrder(stage))
}

func (endarcshapev2grid *EndArcShapeV2Grid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", endarcshapev2grid.GongGetGongstructName(), endarcshapev2grid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (endarcshapev2grid *EndArcShapeV2Grid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", endarcshapev2grid.GongGetGongstructName(), endarcshapev2grid.GongGetOrder(stage))
}

func (explanationtextshape *ExplanationTextShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", explanationtextshape.GongGetGongstructName(), explanationtextshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (explanationtextshape *ExplanationTextShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", explanationtextshape.GongGetGongstructName(), explanationtextshape.GongGetOrder(stage))
}

func (gridpathshape *GridPathShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gridpathshape.GongGetGongstructName(), gridpathshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gridpathshape *GridPathShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gridpathshape.GongGetGongstructName(), gridpathshape.GongGetOrder(stage))
}

func (growthcurvebeziershape *GrowthCurveBezierShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurvebeziershape.GongGetGongstructName(), growthcurvebeziershape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurvebeziershape *GrowthCurveBezierShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurvebeziershape.GongGetGongstructName(), growthcurvebeziershape.GongGetOrder(stage))
}

func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurvebeziershapegrid.GongGetGongstructName(), growthcurvebeziershapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurvebeziershapegrid.GongGetGongstructName(), growthcurvebeziershapegrid.GongGetOrder(stage))
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurverhombusgridshape.GongGetGongstructName(), growthcurverhombusgridshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurverhombusgridshape.GongGetGongstructName(), growthcurverhombusgridshape.GongGetOrder(stage))
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurverhombusshape.GongGetGongstructName(), growthcurverhombusshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurverhombusshape *GrowthCurveRhombusShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurverhombusshape.GongGetGongstructName(), growthcurverhombusshape.GongGetOrder(stage))
}

func (growthvectorshape *GrowthVectorShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthvectorshape.GongGetGongstructName(), growthvectorshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthvectorshape *GrowthVectorShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthvectorshape.GongGetGongstructName(), growthvectorshape.GongGetOrder(stage))
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", initialrhombusgridshape.GongGetGongstructName(), initialrhombusgridshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (initialrhombusgridshape *InitialRhombusGridShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", initialrhombusgridshape.GongGetGongstructName(), initialrhombusgridshape.GongGetOrder(stage))
}

func (initialrhombusshape *InitialRhombusShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", initialrhombusshape.GongGetGongstructName(), initialrhombusshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (initialrhombusshape *InitialRhombusShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", initialrhombusshape.GongGetGongstructName(), initialrhombusshape.GongGetOrder(stage))
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

func (nextcircleshape *NextCircleShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", nextcircleshape.GongGetGongstructName(), nextcircleshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (nextcircleshape *NextCircleShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", nextcircleshape.GongGetGongstructName(), nextcircleshape.GongGetOrder(stage))
}

func (perpendicularvector *PerpendicularVector) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", perpendicularvector.GongGetGongstructName(), perpendicularvector.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (perpendicularvector *PerpendicularVector) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", perpendicularvector.GongGetGongstructName(), perpendicularvector.GongGetOrder(stage))
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", perpendicularvectorgrid.GongGetGongstructName(), perpendicularvectorgrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (perpendicularvectorgrid *PerpendicularVectorGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", perpendicularvectorgrid.GongGetGongstructName(), perpendicularvectorgrid.GongGetOrder(stage))
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", perpendicularvectorgridhalfway.GongGetGongstructName(), perpendicularvectorgridhalfway.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", perpendicularvectorgridhalfway.GongGetGongstructName(), perpendicularvectorgridhalfway.GongGetOrder(stage))
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", perpendicularvectorhalfway.GongGetGongstructName(), perpendicularvectorhalfway.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", perpendicularvectorhalfway.GongGetGongstructName(), perpendicularvectorhalfway.GongGetOrder(stage))
}

func (plant *Plant) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plant.GongGetGongstructName(), plant.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (plant *Plant) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plant.GongGetGongstructName(), plant.GongGetOrder(stage))
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plantcircumferenceshape.GongGetGongstructName(), plantcircumferenceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (plantcircumferenceshape *PlantCircumferenceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plantcircumferenceshape.GongGetGongstructName(), plantcircumferenceshape.GongGetOrder(stage))
}

func (plantdiagram *PlantDiagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plantdiagram.GongGetGongstructName(), plantdiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (plantdiagram *PlantDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plantdiagram.GongGetGongstructName(), plantdiagram.GongGetOrder(stage))
}

func (rendered3dshape *Rendered3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rendered3dshape.GongGetGongstructName(), rendered3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rendered3dshape *Rendered3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rendered3dshape.GongGetGongstructName(), rendered3dshape.GongGetOrder(stage))
}

func (rhombusshape *RhombusShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rhombusshape.GongGetGongstructName(), rhombusshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rhombusshape *RhombusShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rhombusshape.GongGetGongstructName(), rhombusshape.GongGetOrder(stage))
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rotatedrhombusgridshape.GongGetGongstructName(), rotatedrhombusgridshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rotatedrhombusgridshape.GongGetGongstructName(), rotatedrhombusgridshape.GongGetOrder(stage))
}

func (rotatedrhombusshape *RotatedRhombusShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rotatedrhombusshape.GongGetGongstructName(), rotatedrhombusshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rotatedrhombusshape *RotatedRhombusShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rotatedrhombusshape.GongGetGongstructName(), rotatedrhombusshape.GongGetOrder(stage))
}

func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurvebeziershape.GongGetGongstructName(), stackgrowthcurvebeziershape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurvebeziershape.GongGetGongstructName(), stackgrowthcurvebeziershape.GongGetOrder(stage))
}

func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurveendarcshapev2.GongGetGongstructName(), stackgrowthcurveendarcshapev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurveendarcshapev2.GongGetGongstructName(), stackgrowthcurveendarcshapev2.GongGetOrder(stage))
}

func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurvestartarcshapev2.GongGetGongstructName(), stackgrowthcurvestartarcshapev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurvestartarcshapev2.GongGetGongstructName(), stackgrowthcurvestartarcshapev2.GongGetOrder(stage))
}

func (stackofgrowthcurve *StackOfGrowthCurve) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurve.GongGetGongstructName(), stackofgrowthcurve.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofgrowthcurve *StackOfGrowthCurve) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurve.GongGetGongstructName(), stackofgrowthcurve.GongGetOrder(stage))
}

func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurvev2.GongGetGongstructName(), stackofgrowthcurvev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurvev2.GongGetGongstructName(), stackofgrowthcurvev2.GongGetOrder(stage))
}

func (startarcshape *StartArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", startarcshape.GongGetGongstructName(), startarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (startarcshape *StartArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", startarcshape.GongGetGongstructName(), startarcshape.GongGetOrder(stage))
}

func (startarcshapegrid *StartArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", startarcshapegrid.GongGetGongstructName(), startarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (startarcshapegrid *StartArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", startarcshapegrid.GongGetGongstructName(), startarcshapegrid.GongGetOrder(stage))
}

func (startarcshapev2 *StartArcShapeV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", startarcshapev2.GongGetGongstructName(), startarcshapev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (startarcshapev2 *StartArcShapeV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", startarcshapev2.GongGetGongstructName(), startarcshapev2.GongGetOrder(stage))
}

func (startarcshapev2grid *StartArcShapeV2Grid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", startarcshapev2grid.GongGetGongstructName(), startarcshapev2grid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (startarcshapev2grid *StartArcShapeV2Grid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", startarcshapev2grid.GongGetGongstructName(), startarcshapev2grid.GongGetOrder(stage))
}

func (topendarcshapev2 *TopEndArcShapeV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topendarcshapev2.GongGetGongstructName(), topendarcshapev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topendarcshapev2 *TopEndArcShapeV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topendarcshapev2.GongGetGongstructName(), topendarcshapev2.GongGetOrder(stage))
}

func (topendarcshapev2grid *TopEndArcShapeV2Grid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topendarcshapev2grid.GongGetGongstructName(), topendarcshapev2grid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topendarcshapev2grid *TopEndArcShapeV2Grid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topendarcshapev2grid.GongGetGongstructName(), topendarcshapev2grid.GongGetOrder(stage))
}

func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackgrowthcurveendarcshapev2.GongGetGongstructName(), topstackgrowthcurveendarcshapev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackgrowthcurveendarcshapev2.GongGetGongstructName(), topstackgrowthcurveendarcshapev2.GongGetOrder(stage))
}

func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackgrowthcurvestartarcshapev2.GongGetGongstructName(), topstackgrowthcurvestartarcshapev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackgrowthcurvestartarcshapev2.GongGetGongstructName(), topstackgrowthcurvestartarcshapev2.GongGetOrder(stage))
}

func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackofgrowthcurvev2.GongGetGongstructName(), topstackofgrowthcurvev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackofgrowthcurvev2.GongGetGongstructName(), topstackofgrowthcurvev2.GongGetOrder(stage))
}

func (topstartarcshapev2 *TopStartArcShapeV2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstartarcshapev2.GongGetGongstructName(), topstartarcshapev2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstartarcshapev2 *TopStartArcShapeV2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstartarcshapev2.GongGetGongstructName(), topstartarcshapev2.GongGetOrder(stage))
}

func (topstartarcshapev2grid *TopStartArcShapeV2Grid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstartarcshapev2grid.GongGetGongstructName(), topstartarcshapev2grid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstartarcshapev2grid *TopStartArcShapeV2Grid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstartarcshapev2grid.GongGetGongstructName(), topstartarcshapev2grid.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (arcnormalvectorshape *ArcNormalVectorShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arcnormalvectorshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ArcNormalVectorShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(arcnormalvectorshape.Name))
	return
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arcnormalvectorshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ArcNormalVectorShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(arcnormalvectorshapegrid.Name))
	return
}

func (axesshape *AxesShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", axesshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AxesShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(axesshape.Name))
	return
}

func (basevectorshape *BaseVectorShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", basevectorshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BaseVectorShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(basevectorshape.Name))
	return
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", basevectorshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BaseVectorShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(basevectorshapegrid.Name))
	return
}

func (bottomendarcshapev2 *BottomEndArcShapeV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomendarcshapev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BottomEndArcShapeV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bottomendarcshapev2.Name))
	return
}

func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomendarcshapev2grid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BottomEndArcShapeV2Grid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bottomendarcshapev2grid.Name))
	return
}

func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomstackgrowthcurveendarcshapev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BottomStackGrowthCurveEndArcShapeV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bottomstackgrowthcurveendarcshapev2.Name))
	return
}

func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomstackgrowthcurvestartarcshapev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BottomStackGrowthCurveStartArcShapeV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bottomstackgrowthcurvestartarcshapev2.Name))
	return
}

func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomstackofgrowthcurvev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BottomStackOfGrowthCurveV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bottomstackofgrowthcurvev2.Name))
	return
}

func (bottomstartarcshapev2 *BottomStartArcShapeV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomstartarcshapev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BottomStartArcShapeV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bottomstartarcshapev2.Name))
	return
}

func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomstartarcshapev2grid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BottomStartArcShapeV2Grid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(bottomstartarcshapev2grid.Name))
	return
}

func (circlegridshape *CircleGridShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circlegridshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CircleGridShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(circlegridshape.Name))
	return
}

func (endarcshape *EndArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EndArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(endarcshape.Name))
	return
}

func (endarcshapegrid *EndArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EndArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(endarcshapegrid.Name))
	return
}

func (endarcshapev2 *EndArcShapeV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endarcshapev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EndArcShapeV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(endarcshapev2.Name))
	return
}

func (endarcshapev2grid *EndArcShapeV2Grid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endarcshapev2grid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EndArcShapeV2Grid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(endarcshapev2grid.Name))
	return
}

func (explanationtextshape *ExplanationTextShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", explanationtextshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ExplanationTextShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(explanationtextshape.Name))
	return
}

func (gridpathshape *GridPathShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gridpathshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GridPathShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(gridpathshape.Name))
	return
}

func (growthcurvebeziershape *GrowthCurveBezierShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurvebeziershape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GrowthCurveBezierShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(growthcurvebeziershape.Name))
	return
}

func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurvebeziershapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GrowthCurveBezierShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(growthcurvebeziershapegrid.Name))
	return
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurverhombusgridshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GrowthCurveRhombusGridShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(growthcurverhombusgridshape.Name))
	return
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurverhombusshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GrowthCurveRhombusShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(growthcurverhombusshape.Name))
	return
}

func (growthvectorshape *GrowthVectorShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthvectorshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GrowthVectorShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(growthvectorshape.Name))
	return
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", initialrhombusgridshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "InitialRhombusGridShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(initialrhombusgridshape.Name))
	return
}

func (initialrhombusshape *InitialRhombusShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", initialrhombusshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "InitialRhombusShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(initialrhombusshape.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.Name))
	return
}

func (nextcircleshape *NextCircleShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", nextcircleshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NextCircleShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(nextcircleshape.Name))
	return
}

func (perpendicularvector *PerpendicularVector) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", perpendicularvector.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PerpendicularVector")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(perpendicularvector.Name))
	return
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", perpendicularvectorgrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PerpendicularVectorGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(perpendicularvectorgrid.Name))
	return
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", perpendicularvectorgridhalfway.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PerpendicularVectorGridHalfway")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(perpendicularvectorgridhalfway.Name))
	return
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", perpendicularvectorhalfway.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PerpendicularVectorHalfway")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(perpendicularvectorhalfway.Name))
	return
}

func (plant *Plant) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plant.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Plant")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(plant.Name))
	return
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plantcircumferenceshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PlantCircumferenceShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(plantcircumferenceshape.Name))
	return
}

func (plantdiagram *PlantDiagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plantdiagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PlantDiagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(plantdiagram.Name))
	return
}

func (rendered3dshape *Rendered3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rendered3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Rendered3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rendered3dshape.Name))
	return
}

func (rhombusshape *RhombusShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rhombusshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RhombusShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rhombusshape.Name))
	return
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rotatedrhombusgridshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RotatedRhombusGridShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rotatedrhombusgridshape.Name))
	return
}

func (rotatedrhombusshape *RotatedRhombusShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rotatedrhombusshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RotatedRhombusShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rotatedrhombusshape.Name))
	return
}

func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurvebeziershape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackGrowthCurveBezierShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackgrowthcurvebeziershape.Name))
	return
}

func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurveendarcshapev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackGrowthCurveEndArcShapeV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackgrowthcurveendarcshapev2.Name))
	return
}

func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurvestartarcshapev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackGrowthCurveStartArcShapeV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackgrowthcurvestartarcshapev2.Name))
	return
}

func (stackofgrowthcurve *StackOfGrowthCurve) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurve.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfGrowthCurve")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackofgrowthcurve.Name))
	return
}

func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurvev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfGrowthCurveV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackofgrowthcurvev2.Name))
	return
}

func (startarcshape *StartArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", startarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StartArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(startarcshape.Name))
	return
}

func (startarcshapegrid *StartArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", startarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StartArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(startarcshapegrid.Name))
	return
}

func (startarcshapev2 *StartArcShapeV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", startarcshapev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StartArcShapeV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(startarcshapev2.Name))
	return
}

func (startarcshapev2grid *StartArcShapeV2Grid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", startarcshapev2grid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StartArcShapeV2Grid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(startarcshapev2grid.Name))
	return
}

func (topendarcshapev2 *TopEndArcShapeV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendarcshapev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopEndArcShapeV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topendarcshapev2.Name))
	return
}

func (topendarcshapev2grid *TopEndArcShapeV2Grid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendarcshapev2grid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopEndArcShapeV2Grid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topendarcshapev2grid.Name))
	return
}

func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackgrowthcurveendarcshapev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackGrowthCurveEndArcShapeV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstackgrowthcurveendarcshapev2.Name))
	return
}

func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackgrowthcurvestartarcshapev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackGrowthCurveStartArcShapeV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstackgrowthcurvestartarcshapev2.Name))
	return
}

func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofgrowthcurvev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackOfGrowthCurveV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstackofgrowthcurvev2.Name))
	return
}

func (topstartarcshapev2 *TopStartArcShapeV2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstartarcshapev2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStartArcShapeV2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstartarcshapev2.Name))
	return
}

func (topstartarcshapev2grid *TopStartArcShapeV2Grid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstartarcshapev2grid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStartArcShapeV2Grid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstartarcshapev2grid.Name))
	return
}

// insertion point for unstaging
func (arcnormalvectorshape *ArcNormalVectorShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arcnormalvectorshape.GongGetReferenceIdentifier(stage))
	return
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arcnormalvectorshapegrid.GongGetReferenceIdentifier(stage))
	return
}

func (axesshape *AxesShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", axesshape.GongGetReferenceIdentifier(stage))
	return
}

func (basevectorshape *BaseVectorShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", basevectorshape.GongGetReferenceIdentifier(stage))
	return
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", basevectorshapegrid.GongGetReferenceIdentifier(stage))
	return
}

func (bottomendarcshapev2 *BottomEndArcShapeV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomendarcshapev2.GongGetReferenceIdentifier(stage))
	return
}

func (bottomendarcshapev2grid *BottomEndArcShapeV2Grid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomendarcshapev2grid.GongGetReferenceIdentifier(stage))
	return
}

func (bottomstackgrowthcurveendarcshapev2 *BottomStackGrowthCurveEndArcShapeV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomstackgrowthcurveendarcshapev2.GongGetReferenceIdentifier(stage))
	return
}

func (bottomstackgrowthcurvestartarcshapev2 *BottomStackGrowthCurveStartArcShapeV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomstackgrowthcurvestartarcshapev2.GongGetReferenceIdentifier(stage))
	return
}

func (bottomstackofgrowthcurvev2 *BottomStackOfGrowthCurveV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomstackofgrowthcurvev2.GongGetReferenceIdentifier(stage))
	return
}

func (bottomstartarcshapev2 *BottomStartArcShapeV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomstartarcshapev2.GongGetReferenceIdentifier(stage))
	return
}

func (bottomstartarcshapev2grid *BottomStartArcShapeV2Grid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomstartarcshapev2grid.GongGetReferenceIdentifier(stage))
	return
}

func (circlegridshape *CircleGridShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circlegridshape.GongGetReferenceIdentifier(stage))
	return
}

func (endarcshape *EndArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (endarcshapegrid *EndArcShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endarcshapegrid.GongGetReferenceIdentifier(stage))
	return
}

func (endarcshapev2 *EndArcShapeV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endarcshapev2.GongGetReferenceIdentifier(stage))
	return
}

func (endarcshapev2grid *EndArcShapeV2Grid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endarcshapev2grid.GongGetReferenceIdentifier(stage))
	return
}

func (explanationtextshape *ExplanationTextShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", explanationtextshape.GongGetReferenceIdentifier(stage))
	return
}

func (gridpathshape *GridPathShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gridpathshape.GongGetReferenceIdentifier(stage))
	return
}

func (growthcurvebeziershape *GrowthCurveBezierShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurvebeziershape.GongGetReferenceIdentifier(stage))
	return
}

func (growthcurvebeziershapegrid *GrowthCurveBezierShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurvebeziershapegrid.GongGetReferenceIdentifier(stage))
	return
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurverhombusgridshape.GongGetReferenceIdentifier(stage))
	return
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurverhombusshape.GongGetReferenceIdentifier(stage))
	return
}

func (growthvectorshape *GrowthVectorShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthvectorshape.GongGetReferenceIdentifier(stage))
	return
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", initialrhombusgridshape.GongGetReferenceIdentifier(stage))
	return
}

func (initialrhombusshape *InitialRhombusShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", initialrhombusshape.GongGetReferenceIdentifier(stage))
	return
}

func (library *Library) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetReferenceIdentifier(stage))
	return
}

func (nextcircleshape *NextCircleShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", nextcircleshape.GongGetReferenceIdentifier(stage))
	return
}

func (perpendicularvector *PerpendicularVector) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", perpendicularvector.GongGetReferenceIdentifier(stage))
	return
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", perpendicularvectorgrid.GongGetReferenceIdentifier(stage))
	return
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", perpendicularvectorgridhalfway.GongGetReferenceIdentifier(stage))
	return
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", perpendicularvectorhalfway.GongGetReferenceIdentifier(stage))
	return
}

func (plant *Plant) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plant.GongGetReferenceIdentifier(stage))
	return
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plantcircumferenceshape.GongGetReferenceIdentifier(stage))
	return
}

func (plantdiagram *PlantDiagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plantdiagram.GongGetReferenceIdentifier(stage))
	return
}

func (rendered3dshape *Rendered3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rendered3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (rhombusshape *RhombusShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rhombusshape.GongGetReferenceIdentifier(stage))
	return
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rotatedrhombusgridshape.GongGetReferenceIdentifier(stage))
	return
}

func (rotatedrhombusshape *RotatedRhombusShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rotatedrhombusshape.GongGetReferenceIdentifier(stage))
	return
}

func (stackgrowthcurvebeziershape *StackGrowthCurveBezierShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurvebeziershape.GongGetReferenceIdentifier(stage))
	return
}

func (stackgrowthcurveendarcshapev2 *StackGrowthCurveEndArcShapeV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurveendarcshapev2.GongGetReferenceIdentifier(stage))
	return
}

func (stackgrowthcurvestartarcshapev2 *StackGrowthCurveStartArcShapeV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurvestartarcshapev2.GongGetReferenceIdentifier(stage))
	return
}

func (stackofgrowthcurve *StackOfGrowthCurve) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurve.GongGetReferenceIdentifier(stage))
	return
}

func (stackofgrowthcurvev2 *StackOfGrowthCurveV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurvev2.GongGetReferenceIdentifier(stage))
	return
}

func (startarcshape *StartArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", startarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (startarcshapegrid *StartArcShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", startarcshapegrid.GongGetReferenceIdentifier(stage))
	return
}

func (startarcshapev2 *StartArcShapeV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", startarcshapev2.GongGetReferenceIdentifier(stage))
	return
}

func (startarcshapev2grid *StartArcShapeV2Grid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", startarcshapev2grid.GongGetReferenceIdentifier(stage))
	return
}

func (topendarcshapev2 *TopEndArcShapeV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendarcshapev2.GongGetReferenceIdentifier(stage))
	return
}

func (topendarcshapev2grid *TopEndArcShapeV2Grid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendarcshapev2grid.GongGetReferenceIdentifier(stage))
	return
}

func (topstackgrowthcurveendarcshapev2 *TopStackGrowthCurveEndArcShapeV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackgrowthcurveendarcshapev2.GongGetReferenceIdentifier(stage))
	return
}

func (topstackgrowthcurvestartarcshapev2 *TopStackGrowthCurveStartArcShapeV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackgrowthcurvestartarcshapev2.GongGetReferenceIdentifier(stage))
	return
}

func (topstackofgrowthcurvev2 *TopStackOfGrowthCurveV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofgrowthcurvev2.GongGetReferenceIdentifier(stage))
	return
}

func (topstartarcshapev2 *TopStartArcShapeV2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstartarcshapev2.GongGetReferenceIdentifier(stage))
	return
}

func (topstartarcshapev2grid *TopStartArcShapeV2Grid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstartarcshapev2grid.GongGetReferenceIdentifier(stage))
	return
}

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// end of template
