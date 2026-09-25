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

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Angle0Shapes)

	res = __gong__appendInstances(res, stage.ArcNormalVectorShapes)

	res = __gong__appendInstances(res, stage.ArcNormalVectorShapeGrids)

	res = __gong__appendInstances(res, stage.AxesShapes)

	res = __gong__appendInstances(res, stage.BaseVectorShapes)

	res = __gong__appendInstances(res, stage.BaseVectorShapeGrids)

	res = __gong__appendInstances(res, stage.BottomCurvePlane1Shapes)

	res = __gong__appendInstances(res, stage.BottomCurvePlane2Shapes)

	res = __gong__appendInstances(res, stage.ChosenP1P2PairShapes)

	res = __gong__appendInstances(res, stage.CircleGridShapes)

	res = __gong__appendInstances(res, stage.Circumference3DShapes)

	res = __gong__appendInstances(res, stage.Clock2DDiagrams)

	res = __gong__appendInstances(res, stage.Clock3DDiagrams)

	res = __gong__appendInstances(res, stage.ClockTopCurveShapes)

	res = __gong__appendInstances(res, stage.CutLine3DShapes)

	res = __gong__appendInstances(res, stage.EndArcShapes)

	res = __gong__appendInstances(res, stage.EndArcShapeGrids)

	res = __gong__appendInstances(res, stage.EndHalfwayArcShapes)

	res = __gong__appendInstances(res, stage.EndHalfwayArcShapeGrids)

	res = __gong__appendInstances(res, stage.ExplanationTextShapes)

	res = __gong__appendInstances(res, stage.Eye3DShapes)

	res = __gong__appendInstances(res, stage.EyeCornersSampledPoints3DShapes)

	res = __gong__appendInstances(res, stage.EyeSampledPoints3DShapes)

	res = __gong__appendInstances(res, stage.EyeSeatBottomCurveShapes)

	res = __gong__appendInstances(res, stage.EyeStoolBottomCurveShapes)

	res = __gong__appendInstances(res, stage.EyeVolume3DShapes)

	res = __gong__appendInstances(res, stage.GridPathShapes)

	res = __gong__appendInstances(res, stage.GrowthCurve2Ds)

	res = __gong__appendInstances(res, stage.GrowthCurve2DRibbons)

	res = __gong__appendInstances(res, stage.GrowthCurve2DRibbonEndShapes)

	res = __gong__appendInstances(res, stage.GrowthCurve2DRibbonStartShapes)

	res = __gong__appendInstances(res, stage.GrowthCurveRhombusGridShapes)

	res = __gong__appendInstances(res, stage.GrowthCurveRhombusShapes)

	res = __gong__appendInstances(res, stage.GrowthVectorShapes)

	res = __gong__appendInstances(res, stage.InitialRhombusGridShapes)

	res = __gong__appendInstances(res, stage.InitialRhombusShapes)

	res = __gong__appendInstances(res, stage.Key3DShapes)

	res = __gong__appendInstances(res, stage.KeyHole3DShapes)

	res = __gong__appendInstances(res, stage.KeyHoleShapes)

	res = __gong__appendInstances(res, stage.Leaves3DShapes)

	res = __gong__appendInstances(res, stage.Librarys)

	res = __gong__appendInstances(res, stage.MidArcVectorShapes)

	res = __gong__appendInstances(res, stage.MidArcVectorShapeGrids)

	res = __gong__appendInstances(res, stage.OriginalPoints3DShapes)

	res = __gong__appendInstances(res, stage.ParastichyMCurves3DShapes)

	res = __gong__appendInstances(res, stage.ParastichyNCurves3DShapes)

	res = __gong__appendInstances(res, stage.PartiallyGrowthCurve2DRibbons)

	res = __gong__appendInstances(res, stage.PartiallyGrowthCurve2DRibbonEndShapes)

	res = __gong__appendInstances(res, stage.PartiallyGrowthCurve2DRibbonStartShapes)

	res = __gong__appendInstances(res, stage.PartiallyGrowthCurve2DTrajectorys)

	res = __gong__appendInstances(res, stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes)

	res = __gong__appendInstances(res, stage.PartiallyGrowthCurve2DTrajectoryP1P2s)

	res = __gong__appendInstances(res, stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes)

	res = __gong__appendInstances(res, stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes)

	res = __gong__appendInstances(res, stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes)

	res = __gong__appendInstances(res, stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes)

	res = __gong__appendInstances(res, stage.PartiallyGrowthCurve2DTrajectoryShapes)

	res = __gong__appendInstances(res, stage.PartiallyRotatedSeatBottomCurveShapes)

	res = __gong__appendInstances(res, stage.PartiallyRotatedSeatTopCurveShapes)

	res = __gong__appendInstances(res, stage.PartiallyRotatedTorusShapes)

	res = __gong__appendInstances(res, stage.PerpendicularVectors)

	res = __gong__appendInstances(res, stage.PerpendicularVectorGrids)

	res = __gong__appendInstances(res, stage.PerpendicularVectorGridHalfways)

	res = __gong__appendInstances(res, stage.PerpendicularVectorHalfways)

	res = __gong__appendInstances(res, stage.Plant2DDiagrams)

	res = __gong__appendInstances(res, stage.Plant3DDiagrams)

	res = __gong__appendInstances(res, stage.PlantAbstracts)

	res = __gong__appendInstances(res, stage.PlantCircumferenceShapes)

	res = __gong__appendInstances(res, stage.PointsAndLines3DShapes)

	res = __gong__appendInstances(res, stage.PxShapes)

	res = __gong__appendInstances(res, stage.Rendered3DShapes)

	res = __gong__appendInstances(res, stage.RhombusShapes)

	res = __gong__appendInstances(res, stage.RhombusStuffs)

	res = __gong__appendInstances(res, stage.RotatedRhombusGridShapes)

	res = __gong__appendInstances(res, stage.RotatedRhombusShapes)

	res = __gong__appendInstances(res, stage.RotatedSampledPoints3DShapes)

	res = __gong__appendInstances(res, stage.RotatedSeatAndLegs3DShapes)

	res = __gong__appendInstances(res, stage.SampledPoints3DShapes)

	res = __gong__appendInstances(res, stage.Seat3DShapes)

	res = __gong__appendInstances(res, stage.SeatAndLegs3DShapes)

	res = __gong__appendInstances(res, stage.SeatBottomCurveShapes)

	res = __gong__appendInstances(res, stage.SeatTopCurveShapes)

	res = __gong__appendInstances(res, stage.ShiftedBottomTopStartArcShapes)

	res = __gong__appendInstances(res, stage.ShiftedBottomTopStartArcShapeGrids)

	res = __gong__appendInstances(res, stage.ShiftedLeftGrowthCurve2DRibbons)

	res = __gong__appendInstances(res, stage.ShiftedLeftGrowthCurve2DRibbonEndShapes)

	res = __gong__appendInstances(res, stage.ShiftedLeftGrowthCurve2DRibbonStartShapes)

	res = __gong__appendInstances(res, stage.ShiftedLeftPartiallyGrowthCurve2DRibbons)

	res = __gong__appendInstances(res, stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes)

	res = __gong__appendInstances(res, stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes)

	res = __gong__appendInstances(res, stage.ShiftedLeftStackGrowthCurveEndArcShapes)

	res = __gong__appendInstances(res, stage.ShiftedLeftStackGrowthCurveStartArcShapes)

	res = __gong__appendInstances(res, stage.ShiftedLeftStackNormalVectors)

	res = __gong__appendInstances(res, stage.ShiftedLeftStackOfGrowthCurves)

	res = __gong__appendInstances(res, stage.ShiftedLeftStackOfNormalVectors)

	res = __gong__appendInstances(res, stage.ShiftedRightGrowthCurve2DRibbons)

	res = __gong__appendInstances(res, stage.ShiftedRightGrowthCurve2DRibbonEndShapes)

	res = __gong__appendInstances(res, stage.ShiftedRightGrowthCurve2DRibbonStartShapes)

	res = __gong__appendInstances(res, stage.StackGrowthCurve2DEndHalfwayArcShapes)

	res = __gong__appendInstances(res, stage.StackGrowthCurve2DRibbonEndShapes)

	res = __gong__appendInstances(res, stage.StackGrowthCurve2DRibbonStartShapes)

	res = __gong__appendInstances(res, stage.StackGrowthCurve2DStartHalfwayArcShapes)

	res = __gong__appendInstances(res, stage.StackOfGrowthCurve2Ds)

	res = __gong__appendInstances(res, stage.StackOfGrowthCurve2DByGrowthVectors)

	res = __gong__appendInstances(res, stage.StackOfGrowthCurve2DRibbons)

	res = __gong__appendInstances(res, stage.StackOfPartiallyRotatedTorusShapes)

	res = __gong__appendInstances(res, stage.StackOfRotatedGrowthCurve2Ds)

	res = __gong__appendInstances(res, stage.StackOfRotatedGrowthCurve2DRibbons)

	res = __gong__appendInstances(res, stage.StackOfRotatedVaseTrapezeRingsShapes)

	res = __gong__appendInstances(res, stage.StackOfVaseTrapezeRingsShapes)

	res = __gong__appendInstances(res, stage.StackRotatedGrowthCurve2DEndArcShapes)

	res = __gong__appendInstances(res, stage.StackRotatedGrowthCurve2DRibbonEndShapes)

	res = __gong__appendInstances(res, stage.StackRotatedGrowthCurve2DRibbonStartShapes)

	res = __gong__appendInstances(res, stage.StackRotatedGrowthCurve2DStartArcShapes)

	res = __gong__appendInstances(res, stage.StartArcShapes)

	res = __gong__appendInstances(res, stage.StartArcShapeGrids)

	res = __gong__appendInstances(res, stage.StartHalfwayArcShapes)

	res = __gong__appendInstances(res, stage.StartHalfwayArcShapeGrids)

	res = __gong__appendInstances(res, stage.StemCylinder3DShapes)

	res = __gong__appendInstances(res, stage.Stool2DDiagrams)

	res = __gong__appendInstances(res, stage.Stool3DDiagrams)

	res = __gong__appendInstances(res, stage.TiledFloor3DShapes)

	res = __gong__appendInstances(res, stage.TopCurvePlane1Shapes)

	res = __gong__appendInstances(res, stage.TopCurvePlane2Shapes)

	res = __gong__appendInstances(res, stage.TopEndArcShapes)

	res = __gong__appendInstances(res, stage.TopEndArcShapeGrids)

	res = __gong__appendInstances(res, stage.TopEndHalfwayArcShapes)

	res = __gong__appendInstances(res, stage.TopEndHalfwayArcShapeGrids)

	res = __gong__appendInstances(res, stage.TopGrowthCurve2Ds)

	res = __gong__appendInstances(res, stage.TopMidArcVectorShapes)

	res = __gong__appendInstances(res, stage.TopMidArcVectorShapeGrids)

	res = __gong__appendInstances(res, stage.TopStackGrowthCurve2DEndHalfwayArcShapes)

	res = __gong__appendInstances(res, stage.TopStackGrowthCurve2DStartHalfwayArcShapes)

	res = __gong__appendInstances(res, stage.TopStackOfGrowthCurve2Ds)

	res = __gong__appendInstances(res, stage.TopStackOfRotatedGrowthCurve2Ds)

	res = __gong__appendInstances(res, stage.TopStackOfRotatedGrowthCurve2DEndArcShapes)

	res = __gong__appendInstances(res, stage.TopStackOfRotatedGrowthCurve2DStartArcShapes)

	res = __gong__appendInstances(res, stage.TopStartArcShapes)

	res = __gong__appendInstances(res, stage.TopStartArcShapeGrids)

	res = __gong__appendInstances(res, stage.TopStartHalfwayArcShapes)

	res = __gong__appendInstances(res, stage.TopStartHalfwayArcShapeGrids)

	res = __gong__appendInstances(res, stage.Torus3DShapes)

	res = __gong__appendInstances(res, stage.TorusEdge3DShapes)

	res = __gong__appendInstances(res, stage.TorusStackShapes)

	res = __gong__appendInstances(res, stage.TubeVase3DDiagrams)

	res = __gong__appendInstances(res, stage.TubeVaseAbstracts)

	res = __gong__appendInstances(res, stage.Vase2DDiagrams)

	res = __gong__appendInstances(res, stage.VaseTrapezeRingShapes)

	res = __gong__appendInstances(res, stage.VerticalTorusStackShapes)

	res = __gong__appendInstances(res, stage.VolumeKey3DShapes)

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
func (angle0shape *Angle0Shape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, angle0shape)
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, arcnormalvectorshape)
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, arcnormalvectorshapegrid)
}

func (axesshape *AxesShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, axesshape)
}

func (basevectorshape *BaseVectorShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, basevectorshape)
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, basevectorshapegrid)
}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bottomcurveplane1shape)
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bottomcurveplane2shape)
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, chosenp1p2pairshape)
}

func (circlegridshape *CircleGridShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, circlegridshape)
}

func (circumference3dshape *Circumference3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, circumference3dshape)
}

func (clock2ddiagram *Clock2DDiagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, clock2ddiagram)
}

func (clock3ddiagram *Clock3DDiagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, clock3ddiagram)
}

func (clocktopcurveshape *ClockTopCurveShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, clocktopcurveshape)
}

func (cutline3dshape *CutLine3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, cutline3dshape)
}

func (endarcshape *EndArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, endarcshape)
}

func (endarcshapegrid *EndArcShapeGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, endarcshapegrid)
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, endhalfwayarcshape)
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, endhalfwayarcshapegrid)
}

func (explanationtextshape *ExplanationTextShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, explanationtextshape)
}

func (eye3dshape *Eye3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, eye3dshape)
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, eyecornerssampledpoints3dshape)
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, eyesampledpoints3dshape)
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, eyeseatbottomcurveshape)
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, eyestoolbottomcurveshape)
}

func (eyevolume3dshape *EyeVolume3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, eyevolume3dshape)
}

func (gridpathshape *GridPathShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gridpathshape)
}

func (growthcurve2d *GrowthCurve2D) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, growthcurve2d)
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, growthcurve2dribbon)
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, growthcurve2dribbonendshape)
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, growthcurve2dribbonstartshape)
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, growthcurverhombusgridshape)
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, growthcurverhombusshape)
}

func (growthvectorshape *GrowthVectorShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, growthvectorshape)
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, initialrhombusgridshape)
}

func (initialrhombusshape *InitialRhombusShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, initialrhombusshape)
}

func (key3dshape *Key3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, key3dshape)
}

func (keyhole3dshape *KeyHole3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, keyhole3dshape)
}

func (keyholeshape *KeyHoleShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, keyholeshape)
}

func (leaves3dshape *Leaves3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, leaves3dshape)
}

func (library *Library) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, library)
}

func (midarcvectorshape *MidArcVectorShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, midarcvectorshape)
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, midarcvectorshapegrid)
}

func (originalpoints3dshape *OriginalPoints3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, originalpoints3dshape)
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, parastichymcurves3dshape)
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, parastichyncurves3dshape)
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallygrowthcurve2dribbon)
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallygrowthcurve2dribbonendshape)
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallygrowthcurve2dribbonstartshape)
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallygrowthcurve2dtrajectory)
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallygrowthcurve2dtrajectoryp1curveshape)
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallygrowthcurve2dtrajectoryp1p2)
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallygrowthcurve2dtrajectoryp1p2pairlineshape)
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallygrowthcurve2dtrajectoryp1pointshape)
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallygrowthcurve2dtrajectoryp2curveshape)
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallygrowthcurve2dtrajectoryp2pointshape)
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallygrowthcurve2dtrajectoryshape)
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallyrotatedseatbottomcurveshape)
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallyrotatedseattopcurveshape)
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partiallyrotatedtorusshape)
}

func (perpendicularvector *PerpendicularVector) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, perpendicularvector)
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, perpendicularvectorgrid)
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, perpendicularvectorgridhalfway)
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, perpendicularvectorhalfway)
}

func (plant2ddiagram *Plant2DDiagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, plant2ddiagram)
}

func (plant3ddiagram *Plant3DDiagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, plant3ddiagram)
}

func (plantabstract *PlantAbstract) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, plantabstract)
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, plantcircumferenceshape)
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pointsandlines3dshape)
}

func (pxshape *PxShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pxshape)
}

func (rendered3dshape *Rendered3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rendered3dshape)
}

func (rhombusshape *RhombusShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rhombusshape)
}

func (rhombusstuff *RhombusStuff) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rhombusstuff)
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rotatedrhombusgridshape)
}

func (rotatedrhombusshape *RotatedRhombusShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rotatedrhombusshape)
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rotatedsampledpoints3dshape)
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rotatedseatandlegs3dshape)
}

func (sampledpoints3dshape *SampledPoints3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, sampledpoints3dshape)
}

func (seat3dshape *Seat3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, seat3dshape)
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, seatandlegs3dshape)
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, seatbottomcurveshape)
}

func (seattopcurveshape *SeatTopCurveShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, seattopcurveshape)
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedbottomtopstartarcshape)
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedbottomtopstartarcshapegrid)
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedleftgrowthcurve2dribbon)
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedleftgrowthcurve2dribbonendshape)
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedleftgrowthcurve2dribbonstartshape)
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedleftpartiallygrowthcurve2dribbon)
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedleftpartiallygrowthcurve2dribbonendshape)
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedleftpartiallygrowthcurve2dribbonstartshape)
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedleftstackgrowthcurveendarcshape)
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedleftstackgrowthcurvestartarcshape)
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedleftstacknormalvector)
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedleftstackofgrowthcurve)
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedleftstackofnormalvector)
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedrightgrowthcurve2dribbon)
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedrightgrowthcurve2dribbonendshape)
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shiftedrightgrowthcurve2dribbonstartshape)
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackgrowthcurve2dendhalfwayarcshape)
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackgrowthcurve2dribbonendshape)
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackgrowthcurve2dribbonstartshape)
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackgrowthcurve2dstarthalfwayarcshape)
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackofgrowthcurve2d)
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackofgrowthcurve2dbygrowthvector)
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackofgrowthcurve2dribbon)
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackofpartiallyrotatedtorusshape)
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackofrotatedgrowthcurve2d)
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackofrotatedgrowthcurve2dribbon)
}

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackofrotatedvasetrapezeringsshape)
}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackofvasetrapezeringsshape)
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackrotatedgrowthcurve2dendarcshape)
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackrotatedgrowthcurve2dribbonendshape)
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackrotatedgrowthcurve2dribbonstartshape)
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stackrotatedgrowthcurve2dstartarcshape)
}

func (startarcshape *StartArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, startarcshape)
}

func (startarcshapegrid *StartArcShapeGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, startarcshapegrid)
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, starthalfwayarcshape)
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, starthalfwayarcshapegrid)
}

func (stemcylinder3dshape *StemCylinder3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stemcylinder3dshape)
}

func (stool2ddiagram *Stool2DDiagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stool2ddiagram)
}

func (stool3ddiagram *Stool3DDiagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stool3ddiagram)
}

func (tiledfloor3dshape *TiledFloor3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tiledfloor3dshape)
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topcurveplane1shape)
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topcurveplane2shape)
}

func (topendarcshape *TopEndArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topendarcshape)
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topendarcshapegrid)
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topendhalfwayarcshape)
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topendhalfwayarcshapegrid)
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topgrowthcurve2d)
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topmidarcvectorshape)
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topmidarcvectorshapegrid)
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topstackgrowthcurve2dendhalfwayarcshape)
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topstackgrowthcurve2dstarthalfwayarcshape)
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topstackofgrowthcurve2d)
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topstackofrotatedgrowthcurve2d)
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topstackofrotatedgrowthcurve2dendarcshape)
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topstackofrotatedgrowthcurve2dstartarcshape)
}

func (topstartarcshape *TopStartArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topstartarcshape)
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topstartarcshapegrid)
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topstarthalfwayarcshape)
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, topstarthalfwayarcshapegrid)
}

func (torus3dshape *Torus3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, torus3dshape)
}

func (torusedge3dshape *TorusEdge3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, torusedge3dshape)
}

func (torusstackshape *TorusStackShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, torusstackshape)
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tubevase3ddiagram)
}

func (tubevaseabstract *TubeVaseAbstract) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tubevaseabstract)
}

func (vase2ddiagram *Vase2DDiagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, vase2ddiagram)
}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, vasetrapezeringshape)
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, verticaltorusstackshape)
}

