package models

import (
	"fmt"
	"math"
)

func enforceStartArcShapeV2GridHasShapes(stage *Stage, grid *StartArcShapeGrid, pGrid *PerpendicularVectorGrid) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.StartArcShapes) > 0 {
			grid.StartArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.StartArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arc := grid.StartArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if arc == nil || arc.Name != expectedName {
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

			expectedStartX := v1.StartX
			expectedStartY := v1.StartY

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
			R := math.Abs(R_val)
			expectedRadiusX := R
			expectedRadiusY := R

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			expectedEndX := midX
			expectedEndY := midY

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x

			expectedSweepFlag := (model_cross < 0)

			if math.Abs(arc.StartX-expectedStartX) > 1e-4 || math.Abs(arc.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arc.EndX-expectedEndX) > 1e-4 || math.Abs(arc.EndY-expectedEndY) > 1e-4 ||
				math.Abs(arc.RadiusX-expectedRadiusX) > 1e-4 || math.Abs(arc.RadiusY-expectedRadiusY) > 1e-4 ||
				arc.SweepFlag != expectedSweepFlag {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.StartArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.StartArcShapes = make([]*StartArcShape, expectedLen)
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

			expectedStartX := v1.StartX
			expectedStartY := v1.StartY

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
			R := math.Abs(R_val)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			expectedEndX := midX
			expectedEndY := midY

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x

			sweepFlag := (model_cross < 0)

			newArc := new(StartArcShape).Stage(stage)
			newArc.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newArc.StartX = expectedStartX
			newArc.StartY = expectedStartY
			newArc.EndX = expectedEndX
			newArc.EndY = expectedEndY
			newArc.RadiusX = R
			newArc.RadiusY = R
			newArc.SweepFlag = sweepFlag
			newArc.LargeArcFlag = false
			newArc.XAxisRotation = 0

			grid.StartArcShapes[i] = newArc
		}
		needCommit = true
	}
	return needCommit
}

func enforceTopStartArcShapeV2GridHasShapes(stage *Stage, grid *TopStartArcShapeGrid, pGrid *PerpendicularVectorGrid, thickness float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.TopStartArcShapes) > 0 {
			grid.TopStartArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.TopStartArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arc := grid.TopStartArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if arc == nil || arc.Name != expectedName {
				valid = false
				break
			}

			sX, sY, eX, eY, rX, rY, _, _, swp := computeArcV2Geometry(v1, v2, thickness, false)

			expectedStartX := sX
			expectedStartY := sY
			expectedEndX := eX
			expectedEndY := eY
			expectedRadiusX := rX
			expectedRadiusY := rY

			if math.Abs(arc.StartX-expectedStartX) > 1e-4 || math.Abs(arc.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arc.EndX-expectedEndX) > 1e-4 || math.Abs(arc.EndY-expectedEndY) > 1e-4 ||
				math.Abs(arc.RadiusX-expectedRadiusX) > 1e-4 || math.Abs(arc.RadiusY-expectedRadiusY) > 1e-4 ||
				arc.SweepFlag != swp {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.TopStartArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.TopStartArcShapes = make([]*TopStartArcShape, expectedLen)

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp := computeArcV2Geometry(v1, v2, thickness, false)

			newArc := new(TopStartArcShape).Stage(stage)
			newArc.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newArc.StartX = sX
			newArc.StartY = sY
			newArc.EndX = eX
			newArc.EndY = eY
			newArc.RadiusX = rX
			newArc.RadiusY = rY
			newArc.SweepFlag = swp
			newArc.LargeArcFlag = lArc
			newArc.XAxisRotation = xRot

			grid.TopStartArcShapes[i] = newArc
		}
		needCommit = true
	}
	return needCommit
}

