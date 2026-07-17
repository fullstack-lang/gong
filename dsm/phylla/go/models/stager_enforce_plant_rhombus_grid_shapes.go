package models

import (
	"math"
)

func (stager *Stager) enforcePlantRhombusGridShapeHasRhombuses() (needCommit bool) {
	stage := stager.stage

	for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
		angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
		length := plant.RhombusSideLength

		// Cartesian Y-axis points UP
		v1x := length * math.Cos(angleRad/2.0)
		v1y := length * math.Sin(angleRad/2.0)

		v2x := -length * math.Cos(angleRad/2.0)
		v2y := length * math.Sin(angleRad/2.0)

		{
			needCommit = enforceInitialRhombusGridShapeHasRhombuses(stage, plant.InitialRhombusGridShape, plant.N, plant.M, v1x, v1y, v2x, v2y, 0.0) || needCommit
		}
		rotRad := 0.0
		{
			rotRad = -plant.PlantCircumferenceShape.AngleDegree * math.Pi / 180.0
		}
		{
			needCommit = enforceRotatedRhombusGridShapeHasRhombuses(stage, plant.RotatedRhombusGridShape2, plant.N, plant.M, v1x, v1y, v2x, v2y, rotRad) || needCommit
		}
		{
			needCommit = enforceGrowthPathRhombusGridShapeHasRhombuses(stage, plant.GrowthCurveRhombusGridShape, plant.RotatedRhombusGridShape2, plant.PlantCircumferenceShape.AngleDegree, plant.RhombusSideLength, plant.RhombusInsideAngle, plant.PlantCircumferenceShape.Length) || needCommit
		}

		{
			rhombuses := plant.GrowthCurveRhombusGridShape.GrowthCurveRhombusShapes
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
			needCommit = enforcePerpendicularVectorGridHasVectors(stage, plant.PerpendicularVectorGrid, plant.GrowthCurveRhombusGridShape, v1x, v1y, v2x, v2y, rotRad) || needCommit
		}

		{
			needCommit = enforcePerpendicularVectorGridHalfwayHasVectors(stage, plant.PerpendicularVectorGridHalfway, plant.PerpendicularVectorGrid) || needCommit
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
			needCommit = enforceTopStartArcShapeV2GridHasShapes(stage, plant.TopStartArcShapeGrid, plant.PerpendicularVectorGrid, plant.RelativeVerticalThickness*plant.RhombusSideLength) || needCommit
		}

		if plant.ShiftedBottomTopStartArcShapeGrid == nil {
			plant.ShiftedBottomTopStartArcShapeGrid = new(ShiftedBottomTopStartArcShapeGrid).Stage(stage)
			plant.ShiftedBottomTopStartArcShapeGrid.Name = plant.Name + "-ShiftedBottomTopStartArcShapeGrid"
			needCommit = true
		}
		{
			needCommit = enforceShiftedBottomTopStartArcShapeV2GridHasShapes(stage, plant.ShiftedBottomTopStartArcShapeGrid, plant.PerpendicularVectorGrid, plant.RelativeVerticalThickness*plant.RhombusSideLength) || needCommit
		}

		{
			needCommit = enforceEndArcShapeV2GridHasShapes(stage, plant.EndArcShapeGrid, plant.PerpendicularVectorGrid) || needCommit
		}

		{
			needCommit = enforceTopEndArcShapeV2GridHasShapes(stage, plant.TopEndArcShapeGrid, plant.PerpendicularVectorGrid, plant.RelativeVerticalThickness*plant.RhombusSideLength) || needCommit
		}

		if plant.MidArcVectorShapeGrid == nil {
			plant.MidArcVectorShapeGrid = new(MidArcVectorShapeGrid).Stage(stage)
			plant.MidArcVectorShapeGrid.Name = plant.Name + "-MidArcVectorShapeGrid"
			needCommit = true
		}
		{
			needCommit = enforceMidArcVectorShapeGridHasShapes(stage, plant.MidArcVectorShapeGrid, plant.PerpendicularVectorGrid, plant.RelativeVerticalThickness*plant.RhombusSideLength) || needCommit
		}

		if plant.TopMidArcVectorShapeGrid == nil {
			plant.TopMidArcVectorShapeGrid = new(TopMidArcVectorShapeGrid).Stage(stage)
			plant.TopMidArcVectorShapeGrid.Name = plant.Name + "-TopMidArcVectorShapeGrid"
			needCommit = true
		}
		{
			needCommit = enforceTopMidArcVectorShapeGridHasShapes(stage, plant.TopMidArcVectorShapeGrid, plant.PerpendicularVectorGrid, plant.RelativeVerticalThickness*plant.RhombusSideLength) || needCommit
		}

		if plant.StartHalfwayArcShapeGrid == nil {
			plant.StartHalfwayArcShapeGrid = new(StartHalfwayArcShapeGrid).Stage(stage)
			plant.StartHalfwayArcShapeGrid.Name = plant.Name + "-HalfwayArcShapeGrid"
			needCommit = true
		}
		if plant.EndHalfwayArcShapeGrid == nil {
			plant.EndHalfwayArcShapeGrid = new(EndHalfwayArcShapeGrid).Stage(stage)
			plant.EndHalfwayArcShapeGrid.Name = plant.Name + "-EndHalfwayArcShapeGrid"
			needCommit = true
		}
		{
			needCommit = enforceHalfwayArcShapeGridHasShapes(stage, plant.StartHalfwayArcShapeGrid, plant.PerpendicularVectorGrid, plant.RelativeVerticalThickness*plant.RhombusSideLength) || needCommit
			needCommit = enforceEndHalfwayArcShapeGridHasShapes(stage, plant.EndHalfwayArcShapeGrid, plant.PerpendicularVectorGrid, plant.RelativeVerticalThickness*plant.RhombusSideLength) || needCommit
		}

		if plant.GrowthCurve2D == nil {
			plant.GrowthCurve2D = new(GrowthCurve2D).Stage(stage)
			plant.GrowthCurve2D.Name = plant.Name + "-GrowthCurve2D"
			needCommit = true
		}
		{
			if plant.GrowthCurve2D.StartArcShapeGrid != plant.StartArcShapeGrid {
				plant.GrowthCurve2D.StartArcShapeGrid = plant.StartArcShapeGrid
				needCommit = true
			}
			if plant.GrowthCurve2D.EndArcShapeGrid != plant.EndArcShapeGrid {
				plant.GrowthCurve2D.EndArcShapeGrid = plant.EndArcShapeGrid
				needCommit = true
			}
		}

		if plant.TopGrowthCurve2D == nil {
			plant.TopGrowthCurve2D = new(TopGrowthCurve2D).Stage(stage)
			plant.TopGrowthCurve2D.Name = plant.Name + "-TopGrowthCurve2D"
			needCommit = true
		}
		{
			if plant.TopGrowthCurve2D.TopStartArcShapeGrid != plant.TopStartArcShapeGrid {
				plant.TopGrowthCurve2D.TopStartArcShapeGrid = plant.TopStartArcShapeGrid
				needCommit = true
			}
			if plant.TopGrowthCurve2D.TopEndArcShapeGrid != plant.TopEndArcShapeGrid {
				plant.TopGrowthCurve2D.TopEndArcShapeGrid = plant.TopEndArcShapeGrid
				needCommit = true
			}
		}

		{
		}

		{
		}

		circLen := 0.0
		{
			circLen = plant.PlantCircumferenceShape.Length
		}

		{
			needCommit = enforceGrowthCurveBezierShapeGridHasShapes(stage, plant.GrowthCurveBezierShapeGrid, plant.PerpendicularVectorGrid, plant.RhombusSideLength, circLen) || needCommit
		}

		{
			needCommit = enforceStackOfGrowthCurveV2HasShapes(stage, plant.StackOfGrowthCurve, plant.PerpendicularVectorGrid, plant.GrowthVectorShape, plant.StackHeight, circLen, plant.RelativeVerticalThickness*plant.RhombusSideLength) || needCommit
		}
		if plant.TopStackOfGrowthCurve != nil {
			needCommit = enforceTopStackOfGrowthCurveV2HasShapes(stage, plant.TopStackOfGrowthCurve, plant.PerpendicularVectorGrid, plant.GrowthVectorShape, plant.StackHeight, circLen, plant.RelativeVerticalThickness*plant.RhombusSideLength) || needCommit
		}
		if plant.ShiftedLeftStackOfNormalVector == nil {
			plant.ShiftedLeftStackOfNormalVector = new(ShiftedLeftStackOfNormalVector).Stage(stage)
			plant.ShiftedLeftStackOfNormalVector.Name = plant.Name + "-ShiftedLeftStackOfNormalVector"
			needCommit = true
		}
		{
			needCommit = enforceShiftedLeftStackOfNormalVectorHasShapes(stage, plant.ShiftedLeftStackOfNormalVector, plant.ArcNormalVectorShapeGrid, plant.PerpendicularVectorGrid, plant.GrowthVectorShape, plant.StackHeight, circLen, plant.RelativeVerticalThickness*plant.RhombusSideLength) || needCommit
		}
		if plant.ShiftedLeftStackOfGrowthCurve == nil {
			plant.ShiftedLeftStackOfGrowthCurve = new(ShiftedLeftStackOfGrowthCurve).Stage(stage)
			plant.ShiftedLeftStackOfGrowthCurve.Name = plant.Name + "-ShiftedLeftStackOfGrowthCurve"
			needCommit = true
		}
		{
			needCommit = enforceShiftedLeftStackOfGrowthCurveV2HasShapes(stage, plant.ShiftedLeftStackOfGrowthCurve, plant.PerpendicularVectorGrid, plant.GrowthVectorShape, plant.StackHeight, circLen, plant.RelativeVerticalThickness*plant.RhombusSideLength) || needCommit
		}
		{
		}
	}
	return
}
