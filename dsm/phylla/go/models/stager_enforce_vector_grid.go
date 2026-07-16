package models

import (
	"fmt"
	"math"
)

func enforcePerpendicularVectorGridHasVectors(stage *Stage, grid *PerpendicularVectorGrid, sourceGrid *GrowthCurveRhombusGridShape, v1x, v1y, v2x, v2y, rotAngleRad float64) (needCommit bool) {
	// The four vertices relative to the center before rotation
	p1x, p1y := -(v1x+v2x)/2.0, -(v1y+v2y)/2.0
	p2x, p2y := (v1x-v2x)/2.0, (v1y-v2y)/2.0
	p3x, p3y := (v2x-v1x)/2.0, (v2y-v1y)/2.0
	p4x, p4y := (v1x+v2x)/2.0, (v1y+v2y)/2.0

	cosA := math.Cos(rotAngleRad)
	sinA := math.Sin(rotAngleRad)

	pts := []struct{ x, y float64 }{
		{p1x*cosA - p1y*sinA, p1x*sinA + p1y*cosA},
		{p2x*cosA - p2y*sinA, p2x*sinA + p2y*cosA},
		{p3x*cosA - p3y*sinA, p3x*sinA + p3y*cosA},
		{p4x*cosA - p4y*sinA, p4x*sinA + p4y*cosA},
	}

	minY := pts[0].y
	bottomPt := pts[0]
	for i := 1; i < len(pts); i++ {
		if pts[i].y < minY {
			minY = pts[i].y
			bottomPt = pts[i]
		}
	}

	dx := bottomPt.x
	dy := bottomPt.y

	valid := true
	if len(grid.PerpendicularVectors) != len(sourceGrid.GrowthCurveRhombusShapes) {
		valid = false
	} else {
		for i, r := range sourceGrid.GrowthCurveRhombusShapes {
			vec := grid.PerpendicularVectors[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if vec == nil || vec.Name != expectedName {
				valid = false
				break
			}
			expectedStartX := r.X + dx
			expectedStartY := r.Y + dy
			if math.Abs(vec.StartX-expectedStartX) > 1e-4 || math.Abs(vec.StartY-expectedStartY) > 1e-4 ||
				math.Abs(vec.EndX-r.X) > 1e-4 || math.Abs(vec.EndY-r.Y) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		grid.PerpendicularVectors = make([]*PerpendicularVector, len(sourceGrid.GrowthCurveRhombusShapes))
		for i, r := range sourceGrid.GrowthCurveRhombusShapes {
			vec := new(PerpendicularVector).Stage(stage)
			vec.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			vec.StartX = r.X + dx
			vec.StartY = r.Y + dy
			vec.EndX = r.X
			vec.EndY = r.Y
			grid.PerpendicularVectors[i] = vec
		}
		needCommit = true
	} else {
		for i, r := range sourceGrid.GrowthCurveRhombusShapes {
			vec := grid.PerpendicularVectors[i]
			expectedStartX := r.X + dx
			expectedStartY := r.Y + dy
			if vec.StartX != expectedStartX || vec.StartY != expectedStartY || vec.EndX != r.X || vec.EndY != r.Y {
				vec.StartX = expectedStartX
				vec.StartY = expectedStartY
				vec.EndX = r.X
				vec.EndY = r.Y
				needCommit = true
			}
		}
	}

	return
}

func enforceBaseVectorShapeGridHasShapes(stage *Stage, grid *BaseVectorShapeGrid, pGrid *PerpendicularVectorGrid) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.BaseVectorShapes) > 0 {
			grid.BaseVectorShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.BaseVectorShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			base := grid.BaseVectorShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if base == nil || base.Name != expectedName {
				valid = false
				break
			}

			expectedStartX := v1.StartX
			expectedStartY := v1.StartY
			expectedEndX := v2.StartX
			expectedEndY := v2.StartY

			if math.Abs(base.StartX-expectedStartX) > 1e-4 || math.Abs(base.StartY-expectedStartY) > 1e-4 ||
				math.Abs(base.EndX-expectedEndX) > 1e-4 || math.Abs(base.EndY-expectedEndY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.BaseVectorShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.BaseVectorShapes = make([]*BaseVectorShape, expectedLen)
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			expectedStartX := v1.StartX
			expectedStartY := v1.StartY
			expectedEndX := v2.StartX
			expectedEndY := v2.StartY

			newBase := new(BaseVectorShape).Stage(stage)
			newBase.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newBase.StartX = expectedStartX
			newBase.StartY = expectedStartY
			newBase.EndX = expectedEndX
			newBase.EndY = expectedEndY

			grid.BaseVectorShapes[i] = newBase
		}
		needCommit = true
	}
	return needCommit
}