func enforceEndArcShapeV2GridHasShapes(stage *Stage, grid *EndArcShapeGrid, pGrid *PerpendicularVectorGrid) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.EndArcShapes) > 0 {
			grid.EndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.EndArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arc := grid.EndArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if arc == nil || arc.Name != expectedName {
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
			R := math.Abs(R_val)
			expectedRadiusX := R
			expectedRadiusY := R

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
			origSweepFlag := (model_cross < 0)

			expectedStartX := 2*midX - v1.StartX
			expectedStartY := 2*midY - v1.StartY
			expectedEndX := midX
			expectedEndY := midY
			expectedSweepFlag := origSweepFlag

			if math.Abs(arc.StartX-expectedStartX) > 1e-4 || math.Abs(arc.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arc.EndX-expectedEndX) > 1e-4 || math.Abs(arc.EndY-expectedEndY) > 1e-4 ||
				math.Abs(arc.RadiusX-expectedRadiusX) > 1e-4 || math.Abs(arc.RadiusY-expectedRadiusY) > 1e-4 ||
				arc.SweepFlag != expectedSweepFlag {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.EndArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.EndArcShapes = make([]*EndArcShape, expectedLen)
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
			R := math.Abs(R_val)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
			origSweepFlag := (model_cross < 0)

			expectedStartX := 2*midX - v1.StartX
			expectedStartY := 2*midY - v1.StartY
			expectedEndX := midX
			expectedEndY := midY

			newArc := new(EndArcShape).Stage(stage)
			newArc.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newArc.StartX = expectedStartX
			newArc.StartY = expectedStartY
			newArc.EndX = expectedEndX
			newArc.EndY = expectedEndY
			newArc.RadiusX = R
			newArc.RadiusY = R
			newArc.SweepFlag = origSweepFlag
			newArc.LargeArcFlag = false
			newArc.XAxisRotation = 0

			grid.EndArcShapes[i] = newArc
		}
		needCommit = true
	}
	return needCommit
}

func enforceTopEndArcShapeV2GridHasShapes(stage *Stage, grid *TopEndArcShapeGrid, pGrid *PerpendicularVectorGrid, thickness float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.TopEndArcShapes) > 0 {
			grid.TopEndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.TopEndArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arc := grid.TopEndArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if arc == nil || arc.Name != expectedName {
				valid = false
				break
			}

			sX, sY, eX, eY, rX, rY, _, _, swp := computeArcV2Geometry(v1, v2, thickness, true)

			expectedStartX := sX
			expectedStartY := sY
			expectedEndX := eX
			expectedEndY := eY
			expectedRadiusX := rX
			expectedRadiusY := rY

			if math.Abs(arc.StartX-expectedStartX) > 1e-4 || math.Abs(arc.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arc.EndX-expectedEndX) > 1e-4 || math.Abs(arc.EndY-expectedEndY) > 1e-4 ||
				math.Abs(arc.RadiusX-expectedRadiusX) > 1e-4 || math.Abs(arc.RadiusY-expectedRadiusY) > 1e-4 ||
				arc.SweepFlag != swp {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.TopEndArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.TopEndArcShapes = make([]*TopEndArcShape, expectedLen)

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp := computeArcV2Geometry(v1, v2, thickness, true)

			newArc := new(TopEndArcShape).Stage(stage)
			newArc.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newArc.StartX = sX
			newArc.StartY = sY
			newArc.EndX = eX
			newArc.EndY = eY
			newArc.RadiusX = rX
			newArc.RadiusY = rY
			newArc.SweepFlag = swp
			newArc.LargeArcFlag = lArc
			newArc.XAxisRotation = xRot

			grid.TopEndArcShapes[i] = newArc
		}
		needCommit = true
	}
	return needCommit
}

func computeArcV2Geometry(v1, v2 *PerpendicularVector, offset float64, isEndArc bool) (expectedStartX, expectedStartY, expectedEndX, expectedEndY, expectedRadiusX, expectedRadiusY, xAxisRotation float64, largeArcFlag, sweepFlag bool) {
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
	R := math.Abs(R_val)

	cx := v1.StartX + R_val*ux
	cy := v1.StartY + R_val*uy

	AB_model_x := midX - v1.StartX
	AB_model_y := midY - v1.StartY
	AC_model_x := cx - v1.StartX
	AC_model_y := cy - v1.StartY
	model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
	baseSweepFlag := (model_cross < 0)

	if !isEndArc {
		scale := 1.0 - offset/R_val
		R_new := R * math.Abs(scale)

		expectedStartX = cx + scale*(v1.StartX-cx)
		expectedStartY = cy + scale*(v1.StartY-cy)
		expectedEndX = cx + scale*(midX-cx)
		expectedEndY = cy + scale*(midY-cy)

		expectedRadiusX = R_new
		expectedRadiusY = R_new
		xAxisRotation = 0.0
		largeArcFlag = false
		sweepFlag = baseSweepFlag
	} else {
		cx_new := 2*midX - cx
		cy_new := 2*midY - cy

		scale := 1.0 + offset/R_val
		R_new := R * math.Abs(scale)

		expectedStartX = cx_new - scale*(v1.StartX-cx)
		expectedStartY = cy_new - scale*(v1.StartY-cy)
		expectedEndX = cx_new + scale*(midX-cx_new)
		expectedEndY = cy_new + scale*(midY-cy_new)

		expectedRadiusX = R_new
		expectedRadiusY = R_new
		xAxisRotation = 0.0
		largeArcFlag = false
		sweepFlag = baseSweepFlag
	}

	return
}

