package models

import (
	"fmt"
	"time"
)

func enforcePlantHasShape[ShapePointerType PointerToGongstruct](
	stager *Stager,
	newShape func() ShapePointerType,
	getShape func(plant *PlantAbstract) ShapePointerType,
	setShape func(plant *PlantAbstract, shape ShapePointerType),
	isOwned func(plant *PlantAbstract, shape ShapePointerType) bool,
	shapeName string,
) (needCommit bool) {
	stage := stager.stage

	// 1. Ensure each Plant has the shape
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stage) {
		var zero ShapePointerType
		if getShape(plant) == zero {
			shapePointer := newShape()
			shapePointer.StageVoid(stage)

			setShape(plant, shapePointer)

			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Added default %s for %s", shapeName, plant.Name))
			}

			needCommit = true
		}
	}

	// 2. Ensure each Shape belongs to exactly one Plant. If orphaned, remove it.
	for shape := range *GetGongstructInstancesSetFromPointerType[ShapePointerType](stage) {
		hasOwner := false
		for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stage) {
			if isOwned(plant, shape) {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			shape.UnstageVoid(stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Removed orphaned %s %s", shapeName, shape.GetName()))
			}
			needCommit = true
		}
	}

	return
}

func enforcePlantShapeName[ShapePointerType PointerToGongstruct](
	stager *Stager,
	getShape func(plant *PlantAbstract) ShapePointerType,
	shapeNameSuffix string,
) (needCommit bool) {
	stage := stager.stage

	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stage) {
		var zero ShapePointerType
		shape := getShape(plant)
		if shape != zero {
			expectedName := plant.Name + "-" + shapeNameSuffix
			if shape.GetName() != expectedName {
				shape.SetName(expectedName)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Renamed %s to %s", shapeNameSuffix, expectedName))
				}
				needCommit = true
			}
		}
	}

	return
}

// enforceAxesShapeName ensures that the name of the AxesShape matches its owning Plant
func (stager *Stager) enforceAxesShapeName() (needCommit bool) {
	return enforcePlantShapeName[*AxesShape](
		stager,
		func(p *PlantAbstract) *AxesShape { return p.AxesShape },
		"AxesShape",
	)
}

// enforceGridPathShapeName ensures that the name of the GridPathShape matches its owning Plant
func (stager *Stager) enforceGridPathShapeName() (needCommit bool) {
	return enforcePlantShapeName[*GridPathShape](
		stager,
		func(p *PlantAbstract) *GridPathShape { return p.RhombusStuff.GridPathShape },
		"GridPathShape",
	)
}

// enforcePlantCircumferenceShapeName ensures that the name of the PlantCircumferenceShape matches its owning Plant
func (stager *Stager) enforcePlantCircumferenceShapeName() (needCommit bool) {
	return enforcePlantShapeName[*PlantCircumferenceShape](
		stager,
		func(p *PlantAbstract) *PlantCircumferenceShape { return p.RhombusStuff.PlantCircumferenceShape },
		"PlantCircumferenceShape",
	)
}

// enforcePlantHasAxes ensures that each Plant has one and only one Axes that belong to it
func (stager *Stager) enforcePlantHasAxes() (needCommit bool) {
	return enforcePlantHasShape[*AxesShape](
		stager,
		func() *AxesShape { return new(AxesShape) },
		func(p *PlantAbstract) *AxesShape { return p.AxesShape },
		func(p *PlantAbstract, shape *AxesShape) { p.AxesShape = shape },
		func(p *PlantAbstract, shape *AxesShape) bool { return p.AxesShape == shape },
		"AxesShape",
	)
}

// enforcePlantHasGridPathShape ensures that each Plant has one and only one GridPathShape that belong to it
func (stager *Stager) enforcePlantHasGridPathShape() (needCommit bool) {
	return enforcePlantHasShape[*GridPathShape](
		stager,
		func() *GridPathShape { return new(GridPathShape) },
		func(p *PlantAbstract) *GridPathShape { return p.RhombusStuff.GridPathShape },
		func(p *PlantAbstract, shape *GridPathShape) { p.RhombusStuff.GridPathShape = shape },
		func(p *PlantAbstract, shape *GridPathShape) bool {
			return p.RhombusStuff.GridPathShape == shape || p.RhombusStuff.RotatedGridPathShape == shape
		},
		"GridPathShape",
	)
}

// enforcePlantHasInitialRhombusGridShape ensures that each Plant has one and only one InitialRhombusGridShape that belong to it
func (stager *Stager) enforcePlantHasInitialRhombusGridShape() (needCommit bool) {
	return enforcePlantHasShape[*InitialRhombusGridShape](
		stager,
		func() *InitialRhombusGridShape { return new(InitialRhombusGridShape) },
		func(p *PlantAbstract) *InitialRhombusGridShape { return p.RhombusStuff.InitialRhombusGridShape },
		func(p *PlantAbstract, shape *InitialRhombusGridShape) { p.RhombusStuff.InitialRhombusGridShape = shape },
		func(p *PlantAbstract, shape *InitialRhombusGridShape) bool {
			return p.RhombusStuff.InitialRhombusGridShape == shape
		},
		"InitialRhombusGridShape",
	)
}

