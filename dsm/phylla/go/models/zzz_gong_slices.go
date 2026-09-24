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
	// Compute reverse map for named struct Angle0Shape
	// insertion point per field

	// Compute reverse map for named struct ArcNormalVectorShape
	// insertion point per field

	// Compute reverse map for named struct ArcNormalVectorShapeGrid
	// insertion point per field

	// Compute reverse map for named struct AxesShape
	// insertion point per field

	// Compute reverse map for named struct BaseVectorShape
	// insertion point per field

	// Compute reverse map for named struct BaseVectorShapeGrid
	// insertion point per field

	// Compute reverse map for named struct BottomCurvePlane1Shape
	// insertion point per field

	// Compute reverse map for named struct BottomCurvePlane2Shape
	// insertion point per field

	// Compute reverse map for named struct ChosenP1P2PairShape
	// insertion point per field

	// Compute reverse map for named struct CircleGridShape
	// insertion point per field

	// Compute reverse map for named struct Circumference3DShape
	// insertion point per field

	// Compute reverse map for named struct Clock2DDiagram
	// insertion point per field

	// Compute reverse map for named struct Clock3DDiagram
	// insertion point per field

	// Compute reverse map for named struct ClockTopCurveShape
	// insertion point per field

	// Compute reverse map for named struct CutLine3DShape
	// insertion point per field

	// Compute reverse map for named struct EndArcShape
	// insertion point per field

	// Compute reverse map for named struct EndArcShapeGrid
	// insertion point per field

	// Compute reverse map for named struct EndHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct EndHalfwayArcShapeGrid
	// insertion point per field

	// Compute reverse map for named struct ExplanationTextShape
	// insertion point per field

	// Compute reverse map for named struct Eye3DShape
	// insertion point per field

	// Compute reverse map for named struct EyeCornersSampledPoints3DShape
	// insertion point per field

	// Compute reverse map for named struct EyeSampledPoints3DShape
	// insertion point per field

	// Compute reverse map for named struct EyeSeatBottomCurveShape
	// insertion point per field

	// Compute reverse map for named struct EyeStoolBottomCurveShape
	// insertion point per field

	// Compute reverse map for named struct EyeVolume3DShape
	// insertion point per field

	// Compute reverse map for named struct GridPathShape
	// insertion point per field

	// Compute reverse map for named struct GrowthCurve2D
	// insertion point per field

	// Compute reverse map for named struct GrowthCurve2DRibbon
	// insertion point per field

	// Compute reverse map for named struct GrowthCurve2DRibbonEndShape
	// insertion point per field

	// Compute reverse map for named struct GrowthCurve2DRibbonStartShape
	// insertion point per field

	// Compute reverse map for named struct GrowthCurveRhombusGridShape
	// insertion point per field

	// Compute reverse map for named struct GrowthCurveRhombusShape
	// insertion point per field

	// Compute reverse map for named struct GrowthVectorShape
	// insertion point per field

	// Compute reverse map for named struct InitialRhombusGridShape
	// insertion point per field

	// Compute reverse map for named struct InitialRhombusShape
	// insertion point per field

	// Compute reverse map for named struct Key3DShape
	// insertion point per field

	// Compute reverse map for named struct KeyHole3DShape
	// insertion point per field

	// Compute reverse map for named struct KeyHoleShape
	// insertion point per field

	// Compute reverse map for named struct Leaves3DShape
	// insertion point per field

	// Compute reverse map for named struct Library
	// insertion point per field
	stage.Library_Plants_reverseMap = make(map[*PlantAbstract]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _plantabstract := range library.Plants {
			stage.Library_Plants_reverseMap[_plantabstract] = library
		}
	}
	stage.Library_SubLibraries_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibraries {
			stage.Library_SubLibraries_reverseMap[_library] = library
		}
	}

	// Compute reverse map for named struct MidArcVectorShape
	// insertion point per field

	// Compute reverse map for named struct MidArcVectorShapeGrid
	// insertion point per field

	// Compute reverse map for named struct OriginalPoints3DShape
	// insertion point per field

	// Compute reverse map for named struct ParastichyMCurves3DShape
	// insertion point per field

	// Compute reverse map for named struct ParastichyNCurves3DShape
	// insertion point per field

	// Compute reverse map for named struct PartiallyGrowthCurve2DRibbon
	// insertion point per field

	// Compute reverse map for named struct PartiallyGrowthCurve2DRibbonEndShape
	// insertion point per field

	// Compute reverse map for named struct PartiallyGrowthCurve2DRibbonStartShape
	// insertion point per field

	// Compute reverse map for named struct PartiallyGrowthCurve2DTrajectory
	// insertion point per field

	// Compute reverse map for named struct PartiallyGrowthCurve2DTrajectoryP1CurveShape
	// insertion point per field

	// Compute reverse map for named struct PartiallyGrowthCurve2DTrajectoryP1P2
	// insertion point per field

	// Compute reverse map for named struct PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape
	// insertion point per field

	// Compute reverse map for named struct PartiallyGrowthCurve2DTrajectoryP1PointShape
	// insertion point per field

	// Compute reverse map for named struct PartiallyGrowthCurve2DTrajectoryP2CurveShape
	// insertion point per field

	// Compute reverse map for named struct PartiallyGrowthCurve2DTrajectoryP2PointShape
	// insertion point per field

	// Compute reverse map for named struct PartiallyGrowthCurve2DTrajectoryShape
	// insertion point per field

	// Compute reverse map for named struct PartiallyRotatedSeatBottomCurveShape
	// insertion point per field

	// Compute reverse map for named struct PartiallyRotatedSeatTopCurveShape
	// insertion point per field

	// Compute reverse map for named struct PartiallyRotatedTorusShape
	// insertion point per field

	// Compute reverse map for named struct PerpendicularVector
	// insertion point per field

	// Compute reverse map for named struct PerpendicularVectorGrid
	// insertion point per field

	// Compute reverse map for named struct PerpendicularVectorGridHalfway
	// insertion point per field

	// Compute reverse map for named struct PerpendicularVectorHalfway
	// insertion point per field

	// Compute reverse map for named struct Plant2DDiagram
	// insertion point per field

	// Compute reverse map for named struct Plant3DDiagram
	// insertion point per field

	// Compute reverse map for named struct PlantAbstract
	// insertion point per field
	stage.PlantAbstract_Plant2DDiagrams_reverseMap = make(map[*Plant2DDiagram]*PlantAbstract)
	for plantabstract := range stage.PlantAbstracts {
		_ = plantabstract
		for _, _plant2ddiagram := range plantabstract.Plant2DDiagrams {
			stage.PlantAbstract_Plant2DDiagrams_reverseMap[_plant2ddiagram] = plantabstract
		}
	}
	stage.PlantAbstract_Plant3DDiagrams_reverseMap = make(map[*Plant3DDiagram]*PlantAbstract)
	for plantabstract := range stage.PlantAbstracts {
		_ = plantabstract
		for _, _plant3ddiagram := range plantabstract.Plant3DDiagrams {
			stage.PlantAbstract_Plant3DDiagrams_reverseMap[_plant3ddiagram] = plantabstract
		}
	}
	stage.PlantAbstract_Vase2DDiagrams_reverseMap = make(map[*Vase2DDiagram]*PlantAbstract)
	for plantabstract := range stage.PlantAbstracts {
		_ = plantabstract
		for _, _vase2ddiagram := range plantabstract.Vase2DDiagrams {
			stage.PlantAbstract_Vase2DDiagrams_reverseMap[_vase2ddiagram] = plantabstract
		}
	}
	stage.PlantAbstract_TubeVase3DDiagrams_reverseMap = make(map[*TubeVase3DDiagram]*PlantAbstract)
	for plantabstract := range stage.PlantAbstracts {
		_ = plantabstract
		for _, _tubevase3ddiagram := range plantabstract.TubeVase3DDiagrams {
			stage.PlantAbstract_TubeVase3DDiagrams_reverseMap[_tubevase3ddiagram] = plantabstract
		}
	}
	stage.PlantAbstract_Stool2DDiagrams_reverseMap = make(map[*Stool2DDiagram]*PlantAbstract)
	for plantabstract := range stage.PlantAbstracts {
		_ = plantabstract
		for _, _stool2ddiagram := range plantabstract.Stool2DDiagrams {
			stage.PlantAbstract_Stool2DDiagrams_reverseMap[_stool2ddiagram] = plantabstract
		}
	}
	stage.PlantAbstract_Stool3DDiagrams_reverseMap = make(map[*Stool3DDiagram]*PlantAbstract)
	for plantabstract := range stage.PlantAbstracts {
		_ = plantabstract
		for _, _stool3ddiagram := range plantabstract.Stool3DDiagrams {
			stage.PlantAbstract_Stool3DDiagrams_reverseMap[_stool3ddiagram] = plantabstract
		}
	}
	stage.PlantAbstract_Clock2DDiagrams_reverseMap = make(map[*Clock2DDiagram]*PlantAbstract)
	for plantabstract := range stage.PlantAbstracts {
		_ = plantabstract
		for _, _clock2ddiagram := range plantabstract.Clock2DDiagrams {
			stage.PlantAbstract_Clock2DDiagrams_reverseMap[_clock2ddiagram] = plantabstract
		}
	}
	stage.PlantAbstract_Clock3DDiagrams_reverseMap = make(map[*Clock3DDiagram]*PlantAbstract)
	for plantabstract := range stage.PlantAbstracts {
		_ = plantabstract
		for _, _clock3ddiagram := range plantabstract.Clock3DDiagrams {
			stage.PlantAbstract_Clock3DDiagrams_reverseMap[_clock3ddiagram] = plantabstract
		}
	}

	// Compute reverse map for named struct PlantCircumferenceShape
	// insertion point per field

	// Compute reverse map for named struct PointsAndLines3DShape
	// insertion point per field

	// Compute reverse map for named struct PxShape
	// insertion point per field

	// Compute reverse map for named struct Rendered3DShape
	// insertion point per field

	// Compute reverse map for named struct RhombusShape
	// insertion point per field

	// Compute reverse map for named struct RhombusStuff
	// insertion point per field

	// Compute reverse map for named struct RotatedRhombusGridShape
	// insertion point per field

	// Compute reverse map for named struct RotatedRhombusShape
	// insertion point per field

	// Compute reverse map for named struct RotatedSampledPoints3DShape
	// insertion point per field

	// Compute reverse map for named struct RotatedSeatAndLegs3DShape
	// insertion point per field

	// Compute reverse map for named struct SampledPoints3DShape
	// insertion point per field

	// Compute reverse map for named struct Seat3DShape
	// insertion point per field

	// Compute reverse map for named struct SeatAndLegs3DShape
	// insertion point per field

	// Compute reverse map for named struct SeatBottomCurveShape
	// insertion point per field

	// Compute reverse map for named struct SeatTopCurveShape
	// insertion point per field

	// Compute reverse map for named struct ShiftedBottomTopStartArcShape
	// insertion point per field

	// Compute reverse map for named struct ShiftedBottomTopStartArcShapeGrid
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftGrowthCurve2DRibbon
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftGrowthCurve2DRibbonEndShape
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftGrowthCurve2DRibbonStartShape
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftPartiallyGrowthCurve2DRibbon
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftStackGrowthCurveEndArcShape
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftStackGrowthCurveStartArcShape
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftStackNormalVector
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftStackOfGrowthCurve
	// insertion point per field

	// Compute reverse map for named struct ShiftedLeftStackOfNormalVector
	// insertion point per field

	// Compute reverse map for named struct ShiftedRightGrowthCurve2DRibbon
	// insertion point per field

	// Compute reverse map for named struct ShiftedRightGrowthCurve2DRibbonEndShape
	// insertion point per field

	// Compute reverse map for named struct ShiftedRightGrowthCurve2DRibbonStartShape
	// insertion point per field

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

	// Compute reverse map for named struct StackOfGrowthCurve2DByGrowthVector
	// insertion point per field

	// Compute reverse map for named struct StackOfGrowthCurve2DRibbon
	// insertion point per field

	// Compute reverse map for named struct StackOfPartiallyRotatedTorusShape
	// insertion point per field

	// Compute reverse map for named struct StackOfRotatedGrowthCurve2D
	// insertion point per field

	// Compute reverse map for named struct StackOfRotatedGrowthCurve2DRibbon
	// insertion point per field

	// Compute reverse map for named struct StackOfRotatedVaseTrapezeRingsShape
	// insertion point per field

	// Compute reverse map for named struct StackOfVaseTrapezeRingsShape
	// insertion point per field

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

	// Compute reverse map for named struct StartHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct StartHalfwayArcShapeGrid
	// insertion point per field

	// Compute reverse map for named struct StemCylinder3DShape
	// insertion point per field

	// Compute reverse map for named struct Stool2DDiagram
	// insertion point per field

	// Compute reverse map for named struct Stool3DDiagram
	// insertion point per field

	// Compute reverse map for named struct TiledFloor3DShape
	// insertion point per field

	// Compute reverse map for named struct TopCurvePlane1Shape
	// insertion point per field

	// Compute reverse map for named struct TopCurvePlane2Shape
	// insertion point per field

	// Compute reverse map for named struct TopEndArcShape
	// insertion point per field

	// Compute reverse map for named struct TopEndArcShapeGrid
	// insertion point per field

	// Compute reverse map for named struct TopEndHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct TopEndHalfwayArcShapeGrid
	// insertion point per field

	// Compute reverse map for named struct TopGrowthCurve2D
	// insertion point per field

	// Compute reverse map for named struct TopMidArcVectorShape
	// insertion point per field

	// Compute reverse map for named struct TopMidArcVectorShapeGrid
	// insertion point per field

	// Compute reverse map for named struct TopStackGrowthCurve2DEndHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct TopStackGrowthCurve2DStartHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct TopStackOfGrowthCurve2D
	// insertion point per field

	// Compute reverse map for named struct TopStackOfRotatedGrowthCurve2D
	// insertion point per field

	// Compute reverse map for named struct TopStackOfRotatedGrowthCurve2DEndArcShape
	// insertion point per field

	// Compute reverse map for named struct TopStackOfRotatedGrowthCurve2DStartArcShape
	// insertion point per field

	// Compute reverse map for named struct TopStartArcShape
	// insertion point per field

	// Compute reverse map for named struct TopStartArcShapeGrid
	// insertion point per field

	// Compute reverse map for named struct TopStartHalfwayArcShape
	// insertion point per field

	// Compute reverse map for named struct TopStartHalfwayArcShapeGrid
	// insertion point per field

	// Compute reverse map for named struct Torus3DShape
	// insertion point per field

	// Compute reverse map for named struct TorusEdge3DShape
	// insertion point per field

	// Compute reverse map for named struct TorusStackShape
	// insertion point per field

	// Compute reverse map for named struct TubeVase3DDiagram
	// insertion point per field

	// Compute reverse map for named struct TubeVaseAbstract
	// insertion point per field

	// Compute reverse map for named struct Vase2DDiagram
	// insertion point per field

	// Compute reverse map for named struct VaseTrapezeRingShape
	// insertion point per field

	// Compute reverse map for named struct VerticalTorusStackShape
	// insertion point per field

	// Compute reverse map for named struct VolumeKey3DShape
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Angle0Shapes {
		res = append(res, instance)
	}

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

	for instance := range stage.BottomCurvePlane1Shapes {
		res = append(res, instance)
	}

	for instance := range stage.BottomCurvePlane2Shapes {
		res = append(res, instance)
	}

	for instance := range stage.ChosenP1P2PairShapes {
		res = append(res, instance)
	}

	for instance := range stage.CircleGridShapes {
		res = append(res, instance)
	}

	for instance := range stage.Circumference3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.Clock2DDiagrams {
		res = append(res, instance)
	}

	for instance := range stage.Clock3DDiagrams {
		res = append(res, instance)
	}

	for instance := range stage.ClockTopCurveShapes {
		res = append(res, instance)
	}

	for instance := range stage.CutLine3DShapes {
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

	for instance := range stage.Eye3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.EyeCornersSampledPoints3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.EyeSampledPoints3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.EyeSeatBottomCurveShapes {
		res = append(res, instance)
	}

	for instance := range stage.EyeStoolBottomCurveShapes {
		res = append(res, instance)
	}

	for instance := range stage.EyeVolume3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.GridPathShapes {
		res = append(res, instance)
	}

	for instance := range stage.GrowthCurve2Ds {
		res = append(res, instance)
	}

	for instance := range stage.GrowthCurve2DRibbons {
		res = append(res, instance)
	}

	for instance := range stage.GrowthCurve2DRibbonEndShapes {
		res = append(res, instance)
	}

	for instance := range stage.GrowthCurve2DRibbonStartShapes {
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

	for instance := range stage.Key3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.KeyHole3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.KeyHoleShapes {
		res = append(res, instance)
	}

	for instance := range stage.Leaves3DShapes {
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

	for instance := range stage.OriginalPoints3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.ParastichyMCurves3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.ParastichyNCurves3DShapes {
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

	for instance := range stage.PartiallyGrowthCurve2DTrajectorys {
		res = append(res, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes {
		res = append(res, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1P2s {
		res = append(res, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes {
		res = append(res, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes {
		res = append(res, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes {
		res = append(res, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes {
		res = append(res, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryShapes {
		res = append(res, instance)
	}

	for instance := range stage.PartiallyRotatedSeatBottomCurveShapes {
		res = append(res, instance)
	}

	for instance := range stage.PartiallyRotatedSeatTopCurveShapes {
		res = append(res, instance)
	}

	for instance := range stage.PartiallyRotatedTorusShapes {
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

	for instance := range stage.Plant2DDiagrams {
		res = append(res, instance)
	}

	for instance := range stage.Plant3DDiagrams {
		res = append(res, instance)
	}

	for instance := range stage.PlantAbstracts {
		res = append(res, instance)
	}

	for instance := range stage.PlantCircumferenceShapes {
		res = append(res, instance)
	}

	for instance := range stage.PointsAndLines3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.PxShapes {
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

	for instance := range stage.RotatedSampledPoints3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.RotatedSeatAndLegs3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.SampledPoints3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.Seat3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.SeatAndLegs3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.SeatBottomCurveShapes {
		res = append(res, instance)
	}

	for instance := range stage.SeatTopCurveShapes {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedBottomTopStartArcShapes {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedBottomTopStartArcShapeGrids {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedLeftGrowthCurve2DRibbons {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedLeftGrowthCurve2DRibbonEndShapes {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedLeftGrowthCurve2DRibbonStartShapes {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedLeftPartiallyGrowthCurve2DRibbons {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes {
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

	for instance := range stage.ShiftedRightGrowthCurve2DRibbons {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedRightGrowthCurve2DRibbonEndShapes {
		res = append(res, instance)
	}

	for instance := range stage.ShiftedRightGrowthCurve2DRibbonStartShapes {
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

	for instance := range stage.StackOfGrowthCurve2DByGrowthVectors {
		res = append(res, instance)
	}

	for instance := range stage.StackOfGrowthCurve2DRibbons {
		res = append(res, instance)
	}

	for instance := range stage.StackOfPartiallyRotatedTorusShapes {
		res = append(res, instance)
	}

	for instance := range stage.StackOfRotatedGrowthCurve2Ds {
		res = append(res, instance)
	}

	for instance := range stage.StackOfRotatedGrowthCurve2DRibbons {
		res = append(res, instance)
	}

	for instance := range stage.StackOfRotatedVaseTrapezeRingsShapes {
		res = append(res, instance)
	}

	for instance := range stage.StackOfVaseTrapezeRingsShapes {
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

	for instance := range stage.StemCylinder3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.Stool2DDiagrams {
		res = append(res, instance)
	}

	for instance := range stage.Stool3DDiagrams {
		res = append(res, instance)
	}

	for instance := range stage.TiledFloor3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.TopCurvePlane1Shapes {
		res = append(res, instance)
	}

	for instance := range stage.TopCurvePlane2Shapes {
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

	for instance := range stage.Torus3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.TorusEdge3DShapes {
		res = append(res, instance)
	}

	for instance := range stage.TorusStackShapes {
		res = append(res, instance)
	}

	for instance := range stage.TubeVase3DDiagrams {
		res = append(res, instance)
	}

	for instance := range stage.TubeVaseAbstracts {
		res = append(res, instance)
	}

	for instance := range stage.Vase2DDiagrams {
		res = append(res, instance)
	}

	for instance := range stage.VaseTrapezeRingShapes {
		res = append(res, instance)
	}

	for instance := range stage.VerticalTorusStackShapes {
		res = append(res, instance)
	}

	for instance := range stage.VolumeKey3DShapes {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (angle0shape *Angle0Shape) GongCopy() GongstructIF {
	newInstance := new(Angle0Shape)
	angle0shape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongCopy() GongstructIF {
	newInstance := new(ArcNormalVectorShape)
	arcnormalvectorshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongCopy() GongstructIF {
	newInstance := new(ArcNormalVectorShapeGrid)
	arcnormalvectorshapegrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (axesshape *AxesShape) GongCopy() GongstructIF {
	newInstance := new(AxesShape)
	axesshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (basevectorshape *BaseVectorShape) GongCopy() GongstructIF {
	newInstance := new(BaseVectorShape)
	basevectorshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongCopy() GongstructIF {
	newInstance := new(BaseVectorShapeGrid)
	basevectorshapegrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongCopy() GongstructIF {
	newInstance := new(BottomCurvePlane1Shape)
	bottomcurveplane1shape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongCopy() GongstructIF {
	newInstance := new(BottomCurvePlane2Shape)
	bottomcurveplane2shape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongCopy() GongstructIF {
	newInstance := new(ChosenP1P2PairShape)
	chosenp1p2pairshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (circlegridshape *CircleGridShape) GongCopy() GongstructIF {
	newInstance := new(CircleGridShape)
	circlegridshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (circumference3dshape *Circumference3DShape) GongCopy() GongstructIF {
	newInstance := new(Circumference3DShape)
	circumference3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (clock2ddiagram *Clock2DDiagram) GongCopy() GongstructIF {
	newInstance := new(Clock2DDiagram)
	clock2ddiagram.GongCopyBasicFields(newInstance)
	return newInstance
}

func (clock3ddiagram *Clock3DDiagram) GongCopy() GongstructIF {
	newInstance := new(Clock3DDiagram)
	clock3ddiagram.GongCopyBasicFields(newInstance)
	return newInstance
}

func (clocktopcurveshape *ClockTopCurveShape) GongCopy() GongstructIF {
	newInstance := new(ClockTopCurveShape)
	clocktopcurveshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (cutline3dshape *CutLine3DShape) GongCopy() GongstructIF {
	newInstance := new(CutLine3DShape)
	cutline3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (endarcshape *EndArcShape) GongCopy() GongstructIF {
	newInstance := new(EndArcShape)
	endarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (endarcshapegrid *EndArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(EndArcShapeGrid)
	endarcshapegrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(EndHalfwayArcShape)
	endhalfwayarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(EndHalfwayArcShapeGrid)
	endhalfwayarcshapegrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (explanationtextshape *ExplanationTextShape) GongCopy() GongstructIF {
	newInstance := new(ExplanationTextShape)
	explanationtextshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (eye3dshape *Eye3DShape) GongCopy() GongstructIF {
	newInstance := new(Eye3DShape)
	eye3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongCopy() GongstructIF {
	newInstance := new(EyeCornersSampledPoints3DShape)
	eyecornerssampledpoints3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongCopy() GongstructIF {
	newInstance := new(EyeSampledPoints3DShape)
	eyesampledpoints3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongCopy() GongstructIF {
	newInstance := new(EyeSeatBottomCurveShape)
	eyeseatbottomcurveshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongCopy() GongstructIF {
	newInstance := new(EyeStoolBottomCurveShape)
	eyestoolbottomcurveshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (eyevolume3dshape *EyeVolume3DShape) GongCopy() GongstructIF {
	newInstance := new(EyeVolume3DShape)
	eyevolume3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gridpathshape *GridPathShape) GongCopy() GongstructIF {
	newInstance := new(GridPathShape)
	gridpathshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (growthcurve2d *GrowthCurve2D) GongCopy() GongstructIF {
	newInstance := new(GrowthCurve2D)
	growthcurve2d.GongCopyBasicFields(newInstance)
	return newInstance
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongCopy() GongstructIF {
	newInstance := new(GrowthCurve2DRibbon)
	growthcurve2dribbon.GongCopyBasicFields(newInstance)
	return newInstance
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongCopy() GongstructIF {
	newInstance := new(GrowthCurve2DRibbonEndShape)
	growthcurve2dribbonendshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongCopy() GongstructIF {
	newInstance := new(GrowthCurve2DRibbonStartShape)
	growthcurve2dribbonstartshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongCopy() GongstructIF {
	newInstance := new(GrowthCurveRhombusGridShape)
	growthcurverhombusgridshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongCopy() GongstructIF {
	newInstance := new(GrowthCurveRhombusShape)
	growthcurverhombusshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (growthvectorshape *GrowthVectorShape) GongCopy() GongstructIF {
	newInstance := new(GrowthVectorShape)
	growthvectorshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongCopy() GongstructIF {
	newInstance := new(InitialRhombusGridShape)
	initialrhombusgridshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (initialrhombusshape *InitialRhombusShape) GongCopy() GongstructIF {
	newInstance := new(InitialRhombusShape)
	initialrhombusshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (key3dshape *Key3DShape) GongCopy() GongstructIF {
	newInstance := new(Key3DShape)
	key3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (keyhole3dshape *KeyHole3DShape) GongCopy() GongstructIF {
	newInstance := new(KeyHole3DShape)
	keyhole3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (keyholeshape *KeyHoleShape) GongCopy() GongstructIF {
	newInstance := new(KeyHoleShape)
	keyholeshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (leaves3dshape *Leaves3DShape) GongCopy() GongstructIF {
	newInstance := new(Leaves3DShape)
	leaves3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.GongCopyBasicFields(newInstance)
	return newInstance
}

func (midarcvectorshape *MidArcVectorShape) GongCopy() GongstructIF {
	newInstance := new(MidArcVectorShape)
	midarcvectorshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongCopy() GongstructIF {
	newInstance := new(MidArcVectorShapeGrid)
	midarcvectorshapegrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (originalpoints3dshape *OriginalPoints3DShape) GongCopy() GongstructIF {
	newInstance := new(OriginalPoints3DShape)
	originalpoints3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongCopy() GongstructIF {
	newInstance := new(ParastichyMCurves3DShape)
	parastichymcurves3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongCopy() GongstructIF {
	newInstance := new(ParastichyNCurves3DShape)
	parastichyncurves3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DRibbon)
	partiallygrowthcurve2dribbon.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DRibbonEndShape)
	partiallygrowthcurve2dribbonendshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DRibbonStartShape)
	partiallygrowthcurve2dribbonstartshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DTrajectory)
	partiallygrowthcurve2dtrajectory.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DTrajectoryP1CurveShape)
	partiallygrowthcurve2dtrajectoryp1curveshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DTrajectoryP1P2)
	partiallygrowthcurve2dtrajectoryp1p2.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
	partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DTrajectoryP1PointShape)
	partiallygrowthcurve2dtrajectoryp1pointshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DTrajectoryP2CurveShape)
	partiallygrowthcurve2dtrajectoryp2curveshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DTrajectoryP2PointShape)
	partiallygrowthcurve2dtrajectoryp2pointshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongCopy() GongstructIF {
	newInstance := new(PartiallyGrowthCurve2DTrajectoryShape)
	partiallygrowthcurve2dtrajectoryshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongCopy() GongstructIF {
	newInstance := new(PartiallyRotatedSeatBottomCurveShape)
	partiallyrotatedseatbottomcurveshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongCopy() GongstructIF {
	newInstance := new(PartiallyRotatedSeatTopCurveShape)
	partiallyrotatedseattopcurveshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongCopy() GongstructIF {
	newInstance := new(PartiallyRotatedTorusShape)
	partiallyrotatedtorusshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (perpendicularvector *PerpendicularVector) GongCopy() GongstructIF {
	newInstance := new(PerpendicularVector)
	perpendicularvector.GongCopyBasicFields(newInstance)
	return newInstance
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongCopy() GongstructIF {
	newInstance := new(PerpendicularVectorGrid)
	perpendicularvectorgrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongCopy() GongstructIF {
	newInstance := new(PerpendicularVectorGridHalfway)
	perpendicularvectorgridhalfway.GongCopyBasicFields(newInstance)
	return newInstance
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongCopy() GongstructIF {
	newInstance := new(PerpendicularVectorHalfway)
	perpendicularvectorhalfway.GongCopyBasicFields(newInstance)
	return newInstance
}

func (plant2ddiagram *Plant2DDiagram) GongCopy() GongstructIF {
	newInstance := new(Plant2DDiagram)
	plant2ddiagram.GongCopyBasicFields(newInstance)
	return newInstance
}

func (plant3ddiagram *Plant3DDiagram) GongCopy() GongstructIF {
	newInstance := new(Plant3DDiagram)
	plant3ddiagram.GongCopyBasicFields(newInstance)
	return newInstance
}

func (plantabstract *PlantAbstract) GongCopy() GongstructIF {
	newInstance := new(PlantAbstract)
	plantabstract.GongCopyBasicFields(newInstance)
	return newInstance
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongCopy() GongstructIF {
	newInstance := new(PlantCircumferenceShape)
	plantcircumferenceshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongCopy() GongstructIF {
	newInstance := new(PointsAndLines3DShape)
	pointsandlines3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pxshape *PxShape) GongCopy() GongstructIF {
	newInstance := new(PxShape)
	pxshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rendered3dshape *Rendered3DShape) GongCopy() GongstructIF {
	newInstance := new(Rendered3DShape)
	rendered3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rhombusshape *RhombusShape) GongCopy() GongstructIF {
	newInstance := new(RhombusShape)
	rhombusshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rhombusstuff *RhombusStuff) GongCopy() GongstructIF {
	newInstance := new(RhombusStuff)
	rhombusstuff.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongCopy() GongstructIF {
	newInstance := new(RotatedRhombusGridShape)
	rotatedrhombusgridshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rotatedrhombusshape *RotatedRhombusShape) GongCopy() GongstructIF {
	newInstance := new(RotatedRhombusShape)
	rotatedrhombusshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongCopy() GongstructIF {
	newInstance := new(RotatedSampledPoints3DShape)
	rotatedsampledpoints3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongCopy() GongstructIF {
	newInstance := new(RotatedSeatAndLegs3DShape)
	rotatedseatandlegs3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (sampledpoints3dshape *SampledPoints3DShape) GongCopy() GongstructIF {
	newInstance := new(SampledPoints3DShape)
	sampledpoints3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (seat3dshape *Seat3DShape) GongCopy() GongstructIF {
	newInstance := new(Seat3DShape)
	seat3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongCopy() GongstructIF {
	newInstance := new(SeatAndLegs3DShape)
	seatandlegs3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongCopy() GongstructIF {
	newInstance := new(SeatBottomCurveShape)
	seatbottomcurveshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (seattopcurveshape *SeatTopCurveShape) GongCopy() GongstructIF {
	newInstance := new(SeatTopCurveShape)
	seattopcurveshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongCopy() GongstructIF {
	newInstance := new(ShiftedBottomTopStartArcShape)
	shiftedbottomtopstartarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(ShiftedBottomTopStartArcShapeGrid)
	shiftedbottomtopstartarcshapegrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftGrowthCurve2DRibbon)
	shiftedleftgrowthcurve2dribbon.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftGrowthCurve2DRibbonEndShape)
	shiftedleftgrowthcurve2dribbonendshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftGrowthCurve2DRibbonStartShape)
	shiftedleftgrowthcurve2dribbonstartshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftPartiallyGrowthCurve2DRibbon)
	shiftedleftpartiallygrowthcurve2dribbon.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
	shiftedleftpartiallygrowthcurve2dribbonendshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
	shiftedleftpartiallygrowthcurve2dribbonstartshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftStackGrowthCurveEndArcShape)
	shiftedleftstackgrowthcurveendarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftStackGrowthCurveStartArcShape)
	shiftedleftstackgrowthcurvestartarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftStackNormalVector)
	shiftedleftstacknormalvector.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftStackOfGrowthCurve)
	shiftedleftstackofgrowthcurve.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongCopy() GongstructIF {
	newInstance := new(ShiftedLeftStackOfNormalVector)
	shiftedleftstackofnormalvector.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongCopy() GongstructIF {
	newInstance := new(ShiftedRightGrowthCurve2DRibbon)
	shiftedrightgrowthcurve2dribbon.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongCopy() GongstructIF {
	newInstance := new(ShiftedRightGrowthCurve2DRibbonEndShape)
	shiftedrightgrowthcurve2dribbonendshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongCopy() GongstructIF {
	newInstance := new(ShiftedRightGrowthCurve2DRibbonStartShape)
	shiftedrightgrowthcurve2dribbonstartshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(StackGrowthCurve2DEndHalfwayArcShape)
	stackgrowthcurve2dendhalfwayarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongCopy() GongstructIF {
	newInstance := new(StackGrowthCurve2DRibbonEndShape)
	stackgrowthcurve2dribbonendshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongCopy() GongstructIF {
	newInstance := new(StackGrowthCurve2DRibbonStartShape)
	stackgrowthcurve2dribbonstartshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(StackGrowthCurve2DStartHalfwayArcShape)
	stackgrowthcurve2dstarthalfwayarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongCopy() GongstructIF {
	newInstance := new(StackOfGrowthCurve2D)
	stackofgrowthcurve2d.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongCopy() GongstructIF {
	newInstance := new(StackOfGrowthCurve2DByGrowthVector)
	stackofgrowthcurve2dbygrowthvector.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongCopy() GongstructIF {
	newInstance := new(StackOfGrowthCurve2DRibbon)
	stackofgrowthcurve2dribbon.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongCopy() GongstructIF {
	newInstance := new(StackOfPartiallyRotatedTorusShape)
	stackofpartiallyrotatedtorusshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongCopy() GongstructIF {
	newInstance := new(StackOfRotatedGrowthCurve2D)
	stackofrotatedgrowthcurve2d.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongCopy() GongstructIF {
	newInstance := new(StackOfRotatedGrowthCurve2DRibbon)
	stackofrotatedgrowthcurve2dribbon.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongCopy() GongstructIF {
	newInstance := new(StackOfRotatedVaseTrapezeRingsShape)
	stackofrotatedvasetrapezeringsshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongCopy() GongstructIF {
	newInstance := new(StackOfVaseTrapezeRingsShape)
	stackofvasetrapezeringsshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongCopy() GongstructIF {
	newInstance := new(StackRotatedGrowthCurve2DEndArcShape)
	stackrotatedgrowthcurve2dendarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongCopy() GongstructIF {
	newInstance := new(StackRotatedGrowthCurve2DRibbonEndShape)
	stackrotatedgrowthcurve2dribbonendshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongCopy() GongstructIF {
	newInstance := new(StackRotatedGrowthCurve2DRibbonStartShape)
	stackrotatedgrowthcurve2dribbonstartshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongCopy() GongstructIF {
	newInstance := new(StackRotatedGrowthCurve2DStartArcShape)
	stackrotatedgrowthcurve2dstartarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (startarcshape *StartArcShape) GongCopy() GongstructIF {
	newInstance := new(StartArcShape)
	startarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (startarcshapegrid *StartArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(StartArcShapeGrid)
	startarcshapegrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(StartHalfwayArcShape)
	starthalfwayarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(StartHalfwayArcShapeGrid)
	starthalfwayarcshapegrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stemcylinder3dshape *StemCylinder3DShape) GongCopy() GongstructIF {
	newInstance := new(StemCylinder3DShape)
	stemcylinder3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stool2ddiagram *Stool2DDiagram) GongCopy() GongstructIF {
	newInstance := new(Stool2DDiagram)
	stool2ddiagram.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stool3ddiagram *Stool3DDiagram) GongCopy() GongstructIF {
	newInstance := new(Stool3DDiagram)
	stool3ddiagram.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tiledfloor3dshape *TiledFloor3DShape) GongCopy() GongstructIF {
	newInstance := new(TiledFloor3DShape)
	tiledfloor3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongCopy() GongstructIF {
	newInstance := new(TopCurvePlane1Shape)
	topcurveplane1shape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongCopy() GongstructIF {
	newInstance := new(TopCurvePlane2Shape)
	topcurveplane2shape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topendarcshape *TopEndArcShape) GongCopy() GongstructIF {
	newInstance := new(TopEndArcShape)
	topendarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(TopEndArcShapeGrid)
	topendarcshapegrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(TopEndHalfwayArcShape)
	topendhalfwayarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(TopEndHalfwayArcShapeGrid)
	topendhalfwayarcshapegrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongCopy() GongstructIF {
	newInstance := new(TopGrowthCurve2D)
	topgrowthcurve2d.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongCopy() GongstructIF {
	newInstance := new(TopMidArcVectorShape)
	topmidarcvectorshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongCopy() GongstructIF {
	newInstance := new(TopMidArcVectorShapeGrid)
	topmidarcvectorshapegrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(TopStackGrowthCurve2DEndHalfwayArcShape)
	topstackgrowthcurve2dendhalfwayarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(TopStackGrowthCurve2DStartHalfwayArcShape)
	topstackgrowthcurve2dstarthalfwayarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongCopy() GongstructIF {
	newInstance := new(TopStackOfGrowthCurve2D)
	topstackofgrowthcurve2d.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongCopy() GongstructIF {
	newInstance := new(TopStackOfRotatedGrowthCurve2D)
	topstackofrotatedgrowthcurve2d.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongCopy() GongstructIF {
	newInstance := new(TopStackOfRotatedGrowthCurve2DEndArcShape)
	topstackofrotatedgrowthcurve2dendarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongCopy() GongstructIF {
	newInstance := new(TopStackOfRotatedGrowthCurve2DStartArcShape)
	topstackofrotatedgrowthcurve2dstartarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topstartarcshape *TopStartArcShape) GongCopy() GongstructIF {
	newInstance := new(TopStartArcShape)
	topstartarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(TopStartArcShapeGrid)
	topstartarcshapegrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongCopy() GongstructIF {
	newInstance := new(TopStartHalfwayArcShape)
	topstarthalfwayarcshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongCopy() GongstructIF {
	newInstance := new(TopStartHalfwayArcShapeGrid)
	topstarthalfwayarcshapegrid.GongCopyBasicFields(newInstance)
	return newInstance
}

func (torus3dshape *Torus3DShape) GongCopy() GongstructIF {
	newInstance := new(Torus3DShape)
	torus3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (torusedge3dshape *TorusEdge3DShape) GongCopy() GongstructIF {
	newInstance := new(TorusEdge3DShape)
	torusedge3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (torusstackshape *TorusStackShape) GongCopy() GongstructIF {
	newInstance := new(TorusStackShape)
	torusstackshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongCopy() GongstructIF {
	newInstance := new(TubeVase3DDiagram)
	tubevase3ddiagram.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tubevaseabstract *TubeVaseAbstract) GongCopy() GongstructIF {
	newInstance := new(TubeVaseAbstract)
	tubevaseabstract.GongCopyBasicFields(newInstance)
	return newInstance
}

func (vase2ddiagram *Vase2DDiagram) GongCopy() GongstructIF {
	newInstance := new(Vase2DDiagram)
	vase2ddiagram.GongCopyBasicFields(newInstance)
	return newInstance
}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongCopy() GongstructIF {
	newInstance := new(VaseTrapezeRingShape)
	vasetrapezeringshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongCopy() GongstructIF {
	newInstance := new(VerticalTorusStackShape)
	verticaltorusstackshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (volumekey3dshape *VolumeKey3DShape) GongCopy() GongstructIF {
	newInstance := new(VolumeKey3DShape)
	volumekey3dshape.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (angle0shape *Angle0Shape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(angle0shape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(angle0shape), uint64(stage.GetOrder(angle0shape)))
	return
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(arcnormalvectorshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(arcnormalvectorshape), uint64(stage.GetOrder(arcnormalvectorshape)))
	return
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(arcnormalvectorshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(arcnormalvectorshapegrid), uint64(stage.GetOrder(arcnormalvectorshapegrid)))
	return
}

func (axesshape *AxesShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(axesshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(axesshape), uint64(stage.GetOrder(axesshape)))
	return
}

func (basevectorshape *BaseVectorShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(basevectorshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(basevectorshape), uint64(stage.GetOrder(basevectorshape)))
	return
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(basevectorshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(basevectorshapegrid), uint64(stage.GetOrder(basevectorshapegrid)))
	return
}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bottomcurveplane1shape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(bottomcurveplane1shape), uint64(stage.GetOrder(bottomcurveplane1shape)))
	return
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bottomcurveplane2shape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(bottomcurveplane2shape), uint64(stage.GetOrder(bottomcurveplane2shape)))
	return
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(chosenp1p2pairshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(chosenp1p2pairshape), uint64(stage.GetOrder(chosenp1p2pairshape)))
	return
}

func (circlegridshape *CircleGridShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(circlegridshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(circlegridshape), uint64(stage.GetOrder(circlegridshape)))
	return
}

func (circumference3dshape *Circumference3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(circumference3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(circumference3dshape), uint64(stage.GetOrder(circumference3dshape)))
	return
}

func (clock2ddiagram *Clock2DDiagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(clock2ddiagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(clock2ddiagram), uint64(stage.GetOrder(clock2ddiagram)))
	return
}

func (clock3ddiagram *Clock3DDiagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(clock3ddiagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(clock3ddiagram), uint64(stage.GetOrder(clock3ddiagram)))
	return
}

func (clocktopcurveshape *ClockTopCurveShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(clocktopcurveshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(clocktopcurveshape), uint64(stage.GetOrder(clocktopcurveshape)))
	return
}

func (cutline3dshape *CutLine3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cutline3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(cutline3dshape), uint64(stage.GetOrder(cutline3dshape)))
	return
}

func (endarcshape *EndArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(endarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(endarcshape), uint64(stage.GetOrder(endarcshape)))
	return
}

func (endarcshapegrid *EndArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(endarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(endarcshapegrid), uint64(stage.GetOrder(endarcshapegrid)))
	return
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(endhalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(endhalfwayarcshape), uint64(stage.GetOrder(endhalfwayarcshape)))
	return
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(endhalfwayarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(endhalfwayarcshapegrid), uint64(stage.GetOrder(endhalfwayarcshapegrid)))
	return
}

func (explanationtextshape *ExplanationTextShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(explanationtextshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(explanationtextshape), uint64(stage.GetOrder(explanationtextshape)))
	return
}

func (eye3dshape *Eye3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(eye3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(eye3dshape), uint64(stage.GetOrder(eye3dshape)))
	return
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(eyecornerssampledpoints3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(eyecornerssampledpoints3dshape), uint64(stage.GetOrder(eyecornerssampledpoints3dshape)))
	return
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(eyesampledpoints3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(eyesampledpoints3dshape), uint64(stage.GetOrder(eyesampledpoints3dshape)))
	return
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(eyeseatbottomcurveshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(eyeseatbottomcurveshape), uint64(stage.GetOrder(eyeseatbottomcurveshape)))
	return
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(eyestoolbottomcurveshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(eyestoolbottomcurveshape), uint64(stage.GetOrder(eyestoolbottomcurveshape)))
	return
}

func (eyevolume3dshape *EyeVolume3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(eyevolume3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(eyevolume3dshape), uint64(stage.GetOrder(eyevolume3dshape)))
	return
}

func (gridpathshape *GridPathShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(gridpathshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(gridpathshape), uint64(stage.GetOrder(gridpathshape)))
	return
}

func (growthcurve2d *GrowthCurve2D) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(growthcurve2d).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(growthcurve2d), uint64(stage.GetOrder(growthcurve2d)))
	return
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(growthcurve2dribbon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(growthcurve2dribbon), uint64(stage.GetOrder(growthcurve2dribbon)))
	return
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(growthcurve2dribbonendshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(growthcurve2dribbonendshape), uint64(stage.GetOrder(growthcurve2dribbonendshape)))
	return
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(growthcurve2dribbonstartshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(growthcurve2dribbonstartshape), uint64(stage.GetOrder(growthcurve2dribbonstartshape)))
	return
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(growthcurverhombusgridshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(growthcurverhombusgridshape), uint64(stage.GetOrder(growthcurverhombusgridshape)))
	return
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(growthcurverhombusshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(growthcurverhombusshape), uint64(stage.GetOrder(growthcurverhombusshape)))
	return
}

func (growthvectorshape *GrowthVectorShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(growthvectorshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(growthvectorshape), uint64(stage.GetOrder(growthvectorshape)))
	return
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(initialrhombusgridshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(initialrhombusgridshape), uint64(stage.GetOrder(initialrhombusgridshape)))
	return
}

func (initialrhombusshape *InitialRhombusShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(initialrhombusshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(initialrhombusshape), uint64(stage.GetOrder(initialrhombusshape)))
	return
}

func (key3dshape *Key3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(key3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(key3dshape), uint64(stage.GetOrder(key3dshape)))
	return
}

func (keyhole3dshape *KeyHole3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(keyhole3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(keyhole3dshape), uint64(stage.GetOrder(keyhole3dshape)))
	return
}

func (keyholeshape *KeyHoleShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(keyholeshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(keyholeshape), uint64(stage.GetOrder(keyholeshape)))
	return
}

func (leaves3dshape *Leaves3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(leaves3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(leaves3dshape), uint64(stage.GetOrder(leaves3dshape)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(library), uint64(stage.GetOrder(library)))
	return
}

func (midarcvectorshape *MidArcVectorShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(midarcvectorshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(midarcvectorshape), uint64(stage.GetOrder(midarcvectorshape)))
	return
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(midarcvectorshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(midarcvectorshapegrid), uint64(stage.GetOrder(midarcvectorshapegrid)))
	return
}

func (originalpoints3dshape *OriginalPoints3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(originalpoints3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(originalpoints3dshape), uint64(stage.GetOrder(originalpoints3dshape)))
	return
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parastichymcurves3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(parastichymcurves3dshape), uint64(stage.GetOrder(parastichymcurves3dshape)))
	return
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(parastichyncurves3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(parastichyncurves3dshape), uint64(stage.GetOrder(parastichyncurves3dshape)))
	return
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dribbon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallygrowthcurve2dribbon), uint64(stage.GetOrder(partiallygrowthcurve2dribbon)))
	return
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dribbonendshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallygrowthcurve2dribbonendshape), uint64(stage.GetOrder(partiallygrowthcurve2dribbonendshape)))
	return
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dribbonstartshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallygrowthcurve2dribbonstartshape), uint64(stage.GetOrder(partiallygrowthcurve2dribbonstartshape)))
	return
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dtrajectory).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallygrowthcurve2dtrajectory), uint64(stage.GetOrder(partiallygrowthcurve2dtrajectory)))
	return
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dtrajectoryp1curveshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallygrowthcurve2dtrajectoryp1curveshape), uint64(stage.GetOrder(partiallygrowthcurve2dtrajectoryp1curveshape)))
	return
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dtrajectoryp1p2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallygrowthcurve2dtrajectoryp1p2), uint64(stage.GetOrder(partiallygrowthcurve2dtrajectoryp1p2)))
	return
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dtrajectoryp1p2pairlineshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallygrowthcurve2dtrajectoryp1p2pairlineshape), uint64(stage.GetOrder(partiallygrowthcurve2dtrajectoryp1p2pairlineshape)))
	return
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dtrajectoryp1pointshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallygrowthcurve2dtrajectoryp1pointshape), uint64(stage.GetOrder(partiallygrowthcurve2dtrajectoryp1pointshape)))
	return
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dtrajectoryp2curveshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallygrowthcurve2dtrajectoryp2curveshape), uint64(stage.GetOrder(partiallygrowthcurve2dtrajectoryp2curveshape)))
	return
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dtrajectoryp2pointshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallygrowthcurve2dtrajectoryp2pointshape), uint64(stage.GetOrder(partiallygrowthcurve2dtrajectoryp2pointshape)))
	return
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallygrowthcurve2dtrajectoryshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallygrowthcurve2dtrajectoryshape), uint64(stage.GetOrder(partiallygrowthcurve2dtrajectoryshape)))
	return
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallyrotatedseatbottomcurveshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallyrotatedseatbottomcurveshape), uint64(stage.GetOrder(partiallyrotatedseatbottomcurveshape)))
	return
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallyrotatedseattopcurveshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallyrotatedseattopcurveshape), uint64(stage.GetOrder(partiallyrotatedseattopcurveshape)))
	return
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partiallyrotatedtorusshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partiallyrotatedtorusshape), uint64(stage.GetOrder(partiallyrotatedtorusshape)))
	return
}

func (perpendicularvector *PerpendicularVector) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(perpendicularvector).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(perpendicularvector), uint64(stage.GetOrder(perpendicularvector)))
	return
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(perpendicularvectorgrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(perpendicularvectorgrid), uint64(stage.GetOrder(perpendicularvectorgrid)))
	return
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(perpendicularvectorgridhalfway).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(perpendicularvectorgridhalfway), uint64(stage.GetOrder(perpendicularvectorgridhalfway)))
	return
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(perpendicularvectorhalfway).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(perpendicularvectorhalfway), uint64(stage.GetOrder(perpendicularvectorhalfway)))
	return
}

func (plant2ddiagram *Plant2DDiagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(plant2ddiagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(plant2ddiagram), uint64(stage.GetOrder(plant2ddiagram)))
	return
}

func (plant3ddiagram *Plant3DDiagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(plant3ddiagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(plant3ddiagram), uint64(stage.GetOrder(plant3ddiagram)))
	return
}

func (plantabstract *PlantAbstract) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(plantabstract).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(plantabstract), uint64(stage.GetOrder(plantabstract)))
	return
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(plantcircumferenceshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(plantcircumferenceshape), uint64(stage.GetOrder(plantcircumferenceshape)))
	return
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(pointsandlines3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(pointsandlines3dshape), uint64(stage.GetOrder(pointsandlines3dshape)))
	return
}

func (pxshape *PxShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(pxshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(pxshape), uint64(stage.GetOrder(pxshape)))
	return
}

func (rendered3dshape *Rendered3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rendered3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rendered3dshape), uint64(stage.GetOrder(rendered3dshape)))
	return
}

func (rhombusshape *RhombusShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rhombusshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rhombusshape), uint64(stage.GetOrder(rhombusshape)))
	return
}

func (rhombusstuff *RhombusStuff) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rhombusstuff).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rhombusstuff), uint64(stage.GetOrder(rhombusstuff)))
	return
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rotatedrhombusgridshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rotatedrhombusgridshape), uint64(stage.GetOrder(rotatedrhombusgridshape)))
	return
}

func (rotatedrhombusshape *RotatedRhombusShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rotatedrhombusshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rotatedrhombusshape), uint64(stage.GetOrder(rotatedrhombusshape)))
	return
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rotatedsampledpoints3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rotatedsampledpoints3dshape), uint64(stage.GetOrder(rotatedsampledpoints3dshape)))
	return
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rotatedseatandlegs3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rotatedseatandlegs3dshape), uint64(stage.GetOrder(rotatedseatandlegs3dshape)))
	return
}

func (sampledpoints3dshape *SampledPoints3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(sampledpoints3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(sampledpoints3dshape), uint64(stage.GetOrder(sampledpoints3dshape)))
	return
}

func (seat3dshape *Seat3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(seat3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(seat3dshape), uint64(stage.GetOrder(seat3dshape)))
	return
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(seatandlegs3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(seatandlegs3dshape), uint64(stage.GetOrder(seatandlegs3dshape)))
	return
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(seatbottomcurveshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(seatbottomcurveshape), uint64(stage.GetOrder(seatbottomcurveshape)))
	return
}

func (seattopcurveshape *SeatTopCurveShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(seattopcurveshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(seattopcurveshape), uint64(stage.GetOrder(seattopcurveshape)))
	return
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedbottomtopstartarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedbottomtopstartarcshape), uint64(stage.GetOrder(shiftedbottomtopstartarcshape)))
	return
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedbottomtopstartarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedbottomtopstartarcshapegrid), uint64(stage.GetOrder(shiftedbottomtopstartarcshapegrid)))
	return
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftgrowthcurve2dribbon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedleftgrowthcurve2dribbon), uint64(stage.GetOrder(shiftedleftgrowthcurve2dribbon)))
	return
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftgrowthcurve2dribbonendshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedleftgrowthcurve2dribbonendshape), uint64(stage.GetOrder(shiftedleftgrowthcurve2dribbonendshape)))
	return
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftgrowthcurve2dribbonstartshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedleftgrowthcurve2dribbonstartshape), uint64(stage.GetOrder(shiftedleftgrowthcurve2dribbonstartshape)))
	return
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftpartiallygrowthcurve2dribbon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedleftpartiallygrowthcurve2dribbon), uint64(stage.GetOrder(shiftedleftpartiallygrowthcurve2dribbon)))
	return
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftpartiallygrowthcurve2dribbonendshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedleftpartiallygrowthcurve2dribbonendshape), uint64(stage.GetOrder(shiftedleftpartiallygrowthcurve2dribbonendshape)))
	return
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftpartiallygrowthcurve2dribbonstartshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedleftpartiallygrowthcurve2dribbonstartshape), uint64(stage.GetOrder(shiftedleftpartiallygrowthcurve2dribbonstartshape)))
	return
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftstackgrowthcurveendarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedleftstackgrowthcurveendarcshape), uint64(stage.GetOrder(shiftedleftstackgrowthcurveendarcshape)))
	return
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftstackgrowthcurvestartarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedleftstackgrowthcurvestartarcshape), uint64(stage.GetOrder(shiftedleftstackgrowthcurvestartarcshape)))
	return
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftstacknormalvector).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedleftstacknormalvector), uint64(stage.GetOrder(shiftedleftstacknormalvector)))
	return
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftstackofgrowthcurve).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedleftstackofgrowthcurve), uint64(stage.GetOrder(shiftedleftstackofgrowthcurve)))
	return
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedleftstackofnormalvector).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedleftstackofnormalvector), uint64(stage.GetOrder(shiftedleftstackofnormalvector)))
	return
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedrightgrowthcurve2dribbon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedrightgrowthcurve2dribbon), uint64(stage.GetOrder(shiftedrightgrowthcurve2dribbon)))
	return
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedrightgrowthcurve2dribbonendshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedrightgrowthcurve2dribbonendshape), uint64(stage.GetOrder(shiftedrightgrowthcurve2dribbonendshape)))
	return
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shiftedrightgrowthcurve2dribbonstartshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shiftedrightgrowthcurve2dribbonstartshape), uint64(stage.GetOrder(shiftedrightgrowthcurve2dribbonstartshape)))
	return
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackgrowthcurve2dendhalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackgrowthcurve2dendhalfwayarcshape), uint64(stage.GetOrder(stackgrowthcurve2dendhalfwayarcshape)))
	return
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackgrowthcurve2dribbonendshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackgrowthcurve2dribbonendshape), uint64(stage.GetOrder(stackgrowthcurve2dribbonendshape)))
	return
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackgrowthcurve2dribbonstartshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackgrowthcurve2dribbonstartshape), uint64(stage.GetOrder(stackgrowthcurve2dribbonstartshape)))
	return
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackgrowthcurve2dstarthalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackgrowthcurve2dstarthalfwayarcshape), uint64(stage.GetOrder(stackgrowthcurve2dstarthalfwayarcshape)))
	return
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofgrowthcurve2d).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackofgrowthcurve2d), uint64(stage.GetOrder(stackofgrowthcurve2d)))
	return
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofgrowthcurve2dbygrowthvector).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackofgrowthcurve2dbygrowthvector), uint64(stage.GetOrder(stackofgrowthcurve2dbygrowthvector)))
	return
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofgrowthcurve2dribbon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackofgrowthcurve2dribbon), uint64(stage.GetOrder(stackofgrowthcurve2dribbon)))
	return
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofpartiallyrotatedtorusshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackofpartiallyrotatedtorusshape), uint64(stage.GetOrder(stackofpartiallyrotatedtorusshape)))
	return
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofrotatedgrowthcurve2d).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackofrotatedgrowthcurve2d), uint64(stage.GetOrder(stackofrotatedgrowthcurve2d)))
	return
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofrotatedgrowthcurve2dribbon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackofrotatedgrowthcurve2dribbon), uint64(stage.GetOrder(stackofrotatedgrowthcurve2dribbon)))
	return
}

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofrotatedvasetrapezeringsshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackofrotatedvasetrapezeringsshape), uint64(stage.GetOrder(stackofrotatedvasetrapezeringsshape)))
	return
}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackofvasetrapezeringsshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackofvasetrapezeringsshape), uint64(stage.GetOrder(stackofvasetrapezeringsshape)))
	return
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackrotatedgrowthcurve2dendarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackrotatedgrowthcurve2dendarcshape), uint64(stage.GetOrder(stackrotatedgrowthcurve2dendarcshape)))
	return
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackrotatedgrowthcurve2dribbonendshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackrotatedgrowthcurve2dribbonendshape), uint64(stage.GetOrder(stackrotatedgrowthcurve2dribbonendshape)))
	return
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackrotatedgrowthcurve2dribbonstartshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackrotatedgrowthcurve2dribbonstartshape), uint64(stage.GetOrder(stackrotatedgrowthcurve2dribbonstartshape)))
	return
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stackrotatedgrowthcurve2dstartarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stackrotatedgrowthcurve2dstartarcshape), uint64(stage.GetOrder(stackrotatedgrowthcurve2dstartarcshape)))
	return
}

