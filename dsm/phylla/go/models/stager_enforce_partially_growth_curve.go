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

func enforceShiftedLeftPartiallyGrowthCurve2DRibbonHasShapes(
	stage *Stage,
	plant *Plant,
) (needCommit bool) {
	ribbon := plant.ShiftedLeftPartiallyGrowthCurve2DRibbon
	baseRibbonStack := plant.StackOfGrowthCurve2DRibbon

	if ribbon == nil || baseRibbonStack == nil || plant.GrowthVectorShape == nil || plant.RhombusStuff == nil || plant.RhombusStuff.PlantCircumferenceShape == nil {
		if ribbon != nil && (len(ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes) > 0 || len(ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes) > 0) {
			for _, s := range ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes {
				s.Unstage(stage)
			}
			for _, s := range ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes {
				s.Unstage(stage)
			}
			ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes = nil
			ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes = nil
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

	if len(ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes) != len(expectedStart) {
		for _, b := range ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes {
			b.Unstage(stage)
		}
		ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes = make([]*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape).Stage(stage)
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

			ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes[i] = b
		}
		needCommit = true
	} else {
		for i, exp := range expectedStart {
			b := ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes[i]
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

	if len(ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes) != len(expectedEnd) {
		for _, b := range ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes {
			b.Unstage(stage)
		}
		ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes = make([]*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape).Stage(stage)
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

			ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes[i] = b
		}
		needCommit = true
	} else {
		for i, exp := range expectedEnd {
			b := ribbon.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes[i]
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

	circLen := plant.RhombusStuff.PlantCircumferenceShape.Length
	trajOffsetX := plant.RelativeTrajectoryOffsetX * circLen
	trajOffsetY := plant.RelativeTrajectoryOffsetY * circLen

	var prevX, prevY float64
	for step := 0; step <= numSteps; step++ {
		r := float64(step) * 0.01
		_, dy, currentDX := ComputePartiallyGrowthCurveDYForRatio(plant, r)

		x := baseShape.BottomStartX + currentDX + trajOffsetX
		y := baseShape.BottomStartY + dy + trajOffsetY

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

	needCommitP1P2 := enforcePartiallyGrowthCurve2DTrajectoryP1P2HasShapes(stage, plant, pointsX, pointsY)
	needCommit = needCommit || needCommitP1P2

	_, dyCurrent, currentDXCurrent := ComputePartiallyGrowthCurveDY(plant)
	pxX := baseShape.BottomStartX + currentDXCurrent + trajOffsetX
	pxY := baseShape.BottomStartY + dyCurrent + trajOffsetY
	needCommitPx := enforcePxShape(stage, plant, pxX, pxY)
	needCommit = needCommit || needCommitPx

	return needCommit
}

func enforcePxShape(stage *Stage, plant *Plant, pxX, pxY float64) (needCommit bool) {
	if plant.PxShape == nil {
		return false
	}
	expectedName := fmt.Sprintf("%s-Px", plant.Name)
	if plant.PxShape.Name != expectedName || math.Abs(plant.PxShape.X-pxX) > 1e-4 || math.Abs(plant.PxShape.Y-pxY) > 1e-4 {
		plant.PxShape.Name = expectedName
		plant.PxShape.X = pxX
		plant.PxShape.Y = pxY
		needCommit = true
	}
	return needCommit
}


func enforcePartiallyGrowthCurve2DTrajectoryP1P2HasShapes(
	stage *Stage,
	plant *Plant,
	pointsX []float64,
	pointsY []float64,
) (needCommit bool) {
	p1p2 := plant.PartiallyGrowthCurve2DTrajectoryP1P2
	if p1p2 == nil {
		return false
	}

	clearAll := func() bool {
		nc := false
		if len(p1p2.P1PointShapes) > 0 {
			for _, s := range p1p2.P1PointShapes {
				s.Unstage(stage)
			}
			p1p2.P1PointShapes = nil
			nc = true
		}
		if len(p1p2.P2PointShapes) > 0 {
			for _, s := range p1p2.P2PointShapes {
				s.Unstage(stage)
			}
			p1p2.P2PointShapes = nil
			nc = true
		}
		if len(p1p2.P1CurveShapes) > 0 {
			for _, s := range p1p2.P1CurveShapes {
				s.Unstage(stage)
			}
			p1p2.P1CurveShapes = nil
			nc = true
		}
		if len(p1p2.P2CurveShapes) > 0 {
			for _, s := range p1p2.P2CurveShapes {
				s.Unstage(stage)
			}
			p1p2.P2CurveShapes = nil
			nc = true
		}
		if len(p1p2.P1P2PairLineShapes) > 0 {
			for _, s := range p1p2.P1P2PairLineShapes {
				s.Unstage(stage)
			}
			p1p2.P1P2PairLineShapes = nil
			nc = true
		}
		return nc
	}

	numSteps := len(pointsX) - 1
	if numSteps < 2 {
		return clearAll()
	}

	x1, y1 := pointsX[0], pointsY[0]
	x2, y2 := pointsX[numSteps/2], pointsY[numSteps/2]
	x3, y3 := pointsX[numSteps], pointsY[numSteps]

	D := 2 * (x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2))
	if math.Abs(D) <= 1e-6 {
		return clearAll()
	}

	sq1 := x1*x1 + y1*y1
	sq2 := x2*x2 + y2*y2
	sq3 := x3*x3 + y3*y3

	cx := (sq1*(y2-y3) + sq2*(y3-y1) + sq3*(y1-y2)) / D
	cy := (sq1*(x3-x2) + sq2*(x1-x3) + sq3*(x2-x1)) / D
	radius := math.Hypot(x1-cx, y1-cy)

	if radius <= 0 {
		return clearAll()
	}

	maxDev := 0.0
	for step := 0; step <= numSteps; step++ {
		dev := math.Abs(math.Hypot(pointsX[step]-cx, pointsY[step]-cy) - radius)
		if dev > maxDev {
			maxDev = dev
		}
	}

	// Precision condition: maxDev / radius < 0.0001 (precision < 0.01 %)
	if maxDev/radius >= 0.0001 {
		return clearAll()
	}

	// Calculate baseline chord between (x1, y1) and (x3, y3)
	dxChord := x3 - x1
	dyChord := y3 - y1
	chordLen := math.Hypot(dxChord, dyChord)
	if chordLen == 0 {
		return clearAll()
	}

	ux := dxChord / chordLen
	uy := dyChord / chordLen

	// Midpoint of chord
	mx := (x1 + x3) / 2.0
	my := (y1 + y3) / 2.0

	// Normal vector pointing UP towards arc midpoint (x2, y2)
	vxRaw := -uy
	vyRaw := ux
	dotMid := (x2-mx)*vxRaw + (y2-my)*vyRaw
	if dotMid < 0 {
		vxRaw = -vxRaw
		vyRaw = -vyRaw
	}
	vx := vxRaw
	vy := vyRaw

	// R1 is the distance from circle center (cx, cy) to chord line
	R1 := (mx-cx)*vx + (my-cy)*vy
	if R1 < 0 {
		R1 = math.Abs(R1)
	}
	R2 := radius - R1

	if R1 < 0 || R2 <= 0 {
		return clearAll()
	}

	refSteps := plant.NbStepP1P2
	if refSteps <= 0 {
		refSteps = 10
	}
	p1PtsX := make([]float64, refSteps+1)
	p1PtsY := make([]float64, refSteps+1)
	p2PtsX := make([]float64, refSteps+1)
	p2PtsY := make([]float64, refSteps+1)

	for k := 0; k <= refSteps; k++ {
		yVal := float64(k) * R1 / float64(refSteps)
		yLine := -yVal // Cartesian ordinate below chord line

		cK := math.Sqrt(2 * (R1 + yLine) * (R2 - yLine))

		// P1 is on left (-cK*ux), P2 is on right (+cK*ux), both at yLine below chord (-yVal*vx, -yVal*vy)
		p1PtsX[k] = mx - cK*ux - yVal*vx
		p1PtsY[k] = my - cK*uy - yVal*vy

		p2PtsX[k] = mx + cK*ux - yVal*vx
		p2PtsY[k] = my + cK*uy - yVal*vy
	}

	// Reconcile P1PointShapes (refSteps + 1 = 11)
	if len(p1p2.P1PointShapes) != refSteps+1 {
		for _, s := range p1p2.P1PointShapes {
			s.Unstage(stage)
		}
		p1p2.P1PointShapes = make([]*PartiallyGrowthCurve2DTrajectoryP1PointShape, refSteps+1)
		for k := 0; k <= refSteps; k++ {
			s := new(PartiallyGrowthCurve2DTrajectoryP1PointShape).Stage(stage)
			s.Name = fmt.Sprintf("%s-P1-Step-%d", p1p2.Name, k)
			s.X = p1PtsX[k]
			s.Y = p1PtsY[k]
			p1p2.P1PointShapes[k] = s
		}
		needCommit = true
	} else {
		for k := 0; k <= refSteps; k++ {
			s := p1p2.P1PointShapes[k]
			expectedName := fmt.Sprintf("%s-P1-Step-%d", p1p2.Name, k)
			if s.Name != expectedName || math.Abs(s.X-p1PtsX[k]) > 1e-4 || math.Abs(s.Y-p1PtsY[k]) > 1e-4 {
				s.Name = expectedName
				s.X = p1PtsX[k]
				s.Y = p1PtsY[k]
				needCommit = true
			}
		}
	}

	// Reconcile P2PointShapes (refSteps + 1 = 11)
	if len(p1p2.P2PointShapes) != refSteps+1 {
		for _, s := range p1p2.P2PointShapes {
			s.Unstage(stage)
		}
		p1p2.P2PointShapes = make([]*PartiallyGrowthCurve2DTrajectoryP2PointShape, refSteps+1)
		for k := 0; k <= refSteps; k++ {
			s := new(PartiallyGrowthCurve2DTrajectoryP2PointShape).Stage(stage)
			s.Name = fmt.Sprintf("%s-P2-Step-%d", p1p2.Name, k)
			s.X = p2PtsX[k]
			s.Y = p2PtsY[k]
			p1p2.P2PointShapes[k] = s
		}
		needCommit = true
	} else {
		for k := 0; k <= refSteps; k++ {
			s := p1p2.P2PointShapes[k]
			expectedName := fmt.Sprintf("%s-P2-Step-%d", p1p2.Name, k)
			if s.Name != expectedName || math.Abs(s.X-p2PtsX[k]) > 1e-4 || math.Abs(s.Y-p2PtsY[k]) > 1e-4 {
				s.Name = expectedName
				s.X = p2PtsX[k]
				s.Y = p2PtsY[k]
				needCommit = true
			}
		}
	}

	// Reconcile P1CurveShapes (refSteps = 10 line segments)
	if len(p1p2.P1CurveShapes) != refSteps {
		for _, s := range p1p2.P1CurveShapes {
			s.Unstage(stage)
		}
		p1p2.P1CurveShapes = make([]*PartiallyGrowthCurve2DTrajectoryP1CurveShape, refSteps)
		for k := 0; k < refSteps; k++ {
			s := new(PartiallyGrowthCurve2DTrajectoryP1CurveShape).Stage(stage)
			s.Name = fmt.Sprintf("%s-P1-Curve-Seg-%d", p1p2.Name, k)
			s.StartX = p1PtsX[k]
			s.StartY = p1PtsY[k]
			s.EndX = p1PtsX[k+1]
			s.EndY = p1PtsY[k+1]
			p1p2.P1CurveShapes[k] = s
		}
		needCommit = true
	} else {
		for k := 0; k < refSteps; k++ {
			s := p1p2.P1CurveShapes[k]
			expectedName := fmt.Sprintf("%s-P1-Curve-Seg-%d", p1p2.Name, k)
			if s.Name != expectedName || math.Abs(s.StartX-p1PtsX[k]) > 1e-4 || math.Abs(s.StartY-p1PtsY[k]) > 1e-4 || math.Abs(s.EndX-p1PtsX[k+1]) > 1e-4 || math.Abs(s.EndY-p1PtsY[k+1]) > 1e-4 {
				s.Name = expectedName
				s.StartX = p1PtsX[k]
				s.StartY = p1PtsY[k]
				s.EndX = p1PtsX[k+1]
				s.EndY = p1PtsY[k+1]
				needCommit = true
			}
		}
	}

	// Reconcile P2CurveShapes (refSteps = 10 line segments)
	if len(p1p2.P2CurveShapes) != refSteps {
		for _, s := range p1p2.P2CurveShapes {
			s.Unstage(stage)
		}
		p1p2.P2CurveShapes = make([]*PartiallyGrowthCurve2DTrajectoryP2CurveShape, refSteps)
		for k := 0; k < refSteps; k++ {
			s := new(PartiallyGrowthCurve2DTrajectoryP2CurveShape).Stage(stage)
			s.Name = fmt.Sprintf("%s-P2-Curve-Seg-%d", p1p2.Name, k)
			s.StartX = p2PtsX[k]
			s.StartY = p2PtsY[k]
			s.EndX = p2PtsX[k+1]
			s.EndY = p2PtsY[k+1]
			p1p2.P2CurveShapes[k] = s
		}
		needCommit = true
	} else {
		for k := 0; k < refSteps; k++ {
			s := p1p2.P2CurveShapes[k]
			expectedName := fmt.Sprintf("%s-P2-Curve-Seg-%d", p1p2.Name, k)
			if s.Name != expectedName || math.Abs(s.StartX-p2PtsX[k]) > 1e-4 || math.Abs(s.StartY-p2PtsY[k]) > 1e-4 || math.Abs(s.EndX-p2PtsX[k+1]) > 1e-4 || math.Abs(s.EndY-p2PtsY[k+1]) > 1e-4 {
				s.Name = expectedName
				s.StartX = p2PtsX[k]
				s.StartY = p2PtsY[k]
				s.EndX = p2PtsX[k+1]
				s.EndY = p2PtsY[k+1]
				needCommit = true
			}
		}
	}

	// Reconcile P1P2PairLineShapes (refSteps + 1 = 11 line segments connecting P1(k) to P2(k))
	if len(p1p2.P1P2PairLineShapes) != refSteps+1 {

		for _, s := range p1p2.P1P2PairLineShapes {
			s.Unstage(stage)
		}
		p1p2.P1P2PairLineShapes = make([]*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape, refSteps+1)
		for k := 0; k <= refSteps; k++ {
			s := new(PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape).Stage(stage)
			s.Name = fmt.Sprintf("%s-P1P2-Pair-Line-%d", p1p2.Name, k)
			s.StartX = p1PtsX[k]
			s.StartY = p1PtsY[k]
			s.EndX = p2PtsX[k]
			s.EndY = p2PtsY[k]
			p1p2.P1P2PairLineShapes[k] = s
		}
		needCommit = true
	} else {
		for k := 0; k <= refSteps; k++ {
			s := p1p2.P1P2PairLineShapes[k]
			expectedName := fmt.Sprintf("%s-P1P2-Pair-Line-%d", p1p2.Name, k)
			if s.Name != expectedName || math.Abs(s.StartX-p1PtsX[k]) > 1e-4 || math.Abs(s.StartY-p1PtsY[k]) > 1e-4 || math.Abs(s.EndX-p2PtsX[k]) > 1e-4 || math.Abs(s.EndY-p2PtsY[k]) > 1e-4 {
				s.Name = expectedName
				s.StartX = p1PtsX[k]
				s.StartY = p1PtsY[k]
				s.EndX = p2PtsX[k]
				s.EndY = p2PtsY[k]
				needCommit = true
			}
		}
	}

	if plant.ChosenP1P2PairShape != nil {
		chosenK := plant.ChosenStep
		if chosenK < 0 {
			chosenK = 0
		}
		if chosenK > refSteps {
			chosenK = refSteps
		}

		p1x := p1PtsX[chosenK]
		p1y := p1PtsY[chosenK]
		p2x := p2PtsX[chosenK]
		p2y := p2PtsY[chosenK]
		var pxX, pxY float64
		if plant.PxShape != nil {
			pxX = plant.PxShape.X
			pxY = plant.PxShape.Y
		}
		d1 := math.Hypot(p1x-pxX, p1y-pxY)
		d2 := math.Hypot(p2x-pxX, p2y-pxY)
		sum := d1 + d2

		expectedName := fmt.Sprintf("%s-ChosenP1P2Pair", plant.Name)
		if plant.ChosenP1P2PairShape.Name != expectedName ||
			math.Abs(plant.ChosenP1P2PairShape.P1X-p1x) > 1e-4 || math.Abs(plant.ChosenP1P2PairShape.P1Y-p1y) > 1e-4 ||
			math.Abs(plant.ChosenP1P2PairShape.P2X-p2x) > 1e-4 || math.Abs(plant.ChosenP1P2PairShape.P2Y-p2y) > 1e-4 ||
			math.Abs(plant.ChosenP1P2PairShape.PxX-pxX) > 1e-4 || math.Abs(plant.ChosenP1P2PairShape.PxY-pxY) > 1e-4 ||
			math.Abs(plant.ChosenP1P2PairShape.DistanceP1Px-d1) > 1e-4 || math.Abs(plant.ChosenP1P2PairShape.DistanceP2Px-d2) > 1e-4 {
			plant.ChosenP1P2PairShape.Name = expectedName
			plant.ChosenP1P2PairShape.P1X = p1x
			plant.ChosenP1P2PairShape.P1Y = p1y
			plant.ChosenP1P2PairShape.P2X = p2x
			plant.ChosenP1P2PairShape.P2Y = p2y
			plant.ChosenP1P2PairShape.PxX = pxX
			plant.ChosenP1P2PairShape.PxY = pxY
			plant.ChosenP1P2PairShape.DistanceP1Px = d1
			plant.ChosenP1P2PairShape.DistanceP2Px = d2
			plant.ChosenP1P2PairShape.DistanceSum = sum
			needCommit = true
		}
	}

	return needCommit
}


