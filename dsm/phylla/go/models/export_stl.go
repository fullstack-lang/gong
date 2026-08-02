package models

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type vector3 struct {
	X, Y, Z float64
}

func (v vector3) sub(other vector3) vector3 {
	return vector3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

func (v vector3) cross(other vector3) vector3 {
	return vector3{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func (v vector3) normalize() vector3 {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	if length == 0 {
		return vector3{0, 0, 0}
	}
	return vector3{v.X / length, v.Y / length, v.Z / length}
}

func writeFacet(sb *strings.Builder, v1, v2, v3 vector3) {
	normal := v2.sub(v1).cross(v3.sub(v1)).normalize()
	sb.WriteString(fmt.Sprintf("  facet normal %e %e %e\n", normal.X, normal.Y, normal.Z))
	sb.WriteString("    outer loop\n")
	sb.WriteString(fmt.Sprintf("      vertex %e %e %e\n", v1.X, v1.Y, v1.Z))
	sb.WriteString(fmt.Sprintf("      vertex %e %e %e\n", v2.X, v2.Y, v2.Z))
	sb.WriteString(fmt.Sprintf("      vertex %e %e %e\n", v3.X, v3.Y, v3.Z))
	sb.WriteString("    endloop\n")
	sb.WriteString("  endfacet\n")
}

func unwrapAnglesSTL(pts []vector3) (angles []float64, points []vector3) {
	if len(pts) == 0 {
		return nil, nil
	}
	angleToPoint := make(map[float64]vector3)

	firstP := pts[0]
	lastTheta := math.Atan2(firstP.Z, firstP.X)

	accumulated := lastTheta
	for accumulated < 0 {
		accumulated += 2 * math.Pi
	}
	for accumulated >= 2*math.Pi {
		accumulated -= 2 * math.Pi
	}

	angleToPoint[accumulated] = firstP
	lastTheta = math.Atan2(firstP.Z, firstP.X)

	for i := 1; i < len(pts); i++ {
		p := pts[i]
		theta := math.Atan2(p.Z, p.X)
		diff := theta - lastTheta

		for diff < -math.Pi {
			diff += 2 * math.Pi
		}
		for diff > math.Pi {
			diff -= 2 * math.Pi
		}

		if diff < -1e-7 {
			diff = 0
		}

		accumulated += diff
		angleToPoint[accumulated] = p
		lastTheta = theta
	}

	for a := range angleToPoint {
		angles = append(angles, a)
	}
	sort.Float64s(angles)

	for _, a := range angles {
		points = append(points, angleToPoint[a])
	}

	return angles, points
}

func getTargetAnglesSTL(
	originalBottom []vector3,
	originalTop []vector3,
	degreeInterval float64,
	radialRepetitions int,
) (targetAngles []float64, bottomAngles []float64, bottomPoints []vector3, topAngles []float64, topPoints []vector3) {
	if len(originalBottom) == 0 || len(originalTop) == 0 {
		return nil, nil, nil, nil, nil
	}

	bottomAngles, bottomPoints = unwrapAnglesSTL(originalBottom)
	topAngles, topPoints = unwrapAnglesSTL(originalTop)

	radInterval := degreeInterval * math.Pi / 180.0

	if radialRepetitions < 1 {
		radialRepetitions = 1
	}
	expectedDegrees := 360.0 / float64(radialRepetitions)

	nbPoints := int(math.Round(expectedDegrees / degreeInterval))
	for i := 0; i <= nbPoints; i++ {
		targetAngles = append(targetAngles, float64(i)*radInterval)
	}

	return targetAngles, bottomAngles, bottomPoints, topAngles, topPoints
}

func resampleCurveAtAnglesSTL(
	sortedAngles []float64,
	sortedPoints []vector3,
	targetAngles []float64,
	expectedDegrees float64,
) []vector3 {
	var resampled []vector3
	if len(sortedAngles) == 0 {
		return resampled
	}

	for _, target := range targetAngles {
		evalTarget := target
		if len(sortedAngles) > 0 && expectedDegrees > 0 {
			minA := sortedAngles[0]
			maxA := sortedAngles[len(sortedAngles)-1]
			expectedRad := expectedDegrees * math.Pi / 180.0

			for evalTarget < minA {
				evalTarget += expectedRad
			}
			for evalTarget > maxA {
				evalTarget -= expectedRad
			}
			if evalTarget < minA {
				evalTarget = minA
			}
			if evalTarget > maxA {
				evalTarget = maxA
			}
		}

		searchIdx := sort.SearchFloat64s(sortedAngles, evalTarget)
		idx := -1
		if searchIdx > 0 && searchIdx < len(sortedAngles) {
			idx = searchIdx - 1
		} else if searchIdx == 0 && sortedAngles[0] == evalTarget {
			idx = 0
		}

		isExtrapolating := false
		if idx == -1 {
			isExtrapolating = true
			if searchIdx == 0 {
				idx = 0
			} else {
				idx = len(sortedAngles) - 2
			}
		}

		a1 := sortedAngles[idx]
		a2 := sortedAngles[idx+1]
		p1 := sortedPoints[idx]
		p2 := sortedPoints[idx+1]

		t := 0.0
		if a1 != a2 {
			t = (evalTarget - a1) / (a2 - a1)
		}
		if t < 0 {
			t = 0
		}
		if t > 1 {
			t = 1
		}

		var x, y, z float64
		if isExtrapolating {
			var r float64
			if t == 0 {
				r = math.Hypot(p1.X, p1.Z)
				y = p1.Y
			} else {
				r = math.Hypot(p2.X, p2.Z)
				y = p2.Y
			}
			x = r * math.Cos(target)
			z = r * math.Sin(target)
		} else {
			r1 := math.Hypot(p1.X, p1.Z)
			r2 := math.Hypot(p2.X, p2.Z)
			r := r1 + t*(r2-r1)
			y = p1.Y + t*(p2.Y-p1.Y)
			x = r * math.Cos(target)
			z = r * math.Sin(target)
		}

		resampled = append(resampled, vector3{X: x, Y: y, Z: z})
	}

	return resampled
}

func rotateCurveSTL(source []vector3, thetaOffset float64) []vector3 {
	rotated := make([]vector3, len(source))
	for i, p := range source {
		thetaBase := math.Atan2(p.Z, p.X)
		r := math.Hypot(p.X, p.Z)
		newTheta := thetaBase + thetaOffset
		rotated[i] = vector3{
			X: r * math.Cos(newTheta),
			Y: p.Y,
			Z: r * math.Sin(newTheta),
		}
	}
	return rotated
}

func writeRibbonLayerSTL(
	sb *strings.Builder,
	massiveBottom []vector3,
	massiveTop []vector3,
	dy float64,
	thickness float64,
	radialRepetitions int,
) {
	if len(massiveBottom) == 0 || len(massiveBottom) != len(massiveTop) {
		return
	}

	numPointsPerRep := len(massiveBottom) / radialRepetitions
	if numPointsPerRep == 0 {
		numPointsPerRep = 1
	}

	for i := 0; i < len(massiveBottom)-1; i++ {
		// Skip the bridging quad between separate radial repetitions
		if (i+1)%numPointsPerRep == 0 {
			continue
		}

		pB := massiveBottom[i]
		pB_next := massiveBottom[i+1]
		pT := massiveTop[i]
		pT_next := massiveTop[i+1]

		th_i := math.Atan2(pB.Z, pB.X)
		th_next := math.Atan2(pB_next.Z, pB_next.X)

		thTop_i := math.Atan2(pT.Z, pT.X)
		thTop_next := math.Atan2(pT_next.Z, pT_next.X)

		yB := pB.Y + dy
		yB_next := pB_next.Y + dy
		yT := pT.Y + dy
		yT_next := pT_next.Y + dy

		rB := math.Hypot(pB.X, pB.Z)
		rB_outer := rB + thickness
		rB_next := math.Hypot(pB_next.X, pB_next.Z)
		rB_next_outer := rB_next + thickness

		rT := math.Hypot(pT.X, pT.Z)
		rT_outer := rT + thickness
		rT_next := math.Hypot(pT_next.X, pT_next.Z)
		rT_next_outer := rT_next + thickness

		vIn_bottom := vector3{X: rB * math.Cos(th_i), Y: yB, Z: rB * math.Sin(th_i)}
		vOut_bottom := vector3{X: rB_outer * math.Cos(th_i), Y: yB, Z: rB_outer * math.Sin(th_i)}
		vIn_top := vector3{X: rT * math.Cos(thTop_i), Y: yT, Z: rT * math.Sin(thTop_i)}
		vOut_top := vector3{X: rT_outer * math.Cos(thTop_i), Y: yT, Z: rT_outer * math.Sin(thTop_i)}

		vIn_bottom_next := vector3{X: rB_next * math.Cos(th_next), Y: yB_next, Z: rB_next * math.Sin(th_next)}
		vOut_bottom_next := vector3{X: rB_next_outer * math.Cos(th_next), Y: yB_next, Z: rB_next_outer * math.Sin(th_next)}
		vIn_top_next := vector3{X: rT_next * math.Cos(thTop_next), Y: yT_next, Z: rT_next * math.Sin(thTop_next)}
		vOut_top_next := vector3{X: rT_next_outer * math.Cos(thTop_next), Y: yT_next, Z: rT_next_outer * math.Sin(thTop_next)}

		// 1. Outer Surface (facing outward +R)
		writeFacet(sb, vOut_bottom, vOut_top, vOut_bottom_next)
		writeFacet(sb, vOut_bottom_next, vOut_top, vOut_top_next)

		// 2. Inner Surface (facing inward -R)
		writeFacet(sb, vIn_bottom, vIn_bottom_next, vIn_top)
		writeFacet(sb, vIn_top, vIn_bottom_next, vIn_top_next)

		// 3. Bottom Face (facing downward -Y)
		writeFacet(sb, vIn_bottom, vOut_bottom, vIn_bottom_next)
		writeFacet(sb, vOut_bottom, vOut_bottom_next, vIn_bottom_next)

		// 4. Top Face (facing upward +Y)
		writeFacet(sb, vIn_top, vIn_top_next, vOut_top)
		writeFacet(sb, vOut_top, vIn_top_next, vOut_top_next)
	}
}

type stlLayerConfig struct {
	dx, dy, thetaOffset float64
}

func GenerateSTL(plant *Plant) string {
	if plant == nil {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("solid %s\n", plant.Name))

	var globalR float64
	{
		circumference := 10.0
		if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil && plant.RhombusStuff.PlantCircumferenceShape.Length > 0 {
			circumference = plant.RhombusStuff.PlantCircumferenceShape.Length
		} else if plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 0 {
			pGrid := plant.PerpendicularVectorGrid
			first := pGrid.PerpendicularVectors[0]
			last := pGrid.PerpendicularVectors[len(pGrid.PerpendicularVectors)-1]
			circumference = last.StartX - first.StartX
		}
		if circumference <= 0 {
			circumference = 10.0
		}
		threeDModulo := plant.RadialRepetitions
		if threeDModulo < 1 {
			threeDModulo = 1
		}
		globalR = circumference * float64(threeDModulo) / (2 * math.Pi)
	}

	if plant.GrowthCurve2D != nil && plant.TopGrowthCurve2D != nil &&
		plant.GrowthCurve2D.StartHalfwayArcShapeGrid != nil &&
		plant.TopGrowthCurve2D.TopStartHalfwayArcShapeGrid != nil &&
		len(plant.GrowthCurve2D.StartHalfwayArcShapeGrid.StartHalfwayArcShapes) > 0 &&
		len(plant.TopGrowthCurve2D.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes) > 0 {

		gc := plant.GrowthCurve2D
		tgc := plant.TopGrowthCurve2D

		startArcs := gc.StartHalfwayArcShapeGrid.StartHalfwayArcShapes
		var endArcs []*EndHalfwayArcShape
		if gc.EndHalfwayArcShapeGrid != nil {
			endArcs = gc.EndHalfwayArcShapeGrid.EndHalfwayArcShapes
		}

		topStartArcs := tgc.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes
		var topEndArcs []*TopEndHalfwayArcShape
		if tgc.TopEndHalfwayArcShapeGrid != nil {
			topEndArcs = tgc.TopEndHalfwayArcShapeGrid.TopEndHalfwayArcShapes
		}

		thickness := plant.RelativeRadialThickness * plant.RhombusSideLength
		if thickness == 0 {
			thickness = 5.0
		}

		var curvePoints []vector3
		var topCurvePoints []vector3

		appendArcPointsSTL := func(targetCurve *[]vector3, x1, y1, x2, y2, r float64, sweepFlag bool, largeArcFlag bool) {
			dx := (x1 - x2) / 2.0
			dy := (y1 - y2) / 2.0
			d2 := dx*dx + dy*dy
			var cx, cy float64
			if d2 == 0 || r*r < d2 {
				cx = (x1 + x2) / 2.0
				cy = (y1 + y2) / 2.0
				r = math.Sqrt(d2)
			} else {
				root := math.Sqrt(r*r/d2 - 1.0)
				if largeArcFlag == sweepFlag {
					root = -root
				}
				cx = (x1+x2)/2.0 + root*dy
				cy = (y1+y2)/2.0 - root*dx
			}

			startAngle := math.Atan2(y1-cy, x1-cx)
			endAngle := math.Atan2(y2-cy, x2-cx)

			if sweepFlag {
				for endAngle < startAngle {
					endAngle += 2 * math.Pi
				}
			} else {
				for endAngle > startAngle {
					endAngle -= 2 * math.Pi
				}
			}

			steps := 50
			for i := 0; i <= steps; i++ {
				if i == 0 && len(*targetCurve) > 0 {
					continue
				}

				t := float64(i) / float64(steps)
				angle := startAngle + t*(endAngle-startAngle)
				x2d := cx + r*math.Cos(angle)
				y2d := cy + r*math.Sin(angle)

				theta := x2d / globalR
				x3d := globalR * math.Cos(theta)
				z3d := globalR * math.Sin(theta)
				y3d := y2d

				*targetCurve = append(*targetCurve, vector3{X: x3d, Y: y3d, Z: z3d})
			}
		}

		for i := 0; i < len(startArcs); i++ {
			sa := startArcs[i]
			appendArcPointsSTL(&curvePoints, sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag)

			if i < len(endArcs) {
				ea := endArcs[i]
				appendArcPointsSTL(&curvePoints, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
			}
		}

		for i := 0; i < len(topStartArcs); i++ {
			tsa := topStartArcs[i]
			appendArcPointsSTL(&topCurvePoints, tsa.StartX, tsa.StartY, tsa.EndX, tsa.EndY, tsa.RadiusX, !tsa.SweepFlag, tsa.LargeArcFlag)

			if i < len(topEndArcs) {
				ea := topEndArcs[i]
				appendArcPointsSTL(&topCurvePoints, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
			}
		}

		if len(curvePoints) > 1 && len(topCurvePoints) > 1 {
			targetAngles, anglesBottom, bottomPoints, anglesTop, topPoints := getTargetAnglesSTL(curvePoints, topCurvePoints, 0.5, plant.RadialRepetitions)

			expectedDegrees := 360.0
			if plant.RadialRepetitions > 1 {
				expectedDegrees = 360.0 / float64(plant.RadialRepetitions)
			}

			resampledBaseBottom := resampleCurveAtAnglesSTL(anglesBottom, bottomPoints, targetAngles, expectedDegrees)
			resampledBaseTop := resampleCurveAtAnglesSTL(anglesTop, topPoints, targetAngles, expectedDegrees)

			var checkedDiagram *PlantDiagram
			for _, diagram := range plant.PlantDiagrams {
				if diagram.IsChecked {
					checkedDiagram = diagram
					break
				}
			}

			stackHeight := plant.StackHeight

			var activeShapeRuns [][]stlLayerConfig

			if checkedDiagram != nil {
				if !checkedDiagram.IsHiddenTorusStackShape {
					var growthVectorX, growthVectorY float64
					if plant.GrowthVectorShape != nil {
						growthVectorX = plant.GrowthVectorShape.X
						growthVectorY = plant.GrowthVectorShape.Y
					}
					var vx, vy float64
					if plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 0 {
						pGrid := plant.PerpendicularVectorGrid
						vFirst := pGrid.PerpendicularVectors[0]
						vx = vFirst.EndX - vFirst.StartX
						vy = vFirst.EndY - vFirst.StartY
						vLen := math.Hypot(vx, vy)
						if vLen == 0 {
							vLen = 1
						}
						vx, vy = vx/vLen, vy/vLen
					}
					verticalThickness := plant.RelativeVerticalThickness * plant.RhombusSideLength
					rotatedSeparation := plant.RelativeRotatedTorusSeparation * plant.RhombusSideLength

					var run []stlLayerConfig
					for h := 0; h < stackHeight; h++ {
						dx := float64(h)*growthVectorX + float64(h)*verticalThickness*vx
						dy := float64(h)*growthVectorY + float64(h)*verticalThickness*vy + float64(h)*rotatedSeparation
						thetaOffset := dx / globalR
						run = append(run, stlLayerConfig{dx: dx, dy: dy, thetaOffset: thetaOffset})
					}
					activeShapeRuns = append(activeShapeRuns, run)
				}

				if !checkedDiagram.IsHiddenVerticalTorusStackShape {
					var run []stlLayerConfig
					for h := 0; h < stackHeight; h++ {
						dx := 0.0
						dy := float64(h) * plant.RelativeCuttedStackFloorHeight * plant.RhombusSideLength
						thetaOffset := 0.0
						run = append(run, stlLayerConfig{dx: dx, dy: dy, thetaOffset: thetaOffset})
					}
					activeShapeRuns = append(activeShapeRuns, run)
				}

				if !checkedDiagram.IsHiddenPartiallyRotatedTorusShape {
					dx, dy, _ := ComputePartiallyGrowthCurveDY(plant)
					thetaOffset := dx / globalR
					var run []stlLayerConfig
					run = append(run, stlLayerConfig{dx: 0, dy: 0, thetaOffset: 0})
					run = append(run, stlLayerConfig{dx: dx, dy: dy, thetaOffset: thetaOffset})
					activeShapeRuns = append(activeShapeRuns, run)
				}

				if !checkedDiagram.IsHiddenStackOfPartiallyRotatedTorusShape && stackHeight > 0 {
					numSteps := stackHeight - 1
					dxs := make([]float64, stackHeight)
					dys := make([]float64, stackHeight)
					dxs[0] = 0.0
					dys[0] = 0.0

					if numSteps > 0 {
						totalProgress := plant.RotationRatio * float64(numSteps)
						var cumDX, cumDY float64
						for k := 1; k <= numSteps; k++ {
							var r_k float64
							kFloat := float64(numSteps - k + 1)
							if totalProgress >= kFloat {
								r_k = 1.0
							} else if totalProgress <= kFloat-1.0 {
								r_k = 0.0
							} else {
								r_k = totalProgress - (kFloat - 1.0)
							}
							stepDX, stepDY, _ := ComputePartiallyGrowthCurveDYForRatio(plant, r_k)
							cumDX += stepDX
							cumDY += stepDY
							dxs[k] = cumDX
							dys[k] = cumDY
						}
					}

					var run []stlLayerConfig
					for h := 0; h < stackHeight; h++ {
						dx := dxs[h]
						dy := dys[h]
						thetaOffset := dx / globalR
						run = append(run, stlLayerConfig{dx: dx, dy: dy, thetaOffset: thetaOffset})
					}
					activeShapeRuns = append(activeShapeRuns, run)
				}
			}

			if len(activeShapeRuns) == 0 {
				var run []stlLayerConfig
				for h := 0; h < stackHeight; h++ {
					dy := float64(h) * plant.RelativeCuttedStackFloorHeight * plant.RhombusSideLength
					run = append(run, stlLayerConfig{dx: 0, dy: dy, thetaOffset: 0})
				}
				activeShapeRuns = append(activeShapeRuns, run)
			}

			h_horiz := plant.RelativeHorizontalRingsHeight * plant.RhombusSideLength
			if h_horiz == 0 {
				h_horiz = plant.RelativeVerticalThickness * plant.RhombusSideLength
			}

			for _, run := range activeShapeRuns {
				for h, cfg := range run {
					massiveBottom := make([]vector3, 0, len(resampledBaseBottom)*plant.RadialRepetitions)
					massiveTop := make([]vector3, 0, len(resampledBaseTop)*plant.RadialRepetitions)

					for k := 0; k < plant.RadialRepetitions; k++ {
						baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(plant.RadialRepetitions)
						totalThetaOffset := cfg.thetaOffset + baseThetaOffset

						localBottom := rotateCurveSTL(resampledBaseBottom, totalThetaOffset)
						localTop := rotateCurveSTL(resampledBaseTop, totalThetaOffset)

						massiveBottom = append(massiveBottom, localBottom...)
						massiveTop = append(massiveTop, localTop...)
					}

					writeRibbonLayerSTL(&sb, massiveBottom, massiveTop, cfg.dy, thickness, plant.RadialRepetitions)

					if h == 0 && len(massiveBottom) > 0 {
						minY_bottom := math.MaxFloat64
						for _, p := range massiveBottom {
							yVal := p.Y + cfg.dy
							if yVal < minY_bottom {
								minY_bottom = yVal
							}
						}

						horizBottom := make([]vector3, len(massiveBottom))
						horizTop := make([]vector3, len(massiveBottom))
						for idx, p := range massiveBottom {
							horizBottom[idx] = vector3{X: p.X, Y: minY_bottom - cfg.dy, Z: p.Z}
							horizTop[idx] = vector3{X: p.X, Y: (minY_bottom + h_horiz) - cfg.dy, Z: p.Z}
						}
						writeRibbonLayerSTL(&sb, horizBottom, horizTop, cfg.dy, thickness, plant.RadialRepetitions)
					}

					if h == len(run)-1 && len(massiveTop) > 0 {
						maxY_top := -math.MaxFloat64
						for _, p := range massiveTop {
							yVal := p.Y + cfg.dy
							if yVal > maxY_top {
								maxY_top = yVal
							}
						}

						horizBottom := make([]vector3, len(massiveTop))
						horizTop := make([]vector3, len(massiveTop))
						for idx, p := range massiveTop {
							horizBottom[idx] = vector3{X: p.X, Y: (maxY_top - h_horiz) - cfg.dy, Z: p.Z}
							horizTop[idx] = vector3{X: p.X, Y: maxY_top - cfg.dy, Z: p.Z}
						}
						writeRibbonLayerSTL(&sb, horizBottom, horizTop, cfg.dy, thickness, plant.RadialRepetitions)
					}
				}
			}
		}
	}

	sb.WriteString(fmt.Sprintf("endsolid %s\n", plant.Name))

	return sb.String()
}