func (startarcshape *StartArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(startarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(startarcshape), uint64(stage.GetOrder(startarcshape)))
	return
}

func (startarcshapegrid *StartArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(startarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(startarcshapegrid), uint64(stage.GetOrder(startarcshapegrid)))
	return
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(starthalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(starthalfwayarcshape), uint64(stage.GetOrder(starthalfwayarcshape)))
	return
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(starthalfwayarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(starthalfwayarcshapegrid), uint64(stage.GetOrder(starthalfwayarcshapegrid)))
	return
}

func (stemcylinder3dshape *StemCylinder3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stemcylinder3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stemcylinder3dshape), uint64(stage.GetOrder(stemcylinder3dshape)))
	return
}

func (stool2ddiagram *Stool2DDiagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stool2ddiagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stool2ddiagram), uint64(stage.GetOrder(stool2ddiagram)))
	return
}

func (stool3ddiagram *Stool3DDiagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stool3ddiagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stool3ddiagram), uint64(stage.GetOrder(stool3ddiagram)))
	return
}

func (tiledfloor3dshape *TiledFloor3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tiledfloor3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tiledfloor3dshape), uint64(stage.GetOrder(tiledfloor3dshape)))
	return
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topcurveplane1shape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topcurveplane1shape), uint64(stage.GetOrder(topcurveplane1shape)))
	return
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topcurveplane2shape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topcurveplane2shape), uint64(stage.GetOrder(topcurveplane2shape)))
	return
}