func enforceShiftedBottomTopStartArcShapeV2GridHasShapes(stage *Stage, grid *ShiftedBottomTopStartArcShapeGrid, pGrid *PerpendicularVectorGrid, thickness float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.ShiftedBottomTopStartArcShapes) > 0 {
			grid.ShiftedBottomTopStartArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true

	vFirst := pGrid.PerpendicularVectors[0]
	vx := vFirst.EndX - vFirst.StartX
	vy := vFirst.EndY - vFirst.StartY
	vLen := math.Hypot(vx, vy)
	if vLen == 0 {
		vLen = 1
	}
	vx, vy = vx/vLen, vy/vLen

	dx := -thickness * vx
	dy := -thickness * vy

	if len(grid.ShiftedBottomTopStartArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arc := grid.ShiftedBottomTopStartArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if arc == nil || arc.Name != expectedName {
				valid = false
				break
			}

			sX, sY, eX, eY, rX, rY, _, _, swp := computeArcV2Geometry(v1, v2, thickness, false)

			expectedStartX := sX + dx
			expectedStartY := sY + dy
			expectedEndX := eX + dx
			expectedEndY := eY + dy
			expectedRadiusX := rX
			expectedRadiusY := rY

			if math.Abs(arc.StartX-expectedStartX) > 1e-4 || math.Abs(arc.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arc.EndX-expectedEndX) > 1e-4 || math.Abs(arc.EndY-expectedEndY) > 1e-4 ||
				math.Abs(arc.RadiusX-expectedRadiusX) > 1e-4 || math.Abs(arc.RadiusY-expectedRadiusY) > 1e-4 ||
				arc.SweepFlag != swp {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.ShiftedBottomTopStartArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.ShiftedBottomTopStartArcShapes = make([]*ShiftedBottomTopStartArcShape, expectedLen)

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp := computeArcV2Geometry(v1, v2, thickness, false)

			newArc := new(ShiftedBottomTopStartArcShape).Stage(stage)
			newArc.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newArc.StartX = sX + dx
			newArc.StartY = sY + dy
			newArc.EndX = eX + dx
			newArc.EndY = eY + dy
			newArc.RadiusX = rX
			newArc.RadiusY = rY
			newArc.SweepFlag = swp
			newArc.LargeArcFlag = lArc
			newArc.XAxisRotation = xRot

			grid.ShiftedBottomTopStartArcShapes[i] = newArc
		}
		needCommit = true
	}
	return needCommit
}

func enforceMidArcVectorShapeGridHasShapes(stage *Stage, grid *MidArcVectorShapeGrid, pGrid *PerpendicularVectorGrid, thickness float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.MidArcVectorShapes) > 0 {
			grid.MidArcVectorShapes = nil
			return true
		}
		return false
	}

	vFirst := pGrid.PerpendicularVectors[0]
	vx := vFirst.EndX - vFirst.StartX
	vy := vFirst.EndY - vFirst.StartY
	vLen := math.Hypot(vx, vy)
	if vLen == 0 {
		vLen = 1
	}
	vx, vy = vx/vLen, vy/vLen

	dx := -thickness * vx
	dy := -thickness * vy

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.MidArcVectorShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			s := grid.MidArcVectorShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if s == nil || s.Name != expectedName {
				valid = false
				break
			}

			_, _, eX, eY, _, _, _, _, _ := computeArcV2Geometry(v1, v2, thickness, false)
			eX = eX + dx
			eY = eY + dy
			_, _, eX2, eY2, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, true)

			if math.Abs(s.StartX-eX) > 1e-4 || math.Abs(s.StartY-eY) > 1e-4 ||
				math.Abs(s.EndX-eX2) > 1e-4 || math.Abs(s.EndY-eY2) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.MidArcVectorShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.MidArcVectorShapes = make([]*MidArcVectorShape, expectedLen)

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			_, _, eX, eY, _, _, _, _, _ := computeArcV2Geometry(v1, v2, thickness, false)
			eX = eX + dx
			eY = eY + dy
			_, _, eX2, eY2, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, true)

			newShape := new(MidArcVectorShape).Stage(stage)
			newShape.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newShape.StartX = eX
			newShape.StartY = eY
			newShape.EndX = eX2
			newShape.EndY = eY2

			grid.MidArcVectorShapes[i] = newShape
		}
		needCommit = true
	}
	return needCommit
}