func (volumekey3dshape *VolumeKey3DShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, volumekey3dshape)
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
	__gong__computeReferencePass1(stage, stage.Angle0Shapes, &stage.Angle0Shapes_reference, &stage.Angle0Shapes_referenceOrder, &stage.Angle0Shapes_instance)

	__gong__computeReferencePass1(stage, stage.ArcNormalVectorShapes, &stage.ArcNormalVectorShapes_reference, &stage.ArcNormalVectorShapes_referenceOrder, &stage.ArcNormalVectorShapes_instance)

	__gong__computeReferencePass1(stage, stage.ArcNormalVectorShapeGrids, &stage.ArcNormalVectorShapeGrids_reference, &stage.ArcNormalVectorShapeGrids_referenceOrder, &stage.ArcNormalVectorShapeGrids_instance)

	__gong__computeReferencePass1(stage, stage.AxesShapes, &stage.AxesShapes_reference, &stage.AxesShapes_referenceOrder, &stage.AxesShapes_instance)

	__gong__computeReferencePass1(stage, stage.BaseVectorShapes, &stage.BaseVectorShapes_reference, &stage.BaseVectorShapes_referenceOrder, &stage.BaseVectorShapes_instance)

	__gong__computeReferencePass1(stage, stage.BaseVectorShapeGrids, &stage.BaseVectorShapeGrids_reference, &stage.BaseVectorShapeGrids_referenceOrder, &stage.BaseVectorShapeGrids_instance)

	__gong__computeReferencePass1(stage, stage.BottomCurvePlane1Shapes, &stage.BottomCurvePlane1Shapes_reference, &stage.BottomCurvePlane1Shapes_referenceOrder, &stage.BottomCurvePlane1Shapes_instance)

	__gong__computeReferencePass1(stage, stage.BottomCurvePlane2Shapes, &stage.BottomCurvePlane2Shapes_reference, &stage.BottomCurvePlane2Shapes_referenceOrder, &stage.BottomCurvePlane2Shapes_instance)

	__gong__computeReferencePass1(stage, stage.ChosenP1P2PairShapes, &stage.ChosenP1P2PairShapes_reference, &stage.ChosenP1P2PairShapes_referenceOrder, &stage.ChosenP1P2PairShapes_instance)

	__gong__computeReferencePass1(stage, stage.CircleGridShapes, &stage.CircleGridShapes_reference, &stage.CircleGridShapes_referenceOrder, &stage.CircleGridShapes_instance)

	__gong__computeReferencePass1(stage, stage.Circumference3DShapes, &stage.Circumference3DShapes_reference, &stage.Circumference3DShapes_referenceOrder, &stage.Circumference3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.Clock2DDiagrams, &stage.Clock2DDiagrams_reference, &stage.Clock2DDiagrams_referenceOrder, &stage.Clock2DDiagrams_instance)

	__gong__computeReferencePass1(stage, stage.Clock3DDiagrams, &stage.Clock3DDiagrams_reference, &stage.Clock3DDiagrams_referenceOrder, &stage.Clock3DDiagrams_instance)

	__gong__computeReferencePass1(stage, stage.ClockTopCurveShapes, &stage.ClockTopCurveShapes_reference, &stage.ClockTopCurveShapes_referenceOrder, &stage.ClockTopCurveShapes_instance)

	__gong__computeReferencePass1(stage, stage.CutLine3DShapes, &stage.CutLine3DShapes_reference, &stage.CutLine3DShapes_referenceOrder, &stage.CutLine3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.EndArcShapes, &stage.EndArcShapes_reference, &stage.EndArcShapes_referenceOrder, &stage.EndArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.EndArcShapeGrids, &stage.EndArcShapeGrids_reference, &stage.EndArcShapeGrids_referenceOrder, &stage.EndArcShapeGrids_instance)

	__gong__computeReferencePass1(stage, stage.EndHalfwayArcShapes, &stage.EndHalfwayArcShapes_reference, &stage.EndHalfwayArcShapes_referenceOrder, &stage.EndHalfwayArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.EndHalfwayArcShapeGrids, &stage.EndHalfwayArcShapeGrids_reference, &stage.EndHalfwayArcShapeGrids_referenceOrder, &stage.EndHalfwayArcShapeGrids_instance)

	__gong__computeReferencePass1(stage, stage.ExplanationTextShapes, &stage.ExplanationTextShapes_reference, &stage.ExplanationTextShapes_referenceOrder, &stage.ExplanationTextShapes_instance)

	__gong__computeReferencePass1(stage, stage.Eye3DShapes, &stage.Eye3DShapes_reference, &stage.Eye3DShapes_referenceOrder, &stage.Eye3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.EyeCornersSampledPoints3DShapes, &stage.EyeCornersSampledPoints3DShapes_reference, &stage.EyeCornersSampledPoints3DShapes_referenceOrder, &stage.EyeCornersSampledPoints3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.EyeSampledPoints3DShapes, &stage.EyeSampledPoints3DShapes_reference, &stage.EyeSampledPoints3DShapes_referenceOrder, &stage.EyeSampledPoints3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.EyeSeatBottomCurveShapes, &stage.EyeSeatBottomCurveShapes_reference, &stage.EyeSeatBottomCurveShapes_referenceOrder, &stage.EyeSeatBottomCurveShapes_instance)

	__gong__computeReferencePass1(stage, stage.EyeStoolBottomCurveShapes, &stage.EyeStoolBottomCurveShapes_reference, &stage.EyeStoolBottomCurveShapes_referenceOrder, &stage.EyeStoolBottomCurveShapes_instance)

	__gong__computeReferencePass1(stage, stage.EyeVolume3DShapes, &stage.EyeVolume3DShapes_reference, &stage.EyeVolume3DShapes_referenceOrder, &stage.EyeVolume3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.GridPathShapes, &stage.GridPathShapes_reference, &stage.GridPathShapes_referenceOrder, &stage.GridPathShapes_instance)

	__gong__computeReferencePass1(stage, stage.GrowthCurve2Ds, &stage.GrowthCurve2Ds_reference, &stage.GrowthCurve2Ds_referenceOrder, &stage.GrowthCurve2Ds_instance)

	__gong__computeReferencePass1(stage, stage.GrowthCurve2DRibbons, &stage.GrowthCurve2DRibbons_reference, &stage.GrowthCurve2DRibbons_referenceOrder, &stage.GrowthCurve2DRibbons_instance)

	__gong__computeReferencePass1(stage, stage.GrowthCurve2DRibbonEndShapes, &stage.GrowthCurve2DRibbonEndShapes_reference, &stage.GrowthCurve2DRibbonEndShapes_referenceOrder, &stage.GrowthCurve2DRibbonEndShapes_instance)

	__gong__computeReferencePass1(stage, stage.GrowthCurve2DRibbonStartShapes, &stage.GrowthCurve2DRibbonStartShapes_reference, &stage.GrowthCurve2DRibbonStartShapes_referenceOrder, &stage.GrowthCurve2DRibbonStartShapes_instance)

	__gong__computeReferencePass1(stage, stage.GrowthCurveRhombusGridShapes, &stage.GrowthCurveRhombusGridShapes_reference, &stage.GrowthCurveRhombusGridShapes_referenceOrder, &stage.GrowthCurveRhombusGridShapes_instance)

	__gong__computeReferencePass1(stage, stage.GrowthCurveRhombusShapes, &stage.GrowthCurveRhombusShapes_reference, &stage.GrowthCurveRhombusShapes_referenceOrder, &stage.GrowthCurveRhombusShapes_instance)

	__gong__computeReferencePass1(stage, stage.GrowthVectorShapes, &stage.GrowthVectorShapes_reference, &stage.GrowthVectorShapes_referenceOrder, &stage.GrowthVectorShapes_instance)

	__gong__computeReferencePass1(stage, stage.InitialRhombusGridShapes, &stage.InitialRhombusGridShapes_reference, &stage.InitialRhombusGridShapes_referenceOrder, &stage.InitialRhombusGridShapes_instance)

	__gong__computeReferencePass1(stage, stage.InitialRhombusShapes, &stage.InitialRhombusShapes_reference, &stage.InitialRhombusShapes_referenceOrder, &stage.InitialRhombusShapes_instance)

	__gong__computeReferencePass1(stage, stage.Key3DShapes, &stage.Key3DShapes_reference, &stage.Key3DShapes_referenceOrder, &stage.Key3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.KeyHole3DShapes, &stage.KeyHole3DShapes_reference, &stage.KeyHole3DShapes_referenceOrder, &stage.KeyHole3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.KeyHoleShapes, &stage.KeyHoleShapes_reference, &stage.KeyHoleShapes_referenceOrder, &stage.KeyHoleShapes_instance)

	__gong__computeReferencePass1(stage, stage.Leaves3DShapes, &stage.Leaves3DShapes_reference, &stage.Leaves3DShapes_referenceOrder, &stage.Leaves3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.Librarys, &stage.Librarys_reference, &stage.Librarys_referenceOrder, &stage.Librarys_instance)

	__gong__computeReferencePass1(stage, stage.MidArcVectorShapes, &stage.MidArcVectorShapes_reference, &stage.MidArcVectorShapes_referenceOrder, &stage.MidArcVectorShapes_instance)

	__gong__computeReferencePass1(stage, stage.MidArcVectorShapeGrids, &stage.MidArcVectorShapeGrids_reference, &stage.MidArcVectorShapeGrids_referenceOrder, &stage.MidArcVectorShapeGrids_instance)

	__gong__computeReferencePass1(stage, stage.OriginalPoints3DShapes, &stage.OriginalPoints3DShapes_reference, &stage.OriginalPoints3DShapes_referenceOrder, &stage.OriginalPoints3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.ParastichyMCurves3DShapes, &stage.ParastichyMCurves3DShapes_reference, &stage.ParastichyMCurves3DShapes_referenceOrder, &stage.ParastichyMCurves3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.ParastichyNCurves3DShapes, &stage.ParastichyNCurves3DShapes_reference, &stage.ParastichyNCurves3DShapes_referenceOrder, &stage.ParastichyNCurves3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyGrowthCurve2DRibbons, &stage.PartiallyGrowthCurve2DRibbons_reference, &stage.PartiallyGrowthCurve2DRibbons_referenceOrder, &stage.PartiallyGrowthCurve2DRibbons_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyGrowthCurve2DRibbonEndShapes, &stage.PartiallyGrowthCurve2DRibbonEndShapes_reference, &stage.PartiallyGrowthCurve2DRibbonEndShapes_referenceOrder, &stage.PartiallyGrowthCurve2DRibbonEndShapes_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyGrowthCurve2DRibbonStartShapes, &stage.PartiallyGrowthCurve2DRibbonStartShapes_reference, &stage.PartiallyGrowthCurve2DRibbonStartShapes_referenceOrder, &stage.PartiallyGrowthCurve2DRibbonStartShapes_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyGrowthCurve2DTrajectorys, &stage.PartiallyGrowthCurve2DTrajectorys_reference, &stage.PartiallyGrowthCurve2DTrajectorys_referenceOrder, &stage.PartiallyGrowthCurve2DTrajectorys_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes, &stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes_reference, &stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes_referenceOrder, &stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyGrowthCurve2DTrajectoryP1P2s, &stage.PartiallyGrowthCurve2DTrajectoryP1P2s_reference, &stage.PartiallyGrowthCurve2DTrajectoryP1P2s_referenceOrder, &stage.PartiallyGrowthCurve2DTrajectoryP1P2s_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes, &stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes_reference, &stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes_referenceOrder, &stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes, &stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes_reference, &stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes_referenceOrder, &stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes, &stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes_reference, &stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes_referenceOrder, &stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes, &stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes_reference, &stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes_referenceOrder, &stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyGrowthCurve2DTrajectoryShapes, &stage.PartiallyGrowthCurve2DTrajectoryShapes_reference, &stage.PartiallyGrowthCurve2DTrajectoryShapes_referenceOrder, &stage.PartiallyGrowthCurve2DTrajectoryShapes_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyRotatedSeatBottomCurveShapes, &stage.PartiallyRotatedSeatBottomCurveShapes_reference, &stage.PartiallyRotatedSeatBottomCurveShapes_referenceOrder, &stage.PartiallyRotatedSeatBottomCurveShapes_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyRotatedSeatTopCurveShapes, &stage.PartiallyRotatedSeatTopCurveShapes_reference, &stage.PartiallyRotatedSeatTopCurveShapes_referenceOrder, &stage.PartiallyRotatedSeatTopCurveShapes_instance)

	__gong__computeReferencePass1(stage, stage.PartiallyRotatedTorusShapes, &stage.PartiallyRotatedTorusShapes_reference, &stage.PartiallyRotatedTorusShapes_referenceOrder, &stage.PartiallyRotatedTorusShapes_instance)

	__gong__computeReferencePass1(stage, stage.PerpendicularVectors, &stage.PerpendicularVectors_reference, &stage.PerpendicularVectors_referenceOrder, &stage.PerpendicularVectors_instance)

	__gong__computeReferencePass1(stage, stage.PerpendicularVectorGrids, &stage.PerpendicularVectorGrids_reference, &stage.PerpendicularVectorGrids_referenceOrder, &stage.PerpendicularVectorGrids_instance)

	__gong__computeReferencePass1(stage, stage.PerpendicularVectorGridHalfways, &stage.PerpendicularVectorGridHalfways_reference, &stage.PerpendicularVectorGridHalfways_referenceOrder, &stage.PerpendicularVectorGridHalfways_instance)

	__gong__computeReferencePass1(stage, stage.PerpendicularVectorHalfways, &stage.PerpendicularVectorHalfways_reference, &stage.PerpendicularVectorHalfways_referenceOrder, &stage.PerpendicularVectorHalfways_instance)

	__gong__computeReferencePass1(stage, stage.Plant2DDiagrams, &stage.Plant2DDiagrams_reference, &stage.Plant2DDiagrams_referenceOrder, &stage.Plant2DDiagrams_instance)

	__gong__computeReferencePass1(stage, stage.Plant3DDiagrams, &stage.Plant3DDiagrams_reference, &stage.Plant3DDiagrams_referenceOrder, &stage.Plant3DDiagrams_instance)

	__gong__computeReferencePass1(stage, stage.PlantAbstracts, &stage.PlantAbstracts_reference, &stage.PlantAbstracts_referenceOrder, &stage.PlantAbstracts_instance)

	__gong__computeReferencePass1(stage, stage.PlantCircumferenceShapes, &stage.PlantCircumferenceShapes_reference, &stage.PlantCircumferenceShapes_referenceOrder, &stage.PlantCircumferenceShapes_instance)

	__gong__computeReferencePass1(stage, stage.PointsAndLines3DShapes, &stage.PointsAndLines3DShapes_reference, &stage.PointsAndLines3DShapes_referenceOrder, &stage.PointsAndLines3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.PxShapes, &stage.PxShapes_reference, &stage.PxShapes_referenceOrder, &stage.PxShapes_instance)

	__gong__computeReferencePass1(stage, stage.Rendered3DShapes, &stage.Rendered3DShapes_reference, &stage.Rendered3DShapes_referenceOrder, &stage.Rendered3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.RhombusShapes, &stage.RhombusShapes_reference, &stage.RhombusShapes_referenceOrder, &stage.RhombusShapes_instance)

	__gong__computeReferencePass1(stage, stage.RhombusStuffs, &stage.RhombusStuffs_reference, &stage.RhombusStuffs_referenceOrder, &stage.RhombusStuffs_instance)

	__gong__computeReferencePass1(stage, stage.RotatedRhombusGridShapes, &stage.RotatedRhombusGridShapes_reference, &stage.RotatedRhombusGridShapes_referenceOrder, &stage.RotatedRhombusGridShapes_instance)

	__gong__computeReferencePass1(stage, stage.RotatedRhombusShapes, &stage.RotatedRhombusShapes_reference, &stage.RotatedRhombusShapes_referenceOrder, &stage.RotatedRhombusShapes_instance)

	__gong__computeReferencePass1(stage, stage.RotatedSampledPoints3DShapes, &stage.RotatedSampledPoints3DShapes_reference, &stage.RotatedSampledPoints3DShapes_referenceOrder, &stage.RotatedSampledPoints3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.RotatedSeatAndLegs3DShapes, &stage.RotatedSeatAndLegs3DShapes_reference, &stage.RotatedSeatAndLegs3DShapes_referenceOrder, &stage.RotatedSeatAndLegs3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.SampledPoints3DShapes, &stage.SampledPoints3DShapes_reference, &stage.SampledPoints3DShapes_referenceOrder, &stage.SampledPoints3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.Seat3DShapes, &stage.Seat3DShapes_reference, &stage.Seat3DShapes_referenceOrder, &stage.Seat3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.SeatAndLegs3DShapes, &stage.SeatAndLegs3DShapes_reference, &stage.SeatAndLegs3DShapes_referenceOrder, &stage.SeatAndLegs3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.SeatBottomCurveShapes, &stage.SeatBottomCurveShapes_reference, &stage.SeatBottomCurveShapes_referenceOrder, &stage.SeatBottomCurveShapes_instance)

	__gong__computeReferencePass1(stage, stage.SeatTopCurveShapes, &stage.SeatTopCurveShapes_reference, &stage.SeatTopCurveShapes_referenceOrder, &stage.SeatTopCurveShapes_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedBottomTopStartArcShapes, &stage.ShiftedBottomTopStartArcShapes_reference, &stage.ShiftedBottomTopStartArcShapes_referenceOrder, &stage.ShiftedBottomTopStartArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedBottomTopStartArcShapeGrids, &stage.ShiftedBottomTopStartArcShapeGrids_reference, &stage.ShiftedBottomTopStartArcShapeGrids_referenceOrder, &stage.ShiftedBottomTopStartArcShapeGrids_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedLeftGrowthCurve2DRibbons, &stage.ShiftedLeftGrowthCurve2DRibbons_reference, &stage.ShiftedLeftGrowthCurve2DRibbons_referenceOrder, &stage.ShiftedLeftGrowthCurve2DRibbons_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedLeftGrowthCurve2DRibbonEndShapes, &stage.ShiftedLeftGrowthCurve2DRibbonEndShapes_reference, &stage.ShiftedLeftGrowthCurve2DRibbonEndShapes_referenceOrder, &stage.ShiftedLeftGrowthCurve2DRibbonEndShapes_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedLeftGrowthCurve2DRibbonStartShapes, &stage.ShiftedLeftGrowthCurve2DRibbonStartShapes_reference, &stage.ShiftedLeftGrowthCurve2DRibbonStartShapes_referenceOrder, &stage.ShiftedLeftGrowthCurve2DRibbonStartShapes_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedLeftPartiallyGrowthCurve2DRibbons, &stage.ShiftedLeftPartiallyGrowthCurve2DRibbons_reference, &stage.ShiftedLeftPartiallyGrowthCurve2DRibbons_referenceOrder, &stage.ShiftedLeftPartiallyGrowthCurve2DRibbons_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes, &stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes_reference, &stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes_referenceOrder, &stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes, &stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes_reference, &stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes_referenceOrder, &stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedLeftStackGrowthCurveEndArcShapes, &stage.ShiftedLeftStackGrowthCurveEndArcShapes_reference, &stage.ShiftedLeftStackGrowthCurveEndArcShapes_referenceOrder, &stage.ShiftedLeftStackGrowthCurveEndArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedLeftStackGrowthCurveStartArcShapes, &stage.ShiftedLeftStackGrowthCurveStartArcShapes_reference, &stage.ShiftedLeftStackGrowthCurveStartArcShapes_referenceOrder, &stage.ShiftedLeftStackGrowthCurveStartArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedLeftStackNormalVectors, &stage.ShiftedLeftStackNormalVectors_reference, &stage.ShiftedLeftStackNormalVectors_referenceOrder, &stage.ShiftedLeftStackNormalVectors_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedLeftStackOfGrowthCurves, &stage.ShiftedLeftStackOfGrowthCurves_reference, &stage.ShiftedLeftStackOfGrowthCurves_referenceOrder, &stage.ShiftedLeftStackOfGrowthCurves_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedLeftStackOfNormalVectors, &stage.ShiftedLeftStackOfNormalVectors_reference, &stage.ShiftedLeftStackOfNormalVectors_referenceOrder, &stage.ShiftedLeftStackOfNormalVectors_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedRightGrowthCurve2DRibbons, &stage.ShiftedRightGrowthCurve2DRibbons_reference, &stage.ShiftedRightGrowthCurve2DRibbons_referenceOrder, &stage.ShiftedRightGrowthCurve2DRibbons_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedRightGrowthCurve2DRibbonEndShapes, &stage.ShiftedRightGrowthCurve2DRibbonEndShapes_reference, &stage.ShiftedRightGrowthCurve2DRibbonEndShapes_referenceOrder, &stage.ShiftedRightGrowthCurve2DRibbonEndShapes_instance)

	__gong__computeReferencePass1(stage, stage.ShiftedRightGrowthCurve2DRibbonStartShapes, &stage.ShiftedRightGrowthCurve2DRibbonStartShapes_reference, &stage.ShiftedRightGrowthCurve2DRibbonStartShapes_referenceOrder, &stage.ShiftedRightGrowthCurve2DRibbonStartShapes_instance)

	__gong__computeReferencePass1(stage, stage.StackGrowthCurve2DEndHalfwayArcShapes, &stage.StackGrowthCurve2DEndHalfwayArcShapes_reference, &stage.StackGrowthCurve2DEndHalfwayArcShapes_referenceOrder, &stage.StackGrowthCurve2DEndHalfwayArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.StackGrowthCurve2DRibbonEndShapes, &stage.StackGrowthCurve2DRibbonEndShapes_reference, &stage.StackGrowthCurve2DRibbonEndShapes_referenceOrder, &stage.StackGrowthCurve2DRibbonEndShapes_instance)

	__gong__computeReferencePass1(stage, stage.StackGrowthCurve2DRibbonStartShapes, &stage.StackGrowthCurve2DRibbonStartShapes_reference, &stage.StackGrowthCurve2DRibbonStartShapes_referenceOrder, &stage.StackGrowthCurve2DRibbonStartShapes_instance)

	__gong__computeReferencePass1(stage, stage.StackGrowthCurve2DStartHalfwayArcShapes, &stage.StackGrowthCurve2DStartHalfwayArcShapes_reference, &stage.StackGrowthCurve2DStartHalfwayArcShapes_referenceOrder, &stage.StackGrowthCurve2DStartHalfwayArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.StackOfGrowthCurve2Ds, &stage.StackOfGrowthCurve2Ds_reference, &stage.StackOfGrowthCurve2Ds_referenceOrder, &stage.StackOfGrowthCurve2Ds_instance)

	__gong__computeReferencePass1(stage, stage.StackOfGrowthCurve2DByGrowthVectors, &stage.StackOfGrowthCurve2DByGrowthVectors_reference, &stage.StackOfGrowthCurve2DByGrowthVectors_referenceOrder, &stage.StackOfGrowthCurve2DByGrowthVectors_instance)

	__gong__computeReferencePass1(stage, stage.StackOfGrowthCurve2DRibbons, &stage.StackOfGrowthCurve2DRibbons_reference, &stage.StackOfGrowthCurve2DRibbons_referenceOrder, &stage.StackOfGrowthCurve2DRibbons_instance)

	__gong__computeReferencePass1(stage, stage.StackOfPartiallyRotatedTorusShapes, &stage.StackOfPartiallyRotatedTorusShapes_reference, &stage.StackOfPartiallyRotatedTorusShapes_referenceOrder, &stage.StackOfPartiallyRotatedTorusShapes_instance)

	__gong__computeReferencePass1(stage, stage.StackOfRotatedGrowthCurve2Ds, &stage.StackOfRotatedGrowthCurve2Ds_reference, &stage.StackOfRotatedGrowthCurve2Ds_referenceOrder, &stage.StackOfRotatedGrowthCurve2Ds_instance)

	__gong__computeReferencePass1(stage, stage.StackOfRotatedGrowthCurve2DRibbons, &stage.StackOfRotatedGrowthCurve2DRibbons_reference, &stage.StackOfRotatedGrowthCurve2DRibbons_referenceOrder, &stage.StackOfRotatedGrowthCurve2DRibbons_instance)

	__gong__computeReferencePass1(stage, stage.StackOfRotatedVaseTrapezeRingsShapes, &stage.StackOfRotatedVaseTrapezeRingsShapes_reference, &stage.StackOfRotatedVaseTrapezeRingsShapes_referenceOrder, &stage.StackOfRotatedVaseTrapezeRingsShapes_instance)

	__gong__computeReferencePass1(stage, stage.StackOfVaseTrapezeRingsShapes, &stage.StackOfVaseTrapezeRingsShapes_reference, &stage.StackOfVaseTrapezeRingsShapes_referenceOrder, &stage.StackOfVaseTrapezeRingsShapes_instance)

	__gong__computeReferencePass1(stage, stage.StackRotatedGrowthCurve2DEndArcShapes, &stage.StackRotatedGrowthCurve2DEndArcShapes_reference, &stage.StackRotatedGrowthCurve2DEndArcShapes_referenceOrder, &stage.StackRotatedGrowthCurve2DEndArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.StackRotatedGrowthCurve2DRibbonEndShapes, &stage.StackRotatedGrowthCurve2DRibbonEndShapes_reference, &stage.StackRotatedGrowthCurve2DRibbonEndShapes_referenceOrder, &stage.StackRotatedGrowthCurve2DRibbonEndShapes_instance)

	__gong__computeReferencePass1(stage, stage.StackRotatedGrowthCurve2DRibbonStartShapes, &stage.StackRotatedGrowthCurve2DRibbonStartShapes_reference, &stage.StackRotatedGrowthCurve2DRibbonStartShapes_referenceOrder, &stage.StackRotatedGrowthCurve2DRibbonStartShapes_instance)

	__gong__computeReferencePass1(stage, stage.StackRotatedGrowthCurve2DStartArcShapes, &stage.StackRotatedGrowthCurve2DStartArcShapes_reference, &stage.StackRotatedGrowthCurve2DStartArcShapes_referenceOrder, &stage.StackRotatedGrowthCurve2DStartArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.StartArcShapes, &stage.StartArcShapes_reference, &stage.StartArcShapes_referenceOrder, &stage.StartArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.StartArcShapeGrids, &stage.StartArcShapeGrids_reference, &stage.StartArcShapeGrids_referenceOrder, &stage.StartArcShapeGrids_instance)

	__gong__computeReferencePass1(stage, stage.StartHalfwayArcShapes, &stage.StartHalfwayArcShapes_reference, &stage.StartHalfwayArcShapes_referenceOrder, &stage.StartHalfwayArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.StartHalfwayArcShapeGrids, &stage.StartHalfwayArcShapeGrids_reference, &stage.StartHalfwayArcShapeGrids_referenceOrder, &stage.StartHalfwayArcShapeGrids_instance)

	__gong__computeReferencePass1(stage, stage.StemCylinder3DShapes, &stage.StemCylinder3DShapes_reference, &stage.StemCylinder3DShapes_referenceOrder, &stage.StemCylinder3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.Stool2DDiagrams, &stage.Stool2DDiagrams_reference, &stage.Stool2DDiagrams_referenceOrder, &stage.Stool2DDiagrams_instance)

	__gong__computeReferencePass1(stage, stage.Stool3DDiagrams, &stage.Stool3DDiagrams_reference, &stage.Stool3DDiagrams_referenceOrder, &stage.Stool3DDiagrams_instance)

	__gong__computeReferencePass1(stage, stage.TiledFloor3DShapes, &stage.TiledFloor3DShapes_reference, &stage.TiledFloor3DShapes_referenceOrder, &stage.TiledFloor3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.TopCurvePlane1Shapes, &stage.TopCurvePlane1Shapes_reference, &stage.TopCurvePlane1Shapes_referenceOrder, &stage.TopCurvePlane1Shapes_instance)

	__gong__computeReferencePass1(stage, stage.TopCurvePlane2Shapes, &stage.TopCurvePlane2Shapes_reference, &stage.TopCurvePlane2Shapes_referenceOrder, &stage.TopCurvePlane2Shapes_instance)

	__gong__computeReferencePass1(stage, stage.TopEndArcShapes, &stage.TopEndArcShapes_reference, &stage.TopEndArcShapes_referenceOrder, &stage.TopEndArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.TopEndArcShapeGrids, &stage.TopEndArcShapeGrids_reference, &stage.TopEndArcShapeGrids_referenceOrder, &stage.TopEndArcShapeGrids_instance)

	__gong__computeReferencePass1(stage, stage.TopEndHalfwayArcShapes, &stage.TopEndHalfwayArcShapes_reference, &stage.TopEndHalfwayArcShapes_referenceOrder, &stage.TopEndHalfwayArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.TopEndHalfwayArcShapeGrids, &stage.TopEndHalfwayArcShapeGrids_reference, &stage.TopEndHalfwayArcShapeGrids_referenceOrder, &stage.TopEndHalfwayArcShapeGrids_instance)

	__gong__computeReferencePass1(stage, stage.TopGrowthCurve2Ds, &stage.TopGrowthCurve2Ds_reference, &stage.TopGrowthCurve2Ds_referenceOrder, &stage.TopGrowthCurve2Ds_instance)

	__gong__computeReferencePass1(stage, stage.TopMidArcVectorShapes, &stage.TopMidArcVectorShapes_reference, &stage.TopMidArcVectorShapes_referenceOrder, &stage.TopMidArcVectorShapes_instance)

	__gong__computeReferencePass1(stage, stage.TopMidArcVectorShapeGrids, &stage.TopMidArcVectorShapeGrids_reference, &stage.TopMidArcVectorShapeGrids_referenceOrder, &stage.TopMidArcVectorShapeGrids_instance)

	__gong__computeReferencePass1(stage, stage.TopStackGrowthCurve2DEndHalfwayArcShapes, &stage.TopStackGrowthCurve2DEndHalfwayArcShapes_reference, &stage.TopStackGrowthCurve2DEndHalfwayArcShapes_referenceOrder, &stage.TopStackGrowthCurve2DEndHalfwayArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.TopStackGrowthCurve2DStartHalfwayArcShapes, &stage.TopStackGrowthCurve2DStartHalfwayArcShapes_reference, &stage.TopStackGrowthCurve2DStartHalfwayArcShapes_referenceOrder, &stage.TopStackGrowthCurve2DStartHalfwayArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.TopStackOfGrowthCurve2Ds, &stage.TopStackOfGrowthCurve2Ds_reference, &stage.TopStackOfGrowthCurve2Ds_referenceOrder, &stage.TopStackOfGrowthCurve2Ds_instance)

	__gong__computeReferencePass1(stage, stage.TopStackOfRotatedGrowthCurve2Ds, &stage.TopStackOfRotatedGrowthCurve2Ds_reference, &stage.TopStackOfRotatedGrowthCurve2Ds_referenceOrder, &stage.TopStackOfRotatedGrowthCurve2Ds_instance)

	__gong__computeReferencePass1(stage, stage.TopStackOfRotatedGrowthCurve2DEndArcShapes, &stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_reference, &stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_referenceOrder, &stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.TopStackOfRotatedGrowthCurve2DStartArcShapes, &stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_reference, &stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_referenceOrder, &stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.TopStartArcShapes, &stage.TopStartArcShapes_reference, &stage.TopStartArcShapes_referenceOrder, &stage.TopStartArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.TopStartArcShapeGrids, &stage.TopStartArcShapeGrids_reference, &stage.TopStartArcShapeGrids_referenceOrder, &stage.TopStartArcShapeGrids_instance)

	__gong__computeReferencePass1(stage, stage.TopStartHalfwayArcShapes, &stage.TopStartHalfwayArcShapes_reference, &stage.TopStartHalfwayArcShapes_referenceOrder, &stage.TopStartHalfwayArcShapes_instance)

	__gong__computeReferencePass1(stage, stage.TopStartHalfwayArcShapeGrids, &stage.TopStartHalfwayArcShapeGrids_reference, &stage.TopStartHalfwayArcShapeGrids_referenceOrder, &stage.TopStartHalfwayArcShapeGrids_instance)

	__gong__computeReferencePass1(stage, stage.Torus3DShapes, &stage.Torus3DShapes_reference, &stage.Torus3DShapes_referenceOrder, &stage.Torus3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.TorusEdge3DShapes, &stage.TorusEdge3DShapes_reference, &stage.TorusEdge3DShapes_referenceOrder, &stage.TorusEdge3DShapes_instance)

	__gong__computeReferencePass1(stage, stage.TorusStackShapes, &stage.TorusStackShapes_reference, &stage.TorusStackShapes_referenceOrder, &stage.TorusStackShapes_instance)

	__gong__computeReferencePass1(stage, stage.TubeVase3DDiagrams, &stage.TubeVase3DDiagrams_reference, &stage.TubeVase3DDiagrams_referenceOrder, &stage.TubeVase3DDiagrams_instance)

	__gong__computeReferencePass1(stage, stage.TubeVaseAbstracts, &stage.TubeVaseAbstracts_reference, &stage.TubeVaseAbstracts_referenceOrder, &stage.TubeVaseAbstracts_instance)

	__gong__computeReferencePass1(stage, stage.Vase2DDiagrams, &stage.Vase2DDiagrams_reference, &stage.Vase2DDiagrams_referenceOrder, &stage.Vase2DDiagrams_instance)

	__gong__computeReferencePass1(stage, stage.VaseTrapezeRingShapes, &stage.VaseTrapezeRingShapes_reference, &stage.VaseTrapezeRingShapes_referenceOrder, &stage.VaseTrapezeRingShapes_instance)

	__gong__computeReferencePass1(stage, stage.VerticalTorusStackShapes, &stage.VerticalTorusStackShapes_reference, &stage.VerticalTorusStackShapes_referenceOrder, &stage.VerticalTorusStackShapes_instance)

	__gong__computeReferencePass1(stage, stage.VolumeKey3DShapes, &stage.VolumeKey3DShapes_reference, &stage.VolumeKey3DShapes_referenceOrder, &stage.VolumeKey3DShapes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Angle0Shapes, stage.Angle0Shapes_reference, stage)

	__gong__computeReferencePass2(stage.ArcNormalVectorShapes, stage.ArcNormalVectorShapes_reference, stage)

	__gong__computeReferencePass2(stage.ArcNormalVectorShapeGrids, stage.ArcNormalVectorShapeGrids_reference, stage)

	__gong__computeReferencePass2(stage.AxesShapes, stage.AxesShapes_reference, stage)

	__gong__computeReferencePass2(stage.BaseVectorShapes, stage.BaseVectorShapes_reference, stage)

	__gong__computeReferencePass2(stage.BaseVectorShapeGrids, stage.BaseVectorShapeGrids_reference, stage)

	__gong__computeReferencePass2(stage.BottomCurvePlane1Shapes, stage.BottomCurvePlane1Shapes_reference, stage)

	__gong__computeReferencePass2(stage.BottomCurvePlane2Shapes, stage.BottomCurvePlane2Shapes_reference, stage)

	__gong__computeReferencePass2(stage.ChosenP1P2PairShapes, stage.ChosenP1P2PairShapes_reference, stage)

	__gong__computeReferencePass2(stage.CircleGridShapes, stage.CircleGridShapes_reference, stage)

	__gong__computeReferencePass2(stage.Circumference3DShapes, stage.Circumference3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.Clock2DDiagrams, stage.Clock2DDiagrams_reference, stage)

	__gong__computeReferencePass2(stage.Clock3DDiagrams, stage.Clock3DDiagrams_reference, stage)

	__gong__computeReferencePass2(stage.ClockTopCurveShapes, stage.ClockTopCurveShapes_reference, stage)

	__gong__computeReferencePass2(stage.CutLine3DShapes, stage.CutLine3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.EndArcShapes, stage.EndArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.EndArcShapeGrids, stage.EndArcShapeGrids_reference, stage)

	__gong__computeReferencePass2(stage.EndHalfwayArcShapes, stage.EndHalfwayArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.EndHalfwayArcShapeGrids, stage.EndHalfwayArcShapeGrids_reference, stage)

	__gong__computeReferencePass2(stage.ExplanationTextShapes, stage.ExplanationTextShapes_reference, stage)

	__gong__computeReferencePass2(stage.Eye3DShapes, stage.Eye3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.EyeCornersSampledPoints3DShapes, stage.EyeCornersSampledPoints3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.EyeSampledPoints3DShapes, stage.EyeSampledPoints3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.EyeSeatBottomCurveShapes, stage.EyeSeatBottomCurveShapes_reference, stage)

	__gong__computeReferencePass2(stage.EyeStoolBottomCurveShapes, stage.EyeStoolBottomCurveShapes_reference, stage)

	__gong__computeReferencePass2(stage.EyeVolume3DShapes, stage.EyeVolume3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.GridPathShapes, stage.GridPathShapes_reference, stage)

	__gong__computeReferencePass2(stage.GrowthCurve2Ds, stage.GrowthCurve2Ds_reference, stage)

	__gong__computeReferencePass2(stage.GrowthCurve2DRibbons, stage.GrowthCurve2DRibbons_reference, stage)

	__gong__computeReferencePass2(stage.GrowthCurve2DRibbonEndShapes, stage.GrowthCurve2DRibbonEndShapes_reference, stage)

	__gong__computeReferencePass2(stage.GrowthCurve2DRibbonStartShapes, stage.GrowthCurve2DRibbonStartShapes_reference, stage)

	__gong__computeReferencePass2(stage.GrowthCurveRhombusGridShapes, stage.GrowthCurveRhombusGridShapes_reference, stage)

	__gong__computeReferencePass2(stage.GrowthCurveRhombusShapes, stage.GrowthCurveRhombusShapes_reference, stage)

	__gong__computeReferencePass2(stage.GrowthVectorShapes, stage.GrowthVectorShapes_reference, stage)

	__gong__computeReferencePass2(stage.InitialRhombusGridShapes, stage.InitialRhombusGridShapes_reference, stage)

	__gong__computeReferencePass2(stage.InitialRhombusShapes, stage.InitialRhombusShapes_reference, stage)

	__gong__computeReferencePass2(stage.Key3DShapes, stage.Key3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.KeyHole3DShapes, stage.KeyHole3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.KeyHoleShapes, stage.KeyHoleShapes_reference, stage)

	__gong__computeReferencePass2(stage.Leaves3DShapes, stage.Leaves3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.Librarys, stage.Librarys_reference, stage)

	__gong__computeReferencePass2(stage.MidArcVectorShapes, stage.MidArcVectorShapes_reference, stage)

	__gong__computeReferencePass2(stage.MidArcVectorShapeGrids, stage.MidArcVectorShapeGrids_reference, stage)

	__gong__computeReferencePass2(stage.OriginalPoints3DShapes, stage.OriginalPoints3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.ParastichyMCurves3DShapes, stage.ParastichyMCurves3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.ParastichyNCurves3DShapes, stage.ParastichyNCurves3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyGrowthCurve2DRibbons, stage.PartiallyGrowthCurve2DRibbons_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyGrowthCurve2DRibbonEndShapes, stage.PartiallyGrowthCurve2DRibbonEndShapes_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyGrowthCurve2DRibbonStartShapes, stage.PartiallyGrowthCurve2DRibbonStartShapes_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyGrowthCurve2DTrajectorys, stage.PartiallyGrowthCurve2DTrajectorys_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes, stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyGrowthCurve2DTrajectoryP1P2s, stage.PartiallyGrowthCurve2DTrajectoryP1P2s_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes, stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes, stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes, stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes, stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyGrowthCurve2DTrajectoryShapes, stage.PartiallyGrowthCurve2DTrajectoryShapes_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyRotatedSeatBottomCurveShapes, stage.PartiallyRotatedSeatBottomCurveShapes_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyRotatedSeatTopCurveShapes, stage.PartiallyRotatedSeatTopCurveShapes_reference, stage)

	__gong__computeReferencePass2(stage.PartiallyRotatedTorusShapes, stage.PartiallyRotatedTorusShapes_reference, stage)

	__gong__computeReferencePass2(stage.PerpendicularVectors, stage.PerpendicularVectors_reference, stage)

	__gong__computeReferencePass2(stage.PerpendicularVectorGrids, stage.PerpendicularVectorGrids_reference, stage)

	__gong__computeReferencePass2(stage.PerpendicularVectorGridHalfways, stage.PerpendicularVectorGridHalfways_reference, stage)

	__gong__computeReferencePass2(stage.PerpendicularVectorHalfways, stage.PerpendicularVectorHalfways_reference, stage)

	__gong__computeReferencePass2(stage.Plant2DDiagrams, stage.Plant2DDiagrams_reference, stage)

	__gong__computeReferencePass2(stage.Plant3DDiagrams, stage.Plant3DDiagrams_reference, stage)

	__gong__computeReferencePass2(stage.PlantAbstracts, stage.PlantAbstracts_reference, stage)

	__gong__computeReferencePass2(stage.PlantCircumferenceShapes, stage.PlantCircumferenceShapes_reference, stage)

	__gong__computeReferencePass2(stage.PointsAndLines3DShapes, stage.PointsAndLines3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.PxShapes, stage.PxShapes_reference, stage)

	__gong__computeReferencePass2(stage.Rendered3DShapes, stage.Rendered3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.RhombusShapes, stage.RhombusShapes_reference, stage)

	__gong__computeReferencePass2(stage.RhombusStuffs, stage.RhombusStuffs_reference, stage)

	__gong__computeReferencePass2(stage.RotatedRhombusGridShapes, stage.RotatedRhombusGridShapes_reference, stage)

	__gong__computeReferencePass2(stage.RotatedRhombusShapes, stage.RotatedRhombusShapes_reference, stage)

	__gong__computeReferencePass2(stage.RotatedSampledPoints3DShapes, stage.RotatedSampledPoints3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.RotatedSeatAndLegs3DShapes, stage.RotatedSeatAndLegs3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.SampledPoints3DShapes, stage.SampledPoints3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.Seat3DShapes, stage.Seat3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.SeatAndLegs3DShapes, stage.SeatAndLegs3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.SeatBottomCurveShapes, stage.SeatBottomCurveShapes_reference, stage)

	__gong__computeReferencePass2(stage.SeatTopCurveShapes, stage.SeatTopCurveShapes_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedBottomTopStartArcShapes, stage.ShiftedBottomTopStartArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedBottomTopStartArcShapeGrids, stage.ShiftedBottomTopStartArcShapeGrids_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedLeftGrowthCurve2DRibbons, stage.ShiftedLeftGrowthCurve2DRibbons_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedLeftGrowthCurve2DRibbonEndShapes, stage.ShiftedLeftGrowthCurve2DRibbonEndShapes_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedLeftGrowthCurve2DRibbonStartShapes, stage.ShiftedLeftGrowthCurve2DRibbonStartShapes_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedLeftPartiallyGrowthCurve2DRibbons, stage.ShiftedLeftPartiallyGrowthCurve2DRibbons_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes, stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes, stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedLeftStackGrowthCurveEndArcShapes, stage.ShiftedLeftStackGrowthCurveEndArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedLeftStackGrowthCurveStartArcShapes, stage.ShiftedLeftStackGrowthCurveStartArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedLeftStackNormalVectors, stage.ShiftedLeftStackNormalVectors_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedLeftStackOfGrowthCurves, stage.ShiftedLeftStackOfGrowthCurves_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedLeftStackOfNormalVectors, stage.ShiftedLeftStackOfNormalVectors_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedRightGrowthCurve2DRibbons, stage.ShiftedRightGrowthCurve2DRibbons_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedRightGrowthCurve2DRibbonEndShapes, stage.ShiftedRightGrowthCurve2DRibbonEndShapes_reference, stage)

	__gong__computeReferencePass2(stage.ShiftedRightGrowthCurve2DRibbonStartShapes, stage.ShiftedRightGrowthCurve2DRibbonStartShapes_reference, stage)

	__gong__computeReferencePass2(stage.StackGrowthCurve2DEndHalfwayArcShapes, stage.StackGrowthCurve2DEndHalfwayArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.StackGrowthCurve2DRibbonEndShapes, stage.StackGrowthCurve2DRibbonEndShapes_reference, stage)

	__gong__computeReferencePass2(stage.StackGrowthCurve2DRibbonStartShapes, stage.StackGrowthCurve2DRibbonStartShapes_reference, stage)

	__gong__computeReferencePass2(stage.StackGrowthCurve2DStartHalfwayArcShapes, stage.StackGrowthCurve2DStartHalfwayArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.StackOfGrowthCurve2Ds, stage.StackOfGrowthCurve2Ds_reference, stage)

	__gong__computeReferencePass2(stage.StackOfGrowthCurve2DByGrowthVectors, stage.StackOfGrowthCurve2DByGrowthVectors_reference, stage)

	__gong__computeReferencePass2(stage.StackOfGrowthCurve2DRibbons, stage.StackOfGrowthCurve2DRibbons_reference, stage)

	__gong__computeReferencePass2(stage.StackOfPartiallyRotatedTorusShapes, stage.StackOfPartiallyRotatedTorusShapes_reference, stage)

	__gong__computeReferencePass2(stage.StackOfRotatedGrowthCurve2Ds, stage.StackOfRotatedGrowthCurve2Ds_reference, stage)

	__gong__computeReferencePass2(stage.StackOfRotatedGrowthCurve2DRibbons, stage.StackOfRotatedGrowthCurve2DRibbons_reference, stage)

	__gong__computeReferencePass2(stage.StackOfRotatedVaseTrapezeRingsShapes, stage.StackOfRotatedVaseTrapezeRingsShapes_reference, stage)

	__gong__computeReferencePass2(stage.StackOfVaseTrapezeRingsShapes, stage.StackOfVaseTrapezeRingsShapes_reference, stage)

	__gong__computeReferencePass2(stage.StackRotatedGrowthCurve2DEndArcShapes, stage.StackRotatedGrowthCurve2DEndArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.StackRotatedGrowthCurve2DRibbonEndShapes, stage.StackRotatedGrowthCurve2DRibbonEndShapes_reference, stage)

	__gong__computeReferencePass2(stage.StackRotatedGrowthCurve2DRibbonStartShapes, stage.StackRotatedGrowthCurve2DRibbonStartShapes_reference, stage)

	__gong__computeReferencePass2(stage.StackRotatedGrowthCurve2DStartArcShapes, stage.StackRotatedGrowthCurve2DStartArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.StartArcShapes, stage.StartArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.StartArcShapeGrids, stage.StartArcShapeGrids_reference, stage)

	__gong__computeReferencePass2(stage.StartHalfwayArcShapes, stage.StartHalfwayArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.StartHalfwayArcShapeGrids, stage.StartHalfwayArcShapeGrids_reference, stage)

	__gong__computeReferencePass2(stage.StemCylinder3DShapes, stage.StemCylinder3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.Stool2DDiagrams, stage.Stool2DDiagrams_reference, stage)

	__gong__computeReferencePass2(stage.Stool3DDiagrams, stage.Stool3DDiagrams_reference, stage)

	__gong__computeReferencePass2(stage.TiledFloor3DShapes, stage.TiledFloor3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.TopCurvePlane1Shapes, stage.TopCurvePlane1Shapes_reference, stage)

	__gong__computeReferencePass2(stage.TopCurvePlane2Shapes, stage.TopCurvePlane2Shapes_reference, stage)

	__gong__computeReferencePass2(stage.TopEndArcShapes, stage.TopEndArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.TopEndArcShapeGrids, stage.TopEndArcShapeGrids_reference, stage)

	__gong__computeReferencePass2(stage.TopEndHalfwayArcShapes, stage.TopEndHalfwayArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.TopEndHalfwayArcShapeGrids, stage.TopEndHalfwayArcShapeGrids_reference, stage)

	__gong__computeReferencePass2(stage.TopGrowthCurve2Ds, stage.TopGrowthCurve2Ds_reference, stage)

	__gong__computeReferencePass2(stage.TopMidArcVectorShapes, stage.TopMidArcVectorShapes_reference, stage)

	__gong__computeReferencePass2(stage.TopMidArcVectorShapeGrids, stage.TopMidArcVectorShapeGrids_reference, stage)

	__gong__computeReferencePass2(stage.TopStackGrowthCurve2DEndHalfwayArcShapes, stage.TopStackGrowthCurve2DEndHalfwayArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.TopStackGrowthCurve2DStartHalfwayArcShapes, stage.TopStackGrowthCurve2DStartHalfwayArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.TopStackOfGrowthCurve2Ds, stage.TopStackOfGrowthCurve2Ds_reference, stage)

	__gong__computeReferencePass2(stage.TopStackOfRotatedGrowthCurve2Ds, stage.TopStackOfRotatedGrowthCurve2Ds_reference, stage)

	__gong__computeReferencePass2(stage.TopStackOfRotatedGrowthCurve2DEndArcShapes, stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.TopStackOfRotatedGrowthCurve2DStartArcShapes, stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.TopStartArcShapes, stage.TopStartArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.TopStartArcShapeGrids, stage.TopStartArcShapeGrids_reference, stage)

	__gong__computeReferencePass2(stage.TopStartHalfwayArcShapes, stage.TopStartHalfwayArcShapes_reference, stage)

	__gong__computeReferencePass2(stage.TopStartHalfwayArcShapeGrids, stage.TopStartHalfwayArcShapeGrids_reference, stage)

	__gong__computeReferencePass2(stage.Torus3DShapes, stage.Torus3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.TorusEdge3DShapes, stage.TorusEdge3DShapes_reference, stage)

	__gong__computeReferencePass2(stage.TorusStackShapes, stage.TorusStackShapes_reference, stage)

	__gong__computeReferencePass2(stage.TubeVase3DDiagrams, stage.TubeVase3DDiagrams_reference, stage)

	__gong__computeReferencePass2(stage.TubeVaseAbstracts, stage.TubeVaseAbstracts_reference, stage)

	__gong__computeReferencePass2(stage.Vase2DDiagrams, stage.Vase2DDiagrams_reference, stage)

	__gong__computeReferencePass2(stage.VaseTrapezeRingShapes, stage.VaseTrapezeRingShapes_reference, stage)

	__gong__computeReferencePass2(stage.VerticalTorusStackShapes, stage.VerticalTorusStackShapes_reference, stage)

	__gong__computeReferencePass2(stage.VolumeKey3DShapes, stage.VolumeKey3DShapes_reference, stage)

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
	return __gong__getOrder(stage.Angle0Shape_stagedOrder, stage.Angle0Shapes_referenceOrder, angle0shape, "Angle0Shape")
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ArcNormalVectorShape_stagedOrder, stage.ArcNormalVectorShapes_referenceOrder, arcnormalvectorshape, "ArcNormalVectorShape")
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ArcNormalVectorShapeGrid_stagedOrder, stage.ArcNormalVectorShapeGrids_referenceOrder, arcnormalvectorshapegrid, "ArcNormalVectorShapeGrid")
}