func (topendarcshape *TopEndArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topendarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topendarcshape), uint64(stage.GetOrder(topendarcshape)))
	return
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topendarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topendarcshapegrid), uint64(stage.GetOrder(topendarcshapegrid)))
	return
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topendhalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topendhalfwayarcshape), uint64(stage.GetOrder(topendhalfwayarcshape)))
	return
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topendhalfwayarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topendhalfwayarcshapegrid), uint64(stage.GetOrder(topendhalfwayarcshapegrid)))
	return
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topgrowthcurve2d).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topgrowthcurve2d), uint64(stage.GetOrder(topgrowthcurve2d)))
	return
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topmidarcvectorshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topmidarcvectorshape), uint64(stage.GetOrder(topmidarcvectorshape)))
	return
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topmidarcvectorshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topmidarcvectorshapegrid), uint64(stage.GetOrder(topmidarcvectorshapegrid)))
	return
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackgrowthcurve2dendhalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topstackgrowthcurve2dendhalfwayarcshape), uint64(stage.GetOrder(topstackgrowthcurve2dendhalfwayarcshape)))
	return
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackgrowthcurve2dstarthalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topstackgrowthcurve2dstarthalfwayarcshape), uint64(stage.GetOrder(topstackgrowthcurve2dstarthalfwayarcshape)))
	return
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackofgrowthcurve2d).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topstackofgrowthcurve2d), uint64(stage.GetOrder(topstackofgrowthcurve2d)))
	return
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackofrotatedgrowthcurve2d).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topstackofrotatedgrowthcurve2d), uint64(stage.GetOrder(topstackofrotatedgrowthcurve2d)))
	return
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackofrotatedgrowthcurve2dendarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topstackofrotatedgrowthcurve2dendarcshape), uint64(stage.GetOrder(topstackofrotatedgrowthcurve2dendarcshape)))
	return
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstackofrotatedgrowthcurve2dstartarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topstackofrotatedgrowthcurve2dstartarcshape), uint64(stage.GetOrder(topstackofrotatedgrowthcurve2dstartarcshape)))
	return
}

func (topstartarcshape *TopStartArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstartarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topstartarcshape), uint64(stage.GetOrder(topstartarcshape)))
	return
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstartarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topstartarcshapegrid), uint64(stage.GetOrder(topstartarcshapegrid)))
	return
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstarthalfwayarcshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topstarthalfwayarcshape), uint64(stage.GetOrder(topstarthalfwayarcshape)))
	return
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(topstarthalfwayarcshapegrid).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(topstarthalfwayarcshapegrid), uint64(stage.GetOrder(topstarthalfwayarcshapegrid)))
	return
}

func (torus3dshape *Torus3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(torus3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(torus3dshape), uint64(stage.GetOrder(torus3dshape)))
	return
}

func (torusedge3dshape *TorusEdge3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(torusedge3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(torusedge3dshape), uint64(stage.GetOrder(torusedge3dshape)))
	return
}

func (torusstackshape *TorusStackShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(torusstackshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(torusstackshape), uint64(stage.GetOrder(torusstackshape)))
	return
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tubevase3ddiagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tubevase3ddiagram), uint64(stage.GetOrder(tubevase3ddiagram)))
	return
}

func (tubevaseabstract *TubeVaseAbstract) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tubevaseabstract).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tubevaseabstract), uint64(stage.GetOrder(tubevaseabstract)))
	return
}

func (vase2ddiagram *Vase2DDiagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(vase2ddiagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(vase2ddiagram), uint64(stage.GetOrder(vase2ddiagram)))
	return
}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(vasetrapezeringshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(vasetrapezeringshape), uint64(stage.GetOrder(vasetrapezeringshape)))
	return
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(verticaltorusstackshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(verticaltorusstackshape), uint64(stage.GetOrder(verticaltorusstackshape)))
	return
}