func enforceTopMidArcVectorShapeGridHasShapes(stage *Stage, grid *TopMidArcVectorShapeGrid, pGrid *PerpendicularVectorGrid, thickness float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.TopMidArcVectorShapes) > 0 {
			grid.TopMidArcVectorShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	vFirst := pGrid.PerpendicularVectors[0]
	vx := vFirst.EndX - vFirst.StartX
	vy := vFirst.EndY - vFirst.StartY
	vLen := math.Hypot(vx, vy)
	if vLen == 0 {
		vLen = 1
	}
	vx, vy = vx/vLen, vy/vLen

	dx := -thickness * vx
	dy := -thickness * vy

	if len(grid.TopMidArcVectorShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			s := grid.TopMidArcVectorShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if s == nil || s.Name != expectedName {
				valid = false
				break
			}

			// TopMidArcVectorShape is a pure translation of MidArcVectorShape
			// by 'thickness' along the perpendicular vector v2.
			_, _, mX, mY, _, _, _, _, _ := computeArcV2Geometry(v1, v2, thickness, false)
			mX = mX + dx
			mY = mY + dy
			_, _, mX2, mY2, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, true)

			vx2 := v2.EndX - v2.StartX
			vy2 := v2.EndY - v2.StartY
			v2Len := math.Hypot(vx2, vy2)
			if v2Len == 0 {
				v2Len = 1
			}
			vx2, vy2 = vx2/v2Len, vy2/v2Len

			eX := mX + thickness*vx2
			eY := mY + thickness*vy2
			eX2 := mX2 + thickness*vx2
			eY2 := mY2 + thickness*vy2

			if math.Abs(s.StartX-eX) > 1e-4 || math.Abs(s.StartY-eY) > 1e-4 ||
				math.Abs(s.EndX-eX2) > 1e-4 || math.Abs(s.EndY-eY2) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.TopMidArcVectorShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.TopMidArcVectorShapes = make([]*TopMidArcVectorShape, expectedLen)

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			// TopMidArcVectorShape is a pure translation of MidArcVectorShape
			// by 'thickness' along the perpendicular vector v2.
			_, _, mX, mY, _, _, _, _, _ := computeArcV2Geometry(v1, v2, thickness, false)
			mX = mX + dx
			mY = mY + dy
			_, _, mX2, mY2, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, true)

			vx2 := v2.EndX - v2.StartX
			vy2 := v2.EndY - v2.StartY
			v2Len := math.Hypot(vx2, vy2)
			if v2Len == 0 {
				v2Len = 1
			}
			vx2, vy2 = vx2/v2Len, vy2/v2Len

			eX := mX + thickness*vx2
			eY := mY + thickness*vy2
			eX2 := mX2 + thickness*vx2
			eY2 := mY2 + thickness*vy2

			newShape := new(TopMidArcVectorShape).Stage(stage)
			newShape.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newShape.StartX = eX
			newShape.StartY = eY
			newShape.EndX = eX2
			newShape.EndY = eY2

			grid.TopMidArcVectorShapes[i] = newShape
		}
		needCommit = true
	}
	return needCommit
}

