// generated code - do not edit
package models

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/xuri/excelize/v2"
)

func (stage *Stage) SerializeStage(filename string) {
	stage.SerializeStage2(filename, false)
}

func (stage *Stage) SerializeStage2(filename string, addIDs bool) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println("cannot write xl file : ", err)
	}
}

func (stage *Stage) __gong__buildExcelizeFile(addIDs bool) *excelize.File {
	f := excelize.NewFile()
	{
		// insertion point
		{
			var instances []GongstructIF
			for instance := range stage.Angle0Shapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Angle0Shape", instances, (*Angle0Shape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArcNormalVectorShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArcNormalVectorShape", instances, (*ArcNormalVectorShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArcNormalVectorShapeGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArcNormalVectorShapeGrid", instances, (*ArcNormalVectorShapeGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AxesShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AxesShape", instances, (*AxesShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.BaseVectorShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "BaseVectorShape", instances, (*BaseVectorShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.BaseVectorShapeGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "BaseVectorShapeGrid", instances, (*BaseVectorShapeGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ChosenP1P2PairShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ChosenP1P2PairShape", instances, (*ChosenP1P2PairShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CircleGridShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CircleGridShape", instances, (*CircleGridShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Circumference3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Circumference3DShape", instances, (*Circumference3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Clock2DDiagrams {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Clock2DDiagram", instances, (*Clock2DDiagram)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Clock3DDiagrams {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Clock3DDiagram", instances, (*Clock3DDiagram)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ClockAbstracts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ClockAbstract", instances, (*ClockAbstract)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ClockTopCurveShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ClockTopCurveShape", instances, (*ClockTopCurveShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CutLine3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CutLine3DShape", instances, (*CutLine3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EndArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EndArcShape", instances, (*EndArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EndArcShapeGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EndArcShapeGrid", instances, (*EndArcShapeGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EndHalfwayArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EndHalfwayArcShape", instances, (*EndHalfwayArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EndHalfwayArcShapeGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EndHalfwayArcShapeGrid", instances, (*EndHalfwayArcShapeGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ExplanationTextShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ExplanationTextShape", instances, (*ExplanationTextShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Eye3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Eye3DShape", instances, (*Eye3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EyeCornersSampledPoints3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EyeCornersSampledPoints3DShape", instances, (*EyeCornersSampledPoints3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EyeSampledPoints3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EyeSampledPoints3DShape", instances, (*EyeSampledPoints3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EyeSeatBottomCurveShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EyeSeatBottomCurveShape", instances, (*EyeSeatBottomCurveShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EyeStoolBottomCurveShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EyeStoolBottomCurveShape", instances, (*EyeStoolBottomCurveShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EyeVolume3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EyeVolume3DShape", instances, (*EyeVolume3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GridPathShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GridPathShape", instances, (*GridPathShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GrowthCurve2Ds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GrowthCurve2D", instances, (*GrowthCurve2D)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GrowthCurve2DRibbons {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GrowthCurve2DRibbon", instances, (*GrowthCurve2DRibbon)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GrowthCurve2DRibbonEndShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GrowthCurve2DRibbonEndShape", instances, (*GrowthCurve2DRibbonEndShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GrowthCurve2DRibbonStartShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GrowthCurve2DRibbonStartShape", instances, (*GrowthCurve2DRibbonStartShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GrowthCurveRhombusGridShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GrowthCurveRhombusGridShape", instances, (*GrowthCurveRhombusGridShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GrowthCurveRhombusShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GrowthCurveRhombusShape", instances, (*GrowthCurveRhombusShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GrowthVectorShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GrowthVectorShape", instances, (*GrowthVectorShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InitialRhombusGridShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InitialRhombusGridShape", instances, (*InitialRhombusGridShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InitialRhombusShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InitialRhombusShape", instances, (*InitialRhombusShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Key3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Key3DShape", instances, (*Key3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.KeyHole3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "KeyHole3DShape", instances, (*KeyHole3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.KeyHoleShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "KeyHoleShape", instances, (*KeyHoleShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Leaves3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Leaves3DShape", instances, (*Leaves3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Librarys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Library", instances, (*Library)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MidArcVectorShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MidArcVectorShape", instances, (*MidArcVectorShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MidArcVectorShapeGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MidArcVectorShapeGrid", instances, (*MidArcVectorShapeGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MusicAbstracts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MusicAbstract", instances, (*MusicAbstract)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OriginalPoints3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OriginalPoints3DShape", instances, (*OriginalPoints3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ParastichyMCurves3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ParastichyMCurves3DShape", instances, (*ParastichyMCurves3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ParastichyNCurves3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ParastichyNCurves3DShape", instances, (*ParastichyNCurves3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyGrowthCurve2DRibbons {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyGrowthCurve2DRibbon", instances, (*PartiallyGrowthCurve2DRibbon)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyGrowthCurve2DRibbonEndShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyGrowthCurve2DRibbonEndShape", instances, (*PartiallyGrowthCurve2DRibbonEndShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyGrowthCurve2DRibbonStartShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyGrowthCurve2DRibbonStartShape", instances, (*PartiallyGrowthCurve2DRibbonStartShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyGrowthCurve2DTrajectorys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyGrowthCurve2DTrajectory", instances, (*PartiallyGrowthCurve2DTrajectory)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1CurveShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyGrowthCurve2DTrajectoryP1CurveShape", instances, (*PartiallyGrowthCurve2DTrajectoryP1CurveShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1P2s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyGrowthCurve2DTrajectoryP1P2", instances, (*PartiallyGrowthCurve2DTrajectoryP1P2)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape", instances, (*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyGrowthCurve2DTrajectoryP1PointShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyGrowthCurve2DTrajectoryP1PointShape", instances, (*PartiallyGrowthCurve2DTrajectoryP1PointShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyGrowthCurve2DTrajectoryP2CurveShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyGrowthCurve2DTrajectoryP2CurveShape", instances, (*PartiallyGrowthCurve2DTrajectoryP2CurveShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyGrowthCurve2DTrajectoryP2PointShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyGrowthCurve2DTrajectoryP2PointShape", instances, (*PartiallyGrowthCurve2DTrajectoryP2PointShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyGrowthCurve2DTrajectoryShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyGrowthCurve2DTrajectoryShape", instances, (*PartiallyGrowthCurve2DTrajectoryShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyRotatedSeatBottomCurveShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyRotatedSeatBottomCurveShape", instances, (*PartiallyRotatedSeatBottomCurveShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyRotatedSeatTopCurveShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyRotatedSeatTopCurveShape", instances, (*PartiallyRotatedSeatTopCurveShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PartiallyRotatedTorusShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PartiallyRotatedTorusShape", instances, (*PartiallyRotatedTorusShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PerpendicularVectors {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PerpendicularVector", instances, (*PerpendicularVector)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PerpendicularVectorGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PerpendicularVectorGrid", instances, (*PerpendicularVectorGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PerpendicularVectorGridHalfways {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PerpendicularVectorGridHalfway", instances, (*PerpendicularVectorGridHalfway)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PerpendicularVectorHalfways {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PerpendicularVectorHalfway", instances, (*PerpendicularVectorHalfway)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Plant2DDiagrams {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Plant2DDiagram", instances, (*Plant2DDiagram)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Plant3DDiagrams {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Plant3DDiagram", instances, (*Plant3DDiagram)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PlantAbstracts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PlantAbstract", instances, (*PlantAbstract)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PlantCircumferenceShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PlantCircumferenceShape", instances, (*PlantCircumferenceShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PointsAndLines3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PointsAndLines3DShape", instances, (*PointsAndLines3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PxShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PxShape", instances, (*PxShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Rendered3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Rendered3DShape", instances, (*Rendered3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RhombusShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RhombusShape", instances, (*RhombusShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RhombusStuffs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RhombusStuff", instances, (*RhombusStuff)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RotatedRhombusGridShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RotatedRhombusGridShape", instances, (*RotatedRhombusGridShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RotatedRhombusShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RotatedRhombusShape", instances, (*RotatedRhombusShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RotatedSampledPoints3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RotatedSampledPoints3DShape", instances, (*RotatedSampledPoints3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RotatedSeatAndLegs3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RotatedSeatAndLegs3DShape", instances, (*RotatedSeatAndLegs3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SampledPoints3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SampledPoints3DShape", instances, (*SampledPoints3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Seat3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Seat3DShape", instances, (*Seat3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeatAndLegs3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeatAndLegs3DShape", instances, (*SeatAndLegs3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeatBottomCurveShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeatBottomCurveShape", instances, (*SeatBottomCurveShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeatTopCurveShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeatTopCurveShape", instances, (*SeatTopCurveShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedBottomTopStartArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedBottomTopStartArcShape", instances, (*ShiftedBottomTopStartArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedBottomTopStartArcShapeGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedBottomTopStartArcShapeGrid", instances, (*ShiftedBottomTopStartArcShapeGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedLeftGrowthCurve2DRibbons {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedLeftGrowthCurve2DRibbon", instances, (*ShiftedLeftGrowthCurve2DRibbon)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedLeftGrowthCurve2DRibbonEndShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedLeftGrowthCurve2DRibbonEndShape", instances, (*ShiftedLeftGrowthCurve2DRibbonEndShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedLeftGrowthCurve2DRibbonStartShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedLeftGrowthCurve2DRibbonStartShape", instances, (*ShiftedLeftGrowthCurve2DRibbonStartShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedLeftPartiallyGrowthCurve2DRibbons {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedLeftPartiallyGrowthCurve2DRibbon", instances, (*ShiftedLeftPartiallyGrowthCurve2DRibbon)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape", instances, (*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape", instances, (*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedLeftStackGrowthCurveEndArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedLeftStackGrowthCurveEndArcShape", instances, (*ShiftedLeftStackGrowthCurveEndArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedLeftStackGrowthCurveStartArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedLeftStackGrowthCurveStartArcShape", instances, (*ShiftedLeftStackGrowthCurveStartArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedLeftStackNormalVectors {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedLeftStackNormalVector", instances, (*ShiftedLeftStackNormalVector)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedLeftStackOfGrowthCurves {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedLeftStackOfGrowthCurve", instances, (*ShiftedLeftStackOfGrowthCurve)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedLeftStackOfNormalVectors {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedLeftStackOfNormalVector", instances, (*ShiftedLeftStackOfNormalVector)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedRightGrowthCurve2DRibbons {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedRightGrowthCurve2DRibbon", instances, (*ShiftedRightGrowthCurve2DRibbon)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedRightGrowthCurve2DRibbonEndShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedRightGrowthCurve2DRibbonEndShape", instances, (*ShiftedRightGrowthCurve2DRibbonEndShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShiftedRightGrowthCurve2DRibbonStartShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShiftedRightGrowthCurve2DRibbonStartShape", instances, (*ShiftedRightGrowthCurve2DRibbonStartShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackGrowthCurve2DEndHalfwayArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackGrowthCurve2DEndHalfwayArcShape", instances, (*StackGrowthCurve2DEndHalfwayArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackGrowthCurve2DRibbonEndShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackGrowthCurve2DRibbonEndShape", instances, (*StackGrowthCurve2DRibbonEndShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackGrowthCurve2DRibbonStartShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackGrowthCurve2DRibbonStartShape", instances, (*StackGrowthCurve2DRibbonStartShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackGrowthCurve2DStartHalfwayArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackGrowthCurve2DStartHalfwayArcShape", instances, (*StackGrowthCurve2DStartHalfwayArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackOfGrowthCurve2Ds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackOfGrowthCurve2D", instances, (*StackOfGrowthCurve2D)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackOfGrowthCurve2DByGrowthVectors {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackOfGrowthCurve2DByGrowthVector", instances, (*StackOfGrowthCurve2DByGrowthVector)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackOfGrowthCurve2DRibbons {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackOfGrowthCurve2DRibbon", instances, (*StackOfGrowthCurve2DRibbon)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackOfPartiallyRotatedTorusShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackOfPartiallyRotatedTorusShape", instances, (*StackOfPartiallyRotatedTorusShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackOfRotatedGrowthCurve2Ds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackOfRotatedGrowthCurve2D", instances, (*StackOfRotatedGrowthCurve2D)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackOfRotatedGrowthCurve2DRibbons {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackOfRotatedGrowthCurve2DRibbon", instances, (*StackOfRotatedGrowthCurve2DRibbon)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackRotatedGrowthCurve2DEndArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackRotatedGrowthCurve2DEndArcShape", instances, (*StackRotatedGrowthCurve2DEndArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackRotatedGrowthCurve2DRibbonEndShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackRotatedGrowthCurve2DRibbonEndShape", instances, (*StackRotatedGrowthCurve2DRibbonEndShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackRotatedGrowthCurve2DRibbonStartShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackRotatedGrowthCurve2DRibbonStartShape", instances, (*StackRotatedGrowthCurve2DRibbonStartShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StackRotatedGrowthCurve2DStartArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StackRotatedGrowthCurve2DStartArcShape", instances, (*StackRotatedGrowthCurve2DStartArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StartArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StartArcShape", instances, (*StartArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StartArcShapeGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StartArcShapeGrid", instances, (*StartArcShapeGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StartHalfwayArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StartHalfwayArcShape", instances, (*StartHalfwayArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StartHalfwayArcShapeGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StartHalfwayArcShapeGrid", instances, (*StartHalfwayArcShapeGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StemCylinder3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StemCylinder3DShape", instances, (*StemCylinder3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Stool2DDiagrams {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Stool2DDiagram", instances, (*Stool2DDiagram)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Stool3DDiagrams {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Stool3DDiagram", instances, (*Stool3DDiagram)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StoolAbstracts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StoolAbstract", instances, (*StoolAbstract)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TiledFloor3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TiledFloor3DShape", instances, (*TiledFloor3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopEndArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopEndArcShape", instances, (*TopEndArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopEndArcShapeGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopEndArcShapeGrid", instances, (*TopEndArcShapeGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopEndHalfwayArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopEndHalfwayArcShape", instances, (*TopEndHalfwayArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopEndHalfwayArcShapeGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopEndHalfwayArcShapeGrid", instances, (*TopEndHalfwayArcShapeGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopGrowthCurve2Ds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopGrowthCurve2D", instances, (*TopGrowthCurve2D)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopMidArcVectorShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopMidArcVectorShape", instances, (*TopMidArcVectorShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopMidArcVectorShapeGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopMidArcVectorShapeGrid", instances, (*TopMidArcVectorShapeGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopStackGrowthCurve2DEndHalfwayArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopStackGrowthCurve2DEndHalfwayArcShape", instances, (*TopStackGrowthCurve2DEndHalfwayArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopStackGrowthCurve2DStartHalfwayArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopStackGrowthCurve2DStartHalfwayArcShape", instances, (*TopStackGrowthCurve2DStartHalfwayArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopStackOfGrowthCurve2Ds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopStackOfGrowthCurve2D", instances, (*TopStackOfGrowthCurve2D)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopStackOfRotatedGrowthCurve2Ds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopStackOfRotatedGrowthCurve2D", instances, (*TopStackOfRotatedGrowthCurve2D)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopStackOfRotatedGrowthCurve2DEndArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopStackOfRotatedGrowthCurve2DEndArcShape", instances, (*TopStackOfRotatedGrowthCurve2DEndArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopStackOfRotatedGrowthCurve2DStartArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopStackOfRotatedGrowthCurve2DStartArcShape", instances, (*TopStackOfRotatedGrowthCurve2DStartArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopStartArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopStartArcShape", instances, (*TopStartArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopStartArcShapeGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopStartArcShapeGrid", instances, (*TopStartArcShapeGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopStartHalfwayArcShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopStartHalfwayArcShape", instances, (*TopStartHalfwayArcShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopStartHalfwayArcShapeGrids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopStartHalfwayArcShapeGrid", instances, (*TopStartHalfwayArcShapeGrid)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Torus3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Torus3DShape", instances, (*Torus3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TorusEdge3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TorusEdge3DShape", instances, (*TorusEdge3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TorusStackShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TorusStackShape", instances, (*TorusStackShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TubeVase3DDiagrams {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TubeVase3DDiagram", instances, (*TubeVase3DDiagram)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TubeVaseAbstracts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TubeVaseAbstract", instances, (*TubeVaseAbstract)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Vase2DDiagrams {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Vase2DDiagram", instances, (*Vase2DDiagram)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalTorusStackShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalTorusStackShape", instances, (*VerticalTorusStackShape)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VolumeKey3DShapes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VolumeKey3DShape", instances, (*VolumeKey3DShape)(nil).GongGetFieldHeaders(), addIDs)
		}
	}

	// Create a style with wrap text enabled
	wrapStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})
	_ = wrapStyle
	if err != nil {
		fmt.Println("failed to create style:", err)
		return f
	}

	// Create a style with bold text
	boldStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	_ = boldStyle
	if err != nil {
		fmt.Println("failed to create bold style:", err)
		return f
	}

	// Get all sheet names
	sheetList := f.GetSheetList()

	for _, sheet := range sheetList {
		// Use a lazy iterator instead of loading all rows into memory
		rows, err := f.Rows(sheet)
		if err != nil {
			fmt.Printf("failed to get rows iterator for sheet %q: %v\n", sheet, err)
			continue
		}

		// Check if there is at least one row, and move the iterator to it
		if !rows.Next() {
			rows.Close() // Always close iterators
			continue
		}

		// Read ONLY the first row
		firstRow, err := rows.Columns()

		// Close the iterator immediately since we don't need the rest of the sheet
		rows.Close()

		if err != nil {
			fmt.Printf("failed to get columns for sheet %q: %v\n", sheet, err)
			continue
		}

		// If the first row is completely empty, skip
		if len(firstRow) == 0 {
			continue
		}

		// Track the first and last “used” column in the first row,
		// so we can later apply an AutoFilter from the first to last used col
		var firstUsedColIdx, lastUsedColIdx int

		for colIdx, cellValue := range firstRow {
			if cellValue == "" {
				// Skip columns with empty first-row cells
				continue
			}

			// Convert zero-based colIdx to 1-based for Excelize,
			// then get the column name (A, B, C, etc.)
			colName, err := excelize.ColumnNumberToName(colIdx + 1)
			if err != nil {
				fmt.Printf("failed to convert column number: %v\n", err)
				continue
			}

			// Apply wrap-text style to this entire column
			colRange := colName + ":" + colName
			if err := f.SetColStyle(sheet, colRange, wrapStyle); err != nil {
				fmt.Printf("failed to set col style on %s: %v\n", colRange, err)
				continue
			}

			// Make the first row (cell in row 1) bold in this column
			cellRef := fmt.Sprintf("%s1", colName)
			if err := f.SetCellStyle(sheet, cellRef, cellRef, boldStyle); err != nil {
				fmt.Printf("failed to set cell style on %s: %v\n", cellRef, err)
				continue
			}

			// Update our “first used” and “last used” column indices
			if firstUsedColIdx == 0 {
				firstUsedColIdx = colIdx + 1
			}
			if colIdx+1 > lastUsedColIdx {
				lastUsedColIdx = colIdx + 1
			}
		}

		// If we found at least one non-empty column in row 1, enable AutoFilter
		if firstUsedColIdx != 0 && lastUsedColIdx >= firstUsedColIdx {
			startCol, _ := excelize.ColumnNumberToName(firstUsedColIdx)
			endCol, _ := excelize.ColumnNumberToName(lastUsedColIdx)
			styleRange := fmt.Sprintf("%s:%s", startCol, endCol)
			autoFilterRange := fmt.Sprintf("%s1:%s1", startCol, endCol)
			startCellString := fmt.Sprintf("%s1", startCol)
			endCellString := fmt.Sprintf("%s1", endCol)

			if err := f.SetColStyle(sheet, styleRange, wrapStyle); err != nil {
				fmt.Println("failed to set column style:", err)
				return f
			}

			// Apply the bold style to the first row (A1:XFD1)
			if err := f.SetCellStyle(sheet, startCellString, endCellString, boldStyle); err != nil {
				fmt.Println("failed to set bold style:", err)
				return f
			}

			var opts []excelize.AutoFilterOptions
			if err := f.AutoFilter(sheet, autoFilterRange, opts); err != nil {
				fmt.Printf("failed to enable auto filter on range %s: %v\n", autoFilterRange, err)
			}
		}
	}

	var tab ExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
	}
	return f
}

// SerializeStageAsBytes serializes the stage to a pure in-memory Excel file and returns the bytes.
func (stage *Stage) SerializeStageAsBytes(addIDs bool) ([]byte, error) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func __gong__shortenString(s string) string {
	if len(s) > 31 {
		return s[:31]
	}
	return s
}

// Tabulator is an interface for writing to a table strings
type Tabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
}

type ExcelizeTabulator struct {
	f *excelize.File
}

func (tab *ExcelizeTabulator) SetExcelizeFile(f *excelize.File) {
	tab.f = f
}

func (tab *ExcelizeTabulator) AddSheet(sheetName string) {

}

func (tab *ExcelizeTabulator) AddRow(sheetName string) (rowId int) {
	return
}

func (tab *ExcelizeTabulator) AddCell(sheetName string, rowId, columnIndex int, value string) {

}

// SerializeExcelize is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelize(f *excelize.File, name string, instances []GongstructIF, fields []GongFieldHeader, addIDs bool) {
	sheetName := __gong__shortenString(name)

	// Create a new sheet.
	f.NewSheet(sheetName)

	sortedSlice := make([]GongstructIF, len(instances))
	copy(sortedSlice, instances)
	slices.SortFunc(sortedSlice, func(a, b GongstructIF) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	line := 1

	for index, fieldHeader := range fields {
		if !addIDs {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldHeader.Name)
		} else {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldHeader.Name)
			switch fieldHeader.GongFieldValueType {
			case GongFieldValueTypePointer:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":ID")
			case GongFieldValueTypeSliceOfPointers:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":IDs")
			default:
				// if index is 0, this is the ID of the instance
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), "ID")
				} else {
					// one have to put the type of the cell
					header := fieldHeader.Name
					switch fieldHeader.GongFieldValueType {
					case GongFieldValueTypeInt:
						header += ":int"
					case GongFieldValueTypeIntDuration:
						header += ":duration"
					case GongFieldValueTypeFloat:
						header += ":float"
					case GongFieldValueTypeBool:
						header += ":bool"
					case GongFieldValueTypeString:
						header += ":string"
					case GongFieldValueTypeDate:
						header += ":date"
					default:
						header += ":basicType"
					}
					header += ":noID"
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), header)
				}
			}
		}
	}

	// AutoFilter starting from A1
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", GongIntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for _, instance := range sortedSlice {
		line = line + 1

		// 3. Add the ID value in column B

		for index, fieldName := range fields {
			fieldStringValue := stage.GetFieldStringValueFromPointer(instance, fieldName.Name)
			if !addIDs {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldStringValue.GetValueString())
			} else {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldStringValue.GetValueString())
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), instance.GongGetUUID(stage))
				} else {
					switch fieldStringValue.GongFieldValueType {
					case GongFieldValueTypePointer, GongFieldValueTypeSliceOfPointers:
						f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), fieldStringValue.ids)
					}
				}

			}
		}
	}
}

// SerializeExcelizePointer is the Stage method for Excel serialization.
func (stage *Stage) SerializeExcelizePointer[Type PointerToGongstruct](f *excelize.File) {
	stage.SerializeExcelizePointer2[Type](f, false)
}

// SerializeExcelizePointer2 is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelizePointer2[Type PointerToGongstruct](f *excelize.File, addIDs bool) {
	var ret Type
	set := *stage.GetInstancesSet[Type]()
	var instances []GongstructIF
	for key := range set {
		instances = append(instances, key)
	}
	stage.SerializeExcelize(f, ret.GongGetGongstructName(), instances, ret.GongGetFieldHeaders(), addIDs)
}