func (axesshape *AxesShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AxesShape_stagedOrder, stage.AxesShapes_referenceOrder, axesshape, "AxesShape")
}

func (basevectorshape *BaseVectorShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.BaseVectorShape_stagedOrder, stage.BaseVectorShapes_referenceOrder, basevectorshape, "BaseVectorShape")
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.BaseVectorShapeGrid_stagedOrder, stage.BaseVectorShapeGrids_referenceOrder, basevectorshapegrid, "BaseVectorShapeGrid")
}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.BottomCurvePlane1Shape_stagedOrder, stage.BottomCurvePlane1Shapes_referenceOrder, bottomcurveplane1shape, "BottomCurvePlane1Shape")
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.BottomCurvePlane2Shape_stagedOrder, stage.BottomCurvePlane2Shapes_referenceOrder, bottomcurveplane2shape, "BottomCurvePlane2Shape")
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ChosenP1P2PairShape_stagedOrder, stage.ChosenP1P2PairShapes_referenceOrder, chosenp1p2pairshape, "ChosenP1P2PairShape")
}

func (circlegridshape *CircleGridShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CircleGridShape_stagedOrder, stage.CircleGridShapes_referenceOrder, circlegridshape, "CircleGridShape")
}

func (circumference3dshape *Circumference3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Circumference3DShape_stagedOrder, stage.Circumference3DShapes_referenceOrder, circumference3dshape, "Circumference3DShape")
}