func enforceHalfwayArcShapeGridHasShapes(stage *Stage, grid *StartHalfwayArcShapeGrid, pGrid *PerpendicularVectorGrid, thickness float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if grid != nil && len(grid.StartHalfwayArcShapes) > 0 {
			grid.StartHalfwayArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	vFirst := pGrid.PerpendicularVectors[0]
	vx := vFirst.EndX - vFirst.StartX
	vy := vFirst.EndY - vFirst.StartY
	vLen := math.Hypot(vx, vy)
	if vLen == 0 {
		vLen = 1
	}
	vx, vy = vx/vLen, vy/vLen

	dx := -thickness * vx
	dy := -thickness * vy

	if len(grid.StartHalfwayArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			s := grid.StartHalfwayArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if s == nil || s.Name != expectedName {
				valid = false
				break
			}

			sX, sY, _, _, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, false)
			_, _, mX, mY, _, _, _, _, _ := computeArcV2Geometry(v1, v2, thickness, false)
			mX = mX + dx
			mY = mY + dy
			_, _, mX2, mY2, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, true)

			endX := (mX + mX2) / 2.0
			endY := (mY + mY2) / 2.0

			v1_dx := v1.EndX - v1.StartX
			v1_dy := v1.EndY - v1.StartY
			length := math.Hypot(v1_dx, v1_dy)
			if length == 0 {
				length = 1
			}
			ux, uy := v1_dx/length, v1_dy/length

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

			nx := cx - sX
			ny := cy - sY
			nLen := math.Hypot(nx, ny)
			if nLen == 0 {
				nLen = 1
			}
			nx, ny = nx/nLen, ny/nLen

			vX := sX - endX
			vY := sY - endY
			distSq := vX*vX + vY*vY
			dot := vX*nx + vY*ny
			t := 0.0
			if dot != 0 {
				t = -distSq / (2.0 * dot)
			}
			R_new := math.Abs(t)

			cx_new := sX + t*nx
			cy_new := sY + t*ny

			AB_x := endX - sX
			AB_y := endY - sY
			AC_x := cx_new - sX
			AC_y := cy_new - sY
			new_cross := AB_x*AC_y - AB_y*AC_x
			sweepFlag := (new_cross < 0)

			if math.Abs(s.StartX-sX) > 1e-4 || math.Abs(s.StartY-sY) > 1e-4 ||
				math.Abs(s.EndX-endX) > 1e-4 || math.Abs(s.EndY-endY) > 1e-4 ||
				math.Abs(s.RadiusX-R_new) > 1e-4 || s.SweepFlag != sweepFlag {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.StartHalfwayArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.StartHalfwayArcShapes = make([]*StartHalfwayArcShape, expectedLen)

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			sX, sY, _, _, _, _, xAxisRotation, _, _ := computeArcV2Geometry(v1, v2, 0.0, false)

			_, _, mX, mY, _, _, _, _, _ := computeArcV2Geometry(v1, v2, thickness, false)
			mX = mX + dx
			mY = mY + dy
			_, _, mX2, mY2, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, true)

			endX := (mX + mX2) / 2.0
			endY := (mY + mY2) / 2.0

			v1_dx := v1.EndX - v1.StartX
			v1_dy := v1.EndY - v1.StartY
			length := math.Hypot(v1_dx, v1_dy)
			if length == 0 {
				length = 1
			}
			ux, uy := v1_dx/length, v1_dy/length

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

			nx := cx - sX
			ny := cy - sY
			nLen := math.Hypot(nx, ny)
			if nLen == 0 {
				nLen = 1
			}
			nx, ny = nx/nLen, ny/nLen

			vX := sX - endX
			vY := sY - endY
			distSq := vX*vX + vY*vY
			dot := vX*nx + vY*ny
			t := 0.0
			if dot != 0 {
				t = -distSq / (2.0 * dot)
			}
			R_new := math.Abs(t)

			cx_new := sX + t*nx
			cy_new := sY + t*ny

			AB_x := endX - sX
			AB_y := endY - sY
			AC_x := cx_new - sX
			AC_y := cy_new - sY
			new_cross := AB_x*AC_y - AB_y*AC_x
			sweepFlag := (new_cross < 0)

			newShape := new(StartHalfwayArcShape).Stage(stage)
			newShape.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newShape.StartX = sX
			newShape.StartY = sY
			newShape.EndX = endX
			newShape.EndY = endY
			newShape.RadiusX = R_new
			newShape.RadiusY = R_new
			newShape.XAxisRotation = xAxisRotation
			newShape.LargeArcFlag = false
			newShape.SweepFlag = sweepFlag

			grid.StartHalfwayArcShapes[i] = newShape
		}
		needCommit = true
	}
	return needCommit
}

