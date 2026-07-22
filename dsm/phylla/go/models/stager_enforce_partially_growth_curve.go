package models

import (
	"fmt"
	"math"
)

func enforcePartiallyGrowthCurve2DRibbonHasShapes(
	stage *Stage,
	plant *Plant,
) (needCommit bool) {
	ribbon := plant.PartiallyGrowthCurve2DRibbon
	baseRibbonStack := plant.StackOfGrowthCurve2DRibbon

	if ribbon == nil || baseRibbonStack == nil || plant.GrowthVectorShape == nil || plant.RhombusStuff == nil || plant.RhombusStuff.PlantCircumferenceShape == nil {
		if ribbon != nil && (len(ribbon.PartiallyGrowthCurve2DRibbonStartShapes) > 0 || len(ribbon.PartiallyGrowthCurve2DRibbonEndShapes) > 0) {
			ribbon.PartiallyGrowthCurve2DRibbonStartShapes = nil
			ribbon.PartiallyGrowthCurve2DRibbonEndShapes = nil
			return true
		}
		return false
	}

	if len(plant.PerpendicularVectorGrid.PerpendicularVectors) < 2 {
		return false
	}

	expectedLen := len(plant.PerpendicularVectorGrid.PerpendicularVectors) - 1

	if len(baseRibbonStack.StackGrowthCurve2DRibbonStartShapes) < expectedLen || len(baseRibbonStack.StackGrowthCurve2DRibbonEndShapes) < expectedLen {
		return false
	}

	_, dy, currentDX := ComputePartiallyGrowthCurveDY(plant)

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

	for i := 0; i < expectedLen; i++ {
		b := baseRibbonStack.StackGrowthCurve2DRibbonStartShapes[i]
		expectedStart = append(expectedStart, expectedStartShape{
			name:         fmt.Sprintf("%s-start-%d", ribbon.Name, i),
			bottomStartX: b.BottomStartX + currentDX, bottomStartY: b.BottomStartY + dy,
			bottomEndX: b.BottomEndX + currentDX, bottomEndY: b.BottomEndY + dy,
			bottomRadiusX: b.BottomRadiusX, bottomRadiusY: b.BottomRadiusY,
			bottomXAxisRotation: b.BottomXAxisRotation, bottomLargeArcFlag: b.BottomLargeArcFlag, bottomSweepFlag: b.BottomSweepFlag,

			topStartX: b.TopStartX + currentDX, topStartY: b.TopStartY + dy,
			topEndX: b.TopEndX + currentDX, topEndY: b.TopEndY + dy,
			topRadiusX: b.TopRadiusX, topRadiusY: b.TopRadiusY,
			topXAxisRotation: b.TopXAxisRotation, topLargeArcFlag: b.TopLargeArcFlag, topSweepFlag: b.TopSweepFlag,
		})

		e := baseRibbonStack.StackGrowthCurve2DRibbonEndShapes[i]
		expectedEnd = append(expectedEnd, expectedStartShape{
			name:         fmt.Sprintf("%s-end-%d", ribbon.Name, i),
			bottomStartX: e.BottomStartX + currentDX, bottomStartY: e.BottomStartY + dy,
			bottomEndX: e.BottomEndX + currentDX, bottomEndY: e.BottomEndY + dy,
			bottomRadiusX: e.BottomRadiusX, bottomRadiusY: e.BottomRadiusY,
			bottomXAxisRotation: e.BottomXAxisRotation, bottomLargeArcFlag: e.BottomLargeArcFlag, bottomSweepFlag: e.BottomSweepFlag,

			topStartX: e.TopStartX + currentDX, topStartY: e.TopStartY + dy,
			topEndX: e.TopEndX + currentDX, topEndY: e.TopEndY + dy,
			topRadiusX: e.TopRadiusX, topRadiusY: e.TopRadiusY,
			topXAxisRotation: e.TopXAxisRotation, topLargeArcFlag: e.TopLargeArcFlag, topSweepFlag: e.TopSweepFlag,
		})
	}

	if len(ribbon.PartiallyGrowthCurve2DRibbonStartShapes) != len(expectedStart) {
		for _, b := range ribbon.PartiallyGrowthCurve2DRibbonStartShapes {
			b.Unstage(stage)
		}
		ribbon.PartiallyGrowthCurve2DRibbonStartShapes = make([]*PartiallyGrowthCurve2DRibbonStartShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(PartiallyGrowthCurve2DRibbonStartShape).Stage(stage)
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

			ribbon.PartiallyGrowthCurve2DRibbonStartShapes[i] = b
		}
		needCommit = true
	} else {
		for i, exp := range expectedStart {
			b := ribbon.PartiallyGrowthCurve2DRibbonStartShapes[i]
			if b.Name != exp.name ||
				b.BottomStartX != exp.bottomStartX || b.BottomStartY != exp.bottomStartY ||
				b.BottomEndX != exp.bottomEndX || b.BottomEndY != exp.bottomEndY ||
				b.TopStartX != exp.topStartX || b.TopStartY != exp.topStartY ||
				b.TopEndX != exp.topEndX || b.TopEndY != exp.topEndY {

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

	if len(ribbon.PartiallyGrowthCurve2DRibbonEndShapes) != len(expectedEnd) {
		for _, b := range ribbon.PartiallyGrowthCurve2DRibbonEndShapes {
			b.Unstage(stage)
		}
		ribbon.PartiallyGrowthCurve2DRibbonEndShapes = make([]*PartiallyGrowthCurve2DRibbonEndShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(PartiallyGrowthCurve2DRibbonEndShape).Stage(stage)
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

			ribbon.PartiallyGrowthCurve2DRibbonEndShapes[i] = b
		}
		needCommit = true
	} else {
		for i, exp := range expectedEnd {
			b := ribbon.PartiallyGrowthCurve2DRibbonEndShapes[i]
			if b.Name != exp.name ||
				b.BottomStartX != exp.bottomStartX || b.BottomStartY != exp.bottomStartY ||
				b.BottomEndX != exp.bottomEndX || b.BottomEndY != exp.bottomEndY ||
				b.TopStartX != exp.topStartX || b.TopStartY != exp.topStartY ||
				b.TopEndX != exp.topEndX || b.TopEndY != exp.topEndY {

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

func enforcePartiallyGrowthCurve2DTrajectoryHasShapes(
	stage *Stage,
	plant *Plant,
) (needCommit bool) {
	traj := plant.PartiallyGrowthCurve2DTrajectory
	baseRibbonStack := plant.StackOfGrowthCurve2DRibbon

	if traj == nil || baseRibbonStack == nil || plant.GrowthVectorShape == nil || plant.RhombusStuff == nil || plant.RhombusStuff.PlantCircumferenceShape == nil || plant.PerpendicularVectorGrid == nil || len(plant.PerpendicularVectorGrid.PerpendicularVectors) < 2 {
		if traj != nil && len(traj.PartiallyGrowthCurve2DTrajectoryShapes) > 0 {
			for _, s := range traj.PartiallyGrowthCurve2DTrajectoryShapes {
				s.Unstage(stage)
			}
			traj.PartiallyGrowthCurve2DTrajectoryShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(plant.PerpendicularVectorGrid.PerpendicularVectors) - 1
	if len(baseRibbonStack.StackGrowthCurve2DRibbonStartShapes) < expectedLen {
		return false
	}

	baseShape := baseRibbonStack.StackGrowthCurve2DRibbonStartShapes[0]

	type expectedSegment struct {
		name           string
		startX, startY float64
		endX, endY     float64
	}

	var expected []expectedSegment

	numSteps := 100

	pointsX := make([]float64, numSteps+1)
	pointsY := make([]float64, numSteps+1)

	var prevX, prevY float64
	for step := 0; step <= numSteps; step++ {
		r := float64(step) * 0.01
		_, dy, currentDX := ComputePartiallyGrowthCurveDYForRatio(plant, r)

		x := baseShape.BottomStartX + currentDX
		y := baseShape.BottomStartY + dy

		pointsX[step] = x
		pointsY[step] = y

		if step > 0 {
			expected = append(expected, expectedSegment{
				name:   fmt.Sprintf("%s-segment-%d", traj.Name, step-1),
				startX: prevX,
				startY: prevY,
				endX:   x,
				endY:   y,
			})
		}
		prevX, prevY = x, y
	}

	// Arc Verification Computation: Fit circle through P0 (step 0), P50 (step 50), P100 (step 100)
	x1, y1 := pointsX[0], pointsY[0]
	x2, y2 := pointsX[50], pointsY[50]
	x3, y3 := pointsX[100], pointsY[100]

	D := 2 * (x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2))
	if math.Abs(D) > 1e-6 {
		sq1 := x1*x1 + y1*y1
		sq2 := x2*x2 + y2*y2
		sq3 := x3*x3 + y3*y3

		cx := (sq1*(y2-y3) + sq2*(y3-y1) + sq3*(y1-y2)) / D
		cy := (sq1*(x3-x2) + sq2*(x1-x3) + sq3*(x2-x1)) / D
		radius := math.Hypot(x1-cx, y1-cy)

		maxDev := 0.0
		for step := 0; step <= numSteps; step++ {
			dev := math.Abs(math.Hypot(pointsX[step]-cx, pointsY[step]-cy) - radius)
			if dev > maxDev {
				maxDev = dev
			}
		}
		// isArc := maxDev < 0.1
		// log.Println(fmt.Sprintf("[Trajectory Arc Computation] Start=(%.2f, %.2f), Mid=(%.2f, %.2f), End=(%.2f, %.2f) | Fit Center=(%.2f, %.2f), Radius=%.2f | Max Deviation across 101 points=%.6f pixels | Is Exact Circular Arc: %v",
		// x1, y1, x2, y2, x3, y3, cx, cy, radius, maxDev, isArc))
	}

	if len(traj.PartiallyGrowthCurve2DTrajectoryShapes) != len(expected) {
		for _, s := range traj.PartiallyGrowthCurve2DTrajectoryShapes {
			s.Unstage(stage)
		}
		traj.PartiallyGrowthCurve2DTrajectoryShapes = make([]*PartiallyGrowthCurve2DTrajectoryShape, len(expected))
		for i, exp := range expected {
			s := new(PartiallyGrowthCurve2DTrajectoryShape).Stage(stage)
			s.Name = exp.name
			s.StartX = exp.startX
			s.StartY = exp.startY
			s.EndX = exp.endX
			s.EndY = exp.endY
			traj.PartiallyGrowthCurve2DTrajectoryShapes[i] = s
		}
		needCommit = true
	} else {
		for i, exp := range expected {
			s := traj.PartiallyGrowthCurve2DTrajectoryShapes[i]
			if s.Name != exp.name ||
				math.Abs(s.StartX-exp.startX) > 1e-4 || math.Abs(s.StartY-exp.startY) > 1e-4 ||
				math.Abs(s.EndX-exp.endX) > 1e-4 || math.Abs(s.EndY-exp.endY) > 1e-4 {
				s.Name = exp.name
				s.StartX = exp.startX
				s.StartY = exp.startY
				s.EndX = exp.endX
				s.EndY = exp.endY
				needCommit = true
			}
		}
	}

	return needCommit
}
