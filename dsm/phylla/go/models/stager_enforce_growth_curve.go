package models

import (
	"fmt"
	"math"
)

func enforceStackOfGrowthCurveV2HasShapes(stage *Stage, stack *StackOfRotatedGrowthCurve2D, startGrid *StartHalfwayArcShapeGrid, endGrid *EndHalfwayArcShapeGrid, pGrid *PerpendicularVectorGrid, vector *GrowthVectorShape, stackHeight int, circLen float64, thickness float64) (needCommit bool) {
	if stack == nil || startGrid == nil || endGrid == nil || pGrid == nil || vector == nil || stackHeight < 1 || circLen <= 0 || len(pGrid.PerpendicularVectors) < 2 {
		if len(stack.StackRotatedGrowthCurve2DStartArcShapes) > 0 || len(stack.StackRotatedGrowthCurve2DEndArcShapes) > 0 {
			stack.StackRotatedGrowthCurve2DStartArcShapes = nil
			stack.StackRotatedGrowthCurve2DEndArcShapes = nil
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
	if len(stack.StackRotatedGrowthCurve2DStartArcShapes) != len(expectedStart) || len(stack.StackRotatedGrowthCurve2DEndArcShapes) != len(expectedEnd) {
		valid = false
	} else {
		for i, exp := range expectedStart {
			b := stack.StackRotatedGrowthCurve2DStartArcShapes[i]
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
			b := stack.StackRotatedGrowthCurve2DEndArcShapes[i]
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
		stack.StackRotatedGrowthCurve2DStartArcShapes = make([]*StackRotatedGrowthCurve2DStartArcShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(StackRotatedGrowthCurve2DStartArcShape).Stage(stage)
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
			stack.StackRotatedGrowthCurve2DStartArcShapes[i] = b
		}
		stack.StackRotatedGrowthCurve2DEndArcShapes = make([]*StackRotatedGrowthCurve2DEndArcShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(StackRotatedGrowthCurve2DEndArcShape).Stage(stage)
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
			stack.StackRotatedGrowthCurve2DEndArcShapes[i] = b
		}
		needCommit = true
	}

	return needCommit
}

func enforceTopStackOfGrowthCurveV2HasShapes(stage *Stage, stack *TopStackOfRotatedGrowthCurve2D, startGrid *TopStartHalfwayArcShapeGrid, endGrid *TopEndHalfwayArcShapeGrid, pGrid *PerpendicularVectorGrid, vector *GrowthVectorShape, stackHeight int, circLen float64, thickness float64) (needCommit bool) {
	if stack == nil || startGrid == nil || endGrid == nil || pGrid == nil || vector == nil || stackHeight < 1 || circLen <= 0 || len(pGrid.PerpendicularVectors) < 2 {
		if len(stack.TopStackOfRotatedGrowthCurve2DStartArcShapes) > 0 || len(stack.TopStackOfRotatedGrowthCurve2DEndArcShapes) > 0 {
			stack.TopStackOfRotatedGrowthCurve2DStartArcShapes = nil
			stack.TopStackOfRotatedGrowthCurve2DEndArcShapes = nil
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
	if len(stack.TopStackOfRotatedGrowthCurve2DStartArcShapes) != len(expectedStart) || len(stack.TopStackOfRotatedGrowthCurve2DEndArcShapes) != len(expectedEnd) {
		valid = false
	} else {
		for i, exp := range expectedStart {
			b := stack.TopStackOfRotatedGrowthCurve2DStartArcShapes[i]
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
			b := stack.TopStackOfRotatedGrowthCurve2DEndArcShapes[i]
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
		stack.TopStackOfRotatedGrowthCurve2DStartArcShapes = make([]*TopStackOfRotatedGrowthCurve2DStartArcShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(TopStackOfRotatedGrowthCurve2DStartArcShape).Stage(stage)
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
			stack.TopStackOfRotatedGrowthCurve2DStartArcShapes[i] = b
		}
		stack.TopStackOfRotatedGrowthCurve2DEndArcShapes = make([]*TopStackOfRotatedGrowthCurve2DEndArcShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(TopStackOfRotatedGrowthCurve2DEndArcShape).Stage(stage)
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
			stack.TopStackOfRotatedGrowthCurve2DEndArcShapes[i] = b
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

func enforceStackOfGrowthCurve2DHasShapes(stage *Stage, stack *StackOfGrowthCurve2D, startGrid *StartHalfwayArcShapeGrid, endGrid *EndHalfwayArcShapeGrid, pGrid *PerpendicularVectorGrid, stackHeight int, circLen float64, verticalThickness float64) (needCommit bool) {
	if stack == nil || startGrid == nil || endGrid == nil || pGrid == nil || stackHeight < 1 || circLen <= 0 || len(pGrid.PerpendicularVectors) < 2 {
		if len(stack.StackGrowthCurve2DStartHalfwayArcShapes) > 0 || len(stack.StackGrowthCurve2DEndHalfwayArcShapes) > 0 {
			stack.StackGrowthCurve2DStartHalfwayArcShapes = nil
			stack.StackGrowthCurve2DEndHalfwayArcShapes = nil
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

	for h := 0; h < stackHeight; h++ {
		dy := float64(h) * verticalThickness

		for i := 0; i < expectedLen; i++ {
			sStart := startGrid.StartHalfwayArcShapes[i]
			expectedStart = append(expectedStart, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-start-%d-%d", stack.Name, h, i),
				startX: sStart.StartX, startY: sStart.StartY + dy,
				endX: sStart.EndX, endY: sStart.EndY + dy,
				radiusX: sStart.RadiusX, radiusY: sStart.RadiusY,
				xAxisRotation: sStart.XAxisRotation, largeArcFlag: sStart.LargeArcFlag, sweepFlag: sStart.SweepFlag,
			})

			sEnd := endGrid.EndHalfwayArcShapes[i]
			expectedEnd = append(expectedEnd, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-end-%d-%d", stack.Name, h, i),
				startX: sEnd.StartX, startY: sEnd.StartY + dy,
				endX: sEnd.EndX, endY: sEnd.EndY + dy,
				radiusX: sEnd.RadiusX, radiusY: sEnd.RadiusY,
				xAxisRotation: sEnd.XAxisRotation, largeArcFlag: sEnd.LargeArcFlag, sweepFlag: sEnd.SweepFlag,
			})
		}
	}

	valid := true
	if len(stack.StackGrowthCurve2DStartHalfwayArcShapes) != len(expectedStart) || len(stack.StackGrowthCurve2DEndHalfwayArcShapes) != len(expectedEnd) {
		valid = false
	} else {
		for i, exp := range expectedStart {
			b := stack.StackGrowthCurve2DStartHalfwayArcShapes[i]
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
			b := stack.StackGrowthCurve2DEndHalfwayArcShapes[i]
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
		stack.StackGrowthCurve2DStartHalfwayArcShapes = make([]*StackGrowthCurve2DStartHalfwayArcShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(StackGrowthCurve2DStartHalfwayArcShape).Stage(stage)
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
			stack.StackGrowthCurve2DStartHalfwayArcShapes[i] = b
		}
		stack.StackGrowthCurve2DEndHalfwayArcShapes = make([]*StackGrowthCurve2DEndHalfwayArcShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(StackGrowthCurve2DEndHalfwayArcShape).Stage(stage)
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
			stack.StackGrowthCurve2DEndHalfwayArcShapes[i] = b
		}
		needCommit = true
	}

	return needCommit
}

func enforceTopStackOfGrowthCurve2DHasShapes(stage *Stage, stack *TopStackOfGrowthCurve2D, startGrid *TopStartHalfwayArcShapeGrid, endGrid *TopEndHalfwayArcShapeGrid, pGrid *PerpendicularVectorGrid, stackHeight int, circLen float64, verticalThickness float64) (needCommit bool) {
	if stack == nil || startGrid == nil || endGrid == nil || pGrid == nil || stackHeight < 1 || circLen <= 0 || len(pGrid.PerpendicularVectors) < 2 {
		if len(stack.TopStackGrowthCurve2DStartHalfwayArcShapes) > 0 || len(stack.TopStackGrowthCurve2DEndHalfwayArcShapes) > 0 {
			stack.TopStackGrowthCurve2DStartHalfwayArcShapes = nil
			stack.TopStackGrowthCurve2DEndHalfwayArcShapes = nil
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

	for h := 0; h < stackHeight; h++ {
		dy := float64(h) * verticalThickness

		for i := 0; i < expectedLen; i++ {
			sStart := startGrid.TopStartHalfwayArcShapes[i]
			expectedStart = append(expectedStart, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-start-%d-%d", stack.Name, h, i),
				startX: sStart.StartX, startY: sStart.StartY + dy,
				endX: sStart.EndX, endY: sStart.EndY + dy,
				radiusX: sStart.RadiusX, radiusY: sStart.RadiusY,
				xAxisRotation: sStart.XAxisRotation, largeArcFlag: sStart.LargeArcFlag, sweepFlag: sStart.SweepFlag,
			})

			sEnd := endGrid.TopEndHalfwayArcShapes[i]
			expectedEnd = append(expectedEnd, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-end-%d-%d", stack.Name, h, i),
				startX: sEnd.StartX, startY: sEnd.StartY + dy,
				endX: sEnd.EndX, endY: sEnd.EndY + dy,
				radiusX: sEnd.RadiusX, radiusY: sEnd.RadiusY,
				xAxisRotation: sEnd.XAxisRotation, largeArcFlag: sEnd.LargeArcFlag, sweepFlag: sEnd.SweepFlag,
			})
		}
	}

	valid := true
	if len(stack.TopStackGrowthCurve2DStartHalfwayArcShapes) != len(expectedStart) || len(stack.TopStackGrowthCurve2DEndHalfwayArcShapes) != len(expectedEnd) {
		valid = false
	} else {
		for i, exp := range expectedStart {
			b := stack.TopStackGrowthCurve2DStartHalfwayArcShapes[i]
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
			b := stack.TopStackGrowthCurve2DEndHalfwayArcShapes[i]
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
		stack.TopStackGrowthCurve2DStartHalfwayArcShapes = make([]*TopStackGrowthCurve2DStartHalfwayArcShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(TopStackGrowthCurve2DStartHalfwayArcShape).Stage(stage)
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
			stack.TopStackGrowthCurve2DStartHalfwayArcShapes[i] = b
		}
		stack.TopStackGrowthCurve2DEndHalfwayArcShapes = make([]*TopStackGrowthCurve2DEndHalfwayArcShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(TopStackGrowthCurve2DEndHalfwayArcShape).Stage(stage)
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
			stack.TopStackGrowthCurve2DEndHalfwayArcShapes[i] = b
		}
		needCommit = true
	}

	return needCommit
}

func enforceStackOfGrowthCurve2DRibbonHasShapes(
	stage *Stage,
	ribbonStack *StackOfGrowthCurve2DRibbon,
	bottomStack *StackOfGrowthCurve2D,
	topStack *TopStackOfGrowthCurve2D,
) (needCommit bool) {
	if bottomStack == nil || topStack == nil {
		return false
	}

	type expectedStartShape struct {
		name                         string
		bottomStartX, bottomStartY   float64
		bottomEndX, bottomEndY       float64
		bottomRadiusX, bottomRadiusY float64
		bottomXAxisRotation          float64
		bottomLargeArcFlag           bool
		bottomSweepFlag              bool

		topStartX, topStartY   float64
		topEndX, topEndY       float64
		topRadiusX, topRadiusY float64
		topXAxisRotation       float64
		topLargeArcFlag        bool
		topSweepFlag           bool
	}

	var expectedStart []expectedStartShape
	var expectedEnd []expectedStartShape

	if len(bottomStack.StackGrowthCurve2DStartHalfwayArcShapes) == len(topStack.TopStackGrowthCurve2DStartHalfwayArcShapes) &&
		len(bottomStack.StackGrowthCurve2DEndHalfwayArcShapes) == len(topStack.TopStackGrowthCurve2DEndHalfwayArcShapes) {

		for i := range bottomStack.StackGrowthCurve2DStartHalfwayArcShapes {
			b := bottomStack.StackGrowthCurve2DStartHalfwayArcShapes[i]
			t := topStack.TopStackGrowthCurve2DStartHalfwayArcShapes[i]
			expectedStart = append(expectedStart, expectedStartShape{
				name:         fmt.Sprintf("%s-layer-start-%d", ribbonStack.Name, i),
				bottomStartX: b.StartX, bottomStartY: b.StartY,
				bottomEndX: b.EndX, bottomEndY: b.EndY,
				bottomRadiusX: b.RadiusX, bottomRadiusY: b.RadiusY,
				bottomXAxisRotation: b.XAxisRotation, bottomLargeArcFlag: b.LargeArcFlag, bottomSweepFlag: b.SweepFlag,

				topStartX: t.StartX, topStartY: t.StartY,
				topEndX: t.EndX, topEndY: t.EndY,
				topRadiusX: t.RadiusX, topRadiusY: t.RadiusY,
				topXAxisRotation: t.XAxisRotation, topLargeArcFlag: t.LargeArcFlag, topSweepFlag: t.SweepFlag,
			})
		}

		for i := range bottomStack.StackGrowthCurve2DEndHalfwayArcShapes {
			b := bottomStack.StackGrowthCurve2DEndHalfwayArcShapes[i]
			t := topStack.TopStackGrowthCurve2DEndHalfwayArcShapes[i]
			expectedEnd = append(expectedEnd, expectedStartShape{
				name:         fmt.Sprintf("%s-layer-end-%d", ribbonStack.Name, i),
				bottomStartX: b.StartX, bottomStartY: b.StartY,
				bottomEndX: b.EndX, bottomEndY: b.EndY,
				bottomRadiusX: b.RadiusX, bottomRadiusY: b.RadiusY,
				bottomXAxisRotation: b.XAxisRotation, bottomLargeArcFlag: b.LargeArcFlag, bottomSweepFlag: b.SweepFlag,

				topStartX: t.StartX, topStartY: t.StartY,
				topEndX: t.EndX, topEndY: t.EndY,
				topRadiusX: t.RadiusX, topRadiusY: t.RadiusY,
				topXAxisRotation: t.XAxisRotation, topLargeArcFlag: t.LargeArcFlag, topSweepFlag: t.SweepFlag,
			})
		}
	}

	if len(ribbonStack.StackGrowthCurve2DRibbonStartShapes) != len(expectedStart) {
		for _, b := range ribbonStack.StackGrowthCurve2DRibbonStartShapes {
			b.Unstage(stage)
		}
		ribbonStack.StackGrowthCurve2DRibbonStartShapes = make([]*StackGrowthCurve2DRibbonStartShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(StackGrowthCurve2DRibbonStartShape).Stage(stage)
			b.Name = exp.name
			b.BottomStartX = exp.bottomStartX
			b.BottomStartY = exp.bottomStartY
			b.BottomEndX = exp.bottomEndX
			b.BottomEndY = exp.bottomEndY
			b.BottomRadiusX = exp.bottomRadiusX
			b.BottomRadiusY = exp.bottomRadiusY
			b.BottomXAxisRotation = exp.bottomXAxisRotation
			b.BottomLargeArcFlag = exp.bottomLargeArcFlag
			b.BottomSweepFlag = exp.bottomSweepFlag

			b.TopStartX = exp.topStartX
			b.TopStartY = exp.topStartY
			b.TopEndX = exp.topEndX
			b.TopEndY = exp.topEndY
			b.TopRadiusX = exp.topRadiusX
			b.TopRadiusY = exp.topRadiusY
			b.TopXAxisRotation = exp.topXAxisRotation
			b.TopLargeArcFlag = exp.topLargeArcFlag
			b.TopSweepFlag = exp.topSweepFlag

			ribbonStack.StackGrowthCurve2DRibbonStartShapes[i] = b
		}
		needCommit = true
	}

	if len(ribbonStack.StackGrowthCurve2DRibbonEndShapes) != len(expectedEnd) {
		for _, b := range ribbonStack.StackGrowthCurve2DRibbonEndShapes {
			b.Unstage(stage)
		}
		ribbonStack.StackGrowthCurve2DRibbonEndShapes = make([]*StackGrowthCurve2DRibbonEndShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(StackGrowthCurve2DRibbonEndShape).Stage(stage)
			b.Name = exp.name
			b.BottomStartX = exp.bottomStartX
			b.BottomStartY = exp.bottomStartY
			b.BottomEndX = exp.bottomEndX
			b.BottomEndY = exp.bottomEndY
			b.BottomRadiusX = exp.bottomRadiusX
			b.BottomRadiusY = exp.bottomRadiusY
			b.BottomXAxisRotation = exp.bottomXAxisRotation
			b.BottomLargeArcFlag = exp.bottomLargeArcFlag
			b.BottomSweepFlag = exp.bottomSweepFlag

			b.TopStartX = exp.topStartX
			b.TopStartY = exp.topStartY
			b.TopEndX = exp.topEndX
			b.TopEndY = exp.topEndY
			b.TopRadiusX = exp.topRadiusX
			b.TopRadiusY = exp.topRadiusY
			b.TopXAxisRotation = exp.topXAxisRotation
			b.TopLargeArcFlag = exp.topLargeArcFlag
			b.TopSweepFlag = exp.topSweepFlag

			ribbonStack.StackGrowthCurve2DRibbonEndShapes[i] = b
		}
		needCommit = true
	}

	return needCommit
}

func enforceStackOfRotatedGrowthCurve2DRibbonHasShapes(
	stage *Stage,
	ribbonStack *StackOfRotatedGrowthCurve2DRibbon,
	bottomStack *StackOfRotatedGrowthCurve2D,
	topStack *TopStackOfRotatedGrowthCurve2D,
) (needCommit bool) {
	if bottomStack == nil || topStack == nil {
		return false
	}

	type expectedStartShape struct {
		name                         string
		bottomStartX, bottomStartY   float64
		bottomEndX, bottomEndY       float64
		bottomRadiusX, bottomRadiusY float64
		bottomXAxisRotation          float64
		bottomLargeArcFlag           bool
		bottomSweepFlag              bool

		topStartX, topStartY   float64
		topEndX, topEndY       float64
		topRadiusX, topRadiusY float64
		topXAxisRotation       float64
		topLargeArcFlag        bool
		topSweepFlag           bool
	}

	var expectedStart []expectedStartShape
	var expectedEnd []expectedStartShape

	if len(bottomStack.StackRotatedGrowthCurve2DStartArcShapes) == len(topStack.TopStackOfRotatedGrowthCurve2DStartArcShapes) &&
		len(bottomStack.StackRotatedGrowthCurve2DEndArcShapes) == len(topStack.TopStackOfRotatedGrowthCurve2DEndArcShapes) {

		for i := range bottomStack.StackRotatedGrowthCurve2DStartArcShapes {
			b := bottomStack.StackRotatedGrowthCurve2DStartArcShapes[i]
			t := topStack.TopStackOfRotatedGrowthCurve2DStartArcShapes[i]
			expectedStart = append(expectedStart, expectedStartShape{
				name:         fmt.Sprintf("%s-layer-start-%d", ribbonStack.Name, i),
				bottomStartX: b.StartX, bottomStartY: b.StartY,
				bottomEndX: b.EndX, bottomEndY: b.EndY,
				bottomRadiusX: b.RadiusX, bottomRadiusY: b.RadiusY,
				bottomXAxisRotation: b.XAxisRotation, bottomLargeArcFlag: b.LargeArcFlag, bottomSweepFlag: b.SweepFlag,

				topStartX: t.StartX, topStartY: t.StartY,
				topEndX: t.EndX, topEndY: t.EndY,
				topRadiusX: t.RadiusX, topRadiusY: t.RadiusY,
				topXAxisRotation: t.XAxisRotation, topLargeArcFlag: t.LargeArcFlag, topSweepFlag: t.SweepFlag,
			})
		}

		for i := range bottomStack.StackRotatedGrowthCurve2DEndArcShapes {
			b := bottomStack.StackRotatedGrowthCurve2DEndArcShapes[i]
			t := topStack.TopStackOfRotatedGrowthCurve2DEndArcShapes[i]
			expectedEnd = append(expectedEnd, expectedStartShape{
				name:         fmt.Sprintf("%s-layer-end-%d", ribbonStack.Name, i),
				bottomStartX: b.StartX, bottomStartY: b.StartY,
				bottomEndX: b.EndX, bottomEndY: b.EndY,
				bottomRadiusX: b.RadiusX, bottomRadiusY: b.RadiusY,
				bottomXAxisRotation: b.XAxisRotation, bottomLargeArcFlag: b.LargeArcFlag, bottomSweepFlag: b.SweepFlag,

				topStartX: t.StartX, topStartY: t.StartY,
				topEndX: t.EndX, topEndY: t.EndY,
				topRadiusX: t.RadiusX, topRadiusY: t.RadiusY,
				topXAxisRotation: t.XAxisRotation, topLargeArcFlag: t.LargeArcFlag, topSweepFlag: t.SweepFlag,
			})
		}
	}

	if len(ribbonStack.StackRotatedGrowthCurve2DRibbonStartShapes) != len(expectedStart) {
		for _, b := range ribbonStack.StackRotatedGrowthCurve2DRibbonStartShapes {
			b.Unstage(stage)
		}
		ribbonStack.StackRotatedGrowthCurve2DRibbonStartShapes = make([]*StackRotatedGrowthCurve2DRibbonStartShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(StackRotatedGrowthCurve2DRibbonStartShape).Stage(stage)
			b.Name = exp.name
			b.BottomStartX = exp.bottomStartX
			b.BottomStartY = exp.bottomStartY
			b.BottomEndX = exp.bottomEndX
			b.BottomEndY = exp.bottomEndY
			b.BottomRadiusX = exp.bottomRadiusX
			b.BottomRadiusY = exp.bottomRadiusY
			b.BottomXAxisRotation = exp.bottomXAxisRotation
			b.BottomLargeArcFlag = exp.bottomLargeArcFlag
			b.BottomSweepFlag = exp.bottomSweepFlag

			b.TopStartX = exp.topStartX
			b.TopStartY = exp.topStartY
			b.TopEndX = exp.topEndX
			b.TopEndY = exp.topEndY
			b.TopRadiusX = exp.topRadiusX
			b.TopRadiusY = exp.topRadiusY
			b.TopXAxisRotation = exp.topXAxisRotation
			b.TopLargeArcFlag = exp.topLargeArcFlag
			b.TopSweepFlag = exp.topSweepFlag

			ribbonStack.StackRotatedGrowthCurve2DRibbonStartShapes[i] = b
		}
		needCommit = true
	}

	if len(ribbonStack.StackRotatedGrowthCurve2DRibbonEndShapes) != len(expectedEnd) {
		for _, b := range ribbonStack.StackRotatedGrowthCurve2DRibbonEndShapes {
			b.Unstage(stage)
		}
		ribbonStack.StackRotatedGrowthCurve2DRibbonEndShapes = make([]*StackRotatedGrowthCurve2DRibbonEndShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(StackRotatedGrowthCurve2DRibbonEndShape).Stage(stage)
			b.Name = exp.name
			b.BottomStartX = exp.bottomStartX
			b.BottomStartY = exp.bottomStartY
			b.BottomEndX = exp.bottomEndX
			b.BottomEndY = exp.bottomEndY
			b.BottomRadiusX = exp.bottomRadiusX
			b.BottomRadiusY = exp.bottomRadiusY
			b.BottomXAxisRotation = exp.bottomXAxisRotation
			b.BottomLargeArcFlag = exp.bottomLargeArcFlag
			b.BottomSweepFlag = exp.bottomSweepFlag

			b.TopStartX = exp.topStartX
			b.TopStartY = exp.topStartY
			b.TopEndX = exp.topEndX
			b.TopEndY = exp.topEndY
			b.TopRadiusX = exp.topRadiusX
			b.TopRadiusY = exp.topRadiusY
			b.TopXAxisRotation = exp.topXAxisRotation
			b.TopLargeArcFlag = exp.topLargeArcFlag
			b.TopSweepFlag = exp.topSweepFlag

			ribbonStack.StackRotatedGrowthCurve2DRibbonEndShapes[i] = b
		}
		needCommit = true
	}

	return needCommit
}

func enforceStackOfPartiallyRotatedGrowthCurve2DRibbonHasShapes(
	stage *Stage,
	partiallyRotatedRibbonStack *StackOfPartiallyRotatedGrowthCurve2DRibbon,
	growthRibbonStack *StackOfGrowthCurve2DRibbon,
	rotatedRibbonStack *StackOfRotatedGrowthCurve2DRibbon,
	rotationRatio float64,
) (needCommit bool) {
	if growthRibbonStack == nil || rotatedRibbonStack == nil {
		return false
	}

	type expectedStartShape struct {
		name                         string
		bottomStartX, bottomStartY   float64
		bottomEndX, bottomEndY       float64
		bottomRadiusX, bottomRadiusY float64
		bottomXAxisRotation          float64
		bottomLargeArcFlag           bool
		bottomSweepFlag              bool

		topStartX, topStartY   float64
		topEndX, topEndY       float64
		topRadiusX, topRadiusY float64
		topXAxisRotation       float64
		topLargeArcFlag        bool
		topSweepFlag           bool
	}

	var expectedStart []expectedStartShape
	var expectedEnd []expectedStartShape

	if len(growthRibbonStack.StackGrowthCurve2DRibbonStartShapes) == len(rotatedRibbonStack.StackRotatedGrowthCurve2DRibbonStartShapes) &&
		len(growthRibbonStack.StackGrowthCurve2DRibbonEndShapes) == len(rotatedRibbonStack.StackRotatedGrowthCurve2DRibbonEndShapes) {

		for i := range growthRibbonStack.StackGrowthCurve2DRibbonStartShapes {
			g := growthRibbonStack.StackGrowthCurve2DRibbonStartShapes[i]
			r := rotatedRibbonStack.StackRotatedGrowthCurve2DRibbonStartShapes[i]

			expectedStart = append(expectedStart, expectedStartShape{
				name:         fmt.Sprintf("%s-layer-start-%d", partiallyRotatedRibbonStack.Name, i),
				bottomStartX: g.BottomStartX + rotationRatio*(r.BottomStartX-g.BottomStartX),
				bottomStartY: g.BottomStartY + rotationRatio*(r.BottomStartY-g.BottomStartY),
				bottomEndX:   g.BottomEndX + rotationRatio*(r.BottomEndX-g.BottomEndX),
				bottomEndY:   g.BottomEndY + rotationRatio*(r.BottomEndY-g.BottomEndY),
				bottomRadiusX: g.BottomRadiusX + rotationRatio*(r.BottomRadiusX-g.BottomRadiusX),
				bottomRadiusY: g.BottomRadiusY + rotationRatio*(r.BottomRadiusY-g.BottomRadiusY),
				bottomXAxisRotation: g.BottomXAxisRotation + rotationRatio*(r.BottomXAxisRotation-g.BottomXAxisRotation),
				bottomLargeArcFlag:  g.BottomLargeArcFlag,
				bottomSweepFlag:     g.BottomSweepFlag,

				topStartX: g.TopStartX + rotationRatio*(r.TopStartX-g.TopStartX),
				topStartY: g.TopStartY + rotationRatio*(r.TopStartY-g.TopStartY),
				topEndX:   g.TopEndX + rotationRatio*(r.TopEndX-g.TopEndX),
				topEndY:   g.TopEndY + rotationRatio*(r.TopEndY-g.TopEndY),
				topRadiusX: g.TopRadiusX + rotationRatio*(r.TopRadiusX-g.TopRadiusX),
				topRadiusY: g.TopRadiusY + rotationRatio*(r.TopRadiusY-g.TopRadiusY),
				topXAxisRotation: g.TopXAxisRotation + rotationRatio*(r.TopXAxisRotation-g.TopXAxisRotation),
				topLargeArcFlag:  g.TopLargeArcFlag,
				topSweepFlag:     g.TopSweepFlag,
			})
		}

		for i := range growthRibbonStack.StackGrowthCurve2DRibbonEndShapes {
			g := growthRibbonStack.StackGrowthCurve2DRibbonEndShapes[i]
			r := rotatedRibbonStack.StackRotatedGrowthCurve2DRibbonEndShapes[i]

			expectedEnd = append(expectedEnd, expectedStartShape{
				name:         fmt.Sprintf("%s-layer-end-%d", partiallyRotatedRibbonStack.Name, i),
				bottomStartX: g.BottomStartX + rotationRatio*(r.BottomStartX-g.BottomStartX),
				bottomStartY: g.BottomStartY + rotationRatio*(r.BottomStartY-g.BottomStartY),
				bottomEndX:   g.BottomEndX + rotationRatio*(r.BottomEndX-g.BottomEndX),
				bottomEndY:   g.BottomEndY + rotationRatio*(r.BottomEndY-g.BottomEndY),
				bottomRadiusX: g.BottomRadiusX + rotationRatio*(r.BottomRadiusX-g.BottomRadiusX),
				bottomRadiusY: g.BottomRadiusY + rotationRatio*(r.BottomRadiusY-g.BottomRadiusY),
				bottomXAxisRotation: g.BottomXAxisRotation + rotationRatio*(r.BottomXAxisRotation-g.BottomXAxisRotation),
				bottomLargeArcFlag:  g.BottomLargeArcFlag,
				bottomSweepFlag:     g.BottomSweepFlag,

				topStartX: g.TopStartX + rotationRatio*(r.TopStartX-g.TopStartX),
				topStartY: g.TopStartY + rotationRatio*(r.TopStartY-g.TopStartY),
				topEndX:   g.TopEndX + rotationRatio*(r.TopEndX-g.TopEndX),
				topEndY:   g.TopEndY + rotationRatio*(r.TopEndY-g.TopEndY),
				topRadiusX: g.TopRadiusX + rotationRatio*(r.TopRadiusX-g.TopRadiusX),
				topRadiusY: g.TopRadiusY + rotationRatio*(r.TopRadiusY-g.TopRadiusY),
				topXAxisRotation: g.TopXAxisRotation + rotationRatio*(r.TopXAxisRotation-g.TopXAxisRotation),
				topLargeArcFlag:  g.TopLargeArcFlag,
				topSweepFlag:     g.TopSweepFlag,
			})
		}
	}

	if len(partiallyRotatedRibbonStack.StackPartiallyRotatedGrowthCurve2DRibbonStartShapes) != len(expectedStart) {
		for _, b := range partiallyRotatedRibbonStack.StackPartiallyRotatedGrowthCurve2DRibbonStartShapes {
			b.Unstage(stage)
		}
		partiallyRotatedRibbonStack.StackPartiallyRotatedGrowthCurve2DRibbonStartShapes = make([]*StackPartiallyRotatedGrowthCurve2DRibbonStartShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(StackPartiallyRotatedGrowthCurve2DRibbonStartShape).Stage(stage)
			b.Name = exp.name
			b.BottomStartX = exp.bottomStartX
			b.BottomStartY = exp.bottomStartY
			b.BottomEndX = exp.bottomEndX
			b.BottomEndY = exp.bottomEndY
			b.BottomRadiusX = exp.bottomRadiusX
			b.BottomRadiusY = exp.bottomRadiusY
			b.BottomXAxisRotation = exp.bottomXAxisRotation
			b.BottomLargeArcFlag = exp.bottomLargeArcFlag
			b.BottomSweepFlag = exp.bottomSweepFlag

			b.TopStartX = exp.topStartX
			b.TopStartY = exp.topStartY
			b.TopEndX = exp.topEndX
			b.TopEndY = exp.topEndY
			b.TopRadiusX = exp.topRadiusX
			b.TopRadiusY = exp.topRadiusY
			b.TopXAxisRotation = exp.topXAxisRotation
			b.TopLargeArcFlag = exp.topLargeArcFlag
			b.TopSweepFlag = exp.topSweepFlag

			partiallyRotatedRibbonStack.StackPartiallyRotatedGrowthCurve2DRibbonStartShapes[i] = b
		}
		needCommit = true
	} else {
		for i, exp := range expectedStart {
			b := partiallyRotatedRibbonStack.StackPartiallyRotatedGrowthCurve2DRibbonStartShapes[i]
			if b.Name != exp.name ||
				b.BottomStartX != exp.bottomStartX ||
				b.BottomStartY != exp.bottomStartY ||
				b.BottomEndX != exp.bottomEndX ||
				b.BottomEndY != exp.bottomEndY ||
				b.BottomRadiusX != exp.bottomRadiusX ||
				b.BottomRadiusY != exp.bottomRadiusY ||
				b.BottomXAxisRotation != exp.bottomXAxisRotation ||
				b.BottomLargeArcFlag != exp.bottomLargeArcFlag ||
				b.BottomSweepFlag != exp.bottomSweepFlag ||
				b.TopStartX != exp.topStartX ||
				b.TopStartY != exp.topStartY ||
				b.TopEndX != exp.topEndX ||
				b.TopEndY != exp.topEndY ||
				b.TopRadiusX != exp.topRadiusX ||
				b.TopRadiusY != exp.topRadiusY ||
				b.TopXAxisRotation != exp.topXAxisRotation ||
				b.TopLargeArcFlag != exp.topLargeArcFlag ||
				b.TopSweepFlag != exp.topSweepFlag {

				b.Name = exp.name
				b.BottomStartX = exp.bottomStartX
				b.BottomStartY = exp.bottomStartY
				b.BottomEndX = exp.bottomEndX
				b.BottomEndY = exp.bottomEndY
				b.BottomRadiusX = exp.bottomRadiusX
				b.BottomRadiusY = exp.bottomRadiusY
				b.BottomXAxisRotation = exp.bottomXAxisRotation
				b.BottomLargeArcFlag = exp.bottomLargeArcFlag
				b.BottomSweepFlag = exp.bottomSweepFlag

				b.TopStartX = exp.topStartX
				b.TopStartY = exp.topStartY
				b.TopEndX = exp.topEndX
				b.TopEndY = exp.topEndY
				b.TopRadiusX = exp.topRadiusX
				b.TopRadiusY = exp.topRadiusY
				b.TopXAxisRotation = exp.topXAxisRotation
				b.TopLargeArcFlag = exp.topLargeArcFlag
				b.TopSweepFlag = exp.topSweepFlag
				needCommit = true
			}
		}
	}

	if len(partiallyRotatedRibbonStack.StackPartiallyRotatedGrowthCurve2DRibbonEndShapes) != len(expectedEnd) {
		for _, b := range partiallyRotatedRibbonStack.StackPartiallyRotatedGrowthCurve2DRibbonEndShapes {
			b.Unstage(stage)
		}
		partiallyRotatedRibbonStack.StackPartiallyRotatedGrowthCurve2DRibbonEndShapes = make([]*StackPartiallyRotatedGrowthCurve2DRibbonEndShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(StackPartiallyRotatedGrowthCurve2DRibbonEndShape).Stage(stage)
			b.Name = exp.name
			b.BottomStartX = exp.bottomStartX
			b.BottomStartY = exp.bottomStartY
			b.BottomEndX = exp.bottomEndX
			b.BottomEndY = exp.bottomEndY
			b.BottomRadiusX = exp.bottomRadiusX
			b.BottomRadiusY = exp.bottomRadiusY
			b.BottomXAxisRotation = exp.bottomXAxisRotation
			b.BottomLargeArcFlag = exp.bottomLargeArcFlag
			b.BottomSweepFlag = exp.bottomSweepFlag

			b.TopStartX = exp.topStartX
			b.TopStartY = exp.topStartY
			b.TopEndX = exp.topEndX
			b.TopEndY = exp.topEndY
			b.TopRadiusX = exp.topRadiusX
			b.TopRadiusY = exp.topRadiusY
			b.TopXAxisRotation = exp.topXAxisRotation
			b.TopLargeArcFlag = exp.topLargeArcFlag
			b.TopSweepFlag = exp.topSweepFlag

			partiallyRotatedRibbonStack.StackPartiallyRotatedGrowthCurve2DRibbonEndShapes[i] = b
		}
		needCommit = true
	} else {
		for i, exp := range expectedEnd {
			b := partiallyRotatedRibbonStack.StackPartiallyRotatedGrowthCurve2DRibbonEndShapes[i]
			if b.Name != exp.name ||
				b.BottomStartX != exp.bottomStartX ||
				b.BottomStartY != exp.bottomStartY ||
				b.BottomEndX != exp.bottomEndX ||
				b.BottomEndY != exp.bottomEndY ||
				b.BottomRadiusX != exp.bottomRadiusX ||
				b.BottomRadiusY != exp.bottomRadiusY ||
				b.BottomXAxisRotation != exp.bottomXAxisRotation ||
				b.BottomLargeArcFlag != exp.bottomLargeArcFlag ||
				b.BottomSweepFlag != exp.bottomSweepFlag ||
				b.TopStartX != exp.topStartX ||
				b.TopStartY != exp.topStartY ||
				b.TopEndX != exp.topEndX ||
				b.TopEndY != exp.topEndY ||
				b.TopRadiusX != exp.topRadiusX ||
				b.TopRadiusY != exp.topRadiusY ||
				b.TopXAxisRotation != exp.topXAxisRotation ||
				b.TopLargeArcFlag != exp.topLargeArcFlag ||
				b.TopSweepFlag != exp.topSweepFlag {

				b.Name = exp.name
				b.BottomStartX = exp.bottomStartX
				b.BottomStartY = exp.bottomStartY
				b.BottomEndX = exp.bottomEndX
				b.BottomEndY = exp.bottomEndY
				b.BottomRadiusX = exp.bottomRadiusX
				b.BottomRadiusY = exp.bottomRadiusY
				b.BottomXAxisRotation = exp.bottomXAxisRotation
				b.BottomLargeArcFlag = exp.bottomLargeArcFlag
				b.BottomSweepFlag = exp.bottomSweepFlag

				b.TopStartX = exp.topStartX
				b.TopStartY = exp.topStartY
				b.TopEndX = exp.topEndX
				b.TopEndY = exp.topEndY
				b.TopRadiusX = exp.topRadiusX
				b.TopRadiusY = exp.topRadiusY
				b.TopXAxisRotation = exp.topXAxisRotation
				b.TopLargeArcFlag = exp.topLargeArcFlag
				b.TopSweepFlag = exp.topSweepFlag
				needCommit = true
			}
		}
	}

	return needCommit
}