func enforceEndHalfwayArcShapeGridHasShapes(stage *Stage, grid *EndHalfwayArcShapeGrid, pGrid *PerpendicularVectorGrid, thickness float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if grid != nil && len(grid.EndHalfwayArcShapes) > 0 {
			grid.EndHalfwayArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	vFirst := pGrid.PerpendicularVectors[0]
	vx := vFirst.EndX - vFirst.StartX
	vy := vFirst.EndY - vFirst.StartY
	vLen := math.Hypot(vx, vy)
	if vLen == 0 {
		vLen = 1
	}
	vx, vy = vx/vLen, vy/vLen

	dx := -thickness * vx
	dy := -thickness * vy

	if len(grid.EndHalfwayArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			s := grid.EndHalfwayArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if s == nil || s.Name != expectedName {
				valid = false
				break
			}

			// Compute Start Point (A): end point of HalfwayArcShapeGrid
			_, _, mX, mY, _, _, _, _, _ := computeArcV2Geometry(v1, v2, thickness, false)
			mX = mX + dx
			mY = mY + dy
			_, _, mX2, mY2, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, true)
			sX := (mX + mX2) / 2.0
			sY := (mY + mY2) / 2.0

			// Compute End Point (B): start point of EndArcShapeGrid
			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0
			endX := 2*midX - v1.StartX
			endY := 2*midY - v1.StartY

			// Compute normal at End Point (B): pointing to center of EndArcShapeGrid
			v1_dx := v1.EndX - v1.StartX
			v1_dy := v1.EndY - v1.StartY
			length := math.Hypot(v1_dx, v1_dy)
			if length == 0 {
				length = 1
			}
			ux, uy := v1_dx/length, v1_dy/length

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

			cx_end_arc := 2*midX - cx
			cy_end_arc := 2*midY - cy

			nx := cx_end_arc - endX
			ny := cy_end_arc - endY
			nLen := math.Hypot(nx, ny)
			if nLen == 0 {
				nLen = 1
			}
			nx, ny = nx/nLen, ny/nLen

			vX := endX - sX
			vY := endY - sY
			distSq := vX*vX + vY*vY
			dot := vX*nx + vY*ny
			t := 0.0
			if dot != 0 {
				t = -distSq / (2.0 * dot)
			}
			R_new := math.Abs(t)

			cx_new := endX + t*nx
			cy_new := endY + t*ny

			AB_x := endX - sX
			AB_y := endY - sY
			AC_x := cx_new - sX
			AC_y := cy_new - sY
			new_cross := AB_x*AC_y - AB_y*AC_x
			sweepFlag := (new_cross < 0)

			if math.Abs(s.StartX-sX) > 1e-4 || math.Abs(s.StartY-sY) > 1e-4 ||
				math.Abs(s.EndX-endX) > 1e-4 || math.Abs(s.EndY-endY) > 1e-4 ||
				math.Abs(s.RadiusX-R_new) > 1e-4 || s.SweepFlag != sweepFlag {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.EndHalfwayArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.EndHalfwayArcShapes = make([]*EndHalfwayArcShape, expectedLen)

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			_, _, mX, mY, _, _, _, _, _ := computeArcV2Geometry(v1, v2, thickness, false)
			mX = mX + dx
			mY = mY + dy
			_, _, mX2, mY2, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, true)
			sX := (mX + mX2) / 2.0
			sY := (mY + mY2) / 2.0

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0
			endX := 2*midX - v1.StartX
			endY := 2*midY - v1.StartY

			v1_dx := v1.EndX - v1.StartX
			v1_dy := v1.EndY - v1.StartY
			length := math.Hypot(v1_dx, v1_dy)
			if length == 0 {
				length = 1
			}
			ux, uy := v1_dx/length, v1_dy/length

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

			cx_end_arc := 2*midX - cx
			cy_end_arc := 2*midY - cy

			nx := cx_end_arc - endX
			ny := cy_end_arc - endY
			nLen := math.Hypot(nx, ny)
			if nLen == 0 {
				nLen = 1
			}
			nx, ny = nx/nLen, ny/nLen

			vX := endX - sX
			vY := endY - sY
			distSq := vX*vX + vY*vY
			dot := vX*nx + vY*ny
			t := 0.0
			if dot != 0 {
				t = -distSq / (2.0 * dot)
			}
			R_new := math.Abs(t)

			cx_new := endX + t*nx
			cy_new := endY + t*ny

			AB_x := endX - sX
			AB_y := endY - sY
			AC_x := cx_new - sX
			AC_y := cy_new - sY
			new_cross := AB_x*AC_y - AB_y*AC_x
			sweepFlag := (new_cross < 0)

			_, _, _, _, _, _, xAxisRotation, _, _ := computeArcV2Geometry(v1, v2, 0.0, false)

			newShape := new(EndHalfwayArcShape).Stage(stage)
			newShape.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newShape.StartX = sX
			newShape.StartY = sY
			newShape.EndX = endX
			newShape.EndY = endY
			newShape.RadiusX = R_new
			newShape.RadiusY = R_new
			newShape.XAxisRotation = xAxisRotation
			newShape.LargeArcFlag = false
			newShape.SweepFlag = sweepFlag

			grid.EndHalfwayArcShapes[i] = newShape
		}
		needCommit = true
	}
	return needCommit
}