func (clock2ddiagram *Clock2DDiagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Clock2DDiagram_stagedOrder, stage.Clock2DDiagrams_referenceOrder, clock2ddiagram, "Clock2DDiagram")
}

func (clock3ddiagram *Clock3DDiagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Clock3DDiagram_stagedOrder, stage.Clock3DDiagrams_referenceOrder, clock3ddiagram, "Clock3DDiagram")
}

func (clocktopcurveshape *ClockTopCurveShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ClockTopCurveShape_stagedOrder, stage.ClockTopCurveShapes_referenceOrder, clocktopcurveshape, "ClockTopCurveShape")
}

func (cutline3dshape *CutLine3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CutLine3DShape_stagedOrder, stage.CutLine3DShapes_referenceOrder, cutline3dshape, "CutLine3DShape")
}

func (endarcshape *EndArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EndArcShape_stagedOrder, stage.EndArcShapes_referenceOrder, endarcshape, "EndArcShape")
}

func (endarcshapegrid *EndArcShapeGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EndArcShapeGrid_stagedOrder, stage.EndArcShapeGrids_referenceOrder, endarcshapegrid, "EndArcShapeGrid")
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EndHalfwayArcShape_stagedOrder, stage.EndHalfwayArcShapes_referenceOrder, endhalfwayarcshape, "EndHalfwayArcShape")
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EndHalfwayArcShapeGrid_stagedOrder, stage.EndHalfwayArcShapeGrids_referenceOrder, endhalfwayarcshapegrid, "EndHalfwayArcShapeGrid")
}

func (explanationtextshape *ExplanationTextShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ExplanationTextShape_stagedOrder, stage.ExplanationTextShapes_referenceOrder, explanationtextshape, "ExplanationTextShape")
}

func (eye3dshape *Eye3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Eye3DShape_stagedOrder, stage.Eye3DShapes_referenceOrder, eye3dshape, "Eye3DShape")
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EyeCornersSampledPoints3DShape_stagedOrder, stage.EyeCornersSampledPoints3DShapes_referenceOrder, eyecornerssampledpoints3dshape, "EyeCornersSampledPoints3DShape")
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EyeSampledPoints3DShape_stagedOrder, stage.EyeSampledPoints3DShapes_referenceOrder, eyesampledpoints3dshape, "EyeSampledPoints3DShape")
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EyeSeatBottomCurveShape_stagedOrder, stage.EyeSeatBottomCurveShapes_referenceOrder, eyeseatbottomcurveshape, "EyeSeatBottomCurveShape")
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EyeStoolBottomCurveShape_stagedOrder, stage.EyeStoolBottomCurveShapes_referenceOrder, eyestoolbottomcurveshape, "EyeStoolBottomCurveShape")
}

func (eyevolume3dshape *EyeVolume3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EyeVolume3DShape_stagedOrder, stage.EyeVolume3DShapes_referenceOrder, eyevolume3dshape, "EyeVolume3DShape")
}

func (gridpathshape *GridPathShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GridPathShape_stagedOrder, stage.GridPathShapes_referenceOrder, gridpathshape, "GridPathShape")
}

func (growthcurve2d *GrowthCurve2D) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GrowthCurve2D_stagedOrder, stage.GrowthCurve2Ds_referenceOrder, growthcurve2d, "GrowthCurve2D")
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GrowthCurve2DRibbon_stagedOrder, stage.GrowthCurve2DRibbons_referenceOrder, growthcurve2dribbon, "GrowthCurve2DRibbon")
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GrowthCurve2DRibbonEndShape_stagedOrder, stage.GrowthCurve2DRibbonEndShapes_referenceOrder, growthcurve2dribbonendshape, "GrowthCurve2DRibbonEndShape")
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GrowthCurve2DRibbonStartShape_stagedOrder, stage.GrowthCurve2DRibbonStartShapes_referenceOrder, growthcurve2dribbonstartshape, "GrowthCurve2DRibbonStartShape")
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GrowthCurveRhombusGridShape_stagedOrder, stage.GrowthCurveRhombusGridShapes_referenceOrder, growthcurverhombusgridshape, "GrowthCurveRhombusGridShape")
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GrowthCurveRhombusShape_stagedOrder, stage.GrowthCurveRhombusShapes_referenceOrder, growthcurverhombusshape, "GrowthCurveRhombusShape")
}

func (growthvectorshape *GrowthVectorShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GrowthVectorShape_stagedOrder, stage.GrowthVectorShapes_referenceOrder, growthvectorshape, "GrowthVectorShape")
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InitialRhombusGridShape_stagedOrder, stage.InitialRhombusGridShapes_referenceOrder, initialrhombusgridshape, "InitialRhombusGridShape")
}

func (initialrhombusshape *InitialRhombusShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InitialRhombusShape_stagedOrder, stage.InitialRhombusShapes_referenceOrder, initialrhombusshape, "InitialRhombusShape")
}

func (key3dshape *Key3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Key3DShape_stagedOrder, stage.Key3DShapes_referenceOrder, key3dshape, "Key3DShape")
}

func (keyhole3dshape *KeyHole3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.KeyHole3DShape_stagedOrder, stage.KeyHole3DShapes_referenceOrder, keyhole3dshape, "KeyHole3DShape")
}

func (keyholeshape *KeyHoleShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.KeyHoleShape_stagedOrder, stage.KeyHoleShapes_referenceOrder, keyholeshape, "KeyHoleShape")
}

func (leaves3dshape *Leaves3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Leaves3DShape_stagedOrder, stage.Leaves3DShapes_referenceOrder, leaves3dshape, "Leaves3DShape")
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Library_stagedOrder, stage.Librarys_referenceOrder, library, "Library")
}

func (midarcvectorshape *MidArcVectorShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MidArcVectorShape_stagedOrder, stage.MidArcVectorShapes_referenceOrder, midarcvectorshape, "MidArcVectorShape")
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MidArcVectorShapeGrid_stagedOrder, stage.MidArcVectorShapeGrids_referenceOrder, midarcvectorshapegrid, "MidArcVectorShapeGrid")
}

func (originalpoints3dshape *OriginalPoints3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OriginalPoints3DShape_stagedOrder, stage.OriginalPoints3DShapes_referenceOrder, originalpoints3dshape, "OriginalPoints3DShape")
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ParastichyMCurves3DShape_stagedOrder, stage.ParastichyMCurves3DShapes_referenceOrder, parastichymcurves3dshape, "ParastichyMCurves3DShape")
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ParastichyNCurves3DShape_stagedOrder, stage.ParastichyNCurves3DShapes_referenceOrder, parastichyncurves3dshape, "ParastichyNCurves3DShape")
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyGrowthCurve2DRibbon_stagedOrder, stage.PartiallyGrowthCurve2DRibbons_referenceOrder, partiallygrowthcurve2dribbon, "PartiallyGrowthCurve2DRibbon")
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyGrowthCurve2DRibbonEndShape_stagedOrder, stage.PartiallyGrowthCurve2DRibbonEndShapes_referenceOrder, partiallygrowthcurve2dribbonendshape, "PartiallyGrowthCurve2DRibbonEndShape")
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyGrowthCurve2DRibbonStartShape_stagedOrder, stage.PartiallyGrowthCurve2DRibbonStartShapes_referenceOrder, partiallygrowthcurve2dribbonstartshape, "PartiallyGrowthCurve2DRibbonStartShape")
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyGrowthCurve2DTrajectory_stagedOrder, stage.PartiallyGrowthCurve2DTrajectorys_referenceOrder, partiallygrowthcurve2dtrajectory, "PartiallyGrowthCurve2DTrajectory")
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyGrowthCurve2DTrajectoryP1CurveShape_stagedOrder, stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes_referenceOrder, partiallygrowthcurve2dtrajectoryp1curveshape, "PartiallyGrowthCurve2DTrajectoryP1CurveShape")
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyGrowthCurve2DTrajectoryP1P2_stagedOrder, stage.PartiallyGrowthCurve2DTrajectoryP1P2s_referenceOrder, partiallygrowthcurve2dtrajectoryp1p2, "PartiallyGrowthCurve2DTrajectoryP1P2")
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape_stagedOrder, stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes_referenceOrder, partiallygrowthcurve2dtrajectoryp1p2pairlineshape, "PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape")
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyGrowthCurve2DTrajectoryP1PointShape_stagedOrder, stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes_referenceOrder, partiallygrowthcurve2dtrajectoryp1pointshape, "PartiallyGrowthCurve2DTrajectoryP1PointShape")
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyGrowthCurve2DTrajectoryP2CurveShape_stagedOrder, stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes_referenceOrder, partiallygrowthcurve2dtrajectoryp2curveshape, "PartiallyGrowthCurve2DTrajectoryP2CurveShape")
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyGrowthCurve2DTrajectoryP2PointShape_stagedOrder, stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes_referenceOrder, partiallygrowthcurve2dtrajectoryp2pointshape, "PartiallyGrowthCurve2DTrajectoryP2PointShape")
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyGrowthCurve2DTrajectoryShape_stagedOrder, stage.PartiallyGrowthCurve2DTrajectoryShapes_referenceOrder, partiallygrowthcurve2dtrajectoryshape, "PartiallyGrowthCurve2DTrajectoryShape")
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyRotatedSeatBottomCurveShape_stagedOrder, stage.PartiallyRotatedSeatBottomCurveShapes_referenceOrder, partiallyrotatedseatbottomcurveshape, "PartiallyRotatedSeatBottomCurveShape")
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyRotatedSeatTopCurveShape_stagedOrder, stage.PartiallyRotatedSeatTopCurveShapes_referenceOrder, partiallyrotatedseattopcurveshape, "PartiallyRotatedSeatTopCurveShape")
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartiallyRotatedTorusShape_stagedOrder, stage.PartiallyRotatedTorusShapes_referenceOrder, partiallyrotatedtorusshape, "PartiallyRotatedTorusShape")
}

func (perpendicularvector *PerpendicularVector) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PerpendicularVector_stagedOrder, stage.PerpendicularVectors_referenceOrder, perpendicularvector, "PerpendicularVector")
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PerpendicularVectorGrid_stagedOrder, stage.PerpendicularVectorGrids_referenceOrder, perpendicularvectorgrid, "PerpendicularVectorGrid")
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PerpendicularVectorGridHalfway_stagedOrder, stage.PerpendicularVectorGridHalfways_referenceOrder, perpendicularvectorgridhalfway, "PerpendicularVectorGridHalfway")
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PerpendicularVectorHalfway_stagedOrder, stage.PerpendicularVectorHalfways_referenceOrder, perpendicularvectorhalfway, "PerpendicularVectorHalfway")
}

func (plant2ddiagram *Plant2DDiagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Plant2DDiagram_stagedOrder, stage.Plant2DDiagrams_referenceOrder, plant2ddiagram, "Plant2DDiagram")
}

