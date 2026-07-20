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

	// Compute reverse map for named struct EndHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct EndHalfwayArcShapeGrid
	// insertion point per field
	stage.EndHalfwayArcShapeGrid_EndHalfwayArcShapes_reverseMap = make(map[*EndHalfwayArcShape]*EndHalfwayArcShapeGrid)
	for endhalfwayarcshapegrid := range stage.EndHalfwayArcShapeGrids {
		_ = endhalfwayarcshapegrid
		for _, _endhalfwayarcshape := range endhalfwayarcshapegrid.EndHalfwayArcShapes {
			stage.EndHalfwayArcShapeGrid_EndHalfwayArcShapes_reverseMap[_endhalfwayarcshape] = endhalfwayarcshapegrid
		}
	}

	// Compute reverse map for named struct ExplanationTextShape
	// insertion point per field

	// Compute reverse map for named struct GridPathShape
	// insertion point per field

	// Compute reverse map for named struct GrowthCurve2D
	// insertion point per field

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

	// Compute reverse map for named struct MidArcVectorShape
	// insertion point per field

	// Compute reverse map for named struct MidArcVectorShapeGrid
	// insertion point per field
	stage.MidArcVectorShapeGrid_MidArcVectorShapes_reverseMap = make(map[*MidArcVectorShape]*MidArcVectorShapeGrid)
	for midarcvectorshapegrid := range stage.MidArcVectorShapeGrids {
		_ = midarcvectorshapegrid
		for _, _midarcvectorshape := range midarcvectorshapegrid.MidArcVectorShapes {
			stage.MidArcVectorShapeGrid_MidArcVectorShapes_reverseMap[_midarcvectorshape] = midarcvectorshapegrid
		}
	}

	// Compute reverse map for named struct PartiallyGrowthCurve2DRibbon
	// insertion point per field
	stage.PartiallyGrowthCurve2DRibbon_PartiallyGrowthCurve2DRibbonStartShapes_reverseMap = make(map[*PartiallyGrowthCurve2DRibbonStartShape]*PartiallyGrowthCurve2DRibbon)
	for partiallygrowthcurve2dribbon := range stage.PartiallyGrowthCurve2DRibbons {
		_ = partiallygrowthcurve2dribbon
		for _, _partiallygrowthcurve2dribbonstartshape := range partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonStartShapes {
			stage.PartiallyGrowthCurve2DRibbon_PartiallyGrowthCurve2DRibbonStartShapes_reverseMap[_partiallygrowthcurve2dribbonstartshape] = partiallygrowthcurve2dribbon
		}
	}
	stage.PartiallyGrowthCurve2DRibbon_PartiallyGrowthCurve2DRibbonEndShapes_reverseMap = make(map[*PartiallyGrowthCurve2DRibbonEndShape]*PartiallyGrowthCurve2DRibbon)
	for partiallygrowthcurve2dribbon := range stage.PartiallyGrowthCurve2DRibbons {
		_ = partiallygrowthcurve2dribbon
		for _, _partiallygrowthcurve2dribbonendshape := range partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonEndShapes {
			stage.PartiallyGrowthCurve2DRibbon_PartiallyGrowthCurve2DRibbonEndShapes_reverseMap[_partiallygrowthcurve2dribbonendshape] = partiallygrowthcurve2dribbon
		}
	}

	// Compute reverse map for named struct PartiallyGrowthCurve2DRibbonEndShape
	// insertion point per field

	// Compute reverse map for named struct PartiallyGrowthCurve2DRibbonStartShape
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

	// Compute reverse map for named struct RhombusStuff
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

	// Compute reverse map for named struct ShiftedBottomTopStartArcShape
	// insertion point per field

	// Compute reverse map for named struct ShiftedBottomTopStartArcShapeGrid
	// insertion point per field
	stage.ShiftedBottomTopStartArcShapeGrid_ShiftedBottomTopStartArcShapes_reverseMap = make(map[*ShiftedBottomTopStartArcShape]*ShiftedBottomTopStartArcShapeGrid)
	for shiftedbottomtopstartarcshapegrid := range stage.ShiftedBottomTopStartArcShapeGrids {
		_ = shiftedbottomtopstartarcshapegrid
		for _, _shiftedbottomtopstartarcshape := range shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes {
			stage.ShiftedBottomTopStartArcShapeGrid_ShiftedBottomTopStartArcShapes_reverseMap[_shiftedbottomtopstartarcshape] = shiftedbottomtopstartarcshapegrid
		}
	}

	// Compute reverse map for named struct ShiftedLeftStackGrowthCurveEndArcShape
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftStackGrowthCurveStartArcShape
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftStackNormalVector
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftStackOfGrowthCurve
	// insertion point per field
	stage.ShiftedLeftStackOfGrowthCurve_ShiftedLeftStackGrowthCurveStartArcShapes_reverseMap = make(map[*ShiftedLeftStackGrowthCurveStartArcShape]*ShiftedLeftStackOfGrowthCurve)
	for shiftedleftstackofgrowthcurve := range stage.ShiftedLeftStackOfGrowthCurves {
		_ = shiftedleftstackofgrowthcurve
		for _, _shiftedleftstackgrowthcurvestartarcshape := range shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes {
			stage.ShiftedLeftStackOfGrowthCurve_ShiftedLeftStackGrowthCurveStartArcShapes_reverseMap[_shiftedleftstackgrowthcurvestartarcshape] = shiftedleftstackofgrowthcurve
		}
	}
	stage.ShiftedLeftStackOfGrowthCurve_ShiftedLeftStackGrowthCurveEndArcShapes_reverseMap = make(map[*ShiftedLeftStackGrowthCurveEndArcShape]*ShiftedLeftStackOfGrowthCurve)
	for shiftedleftstackofgrowthcurve := range stage.ShiftedLeftStackOfGrowthCurves {
		_ = shiftedleftstackofgrowthcurve
		for _, _shiftedleftstackgrowthcurveendarcshape := range shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes {
			stage.ShiftedLeftStackOfGrowthCurve_ShiftedLeftStackGrowthCurveEndArcShapes_reverseMap[_shiftedleftstackgrowthcurveendarcshape] = shiftedleftstackofgrowthcurve
		}
	}

	// Compute reverse map for named struct ShiftedLeftStackOfNormalVector
	// insertion point per field
	stage.ShiftedLeftStackOfNormalVector_ShiftedLeftStackNormalVectors_reverseMap = make(map[*ShiftedLeftStackNormalVector]*ShiftedLeftStackOfNormalVector)
	for shiftedleftstackofnormalvector := range stage.ShiftedLeftStackOfNormalVectors {
		_ = shiftedleftstackofnormalvector
		for _, _shiftedleftstacknormalvector := range shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors {
			stage.ShiftedLeftStackOfNormalVector_ShiftedLeftStackNormalVectors_reverseMap[_shiftedleftstacknormalvector] = shiftedleftstackofnormalvector
		}
	}

	// Compute reverse map for named struct StackGrowthCurve2DEndHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct StackGrowthCurve2DRibbonEndShape
	// insertion point per field

	// Compute reverse map for named struct StackGrowthCurve2DRibbonStartShape
	// insertion point per field

	// Compute reverse map for named struct StackGrowthCurve2DStartHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct StackOfGrowthCurve2D
	// insertion point per field
	stage.StackOfGrowthCurve2D_StackGrowthCurve2DStartHalfwayArcShapes_reverseMap = make(map[*StackGrowthCurve2DStartHalfwayArcShape]*StackOfGrowthCurve2D)
	for stackofgrowthcurve2d := range stage.StackOfGrowthCurve2Ds {
		_ = stackofgrowthcurve2d
		for _, _stackgrowthcurve2dstarthalfwayarcshape := range stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes {
			stage.StackOfGrowthCurve2D_StackGrowthCurve2DStartHalfwayArcShapes_reverseMap[_stackgrowthcurve2dstarthalfwayarcshape] = stackofgrowthcurve2d
		}
	}
	stage.StackOfGrowthCurve2D_StackGrowthCurve2DEndHalfwayArcShapes_reverseMap = make(map[*StackGrowthCurve2DEndHalfwayArcShape]*StackOfGrowthCurve2D)
	for stackofgrowthcurve2d := range stage.StackOfGrowthCurve2Ds {
		_ = stackofgrowthcurve2d
		for _, _stackgrowthcurve2dendhalfwayarcshape := range stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes {
			stage.StackOfGrowthCurve2D_StackGrowthCurve2DEndHalfwayArcShapes_reverseMap[_stackgrowthcurve2dendhalfwayarcshape] = stackofgrowthcurve2d
		}
	}

	// Compute reverse map for named struct StackOfGrowthCurve2DRibbon
	// insertion point per field
	stage.StackOfGrowthCurve2DRibbon_StackGrowthCurve2DRibbonStartShapes_reverseMap = make(map[*StackGrowthCurve2DRibbonStartShape]*StackOfGrowthCurve2DRibbon)
	for stackofgrowthcurve2dribbon := range stage.StackOfGrowthCurve2DRibbons {
		_ = stackofgrowthcurve2dribbon
		for _, _stackgrowthcurve2dribbonstartshape := range stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes {
			stage.StackOfGrowthCurve2DRibbon_StackGrowthCurve2DRibbonStartShapes_reverseMap[_stackgrowthcurve2dribbonstartshape] = stackofgrowthcurve2dribbon
		}
	}
	stage.StackOfGrowthCurve2DRibbon_StackGrowthCurve2DRibbonEndShapes_reverseMap = make(map[*StackGrowthCurve2DRibbonEndShape]*StackOfGrowthCurve2DRibbon)
	for stackofgrowthcurve2dribbon := range stage.StackOfGrowthCurve2DRibbons {
		_ = stackofgrowthcurve2dribbon
		for _, _stackgrowthcurve2dribbonendshape := range stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes {
			stage.StackOfGrowthCurve2DRibbon_StackGrowthCurve2DRibbonEndShapes_reverseMap[_stackgrowthcurve2dribbonendshape] = stackofgrowthcurve2dribbon
		}
	}

	// Compute reverse map for named struct StackOfRotatedGrowthCurve2D
	// insertion point per field
	stage.StackOfRotatedGrowthCurve2D_StackRotatedGrowthCurve2DStartArcShapes_reverseMap = make(map[*StackRotatedGrowthCurve2DStartArcShape]*StackOfRotatedGrowthCurve2D)
	for stackofrotatedgrowthcurve2d := range stage.StackOfRotatedGrowthCurve2Ds {
		_ = stackofrotatedgrowthcurve2d
		for _, _stackrotatedgrowthcurve2dstartarcshape := range stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes {
			stage.StackOfRotatedGrowthCurve2D_StackRotatedGrowthCurve2DStartArcShapes_reverseMap[_stackrotatedgrowthcurve2dstartarcshape] = stackofrotatedgrowthcurve2d
		}
	}
	stage.StackOfRotatedGrowthCurve2D_StackRotatedGrowthCurve2DEndArcShapes_reverseMap = make(map[*StackRotatedGrowthCurve2DEndArcShape]*StackOfRotatedGrowthCurve2D)
	for stackofrotatedgrowthcurve2d := range stage.StackOfRotatedGrowthCurve2Ds {
		_ = stackofrotatedgrowthcurve2d
		for _, _stackrotatedgrowthcurve2dendarcshape := range stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes {
			stage.StackOfRotatedGrowthCurve2D_StackRotatedGrowthCurve2DEndArcShapes_reverseMap[_stackrotatedgrowthcurve2dendarcshape] = stackofrotatedgrowthcurve2d
		}
	}

	// Compute reverse map for named struct StackOfRotatedGrowthCurve2DRibbon
	// insertion point per field
	stage.StackOfRotatedGrowthCurve2DRibbon_StackRotatedGrowthCurve2DRibbonStartShapes_reverseMap = make(map[*StackRotatedGrowthCurve2DRibbonStartShape]*StackOfRotatedGrowthCurve2DRibbon)
	for stackofrotatedgrowthcurve2dribbon := range stage.StackOfRotatedGrowthCurve2DRibbons {
		_ = stackofrotatedgrowthcurve2dribbon
		for _, _stackrotatedgrowthcurve2dribbonstartshape := range stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonStartShapes {
			stage.StackOfRotatedGrowthCurve2DRibbon_StackRotatedGrowthCurve2DRibbonStartShapes_reverseMap[_stackrotatedgrowthcurve2dribbonstartshape] = stackofrotatedgrowthcurve2dribbon
		}
	}
	stage.StackOfRotatedGrowthCurve2DRibbon_StackRotatedGrowthCurve2DRibbonEndShapes_reverseMap = make(map[*StackRotatedGrowthCurve2DRibbonEndShape]*StackOfRotatedGrowthCurve2DRibbon)
	for stackofrotatedgrowthcurve2dribbon := range stage.StackOfRotatedGrowthCurve2DRibbons {
		_ = stackofrotatedgrowthcurve2dribbon
		for _, _stackrotatedgrowthcurve2dribbonendshape := range stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonEndShapes {
			stage.StackOfRotatedGrowthCurve2DRibbon_StackRotatedGrowthCurve2DRibbonEndShapes_reverseMap[_stackrotatedgrowthcurve2dribbonendshape] = stackofrotatedgrowthcurve2dribbon
		}
	}

	// Compute reverse map for named struct StackRotatedGrowthCurve2DEndArcShape
	// insertion point per field

	// Compute reverse map for named struct StackRotatedGrowthCurve2DRibbonEndShape
	// insertion point per field

	// Compute reverse map for named struct StackRotatedGrowthCurve2DRibbonStartShape
	// insertion point per field

	// Compute reverse map for named struct StackRotatedGrowthCurve2DStartArcShape
	// insertion point per field

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

	// Compute reverse map for named struct StartHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct StartHalfwayArcShapeGrid
	// insertion point per field
	stage.StartHalfwayArcShapeGrid_StartHalfwayArcShapes_reverseMap = make(map[*StartHalfwayArcShape]*StartHalfwayArcShapeGrid)
	for starthalfwayarcshapegrid := range stage.StartHalfwayArcShapeGrids {
		_ = starthalfwayarcshapegrid
		for _, _starthalfwayarcshape := range starthalfwayarcshapegrid.StartHalfwayArcShapes {
			stage.StartHalfwayArcShapeGrid_StartHalfwayArcShapes_reverseMap[_starthalfwayarcshape] = starthalfwayarcshapegrid
		}
	}

	// Compute reverse map for named struct TopEndArcShape
	// insertion point per field

	// Compute reverse map for named struct TopEndArcShapeGrid
	// insertion point per field
	stage.TopEndArcShapeGrid_TopEndArcShapes_reverseMap = make(map[*TopEndArcShape]*TopEndArcShapeGrid)
	for topendarcshapegrid := range stage.TopEndArcShapeGrids {
		_ = topendarcshapegrid
		for _, _topendarcshape := range topendarcshapegrid.TopEndArcShapes {
			stage.TopEndArcShapeGrid_TopEndArcShapes_reverseMap[_topendarcshape] = topendarcshapegrid
		}
	}

	// Compute reverse map for named struct TopEndHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct TopEndHalfwayArcShapeGrid
	// insertion point per field
	stage.TopEndHalfwayArcShapeGrid_TopEndHalfwayArcShapes_reverseMap = make(map[*TopEndHalfwayArcShape]*TopEndHalfwayArcShapeGrid)
	for topendhalfwayarcshapegrid := range stage.TopEndHalfwayArcShapeGrids {
		_ = topendhalfwayarcshapegrid
		for _, _topendhalfwayarcshape := range topendhalfwayarcshapegrid.TopEndHalfwayArcShapes {
			stage.TopEndHalfwayArcShapeGrid_TopEndHalfwayArcShapes_reverseMap[_topendhalfwayarcshape] = topendhalfwayarcshapegrid
		}
	}

	// Compute reverse map for named struct TopGrowthCurve2D
	// insertion point per field

	// Compute reverse map for named struct TopMidArcVectorShape
	// insertion point per field

	// Compute reverse map for named struct TopMidArcVectorShapeGrid
	// insertion point per field
	stage.TopMidArcVectorShapeGrid_TopMidArcVectorShapes_reverseMap = make(map[*TopMidArcVectorShape]*TopMidArcVectorShapeGrid)
	for topmidarcvectorshapegrid := range stage.TopMidArcVectorShapeGrids {
		_ = topmidarcvectorshapegrid
		for _, _topmidarcvectorshape := range topmidarcvectorshapegrid.TopMidArcVectorShapes {
			stage.TopMidArcVectorShapeGrid_TopMidArcVectorShapes_reverseMap[_topmidarcvectorshape] = topmidarcvectorshapegrid
		}
	}

	// Compute reverse map for named struct TopStackGrowthCurve2DEndHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct TopStackGrowthCurve2DStartHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct TopStackOfGrowthCurve2D
	// insertion point per field
	stage.TopStackOfGrowthCurve2D_TopStackGrowthCurve2DStartHalfwayArcShapes_reverseMap = make(map[*TopStackGrowthCurve2DStartHalfwayArcShape]*TopStackOfGrowthCurve2D)
	for topstackofgrowthcurve2d := range stage.TopStackOfGrowthCurve2Ds {
		_ = topstackofgrowthcurve2d
		for _, _topstackgrowthcurve2dstarthalfwayarcshape := range topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes {
			stage.TopStackOfGrowthCurve2D_TopStackGrowthCurve2DStartHalfwayArcShapes_reverseMap[_topstackgrowthcurve2dstarthalfwayarcshape] = topstackofgrowthcurve2d
		}
	}
	stage.TopStackOfGrowthCurve2D_TopStackGrowthCurve2DEndHalfwayArcShapes_reverseMap = make(map[*TopStackGrowthCurve2DEndHalfwayArcShape]*TopStackOfGrowthCurve2D)
	for topstackofgrowthcurve2d := range stage.TopStackOfGrowthCurve2Ds {
		_ = topstackofgrowthcurve2d
		for _, _topstackgrowthcurve2dendhalfwayarcshape := range topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes {
			stage.TopStackOfGrowthCurve2D_TopStackGrowthCurve2DEndHalfwayArcShapes_reverseMap[_topstackgrowthcurve2dendhalfwayarcshape] = topstackofgrowthcurve2d
		}
	}

	// Compute reverse map for named struct TopStackOfRotatedGrowthCurve2D
	// insertion point per field
	stage.TopStackOfRotatedGrowthCurve2D_TopStackOfRotatedGrowthCurve2DStartArcShapes_reverseMap = make(map[*TopStackOfRotatedGrowthCurve2DStartArcShape]*TopStackOfRotatedGrowthCurve2D)
	for topstackofrotatedgrowthcurve2d := range stage.TopStackOfRotatedGrowthCurve2Ds {
		_ = topstackofrotatedgrowthcurve2d
		for _, _topstackofrotatedgrowthcurve2dstartarcshape := range topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes {
			stage.TopStackOfRotatedGrowthCurve2D_TopStackOfRotatedGrowthCurve2DStartArcShapes_reverseMap[_topstackofrotatedgrowthcurve2dstartarcshape] = topstackofrotatedgrowthcurve2d
		}
	}
	stage.TopStackOfRotatedGrowthCurve2D_TopStackOfRotatedGrowthCurve2DEndArcShapes_reverseMap = make(map[*TopStackOfRotatedGrowthCurve2DEndArcShape]*TopStackOfRotatedGrowthCurve2D)
	for topstackofrotatedgrowthcurve2d := range stage.TopStackOfRotatedGrowthCurve2Ds {
		_ = topstackofrotatedgrowthcurve2d
		for _, _topstackofrotatedgrowthcurve2dendarcshape := range topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes {
			stage.TopStackOfRotatedGrowthCurve2D_TopStackOfRotatedGrowthCurve2DEndArcShapes_reverseMap[_topstackofrotatedgrowthcurve2dendarcshape] = topstackofrotatedgrowthcurve2d
		}
	}

	// Compute reverse map for named struct TopStackOfRotatedGrowthCurve2DEndArcShape
	// insertion point per field

	// Compute reverse map for named struct TopStackOfRotatedGrowthCurve2DStartArcShape
	// insertion point per field

	// Compute reverse map for named struct TopStartArcShape
	// insertion point per field

	// Compute reverse map for named struct TopStartArcShapeGrid
	// insertion point per field
	stage.TopStartArcShapeGrid_TopStartArcShapes_reverseMap = make(map[*TopStartArcShape]*TopStartArcShapeGrid)
	for topstartarcshapegrid := range stage.TopStartArcShapeGrids {
		_ = topstartarcshapegrid
		for _, _topstartarcshape := range topstartarcshapegrid.TopStartArcShapes {
			stage.TopStartArcShapeGrid_TopStartArcShapes_reverseMap[_topstartarcshape] = topstartarcshapegrid
		}
	}

	// Compute reverse map for named struct TopStartHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct TopStartHalfwayArcShapeGrid
	// insertion point per field
	stage.TopStartHalfwayArcShapeGrid_TopStartHalfwayArcShapes_reverseMap = make(map[*TopStartHalfwayArcShape]*TopStartHalfwayArcShapeGrid)
	for topstarthalfwayarcshapegrid := range stage.TopStartHalfwayArcShapeGrids {
		_ = topstarthalfwayarcshapegrid
		for _, _topstarthalfwayarcshape := range topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes {
			stage.TopStartHalfwayArcShapeGrid_TopStartHalfwayArcShapes_reverseMap[_topstarthalfwayarcshape] = topstarthalfwayarcshapegrid
		}
	}

	// Compute reverse map for named struct TorusStackShape
	// insertion point per field

	// Compute reverse map for named struct VerticalTorusStackShape
	// insertion point per field

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

	for instance := range stage.CircleGridShapes {
		res = append(res, instance)
	}

	for instance := range stage.EndArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.EndArcShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.EndHalfwayArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.EndHalfwayArcShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.ExplanationTextShapes {
		res = append(res, instance)
	}

	for instance := range stage.GridPathShapes {
		res = append(res, instance)
	}

	for instance := range stage.GrowthCurve2Ds {
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

	for instance := range stage.MidArcVectorShapes {
		res = append(res, instance)
	}

	for instance := range stage.MidArcVectorShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DRibbons {
		res = append(res, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DRibbonEndShapes {
		res = append(res, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DRibbonStartShapes {
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

	for instance := range stage.RhombusStuffs {
		res = append(res, instance)
	}

	for instance := range stage.RotatedRhombusGridShapes {
		res = append(res, instance)
	}

	for instance := range stage.RotatedRhombusShapes {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedBottomTopStartArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedBottomTopStartArcShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedLeftStackGrowthCurveEndArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedLeftStackGrowthCurveStartArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedLeftStackNormalVectors {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedLeftStackOfGrowthCurves {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedLeftStackOfNormalVectors {
		res = append(res, instance)
	}

	for instance := range stage.StackGrowthCurve2DEndHalfwayArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.StackGrowthCurve2DRibbonEndShapes {
		res = append(res, instance)
	}

	for instance := range stage.StackGrowthCurve2DRibbonStartShapes {
		res = append(res, instance)
	}

	for instance := range stage.StackGrowthCurve2DStartHalfwayArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.StackOfGrowthCurve2Ds {
		res = append(res, instance)
	}

	for instance := range stage.StackOfGrowthCurve2DRibbons {
		res = append(res, instance)
	}

	for instance := range stage.StackOfRotatedGrowthCurve2Ds {
		res = append(res, instance)
	}

	for instance := range stage.StackOfRotatedGrowthCurve2DRibbons {
		res = append(res, instance)
	}

	for instance := range stage.StackRotatedGrowthCurve2DEndArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.StackRotatedGrowthCurve2DRibbonEndShapes {
		res = append(res, instance)
	}

	for instance := range stage.StackRotatedGrowthCurve2DRibbonStartShapes {
		res = append(res, instance)
	}

	for instance := range stage.StackRotatedGrowthCurve2DStartArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.StartArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.StartArcShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.StartHalfwayArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.StartHalfwayArcShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.TopEndArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.TopEndArcShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.TopEndHalfwayArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.TopEndHalfwayArcShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.TopGrowthCurve2Ds {
		res = append(res, instance)
	}

	for instance := range stage.TopMidArcVectorShapes {
		res = append(res, instance)
	}

	for instance := range stage.TopMidArcVectorShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.TopStackGrowthCurve2DEndHalfwayArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.TopStackGrowthCurve2DStartHalfwayArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.TopStackOfGrowthCurve2Ds {
		res = append(res, instance)
	}

	for instance := range stage.TopStackOfRotatedGrowthCurve2Ds {
		res = append(res, instance)
	}

	for instance := range stage.TopStackOfRotatedGrowthCurve2DEndArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.TopStackOfRotatedGrowthCurve2DStartArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.TopStartArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.TopStartArcShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.TopStartHalfwayArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.TopStartHalfwayArcShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.TorusStackShapes {
		res = append(res, instance)
	}

	for instance := range stage.VerticalTorusStackShapes {
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

func (endhalfwayarcshape *EndHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(EndHalfwayArcShape)
	endhalfwayarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(EndHalfwayArcShapeGrid)
	endhalfwayarcshapegrid.CopyBasicFields(newInstance)
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

func (growthcurve2d *GrowthCurve2D) GongCopy() GongstructIF {
	newInstance := new(GrowthCurve2D)
	growthcurve2d.CopyBasicFields(newInstance)
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

func (midarcvectorshape *MidArcVectorShape) GongCopy() GongstructIF {
	newInstance := new(MidArcVectorShape)
	midarcvectorshape.CopyBasicFields(newInstance)
	return newInstance
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongCopy() GongstructIF {
	newInstance := new(MidArcVectorShapeGrid)
	midarcvectorshapegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DRibbon)
	partiallygrowthcurve2dribbon.CopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DRibbonEndShape)
	partiallygrowthcurve2dribbonendshape.CopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DRibbonStartShape)
	partiallygrowthcurve2dribbonstartshape.CopyBasicFields(newInstance)
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

func (rhombusstuff *RhombusStuff) GongCopy() GongstructIF {
	newInstance := new(RhombusStuff)
	rhombusstuff.CopyBasicFields(newInstance)
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

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongCopy() GongstructIF {
	newInstance := new(ShiftedBottomTopStartArcShape)
	shiftedbottomtopstartarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(ShiftedBottomTopStartArcShapeGrid)
	shiftedbottomtopstartarcshapegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftStackGrowthCurveEndArcShape)
	shiftedleftstackgrowthcurveendarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftStackGrowthCurveStartArcShape)
	shiftedleftstackgrowthcurvestartarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftStackNormalVector)
	shiftedleftstacknormalvector.CopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftStackOfGrowthCurve)
	shiftedleftstackofgrowthcurve.CopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftStackOfNormalVector)
	shiftedleftstackofnormalvector.CopyBasicFields(newInstance)
	return newInstance
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(StackGrowthCurve2DEndHalfwayArcShape)
	stackgrowthcurve2dendhalfwayarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongCopy() GongstructIF {
	newInstance := new(StackGrowthCurve2DRibbonEndShape)
	stackgrowthcurve2dribbonendshape.CopyBasicFields(newInstance)
	return newInstance
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongCopy() GongstructIF {
	newInstance := new(StackGrowthCurve2DRibbonStartShape)
	stackgrowthcurve2dribbonstartshape.CopyBasicFields(newInstance)
	return newInstance
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(StackGrowthCurve2DStartHalfwayArcShape)
	stackgrowthcurve2dstarthalfwayarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongCopy() GongstructIF {
	newInstance := new(StackOfGrowthCurve2D)
	stackofgrowthcurve2d.CopyBasicFields(newInstance)
	return newInstance
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongCopy() GongstructIF {
	newInstance := new(StackOfGrowthCurve2DRibbon)
	stackofgrowthcurve2dribbon.CopyBasicFields(newInstance)
	return newInstance
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongCopy() GongstructIF {
	newInstance := new(StackOfRotatedGrowthCurve2D)
	stackofrotatedgrowthcurve2d.CopyBasicFields(newInstance)
	return newInstance
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongCopy() GongstructIF {
	newInstance := new(StackOfRotatedGrowthCurve2DRibbon)
	stackofrotatedgrowthcurve2dribbon.CopyBasicFields(newInstance)
	return newInstance
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongCopy() GongstructIF {
	newInstance := new(StackRotatedGrowthCurve2DEndArcShape)
	stackrotatedgrowthcurve2dendarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongCopy() GongstructIF {
	newInstance := new(StackRotatedGrowthCurve2DRibbonEndShape)
	stackrotatedgrowthcurve2dribbonendshape.CopyBasicFields(newInstance)
	return newInstance
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongCopy() GongstructIF {
	newInstance := new(StackRotatedGrowthCurve2DRibbonStartShape)
	stackrotatedgrowthcurve2dribbonstartshape.CopyBasicFields(newInstance)
	return newInstance
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongCopy() GongstructIF {
	newInstance := new(StackRotatedGrowthCurve2DStartArcShape)
	stackrotatedgrowthcurve2dstartarcshape.CopyBasicFields(newInstance)
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

func (starthalfwayarcshape *StartHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(StartHalfwayArcShape)
	starthalfwayarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(StartHalfwayArcShapeGrid)
	starthalfwayarcshapegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (topendarcshape *TopEndArcShape) GongCopy() GongstructIF {
	newInstance := new(TopEndArcShape)
	topendarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(TopEndArcShapeGrid)
	topendarcshapegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(TopEndHalfwayArcShape)
	topendhalfwayarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(TopEndHalfwayArcShapeGrid)
	topendhalfwayarcshapegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongCopy() GongstructIF {
	newInstance := new(TopGrowthCurve2D)
	topgrowthcurve2d.CopyBasicFields(newInstance)
	return newInstance
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongCopy() GongstructIF {
	newInstance := new(TopMidArcVectorShape)
	topmidarcvectorshape.CopyBasicFields(newInstance)
	return newInstance
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongCopy() GongstructIF {
	newInstance := new(TopMidArcVectorShapeGrid)
	topmidarcvectorshapegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(TopStackGrowthCurve2DEndHalfwayArcShape)
	topstackgrowthcurve2dendhalfwayarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(TopStackGrowthCurve2DStartHalfwayArcShape)
	topstackgrowthcurve2dstarthalfwayarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongCopy() GongstructIF {
	newInstance := new(TopStackOfGrowthCurve2D)
	topstackofgrowthcurve2d.CopyBasicFields(newInstance)
	return newInstance
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongCopy() GongstructIF {
	newInstance := new(TopStackOfRotatedGrowthCurve2D)
	topstackofrotatedgrowthcurve2d.CopyBasicFields(newInstance)
	return newInstance
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongCopy() GongstructIF {
	newInstance := new(TopStackOfRotatedGrowthCurve2DEndArcShape)
	topstackofrotatedgrowthcurve2dendarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongCopy() GongstructIF {
	newInstance := new(TopStackOfRotatedGrowthCurve2DStartArcShape)
	topstackofrotatedgrowthcurve2dstartarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (topstartarcshape *TopStartArcShape) GongCopy() GongstructIF {
	newInstance := new(TopStartArcShape)
	topstartarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(TopStartArcShapeGrid)
	topstartarcshapegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(TopStartHalfwayArcShape)
	topstarthalfwayarcshape.CopyBasicFields(newInstance)
	return newInstance
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(TopStartHalfwayArcShapeGrid)
	topstarthalfwayarcshapegrid.CopyBasicFields(newInstance)
	return newInstance
}

func (torusstackshape *TorusStackShape) GongCopy() GongstructIF {
	newInstance := new(TorusStackShape)
	torusstackshape.CopyBasicFields(newInstance)
	return newInstance
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongCopy() GongstructIF {
	newInstance := new(VerticalTorusStackShape)
	verticaltorusstackshape.CopyBasicFields(newInstance)
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

func (endhalfwayarcshape *EndHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(endhalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(endhalfwayarcshape), uint64(GetOrderPointerGongstruct(stage, endhalfwayarcshape)))
	return
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(endhalfwayarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(endhalfwayarcshapegrid), uint64(GetOrderPointerGongstruct(stage, endhalfwayarcshapegrid)))
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

func (growthcurve2d *GrowthCurve2D) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(growthcurve2d).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(growthcurve2d), uint64(GetOrderPointerGongstruct(stage, growthcurve2d)))
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

func (midarcvectorshape *MidArcVectorShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(midarcvectorshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(midarcvectorshape), uint64(GetOrderPointerGongstruct(stage, midarcvectorshape)))
	return
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(midarcvectorshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(midarcvectorshapegrid), uint64(GetOrderPointerGongstruct(stage, midarcvectorshapegrid)))
	return
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dribbon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(partiallygrowthcurve2dribbon), uint64(GetOrderPointerGongstruct(stage, partiallygrowthcurve2dribbon)))
	return
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dribbonendshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(partiallygrowthcurve2dribbonendshape), uint64(GetOrderPointerGongstruct(stage, partiallygrowthcurve2dribbonendshape)))
	return
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dribbonstartshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(partiallygrowthcurve2dribbonstartshape), uint64(GetOrderPointerGongstruct(stage, partiallygrowthcurve2dribbonstartshape)))
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

func (rhombusstuff *RhombusStuff) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rhombusstuff).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(rhombusstuff), uint64(GetOrderPointerGongstruct(stage, rhombusstuff)))
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

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedbottomtopstartarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(shiftedbottomtopstartarcshape), uint64(GetOrderPointerGongstruct(stage, shiftedbottomtopstartarcshape)))
	return
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedbottomtopstartarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(shiftedbottomtopstartarcshapegrid), uint64(GetOrderPointerGongstruct(stage, shiftedbottomtopstartarcshapegrid)))
	return
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftstackgrowthcurveendarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(shiftedleftstackgrowthcurveendarcshape), uint64(GetOrderPointerGongstruct(stage, shiftedleftstackgrowthcurveendarcshape)))
	return
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftstackgrowthcurvestartarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(shiftedleftstackgrowthcurvestartarcshape), uint64(GetOrderPointerGongstruct(stage, shiftedleftstackgrowthcurvestartarcshape)))
	return
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftstacknormalvector).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(shiftedleftstacknormalvector), uint64(GetOrderPointerGongstruct(stage, shiftedleftstacknormalvector)))
	return
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftstackofgrowthcurve).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(shiftedleftstackofgrowthcurve), uint64(GetOrderPointerGongstruct(stage, shiftedleftstackofgrowthcurve)))
	return
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftstackofnormalvector).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(shiftedleftstackofnormalvector), uint64(GetOrderPointerGongstruct(stage, shiftedleftstackofnormalvector)))
	return
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackgrowthcurve2dendhalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackgrowthcurve2dendhalfwayarcshape), uint64(GetOrderPointerGongstruct(stage, stackgrowthcurve2dendhalfwayarcshape)))
	return
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackgrowthcurve2dribbonendshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackgrowthcurve2dribbonendshape), uint64(GetOrderPointerGongstruct(stage, stackgrowthcurve2dribbonendshape)))
	return
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackgrowthcurve2dribbonstartshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackgrowthcurve2dribbonstartshape), uint64(GetOrderPointerGongstruct(stage, stackgrowthcurve2dribbonstartshape)))
	return
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackgrowthcurve2dstarthalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackgrowthcurve2dstarthalfwayarcshape), uint64(GetOrderPointerGongstruct(stage, stackgrowthcurve2dstarthalfwayarcshape)))
	return
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofgrowthcurve2d).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackofgrowthcurve2d), uint64(GetOrderPointerGongstruct(stage, stackofgrowthcurve2d)))
	return
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofgrowthcurve2dribbon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackofgrowthcurve2dribbon), uint64(GetOrderPointerGongstruct(stage, stackofgrowthcurve2dribbon)))
	return
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofrotatedgrowthcurve2d).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackofrotatedgrowthcurve2d), uint64(GetOrderPointerGongstruct(stage, stackofrotatedgrowthcurve2d)))
	return
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofrotatedgrowthcurve2dribbon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackofrotatedgrowthcurve2dribbon), uint64(GetOrderPointerGongstruct(stage, stackofrotatedgrowthcurve2dribbon)))
	return
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackrotatedgrowthcurve2dendarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackrotatedgrowthcurve2dendarcshape), uint64(GetOrderPointerGongstruct(stage, stackrotatedgrowthcurve2dendarcshape)))
	return
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackrotatedgrowthcurve2dribbonendshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackrotatedgrowthcurve2dribbonendshape), uint64(GetOrderPointerGongstruct(stage, stackrotatedgrowthcurve2dribbonendshape)))
	return
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackrotatedgrowthcurve2dribbonstartshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackrotatedgrowthcurve2dribbonstartshape), uint64(GetOrderPointerGongstruct(stage, stackrotatedgrowthcurve2dribbonstartshape)))
	return
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackrotatedgrowthcurve2dstartarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stackrotatedgrowthcurve2dstartarcshape), uint64(GetOrderPointerGongstruct(stage, stackrotatedgrowthcurve2dstartarcshape)))
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