// enforcePlantHasExplanationTextShape ensures that each Plant has one and only one ExplanationTextShape that belong to it
func (stager *Stager) enforcePlantHasExplanationTextShape() (needCommit bool) {
	return enforcePlantHasShape[*ExplanationTextShape](
		stager,
		func() *ExplanationTextShape { return new(ExplanationTextShape) },
		func(p *PlantAbstract) *ExplanationTextShape { return p.RhombusStuff.ExplanationTextShape },
		func(p *PlantAbstract, shape *ExplanationTextShape) { p.RhombusStuff.ExplanationTextShape = shape },
		func(p *PlantAbstract, shape *ExplanationTextShape) bool {
			return p.RhombusStuff.ExplanationTextShape == shape
		},
		"ExplanationTextShape",
	)
}

// enforcePlantHasPlantCircumferenceShape ensures that each Plant has one and only one PlantCircumferenceShape that belong to it
func (stager *Stager) enforcePlantHasPlantCircumferenceShape() (needCommit bool) {
	return enforcePlantHasShape[*PlantCircumferenceShape](
		stager,
		func() *PlantCircumferenceShape { return new(PlantCircumferenceShape) },
		func(p *PlantAbstract) *PlantCircumferenceShape { return p.RhombusStuff.PlantCircumferenceShape },
		func(p *PlantAbstract, shape *PlantCircumferenceShape) { p.RhombusStuff.PlantCircumferenceShape = shape },
		func(p *PlantAbstract, shape *PlantCircumferenceShape) bool {
			return p.RhombusStuff.PlantCircumferenceShape == shape || p.RhombusStuff.RotatedPlantCircumferenceShape == shape
		},
		"PlantCircumferenceShape",
	)
}

// enforcePlantHasReferenceRhombus ensures that each Plant has one and only one ReferenceRhombus that belong to it
func (stager *Stager) enforcePlantHasReferenceRhombus() (needCommit bool) {
	return enforcePlantHasShape[*RhombusShape](
		stager,
		func() *RhombusShape { return new(RhombusShape) },
		func(p *PlantAbstract) *RhombusShape { return p.RhombusStuff.ReferenceRhombus },
		func(p *PlantAbstract, shape *RhombusShape) { p.RhombusStuff.ReferenceRhombus = shape },
		func(p *PlantAbstract, shape *RhombusShape) bool {
			return isRhombusShapeOwnedByPlant(p, shape)
		},
		"ReferenceRhombus",
	)
}