func (volumekey3dshape *VolumeKey3DShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(volumekey3dshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(volumekey3dshape), uint64(stage.GetOrder(volumekey3dshape)))
	return
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
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
	computeCommitsForType(
		stage,
		stage.Angle0Shapes,
		stage.Angle0Shape_stagedOrder,
		stage.Angle0Shapes_reference,
		&stage.Angle0Shapes_referenceOrder,
		stage.Angle0Shapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.BottomCurvePlane1Shapes,
		stage.BottomCurvePlane1Shape_stagedOrder,
		stage.BottomCurvePlane1Shapes_reference,
		&stage.BottomCurvePlane1Shapes_referenceOrder,
		stage.BottomCurvePlane1Shapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.BottomCurvePlane2Shapes,
		stage.BottomCurvePlane2Shape_stagedOrder,
		stage.BottomCurvePlane2Shapes_reference,
		&stage.BottomCurvePlane2Shapes_referenceOrder,
		stage.BottomCurvePlane2Shapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Circumference3DShapes,
		stage.Circumference3DShape_stagedOrder,
		stage.Circumference3DShapes_reference,
		&stage.Circumference3DShapes_referenceOrder,
		stage.Circumference3DShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Clock2DDiagrams,
		stage.Clock2DDiagram_stagedOrder,
		stage.Clock2DDiagrams_reference,
		&stage.Clock2DDiagrams_referenceOrder,
		stage.Clock2DDiagrams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Clock3DDiagrams,
		stage.Clock3DDiagram_stagedOrder,
		stage.Clock3DDiagrams_reference,
		&stage.Clock3DDiagrams_referenceOrder,
		stage.Clock3DDiagrams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CutLine3DShapes,
		stage.CutLine3DShape_stagedOrder,
		stage.CutLine3DShapes_reference,
		&stage.CutLine3DShapes_referenceOrder,
		stage.CutLine3DShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Leaves3DShapes,
		stage.Leaves3DShape_stagedOrder,
		stage.Leaves3DShapes_reference,
		&stage.Leaves3DShapes_referenceOrder,
		stage.Leaves3DShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Librarys,
		stage.Library_stagedOrder,
		stage.Librarys_reference,
		&stage.Librarys_referenceOrder,
		stage.Librarys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.OriginalPoints3DShapes,
		stage.OriginalPoints3DShape_stagedOrder,
		stage.OriginalPoints3DShapes_reference,
		&stage.OriginalPoints3DShapes_referenceOrder,
		stage.OriginalPoints3DShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ParastichyMCurves3DShapes,
		stage.ParastichyMCurves3DShape_stagedOrder,
		stage.ParastichyMCurves3DShapes_reference,
		&stage.ParastichyMCurves3DShapes_referenceOrder,
		stage.ParastichyMCurves3DShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ParastichyNCurves3DShapes,
		stage.ParastichyNCurves3DShape_stagedOrder,
		stage.ParastichyNCurves3DShapes_reference,
		&stage.ParastichyNCurves3DShapes_referenceOrder,
		stage.ParastichyNCurves3DShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Plant2DDiagrams,
		stage.Plant2DDiagram_stagedOrder,
		stage.Plant2DDiagrams_reference,
		&stage.Plant2DDiagrams_referenceOrder,
		stage.Plant2DDiagrams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Plant3DDiagrams,
		stage.Plant3DDiagram_stagedOrder,
		stage.Plant3DDiagrams_reference,
		&stage.Plant3DDiagrams_referenceOrder,
		stage.Plant3DDiagrams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PlantAbstracts,
		stage.PlantAbstract_stagedOrder,
		stage.PlantAbstracts_reference,
		&stage.PlantAbstracts_referenceOrder,
		stage.PlantAbstracts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Rendered3DShapes,
		stage.Rendered3DShape_stagedOrder,
		stage.Rendered3DShapes_reference,
		&stage.Rendered3DShapes_referenceOrder,
		stage.Rendered3DShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SampledPoints3DShapes,
		stage.SampledPoints3DShape_stagedOrder,
		stage.SampledPoints3DShapes_reference,
		&stage.SampledPoints3DShapes_referenceOrder,
		stage.SampledPoints3DShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StackOfRotatedVaseTrapezeRingsShapes,
		stage.StackOfRotatedVaseTrapezeRingsShape_stagedOrder,
		stage.StackOfRotatedVaseTrapezeRingsShapes_reference,
		&stage.StackOfRotatedVaseTrapezeRingsShapes_referenceOrder,
		stage.StackOfRotatedVaseTrapezeRingsShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StackOfVaseTrapezeRingsShapes,
		stage.StackOfVaseTrapezeRingsShape_stagedOrder,
		stage.StackOfVaseTrapezeRingsShapes_reference,
		&stage.StackOfVaseTrapezeRingsShapes_referenceOrder,
		stage.StackOfVaseTrapezeRingsShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StemCylinder3DShapes,
		stage.StemCylinder3DShape_stagedOrder,
		stage.StemCylinder3DShapes_reference,
		&stage.StemCylinder3DShapes_referenceOrder,
		stage.StemCylinder3DShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Stool2DDiagrams,
		stage.Stool2DDiagram_stagedOrder,
		stage.Stool2DDiagrams_reference,
		&stage.Stool2DDiagrams_referenceOrder,
		stage.Stool2DDiagrams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Stool3DDiagrams,
		stage.Stool3DDiagram_stagedOrder,
		stage.Stool3DDiagrams_reference,
		&stage.Stool3DDiagrams_referenceOrder,
		stage.Stool3DDiagrams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TopCurvePlane1Shapes,
		stage.TopCurvePlane1Shape_stagedOrder,
		stage.TopCurvePlane1Shapes_reference,
		&stage.TopCurvePlane1Shapes_referenceOrder,
		stage.TopCurvePlane1Shapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TopCurvePlane2Shapes,
		stage.TopCurvePlane2Shape_stagedOrder,
		stage.TopCurvePlane2Shapes_reference,
		&stage.TopCurvePlane2Shapes_referenceOrder,
		stage.TopCurvePlane2Shapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TubeVase3DDiagrams,
		stage.TubeVase3DDiagram_stagedOrder,
		stage.TubeVase3DDiagrams_reference,
		&stage.TubeVase3DDiagrams_referenceOrder,
		stage.TubeVase3DDiagrams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TubeVaseAbstracts,
		stage.TubeVaseAbstract_stagedOrder,
		stage.TubeVaseAbstracts_reference,
		&stage.TubeVaseAbstracts_referenceOrder,
		stage.TubeVaseAbstracts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Vase2DDiagrams,
		stage.Vase2DDiagram_stagedOrder,
		stage.Vase2DDiagrams_reference,
		&stage.Vase2DDiagrams_referenceOrder,
		stage.Vase2DDiagrams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.VaseTrapezeRingShapes,
		stage.VaseTrapezeRingShape_stagedOrder,
		stage.VaseTrapezeRingShapes_reference,
		&stage.VaseTrapezeRingShapes_referenceOrder,
		stage.VaseTrapezeRingShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

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
	stage.Angle0Shapes_reference = make(map[*Angle0Shape]*Angle0Shape)
	stage.Angle0Shapes_referenceOrder = make(map[*Angle0Shape]uint) // diff Unstage needs the reference order
	stage.Angle0Shapes_instance = make(map[*Angle0Shape]*Angle0Shape)
	for instance := range stage.Angle0Shapes {
		_copy := instance.GongCopy().(*Angle0Shape)
		stage.Angle0Shapes_reference[instance] = _copy
		stage.Angle0Shapes_instance[_copy] = instance
		stage.Angle0Shapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

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

	stage.BottomCurvePlane1Shapes_reference = make(map[*BottomCurvePlane1Shape]*BottomCurvePlane1Shape)
	stage.BottomCurvePlane1Shapes_referenceOrder = make(map[*BottomCurvePlane1Shape]uint) // diff Unstage needs the reference order
	stage.BottomCurvePlane1Shapes_instance = make(map[*BottomCurvePlane1Shape]*BottomCurvePlane1Shape)
	for instance := range stage.BottomCurvePlane1Shapes {
		_copy := instance.GongCopy().(*BottomCurvePlane1Shape)
		stage.BottomCurvePlane1Shapes_reference[instance] = _copy
		stage.BottomCurvePlane1Shapes_instance[_copy] = instance
		stage.BottomCurvePlane1Shapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BottomCurvePlane2Shapes_reference = make(map[*BottomCurvePlane2Shape]*BottomCurvePlane2Shape)
	stage.BottomCurvePlane2Shapes_referenceOrder = make(map[*BottomCurvePlane2Shape]uint) // diff Unstage needs the reference order
	stage.BottomCurvePlane2Shapes_instance = make(map[*BottomCurvePlane2Shape]*BottomCurvePlane2Shape)
	for instance := range stage.BottomCurvePlane2Shapes {
		_copy := instance.GongCopy().(*BottomCurvePlane2Shape)
		stage.BottomCurvePlane2Shapes_reference[instance] = _copy
		stage.BottomCurvePlane2Shapes_instance[_copy] = instance
		stage.BottomCurvePlane2Shapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ChosenP1P2PairShapes_reference = make(map[*ChosenP1P2PairShape]*ChosenP1P2PairShape)
	stage.ChosenP1P2PairShapes_referenceOrder = make(map[*ChosenP1P2PairShape]uint) // diff Unstage needs the reference order
	stage.ChosenP1P2PairShapes_instance = make(map[*ChosenP1P2PairShape]*ChosenP1P2PairShape)
	for instance := range stage.ChosenP1P2PairShapes {
		_copy := instance.GongCopy().(*ChosenP1P2PairShape)
		stage.ChosenP1P2PairShapes_reference[instance] = _copy
		stage.ChosenP1P2PairShapes_instance[_copy] = instance
		stage.ChosenP1P2PairShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.Circumference3DShapes_reference = make(map[*Circumference3DShape]*Circumference3DShape)
	stage.Circumference3DShapes_referenceOrder = make(map[*Circumference3DShape]uint) // diff Unstage needs the reference order
	stage.Circumference3DShapes_instance = make(map[*Circumference3DShape]*Circumference3DShape)
	for instance := range stage.Circumference3DShapes {
		_copy := instance.GongCopy().(*Circumference3DShape)
		stage.Circumference3DShapes_reference[instance] = _copy
		stage.Circumference3DShapes_instance[_copy] = instance
		stage.Circumference3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Clock2DDiagrams_reference = make(map[*Clock2DDiagram]*Clock2DDiagram)
	stage.Clock2DDiagrams_referenceOrder = make(map[*Clock2DDiagram]uint) // diff Unstage needs the reference order
	stage.Clock2DDiagrams_instance = make(map[*Clock2DDiagram]*Clock2DDiagram)
	for instance := range stage.Clock2DDiagrams {
		_copy := instance.GongCopy().(*Clock2DDiagram)
		stage.Clock2DDiagrams_reference[instance] = _copy
		stage.Clock2DDiagrams_instance[_copy] = instance
		stage.Clock2DDiagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Clock3DDiagrams_reference = make(map[*Clock3DDiagram]*Clock3DDiagram)
	stage.Clock3DDiagrams_referenceOrder = make(map[*Clock3DDiagram]uint) // diff Unstage needs the reference order
	stage.Clock3DDiagrams_instance = make(map[*Clock3DDiagram]*Clock3DDiagram)
	for instance := range stage.Clock3DDiagrams {
		_copy := instance.GongCopy().(*Clock3DDiagram)
		stage.Clock3DDiagrams_reference[instance] = _copy
		stage.Clock3DDiagrams_instance[_copy] = instance
		stage.Clock3DDiagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ClockTopCurveShapes_reference = make(map[*ClockTopCurveShape]*ClockTopCurveShape)
	stage.ClockTopCurveShapes_referenceOrder = make(map[*ClockTopCurveShape]uint) // diff Unstage needs the reference order
	stage.ClockTopCurveShapes_instance = make(map[*ClockTopCurveShape]*ClockTopCurveShape)
	for instance := range stage.ClockTopCurveShapes {
		_copy := instance.GongCopy().(*ClockTopCurveShape)
		stage.ClockTopCurveShapes_reference[instance] = _copy
		stage.ClockTopCurveShapes_instance[_copy] = instance
		stage.ClockTopCurveShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.CutLine3DShapes_reference = make(map[*CutLine3DShape]*CutLine3DShape)
	stage.CutLine3DShapes_referenceOrder = make(map[*CutLine3DShape]uint) // diff Unstage needs the reference order
	stage.CutLine3DShapes_instance = make(map[*CutLine3DShape]*CutLine3DShape)
	for instance := range stage.CutLine3DShapes {
		_copy := instance.GongCopy().(*CutLine3DShape)
		stage.CutLine3DShapes_reference[instance] = _copy
		stage.CutLine3DShapes_instance[_copy] = instance
		stage.CutLine3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.Eye3DShapes_reference = make(map[*Eye3DShape]*Eye3DShape)
	stage.Eye3DShapes_referenceOrder = make(map[*Eye3DShape]uint) // diff Unstage needs the reference order
	stage.Eye3DShapes_instance = make(map[*Eye3DShape]*Eye3DShape)
	for instance := range stage.Eye3DShapes {
		_copy := instance.GongCopy().(*Eye3DShape)
		stage.Eye3DShapes_reference[instance] = _copy
		stage.Eye3DShapes_instance[_copy] = instance
		stage.Eye3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EyeCornersSampledPoints3DShapes_reference = make(map[*EyeCornersSampledPoints3DShape]*EyeCornersSampledPoints3DShape)
	stage.EyeCornersSampledPoints3DShapes_referenceOrder = make(map[*EyeCornersSampledPoints3DShape]uint) // diff Unstage needs the reference order
	stage.EyeCornersSampledPoints3DShapes_instance = make(map[*EyeCornersSampledPoints3DShape]*EyeCornersSampledPoints3DShape)
	for instance := range stage.EyeCornersSampledPoints3DShapes {
		_copy := instance.GongCopy().(*EyeCornersSampledPoints3DShape)
		stage.EyeCornersSampledPoints3DShapes_reference[instance] = _copy
		stage.EyeCornersSampledPoints3DShapes_instance[_copy] = instance
		stage.EyeCornersSampledPoints3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EyeSampledPoints3DShapes_reference = make(map[*EyeSampledPoints3DShape]*EyeSampledPoints3DShape)
	stage.EyeSampledPoints3DShapes_referenceOrder = make(map[*EyeSampledPoints3DShape]uint) // diff Unstage needs the reference order
	stage.EyeSampledPoints3DShapes_instance = make(map[*EyeSampledPoints3DShape]*EyeSampledPoints3DShape)
	for instance := range stage.EyeSampledPoints3DShapes {
		_copy := instance.GongCopy().(*EyeSampledPoints3DShape)
		stage.EyeSampledPoints3DShapes_reference[instance] = _copy
		stage.EyeSampledPoints3DShapes_instance[_copy] = instance
		stage.EyeSampledPoints3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EyeSeatBottomCurveShapes_reference = make(map[*EyeSeatBottomCurveShape]*EyeSeatBottomCurveShape)
	stage.EyeSeatBottomCurveShapes_referenceOrder = make(map[*EyeSeatBottomCurveShape]uint) // diff Unstage needs the reference order
	stage.EyeSeatBottomCurveShapes_instance = make(map[*EyeSeatBottomCurveShape]*EyeSeatBottomCurveShape)
	for instance := range stage.EyeSeatBottomCurveShapes {
		_copy := instance.GongCopy().(*EyeSeatBottomCurveShape)
		stage.EyeSeatBottomCurveShapes_reference[instance] = _copy
		stage.EyeSeatBottomCurveShapes_instance[_copy] = instance
		stage.EyeSeatBottomCurveShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EyeStoolBottomCurveShapes_reference = make(map[*EyeStoolBottomCurveShape]*EyeStoolBottomCurveShape)
	stage.EyeStoolBottomCurveShapes_referenceOrder = make(map[*EyeStoolBottomCurveShape]uint) // diff Unstage needs the reference order
	stage.EyeStoolBottomCurveShapes_instance = make(map[*EyeStoolBottomCurveShape]*EyeStoolBottomCurveShape)
	for instance := range stage.EyeStoolBottomCurveShapes {
		_copy := instance.GongCopy().(*EyeStoolBottomCurveShape)
		stage.EyeStoolBottomCurveShapes_reference[instance] = _copy
		stage.EyeStoolBottomCurveShapes_instance[_copy] = instance
		stage.EyeStoolBottomCurveShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EyeVolume3DShapes_reference = make(map[*EyeVolume3DShape]*EyeVolume3DShape)
	stage.EyeVolume3DShapes_referenceOrder = make(map[*EyeVolume3DShape]uint) // diff Unstage needs the reference order
	stage.EyeVolume3DShapes_instance = make(map[*EyeVolume3DShape]*EyeVolume3DShape)
	for instance := range stage.EyeVolume3DShapes {
		_copy := instance.GongCopy().(*EyeVolume3DShape)
		stage.EyeVolume3DShapes_reference[instance] = _copy
		stage.EyeVolume3DShapes_instance[_copy] = instance
		stage.EyeVolume3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.GrowthCurve2DRibbons_reference = make(map[*GrowthCurve2DRibbon]*GrowthCurve2DRibbon)
	stage.GrowthCurve2DRibbons_referenceOrder = make(map[*GrowthCurve2DRibbon]uint) // diff Unstage needs the reference order
	stage.GrowthCurve2DRibbons_instance = make(map[*GrowthCurve2DRibbon]*GrowthCurve2DRibbon)
	for instance := range stage.GrowthCurve2DRibbons {
		_copy := instance.GongCopy().(*GrowthCurve2DRibbon)
		stage.GrowthCurve2DRibbons_reference[instance] = _copy
		stage.GrowthCurve2DRibbons_instance[_copy] = instance
		stage.GrowthCurve2DRibbons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GrowthCurve2DRibbonEndShapes_reference = make(map[*GrowthCurve2DRibbonEndShape]*GrowthCurve2DRibbonEndShape)
	stage.GrowthCurve2DRibbonEndShapes_referenceOrder = make(map[*GrowthCurve2DRibbonEndShape]uint) // diff Unstage needs the reference order
	stage.GrowthCurve2DRibbonEndShapes_instance = make(map[*GrowthCurve2DRibbonEndShape]*GrowthCurve2DRibbonEndShape)
	for instance := range stage.GrowthCurve2DRibbonEndShapes {
		_copy := instance.GongCopy().(*GrowthCurve2DRibbonEndShape)
		stage.GrowthCurve2DRibbonEndShapes_reference[instance] = _copy
		stage.GrowthCurve2DRibbonEndShapes_instance[_copy] = instance
		stage.GrowthCurve2DRibbonEndShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GrowthCurve2DRibbonStartShapes_reference = make(map[*GrowthCurve2DRibbonStartShape]*GrowthCurve2DRibbonStartShape)
	stage.GrowthCurve2DRibbonStartShapes_referenceOrder = make(map[*GrowthCurve2DRibbonStartShape]uint) // diff Unstage needs the reference order
	stage.GrowthCurve2DRibbonStartShapes_instance = make(map[*GrowthCurve2DRibbonStartShape]*GrowthCurve2DRibbonStartShape)
	for instance := range stage.GrowthCurve2DRibbonStartShapes {
		_copy := instance.GongCopy().(*GrowthCurve2DRibbonStartShape)
		stage.GrowthCurve2DRibbonStartShapes_reference[instance] = _copy
		stage.GrowthCurve2DRibbonStartShapes_instance[_copy] = instance
		stage.GrowthCurve2DRibbonStartShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.Key3DShapes_reference = make(map[*Key3DShape]*Key3DShape)
	stage.Key3DShapes_referenceOrder = make(map[*Key3DShape]uint) // diff Unstage needs the reference order
	stage.Key3DShapes_instance = make(map[*Key3DShape]*Key3DShape)
	for instance := range stage.Key3DShapes {
		_copy := instance.GongCopy().(*Key3DShape)
		stage.Key3DShapes_reference[instance] = _copy
		stage.Key3DShapes_instance[_copy] = instance
		stage.Key3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.KeyHole3DShapes_reference = make(map[*KeyHole3DShape]*KeyHole3DShape)
	stage.KeyHole3DShapes_referenceOrder = make(map[*KeyHole3DShape]uint) // diff Unstage needs the reference order
	stage.KeyHole3DShapes_instance = make(map[*KeyHole3DShape]*KeyHole3DShape)
	for instance := range stage.KeyHole3DShapes {
		_copy := instance.GongCopy().(*KeyHole3DShape)
		stage.KeyHole3DShapes_reference[instance] = _copy
		stage.KeyHole3DShapes_instance[_copy] = instance
		stage.KeyHole3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.KeyHoleShapes_reference = make(map[*KeyHoleShape]*KeyHoleShape)
	stage.KeyHoleShapes_referenceOrder = make(map[*KeyHoleShape]uint) // diff Unstage needs the reference order
	stage.KeyHoleShapes_instance = make(map[*KeyHoleShape]*KeyHoleShape)
	for instance := range stage.KeyHoleShapes {
		_copy := instance.GongCopy().(*KeyHoleShape)
		stage.KeyHoleShapes_reference[instance] = _copy
		stage.KeyHoleShapes_instance[_copy] = instance
		stage.KeyHoleShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Leaves3DShapes_reference = make(map[*Leaves3DShape]*Leaves3DShape)
	stage.Leaves3DShapes_referenceOrder = make(map[*Leaves3DShape]uint) // diff Unstage needs the reference order
	stage.Leaves3DShapes_instance = make(map[*Leaves3DShape]*Leaves3DShape)
	for instance := range stage.Leaves3DShapes {
		_copy := instance.GongCopy().(*Leaves3DShape)
		stage.Leaves3DShapes_reference[instance] = _copy
		stage.Leaves3DShapes_instance[_copy] = instance
		stage.Leaves3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.OriginalPoints3DShapes_reference = make(map[*OriginalPoints3DShape]*OriginalPoints3DShape)
	stage.OriginalPoints3DShapes_referenceOrder = make(map[*OriginalPoints3DShape]uint) // diff Unstage needs the reference order
	stage.OriginalPoints3DShapes_instance = make(map[*OriginalPoints3DShape]*OriginalPoints3DShape)
	for instance := range stage.OriginalPoints3DShapes {
		_copy := instance.GongCopy().(*OriginalPoints3DShape)
		stage.OriginalPoints3DShapes_reference[instance] = _copy
		stage.OriginalPoints3DShapes_instance[_copy] = instance
		stage.OriginalPoints3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ParastichyMCurves3DShapes_reference = make(map[*ParastichyMCurves3DShape]*ParastichyMCurves3DShape)
	stage.ParastichyMCurves3DShapes_referenceOrder = make(map[*ParastichyMCurves3DShape]uint) // diff Unstage needs the reference order
	stage.ParastichyMCurves3DShapes_instance = make(map[*ParastichyMCurves3DShape]*ParastichyMCurves3DShape)
	for instance := range stage.ParastichyMCurves3DShapes {
		_copy := instance.GongCopy().(*ParastichyMCurves3DShape)
		stage.ParastichyMCurves3DShapes_reference[instance] = _copy
		stage.ParastichyMCurves3DShapes_instance[_copy] = instance
		stage.ParastichyMCurves3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ParastichyNCurves3DShapes_reference = make(map[*ParastichyNCurves3DShape]*ParastichyNCurves3DShape)
	stage.ParastichyNCurves3DShapes_referenceOrder = make(map[*ParastichyNCurves3DShape]uint) // diff Unstage needs the reference order
	stage.ParastichyNCurves3DShapes_instance = make(map[*ParastichyNCurves3DShape]*ParastichyNCurves3DShape)
	for instance := range stage.ParastichyNCurves3DShapes {
		_copy := instance.GongCopy().(*ParastichyNCurves3DShape)
		stage.ParastichyNCurves3DShapes_reference[instance] = _copy
		stage.ParastichyNCurves3DShapes_instance[_copy] = instance
		stage.ParastichyNCurves3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.PartiallyGrowthCurve2DTrajectorys_reference = make(map[*PartiallyGrowthCurve2DTrajectory]*PartiallyGrowthCurve2DTrajectory)
	stage.PartiallyGrowthCurve2DTrajectorys_referenceOrder = make(map[*PartiallyGrowthCurve2DTrajectory]uint) // diff Unstage needs the reference order
	stage.PartiallyGrowthCurve2DTrajectorys_instance = make(map[*PartiallyGrowthCurve2DTrajectory]*PartiallyGrowthCurve2DTrajectory)
	for instance := range stage.PartiallyGrowthCurve2DTrajectorys {
		_copy := instance.GongCopy().(*PartiallyGrowthCurve2DTrajectory)
		stage.PartiallyGrowthCurve2DTrajectorys_reference[instance] = _copy
		stage.PartiallyGrowthCurve2DTrajectorys_instance[_copy] = instance
		stage.PartiallyGrowthCurve2DTrajectorys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes_reference = make(map[*PartiallyGrowthCurve2DTrajectoryP1CurveShape]*PartiallyGrowthCurve2DTrajectoryP1CurveShape)
	stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes_referenceOrder = make(map[*PartiallyGrowthCurve2DTrajectoryP1CurveShape]uint) // diff Unstage needs the reference order
	stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes_instance = make(map[*PartiallyGrowthCurve2DTrajectoryP1CurveShape]*PartiallyGrowthCurve2DTrajectoryP1CurveShape)
	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes {
		_copy := instance.GongCopy().(*PartiallyGrowthCurve2DTrajectoryP1CurveShape)
		stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes_reference[instance] = _copy
		stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes_instance[_copy] = instance
		stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartiallyGrowthCurve2DTrajectoryP1P2s_reference = make(map[*PartiallyGrowthCurve2DTrajectoryP1P2]*PartiallyGrowthCurve2DTrajectoryP1P2)
	stage.PartiallyGrowthCurve2DTrajectoryP1P2s_referenceOrder = make(map[*PartiallyGrowthCurve2DTrajectoryP1P2]uint) // diff Unstage needs the reference order
	stage.PartiallyGrowthCurve2DTrajectoryP1P2s_instance = make(map[*PartiallyGrowthCurve2DTrajectoryP1P2]*PartiallyGrowthCurve2DTrajectoryP1P2)
	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1P2s {
		_copy := instance.GongCopy().(*PartiallyGrowthCurve2DTrajectoryP1P2)
		stage.PartiallyGrowthCurve2DTrajectoryP1P2s_reference[instance] = _copy
		stage.PartiallyGrowthCurve2DTrajectoryP1P2s_instance[_copy] = instance
		stage.PartiallyGrowthCurve2DTrajectoryP1P2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes_reference = make(map[*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape]*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
	stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes_referenceOrder = make(map[*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape]uint) // diff Unstage needs the reference order
	stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes_instance = make(map[*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape]*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes {
		_copy := instance.GongCopy().(*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
		stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes_reference[instance] = _copy
		stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes_instance[_copy] = instance
		stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes_reference = make(map[*PartiallyGrowthCurve2DTrajectoryP1PointShape]*PartiallyGrowthCurve2DTrajectoryP1PointShape)
	stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes_referenceOrder = make(map[*PartiallyGrowthCurve2DTrajectoryP1PointShape]uint) // diff Unstage needs the reference order
	stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes_instance = make(map[*PartiallyGrowthCurve2DTrajectoryP1PointShape]*PartiallyGrowthCurve2DTrajectoryP1PointShape)
	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes {
		_copy := instance.GongCopy().(*PartiallyGrowthCurve2DTrajectoryP1PointShape)
		stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes_reference[instance] = _copy
		stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes_instance[_copy] = instance
		stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes_reference = make(map[*PartiallyGrowthCurve2DTrajectoryP2CurveShape]*PartiallyGrowthCurve2DTrajectoryP2CurveShape)
	stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes_referenceOrder = make(map[*PartiallyGrowthCurve2DTrajectoryP2CurveShape]uint) // diff Unstage needs the reference order
	stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes_instance = make(map[*PartiallyGrowthCurve2DTrajectoryP2CurveShape]*PartiallyGrowthCurve2DTrajectoryP2CurveShape)
	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes {
		_copy := instance.GongCopy().(*PartiallyGrowthCurve2DTrajectoryP2CurveShape)
		stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes_reference[instance] = _copy
		stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes_instance[_copy] = instance
		stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes_reference = make(map[*PartiallyGrowthCurve2DTrajectoryP2PointShape]*PartiallyGrowthCurve2DTrajectoryP2PointShape)
	stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes_referenceOrder = make(map[*PartiallyGrowthCurve2DTrajectoryP2PointShape]uint) // diff Unstage needs the reference order
	stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes_instance = make(map[*PartiallyGrowthCurve2DTrajectoryP2PointShape]*PartiallyGrowthCurve2DTrajectoryP2PointShape)
	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes {
		_copy := instance.GongCopy().(*PartiallyGrowthCurve2DTrajectoryP2PointShape)
		stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes_reference[instance] = _copy
		stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes_instance[_copy] = instance
		stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartiallyGrowthCurve2DTrajectoryShapes_reference = make(map[*PartiallyGrowthCurve2DTrajectoryShape]*PartiallyGrowthCurve2DTrajectoryShape)
	stage.PartiallyGrowthCurve2DTrajectoryShapes_referenceOrder = make(map[*PartiallyGrowthCurve2DTrajectoryShape]uint) // diff Unstage needs the reference order
	stage.PartiallyGrowthCurve2DTrajectoryShapes_instance = make(map[*PartiallyGrowthCurve2DTrajectoryShape]*PartiallyGrowthCurve2DTrajectoryShape)
	for instance := range stage.PartiallyGrowthCurve2DTrajectoryShapes {
		_copy := instance.GongCopy().(*PartiallyGrowthCurve2DTrajectoryShape)
		stage.PartiallyGrowthCurve2DTrajectoryShapes_reference[instance] = _copy
		stage.PartiallyGrowthCurve2DTrajectoryShapes_instance[_copy] = instance
		stage.PartiallyGrowthCurve2DTrajectoryShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartiallyRotatedSeatBottomCurveShapes_reference = make(map[*PartiallyRotatedSeatBottomCurveShape]*PartiallyRotatedSeatBottomCurveShape)
	stage.PartiallyRotatedSeatBottomCurveShapes_referenceOrder = make(map[*PartiallyRotatedSeatBottomCurveShape]uint) // diff Unstage needs the reference order
	stage.PartiallyRotatedSeatBottomCurveShapes_instance = make(map[*PartiallyRotatedSeatBottomCurveShape]*PartiallyRotatedSeatBottomCurveShape)
	for instance := range stage.PartiallyRotatedSeatBottomCurveShapes {
		_copy := instance.GongCopy().(*PartiallyRotatedSeatBottomCurveShape)
		stage.PartiallyRotatedSeatBottomCurveShapes_reference[instance] = _copy
		stage.PartiallyRotatedSeatBottomCurveShapes_instance[_copy] = instance
		stage.PartiallyRotatedSeatBottomCurveShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartiallyRotatedSeatTopCurveShapes_reference = make(map[*PartiallyRotatedSeatTopCurveShape]*PartiallyRotatedSeatTopCurveShape)
	stage.PartiallyRotatedSeatTopCurveShapes_referenceOrder = make(map[*PartiallyRotatedSeatTopCurveShape]uint) // diff Unstage needs the reference order
	stage.PartiallyRotatedSeatTopCurveShapes_instance = make(map[*PartiallyRotatedSeatTopCurveShape]*PartiallyRotatedSeatTopCurveShape)
	for instance := range stage.PartiallyRotatedSeatTopCurveShapes {
		_copy := instance.GongCopy().(*PartiallyRotatedSeatTopCurveShape)
		stage.PartiallyRotatedSeatTopCurveShapes_reference[instance] = _copy
		stage.PartiallyRotatedSeatTopCurveShapes_instance[_copy] = instance
		stage.PartiallyRotatedSeatTopCurveShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartiallyRotatedTorusShapes_reference = make(map[*PartiallyRotatedTorusShape]*PartiallyRotatedTorusShape)
	stage.PartiallyRotatedTorusShapes_referenceOrder = make(map[*PartiallyRotatedTorusShape]uint) // diff Unstage needs the reference order
	stage.PartiallyRotatedTorusShapes_instance = make(map[*PartiallyRotatedTorusShape]*PartiallyRotatedTorusShape)
	for instance := range stage.PartiallyRotatedTorusShapes {
		_copy := instance.GongCopy().(*PartiallyRotatedTorusShape)
		stage.PartiallyRotatedTorusShapes_reference[instance] = _copy
		stage.PartiallyRotatedTorusShapes_instance[_copy] = instance
		stage.PartiallyRotatedTorusShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.Plant2DDiagrams_reference = make(map[*Plant2DDiagram]*Plant2DDiagram)
	stage.Plant2DDiagrams_referenceOrder = make(map[*Plant2DDiagram]uint) // diff Unstage needs the reference order
	stage.Plant2DDiagrams_instance = make(map[*Plant2DDiagram]*Plant2DDiagram)
	for instance := range stage.Plant2DDiagrams {
		_copy := instance.GongCopy().(*Plant2DDiagram)
		stage.Plant2DDiagrams_reference[instance] = _copy
		stage.Plant2DDiagrams_instance[_copy] = instance
		stage.Plant2DDiagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Plant3DDiagrams_reference = make(map[*Plant3DDiagram]*Plant3DDiagram)
	stage.Plant3DDiagrams_referenceOrder = make(map[*Plant3DDiagram]uint) // diff Unstage needs the reference order
	stage.Plant3DDiagrams_instance = make(map[*Plant3DDiagram]*Plant3DDiagram)
	for instance := range stage.Plant3DDiagrams {
		_copy := instance.GongCopy().(*Plant3DDiagram)
		stage.Plant3DDiagrams_reference[instance] = _copy
		stage.Plant3DDiagrams_instance[_copy] = instance
		stage.Plant3DDiagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PlantAbstracts_reference = make(map[*PlantAbstract]*PlantAbstract)
	stage.PlantAbstracts_referenceOrder = make(map[*PlantAbstract]uint) // diff Unstage needs the reference order
	stage.PlantAbstracts_instance = make(map[*PlantAbstract]*PlantAbstract)
	for instance := range stage.PlantAbstracts {
		_copy := instance.GongCopy().(*PlantAbstract)
		stage.PlantAbstracts_reference[instance] = _copy
		stage.PlantAbstracts_instance[_copy] = instance
		stage.PlantAbstracts_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.PointsAndLines3DShapes_reference = make(map[*PointsAndLines3DShape]*PointsAndLines3DShape)
	stage.PointsAndLines3DShapes_referenceOrder = make(map[*PointsAndLines3DShape]uint) // diff Unstage needs the reference order
	stage.PointsAndLines3DShapes_instance = make(map[*PointsAndLines3DShape]*PointsAndLines3DShape)
	for instance := range stage.PointsAndLines3DShapes {
		_copy := instance.GongCopy().(*PointsAndLines3DShape)
		stage.PointsAndLines3DShapes_reference[instance] = _copy
		stage.PointsAndLines3DShapes_instance[_copy] = instance
		stage.PointsAndLines3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PxShapes_reference = make(map[*PxShape]*PxShape)
	stage.PxShapes_referenceOrder = make(map[*PxShape]uint) // diff Unstage needs the reference order
	stage.PxShapes_instance = make(map[*PxShape]*PxShape)
	for instance := range stage.PxShapes {
		_copy := instance.GongCopy().(*PxShape)
		stage.PxShapes_reference[instance] = _copy
		stage.PxShapes_instance[_copy] = instance
		stage.PxShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.RotatedSampledPoints3DShapes_reference = make(map[*RotatedSampledPoints3DShape]*RotatedSampledPoints3DShape)
	stage.RotatedSampledPoints3DShapes_referenceOrder = make(map[*RotatedSampledPoints3DShape]uint) // diff Unstage needs the reference order
	stage.RotatedSampledPoints3DShapes_instance = make(map[*RotatedSampledPoints3DShape]*RotatedSampledPoints3DShape)
	for instance := range stage.RotatedSampledPoints3DShapes {
		_copy := instance.GongCopy().(*RotatedSampledPoints3DShape)
		stage.RotatedSampledPoints3DShapes_reference[instance] = _copy
		stage.RotatedSampledPoints3DShapes_instance[_copy] = instance
		stage.RotatedSampledPoints3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RotatedSeatAndLegs3DShapes_reference = make(map[*RotatedSeatAndLegs3DShape]*RotatedSeatAndLegs3DShape)
	stage.RotatedSeatAndLegs3DShapes_referenceOrder = make(map[*RotatedSeatAndLegs3DShape]uint) // diff Unstage needs the reference order
	stage.RotatedSeatAndLegs3DShapes_instance = make(map[*RotatedSeatAndLegs3DShape]*RotatedSeatAndLegs3DShape)
	for instance := range stage.RotatedSeatAndLegs3DShapes {
		_copy := instance.GongCopy().(*RotatedSeatAndLegs3DShape)
		stage.RotatedSeatAndLegs3DShapes_reference[instance] = _copy
		stage.RotatedSeatAndLegs3DShapes_instance[_copy] = instance
		stage.RotatedSeatAndLegs3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SampledPoints3DShapes_reference = make(map[*SampledPoints3DShape]*SampledPoints3DShape)
	stage.SampledPoints3DShapes_referenceOrder = make(map[*SampledPoints3DShape]uint) // diff Unstage needs the reference order
	stage.SampledPoints3DShapes_instance = make(map[*SampledPoints3DShape]*SampledPoints3DShape)
	for instance := range stage.SampledPoints3DShapes {
		_copy := instance.GongCopy().(*SampledPoints3DShape)
		stage.SampledPoints3DShapes_reference[instance] = _copy
		stage.SampledPoints3DShapes_instance[_copy] = instance
		stage.SampledPoints3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Seat3DShapes_reference = make(map[*Seat3DShape]*Seat3DShape)
	stage.Seat3DShapes_referenceOrder = make(map[*Seat3DShape]uint) // diff Unstage needs the reference order
	stage.Seat3DShapes_instance = make(map[*Seat3DShape]*Seat3DShape)
	for instance := range stage.Seat3DShapes {
		_copy := instance.GongCopy().(*Seat3DShape)
		stage.Seat3DShapes_reference[instance] = _copy
		stage.Seat3DShapes_instance[_copy] = instance
		stage.Seat3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SeatAndLegs3DShapes_reference = make(map[*SeatAndLegs3DShape]*SeatAndLegs3DShape)
	stage.SeatAndLegs3DShapes_referenceOrder = make(map[*SeatAndLegs3DShape]uint) // diff Unstage needs the reference order
	stage.SeatAndLegs3DShapes_instance = make(map[*SeatAndLegs3DShape]*SeatAndLegs3DShape)
	for instance := range stage.SeatAndLegs3DShapes {
		_copy := instance.GongCopy().(*SeatAndLegs3DShape)
		stage.SeatAndLegs3DShapes_reference[instance] = _copy
		stage.SeatAndLegs3DShapes_instance[_copy] = instance
		stage.SeatAndLegs3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SeatBottomCurveShapes_reference = make(map[*SeatBottomCurveShape]*SeatBottomCurveShape)
	stage.SeatBottomCurveShapes_referenceOrder = make(map[*SeatBottomCurveShape]uint) // diff Unstage needs the reference order
	stage.SeatBottomCurveShapes_instance = make(map[*SeatBottomCurveShape]*SeatBottomCurveShape)
	for instance := range stage.SeatBottomCurveShapes {
		_copy := instance.GongCopy().(*SeatBottomCurveShape)
		stage.SeatBottomCurveShapes_reference[instance] = _copy
		stage.SeatBottomCurveShapes_instance[_copy] = instance
		stage.SeatBottomCurveShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SeatTopCurveShapes_reference = make(map[*SeatTopCurveShape]*SeatTopCurveShape)
	stage.SeatTopCurveShapes_referenceOrder = make(map[*SeatTopCurveShape]uint) // diff Unstage needs the reference order
	stage.SeatTopCurveShapes_instance = make(map[*SeatTopCurveShape]*SeatTopCurveShape)
	for instance := range stage.SeatTopCurveShapes {
		_copy := instance.GongCopy().(*SeatTopCurveShape)
		stage.SeatTopCurveShapes_reference[instance] = _copy
		stage.SeatTopCurveShapes_instance[_copy] = instance
		stage.SeatTopCurveShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.ShiftedLeftGrowthCurve2DRibbons_reference = make(map[*ShiftedLeftGrowthCurve2DRibbon]*ShiftedLeftGrowthCurve2DRibbon)
	stage.ShiftedLeftGrowthCurve2DRibbons_referenceOrder = make(map[*ShiftedLeftGrowthCurve2DRibbon]uint) // diff Unstage needs the reference order
	stage.ShiftedLeftGrowthCurve2DRibbons_instance = make(map[*ShiftedLeftGrowthCurve2DRibbon]*ShiftedLeftGrowthCurve2DRibbon)
	for instance := range stage.ShiftedLeftGrowthCurve2DRibbons {
		_copy := instance.GongCopy().(*ShiftedLeftGrowthCurve2DRibbon)
		stage.ShiftedLeftGrowthCurve2DRibbons_reference[instance] = _copy
		stage.ShiftedLeftGrowthCurve2DRibbons_instance[_copy] = instance
		stage.ShiftedLeftGrowthCurve2DRibbons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShiftedLeftGrowthCurve2DRibbonEndShapes_reference = make(map[*ShiftedLeftGrowthCurve2DRibbonEndShape]*ShiftedLeftGrowthCurve2DRibbonEndShape)
	stage.ShiftedLeftGrowthCurve2DRibbonEndShapes_referenceOrder = make(map[*ShiftedLeftGrowthCurve2DRibbonEndShape]uint) // diff Unstage needs the reference order
	stage.ShiftedLeftGrowthCurve2DRibbonEndShapes_instance = make(map[*ShiftedLeftGrowthCurve2DRibbonEndShape]*ShiftedLeftGrowthCurve2DRibbonEndShape)
	for instance := range stage.ShiftedLeftGrowthCurve2DRibbonEndShapes {
		_copy := instance.GongCopy().(*ShiftedLeftGrowthCurve2DRibbonEndShape)
		stage.ShiftedLeftGrowthCurve2DRibbonEndShapes_reference[instance] = _copy
		stage.ShiftedLeftGrowthCurve2DRibbonEndShapes_instance[_copy] = instance
		stage.ShiftedLeftGrowthCurve2DRibbonEndShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShiftedLeftGrowthCurve2DRibbonStartShapes_reference = make(map[*ShiftedLeftGrowthCurve2DRibbonStartShape]*ShiftedLeftGrowthCurve2DRibbonStartShape)
	stage.ShiftedLeftGrowthCurve2DRibbonStartShapes_referenceOrder = make(map[*ShiftedLeftGrowthCurve2DRibbonStartShape]uint) // diff Unstage needs the reference order
	stage.ShiftedLeftGrowthCurve2DRibbonStartShapes_instance = make(map[*ShiftedLeftGrowthCurve2DRibbonStartShape]*ShiftedLeftGrowthCurve2DRibbonStartShape)
	for instance := range stage.ShiftedLeftGrowthCurve2DRibbonStartShapes {
		_copy := instance.GongCopy().(*ShiftedLeftGrowthCurve2DRibbonStartShape)
		stage.ShiftedLeftGrowthCurve2DRibbonStartShapes_reference[instance] = _copy
		stage.ShiftedLeftGrowthCurve2DRibbonStartShapes_instance[_copy] = instance
		stage.ShiftedLeftGrowthCurve2DRibbonStartShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShiftedLeftPartiallyGrowthCurve2DRibbons_reference = make(map[*ShiftedLeftPartiallyGrowthCurve2DRibbon]*ShiftedLeftPartiallyGrowthCurve2DRibbon)
	stage.ShiftedLeftPartiallyGrowthCurve2DRibbons_referenceOrder = make(map[*ShiftedLeftPartiallyGrowthCurve2DRibbon]uint) // diff Unstage needs the reference order
	stage.ShiftedLeftPartiallyGrowthCurve2DRibbons_instance = make(map[*ShiftedLeftPartiallyGrowthCurve2DRibbon]*ShiftedLeftPartiallyGrowthCurve2DRibbon)
	for instance := range stage.ShiftedLeftPartiallyGrowthCurve2DRibbons {
		_copy := instance.GongCopy().(*ShiftedLeftPartiallyGrowthCurve2DRibbon)
		stage.ShiftedLeftPartiallyGrowthCurve2DRibbons_reference[instance] = _copy
		stage.ShiftedLeftPartiallyGrowthCurve2DRibbons_instance[_copy] = instance
		stage.ShiftedLeftPartiallyGrowthCurve2DRibbons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes_reference = make(map[*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape]*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
	stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes_referenceOrder = make(map[*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape]uint) // diff Unstage needs the reference order
	stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes_instance = make(map[*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape]*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
	for instance := range stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes {
		_copy := instance.GongCopy().(*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
		stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes_reference[instance] = _copy
		stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes_instance[_copy] = instance
		stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes_reference = make(map[*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape]*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
	stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes_referenceOrder = make(map[*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape]uint) // diff Unstage needs the reference order
	stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes_instance = make(map[*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape]*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
	for instance := range stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes {
		_copy := instance.GongCopy().(*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
		stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes_reference[instance] = _copy
		stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes_instance[_copy] = instance
		stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.ShiftedRightGrowthCurve2DRibbons_reference = make(map[*ShiftedRightGrowthCurve2DRibbon]*ShiftedRightGrowthCurve2DRibbon)
	stage.ShiftedRightGrowthCurve2DRibbons_referenceOrder = make(map[*ShiftedRightGrowthCurve2DRibbon]uint) // diff Unstage needs the reference order
	stage.ShiftedRightGrowthCurve2DRibbons_instance = make(map[*ShiftedRightGrowthCurve2DRibbon]*ShiftedRightGrowthCurve2DRibbon)
	for instance := range stage.ShiftedRightGrowthCurve2DRibbons {
		_copy := instance.GongCopy().(*ShiftedRightGrowthCurve2DRibbon)
		stage.ShiftedRightGrowthCurve2DRibbons_reference[instance] = _copy
		stage.ShiftedRightGrowthCurve2DRibbons_instance[_copy] = instance
		stage.ShiftedRightGrowthCurve2DRibbons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShiftedRightGrowthCurve2DRibbonEndShapes_reference = make(map[*ShiftedRightGrowthCurve2DRibbonEndShape]*ShiftedRightGrowthCurve2DRibbonEndShape)
	stage.ShiftedRightGrowthCurve2DRibbonEndShapes_referenceOrder = make(map[*ShiftedRightGrowthCurve2DRibbonEndShape]uint) // diff Unstage needs the reference order
	stage.ShiftedRightGrowthCurve2DRibbonEndShapes_instance = make(map[*ShiftedRightGrowthCurve2DRibbonEndShape]*ShiftedRightGrowthCurve2DRibbonEndShape)
	for instance := range stage.ShiftedRightGrowthCurve2DRibbonEndShapes {
		_copy := instance.GongCopy().(*ShiftedRightGrowthCurve2DRibbonEndShape)
		stage.ShiftedRightGrowthCurve2DRibbonEndShapes_reference[instance] = _copy
		stage.ShiftedRightGrowthCurve2DRibbonEndShapes_instance[_copy] = instance
		stage.ShiftedRightGrowthCurve2DRibbonEndShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ShiftedRightGrowthCurve2DRibbonStartShapes_reference = make(map[*ShiftedRightGrowthCurve2DRibbonStartShape]*ShiftedRightGrowthCurve2DRibbonStartShape)
	stage.ShiftedRightGrowthCurve2DRibbonStartShapes_referenceOrder = make(map[*ShiftedRightGrowthCurve2DRibbonStartShape]uint) // diff Unstage needs the reference order
	stage.ShiftedRightGrowthCurve2DRibbonStartShapes_instance = make(map[*ShiftedRightGrowthCurve2DRibbonStartShape]*ShiftedRightGrowthCurve2DRibbonStartShape)
	for instance := range stage.ShiftedRightGrowthCurve2DRibbonStartShapes {
		_copy := instance.GongCopy().(*ShiftedRightGrowthCurve2DRibbonStartShape)
		stage.ShiftedRightGrowthCurve2DRibbonStartShapes_reference[instance] = _copy
		stage.ShiftedRightGrowthCurve2DRibbonStartShapes_instance[_copy] = instance
		stage.ShiftedRightGrowthCurve2DRibbonStartShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.StackOfGrowthCurve2DByGrowthVectors_reference = make(map[*StackOfGrowthCurve2DByGrowthVector]*StackOfGrowthCurve2DByGrowthVector)
	stage.StackOfGrowthCurve2DByGrowthVectors_referenceOrder = make(map[*StackOfGrowthCurve2DByGrowthVector]uint) // diff Unstage needs the reference order
	stage.StackOfGrowthCurve2DByGrowthVectors_instance = make(map[*StackOfGrowthCurve2DByGrowthVector]*StackOfGrowthCurve2DByGrowthVector)
	for instance := range stage.StackOfGrowthCurve2DByGrowthVectors {
		_copy := instance.GongCopy().(*StackOfGrowthCurve2DByGrowthVector)
		stage.StackOfGrowthCurve2DByGrowthVectors_reference[instance] = _copy
		stage.StackOfGrowthCurve2DByGrowthVectors_instance[_copy] = instance
		stage.StackOfGrowthCurve2DByGrowthVectors_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.StackOfPartiallyRotatedTorusShapes_reference = make(map[*StackOfPartiallyRotatedTorusShape]*StackOfPartiallyRotatedTorusShape)
	stage.StackOfPartiallyRotatedTorusShapes_referenceOrder = make(map[*StackOfPartiallyRotatedTorusShape]uint) // diff Unstage needs the reference order
	stage.StackOfPartiallyRotatedTorusShapes_instance = make(map[*StackOfPartiallyRotatedTorusShape]*StackOfPartiallyRotatedTorusShape)
	for instance := range stage.StackOfPartiallyRotatedTorusShapes {
		_copy := instance.GongCopy().(*StackOfPartiallyRotatedTorusShape)
		stage.StackOfPartiallyRotatedTorusShapes_reference[instance] = _copy
		stage.StackOfPartiallyRotatedTorusShapes_instance[_copy] = instance
		stage.StackOfPartiallyRotatedTorusShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.StackOfRotatedVaseTrapezeRingsShapes_reference = make(map[*StackOfRotatedVaseTrapezeRingsShape]*StackOfRotatedVaseTrapezeRingsShape)
	stage.StackOfRotatedVaseTrapezeRingsShapes_referenceOrder = make(map[*StackOfRotatedVaseTrapezeRingsShape]uint) // diff Unstage needs the reference order
	stage.StackOfRotatedVaseTrapezeRingsShapes_instance = make(map[*StackOfRotatedVaseTrapezeRingsShape]*StackOfRotatedVaseTrapezeRingsShape)
	for instance := range stage.StackOfRotatedVaseTrapezeRingsShapes {
		_copy := instance.GongCopy().(*StackOfRotatedVaseTrapezeRingsShape)
		stage.StackOfRotatedVaseTrapezeRingsShapes_reference[instance] = _copy
		stage.StackOfRotatedVaseTrapezeRingsShapes_instance[_copy] = instance
		stage.StackOfRotatedVaseTrapezeRingsShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StackOfVaseTrapezeRingsShapes_reference = make(map[*StackOfVaseTrapezeRingsShape]*StackOfVaseTrapezeRingsShape)
	stage.StackOfVaseTrapezeRingsShapes_referenceOrder = make(map[*StackOfVaseTrapezeRingsShape]uint) // diff Unstage needs the reference order
	stage.StackOfVaseTrapezeRingsShapes_instance = make(map[*StackOfVaseTrapezeRingsShape]*StackOfVaseTrapezeRingsShape)
	for instance := range stage.StackOfVaseTrapezeRingsShapes {
		_copy := instance.GongCopy().(*StackOfVaseTrapezeRingsShape)
		stage.StackOfVaseTrapezeRingsShapes_reference[instance] = _copy
		stage.StackOfVaseTrapezeRingsShapes_instance[_copy] = instance
		stage.StackOfVaseTrapezeRingsShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.StemCylinder3DShapes_reference = make(map[*StemCylinder3DShape]*StemCylinder3DShape)
	stage.StemCylinder3DShapes_referenceOrder = make(map[*StemCylinder3DShape]uint) // diff Unstage needs the reference order
	stage.StemCylinder3DShapes_instance = make(map[*StemCylinder3DShape]*StemCylinder3DShape)
	for instance := range stage.StemCylinder3DShapes {
		_copy := instance.GongCopy().(*StemCylinder3DShape)
		stage.StemCylinder3DShapes_reference[instance] = _copy
		stage.StemCylinder3DShapes_instance[_copy] = instance
		stage.StemCylinder3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Stool2DDiagrams_reference = make(map[*Stool2DDiagram]*Stool2DDiagram)
	stage.Stool2DDiagrams_referenceOrder = make(map[*Stool2DDiagram]uint) // diff Unstage needs the reference order
	stage.Stool2DDiagrams_instance = make(map[*Stool2DDiagram]*Stool2DDiagram)
	for instance := range stage.Stool2DDiagrams {
		_copy := instance.GongCopy().(*Stool2DDiagram)
		stage.Stool2DDiagrams_reference[instance] = _copy
		stage.Stool2DDiagrams_instance[_copy] = instance
		stage.Stool2DDiagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Stool3DDiagrams_reference = make(map[*Stool3DDiagram]*Stool3DDiagram)
	stage.Stool3DDiagrams_referenceOrder = make(map[*Stool3DDiagram]uint) // diff Unstage needs the reference order
	stage.Stool3DDiagrams_instance = make(map[*Stool3DDiagram]*Stool3DDiagram)
	for instance := range stage.Stool3DDiagrams {
		_copy := instance.GongCopy().(*Stool3DDiagram)
		stage.Stool3DDiagrams_reference[instance] = _copy
		stage.Stool3DDiagrams_instance[_copy] = instance
		stage.Stool3DDiagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TiledFloor3DShapes_reference = make(map[*TiledFloor3DShape]*TiledFloor3DShape)
	stage.TiledFloor3DShapes_referenceOrder = make(map[*TiledFloor3DShape]uint) // diff Unstage needs the reference order
	stage.TiledFloor3DShapes_instance = make(map[*TiledFloor3DShape]*TiledFloor3DShape)
	for instance := range stage.TiledFloor3DShapes {
		_copy := instance.GongCopy().(*TiledFloor3DShape)
		stage.TiledFloor3DShapes_reference[instance] = _copy
		stage.TiledFloor3DShapes_instance[_copy] = instance
		stage.TiledFloor3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopCurvePlane1Shapes_reference = make(map[*TopCurvePlane1Shape]*TopCurvePlane1Shape)
	stage.TopCurvePlane1Shapes_referenceOrder = make(map[*TopCurvePlane1Shape]uint) // diff Unstage needs the reference order
	stage.TopCurvePlane1Shapes_instance = make(map[*TopCurvePlane1Shape]*TopCurvePlane1Shape)
	for instance := range stage.TopCurvePlane1Shapes {
		_copy := instance.GongCopy().(*TopCurvePlane1Shape)
		stage.TopCurvePlane1Shapes_reference[instance] = _copy
		stage.TopCurvePlane1Shapes_instance[_copy] = instance
		stage.TopCurvePlane1Shapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TopCurvePlane2Shapes_reference = make(map[*TopCurvePlane2Shape]*TopCurvePlane2Shape)
	stage.TopCurvePlane2Shapes_referenceOrder = make(map[*TopCurvePlane2Shape]uint) // diff Unstage needs the reference order
	stage.TopCurvePlane2Shapes_instance = make(map[*TopCurvePlane2Shape]*TopCurvePlane2Shape)
	for instance := range stage.TopCurvePlane2Shapes {
		_copy := instance.GongCopy().(*TopCurvePlane2Shape)
		stage.TopCurvePlane2Shapes_reference[instance] = _copy
		stage.TopCurvePlane2Shapes_instance[_copy] = instance
		stage.TopCurvePlane2Shapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.Torus3DShapes_reference = make(map[*Torus3DShape]*Torus3DShape)
	stage.Torus3DShapes_referenceOrder = make(map[*Torus3DShape]uint) // diff Unstage needs the reference order
	stage.Torus3DShapes_instance = make(map[*Torus3DShape]*Torus3DShape)
	for instance := range stage.Torus3DShapes {
		_copy := instance.GongCopy().(*Torus3DShape)
		stage.Torus3DShapes_reference[instance] = _copy
		stage.Torus3DShapes_instance[_copy] = instance
		stage.Torus3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TorusEdge3DShapes_reference = make(map[*TorusEdge3DShape]*TorusEdge3DShape)
	stage.TorusEdge3DShapes_referenceOrder = make(map[*TorusEdge3DShape]uint) // diff Unstage needs the reference order
	stage.TorusEdge3DShapes_instance = make(map[*TorusEdge3DShape]*TorusEdge3DShape)
	for instance := range stage.TorusEdge3DShapes {
		_copy := instance.GongCopy().(*TorusEdge3DShape)
		stage.TorusEdge3DShapes_reference[instance] = _copy
		stage.TorusEdge3DShapes_instance[_copy] = instance
		stage.TorusEdge3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.TubeVase3DDiagrams_reference = make(map[*TubeVase3DDiagram]*TubeVase3DDiagram)
	stage.TubeVase3DDiagrams_referenceOrder = make(map[*TubeVase3DDiagram]uint) // diff Unstage needs the reference order
	stage.TubeVase3DDiagrams_instance = make(map[*TubeVase3DDiagram]*TubeVase3DDiagram)
	for instance := range stage.TubeVase3DDiagrams {
		_copy := instance.GongCopy().(*TubeVase3DDiagram)
		stage.TubeVase3DDiagrams_reference[instance] = _copy
		stage.TubeVase3DDiagrams_instance[_copy] = instance
		stage.TubeVase3DDiagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TubeVaseAbstracts_reference = make(map[*TubeVaseAbstract]*TubeVaseAbstract)
	stage.TubeVaseAbstracts_referenceOrder = make(map[*TubeVaseAbstract]uint) // diff Unstage needs the reference order
	stage.TubeVaseAbstracts_instance = make(map[*TubeVaseAbstract]*TubeVaseAbstract)
	for instance := range stage.TubeVaseAbstracts {
		_copy := instance.GongCopy().(*TubeVaseAbstract)
		stage.TubeVaseAbstracts_reference[instance] = _copy
		stage.TubeVaseAbstracts_instance[_copy] = instance
		stage.TubeVaseAbstracts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Vase2DDiagrams_reference = make(map[*Vase2DDiagram]*Vase2DDiagram)
	stage.Vase2DDiagrams_referenceOrder = make(map[*Vase2DDiagram]uint) // diff Unstage needs the reference order
	stage.Vase2DDiagrams_instance = make(map[*Vase2DDiagram]*Vase2DDiagram)
	for instance := range stage.Vase2DDiagrams {
		_copy := instance.GongCopy().(*Vase2DDiagram)
		stage.Vase2DDiagrams_reference[instance] = _copy
		stage.Vase2DDiagrams_instance[_copy] = instance
		stage.Vase2DDiagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.VaseTrapezeRingShapes_reference = make(map[*VaseTrapezeRingShape]*VaseTrapezeRingShape)
	stage.VaseTrapezeRingShapes_referenceOrder = make(map[*VaseTrapezeRingShape]uint) // diff Unstage needs the reference order
	stage.VaseTrapezeRingShapes_instance = make(map[*VaseTrapezeRingShape]*VaseTrapezeRingShape)
	for instance := range stage.VaseTrapezeRingShapes {
		_copy := instance.GongCopy().(*VaseTrapezeRingShape)
		stage.VaseTrapezeRingShapes_reference[instance] = _copy
		stage.VaseTrapezeRingShapes_instance[_copy] = instance
		stage.VaseTrapezeRingShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.VolumeKey3DShapes_reference = make(map[*VolumeKey3DShape]*VolumeKey3DShape)
	stage.VolumeKey3DShapes_referenceOrder = make(map[*VolumeKey3DShape]uint) // diff Unstage needs the reference order
	stage.VolumeKey3DShapes_instance = make(map[*VolumeKey3DShape]*VolumeKey3DShape)
	for instance := range stage.VolumeKey3DShapes {
		_copy := instance.GongCopy().(*VolumeKey3DShape)
		stage.VolumeKey3DShapes_reference[instance] = _copy
		stage.VolumeKey3DShapes_instance[_copy] = instance
		stage.VolumeKey3DShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Angle0Shapes {
		reference := stage.Angle0Shapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

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

	for instance := range stage.BottomCurvePlane1Shapes {
		reference := stage.BottomCurvePlane1Shapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BottomCurvePlane2Shapes {
		reference := stage.BottomCurvePlane2Shapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ChosenP1P2PairShapes {
		reference := stage.ChosenP1P2PairShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.CircleGridShapes {
		reference := stage.CircleGridShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Circumference3DShapes {
		reference := stage.Circumference3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Clock2DDiagrams {
		reference := stage.Clock2DDiagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Clock3DDiagrams {
		reference := stage.Clock3DDiagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ClockTopCurveShapes {
		reference := stage.ClockTopCurveShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.CutLine3DShapes {
		reference := stage.CutLine3DShapes_reference[instance]
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

	for instance := range stage.Eye3DShapes {
		reference := stage.Eye3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EyeCornersSampledPoints3DShapes {
		reference := stage.EyeCornersSampledPoints3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EyeSampledPoints3DShapes {
		reference := stage.EyeSampledPoints3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EyeSeatBottomCurveShapes {
		reference := stage.EyeSeatBottomCurveShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EyeStoolBottomCurveShapes {
		reference := stage.EyeStoolBottomCurveShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EyeVolume3DShapes {
		reference := stage.EyeVolume3DShapes_reference[instance]
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

	for instance := range stage.GrowthCurve2DRibbons {
		reference := stage.GrowthCurve2DRibbons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GrowthCurve2DRibbonEndShapes {
		reference := stage.GrowthCurve2DRibbonEndShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GrowthCurve2DRibbonStartShapes {
		reference := stage.GrowthCurve2DRibbonStartShapes_reference[instance]
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

	for instance := range stage.Key3DShapes {
		reference := stage.Key3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.KeyHole3DShapes {
		reference := stage.KeyHole3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.KeyHoleShapes {
		reference := stage.KeyHoleShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Leaves3DShapes {
		reference := stage.Leaves3DShapes_reference[instance]
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

	for instance := range stage.OriginalPoints3DShapes {
		reference := stage.OriginalPoints3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ParastichyMCurves3DShapes {
		reference := stage.ParastichyMCurves3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ParastichyNCurves3DShapes {
		reference := stage.ParastichyNCurves3DShapes_reference[instance]
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

	for instance := range stage.PartiallyGrowthCurve2DTrajectorys {
		reference := stage.PartiallyGrowthCurve2DTrajectorys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes {
		reference := stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1P2s {
		reference := stage.PartiallyGrowthCurve2DTrajectoryP1P2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes {
		reference := stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes {
		reference := stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes {
		reference := stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes {
		reference := stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartiallyGrowthCurve2DTrajectoryShapes {
		reference := stage.PartiallyGrowthCurve2DTrajectoryShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartiallyRotatedSeatBottomCurveShapes {
		reference := stage.PartiallyRotatedSeatBottomCurveShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartiallyRotatedSeatTopCurveShapes {
		reference := stage.PartiallyRotatedSeatTopCurveShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartiallyRotatedTorusShapes {
		reference := stage.PartiallyRotatedTorusShapes_reference[instance]
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

	for instance := range stage.Plant2DDiagrams {
		reference := stage.Plant2DDiagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Plant3DDiagrams {
		reference := stage.Plant3DDiagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PlantAbstracts {
		reference := stage.PlantAbstracts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PlantCircumferenceShapes {
		reference := stage.PlantCircumferenceShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PointsAndLines3DShapes {
		reference := stage.PointsAndLines3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PxShapes {
		reference := stage.PxShapes_reference[instance]
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

	for instance := range stage.RotatedSampledPoints3DShapes {
		reference := stage.RotatedSampledPoints3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RotatedSeatAndLegs3DShapes {
		reference := stage.RotatedSeatAndLegs3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SampledPoints3DShapes {
		reference := stage.SampledPoints3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Seat3DShapes {
		reference := stage.Seat3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SeatAndLegs3DShapes {
		reference := stage.SeatAndLegs3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SeatBottomCurveShapes {
		reference := stage.SeatBottomCurveShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SeatTopCurveShapes {
		reference := stage.SeatTopCurveShapes_reference[instance]
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

	for instance := range stage.ShiftedLeftGrowthCurve2DRibbons {
		reference := stage.ShiftedLeftGrowthCurve2DRibbons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShiftedLeftGrowthCurve2DRibbonEndShapes {
		reference := stage.ShiftedLeftGrowthCurve2DRibbonEndShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShiftedLeftGrowthCurve2DRibbonStartShapes {
		reference := stage.ShiftedLeftGrowthCurve2DRibbonStartShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShiftedLeftPartiallyGrowthCurve2DRibbons {
		reference := stage.ShiftedLeftPartiallyGrowthCurve2DRibbons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes {
		reference := stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes {
		reference := stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes_reference[instance]
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

	for instance := range stage.ShiftedRightGrowthCurve2DRibbons {
		reference := stage.ShiftedRightGrowthCurve2DRibbons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShiftedRightGrowthCurve2DRibbonEndShapes {
		reference := stage.ShiftedRightGrowthCurve2DRibbonEndShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ShiftedRightGrowthCurve2DRibbonStartShapes {
		reference := stage.ShiftedRightGrowthCurve2DRibbonStartShapes_reference[instance]
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

	for instance := range stage.StackOfGrowthCurve2DByGrowthVectors {
		reference := stage.StackOfGrowthCurve2DByGrowthVectors_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackOfGrowthCurve2DRibbons {
		reference := stage.StackOfGrowthCurve2DRibbons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackOfPartiallyRotatedTorusShapes {
		reference := stage.StackOfPartiallyRotatedTorusShapes_reference[instance]
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

	for instance := range stage.StackOfRotatedVaseTrapezeRingsShapes {
		reference := stage.StackOfRotatedVaseTrapezeRingsShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StackOfVaseTrapezeRingsShapes {
		reference := stage.StackOfVaseTrapezeRingsShapes_reference[instance]
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

	for instance := range stage.StemCylinder3DShapes {
		reference := stage.StemCylinder3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Stool2DDiagrams {
		reference := stage.Stool2DDiagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Stool3DDiagrams {
		reference := stage.Stool3DDiagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TiledFloor3DShapes {
		reference := stage.TiledFloor3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopCurvePlane1Shapes {
		reference := stage.TopCurvePlane1Shapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TopCurvePlane2Shapes {
		reference := stage.TopCurvePlane2Shapes_reference[instance]
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

	for instance := range stage.Torus3DShapes {
		reference := stage.Torus3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TorusEdge3DShapes {
		reference := stage.TorusEdge3DShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TorusStackShapes {
		reference := stage.TorusStackShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TubeVase3DDiagrams {
		reference := stage.TubeVase3DDiagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TubeVaseAbstracts {
		reference := stage.TubeVaseAbstracts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Vase2DDiagrams {
		reference := stage.Vase2DDiagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.VaseTrapezeRingShapes {
		reference := stage.VaseTrapezeRingShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.VerticalTorusStackShapes {
		reference := stage.VerticalTorusStackShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.VolumeKey3DShapes {
		reference := stage.VolumeKey3DShapes_reference[instance]
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
func (angle0shape *Angle0Shape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Angle0Shape_stagedOrder[angle0shape]; ok {
		return order
	}
	if order, ok := stage.Angle0Shapes_referenceOrder[angle0shape]; ok {
		return order
	} else {
		log.Printf("instance %p of type Angle0Shape was not staged and does not have a reference order", angle0shape)
		return 0
	}
}

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

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BottomCurvePlane1Shape_stagedOrder[bottomcurveplane1shape]; ok {
		return order
	}
	if order, ok := stage.BottomCurvePlane1Shapes_referenceOrder[bottomcurveplane1shape]; ok {
		return order
	} else {
		log.Printf("instance %p of type BottomCurvePlane1Shape was not staged and does not have a reference order", bottomcurveplane1shape)
		return 0
	}
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BottomCurvePlane2Shape_stagedOrder[bottomcurveplane2shape]; ok {
		return order
	}
	if order, ok := stage.BottomCurvePlane2Shapes_referenceOrder[bottomcurveplane2shape]; ok {
		return order
	} else {
		log.Printf("instance %p of type BottomCurvePlane2Shape was not staged and does not have a reference order", bottomcurveplane2shape)
		return 0
	}
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ChosenP1P2PairShape_stagedOrder[chosenp1p2pairshape]; ok {
		return order
	}
	if order, ok := stage.ChosenP1P2PairShapes_referenceOrder[chosenp1p2pairshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ChosenP1P2PairShape was not staged and does not have a reference order", chosenp1p2pairshape)
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

func (circumference3dshape *Circumference3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Circumference3DShape_stagedOrder[circumference3dshape]; ok {
		return order
	}
	if order, ok := stage.Circumference3DShapes_referenceOrder[circumference3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type Circumference3DShape was not staged and does not have a reference order", circumference3dshape)
		return 0
	}
}

func (clock2ddiagram *Clock2DDiagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Clock2DDiagram_stagedOrder[clock2ddiagram]; ok {
		return order
	}
	if order, ok := stage.Clock2DDiagrams_referenceOrder[clock2ddiagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Clock2DDiagram was not staged and does not have a reference order", clock2ddiagram)
		return 0
	}
}

func (clock3ddiagram *Clock3DDiagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Clock3DDiagram_stagedOrder[clock3ddiagram]; ok {
		return order
	}
	if order, ok := stage.Clock3DDiagrams_referenceOrder[clock3ddiagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Clock3DDiagram was not staged and does not have a reference order", clock3ddiagram)
		return 0
	}
}

func (clocktopcurveshape *ClockTopCurveShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ClockTopCurveShape_stagedOrder[clocktopcurveshape]; ok {
		return order
	}
	if order, ok := stage.ClockTopCurveShapes_referenceOrder[clocktopcurveshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ClockTopCurveShape was not staged and does not have a reference order", clocktopcurveshape)
		return 0
	}
}

func (cutline3dshape *CutLine3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CutLine3DShape_stagedOrder[cutline3dshape]; ok {
		return order
	}
	if order, ok := stage.CutLine3DShapes_referenceOrder[cutline3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type CutLine3DShape was not staged and does not have a reference order", cutline3dshape)
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

func (eye3dshape *Eye3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Eye3DShape_stagedOrder[eye3dshape]; ok {
		return order
	}
	if order, ok := stage.Eye3DShapes_referenceOrder[eye3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type Eye3DShape was not staged and does not have a reference order", eye3dshape)
		return 0
	}
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EyeCornersSampledPoints3DShape_stagedOrder[eyecornerssampledpoints3dshape]; ok {
		return order
	}
	if order, ok := stage.EyeCornersSampledPoints3DShapes_referenceOrder[eyecornerssampledpoints3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type EyeCornersSampledPoints3DShape was not staged and does not have a reference order", eyecornerssampledpoints3dshape)
		return 0
	}
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EyeSampledPoints3DShape_stagedOrder[eyesampledpoints3dshape]; ok {
		return order
	}
	if order, ok := stage.EyeSampledPoints3DShapes_referenceOrder[eyesampledpoints3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type EyeSampledPoints3DShape was not staged and does not have a reference order", eyesampledpoints3dshape)
		return 0
	}
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EyeSeatBottomCurveShape_stagedOrder[eyeseatbottomcurveshape]; ok {
		return order
	}
	if order, ok := stage.EyeSeatBottomCurveShapes_referenceOrder[eyeseatbottomcurveshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type EyeSeatBottomCurveShape was not staged and does not have a reference order", eyeseatbottomcurveshape)
		return 0
	}
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EyeStoolBottomCurveShape_stagedOrder[eyestoolbottomcurveshape]; ok {
		return order
	}
	if order, ok := stage.EyeStoolBottomCurveShapes_referenceOrder[eyestoolbottomcurveshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type EyeStoolBottomCurveShape was not staged and does not have a reference order", eyestoolbottomcurveshape)
		return 0
	}
}

func (eyevolume3dshape *EyeVolume3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EyeVolume3DShape_stagedOrder[eyevolume3dshape]; ok {
		return order
	}
	if order, ok := stage.EyeVolume3DShapes_referenceOrder[eyevolume3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type EyeVolume3DShape was not staged and does not have a reference order", eyevolume3dshape)
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

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GrowthCurve2DRibbon_stagedOrder[growthcurve2dribbon]; ok {
		return order
	}
	if order, ok := stage.GrowthCurve2DRibbons_referenceOrder[growthcurve2dribbon]; ok {
		return order
	} else {
		log.Printf("instance %p of type GrowthCurve2DRibbon was not staged and does not have a reference order", growthcurve2dribbon)
		return 0
	}
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GrowthCurve2DRibbonEndShape_stagedOrder[growthcurve2dribbonendshape]; ok {
		return order
	}
	if order, ok := stage.GrowthCurve2DRibbonEndShapes_referenceOrder[growthcurve2dribbonendshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type GrowthCurve2DRibbonEndShape was not staged and does not have a reference order", growthcurve2dribbonendshape)
		return 0
	}
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GrowthCurve2DRibbonStartShape_stagedOrder[growthcurve2dribbonstartshape]; ok {
		return order
	}
	if order, ok := stage.GrowthCurve2DRibbonStartShapes_referenceOrder[growthcurve2dribbonstartshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type GrowthCurve2DRibbonStartShape was not staged and does not have a reference order", growthcurve2dribbonstartshape)
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

func (key3dshape *Key3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Key3DShape_stagedOrder[key3dshape]; ok {
		return order
	}
	if order, ok := stage.Key3DShapes_referenceOrder[key3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type Key3DShape was not staged and does not have a reference order", key3dshape)
		return 0
	}
}

func (keyhole3dshape *KeyHole3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.KeyHole3DShape_stagedOrder[keyhole3dshape]; ok {
		return order
	}
	if order, ok := stage.KeyHole3DShapes_referenceOrder[keyhole3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type KeyHole3DShape was not staged and does not have a reference order", keyhole3dshape)
		return 0
	}
}

func (keyholeshape *KeyHoleShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.KeyHoleShape_stagedOrder[keyholeshape]; ok {
		return order
	}
	if order, ok := stage.KeyHoleShapes_referenceOrder[keyholeshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type KeyHoleShape was not staged and does not have a reference order", keyholeshape)
		return 0
	}
}

func (leaves3dshape *Leaves3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Leaves3DShape_stagedOrder[leaves3dshape]; ok {
		return order
	}
	if order, ok := stage.Leaves3DShapes_referenceOrder[leaves3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type Leaves3DShape was not staged and does not have a reference order", leaves3dshape)
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

func (originalpoints3dshape *OriginalPoints3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.OriginalPoints3DShape_stagedOrder[originalpoints3dshape]; ok {
		return order
	}
	if order, ok := stage.OriginalPoints3DShapes_referenceOrder[originalpoints3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type OriginalPoints3DShape was not staged and does not have a reference order", originalpoints3dshape)
		return 0
	}
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ParastichyMCurves3DShape_stagedOrder[parastichymcurves3dshape]; ok {
		return order
	}
	if order, ok := stage.ParastichyMCurves3DShapes_referenceOrder[parastichymcurves3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ParastichyMCurves3DShape was not staged and does not have a reference order", parastichymcurves3dshape)
		return 0
	}
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ParastichyNCurves3DShape_stagedOrder[parastichyncurves3dshape]; ok {
		return order
	}
	if order, ok := stage.ParastichyNCurves3DShapes_referenceOrder[parastichyncurves3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ParastichyNCurves3DShape was not staged and does not have a reference order", parastichyncurves3dshape)
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

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyGrowthCurve2DTrajectory_stagedOrder[partiallygrowthcurve2dtrajectory]; ok {
		return order
	}
	if order, ok := stage.PartiallyGrowthCurve2DTrajectorys_referenceOrder[partiallygrowthcurve2dtrajectory]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyGrowthCurve2DTrajectory was not staged and does not have a reference order", partiallygrowthcurve2dtrajectory)
		return 0
	}
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryP1CurveShape_stagedOrder[partiallygrowthcurve2dtrajectoryp1curveshape]; ok {
		return order
	}
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes_referenceOrder[partiallygrowthcurve2dtrajectoryp1curveshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyGrowthCurve2DTrajectoryP1CurveShape was not staged and does not have a reference order", partiallygrowthcurve2dtrajectoryp1curveshape)
		return 0
	}
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryP1P2_stagedOrder[partiallygrowthcurve2dtrajectoryp1p2]; ok {
		return order
	}
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryP1P2s_referenceOrder[partiallygrowthcurve2dtrajectoryp1p2]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyGrowthCurve2DTrajectoryP1P2 was not staged and does not have a reference order", partiallygrowthcurve2dtrajectoryp1p2)
		return 0
	}
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape_stagedOrder[partiallygrowthcurve2dtrajectoryp1p2pairlineshape]; ok {
		return order
	}
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes_referenceOrder[partiallygrowthcurve2dtrajectoryp1p2pairlineshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape was not staged and does not have a reference order", partiallygrowthcurve2dtrajectoryp1p2pairlineshape)
		return 0
	}
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryP1PointShape_stagedOrder[partiallygrowthcurve2dtrajectoryp1pointshape]; ok {
		return order
	}
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes_referenceOrder[partiallygrowthcurve2dtrajectoryp1pointshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyGrowthCurve2DTrajectoryP1PointShape was not staged and does not have a reference order", partiallygrowthcurve2dtrajectoryp1pointshape)
		return 0
	}
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryP2CurveShape_stagedOrder[partiallygrowthcurve2dtrajectoryp2curveshape]; ok {
		return order
	}
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes_referenceOrder[partiallygrowthcurve2dtrajectoryp2curveshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyGrowthCurve2DTrajectoryP2CurveShape was not staged and does not have a reference order", partiallygrowthcurve2dtrajectoryp2curveshape)
		return 0
	}
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryP2PointShape_stagedOrder[partiallygrowthcurve2dtrajectoryp2pointshape]; ok {
		return order
	}
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes_referenceOrder[partiallygrowthcurve2dtrajectoryp2pointshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyGrowthCurve2DTrajectoryP2PointShape was not staged and does not have a reference order", partiallygrowthcurve2dtrajectoryp2pointshape)
		return 0
	}
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryShape_stagedOrder[partiallygrowthcurve2dtrajectoryshape]; ok {
		return order
	}
	if order, ok := stage.PartiallyGrowthCurve2DTrajectoryShapes_referenceOrder[partiallygrowthcurve2dtrajectoryshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyGrowthCurve2DTrajectoryShape was not staged and does not have a reference order", partiallygrowthcurve2dtrajectoryshape)
		return 0
	}
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyRotatedSeatBottomCurveShape_stagedOrder[partiallyrotatedseatbottomcurveshape]; ok {
		return order
	}
	if order, ok := stage.PartiallyRotatedSeatBottomCurveShapes_referenceOrder[partiallyrotatedseatbottomcurveshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyRotatedSeatBottomCurveShape was not staged and does not have a reference order", partiallyrotatedseatbottomcurveshape)
		return 0
	}
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyRotatedSeatTopCurveShape_stagedOrder[partiallyrotatedseattopcurveshape]; ok {
		return order
	}
	if order, ok := stage.PartiallyRotatedSeatTopCurveShapes_referenceOrder[partiallyrotatedseattopcurveshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyRotatedSeatTopCurveShape was not staged and does not have a reference order", partiallyrotatedseattopcurveshape)
		return 0
	}
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartiallyRotatedTorusShape_stagedOrder[partiallyrotatedtorusshape]; ok {
		return order
	}
	if order, ok := stage.PartiallyRotatedTorusShapes_referenceOrder[partiallyrotatedtorusshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartiallyRotatedTorusShape was not staged and does not have a reference order", partiallyrotatedtorusshape)
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

func (plant2ddiagram *Plant2DDiagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Plant2DDiagram_stagedOrder[plant2ddiagram]; ok {
		return order
	}
	if order, ok := stage.Plant2DDiagrams_referenceOrder[plant2ddiagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Plant2DDiagram was not staged and does not have a reference order", plant2ddiagram)
		return 0
	}
}

func (plant3ddiagram *Plant3DDiagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Plant3DDiagram_stagedOrder[plant3ddiagram]; ok {
		return order
	}
	if order, ok := stage.Plant3DDiagrams_referenceOrder[plant3ddiagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Plant3DDiagram was not staged and does not have a reference order", plant3ddiagram)
		return 0
	}
}

func (plantabstract *PlantAbstract) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PlantAbstract_stagedOrder[plantabstract]; ok {
		return order
	}
	if order, ok := stage.PlantAbstracts_referenceOrder[plantabstract]; ok {
		return order
	} else {
		log.Printf("instance %p of type PlantAbstract was not staged and does not have a reference order", plantabstract)
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

func (pointsandlines3dshape *PointsAndLines3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PointsAndLines3DShape_stagedOrder[pointsandlines3dshape]; ok {
		return order
	}
	if order, ok := stage.PointsAndLines3DShapes_referenceOrder[pointsandlines3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PointsAndLines3DShape was not staged and does not have a reference order", pointsandlines3dshape)
		return 0
	}
}

func (pxshape *PxShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PxShape_stagedOrder[pxshape]; ok {
		return order
	}
	if order, ok := stage.PxShapes_referenceOrder[pxshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PxShape was not staged and does not have a reference order", pxshape)
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

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RotatedSampledPoints3DShape_stagedOrder[rotatedsampledpoints3dshape]; ok {
		return order
	}
	if order, ok := stage.RotatedSampledPoints3DShapes_referenceOrder[rotatedsampledpoints3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type RotatedSampledPoints3DShape was not staged and does not have a reference order", rotatedsampledpoints3dshape)
		return 0
	}
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RotatedSeatAndLegs3DShape_stagedOrder[rotatedseatandlegs3dshape]; ok {
		return order
	}
	if order, ok := stage.RotatedSeatAndLegs3DShapes_referenceOrder[rotatedseatandlegs3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type RotatedSeatAndLegs3DShape was not staged and does not have a reference order", rotatedseatandlegs3dshape)
		return 0
	}
}

func (sampledpoints3dshape *SampledPoints3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SampledPoints3DShape_stagedOrder[sampledpoints3dshape]; ok {
		return order
	}
	if order, ok := stage.SampledPoints3DShapes_referenceOrder[sampledpoints3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type SampledPoints3DShape was not staged and does not have a reference order", sampledpoints3dshape)
		return 0
	}
}

func (seat3dshape *Seat3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Seat3DShape_stagedOrder[seat3dshape]; ok {
		return order
	}
	if order, ok := stage.Seat3DShapes_referenceOrder[seat3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type Seat3DShape was not staged and does not have a reference order", seat3dshape)
		return 0
	}
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SeatAndLegs3DShape_stagedOrder[seatandlegs3dshape]; ok {
		return order
	}
	if order, ok := stage.SeatAndLegs3DShapes_referenceOrder[seatandlegs3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type SeatAndLegs3DShape was not staged and does not have a reference order", seatandlegs3dshape)
		return 0
	}
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SeatBottomCurveShape_stagedOrder[seatbottomcurveshape]; ok {
		return order
	}
	if order, ok := stage.SeatBottomCurveShapes_referenceOrder[seatbottomcurveshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type SeatBottomCurveShape was not staged and does not have a reference order", seatbottomcurveshape)
		return 0
	}
}

func (seattopcurveshape *SeatTopCurveShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SeatTopCurveShape_stagedOrder[seattopcurveshape]; ok {
		return order
	}
	if order, ok := stage.SeatTopCurveShapes_referenceOrder[seattopcurveshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type SeatTopCurveShape was not staged and does not have a reference order", seattopcurveshape)
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

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedLeftGrowthCurve2DRibbon_stagedOrder[shiftedleftgrowthcurve2dribbon]; ok {
		return order
	}
	if order, ok := stage.ShiftedLeftGrowthCurve2DRibbons_referenceOrder[shiftedleftgrowthcurve2dribbon]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedLeftGrowthCurve2DRibbon was not staged and does not have a reference order", shiftedleftgrowthcurve2dribbon)
		return 0
	}
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedLeftGrowthCurve2DRibbonEndShape_stagedOrder[shiftedleftgrowthcurve2dribbonendshape]; ok {
		return order
	}
	if order, ok := stage.ShiftedLeftGrowthCurve2DRibbonEndShapes_referenceOrder[shiftedleftgrowthcurve2dribbonendshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedLeftGrowthCurve2DRibbonEndShape was not staged and does not have a reference order", shiftedleftgrowthcurve2dribbonendshape)
		return 0
	}
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedLeftGrowthCurve2DRibbonStartShape_stagedOrder[shiftedleftgrowthcurve2dribbonstartshape]; ok {
		return order
	}
	if order, ok := stage.ShiftedLeftGrowthCurve2DRibbonStartShapes_referenceOrder[shiftedleftgrowthcurve2dribbonstartshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedLeftGrowthCurve2DRibbonStartShape was not staged and does not have a reference order", shiftedleftgrowthcurve2dribbonstartshape)
		return 0
	}
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedLeftPartiallyGrowthCurve2DRibbon_stagedOrder[shiftedleftpartiallygrowthcurve2dribbon]; ok {
		return order
	}
	if order, ok := stage.ShiftedLeftPartiallyGrowthCurve2DRibbons_referenceOrder[shiftedleftpartiallygrowthcurve2dribbon]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedLeftPartiallyGrowthCurve2DRibbon was not staged and does not have a reference order", shiftedleftpartiallygrowthcurve2dribbon)
		return 0
	}
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape_stagedOrder[shiftedleftpartiallygrowthcurve2dribbonendshape]; ok {
		return order
	}
	if order, ok := stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes_referenceOrder[shiftedleftpartiallygrowthcurve2dribbonendshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape was not staged and does not have a reference order", shiftedleftpartiallygrowthcurve2dribbonendshape)
		return 0
	}
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape_stagedOrder[shiftedleftpartiallygrowthcurve2dribbonstartshape]; ok {
		return order
	}
	if order, ok := stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes_referenceOrder[shiftedleftpartiallygrowthcurve2dribbonstartshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape was not staged and does not have a reference order", shiftedleftpartiallygrowthcurve2dribbonstartshape)
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

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedRightGrowthCurve2DRibbon_stagedOrder[shiftedrightgrowthcurve2dribbon]; ok {
		return order
	}
	if order, ok := stage.ShiftedRightGrowthCurve2DRibbons_referenceOrder[shiftedrightgrowthcurve2dribbon]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedRightGrowthCurve2DRibbon was not staged and does not have a reference order", shiftedrightgrowthcurve2dribbon)
		return 0
	}
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedRightGrowthCurve2DRibbonEndShape_stagedOrder[shiftedrightgrowthcurve2dribbonendshape]; ok {
		return order
	}
	if order, ok := stage.ShiftedRightGrowthCurve2DRibbonEndShapes_referenceOrder[shiftedrightgrowthcurve2dribbonendshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedRightGrowthCurve2DRibbonEndShape was not staged and does not have a reference order", shiftedrightgrowthcurve2dribbonendshape)
		return 0
	}
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ShiftedRightGrowthCurve2DRibbonStartShape_stagedOrder[shiftedrightgrowthcurve2dribbonstartshape]; ok {
		return order
	}
	if order, ok := stage.ShiftedRightGrowthCurve2DRibbonStartShapes_referenceOrder[shiftedrightgrowthcurve2dribbonstartshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ShiftedRightGrowthCurve2DRibbonStartShape was not staged and does not have a reference order", shiftedrightgrowthcurve2dribbonstartshape)
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

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackOfGrowthCurve2DByGrowthVector_stagedOrder[stackofgrowthcurve2dbygrowthvector]; ok {
		return order
	}
	if order, ok := stage.StackOfGrowthCurve2DByGrowthVectors_referenceOrder[stackofgrowthcurve2dbygrowthvector]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackOfGrowthCurve2DByGrowthVector was not staged and does not have a reference order", stackofgrowthcurve2dbygrowthvector)
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

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackOfPartiallyRotatedTorusShape_stagedOrder[stackofpartiallyrotatedtorusshape]; ok {
		return order
	}
	if order, ok := stage.StackOfPartiallyRotatedTorusShapes_referenceOrder[stackofpartiallyrotatedtorusshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackOfPartiallyRotatedTorusShape was not staged and does not have a reference order", stackofpartiallyrotatedtorusshape)
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

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackOfRotatedVaseTrapezeRingsShape_stagedOrder[stackofrotatedvasetrapezeringsshape]; ok {
		return order
	}
	if order, ok := stage.StackOfRotatedVaseTrapezeRingsShapes_referenceOrder[stackofrotatedvasetrapezeringsshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackOfRotatedVaseTrapezeRingsShape was not staged and does not have a reference order", stackofrotatedvasetrapezeringsshape)
		return 0
	}
}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StackOfVaseTrapezeRingsShape_stagedOrder[stackofvasetrapezeringsshape]; ok {
		return order
	}
	if order, ok := stage.StackOfVaseTrapezeRingsShapes_referenceOrder[stackofvasetrapezeringsshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StackOfVaseTrapezeRingsShape was not staged and does not have a reference order", stackofvasetrapezeringsshape)
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

func (stemcylinder3dshape *StemCylinder3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StemCylinder3DShape_stagedOrder[stemcylinder3dshape]; ok {
		return order
	}
	if order, ok := stage.StemCylinder3DShapes_referenceOrder[stemcylinder3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StemCylinder3DShape was not staged and does not have a reference order", stemcylinder3dshape)
		return 0
	}
}

func (stool2ddiagram *Stool2DDiagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Stool2DDiagram_stagedOrder[stool2ddiagram]; ok {
		return order
	}
	if order, ok := stage.Stool2DDiagrams_referenceOrder[stool2ddiagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Stool2DDiagram was not staged and does not have a reference order", stool2ddiagram)
		return 0
	}
}

func (stool3ddiagram *Stool3DDiagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Stool3DDiagram_stagedOrder[stool3ddiagram]; ok {
		return order
	}
	if order, ok := stage.Stool3DDiagrams_referenceOrder[stool3ddiagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Stool3DDiagram was not staged and does not have a reference order", stool3ddiagram)
		return 0
	}
}

func (tiledfloor3dshape *TiledFloor3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TiledFloor3DShape_stagedOrder[tiledfloor3dshape]; ok {
		return order
	}
	if order, ok := stage.TiledFloor3DShapes_referenceOrder[tiledfloor3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TiledFloor3DShape was not staged and does not have a reference order", tiledfloor3dshape)
		return 0
	}
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopCurvePlane1Shape_stagedOrder[topcurveplane1shape]; ok {
		return order
	}
	if order, ok := stage.TopCurvePlane1Shapes_referenceOrder[topcurveplane1shape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopCurvePlane1Shape was not staged and does not have a reference order", topcurveplane1shape)
		return 0
	}
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TopCurvePlane2Shape_stagedOrder[topcurveplane2shape]; ok {
		return order
	}
	if order, ok := stage.TopCurvePlane2Shapes_referenceOrder[topcurveplane2shape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TopCurvePlane2Shape was not staged and does not have a reference order", topcurveplane2shape)
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

func (torus3dshape *Torus3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Torus3DShape_stagedOrder[torus3dshape]; ok {
		return order
	}
	if order, ok := stage.Torus3DShapes_referenceOrder[torus3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type Torus3DShape was not staged and does not have a reference order", torus3dshape)
		return 0
	}
}

func (torusedge3dshape *TorusEdge3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TorusEdge3DShape_stagedOrder[torusedge3dshape]; ok {
		return order
	}
	if order, ok := stage.TorusEdge3DShapes_referenceOrder[torusedge3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TorusEdge3DShape was not staged and does not have a reference order", torusedge3dshape)
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

func (tubevase3ddiagram *TubeVase3DDiagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TubeVase3DDiagram_stagedOrder[tubevase3ddiagram]; ok {
		return order
	}
	if order, ok := stage.TubeVase3DDiagrams_referenceOrder[tubevase3ddiagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type TubeVase3DDiagram was not staged and does not have a reference order", tubevase3ddiagram)
		return 0
	}
}

func (tubevaseabstract *TubeVaseAbstract) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TubeVaseAbstract_stagedOrder[tubevaseabstract]; ok {
		return order
	}
	if order, ok := stage.TubeVaseAbstracts_referenceOrder[tubevaseabstract]; ok {
		return order
	} else {
		log.Printf("instance %p of type TubeVaseAbstract was not staged and does not have a reference order", tubevaseabstract)
		return 0
	}
}

func (vase2ddiagram *Vase2DDiagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Vase2DDiagram_stagedOrder[vase2ddiagram]; ok {
		return order
	}
	if order, ok := stage.Vase2DDiagrams_referenceOrder[vase2ddiagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Vase2DDiagram was not staged and does not have a reference order", vase2ddiagram)
		return 0
	}
}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.VaseTrapezeRingShape_stagedOrder[vasetrapezeringshape]; ok {
		return order
	}
	if order, ok := stage.VaseTrapezeRingShapes_referenceOrder[vasetrapezeringshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type VaseTrapezeRingShape was not staged and does not have a reference order", vasetrapezeringshape)
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

func (volumekey3dshape *VolumeKey3DShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.VolumeKey3DShape_stagedOrder[volumekey3dshape]; ok {
		return order
	}
	if order, ok := stage.VolumeKey3DShapes_referenceOrder[volumekey3dshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type VolumeKey3DShape was not staged and does not have a reference order", volumekey3dshape)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (angle0shape *Angle0Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", angle0shape.GongGetGongstructName(), angle0shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (angle0shape *Angle0Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", angle0shape.GongGetGongstructName(), angle0shape.GongGetOrder(stage))
}

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

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomcurveplane1shape.GongGetGongstructName(), bottomcurveplane1shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomcurveplane1shape.GongGetGongstructName(), bottomcurveplane1shape.GongGetOrder(stage))
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomcurveplane2shape.GongGetGongstructName(), bottomcurveplane2shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bottomcurveplane2shape.GongGetGongstructName(), bottomcurveplane2shape.GongGetOrder(stage))
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", chosenp1p2pairshape.GongGetGongstructName(), chosenp1p2pairshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (chosenp1p2pairshape *ChosenP1P2PairShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", chosenp1p2pairshape.GongGetGongstructName(), chosenp1p2pairshape.GongGetOrder(stage))
}

func (circlegridshape *CircleGridShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circlegridshape.GongGetGongstructName(), circlegridshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circlegridshape *CircleGridShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circlegridshape.GongGetGongstructName(), circlegridshape.GongGetOrder(stage))
}

func (circumference3dshape *Circumference3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circumference3dshape.GongGetGongstructName(), circumference3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circumference3dshape *Circumference3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", circumference3dshape.GongGetGongstructName(), circumference3dshape.GongGetOrder(stage))
}

func (clock2ddiagram *Clock2DDiagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", clock2ddiagram.GongGetGongstructName(), clock2ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (clock2ddiagram *Clock2DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", clock2ddiagram.GongGetGongstructName(), clock2ddiagram.GongGetOrder(stage))
}

func (clock3ddiagram *Clock3DDiagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", clock3ddiagram.GongGetGongstructName(), clock3ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (clock3ddiagram *Clock3DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", clock3ddiagram.GongGetGongstructName(), clock3ddiagram.GongGetOrder(stage))
}

func (clocktopcurveshape *ClockTopCurveShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", clocktopcurveshape.GongGetGongstructName(), clocktopcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (clocktopcurveshape *ClockTopCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", clocktopcurveshape.GongGetGongstructName(), clocktopcurveshape.GongGetOrder(stage))
}

func (cutline3dshape *CutLine3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cutline3dshape.GongGetGongstructName(), cutline3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cutline3dshape *CutLine3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cutline3dshape.GongGetGongstructName(), cutline3dshape.GongGetOrder(stage))
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

func (eye3dshape *Eye3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", eye3dshape.GongGetGongstructName(), eye3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (eye3dshape *Eye3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", eye3dshape.GongGetGongstructName(), eye3dshape.GongGetOrder(stage))
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", eyecornerssampledpoints3dshape.GongGetGongstructName(), eyecornerssampledpoints3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", eyecornerssampledpoints3dshape.GongGetGongstructName(), eyecornerssampledpoints3dshape.GongGetOrder(stage))
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", eyesampledpoints3dshape.GongGetGongstructName(), eyesampledpoints3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", eyesampledpoints3dshape.GongGetGongstructName(), eyesampledpoints3dshape.GongGetOrder(stage))
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", eyeseatbottomcurveshape.GongGetGongstructName(), eyeseatbottomcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", eyeseatbottomcurveshape.GongGetGongstructName(), eyeseatbottomcurveshape.GongGetOrder(stage))
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", eyestoolbottomcurveshape.GongGetGongstructName(), eyestoolbottomcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", eyestoolbottomcurveshape.GongGetGongstructName(), eyestoolbottomcurveshape.GongGetOrder(stage))
}

func (eyevolume3dshape *EyeVolume3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", eyevolume3dshape.GongGetGongstructName(), eyevolume3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (eyevolume3dshape *EyeVolume3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", eyevolume3dshape.GongGetGongstructName(), eyevolume3dshape.GongGetOrder(stage))
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

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurve2dribbon.GongGetGongstructName(), growthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurve2dribbon *GrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurve2dribbon.GongGetGongstructName(), growthcurve2dribbon.GongGetOrder(stage))
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurve2dribbonendshape.GongGetGongstructName(), growthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurve2dribbonendshape.GongGetGongstructName(), growthcurve2dribbonendshape.GongGetOrder(stage))
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurve2dribbonstartshape.GongGetGongstructName(), growthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", growthcurve2dribbonstartshape.GongGetGongstructName(), growthcurve2dribbonstartshape.GongGetOrder(stage))
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

func (key3dshape *Key3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", key3dshape.GongGetGongstructName(), key3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (key3dshape *Key3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", key3dshape.GongGetGongstructName(), key3dshape.GongGetOrder(stage))
}

func (keyhole3dshape *KeyHole3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", keyhole3dshape.GongGetGongstructName(), keyhole3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (keyhole3dshape *KeyHole3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", keyhole3dshape.GongGetGongstructName(), keyhole3dshape.GongGetOrder(stage))
}

func (keyholeshape *KeyHoleShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", keyholeshape.GongGetGongstructName(), keyholeshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (keyholeshape *KeyHoleShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", keyholeshape.GongGetGongstructName(), keyholeshape.GongGetOrder(stage))
}

func (leaves3dshape *Leaves3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", leaves3dshape.GongGetGongstructName(), leaves3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (leaves3dshape *Leaves3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", leaves3dshape.GongGetGongstructName(), leaves3dshape.GongGetOrder(stage))
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

func (originalpoints3dshape *OriginalPoints3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", originalpoints3dshape.GongGetGongstructName(), originalpoints3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (originalpoints3dshape *OriginalPoints3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", originalpoints3dshape.GongGetGongstructName(), originalpoints3dshape.GongGetOrder(stage))
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parastichymcurves3dshape.GongGetGongstructName(), parastichymcurves3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parastichymcurves3dshape.GongGetGongstructName(), parastichymcurves3dshape.GongGetOrder(stage))
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parastichyncurves3dshape.GongGetGongstructName(), parastichyncurves3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", parastichyncurves3dshape.GongGetGongstructName(), parastichyncurves3dshape.GongGetOrder(stage))
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

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectory.GongGetGongstructName(), partiallygrowthcurve2dtrajectory.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectory.GongGetGongstructName(), partiallygrowthcurve2dtrajectory.GongGetOrder(stage))
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryp1curveshape.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryp1curveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryp1curveshape.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryp1curveshape.GongGetOrder(stage))
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryp1p2.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryp1p2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryp1p2.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryp1p2.GongGetOrder(stage))
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongGetOrder(stage))
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryp1pointshape.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryp1pointshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryp1pointshape.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryp1pointshape.GongGetOrder(stage))
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryp2curveshape.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryp2curveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryp2curveshape.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryp2curveshape.GongGetOrder(stage))
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryp2pointshape.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryp2pointshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryp2pointshape.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryp2pointshape.GongGetOrder(stage))
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryshape.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallygrowthcurve2dtrajectoryshape.GongGetGongstructName(), partiallygrowthcurve2dtrajectoryshape.GongGetOrder(stage))
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallyrotatedseatbottomcurveshape.GongGetGongstructName(), partiallyrotatedseatbottomcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallyrotatedseatbottomcurveshape.GongGetGongstructName(), partiallyrotatedseatbottomcurveshape.GongGetOrder(stage))
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallyrotatedseattopcurveshape.GongGetGongstructName(), partiallyrotatedseattopcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallyrotatedseattopcurveshape.GongGetGongstructName(), partiallyrotatedseattopcurveshape.GongGetOrder(stage))
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallyrotatedtorusshape.GongGetGongstructName(), partiallyrotatedtorusshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partiallyrotatedtorusshape.GongGetGongstructName(), partiallyrotatedtorusshape.GongGetOrder(stage))
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

func (plant2ddiagram *Plant2DDiagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plant2ddiagram.GongGetGongstructName(), plant2ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (plant2ddiagram *Plant2DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plant2ddiagram.GongGetGongstructName(), plant2ddiagram.GongGetOrder(stage))
}

func (plant3ddiagram *Plant3DDiagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plant3ddiagram.GongGetGongstructName(), plant3ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (plant3ddiagram *Plant3DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plant3ddiagram.GongGetGongstructName(), plant3ddiagram.GongGetOrder(stage))
}

func (plantabstract *PlantAbstract) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plantabstract.GongGetGongstructName(), plantabstract.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (plantabstract *PlantAbstract) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plantabstract.GongGetGongstructName(), plantabstract.GongGetOrder(stage))
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plantcircumferenceshape.GongGetGongstructName(), plantcircumferenceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (plantcircumferenceshape *PlantCircumferenceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plantcircumferenceshape.GongGetGongstructName(), plantcircumferenceshape.GongGetOrder(stage))
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pointsandlines3dshape.GongGetGongstructName(), pointsandlines3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointsandlines3dshape *PointsAndLines3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pointsandlines3dshape.GongGetGongstructName(), pointsandlines3dshape.GongGetOrder(stage))
}

func (pxshape *PxShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pxshape.GongGetGongstructName(), pxshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pxshape *PxShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pxshape.GongGetGongstructName(), pxshape.GongGetOrder(stage))
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

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rotatedsampledpoints3dshape.GongGetGongstructName(), rotatedsampledpoints3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rotatedsampledpoints3dshape.GongGetGongstructName(), rotatedsampledpoints3dshape.GongGetOrder(stage))
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rotatedseatandlegs3dshape.GongGetGongstructName(), rotatedseatandlegs3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rotatedseatandlegs3dshape.GongGetGongstructName(), rotatedseatandlegs3dshape.GongGetOrder(stage))
}

func (sampledpoints3dshape *SampledPoints3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sampledpoints3dshape.GongGetGongstructName(), sampledpoints3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sampledpoints3dshape *SampledPoints3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sampledpoints3dshape.GongGetGongstructName(), sampledpoints3dshape.GongGetOrder(stage))
}

func (seat3dshape *Seat3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", seat3dshape.GongGetGongstructName(), seat3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (seat3dshape *Seat3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", seat3dshape.GongGetGongstructName(), seat3dshape.GongGetOrder(stage))
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", seatandlegs3dshape.GongGetGongstructName(), seatandlegs3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (seatandlegs3dshape *SeatAndLegs3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", seatandlegs3dshape.GongGetGongstructName(), seatandlegs3dshape.GongGetOrder(stage))
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", seatbottomcurveshape.GongGetGongstructName(), seatbottomcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (seatbottomcurveshape *SeatBottomCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", seatbottomcurveshape.GongGetGongstructName(), seatbottomcurveshape.GongGetOrder(stage))
}

func (seattopcurveshape *SeatTopCurveShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", seattopcurveshape.GongGetGongstructName(), seattopcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (seattopcurveshape *SeatTopCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", seattopcurveshape.GongGetGongstructName(), seattopcurveshape.GongGetOrder(stage))
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

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftgrowthcurve2dribbon.GongGetGongstructName(), shiftedleftgrowthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftgrowthcurve2dribbon.GongGetGongstructName(), shiftedleftgrowthcurve2dribbon.GongGetOrder(stage))
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftgrowthcurve2dribbonendshape.GongGetGongstructName(), shiftedleftgrowthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftgrowthcurve2dribbonendshape.GongGetGongstructName(), shiftedleftgrowthcurve2dribbonendshape.GongGetOrder(stage))
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftgrowthcurve2dribbonstartshape.GongGetGongstructName(), shiftedleftgrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftgrowthcurve2dribbonstartshape.GongGetGongstructName(), shiftedleftgrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftpartiallygrowthcurve2dribbon.GongGetGongstructName(), shiftedleftpartiallygrowthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftpartiallygrowthcurve2dribbon.GongGetGongstructName(), shiftedleftpartiallygrowthcurve2dribbon.GongGetOrder(stage))
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftpartiallygrowthcurve2dribbonendshape.GongGetGongstructName(), shiftedleftpartiallygrowthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftpartiallygrowthcurve2dribbonendshape.GongGetGongstructName(), shiftedleftpartiallygrowthcurve2dribbonendshape.GongGetOrder(stage))
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftpartiallygrowthcurve2dribbonstartshape.GongGetGongstructName(), shiftedleftpartiallygrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedleftpartiallygrowthcurve2dribbonstartshape.GongGetGongstructName(), shiftedleftpartiallygrowthcurve2dribbonstartshape.GongGetOrder(stage))
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

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedrightgrowthcurve2dribbon.GongGetGongstructName(), shiftedrightgrowthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedrightgrowthcurve2dribbon.GongGetGongstructName(), shiftedrightgrowthcurve2dribbon.GongGetOrder(stage))
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedrightgrowthcurve2dribbonendshape.GongGetGongstructName(), shiftedrightgrowthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedrightgrowthcurve2dribbonendshape.GongGetGongstructName(), shiftedrightgrowthcurve2dribbonendshape.GongGetOrder(stage))
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedrightgrowthcurve2dribbonstartshape.GongGetGongstructName(), shiftedrightgrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shiftedrightgrowthcurve2dribbonstartshape.GongGetGongstructName(), shiftedrightgrowthcurve2dribbonstartshape.GongGetOrder(stage))
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

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurve2dbygrowthvector.GongGetGongstructName(), stackofgrowthcurve2dbygrowthvector.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurve2dbygrowthvector.GongGetGongstructName(), stackofgrowthcurve2dbygrowthvector.GongGetOrder(stage))
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurve2dribbon.GongGetGongstructName(), stackofgrowthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofgrowthcurve2dribbon.GongGetGongstructName(), stackofgrowthcurve2dribbon.GongGetOrder(stage))
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofpartiallyrotatedtorusshape.GongGetGongstructName(), stackofpartiallyrotatedtorusshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofpartiallyrotatedtorusshape.GongGetGongstructName(), stackofpartiallyrotatedtorusshape.GongGetOrder(stage))
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

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofrotatedvasetrapezeringsshape.GongGetGongstructName(), stackofrotatedvasetrapezeringsshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofrotatedvasetrapezeringsshape.GongGetGongstructName(), stackofrotatedvasetrapezeringsshape.GongGetOrder(stage))
}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofvasetrapezeringsshape.GongGetGongstructName(), stackofvasetrapezeringsshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stackofvasetrapezeringsshape.GongGetGongstructName(), stackofvasetrapezeringsshape.GongGetOrder(stage))
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

func (stemcylinder3dshape *StemCylinder3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stemcylinder3dshape.GongGetGongstructName(), stemcylinder3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stemcylinder3dshape *StemCylinder3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stemcylinder3dshape.GongGetGongstructName(), stemcylinder3dshape.GongGetOrder(stage))
}

func (stool2ddiagram *Stool2DDiagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stool2ddiagram.GongGetGongstructName(), stool2ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stool2ddiagram *Stool2DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stool2ddiagram.GongGetGongstructName(), stool2ddiagram.GongGetOrder(stage))
}

func (stool3ddiagram *Stool3DDiagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stool3ddiagram.GongGetGongstructName(), stool3ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stool3ddiagram *Stool3DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stool3ddiagram.GongGetGongstructName(), stool3ddiagram.GongGetOrder(stage))
}

func (tiledfloor3dshape *TiledFloor3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tiledfloor3dshape.GongGetGongstructName(), tiledfloor3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tiledfloor3dshape *TiledFloor3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tiledfloor3dshape.GongGetGongstructName(), tiledfloor3dshape.GongGetOrder(stage))
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topcurveplane1shape.GongGetGongstructName(), topcurveplane1shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topcurveplane1shape *TopCurvePlane1Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topcurveplane1shape.GongGetGongstructName(), topcurveplane1shape.GongGetOrder(stage))
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topcurveplane2shape.GongGetGongstructName(), topcurveplane2shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topcurveplane2shape *TopCurvePlane2Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", topcurveplane2shape.GongGetGongstructName(), topcurveplane2shape.GongGetOrder(stage))
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

func (torus3dshape *Torus3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", torus3dshape.GongGetGongstructName(), torus3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (torus3dshape *Torus3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", torus3dshape.GongGetGongstructName(), torus3dshape.GongGetOrder(stage))
}

func (torusedge3dshape *TorusEdge3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", torusedge3dshape.GongGetGongstructName(), torusedge3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (torusedge3dshape *TorusEdge3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", torusedge3dshape.GongGetGongstructName(), torusedge3dshape.GongGetOrder(stage))
}

func (torusstackshape *TorusStackShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", torusstackshape.GongGetGongstructName(), torusstackshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (torusstackshape *TorusStackShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", torusstackshape.GongGetGongstructName(), torusstackshape.GongGetOrder(stage))
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tubevase3ddiagram.GongGetGongstructName(), tubevase3ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tubevase3ddiagram *TubeVase3DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tubevase3ddiagram.GongGetGongstructName(), tubevase3ddiagram.GongGetOrder(stage))
}

func (tubevaseabstract *TubeVaseAbstract) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tubevaseabstract.GongGetGongstructName(), tubevaseabstract.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tubevaseabstract *TubeVaseAbstract) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tubevaseabstract.GongGetGongstructName(), tubevaseabstract.GongGetOrder(stage))
}

func (vase2ddiagram *Vase2DDiagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", vase2ddiagram.GongGetGongstructName(), vase2ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vase2ddiagram *Vase2DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", vase2ddiagram.GongGetGongstructName(), vase2ddiagram.GongGetOrder(stage))
}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", vasetrapezeringshape.GongGetGongstructName(), vasetrapezeringshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vasetrapezeringshape *VaseTrapezeRingShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", vasetrapezeringshape.GongGetGongstructName(), vasetrapezeringshape.GongGetOrder(stage))
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", verticaltorusstackshape.GongGetGongstructName(), verticaltorusstackshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticaltorusstackshape *VerticalTorusStackShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", verticaltorusstackshape.GongGetGongstructName(), verticaltorusstackshape.GongGetOrder(stage))
}

func (volumekey3dshape *VolumeKey3DShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", volumekey3dshape.GongGetGongstructName(), volumekey3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (volumekey3dshape *VolumeKey3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", volumekey3dshape.GongGetGongstructName(), volumekey3dshape.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (angle0shape *Angle0Shape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", angle0shape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Angle0Shape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(angle0shape.Name))
	return
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arcnormalvectorshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ArcNormalVectorShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(arcnormalvectorshape.Name))
	return
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arcnormalvectorshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ArcNormalVectorShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(arcnormalvectorshapegrid.Name))
	return
}

func (axesshape *AxesShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", axesshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AxesShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(axesshape.Name))
	return
}

func (basevectorshape *BaseVectorShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", basevectorshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BaseVectorShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(basevectorshape.Name))
	return
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", basevectorshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BaseVectorShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(basevectorshapegrid.Name))
	return
}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomcurveplane1shape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BottomCurvePlane1Shape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(bottomcurveplane1shape.Name))
	return
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomcurveplane2shape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BottomCurvePlane2Shape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(bottomcurveplane2shape.Name))
	return
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", chosenp1p2pairshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ChosenP1P2PairShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(chosenp1p2pairshape.Name))
	return
}

func (circlegridshape *CircleGridShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circlegridshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CircleGridShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(circlegridshape.Name))
	return
}

func (circumference3dshape *Circumference3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circumference3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Circumference3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(circumference3dshape.Name))
	return
}

func (clock2ddiagram *Clock2DDiagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", clock2ddiagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Clock2DDiagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(clock2ddiagram.Name))
	return
}

func (clock3ddiagram *Clock3DDiagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", clock3ddiagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Clock3DDiagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(clock3ddiagram.Name))
	return
}

func (clocktopcurveshape *ClockTopCurveShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", clocktopcurveshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ClockTopCurveShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(clocktopcurveshape.Name))
	return
}

func (cutline3dshape *CutLine3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cutline3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CutLine3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(cutline3dshape.Name))
	return
}

func (endarcshape *EndArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EndArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(endarcshape.Name))
	return
}

func (endarcshapegrid *EndArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EndArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(endarcshapegrid.Name))
	return
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endhalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EndHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(endhalfwayarcshape.Name))
	return
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", endhalfwayarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EndHalfwayArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(endhalfwayarcshapegrid.Name))
	return
}

func (explanationtextshape *ExplanationTextShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", explanationtextshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ExplanationTextShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(explanationtextshape.Name))
	return
}

func (eye3dshape *Eye3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", eye3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Eye3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(eye3dshape.Name))
	return
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", eyecornerssampledpoints3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EyeCornersSampledPoints3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(eyecornerssampledpoints3dshape.Name))
	return
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", eyesampledpoints3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EyeSampledPoints3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(eyesampledpoints3dshape.Name))
	return
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", eyeseatbottomcurveshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EyeSeatBottomCurveShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(eyeseatbottomcurveshape.Name))
	return
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", eyestoolbottomcurveshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EyeStoolBottomCurveShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(eyestoolbottomcurveshape.Name))
	return
}

func (eyevolume3dshape *EyeVolume3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", eyevolume3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EyeVolume3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(eyevolume3dshape.Name))
	return
}

func (gridpathshape *GridPathShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gridpathshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GridPathShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(gridpathshape.Name))
	return
}

func (growthcurve2d *GrowthCurve2D) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurve2d.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GrowthCurve2D")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(growthcurve2d.Name))
	return
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurve2dribbon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GrowthCurve2DRibbon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(growthcurve2dribbon.Name))
	return
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurve2dribbonendshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GrowthCurve2DRibbonEndShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(growthcurve2dribbonendshape.Name))
	return
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurve2dribbonstartshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GrowthCurve2DRibbonStartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(growthcurve2dribbonstartshape.Name))
	return
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurverhombusgridshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GrowthCurveRhombusGridShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(growthcurverhombusgridshape.Name))
	return
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurverhombusshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GrowthCurveRhombusShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(growthcurverhombusshape.Name))
	return
}

func (growthvectorshape *GrowthVectorShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthvectorshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GrowthVectorShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(growthvectorshape.Name))
	return
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", initialrhombusgridshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "InitialRhombusGridShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(initialrhombusgridshape.Name))
	return
}

func (initialrhombusshape *InitialRhombusShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", initialrhombusshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "InitialRhombusShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(initialrhombusshape.Name))
	return
}

func (key3dshape *Key3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", key3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Key3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(key3dshape.Name))
	return
}

func (keyhole3dshape *KeyHole3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", keyhole3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "KeyHole3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(keyhole3dshape.Name))
	return
}

func (keyholeshape *KeyHoleShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", keyholeshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "KeyHoleShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(keyholeshape.Name))
	return
}

func (leaves3dshape *Leaves3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", leaves3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Leaves3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(leaves3dshape.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(library.Name))
	return
}

func (midarcvectorshape *MidArcVectorShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", midarcvectorshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MidArcVectorShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(midarcvectorshape.Name))
	return
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", midarcvectorshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MidArcVectorShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(midarcvectorshapegrid.Name))
	return
}

func (originalpoints3dshape *OriginalPoints3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", originalpoints3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "OriginalPoints3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(originalpoints3dshape.Name))
	return
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parastichymcurves3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParastichyMCurves3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(parastichymcurves3dshape.Name))
	return
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parastichyncurves3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParastichyNCurves3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(parastichyncurves3dshape.Name))
	return
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dribbon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DRibbon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallygrowthcurve2dribbon.Name))
	return
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dribbonendshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DRibbonEndShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallygrowthcurve2dribbonendshape.Name))
	return
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dribbonstartshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DRibbonStartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallygrowthcurve2dribbonstartshape.Name))
	return
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectory.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DTrajectory")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallygrowthcurve2dtrajectory.Name))
	return
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryp1curveshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DTrajectoryP1CurveShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallygrowthcurve2dtrajectoryp1curveshape.Name))
	return
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryp1p2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DTrajectoryP1P2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallygrowthcurve2dtrajectoryp1p2.Name))
	return
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallygrowthcurve2dtrajectoryp1p2pairlineshape.Name))
	return
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryp1pointshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DTrajectoryP1PointShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallygrowthcurve2dtrajectoryp1pointshape.Name))
	return
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryp2curveshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DTrajectoryP2CurveShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallygrowthcurve2dtrajectoryp2curveshape.Name))
	return
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryp2pointshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DTrajectoryP2PointShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallygrowthcurve2dtrajectoryp2pointshape.Name))
	return
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyGrowthCurve2DTrajectoryShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallygrowthcurve2dtrajectoryshape.Name))
	return
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallyrotatedseatbottomcurveshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyRotatedSeatBottomCurveShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallyrotatedseatbottomcurveshape.Name))
	return
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallyrotatedseattopcurveshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyRotatedSeatTopCurveShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallyrotatedseattopcurveshape.Name))
	return
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallyrotatedtorusshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartiallyRotatedTorusShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partiallyrotatedtorusshape.Name))
	return
}

func (perpendicularvector *PerpendicularVector) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", perpendicularvector.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PerpendicularVector")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(perpendicularvector.Name))
	return
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", perpendicularvectorgrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PerpendicularVectorGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(perpendicularvectorgrid.Name))
	return
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", perpendicularvectorgridhalfway.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PerpendicularVectorGridHalfway")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(perpendicularvectorgridhalfway.Name))
	return
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", perpendicularvectorhalfway.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PerpendicularVectorHalfway")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(perpendicularvectorhalfway.Name))
	return
}

func (plant2ddiagram *Plant2DDiagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plant2ddiagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Plant2DDiagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(plant2ddiagram.Name))
	return
}

func (plant3ddiagram *Plant3DDiagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plant3ddiagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Plant3DDiagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(plant3ddiagram.Name))
	return
}

func (plantabstract *PlantAbstract) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plantabstract.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PlantAbstract")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(plantabstract.Name))
	return
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plantcircumferenceshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PlantCircumferenceShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(plantcircumferenceshape.Name))
	return
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pointsandlines3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PointsAndLines3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(pointsandlines3dshape.Name))
	return
}

func (pxshape *PxShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pxshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PxShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(pxshape.Name))
	return
}

func (rendered3dshape *Rendered3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rendered3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Rendered3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rendered3dshape.Name))
	return
}

func (rhombusshape *RhombusShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rhombusshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RhombusShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rhombusshape.Name))
	return
}

func (rhombusstuff *RhombusStuff) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rhombusstuff.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RhombusStuff")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rhombusstuff.Name))
	return
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rotatedrhombusgridshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RotatedRhombusGridShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rotatedrhombusgridshape.Name))
	return
}

func (rotatedrhombusshape *RotatedRhombusShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rotatedrhombusshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RotatedRhombusShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rotatedrhombusshape.Name))
	return
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rotatedsampledpoints3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RotatedSampledPoints3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rotatedsampledpoints3dshape.Name))
	return
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rotatedseatandlegs3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RotatedSeatAndLegs3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rotatedseatandlegs3dshape.Name))
	return
}

func (sampledpoints3dshape *SampledPoints3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sampledpoints3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SampledPoints3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(sampledpoints3dshape.Name))
	return
}

func (seat3dshape *Seat3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", seat3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Seat3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(seat3dshape.Name))
	return
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", seatandlegs3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SeatAndLegs3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(seatandlegs3dshape.Name))
	return
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", seatbottomcurveshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SeatBottomCurveShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(seatbottomcurveshape.Name))
	return
}

func (seattopcurveshape *SeatTopCurveShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", seattopcurveshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SeatTopCurveShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(seattopcurveshape.Name))
	return
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedbottomtopstartarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedBottomTopStartArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedbottomtopstartarcshape.Name))
	return
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedbottomtopstartarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedBottomTopStartArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedbottomtopstartarcshapegrid.Name))
	return
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftgrowthcurve2dribbon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftGrowthCurve2DRibbon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedleftgrowthcurve2dribbon.Name))
	return
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftgrowthcurve2dribbonendshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftGrowthCurve2DRibbonEndShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedleftgrowthcurve2dribbonendshape.Name))
	return
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftgrowthcurve2dribbonstartshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftGrowthCurve2DRibbonStartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedleftgrowthcurve2dribbonstartshape.Name))
	return
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftpartiallygrowthcurve2dribbon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftPartiallyGrowthCurve2DRibbon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedleftpartiallygrowthcurve2dribbon.Name))
	return
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftpartiallygrowthcurve2dribbonendshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedleftpartiallygrowthcurve2dribbonendshape.Name))
	return
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftpartiallygrowthcurve2dribbonstartshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedleftpartiallygrowthcurve2dribbonstartshape.Name))
	return
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstackgrowthcurveendarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftStackGrowthCurveEndArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedleftstackgrowthcurveendarcshape.Name))
	return
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstackgrowthcurvestartarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftStackGrowthCurveStartArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedleftstackgrowthcurvestartarcshape.Name))
	return
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstacknormalvector.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftStackNormalVector")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedleftstacknormalvector.Name))
	return
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstackofgrowthcurve.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftStackOfGrowthCurve")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedleftstackofgrowthcurve.Name))
	return
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftstackofnormalvector.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedLeftStackOfNormalVector")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedleftstackofnormalvector.Name))
	return
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedrightgrowthcurve2dribbon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedRightGrowthCurve2DRibbon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedrightgrowthcurve2dribbon.Name))
	return
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedrightgrowthcurve2dribbonendshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedRightGrowthCurve2DRibbonEndShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedrightgrowthcurve2dribbonendshape.Name))
	return
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedrightgrowthcurve2dribbonstartshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ShiftedRightGrowthCurve2DRibbonStartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shiftedrightgrowthcurve2dribbonstartshape.Name))
	return
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurve2dendhalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackGrowthCurve2DEndHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackgrowthcurve2dendhalfwayarcshape.Name))
	return
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurve2dribbonendshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackGrowthCurve2DRibbonEndShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackgrowthcurve2dribbonendshape.Name))
	return
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurve2dribbonstartshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackGrowthCurve2DRibbonStartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackgrowthcurve2dribbonstartshape.Name))
	return
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackgrowthcurve2dstarthalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackGrowthCurve2DStartHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackgrowthcurve2dstarthalfwayarcshape.Name))
	return
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurve2d.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfGrowthCurve2D")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackofgrowthcurve2d.Name))
	return
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurve2dbygrowthvector.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfGrowthCurve2DByGrowthVector")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackofgrowthcurve2dbygrowthvector.Name))
	return
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurve2dribbon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfGrowthCurve2DRibbon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackofgrowthcurve2dribbon.Name))
	return
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofpartiallyrotatedtorusshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfPartiallyRotatedTorusShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackofpartiallyrotatedtorusshape.Name))
	return
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofrotatedgrowthcurve2d.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfRotatedGrowthCurve2D")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackofrotatedgrowthcurve2d.Name))
	return
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofrotatedgrowthcurve2dribbon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfRotatedGrowthCurve2DRibbon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackofrotatedgrowthcurve2dribbon.Name))
	return
}

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofrotatedvasetrapezeringsshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfRotatedVaseTrapezeRingsShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackofrotatedvasetrapezeringsshape.Name))
	return
}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofvasetrapezeringsshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackOfVaseTrapezeRingsShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackofvasetrapezeringsshape.Name))
	return
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackrotatedgrowthcurve2dendarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackRotatedGrowthCurve2DEndArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackrotatedgrowthcurve2dendarcshape.Name))
	return
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackrotatedgrowthcurve2dribbonendshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackRotatedGrowthCurve2DRibbonEndShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackrotatedgrowthcurve2dribbonendshape.Name))
	return
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackrotatedgrowthcurve2dribbonstartshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackRotatedGrowthCurve2DRibbonStartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackrotatedgrowthcurve2dribbonstartshape.Name))
	return
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackrotatedgrowthcurve2dstartarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StackRotatedGrowthCurve2DStartArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stackrotatedgrowthcurve2dstartarcshape.Name))
	return
}

