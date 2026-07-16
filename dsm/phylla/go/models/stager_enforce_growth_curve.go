package models

import (
	"fmt"
	"math"
)

func enforceGrowthCurveBezierShapeGridHasShapes(stage *Stage, grid *GrowthCurveBezierShapeGrid, pGrid *PerpendicularVectorGrid, sideLength float64, circumferenceLength float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.GrowthCurveBezierShapes) > 0 {
			grid.GrowthCurveBezierShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	expectedTotalLen := expectedLen * 2
	valid := true
	if len(grid.GrowthCurveBezierShapes) != expectedTotalLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			b := grid.GrowthCurveBezierShapes[i]
			bTrans := grid.GrowthCurveBezierShapes[i+expectedLen]

			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			expectedTransName := fmt.Sprintf("%s-translated-%d", grid.Name, i)
			if b == nil || b.Name != expectedName || bTrans == nil || bTrans.Name != expectedTransName {
				valid = false
				break
			}
			angleRad := math.Atan2(v1.EndY-v1.StartY, v1.EndX-v1.StartX) - math.Pi/2.0

			// wait, if we consider length ratio = 0.33
			ratio := 0.33
			expectedCtrlStartX := v1.StartX + sideLength*ratio*math.Cos(angleRad)
			expectedCtrlStartY := v1.StartY + sideLength*ratio*math.Sin(angleRad)
			expectedCtrlEndX := v2.StartX + sideLength*ratio*math.Cos(angleRad+math.Pi)
			expectedCtrlEndY := v2.StartY + sideLength*ratio*math.Sin(angleRad+math.Pi)

			if math.Abs(b.StartX-v1.StartX) > 1e-4 || math.Abs(b.StartY-v1.StartY) > 1e-4 ||
				math.Abs(b.EndX-v2.StartX) > 1e-4 || math.Abs(b.EndY-v2.StartY) > 1e-4 ||
				math.Abs(b.ControlPointStartX-expectedCtrlStartX) > 1e-4 || math.Abs(b.ControlPointStartY-expectedCtrlStartY) > 1e-4 ||
				math.Abs(b.ControlPointEndX-expectedCtrlEndX) > 1e-4 || math.Abs(b.ControlPointEndY-expectedCtrlEndY) > 1e-4 {
				valid = false
				break
			}

			if math.Abs(bTrans.StartX-(v1.StartX+circumferenceLength)) > 1e-4 || math.Abs(bTrans.StartY-v1.StartY) > 1e-4 ||
				math.Abs(bTrans.EndX-(v2.StartX+circumferenceLength)) > 1e-4 || math.Abs(bTrans.EndY-v2.StartY) > 1e-4 ||
				math.Abs(bTrans.ControlPointStartX-(expectedCtrlStartX+circumferenceLength)) > 1e-4 || math.Abs(bTrans.ControlPointStartY-expectedCtrlStartY) > 1e-4 ||
				math.Abs(bTrans.ControlPointEndX-(expectedCtrlEndX+circumferenceLength)) > 1e-4 || math.Abs(bTrans.ControlPointEndY-expectedCtrlEndY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		grid.GrowthCurveBezierShapes = make([]*GrowthCurveBezierShape, expectedTotalLen)
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			b := new(GrowthCurveBezierShape).Stage(stage)
			b.Name = fmt.Sprintf("%s-%d", grid.Name, i)

			b.StartX = v1.StartX
			b.StartY = v1.StartY
			b.EndX = v2.StartX
			b.EndY = v2.StartY

			angleRad := math.Atan2(v1.EndY-v1.StartY, v1.EndX-v1.StartX) - math.Pi/2.0
			ratio := 0.33
			b.ControlPointStartX = b.StartX + sideLength*ratio*math.Cos(angleRad)
			b.ControlPointStartY = b.StartY + sideLength*ratio*math.Sin(angleRad)
			b.ControlPointEndX = b.EndX + sideLength*ratio*math.Cos(angleRad+math.Pi)
			b.ControlPointEndY = b.EndY + sideLength*ratio*math.Sin(angleRad+math.Pi)

			grid.GrowthCurveBezierShapes[i] = b

			bTrans := new(GrowthCurveBezierShape).Stage(stage)
			bTrans.Name = fmt.Sprintf("%s-translated-%d", grid.Name, i)

			bTrans.StartX = b.StartX + circumferenceLength
			bTrans.StartY = b.StartY
			bTrans.EndX = b.EndX + circumferenceLength
			bTrans.EndY = b.EndY
			bTrans.ControlPointStartX = b.ControlPointStartX + circumferenceLength
			bTrans.ControlPointStartY = b.ControlPointStartY
			bTrans.ControlPointEndX = b.ControlPointEndX + circumferenceLength
			bTrans.ControlPointEndY = b.ControlPointEndY

			grid.GrowthCurveBezierShapes[i+expectedLen] = bTrans
		}
		needCommit = true
	}

	return needCommit
}

