package models

import (
	"fmt"
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