func (plant3ddiagram *Plant3DDiagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Plant3DDiagram_stagedOrder, stage.Plant3DDiagrams_referenceOrder, plant3ddiagram, "Plant3DDiagram")
}

func (plantabstract *PlantAbstract) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PlantAbstract_stagedOrder, stage.PlantAbstracts_referenceOrder, plantabstract, "PlantAbstract")
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PlantCircumferenceShape_stagedOrder, stage.PlantCircumferenceShapes_referenceOrder, plantcircumferenceshape, "PlantCircumferenceShape")
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PointsAndLines3DShape_stagedOrder, stage.PointsAndLines3DShapes_referenceOrder, pointsandlines3dshape, "PointsAndLines3DShape")
}

func (pxshape *PxShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PxShape_stagedOrder, stage.PxShapes_referenceOrder, pxshape, "PxShape")
}

func (rendered3dshape *Rendered3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Rendered3DShape_stagedOrder, stage.Rendered3DShapes_referenceOrder, rendered3dshape, "Rendered3DShape")
}

func (rhombusshape *RhombusShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RhombusShape_stagedOrder, stage.RhombusShapes_referenceOrder, rhombusshape, "RhombusShape")
}

func (rhombusstuff *RhombusStuff) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RhombusStuff_stagedOrder, stage.RhombusStuffs_referenceOrder, rhombusstuff, "RhombusStuff")
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RotatedRhombusGridShape_stagedOrder, stage.RotatedRhombusGridShapes_referenceOrder, rotatedrhombusgridshape, "RotatedRhombusGridShape")
}

func (rotatedrhombusshape *RotatedRhombusShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RotatedRhombusShape_stagedOrder, stage.RotatedRhombusShapes_referenceOrder, rotatedrhombusshape, "RotatedRhombusShape")
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RotatedSampledPoints3DShape_stagedOrder, stage.RotatedSampledPoints3DShapes_referenceOrder, rotatedsampledpoints3dshape, "RotatedSampledPoints3DShape")
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RotatedSeatAndLegs3DShape_stagedOrder, stage.RotatedSeatAndLegs3DShapes_referenceOrder, rotatedseatandlegs3dshape, "RotatedSeatAndLegs3DShape")
}

func (sampledpoints3dshape *SampledPoints3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SampledPoints3DShape_stagedOrder, stage.SampledPoints3DShapes_referenceOrder, sampledpoints3dshape, "SampledPoints3DShape")
}

func (seat3dshape *Seat3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Seat3DShape_stagedOrder, stage.Seat3DShapes_referenceOrder, seat3dshape, "Seat3DShape")
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SeatAndLegs3DShape_stagedOrder, stage.SeatAndLegs3DShapes_referenceOrder, seatandlegs3dshape, "SeatAndLegs3DShape")
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SeatBottomCurveShape_stagedOrder, stage.SeatBottomCurveShapes_referenceOrder, seatbottomcurveshape, "SeatBottomCurveShape")
}

func (seattopcurveshape *SeatTopCurveShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SeatTopCurveShape_stagedOrder, stage.SeatTopCurveShapes_referenceOrder, seattopcurveshape, "SeatTopCurveShape")
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedBottomTopStartArcShape_stagedOrder, stage.ShiftedBottomTopStartArcShapes_referenceOrder, shiftedbottomtopstartarcshape, "ShiftedBottomTopStartArcShape")
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedBottomTopStartArcShapeGrid_stagedOrder, stage.ShiftedBottomTopStartArcShapeGrids_referenceOrder, shiftedbottomtopstartarcshapegrid, "ShiftedBottomTopStartArcShapeGrid")
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedLeftGrowthCurve2DRibbon_stagedOrder, stage.ShiftedLeftGrowthCurve2DRibbons_referenceOrder, shiftedleftgrowthcurve2dribbon, "ShiftedLeftGrowthCurve2DRibbon")
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedLeftGrowthCurve2DRibbonEndShape_stagedOrder, stage.ShiftedLeftGrowthCurve2DRibbonEndShapes_referenceOrder, shiftedleftgrowthcurve2dribbonendshape, "ShiftedLeftGrowthCurve2DRibbonEndShape")
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedLeftGrowthCurve2DRibbonStartShape_stagedOrder, stage.ShiftedLeftGrowthCurve2DRibbonStartShapes_referenceOrder, shiftedleftgrowthcurve2dribbonstartshape, "ShiftedLeftGrowthCurve2DRibbonStartShape")
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedLeftPartiallyGrowthCurve2DRibbon_stagedOrder, stage.ShiftedLeftPartiallyGrowthCurve2DRibbons_referenceOrder, shiftedleftpartiallygrowthcurve2dribbon, "ShiftedLeftPartiallyGrowthCurve2DRibbon")
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape_stagedOrder, stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes_referenceOrder, shiftedleftpartiallygrowthcurve2dribbonendshape, "ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape")
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape_stagedOrder, stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes_referenceOrder, shiftedleftpartiallygrowthcurve2dribbonstartshape, "ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape")
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedLeftStackGrowthCurveEndArcShape_stagedOrder, stage.ShiftedLeftStackGrowthCurveEndArcShapes_referenceOrder, shiftedleftstackgrowthcurveendarcshape, "ShiftedLeftStackGrowthCurveEndArcShape")
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedLeftStackGrowthCurveStartArcShape_stagedOrder, stage.ShiftedLeftStackGrowthCurveStartArcShapes_referenceOrder, shiftedleftstackgrowthcurvestartarcshape, "ShiftedLeftStackGrowthCurveStartArcShape")
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedLeftStackNormalVector_stagedOrder, stage.ShiftedLeftStackNormalVectors_referenceOrder, shiftedleftstacknormalvector, "ShiftedLeftStackNormalVector")
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedLeftStackOfGrowthCurve_stagedOrder, stage.ShiftedLeftStackOfGrowthCurves_referenceOrder, shiftedleftstackofgrowthcurve, "ShiftedLeftStackOfGrowthCurve")
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedLeftStackOfNormalVector_stagedOrder, stage.ShiftedLeftStackOfNormalVectors_referenceOrder, shiftedleftstackofnormalvector, "ShiftedLeftStackOfNormalVector")
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedRightGrowthCurve2DRibbon_stagedOrder, stage.ShiftedRightGrowthCurve2DRibbons_referenceOrder, shiftedrightgrowthcurve2dribbon, "ShiftedRightGrowthCurve2DRibbon")
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedRightGrowthCurve2DRibbonEndShape_stagedOrder, stage.ShiftedRightGrowthCurve2DRibbonEndShapes_referenceOrder, shiftedrightgrowthcurve2dribbonendshape, "ShiftedRightGrowthCurve2DRibbonEndShape")
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ShiftedRightGrowthCurve2DRibbonStartShape_stagedOrder, stage.ShiftedRightGrowthCurve2DRibbonStartShapes_referenceOrder, shiftedrightgrowthcurve2dribbonstartshape, "ShiftedRightGrowthCurve2DRibbonStartShape")
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackGrowthCurve2DEndHalfwayArcShape_stagedOrder, stage.StackGrowthCurve2DEndHalfwayArcShapes_referenceOrder, stackgrowthcurve2dendhalfwayarcshape, "StackGrowthCurve2DEndHalfwayArcShape")
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackGrowthCurve2DRibbonEndShape_stagedOrder, stage.StackGrowthCurve2DRibbonEndShapes_referenceOrder, stackgrowthcurve2dribbonendshape, "StackGrowthCurve2DRibbonEndShape")
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackGrowthCurve2DRibbonStartShape_stagedOrder, stage.StackGrowthCurve2DRibbonStartShapes_referenceOrder, stackgrowthcurve2dribbonstartshape, "StackGrowthCurve2DRibbonStartShape")
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackGrowthCurve2DStartHalfwayArcShape_stagedOrder, stage.StackGrowthCurve2DStartHalfwayArcShapes_referenceOrder, stackgrowthcurve2dstarthalfwayarcshape, "StackGrowthCurve2DStartHalfwayArcShape")
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackOfGrowthCurve2D_stagedOrder, stage.StackOfGrowthCurve2Ds_referenceOrder, stackofgrowthcurve2d, "StackOfGrowthCurve2D")
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackOfGrowthCurve2DByGrowthVector_stagedOrder, stage.StackOfGrowthCurve2DByGrowthVectors_referenceOrder, stackofgrowthcurve2dbygrowthvector, "StackOfGrowthCurve2DByGrowthVector")
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackOfGrowthCurve2DRibbon_stagedOrder, stage.StackOfGrowthCurve2DRibbons_referenceOrder, stackofgrowthcurve2dribbon, "StackOfGrowthCurve2DRibbon")
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackOfPartiallyRotatedTorusShape_stagedOrder, stage.StackOfPartiallyRotatedTorusShapes_referenceOrder, stackofpartiallyrotatedtorusshape, "StackOfPartiallyRotatedTorusShape")
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackOfRotatedGrowthCurve2D_stagedOrder, stage.StackOfRotatedGrowthCurve2Ds_referenceOrder, stackofrotatedgrowthcurve2d, "StackOfRotatedGrowthCurve2D")
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackOfRotatedGrowthCurve2DRibbon_stagedOrder, stage.StackOfRotatedGrowthCurve2DRibbons_referenceOrder, stackofrotatedgrowthcurve2dribbon, "StackOfRotatedGrowthCurve2DRibbon")
}

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackOfRotatedVaseTrapezeRingsShape_stagedOrder, stage.StackOfRotatedVaseTrapezeRingsShapes_referenceOrder, stackofrotatedvasetrapezeringsshape, "StackOfRotatedVaseTrapezeRingsShape")
}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackOfVaseTrapezeRingsShape_stagedOrder, stage.StackOfVaseTrapezeRingsShapes_referenceOrder, stackofvasetrapezeringsshape, "StackOfVaseTrapezeRingsShape")
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackRotatedGrowthCurve2DEndArcShape_stagedOrder, stage.StackRotatedGrowthCurve2DEndArcShapes_referenceOrder, stackrotatedgrowthcurve2dendarcshape, "StackRotatedGrowthCurve2DEndArcShape")
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackRotatedGrowthCurve2DRibbonEndShape_stagedOrder, stage.StackRotatedGrowthCurve2DRibbonEndShapes_referenceOrder, stackrotatedgrowthcurve2dribbonendshape, "StackRotatedGrowthCurve2DRibbonEndShape")
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackRotatedGrowthCurve2DRibbonStartShape_stagedOrder, stage.StackRotatedGrowthCurve2DRibbonStartShapes_referenceOrder, stackrotatedgrowthcurve2dribbonstartshape, "StackRotatedGrowthCurve2DRibbonStartShape")
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StackRotatedGrowthCurve2DStartArcShape_stagedOrder, stage.StackRotatedGrowthCurve2DStartArcShapes_referenceOrder, stackrotatedgrowthcurve2dstartarcshape, "StackRotatedGrowthCurve2DStartArcShape")
}

func (startarcshape *StartArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StartArcShape_stagedOrder, stage.StartArcShapes_referenceOrder, startarcshape, "StartArcShape")
}

func (startarcshapegrid *StartArcShapeGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StartArcShapeGrid_stagedOrder, stage.StartArcShapeGrids_referenceOrder, startarcshapegrid, "StartArcShapeGrid")
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StartHalfwayArcShape_stagedOrder, stage.StartHalfwayArcShapes_referenceOrder, starthalfwayarcshape, "StartHalfwayArcShape")
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StartHalfwayArcShapeGrid_stagedOrder, stage.StartHalfwayArcShapeGrids_referenceOrder, starthalfwayarcshapegrid, "StartHalfwayArcShapeGrid")
}

func (stemcylinder3dshape *StemCylinder3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StemCylinder3DShape_stagedOrder, stage.StemCylinder3DShapes_referenceOrder, stemcylinder3dshape, "StemCylinder3DShape")
}

func (stool2ddiagram *Stool2DDiagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Stool2DDiagram_stagedOrder, stage.Stool2DDiagrams_referenceOrder, stool2ddiagram, "Stool2DDiagram")
}

func (stool3ddiagram *Stool3DDiagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Stool3DDiagram_stagedOrder, stage.Stool3DDiagrams_referenceOrder, stool3ddiagram, "Stool3DDiagram")
}

func (tiledfloor3dshape *TiledFloor3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TiledFloor3DShape_stagedOrder, stage.TiledFloor3DShapes_referenceOrder, tiledfloor3dshape, "TiledFloor3DShape")
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopCurvePlane1Shape_stagedOrder, stage.TopCurvePlane1Shapes_referenceOrder, topcurveplane1shape, "TopCurvePlane1Shape")
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopCurvePlane2Shape_stagedOrder, stage.TopCurvePlane2Shapes_referenceOrder, topcurveplane2shape, "TopCurvePlane2Shape")
}

func (topendarcshape *TopEndArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopEndArcShape_stagedOrder, stage.TopEndArcShapes_referenceOrder, topendarcshape, "TopEndArcShape")
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopEndArcShapeGrid_stagedOrder, stage.TopEndArcShapeGrids_referenceOrder, topendarcshapegrid, "TopEndArcShapeGrid")
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopEndHalfwayArcShape_stagedOrder, stage.TopEndHalfwayArcShapes_referenceOrder, topendhalfwayarcshape, "TopEndHalfwayArcShape")
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopEndHalfwayArcShapeGrid_stagedOrder, stage.TopEndHalfwayArcShapeGrids_referenceOrder, topendhalfwayarcshapegrid, "TopEndHalfwayArcShapeGrid")
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopGrowthCurve2D_stagedOrder, stage.TopGrowthCurve2Ds_referenceOrder, topgrowthcurve2d, "TopGrowthCurve2D")
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopMidArcVectorShape_stagedOrder, stage.TopMidArcVectorShapes_referenceOrder, topmidarcvectorshape, "TopMidArcVectorShape")
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopMidArcVectorShapeGrid_stagedOrder, stage.TopMidArcVectorShapeGrids_referenceOrder, topmidarcvectorshapegrid, "TopMidArcVectorShapeGrid")
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopStackGrowthCurve2DEndHalfwayArcShape_stagedOrder, stage.TopStackGrowthCurve2DEndHalfwayArcShapes_referenceOrder, topstackgrowthcurve2dendhalfwayarcshape, "TopStackGrowthCurve2DEndHalfwayArcShape")
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopStackGrowthCurve2DStartHalfwayArcShape_stagedOrder, stage.TopStackGrowthCurve2DStartHalfwayArcShapes_referenceOrder, topstackgrowthcurve2dstarthalfwayarcshape, "TopStackGrowthCurve2DStartHalfwayArcShape")
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopStackOfGrowthCurve2D_stagedOrder, stage.TopStackOfGrowthCurve2Ds_referenceOrder, topstackofgrowthcurve2d, "TopStackOfGrowthCurve2D")
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopStackOfRotatedGrowthCurve2D_stagedOrder, stage.TopStackOfRotatedGrowthCurve2Ds_referenceOrder, topstackofrotatedgrowthcurve2d, "TopStackOfRotatedGrowthCurve2D")
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopStackOfRotatedGrowthCurve2DEndArcShape_stagedOrder, stage.TopStackOfRotatedGrowthCurve2DEndArcShapes_referenceOrder, topstackofrotatedgrowthcurve2dendarcshape, "TopStackOfRotatedGrowthCurve2DEndArcShape")
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopStackOfRotatedGrowthCurve2DStartArcShape_stagedOrder, stage.TopStackOfRotatedGrowthCurve2DStartArcShapes_referenceOrder, topstackofrotatedgrowthcurve2dstartarcshape, "TopStackOfRotatedGrowthCurve2DStartArcShape")
}

func (topstartarcshape *TopStartArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopStartArcShape_stagedOrder, stage.TopStartArcShapes_referenceOrder, topstartarcshape, "TopStartArcShape")
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopStartArcShapeGrid_stagedOrder, stage.TopStartArcShapeGrids_referenceOrder, topstartarcshapegrid, "TopStartArcShapeGrid")
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopStartHalfwayArcShape_stagedOrder, stage.TopStartHalfwayArcShapes_referenceOrder, topstarthalfwayarcshape, "TopStartHalfwayArcShape")
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TopStartHalfwayArcShapeGrid_stagedOrder, stage.TopStartHalfwayArcShapeGrids_referenceOrder, topstarthalfwayarcshapegrid, "TopStartHalfwayArcShapeGrid")
}

func (torus3dshape *Torus3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Torus3DShape_stagedOrder, stage.Torus3DShapes_referenceOrder, torus3dshape, "Torus3DShape")
}

func (torusedge3dshape *TorusEdge3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TorusEdge3DShape_stagedOrder, stage.TorusEdge3DShapes_referenceOrder, torusedge3dshape, "TorusEdge3DShape")
}

func (torusstackshape *TorusStackShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TorusStackShape_stagedOrder, stage.TorusStackShapes_referenceOrder, torusstackshape, "TorusStackShape")
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TubeVase3DDiagram_stagedOrder, stage.TubeVase3DDiagrams_referenceOrder, tubevase3ddiagram, "TubeVase3DDiagram")
}

func (tubevaseabstract *TubeVaseAbstract) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TubeVaseAbstract_stagedOrder, stage.TubeVaseAbstracts_referenceOrder, tubevaseabstract, "TubeVaseAbstract")
}

func (vase2ddiagram *Vase2DDiagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Vase2DDiagram_stagedOrder, stage.Vase2DDiagrams_referenceOrder, vase2ddiagram, "Vase2DDiagram")
}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VaseTrapezeRingShape_stagedOrder, stage.VaseTrapezeRingShapes_referenceOrder, vasetrapezeringshape, "VaseTrapezeRingShape")
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VerticalTorusStackShape_stagedOrder, stage.VerticalTorusStackShapes_referenceOrder, verticaltorusstackshape, "VerticalTorusStackShape")
}

