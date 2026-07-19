package models

import (
	"fmt"
	"math"
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
		globalR = circumference / (2 * math.Pi)
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

		thickness := plant.RadialThickness
		if thickness == 0 {
			thickness = 5.0
		}

		var curvePoints []vector3
		var topCurvePoints []vector3

		appendArcPoints := func(targetCurve *[]vector3, x1, y1, x2, y2, r float64, sweepFlag bool, largeArcFlag bool) {
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
			appendArcPoints(&curvePoints, sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag)

			if i < len(endArcs) {
				ea := endArcs[i]
				appendArcPoints(&curvePoints, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
			}
		}

		for i := 0; i < len(topStartArcs); i++ {
			tsa := topStartArcs[i]
			appendArcPoints(&topCurvePoints, tsa.StartX, tsa.StartY, tsa.EndX, tsa.EndY, tsa.RadiusX, !tsa.SweepFlag, tsa.LargeArcFlag)

			if i < len(topEndArcs) {
				ea := topEndArcs[i]
				appendArcPoints(&topCurvePoints, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
			}
		}

		if len(curvePoints) > 1 && len(topCurvePoints) > 1 {
			createFaceMeshSTL := func(edges [][2]vector3, reverseWinding bool) {
				for i := 0; i < len(edges)-1; i++ {
					v1 := edges[i][0]
					v2 := edges[i][1]
					v3 := edges[i+1][0]
					v4 := edges[i+1][1]

					if reverseWinding {
						writeFacet(&sb, v1, v3, v2)
						writeFacet(&sb, v2, v3, v4)
					} else {
						writeFacet(&sb, v1, v2, v3)
						writeFacet(&sb, v2, v4, v3)
					}
				}
			}

			stackHeight := plant.StackHeight

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

			for h := 0; h < stackHeight; h++ {
				dx := float64(h)*growthVectorX + float64(h)*verticalThickness*vx
				dy := float64(h)*growthVectorY + float64(h)*verticalThickness*vy

				thetaOffset := dx / globalR

				var bottomEdges, topEdges, innerEdges, outerEdges [][2]vector3

				for i := 0; i < len(curvePoints) && i < len(topCurvePoints); i++ {
					p := curvePoints[i]
					pTop := topCurvePoints[i]

					thetaBase := math.Atan2(p.Z, p.X)
					theta := thetaBase + thetaOffset

					thetaBaseTop := math.Atan2(pTop.Z, pTop.X)
					thetaTop := thetaBaseTop + thetaOffset

					yBase := p.Y + dy
					yBaseTop := pTop.Y + dy

					rBase := math.Sqrt(p.X*p.X + p.Z*p.Z)
					rOuter := rBase + thickness

					rBaseTop := math.Sqrt(pTop.X*pTop.X + pTop.Z*pTop.Z)
					rOuterTop := rBaseTop + thickness

					xBL := rBase * math.Cos(theta)
					zBL := rBase * math.Sin(theta)
					yBL := yBase

					xBR := rOuter * math.Cos(theta)
					zBR := rOuter * math.Sin(theta)
					yBR := yBase

					xTL := rBaseTop * math.Cos(thetaTop)
					zTL := rBaseTop * math.Sin(thetaTop)
					yTL := yBaseTop

					xTR := rOuterTop * math.Cos(thetaTop)
					zTR := rOuterTop * math.Sin(thetaTop)
					yTR := yBaseTop

					vBL := vector3{X: xBL, Y: yBL, Z: zBL}
					vBR := vector3{X: xBR, Y: yBR, Z: zBR}
					vTL := vector3{X: xTL, Y: yTL, Z: zTL}
					vTR := vector3{X: xTR, Y: yTR, Z: zTR}

					bottomEdges = append(bottomEdges, [2]vector3{vBL, vBR})
					topEdges = append(topEdges, [2]vector3{vTL, vTR})
					innerEdges = append(innerEdges, [2]vector3{vBL, vTL})
					outerEdges = append(outerEdges, [2]vector3{vBR, vTR})
				}

				createFaceMeshSTL(bottomEdges, false)
				createFaceMeshSTL(topEdges, true)
				createFaceMeshSTL(innerEdges, true)
				createFaceMeshSTL(outerEdges, false)
			}
		}
	}

	sb.WriteString(fmt.Sprintf("endsolid %s\n", plant.Name))

	return sb.String()
}