func enforceArcNormalVectorShapeGridHasShapes(stage *Stage, grid *ArcNormalVectorShapeGrid, pGrid *PerpendicularVectorGrid) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.ArcNormalVectorShapes) > 0 {
			grid.ArcNormalVectorShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.ArcNormalVectorShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arcNorm := grid.ArcNormalVectorShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if arcNorm == nil || arcNorm.Name != expectedName {
				valid = false
				break
			}

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
			origSweepFlag := (model_cross < 0)

			radX := midX - cx
			radY := midY - cy
			var tx, ty float64
			if origSweepFlag {
				tx = -radY
				ty = radX
			} else {
				tx = radY
				ty = -radX
			}

			tLen := math.Hypot(tx, ty)
			if tLen != 0 {
				tx /= tLen
				ty /= tLen
			}

			dirX := ty
			dirY := -tx

			vLen := length

			expectedStartX := midX
			expectedStartY := midY
			expectedEndX := midX + dirX*vLen
			expectedEndY := midY + dirY*vLen

			if math.Abs(arcNorm.StartX-expectedStartX) > 1e-4 || math.Abs(arcNorm.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arcNorm.EndX-expectedEndX) > 1e-4 || math.Abs(arcNorm.EndY-expectedEndY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.ArcNormalVectorShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.ArcNormalVectorShapes = make([]*ArcNormalVectorShape, expectedLen)
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
			origSweepFlag := (model_cross < 0)

			radX := midX - cx
			radY := midY - cy
			var tx, ty float64
			if origSweepFlag {
				tx = -radY
				ty = radX
			} else {
				tx = radY
				ty = -radX
			}

			tLen := math.Hypot(tx, ty)
			if tLen != 0 {
				tx /= tLen
				ty /= tLen
			}

			dirX := ty
			dirY := -tx

			vLen := length

			expectedStartX := midX
			expectedStartY := midY
			expectedEndX := midX + dirX*vLen
			expectedEndY := midY + dirY*vLen

			newVec := new(ArcNormalVectorShape).Stage(stage)
			newVec.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newVec.StartX = expectedStartX
			newVec.StartY = expectedStartY
			newVec.EndX = expectedEndX
			newVec.EndY = expectedEndY

			grid.ArcNormalVectorShapes[i] = newVec
		}
		needCommit = true
	}
	return needCommit
}

func enforcePerpendicularVectorGridHalfwayHasVectors(stage *Stage, grid *PerpendicularVectorGridHalfway, sourceGrid *PerpendicularVectorGrid) (needCommit bool) {
	if sourceGrid == nil || grid == nil || len(sourceGrid.PerpendicularVectors) < 2 {
		if len(grid.PerpendicularVectorHalfways) > 0 {
			grid.PerpendicularVectorHalfways = nil
			return true
		}
		return false
	}

	expectedLen := len(sourceGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.PerpendicularVectorHalfways) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := sourceGrid.PerpendicularVectors[i]
			v2 := sourceGrid.PerpendicularVectors[i+1]
			vHalfway := grid.PerpendicularVectorHalfways[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if vHalfway == nil || vHalfway.Name != expectedName {
				valid = false
				break
			}
			expectedStartX := (v1.StartX + v2.StartX) / 2.0
			expectedStartY := (v1.StartY + v2.StartY) / 2.0

			// Translation of v1
			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY

			expectedEndX := expectedStartX + dx
			expectedEndY := expectedStartY + dy

			if math.Abs(vHalfway.StartX-expectedStartX) > 1e-4 || math.Abs(vHalfway.StartY-expectedStartY) > 1e-4 ||
				math.Abs(vHalfway.EndX-expectedEndX) > 1e-4 || math.Abs(vHalfway.EndY-expectedEndY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, v := range grid.PerpendicularVectorHalfways {
			if v != nil {
				v.Unstage(stage)
			}
		}
		grid.PerpendicularVectorHalfways = make([]*PerpendicularVectorHalfway, expectedLen)
		for i := 0; i < expectedLen; i++ {
			v1 := sourceGrid.PerpendicularVectors[i]
			v2 := sourceGrid.PerpendicularVectors[i+1]

			expectedStartX := (v1.StartX + v2.StartX) / 2.0
			expectedStartY := (v1.StartY + v2.StartY) / 2.0

			// Translation of v1
			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY

			expectedEndX := expectedStartX + dx
			expectedEndY := expectedStartY + dy

			newVec := new(PerpendicularVectorHalfway).Stage(stage)
			newVec.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newVec.StartX = expectedStartX
			newVec.StartY = expectedStartY
			newVec.EndX = expectedEndX
			newVec.EndY = expectedEndY
			grid.PerpendicularVectorHalfways[i] = newVec
		}
		needCommit = true
	}
	return needCommit
}

func enforceShiftedLeftStackOfNormalVectorHasShapes(stage *Stage, stack *ShiftedLeftStackOfNormalVector, nGrid *ArcNormalVectorShapeGrid, pGrid *PerpendicularVectorGrid, vector *GrowthVectorShape, stackHeight int, circLen float64, thickness float64) (needCommit bool) {
	if stack == nil || nGrid == nil || pGrid == nil || vector == nil || stackHeight < 1 || circLen <= 0 || len(nGrid.ArcNormalVectorShapes) == 0 || len(pGrid.PerpendicularVectors) == 0 {
		if len(stack.ShiftedLeftStackNormalVectors) > 0 {
			stack.ShiftedLeftStackNormalVectors = nil
			return true
		}
		return false
	}

	expectedLen := len(nGrid.ArcNormalVectorShapes)

	type expectedShape struct {
		name                       string
		startX, startY, endX, endY float64
	}
	var expected []expectedShape

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
			vec := nGrid.ArcNormalVectorShapes[i]

			currentDX := math.Mod(dx, circLen)
			if currentDX < 0 {
				currentDX += circLen
			}

			expected = append(expected, expectedShape{
				name:   fmt.Sprintf("%s-layer-%d-%d", stack.Name, h, i),
				startX: vec.StartX + currentDX - circLen, startY: vec.StartY + dy,
				endX: vec.EndX + currentDX - circLen, endY: vec.EndY + dy,
			})
		}
	}

	valid := true
	if len(stack.ShiftedLeftStackNormalVectors) != len(expected) {
		valid = false
	} else {
		for i, exp := range expected {
			b := stack.ShiftedLeftStackNormalVectors[i]
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
		stack.ShiftedLeftStackNormalVectors = make([]*ShiftedLeftStackNormalVector, len(expected))
		for i, exp := range expected {
			b := new(ShiftedLeftStackNormalVector).Stage(stage)
			b.Name = exp.name
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			stack.ShiftedLeftStackNormalVectors[i] = b
		}
		needCommit = true
	}

	return needCommit
}