func enforceTopStartHalfwayArcShapeGridHasShapes(stage *Stage, grid *TopStartHalfwayArcShapeGrid, pGrid *PerpendicularVectorGrid, thickness float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if grid != nil && len(grid.TopStartHalfwayArcShapes) > 0 {
			grid.TopStartHalfwayArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	vFirst := pGrid.PerpendicularVectors[0]
	vx := vFirst.EndX - vFirst.StartX
	vy := vFirst.EndY - vFirst.StartY
	vLen := math.Hypot(vx, vy)
	if vLen == 0 {
		vLen = 1
	}
	vx, vy = vx/vLen, vy/vLen

	dx := -thickness * vx
	dy := -thickness * vy

	if len(grid.TopStartHalfwayArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			s := grid.TopStartHalfwayArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if s == nil || s.Name != expectedName {
				valid = false
				break
			}

			sX, sY, _, _, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, false)
			_, _, mX, mY, _, _, _, _, _ := computeArcV2Geometry(v1, v2, thickness, false)
			mX = mX + dx
			mY = mY + dy
			_, _, mX2, mY2, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, true)

			endX := (mX + mX2) / 2.0
			endY := (mY + mY2) / 2.0

			v1_dx := v1.EndX - v1.StartX
			v1_dy := v1.EndY - v1.StartY
			length := math.Hypot(v1_dx, v1_dy)
			if length == 0 {
				length = 1
			}
			ux, uy := v1_dx/length, v1_dy/length

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

			nx := cx - sX
			ny := cy - sY
			nLen := math.Hypot(nx, ny)
			if nLen == 0 {
				nLen = 1
			}
			nx, ny = nx/nLen, ny/nLen

			vX := sX - endX
			vY := sY - endY
			distSq := vX*vX + vY*vY
			dot := vX*nx + vY*ny
			t := 0.0
			if dot != 0 {
				t = -distSq / (2.0 * dot)
			}
			R_new := math.Abs(t)

			cx_new := sX + t*nx
			cy_new := sY + t*ny

			AB_x := endX - sX
			AB_y := endY - sY
			AC_x := cx_new - sX
			AC_y := cy_new - sY
			new_cross := AB_x*AC_y - AB_y*AC_x
			sweepFlag := (new_cross < 0)
			sX, sY = sX+dx, sY+dy
			endX, endY = endX+dx, endY+dy

			if math.Abs(s.StartX-sX) > 1e-4 || math.Abs(s.StartY-sY) > 1e-4 ||
				math.Abs(s.EndX-endX) > 1e-4 || math.Abs(s.EndY-endY) > 1e-4 ||
				math.Abs(s.RadiusX-R_new) > 1e-4 || s.SweepFlag != sweepFlag {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.TopStartHalfwayArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.TopStartHalfwayArcShapes = make([]*TopStartHalfwayArcShape, expectedLen)

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			sX, sY, _, _, _, _, xAxisRotation, _, _ := computeArcV2Geometry(v1, v2, 0.0, false)

			_, _, mX, mY, _, _, _, _, _ := computeArcV2Geometry(v1, v2, thickness, false)
			mX = mX + dx
			mY = mY + dy
			_, _, mX2, mY2, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, true)

			endX := (mX + mX2) / 2.0
			endY := (mY + mY2) / 2.0

			v1_dx := v1.EndX - v1.StartX
			v1_dy := v1.EndY - v1.StartY
			length := math.Hypot(v1_dx, v1_dy)
			if length == 0 {
				length = 1
			}
			ux, uy := v1_dx/length, v1_dy/length

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

			nx := cx - sX
			ny := cy - sY
			nLen := math.Hypot(nx, ny)
			if nLen == 0 {
				nLen = 1
			}
			nx, ny = nx/nLen, ny/nLen

			vX := sX - endX
			vY := sY - endY
			distSq := vX*vX + vY*vY
			dot := vX*nx + vY*ny
			t := 0.0
			if dot != 0 {
				t = -distSq / (2.0 * dot)
			}
			R_new := math.Abs(t)

			cx_new := sX + t*nx
			cy_new := sY + t*ny

			AB_x := endX - sX
			AB_y := endY - sY
			AC_x := cx_new - sX
			AC_y := cy_new - sY
			new_cross := AB_x*AC_y - AB_y*AC_x
			sweepFlag := (new_cross < 0)
			sX, sY = sX+dx, sY+dy
			endX, endY = endX+dx, endY+dy

			newShape := new(TopStartHalfwayArcShape).Stage(stage)
			newShape.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newShape.StartX = sX
			newShape.StartY = sY
			newShape.EndX = endX
			newShape.EndY = endY
			newShape.RadiusX = R_new
			newShape.RadiusY = R_new
			newShape.XAxisRotation = xAxisRotation
			newShape.LargeArcFlag = false
			newShape.SweepFlag = sweepFlag

			grid.TopStartHalfwayArcShapes[i] = newShape
		}
		needCommit = true
	}
	return needCommit
}

