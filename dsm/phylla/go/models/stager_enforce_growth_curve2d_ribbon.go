package models

import (
	"fmt"
)

func enforceGrowthCurve2DRibbonHasShapes(
	stage *Stage,
	plant *Plant,
) (needCommit bool) {
	ribbon := plant.GrowthCurve2DRibbon
	stack := plant.StackOfRotatedGrowthCurve2DRibbon

	if ribbon == nil || stack == nil || plant.PerpendicularVectorGrid == nil || len(plant.PerpendicularVectorGrid.PerpendicularVectors) < 2 {
		if ribbon != nil && (len(ribbon.GrowthCurve2DRibbonStartShapes) > 0 || len(ribbon.GrowthCurve2DRibbonEndShapes) > 0) {
			for _, s := range ribbon.GrowthCurve2DRibbonStartShapes {
				s.Unstage(stage)
			}
			for _, s := range ribbon.GrowthCurve2DRibbonEndShapes {
				s.Unstage(stage)
			}
			ribbon.GrowthCurve2DRibbonStartShapes = nil
			ribbon.GrowthCurve2DRibbonEndShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(plant.PerpendicularVectorGrid.PerpendicularVectors) - 1

	if len(stack.StackRotatedGrowthCurve2DRibbonStartShapes) < expectedLen || len(stack.StackRotatedGrowthCurve2DRibbonEndShapes) < expectedLen {
		return false
	}

	type expectedShape struct {
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

	var expectedStart []expectedShape
	var expectedEnd []expectedShape

	for i := 0; i < expectedLen; i++ {
		s := stack.StackRotatedGrowthCurve2DRibbonStartShapes[i]
		expectedStart = append(expectedStart, expectedShape{
			name:         fmt.Sprintf("%s-start-%d", ribbon.Name, i),
			bottomStartX: s.BottomStartX, bottomStartY: s.BottomStartY,
			bottomEndX: s.BottomEndX, bottomEndY: s.BottomEndY,
			bottomRadiusX: s.BottomRadiusX, bottomRadiusY: s.BottomRadiusY,
			bottomXAxisRotation: s.BottomXAxisRotation, bottomLargeArcFlag: s.BottomLargeArcFlag, bottomSweepFlag: s.BottomSweepFlag,

			topStartX: s.TopStartX, topStartY: s.TopStartY,
			topEndX: s.TopEndX, topEndY: s.TopEndY,
			topRadiusX: s.TopRadiusX, topRadiusY: s.TopRadiusY,
			topXAxisRotation: s.TopXAxisRotation, topLargeArcFlag: s.TopLargeArcFlag, topSweepFlag: s.TopSweepFlag,
		})

		e := stack.StackRotatedGrowthCurve2DRibbonEndShapes[i]
		expectedEnd = append(expectedEnd, expectedShape{
			name:         fmt.Sprintf("%s-end-%d", ribbon.Name, i),
			bottomStartX: e.BottomStartX, bottomStartY: e.BottomStartY,
			bottomEndX: e.BottomEndX, bottomEndY: e.BottomEndY,
			bottomRadiusX: e.BottomRadiusX, bottomRadiusY: e.BottomRadiusY,
			bottomXAxisRotation: e.BottomXAxisRotation, bottomLargeArcFlag: e.BottomLargeArcFlag, bottomSweepFlag: e.BottomSweepFlag,

			topStartX: e.TopStartX, topStartY: e.TopStartY,
			topEndX: e.TopEndX, topEndY: e.TopEndY,
			topRadiusX: e.TopRadiusX, topRadiusY: e.TopRadiusY,
			topXAxisRotation: e.TopXAxisRotation, topLargeArcFlag: e.TopLargeArcFlag, topSweepFlag: e.TopSweepFlag,
		})
	}

	if len(ribbon.GrowthCurve2DRibbonStartShapes) != len(expectedStart) {
		for _, b := range ribbon.GrowthCurve2DRibbonStartShapes {
			b.Unstage(stage)
		}
		ribbon.GrowthCurve2DRibbonStartShapes = make([]*GrowthCurve2DRibbonStartShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(GrowthCurve2DRibbonStartShape).Stage(stage)
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

			ribbon.GrowthCurve2DRibbonStartShapes[i] = b
		}
		needCommit = true
	} else {
		for i, exp := range expectedStart {
			b := ribbon.GrowthCurve2DRibbonStartShapes[i]
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

	if len(ribbon.GrowthCurve2DRibbonEndShapes) != len(expectedEnd) {
		for _, b := range ribbon.GrowthCurve2DRibbonEndShapes {
			b.Unstage(stage)
		}
		ribbon.GrowthCurve2DRibbonEndShapes = make([]*GrowthCurve2DRibbonEndShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(GrowthCurve2DRibbonEndShape).Stage(stage)
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

			ribbon.GrowthCurve2DRibbonEndShapes[i] = b
		}
		needCommit = true
	} else {
		for i, exp := range expectedEnd {
			b := ribbon.GrowthCurve2DRibbonEndShapes[i]
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

func enforceShiftedRightGrowthCurve2DRibbonHasShapes(
	stage *Stage,
	plant *Plant,
) (needCommit bool) {
	ribbon := plant.ShiftedRightGrowthCurve2DRibbon
	stack := plant.StackOfRotatedGrowthCurve2DRibbon

	if ribbon == nil || stack == nil || plant.PerpendicularVectorGrid == nil || len(plant.PerpendicularVectorGrid.PerpendicularVectors) < 2 {
		if ribbon != nil && (len(ribbon.ShiftedRightGrowthCurve2DRibbonStartShapes) > 0 || len(ribbon.ShiftedRightGrowthCurve2DRibbonEndShapes) > 0) {
			for _, s := range ribbon.ShiftedRightGrowthCurve2DRibbonStartShapes {
				s.Unstage(stage)
			}
			for _, s := range ribbon.ShiftedRightGrowthCurve2DRibbonEndShapes {
				s.Unstage(stage)
			}
			ribbon.ShiftedRightGrowthCurve2DRibbonStartShapes = nil
			ribbon.ShiftedRightGrowthCurve2DRibbonEndShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(plant.PerpendicularVectorGrid.PerpendicularVectors) - 1

	if len(stack.StackRotatedGrowthCurve2DRibbonStartShapes) < expectedLen || len(stack.StackRotatedGrowthCurve2DRibbonEndShapes) < expectedLen {
		return false
	}

	type expectedShape struct {
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

	var expectedStart []expectedShape
	var expectedEnd []expectedShape

	for i := 0; i < expectedLen; i++ {
		s := stack.StackRotatedGrowthCurve2DRibbonStartShapes[i]
		expectedStart = append(expectedStart, expectedShape{
			name:         fmt.Sprintf("%s-start-%d", ribbon.Name, i),
			bottomStartX: s.BottomStartX, bottomStartY: s.BottomStartY,
			bottomEndX: s.BottomEndX, bottomEndY: s.BottomEndY,
			bottomRadiusX: s.BottomRadiusX, bottomRadiusY: s.BottomRadiusY,
			bottomXAxisRotation: s.BottomXAxisRotation, bottomLargeArcFlag: s.BottomLargeArcFlag, bottomSweepFlag: s.BottomSweepFlag,

			topStartX: s.TopStartX, topStartY: s.TopStartY,
			topEndX: s.TopEndX, topEndY: s.TopEndY,
			topRadiusX: s.TopRadiusX, topRadiusY: s.TopRadiusY,
			topXAxisRotation: s.TopXAxisRotation, topLargeArcFlag: s.TopLargeArcFlag, topSweepFlag: s.TopSweepFlag,
		})

		e := stack.StackRotatedGrowthCurve2DRibbonEndShapes[i]
		expectedEnd = append(expectedEnd, expectedShape{
			name:         fmt.Sprintf("%s-end-%d", ribbon.Name, i),
			bottomStartX: e.BottomStartX, bottomStartY: e.BottomStartY,
			bottomEndX: e.BottomEndX, bottomEndY: e.BottomEndY,
			bottomRadiusX: e.BottomRadiusX, bottomRadiusY: e.BottomRadiusY,
			bottomXAxisRotation: e.BottomXAxisRotation, bottomLargeArcFlag: e.BottomLargeArcFlag, bottomSweepFlag: e.BottomSweepFlag,

			topStartX: e.TopStartX, topStartY: e.TopStartY,
			topEndX: e.TopEndX, topEndY: e.TopEndY,
			topRadiusX: e.TopRadiusX, topRadiusY: e.TopRadiusY,
			topXAxisRotation: e.TopXAxisRotation, topLargeArcFlag: e.TopLargeArcFlag, topSweepFlag: e.TopSweepFlag,
		})
	}

	if len(ribbon.ShiftedRightGrowthCurve2DRibbonStartShapes) != len(expectedStart) {
		for _, b := range ribbon.ShiftedRightGrowthCurve2DRibbonStartShapes {
			b.Unstage(stage)
		}
		ribbon.ShiftedRightGrowthCurve2DRibbonStartShapes = make([]*ShiftedRightGrowthCurve2DRibbonStartShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(ShiftedRightGrowthCurve2DRibbonStartShape).Stage(stage)
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

			ribbon.ShiftedRightGrowthCurve2DRibbonStartShapes[i] = b
		}
		needCommit = true
	} else {
		for i, exp := range expectedStart {
			b := ribbon.ShiftedRightGrowthCurve2DRibbonStartShapes[i]
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

	if len(ribbon.ShiftedRightGrowthCurve2DRibbonEndShapes) != len(expectedEnd) {
		for _, b := range ribbon.ShiftedRightGrowthCurve2DRibbonEndShapes {
			b.Unstage(stage)
		}
		ribbon.ShiftedRightGrowthCurve2DRibbonEndShapes = make([]*ShiftedRightGrowthCurve2DRibbonEndShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(ShiftedRightGrowthCurve2DRibbonEndShape).Stage(stage)
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

			ribbon.ShiftedRightGrowthCurve2DRibbonEndShapes[i] = b
		}
		needCommit = true
	} else {
		for i, exp := range expectedEnd {
			b := ribbon.ShiftedRightGrowthCurve2DRibbonEndShapes[i]
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