func (volumekey3dshape *VolumeKey3DShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VolumeKey3DShape_stagedOrder, stage.VolumeKey3DShapes_referenceOrder, volumekey3dshape, "VolumeKey3DShape")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (angle0shape *Angle0Shape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(angle0shape, angle0shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (angle0shape *Angle0Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return angle0shape.GongGetIdentifier(stage)
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(arcnormalvectorshape, arcnormalvectorshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arcnormalvectorshape *ArcNormalVectorShape) GongGetReferenceIdentifier(stage *Stage) string {
	return arcnormalvectorshape.GongGetIdentifier(stage)
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(arcnormalvectorshapegrid, arcnormalvectorshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return arcnormalvectorshapegrid.GongGetIdentifier(stage)
}

func (axesshape *AxesShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(axesshape, axesshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (axesshape *AxesShape) GongGetReferenceIdentifier(stage *Stage) string {
	return axesshape.GongGetIdentifier(stage)
}

func (basevectorshape *BaseVectorShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(basevectorshape, basevectorshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (basevectorshape *BaseVectorShape) GongGetReferenceIdentifier(stage *Stage) string {
	return basevectorshape.GongGetIdentifier(stage)
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(basevectorshapegrid, basevectorshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (basevectorshapegrid *BaseVectorShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return basevectorshapegrid.GongGetIdentifier(stage)
}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bottomcurveplane1shape, bottomcurveplane1shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return bottomcurveplane1shape.GongGetIdentifier(stage)
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bottomcurveplane2shape, bottomcurveplane2shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return bottomcurveplane2shape.GongGetIdentifier(stage)
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(chosenp1p2pairshape, chosenp1p2pairshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (chosenp1p2pairshape *ChosenP1P2PairShape) GongGetReferenceIdentifier(stage *Stage) string {
	return chosenp1p2pairshape.GongGetIdentifier(stage)
}

func (circlegridshape *CircleGridShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(circlegridshape, circlegridshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circlegridshape *CircleGridShape) GongGetReferenceIdentifier(stage *Stage) string {
	return circlegridshape.GongGetIdentifier(stage)
}

func (circumference3dshape *Circumference3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(circumference3dshape, circumference3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (circumference3dshape *Circumference3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return circumference3dshape.GongGetIdentifier(stage)
}

func (clock2ddiagram *Clock2DDiagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(clock2ddiagram, clock2ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (clock2ddiagram *Clock2DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return clock2ddiagram.GongGetIdentifier(stage)
}

func (clock3ddiagram *Clock3DDiagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(clock3ddiagram, clock3ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (clock3ddiagram *Clock3DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return clock3ddiagram.GongGetIdentifier(stage)
}

func (clocktopcurveshape *ClockTopCurveShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(clocktopcurveshape, clocktopcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (clocktopcurveshape *ClockTopCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return clocktopcurveshape.GongGetIdentifier(stage)
}

func (cutline3dshape *CutLine3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(cutline3dshape, cutline3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cutline3dshape *CutLine3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return cutline3dshape.GongGetIdentifier(stage)
}

func (endarcshape *EndArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(endarcshape, endarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (endarcshape *EndArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return endarcshape.GongGetIdentifier(stage)
}

func (endarcshapegrid *EndArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(endarcshapegrid, endarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (endarcshapegrid *EndArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return endarcshapegrid.GongGetIdentifier(stage)
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(endhalfwayarcshape, endhalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (endhalfwayarcshape *EndHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return endhalfwayarcshape.GongGetIdentifier(stage)
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(endhalfwayarcshapegrid, endhalfwayarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return endhalfwayarcshapegrid.GongGetIdentifier(stage)
}

func (explanationtextshape *ExplanationTextShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(explanationtextshape, explanationtextshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (explanationtextshape *ExplanationTextShape) GongGetReferenceIdentifier(stage *Stage) string {
	return explanationtextshape.GongGetIdentifier(stage)
}

func (eye3dshape *Eye3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(eye3dshape, eye3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (eye3dshape *Eye3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return eye3dshape.GongGetIdentifier(stage)
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(eyecornerssampledpoints3dshape, eyecornerssampledpoints3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return eyecornerssampledpoints3dshape.GongGetIdentifier(stage)
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(eyesampledpoints3dshape, eyesampledpoints3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return eyesampledpoints3dshape.GongGetIdentifier(stage)
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(eyeseatbottomcurveshape, eyeseatbottomcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return eyeseatbottomcurveshape.GongGetIdentifier(stage)
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(eyestoolbottomcurveshape, eyestoolbottomcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return eyestoolbottomcurveshape.GongGetIdentifier(stage)
}

func (eyevolume3dshape *EyeVolume3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(eyevolume3dshape, eyevolume3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (eyevolume3dshape *EyeVolume3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return eyevolume3dshape.GongGetIdentifier(stage)
}

func (gridpathshape *GridPathShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gridpathshape, gridpathshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gridpathshape *GridPathShape) GongGetReferenceIdentifier(stage *Stage) string {
	return gridpathshape.GongGetIdentifier(stage)
}

func (growthcurve2d *GrowthCurve2D) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(growthcurve2d, growthcurve2d.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurve2d *GrowthCurve2D) GongGetReferenceIdentifier(stage *Stage) string {
	return growthcurve2d.GongGetIdentifier(stage)
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(growthcurve2dribbon, growthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurve2dribbon *GrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return growthcurve2dribbon.GongGetIdentifier(stage)
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(growthcurve2dribbonendshape, growthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return growthcurve2dribbonendshape.GongGetIdentifier(stage)
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(growthcurve2dribbonstartshape, growthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return growthcurve2dribbonstartshape.GongGetIdentifier(stage)
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(growthcurverhombusgridshape, growthcurverhombusgridshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongGetReferenceIdentifier(stage *Stage) string {
	return growthcurverhombusgridshape.GongGetIdentifier(stage)
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(growthcurverhombusshape, growthcurverhombusshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthcurverhombusshape *GrowthCurveRhombusShape) GongGetReferenceIdentifier(stage *Stage) string {
	return growthcurverhombusshape.GongGetIdentifier(stage)
}

func (growthvectorshape *GrowthVectorShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(growthvectorshape, growthvectorshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (growthvectorshape *GrowthVectorShape) GongGetReferenceIdentifier(stage *Stage) string {
	return growthvectorshape.GongGetIdentifier(stage)
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(initialrhombusgridshape, initialrhombusgridshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (initialrhombusgridshape *InitialRhombusGridShape) GongGetReferenceIdentifier(stage *Stage) string {
	return initialrhombusgridshape.GongGetIdentifier(stage)
}

func (initialrhombusshape *InitialRhombusShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(initialrhombusshape, initialrhombusshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (initialrhombusshape *InitialRhombusShape) GongGetReferenceIdentifier(stage *Stage) string {
	return initialrhombusshape.GongGetIdentifier(stage)
}

func (key3dshape *Key3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(key3dshape, key3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (key3dshape *Key3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return key3dshape.GongGetIdentifier(stage)
}

func (keyhole3dshape *KeyHole3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(keyhole3dshape, keyhole3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (keyhole3dshape *KeyHole3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return keyhole3dshape.GongGetIdentifier(stage)
}

func (keyholeshape *KeyHoleShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(keyholeshape, keyholeshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (keyholeshape *KeyHoleShape) GongGetReferenceIdentifier(stage *Stage) string {
	return keyholeshape.GongGetIdentifier(stage)
}

func (leaves3dshape *Leaves3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(leaves3dshape, leaves3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (leaves3dshape *Leaves3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return leaves3dshape.GongGetIdentifier(stage)
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(library, library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return library.GongGetIdentifier(stage)
}

func (midarcvectorshape *MidArcVectorShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(midarcvectorshape, midarcvectorshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (midarcvectorshape *MidArcVectorShape) GongGetReferenceIdentifier(stage *Stage) string {
	return midarcvectorshape.GongGetIdentifier(stage)
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(midarcvectorshapegrid, midarcvectorshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return midarcvectorshapegrid.GongGetIdentifier(stage)
}

func (originalpoints3dshape *OriginalPoints3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(originalpoints3dshape, originalpoints3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (originalpoints3dshape *OriginalPoints3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return originalpoints3dshape.GongGetIdentifier(stage)
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(parastichymcurves3dshape, parastichymcurves3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return parastichymcurves3dshape.GongGetIdentifier(stage)
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(parastichyncurves3dshape, parastichyncurves3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return parastichyncurves3dshape.GongGetIdentifier(stage)
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallygrowthcurve2dribbon, partiallygrowthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallygrowthcurve2dribbon.GongGetIdentifier(stage)
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallygrowthcurve2dribbonendshape, partiallygrowthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallygrowthcurve2dribbonendshape.GongGetIdentifier(stage)
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallygrowthcurve2dribbonstartshape, partiallygrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallygrowthcurve2dribbonstartshape.GongGetIdentifier(stage)
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallygrowthcurve2dtrajectory, partiallygrowthcurve2dtrajectory.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallygrowthcurve2dtrajectory.GongGetIdentifier(stage)
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallygrowthcurve2dtrajectoryp1curveshape, partiallygrowthcurve2dtrajectoryp1curveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallygrowthcurve2dtrajectoryp1curveshape.GongGetIdentifier(stage)
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallygrowthcurve2dtrajectoryp1p2, partiallygrowthcurve2dtrajectoryp1p2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallygrowthcurve2dtrajectoryp1p2.GongGetIdentifier(stage)
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallygrowthcurve2dtrajectoryp1p2pairlineshape, partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongGetIdentifier(stage)
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallygrowthcurve2dtrajectoryp1pointshape, partiallygrowthcurve2dtrajectoryp1pointshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallygrowthcurve2dtrajectoryp1pointshape.GongGetIdentifier(stage)
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallygrowthcurve2dtrajectoryp2curveshape, partiallygrowthcurve2dtrajectoryp2curveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallygrowthcurve2dtrajectoryp2curveshape.GongGetIdentifier(stage)
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallygrowthcurve2dtrajectoryp2pointshape, partiallygrowthcurve2dtrajectoryp2pointshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallygrowthcurve2dtrajectoryp2pointshape.GongGetIdentifier(stage)
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallygrowthcurve2dtrajectoryshape, partiallygrowthcurve2dtrajectoryshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallygrowthcurve2dtrajectoryshape.GongGetIdentifier(stage)
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallyrotatedseatbottomcurveshape, partiallyrotatedseatbottomcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallyrotatedseatbottomcurveshape.GongGetIdentifier(stage)
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallyrotatedseattopcurveshape, partiallyrotatedseattopcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallyrotatedseattopcurveshape.GongGetIdentifier(stage)
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partiallyrotatedtorusshape, partiallyrotatedtorusshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongGetReferenceIdentifier(stage *Stage) string {
	return partiallyrotatedtorusshape.GongGetIdentifier(stage)
}

func (perpendicularvector *PerpendicularVector) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(perpendicularvector, perpendicularvector.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (perpendicularvector *PerpendicularVector) GongGetReferenceIdentifier(stage *Stage) string {
	return perpendicularvector.GongGetIdentifier(stage)
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(perpendicularvectorgrid, perpendicularvectorgrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (perpendicularvectorgrid *PerpendicularVectorGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return perpendicularvectorgrid.GongGetIdentifier(stage)
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(perpendicularvectorgridhalfway, perpendicularvectorgridhalfway.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongGetReferenceIdentifier(stage *Stage) string {
	return perpendicularvectorgridhalfway.GongGetIdentifier(stage)
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(perpendicularvectorhalfway, perpendicularvectorhalfway.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongGetReferenceIdentifier(stage *Stage) string {
	return perpendicularvectorhalfway.GongGetIdentifier(stage)
}

func (plant2ddiagram *Plant2DDiagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(plant2ddiagram, plant2ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (plant2ddiagram *Plant2DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return plant2ddiagram.GongGetIdentifier(stage)
}

func (plant3ddiagram *Plant3DDiagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(plant3ddiagram, plant3ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (plant3ddiagram *Plant3DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return plant3ddiagram.GongGetIdentifier(stage)
}

func (plantabstract *PlantAbstract) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(plantabstract, plantabstract.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (plantabstract *PlantAbstract) GongGetReferenceIdentifier(stage *Stage) string {
	return plantabstract.GongGetIdentifier(stage)
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(plantcircumferenceshape, plantcircumferenceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (plantcircumferenceshape *PlantCircumferenceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return plantcircumferenceshape.GongGetIdentifier(stage)
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pointsandlines3dshape, pointsandlines3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointsandlines3dshape *PointsAndLines3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return pointsandlines3dshape.GongGetIdentifier(stage)
}

func (pxshape *PxShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pxshape, pxshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pxshape *PxShape) GongGetReferenceIdentifier(stage *Stage) string {
	return pxshape.GongGetIdentifier(stage)
}

func (rendered3dshape *Rendered3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rendered3dshape, rendered3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rendered3dshape *Rendered3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return rendered3dshape.GongGetIdentifier(stage)
}

func (rhombusshape *RhombusShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rhombusshape, rhombusshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rhombusshape *RhombusShape) GongGetReferenceIdentifier(stage *Stage) string {
	return rhombusshape.GongGetIdentifier(stage)
}

func (rhombusstuff *RhombusStuff) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rhombusstuff, rhombusstuff.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rhombusstuff *RhombusStuff) GongGetReferenceIdentifier(stage *Stage) string {
	return rhombusstuff.GongGetIdentifier(stage)
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rotatedrhombusgridshape, rotatedrhombusgridshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongGetReferenceIdentifier(stage *Stage) string {
	return rotatedrhombusgridshape.GongGetIdentifier(stage)
}

func (rotatedrhombusshape *RotatedRhombusShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rotatedrhombusshape, rotatedrhombusshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rotatedrhombusshape *RotatedRhombusShape) GongGetReferenceIdentifier(stage *Stage) string {
	return rotatedrhombusshape.GongGetIdentifier(stage)
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rotatedsampledpoints3dshape, rotatedsampledpoints3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return rotatedsampledpoints3dshape.GongGetIdentifier(stage)
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rotatedseatandlegs3dshape, rotatedseatandlegs3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return rotatedseatandlegs3dshape.GongGetIdentifier(stage)
}

func (sampledpoints3dshape *SampledPoints3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(sampledpoints3dshape, sampledpoints3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sampledpoints3dshape *SampledPoints3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return sampledpoints3dshape.GongGetIdentifier(stage)
}

func (seat3dshape *Seat3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(seat3dshape, seat3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (seat3dshape *Seat3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return seat3dshape.GongGetIdentifier(stage)
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(seatandlegs3dshape, seatandlegs3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (seatandlegs3dshape *SeatAndLegs3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return seatandlegs3dshape.GongGetIdentifier(stage)
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(seatbottomcurveshape, seatbottomcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (seatbottomcurveshape *SeatBottomCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return seatbottomcurveshape.GongGetIdentifier(stage)
}

func (seattopcurveshape *SeatTopCurveShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(seattopcurveshape, seattopcurveshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (seattopcurveshape *SeatTopCurveShape) GongGetReferenceIdentifier(stage *Stage) string {
	return seattopcurveshape.GongGetIdentifier(stage)
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedbottomtopstartarcshape, shiftedbottomtopstartarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedbottomtopstartarcshape.GongGetIdentifier(stage)
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedbottomtopstartarcshapegrid, shiftedbottomtopstartarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedbottomtopstartarcshapegrid.GongGetIdentifier(stage)
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedleftgrowthcurve2dribbon, shiftedleftgrowthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedleftgrowthcurve2dribbon.GongGetIdentifier(stage)
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedleftgrowthcurve2dribbonendshape, shiftedleftgrowthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedleftgrowthcurve2dribbonendshape.GongGetIdentifier(stage)
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedleftgrowthcurve2dribbonstartshape, shiftedleftgrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedleftgrowthcurve2dribbonstartshape.GongGetIdentifier(stage)
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedleftpartiallygrowthcurve2dribbon, shiftedleftpartiallygrowthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedleftpartiallygrowthcurve2dribbon.GongGetIdentifier(stage)
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedleftpartiallygrowthcurve2dribbonendshape, shiftedleftpartiallygrowthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedleftpartiallygrowthcurve2dribbonendshape.GongGetIdentifier(stage)
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedleftpartiallygrowthcurve2dribbonstartshape, shiftedleftpartiallygrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedleftpartiallygrowthcurve2dribbonstartshape.GongGetIdentifier(stage)
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedleftstackgrowthcurveendarcshape, shiftedleftstackgrowthcurveendarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedleftstackgrowthcurveendarcshape.GongGetIdentifier(stage)
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedleftstackgrowthcurvestartarcshape, shiftedleftstackgrowthcurvestartarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedleftstackgrowthcurvestartarcshape.GongGetIdentifier(stage)
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedleftstacknormalvector, shiftedleftstacknormalvector.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedleftstacknormalvector.GongGetIdentifier(stage)
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedleftstackofgrowthcurve, shiftedleftstackofgrowthcurve.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedleftstackofgrowthcurve.GongGetIdentifier(stage)
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedleftstackofnormalvector, shiftedleftstackofnormalvector.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedleftstackofnormalvector.GongGetIdentifier(stage)
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedrightgrowthcurve2dribbon, shiftedrightgrowthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedrightgrowthcurve2dribbon.GongGetIdentifier(stage)
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedrightgrowthcurve2dribbonendshape, shiftedrightgrowthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedrightgrowthcurve2dribbonendshape.GongGetIdentifier(stage)
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shiftedrightgrowthcurve2dribbonstartshape, shiftedrightgrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return shiftedrightgrowthcurve2dribbonstartshape.GongGetIdentifier(stage)
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackgrowthcurve2dendhalfwayarcshape, stackgrowthcurve2dendhalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stackgrowthcurve2dendhalfwayarcshape.GongGetIdentifier(stage)
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackgrowthcurve2dribbonendshape, stackgrowthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stackgrowthcurve2dribbonendshape.GongGetIdentifier(stage)
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackgrowthcurve2dribbonstartshape, stackgrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stackgrowthcurve2dribbonstartshape.GongGetIdentifier(stage)
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackgrowthcurve2dstarthalfwayarcshape, stackgrowthcurve2dstarthalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stackgrowthcurve2dstarthalfwayarcshape.GongGetIdentifier(stage)
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackofgrowthcurve2d, stackofgrowthcurve2d.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongGetReferenceIdentifier(stage *Stage) string {
	return stackofgrowthcurve2d.GongGetIdentifier(stage)
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackofgrowthcurve2dbygrowthvector, stackofgrowthcurve2dbygrowthvector.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongGetReferenceIdentifier(stage *Stage) string {
	return stackofgrowthcurve2dbygrowthvector.GongGetIdentifier(stage)
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackofgrowthcurve2dribbon, stackofgrowthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return stackofgrowthcurve2dribbon.GongGetIdentifier(stage)
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackofpartiallyrotatedtorusshape, stackofpartiallyrotatedtorusshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stackofpartiallyrotatedtorusshape.GongGetIdentifier(stage)
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackofrotatedgrowthcurve2d, stackofrotatedgrowthcurve2d.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongGetReferenceIdentifier(stage *Stage) string {
	return stackofrotatedgrowthcurve2d.GongGetIdentifier(stage)
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackofrotatedgrowthcurve2dribbon, stackofrotatedgrowthcurve2dribbon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongGetReferenceIdentifier(stage *Stage) string {
	return stackofrotatedgrowthcurve2dribbon.GongGetIdentifier(stage)
}

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackofrotatedvasetrapezeringsshape, stackofrotatedvasetrapezeringsshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stackofrotatedvasetrapezeringsshape.GongGetIdentifier(stage)
}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackofvasetrapezeringsshape, stackofvasetrapezeringsshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stackofvasetrapezeringsshape.GongGetIdentifier(stage)
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackrotatedgrowthcurve2dendarcshape, stackrotatedgrowthcurve2dendarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stackrotatedgrowthcurve2dendarcshape.GongGetIdentifier(stage)
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackrotatedgrowthcurve2dribbonendshape, stackrotatedgrowthcurve2dribbonendshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stackrotatedgrowthcurve2dribbonendshape.GongGetIdentifier(stage)
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackrotatedgrowthcurve2dribbonstartshape, stackrotatedgrowthcurve2dribbonstartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stackrotatedgrowthcurve2dribbonstartshape.GongGetIdentifier(stage)
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stackrotatedgrowthcurve2dstartarcshape, stackrotatedgrowthcurve2dstartarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stackrotatedgrowthcurve2dstartarcshape.GongGetIdentifier(stage)
}

func (startarcshape *StartArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(startarcshape, startarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (startarcshape *StartArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return startarcshape.GongGetIdentifier(stage)
}

func (startarcshapegrid *StartArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(startarcshapegrid, startarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (startarcshapegrid *StartArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return startarcshapegrid.GongGetIdentifier(stage)
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(starthalfwayarcshape, starthalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (starthalfwayarcshape *StartHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return starthalfwayarcshape.GongGetIdentifier(stage)
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(starthalfwayarcshapegrid, starthalfwayarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return starthalfwayarcshapegrid.GongGetIdentifier(stage)
}

func (stemcylinder3dshape *StemCylinder3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stemcylinder3dshape, stemcylinder3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stemcylinder3dshape *StemCylinder3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stemcylinder3dshape.GongGetIdentifier(stage)
}

func (stool2ddiagram *Stool2DDiagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stool2ddiagram, stool2ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stool2ddiagram *Stool2DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return stool2ddiagram.GongGetIdentifier(stage)
}

func (stool3ddiagram *Stool3DDiagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stool3ddiagram, stool3ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stool3ddiagram *Stool3DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return stool3ddiagram.GongGetIdentifier(stage)
}

func (tiledfloor3dshape *TiledFloor3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tiledfloor3dshape, tiledfloor3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tiledfloor3dshape *TiledFloor3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return tiledfloor3dshape.GongGetIdentifier(stage)
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topcurveplane1shape, topcurveplane1shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topcurveplane1shape *TopCurvePlane1Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return topcurveplane1shape.GongGetIdentifier(stage)
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topcurveplane2shape, topcurveplane2shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topcurveplane2shape *TopCurvePlane2Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return topcurveplane2shape.GongGetIdentifier(stage)
}

func (topendarcshape *TopEndArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topendarcshape, topendarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topendarcshape *TopEndArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return topendarcshape.GongGetIdentifier(stage)
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topendarcshapegrid, topendarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topendarcshapegrid *TopEndArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return topendarcshapegrid.GongGetIdentifier(stage)
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topendhalfwayarcshape, topendhalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return topendhalfwayarcshape.GongGetIdentifier(stage)
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topendhalfwayarcshapegrid, topendhalfwayarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return topendhalfwayarcshapegrid.GongGetIdentifier(stage)
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topgrowthcurve2d, topgrowthcurve2d.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topgrowthcurve2d *TopGrowthCurve2D) GongGetReferenceIdentifier(stage *Stage) string {
	return topgrowthcurve2d.GongGetIdentifier(stage)
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topmidarcvectorshape, topmidarcvectorshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topmidarcvectorshape *TopMidArcVectorShape) GongGetReferenceIdentifier(stage *Stage) string {
	return topmidarcvectorshape.GongGetIdentifier(stage)
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topmidarcvectorshapegrid, topmidarcvectorshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return topmidarcvectorshapegrid.GongGetIdentifier(stage)
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topstackgrowthcurve2dendhalfwayarcshape, topstackgrowthcurve2dendhalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return topstackgrowthcurve2dendhalfwayarcshape.GongGetIdentifier(stage)
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topstackgrowthcurve2dstarthalfwayarcshape, topstackgrowthcurve2dstarthalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return topstackgrowthcurve2dstarthalfwayarcshape.GongGetIdentifier(stage)
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topstackofgrowthcurve2d, topstackofgrowthcurve2d.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongGetReferenceIdentifier(stage *Stage) string {
	return topstackofgrowthcurve2d.GongGetIdentifier(stage)
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topstackofrotatedgrowthcurve2d, topstackofrotatedgrowthcurve2d.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongGetReferenceIdentifier(stage *Stage) string {
	return topstackofrotatedgrowthcurve2d.GongGetIdentifier(stage)
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topstackofrotatedgrowthcurve2dendarcshape, topstackofrotatedgrowthcurve2dendarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return topstackofrotatedgrowthcurve2dendarcshape.GongGetIdentifier(stage)
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topstackofrotatedgrowthcurve2dstartarcshape, topstackofrotatedgrowthcurve2dstartarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return topstackofrotatedgrowthcurve2dstartarcshape.GongGetIdentifier(stage)
}

func (topstartarcshape *TopStartArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topstartarcshape, topstartarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstartarcshape *TopStartArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return topstartarcshape.GongGetIdentifier(stage)
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topstartarcshapegrid, topstartarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstartarcshapegrid *TopStartArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return topstartarcshapegrid.GongGetIdentifier(stage)
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topstarthalfwayarcshape, topstarthalfwayarcshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongGetReferenceIdentifier(stage *Stage) string {
	return topstarthalfwayarcshape.GongGetIdentifier(stage)
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(topstarthalfwayarcshapegrid, topstarthalfwayarcshapegrid.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongGetReferenceIdentifier(stage *Stage) string {
	return topstarthalfwayarcshapegrid.GongGetIdentifier(stage)
}

func (torus3dshape *Torus3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(torus3dshape, torus3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (torus3dshape *Torus3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return torus3dshape.GongGetIdentifier(stage)
}

func (torusedge3dshape *TorusEdge3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(torusedge3dshape, torusedge3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (torusedge3dshape *TorusEdge3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return torusedge3dshape.GongGetIdentifier(stage)
}

func (torusstackshape *TorusStackShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(torusstackshape, torusstackshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (torusstackshape *TorusStackShape) GongGetReferenceIdentifier(stage *Stage) string {
	return torusstackshape.GongGetIdentifier(stage)
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tubevase3ddiagram, tubevase3ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tubevase3ddiagram *TubeVase3DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return tubevase3ddiagram.GongGetIdentifier(stage)
}

func (tubevaseabstract *TubeVaseAbstract) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tubevaseabstract, tubevaseabstract.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tubevaseabstract *TubeVaseAbstract) GongGetReferenceIdentifier(stage *Stage) string {
	return tubevaseabstract.GongGetIdentifier(stage)
}

func (vase2ddiagram *Vase2DDiagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(vase2ddiagram, vase2ddiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vase2ddiagram *Vase2DDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return vase2ddiagram.GongGetIdentifier(stage)
}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(vasetrapezeringshape, vasetrapezeringshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vasetrapezeringshape *VaseTrapezeRingShape) GongGetReferenceIdentifier(stage *Stage) string {
	return vasetrapezeringshape.GongGetIdentifier(stage)
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(verticaltorusstackshape, verticaltorusstackshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticaltorusstackshape *VerticalTorusStackShape) GongGetReferenceIdentifier(stage *Stage) string {
	return verticaltorusstackshape.GongGetIdentifier(stage)
}

func (volumekey3dshape *VolumeKey3DShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(volumekey3dshape, volumekey3dshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (volumekey3dshape *VolumeKey3DShape) GongGetReferenceIdentifier(stage *Stage) string {
	return volumekey3dshape.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (angle0shape *Angle0Shape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(angle0shape.GongGetIdentifier(stage), "Angle0Shape", angle0shape.Name)
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(arcnormalvectorshape.GongGetIdentifier(stage), "ArcNormalVectorShape", arcnormalvectorshape.Name)
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(arcnormalvectorshapegrid.GongGetIdentifier(stage), "ArcNormalVectorShapeGrid", arcnormalvectorshapegrid.Name)
}

func (axesshape *AxesShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(axesshape.GongGetIdentifier(stage), "AxesShape", axesshape.Name)
}

func (basevectorshape *BaseVectorShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(basevectorshape.GongGetIdentifier(stage), "BaseVectorShape", basevectorshape.Name)
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(basevectorshapegrid.GongGetIdentifier(stage), "BaseVectorShapeGrid", basevectorshapegrid.Name)
}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bottomcurveplane1shape.GongGetIdentifier(stage), "BottomCurvePlane1Shape", bottomcurveplane1shape.Name)
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bottomcurveplane2shape.GongGetIdentifier(stage), "BottomCurvePlane2Shape", bottomcurveplane2shape.Name)
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(chosenp1p2pairshape.GongGetIdentifier(stage), "ChosenP1P2PairShape", chosenp1p2pairshape.Name)
}

func (circlegridshape *CircleGridShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(circlegridshape.GongGetIdentifier(stage), "CircleGridShape", circlegridshape.Name)
}

func (circumference3dshape *Circumference3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(circumference3dshape.GongGetIdentifier(stage), "Circumference3DShape", circumference3dshape.Name)
}

func (clock2ddiagram *Clock2DDiagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(clock2ddiagram.GongGetIdentifier(stage), "Clock2DDiagram", clock2ddiagram.Name)
}

func (clock3ddiagram *Clock3DDiagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(clock3ddiagram.GongGetIdentifier(stage), "Clock3DDiagram", clock3ddiagram.Name)
}

func (clocktopcurveshape *ClockTopCurveShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(clocktopcurveshape.GongGetIdentifier(stage), "ClockTopCurveShape", clocktopcurveshape.Name)
}

func (cutline3dshape *CutLine3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(cutline3dshape.GongGetIdentifier(stage), "CutLine3DShape", cutline3dshape.Name)
}

func (endarcshape *EndArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(endarcshape.GongGetIdentifier(stage), "EndArcShape", endarcshape.Name)
}

func (endarcshapegrid *EndArcShapeGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(endarcshapegrid.GongGetIdentifier(stage), "EndArcShapeGrid", endarcshapegrid.Name)
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(endhalfwayarcshape.GongGetIdentifier(stage), "EndHalfwayArcShape", endhalfwayarcshape.Name)
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(endhalfwayarcshapegrid.GongGetIdentifier(stage), "EndHalfwayArcShapeGrid", endhalfwayarcshapegrid.Name)
}

func (explanationtextshape *ExplanationTextShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(explanationtextshape.GongGetIdentifier(stage), "ExplanationTextShape", explanationtextshape.Name)
}

func (eye3dshape *Eye3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(eye3dshape.GongGetIdentifier(stage), "Eye3DShape", eye3dshape.Name)
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(eyecornerssampledpoints3dshape.GongGetIdentifier(stage), "EyeCornersSampledPoints3DShape", eyecornerssampledpoints3dshape.Name)
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(eyesampledpoints3dshape.GongGetIdentifier(stage), "EyeSampledPoints3DShape", eyesampledpoints3dshape.Name)
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(eyeseatbottomcurveshape.GongGetIdentifier(stage), "EyeSeatBottomCurveShape", eyeseatbottomcurveshape.Name)
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(eyestoolbottomcurveshape.GongGetIdentifier(stage), "EyeStoolBottomCurveShape", eyestoolbottomcurveshape.Name)
}

func (eyevolume3dshape *EyeVolume3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(eyevolume3dshape.GongGetIdentifier(stage), "EyeVolume3DShape", eyevolume3dshape.Name)
}

func (gridpathshape *GridPathShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gridpathshape.GongGetIdentifier(stage), "GridPathShape", gridpathshape.Name)
}

func (growthcurve2d *GrowthCurve2D) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(growthcurve2d.GongGetIdentifier(stage), "GrowthCurve2D", growthcurve2d.Name)
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(growthcurve2dribbon.GongGetIdentifier(stage), "GrowthCurve2DRibbon", growthcurve2dribbon.Name)
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(growthcurve2dribbonendshape.GongGetIdentifier(stage), "GrowthCurve2DRibbonEndShape", growthcurve2dribbonendshape.Name)
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(growthcurve2dribbonstartshape.GongGetIdentifier(stage), "GrowthCurve2DRibbonStartShape", growthcurve2dribbonstartshape.Name)
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(growthcurverhombusgridshape.GongGetIdentifier(stage), "GrowthCurveRhombusGridShape", growthcurverhombusgridshape.Name)
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(growthcurverhombusshape.GongGetIdentifier(stage), "GrowthCurveRhombusShape", growthcurverhombusshape.Name)
}

func (growthvectorshape *GrowthVectorShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(growthvectorshape.GongGetIdentifier(stage), "GrowthVectorShape", growthvectorshape.Name)
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(initialrhombusgridshape.GongGetIdentifier(stage), "InitialRhombusGridShape", initialrhombusgridshape.Name)
}

func (initialrhombusshape *InitialRhombusShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(initialrhombusshape.GongGetIdentifier(stage), "InitialRhombusShape", initialrhombusshape.Name)
}

func (key3dshape *Key3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(key3dshape.GongGetIdentifier(stage), "Key3DShape", key3dshape.Name)
}

func (keyhole3dshape *KeyHole3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(keyhole3dshape.GongGetIdentifier(stage), "KeyHole3DShape", keyhole3dshape.Name)
}

func (keyholeshape *KeyHoleShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(keyholeshape.GongGetIdentifier(stage), "KeyHoleShape", keyholeshape.Name)
}

func (leaves3dshape *Leaves3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(leaves3dshape.GongGetIdentifier(stage), "Leaves3DShape", leaves3dshape.Name)
}

func (library *Library) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(library.GongGetIdentifier(stage), "Library", library.Name)
}

func (midarcvectorshape *MidArcVectorShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(midarcvectorshape.GongGetIdentifier(stage), "MidArcVectorShape", midarcvectorshape.Name)
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(midarcvectorshapegrid.GongGetIdentifier(stage), "MidArcVectorShapeGrid", midarcvectorshapegrid.Name)
}

func (originalpoints3dshape *OriginalPoints3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(originalpoints3dshape.GongGetIdentifier(stage), "OriginalPoints3DShape", originalpoints3dshape.Name)
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(parastichymcurves3dshape.GongGetIdentifier(stage), "ParastichyMCurves3DShape", parastichymcurves3dshape.Name)
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(parastichyncurves3dshape.GongGetIdentifier(stage), "ParastichyNCurves3DShape", parastichyncurves3dshape.Name)
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallygrowthcurve2dribbon.GongGetIdentifier(stage), "PartiallyGrowthCurve2DRibbon", partiallygrowthcurve2dribbon.Name)
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallygrowthcurve2dribbonendshape.GongGetIdentifier(stage), "PartiallyGrowthCurve2DRibbonEndShape", partiallygrowthcurve2dribbonendshape.Name)
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallygrowthcurve2dribbonstartshape.GongGetIdentifier(stage), "PartiallyGrowthCurve2DRibbonStartShape", partiallygrowthcurve2dribbonstartshape.Name)
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallygrowthcurve2dtrajectory.GongGetIdentifier(stage), "PartiallyGrowthCurve2DTrajectory", partiallygrowthcurve2dtrajectory.Name)
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallygrowthcurve2dtrajectoryp1curveshape.GongGetIdentifier(stage), "PartiallyGrowthCurve2DTrajectoryP1CurveShape", partiallygrowthcurve2dtrajectoryp1curveshape.Name)
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallygrowthcurve2dtrajectoryp1p2.GongGetIdentifier(stage), "PartiallyGrowthCurve2DTrajectoryP1P2", partiallygrowthcurve2dtrajectoryp1p2.Name)
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongGetIdentifier(stage), "PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape", partiallygrowthcurve2dtrajectoryp1p2pairlineshape.Name)
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallygrowthcurve2dtrajectoryp1pointshape.GongGetIdentifier(stage), "PartiallyGrowthCurve2DTrajectoryP1PointShape", partiallygrowthcurve2dtrajectoryp1pointshape.Name)
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallygrowthcurve2dtrajectoryp2curveshape.GongGetIdentifier(stage), "PartiallyGrowthCurve2DTrajectoryP2CurveShape", partiallygrowthcurve2dtrajectoryp2curveshape.Name)
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallygrowthcurve2dtrajectoryp2pointshape.GongGetIdentifier(stage), "PartiallyGrowthCurve2DTrajectoryP2PointShape", partiallygrowthcurve2dtrajectoryp2pointshape.Name)
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallygrowthcurve2dtrajectoryshape.GongGetIdentifier(stage), "PartiallyGrowthCurve2DTrajectoryShape", partiallygrowthcurve2dtrajectoryshape.Name)
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallyrotatedseatbottomcurveshape.GongGetIdentifier(stage), "PartiallyRotatedSeatBottomCurveShape", partiallyrotatedseatbottomcurveshape.Name)
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallyrotatedseattopcurveshape.GongGetIdentifier(stage), "PartiallyRotatedSeatTopCurveShape", partiallyrotatedseattopcurveshape.Name)
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partiallyrotatedtorusshape.GongGetIdentifier(stage), "PartiallyRotatedTorusShape", partiallyrotatedtorusshape.Name)
}

func (perpendicularvector *PerpendicularVector) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(perpendicularvector.GongGetIdentifier(stage), "PerpendicularVector", perpendicularvector.Name)
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(perpendicularvectorgrid.GongGetIdentifier(stage), "PerpendicularVectorGrid", perpendicularvectorgrid.Name)
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(perpendicularvectorgridhalfway.GongGetIdentifier(stage), "PerpendicularVectorGridHalfway", perpendicularvectorgridhalfway.Name)
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(perpendicularvectorhalfway.GongGetIdentifier(stage), "PerpendicularVectorHalfway", perpendicularvectorhalfway.Name)
}

func (plant2ddiagram *Plant2DDiagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(plant2ddiagram.GongGetIdentifier(stage), "Plant2DDiagram", plant2ddiagram.Name)
}

func (plant3ddiagram *Plant3DDiagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(plant3ddiagram.GongGetIdentifier(stage), "Plant3DDiagram", plant3ddiagram.Name)
}

func (plantabstract *PlantAbstract) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(plantabstract.GongGetIdentifier(stage), "PlantAbstract", plantabstract.Name)
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(plantcircumferenceshape.GongGetIdentifier(stage), "PlantCircumferenceShape", plantcircumferenceshape.Name)
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pointsandlines3dshape.GongGetIdentifier(stage), "PointsAndLines3DShape", pointsandlines3dshape.Name)
}

func (pxshape *PxShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pxshape.GongGetIdentifier(stage), "PxShape", pxshape.Name)
}

func (rendered3dshape *Rendered3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rendered3dshape.GongGetIdentifier(stage), "Rendered3DShape", rendered3dshape.Name)
}

func (rhombusshape *RhombusShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rhombusshape.GongGetIdentifier(stage), "RhombusShape", rhombusshape.Name)
}

func (rhombusstuff *RhombusStuff) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rhombusstuff.GongGetIdentifier(stage), "RhombusStuff", rhombusstuff.Name)
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rotatedrhombusgridshape.GongGetIdentifier(stage), "RotatedRhombusGridShape", rotatedrhombusgridshape.Name)
}

func (rotatedrhombusshape *RotatedRhombusShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rotatedrhombusshape.GongGetIdentifier(stage), "RotatedRhombusShape", rotatedrhombusshape.Name)
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rotatedsampledpoints3dshape.GongGetIdentifier(stage), "RotatedSampledPoints3DShape", rotatedsampledpoints3dshape.Name)
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rotatedseatandlegs3dshape.GongGetIdentifier(stage), "RotatedSeatAndLegs3DShape", rotatedseatandlegs3dshape.Name)
}

func (sampledpoints3dshape *SampledPoints3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(sampledpoints3dshape.GongGetIdentifier(stage), "SampledPoints3DShape", sampledpoints3dshape.Name)
}

func (seat3dshape *Seat3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(seat3dshape.GongGetIdentifier(stage), "Seat3DShape", seat3dshape.Name)
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(seatandlegs3dshape.GongGetIdentifier(stage), "SeatAndLegs3DShape", seatandlegs3dshape.Name)
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(seatbottomcurveshape.GongGetIdentifier(stage), "SeatBottomCurveShape", seatbottomcurveshape.Name)
}

func (seattopcurveshape *SeatTopCurveShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(seattopcurveshape.GongGetIdentifier(stage), "SeatTopCurveShape", seattopcurveshape.Name)
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedbottomtopstartarcshape.GongGetIdentifier(stage), "ShiftedBottomTopStartArcShape", shiftedbottomtopstartarcshape.Name)
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedbottomtopstartarcshapegrid.GongGetIdentifier(stage), "ShiftedBottomTopStartArcShapeGrid", shiftedbottomtopstartarcshapegrid.Name)
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedleftgrowthcurve2dribbon.GongGetIdentifier(stage), "ShiftedLeftGrowthCurve2DRibbon", shiftedleftgrowthcurve2dribbon.Name)
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedleftgrowthcurve2dribbonendshape.GongGetIdentifier(stage), "ShiftedLeftGrowthCurve2DRibbonEndShape", shiftedleftgrowthcurve2dribbonendshape.Name)
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedleftgrowthcurve2dribbonstartshape.GongGetIdentifier(stage), "ShiftedLeftGrowthCurve2DRibbonStartShape", shiftedleftgrowthcurve2dribbonstartshape.Name)
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedleftpartiallygrowthcurve2dribbon.GongGetIdentifier(stage), "ShiftedLeftPartiallyGrowthCurve2DRibbon", shiftedleftpartiallygrowthcurve2dribbon.Name)
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedleftpartiallygrowthcurve2dribbonendshape.GongGetIdentifier(stage), "ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape", shiftedleftpartiallygrowthcurve2dribbonendshape.Name)
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongGetIdentifier(stage), "ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape", shiftedleftpartiallygrowthcurve2dribbonstartshape.Name)
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedleftstackgrowthcurveendarcshape.GongGetIdentifier(stage), "ShiftedLeftStackGrowthCurveEndArcShape", shiftedleftstackgrowthcurveendarcshape.Name)
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedleftstackgrowthcurvestartarcshape.GongGetIdentifier(stage), "ShiftedLeftStackGrowthCurveStartArcShape", shiftedleftstackgrowthcurvestartarcshape.Name)
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedleftstacknormalvector.GongGetIdentifier(stage), "ShiftedLeftStackNormalVector", shiftedleftstacknormalvector.Name)
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedleftstackofgrowthcurve.GongGetIdentifier(stage), "ShiftedLeftStackOfGrowthCurve", shiftedleftstackofgrowthcurve.Name)
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedleftstackofnormalvector.GongGetIdentifier(stage), "ShiftedLeftStackOfNormalVector", shiftedleftstackofnormalvector.Name)
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedrightgrowthcurve2dribbon.GongGetIdentifier(stage), "ShiftedRightGrowthCurve2DRibbon", shiftedrightgrowthcurve2dribbon.Name)
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedrightgrowthcurve2dribbonendshape.GongGetIdentifier(stage), "ShiftedRightGrowthCurve2DRibbonEndShape", shiftedrightgrowthcurve2dribbonendshape.Name)
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shiftedrightgrowthcurve2dribbonstartshape.GongGetIdentifier(stage), "ShiftedRightGrowthCurve2DRibbonStartShape", shiftedrightgrowthcurve2dribbonstartshape.Name)
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackgrowthcurve2dendhalfwayarcshape.GongGetIdentifier(stage), "StackGrowthCurve2DEndHalfwayArcShape", stackgrowthcurve2dendhalfwayarcshape.Name)
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackgrowthcurve2dribbonendshape.GongGetIdentifier(stage), "StackGrowthCurve2DRibbonEndShape", stackgrowthcurve2dribbonendshape.Name)
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackgrowthcurve2dribbonstartshape.GongGetIdentifier(stage), "StackGrowthCurve2DRibbonStartShape", stackgrowthcurve2dribbonstartshape.Name)
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackgrowthcurve2dstarthalfwayarcshape.GongGetIdentifier(stage), "StackGrowthCurve2DStartHalfwayArcShape", stackgrowthcurve2dstarthalfwayarcshape.Name)
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackofgrowthcurve2d.GongGetIdentifier(stage), "StackOfGrowthCurve2D", stackofgrowthcurve2d.Name)
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackofgrowthcurve2dbygrowthvector.GongGetIdentifier(stage), "StackOfGrowthCurve2DByGrowthVector", stackofgrowthcurve2dbygrowthvector.Name)
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackofgrowthcurve2dribbon.GongGetIdentifier(stage), "StackOfGrowthCurve2DRibbon", stackofgrowthcurve2dribbon.Name)
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackofpartiallyrotatedtorusshape.GongGetIdentifier(stage), "StackOfPartiallyRotatedTorusShape", stackofpartiallyrotatedtorusshape.Name)
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackofrotatedgrowthcurve2d.GongGetIdentifier(stage), "StackOfRotatedGrowthCurve2D", stackofrotatedgrowthcurve2d.Name)
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackofrotatedgrowthcurve2dribbon.GongGetIdentifier(stage), "StackOfRotatedGrowthCurve2DRibbon", stackofrotatedgrowthcurve2dribbon.Name)
}

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackofrotatedvasetrapezeringsshape.GongGetIdentifier(stage), "StackOfRotatedVaseTrapezeRingsShape", stackofrotatedvasetrapezeringsshape.Name)
}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackofvasetrapezeringsshape.GongGetIdentifier(stage), "StackOfVaseTrapezeRingsShape", stackofvasetrapezeringsshape.Name)
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackrotatedgrowthcurve2dendarcshape.GongGetIdentifier(stage), "StackRotatedGrowthCurve2DEndArcShape", stackrotatedgrowthcurve2dendarcshape.Name)
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackrotatedgrowthcurve2dribbonendshape.GongGetIdentifier(stage), "StackRotatedGrowthCurve2DRibbonEndShape", stackrotatedgrowthcurve2dribbonendshape.Name)
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackrotatedgrowthcurve2dribbonstartshape.GongGetIdentifier(stage), "StackRotatedGrowthCurve2DRibbonStartShape", stackrotatedgrowthcurve2dribbonstartshape.Name)
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stackrotatedgrowthcurve2dstartarcshape.GongGetIdentifier(stage), "StackRotatedGrowthCurve2DStartArcShape", stackrotatedgrowthcurve2dstartarcshape.Name)
}

func (startarcshape *StartArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(startarcshape.GongGetIdentifier(stage), "StartArcShape", startarcshape.Name)
}

func (startarcshapegrid *StartArcShapeGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(startarcshapegrid.GongGetIdentifier(stage), "StartArcShapeGrid", startarcshapegrid.Name)
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(starthalfwayarcshape.GongGetIdentifier(stage), "StartHalfwayArcShape", starthalfwayarcshape.Name)
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(starthalfwayarcshapegrid.GongGetIdentifier(stage), "StartHalfwayArcShapeGrid", starthalfwayarcshapegrid.Name)
}

func (stemcylinder3dshape *StemCylinder3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stemcylinder3dshape.GongGetIdentifier(stage), "StemCylinder3DShape", stemcylinder3dshape.Name)
}

func (stool2ddiagram *Stool2DDiagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stool2ddiagram.GongGetIdentifier(stage), "Stool2DDiagram", stool2ddiagram.Name)
}

func (stool3ddiagram *Stool3DDiagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stool3ddiagram.GongGetIdentifier(stage), "Stool3DDiagram", stool3ddiagram.Name)
}

func (tiledfloor3dshape *TiledFloor3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tiledfloor3dshape.GongGetIdentifier(stage), "TiledFloor3DShape", tiledfloor3dshape.Name)
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topcurveplane1shape.GongGetIdentifier(stage), "TopCurvePlane1Shape", topcurveplane1shape.Name)
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topcurveplane2shape.GongGetIdentifier(stage), "TopCurvePlane2Shape", topcurveplane2shape.Name)
}

func (topendarcshape *TopEndArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topendarcshape.GongGetIdentifier(stage), "TopEndArcShape", topendarcshape.Name)
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topendarcshapegrid.GongGetIdentifier(stage), "TopEndArcShapeGrid", topendarcshapegrid.Name)
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topendhalfwayarcshape.GongGetIdentifier(stage), "TopEndHalfwayArcShape", topendhalfwayarcshape.Name)
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topendhalfwayarcshapegrid.GongGetIdentifier(stage), "TopEndHalfwayArcShapeGrid", topendhalfwayarcshapegrid.Name)
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topgrowthcurve2d.GongGetIdentifier(stage), "TopGrowthCurve2D", topgrowthcurve2d.Name)
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topmidarcvectorshape.GongGetIdentifier(stage), "TopMidArcVectorShape", topmidarcvectorshape.Name)
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topmidarcvectorshapegrid.GongGetIdentifier(stage), "TopMidArcVectorShapeGrid", topmidarcvectorshapegrid.Name)
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topstackgrowthcurve2dendhalfwayarcshape.GongGetIdentifier(stage), "TopStackGrowthCurve2DEndHalfwayArcShape", topstackgrowthcurve2dendhalfwayarcshape.Name)
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topstackgrowthcurve2dstarthalfwayarcshape.GongGetIdentifier(stage), "TopStackGrowthCurve2DStartHalfwayArcShape", topstackgrowthcurve2dstarthalfwayarcshape.Name)
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topstackofgrowthcurve2d.GongGetIdentifier(stage), "TopStackOfGrowthCurve2D", topstackofgrowthcurve2d.Name)
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topstackofrotatedgrowthcurve2d.GongGetIdentifier(stage), "TopStackOfRotatedGrowthCurve2D", topstackofrotatedgrowthcurve2d.Name)
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topstackofrotatedgrowthcurve2dendarcshape.GongGetIdentifier(stage), "TopStackOfRotatedGrowthCurve2DEndArcShape", topstackofrotatedgrowthcurve2dendarcshape.Name)
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topstackofrotatedgrowthcurve2dstartarcshape.GongGetIdentifier(stage), "TopStackOfRotatedGrowthCurve2DStartArcShape", topstackofrotatedgrowthcurve2dstartarcshape.Name)
}

func (topstartarcshape *TopStartArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topstartarcshape.GongGetIdentifier(stage), "TopStartArcShape", topstartarcshape.Name)
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topstartarcshapegrid.GongGetIdentifier(stage), "TopStartArcShapeGrid", topstartarcshapegrid.Name)
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topstarthalfwayarcshape.GongGetIdentifier(stage), "TopStartHalfwayArcShape", topstarthalfwayarcshape.Name)
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(topstarthalfwayarcshapegrid.GongGetIdentifier(stage), "TopStartHalfwayArcShapeGrid", topstarthalfwayarcshapegrid.Name)
}

func (torus3dshape *Torus3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(torus3dshape.GongGetIdentifier(stage), "Torus3DShape", torus3dshape.Name)
}

func (torusedge3dshape *TorusEdge3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(torusedge3dshape.GongGetIdentifier(stage), "TorusEdge3DShape", torusedge3dshape.Name)
}

func (torusstackshape *TorusStackShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(torusstackshape.GongGetIdentifier(stage), "TorusStackShape", torusstackshape.Name)
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tubevase3ddiagram.GongGetIdentifier(stage), "TubeVase3DDiagram", tubevase3ddiagram.Name)
}

func (tubevaseabstract *TubeVaseAbstract) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tubevaseabstract.GongGetIdentifier(stage), "TubeVaseAbstract", tubevaseabstract.Name)
}

func (vase2ddiagram *Vase2DDiagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(vase2ddiagram.GongGetIdentifier(stage), "Vase2DDiagram", vase2ddiagram.Name)
}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(vasetrapezeringshape.GongGetIdentifier(stage), "VaseTrapezeRingShape", vasetrapezeringshape.Name)
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(verticaltorusstackshape.GongGetIdentifier(stage), "VerticalTorusStackShape", verticaltorusstackshape.Name)
}

func (volumekey3dshape *VolumeKey3DShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(volumekey3dshape.GongGetIdentifier(stage), "VolumeKey3DShape", volumekey3dshape.Name)
}

// insertion point for unstaging
func (angle0shape *Angle0Shape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(angle0shape.GongGetReferenceIdentifier(stage))
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(arcnormalvectorshape.GongGetReferenceIdentifier(stage))
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(arcnormalvectorshapegrid.GongGetReferenceIdentifier(stage))
}

func (axesshape *AxesShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(axesshape.GongGetReferenceIdentifier(stage))
}

func (basevectorshape *BaseVectorShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(basevectorshape.GongGetReferenceIdentifier(stage))
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(basevectorshapegrid.GongGetReferenceIdentifier(stage))
}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bottomcurveplane1shape.GongGetReferenceIdentifier(stage))
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bottomcurveplane2shape.GongGetReferenceIdentifier(stage))
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(chosenp1p2pairshape.GongGetReferenceIdentifier(stage))
}

func (circlegridshape *CircleGridShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(circlegridshape.GongGetReferenceIdentifier(stage))
}

func (circumference3dshape *Circumference3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(circumference3dshape.GongGetReferenceIdentifier(stage))
}

func (clock2ddiagram *Clock2DDiagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(clock2ddiagram.GongGetReferenceIdentifier(stage))
}

func (clock3ddiagram *Clock3DDiagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(clock3ddiagram.GongGetReferenceIdentifier(stage))
}

func (clocktopcurveshape *ClockTopCurveShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(clocktopcurveshape.GongGetReferenceIdentifier(stage))
}

func (cutline3dshape *CutLine3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(cutline3dshape.GongGetReferenceIdentifier(stage))
}

func (endarcshape *EndArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(endarcshape.GongGetReferenceIdentifier(stage))
}

func (endarcshapegrid *EndArcShapeGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(endarcshapegrid.GongGetReferenceIdentifier(stage))
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(endhalfwayarcshape.GongGetReferenceIdentifier(stage))
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(endhalfwayarcshapegrid.GongGetReferenceIdentifier(stage))
}

func (explanationtextshape *ExplanationTextShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(explanationtextshape.GongGetReferenceIdentifier(stage))
}

func (eye3dshape *Eye3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(eye3dshape.GongGetReferenceIdentifier(stage))
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(eyecornerssampledpoints3dshape.GongGetReferenceIdentifier(stage))
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(eyesampledpoints3dshape.GongGetReferenceIdentifier(stage))
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(eyeseatbottomcurveshape.GongGetReferenceIdentifier(stage))
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(eyestoolbottomcurveshape.GongGetReferenceIdentifier(stage))
}

func (eyevolume3dshape *EyeVolume3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(eyevolume3dshape.GongGetReferenceIdentifier(stage))
}

func (gridpathshape *GridPathShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gridpathshape.GongGetReferenceIdentifier(stage))
}

func (growthcurve2d *GrowthCurve2D) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(growthcurve2d.GongGetReferenceIdentifier(stage))
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(growthcurve2dribbon.GongGetReferenceIdentifier(stage))
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(growthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(growthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(growthcurverhombusgridshape.GongGetReferenceIdentifier(stage))
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(growthcurverhombusshape.GongGetReferenceIdentifier(stage))
}

func (growthvectorshape *GrowthVectorShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(growthvectorshape.GongGetReferenceIdentifier(stage))
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(initialrhombusgridshape.GongGetReferenceIdentifier(stage))
}

func (initialrhombusshape *InitialRhombusShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(initialrhombusshape.GongGetReferenceIdentifier(stage))
}

func (key3dshape *Key3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(key3dshape.GongGetReferenceIdentifier(stage))
}

func (keyhole3dshape *KeyHole3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(keyhole3dshape.GongGetReferenceIdentifier(stage))
}

func (keyholeshape *KeyHoleShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(keyholeshape.GongGetReferenceIdentifier(stage))
}

func (leaves3dshape *Leaves3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(leaves3dshape.GongGetReferenceIdentifier(stage))
}

func (library *Library) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(library.GongGetReferenceIdentifier(stage))
}

func (midarcvectorshape *MidArcVectorShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(midarcvectorshape.GongGetReferenceIdentifier(stage))
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(midarcvectorshapegrid.GongGetReferenceIdentifier(stage))
}

func (originalpoints3dshape *OriginalPoints3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(originalpoints3dshape.GongGetReferenceIdentifier(stage))
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(parastichymcurves3dshape.GongGetReferenceIdentifier(stage))
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(parastichyncurves3dshape.GongGetReferenceIdentifier(stage))
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallygrowthcurve2dribbon.GongGetReferenceIdentifier(stage))
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallygrowthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallygrowthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallygrowthcurve2dtrajectory.GongGetReferenceIdentifier(stage))
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallygrowthcurve2dtrajectoryp1curveshape.GongGetReferenceIdentifier(stage))
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallygrowthcurve2dtrajectoryp1p2.GongGetReferenceIdentifier(stage))
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallygrowthcurve2dtrajectoryp1p2pairlineshape.GongGetReferenceIdentifier(stage))
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallygrowthcurve2dtrajectoryp1pointshape.GongGetReferenceIdentifier(stage))
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallygrowthcurve2dtrajectoryp2curveshape.GongGetReferenceIdentifier(stage))
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallygrowthcurve2dtrajectoryp2pointshape.GongGetReferenceIdentifier(stage))
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallygrowthcurve2dtrajectoryshape.GongGetReferenceIdentifier(stage))
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallyrotatedseatbottomcurveshape.GongGetReferenceIdentifier(stage))
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallyrotatedseattopcurveshape.GongGetReferenceIdentifier(stage))
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partiallyrotatedtorusshape.GongGetReferenceIdentifier(stage))
}

func (perpendicularvector *PerpendicularVector) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(perpendicularvector.GongGetReferenceIdentifier(stage))
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(perpendicularvectorgrid.GongGetReferenceIdentifier(stage))
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(perpendicularvectorgridhalfway.GongGetReferenceIdentifier(stage))
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(perpendicularvectorhalfway.GongGetReferenceIdentifier(stage))
}

func (plant2ddiagram *Plant2DDiagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(plant2ddiagram.GongGetReferenceIdentifier(stage))
}

func (plant3ddiagram *Plant3DDiagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(plant3ddiagram.GongGetReferenceIdentifier(stage))
}

func (plantabstract *PlantAbstract) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(plantabstract.GongGetReferenceIdentifier(stage))
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(plantcircumferenceshape.GongGetReferenceIdentifier(stage))
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pointsandlines3dshape.GongGetReferenceIdentifier(stage))
}

func (pxshape *PxShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pxshape.GongGetReferenceIdentifier(stage))
}

func (rendered3dshape *Rendered3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rendered3dshape.GongGetReferenceIdentifier(stage))
}

func (rhombusshape *RhombusShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rhombusshape.GongGetReferenceIdentifier(stage))
}

