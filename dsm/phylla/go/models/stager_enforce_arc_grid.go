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

			sweepFlag := (model_cross < 0)

			scale := 1.0 - thickness/R_val
			R_new := R * math.Abs(scale)

			expectedStartX := cx + scale*(v1.StartX-cx)
			expectedStartY := cy + scale*(v1.StartY-cy)
			expectedEndX := cx + scale*(midX-cx)
			expectedEndY := cy + scale*(midY-cy)

			expectedRadiusX := R_new
			expectedRadiusY := R_new

			if math.Abs(arc.StartX-expectedStartX) > 1e-4 || math.Abs(arc.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arc.EndX-expectedEndX) > 1e-4 || math.Abs(arc.EndY-expectedEndY) > 1e-4 ||
				math.Abs(arc.RadiusX-expectedRadiusX) > 1e-4 || math.Abs(arc.RadiusY-expectedRadiusY) > 1e-4 ||
				arc.SweepFlag != sweepFlag {
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
		vFirst := pGrid.PerpendicularVectors[0]
		vx := vFirst.EndX - vFirst.StartX
		vy := vFirst.EndY - vFirst.StartY
		vLen := math.Hypot(vx, vy)
		if vLen == 0 {
			vLen = 1
		}
		vx, vy = vx/vLen, vy/vLen
		dx_thick := thickness * vx
		dy_thick := thickness * vy

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

			sweepFlag := (model_cross < 0)

			expectedStartX := v1.StartX + dx_thick
			expectedStartY := v1.StartY + dy_thick
			expectedEndX := midX + dx_thick
			expectedEndY := midY + dy_thick

			expectedRadiusX := R
			expectedRadiusY := R

			newArc := new(TopStartArcShape).Stage(stage)
			newArc.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newArc.StartX = expectedStartX
			newArc.StartY = expectedStartY
			newArc.EndX = expectedEndX
			newArc.EndY = expectedEndY
			newArc.RadiusX = expectedRadiusX
			newArc.RadiusY = expectedRadiusY
			newArc.SweepFlag = sweepFlag
			newArc.LargeArcFlag = false
			newArc.XAxisRotation = 0

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
		vFirst := pGrid.PerpendicularVectors[0]
		vx := vFirst.EndX - vFirst.StartX
		vy := vFirst.EndY - vFirst.StartY
		vLen := math.Hypot(vx, vy)
		if vLen == 0 {
			vLen = 1
		}
		vx, vy = vx/vLen, vy/vLen
		dx_thick := thickness * vx
		dy_thick := thickness * vy

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arc := grid.TopEndArcShapes[i]
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

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
			origSweepFlag := (model_cross < 0)

			expectedStartX := 2*midX - v1.StartX + dx_thick
			expectedStartY := 2*midY - v1.StartY + dy_thick
			expectedEndX := midX + dx_thick
			expectedEndY := midY + dy_thick

			expectedRadiusX := R
			expectedRadiusY := R

			if math.Abs(arc.StartX-expectedStartX) > 1e-4 || math.Abs(arc.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arc.EndX-expectedEndX) > 1e-4 || math.Abs(arc.EndY-expectedEndY) > 1e-4 ||
				math.Abs(arc.RadiusX-expectedRadiusX) > 1e-4 || math.Abs(arc.RadiusY-expectedRadiusY) > 1e-4 ||
				arc.SweepFlag != origSweepFlag {
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
		vFirst := pGrid.PerpendicularVectors[0]
		vx := vFirst.EndX - vFirst.StartX
		vy := vFirst.EndY - vFirst.StartY
		vLen := math.Hypot(vx, vy)
		if vLen == 0 {
			vLen = 1
		}
		vx, vy = vx/vLen, vy/vLen
		dx_thick := thickness * vx
		dy_thick := thickness * vy

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

			expectedStartX := 2*midX - v1.StartX + dx_thick
			expectedStartY := 2*midY - v1.StartY + dy_thick
			expectedEndX := midX + dx_thick
			expectedEndY := midY + dy_thick

			expectedRadiusX := R
			expectedRadiusY := R

			newArc := new(TopEndArcShape).Stage(stage)
			newArc.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newArc.StartX = expectedStartX
			newArc.StartY = expectedStartY
			newArc.EndX = expectedEndX
			newArc.EndY = expectedEndY
			newArc.RadiusX = expectedRadiusX
			newArc.RadiusY = expectedRadiusY
			newArc.SweepFlag = origSweepFlag
			newArc.LargeArcFlag = false
			newArc.XAxisRotation = 0

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