func enforceTopEndHalfwayArcShapeGridHasShapes(stage *Stage, grid *TopEndHalfwayArcShapeGrid, pGrid *PerpendicularVectorGrid, thickness float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if grid != nil && len(grid.TopEndHalfwayArcShapes) > 0 {
			grid.TopEndHalfwayArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	vFirst := pGrid.PerpendicularVectors[0]
	vx := vFirst.EndX - vFirst.StartX
	vy := vFirst.EndY - vFirst.StartY
	vLen := math.Hypot(vx, vy)
	if vLen == 0 {
		vLen = 1
	}
	vx, vy = vx/vLen, vy/vLen

	dx := -thickness * vx
	dy := -thickness * vy

	if len(grid.TopEndHalfwayArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			s := grid.TopEndHalfwayArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if s == nil || s.Name != expectedName {
				valid = false
				break
			}

			// Compute Start Point (A): end point of HalfwayArcShapeGrid
			_, _, mX, mY, _, _, _, _, _ := computeArcV2Geometry(v1, v2, thickness, false)
			mX = mX + dx
			mY = mY + dy
			_, _, mX2, mY2, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, true)
			sX := (mX + mX2) / 2.0
			sY := (mY + mY2) / 2.0

			// Compute End Point (B): start point of EndArcShapeGrid
			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0
			endX := 2*midX - v1.StartX
			endY := 2*midY - v1.StartY

			// Compute normal at End Point (B): pointing to center of EndArcShapeGrid
			v1_dx := v1.EndX - v1.StartX
			v1_dy := v1.EndY - v1.StartY
			length := math.Hypot(v1_dx, v1_dy)
			if length == 0 {
				length = 1
			}
			ux, uy := v1_dx/length, v1_dy/length

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

			cx_end_arc := 2*midX - cx
			cy_end_arc := 2*midY - cy

			nx := cx_end_arc - endX
			ny := cy_end_arc - endY
			nLen := math.Hypot(nx, ny)
			if nLen == 0 {
				nLen = 1
			}
			nx, ny = nx/nLen, ny/nLen

			vX := endX - sX
			vY := endY - sY
			distSq := vX*vX + vY*vY
			dot := vX*nx + vY*ny
			t := 0.0
			if dot != 0 {
				t = -distSq / (2.0 * dot)
			}
			R_new := math.Abs(t)

			cx_new := endX + t*nx
			cy_new := endY + t*ny

			AB_x := endX - sX
			AB_y := endY - sY
			AC_x := cx_new - sX
			AC_y := cy_new - sY
			new_cross := AB_x*AC_y - AB_y*AC_x
			sweepFlag := (new_cross < 0)
			sX, sY = sX+dx, sY+dy
			endX, endY = endX+dx, endY+dy

			if math.Abs(s.StartX-sX) > 1e-4 || math.Abs(s.StartY-sY) > 1e-4 ||
				math.Abs(s.EndX-endX) > 1e-4 || math.Abs(s.EndY-endY) > 1e-4 ||
				math.Abs(s.RadiusX-R_new) > 1e-4 || s.SweepFlag != sweepFlag {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.TopEndHalfwayArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.TopEndHalfwayArcShapes = make([]*TopEndHalfwayArcShape, expectedLen)

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			_, _, mX, mY, _, _, _, _, _ := computeArcV2Geometry(v1, v2, thickness, false)
			mX = mX + dx
			mY = mY + dy
			_, _, mX2, mY2, _, _, _, _, _ := computeArcV2Geometry(v1, v2, 0.0, true)
			sX := (mX + mX2) / 2.0
			sY := (mY + mY2) / 2.0

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0
			endX := 2*midX - v1.StartX
			endY := 2*midY - v1.StartY

			v1_dx := v1.EndX - v1.StartX
			v1_dy := v1.EndY - v1.StartY
			length := math.Hypot(v1_dx, v1_dy)
			if length == 0 {
				length = 1
			}
			ux, uy := v1_dx/length, v1_dy/length

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

			cx_end_arc := 2*midX - cx
			cy_end_arc := 2*midY - cy

			nx := cx_end_arc - endX
			ny := cy_end_arc - endY
			nLen := math.Hypot(nx, ny)
			if nLen == 0 {
				nLen = 1
			}
			nx, ny = nx/nLen, ny/nLen

			vX := endX - sX
			vY := endY - sY
			distSq := vX*vX + vY*vY
			dot := vX*nx + vY*ny
			t := 0.0
			if dot != 0 {
				t = -distSq / (2.0 * dot)
			}
			R_new := math.Abs(t)

			cx_new := endX + t*nx
			cy_new := endY + t*ny

			AB_x := endX - sX
			AB_y := endY - sY
			AC_x := cx_new - sX
			AC_y := cy_new - sY
			new_cross := AB_x*AC_y - AB_y*AC_x
			sweepFlag := (new_cross < 0)
			sX, sY = sX+dx, sY+dy
			endX, endY = endX+dx, endY+dy

			_, _, _, _, _, _, xAxisRotation, _, _ := computeArcV2Geometry(v1, v2, 0.0, false)

			newShape := new(TopEndHalfwayArcShape).Stage(stage)
			newShape.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newShape.StartX = sX
			newShape.StartY = sY
			newShape.EndX = endX
			newShape.EndY = endY
			newShape.RadiusX = R_new
			newShape.RadiusY = R_new
			newShape.XAxisRotation = xAxisRotation
			newShape.LargeArcFlag = false
			newShape.SweepFlag = sweepFlag

			grid.TopEndHalfwayArcShapes[i] = newShape
		}
		needCommit = true
	}
	return needCommit
}