func (rhombusstuff *RhombusStuff) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rhombusstuff.GongGetReferenceIdentifier(stage))
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rotatedrhombusgridshape.GongGetReferenceIdentifier(stage))
}

func (rotatedrhombusshape *RotatedRhombusShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rotatedrhombusshape.GongGetReferenceIdentifier(stage))
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rotatedsampledpoints3dshape.GongGetReferenceIdentifier(stage))
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rotatedseatandlegs3dshape.GongGetReferenceIdentifier(stage))
}

func (sampledpoints3dshape *SampledPoints3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(sampledpoints3dshape.GongGetReferenceIdentifier(stage))
}

func (seat3dshape *Seat3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(seat3dshape.GongGetReferenceIdentifier(stage))
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(seatandlegs3dshape.GongGetReferenceIdentifier(stage))
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(seatbottomcurveshape.GongGetReferenceIdentifier(stage))
}

func (seattopcurveshape *SeatTopCurveShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(seattopcurveshape.GongGetReferenceIdentifier(stage))
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedbottomtopstartarcshape.GongGetReferenceIdentifier(stage))
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedbottomtopstartarcshapegrid.GongGetReferenceIdentifier(stage))
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedleftgrowthcurve2dribbon.GongGetReferenceIdentifier(stage))
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedleftgrowthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedleftgrowthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedleftpartiallygrowthcurve2dribbon.GongGetReferenceIdentifier(stage))
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedleftpartiallygrowthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedleftpartiallygrowthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedleftstackgrowthcurveendarcshape.GongGetReferenceIdentifier(stage))
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedleftstackgrowthcurvestartarcshape.GongGetReferenceIdentifier(stage))
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedleftstacknormalvector.GongGetReferenceIdentifier(stage))
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedleftstackofgrowthcurve.GongGetReferenceIdentifier(stage))
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedleftstackofnormalvector.GongGetReferenceIdentifier(stage))
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedrightgrowthcurve2dribbon.GongGetReferenceIdentifier(stage))
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedrightgrowthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shiftedrightgrowthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackgrowthcurve2dendhalfwayarcshape.GongGetReferenceIdentifier(stage))
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackgrowthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackgrowthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackgrowthcurve2dstarthalfwayarcshape.GongGetReferenceIdentifier(stage))
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackofgrowthcurve2d.GongGetReferenceIdentifier(stage))
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackofgrowthcurve2dbygrowthvector.GongGetReferenceIdentifier(stage))
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackofgrowthcurve2dribbon.GongGetReferenceIdentifier(stage))
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackofpartiallyrotatedtorusshape.GongGetReferenceIdentifier(stage))
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackofrotatedgrowthcurve2d.GongGetReferenceIdentifier(stage))
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackofrotatedgrowthcurve2dribbon.GongGetReferenceIdentifier(stage))
}

