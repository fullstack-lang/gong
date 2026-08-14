package models

import (
	"math"
)

func (stager *Stager) enforcePlantRhombusGridShapeHasRhombuses() (needCommit bool) {
	stage := stager.stage

	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stage) {
		angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
		length := plant.RhombusSideLength * stager.GetPlant2DZoom(plant)

		// Cartesian Y-axis points UP
		v1x := length * math.Cos(angleRad/2.0)
		v1y := length * math.Sin(angleRad/2.0)

		v2x := -length * math.Cos(angleRad/2.0)
		v2y := length * math.Sin(angleRad/2.0)

		{
			needCommit = enforceInitialRhombusGridShapeHasRhombuses(stage, plant.RhombusStuff.InitialRhombusGridShape, plant.N, plant.M, v1x, v1y, v2x, v2y, 0.0) || needCommit
		}
		rotRad := 0.0
		{
			rotRad = -plant.RhombusStuff.PlantCircumferenceShape.AngleDegree * math.Pi / 180.0
		}
		{
			needCommit = enforceRotatedRhombusGridShapeHasRhombuses(stage, plant.RhombusStuff.RotatedRhombusGridShape2, plant.N, plant.M, v1x, v1y, v2x, v2y, rotRad) || needCommit
		}
		{
			needCommit = enforceGrowthPathRhombusGridShapeHasRhombuses(stage, plant.RhombusStuff.GrowthCurveRhombusGridShape, plant.RhombusStuff.RotatedRhombusGridShape2, plant.RhombusStuff.PlantCircumferenceShape.AngleDegree, length, plant.RhombusInsideAngle, plant.RhombusStuff.PlantCircumferenceShape.Length) || needCommit
		}

		{
			rhombuses := plant.RhombusStuff.GrowthCurveRhombusGridShape.GrowthCurveRhombusShapes
			if len(rhombuses) >= 2 {
				first := rhombuses[0]
				var endRhombus *GrowthCurveRhombusShape
				minDiffY := math.MaxFloat64
				for i := 1; i < len(rhombuses); i++ {
					r := rhombuses[i]
					if r.Y > first.Y {
						diff := r.Y - first.Y
						if diff < minDiffY {
							minDiffY = diff
							endRhombus = r
						}
					}
				}

				if endRhombus != nil {
					vx := endRhombus.X - first.X
					vy := endRhombus.Y - first.Y
					if plant.GrowthVectorShape.X != vx || plant.GrowthVectorShape.Y != vy {
						plant.GrowthVectorShape.X = vx
						plant.GrowthVectorShape.Y = vy
						needCommit = true
					}
				} else {
					if plant.GrowthVectorShape.X != 0 || plant.GrowthVectorShape.Y != 0 {
						plant.GrowthVectorShape.X = 0
						plant.GrowthVectorShape.Y = 0
						needCommit = true
					}
				}
			} else {
				if plant.GrowthVectorShape.X != 0 || plant.GrowthVectorShape.Y != 0 {
					plant.GrowthVectorShape.X = 0
					plant.GrowthVectorShape.Y = 0
					needCommit = true
				}
			}
		}

		{
			needCommit = enforcePerpendicularVectorGridHasVectors(stage, plant.PerpendicularVectorGrid, plant.RhombusStuff.GrowthCurveRhombusGridShape, v1x, v1y, v2x, v2y, rotRad) || needCommit
		}

		{
			needCommit = enforceBaseVectorShapeGridHasShapes(stage, plant.BaseVectorShapeGrid, plant.PerpendicularVectorGrid) || needCommit
		}

		{
			needCommit = enforceArcNormalVectorShapeGridHasShapes(stage, plant.ArcNormalVectorShapeGrid, plant.PerpendicularVectorGrid) || needCommit
		}

		{
			needCommit = enforceStartArcShapeV2GridHasShapes(stage, plant.StartArcShapeGrid, plant.PerpendicularVectorGrid) || needCommit
		}

		{
			needCommit = enforceEndArcShapeV2GridHasShapes(stage, plant.EndArcShapeGrid, plant.PerpendicularVectorGrid) || needCommit
		}

		if plant.MidArcVectorShapeGrid == nil {
			plant.MidArcVectorShapeGrid = new(MidArcVectorShapeGrid).Stage(stage)
			plant.MidArcVectorShapeGrid.Name = plant.Name + "-MidArcVectorShapeGrid"
			needCommit = true
		}
		{
			vThicknessMid := 0.0
			if plant.PlantType == Vase && plant.VaseAbstract != nil {
				vThicknessMid = plant.VaseAbstract.RelativeVerticalThickness * plant.RhombusSideLength
			}
			needCommit = enforceMidArcVectorShapeGridHasShapes(stage, plant.MidArcVectorShapeGrid, plant.PerpendicularVectorGrid, vThicknessMid) || needCommit
		}

		if plant.GrowthCurve2D == nil {
			plant.GrowthCurve2D = new(GrowthCurve2D).Stage(stage)
			plant.GrowthCurve2D.Name = plant.Name + "-GrowthCurve2D"
			needCommit = true
		}

		if plant.PlantType == Vase && plant.VaseAbstract != nil {
			vase := plant.VaseAbstract
			vThickness := vase.RelativeVerticalThickness * plant.RhombusSideLength
			cHeight := vase.RelativeCuttedStackFloorHeight * plant.RhombusSideLength

			{
				needCommit = enforcePerpendicularVectorGridHalfwayHasVectors(stage, vase.PerpendicularVectorGridHalfway, plant.PerpendicularVectorGrid) || needCommit
			}

			{
				needCommit = enforceTopStartArcShapeV2GridHasShapes(stage, vase.TopStartArcShapeGrid, plant.PerpendicularVectorGrid, vThickness) || needCommit
			}

			if vase.ShiftedBottomTopStartArcShapeGrid == nil {
				vase.ShiftedBottomTopStartArcShapeGrid = new(ShiftedBottomTopStartArcShapeGrid).Stage(stage)
				vase.ShiftedBottomTopStartArcShapeGrid.Name = plant.Name + "-ShiftedBottomTopStartArcShapeGrid"
				needCommit = true
			}
			{
				needCommit = enforceShiftedBottomTopStartArcShapeV2GridHasShapes(stage, vase.ShiftedBottomTopStartArcShapeGrid, plant.PerpendicularVectorGrid, vThickness) || needCommit
			}

			{
				needCommit = enforceTopEndArcShapeV2GridHasShapes(stage, vase.TopEndArcShapeGrid, plant.PerpendicularVectorGrid, vThickness) || needCommit
			}

			if vase.TopMidArcVectorShapeGrid == nil {
				vase.TopMidArcVectorShapeGrid = new(TopMidArcVectorShapeGrid).Stage(stage)
				vase.TopMidArcVectorShapeGrid.Name = plant.Name + "-TopMidArcVectorShapeGrid"
				needCommit = true
			}
			{
				needCommit = enforceTopMidArcVectorShapeGridHasShapes(stage, vase.TopMidArcVectorShapeGrid, plant.PerpendicularVectorGrid, vThickness) || needCommit
			}

			if vase.StartHalfwayArcShapeGrid == nil {
				vase.StartHalfwayArcShapeGrid = new(StartHalfwayArcShapeGrid).Stage(stage)
				vase.StartHalfwayArcShapeGrid.Name = plant.Name + "-StartHalfwayArcShapeGrid"
				needCommit = true
			}
			if vase.TopStartHalfwayArcShapeGrid == nil {
				vase.TopStartHalfwayArcShapeGrid = new(TopStartHalfwayArcShapeGrid).Stage(stage)
				vase.TopStartHalfwayArcShapeGrid.Name = plant.Name + "-TopStartHalfwayArcShapeGrid"
				needCommit = true
			}
			if vase.EndHalfwayArcShapeGrid == nil {
				vase.EndHalfwayArcShapeGrid = new(EndHalfwayArcShapeGrid).Stage(stage)
				vase.EndHalfwayArcShapeGrid.Name = plant.Name + "-EndHalfwayArcShapeGrid"
				needCommit = true
			}
			if vase.TopEndHalfwayArcShapeGrid == nil {
				vase.TopEndHalfwayArcShapeGrid = new(TopEndHalfwayArcShapeGrid).Stage(stage)
				vase.TopEndHalfwayArcShapeGrid.Name = plant.Name + "-TopEndHalfwayArcShapeGrid"
				needCommit = true
			}
			{
				needCommit = enforceHalfwayArcShapeGridHasShapes(stage, vase.StartHalfwayArcShapeGrid, plant.PerpendicularVectorGrid, vThickness) || needCommit
				needCommit = enforceTopStartHalfwayArcShapeGridHasShapes(stage, vase.TopStartHalfwayArcShapeGrid, plant.PerpendicularVectorGrid, vThickness) || needCommit
				needCommit = enforceEndHalfwayArcShapeGridHasShapes(stage, vase.EndHalfwayArcShapeGrid, plant.PerpendicularVectorGrid, vThickness) || needCommit
				needCommit = enforceTopEndHalfwayArcShapeGridHasShapes(stage, vase.TopEndHalfwayArcShapeGrid, plant.PerpendicularVectorGrid, vThickness) || needCommit
			}

			{
				if plant.GrowthCurve2D.StartHalfwayArcShapeGrid != vase.StartHalfwayArcShapeGrid {
					plant.GrowthCurve2D.StartHalfwayArcShapeGrid = vase.StartHalfwayArcShapeGrid
					needCommit = true
				}
				if plant.GrowthCurve2D.EndHalfwayArcShapeGrid != vase.EndHalfwayArcShapeGrid {
					plant.GrowthCurve2D.EndHalfwayArcShapeGrid = vase.EndHalfwayArcShapeGrid
					needCommit = true
				}
			}

			if vase.TopGrowthCurve2D == nil {
				vase.TopGrowthCurve2D = new(TopGrowthCurve2D).Stage(stage)
				vase.TopGrowthCurve2D.Name = plant.Name + "-TopGrowthCurve2D"
				needCommit = true
			}
			{
				if vase.TopGrowthCurve2D.TopStartHalfwayArcShapeGrid != vase.TopStartHalfwayArcShapeGrid {
					vase.TopGrowthCurve2D.TopStartHalfwayArcShapeGrid = vase.TopStartHalfwayArcShapeGrid
					needCommit = true
				}
				if vase.TopGrowthCurve2D.TopEndHalfwayArcShapeGrid != vase.TopEndHalfwayArcShapeGrid {
					vase.TopGrowthCurve2D.TopEndHalfwayArcShapeGrid = vase.TopEndHalfwayArcShapeGrid
					needCommit = true
				}
			}

			circLen := 0.0
			if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil {
				circLen = plant.RhombusStuff.PlantCircumferenceShape.Length
			}

			{
				needCommit = enforceStackOfGrowthCurveV2HasShapes(stage, vase.StackOfRotatedGrowthCurve2D, vase.StartHalfwayArcShapeGrid, vase.EndHalfwayArcShapeGrid, plant.PerpendicularVectorGrid, plant.GrowthVectorShape, plant.StackHeight, circLen, vThickness) || needCommit
			}
			if vase.TopStackOfRotatedGrowthCurve2D != nil {
				needCommit = enforceTopStackOfGrowthCurveV2HasShapes(stage, vase.TopStackOfRotatedGrowthCurve2D, vase.TopStartHalfwayArcShapeGrid, vase.TopEndHalfwayArcShapeGrid, plant.PerpendicularVectorGrid, plant.GrowthVectorShape, plant.StackHeight, circLen, vThickness) || needCommit
			}
			if vase.StackOfGrowthCurve2D != nil {
				needCommit = enforceStackOfGrowthCurve2DHasShapes(stage, vase.StackOfGrowthCurve2D, vase.StartHalfwayArcShapeGrid, vase.EndHalfwayArcShapeGrid, plant.PerpendicularVectorGrid, plant.StackHeight, circLen, cHeight) || needCommit
			}
			if vase.TopStackOfGrowthCurve2D != nil {
				needCommit = enforceTopStackOfGrowthCurve2DHasShapes(stage, vase.TopStackOfGrowthCurve2D, vase.TopStartHalfwayArcShapeGrid, vase.TopEndHalfwayArcShapeGrid, plant.PerpendicularVectorGrid, plant.StackHeight, circLen, cHeight) || needCommit
			}
			if vase.StackOfGrowthCurve2DRibbon != nil {
				needCommit = enforceStackOfGrowthCurve2DRibbonHasShapes(stage, vase.StackOfGrowthCurve2DRibbon, vase.StackOfGrowthCurve2D, vase.TopStackOfGrowthCurve2D) || needCommit
			}
			if vase.StackOfRotatedGrowthCurve2DRibbon != nil {
				needCommit = enforceStackOfRotatedGrowthCurve2DRibbonHasShapes(stage, vase.StackOfRotatedGrowthCurve2DRibbon, vase.StackOfRotatedGrowthCurve2D, vase.TopStackOfRotatedGrowthCurve2D) || needCommit
			}
			if vase.PartiallyGrowthCurve2DRibbon != nil {
				needCommit = enforcePartiallyGrowthCurve2DRibbonHasShapes(stage, plant) || needCommit
			}
			if vase.ShiftedLeftPartiallyGrowthCurve2DRibbon != nil {
				needCommit = enforceShiftedLeftPartiallyGrowthCurve2DRibbonHasShapes(stage, plant) || needCommit
			}
			if vase.PartiallyGrowthCurve2DTrajectory != nil {
				needCommit = enforcePartiallyGrowthCurve2DTrajectoryHasShapes(stage, plant) || needCommit
			}
			if vase.GrowthCurve2DRibbon != nil {
				needCommit = enforceGrowthCurve2DRibbonHasShapes(stage, plant) || needCommit
			}
			if vase.ShiftedRightGrowthCurve2DRibbon != nil {
				needCommit = enforceShiftedRightGrowthCurve2DRibbonHasShapes(stage, plant) || needCommit
			}
			if vase.ShiftedLeftGrowthCurve2DRibbon != nil {
				needCommit = enforceShiftedLeftGrowthCurve2DRibbonHasShapes(stage, plant) || needCommit
			}
		}

	}
	return
}