func (startarcshape *StartArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", startarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StartArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(startarcshape.Name))
	return
}

func (startarcshapegrid *StartArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", startarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StartArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(startarcshapegrid.Name))
	return
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", starthalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StartHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(starthalfwayarcshape.Name))
	return
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", starthalfwayarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StartHalfwayArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(starthalfwayarcshapegrid.Name))
	return
}

func (stemcylinder3dshape *StemCylinder3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stemcylinder3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StemCylinder3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stemcylinder3dshape.Name))
	return
}

func (stool2ddiagram *Stool2DDiagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stool2ddiagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Stool2DDiagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stool2ddiagram.Name))
	return
}

func (stool3ddiagram *Stool3DDiagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stool3ddiagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Stool3DDiagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stool3ddiagram.Name))
	return
}

func (tiledfloor3dshape *TiledFloor3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tiledfloor3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TiledFloor3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tiledfloor3dshape.Name))
	return
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topcurveplane1shape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopCurvePlane1Shape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topcurveplane1shape.Name))
	return
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topcurveplane2shape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopCurvePlane2Shape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topcurveplane2shape.Name))
	return
}

func (topendarcshape *TopEndArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopEndArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topendarcshape.Name))
	return
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopEndArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topendarcshapegrid.Name))
	return
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendhalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopEndHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topendhalfwayarcshape.Name))
	return
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topendhalfwayarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopEndHalfwayArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topendhalfwayarcshapegrid.Name))
	return
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topgrowthcurve2d.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopGrowthCurve2D")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topgrowthcurve2d.Name))
	return
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topmidarcvectorshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopMidArcVectorShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topmidarcvectorshape.Name))
	return
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topmidarcvectorshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopMidArcVectorShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topmidarcvectorshapegrid.Name))
	return
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackgrowthcurve2dendhalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackGrowthCurve2DEndHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topstackgrowthcurve2dendhalfwayarcshape.Name))
	return
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackgrowthcurve2dstarthalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackGrowthCurve2DStartHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topstackgrowthcurve2dstarthalfwayarcshape.Name))
	return
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofgrowthcurve2d.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackOfGrowthCurve2D")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topstackofgrowthcurve2d.Name))
	return
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofrotatedgrowthcurve2d.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackOfRotatedGrowthCurve2D")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topstackofrotatedgrowthcurve2d.Name))
	return
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofrotatedgrowthcurve2dendarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackOfRotatedGrowthCurve2DEndArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topstackofrotatedgrowthcurve2dendarcshape.Name))
	return
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstackofrotatedgrowthcurve2dstartarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStackOfRotatedGrowthCurve2DStartArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topstackofrotatedgrowthcurve2dstartarcshape.Name))
	return
}

