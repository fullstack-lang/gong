package models

import (
	"fmt"
	"math"
)

func enforceGrowthPathRhombusGridShapeHasRhombuses(stage *Stage, grid *GrowthCurveRhombusGridShape, sourceGrid *RotatedRhombusGridShape, angleDegree float64, sideLength float64, rhombusInsideAngle float64) (needCommit bool) {
	if sourceGrid == nil || grid == nil {
		return false
	}

	rotRad := angleDegree * math.Pi / 180.0
	cosA := math.Cos(rotRad)
	sinA := math.Sin(rotRad)

	insideAngleRad := rhombusInsideAngle * math.Pi / 180.0

	v1x := sideLength * math.Cos(insideAngleRad/2.0)
	v1y := -sideLength * math.Sin(insideAngleRad/2.0)
	v2x := -sideLength * math.Cos(insideAngleRad/2.0)
	v2y := -sideLength * math.Sin(insideAngleRad/2.0)

	v1_rot_x := v1x*cosA - v1y*sinA
	v1_rot_y := v1x*sinA + v1y*cosA
	v2_rot_x := v2x*cosA - v2y*sinA
	v2_rot_y := v2x*sinA + v2y*cosA

	type rotatedRhombus struct {
		r        *RotatedRhombusShape
		rotatedX float64
		rotatedY float64
	}

	var candidates []rotatedRhombus
	for _, r := range sourceGrid.RotatedRhombusShapes {
		// r.X and r.Y are now pre-rotated by the stager!
		// Calculate the rotated center of the rhombus
		cx := r.X + (v1_rot_x+v2_rot_x)/2.0
		cy := r.Y + (v1_rot_y+v2_rot_y)/2.0

		// Keep rhombuses where the CENTER is visually above or on the origin line (SVG Y <= 0)
		if cy <= 0.0001 {
			candidates = append(candidates, rotatedRhombus{r: r, rotatedX: cx, rotatedY: cy})
		}
	}

	if len(candidates) == 0 {
		if len(grid.GrowthCurveRhombusShapes) > 0 {
			grid.GrowthCurveRhombusShapes = nil
			return true
		}
		return false
	}

	leftmost := candidates[0]
	for _, cand := range candidates {
		if cand.rotatedX < leftmost.rotatedX {
			leftmost = cand
		}
	}

	var ordered []rotatedRhombus
	ordered = append(ordered, leftmost)
	current := leftmost

	for {
		var next *rotatedRhombus
		for i := range candidates {
			cand := &candidates[i]
			if cand.r == current.r {
				continue
			}

			// adjacent on the right
			if cand.rotatedX > current.rotatedX+1e-4 {
				dx := cand.rotatedX - current.rotatedX
				dy := cand.rotatedY - current.rotatedY
				dist := math.Sqrt(dx*dx + dy*dy)
				if math.Abs(dist-sideLength) < 1e-4 {
					// The user wants the one "below" if there are two adjacent to the right.
					// In SVG space, "below" means a higher Y coordinate.
					if next == nil || cand.rotatedY > next.rotatedY {
						next = cand
					}
				}
			}
		}

		if next == nil {
			break
		}

		ordered = append(ordered, *next)
		current = *next
	}

	// Now check if grid.GrowthCurveRhombusShapes matches ordered
	valid := true
	if len(grid.GrowthCurveRhombusShapes) != len(ordered) {
		valid = false
	} else {
		for i, cand := range ordered {
			r := grid.GrowthCurveRhombusShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if r == nil || r.Name != expectedName || math.Abs(r.X-cand.r.X) > 1e-4 || math.Abs(r.Y-cand.r.Y) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		grid.GrowthCurveRhombusShapes = make([]*GrowthCurveRhombusShape, len(ordered))
		for i, cand := range ordered {
			r := new(GrowthCurveRhombusShape).Stage(stage)
			r.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			r.X = cand.r.X
			r.Y = cand.r.Y
			grid.GrowthCurveRhombusShapes[i] = r
		}
		needCommit = true
	}

	return needCommit
}
