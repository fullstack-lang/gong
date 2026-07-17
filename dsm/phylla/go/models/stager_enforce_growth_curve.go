package models

import (
	"fmt"
	"math"
)

func enforceStackOfGrowthCurveV2HasShapes(stage *Stage, stack *StackOfGrowthCurve, startGrid *StartHalfwayArcShapeGrid, endGrid *EndHalfwayArcShapeGrid, pGrid *PerpendicularVectorGrid, vector *GrowthVectorShape, stackHeight int, circLen float64, thickness float64) (needCommit bool) {
	if stack == nil || startGrid == nil || endGrid == nil || pGrid == nil || vector == nil || stackHeight < 1 || circLen <= 0 || len(pGrid.PerpendicularVectors) < 2 {
		if len(stack.StackGrowthCurveStartArcShapes) > 0 || len(stack.StackGrowthCurveEndArcShapes) > 0 {
			stack.StackGrowthCurveStartArcShapes = nil
			stack.StackGrowthCurveEndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	if len(startGrid.StartHalfwayArcShapes) != expectedLen || len(endGrid.EndHalfwayArcShapes) != expectedLen {
		return false
	}

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

		for i := 0; i < expectedLen; i++ {
			currentDX := math.Mod(dx, circLen)
			if currentDX < 0 {
				currentDX += circLen
			}

			sStart := startGrid.StartHalfwayArcShapes[i]
			expectedStart = append(expectedStart, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-start-%d-%d", stack.Name, h, i),
				startX: sStart.StartX + currentDX, startY: sStart.StartY + dy,
				endX: sStart.EndX + currentDX, endY: sStart.EndY + dy,
				radiusX: sStart.RadiusX, radiusY: sStart.RadiusY,
				xAxisRotation: sStart.XAxisRotation, largeArcFlag: sStart.LargeArcFlag, sweepFlag: sStart.SweepFlag,
			})

			sEnd := endGrid.EndHalfwayArcShapes[i]
			expectedEnd = append(expectedEnd, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-end-%d-%d", stack.Name, h, i),
				startX: sEnd.StartX + currentDX, startY: sEnd.StartY + dy,
				endX: sEnd.EndX + currentDX, endY: sEnd.EndY + dy,
				radiusX: sEnd.RadiusX, radiusY: sEnd.RadiusY,
				xAxisRotation: sEnd.XAxisRotation, largeArcFlag: sEnd.LargeArcFlag, sweepFlag: sEnd.SweepFlag,
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

func enforceTopStackOfGrowthCurveV2HasShapes(stage *Stage, stack *TopStackOfGrowthCurve, startGrid *TopStartHalfwayArcShapeGrid, endGrid *TopEndHalfwayArcShapeGrid, pGrid *PerpendicularVectorGrid, vector *GrowthVectorShape, stackHeight int, circLen float64, thickness float64) (needCommit bool) {
	if stack == nil || startGrid == nil || endGrid == nil || pGrid == nil || vector == nil || stackHeight < 1 || circLen <= 0 || len(pGrid.PerpendicularVectors) < 2 {
		if len(stack.TopStackGrowthCurveStartArcShapes) > 0 || len(stack.TopStackGrowthCurveEndArcShapes) > 0 {
			stack.TopStackGrowthCurveStartArcShapes = nil
			stack.TopStackGrowthCurveEndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	if len(startGrid.TopStartHalfwayArcShapes) != expectedLen || len(endGrid.TopEndHalfwayArcShapes) != expectedLen {
		return false
	}

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

		for i := 0; i < expectedLen; i++ {
			currentDX := math.Mod(dx, circLen)
			if currentDX < 0 {
				currentDX += circLen
			}

			sStart := startGrid.TopStartHalfwayArcShapes[i]
			expectedStart = append(expectedStart, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-start-%d-%d", stack.Name, h, i),
				startX: sStart.StartX + currentDX, startY: sStart.StartY + dy,
				endX: sStart.EndX + currentDX, endY: sStart.EndY + dy,
				radiusX: sStart.RadiusX, radiusY: sStart.RadiusY,
				xAxisRotation: sStart.XAxisRotation, largeArcFlag: sStart.LargeArcFlag, sweepFlag: sStart.SweepFlag,
			})

			sEnd := endGrid.TopEndHalfwayArcShapes[i]
			expectedEnd = append(expectedEnd, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-end-%d-%d", stack.Name, h, i),
				startX: sEnd.StartX + currentDX, startY: sEnd.StartY + dy,
				endX: sEnd.EndX + currentDX, endY: sEnd.EndY + dy,
				radiusX: sEnd.RadiusX, radiusY: sEnd.RadiusY,
				xAxisRotation: sEnd.XAxisRotation, largeArcFlag: sEnd.LargeArcFlag, sweepFlag: sEnd.SweepFlag,
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

func enforceShiftedLeftStackOfGrowthCurveV2HasShapes(stage *Stage, stack *ShiftedLeftStackOfGrowthCurve, startGrid *StartHalfwayArcShapeGrid, endGrid *EndHalfwayArcShapeGrid, pGrid *PerpendicularVectorGrid, vector *GrowthVectorShape, stackHeight int, circLen float64, thickness float64) (needCommit bool) {
	if stack == nil || startGrid == nil || endGrid == nil || pGrid == nil || vector == nil || stackHeight < 1 || circLen <= 0 || len(pGrid.PerpendicularVectors) < 2 {
		if len(stack.ShiftedLeftStackGrowthCurveStartArcShapes) > 0 || len(stack.ShiftedLeftStackGrowthCurveEndArcShapes) > 0 {
			stack.ShiftedLeftStackGrowthCurveStartArcShapes = nil
			stack.ShiftedLeftStackGrowthCurveEndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	if len(startGrid.StartHalfwayArcShapes) != expectedLen || len(endGrid.EndHalfwayArcShapes) != expectedLen {
		return false
	}

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

		for i := 0; i < expectedLen; i++ {
			currentDX := math.Mod(dx, circLen)
			if currentDX < 0 {
				currentDX += circLen
			}

			sStart := startGrid.StartHalfwayArcShapes[i]
			expectedStart = append(expectedStart, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-start-%d-%d", stack.Name, h, i),
				startX: sStart.StartX + currentDX - circLen, startY: sStart.StartY + dy,
				endX: sStart.EndX + currentDX - circLen, endY: sStart.EndY + dy,
				radiusX: sStart.RadiusX, radiusY: sStart.RadiusY,
				xAxisRotation: sStart.XAxisRotation, largeArcFlag: sStart.LargeArcFlag, sweepFlag: sStart.SweepFlag,
			})

			sEnd := endGrid.EndHalfwayArcShapes[i]
			expectedEnd = append(expectedEnd, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-end-%d-%d", stack.Name, h, i),
				startX: sEnd.StartX + currentDX - circLen, startY: sEnd.StartY + dy,
				endX: sEnd.EndX + currentDX - circLen, endY: sEnd.EndY + dy,
				radiusX: sEnd.RadiusX, radiusY: sEnd.RadiusY,
				xAxisRotation: sEnd.XAxisRotation, largeArcFlag: sEnd.LargeArcFlag, sweepFlag: sEnd.SweepFlag,
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