func (topstartarcshape *TopStartArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstartarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStartArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topstartarcshape.Name))
	return
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstartarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStartArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topstartarcshapegrid.Name))
	return
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstarthalfwayarcshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStartHalfwayArcShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topstarthalfwayarcshape.Name))
	return
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topstarthalfwayarcshapegrid.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TopStartHalfwayArcShapeGrid")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(topstarthalfwayarcshapegrid.Name))
	return
}

func (torus3dshape *Torus3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", torus3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Torus3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(torus3dshape.Name))
	return
}

func (torusedge3dshape *TorusEdge3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", torusedge3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TorusEdge3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(torusedge3dshape.Name))
	return
}

func (torusstackshape *TorusStackShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", torusstackshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TorusStackShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(torusstackshape.Name))
	return
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tubevase3ddiagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TubeVase3DDiagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tubevase3ddiagram.Name))
	return
}

func (tubevaseabstract *TubeVaseAbstract) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tubevaseabstract.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TubeVaseAbstract")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tubevaseabstract.Name))
	return
}

func (vase2ddiagram *Vase2DDiagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", vase2ddiagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Vase2DDiagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(vase2ddiagram.Name))
	return
}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", vasetrapezeringshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "VaseTrapezeRingShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(vasetrapezeringshape.Name))
	return
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", verticaltorusstackshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "VerticalTorusStackShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(verticaltorusstackshape.Name))
	return
}