func (stackofrotatedvasetrapezeringsshape *StackOfRotatedVaseTrapezeRingsShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackofrotatedvasetrapezeringsshape.GongGetReferenceIdentifier(stage))
}

func (stackofvasetrapezeringsshape *StackOfVaseTrapezeRingsShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackofvasetrapezeringsshape.GongGetReferenceIdentifier(stage))
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackrotatedgrowthcurve2dendarcshape.GongGetReferenceIdentifier(stage))
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackrotatedgrowthcurve2dribbonendshape.GongGetReferenceIdentifier(stage))
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackrotatedgrowthcurve2dribbonstartshape.GongGetReferenceIdentifier(stage))
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stackrotatedgrowthcurve2dstartarcshape.GongGetReferenceIdentifier(stage))
}

func (startarcshape *StartArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(startarcshape.GongGetReferenceIdentifier(stage))
}

func (startarcshapegrid *StartArcShapeGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(startarcshapegrid.GongGetReferenceIdentifier(stage))
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(starthalfwayarcshape.GongGetReferenceIdentifier(stage))
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(starthalfwayarcshapegrid.GongGetReferenceIdentifier(stage))
}

func (stemcylinder3dshape *StemCylinder3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stemcylinder3dshape.GongGetReferenceIdentifier(stage))
}

func (stool2ddiagram *Stool2DDiagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stool2ddiagram.GongGetReferenceIdentifier(stage))
}

func (stool3ddiagram *Stool3DDiagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stool3ddiagram.GongGetReferenceIdentifier(stage))
}

func (tiledfloor3dshape *TiledFloor3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tiledfloor3dshape.GongGetReferenceIdentifier(stage))
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topcurveplane1shape.GongGetReferenceIdentifier(stage))
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topcurveplane2shape.GongGetReferenceIdentifier(stage))
}

func (topendarcshape *TopEndArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topendarcshape.GongGetReferenceIdentifier(stage))
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topendarcshapegrid.GongGetReferenceIdentifier(stage))
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topendhalfwayarcshape.GongGetReferenceIdentifier(stage))
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topendhalfwayarcshapegrid.GongGetReferenceIdentifier(stage))
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topgrowthcurve2d.GongGetReferenceIdentifier(stage))
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topmidarcvectorshape.GongGetReferenceIdentifier(stage))
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topmidarcvectorshapegrid.GongGetReferenceIdentifier(stage))
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topstackgrowthcurve2dendhalfwayarcshape.GongGetReferenceIdentifier(stage))
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topstackgrowthcurve2dstarthalfwayarcshape.GongGetReferenceIdentifier(stage))
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topstackofgrowthcurve2d.GongGetReferenceIdentifier(stage))
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topstackofrotatedgrowthcurve2d.GongGetReferenceIdentifier(stage))
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topstackofrotatedgrowthcurve2dendarcshape.GongGetReferenceIdentifier(stage))
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topstackofrotatedgrowthcurve2dstartarcshape.GongGetReferenceIdentifier(stage))
}

func (topstartarcshape *TopStartArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topstartarcshape.GongGetReferenceIdentifier(stage))
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topstartarcshapegrid.GongGetReferenceIdentifier(stage))
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topstarthalfwayarcshape.GongGetReferenceIdentifier(stage))
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(topstarthalfwayarcshapegrid.GongGetReferenceIdentifier(stage))
}

func (torus3dshape *Torus3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(torus3dshape.GongGetReferenceIdentifier(stage))
}

func (torusedge3dshape *TorusEdge3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(torusedge3dshape.GongGetReferenceIdentifier(stage))
}

func (torusstackshape *TorusStackShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(torusstackshape.GongGetReferenceIdentifier(stage))
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tubevase3ddiagram.GongGetReferenceIdentifier(stage))
}

func (tubevaseabstract *TubeVaseAbstract) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tubevaseabstract.GongGetReferenceIdentifier(stage))
}

func (vase2ddiagram *Vase2DDiagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(vase2ddiagram.GongGetReferenceIdentifier(stage))
}

func (vasetrapezeringshape *VaseTrapezeRingShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(vasetrapezeringshape.GongGetReferenceIdentifier(stage))
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(verticaltorusstackshape.GongGetReferenceIdentifier(stage))
}

func (volumekey3dshape *VolumeKey3DShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(volumekey3dshape.GongGetReferenceIdentifier(stage))
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

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