func enforceStackOfGrowthCurveV2HasShapes(stage *Stage, stack *StackOfGrowthCurve, pGrid *PerpendicularVectorGrid, vector *GrowthVectorShape, stackHeight int, circLen float64, thickness float64) (needCommit bool) {
	if stack == nil || pGrid == nil || vector == nil || stackHeight < 1 || circLen <= 0 || len(pGrid.PerpendicularVectors) < 2 {
		if len(stack.StackGrowthCurveStartArcShapes) > 0 || len(stack.StackGrowthCurveEndArcShapes) > 0 {
			stack.StackGrowthCurveStartArcShapes = nil
			stack.StackGrowthCurveEndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1

	type expectedStartShape struct {
		name                       string
		startX, startY, endX, endY float64
		radiusX, radiusY           float64
		xAxisRotation              float64
		largeArcFlag, sweepFlag    bool
	}
	var expectedStart []expectedStartShape
	var expectedEnd []expectedStartShape

	vFirst := pGrid.PerpendicularVectors[0]
	vx := vFirst.EndX - vFirst.StartX
	vy := vFirst.EndY - vFirst.StartY
	vLen := math.Hypot(vx, vy)
	if vLen == 0 {
		vLen = 1
	}
	vx, vy = vx/vLen, vy/vLen

	for h := 0; h < stackHeight; h++ {
		dx := float64(h)*vector.X + float64(h)*thickness*vx
		dy := float64(h)*vector.Y + float64(h)*thickness*vy
		offset := 0.0

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			currentDX := math.Mod(dx, circLen)
			if currentDX < 0 {
				currentDX += circLen
			}

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp := computeArcV2Geometry(v1, v2, offset, false)
			expectedStart = append(expectedStart, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-start-%d-%d", stack.Name, h, i),
				startX: sX + currentDX, startY: sY + dy,
				endX: eX + currentDX, endY: eY + dy,
				radiusX: rX, radiusY: rY,
				xAxisRotation: xRot, largeArcFlag: lArc, sweepFlag: swp,
			})

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp = computeArcV2Geometry(v1, v2, offset, true)
			expectedEnd = append(expectedEnd, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-end-%d-%d", stack.Name, h, i),
				startX: sX + currentDX, startY: sY + dy,
				endX: eX + currentDX, endY: eY + dy,
				radiusX: rX, radiusY: rY,
				xAxisRotation: xRot, largeArcFlag: lArc, sweepFlag: swp,
			})
		}
	}

	valid := true
	if len(stack.StackGrowthCurveStartArcShapes) != len(expectedStart) || len(stack.StackGrowthCurveEndArcShapes) != len(expectedEnd) {
		valid = false
	} else {
		for i, exp := range expectedStart {
			b := stack.StackGrowthCurveStartArcShapes[i]
			if b == nil || b.Name != exp.name {
				valid = false
				break
			}
			if math.Abs(b.StartX-exp.startX) > 1e-4 || math.Abs(b.StartY-exp.startY) > 1e-4 ||
				math.Abs(b.EndX-exp.endX) > 1e-4 || math.Abs(b.EndY-exp.endY) > 1e-4 {
				valid = false
				break
			}
		}
		for i, exp := range expectedEnd {
			b := stack.StackGrowthCurveEndArcShapes[i]
			if b == nil || b.Name != exp.name {
				valid = false
				break
			}
			if math.Abs(b.StartX-exp.startX) > 1e-4 || math.Abs(b.StartY-exp.startY) > 1e-4 ||
				math.Abs(b.EndX-exp.endX) > 1e-4 || math.Abs(b.EndY-exp.endY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		stack.StackGrowthCurveStartArcShapes = make([]*StackGrowthCurveStartArcShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(StackGrowthCurveStartArcShape).Stage(stage)
			b.Name = exp.name
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			b.RadiusX = exp.radiusX
			b.RadiusY = exp.radiusY
			b.XAxisRotation = exp.xAxisRotation
			b.LargeArcFlag = exp.largeArcFlag
			b.SweepFlag = exp.sweepFlag
			stack.StackGrowthCurveStartArcShapes[i] = b
		}
		stack.StackGrowthCurveEndArcShapes = make([]*StackGrowthCurveEndArcShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(StackGrowthCurveEndArcShape).Stage(stage)
			b.Name = exp.name
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			b.RadiusX = exp.radiusX
			b.RadiusY = exp.radiusY
			b.XAxisRotation = exp.xAxisRotation
			b.LargeArcFlag = exp.largeArcFlag
			b.SweepFlag = exp.sweepFlag
			stack.StackGrowthCurveEndArcShapes[i] = b
		}
		needCommit = true
	}

	return needCommit
}

func enforceTopStackOfGrowthCurveV2HasShapes(stage *Stage, stack *TopStackOfGrowthCurve, pGrid *PerpendicularVectorGrid, vector *GrowthVectorShape, stackHeight int, circLen float64, thickness float64) (needCommit bool) {
	if stack == nil || pGrid == nil || vector == nil || stackHeight < 1 || circLen <= 0 || len(pGrid.PerpendicularVectors) < 2 {
		if len(stack.TopStackGrowthCurveStartArcShapes) > 0 || len(stack.TopStackGrowthCurveEndArcShapes) > 0 {
			stack.TopStackGrowthCurveStartArcShapes = nil
			stack.TopStackGrowthCurveEndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1

	type expectedStartShape struct {
		name                       string
		startX, startY, endX, endY float64
		radiusX, radiusY           float64
		xAxisRotation              float64
		largeArcFlag, sweepFlag    bool
	}
	var expectedStart []expectedStartShape
	var expectedEnd []expectedStartShape

	vFirst := pGrid.PerpendicularVectors[0]
	vx := vFirst.EndX - vFirst.StartX
	vy := vFirst.EndY - vFirst.StartY
	vLen := math.Hypot(vx, vy)
	if vLen == 0 {
		vLen = 1
	}
	vx, vy = vx/vLen, vy/vLen

	for h := 0; h < stackHeight; h++ {
		dx := float64(h)*vector.X + float64(h)*thickness*vx
		dy := float64(h)*vector.Y + float64(h)*thickness*vy
		offset := thickness

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			currentDX := math.Mod(dx, circLen)
			if currentDX < 0 {
				currentDX += circLen
			}

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp := computeArcV2Geometry(v1, v2, offset, false)
			expectedStart = append(expectedStart, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-start-%d-%d", stack.Name, h, i),
				startX: sX + currentDX, startY: sY + dy,
				endX: eX + currentDX, endY: eY + dy,
				radiusX: rX, radiusY: rY,
				xAxisRotation: xRot, largeArcFlag: lArc, sweepFlag: swp,
			})

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp = computeArcV2Geometry(v1, v2, offset, true)
			expectedEnd = append(expectedEnd, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-end-%d-%d", stack.Name, h, i),
				startX: sX + currentDX, startY: sY + dy,
				endX: eX + currentDX, endY: eY + dy,
				radiusX: rX, radiusY: rY,
				xAxisRotation: xRot, largeArcFlag: lArc, sweepFlag: swp,
			})
		}
	}

	valid := true
	if len(stack.TopStackGrowthCurveStartArcShapes) != len(expectedStart) || len(stack.TopStackGrowthCurveEndArcShapes) != len(expectedEnd) {
		valid = false
	} else {
		for i, exp := range expectedStart {
			b := stack.TopStackGrowthCurveStartArcShapes[i]
			if b == nil || b.Name != exp.name {
				valid = false
				break
			}
			if math.Abs(b.StartX-exp.startX) > 1e-4 || math.Abs(b.StartY-exp.startY) > 1e-4 ||
				math.Abs(b.EndX-exp.endX) > 1e-4 || math.Abs(b.EndY-exp.endY) > 1e-4 {
				valid = false
				break
			}
		}
		for i, exp := range expectedEnd {
			b := stack.TopStackGrowthCurveEndArcShapes[i]
			if b == nil || b.Name != exp.name {
				valid = false
				break
			}
			if math.Abs(b.StartX-exp.startX) > 1e-4 || math.Abs(b.StartY-exp.startY) > 1e-4 ||
				math.Abs(b.EndX-exp.endX) > 1e-4 || math.Abs(b.EndY-exp.endY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		stack.TopStackGrowthCurveStartArcShapes = make([]*TopStackGrowthCurveStartArcShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(TopStackGrowthCurveStartArcShape).Stage(stage)
			b.Name = exp.name
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			b.RadiusX = exp.radiusX
			b.RadiusY = exp.radiusY
			b.XAxisRotation = exp.xAxisRotation
			b.LargeArcFlag = exp.largeArcFlag
			b.SweepFlag = exp.sweepFlag
			stack.TopStackGrowthCurveStartArcShapes[i] = b
		}
		stack.TopStackGrowthCurveEndArcShapes = make([]*TopStackGrowthCurveEndArcShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(TopStackGrowthCurveEndArcShape).Stage(stage)
			b.Name = exp.name
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			b.RadiusX = exp.radiusX
			b.RadiusY = exp.radiusY
			b.XAxisRotation = exp.xAxisRotation
			b.LargeArcFlag = exp.largeArcFlag
			b.SweepFlag = exp.sweepFlag
			stack.TopStackGrowthCurveEndArcShapes[i] = b
		}
		needCommit = true
	}

	return needCommit
}

func enforceShiftedLeftStackOfGrowthCurveV2HasShapes(stage *Stage, stack *ShiftedLeftStackOfGrowthCurve, pGrid *PerpendicularVectorGrid, vector *GrowthVectorShape, stackHeight int, circLen float64, thickness float64) (needCommit bool) {
	if stack == nil || pGrid == nil || vector == nil || stackHeight < 1 || circLen <= 0 || len(pGrid.PerpendicularVectors) < 2 {
		if len(stack.ShiftedLeftStackGrowthCurveStartArcShapes) > 0 || len(stack.ShiftedLeftStackGrowthCurveEndArcShapes) > 0 {
			stack.ShiftedLeftStackGrowthCurveStartArcShapes = nil
			stack.ShiftedLeftStackGrowthCurveEndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1

	type expectedStartShape struct {
		name                       string
		startX, startY, endX, endY float64
		radiusX, radiusY           float64
		xAxisRotation              float64
		largeArcFlag, sweepFlag    bool
	}
	var expectedStart []expectedStartShape
	var expectedEnd []expectedStartShape

	vFirst := pGrid.PerpendicularVectors[0]
	vx := vFirst.EndX - vFirst.StartX
	vy := vFirst.EndY - vFirst.StartY
	vLen := math.Hypot(vx, vy)
	if vLen == 0 {
		vLen = 1
	}
	vx, vy = vx/vLen, vy/vLen

	for h := 0; h < stackHeight; h++ {
		dx := float64(h)*vector.X + float64(h)*thickness*vx
		dy := float64(h)*vector.Y + float64(h)*thickness*vy
		offset := 0.0

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			currentDX := math.Mod(dx, circLen)
			if currentDX < 0 {
				currentDX += circLen
			}

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp := computeArcV2Geometry(v1, v2, offset, false)
			expectedStart = append(expectedStart, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-start-%d-%d", stack.Name, h, i),
				startX: sX + currentDX - circLen, startY: sY + dy,
				endX: eX + currentDX - circLen, endY: eY + dy,
				radiusX: rX, radiusY: rY,
				xAxisRotation: xRot, largeArcFlag: lArc, sweepFlag: swp,
			})

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp = computeArcV2Geometry(v1, v2, offset, true)
			expectedEnd = append(expectedEnd, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-end-%d-%d", stack.Name, h, i),
				startX: sX + currentDX - circLen, startY: sY + dy,
				endX: eX + currentDX - circLen, endY: eY + dy,
				radiusX: rX, radiusY: rY,
				xAxisRotation: xRot, largeArcFlag: lArc, sweepFlag: swp,
			})
		}
	}

	valid := true
	if len(stack.ShiftedLeftStackGrowthCurveStartArcShapes) != len(expectedStart) || len(stack.ShiftedLeftStackGrowthCurveEndArcShapes) != len(expectedEnd) {
		valid = false
	} else {
		for i, exp := range expectedStart {
			b := stack.ShiftedLeftStackGrowthCurveStartArcShapes[i]
			if b == nil || b.Name != exp.name {
				valid = false
				break
			}
			if math.Abs(b.StartX-exp.startX) > 1e-4 || math.Abs(b.StartY-exp.startY) > 1e-4 ||
				math.Abs(b.EndX-exp.endX) > 1e-4 || math.Abs(b.EndY-exp.endY) > 1e-4 {
				valid = false
				break
			}
		}
		for i, exp := range expectedEnd {
			b := stack.ShiftedLeftStackGrowthCurveEndArcShapes[i]
			if b == nil || b.Name != exp.name {
				valid = false
				break
			}
			if math.Abs(b.StartX-exp.startX) > 1e-4 || math.Abs(b.StartY-exp.startY) > 1e-4 ||
				math.Abs(b.EndX-exp.endX) > 1e-4 || math.Abs(b.EndY-exp.endY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		stack.ShiftedLeftStackGrowthCurveStartArcShapes = make([]*ShiftedLeftStackGrowthCurveStartArcShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(ShiftedLeftStackGrowthCurveStartArcShape).Stage(stage)
			b.Name = exp.name
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			b.RadiusX = exp.radiusX
			b.RadiusY = exp.radiusY
			b.XAxisRotation = exp.xAxisRotation
			b.LargeArcFlag = exp.largeArcFlag
			b.SweepFlag = exp.sweepFlag
			stack.ShiftedLeftStackGrowthCurveStartArcShapes[i] = b
		}
		stack.ShiftedLeftStackGrowthCurveEndArcShapes = make([]*ShiftedLeftStackGrowthCurveEndArcShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(ShiftedLeftStackGrowthCurveEndArcShape).Stage(stage)
			b.Name = exp.name
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			b.RadiusX = exp.radiusX
			b.RadiusY = exp.radiusY
			b.XAxisRotation = exp.xAxisRotation
			b.LargeArcFlag = exp.largeArcFlag
			b.SweepFlag = exp.sweepFlag
			stack.ShiftedLeftStackGrowthCurveEndArcShapes[i] = b
		}
		needCommit = true
	}

	return needCommit
}