func (volumekey3dshape *VolumeKey3DShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", volumekey3dshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "VolumeKey3DShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(volumekey3dshape.Name))
	return
}

// insertion point for unstaging
func (angle0shape *Angle0Shape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", angle0shape.GongGetReferenceIdentifier(stage))
	return
}

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

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomcurveplane1shape.GongGetReferenceIdentifier(stage))
	return
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bottomcurveplane2shape.GongGetReferenceIdentifier(stage))
	return
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", chosenp1p2pairshape.GongGetReferenceIdentifier(stage))
	return
}

func (circlegridshape *CircleGridShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circlegridshape.GongGetReferenceIdentifier(stage))
	return
}

func (circumference3dshape *Circumference3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", circumference3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (clock2ddiagram *Clock2DDiagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", clock2ddiagram.GongGetReferenceIdentifier(stage))
	return
}

func (clock3ddiagram *Clock3DDiagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", clock3ddiagram.GongGetReferenceIdentifier(stage))
	return
}

func (clocktopcurveshape *ClockTopCurveShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", clocktopcurveshape.GongGetReferenceIdentifier(stage))
	return
}

func (cutline3dshape *CutLine3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cutline3dshape.GongGetReferenceIdentifier(stage))
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

func (eye3dshape *Eye3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", eye3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", eyecornerssampledpoints3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", eyesampledpoints3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", eyeseatbottomcurveshape.GongGetReferenceIdentifier(stage))
	return
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", eyestoolbottomcurveshape.GongGetReferenceIdentifier(stage))
	return
}