func (starthalfwayarcshape *StartHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(starthalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(starthalfwayarcshape), uint64(GetOrderPointerGongstruct(stage, starthalfwayarcshape)))
	return
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(starthalfwayarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(starthalfwayarcshapegrid), uint64(GetOrderPointerGongstruct(stage, starthalfwayarcshapegrid)))
	return
}

func (topendarcshape *TopEndArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topendarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topendarcshape), uint64(GetOrderPointerGongstruct(stage, topendarcshape)))
	return
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topendarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topendarcshapegrid), uint64(GetOrderPointerGongstruct(stage, topendarcshapegrid)))
	return
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topendhalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topendhalfwayarcshape), uint64(GetOrderPointerGongstruct(stage, topendhalfwayarcshape)))
	return
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topendhalfwayarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topendhalfwayarcshapegrid), uint64(GetOrderPointerGongstruct(stage, topendhalfwayarcshapegrid)))
	return
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topgrowthcurve2d).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topgrowthcurve2d), uint64(GetOrderPointerGongstruct(stage, topgrowthcurve2d)))
	return
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topmidarcvectorshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topmidarcvectorshape), uint64(GetOrderPointerGongstruct(stage, topmidarcvectorshape)))
	return
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topmidarcvectorshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topmidarcvectorshapegrid), uint64(GetOrderPointerGongstruct(stage, topmidarcvectorshapegrid)))
	return
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackgrowthcurve2dendhalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstackgrowthcurve2dendhalfwayarcshape), uint64(GetOrderPointerGongstruct(stage, topstackgrowthcurve2dendhalfwayarcshape)))
	return
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackgrowthcurve2dstarthalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstackgrowthcurve2dstarthalfwayarcshape), uint64(GetOrderPointerGongstruct(stage, topstackgrowthcurve2dstarthalfwayarcshape)))
	return
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackofgrowthcurve2d).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstackofgrowthcurve2d), uint64(GetOrderPointerGongstruct(stage, topstackofgrowthcurve2d)))
	return
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackofrotatedgrowthcurve2d).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstackofrotatedgrowthcurve2d), uint64(GetOrderPointerGongstruct(stage, topstackofrotatedgrowthcurve2d)))
	return
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackofrotatedgrowthcurve2dendarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstackofrotatedgrowthcurve2dendarcshape), uint64(GetOrderPointerGongstruct(stage, topstackofrotatedgrowthcurve2dendarcshape)))
	return
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackofrotatedgrowthcurve2dstartarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstackofrotatedgrowthcurve2dstartarcshape), uint64(GetOrderPointerGongstruct(stage, topstackofrotatedgrowthcurve2dstartarcshape)))
	return
}

func (topstartarcshape *TopStartArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstartarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstartarcshape), uint64(GetOrderPointerGongstruct(stage, topstartarcshape)))
	return
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstartarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstartarcshapegrid), uint64(GetOrderPointerGongstruct(stage, topstartarcshapegrid)))
	return
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstarthalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstarthalfwayarcshape), uint64(GetOrderPointerGongstruct(stage, topstarthalfwayarcshape)))
	return
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstarthalfwayarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(topstarthalfwayarcshapegrid), uint64(GetOrderPointerGongstruct(stage, topstarthalfwayarcshapegrid)))
	return
}

