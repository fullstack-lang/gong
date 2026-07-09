package models

import (
	"fmt"
	"math"
)

func enforceGrowthPathRhombusGridShapeHasRhombuses(stage *Stage, grid *GrowthCurveRhombusGridShape, sourceGrid *RotatedRhombusGridShape, angleDegree float64, sideLength float64, rhombusInsideAngle float64) (needCommit bool) {
	if sourceGrid == nil || grid == nil {
		return false
	}





	type rotatedRhombus struct {
		r        *RotatedRhombusShape
		rotatedX float64
		rotatedY float64
	}

	var candidates []rotatedRhombus
	for _, r := range sourceGrid.RotatedRhombusShapes {
		// r.X and r.Y are now pre-rotated Cartesian centers by the stager!
		cx := r.X
		cy := r.Y

		// Keep rhombuses where the CENTER is on or above the origin line (Cartesian Y >= 0)
		if cy >= -0.0001 {
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
					// In Cartesian space, "below" means a lower Y coordinate.
					if next == nil || cand.rotatedY < next.rotatedY {
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