func (eyevolume3dshape *EyeVolume3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", eyevolume3dshape.GongGetReferenceIdentifier(stage))
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

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurve2dribbon.GongGetReferenceIdentifier(stage))
	return
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
	return
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", growthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
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

func (key3dshape *Key3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", key3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (keyhole3dshape *KeyHole3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", keyhole3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (keyholeshape *KeyHoleShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", keyholeshape.GongGetReferenceIdentifier(stage))
	return
}

func (leaves3dshape *Leaves3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", leaves3dshape.GongGetReferenceIdentifier(stage))
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

func (originalpoints3dshape *OriginalPoints3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", originalpoints3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parastichymcurves3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", parastichyncurves3dshape.GongGetReferenceIdentifier(stage))
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

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectory.GongGetReferenceIdentifier(stage))
	return
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryp1curveshape.GongGetReferenceIdentifier(stage))
	return
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryp1p2.GongGetReferenceIdentifier(stage))
	return
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongGetReferenceIdentifier(stage))
	return
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryp1pointshape.GongGetReferenceIdentifier(stage))
	return
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryp2curveshape.GongGetReferenceIdentifier(stage))
	return
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryp2pointshape.GongGetReferenceIdentifier(stage))
	return
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallygrowthcurve2dtrajectoryshape.GongGetReferenceIdentifier(stage))
	return
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallyrotatedseatbottomcurveshape.GongGetReferenceIdentifier(stage))
	return
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallyrotatedseattopcurveshape.GongGetReferenceIdentifier(stage))
	return
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partiallyrotatedtorusshape.GongGetReferenceIdentifier(stage))
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

func (plant2ddiagram *Plant2DDiagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plant2ddiagram.GongGetReferenceIdentifier(stage))
	return
}

func (plant3ddiagram *Plant3DDiagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plant3ddiagram.GongGetReferenceIdentifier(stage))
	return
}

func (plantabstract *PlantAbstract) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plantabstract.GongGetReferenceIdentifier(stage))
	return
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plantcircumferenceshape.GongGetReferenceIdentifier(stage))
	return
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pointsandlines3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (pxshape *PxShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pxshape.GongGetReferenceIdentifier(stage))
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

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rotatedsampledpoints3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rotatedseatandlegs3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (sampledpoints3dshape *SampledPoints3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sampledpoints3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (seat3dshape *Seat3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", seat3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", seatandlegs3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", seatbottomcurveshape.GongGetReferenceIdentifier(stage))
	return
}

func (seattopcurveshape *SeatTopCurveShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", seattopcurveshape.GongGetReferenceIdentifier(stage))
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

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftgrowthcurve2dribbon.GongGetReferenceIdentifier(stage))
	return
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftgrowthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
	return
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftgrowthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
	return
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftpartiallygrowthcurve2dribbon.GongGetReferenceIdentifier(stage))
	return
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftpartiallygrowthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
	return
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedleftpartiallygrowthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
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

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedrightgrowthcurve2dribbon.GongGetReferenceIdentifier(stage))
	return
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedrightgrowthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
	return
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shiftedrightgrowthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
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

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurve2dbygrowthvector.GongGetReferenceIdentifier(stage))
	return
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofgrowthcurve2dribbon.GongGetReferenceIdentifier(stage))
	return
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofpartiallyrotatedtorusshape.GongGetReferenceIdentifier(stage))
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

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofrotatedvasetrapezeringsshape.GongGetReferenceIdentifier(stage))
	return
}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stackofvasetrapezeringsshape.GongGetReferenceIdentifier(stage))
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

func (stemcylinder3dshape *StemCylinder3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stemcylinder3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (stool2ddiagram *Stool2DDiagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stool2ddiagram.GongGetReferenceIdentifier(stage))
	return
}

func (stool3ddiagram *Stool3DDiagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stool3ddiagram.GongGetReferenceIdentifier(stage))
	return
}

func (tiledfloor3dshape *TiledFloor3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tiledfloor3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topcurveplane1shape.GongGetReferenceIdentifier(stage))
	return
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", topcurveplane2shape.GongGetReferenceIdentifier(stage))
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

func (torus3dshape *Torus3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", torus3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (torusedge3dshape *TorusEdge3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", torusedge3dshape.GongGetReferenceIdentifier(stage))
	return
}

func (torusstackshape *TorusStackShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", torusstackshape.GongGetReferenceIdentifier(stage))
	return
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tubevase3ddiagram.GongGetReferenceIdentifier(stage))
	return
}

func (tubevaseabstract *TubeVaseAbstract) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tubevaseabstract.GongGetReferenceIdentifier(stage))
	return
}

func (vase2ddiagram *Vase2DDiagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", vase2ddiagram.GongGetReferenceIdentifier(stage))
	return
}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", vasetrapezeringshape.GongGetReferenceIdentifier(stage))
	return
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", verticaltorusstackshape.GongGetReferenceIdentifier(stage))
	return
}

func (volumekey3dshape *VolumeKey3DShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", volumekey3dshape.GongGetReferenceIdentifier(stage))
	return
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
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