func (torusstackshape *TorusStackShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(torusstackshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(torusstackshape), uint64(GetOrderPointerGongstruct(stage, torusstackshape)))
	return
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(verticaltorusstackshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(verticaltorusstackshape), uint64(GetOrderPointerGongstruct(stage, verticaltorusstackshape)))
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
	var endhalfwayarcshapes_newInstances []*EndHalfwayArcShape
	var endhalfwayarcshapes_deletedInstances []*EndHalfwayArcShape

	// parse all staged instances and check if they have a reference
	for endhalfwayarcshape := range stage.EndHalfwayArcShapes {
		if ref, ok := stage.EndHalfwayArcShapes_reference[endhalfwayarcshape]; !ok {
			endhalfwayarcshapes_newInstances = append(endhalfwayarcshapes_newInstances, endhalfwayarcshape)
			newInstancesSlice = append(newInstancesSlice, endhalfwayarcshape.GongMarshallIdentifier(stage))
			if stage.EndHalfwayArcShapes_referenceOrder == nil {
				stage.EndHalfwayArcShapes_referenceOrder = make(map[*EndHalfwayArcShape]uint)
			}
			stage.EndHalfwayArcShapes_referenceOrder[endhalfwayarcshape] = stage.EndHalfwayArcShape_stagedOrder[endhalfwayarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, endhalfwayarcshape.GongMarshallUnstaging(stage))
			// delete(stage.EndHalfwayArcShapes_referenceOrder, endhalfwayarcshape)
			fieldInitializers, pointersInitializations := endhalfwayarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EndHalfwayArcShape_stagedOrder[ref] = stage.EndHalfwayArcShape_stagedOrder[endhalfwayarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := endhalfwayarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, endhalfwayarcshape)
			// delete(stage.EndHalfwayArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if endhalfwayarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", endhalfwayarcshape.GetName())
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
	for _, ref := range stage.EndHalfwayArcShapes_reference {
		instance := stage.EndHalfwayArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.EndHalfwayArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			endhalfwayarcshapes_deletedInstances = append(endhalfwayarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(endhalfwayarcshapes_newInstances)
	lenDeletedInstances += len(endhalfwayarcshapes_deletedInstances)
	var endhalfwayarcshapegrids_newInstances []*EndHalfwayArcShapeGrid
	var endhalfwayarcshapegrids_deletedInstances []*EndHalfwayArcShapeGrid

	// parse all staged instances and check if they have a reference
	for endhalfwayarcshapegrid := range stage.EndHalfwayArcShapeGrids {
		if ref, ok := stage.EndHalfwayArcShapeGrids_reference[endhalfwayarcshapegrid]; !ok {
			endhalfwayarcshapegrids_newInstances = append(endhalfwayarcshapegrids_newInstances, endhalfwayarcshapegrid)
			newInstancesSlice = append(newInstancesSlice, endhalfwayarcshapegrid.GongMarshallIdentifier(stage))
			if stage.EndHalfwayArcShapeGrids_referenceOrder == nil {
				stage.EndHalfwayArcShapeGrids_referenceOrder = make(map[*EndHalfwayArcShapeGrid]uint)
			}
			stage.EndHalfwayArcShapeGrids_referenceOrder[endhalfwayarcshapegrid] = stage.EndHalfwayArcShapeGrid_stagedOrder[endhalfwayarcshapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, endhalfwayarcshapegrid.GongMarshallUnstaging(stage))
			// delete(stage.EndHalfwayArcShapeGrids_referenceOrder, endhalfwayarcshapegrid)
			fieldInitializers, pointersInitializations := endhalfwayarcshapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EndHalfwayArcShapeGrid_stagedOrder[ref] = stage.EndHalfwayArcShapeGrid_stagedOrder[endhalfwayarcshapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := endhalfwayarcshapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, endhalfwayarcshapegrid)
			// delete(stage.EndHalfwayArcShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if endhalfwayarcshapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", endhalfwayarcshapegrid.GetName())
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
	for _, ref := range stage.EndHalfwayArcShapeGrids_reference {
		instance := stage.EndHalfwayArcShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.EndHalfwayArcShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			endhalfwayarcshapegrids_deletedInstances = append(endhalfwayarcshapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(endhalfwayarcshapegrids_newInstances)
	lenDeletedInstances += len(endhalfwayarcshapegrids_deletedInstances)
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
	var growthcurve2ds_newInstances []*GrowthCurve2D
	var growthcurve2ds_deletedInstances []*GrowthCurve2D

	// parse all staged instances and check if they have a reference
	for growthcurve2d := range stage.GrowthCurve2Ds {
		if ref, ok := stage.GrowthCurve2Ds_reference[growthcurve2d]; !ok {
			growthcurve2ds_newInstances = append(growthcurve2ds_newInstances, growthcurve2d)
			newInstancesSlice = append(newInstancesSlice, growthcurve2d.GongMarshallIdentifier(stage))
			if stage.GrowthCurve2Ds_referenceOrder == nil {
				stage.GrowthCurve2Ds_referenceOrder = make(map[*GrowthCurve2D]uint)
			}
			stage.GrowthCurve2Ds_referenceOrder[growthcurve2d] = stage.GrowthCurve2D_stagedOrder[growthcurve2d]
			newInstancesReverseSlice = append(newInstancesReverseSlice, growthcurve2d.GongMarshallUnstaging(stage))
			// delete(stage.GrowthCurve2Ds_referenceOrder, growthcurve2d)
			fieldInitializers, pointersInitializations := growthcurve2d.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GrowthCurve2D_stagedOrder[ref] = stage.GrowthCurve2D_stagedOrder[growthcurve2d]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := growthcurve2d.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, growthcurve2d)
			// delete(stage.GrowthCurve2D_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if growthcurve2d.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", growthcurve2d.GetName())
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
	for _, ref := range stage.GrowthCurve2Ds_reference {
		instance := stage.GrowthCurve2Ds_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GrowthCurve2Ds[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			growthcurve2ds_deletedInstances = append(growthcurve2ds_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(growthcurve2ds_newInstances)
	lenDeletedInstances += len(growthcurve2ds_deletedInstances)
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
	var midarcvectorshapes_newInstances []*MidArcVectorShape
	var midarcvectorshapes_deletedInstances []*MidArcVectorShape

	// parse all staged instances and check if they have a reference
	for midarcvectorshape := range stage.MidArcVectorShapes {
		if ref, ok := stage.MidArcVectorShapes_reference[midarcvectorshape]; !ok {
			midarcvectorshapes_newInstances = append(midarcvectorshapes_newInstances, midarcvectorshape)
			newInstancesSlice = append(newInstancesSlice, midarcvectorshape.GongMarshallIdentifier(stage))
			if stage.MidArcVectorShapes_referenceOrder == nil {
				stage.MidArcVectorShapes_referenceOrder = make(map[*MidArcVectorShape]uint)
			}
			stage.MidArcVectorShapes_referenceOrder[midarcvectorshape] = stage.MidArcVectorShape_stagedOrder[midarcvectorshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, midarcvectorshape.GongMarshallUnstaging(stage))
			// delete(stage.MidArcVectorShapes_referenceOrder, midarcvectorshape)
			fieldInitializers, pointersInitializations := midarcvectorshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MidArcVectorShape_stagedOrder[ref] = stage.MidArcVectorShape_stagedOrder[midarcvectorshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := midarcvectorshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, midarcvectorshape)
			// delete(stage.MidArcVectorShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if midarcvectorshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", midarcvectorshape.GetName())
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
	for _, ref := range stage.MidArcVectorShapes_reference {
		instance := stage.MidArcVectorShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.MidArcVectorShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			midarcvectorshapes_deletedInstances = append(midarcvectorshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(midarcvectorshapes_newInstances)
	lenDeletedInstances += len(midarcvectorshapes_deletedInstances)
	var midarcvectorshapegrids_newInstances []*MidArcVectorShapeGrid
	var midarcvectorshapegrids_deletedInstances []*MidArcVectorShapeGrid

	// parse all staged instances and check if they have a reference
	for midarcvectorshapegrid := range stage.MidArcVectorShapeGrids {
		if ref, ok := stage.MidArcVectorShapeGrids_reference[midarcvectorshapegrid]; !ok {
			midarcvectorshapegrids_newInstances = append(midarcvectorshapegrids_newInstances, midarcvectorshapegrid)
			newInstancesSlice = append(newInstancesSlice, midarcvectorshapegrid.GongMarshallIdentifier(stage))
			if stage.MidArcVectorShapeGrids_referenceOrder == nil {
				stage.MidArcVectorShapeGrids_referenceOrder = make(map[*MidArcVectorShapeGrid]uint)
			}
			stage.MidArcVectorShapeGrids_referenceOrder[midarcvectorshapegrid] = stage.MidArcVectorShapeGrid_stagedOrder[midarcvectorshapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, midarcvectorshapegrid.GongMarshallUnstaging(stage))
			// delete(stage.MidArcVectorShapeGrids_referenceOrder, midarcvectorshapegrid)
			fieldInitializers, pointersInitializations := midarcvectorshapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MidArcVectorShapeGrid_stagedOrder[ref] = stage.MidArcVectorShapeGrid_stagedOrder[midarcvectorshapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := midarcvectorshapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, midarcvectorshapegrid)
			// delete(stage.MidArcVectorShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if midarcvectorshapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", midarcvectorshapegrid.GetName())
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
	for _, ref := range stage.MidArcVectorShapeGrids_reference {
		instance := stage.MidArcVectorShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.MidArcVectorShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			midarcvectorshapegrids_deletedInstances = append(midarcvectorshapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(midarcvectorshapegrids_newInstances)
	lenDeletedInstances += len(midarcvectorshapegrids_deletedInstances)
	var partiallygrowthcurve2dribbons_newInstances []*PartiallyGrowthCurve2DRibbon
	var partiallygrowthcurve2dribbons_deletedInstances []*PartiallyGrowthCurve2DRibbon

	// parse all staged instances and check if they have a reference
	for partiallygrowthcurve2dribbon := range stage.PartiallyGrowthCurve2DRibbons {
		if ref, ok := stage.PartiallyGrowthCurve2DRibbons_reference[partiallygrowthcurve2dribbon]; !ok {
			partiallygrowthcurve2dribbons_newInstances = append(partiallygrowthcurve2dribbons_newInstances, partiallygrowthcurve2dribbon)
			newInstancesSlice = append(newInstancesSlice, partiallygrowthcurve2dribbon.GongMarshallIdentifier(stage))
			if stage.PartiallyGrowthCurve2DRibbons_referenceOrder == nil {
				stage.PartiallyGrowthCurve2DRibbons_referenceOrder = make(map[*PartiallyGrowthCurve2DRibbon]uint)
			}
			stage.PartiallyGrowthCurve2DRibbons_referenceOrder[partiallygrowthcurve2dribbon] = stage.PartiallyGrowthCurve2DRibbon_stagedOrder[partiallygrowthcurve2dribbon]
			newInstancesReverseSlice = append(newInstancesReverseSlice, partiallygrowthcurve2dribbon.GongMarshallUnstaging(stage))
			// delete(stage.PartiallyGrowthCurve2DRibbons_referenceOrder, partiallygrowthcurve2dribbon)
			fieldInitializers, pointersInitializations := partiallygrowthcurve2dribbon.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PartiallyGrowthCurve2DRibbon_stagedOrder[ref] = stage.PartiallyGrowthCurve2DRibbon_stagedOrder[partiallygrowthcurve2dribbon]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := partiallygrowthcurve2dribbon.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, partiallygrowthcurve2dribbon)
			// delete(stage.PartiallyGrowthCurve2DRibbon_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if partiallygrowthcurve2dribbon.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", partiallygrowthcurve2dribbon.GetName())
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
	for _, ref := range stage.PartiallyGrowthCurve2DRibbons_reference {
		instance := stage.PartiallyGrowthCurve2DRibbons_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PartiallyGrowthCurve2DRibbons[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			partiallygrowthcurve2dribbons_deletedInstances = append(partiallygrowthcurve2dribbons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(partiallygrowthcurve2dribbons_newInstances)
	lenDeletedInstances += len(partiallygrowthcurve2dribbons_deletedInstances)
	var partiallygrowthcurve2dribbonendshapes_newInstances []*PartiallyGrowthCurve2DRibbonEndShape
	var partiallygrowthcurve2dribbonendshapes_deletedInstances []*PartiallyGrowthCurve2DRibbonEndShape

	// parse all staged instances and check if they have a reference
	for partiallygrowthcurve2dribbonendshape := range stage.PartiallyGrowthCurve2DRibbonEndShapes {
		if ref, ok := stage.PartiallyGrowthCurve2DRibbonEndShapes_reference[partiallygrowthcurve2dribbonendshape]; !ok {
			partiallygrowthcurve2dribbonendshapes_newInstances = append(partiallygrowthcurve2dribbonendshapes_newInstances, partiallygrowthcurve2dribbonendshape)
			newInstancesSlice = append(newInstancesSlice, partiallygrowthcurve2dribbonendshape.GongMarshallIdentifier(stage))
			if stage.PartiallyGrowthCurve2DRibbonEndShapes_referenceOrder == nil {
				stage.PartiallyGrowthCurve2DRibbonEndShapes_referenceOrder = make(map[*PartiallyGrowthCurve2DRibbonEndShape]uint)
			}
			stage.PartiallyGrowthCurve2DRibbonEndShapes_referenceOrder[partiallygrowthcurve2dribbonendshape] = stage.PartiallyGrowthCurve2DRibbonEndShape_stagedOrder[partiallygrowthcurve2dribbonendshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, partiallygrowthcurve2dribbonendshape.GongMarshallUnstaging(stage))
			// delete(stage.PartiallyGrowthCurve2DRibbonEndShapes_referenceOrder, partiallygrowthcurve2dribbonendshape)
			fieldInitializers, pointersInitializations := partiallygrowthcurve2dribbonendshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PartiallyGrowthCurve2DRibbonEndShape_stagedOrder[ref] = stage.PartiallyGrowthCurve2DRibbonEndShape_stagedOrder[partiallygrowthcurve2dribbonendshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := partiallygrowthcurve2dribbonendshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, partiallygrowthcurve2dribbonendshape)
			// delete(stage.PartiallyGrowthCurve2DRibbonEndShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if partiallygrowthcurve2dribbonendshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", partiallygrowthcurve2dribbonendshape.GetName())
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
	for _, ref := range stage.PartiallyGrowthCurve2DRibbonEndShapes_reference {
		instance := stage.PartiallyGrowthCurve2DRibbonEndShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PartiallyGrowthCurve2DRibbonEndShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			partiallygrowthcurve2dribbonendshapes_deletedInstances = append(partiallygrowthcurve2dribbonendshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(partiallygrowthcurve2dribbonendshapes_newInstances)
	lenDeletedInstances += len(partiallygrowthcurve2dribbonendshapes_deletedInstances)
	var partiallygrowthcurve2dribbonstartshapes_newInstances []*PartiallyGrowthCurve2DRibbonStartShape
	var partiallygrowthcurve2dribbonstartshapes_deletedInstances []*PartiallyGrowthCurve2DRibbonStartShape

	// parse all staged instances and check if they have a reference
	for partiallygrowthcurve2dribbonstartshape := range stage.PartiallyGrowthCurve2DRibbonStartShapes {
		if ref, ok := stage.PartiallyGrowthCurve2DRibbonStartShapes_reference[partiallygrowthcurve2dribbonstartshape]; !ok {
			partiallygrowthcurve2dribbonstartshapes_newInstances = append(partiallygrowthcurve2dribbonstartshapes_newInstances, partiallygrowthcurve2dribbonstartshape)
			newInstancesSlice = append(newInstancesSlice, partiallygrowthcurve2dribbonstartshape.GongMarshallIdentifier(stage))
			if stage.PartiallyGrowthCurve2DRibbonStartShapes_referenceOrder == nil {
				stage.PartiallyGrowthCurve2DRibbonStartShapes_referenceOrder = make(map[*PartiallyGrowthCurve2DRibbonStartShape]uint)
			}
			stage.PartiallyGrowthCurve2DRibbonStartShapes_referenceOrder[partiallygrowthcurve2dribbonstartshape] = stage.PartiallyGrowthCurve2DRibbonStartShape_stagedOrder[partiallygrowthcurve2dribbonstartshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, partiallygrowthcurve2dribbonstartshape.GongMarshallUnstaging(stage))
			// delete(stage.PartiallyGrowthCurve2DRibbonStartShapes_referenceOrder, partiallygrowthcurve2dribbonstartshape)
			fieldInitializers, pointersInitializations := partiallygrowthcurve2dribbonstartshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PartiallyGrowthCurve2DRibbonStartShape_stagedOrder[ref] = stage.PartiallyGrowthCurve2DRibbonStartShape_stagedOrder[partiallygrowthcurve2dribbonstartshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := partiallygrowthcurve2dribbonstartshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, partiallygrowthcurve2dribbonstartshape)
			// delete(stage.PartiallyGrowthCurve2DRibbonStartShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if partiallygrowthcurve2dribbonstartshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", partiallygrowthcurve2dribbonstartshape.GetName())
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
	for _, ref := range stage.PartiallyGrowthCurve2DRibbonStartShapes_reference {
		instance := stage.PartiallyGrowthCurve2DRibbonStartShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PartiallyGrowthCurve2DRibbonStartShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			partiallygrowthcurve2dribbonstartshapes_deletedInstances = append(partiallygrowthcurve2dribbonstartshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(partiallygrowthcurve2dribbonstartshapes_newInstances)
	lenDeletedInstances += len(partiallygrowthcurve2dribbonstartshapes_deletedInstances)
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
	var rhombusstuffs_newInstances []*RhombusStuff
	var rhombusstuffs_deletedInstances []*RhombusStuff

	// parse all staged instances and check if they have a reference
	for rhombusstuff := range stage.RhombusStuffs {
		if ref, ok := stage.RhombusStuffs_reference[rhombusstuff]; !ok {
			rhombusstuffs_newInstances = append(rhombusstuffs_newInstances, rhombusstuff)
			newInstancesSlice = append(newInstancesSlice, rhombusstuff.GongMarshallIdentifier(stage))
			if stage.RhombusStuffs_referenceOrder == nil {
				stage.RhombusStuffs_referenceOrder = make(map[*RhombusStuff]uint)
			}
			stage.RhombusStuffs_referenceOrder[rhombusstuff] = stage.RhombusStuff_stagedOrder[rhombusstuff]
			newInstancesReverseSlice = append(newInstancesReverseSlice, rhombusstuff.GongMarshallUnstaging(stage))
			// delete(stage.RhombusStuffs_referenceOrder, rhombusstuff)
			fieldInitializers, pointersInitializations := rhombusstuff.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RhombusStuff_stagedOrder[ref] = stage.RhombusStuff_stagedOrder[rhombusstuff]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := rhombusstuff.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, rhombusstuff)
			// delete(stage.RhombusStuff_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if rhombusstuff.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", rhombusstuff.GetName())
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
	for _, ref := range stage.RhombusStuffs_reference {
		instance := stage.RhombusStuffs_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.RhombusStuffs[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			rhombusstuffs_deletedInstances = append(rhombusstuffs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(rhombusstuffs_newInstances)
	lenDeletedInstances += len(rhombusstuffs_deletedInstances)
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
	var shiftedbottomtopstartarcshapes_newInstances []*ShiftedBottomTopStartArcShape
	var shiftedbottomtopstartarcshapes_deletedInstances []*ShiftedBottomTopStartArcShape

	// parse all staged instances and check if they have a reference
	for shiftedbottomtopstartarcshape := range stage.ShiftedBottomTopStartArcShapes {
		if ref, ok := stage.ShiftedBottomTopStartArcShapes_reference[shiftedbottomtopstartarcshape]; !ok {
			shiftedbottomtopstartarcshapes_newInstances = append(shiftedbottomtopstartarcshapes_newInstances, shiftedbottomtopstartarcshape)
			newInstancesSlice = append(newInstancesSlice, shiftedbottomtopstartarcshape.GongMarshallIdentifier(stage))
			if stage.ShiftedBottomTopStartArcShapes_referenceOrder == nil {
				stage.ShiftedBottomTopStartArcShapes_referenceOrder = make(map[*ShiftedBottomTopStartArcShape]uint)
			}
			stage.ShiftedBottomTopStartArcShapes_referenceOrder[shiftedbottomtopstartarcshape] = stage.ShiftedBottomTopStartArcShape_stagedOrder[shiftedbottomtopstartarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, shiftedbottomtopstartarcshape.GongMarshallUnstaging(stage))
			// delete(stage.ShiftedBottomTopStartArcShapes_referenceOrder, shiftedbottomtopstartarcshape)
			fieldInitializers, pointersInitializations := shiftedbottomtopstartarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ShiftedBottomTopStartArcShape_stagedOrder[ref] = stage.ShiftedBottomTopStartArcShape_stagedOrder[shiftedbottomtopstartarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := shiftedbottomtopstartarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, shiftedbottomtopstartarcshape)
			// delete(stage.ShiftedBottomTopStartArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if shiftedbottomtopstartarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", shiftedbottomtopstartarcshape.GetName())
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
	for _, ref := range stage.ShiftedBottomTopStartArcShapes_reference {
		instance := stage.ShiftedBottomTopStartArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ShiftedBottomTopStartArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			shiftedbottomtopstartarcshapes_deletedInstances = append(shiftedbottomtopstartarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(shiftedbottomtopstartarcshapes_newInstances)
	lenDeletedInstances += len(shiftedbottomtopstartarcshapes_deletedInstances)
	var shiftedbottomtopstartarcshapegrids_newInstances []*ShiftedBottomTopStartArcShapeGrid
	var shiftedbottomtopstartarcshapegrids_deletedInstances []*ShiftedBottomTopStartArcShapeGrid

	// parse all staged instances and check if they have a reference
	for shiftedbottomtopstartarcshapegrid := range stage.ShiftedBottomTopStartArcShapeGrids {
		if ref, ok := stage.ShiftedBottomTopStartArcShapeGrids_reference[shiftedbottomtopstartarcshapegrid]; !ok {
			shiftedbottomtopstartarcshapegrids_newInstances = append(shiftedbottomtopstartarcshapegrids_newInstances, shiftedbottomtopstartarcshapegrid)
			newInstancesSlice = append(newInstancesSlice, shiftedbottomtopstartarcshapegrid.GongMarshallIdentifier(stage))
			if stage.ShiftedBottomTopStartArcShapeGrids_referenceOrder == nil {
				stage.ShiftedBottomTopStartArcShapeGrids_referenceOrder = make(map[*ShiftedBottomTopStartArcShapeGrid]uint)
			}
			stage.ShiftedBottomTopStartArcShapeGrids_referenceOrder[shiftedbottomtopstartarcshapegrid] = stage.ShiftedBottomTopStartArcShapeGrid_stagedOrder[shiftedbottomtopstartarcshapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, shiftedbottomtopstartarcshapegrid.GongMarshallUnstaging(stage))
			// delete(stage.ShiftedBottomTopStartArcShapeGrids_referenceOrder, shiftedbottomtopstartarcshapegrid)
			fieldInitializers, pointersInitializations := shiftedbottomtopstartarcshapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ShiftedBottomTopStartArcShapeGrid_stagedOrder[ref] = stage.ShiftedBottomTopStartArcShapeGrid_stagedOrder[shiftedbottomtopstartarcshapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := shiftedbottomtopstartarcshapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, shiftedbottomtopstartarcshapegrid)
			// delete(stage.ShiftedBottomTopStartArcShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if shiftedbottomtopstartarcshapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", shiftedbottomtopstartarcshapegrid.GetName())
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
	for _, ref := range stage.ShiftedBottomTopStartArcShapeGrids_reference {
		instance := stage.ShiftedBottomTopStartArcShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ShiftedBottomTopStartArcShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			shiftedbottomtopstartarcshapegrids_deletedInstances = append(shiftedbottomtopstartarcshapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(shiftedbottomtopstartarcshapegrids_newInstances)
	lenDeletedInstances += len(shiftedbottomtopstartarcshapegrids_deletedInstances)
	var shiftedleftstackgrowthcurveendarcshapes_newInstances []*ShiftedLeftStackGrowthCurveEndArcShape
	var shiftedleftstackgrowthcurveendarcshapes_deletedInstances []*ShiftedLeftStackGrowthCurveEndArcShape

	// parse all staged instances and check if they have a reference
	for shiftedleftstackgrowthcurveendarcshape := range stage.ShiftedLeftStackGrowthCurveEndArcShapes {
		if ref, ok := stage.ShiftedLeftStackGrowthCurveEndArcShapes_reference[shiftedleftstackgrowthcurveendarcshape]; !ok {
			shiftedleftstackgrowthcurveendarcshapes_newInstances = append(shiftedleftstackgrowthcurveendarcshapes_newInstances, shiftedleftstackgrowthcurveendarcshape)
			newInstancesSlice = append(newInstancesSlice, shiftedleftstackgrowthcurveendarcshape.GongMarshallIdentifier(stage))
			if stage.ShiftedLeftStackGrowthCurveEndArcShapes_referenceOrder == nil {
				stage.ShiftedLeftStackGrowthCurveEndArcShapes_referenceOrder = make(map[*ShiftedLeftStackGrowthCurveEndArcShape]uint)
			}
			stage.ShiftedLeftStackGrowthCurveEndArcShapes_referenceOrder[shiftedleftstackgrowthcurveendarcshape] = stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder[shiftedleftstackgrowthcurveendarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, shiftedleftstackgrowthcurveendarcshape.GongMarshallUnstaging(stage))
			// delete(stage.ShiftedLeftStackGrowthCurveEndArcShapes_referenceOrder, shiftedleftstackgrowthcurveendarcshape)
			fieldInitializers, pointersInitializations := shiftedleftstackgrowthcurveendarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder[ref] = stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder[shiftedleftstackgrowthcurveendarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := shiftedleftstackgrowthcurveendarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, shiftedleftstackgrowthcurveendarcshape)
			// delete(stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if shiftedleftstackgrowthcurveendarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", shiftedleftstackgrowthcurveendarcshape.GetName())
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
	for _, ref := range stage.ShiftedLeftStackGrowthCurveEndArcShapes_reference {
		instance := stage.ShiftedLeftStackGrowthCurveEndArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ShiftedLeftStackGrowthCurveEndArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			shiftedleftstackgrowthcurveendarcshapes_deletedInstances = append(shiftedleftstackgrowthcurveendarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(shiftedleftstackgrowthcurveendarcshapes_newInstances)
	lenDeletedInstances += len(shiftedleftstackgrowthcurveendarcshapes_deletedInstances)
	var shiftedleftstackgrowthcurvestartarcshapes_newInstances []*ShiftedLeftStackGrowthCurveStartArcShape
	var shiftedleftstackgrowthcurvestartarcshapes_deletedInstances []*ShiftedLeftStackGrowthCurveStartArcShape

	// parse all staged instances and check if they have a reference
	for shiftedleftstackgrowthcurvestartarcshape := range stage.ShiftedLeftStackGrowthCurveStartArcShapes {
		if ref, ok := stage.ShiftedLeftStackGrowthCurveStartArcShapes_reference[shiftedleftstackgrowthcurvestartarcshape]; !ok {
			shiftedleftstackgrowthcurvestartarcshapes_newInstances = append(shiftedleftstackgrowthcurvestartarcshapes_newInstances, shiftedleftstackgrowthcurvestartarcshape)
			newInstancesSlice = append(newInstancesSlice, shiftedleftstackgrowthcurvestartarcshape.GongMarshallIdentifier(stage))
			if stage.ShiftedLeftStackGrowthCurveStartArcShapes_referenceOrder == nil {
				stage.ShiftedLeftStackGrowthCurveStartArcShapes_referenceOrder = make(map[*ShiftedLeftStackGrowthCurveStartArcShape]uint)
			}
			stage.ShiftedLeftStackGrowthCurveStartArcShapes_referenceOrder[shiftedleftstackgrowthcurvestartarcshape] = stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder[shiftedleftstackgrowthcurvestartarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, shiftedleftstackgrowthcurvestartarcshape.GongMarshallUnstaging(stage))
			// delete(stage.ShiftedLeftStackGrowthCurveStartArcShapes_referenceOrder, shiftedleftstackgrowthcurvestartarcshape)
			fieldInitializers, pointersInitializations := shiftedleftstackgrowthcurvestartarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder[ref] = stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder[shiftedleftstackgrowthcurvestartarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := shiftedleftstackgrowthcurvestartarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, shiftedleftstackgrowthcurvestartarcshape)
			// delete(stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if shiftedleftstackgrowthcurvestartarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", shiftedleftstackgrowthcurvestartarcshape.GetName())
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
	for _, ref := range stage.ShiftedLeftStackGrowthCurveStartArcShapes_reference {
		instance := stage.ShiftedLeftStackGrowthCurveStartArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ShiftedLeftStackGrowthCurveStartArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			shiftedleftstackgrowthcurvestartarcshapes_deletedInstances = append(shiftedleftstackgrowthcurvestartarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(shiftedleftstackgrowthcurvestartarcshapes_newInstances)
	lenDeletedInstances += len(shiftedleftstackgrowthcurvestartarcshapes_deletedInstances)
	var shiftedleftstacknormalvectors_newInstances []*ShiftedLeftStackNormalVector
	var shiftedleftstacknormalvectors_deletedInstances []*ShiftedLeftStackNormalVector

	// parse all staged instances and check if they have a reference
	for shiftedleftstacknormalvector := range stage.ShiftedLeftStackNormalVectors {
		if ref, ok := stage.ShiftedLeftStackNormalVectors_reference[shiftedleftstacknormalvector]; !ok {
			shiftedleftstacknormalvectors_newInstances = append(shiftedleftstacknormalvectors_newInstances, shiftedleftstacknormalvector)
			newInstancesSlice = append(newInstancesSlice, shiftedleftstacknormalvector.GongMarshallIdentifier(stage))
			if stage.ShiftedLeftStackNormalVectors_referenceOrder == nil {
				stage.ShiftedLeftStackNormalVectors_referenceOrder = make(map[*ShiftedLeftStackNormalVector]uint)
			}
			stage.ShiftedLeftStackNormalVectors_referenceOrder[shiftedleftstacknormalvector] = stage.ShiftedLeftStackNormalVector_stagedOrder[shiftedleftstacknormalvector]
			newInstancesReverseSlice = append(newInstancesReverseSlice, shiftedleftstacknormalvector.GongMarshallUnstaging(stage))
			// delete(stage.ShiftedLeftStackNormalVectors_referenceOrder, shiftedleftstacknormalvector)
			fieldInitializers, pointersInitializations := shiftedleftstacknormalvector.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ShiftedLeftStackNormalVector_stagedOrder[ref] = stage.ShiftedLeftStackNormalVector_stagedOrder[shiftedleftstacknormalvector]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := shiftedleftstacknormalvector.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, shiftedleftstacknormalvector)
			// delete(stage.ShiftedLeftStackNormalVector_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if shiftedleftstacknormalvector.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", shiftedleftstacknormalvector.GetName())
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
	for _, ref := range stage.ShiftedLeftStackNormalVectors_reference {
		instance := stage.ShiftedLeftStackNormalVectors_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ShiftedLeftStackNormalVectors[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			shiftedleftstacknormalvectors_deletedInstances = append(shiftedleftstacknormalvectors_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(shiftedleftstacknormalvectors_newInstances)
	lenDeletedInstances += len(shiftedleftstacknormalvectors_deletedInstances)
	var shiftedleftstackofgrowthcurves_newInstances []*ShiftedLeftStackOfGrowthCurve
	var shiftedleftstackofgrowthcurves_deletedInstances []*ShiftedLeftStackOfGrowthCurve

	// parse all staged instances and check if they have a reference
	for shiftedleftstackofgrowthcurve := range stage.ShiftedLeftStackOfGrowthCurves {
		if ref, ok := stage.ShiftedLeftStackOfGrowthCurves_reference[shiftedleftstackofgrowthcurve]; !ok {
			shiftedleftstackofgrowthcurves_newInstances = append(shiftedleftstackofgrowthcurves_newInstances, shiftedleftstackofgrowthcurve)
			newInstancesSlice = append(newInstancesSlice, shiftedleftstackofgrowthcurve.GongMarshallIdentifier(stage))
			if stage.ShiftedLeftStackOfGrowthCurves_referenceOrder == nil {
				stage.ShiftedLeftStackOfGrowthCurves_referenceOrder = make(map[*ShiftedLeftStackOfGrowthCurve]uint)
			}
			stage.ShiftedLeftStackOfGrowthCurves_referenceOrder[shiftedleftstackofgrowthcurve] = stage.ShiftedLeftStackOfGrowthCurve_stagedOrder[shiftedleftstackofgrowthcurve]
			newInstancesReverseSlice = append(newInstancesReverseSlice, shiftedleftstackofgrowthcurve.GongMarshallUnstaging(stage))
			// delete(stage.ShiftedLeftStackOfGrowthCurves_referenceOrder, shiftedleftstackofgrowthcurve)
			fieldInitializers, pointersInitializations := shiftedleftstackofgrowthcurve.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ShiftedLeftStackOfGrowthCurve_stagedOrder[ref] = stage.ShiftedLeftStackOfGrowthCurve_stagedOrder[shiftedleftstackofgrowthcurve]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := shiftedleftstackofgrowthcurve.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, shiftedleftstackofgrowthcurve)
			// delete(stage.ShiftedLeftStackOfGrowthCurve_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if shiftedleftstackofgrowthcurve.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", shiftedleftstackofgrowthcurve.GetName())
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
	for _, ref := range stage.ShiftedLeftStackOfGrowthCurves_reference {
		instance := stage.ShiftedLeftStackOfGrowthCurves_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ShiftedLeftStackOfGrowthCurves[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			shiftedleftstackofgrowthcurves_deletedInstances = append(shiftedleftstackofgrowthcurves_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(shiftedleftstackofgrowthcurves_newInstances)
	lenDeletedInstances += len(shiftedleftstackofgrowthcurves_deletedInstances)
	var shiftedleftstackofnormalvectors_newInstances []*ShiftedLeftStackOfNormalVector
	var shiftedleftstackofnormalvectors_deletedInstances []*ShiftedLeftStackOfNormalVector

	// parse all staged instances and check if they have a reference
	for shiftedleftstackofnormalvector := range stage.ShiftedLeftStackOfNormalVectors {
		if ref, ok := stage.ShiftedLeftStackOfNormalVectors_reference[shiftedleftstackofnormalvector]; !ok {
			shiftedleftstackofnormalvectors_newInstances = append(shiftedleftstackofnormalvectors_newInstances, shiftedleftstackofnormalvector)
			newInstancesSlice = append(newInstancesSlice, shiftedleftstackofnormalvector.GongMarshallIdentifier(stage))
			if stage.ShiftedLeftStackOfNormalVectors_referenceOrder == nil {
				stage.ShiftedLeftStackOfNormalVectors_referenceOrder = make(map[*ShiftedLeftStackOfNormalVector]uint)
			}
			stage.ShiftedLeftStackOfNormalVectors_referenceOrder[shiftedleftstackofnormalvector] = stage.ShiftedLeftStackOfNormalVector_stagedOrder[shiftedleftstackofnormalvector]
			newInstancesReverseSlice = append(newInstancesReverseSlice, shiftedleftstackofnormalvector.GongMarshallUnstaging(stage))
			// delete(stage.ShiftedLeftStackOfNormalVectors_referenceOrder, shiftedleftstackofnormalvector)
			fieldInitializers, pointersInitializations := shiftedleftstackofnormalvector.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ShiftedLeftStackOfNormalVector_stagedOrder[ref] = stage.ShiftedLeftStackOfNormalVector_stagedOrder[shiftedleftstackofnormalvector]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := shiftedleftstackofnormalvector.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, shiftedleftstackofnormalvector)
			// delete(stage.ShiftedLeftStackOfNormalVector_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if shiftedleftstackofnormalvector.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", shiftedleftstackofnormalvector.GetName())
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
	for _, ref := range stage.ShiftedLeftStackOfNormalVectors_reference {
		instance := stage.ShiftedLeftStackOfNormalVectors_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ShiftedLeftStackOfNormalVectors[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			shiftedleftstackofnormalvectors_deletedInstances = append(shiftedleftstackofnormalvectors_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(shiftedleftstackofnormalvectors_newInstances)
	lenDeletedInstances += len(shiftedleftstackofnormalvectors_deletedInstances)
	var stackgrowthcurve2dendhalfwayarcshapes_newInstances []*StackGrowthCurve2DEndHalfwayArcShape
	var stackgrowthcurve2dendhalfwayarcshapes_deletedInstances []*StackGrowthCurve2DEndHalfwayArcShape

	// parse all staged instances and check if they have a reference
	for stackgrowthcurve2dendhalfwayarcshape := range stage.StackGrowthCurve2DEndHalfwayArcShapes {
		if ref, ok := stage.StackGrowthCurve2DEndHalfwayArcShapes_reference[stackgrowthcurve2dendhalfwayarcshape]; !ok {
			stackgrowthcurve2dendhalfwayarcshapes_newInstances = append(stackgrowthcurve2dendhalfwayarcshapes_newInstances, stackgrowthcurve2dendhalfwayarcshape)
			newInstancesSlice = append(newInstancesSlice, stackgrowthcurve2dendhalfwayarcshape.GongMarshallIdentifier(stage))
			if stage.StackGrowthCurve2DEndHalfwayArcShapes_referenceOrder == nil {
				stage.StackGrowthCurve2DEndHalfwayArcShapes_referenceOrder = make(map[*StackGrowthCurve2DEndHalfwayArcShape]uint)
			}
			stage.StackGrowthCurve2DEndHalfwayArcShapes_referenceOrder[stackgrowthcurve2dendhalfwayarcshape] = stage.StackGrowthCurve2DEndHalfwayArcShape_stagedOrder[stackgrowthcurve2dendhalfwayarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackgrowthcurve2dendhalfwayarcshape.GongMarshallUnstaging(stage))
			// delete(stage.StackGrowthCurve2DEndHalfwayArcShapes_referenceOrder, stackgrowthcurve2dendhalfwayarcshape)
			fieldInitializers, pointersInitializations := stackgrowthcurve2dendhalfwayarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackGrowthCurve2DEndHalfwayArcShape_stagedOrder[ref] = stage.StackGrowthCurve2DEndHalfwayArcShape_stagedOrder[stackgrowthcurve2dendhalfwayarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackgrowthcurve2dendhalfwayarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackgrowthcurve2dendhalfwayarcshape)
			// delete(stage.StackGrowthCurve2DEndHalfwayArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackgrowthcurve2dendhalfwayarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackgrowthcurve2dendhalfwayarcshape.GetName())
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
	for _, ref := range stage.StackGrowthCurve2DEndHalfwayArcShapes_reference {
		instance := stage.StackGrowthCurve2DEndHalfwayArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackGrowthCurve2DEndHalfwayArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackgrowthcurve2dendhalfwayarcshapes_deletedInstances = append(stackgrowthcurve2dendhalfwayarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackgrowthcurve2dendhalfwayarcshapes_newInstances)
	lenDeletedInstances += len(stackgrowthcurve2dendhalfwayarcshapes_deletedInstances)
	var stackgrowthcurve2dribbonendshapes_newInstances []*StackGrowthCurve2DRibbonEndShape
	var stackgrowthcurve2dribbonendshapes_deletedInstances []*StackGrowthCurve2DRibbonEndShape

	// parse all staged instances and check if they have a reference
	for stackgrowthcurve2dribbonendshape := range stage.StackGrowthCurve2DRibbonEndShapes {
		if ref, ok := stage.StackGrowthCurve2DRibbonEndShapes_reference[stackgrowthcurve2dribbonendshape]; !ok {
			stackgrowthcurve2dribbonendshapes_newInstances = append(stackgrowthcurve2dribbonendshapes_newInstances, stackgrowthcurve2dribbonendshape)
			newInstancesSlice = append(newInstancesSlice, stackgrowthcurve2dribbonendshape.GongMarshallIdentifier(stage))
			if stage.StackGrowthCurve2DRibbonEndShapes_referenceOrder == nil {
				stage.StackGrowthCurve2DRibbonEndShapes_referenceOrder = make(map[*StackGrowthCurve2DRibbonEndShape]uint)
			}
			stage.StackGrowthCurve2DRibbonEndShapes_referenceOrder[stackgrowthcurve2dribbonendshape] = stage.StackGrowthCurve2DRibbonEndShape_stagedOrder[stackgrowthcurve2dribbonendshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackgrowthcurve2dribbonendshape.GongMarshallUnstaging(stage))
			// delete(stage.StackGrowthCurve2DRibbonEndShapes_referenceOrder, stackgrowthcurve2dribbonendshape)
			fieldInitializers, pointersInitializations := stackgrowthcurve2dribbonendshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackGrowthCurve2DRibbonEndShape_stagedOrder[ref] = stage.StackGrowthCurve2DRibbonEndShape_stagedOrder[stackgrowthcurve2dribbonendshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackgrowthcurve2dribbonendshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackgrowthcurve2dribbonendshape)
			// delete(stage.StackGrowthCurve2DRibbonEndShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackgrowthcurve2dribbonendshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackgrowthcurve2dribbonendshape.GetName())
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
	for _, ref := range stage.StackGrowthCurve2DRibbonEndShapes_reference {
		instance := stage.StackGrowthCurve2DRibbonEndShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackGrowthCurve2DRibbonEndShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackgrowthcurve2dribbonendshapes_deletedInstances = append(stackgrowthcurve2dribbonendshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackgrowthcurve2dribbonendshapes_newInstances)
	lenDeletedInstances += len(stackgrowthcurve2dribbonendshapes_deletedInstances)
	var stackgrowthcurve2dribbonstartshapes_newInstances []*StackGrowthCurve2DRibbonStartShape
	var stackgrowthcurve2dribbonstartshapes_deletedInstances []*StackGrowthCurve2DRibbonStartShape

	// parse all staged instances and check if they have a reference
	for stackgrowthcurve2dribbonstartshape := range stage.StackGrowthCurve2DRibbonStartShapes {
		if ref, ok := stage.StackGrowthCurve2DRibbonStartShapes_reference[stackgrowthcurve2dribbonstartshape]; !ok {
			stackgrowthcurve2dribbonstartshapes_newInstances = append(stackgrowthcurve2dribbonstartshapes_newInstances, stackgrowthcurve2dribbonstartshape)
			newInstancesSlice = append(newInstancesSlice, stackgrowthcurve2dribbonstartshape.GongMarshallIdentifier(stage))
			if stage.StackGrowthCurve2DRibbonStartShapes_referenceOrder == nil {
				stage.StackGrowthCurve2DRibbonStartShapes_referenceOrder = make(map[*StackGrowthCurve2DRibbonStartShape]uint)
			}
			stage.StackGrowthCurve2DRibbonStartShapes_referenceOrder[stackgrowthcurve2dribbonstartshape] = stage.StackGrowthCurve2DRibbonStartShape_stagedOrder[stackgrowthcurve2dribbonstartshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackgrowthcurve2dribbonstartshape.GongMarshallUnstaging(stage))
			// delete(stage.StackGrowthCurve2DRibbonStartShapes_referenceOrder, stackgrowthcurve2dribbonstartshape)
			fieldInitializers, pointersInitializations := stackgrowthcurve2dribbonstartshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackGrowthCurve2DRibbonStartShape_stagedOrder[ref] = stage.StackGrowthCurve2DRibbonStartShape_stagedOrder[stackgrowthcurve2dribbonstartshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackgrowthcurve2dribbonstartshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackgrowthcurve2dribbonstartshape)
			// delete(stage.StackGrowthCurve2DRibbonStartShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackgrowthcurve2dribbonstartshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackgrowthcurve2dribbonstartshape.GetName())
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
	for _, ref := range stage.StackGrowthCurve2DRibbonStartShapes_reference {
		instance := stage.StackGrowthCurve2DRibbonStartShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackGrowthCurve2DRibbonStartShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackgrowthcurve2dribbonstartshapes_deletedInstances = append(stackgrowthcurve2dribbonstartshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackgrowthcurve2dribbonstartshapes_newInstances)
	lenDeletedInstances += len(stackgrowthcurve2dribbonstartshapes_deletedInstances)
	var stackgrowthcurve2dstarthalfwayarcshapes_newInstances []*StackGrowthCurve2DStartHalfwayArcShape
	var stackgrowthcurve2dstarthalfwayarcshapes_deletedInstances []*StackGrowthCurve2DStartHalfwayArcShape

	// parse all staged instances and check if they have a reference
	for stackgrowthcurve2dstarthalfwayarcshape := range stage.StackGrowthCurve2DStartHalfwayArcShapes {
		if ref, ok := stage.StackGrowthCurve2DStartHalfwayArcShapes_reference[stackgrowthcurve2dstarthalfwayarcshape]; !ok {
			stackgrowthcurve2dstarthalfwayarcshapes_newInstances = append(stackgrowthcurve2dstarthalfwayarcshapes_newInstances, stackgrowthcurve2dstarthalfwayarcshape)
			newInstancesSlice = append(newInstancesSlice, stackgrowthcurve2dstarthalfwayarcshape.GongMarshallIdentifier(stage))
			if stage.StackGrowthCurve2DStartHalfwayArcShapes_referenceOrder == nil {
				stage.StackGrowthCurve2DStartHalfwayArcShapes_referenceOrder = make(map[*StackGrowthCurve2DStartHalfwayArcShape]uint)
			}
			stage.StackGrowthCurve2DStartHalfwayArcShapes_referenceOrder[stackgrowthcurve2dstarthalfwayarcshape] = stage.StackGrowthCurve2DStartHalfwayArcShape_stagedOrder[stackgrowthcurve2dstarthalfwayarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackgrowthcurve2dstarthalfwayarcshape.GongMarshallUnstaging(stage))
			// delete(stage.StackGrowthCurve2DStartHalfwayArcShapes_referenceOrder, stackgrowthcurve2dstarthalfwayarcshape)
			fieldInitializers, pointersInitializations := stackgrowthcurve2dstarthalfwayarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackGrowthCurve2DStartHalfwayArcShape_stagedOrder[ref] = stage.StackGrowthCurve2DStartHalfwayArcShape_stagedOrder[stackgrowthcurve2dstarthalfwayarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackgrowthcurve2dstarthalfwayarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackgrowthcurve2dstarthalfwayarcshape)
			// delete(stage.StackGrowthCurve2DStartHalfwayArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackgrowthcurve2dstarthalfwayarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackgrowthcurve2dstarthalfwayarcshape.GetName())
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
	for _, ref := range stage.StackGrowthCurve2DStartHalfwayArcShapes_reference {
		instance := stage.StackGrowthCurve2DStartHalfwayArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackGrowthCurve2DStartHalfwayArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackgrowthcurve2dstarthalfwayarcshapes_deletedInstances = append(stackgrowthcurve2dstarthalfwayarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackgrowthcurve2dstarthalfwayarcshapes_newInstances)
	lenDeletedInstances += len(stackgrowthcurve2dstarthalfwayarcshapes_deletedInstances)
	var stackofgrowthcurve2ds_newInstances []*StackOfGrowthCurve2D
	var stackofgrowthcurve2ds_deletedInstances []*StackOfGrowthCurve2D

	// parse all staged instances and check if they have a reference
	for stackofgrowthcurve2d := range stage.StackOfGrowthCurve2Ds {
		if ref, ok := stage.StackOfGrowthCurve2Ds_reference[stackofgrowthcurve2d]; !ok {
			stackofgrowthcurve2ds_newInstances = append(stackofgrowthcurve2ds_newInstances, stackofgrowthcurve2d)
			newInstancesSlice = append(newInstancesSlice, stackofgrowthcurve2d.GongMarshallIdentifier(stage))
			if stage.StackOfGrowthCurve2Ds_referenceOrder == nil {
				stage.StackOfGrowthCurve2Ds_referenceOrder = make(map[*StackOfGrowthCurve2D]uint)
			}
			stage.StackOfGrowthCurve2Ds_referenceOrder[stackofgrowthcurve2d] = stage.StackOfGrowthCurve2D_stagedOrder[stackofgrowthcurve2d]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackofgrowthcurve2d.GongMarshallUnstaging(stage))
			// delete(stage.StackOfGrowthCurve2Ds_referenceOrder, stackofgrowthcurve2d)
			fieldInitializers, pointersInitializations := stackofgrowthcurve2d.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackOfGrowthCurve2D_stagedOrder[ref] = stage.StackOfGrowthCurve2D_stagedOrder[stackofgrowthcurve2d]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackofgrowthcurve2d.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackofgrowthcurve2d)
			// delete(stage.StackOfGrowthCurve2D_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackofgrowthcurve2d.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackofgrowthcurve2d.GetName())
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
	for _, ref := range stage.StackOfGrowthCurve2Ds_reference {
		instance := stage.StackOfGrowthCurve2Ds_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackOfGrowthCurve2Ds[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackofgrowthcurve2ds_deletedInstances = append(stackofgrowthcurve2ds_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackofgrowthcurve2ds_newInstances)
	lenDeletedInstances += len(stackofgrowthcurve2ds_deletedInstances)
	var stackofgrowthcurve2dribbons_newInstances []*StackOfGrowthCurve2DRibbon
	var stackofgrowthcurve2dribbons_deletedInstances []*StackOfGrowthCurve2DRibbon

	// parse all staged instances and check if they have a reference
	for stackofgrowthcurve2dribbon := range stage.StackOfGrowthCurve2DRibbons {
		if ref, ok := stage.StackOfGrowthCurve2DRibbons_reference[stackofgrowthcurve2dribbon]; !ok {
			stackofgrowthcurve2dribbons_newInstances = append(stackofgrowthcurve2dribbons_newInstances, stackofgrowthcurve2dribbon)
			newInstancesSlice = append(newInstancesSlice, stackofgrowthcurve2dribbon.GongMarshallIdentifier(stage))
			if stage.StackOfGrowthCurve2DRibbons_referenceOrder == nil {
				stage.StackOfGrowthCurve2DRibbons_referenceOrder = make(map[*StackOfGrowthCurve2DRibbon]uint)
			}
			stage.StackOfGrowthCurve2DRibbons_referenceOrder[stackofgrowthcurve2dribbon] = stage.StackOfGrowthCurve2DRibbon_stagedOrder[stackofgrowthcurve2dribbon]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackofgrowthcurve2dribbon.GongMarshallUnstaging(stage))
			// delete(stage.StackOfGrowthCurve2DRibbons_referenceOrder, stackofgrowthcurve2dribbon)
			fieldInitializers, pointersInitializations := stackofgrowthcurve2dribbon.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackOfGrowthCurve2DRibbon_stagedOrder[ref] = stage.StackOfGrowthCurve2DRibbon_stagedOrder[stackofgrowthcurve2dribbon]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackofgrowthcurve2dribbon.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackofgrowthcurve2dribbon)
			// delete(stage.StackOfGrowthCurve2DRibbon_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackofgrowthcurve2dribbon.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackofgrowthcurve2dribbon.GetName())
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
	for _, ref := range stage.StackOfGrowthCurve2DRibbons_reference {
		instance := stage.StackOfGrowthCurve2DRibbons_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackOfGrowthCurve2DRibbons[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackofgrowthcurve2dribbons_deletedInstances = append(stackofgrowthcurve2dribbons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackofgrowthcurve2dribbons_newInstances)
	lenDeletedInstances += len(stackofgrowthcurve2dribbons_deletedInstances)
	var stackofrotatedgrowthcurve2ds_newInstances []*StackOfRotatedGrowthCurve2D
	var stackofrotatedgrowthcurve2ds_deletedInstances []*StackOfRotatedGrowthCurve2D

	// parse all staged instances and check if they have a reference
	for stackofrotatedgrowthcurve2d := range stage.StackOfRotatedGrowthCurve2Ds {
		if ref, ok := stage.StackOfRotatedGrowthCurve2Ds_reference[stackofrotatedgrowthcurve2d]; !ok {
			stackofrotatedgrowthcurve2ds_newInstances = append(stackofrotatedgrowthcurve2ds_newInstances, stackofrotatedgrowthcurve2d)
			newInstancesSlice = append(newInstancesSlice, stackofrotatedgrowthcurve2d.GongMarshallIdentifier(stage))
			if stage.StackOfRotatedGrowthCurve2Ds_referenceOrder == nil {
				stage.StackOfRotatedGrowthCurve2Ds_referenceOrder = make(map[*StackOfRotatedGrowthCurve2D]uint)
			}
			stage.StackOfRotatedGrowthCurve2Ds_referenceOrder[stackofrotatedgrowthcurve2d] = stage.StackOfRotatedGrowthCurve2D_stagedOrder[stackofrotatedgrowthcurve2d]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackofrotatedgrowthcurve2d.GongMarshallUnstaging(stage))
			// delete(stage.StackOfRotatedGrowthCurve2Ds_referenceOrder, stackofrotatedgrowthcurve2d)
			fieldInitializers, pointersInitializations := stackofrotatedgrowthcurve2d.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackOfRotatedGrowthCurve2D_stagedOrder[ref] = stage.StackOfRotatedGrowthCurve2D_stagedOrder[stackofrotatedgrowthcurve2d]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackofrotatedgrowthcurve2d.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackofrotatedgrowthcurve2d)
			// delete(stage.StackOfRotatedGrowthCurve2D_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackofrotatedgrowthcurve2d.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackofrotatedgrowthcurve2d.GetName())
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
	for _, ref := range stage.StackOfRotatedGrowthCurve2Ds_reference {
		instance := stage.StackOfRotatedGrowthCurve2Ds_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackOfRotatedGrowthCurve2Ds[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackofrotatedgrowthcurve2ds_deletedInstances = append(stackofrotatedgrowthcurve2ds_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackofrotatedgrowthcurve2ds_newInstances)
	lenDeletedInstances += len(stackofrotatedgrowthcurve2ds_deletedInstances)
	var stackofrotatedgrowthcurve2dribbons_newInstances []*StackOfRotatedGrowthCurve2DRibbon
	var stackofrotatedgrowthcurve2dribbons_deletedInstances []*StackOfRotatedGrowthCurve2DRibbon

	// parse all staged instances and check if they have a reference
	for stackofrotatedgrowthcurve2dribbon := range stage.StackOfRotatedGrowthCurve2DRibbons {
		if ref, ok := stage.StackOfRotatedGrowthCurve2DRibbons_reference[stackofrotatedgrowthcurve2dribbon]; !ok {
			stackofrotatedgrowthcurve2dribbons_newInstances = append(stackofrotatedgrowthcurve2dribbons_newInstances, stackofrotatedgrowthcurve2dribbon)
			newInstancesSlice = append(newInstancesSlice, stackofrotatedgrowthcurve2dribbon.GongMarshallIdentifier(stage))
			if stage.StackOfRotatedGrowthCurve2DRibbons_referenceOrder == nil {
				stage.StackOfRotatedGrowthCurve2DRibbons_referenceOrder = make(map[*StackOfRotatedGrowthCurve2DRibbon]uint)
			}
			stage.StackOfRotatedGrowthCurve2DRibbons_referenceOrder[stackofrotatedgrowthcurve2dribbon] = stage.StackOfRotatedGrowthCurve2DRibbon_stagedOrder[stackofrotatedgrowthcurve2dribbon]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackofrotatedgrowthcurve2dribbon.GongMarshallUnstaging(stage))
			// delete(stage.StackOfRotatedGrowthCurve2DRibbons_referenceOrder, stackofrotatedgrowthcurve2dribbon)
			fieldInitializers, pointersInitializations := stackofrotatedgrowthcurve2dribbon.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackOfRotatedGrowthCurve2DRibbon_stagedOrder[ref] = stage.StackOfRotatedGrowthCurve2DRibbon_stagedOrder[stackofrotatedgrowthcurve2dribbon]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackofrotatedgrowthcurve2dribbon.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackofrotatedgrowthcurve2dribbon)
			// delete(stage.StackOfRotatedGrowthCurve2DRibbon_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackofrotatedgrowthcurve2dribbon.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackofrotatedgrowthcurve2dribbon.GetName())
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
	for _, ref := range stage.StackOfRotatedGrowthCurve2DRibbons_reference {
		instance := stage.StackOfRotatedGrowthCurve2DRibbons_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackOfRotatedGrowthCurve2DRibbons[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackofrotatedgrowthcurve2dribbons_deletedInstances = append(stackofrotatedgrowthcurve2dribbons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackofrotatedgrowthcurve2dribbons_newInstances)
	lenDeletedInstances += len(stackofrotatedgrowthcurve2dribbons_deletedInstances)
	var stackrotatedgrowthcurve2dendarcshapes_newInstances []*StackRotatedGrowthCurve2DEndArcShape
	var stackrotatedgrowthcurve2dendarcshapes_deletedInstances []*StackRotatedGrowthCurve2DEndArcShape

	// parse all staged instances and check if they have a reference
	for stackrotatedgrowthcurve2dendarcshape := range stage.StackRotatedGrowthCurve2DEndArcShapes {
		if ref, ok := stage.StackRotatedGrowthCurve2DEndArcShapes_reference[stackrotatedgrowthcurve2dendarcshape]; !ok {
			stackrotatedgrowthcurve2dendarcshapes_newInstances = append(stackrotatedgrowthcurve2dendarcshapes_newInstances, stackrotatedgrowthcurve2dendarcshape)
			newInstancesSlice = append(newInstancesSlice, stackrotatedgrowthcurve2dendarcshape.GongMarshallIdentifier(stage))
			if stage.StackRotatedGrowthCurve2DEndArcShapes_referenceOrder == nil {
				stage.StackRotatedGrowthCurve2DEndArcShapes_referenceOrder = make(map[*StackRotatedGrowthCurve2DEndArcShape]uint)
			}
			stage.StackRotatedGrowthCurve2DEndArcShapes_referenceOrder[stackrotatedgrowthcurve2dendarcshape] = stage.StackRotatedGrowthCurve2DEndArcShape_stagedOrder[stackrotatedgrowthcurve2dendarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackrotatedgrowthcurve2dendarcshape.GongMarshallUnstaging(stage))
			// delete(stage.StackRotatedGrowthCurve2DEndArcShapes_referenceOrder, stackrotatedgrowthcurve2dendarcshape)
			fieldInitializers, pointersInitializations := stackrotatedgrowthcurve2dendarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackRotatedGrowthCurve2DEndArcShape_stagedOrder[ref] = stage.StackRotatedGrowthCurve2DEndArcShape_stagedOrder[stackrotatedgrowthcurve2dendarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackrotatedgrowthcurve2dendarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackrotatedgrowthcurve2dendarcshape)
			// delete(stage.StackRotatedGrowthCurve2DEndArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackrotatedgrowthcurve2dendarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackrotatedgrowthcurve2dendarcshape.GetName())
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
	for _, ref := range stage.StackRotatedGrowthCurve2DEndArcShapes_reference {
		instance := stage.StackRotatedGrowthCurve2DEndArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackRotatedGrowthCurve2DEndArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackrotatedgrowthcurve2dendarcshapes_deletedInstances = append(stackrotatedgrowthcurve2dendarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackrotatedgrowthcurve2dendarcshapes_newInstances)
	lenDeletedInstances += len(stackrotatedgrowthcurve2dendarcshapes_deletedInstances)
	var stackrotatedgrowthcurve2dribbonendshapes_newInstances []*StackRotatedGrowthCurve2DRibbonEndShape
	var stackrotatedgrowthcurve2dribbonendshapes_deletedInstances []*StackRotatedGrowthCurve2DRibbonEndShape

	// parse all staged instances and check if they have a reference
	for stackrotatedgrowthcurve2dribbonendshape := range stage.StackRotatedGrowthCurve2DRibbonEndShapes {
		if ref, ok := stage.StackRotatedGrowthCurve2DRibbonEndShapes_reference[stackrotatedgrowthcurve2dribbonendshape]; !ok {
			stackrotatedgrowthcurve2dribbonendshapes_newInstances = append(stackrotatedgrowthcurve2dribbonendshapes_newInstances, stackrotatedgrowthcurve2dribbonendshape)
			newInstancesSlice = append(newInstancesSlice, stackrotatedgrowthcurve2dribbonendshape.GongMarshallIdentifier(stage))
			if stage.StackRotatedGrowthCurve2DRibbonEndShapes_referenceOrder == nil {
				stage.StackRotatedGrowthCurve2DRibbonEndShapes_referenceOrder = make(map[*StackRotatedGrowthCurve2DRibbonEndShape]uint)
			}
			stage.StackRotatedGrowthCurve2DRibbonEndShapes_referenceOrder[stackrotatedgrowthcurve2dribbonendshape] = stage.StackRotatedGrowthCurve2DRibbonEndShape_stagedOrder[stackrotatedgrowthcurve2dribbonendshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackrotatedgrowthcurve2dribbonendshape.GongMarshallUnstaging(stage))
			// delete(stage.StackRotatedGrowthCurve2DRibbonEndShapes_referenceOrder, stackrotatedgrowthcurve2dribbonendshape)
			fieldInitializers, pointersInitializations := stackrotatedgrowthcurve2dribbonendshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackRotatedGrowthCurve2DRibbonEndShape_stagedOrder[ref] = stage.StackRotatedGrowthCurve2DRibbonEndShape_stagedOrder[stackrotatedgrowthcurve2dribbonendshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackrotatedgrowthcurve2dribbonendshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackrotatedgrowthcurve2dribbonendshape)
			// delete(stage.StackRotatedGrowthCurve2DRibbonEndShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackrotatedgrowthcurve2dribbonendshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackrotatedgrowthcurve2dribbonendshape.GetName())
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
	for _, ref := range stage.StackRotatedGrowthCurve2DRibbonEndShapes_reference {
		instance := stage.StackRotatedGrowthCurve2DRibbonEndShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackRotatedGrowthCurve2DRibbonEndShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackrotatedgrowthcurve2dribbonendshapes_deletedInstances = append(stackrotatedgrowthcurve2dribbonendshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackrotatedgrowthcurve2dribbonendshapes_newInstances)
	lenDeletedInstances += len(stackrotatedgrowthcurve2dribbonendshapes_deletedInstances)
	var stackrotatedgrowthcurve2dribbonstartshapes_newInstances []*StackRotatedGrowthCurve2DRibbonStartShape
	var stackrotatedgrowthcurve2dribbonstartshapes_deletedInstances []*StackRotatedGrowthCurve2DRibbonStartShape

	// parse all staged instances and check if they have a reference
	for stackrotatedgrowthcurve2dribbonstartshape := range stage.StackRotatedGrowthCurve2DRibbonStartShapes {
		if ref, ok := stage.StackRotatedGrowthCurve2DRibbonStartShapes_reference[stackrotatedgrowthcurve2dribbonstartshape]; !ok {
			stackrotatedgrowthcurve2dribbonstartshapes_newInstances = append(stackrotatedgrowthcurve2dribbonstartshapes_newInstances, stackrotatedgrowthcurve2dribbonstartshape)
			newInstancesSlice = append(newInstancesSlice, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallIdentifier(stage))
			if stage.StackRotatedGrowthCurve2DRibbonStartShapes_referenceOrder == nil {
				stage.StackRotatedGrowthCurve2DRibbonStartShapes_referenceOrder = make(map[*StackRotatedGrowthCurve2DRibbonStartShape]uint)
			}
			stage.StackRotatedGrowthCurve2DRibbonStartShapes_referenceOrder[stackrotatedgrowthcurve2dribbonstartshape] = stage.StackRotatedGrowthCurve2DRibbonStartShape_stagedOrder[stackrotatedgrowthcurve2dribbonstartshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackrotatedgrowthcurve2dribbonstartshape.GongMarshallUnstaging(stage))
			// delete(stage.StackRotatedGrowthCurve2DRibbonStartShapes_referenceOrder, stackrotatedgrowthcurve2dribbonstartshape)
			fieldInitializers, pointersInitializations := stackrotatedgrowthcurve2dribbonstartshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackRotatedGrowthCurve2DRibbonStartShape_stagedOrder[ref] = stage.StackRotatedGrowthCurve2DRibbonStartShape_stagedOrder[stackrotatedgrowthcurve2dribbonstartshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackrotatedgrowthcurve2dribbonstartshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackrotatedgrowthcurve2dribbonstartshape)
			// delete(stage.StackRotatedGrowthCurve2DRibbonStartShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackrotatedgrowthcurve2dribbonstartshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackrotatedgrowthcurve2dribbonstartshape.GetName())
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
	for _, ref := range stage.StackRotatedGrowthCurve2DRibbonStartShapes_reference {
		instance := stage.StackRotatedGrowthCurve2DRibbonStartShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackRotatedGrowthCurve2DRibbonStartShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackrotatedgrowthcurve2dribbonstartshapes_deletedInstances = append(stackrotatedgrowthcurve2dribbonstartshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackrotatedgrowthcurve2dribbonstartshapes_newInstances)
	lenDeletedInstances += len(stackrotatedgrowthcurve2dribbonstartshapes_deletedInstances)
	var stackrotatedgrowthcurve2dstartarcshapes_newInstances []*StackRotatedGrowthCurve2DStartArcShape
	var stackrotatedgrowthcurve2dstartarcshapes_deletedInstances []*StackRotatedGrowthCurve2DStartArcShape

	// parse all staged instances and check if they have a reference
	for stackrotatedgrowthcurve2dstartarcshape := range stage.StackRotatedGrowthCurve2DStartArcShapes {
		if ref, ok := stage.StackRotatedGrowthCurve2DStartArcShapes_reference[stackrotatedgrowthcurve2dstartarcshape]; !ok {
			stackrotatedgrowthcurve2dstartarcshapes_newInstances = append(stackrotatedgrowthcurve2dstartarcshapes_newInstances, stackrotatedgrowthcurve2dstartarcshape)
			newInstancesSlice = append(newInstancesSlice, stackrotatedgrowthcurve2dstartarcshape.GongMarshallIdentifier(stage))
			if stage.StackRotatedGrowthCurve2DStartArcShapes_referenceOrder == nil {
				stage.StackRotatedGrowthCurve2DStartArcShapes_referenceOrder = make(map[*StackRotatedGrowthCurve2DStartArcShape]uint)
			}
			stage.StackRotatedGrowthCurve2DStartArcShapes_referenceOrder[stackrotatedgrowthcurve2dstartarcshape] = stage.StackRotatedGrowthCurve2DStartArcShape_stagedOrder[stackrotatedgrowthcurve2dstartarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stackrotatedgrowthcurve2dstartarcshape.GongMarshallUnstaging(stage))
			// delete(stage.StackRotatedGrowthCurve2DStartArcShapes_referenceOrder, stackrotatedgrowthcurve2dstartarcshape)
			fieldInitializers, pointersInitializations := stackrotatedgrowthcurve2dstartarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StackRotatedGrowthCurve2DStartArcShape_stagedOrder[ref] = stage.StackRotatedGrowthCurve2DStartArcShape_stagedOrder[stackrotatedgrowthcurve2dstartarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stackrotatedgrowthcurve2dstartarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stackrotatedgrowthcurve2dstartarcshape)
			// delete(stage.StackRotatedGrowthCurve2DStartArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stackrotatedgrowthcurve2dstartarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stackrotatedgrowthcurve2dstartarcshape.GetName())
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
	for _, ref := range stage.StackRotatedGrowthCurve2DStartArcShapes_reference {
		instance := stage.StackRotatedGrowthCurve2DStartArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StackRotatedGrowthCurve2DStartArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stackrotatedgrowthcurve2dstartarcshapes_deletedInstances = append(stackrotatedgrowthcurve2dstartarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stackrotatedgrowthcurve2dstartarcshapes_newInstances)
	lenDeletedInstances += len(stackrotatedgrowthcurve2dstartarcshapes_deletedInstances)
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
	var starthalfwayarcshapes_newInstances []*StartHalfwayArcShape
	var starthalfwayarcshapes_deletedInstances []*StartHalfwayArcShape

	// parse all staged instances and check if they have a reference
	for starthalfwayarcshape := range stage.StartHalfwayArcShapes {
		if ref, ok := stage.StartHalfwayArcShapes_reference[starthalfwayarcshape]; !ok {
			starthalfwayarcshapes_newInstances = append(starthalfwayarcshapes_newInstances, starthalfwayarcshape)
			newInstancesSlice = append(newInstancesSlice, starthalfwayarcshape.GongMarshallIdentifier(stage))
			if stage.StartHalfwayArcShapes_referenceOrder == nil {
				stage.StartHalfwayArcShapes_referenceOrder = make(map[*StartHalfwayArcShape]uint)
			}
			stage.StartHalfwayArcShapes_referenceOrder[starthalfwayarcshape] = stage.StartHalfwayArcShape_stagedOrder[starthalfwayarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, starthalfwayarcshape.GongMarshallUnstaging(stage))
			// delete(stage.StartHalfwayArcShapes_referenceOrder, starthalfwayarcshape)
			fieldInitializers, pointersInitializations := starthalfwayarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StartHalfwayArcShape_stagedOrder[ref] = stage.StartHalfwayArcShape_stagedOrder[starthalfwayarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := starthalfwayarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, starthalfwayarcshape)
			// delete(stage.StartHalfwayArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if starthalfwayarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", starthalfwayarcshape.GetName())
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
	for _, ref := range stage.StartHalfwayArcShapes_reference {
		instance := stage.StartHalfwayArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StartHalfwayArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			starthalfwayarcshapes_deletedInstances = append(starthalfwayarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(starthalfwayarcshapes_newInstances)
	lenDeletedInstances += len(starthalfwayarcshapes_deletedInstances)
	var starthalfwayarcshapegrids_newInstances []*StartHalfwayArcShapeGrid
	var starthalfwayarcshapegrids_deletedInstances []*StartHalfwayArcShapeGrid

	// parse all staged instances and check if they have a reference
	for starthalfwayarcshapegrid := range stage.StartHalfwayArcShapeGrids {
		if ref, ok := stage.StartHalfwayArcShapeGrids_reference[starthalfwayarcshapegrid]; !ok {
			starthalfwayarcshapegrids_newInstances = append(starthalfwayarcshapegrids_newInstances, starthalfwayarcshapegrid)
			newInstancesSlice = append(newInstancesSlice, starthalfwayarcshapegrid.GongMarshallIdentifier(stage))
			if stage.StartHalfwayArcShapeGrids_referenceOrder == nil {
				stage.StartHalfwayArcShapeGrids_referenceOrder = make(map[*StartHalfwayArcShapeGrid]uint)
			}
			stage.StartHalfwayArcShapeGrids_referenceOrder[starthalfwayarcshapegrid] = stage.StartHalfwayArcShapeGrid_stagedOrder[starthalfwayarcshapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, starthalfwayarcshapegrid.GongMarshallUnstaging(stage))
			// delete(stage.StartHalfwayArcShapeGrids_referenceOrder, starthalfwayarcshapegrid)
			fieldInitializers, pointersInitializations := starthalfwayarcshapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StartHalfwayArcShapeGrid_stagedOrder[ref] = stage.StartHalfwayArcShapeGrid_stagedOrder[starthalfwayarcshapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := starthalfwayarcshapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, starthalfwayarcshapegrid)
			// delete(stage.StartHalfwayArcShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if starthalfwayarcshapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", starthalfwayarcshapegrid.GetName())
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
	for _, ref := range stage.StartHalfwayArcShapeGrids_reference {
		instance := stage.StartHalfwayArcShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StartHalfwayArcShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			starthalfwayarcshapegrids_deletedInstances = append(starthalfwayarcshapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(starthalfwayarcshapegrids_newInstances)
	lenDeletedInstances += len(starthalfwayarcshapegrids_deletedInstances)
	var topendarcshapes_newInstances []*TopEndArcShape
	var topendarcshapes_deletedInstances []*TopEndArcShape

	// parse all staged instances and check if they have a reference
	for topendarcshape := range stage.TopEndArcShapes {
		if ref, ok := stage.TopEndArcShapes_reference[topendarcshape]; !ok {
			topendarcshapes_newInstances = append(topendarcshapes_newInstances, topendarcshape)
			newInstancesSlice = append(newInstancesSlice, topendarcshape.GongMarshallIdentifier(stage))
			if stage.TopEndArcShapes_referenceOrder == nil {
				stage.TopEndArcShapes_referenceOrder = make(map[*TopEndArcShape]uint)
			}
			stage.TopEndArcShapes_referenceOrder[topendarcshape] = stage.TopEndArcShape_stagedOrder[topendarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topendarcshape.GongMarshallUnstaging(stage))
			// delete(stage.TopEndArcShapes_referenceOrder, topendarcshape)
			fieldInitializers, pointersInitializations := topendarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopEndArcShape_stagedOrder[ref] = stage.TopEndArcShape_stagedOrder[topendarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topendarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topendarcshape)
			// delete(stage.TopEndArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topendarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topendarcshape.GetName())
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
	for _, ref := range stage.TopEndArcShapes_reference {
		instance := stage.TopEndArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopEndArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topendarcshapes_deletedInstances = append(topendarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topendarcshapes_newInstances)
	lenDeletedInstances += len(topendarcshapes_deletedInstances)
	var topendarcshapegrids_newInstances []*TopEndArcShapeGrid
	var topendarcshapegrids_deletedInstances []*TopEndArcShapeGrid

	// parse all staged instances and check if they have a reference
	for topendarcshapegrid := range stage.TopEndArcShapeGrids {
		if ref, ok := stage.TopEndArcShapeGrids_reference[topendarcshapegrid]; !ok {
			topendarcshapegrids_newInstances = append(topendarcshapegrids_newInstances, topendarcshapegrid)
			newInstancesSlice = append(newInstancesSlice, topendarcshapegrid.GongMarshallIdentifier(stage))
			if stage.TopEndArcShapeGrids_referenceOrder == nil {
				stage.TopEndArcShapeGrids_referenceOrder = make(map[*TopEndArcShapeGrid]uint)
			}
			stage.TopEndArcShapeGrids_referenceOrder[topendarcshapegrid] = stage.TopEndArcShapeGrid_stagedOrder[topendarcshapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topendarcshapegrid.GongMarshallUnstaging(stage))
			// delete(stage.TopEndArcShapeGrids_referenceOrder, topendarcshapegrid)
			fieldInitializers, pointersInitializations := topendarcshapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopEndArcShapeGrid_stagedOrder[ref] = stage.TopEndArcShapeGrid_stagedOrder[topendarcshapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topendarcshapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topendarcshapegrid)
			// delete(stage.TopEndArcShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topendarcshapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topendarcshapegrid.GetName())
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
	for _, ref := range stage.TopEndArcShapeGrids_reference {
		instance := stage.TopEndArcShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopEndArcShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topendarcshapegrids_deletedInstances = append(topendarcshapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topendarcshapegrids_newInstances)
	lenDeletedInstances += len(topendarcshapegrids_deletedInstances)
	var topendhalfwayarcshapes_newInstances []*TopEndHalfwayArcShape
	var topendhalfwayarcshapes_deletedInstances []*TopEndHalfwayArcShape

	// parse all staged instances and check if they have a reference
	for topendhalfwayarcshape := range stage.TopEndHalfwayArcShapes {
		if ref, ok := stage.TopEndHalfwayArcShapes_reference[topendhalfwayarcshape]; !ok {
			topendhalfwayarcshapes_newInstances = append(topendhalfwayarcshapes_newInstances, topendhalfwayarcshape)
			newInstancesSlice = append(newInstancesSlice, topendhalfwayarcshape.GongMarshallIdentifier(stage))
			if stage.TopEndHalfwayArcShapes_referenceOrder == nil {
				stage.TopEndHalfwayArcShapes_referenceOrder = make(map[*TopEndHalfwayArcShape]uint)
			}
			stage.TopEndHalfwayArcShapes_referenceOrder[topendhalfwayarcshape] = stage.TopEndHalfwayArcShape_stagedOrder[topendhalfwayarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topendhalfwayarcshape.GongMarshallUnstaging(stage))
			// delete(stage.TopEndHalfwayArcShapes_referenceOrder, topendhalfwayarcshape)
			fieldInitializers, pointersInitializations := topendhalfwayarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopEndHalfwayArcShape_stagedOrder[ref] = stage.TopEndHalfwayArcShape_stagedOrder[topendhalfwayarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topendhalfwayarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topendhalfwayarcshape)
			// delete(stage.TopEndHalfwayArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topendhalfwayarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topendhalfwayarcshape.GetName())
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
	for _, ref := range stage.TopEndHalfwayArcShapes_reference {
		instance := stage.TopEndHalfwayArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopEndHalfwayArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topendhalfwayarcshapes_deletedInstances = append(topendhalfwayarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topendhalfwayarcshapes_newInstances)
	lenDeletedInstances += len(topendhalfwayarcshapes_deletedInstances)
	var topendhalfwayarcshapegrids_newInstances []*TopEndHalfwayArcShapeGrid
	var topendhalfwayarcshapegrids_deletedInstances []*TopEndHalfwayArcShapeGrid

	// parse all staged instances and check if they have a reference
	for topendhalfwayarcshapegrid := range stage.TopEndHalfwayArcShapeGrids {
		if ref, ok := stage.TopEndHalfwayArcShapeGrids_reference[topendhalfwayarcshapegrid]; !ok {
			topendhalfwayarcshapegrids_newInstances = append(topendhalfwayarcshapegrids_newInstances, topendhalfwayarcshapegrid)
			newInstancesSlice = append(newInstancesSlice, topendhalfwayarcshapegrid.GongMarshallIdentifier(stage))
			if stage.TopEndHalfwayArcShapeGrids_referenceOrder == nil {
				stage.TopEndHalfwayArcShapeGrids_referenceOrder = make(map[*TopEndHalfwayArcShapeGrid]uint)
			}
			stage.TopEndHalfwayArcShapeGrids_referenceOrder[topendhalfwayarcshapegrid] = stage.TopEndHalfwayArcShapeGrid_stagedOrder[topendhalfwayarcshapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topendhalfwayarcshapegrid.GongMarshallUnstaging(stage))
			// delete(stage.TopEndHalfwayArcShapeGrids_referenceOrder, topendhalfwayarcshapegrid)
			fieldInitializers, pointersInitializations := topendhalfwayarcshapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopEndHalfwayArcShapeGrid_stagedOrder[ref] = stage.TopEndHalfwayArcShapeGrid_stagedOrder[topendhalfwayarcshapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topendhalfwayarcshapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topendhalfwayarcshapegrid)
			// delete(stage.TopEndHalfwayArcShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topendhalfwayarcshapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topendhalfwayarcshapegrid.GetName())
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
	for _, ref := range stage.TopEndHalfwayArcShapeGrids_reference {
		instance := stage.TopEndHalfwayArcShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopEndHalfwayArcShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topendhalfwayarcshapegrids_deletedInstances = append(topendhalfwayarcshapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topendhalfwayarcshapegrids_newInstances)
	lenDeletedInstances += len(topendhalfwayarcshapegrids_deletedInstances)
	var topgrowthcurve2ds_newInstances []*TopGrowthCurve2D
	var topgrowthcurve2ds_deletedInstances []*TopGrowthCurve2D

	// parse all staged instances and check if they have a reference
	for topgrowthcurve2d := range stage.TopGrowthCurve2Ds {
		if ref, ok := stage.TopGrowthCurve2Ds_reference[topgrowthcurve2d]; !ok {
			topgrowthcurve2ds_newInstances = append(topgrowthcurve2ds_newInstances, topgrowthcurve2d)
			newInstancesSlice = append(newInstancesSlice, topgrowthcurve2d.GongMarshallIdentifier(stage))
			if stage.TopGrowthCurve2Ds_referenceOrder == nil {
				stage.TopGrowthCurve2Ds_referenceOrder = make(map[*TopGrowthCurve2D]uint)
			}
			stage.TopGrowthCurve2Ds_referenceOrder[topgrowthcurve2d] = stage.TopGrowthCurve2D_stagedOrder[topgrowthcurve2d]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topgrowthcurve2d.GongMarshallUnstaging(stage))
			// delete(stage.TopGrowthCurve2Ds_referenceOrder, topgrowthcurve2d)
			fieldInitializers, pointersInitializations := topgrowthcurve2d.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopGrowthCurve2D_stagedOrder[ref] = stage.TopGrowthCurve2D_stagedOrder[topgrowthcurve2d]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topgrowthcurve2d.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topgrowthcurve2d)
			// delete(stage.TopGrowthCurve2D_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topgrowthcurve2d.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topgrowthcurve2d.GetName())
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
	for _, ref := range stage.TopGrowthCurve2Ds_reference {
		instance := stage.TopGrowthCurve2Ds_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopGrowthCurve2Ds[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topgrowthcurve2ds_deletedInstances = append(topgrowthcurve2ds_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topgrowthcurve2ds_newInstances)
	lenDeletedInstances += len(topgrowthcurve2ds_deletedInstances)
	var topmidarcvectorshapes_newInstances []*TopMidArcVectorShape
	var topmidarcvectorshapes_deletedInstances []*TopMidArcVectorShape

	// parse all staged instances and check if they have a reference
	for topmidarcvectorshape := range stage.TopMidArcVectorShapes {
		if ref, ok := stage.TopMidArcVectorShapes_reference[topmidarcvectorshape]; !ok {
			topmidarcvectorshapes_newInstances = append(topmidarcvectorshapes_newInstances, topmidarcvectorshape)
			newInstancesSlice = append(newInstancesSlice, topmidarcvectorshape.GongMarshallIdentifier(stage))
			if stage.TopMidArcVectorShapes_referenceOrder == nil {
				stage.TopMidArcVectorShapes_referenceOrder = make(map[*TopMidArcVectorShape]uint)
			}
			stage.TopMidArcVectorShapes_referenceOrder[topmidarcvectorshape] = stage.TopMidArcVectorShape_stagedOrder[topmidarcvectorshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topmidarcvectorshape.GongMarshallUnstaging(stage))
			// delete(stage.TopMidArcVectorShapes_referenceOrder, topmidarcvectorshape)
			fieldInitializers, pointersInitializations := topmidarcvectorshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopMidArcVectorShape_stagedOrder[ref] = stage.TopMidArcVectorShape_stagedOrder[topmidarcvectorshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topmidarcvectorshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topmidarcvectorshape)
			// delete(stage.TopMidArcVectorShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topmidarcvectorshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topmidarcvectorshape.GetName())
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
	for _, ref := range stage.TopMidArcVectorShapes_reference {
		instance := stage.TopMidArcVectorShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopMidArcVectorShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topmidarcvectorshapes_deletedInstances = append(topmidarcvectorshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topmidarcvectorshapes_newInstances)
	lenDeletedInstances += len(topmidarcvectorshapes_deletedInstances)
	var topmidarcvectorshapegrids_newInstances []*TopMidArcVectorShapeGrid
	var topmidarcvectorshapegrids_deletedInstances []*TopMidArcVectorShapeGrid

	// parse all staged instances and check if they have a reference
	for topmidarcvectorshapegrid := range stage.TopMidArcVectorShapeGrids {
		if ref, ok := stage.TopMidArcVectorShapeGrids_reference[topmidarcvectorshapegrid]; !ok {
			topmidarcvectorshapegrids_newInstances = append(topmidarcvectorshapegrids_newInstances, topmidarcvectorshapegrid)
			newInstancesSlice = append(newInstancesSlice, topmidarcvectorshapegrid.GongMarshallIdentifier(stage))
			if stage.TopMidArcVectorShapeGrids_referenceOrder == nil {
				stage.TopMidArcVectorShapeGrids_referenceOrder = make(map[*TopMidArcVectorShapeGrid]uint)
			}
			stage.TopMidArcVectorShapeGrids_referenceOrder[topmidarcvectorshapegrid] = stage.TopMidArcVectorShapeGrid_stagedOrder[topmidarcvectorshapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topmidarcvectorshapegrid.GongMarshallUnstaging(stage))
			// delete(stage.TopMidArcVectorShapeGrids_referenceOrder, topmidarcvectorshapegrid)
			fieldInitializers, pointersInitializations := topmidarcvectorshapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopMidArcVectorShapeGrid_stagedOrder[ref] = stage.TopMidArcVectorShapeGrid_stagedOrder[topmidarcvectorshapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topmidarcvectorshapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topmidarcvectorshapegrid)
			// delete(stage.TopMidArcVectorShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topmidarcvectorshapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topmidarcvectorshapegrid.GetName())
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
	for _, ref := range stage.TopMidArcVectorShapeGrids_reference {
		instance := stage.TopMidArcVectorShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopMidArcVectorShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topmidarcvectorshapegrids_deletedInstances = append(topmidarcvectorshapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topmidarcvectorshapegrids_newInstances)
	lenDeletedInstances += len(topmidarcvectorshapegrids_deletedInstances)
	var topstackgrowthcurve2dendhalfwayarcshapes_newInstances []*TopStackGrowthCurve2DEndHalfwayArcShape
	var topstackgrowthcurve2dendhalfwayarcshapes_deletedInstances []*TopStackGrowthCurve2DEndHalfwayArcShape

	// parse all staged instances and check if they have a reference
	for topstackgrowthcurve2dendhalfwayarcshape := range stage.TopStackGrowthCurve2DEndHalfwayArcShapes {
		if ref, ok := stage.TopStackGrowthCurve2DEndHalfwayArcShapes_reference[topstackgrowthcurve2dendhalfwayarcshape]; !ok {
			topstackgrowthcurve2dendhalfwayarcshapes_newInstances = append(topstackgrowthcurve2dendhalfwayarcshapes_newInstances, topstackgrowthcurve2dendhalfwayarcshape)
			newInstancesSlice = append(newInstancesSlice, topstackgrowthcurve2dendhalfwayarcshape.GongMarshallIdentifier(stage))
			if stage.TopStackGrowthCurve2DEndHalfwayArcShapes_referenceOrder == nil {
				stage.TopStackGrowthCurve2DEndHalfwayArcShapes_referenceOrder = make(map[*TopStackGrowthCurve2DEndHalfwayArcShape]uint)
			}
			stage.TopStackGrowthCurve2DEndHalfwayArcShapes_referenceOrder[topstackgrowthcurve2dendhalfwayarcshape] = stage.TopStackGrowthCurve2DEndHalfwayArcShape_stagedOrder[topstackgrowthcurve2dendhalfwayarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstackgrowthcurve2dendhalfwayarcshape.GongMarshallUnstaging(stage))
			// delete(stage.TopStackGrowthCurve2DEndHalfwayArcShapes_referenceOrder, topstackgrowthcurve2dendhalfwayarcshape)
			fieldInitializers, pointersInitializations := topstackgrowthcurve2dendhalfwayarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStackGrowthCurve2DEndHalfwayArcShape_stagedOrder[ref] = stage.TopStackGrowthCurve2DEndHalfwayArcShape_stagedOrder[topstackgrowthcurve2dendhalfwayarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstackgrowthcurve2dendhalfwayarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstackgrowthcurve2dendhalfwayarcshape)
			// delete(stage.TopStackGrowthCurve2DEndHalfwayArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstackgrowthcurve2dendhalfwayarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstackgrowthcurve2dendhalfwayarcshape.GetName())
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
	for _, ref := range stage.TopStackGrowthCurve2DEndHalfwayArcShapes_reference {
		instance := stage.TopStackGrowthCurve2DEndHalfwayArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStackGrowthCurve2DEndHalfwayArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstackgrowthcurve2dendhalfwayarcshapes_deletedInstances = append(topstackgrowthcurve2dendhalfwayarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstackgrowthcurve2dendhalfwayarcshapes_newInstances)
	lenDeletedInstances += len(topstackgrowthcurve2dendhalfwayarcshapes_deletedInstances)
	var topstackgrowthcurve2dstarthalfwayarcshapes_newInstances []*TopStackGrowthCurve2DStartHalfwayArcShape
	var topstackgrowthcurve2dstarthalfwayarcshapes_deletedInstances []*TopStackGrowthCurve2DStartHalfwayArcShape

	// parse all staged instances and check if they have a reference
	for topstackgrowthcurve2dstarthalfwayarcshape := range stage.TopStackGrowthCurve2DStartHalfwayArcShapes {
		if ref, ok := stage.TopStackGrowthCurve2DStartHalfwayArcShapes_reference[topstackgrowthcurve2dstarthalfwayarcshape]; !ok {
			topstackgrowthcurve2dstarthalfwayarcshapes_newInstances = append(topstackgrowthcurve2dstarthalfwayarcshapes_newInstances, topstackgrowthcurve2dstarthalfwayarcshape)
			newInstancesSlice = append(newInstancesSlice, topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallIdentifier(stage))
			if stage.TopStackGrowthCurve2DStartHalfwayArcShapes_referenceOrder == nil {
				stage.TopStackGrowthCurve2DStartHalfwayArcShapes_referenceOrder = make(map[*TopStackGrowthCurve2DStartHalfwayArcShape]uint)
			}
			stage.TopStackGrowthCurve2DStartHalfwayArcShapes_referenceOrder[topstackgrowthcurve2dstarthalfwayarcshape] = stage.TopStackGrowthCurve2DStartHalfwayArcShape_stagedOrder[topstackgrowthcurve2dstarthalfwayarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallUnstaging(stage))
			// delete(stage.TopStackGrowthCurve2DStartHalfwayArcShapes_referenceOrder, topstackgrowthcurve2dstarthalfwayarcshape)
			fieldInitializers, pointersInitializations := topstackgrowthcurve2dstarthalfwayarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStackGrowthCurve2DStartHalfwayArcShape_stagedOrder[ref] = stage.TopStackGrowthCurve2DStartHalfwayArcShape_stagedOrder[topstackgrowthcurve2dstarthalfwayarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstackgrowthcurve2dstarthalfwayarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstackgrowthcurve2dstarthalfwayarcshape)
			// delete(stage.TopStackGrowthCurve2DStartHalfwayArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstackgrowthcurve2dstarthalfwayarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstackgrowthcurve2dstarthalfwayarcshape.GetName())
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
	for _, ref := range stage.TopStackGrowthCurve2DStartHalfwayArcShapes_reference {
		instance := stage.TopStackGrowthCurve2DStartHalfwayArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStackGrowthCurve2DStartHalfwayArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstackgrowthcurve2dstarthalfwayarcshapes_deletedInstances = append(topstackgrowthcurve2dstarthalfwayarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstackgrowthcurve2dstarthalfwayarcshapes_newInstances)
	lenDeletedInstances += len(topstackgrowthcurve2dstarthalfwayarcshapes_deletedInstances)
	var topstackofgrowthcurve2ds_newInstances []*TopStackOfGrowthCurve2D
	var topstackofgrowthcurve2ds_deletedInstances []*TopStackOfGrowthCurve2D

	// parse all staged instances and check if they have a reference
	for topstackofgrowthcurve2d := range stage.TopStackOfGrowthCurve2Ds {
		if ref, ok := stage.TopStackOfGrowthCurve2Ds_reference[topstackofgrowthcurve2d]; !ok {
			topstackofgrowthcurve2ds_newInstances = append(topstackofgrowthcurve2ds_newInstances, topstackofgrowthcurve2d)
			newInstancesSlice = append(newInstancesSlice, topstackofgrowthcurve2d.GongMarshallIdentifier(stage))
			if stage.TopStackOfGrowthCurve2Ds_referenceOrder == nil {
				stage.TopStackOfGrowthCurve2Ds_referenceOrder = make(map[*TopStackOfGrowthCurve2D]uint)
			}
			stage.TopStackOfGrowthCurve2Ds_referenceOrder[topstackofgrowthcurve2d] = stage.TopStackOfGrowthCurve2D_stagedOrder[topstackofgrowthcurve2d]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstackofgrowthcurve2d.GongMarshallUnstaging(stage))
			// delete(stage.TopStackOfGrowthCurve2Ds_referenceOrder, topstackofgrowthcurve2d)
			fieldInitializers, pointersInitializations := topstackofgrowthcurve2d.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStackOfGrowthCurve2D_stagedOrder[ref] = stage.TopStackOfGrowthCurve2D_stagedOrder[topstackofgrowthcurve2d]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstackofgrowthcurve2d.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstackofgrowthcurve2d)
			// delete(stage.TopStackOfGrowthCurve2D_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstackofgrowthcurve2d.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstackofgrowthcurve2d.GetName())
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
	for _, ref := range stage.TopStackOfGrowthCurve2Ds_reference {
		instance := stage.TopStackOfGrowthCurve2Ds_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStackOfGrowthCurve2Ds[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstackofgrowthcurve2ds_deletedInstances = append(topstackofgrowthcurve2ds_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstackofgrowthcurve2ds_newInstances)
	lenDeletedInstances += len(topstackofgrowthcurve2ds_deletedInstances)
	var topstackofrotatedgrowthcurve2ds_newInstances []*TopStackOfRotatedGrowthCurve2D
	var topstackofrotatedgrowthcurve2ds_deletedInstances []*TopStackOfRotatedGrowthCurve2D

	// parse all staged instances and check if they have a reference
	for topstackofrotatedgrowthcurve2d := range stage.TopStackOfRotatedGrowthCurve2Ds {
		if ref, ok := stage.TopStackOfRotatedGrowthCurve2Ds_reference[topstackofrotatedgrowthcurve2d]; !ok {
			topstackofrotatedgrowthcurve2ds_newInstances = append(topstackofrotatedgrowthcurve2ds_newInstances, topstackofrotatedgrowthcurve2d)
			newInstancesSlice = append(newInstancesSlice, topstackofrotatedgrowthcurve2d.GongMarshallIdentifier(stage))
			if stage.TopStackOfRotatedGrowthCurve2Ds_referenceOrder == nil {
				stage.TopStackOfRotatedGrowthCurve2Ds_referenceOrder = make(map[*TopStackOfRotatedGrowthCurve2D]uint)
			}
			stage.TopStackOfRotatedGrowthCurve2Ds_referenceOrder[topstackofrotatedgrowthcurve2d] = stage.TopStackOfRotatedGrowthCurve2D_stagedOrder[topstackofrotatedgrowthcurve2d]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstackofrotatedgrowthcurve2d.GongMarshallUnstaging(stage))
			// delete(stage.TopStackOfRotatedGrowthCurve2Ds_referenceOrder, topstackofrotatedgrowthcurve2d)
			fieldInitializers, pointersInitializations := topstackofrotatedgrowthcurve2d.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStackOfRotatedGrowthCurve2D_stagedOrder[ref] = stage.TopStackOfRotatedGrowthCurve2D_stagedOrder[topstackofrotatedgrowthcurve2d]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstackofrotatedgrowthcurve2d.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstackofrotatedgrowthcurve2d)
			// delete(stage.TopStackOfRotatedGrowthCurve2D_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstackofrotatedgrowthcurve2d.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstackofrotatedgrowthcurve2d.GetName())
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
	for _, ref := range stage.TopStackOfRotatedGrowthCurve2Ds_reference {
		instance := stage.TopStackOfRotatedGrowthCurve2Ds_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStackOfRotatedGrowthCurve2Ds[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstackofrotatedgrowthcurve2ds_deletedInstances = append(topstackofrotatedgrowthcurve2ds_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstackofrotatedgrowthcurve2ds_newInstances)
	lenDeletedInstances += len(topstackofrotatedgrowthcurve2ds_deletedInstances)
	var topstackofrotatedgrowthcurve2dendarcshapes_newInstances []*TopStackOfRotatedGrowthCurve2DEndArcShape
	var topstackofrotatedgrowthcurve2dendarcshapes_deletedInstances []*TopStackOfRotatedGrowthCurve2DEndArcShape

	// parse all staged instances and check if they have a reference
	for topstackofrotatedgrowthcurve2dendarcshape := range stage.TopStackOfRotatedGrowthCurve2DEndArcShapes {
		if ref, ok := stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_reference[topstackofrotatedgrowthcurve2dendarcshape]; !ok {
			topstackofrotatedgrowthcurve2dendarcshapes_newInstances = append(topstackofrotatedgrowthcurve2dendarcshapes_newInstances, topstackofrotatedgrowthcurve2dendarcshape)
			newInstancesSlice = append(newInstancesSlice, topstackofrotatedgrowthcurve2dendarcshape.GongMarshallIdentifier(stage))
			if stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_referenceOrder == nil {
				stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_referenceOrder = make(map[*TopStackOfRotatedGrowthCurve2DEndArcShape]uint)
			}
			stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_referenceOrder[topstackofrotatedgrowthcurve2dendarcshape] = stage.TopStackOfRotatedGrowthCurve2DEndArcShape_stagedOrder[topstackofrotatedgrowthcurve2dendarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstackofrotatedgrowthcurve2dendarcshape.GongMarshallUnstaging(stage))
			// delete(stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_referenceOrder, topstackofrotatedgrowthcurve2dendarcshape)
			fieldInitializers, pointersInitializations := topstackofrotatedgrowthcurve2dendarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStackOfRotatedGrowthCurve2DEndArcShape_stagedOrder[ref] = stage.TopStackOfRotatedGrowthCurve2DEndArcShape_stagedOrder[topstackofrotatedgrowthcurve2dendarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstackofrotatedgrowthcurve2dendarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstackofrotatedgrowthcurve2dendarcshape)
			// delete(stage.TopStackOfRotatedGrowthCurve2DEndArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstackofrotatedgrowthcurve2dendarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstackofrotatedgrowthcurve2dendarcshape.GetName())
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
	for _, ref := range stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_reference {
		instance := stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStackOfRotatedGrowthCurve2DEndArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstackofrotatedgrowthcurve2dendarcshapes_deletedInstances = append(topstackofrotatedgrowthcurve2dendarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstackofrotatedgrowthcurve2dendarcshapes_newInstances)
	lenDeletedInstances += len(topstackofrotatedgrowthcurve2dendarcshapes_deletedInstances)
	var topstackofrotatedgrowthcurve2dstartarcshapes_newInstances []*TopStackOfRotatedGrowthCurve2DStartArcShape
	var topstackofrotatedgrowthcurve2dstartarcshapes_deletedInstances []*TopStackOfRotatedGrowthCurve2DStartArcShape

	// parse all staged instances and check if they have a reference
	for topstackofrotatedgrowthcurve2dstartarcshape := range stage.TopStackOfRotatedGrowthCurve2DStartArcShapes {
		if ref, ok := stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_reference[topstackofrotatedgrowthcurve2dstartarcshape]; !ok {
			topstackofrotatedgrowthcurve2dstartarcshapes_newInstances = append(topstackofrotatedgrowthcurve2dstartarcshapes_newInstances, topstackofrotatedgrowthcurve2dstartarcshape)
			newInstancesSlice = append(newInstancesSlice, topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallIdentifier(stage))
			if stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_referenceOrder == nil {
				stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_referenceOrder = make(map[*TopStackOfRotatedGrowthCurve2DStartArcShape]uint)
			}
			stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_referenceOrder[topstackofrotatedgrowthcurve2dstartarcshape] = stage.TopStackOfRotatedGrowthCurve2DStartArcShape_stagedOrder[topstackofrotatedgrowthcurve2dstartarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallUnstaging(stage))
			// delete(stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_referenceOrder, topstackofrotatedgrowthcurve2dstartarcshape)
			fieldInitializers, pointersInitializations := topstackofrotatedgrowthcurve2dstartarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStackOfRotatedGrowthCurve2DStartArcShape_stagedOrder[ref] = stage.TopStackOfRotatedGrowthCurve2DStartArcShape_stagedOrder[topstackofrotatedgrowthcurve2dstartarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstackofrotatedgrowthcurve2dstartarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstackofrotatedgrowthcurve2dstartarcshape)
			// delete(stage.TopStackOfRotatedGrowthCurve2DStartArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstackofrotatedgrowthcurve2dstartarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstackofrotatedgrowthcurve2dstartarcshape.GetName())
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
	for _, ref := range stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_reference {
		instance := stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStackOfRotatedGrowthCurve2DStartArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstackofrotatedgrowthcurve2dstartarcshapes_deletedInstances = append(topstackofrotatedgrowthcurve2dstartarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstackofrotatedgrowthcurve2dstartarcshapes_newInstances)
	lenDeletedInstances += len(topstackofrotatedgrowthcurve2dstartarcshapes_deletedInstances)
	var topstartarcshapes_newInstances []*TopStartArcShape
	var topstartarcshapes_deletedInstances []*TopStartArcShape

	// parse all staged instances and check if they have a reference
	for topstartarcshape := range stage.TopStartArcShapes {
		if ref, ok := stage.TopStartArcShapes_reference[topstartarcshape]; !ok {
			topstartarcshapes_newInstances = append(topstartarcshapes_newInstances, topstartarcshape)
			newInstancesSlice = append(newInstancesSlice, topstartarcshape.GongMarshallIdentifier(stage))
			if stage.TopStartArcShapes_referenceOrder == nil {
				stage.TopStartArcShapes_referenceOrder = make(map[*TopStartArcShape]uint)
			}
			stage.TopStartArcShapes_referenceOrder[topstartarcshape] = stage.TopStartArcShape_stagedOrder[topstartarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstartarcshape.GongMarshallUnstaging(stage))
			// delete(stage.TopStartArcShapes_referenceOrder, topstartarcshape)
			fieldInitializers, pointersInitializations := topstartarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStartArcShape_stagedOrder[ref] = stage.TopStartArcShape_stagedOrder[topstartarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstartarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstartarcshape)
			// delete(stage.TopStartArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstartarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstartarcshape.GetName())
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
	for _, ref := range stage.TopStartArcShapes_reference {
		instance := stage.TopStartArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStartArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstartarcshapes_deletedInstances = append(topstartarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstartarcshapes_newInstances)
	lenDeletedInstances += len(topstartarcshapes_deletedInstances)
	var topstartarcshapegrids_newInstances []*TopStartArcShapeGrid
	var topstartarcshapegrids_deletedInstances []*TopStartArcShapeGrid

	// parse all staged instances and check if they have a reference
	for topstartarcshapegrid := range stage.TopStartArcShapeGrids {
		if ref, ok := stage.TopStartArcShapeGrids_reference[topstartarcshapegrid]; !ok {
			topstartarcshapegrids_newInstances = append(topstartarcshapegrids_newInstances, topstartarcshapegrid)
			newInstancesSlice = append(newInstancesSlice, topstartarcshapegrid.GongMarshallIdentifier(stage))
			if stage.TopStartArcShapeGrids_referenceOrder == nil {
				stage.TopStartArcShapeGrids_referenceOrder = make(map[*TopStartArcShapeGrid]uint)
			}
			stage.TopStartArcShapeGrids_referenceOrder[topstartarcshapegrid] = stage.TopStartArcShapeGrid_stagedOrder[topstartarcshapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstartarcshapegrid.GongMarshallUnstaging(stage))
			// delete(stage.TopStartArcShapeGrids_referenceOrder, topstartarcshapegrid)
			fieldInitializers, pointersInitializations := topstartarcshapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStartArcShapeGrid_stagedOrder[ref] = stage.TopStartArcShapeGrid_stagedOrder[topstartarcshapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstartarcshapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstartarcshapegrid)
			// delete(stage.TopStartArcShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstartarcshapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstartarcshapegrid.GetName())
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
	for _, ref := range stage.TopStartArcShapeGrids_reference {
		instance := stage.TopStartArcShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStartArcShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstartarcshapegrids_deletedInstances = append(topstartarcshapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstartarcshapegrids_newInstances)
	lenDeletedInstances += len(topstartarcshapegrids_deletedInstances)
	var topstarthalfwayarcshapes_newInstances []*TopStartHalfwayArcShape
	var topstarthalfwayarcshapes_deletedInstances []*TopStartHalfwayArcShape

	// parse all staged instances and check if they have a reference
	for topstarthalfwayarcshape := range stage.TopStartHalfwayArcShapes {
		if ref, ok := stage.TopStartHalfwayArcShapes_reference[topstarthalfwayarcshape]; !ok {
			topstarthalfwayarcshapes_newInstances = append(topstarthalfwayarcshapes_newInstances, topstarthalfwayarcshape)
			newInstancesSlice = append(newInstancesSlice, topstarthalfwayarcshape.GongMarshallIdentifier(stage))
			if stage.TopStartHalfwayArcShapes_referenceOrder == nil {
				stage.TopStartHalfwayArcShapes_referenceOrder = make(map[*TopStartHalfwayArcShape]uint)
			}
			stage.TopStartHalfwayArcShapes_referenceOrder[topstarthalfwayarcshape] = stage.TopStartHalfwayArcShape_stagedOrder[topstarthalfwayarcshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstarthalfwayarcshape.GongMarshallUnstaging(stage))
			// delete(stage.TopStartHalfwayArcShapes_referenceOrder, topstarthalfwayarcshape)
			fieldInitializers, pointersInitializations := topstarthalfwayarcshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStartHalfwayArcShape_stagedOrder[ref] = stage.TopStartHalfwayArcShape_stagedOrder[topstarthalfwayarcshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstarthalfwayarcshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstarthalfwayarcshape)
			// delete(stage.TopStartHalfwayArcShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstarthalfwayarcshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstarthalfwayarcshape.GetName())
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
	for _, ref := range stage.TopStartHalfwayArcShapes_reference {
		instance := stage.TopStartHalfwayArcShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStartHalfwayArcShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstarthalfwayarcshapes_deletedInstances = append(topstarthalfwayarcshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstarthalfwayarcshapes_newInstances)
	lenDeletedInstances += len(topstarthalfwayarcshapes_deletedInstances)
	var topstarthalfwayarcshapegrids_newInstances []*TopStartHalfwayArcShapeGrid
	var topstarthalfwayarcshapegrids_deletedInstances []*TopStartHalfwayArcShapeGrid

	// parse all staged instances and check if they have a reference
	for topstarthalfwayarcshapegrid := range stage.TopStartHalfwayArcShapeGrids {
		if ref, ok := stage.TopStartHalfwayArcShapeGrids_reference[topstarthalfwayarcshapegrid]; !ok {
			topstarthalfwayarcshapegrids_newInstances = append(topstarthalfwayarcshapegrids_newInstances, topstarthalfwayarcshapegrid)
			newInstancesSlice = append(newInstancesSlice, topstarthalfwayarcshapegrid.GongMarshallIdentifier(stage))
			if stage.TopStartHalfwayArcShapeGrids_referenceOrder == nil {
				stage.TopStartHalfwayArcShapeGrids_referenceOrder = make(map[*TopStartHalfwayArcShapeGrid]uint)
			}
			stage.TopStartHalfwayArcShapeGrids_referenceOrder[topstarthalfwayarcshapegrid] = stage.TopStartHalfwayArcShapeGrid_stagedOrder[topstarthalfwayarcshapegrid]
			newInstancesReverseSlice = append(newInstancesReverseSlice, topstarthalfwayarcshapegrid.GongMarshallUnstaging(stage))
			// delete(stage.TopStartHalfwayArcShapeGrids_referenceOrder, topstarthalfwayarcshapegrid)
			fieldInitializers, pointersInitializations := topstarthalfwayarcshapegrid.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TopStartHalfwayArcShapeGrid_stagedOrder[ref] = stage.TopStartHalfwayArcShapeGrid_stagedOrder[topstarthalfwayarcshapegrid]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := topstarthalfwayarcshapegrid.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, topstarthalfwayarcshapegrid)
			// delete(stage.TopStartHalfwayArcShapeGrid_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if topstarthalfwayarcshapegrid.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", topstarthalfwayarcshapegrid.GetName())
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
	for _, ref := range stage.TopStartHalfwayArcShapeGrids_reference {
		instance := stage.TopStartHalfwayArcShapeGrids_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TopStartHalfwayArcShapeGrids[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			topstarthalfwayarcshapegrids_deletedInstances = append(topstarthalfwayarcshapegrids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(topstarthalfwayarcshapegrids_newInstances)
	lenDeletedInstances += len(topstarthalfwayarcshapegrids_deletedInstances)
	var torusstackshapes_newInstances []*TorusStackShape
	var torusstackshapes_deletedInstances []*TorusStackShape

	// parse all staged instances and check if they have a reference
	for torusstackshape := range stage.TorusStackShapes {
		if ref, ok := stage.TorusStackShapes_reference[torusstackshape]; !ok {
			torusstackshapes_newInstances = append(torusstackshapes_newInstances, torusstackshape)
			newInstancesSlice = append(newInstancesSlice, torusstackshape.GongMarshallIdentifier(stage))
			if stage.TorusStackShapes_referenceOrder == nil {
				stage.TorusStackShapes_referenceOrder = make(map[*TorusStackShape]uint)
			}
			stage.TorusStackShapes_referenceOrder[torusstackshape] = stage.TorusStackShape_stagedOrder[torusstackshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, torusstackshape.GongMarshallUnstaging(stage))
			// delete(stage.TorusStackShapes_referenceOrder, torusstackshape)
			fieldInitializers, pointersInitializations := torusstackshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TorusStackShape_stagedOrder[ref] = stage.TorusStackShape_stagedOrder[torusstackshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := torusstackshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, torusstackshape)
			// delete(stage.TorusStackShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if torusstackshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", torusstackshape.GetName())
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
	for _, ref := range stage.TorusStackShapes_reference {
		instance := stage.TorusStackShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TorusStackShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			torusstackshapes_deletedInstances = append(torusstackshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(torusstackshapes_newInstances)
	lenDeletedInstances += len(torusstackshapes_deletedInstances)
	var verticaltorusstackshapes_newInstances []*VerticalTorusStackShape
	var verticaltorusstackshapes_deletedInstances []*VerticalTorusStackShape

	// parse all staged instances and check if they have a reference
	for verticaltorusstackshape := range stage.VerticalTorusStackShapes {
		if ref, ok := stage.VerticalTorusStackShapes_reference[verticaltorusstackshape]; !ok {
			verticaltorusstackshapes_newInstances = append(verticaltorusstackshapes_newInstances, verticaltorusstackshape)
			newInstancesSlice = append(newInstancesSlice, verticaltorusstackshape.GongMarshallIdentifier(stage))
			if stage.VerticalTorusStackShapes_referenceOrder == nil {
				stage.VerticalTorusStackShapes_referenceOrder = make(map[*VerticalTorusStackShape]uint)
			}
			stage.VerticalTorusStackShapes_referenceOrder[verticaltorusstackshape] = stage.VerticalTorusStackShape_stagedOrder[verticaltorusstackshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, verticaltorusstackshape.GongMarshallUnstaging(stage))
			// delete(stage.VerticalTorusStackShapes_referenceOrder, verticaltorusstackshape)
			fieldInitializers, pointersInitializations := verticaltorusstackshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.VerticalTorusStackShape_stagedOrder[ref] = stage.VerticalTorusStackShape_stagedOrder[verticaltorusstackshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := verticaltorusstackshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, verticaltorusstackshape)
			// delete(stage.VerticalTorusStackShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if verticaltorusstackshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", verticaltorusstackshape.GetName())
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
	for _, ref := range stage.VerticalTorusStackShapes_reference {
		instance := stage.VerticalTorusStackShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.VerticalTorusStackShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			verticaltorusstackshapes_deletedInstances = append(verticaltorusstackshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(verticaltorusstackshapes_newInstances)
	lenDeletedInstances += len(verticaltorusstackshapes_deletedInstances)

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

	stage.EndHalfwayArcShapes_reference = make(map[*EndHalfwayArcShape]*EndHalfwayArcShape)
	stage.EndHalfwayArcShapes_referenceOrder = make(map[*EndHalfwayArcShape]uint) // diff Unstage needs the reference order
	stage.EndHalfwayArcShapes_instance = make(map[*EndHalfwayArcShape]*EndHalfwayArcShape)
	for instance := range stage.EndHalfwayArcShapes {
		_copy := instance.GongCopy().(*EndHalfwayArcShape)
		stage.EndHalfwayArcShapes_reference[instance] = _copy
		stage.EndHalfwayArcShapes_instance[_copy] = instance
		stage.EndHalfwayArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EndHalfwayArcShapeGrids_reference = make(map[*EndHalfwayArcShapeGrid]*EndHalfwayArcShapeGrid)
	stage.EndHalfwayArcShapeGrids_referenceOrder = make(map[*EndHalfwayArcShapeGrid]uint) // diff Unstage needs the reference order
	stage.EndHalfwayArcShapeGrids_instance = make(map[*EndHalfwayArcShapeGrid]*EndHalfwayArcShapeGrid)
	for instance := range stage.EndHalfwayArcShapeGrids {
		_copy := instance.GongCopy().(*EndHalfwayArcShapeGrid)
		stage.EndHalfwayArcShapeGrids_reference[instance] = _copy
		stage.EndHalfwayArcShapeGrids_instance[_copy] = instance
		stage.EndHalfwayArcShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.GrowthCurve2Ds_reference = make(map[*GrowthCurve2D]*GrowthCurve2D)
	stage.GrowthCurve2Ds_referenceOrder = make(map[*GrowthCurve2D]uint) // diff Unstage needs the reference order
	stage.GrowthCurve2Ds_instance = make(map[*GrowthCurve2D]*GrowthCurve2D)
	for instance := range stage.GrowthCurve2Ds {
		_copy := instance.GongCopy().(*GrowthCurve2D)
		stage.GrowthCurve2Ds_reference[instance] = _copy
		stage.GrowthCurve2Ds_instance[_copy] = instance
		stage.GrowthCurve2Ds_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.MidArcVectorShapes_reference = make(map[*MidArcVectorShape]*MidArcVectorShape)
	stage.MidArcVectorShapes_referenceOrder = make(map[*MidArcVectorShape]uint) // diff Unstage needs the reference order
	stage.MidArcVectorShapes_instance = make(map[*MidArcVectorShape]*MidArcVectorShape)
	for instance := range stage.MidArcVectorShapes {
		_copy := instance.GongCopy().(*MidArcVectorShape)
		stage.MidArcVectorShapes_reference[instance] = _copy
		stage.MidArcVectorShapes_instance[_copy] = instance
		stage.MidArcVectorShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.MidArcVectorShapeGrids_reference = make(map[*MidArcVectorShapeGrid]*MidArcVectorShapeGrid)
	stage.MidArcVectorShapeGrids_referenceOrder = make(map[*MidArcVectorShapeGrid]uint) // diff Unstage needs the reference order
	stage.MidArcVectorShapeGrids_instance = make(map[*MidArcVectorShapeGrid]*MidArcVectorShapeGrid)
	for instance := range stage.MidArcVectorShapeGrids {
		_copy := instance.GongCopy().(*MidArcVectorShapeGrid)
		stage.MidArcVectorShapeGrids_reference[instance] = _copy
		stage.MidArcVectorShapeGrids_instance[_copy] = instance
		stage.MidArcVectorShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartiallyGrowthCurve2DRibbons_reference = make(map[*PartiallyGrowthCurve2DRibbon]*PartiallyGrowthCurve2DRibbon)
	stage.PartiallyGrowthCurve2DRibbons_referenceOrder = make(map[*PartiallyGrowthCurve2DRibbon]uint) // diff Unstage needs the reference order
	stage.PartiallyGrowthCurve2DRibbons_instance = make(map[*PartiallyGrowthCurve2DRibbon]*PartiallyGrowthCurve2DRibbon)
	for instance := range stage.PartiallyGrowthCurve2DRibbons {
		_copy := instance.GongCopy().(*PartiallyGrowthCurve2DRibbon)
		stage.PartiallyGrowthCurve2DRibbons_reference[instance] = _copy
		stage.PartiallyGrowthCurve2DRibbons_instance[_copy] = instance
		stage.PartiallyGrowthCurve2DRibbons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartiallyGrowthCurve2DRibbonEndShapes_reference = make(map[*PartiallyGrowthCurve2DRibbonEndShape]*PartiallyGrowthCurve2DRibbonEndShape)
	stage.PartiallyGrowthCurve2DRibbonEndShapes_referenceOrder = make(map[*PartiallyGrowthCurve2DRibbonEndShape]uint) // diff Unstage needs the reference order
	stage.PartiallyGrowthCurve2DRibbonEndShapes_instance = make(map[*PartiallyGrowthCurve2DRibbonEndShape]*PartiallyGrowthCurve2DRibbonEndShape)
	for instance := range stage.PartiallyGrowthCurve2DRibbonEndShapes {
		_copy := instance.GongCopy().(*PartiallyGrowthCurve2DRibbonEndShape)
		stage.PartiallyGrowthCurve2DRibbonEndShapes_reference[instance] = _copy
		stage.PartiallyGrowthCurve2DRibbonEndShapes_instance[_copy] = instance
		stage.PartiallyGrowthCurve2DRibbonEndShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartiallyGrowthCurve2DRibbonStartShapes_reference = make(map[*PartiallyGrowthCurve2DRibbonStartShape]*PartiallyGrowthCurve2DRibbonStartShape)
	stage.PartiallyGrowthCurve2DRibbonStartShapes_referenceOrder = make(map[*PartiallyGrowthCurve2DRibbonStartShape]uint) // diff Unstage needs the reference order
	stage.PartiallyGrowthCurve2DRibbonStartShapes_instance = make(map[*PartiallyGrowthCurve2DRibbonStartShape]*PartiallyGrowthCurve2DRibbonStartShape)
	for instance := range stage.PartiallyGrowthCurve2DRibbonStartShapes {
		_copy := instance.GongCopy().(*PartiallyGrowthCurve2DRibbonStartShape)
		stage.PartiallyGrowthCurve2DRibbonStartShapes_reference[instance] = _copy
		stage.PartiallyGrowthCurve2DRibbonStartShapes_instance[_copy] = instance
		stage.PartiallyGrowthCurve2DRibbonStartShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.RhombusStuffs_reference = make(map[*RhombusStuff]*RhombusStuff)
	stage.RhombusStuffs_referenceOrder = make(map[*RhombusStuff]uint) // diff Unstage needs the reference order
	stage.RhombusStuffs_instance = make(map[*RhombusStuff]*RhombusStuff)
	for instance := range stage.RhombusStuffs {
		_copy := instance.GongCopy().(*RhombusStuff)
		stage.RhombusStuffs_reference[instance] = _copy
		stage.RhombusStuffs_instance[_copy] = instance
		stage.RhombusStuffs_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.ShiftedBottomTopStartArcShapes_reference = make(map[*ShiftedBottomTopStartArcShape]*ShiftedBottomTopStartArcShape)
	stage.ShiftedBottomTopStartArcShapes_referenceOrder = make(map[*ShiftedBottomTopStartArcShape]uint) // diff Unstage needs the reference order
	stage.ShiftedBottomTopStartArcShapes_instance = make(map[*ShiftedBottomTopStartArcShape]*ShiftedBottomTopStartArcShape)
	for instance := range stage.ShiftedBottomTopStartArcShapes {
		_copy := instance.GongCopy().(*ShiftedBottomTopStartArcShape)
		stage.ShiftedBottomTopStartArcShapes_reference[instance] = _copy
		stage.ShiftedBottomTopStartArcShapes_instance[_copy] = instance
		stage.ShiftedBottomTopStartArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShiftedBottomTopStartArcShapeGrids_reference = make(map[*ShiftedBottomTopStartArcShapeGrid]*ShiftedBottomTopStartArcShapeGrid)
	stage.ShiftedBottomTopStartArcShapeGrids_referenceOrder = make(map[*ShiftedBottomTopStartArcShapeGrid]uint) // diff Unstage needs the reference order
	stage.ShiftedBottomTopStartArcShapeGrids_instance = make(map[*ShiftedBottomTopStartArcShapeGrid]*ShiftedBottomTopStartArcShapeGrid)
	for instance := range stage.ShiftedBottomTopStartArcShapeGrids {
		_copy := instance.GongCopy().(*ShiftedBottomTopStartArcShapeGrid)
		stage.ShiftedBottomTopStartArcShapeGrids_reference[instance] = _copy
		stage.ShiftedBottomTopStartArcShapeGrids_instance[_copy] = instance
		stage.ShiftedBottomTopStartArcShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShiftedLeftStackGrowthCurveEndArcShapes_reference = make(map[*ShiftedLeftStackGrowthCurveEndArcShape]*ShiftedLeftStackGrowthCurveEndArcShape)
	stage.ShiftedLeftStackGrowthCurveEndArcShapes_referenceOrder = make(map[*ShiftedLeftStackGrowthCurveEndArcShape]uint) // diff Unstage needs the reference order
	stage.ShiftedLeftStackGrowthCurveEndArcShapes_instance = make(map[*ShiftedLeftStackGrowthCurveEndArcShape]*ShiftedLeftStackGrowthCurveEndArcShape)
	for instance := range stage.ShiftedLeftStackGrowthCurveEndArcShapes {
		_copy := instance.GongCopy().(*ShiftedLeftStackGrowthCurveEndArcShape)
		stage.ShiftedLeftStackGrowthCurveEndArcShapes_reference[instance] = _copy
		stage.ShiftedLeftStackGrowthCurveEndArcShapes_instance[_copy] = instance
		stage.ShiftedLeftStackGrowthCurveEndArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShiftedLeftStackGrowthCurveStartArcShapes_reference = make(map[*ShiftedLeftStackGrowthCurveStartArcShape]*ShiftedLeftStackGrowthCurveStartArcShape)
	stage.ShiftedLeftStackGrowthCurveStartArcShapes_referenceOrder = make(map[*ShiftedLeftStackGrowthCurveStartArcShape]uint) // diff Unstage needs the reference order
	stage.ShiftedLeftStackGrowthCurveStartArcShapes_instance = make(map[*ShiftedLeftStackGrowthCurveStartArcShape]*ShiftedLeftStackGrowthCurveStartArcShape)
	for instance := range stage.ShiftedLeftStackGrowthCurveStartArcShapes {
		_copy := instance.GongCopy().(*ShiftedLeftStackGrowthCurveStartArcShape)
		stage.ShiftedLeftStackGrowthCurveStartArcShapes_reference[instance] = _copy
		stage.ShiftedLeftStackGrowthCurveStartArcShapes_instance[_copy] = instance
		stage.ShiftedLeftStackGrowthCurveStartArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShiftedLeftStackNormalVectors_reference = make(map[*ShiftedLeftStackNormalVector]*ShiftedLeftStackNormalVector)
	stage.ShiftedLeftStackNormalVectors_referenceOrder = make(map[*ShiftedLeftStackNormalVector]uint) // diff Unstage needs the reference order
	stage.ShiftedLeftStackNormalVectors_instance = make(map[*ShiftedLeftStackNormalVector]*ShiftedLeftStackNormalVector)
	for instance := range stage.ShiftedLeftStackNormalVectors {
		_copy := instance.GongCopy().(*ShiftedLeftStackNormalVector)
		stage.ShiftedLeftStackNormalVectors_reference[instance] = _copy
		stage.ShiftedLeftStackNormalVectors_instance[_copy] = instance
		stage.ShiftedLeftStackNormalVectors_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShiftedLeftStackOfGrowthCurves_reference = make(map[*ShiftedLeftStackOfGrowthCurve]*ShiftedLeftStackOfGrowthCurve)
	stage.ShiftedLeftStackOfGrowthCurves_referenceOrder = make(map[*ShiftedLeftStackOfGrowthCurve]uint) // diff Unstage needs the reference order
	stage.ShiftedLeftStackOfGrowthCurves_instance = make(map[*ShiftedLeftStackOfGrowthCurve]*ShiftedLeftStackOfGrowthCurve)
	for instance := range stage.ShiftedLeftStackOfGrowthCurves {
		_copy := instance.GongCopy().(*ShiftedLeftStackOfGrowthCurve)
		stage.ShiftedLeftStackOfGrowthCurves_reference[instance] = _copy
		stage.ShiftedLeftStackOfGrowthCurves_instance[_copy] = instance
		stage.ShiftedLeftStackOfGrowthCurves_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShiftedLeftStackOfNormalVectors_reference = make(map[*ShiftedLeftStackOfNormalVector]*ShiftedLeftStackOfNormalVector)
	stage.ShiftedLeftStackOfNormalVectors_referenceOrder = make(map[*ShiftedLeftStackOfNormalVector]uint) // diff Unstage needs the reference order
	stage.ShiftedLeftStackOfNormalVectors_instance = make(map[*ShiftedLeftStackOfNormalVector]*ShiftedLeftStackOfNormalVector)
	for instance := range stage.ShiftedLeftStackOfNormalVectors {
		_copy := instance.GongCopy().(*ShiftedLeftStackOfNormalVector)
		stage.ShiftedLeftStackOfNormalVectors_reference[instance] = _copy
		stage.ShiftedLeftStackOfNormalVectors_instance[_copy] = instance
		stage.ShiftedLeftStackOfNormalVectors_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackGrowthCurve2DEndHalfwayArcShapes_reference = make(map[*StackGrowthCurve2DEndHalfwayArcShape]*StackGrowthCurve2DEndHalfwayArcShape)
	stage.StackGrowthCurve2DEndHalfwayArcShapes_referenceOrder = make(map[*StackGrowthCurve2DEndHalfwayArcShape]uint) // diff Unstage needs the reference order
	stage.StackGrowthCurve2DEndHalfwayArcShapes_instance = make(map[*StackGrowthCurve2DEndHalfwayArcShape]*StackGrowthCurve2DEndHalfwayArcShape)
	for instance := range stage.StackGrowthCurve2DEndHalfwayArcShapes {
		_copy := instance.GongCopy().(*StackGrowthCurve2DEndHalfwayArcShape)
		stage.StackGrowthCurve2DEndHalfwayArcShapes_reference[instance] = _copy
		stage.StackGrowthCurve2DEndHalfwayArcShapes_instance[_copy] = instance
		stage.StackGrowthCurve2DEndHalfwayArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackGrowthCurve2DRibbonEndShapes_reference = make(map[*StackGrowthCurve2DRibbonEndShape]*StackGrowthCurve2DRibbonEndShape)
	stage.StackGrowthCurve2DRibbonEndShapes_referenceOrder = make(map[*StackGrowthCurve2DRibbonEndShape]uint) // diff Unstage needs the reference order
	stage.StackGrowthCurve2DRibbonEndShapes_instance = make(map[*StackGrowthCurve2DRibbonEndShape]*StackGrowthCurve2DRibbonEndShape)
	for instance := range stage.StackGrowthCurve2DRibbonEndShapes {
		_copy := instance.GongCopy().(*StackGrowthCurve2DRibbonEndShape)
		stage.StackGrowthCurve2DRibbonEndShapes_reference[instance] = _copy
		stage.StackGrowthCurve2DRibbonEndShapes_instance[_copy] = instance
		stage.StackGrowthCurve2DRibbonEndShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackGrowthCurve2DRibbonStartShapes_reference = make(map[*StackGrowthCurve2DRibbonStartShape]*StackGrowthCurve2DRibbonStartShape)
	stage.StackGrowthCurve2DRibbonStartShapes_referenceOrder = make(map[*StackGrowthCurve2DRibbonStartShape]uint) // diff Unstage needs the reference order
	stage.StackGrowthCurve2DRibbonStartShapes_instance = make(map[*StackGrowthCurve2DRibbonStartShape]*StackGrowthCurve2DRibbonStartShape)
	for instance := range stage.StackGrowthCurve2DRibbonStartShapes {
		_copy := instance.GongCopy().(*StackGrowthCurve2DRibbonStartShape)
		stage.StackGrowthCurve2DRibbonStartShapes_reference[instance] = _copy
		stage.StackGrowthCurve2DRibbonStartShapes_instance[_copy] = instance
		stage.StackGrowthCurve2DRibbonStartShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackGrowthCurve2DStartHalfwayArcShapes_reference = make(map[*StackGrowthCurve2DStartHalfwayArcShape]*StackGrowthCurve2DStartHalfwayArcShape)
	stage.StackGrowthCurve2DStartHalfwayArcShapes_referenceOrder = make(map[*StackGrowthCurve2DStartHalfwayArcShape]uint) // diff Unstage needs the reference order
	stage.StackGrowthCurve2DStartHalfwayArcShapes_instance = make(map[*StackGrowthCurve2DStartHalfwayArcShape]*StackGrowthCurve2DStartHalfwayArcShape)
	for instance := range stage.StackGrowthCurve2DStartHalfwayArcShapes {
		_copy := instance.GongCopy().(*StackGrowthCurve2DStartHalfwayArcShape)
		stage.StackGrowthCurve2DStartHalfwayArcShapes_reference[instance] = _copy
		stage.StackGrowthCurve2DStartHalfwayArcShapes_instance[_copy] = instance
		stage.StackGrowthCurve2DStartHalfwayArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackOfGrowthCurve2Ds_reference = make(map[*StackOfGrowthCurve2D]*StackOfGrowthCurve2D)
	stage.StackOfGrowthCurve2Ds_referenceOrder = make(map[*StackOfGrowthCurve2D]uint) // diff Unstage needs the reference order
	stage.StackOfGrowthCurve2Ds_instance = make(map[*StackOfGrowthCurve2D]*StackOfGrowthCurve2D)
	for instance := range stage.StackOfGrowthCurve2Ds {
		_copy := instance.GongCopy().(*StackOfGrowthCurve2D)
		stage.StackOfGrowthCurve2Ds_reference[instance] = _copy
		stage.StackOfGrowthCurve2Ds_instance[_copy] = instance
		stage.StackOfGrowthCurve2Ds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackOfGrowthCurve2DRibbons_reference = make(map[*StackOfGrowthCurve2DRibbon]*StackOfGrowthCurve2DRibbon)
	stage.StackOfGrowthCurve2DRibbons_referenceOrder = make(map[*StackOfGrowthCurve2DRibbon]uint) // diff Unstage needs the reference order
	stage.StackOfGrowthCurve2DRibbons_instance = make(map[*StackOfGrowthCurve2DRibbon]*StackOfGrowthCurve2DRibbon)
	for instance := range stage.StackOfGrowthCurve2DRibbons {
		_copy := instance.GongCopy().(*StackOfGrowthCurve2DRibbon)
		stage.StackOfGrowthCurve2DRibbons_reference[instance] = _copy
		stage.StackOfGrowthCurve2DRibbons_instance[_copy] = instance
		stage.StackOfGrowthCurve2DRibbons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackOfRotatedGrowthCurve2Ds_reference = make(map[*StackOfRotatedGrowthCurve2D]*StackOfRotatedGrowthCurve2D)
	stage.StackOfRotatedGrowthCurve2Ds_referenceOrder = make(map[*StackOfRotatedGrowthCurve2D]uint) // diff Unstage needs the reference order
	stage.StackOfRotatedGrowthCurve2Ds_instance = make(map[*StackOfRotatedGrowthCurve2D]*StackOfRotatedGrowthCurve2D)
	for instance := range stage.StackOfRotatedGrowthCurve2Ds {
		_copy := instance.GongCopy().(*StackOfRotatedGrowthCurve2D)
		stage.StackOfRotatedGrowthCurve2Ds_reference[instance] = _copy
		stage.StackOfRotatedGrowthCurve2Ds_instance[_copy] = instance
		stage.StackOfRotatedGrowthCurve2Ds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackOfRotatedGrowthCurve2DRibbons_reference = make(map[*StackOfRotatedGrowthCurve2DRibbon]*StackOfRotatedGrowthCurve2DRibbon)
	stage.StackOfRotatedGrowthCurve2DRibbons_referenceOrder = make(map[*StackOfRotatedGrowthCurve2DRibbon]uint) // diff Unstage needs the reference order
	stage.StackOfRotatedGrowthCurve2DRibbons_instance = make(map[*StackOfRotatedGrowthCurve2DRibbon]*StackOfRotatedGrowthCurve2DRibbon)
	for instance := range stage.StackOfRotatedGrowthCurve2DRibbons {
		_copy := instance.GongCopy().(*StackOfRotatedGrowthCurve2DRibbon)
		stage.StackOfRotatedGrowthCurve2DRibbons_reference[instance] = _copy
		stage.StackOfRotatedGrowthCurve2DRibbons_instance[_copy] = instance
		stage.StackOfRotatedGrowthCurve2DRibbons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackRotatedGrowthCurve2DEndArcShapes_reference = make(map[*StackRotatedGrowthCurve2DEndArcShape]*StackRotatedGrowthCurve2DEndArcShape)
	stage.StackRotatedGrowthCurve2DEndArcShapes_referenceOrder = make(map[*StackRotatedGrowthCurve2DEndArcShape]uint) // diff Unstage needs the reference order
	stage.StackRotatedGrowthCurve2DEndArcShapes_instance = make(map[*StackRotatedGrowthCurve2DEndArcShape]*StackRotatedGrowthCurve2DEndArcShape)
	for instance := range stage.StackRotatedGrowthCurve2DEndArcShapes {
		_copy := instance.GongCopy().(*StackRotatedGrowthCurve2DEndArcShape)
		stage.StackRotatedGrowthCurve2DEndArcShapes_reference[instance] = _copy
		stage.StackRotatedGrowthCurve2DEndArcShapes_instance[_copy] = instance
		stage.StackRotatedGrowthCurve2DEndArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackRotatedGrowthCurve2DRibbonEndShapes_reference = make(map[*StackRotatedGrowthCurve2DRibbonEndShape]*StackRotatedGrowthCurve2DRibbonEndShape)
	stage.StackRotatedGrowthCurve2DRibbonEndShapes_referenceOrder = make(map[*StackRotatedGrowthCurve2DRibbonEndShape]uint) // diff Unstage needs the reference order
	stage.StackRotatedGrowthCurve2DRibbonEndShapes_instance = make(map[*StackRotatedGrowthCurve2DRibbonEndShape]*StackRotatedGrowthCurve2DRibbonEndShape)
	for instance := range stage.StackRotatedGrowthCurve2DRibbonEndShapes {
		_copy := instance.GongCopy().(*StackRotatedGrowthCurve2DRibbonEndShape)
		stage.StackRotatedGrowthCurve2DRibbonEndShapes_reference[instance] = _copy
		stage.StackRotatedGrowthCurve2DRibbonEndShapes_instance[_copy] = instance
		stage.StackRotatedGrowthCurve2DRibbonEndShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackRotatedGrowthCurve2DRibbonStartShapes_reference = make(map[*StackRotatedGrowthCurve2DRibbonStartShape]*StackRotatedGrowthCurve2DRibbonStartShape)
	stage.StackRotatedGrowthCurve2DRibbonStartShapes_referenceOrder = make(map[*StackRotatedGrowthCurve2DRibbonStartShape]uint) // diff Unstage needs the reference order
	stage.StackRotatedGrowthCurve2DRibbonStartShapes_instance = make(map[*StackRotatedGrowthCurve2DRibbonStartShape]*StackRotatedGrowthCurve2DRibbonStartShape)
	for instance := range stage.StackRotatedGrowthCurve2DRibbonStartShapes {
		_copy := instance.GongCopy().(*StackRotatedGrowthCurve2DRibbonStartShape)
		stage.StackRotatedGrowthCurve2DRibbonStartShapes_reference[instance] = _copy
		stage.StackRotatedGrowthCurve2DRibbonStartShapes_instance[_copy] = instance
		stage.StackRotatedGrowthCurve2DRibbonStartShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackRotatedGrowthCurve2DStartArcShapes_reference = make(map[*StackRotatedGrowthCurve2DStartArcShape]*StackRotatedGrowthCurve2DStartArcShape)
	stage.StackRotatedGrowthCurve2DStartArcShapes_referenceOrder = make(map[*StackRotatedGrowthCurve2DStartArcShape]uint) // diff Unstage needs the reference order
	stage.StackRotatedGrowthCurve2DStartArcShapes_instance = make(map[*StackRotatedGrowthCurve2DStartArcShape]*StackRotatedGrowthCurve2DStartArcShape)
	for instance := range stage.StackRotatedGrowthCurve2DStartArcShapes {
		_copy := instance.GongCopy().(*StackRotatedGrowthCurve2DStartArcShape)
		stage.StackRotatedGrowthCurve2DStartArcShapes_reference[instance] = _copy
		stage.StackRotatedGrowthCurve2DStartArcShapes_instance[_copy] = instance
		stage.StackRotatedGrowthCurve2DStartArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.StartHalfwayArcShapes_reference = make(map[*StartHalfwayArcShape]*StartHalfwayArcShape)
	stage.StartHalfwayArcShapes_referenceOrder = make(map[*StartHalfwayArcShape]uint) // diff Unstage needs the reference order
	stage.StartHalfwayArcShapes_instance = make(map[*StartHalfwayArcShape]*StartHalfwayArcShape)
	for instance := range stage.StartHalfwayArcShapes {
		_copy := instance.GongCopy().(*StartHalfwayArcShape)
		stage.StartHalfwayArcShapes_reference[instance] = _copy
		stage.StartHalfwayArcShapes_instance[_copy] = instance
		stage.StartHalfwayArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StartHalfwayArcShapeGrids_reference = make(map[*StartHalfwayArcShapeGrid]*StartHalfwayArcShapeGrid)
	stage.StartHalfwayArcShapeGrids_referenceOrder = make(map[*StartHalfwayArcShapeGrid]uint) // diff Unstage needs the reference order
	stage.StartHalfwayArcShapeGrids_instance = make(map[*StartHalfwayArcShapeGrid]*StartHalfwayArcShapeGrid)
	for instance := range stage.StartHalfwayArcShapeGrids {
		_copy := instance.GongCopy().(*StartHalfwayArcShapeGrid)
		stage.StartHalfwayArcShapeGrids_reference[instance] = _copy
		stage.StartHalfwayArcShapeGrids_instance[_copy] = instance
		stage.StartHalfwayArcShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopEndArcShapes_reference = make(map[*TopEndArcShape]*TopEndArcShape)
	stage.TopEndArcShapes_referenceOrder = make(map[*TopEndArcShape]uint) // diff Unstage needs the reference order
	stage.TopEndArcShapes_instance = make(map[*TopEndArcShape]*TopEndArcShape)
	for instance := range stage.TopEndArcShapes {
		_copy := instance.GongCopy().(*TopEndArcShape)
		stage.TopEndArcShapes_reference[instance] = _copy
		stage.TopEndArcShapes_instance[_copy] = instance
		stage.TopEndArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopEndArcShapeGrids_reference = make(map[*TopEndArcShapeGrid]*TopEndArcShapeGrid)
	stage.TopEndArcShapeGrids_referenceOrder = make(map[*TopEndArcShapeGrid]uint) // diff Unstage needs the reference order
	stage.TopEndArcShapeGrids_instance = make(map[*TopEndArcShapeGrid]*TopEndArcShapeGrid)
	for instance := range stage.TopEndArcShapeGrids {
		_copy := instance.GongCopy().(*TopEndArcShapeGrid)
		stage.TopEndArcShapeGrids_reference[instance] = _copy
		stage.TopEndArcShapeGrids_instance[_copy] = instance
		stage.TopEndArcShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopEndHalfwayArcShapes_reference = make(map[*TopEndHalfwayArcShape]*TopEndHalfwayArcShape)
	stage.TopEndHalfwayArcShapes_referenceOrder = make(map[*TopEndHalfwayArcShape]uint) // diff Unstage needs the reference order
	stage.TopEndHalfwayArcShapes_instance = make(map[*TopEndHalfwayArcShape]*TopEndHalfwayArcShape)
	for instance := range stage.TopEndHalfwayArcShapes {
		_copy := instance.GongCopy().(*TopEndHalfwayArcShape)
		stage.TopEndHalfwayArcShapes_reference[instance] = _copy
		stage.TopEndHalfwayArcShapes_instance[_copy] = instance
		stage.TopEndHalfwayArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopEndHalfwayArcShapeGrids_reference = make(map[*TopEndHalfwayArcShapeGrid]*TopEndHalfwayArcShapeGrid)
	stage.TopEndHalfwayArcShapeGrids_referenceOrder = make(map[*TopEndHalfwayArcShapeGrid]uint) // diff Unstage needs the reference order
	stage.TopEndHalfwayArcShapeGrids_instance = make(map[*TopEndHalfwayArcShapeGrid]*TopEndHalfwayArcShapeGrid)
	for instance := range stage.TopEndHalfwayArcShapeGrids {
		_copy := instance.GongCopy().(*TopEndHalfwayArcShapeGrid)
		stage.TopEndHalfwayArcShapeGrids_reference[instance] = _copy
		stage.TopEndHalfwayArcShapeGrids_instance[_copy] = instance
		stage.TopEndHalfwayArcShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopGrowthCurve2Ds_reference = make(map[*TopGrowthCurve2D]*TopGrowthCurve2D)
	stage.TopGrowthCurve2Ds_referenceOrder = make(map[*TopGrowthCurve2D]uint) // diff Unstage needs the reference order
	stage.TopGrowthCurve2Ds_instance = make(map[*TopGrowthCurve2D]*TopGrowthCurve2D)
	for instance := range stage.TopGrowthCurve2Ds {
		_copy := instance.GongCopy().(*TopGrowthCurve2D)
		stage.TopGrowthCurve2Ds_reference[instance] = _copy
		stage.TopGrowthCurve2Ds_instance[_copy] = instance
		stage.TopGrowthCurve2Ds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopMidArcVectorShapes_reference = make(map[*TopMidArcVectorShape]*TopMidArcVectorShape)
	stage.TopMidArcVectorShapes_referenceOrder = make(map[*TopMidArcVectorShape]uint) // diff Unstage needs the reference order
	stage.TopMidArcVectorShapes_instance = make(map[*TopMidArcVectorShape]*TopMidArcVectorShape)
	for instance := range stage.TopMidArcVectorShapes {
		_copy := instance.GongCopy().(*TopMidArcVectorShape)
		stage.TopMidArcVectorShapes_reference[instance] = _copy
		stage.TopMidArcVectorShapes_instance[_copy] = instance
		stage.TopMidArcVectorShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopMidArcVectorShapeGrids_reference = make(map[*TopMidArcVectorShapeGrid]*TopMidArcVectorShapeGrid)
	stage.TopMidArcVectorShapeGrids_referenceOrder = make(map[*TopMidArcVectorShapeGrid]uint) // diff Unstage needs the reference order
	stage.TopMidArcVectorShapeGrids_instance = make(map[*TopMidArcVectorShapeGrid]*TopMidArcVectorShapeGrid)
	for instance := range stage.TopMidArcVectorShapeGrids {
		_copy := instance.GongCopy().(*TopMidArcVectorShapeGrid)
		stage.TopMidArcVectorShapeGrids_reference[instance] = _copy
		stage.TopMidArcVectorShapeGrids_instance[_copy] = instance
		stage.TopMidArcVectorShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStackGrowthCurve2DEndHalfwayArcShapes_reference = make(map[*TopStackGrowthCurve2DEndHalfwayArcShape]*TopStackGrowthCurve2DEndHalfwayArcShape)
	stage.TopStackGrowthCurve2DEndHalfwayArcShapes_referenceOrder = make(map[*TopStackGrowthCurve2DEndHalfwayArcShape]uint) // diff Unstage needs the reference order
	stage.TopStackGrowthCurve2DEndHalfwayArcShapes_instance = make(map[*TopStackGrowthCurve2DEndHalfwayArcShape]*TopStackGrowthCurve2DEndHalfwayArcShape)
	for instance := range stage.TopStackGrowthCurve2DEndHalfwayArcShapes {
		_copy := instance.GongCopy().(*TopStackGrowthCurve2DEndHalfwayArcShape)
		stage.TopStackGrowthCurve2DEndHalfwayArcShapes_reference[instance] = _copy
		stage.TopStackGrowthCurve2DEndHalfwayArcShapes_instance[_copy] = instance
		stage.TopStackGrowthCurve2DEndHalfwayArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStackGrowthCurve2DStartHalfwayArcShapes_reference = make(map[*TopStackGrowthCurve2DStartHalfwayArcShape]*TopStackGrowthCurve2DStartHalfwayArcShape)
	stage.TopStackGrowthCurve2DStartHalfwayArcShapes_referenceOrder = make(map[*TopStackGrowthCurve2DStartHalfwayArcShape]uint) // diff Unstage needs the reference order
	stage.TopStackGrowthCurve2DStartHalfwayArcShapes_instance = make(map[*TopStackGrowthCurve2DStartHalfwayArcShape]*TopStackGrowthCurve2DStartHalfwayArcShape)
	for instance := range stage.TopStackGrowthCurve2DStartHalfwayArcShapes {
		_copy := instance.GongCopy().(*TopStackGrowthCurve2DStartHalfwayArcShape)
		stage.TopStackGrowthCurve2DStartHalfwayArcShapes_reference[instance] = _copy
		stage.TopStackGrowthCurve2DStartHalfwayArcShapes_instance[_copy] = instance
		stage.TopStackGrowthCurve2DStartHalfwayArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStackOfGrowthCurve2Ds_reference = make(map[*TopStackOfGrowthCurve2D]*TopStackOfGrowthCurve2D)
	stage.TopStackOfGrowthCurve2Ds_referenceOrder = make(map[*TopStackOfGrowthCurve2D]uint) // diff Unstage needs the reference order
	stage.TopStackOfGrowthCurve2Ds_instance = make(map[*TopStackOfGrowthCurve2D]*TopStackOfGrowthCurve2D)
	for instance := range stage.TopStackOfGrowthCurve2Ds {
		_copy := instance.GongCopy().(*TopStackOfGrowthCurve2D)
		stage.TopStackOfGrowthCurve2Ds_reference[instance] = _copy
		stage.TopStackOfGrowthCurve2Ds_instance[_copy] = instance
		stage.TopStackOfGrowthCurve2Ds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStackOfRotatedGrowthCurve2Ds_reference = make(map[*TopStackOfRotatedGrowthCurve2D]*TopStackOfRotatedGrowthCurve2D)
	stage.TopStackOfRotatedGrowthCurve2Ds_referenceOrder = make(map[*TopStackOfRotatedGrowthCurve2D]uint) // diff Unstage needs the reference order
	stage.TopStackOfRotatedGrowthCurve2Ds_instance = make(map[*TopStackOfRotatedGrowthCurve2D]*TopStackOfRotatedGrowthCurve2D)
	for instance := range stage.TopStackOfRotatedGrowthCurve2Ds {
		_copy := instance.GongCopy().(*TopStackOfRotatedGrowthCurve2D)
		stage.TopStackOfRotatedGrowthCurve2Ds_reference[instance] = _copy
		stage.TopStackOfRotatedGrowthCurve2Ds_instance[_copy] = instance
		stage.TopStackOfRotatedGrowthCurve2Ds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_reference = make(map[*TopStackOfRotatedGrowthCurve2DEndArcShape]*TopStackOfRotatedGrowthCurve2DEndArcShape)
	stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_referenceOrder = make(map[*TopStackOfRotatedGrowthCurve2DEndArcShape]uint) // diff Unstage needs the reference order
	stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_instance = make(map[*TopStackOfRotatedGrowthCurve2DEndArcShape]*TopStackOfRotatedGrowthCurve2DEndArcShape)
	for instance := range stage.TopStackOfRotatedGrowthCurve2DEndArcShapes {
		_copy := instance.GongCopy().(*TopStackOfRotatedGrowthCurve2DEndArcShape)
		stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_reference[instance] = _copy
		stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_instance[_copy] = instance
		stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_reference = make(map[*TopStackOfRotatedGrowthCurve2DStartArcShape]*TopStackOfRotatedGrowthCurve2DStartArcShape)
	stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_referenceOrder = make(map[*TopStackOfRotatedGrowthCurve2DStartArcShape]uint) // diff Unstage needs the reference order
	stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_instance = make(map[*TopStackOfRotatedGrowthCurve2DStartArcShape]*TopStackOfRotatedGrowthCurve2DStartArcShape)
	for instance := range stage.TopStackOfRotatedGrowthCurve2DStartArcShapes {
		_copy := instance.GongCopy().(*TopStackOfRotatedGrowthCurve2DStartArcShape)
		stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_reference[instance] = _copy
		stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_instance[_copy] = instance
		stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStartArcShapes_reference = make(map[*TopStartArcShape]*TopStartArcShape)
	stage.TopStartArcShapes_referenceOrder = make(map[*TopStartArcShape]uint) // diff Unstage needs the reference order
	stage.TopStartArcShapes_instance = make(map[*TopStartArcShape]*TopStartArcShape)
	for instance := range stage.TopStartArcShapes {
		_copy := instance.GongCopy().(*TopStartArcShape)
		stage.TopStartArcShapes_reference[instance] = _copy
		stage.TopStartArcShapes_instance[_copy] = instance
		stage.TopStartArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStartArcShapeGrids_reference = make(map[*TopStartArcShapeGrid]*TopStartArcShapeGrid)
	stage.TopStartArcShapeGrids_referenceOrder = make(map[*TopStartArcShapeGrid]uint) // diff Unstage needs the reference order
	stage.TopStartArcShapeGrids_instance = make(map[*TopStartArcShapeGrid]*TopStartArcShapeGrid)
	for instance := range stage.TopStartArcShapeGrids {
		_copy := instance.GongCopy().(*TopStartArcShapeGrid)
		stage.TopStartArcShapeGrids_reference[instance] = _copy
		stage.TopStartArcShapeGrids_instance[_copy] = instance
		stage.TopStartArcShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStartHalfwayArcShapes_reference = make(map[*TopStartHalfwayArcShape]*TopStartHalfwayArcShape)
	stage.TopStartHalfwayArcShapes_referenceOrder = make(map[*TopStartHalfwayArcShape]uint) // diff Unstage needs the reference order
	stage.TopStartHalfwayArcShapes_instance = make(map[*TopStartHalfwayArcShape]*TopStartHalfwayArcShape)
	for instance := range stage.TopStartHalfwayArcShapes {
		_copy := instance.GongCopy().(*TopStartHalfwayArcShape)
		stage.TopStartHalfwayArcShapes_reference[instance] = _copy
		stage.TopStartHalfwayArcShapes_instance[_copy] = instance
		stage.TopStartHalfwayArcShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopStartHalfwayArcShapeGrids_reference = make(map[*TopStartHalfwayArcShapeGrid]*TopStartHalfwayArcShapeGrid)
	stage.TopStartHalfwayArcShapeGrids_referenceOrder = make(map[*TopStartHalfwayArcShapeGrid]uint) // diff Unstage needs the reference order
	stage.TopStartHalfwayArcShapeGrids_instance = make(map[*TopStartHalfwayArcShapeGrid]*TopStartHalfwayArcShapeGrid)
	for instance := range stage.TopStartHalfwayArcShapeGrids {
		_copy := instance.GongCopy().(*TopStartHalfwayArcShapeGrid)
		stage.TopStartHalfwayArcShapeGrids_reference[instance] = _copy
		stage.TopStartHalfwayArcShapeGrids_instance[_copy] = instance
		stage.TopStartHalfwayArcShapeGrids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TorusStackShapes_reference = make(map[*TorusStackShape]*TorusStackShape)
	stage.TorusStackShapes_referenceOrder = make(map[*TorusStackShape]uint) // diff Unstage needs the reference order
	stage.TorusStackShapes_instance = make(map[*TorusStackShape]*TorusStackShape)
	for instance := range stage.TorusStackShapes {
		_copy := instance.GongCopy().(*TorusStackShape)
		stage.TorusStackShapes_reference[instance] = _copy
		stage.TorusStackShapes_instance[_copy] = instance
		stage.TorusStackShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.VerticalTorusStackShapes_reference = make(map[*VerticalTorusStackShape]*VerticalTorusStackShape)
	stage.VerticalTorusStackShapes_referenceOrder = make(map[*VerticalTorusStackShape]uint) // diff Unstage needs the reference order
	stage.VerticalTorusStackShapes_instance = make(map[*VerticalTorusStackShape]*VerticalTorusStackShape)
	for instance := range stage.VerticalTorusStackShapes {
		_copy := instance.GongCopy().(*VerticalTorusStackShape)
		stage.VerticalTorusStackShapes_reference[instance] = _copy
		stage.VerticalTorusStackShapes_instance[_copy] = instance
		stage.VerticalTorusStackShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	for instance := range stage.EndHalfwayArcShapes {
		reference := stage.EndHalfwayArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EndHalfwayArcShapeGrids {
		reference := stage.EndHalfwayArcShapeGrids_reference[instance]
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

	for instance := range stage.GrowthCurve2Ds {
		reference := stage.GrowthCurve2Ds_reference[instance]
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

	for instance := range stage.MidArcVectorShapes {
		reference := stage.MidArcVectorShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MidArcVectorShapeGrids {
		reference := stage.MidArcVectorShapeGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DRibbons {
		reference := stage.PartiallyGrowthCurve2DRibbons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DRibbonEndShapes {
		reference := stage.PartiallyGrowthCurve2DRibbonEndShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DRibbonStartShapes {
		reference := stage.PartiallyGrowthCurve2DRibbonStartShapes_reference[instance]
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

	for instance := range stage.RhombusStuffs {
		reference := stage.RhombusStuffs_reference[instance]
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

	for instance := range stage.ShiftedBottomTopStartArcShapes {
		reference := stage.ShiftedBottomTopStartArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShiftedBottomTopStartArcShapeGrids {
		reference := stage.ShiftedBottomTopStartArcShapeGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShiftedLeftStackGrowthCurveEndArcShapes {
		reference := stage.ShiftedLeftStackGrowthCurveEndArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShiftedLeftStackGrowthCurveStartArcShapes {
		reference := stage.ShiftedLeftStackGrowthCurveStartArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShiftedLeftStackNormalVectors {
		reference := stage.ShiftedLeftStackNormalVectors_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShiftedLeftStackOfGrowthCurves {
		reference := stage.ShiftedLeftStackOfGrowthCurves_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShiftedLeftStackOfNormalVectors {
		reference := stage.ShiftedLeftStackOfNormalVectors_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackGrowthCurve2DEndHalfwayArcShapes {
		reference := stage.StackGrowthCurve2DEndHalfwayArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackGrowthCurve2DRibbonEndShapes {
		reference := stage.StackGrowthCurve2DRibbonEndShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackGrowthCurve2DRibbonStartShapes {
		reference := stage.StackGrowthCurve2DRibbonStartShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackGrowthCurve2DStartHalfwayArcShapes {
		reference := stage.StackGrowthCurve2DStartHalfwayArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackOfGrowthCurve2Ds {
		reference := stage.StackOfGrowthCurve2Ds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackOfGrowthCurve2DRibbons {
		reference := stage.StackOfGrowthCurve2DRibbons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackOfRotatedGrowthCurve2Ds {
		reference := stage.StackOfRotatedGrowthCurve2Ds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackOfRotatedGrowthCurve2DRibbons {
		reference := stage.StackOfRotatedGrowthCurve2DRibbons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackRotatedGrowthCurve2DEndArcShapes {
		reference := stage.StackRotatedGrowthCurve2DEndArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackRotatedGrowthCurve2DRibbonEndShapes {
		reference := stage.StackRotatedGrowthCurve2DRibbonEndShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackRotatedGrowthCurve2DRibbonStartShapes {
		reference := stage.StackRotatedGrowthCurve2DRibbonStartShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackRotatedGrowthCurve2DStartArcShapes {
		reference := stage.StackRotatedGrowthCurve2DStartArcShapes_reference[instance]
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

	for instance := range stage.StartHalfwayArcShapes {
		reference := stage.StartHalfwayArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StartHalfwayArcShapeGrids {
		reference := stage.StartHalfwayArcShapeGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopEndArcShapes {
		reference := stage.TopEndArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopEndArcShapeGrids {
		reference := stage.TopEndArcShapeGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopEndHalfwayArcShapes {
		reference := stage.TopEndHalfwayArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopEndHalfwayArcShapeGrids {
		reference := stage.TopEndHalfwayArcShapeGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopGrowthCurve2Ds {
		reference := stage.TopGrowthCurve2Ds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopMidArcVectorShapes {
		reference := stage.TopMidArcVectorShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopMidArcVectorShapeGrids {
		reference := stage.TopMidArcVectorShapeGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStackGrowthCurve2DEndHalfwayArcShapes {
		reference := stage.TopStackGrowthCurve2DEndHalfwayArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStackGrowthCurve2DStartHalfwayArcShapes {
		reference := stage.TopStackGrowthCurve2DStartHalfwayArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStackOfGrowthCurve2Ds {
		reference := stage.TopStackOfGrowthCurve2Ds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStackOfRotatedGrowthCurve2Ds {
		reference := stage.TopStackOfRotatedGrowthCurve2Ds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStackOfRotatedGrowthCurve2DEndArcShapes {
		reference := stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStackOfRotatedGrowthCurve2DStartArcShapes {
		reference := stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStartArcShapes {
		reference := stage.TopStartArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStartArcShapeGrids {
		reference := stage.TopStartArcShapeGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStartHalfwayArcShapes {
		reference := stage.TopStartHalfwayArcShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopStartHalfwayArcShapeGrids {
		reference := stage.TopStartHalfwayArcShapeGrids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TorusStackShapes {
		reference := stage.TorusStackShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.VerticalTorusStackShapes {
		reference := stage.VerticalTorusStackShapes_reference[instance]
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

func (endhalfwayarcshape *EndHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EndHalfwayArcShape_stagedOrder[endhalfwayarcshape]; ok {
		return order
	}
	if order, ok := stage.EndHalfwayArcShapes_referenceOrder[endhalfwayarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type EndHalfwayArcShape was not staged and does not have a reference order", endhalfwayarcshape)
		return 0
	}
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EndHalfwayArcShapeGrid_stagedOrder[endhalfwayarcshapegrid]; ok {
		return order
	}
	if order, ok := stage.EndHalfwayArcShapeGrids_referenceOrder[endhalfwayarcshapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type EndHalfwayArcShapeGrid was not staged and does not have a reference order", endhalfwayarcshapegrid)
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

func (growthcurve2d *GrowthCurve2D) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GrowthCurve2D_stagedOrder[growthcurve2d]; ok {
		return order
	}
	if order, ok := stage.GrowthCurve2Ds_referenceOrder[growthcurve2d]; ok {
		return order
	} else {
		log.Printf("instance %p of type GrowthCurve2D was not staged and does not have a reference order", growthcurve2d)
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

func (midarcvectorshape *MidArcVectorShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MidArcVectorShape_stagedOrder[midarcvectorshape]; ok {
		return order
	}
	if order, ok := stage.MidArcVectorShapes_referenceOrder[midarcvectorshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type MidArcVectorShape was not staged and does not have a reference order", midarcvectorshape)
		return 0
	}
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MidArcVectorShapeGrid_stagedOrder[midarcvectorshapegrid]; ok {
		return order
	}
	if order, ok := stage.MidArcVectorShapeGrids_referenceOrder[midarcvectorshapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type MidArcVectorShapeGrid was not staged and does not have a reference order", midarcvectorshapegrid)
		return 0
	}
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyGrowthCurve2DRibbon_stagedOrder[partiallygrowthcurve2dribbon]; ok {
		return order
	}
	if order, ok := stage.PartiallyGrowthCurve2DRibbons_referenceOrder[partiallygrowthcurve2dribbon]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyGrowthCurve2DRibbon was not staged and does not have a reference order", partiallygrowthcurve2dribbon)
		return 0
	}
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyGrowthCurve2DRibbonEndShape_stagedOrder[partiallygrowthcurve2dribbonendshape]; ok {
		return order
	}
	if order, ok := stage.PartiallyGrowthCurve2DRibbonEndShapes_referenceOrder[partiallygrowthcurve2dribbonendshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyGrowthCurve2DRibbonEndShape was not staged and does not have a reference order", partiallygrowthcurve2dribbonendshape)
		return 0
	}
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyGrowthCurve2DRibbonStartShape_stagedOrder[partiallygrowthcurve2dribbonstartshape]; ok {
		return order
	}
	if order, ok := stage.PartiallyGrowthCurve2DRibbonStartShapes_referenceOrder[partiallygrowthcurve2dribbonstartshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyGrowthCurve2DRibbonStartShape was not staged and does not have a reference order", partiallygrowthcurve2dribbonstartshape)
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

func (rhombusstuff *RhombusStuff) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RhombusStuff_stagedOrder[rhombusstuff]; ok {
		return order
	}
	if order, ok := stage.RhombusStuffs_referenceOrder[rhombusstuff]; ok {
		return order
	} else {
		log.Printf("instance %p of type RhombusStuff was not staged and does not have a reference order", rhombusstuff)
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

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedBottomTopStartArcShape_stagedOrder[shiftedbottomtopstartarcshape]; ok {
		return order
	}
	if order, ok := stage.ShiftedBottomTopStartArcShapes_referenceOrder[shiftedbottomtopstartarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedBottomTopStartArcShape was not staged and does not have a reference order", shiftedbottomtopstartarcshape)
		return 0
	}
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedBottomTopStartArcShapeGrid_stagedOrder[shiftedbottomtopstartarcshapegrid]; ok {
		return order
	}
	if order, ok := stage.ShiftedBottomTopStartArcShapeGrids_referenceOrder[shiftedbottomtopstartarcshapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedBottomTopStartArcShapeGrid was not staged and does not have a reference order", shiftedbottomtopstartarcshapegrid)
		return 0
	}
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder[shiftedleftstackgrowthcurveendarcshape]; ok {
		return order
	}
	if order, ok := stage.ShiftedLeftStackGrowthCurveEndArcShapes_referenceOrder[shiftedleftstackgrowthcurveendarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedLeftStackGrowthCurveEndArcShape was not staged and does not have a reference order", shiftedleftstackgrowthcurveendarcshape)
		return 0
	}
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder[shiftedleftstackgrowthcurvestartarcshape]; ok {
		return order
	}
	if order, ok := stage.ShiftedLeftStackGrowthCurveStartArcShapes_referenceOrder[shiftedleftstackgrowthcurvestartarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedLeftStackGrowthCurveStartArcShape was not staged and does not have a reference order", shiftedleftstackgrowthcurvestartarcshape)
		return 0
	}
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedLeftStackNormalVector_stagedOrder[shiftedleftstacknormalvector]; ok {
		return order
	}
	if order, ok := stage.ShiftedLeftStackNormalVectors_referenceOrder[shiftedleftstacknormalvector]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedLeftStackNormalVector was not staged and does not have a reference order", shiftedleftstacknormalvector)
		return 0
	}
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedLeftStackOfGrowthCurve_stagedOrder[shiftedleftstackofgrowthcurve]; ok {
		return order
	}
	if order, ok := stage.ShiftedLeftStackOfGrowthCurves_referenceOrder[shiftedleftstackofgrowthcurve]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedLeftStackOfGrowthCurve was not staged and does not have a reference order", shiftedleftstackofgrowthcurve)
		return 0
	}
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedLeftStackOfNormalVector_stagedOrder[shiftedleftstackofnormalvector]; ok {
		return order
	}
	if order, ok := stage.ShiftedLeftStackOfNormalVectors_referenceOrder[shiftedleftstackofnormalvector]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedLeftStackOfNormalVector was not staged and does not have a reference order", shiftedleftstackofnormalvector)
		return 0
	}
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackGrowthCurve2DEndHalfwayArcShape_stagedOrder[stackgrowthcurve2dendhalfwayarcshape]; ok {
		return order
	}
	if order, ok := stage.StackGrowthCurve2DEndHalfwayArcShapes_referenceOrder[stackgrowthcurve2dendhalfwayarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackGrowthCurve2DEndHalfwayArcShape was not staged and does not have a reference order", stackgrowthcurve2dendhalfwayarcshape)
		return 0
	}
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackGrowthCurve2DRibbonEndShape_stagedOrder[stackgrowthcurve2dribbonendshape]; ok {
		return order
	}
	if order, ok := stage.StackGrowthCurve2DRibbonEndShapes_referenceOrder[stackgrowthcurve2dribbonendshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackGrowthCurve2DRibbonEndShape was not staged and does not have a reference order", stackgrowthcurve2dribbonendshape)
		return 0
	}
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackGrowthCurve2DRibbonStartShape_stagedOrder[stackgrowthcurve2dribbonstartshape]; ok {
		return order
	}
	if order, ok := stage.StackGrowthCurve2DRibbonStartShapes_referenceOrder[stackgrowthcurve2dribbonstartshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackGrowthCurve2DRibbonStartShape was not staged and does not have a reference order", stackgrowthcurve2dribbonstartshape)
		return 0
	}
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackGrowthCurve2DStartHalfwayArcShape_stagedOrder[stackgrowthcurve2dstarthalfwayarcshape]; ok {
		return order
	}
	if order, ok := stage.StackGrowthCurve2DStartHalfwayArcShapes_referenceOrder[stackgrowthcurve2dstarthalfwayarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackGrowthCurve2DStartHalfwayArcShape was not staged and does not have a reference order", stackgrowthcurve2dstarthalfwayarcshape)
		return 0
	}
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackOfGrowthCurve2D_stagedOrder[stackofgrowthcurve2d]; ok {
		return order
	}
	if order, ok := stage.StackOfGrowthCurve2Ds_referenceOrder[stackofgrowthcurve2d]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackOfGrowthCurve2D was not staged and does not have a reference order", stackofgrowthcurve2d)
		return 0
	}
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackOfGrowthCurve2DRibbon_stagedOrder[stackofgrowthcurve2dribbon]; ok {
		return order
	}
	if order, ok := stage.StackOfGrowthCurve2DRibbons_referenceOrder[stackofgrowthcurve2dribbon]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackOfGrowthCurve2DRibbon was not staged and does not have a reference order", stackofgrowthcurve2dribbon)
		return 0
	}
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackOfRotatedGrowthCurve2D_stagedOrder[stackofrotatedgrowthcurve2d]; ok {
		return order
	}
	if order, ok := stage.StackOfRotatedGrowthCurve2Ds_referenceOrder[stackofrotatedgrowthcurve2d]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackOfRotatedGrowthCurve2D was not staged and does not have a reference order", stackofrotatedgrowthcurve2d)
		return 0
	}
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackOfRotatedGrowthCurve2DRibbon_stagedOrder[stackofrotatedgrowthcurve2dribbon]; ok {
		return order
	}
	if order, ok := stage.StackOfRotatedGrowthCurve2DRibbons_referenceOrder[stackofrotatedgrowthcurve2dribbon]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackOfRotatedGrowthCurve2DRibbon was not staged and does not have a reference order", stackofrotatedgrowthcurve2dribbon)
		return 0
	}
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackRotatedGrowthCurve2DEndArcShape_stagedOrder[stackrotatedgrowthcurve2dendarcshape]; ok {
		return order
	}
	if order, ok := stage.StackRotatedGrowthCurve2DEndArcShapes_referenceOrder[stackrotatedgrowthcurve2dendarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackRotatedGrowthCurve2DEndArcShape was not staged and does not have a reference order", stackrotatedgrowthcurve2dendarcshape)
		return 0
	}
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackRotatedGrowthCurve2DRibbonEndShape_stagedOrder[stackrotatedgrowthcurve2dribbonendshape]; ok {
		return order
	}
	if order, ok := stage.StackRotatedGrowthCurve2DRibbonEndShapes_referenceOrder[stackrotatedgrowthcurve2dribbonendshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackRotatedGrowthCurve2DRibbonEndShape was not staged and does not have a reference order", stackrotatedgrowthcurve2dribbonendshape)
		return 0
	}
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackRotatedGrowthCurve2DRibbonStartShape_stagedOrder[stackrotatedgrowthcurve2dribbonstartshape]; ok {
		return order
	}
	if order, ok := stage.StackRotatedGrowthCurve2DRibbonStartShapes_referenceOrder[stackrotatedgrowthcurve2dribbonstartshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackRotatedGrowthCurve2DRibbonStartShape was not staged and does not have a reference order", stackrotatedgrowthcurve2dribbonstartshape)
		return 0
	}
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackRotatedGrowthCurve2DStartArcShape_stagedOrder[stackrotatedgrowthcurve2dstartarcshape]; ok {
		return order
	}
	if order, ok := stage.StackRotatedGrowthCurve2DStartArcShapes_referenceOrder[stackrotatedgrowthcurve2dstartarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackRotatedGrowthCurve2DStartArcShape was not staged and does not have a reference order", stackrotatedgrowthcurve2dstartarcshape)
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

func (starthalfwayarcshape *StartHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StartHalfwayArcShape_stagedOrder[starthalfwayarcshape]; ok {
		return order
	}
	if order, ok := stage.StartHalfwayArcShapes_referenceOrder[starthalfwayarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StartHalfwayArcShape was not staged and does not have a reference order", starthalfwayarcshape)
		return 0
	}
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StartHalfwayArcShapeGrid_stagedOrder[starthalfwayarcshapegrid]; ok {
		return order
	}
	if order, ok := stage.StartHalfwayArcShapeGrids_referenceOrder[starthalfwayarcshapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type StartHalfwayArcShapeGrid was not staged and does not have a reference order", starthalfwayarcshapegrid)
		return 0
	}
}

func (topendarcshape *TopEndArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopEndArcShape_stagedOrder[topendarcshape]; ok {
		return order
	}
	if order, ok := stage.TopEndArcShapes_referenceOrder[topendarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopEndArcShape was not staged and does not have a reference order", topendarcshape)
		return 0
	}
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopEndArcShapeGrid_stagedOrder[topendarcshapegrid]; ok {
		return order
	}
	if order, ok := stage.TopEndArcShapeGrids_referenceOrder[topendarcshapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopEndArcShapeGrid was not staged and does not have a reference order", topendarcshapegrid)
		return 0
	}
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopEndHalfwayArcShape_stagedOrder[topendhalfwayarcshape]; ok {
		return order
	}
	if order, ok := stage.TopEndHalfwayArcShapes_referenceOrder[topendhalfwayarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopEndHalfwayArcShape was not staged and does not have a reference order", topendhalfwayarcshape)
		return 0
	}
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopEndHalfwayArcShapeGrid_stagedOrder[topendhalfwayarcshapegrid]; ok {
		return order
	}
	if order, ok := stage.TopEndHalfwayArcShapeGrids_referenceOrder[topendhalfwayarcshapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopEndHalfwayArcShapeGrid was not staged and does not have a reference order", topendhalfwayarcshapegrid)
		return 0
	}
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopGrowthCurve2D_stagedOrder[topgrowthcurve2d]; ok {
		return order
	}
	if order, ok := stage.TopGrowthCurve2Ds_referenceOrder[topgrowthcurve2d]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopGrowthCurve2D was not staged and does not have a reference order", topgrowthcurve2d)
		return 0
	}
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopMidArcVectorShape_stagedOrder[topmidarcvectorshape]; ok {
		return order
	}
	if order, ok := stage.TopMidArcVectorShapes_referenceOrder[topmidarcvectorshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopMidArcVectorShape was not staged and does not have a reference order", topmidarcvectorshape)
		return 0
	}
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopMidArcVectorShapeGrid_stagedOrder[topmidarcvectorshapegrid]; ok {
		return order
	}
	if order, ok := stage.TopMidArcVectorShapeGrids_referenceOrder[topmidarcvectorshapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopMidArcVectorShapeGrid was not staged and does not have a reference order", topmidarcvectorshapegrid)
		return 0
	}
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStackGrowthCurve2DEndHalfwayArcShape_stagedOrder[topstackgrowthcurve2dendhalfwayarcshape]; ok {
		return order
	}
	if order, ok := stage.TopStackGrowthCurve2DEndHalfwayArcShapes_referenceOrder[topstackgrowthcurve2dendhalfwayarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStackGrowthCurve2DEndHalfwayArcShape was not staged and does not have a reference order", topstackgrowthcurve2dendhalfwayarcshape)
		return 0
	}
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStackGrowthCurve2DStartHalfwayArcShape_stagedOrder[topstackgrowthcurve2dstarthalfwayarcshape]; ok {
		return order
	}
	if order, ok := stage.TopStackGrowthCurve2DStartHalfwayArcShapes_referenceOrder[topstackgrowthcurve2dstarthalfwayarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStackGrowthCurve2DStartHalfwayArcShape was not staged and does not have a reference order", topstackgrowthcurve2dstarthalfwayarcshape)
		return 0
	}
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStackOfGrowthCurve2D_stagedOrder[topstackofgrowthcurve2d]; ok {
		return order
	}
	if order, ok := stage.TopStackOfGrowthCurve2Ds_referenceOrder[topstackofgrowthcurve2d]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStackOfGrowthCurve2D was not staged and does not have a reference order", topstackofgrowthcurve2d)
		return 0
	}
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStackOfRotatedGrowthCurve2D_stagedOrder[topstackofrotatedgrowthcurve2d]; ok {
		return order
	}
	if order, ok := stage.TopStackOfRotatedGrowthCurve2Ds_referenceOrder[topstackofrotatedgrowthcurve2d]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStackOfRotatedGrowthCurve2D was not staged and does not have a reference order", topstackofrotatedgrowthcurve2d)
		return 0
	}
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStackOfRotatedGrowthCurve2DEndArcShape_stagedOrder[topstackofrotatedgrowthcurve2dendarcshape]; ok {
		return order
	}
	if order, ok := stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_referenceOrder[topstackofrotatedgrowthcurve2dendarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStackOfRotatedGrowthCurve2DEndArcShape was not staged and does not have a reference order", topstackofrotatedgrowthcurve2dendarcshape)
		return 0
	}
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStackOfRotatedGrowthCurve2DStartArcShape_stagedOrder[topstackofrotatedgrowthcurve2dstartarcshape]; ok {
		return order
	}
	if order, ok := stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_referenceOrder[topstackofrotatedgrowthcurve2dstartarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStackOfRotatedGrowthCurve2DStartArcShape was not staged and does not have a reference order", topstackofrotatedgrowthcurve2dstartarcshape)
		return 0
	}
}

func (topstartarcshape *TopStartArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStartArcShape_stagedOrder[topstartarcshape]; ok {
		return order
	}
	if order, ok := stage.TopStartArcShapes_referenceOrder[topstartarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStartArcShape was not staged and does not have a reference order", topstartarcshape)
		return 0
	}
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStartArcShapeGrid_stagedOrder[topstartarcshapegrid]; ok {
		return order
	}
	if order, ok := stage.TopStartArcShapeGrids_referenceOrder[topstartarcshapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStartArcShapeGrid was not staged and does not have a reference order", topstartarcshapegrid)
		return 0
	}
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStartHalfwayArcShape_stagedOrder[topstarthalfwayarcshape]; ok {
		return order
	}
	if order, ok := stage.TopStartHalfwayArcShapes_referenceOrder[topstarthalfwayarcshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStartHalfwayArcShape was not staged and does not have a reference order", topstarthalfwayarcshape)
		return 0
	}
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopStartHalfwayArcShapeGrid_stagedOrder[topstarthalfwayarcshapegrid]; ok {
		return order
	}
	if order, ok := stage.TopStartHalfwayArcShapeGrids_referenceOrder[topstarthalfwayarcshapegrid]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopStartHalfwayArcShapeGrid was not staged and does not have a reference order", topstarthalfwayarcshapegrid)
		return 0
	}
}

func (torusstackshape *TorusStackShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TorusStackShape_stagedOrder[torusstackshape]; ok {
		return order
	}
	if order, ok := stage.TorusStackShapes_referenceOrder[torusstackshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TorusStackShape was not staged and does not have a reference order", torusstackshape)
		return 0
	}
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.VerticalTorusStackShape_stagedOrder[verticaltorusstackshape]; ok {
		return order
	}
	if order, ok := stage.VerticalTorusStackShapes_referenceOrder[verticaltorusstackshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type VerticalTorusStackShape was not staged and does not have a reference order", verticaltorusstackshape)
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

func (endhalfwayarcshape *EndHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", endhalfwayarcshape.GongGetGongstructName(), endhalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (endhalfwayarcshape *EndHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", endhalfwayarcshape.GongGetGongstructName(), endhalfwayarcshape.GongGetOrder(stage))
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", endhalfwayarcshapegrid.GongGetGongstructName(), endhalfwayarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", endhalfwayarcshapegrid.GongGetGongstructName(), endhalfwayarcshapegrid.GongGetOrder(stage))
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

func (growthcurve2d *GrowthCurve2D) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurve2d.GongGetGongstructName(), growthcurve2d.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurve2d *GrowthCurve2D) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurve2d.GongGetGongstructName(), growthcurve2d.GongGetOrder(stage))
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

func (midarcvectorshape *MidArcVectorShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", midarcvectorshape.GongGetGongstructName(), midarcvectorshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (midarcvectorshape *MidArcVectorShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", midarcvectorshape.GongGetGongstructName(), midarcvectorshape.GongGetOrder(stage))
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", midarcvectorshapegrid.GongGetGongstructName(), midarcvectorshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", midarcvectorshapegrid.GongGetGongstructName(), midarcvectorshapegrid.GongGetOrder(stage))
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dribbon.GongGetGongstructName(), partiallygrowthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dribbon.GongGetGongstructName(), partiallygrowthcurve2dribbon.GongGetOrder(stage))
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dribbonendshape.GongGetGongstructName(), partiallygrowthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dribbonendshape.GongGetGongstructName(), partiallygrowthcurve2dribbonendshape.GongGetOrder(stage))
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dribbonstartshape.GongGetGongstructName(), partiallygrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dribbonstartshape.GongGetGongstructName(), partiallygrowthcurve2dribbonstartshape.GongGetOrder(stage))
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

func (rhombusstuff *RhombusStuff) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rhombusstuff.GongGetGongstructName(), rhombusstuff.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rhombusstuff *RhombusStuff) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rhombusstuff.GongGetGongstructName(), rhombusstuff.GongGetOrder(stage))
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

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedbottomtopstartarcshape.GongGetGongstructName(), shiftedbottomtopstartarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedbottomtopstartarcshape.GongGetGongstructName(), shiftedbottomtopstartarcshape.GongGetOrder(stage))
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedbottomtopstartarcshapegrid.GongGetGongstructName(), shiftedbottomtopstartarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedbottomtopstartarcshapegrid.GongGetGongstructName(), shiftedbottomtopstartarcshapegrid.GongGetOrder(stage))
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftstackgrowthcurveendarcshape.GongGetGongstructName(), shiftedleftstackgrowthcurveendarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftstackgrowthcurveendarcshape.GongGetGongstructName(), shiftedleftstackgrowthcurveendarcshape.GongGetOrder(stage))
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftstackgrowthcurvestartarcshape.GongGetGongstructName(), shiftedleftstackgrowthcurvestartarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftstackgrowthcurvestartarcshape.GongGetGongstructName(), shiftedleftstackgrowthcurvestartarcshape.GongGetOrder(stage))
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftstacknormalvector.GongGetGongstructName(), shiftedleftstacknormalvector.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftstacknormalvector.GongGetGongstructName(), shiftedleftstacknormalvector.GongGetOrder(stage))
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftstackofgrowthcurve.GongGetGongstructName(), shiftedleftstackofgrowthcurve.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftstackofgrowthcurve.GongGetGongstructName(), shiftedleftstackofgrowthcurve.GongGetOrder(stage))
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftstackofnormalvector.GongGetGongstructName(), shiftedleftstackofnormalvector.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftstackofnormalvector.GongGetGongstructName(), shiftedleftstackofnormalvector.GongGetOrder(stage))
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurve2dendhalfwayarcshape.GongGetGongstructName(), stackgrowthcurve2dendhalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurve2dendhalfwayarcshape.GongGetGongstructName(), stackgrowthcurve2dendhalfwayarcshape.GongGetOrder(stage))
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurve2dribbonendshape.GongGetGongstructName(), stackgrowthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurve2dribbonendshape.GongGetGongstructName(), stackgrowthcurve2dribbonendshape.GongGetOrder(stage))
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurve2dribbonstartshape.GongGetGongstructName(), stackgrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurve2dribbonstartshape.GongGetGongstructName(), stackgrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurve2dstarthalfwayarcshape.GongGetGongstructName(), stackgrowthcurve2dstarthalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackgrowthcurve2dstarthalfwayarcshape.GongGetGongstructName(), stackgrowthcurve2dstarthalfwayarcshape.GongGetOrder(stage))
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurve2d.GongGetGongstructName(), stackofgrowthcurve2d.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurve2d.GongGetGongstructName(), stackofgrowthcurve2d.GongGetOrder(stage))
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurve2dribbon.GongGetGongstructName(), stackofgrowthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurve2dribbon.GongGetGongstructName(), stackofgrowthcurve2dribbon.GongGetOrder(stage))
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofrotatedgrowthcurve2d.GongGetGongstructName(), stackofrotatedgrowthcurve2d.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofrotatedgrowthcurve2d.GongGetGongstructName(), stackofrotatedgrowthcurve2d.GongGetOrder(stage))
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofrotatedgrowthcurve2dribbon.GongGetGongstructName(), stackofrotatedgrowthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofrotatedgrowthcurve2dribbon.GongGetGongstructName(), stackofrotatedgrowthcurve2dribbon.GongGetOrder(stage))
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackrotatedgrowthcurve2dendarcshape.GongGetGongstructName(), stackrotatedgrowthcurve2dendarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackrotatedgrowthcurve2dendarcshape.GongGetGongstructName(), stackrotatedgrowthcurve2dendarcshape.GongGetOrder(stage))
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackrotatedgrowthcurve2dribbonendshape.GongGetGongstructName(), stackrotatedgrowthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackrotatedgrowthcurve2dribbonendshape.GongGetGongstructName(), stackrotatedgrowthcurve2dribbonendshape.GongGetOrder(stage))
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackrotatedgrowthcurve2dribbonstartshape.GongGetGongstructName(), stackrotatedgrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackrotatedgrowthcurve2dribbonstartshape.GongGetGongstructName(), stackrotatedgrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackrotatedgrowthcurve2dstartarcshape.GongGetGongstructName(), stackrotatedgrowthcurve2dstartarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackrotatedgrowthcurve2dstartarcshape.GongGetGongstructName(), stackrotatedgrowthcurve2dstartarcshape.GongGetOrder(stage))
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

func (starthalfwayarcshape *StartHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", starthalfwayarcshape.GongGetGongstructName(), starthalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (starthalfwayarcshape *StartHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", starthalfwayarcshape.GongGetGongstructName(), starthalfwayarcshape.GongGetOrder(stage))
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", starthalfwayarcshapegrid.GongGetGongstructName(), starthalfwayarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", starthalfwayarcshapegrid.GongGetGongstructName(), starthalfwayarcshapegrid.GongGetOrder(stage))
}

func (topendarcshape *TopEndArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topendarcshape.GongGetGongstructName(), topendarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topendarcshape *TopEndArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topendarcshape.GongGetGongstructName(), topendarcshape.GongGetOrder(stage))
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topendarcshapegrid.GongGetGongstructName(), topendarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topendarcshapegrid *TopEndArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topendarcshapegrid.GongGetGongstructName(), topendarcshapegrid.GongGetOrder(stage))
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topendhalfwayarcshape.GongGetGongstructName(), topendhalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topendhalfwayarcshape.GongGetGongstructName(), topendhalfwayarcshape.GongGetOrder(stage))
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topendhalfwayarcshapegrid.GongGetGongstructName(), topendhalfwayarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topendhalfwayarcshapegrid.GongGetGongstructName(), topendhalfwayarcshapegrid.GongGetOrder(stage))
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topgrowthcurve2d.GongGetGongstructName(), topgrowthcurve2d.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topgrowthcurve2d *TopGrowthCurve2D) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topgrowthcurve2d.GongGetGongstructName(), topgrowthcurve2d.GongGetOrder(stage))
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topmidarcvectorshape.GongGetGongstructName(), topmidarcvectorshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topmidarcvectorshape *TopMidArcVectorShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topmidarcvectorshape.GongGetGongstructName(), topmidarcvectorshape.GongGetOrder(stage))
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topmidarcvectorshapegrid.GongGetGongstructName(), topmidarcvectorshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topmidarcvectorshapegrid.GongGetGongstructName(), topmidarcvectorshapegrid.GongGetOrder(stage))
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackgrowthcurve2dendhalfwayarcshape.GongGetGongstructName(), topstackgrowthcurve2dendhalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackgrowthcurve2dendhalfwayarcshape.GongGetGongstructName(), topstackgrowthcurve2dendhalfwayarcshape.GongGetOrder(stage))
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackgrowthcurve2dstarthalfwayarcshape.GongGetGongstructName(), topstackgrowthcurve2dstarthalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackgrowthcurve2dstarthalfwayarcshape.GongGetGongstructName(), topstackgrowthcurve2dstarthalfwayarcshape.GongGetOrder(stage))
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackofgrowthcurve2d.GongGetGongstructName(), topstackofgrowthcurve2d.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackofgrowthcurve2d.GongGetGongstructName(), topstackofgrowthcurve2d.GongGetOrder(stage))
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackofrotatedgrowthcurve2d.GongGetGongstructName(), topstackofrotatedgrowthcurve2d.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackofrotatedgrowthcurve2d.GongGetGongstructName(), topstackofrotatedgrowthcurve2d.GongGetOrder(stage))
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackofrotatedgrowthcurve2dendarcshape.GongGetGongstructName(), topstackofrotatedgrowthcurve2dendarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackofrotatedgrowthcurve2dendarcshape.GongGetGongstructName(), topstackofrotatedgrowthcurve2dendarcshape.GongGetOrder(stage))
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackofrotatedgrowthcurve2dstartarcshape.GongGetGongstructName(), topstackofrotatedgrowthcurve2dstartarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstackofrotatedgrowthcurve2dstartarcshape.GongGetGongstructName(), topstackofrotatedgrowthcurve2dstartarcshape.GongGetOrder(stage))
}

func (topstartarcshape *TopStartArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstartarcshape.GongGetGongstructName(), topstartarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstartarcshape *TopStartArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstartarcshape.GongGetGongstructName(), topstartarcshape.GongGetOrder(stage))
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstartarcshapegrid.GongGetGongstructName(), topstartarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstartarcshapegrid *TopStartArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstartarcshapegrid.GongGetGongstructName(), topstartarcshapegrid.GongGetOrder(stage))
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstarthalfwayarcshape.GongGetGongstructName(), topstarthalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstarthalfwayarcshape.GongGetGongstructName(), topstarthalfwayarcshape.GongGetOrder(stage))
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstarthalfwayarcshapegrid.GongGetGongstructName(), topstarthalfwayarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topstarthalfwayarcshapegrid.GongGetGongstructName(), topstarthalfwayarcshapegrid.GongGetOrder(stage))
}

func (torusstackshape *TorusStackShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", torusstackshape.GongGetGongstructName(), torusstackshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (torusstackshape *TorusStackShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", torusstackshape.GongGetGongstructName(), torusstackshape.GongGetOrder(stage))
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", verticaltorusstackshape.GongGetGongstructName(), verticaltorusstackshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticaltorusstackshape *VerticalTorusStackShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", verticaltorusstackshape.GongGetGongstructName(), verticaltorusstackshape.GongGetOrder(stage))
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

func (endhalfwayarcshape *EndHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endhalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EndHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(endhalfwayarcshape.Name))
	return
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endhalfwayarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EndHalfwayArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(endhalfwayarcshapegrid.Name))
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

func (growthcurve2d *GrowthCurve2D) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurve2d.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GrowthCurve2D")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(growthcurve2d.Name))
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

func (midarcvectorshape *MidArcVectorShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", midarcvectorshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MidArcVectorShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(midarcvectorshape.Name))
	return
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", midarcvectorshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MidArcVectorShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(midarcvectorshapegrid.Name))
	return
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dribbon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DRibbon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(partiallygrowthcurve2dribbon.Name))
	return
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dribbonendshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DRibbonEndShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(partiallygrowthcurve2dribbonendshape.Name))
	return
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dribbonstartshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DRibbonStartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(partiallygrowthcurve2dribbonstartshape.Name))
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

func (rhombusstuff *RhombusStuff) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rhombusstuff.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RhombusStuff")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(rhombusstuff.Name))
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

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedbottomtopstartarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedBottomTopStartArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(shiftedbottomtopstartarcshape.Name))
	return
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedbottomtopstartarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedBottomTopStartArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(shiftedbottomtopstartarcshapegrid.Name))
	return
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstackgrowthcurveendarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftStackGrowthCurveEndArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(shiftedleftstackgrowthcurveendarcshape.Name))
	return
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstackgrowthcurvestartarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftStackGrowthCurveStartArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(shiftedleftstackgrowthcurvestartarcshape.Name))
	return
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstacknormalvector.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftStackNormalVector")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(shiftedleftstacknormalvector.Name))
	return
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstackofgrowthcurve.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftStackOfGrowthCurve")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(shiftedleftstackofgrowthcurve.Name))
	return
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstackofnormalvector.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftStackOfNormalVector")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(shiftedleftstackofnormalvector.Name))
	return
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurve2dendhalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackGrowthCurve2DEndHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackgrowthcurve2dendhalfwayarcshape.Name))
	return
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurve2dribbonendshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackGrowthCurve2DRibbonEndShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackgrowthcurve2dribbonendshape.Name))
	return
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurve2dribbonstartshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackGrowthCurve2DRibbonStartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackgrowthcurve2dribbonstartshape.Name))
	return
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurve2dstarthalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackGrowthCurve2DStartHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackgrowthcurve2dstarthalfwayarcshape.Name))
	return
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurve2d.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfGrowthCurve2D")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackofgrowthcurve2d.Name))
	return
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurve2dribbon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfGrowthCurve2DRibbon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackofgrowthcurve2dribbon.Name))
	return
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofrotatedgrowthcurve2d.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfRotatedGrowthCurve2D")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackofrotatedgrowthcurve2d.Name))
	return
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofrotatedgrowthcurve2dribbon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfRotatedGrowthCurve2DRibbon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackofrotatedgrowthcurve2dribbon.Name))
	return
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackrotatedgrowthcurve2dendarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackRotatedGrowthCurve2DEndArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackrotatedgrowthcurve2dendarcshape.Name))
	return
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackrotatedgrowthcurve2dribbonendshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackRotatedGrowthCurve2DRibbonEndShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackrotatedgrowthcurve2dribbonendshape.Name))
	return
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackrotatedgrowthcurve2dribbonstartshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackRotatedGrowthCurve2DRibbonStartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackrotatedgrowthcurve2dribbonstartshape.Name))
	return
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackrotatedgrowthcurve2dstartarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackRotatedGrowthCurve2DStartArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stackrotatedgrowthcurve2dstartarcshape.Name))
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

func (starthalfwayarcshape *StartHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", starthalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StartHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(starthalfwayarcshape.Name))
	return
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", starthalfwayarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StartHalfwayArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(starthalfwayarcshapegrid.Name))
	return
}

func (topendarcshape *TopEndArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopEndArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topendarcshape.Name))
	return
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopEndArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topendarcshapegrid.Name))
	return
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendhalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopEndHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topendhalfwayarcshape.Name))
	return
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendhalfwayarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopEndHalfwayArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topendhalfwayarcshapegrid.Name))
	return
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topgrowthcurve2d.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopGrowthCurve2D")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topgrowthcurve2d.Name))
	return
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topmidarcvectorshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopMidArcVectorShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topmidarcvectorshape.Name))
	return
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topmidarcvectorshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopMidArcVectorShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topmidarcvectorshapegrid.Name))
	return
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackgrowthcurve2dendhalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackGrowthCurve2DEndHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstackgrowthcurve2dendhalfwayarcshape.Name))
	return
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackgrowthcurve2dstarthalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackGrowthCurve2DStartHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstackgrowthcurve2dstarthalfwayarcshape.Name))
	return
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofgrowthcurve2d.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackOfGrowthCurve2D")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstackofgrowthcurve2d.Name))
	return
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofrotatedgrowthcurve2d.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackOfRotatedGrowthCurve2D")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstackofrotatedgrowthcurve2d.Name))
	return
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofrotatedgrowthcurve2dendarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackOfRotatedGrowthCurve2DEndArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstackofrotatedgrowthcurve2dendarcshape.Name))
	return
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofrotatedgrowthcurve2dstartarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackOfRotatedGrowthCurve2DStartArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstackofrotatedgrowthcurve2dstartarcshape.Name))
	return
}

func (topstartarcshape *TopStartArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstartarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStartArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstartarcshape.Name))
	return
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstartarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStartArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstartarcshapegrid.Name))
	return
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstarthalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStartHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstarthalfwayarcshape.Name))
	return
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstarthalfwayarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStartHalfwayArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(topstarthalfwayarcshapegrid.Name))
	return
}

func (torusstackshape *TorusStackShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", torusstackshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TorusStackShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(torusstackshape.Name))
	return
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", verticaltorusstackshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "VerticalTorusStackShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(verticaltorusstackshape.Name))
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

func (endhalfwayarcshape *EndHalfwayArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endhalfwayarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endhalfwayarcshapegrid.GongGetReferenceIdentifier(stage))
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

func (growthcurve2d *GrowthCurve2D) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurve2d.GongGetReferenceIdentifier(stage))
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

func (midarcvectorshape *MidArcVectorShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", midarcvectorshape.GongGetReferenceIdentifier(stage))
	return
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", midarcvectorshapegrid.GongGetReferenceIdentifier(stage))
	return
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dribbon.GongGetReferenceIdentifier(stage))
	return
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
	return
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
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

func (rhombusstuff *RhombusStuff) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rhombusstuff.GongGetReferenceIdentifier(stage))
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

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedbottomtopstartarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedbottomtopstartarcshapegrid.GongGetReferenceIdentifier(stage))
	return
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstackgrowthcurveendarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstackgrowthcurvestartarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstacknormalvector.GongGetReferenceIdentifier(stage))
	return
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstackofgrowthcurve.GongGetReferenceIdentifier(stage))
	return
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstackofnormalvector.GongGetReferenceIdentifier(stage))
	return
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurve2dendhalfwayarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
	return
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
	return
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurve2dstarthalfwayarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurve2d.GongGetReferenceIdentifier(stage))
	return
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurve2dribbon.GongGetReferenceIdentifier(stage))
	return
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofrotatedgrowthcurve2d.GongGetReferenceIdentifier(stage))
	return
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofrotatedgrowthcurve2dribbon.GongGetReferenceIdentifier(stage))
	return
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackrotatedgrowthcurve2dendarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackrotatedgrowthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
	return
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackrotatedgrowthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
	return
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackrotatedgrowthcurve2dstartarcshape.GongGetReferenceIdentifier(stage))
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

func (starthalfwayarcshape *StartHalfwayArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", starthalfwayarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", starthalfwayarcshapegrid.GongGetReferenceIdentifier(stage))
	return
}

func (topendarcshape *TopEndArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendarcshapegrid.GongGetReferenceIdentifier(stage))
	return
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendhalfwayarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendhalfwayarcshapegrid.GongGetReferenceIdentifier(stage))
	return
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topgrowthcurve2d.GongGetReferenceIdentifier(stage))
	return
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topmidarcvectorshape.GongGetReferenceIdentifier(stage))
	return
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topmidarcvectorshapegrid.GongGetReferenceIdentifier(stage))
	return
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackgrowthcurve2dendhalfwayarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackgrowthcurve2dstarthalfwayarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofgrowthcurve2d.GongGetReferenceIdentifier(stage))
	return
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofrotatedgrowthcurve2d.GongGetReferenceIdentifier(stage))
	return
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofrotatedgrowthcurve2dendarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofrotatedgrowthcurve2dstartarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (topstartarcshape *TopStartArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstartarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstartarcshapegrid.GongGetReferenceIdentifier(stage))
	return
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstarthalfwayarcshape.GongGetReferenceIdentifier(stage))
	return
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstarthalfwayarcshapegrid.GongGetReferenceIdentifier(stage))
	return
}

func (torusstackshape *TorusStackShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", torusstackshape.GongGetReferenceIdentifier(stage))
	return
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", verticaltorusstackshape.GongGetReferenceIdentifier(stage))
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