// enforcePlantHasRotatedShapes ensures that each Plant has its Rotated shapes
func (stager *Stager) enforcePlantHasRotatedShapes() (needCommit bool) {
	n1 := enforcePlantHasShape[*RhombusShape](
		stager,
		func() *RhombusShape { return new(RhombusShape) },
		func(p *PlantAbstract) *RhombusShape { return p.RhombusStuff.RotatedReferenceRhombus },
		func(p *PlantAbstract, shape *RhombusShape) { p.RhombusStuff.RotatedReferenceRhombus = shape },
		func(p *PlantAbstract, shape *RhombusShape) bool {
			return isRhombusShapeOwnedByPlant(p, shape)
		},
		"RotatedReferenceRhombus",
	)

	n2 := enforcePlantHasShape[*PlantCircumferenceShape](
		stager,
		func() *PlantCircumferenceShape { return new(PlantCircumferenceShape) },
		func(p *PlantAbstract) *PlantCircumferenceShape { return p.RhombusStuff.RotatedPlantCircumferenceShape },
		func(p *PlantAbstract, shape *PlantCircumferenceShape) {
			p.RhombusStuff.RotatedPlantCircumferenceShape = shape
		},
		func(p *PlantAbstract, shape *PlantCircumferenceShape) bool {
			return p.RhombusStuff.PlantCircumferenceShape == shape || p.RhombusStuff.RotatedPlantCircumferenceShape == shape
		},
		"RotatedPlantCircumferenceShape",
	)

	n3 := enforcePlantHasShape[*GridPathShape](
		stager,
		func() *GridPathShape { return new(GridPathShape) },
		func(p *PlantAbstract) *GridPathShape { return p.RhombusStuff.RotatedGridPathShape },
		func(p *PlantAbstract, shape *GridPathShape) { p.RhombusStuff.RotatedGridPathShape = shape },
		func(p *PlantAbstract, shape *GridPathShape) bool {
			return p.RhombusStuff.GridPathShape == shape || p.RhombusStuff.RotatedGridPathShape == shape
		},
		"RotatedGridPathShape",
	)

	n4 := enforcePlantHasShape[*RotatedRhombusGridShape](
		stager,
		func() *RotatedRhombusGridShape { return new(RotatedRhombusGridShape) },
		func(p *PlantAbstract) *RotatedRhombusGridShape { return p.RhombusStuff.RotatedRhombusGridShape2 },
		func(p *PlantAbstract, shape *RotatedRhombusGridShape) {
			p.RhombusStuff.RotatedRhombusGridShape2 = shape
		},
		func(p *PlantAbstract, shape *RotatedRhombusGridShape) bool {
			return p.RhombusStuff.RotatedRhombusGridShape2 == shape
		},
		"RotatedRhombusGridShape",
	)

	n5 := enforcePlantHasShape[*GrowthCurveRhombusGridShape](
		stager,
		func() *GrowthCurveRhombusGridShape { return new(GrowthCurveRhombusGridShape) },
		func(p *PlantAbstract) *GrowthCurveRhombusGridShape { return p.RhombusStuff.GrowthCurveRhombusGridShape },
		func(p *PlantAbstract, shape *GrowthCurveRhombusGridShape) {
			p.RhombusStuff.GrowthCurveRhombusGridShape = shape
		},
		func(p *PlantAbstract, shape *GrowthCurveRhombusGridShape) bool {
			return p.RhombusStuff.GrowthCurveRhombusGridShape == shape
		},
		"GrowthCurveRhombusGridShape",
	)

	n6 := enforcePlantHasShape[*GrowthVectorShape](
		stager,
		func() *GrowthVectorShape { return new(GrowthVectorShape) },
		func(p *PlantAbstract) *GrowthVectorShape { return p.GrowthVectorShape },
		func(p *PlantAbstract, shape *GrowthVectorShape) { p.GrowthVectorShape = shape },
		func(p *PlantAbstract, shape *GrowthVectorShape) bool {
			return p.GrowthVectorShape == shape
		},
		"GrowthVectorShape",
	)

	n7 := enforcePlantHasShape[*PerpendicularVectorGrid](
		stager,
		func() *PerpendicularVectorGrid { return new(PerpendicularVectorGrid) },
		func(p *PlantAbstract) *PerpendicularVectorGrid { return p.PerpendicularVectorGrid },
		func(p *PlantAbstract, shape *PerpendicularVectorGrid) { p.PerpendicularVectorGrid = shape },
		func(p *PlantAbstract, shape *PerpendicularVectorGrid) bool {
			return p.PerpendicularVectorGrid == shape
		},
		"PerpendicularVectorGrid",
	)
	needCommit = n7 || needCommit

	n7_halfway := enforcePlantHasShape[*PerpendicularVectorGridHalfway](
		stager,
		func() *PerpendicularVectorGridHalfway { return new(PerpendicularVectorGridHalfway) },
		func(p *PlantAbstract) *PerpendicularVectorGridHalfway { return p.PerpendicularVectorGridHalfway },
		func(p *PlantAbstract, shape *PerpendicularVectorGridHalfway) {
			p.PerpendicularVectorGridHalfway = shape
		},
		func(p *PlantAbstract, shape *PerpendicularVectorGridHalfway) bool {
			return p.PerpendicularVectorGridHalfway == shape
		},
		"PerpendicularVectorGridHalfway",
	)
	needCommit = n7_halfway || needCommit

	n7_base := enforcePlantHasShape[*BaseVectorShapeGrid](
		stager,
		func() *BaseVectorShapeGrid { return new(BaseVectorShapeGrid) },
		func(p *PlantAbstract) *BaseVectorShapeGrid { return p.BaseVectorShapeGrid },
		func(p *PlantAbstract, shape *BaseVectorShapeGrid) { p.BaseVectorShapeGrid = shape },
		func(p *PlantAbstract, shape *BaseVectorShapeGrid) bool {
			return p.BaseVectorShapeGrid == shape
		},
		"BaseVectorShapeGrid",
	)
	needCommit = n7_base || needCommit

	n7_arc_normal := enforcePlantHasShape[*ArcNormalVectorShapeGrid](
		stager,
		func() *ArcNormalVectorShapeGrid { return new(ArcNormalVectorShapeGrid) },
		func(p *PlantAbstract) *ArcNormalVectorShapeGrid { return p.ArcNormalVectorShapeGrid },
		func(p *PlantAbstract, shape *ArcNormalVectorShapeGrid) { p.ArcNormalVectorShapeGrid = shape },
		func(p *PlantAbstract, shape *ArcNormalVectorShapeGrid) bool {
			return p.ArcNormalVectorShapeGrid == shape
		},
		"ArcNormalVectorShapeGrid",
	)
	needCommit = n7_arc_normal || needCommit

	n7_arc_v2 := enforcePlantHasShape[*StartArcShapeGrid](
		stager,
		func() *StartArcShapeGrid { return new(StartArcShapeGrid) },
		func(p *PlantAbstract) *StartArcShapeGrid { return p.StartArcShapeGrid },
		func(p *PlantAbstract, shape *StartArcShapeGrid) { p.StartArcShapeGrid = shape },
		func(p *PlantAbstract, shape *StartArcShapeGrid) bool {
			return p.StartArcShapeGrid == shape
		},
		"StartArcShapeV2Grid",
	)
	needCommit = n7_arc_v2 || needCommit

	n7_top_arc_v2 := enforcePlantHasShape[*TopStartArcShapeGrid](
		stager,
		func() *TopStartArcShapeGrid { return new(TopStartArcShapeGrid) },
		func(p *PlantAbstract) *TopStartArcShapeGrid { return p.TopStartArcShapeGrid },
		func(p *PlantAbstract, shape *TopStartArcShapeGrid) { p.TopStartArcShapeGrid = shape },
		func(p *PlantAbstract, shape *TopStartArcShapeGrid) bool {
			return p.TopStartArcShapeGrid == shape
		},
		"TopStartArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2 || needCommit

	n7_arc_v2_end := enforcePlantHasShape[*EndArcShapeGrid](
		stager,
		func() *EndArcShapeGrid { return new(EndArcShapeGrid) },
		func(p *PlantAbstract) *EndArcShapeGrid { return p.EndArcShapeGrid },
		func(p *PlantAbstract, shape *EndArcShapeGrid) { p.EndArcShapeGrid = shape },
		func(p *PlantAbstract, shape *EndArcShapeGrid) bool {
			return p.EndArcShapeGrid == shape
		},
		"EndArcShapeV2Grid",
	)
	needCommit = n7_arc_v2_end || needCommit

	n7_top_arc_v2_end := enforcePlantHasShape[*TopEndArcShapeGrid](
		stager,
		func() *TopEndArcShapeGrid { return new(TopEndArcShapeGrid) },
		func(p *PlantAbstract) *TopEndArcShapeGrid { return p.TopEndArcShapeGrid },
		func(p *PlantAbstract, shape *TopEndArcShapeGrid) { p.TopEndArcShapeGrid = shape },
		func(p *PlantAbstract, shape *TopEndArcShapeGrid) bool {
			return p.TopEndArcShapeGrid == shape
		},
		"TopEndArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2_end || needCommit

	n10 := enforcePlantHasShape[*StackOfRotatedGrowthCurve2D](
		stager,
		func() *StackOfRotatedGrowthCurve2D { return new(StackOfRotatedGrowthCurve2D) },
		func(p *PlantAbstract) *StackOfRotatedGrowthCurve2D { return p.StackOfRotatedGrowthCurve2D },
		func(p *PlantAbstract, shape *StackOfRotatedGrowthCurve2D) { p.StackOfRotatedGrowthCurve2D = shape },
		func(p *PlantAbstract, shape *StackOfRotatedGrowthCurve2D) bool {
			return p.StackOfRotatedGrowthCurve2D == shape
		},
		"StackOfGrowthCurveV2",
	)
	needCommit = n10 || needCommit

	n11 := enforcePlantHasShape[*TopStackOfRotatedGrowthCurve2D](
		stager,
		func() *TopStackOfRotatedGrowthCurve2D { return new(TopStackOfRotatedGrowthCurve2D) },
		func(p *PlantAbstract) *TopStackOfRotatedGrowthCurve2D { return p.TopStackOfRotatedGrowthCurve2D },
		func(p *PlantAbstract, shape *TopStackOfRotatedGrowthCurve2D) {
			p.TopStackOfRotatedGrowthCurve2D = shape
		},
		func(p *PlantAbstract, shape *TopStackOfRotatedGrowthCurve2D) bool {
			return p.TopStackOfRotatedGrowthCurve2D == shape
		},
		"TopStackOfGrowthCurveV2",
	)
	needCommit = n11 || needCommit

	n15 := enforcePlantHasShape[*MidArcVectorShapeGrid](
		stager,
		func() *MidArcVectorShapeGrid { return new(MidArcVectorShapeGrid) },
		func(p *PlantAbstract) *MidArcVectorShapeGrid { return p.MidArcVectorShapeGrid },
		func(p *PlantAbstract, shape *MidArcVectorShapeGrid) { p.MidArcVectorShapeGrid = shape },
		func(p *PlantAbstract, shape *MidArcVectorShapeGrid) bool {
			return p.MidArcVectorShapeGrid == shape
		},
		"MidArcVectorShapeGrid",
	)
	needCommit = n15 || needCommit

	n16 := enforcePlantHasShape[*TopMidArcVectorShapeGrid](
		stager,
		func() *TopMidArcVectorShapeGrid { return new(TopMidArcVectorShapeGrid) },
		func(p *PlantAbstract) *TopMidArcVectorShapeGrid { return p.TopMidArcVectorShapeGrid },
		func(p *PlantAbstract, shape *TopMidArcVectorShapeGrid) { p.TopMidArcVectorShapeGrid = shape },
		func(p *PlantAbstract, shape *TopMidArcVectorShapeGrid) bool {
			return p.TopMidArcVectorShapeGrid == shape
		},
		"TopMidArcVectorShapeGrid",
	)
	needCommit = n16 || needCommit

	n14 := enforcePlantHasShape[*ShiftedBottomTopStartArcShapeGrid](
		stager,
		func() *ShiftedBottomTopStartArcShapeGrid { return new(ShiftedBottomTopStartArcShapeGrid) },
		func(p *PlantAbstract) *ShiftedBottomTopStartArcShapeGrid { return p.ShiftedBottomTopStartArcShapeGrid },
		func(p *PlantAbstract, shape *ShiftedBottomTopStartArcShapeGrid) {
			p.ShiftedBottomTopStartArcShapeGrid = shape
		},
		func(p *PlantAbstract, shape *ShiftedBottomTopStartArcShapeGrid) bool {
			return p.ShiftedBottomTopStartArcShapeGrid == shape
		},
		"ShiftedBottomTopStartArcShapeGrid",
	)
	needCommit = n14 || needCommit

	n17 := enforcePlantHasShape[*StackOfGrowthCurve2D](
		stager,
		func() *StackOfGrowthCurve2D { return new(StackOfGrowthCurve2D) },
		func(p *PlantAbstract) *StackOfGrowthCurve2D { return p.StackOfGrowthCurve2D },
		func(p *PlantAbstract, shape *StackOfGrowthCurve2D) { p.StackOfGrowthCurve2D = shape },
		func(p *PlantAbstract, shape *StackOfGrowthCurve2D) bool {
			return p.StackOfGrowthCurve2D == shape
		},
		"StackOfGrowthCurve2D",
	)
	needCommit = n17 || needCommit

	n18 := enforcePlantHasShape[*TopStackOfGrowthCurve2D](
		stager,
		func() *TopStackOfGrowthCurve2D { return new(TopStackOfGrowthCurve2D) },
		func(p *PlantAbstract) *TopStackOfGrowthCurve2D { return p.TopStackOfGrowthCurve2D },
		func(p *PlantAbstract, shape *TopStackOfGrowthCurve2D) { p.TopStackOfGrowthCurve2D = shape },
		func(p *PlantAbstract, shape *TopStackOfGrowthCurve2D) bool {
			return p.TopStackOfGrowthCurve2D == shape
		},
		"TopStackOfGrowthCurve2D",
	)
	needCommit = n18 || needCommit

	n19 := enforcePlantHasShape[*StackOfGrowthCurve2DRibbon](
		stager,
		func() *StackOfGrowthCurve2DRibbon { return new(StackOfGrowthCurve2DRibbon) },
		func(p *PlantAbstract) *StackOfGrowthCurve2DRibbon { return p.StackOfGrowthCurve2DRibbon },
		func(p *PlantAbstract, shape *StackOfGrowthCurve2DRibbon) { p.StackOfGrowthCurve2DRibbon = shape },
		func(p *PlantAbstract, shape *StackOfGrowthCurve2DRibbon) bool {
			return p.StackOfGrowthCurve2DRibbon == shape
		},
		"StackOfGrowthCurve2DRibbon",
	)
	needCommit = n19 || needCommit

	n20 := enforcePlantHasShape[*StackOfRotatedGrowthCurve2DRibbon](
		stager,
		func() *StackOfRotatedGrowthCurve2DRibbon { return new(StackOfRotatedGrowthCurve2DRibbon) },
		func(p *PlantAbstract) *StackOfRotatedGrowthCurve2DRibbon { return p.StackOfRotatedGrowthCurve2DRibbon },
		func(p *PlantAbstract, shape *StackOfRotatedGrowthCurve2DRibbon) {
			p.StackOfRotatedGrowthCurve2DRibbon = shape
		},
		func(p *PlantAbstract, shape *StackOfRotatedGrowthCurve2DRibbon) bool {
			return p.StackOfRotatedGrowthCurve2DRibbon == shape
		},
		"StackOfRotatedGrowthCurve2DRibbon",
	)
	needCommit = n20 || needCommit

	n21 := enforcePlantHasShape[*PartiallyGrowthCurve2DRibbon](
		stager,
		func() *PartiallyGrowthCurve2DRibbon { return new(PartiallyGrowthCurve2DRibbon) },
		func(p *PlantAbstract) *PartiallyGrowthCurve2DRibbon { return p.PartiallyGrowthCurve2DRibbon },
		func(p *PlantAbstract, shape *PartiallyGrowthCurve2DRibbon) { p.PartiallyGrowthCurve2DRibbon = shape },
		func(p *PlantAbstract, shape *PartiallyGrowthCurve2DRibbon) bool {
			return p.PartiallyGrowthCurve2DRibbon == shape
		},
		"PartiallyGrowthCurve2DRibbon",
	)
	needCommit = n21 || needCommit

	n21_shiftedleft_partially := enforcePlantHasShape[*ShiftedLeftPartiallyGrowthCurve2DRibbon](
		stager,
		func() *ShiftedLeftPartiallyGrowthCurve2DRibbon { return new(ShiftedLeftPartiallyGrowthCurve2DRibbon) },
		func(p *PlantAbstract) *ShiftedLeftPartiallyGrowthCurve2DRibbon {
			return p.ShiftedLeftPartiallyGrowthCurve2DRibbon
		},
		func(p *PlantAbstract, shape *ShiftedLeftPartiallyGrowthCurve2DRibbon) {
			p.ShiftedLeftPartiallyGrowthCurve2DRibbon = shape
		},
		func(p *PlantAbstract, shape *ShiftedLeftPartiallyGrowthCurve2DRibbon) bool {
			return p.ShiftedLeftPartiallyGrowthCurve2DRibbon == shape
		},
		"ShiftedLeftPartiallyGrowthCurve2DRibbon",
	)
	needCommit = n21_shiftedleft_partially || needCommit

	n21_traj := enforcePlantHasShape[*PartiallyGrowthCurve2DTrajectory](
		stager,
		func() *PartiallyGrowthCurve2DTrajectory { return new(PartiallyGrowthCurve2DTrajectory) },
		func(p *PlantAbstract) *PartiallyGrowthCurve2DTrajectory { return p.PartiallyGrowthCurve2DTrajectory },
		func(p *PlantAbstract, shape *PartiallyGrowthCurve2DTrajectory) {
			p.PartiallyGrowthCurve2DTrajectory = shape
		},
		func(p *PlantAbstract, shape *PartiallyGrowthCurve2DTrajectory) bool {
			return p.PartiallyGrowthCurve2DTrajectory == shape
		},
		"PartiallyGrowthCurve2DTrajectory",
	)
	needCommit = n21_traj || needCommit

	n21_trajP1P2 := enforcePlantHasShape[*PartiallyGrowthCurve2DTrajectoryP1P2](
		stager,
		func() *PartiallyGrowthCurve2DTrajectoryP1P2 { return new(PartiallyGrowthCurve2DTrajectoryP1P2) },
		func(p *PlantAbstract) *PartiallyGrowthCurve2DTrajectoryP1P2 {
			return p.PartiallyGrowthCurve2DTrajectoryP1P2
		},
		func(p *PlantAbstract, shape *PartiallyGrowthCurve2DTrajectoryP1P2) {
			p.PartiallyGrowthCurve2DTrajectoryP1P2 = shape
		},
		func(p *PlantAbstract, shape *PartiallyGrowthCurve2DTrajectoryP1P2) bool {
			return p.PartiallyGrowthCurve2DTrajectoryP1P2 == shape
		},
		"PartiallyGrowthCurve2DTrajectoryP1P2",
	)
	needCommit = n21_trajP1P2 || needCommit

	n21_px := enforcePlantHasShape[*PxShape](
		stager,
		func() *PxShape { return new(PxShape) },
		func(p *PlantAbstract) *PxShape { return p.PxShape },
		func(p *PlantAbstract, shape *PxShape) { p.PxShape = shape },
		func(p *PlantAbstract, shape *PxShape) bool {
			return p.PxShape == shape
		},
		"PxShape",
	)
	needCommit = n21_px || needCommit

	n21_chosenP1P2 := enforcePlantHasShape[*ChosenP1P2PairShape](
		stager,
		func() *ChosenP1P2PairShape { return new(ChosenP1P2PairShape) },
		func(p *PlantAbstract) *ChosenP1P2PairShape { return p.ChosenP1P2PairShape },
		func(p *PlantAbstract, shape *ChosenP1P2PairShape) { p.ChosenP1P2PairShape = shape },
		func(p *PlantAbstract, shape *ChosenP1P2PairShape) bool {
			return p.ChosenP1P2PairShape == shape
		},
		"ChosenP1P2PairShape",
	)
	needCommit = n21_chosenP1P2 || needCommit

	n21_keyHole := enforcePlantHasShape[*KeyHoleShape](
		stager,
		func() *KeyHoleShape { return new(KeyHoleShape) },
		func(p *PlantAbstract) *KeyHoleShape { return p.KeyHoleShape },
		func(p *PlantAbstract, shape *KeyHoleShape) { p.KeyHoleShape = shape },
		func(p *PlantAbstract, shape *KeyHoleShape) bool {
			return p.KeyHoleShape == shape
		},
		"KeyHoleShape",
	)
	needCommit = n21_keyHole || needCommit

	n22 := enforcePlantHasShape[*GrowthCurve2DRibbon](
		stager,
		func() *GrowthCurve2DRibbon { return new(GrowthCurve2DRibbon) },
		func(p *PlantAbstract) *GrowthCurve2DRibbon { return p.GrowthCurve2DRibbon },
		func(p *PlantAbstract, shape *GrowthCurve2DRibbon) { p.GrowthCurve2DRibbon = shape },
		func(p *PlantAbstract, shape *GrowthCurve2DRibbon) bool {
			return p.GrowthCurve2DRibbon == shape
		},
		"GrowthCurve2DRibbon",
	)
	needCommit = n22 || needCommit

	n23 := enforcePlantHasShape[*ShiftedRightGrowthCurve2DRibbon](
		stager,
		func() *ShiftedRightGrowthCurve2DRibbon { return new(ShiftedRightGrowthCurve2DRibbon) },
		func(p *PlantAbstract) *ShiftedRightGrowthCurve2DRibbon { return p.ShiftedRightGrowthCurve2DRibbon },
		func(p *PlantAbstract, shape *ShiftedRightGrowthCurve2DRibbon) {
			p.ShiftedRightGrowthCurve2DRibbon = shape
		},
		func(p *PlantAbstract, shape *ShiftedRightGrowthCurve2DRibbon) bool {
			return p.ShiftedRightGrowthCurve2DRibbon == shape
		},
		"ShiftedRightGrowthCurve2DRibbon",
	)
	needCommit = n23 || needCommit

	n24 := enforcePlantHasShape[*ShiftedLeftGrowthCurve2DRibbon](
		stager,
		func() *ShiftedLeftGrowthCurve2DRibbon { return new(ShiftedLeftGrowthCurve2DRibbon) },
		func(p *PlantAbstract) *ShiftedLeftGrowthCurve2DRibbon { return p.ShiftedLeftGrowthCurve2DRibbon },
		func(p *PlantAbstract, shape *ShiftedLeftGrowthCurve2DRibbon) {
			p.ShiftedLeftGrowthCurve2DRibbon = shape
		},
		func(p *PlantAbstract, shape *ShiftedLeftGrowthCurve2DRibbon) bool {
			return p.ShiftedLeftGrowthCurve2DRibbon == shape
		},
		"ShiftedLeftGrowthCurve2DRibbon",
	)
	needCommit = n24 || needCommit

	return n1 || n2 || n3 || n4 || n5 || n6 || n7 || n7_halfway || n7_base || n7_arc_normal || n7_arc_v2 || n7_top_arc_v2 || n7_arc_v2_end || n7_top_arc_v2_end || n10 || n11 || n14 || n15 || n16 || n17 || n18 || n19 || n20 || n21 || n22 || n23 || n24
}

// enforceReferenceRhombusName ensures that the name of the ReferenceRhombus matches its owning Plant
func (stager *Stager) enforceReferenceRhombusName() (needCommit bool) {
	return enforcePlantShapeName[*RhombusShape](
		stager,
		func(p *PlantAbstract) *RhombusShape { return p.RhombusStuff.ReferenceRhombus },
		"ReferenceRhombus",
	)
}

// enforceInitialRhombusGridShapeName ensures that the name of the InitialRhombusGridShape matches its owning Plant
func (stager *Stager) enforceInitialRhombusGridShapeName() (needCommit bool) {
	return enforcePlantShapeName[*InitialRhombusGridShape](
		stager,
		func(p *PlantAbstract) *InitialRhombusGridShape { return p.RhombusStuff.InitialRhombusGridShape },
		"InitialRhombusGridShape",
	)
}

// enforceExplanationTextShapeName ensures that the name of the ExplanationTextShape matches its owning Plant
func (stager *Stager) enforceExplanationTextShapeName() (needCommit bool) {
	return enforcePlantShapeName[*ExplanationTextShape](
		stager,
		func(p *PlantAbstract) *ExplanationTextShape { return p.RhombusStuff.ExplanationTextShape },
		"ExplanationTextShape",
	)
}

// enforceRotatedShapesNames ensures that the name of the Rotated shapes match their owning Plant
func (stager *Stager) enforceRotatedShapesNames() (needCommit bool) {
	n1 := enforcePlantShapeName[*RhombusShape](
		stager,
		func(p *PlantAbstract) *RhombusShape { return p.RhombusStuff.RotatedReferenceRhombus },
		"RotatedReferenceRhombus",
	)

	n2 := enforcePlantShapeName[*PlantCircumferenceShape](
		stager,
		func(p *PlantAbstract) *PlantCircumferenceShape { return p.RhombusStuff.RotatedPlantCircumferenceShape },
		"RotatedPlantCircumferenceShape",
	)

	n3 := enforcePlantShapeName[*GridPathShape](
		stager,
		func(p *PlantAbstract) *GridPathShape { return p.RhombusStuff.RotatedGridPathShape },
		"RotatedGridPathShape",
	)

	n4 := enforcePlantShapeName[*RotatedRhombusGridShape](
		stager,
		func(p *PlantAbstract) *RotatedRhombusGridShape { return p.RhombusStuff.RotatedRhombusGridShape2 },
		"RotatedRhombusGridShape",
	)

	n5 := enforcePlantShapeName[*GrowthCurveRhombusGridShape](
		stager,
		func(p *PlantAbstract) *GrowthCurveRhombusGridShape { return p.RhombusStuff.GrowthCurveRhombusGridShape },
		"GrowthCurveRhombusGridShape",
	)

	n6 := enforcePlantShapeName[*GrowthVectorShape](
		stager,
		func(p *PlantAbstract) *GrowthVectorShape { return p.GrowthVectorShape },
		"GrowthVectorShape",
	)

	n7 := enforcePlantShapeName[*PerpendicularVectorGrid](
		stager,
		func(p *PlantAbstract) *PerpendicularVectorGrid { return p.PerpendicularVectorGrid },
		"PerpendicularVectorGrid",
	)
	needCommit = n7 || needCommit

	n7_halfway := enforcePlantShapeName[*PerpendicularVectorGridHalfway](
		stager,
		func(p *PlantAbstract) *PerpendicularVectorGridHalfway { return p.PerpendicularVectorGridHalfway },
		"PerpendicularVectorGridHalfway",
	)
	needCommit = n7_halfway || needCommit

	n7_base := enforcePlantShapeName[*BaseVectorShapeGrid](
		stager,
		func(p *PlantAbstract) *BaseVectorShapeGrid { return p.BaseVectorShapeGrid },
		"BaseVectorShapeGrid",
	)
	needCommit = n7_base || needCommit

	n7_arc_normal := enforcePlantShapeName[*ArcNormalVectorShapeGrid](
		stager,
		func(p *PlantAbstract) *ArcNormalVectorShapeGrid { return p.ArcNormalVectorShapeGrid },
		"ArcNormalVectorShapeGrid",
	)
	needCommit = n7_arc_normal || needCommit

	n7_arc_v2 := enforcePlantShapeName[*StartArcShapeGrid](
		stager,
		func(p *PlantAbstract) *StartArcShapeGrid { return p.StartArcShapeGrid },
		"StartArcShapeV2Grid",
	)
	needCommit = n7_arc_v2 || needCommit

	n7_top_arc_v2 := enforcePlantShapeName[*TopStartArcShapeGrid](
		stager,
		func(p *PlantAbstract) *TopStartArcShapeGrid { return p.TopStartArcShapeGrid },
		"TopStartArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2 || needCommit

	n7_arc_v2_end := enforcePlantShapeName[*EndArcShapeGrid](
		stager,
		func(p *PlantAbstract) *EndArcShapeGrid { return p.EndArcShapeGrid },
		"EndArcShapeV2Grid",
	)
	needCommit = n7_arc_v2_end || needCommit

	n7_top_arc_v2_end := enforcePlantShapeName[*TopEndArcShapeGrid](
		stager,
		func(p *PlantAbstract) *TopEndArcShapeGrid { return p.TopEndArcShapeGrid },
		"TopEndArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2_end || needCommit

	n10 := enforcePlantShapeName[*StackOfRotatedGrowthCurve2D](
		stager,
		func(p *PlantAbstract) *StackOfRotatedGrowthCurve2D { return p.StackOfRotatedGrowthCurve2D },
		"StackOfGrowthCurveV2",
	)
	needCommit = n10 || needCommit

	n11 := enforcePlantShapeName[*TopStackOfRotatedGrowthCurve2D](
		stager,
		func(p *PlantAbstract) *TopStackOfRotatedGrowthCurve2D { return p.TopStackOfRotatedGrowthCurve2D },
		"TopStackOfGrowthCurveV2",
	)
	needCommit = n11 || needCommit

	n12 := enforcePlantShapeName[*StackOfGrowthCurve2D](
		stager,
		func(p *PlantAbstract) *StackOfGrowthCurve2D { return p.StackOfGrowthCurve2D },
		"StackOfGrowthCurve2D",
	)
	needCommit = n12 || needCommit

	n13 := enforcePlantShapeName[*TopStackOfGrowthCurve2D](
		stager,
		func(p *PlantAbstract) *TopStackOfGrowthCurve2D { return p.TopStackOfGrowthCurve2D },
		"TopStackOfGrowthCurve2D",
	)
	needCommit = n13 || needCommit

	n14 := enforcePlantShapeName[*StackOfGrowthCurve2DRibbon](
		stager,
		func(p *PlantAbstract) *StackOfGrowthCurve2DRibbon { return p.StackOfGrowthCurve2DRibbon },
		"StackOfGrowthCurve2DRibbon",
	)
	needCommit = n14 || needCommit

	n15 := enforcePlantShapeName[*StackOfRotatedGrowthCurve2DRibbon](
		stager,
		func(p *PlantAbstract) *StackOfRotatedGrowthCurve2DRibbon { return p.StackOfRotatedGrowthCurve2DRibbon },
		"StackOfRotatedGrowthCurve2DRibbon",
	)
	needCommit = n15 || needCommit

	n16_r := enforcePlantShapeName[*GrowthCurve2DRibbon](
		stager,
		func(p *PlantAbstract) *GrowthCurve2DRibbon { return p.GrowthCurve2DRibbon },
		"GrowthCurve2DRibbon",
	)
	needCommit = n16_r || needCommit

	n17_r := enforcePlantShapeName[*ShiftedRightGrowthCurve2DRibbon](
		stager,
		func(p *PlantAbstract) *ShiftedRightGrowthCurve2DRibbon { return p.ShiftedRightGrowthCurve2DRibbon },
		"ShiftedRightGrowthCurve2DRibbon",
	)
	needCommit = n17_r || needCommit

	n18_r := enforcePlantShapeName[*ShiftedLeftGrowthCurve2DRibbon](
		stager,
		func(p *PlantAbstract) *ShiftedLeftGrowthCurve2DRibbon { return p.ShiftedLeftGrowthCurve2DRibbon },
		"ShiftedLeftGrowthCurve2DRibbon",
	)
	needCommit = n18_r || needCommit

	n19_r := enforcePlantShapeName[*ShiftedLeftPartiallyGrowthCurve2DRibbon](
		stager,
		func(p *PlantAbstract) *ShiftedLeftPartiallyGrowthCurve2DRibbon {
			return p.ShiftedLeftPartiallyGrowthCurve2DRibbon
		},
		"ShiftedLeftPartiallyGrowthCurve2DRibbon",
	)
	needCommit = n19_r || needCommit

	n20_keyHole := enforcePlantShapeName[*KeyHoleShape](
		stager,
		func(p *PlantAbstract) *KeyHoleShape { return p.KeyHoleShape },
		"KeyHoleShape",
	)
	needCommit = n20_keyHole || needCommit

	return n1 || n2 || n3 || n4 || n5 || n6 || n7 || n7_halfway || n7_base || n7_arc_normal || n7_arc_v2 || n7_top_arc_v2 || n7_arc_v2_end || n7_top_arc_v2_end || n10 || n11 || n12 || n13 || n14 || n15 || n16_r || n17_r || n18_r || n19_r || n20_keyHole

}

// enforcePlantRhombusGridShapeHasRhombuses ensures that each RhombusGridShape has the correct number of RhombusShapes and their X,Y fields are correctly computed
func isRhombusShapeOwnedByPlant(p *PlantAbstract, shape *RhombusShape) bool {
	if p.RhombusStuff.ReferenceRhombus == shape || p.RhombusStuff.RotatedReferenceRhombus == shape {
		return true
	}
	// Initial, Rotated and Growth grids no longer use generic RhombusShape.
	return false
}
