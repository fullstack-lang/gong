package models

import (
	"fmt"
	"math"
)

func enforceInitialRhombusGridShapeHasRhombuses(stage *Stage, grid *InitialRhombusGridShape, N, M int, v1x, v1y, v2x, v2y, rotAngleRad float64) (needCommit bool) {
	cosA := math.Cos(rotAngleRad)
	sinA := math.Sin(rotAngleRad)
	valid := true
	if len(grid.InitialRhombusShapes) != (N+1)*M {
		valid = false
	} else {
		seen := make(map[*InitialRhombusShape]bool)
		idx := 0
		for i := -1; i < N; i++ {
			for j := 0; j < M; j++ {
				r := grid.InitialRhombusShapes[idx]
				if seen[r] {
					valid = false
					break
				}
				seen[r] = true
				expectedName := fmt.Sprintf("%s-%d-%d", grid.Name, i, j)
				if r == nil || r.Name != expectedName {
					valid = false
					break
				}
				idx++
			}
			if !valid {
				break
			}
		}
	}

	if !valid {
		grid.InitialRhombusShapes = make([]*InitialRhombusShape, 0, (N+1)*M)
		for i := -1; i < N; i++ {
			for j := 0; j < M; j++ {
				r := new(InitialRhombusShape).Stage(stage)
				r.Name = fmt.Sprintf("%s-%d-%d", grid.Name, i, j)
				origX := float64(i)*v1x + float64(j)*v2x + (v1x+v2x)/2.0
				origY := float64(i)*v1y + float64(j)*v2y + (v1y+v2y)/2.0
				r.X = origX*cosA - origY*sinA
				r.Y = origX*sinA + origY*cosA
				grid.InitialRhombusShapes = append(grid.InitialRhombusShapes, r)
			}
		}
		needCommit = true
	} else {
		idx := 0
		for i := -1; i < N; i++ {
			for j := 0; j < M; j++ {
				r := grid.InitialRhombusShapes[idx]
				origX := float64(i)*v1x + float64(j)*v2x + (v1x+v2x)/2.0
				origY := float64(i)*v1y + float64(j)*v2y + (v1y+v2y)/2.0
				expectedX := origX*cosA - origY*sinA
				expectedY := origX*sinA + origY*cosA
				if r.X != expectedX || r.Y != expectedY {
					r.X = expectedX
					r.Y = expectedY
					needCommit = true
				}
				idx++
			}
		}
	}
	return
}

func enforceRotatedRhombusGridShapeHasRhombuses(stage *Stage, grid *RotatedRhombusGridShape, N, M int, v1x, v1y, v2x, v2y, rotAngleRad float64) (needCommit bool) {
	cosA := math.Cos(rotAngleRad)
	sinA := math.Sin(rotAngleRad)
	valid := true
	if len(grid.RotatedRhombusShapes) != (N+1)*M {
		valid = false
	} else {
		seen := make(map[*RotatedRhombusShape]bool)
		idx := 0
		for i := -1; i < N; i++ {
			for j := 0; j < M; j++ {
				r := grid.RotatedRhombusShapes[idx]
				if seen[r] {
					valid = false
					break
				}
				seen[r] = true
				expectedName := fmt.Sprintf("%s-%d-%d", grid.Name, i, j)
				if r == nil || r.Name != expectedName {
					valid = false
					break
				}
				idx++
			}
			if !valid {
				break
			}
		}
	}

	if !valid {
		grid.RotatedRhombusShapes = make([]*RotatedRhombusShape, 0, (N+1)*M)
		for i := -1; i < N; i++ {
			for j := 0; j < M; j++ {
				r := new(RotatedRhombusShape).Stage(stage)
				r.Name = fmt.Sprintf("%s-%d-%d", grid.Name, i, j)
				origX := float64(i)*v1x + float64(j)*v2x + (v1x+v2x)/2.0
				origY := float64(i)*v1y + float64(j)*v2y + (v1y+v2y)/2.0
				r.X = origX*cosA - origY*sinA
				r.Y = origX*sinA + origY*cosA
				grid.RotatedRhombusShapes = append(grid.RotatedRhombusShapes, r)
			}
		}
		needCommit = true
	} else {
		idx := 0
		for i := -1; i < N; i++ {
			for j := 0; j < M; j++ {
				r := grid.RotatedRhombusShapes[idx]
				origX := float64(i)*v1x + float64(j)*v2x + (v1x+v2x)/2.0
				origY := float64(i)*v1y + float64(j)*v2y + (v1y+v2y)/2.0
				expectedX := origX*cosA - origY*sinA
				expectedY := origX*sinA + origY*cosA
				if r.X != expectedX || r.Y != expectedY {
					r.X = expectedX
					r.Y = expectedY
					needCommit = true
				}
				idx++
			}
		}
	}
	return
}
