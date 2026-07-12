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

	// Compute reverse map for named struct StackOfGrowthCurve
	// insertion point per field
	stage.StackOfGrowthCurve_StackGrowthCurveBezierShapes_reverseMap = make(map[*StackGrowthCurveBezierShape]*StackOfGrowthCurve)
	for stackofgrowthcurve := range stage.StackOfGrowthCurves {
		_ = stackofgrowthcurve
		for _, _stackgrowthcurvebeziershape := range stackofgrowthcurve.StackGrowthCurveBezierShapes {
			stage.StackOfGrowthCurve_StackGrowthCurveBezierShapes_reverseMap[_stackgrowthcurvebeziershape] = stackofgrowthcurve
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

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.AxesShapes {
		res = append(res, instance)
	}

	for instance := range stage.BaseVectorShapes {
		res = append(res, instance)
	}

	for instance := range stage.BaseVectorShapeGrids {
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

	for instance := range stage.StackOfGrowthCurves {
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

	return
}

// insertion point per named struct
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

func (stackofgrowthcurve *StackOfGrowthCurve) GongCopy() GongstructIF {
	newInstance := new(StackOfGrowthCurve)
	stackofgrowthcurve.CopyBasicFields(newInstance)
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

// insertion point per named struct
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

func (stackofgrowthcurve *StackOfGrowthCurve) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofgrowthcurve).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackofgrowthcurve), uint64(GetOrderPointerGongstruct(stage, stackofgrowthcurve)))
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

	stage.StackOfGrowthCurves_reference = make(map[*StackOfGrowthCurve]*StackOfGrowthCurve)
	stage.StackOfGrowthCurves_referenceOrder = make(map[*StackOfGrowthCurve]uint) // diff Unstage needs the reference order
	stage.StackOfGrowthCurves_instance = make(map[*StackOfGrowthCurve]*StackOfGrowthCurve)
	for instance := range stage.StackOfGrowthCurves {
		_copy := instance.GongCopy().(*StackOfGrowthCurve)
		stage.StackOfGrowthCurves_reference[instance] = _copy
		stage.StackOfGrowthCurves_instance[_copy] = instance
		stage.StackOfGrowthCurves_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	// insertion point per named struct
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

	for instance := range stage.StackOfGrowthCurves {
		reference := stage.StackOfGrowthCurves_reference[instance]
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

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
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

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
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

func (stackofgrowthcurve *StackOfGrowthCurve) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurve.GongGetGongstructName(), stackofgrowthcurve.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofgrowthcurve *StackOfGrowthCurve) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurve.GongGetGongstructName(), stackofgrowthcurve.GongGetOrder(stage))
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

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
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

func (stackofgrowthcurve *StackOfGrowthCurve) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurve.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfGrowthCurve")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackofgrowthcurve.Name))
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

// insertion point for unstaging
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

func (stackofgrowthcurve *StackOfGrowthCurve) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurve.GongGetReferenceIdentifier(stage))
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
