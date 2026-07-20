package models

import (
	"fmt"
	"time"
)

func enforcePlantHasShape[ShapePointerType PointerToGongstruct](
	stager *Stager,
	newShape func() ShapePointerType,
	getShape func(plant *Plant) ShapePointerType,
	setShape func(plant *Plant, shape ShapePointerType),
	isOwned func(plant *Plant, shape ShapePointerType) bool,
	shapeName string,
) (needCommit bool) {
	stage := stager.stage

	// 1. Ensure each Plant has the shape
	for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
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
		for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
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
	getShape func(plant *Plant) ShapePointerType,
	shapeNameSuffix string,
) (needCommit bool) {
	stage := stager.stage

	for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
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
		func(p *Plant) *AxesShape { return p.AxesShape },
		"AxesShape",
	)
}

// enforceGridPathShapeName ensures that the name of the GridPathShape matches its owning Plant
func (stager *Stager) enforceGridPathShapeName() (needCommit bool) {
	return enforcePlantShapeName[*GridPathShape](
		stager,
		func(p *Plant) *GridPathShape { return p.RhombusStuff.GridPathShape },
		"GridPathShape",
	)
}

// enforcePlantCircumferenceShapeName ensures that the name of the PlantCircumferenceShape matches its owning Plant
func (stager *Stager) enforcePlantCircumferenceShapeName() (needCommit bool) {
	return enforcePlantShapeName[*PlantCircumferenceShape](
		stager,
		func(p *Plant) *PlantCircumferenceShape { return p.RhombusStuff.PlantCircumferenceShape },
		"PlantCircumferenceShape",
	)
}

// enforcePlantHasAxes ensures that each Plant has one and only one Axes that belong to it
func (stager *Stager) enforcePlantHasAxes() (needCommit bool) {
	return enforcePlantHasShape[*AxesShape](
		stager,
		func() *AxesShape { return new(AxesShape) },
		func(p *Plant) *AxesShape { return p.AxesShape },
		func(p *Plant, shape *AxesShape) { p.AxesShape = shape },
		func(p *Plant, shape *AxesShape) bool { return p.AxesShape == shape },
		"AxesShape",
	)
}

// enforcePlantHasGridPathShape ensures that each Plant has one and only one GridPathShape that belong to it
func (stager *Stager) enforcePlantHasGridPathShape() (needCommit bool) {
	return enforcePlantHasShape[*GridPathShape](
		stager,
		func() *GridPathShape { return new(GridPathShape) },
		func(p *Plant) *GridPathShape { return p.RhombusStuff.GridPathShape },
		func(p *Plant, shape *GridPathShape) { p.RhombusStuff.GridPathShape = shape },
		func(p *Plant, shape *GridPathShape) bool {
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
		func(p *Plant) *InitialRhombusGridShape { return p.RhombusStuff.InitialRhombusGridShape },
		func(p *Plant, shape *InitialRhombusGridShape) { p.RhombusStuff.InitialRhombusGridShape = shape },
		func(p *Plant, shape *InitialRhombusGridShape) bool {
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
		func(p *Plant) *ExplanationTextShape { return p.RhombusStuff.ExplanationTextShape },
		func(p *Plant, shape *ExplanationTextShape) { p.RhombusStuff.ExplanationTextShape = shape },
		func(p *Plant, shape *ExplanationTextShape) bool {
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
		func(p *Plant) *PlantCircumferenceShape { return p.RhombusStuff.PlantCircumferenceShape },
		func(p *Plant, shape *PlantCircumferenceShape) { p.RhombusStuff.PlantCircumferenceShape = shape },
		func(p *Plant, shape *PlantCircumferenceShape) bool {
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
		func(p *Plant) *RhombusShape { return p.RhombusStuff.ReferenceRhombus },
		func(p *Plant, shape *RhombusShape) { p.RhombusStuff.ReferenceRhombus = shape },
		func(p *Plant, shape *RhombusShape) bool {
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
		func(p *Plant) *RhombusShape { return p.RhombusStuff.RotatedReferenceRhombus },
		func(p *Plant, shape *RhombusShape) { p.RhombusStuff.RotatedReferenceRhombus = shape },
		func(p *Plant, shape *RhombusShape) bool {
			return isRhombusShapeOwnedByPlant(p, shape)
		},
		"RotatedReferenceRhombus",
	)

	n2 := enforcePlantHasShape[*PlantCircumferenceShape](
		stager,
		func() *PlantCircumferenceShape { return new(PlantCircumferenceShape) },
		func(p *Plant) *PlantCircumferenceShape { return p.RhombusStuff.RotatedPlantCircumferenceShape },
		func(p *Plant, shape *PlantCircumferenceShape) { p.RhombusStuff.RotatedPlantCircumferenceShape = shape },
		func(p *Plant, shape *PlantCircumferenceShape) bool {
			return p.RhombusStuff.PlantCircumferenceShape == shape || p.RhombusStuff.RotatedPlantCircumferenceShape == shape
		},
		"RotatedPlantCircumferenceShape",
	)

	n3 := enforcePlantHasShape[*GridPathShape](
		stager,
		func() *GridPathShape { return new(GridPathShape) },
		func(p *Plant) *GridPathShape { return p.RhombusStuff.RotatedGridPathShape },
		func(p *Plant, shape *GridPathShape) { p.RhombusStuff.RotatedGridPathShape = shape },
		func(p *Plant, shape *GridPathShape) bool {
			return p.RhombusStuff.GridPathShape == shape || p.RhombusStuff.RotatedGridPathShape == shape
		},
		"RotatedGridPathShape",
	)

	n4 := enforcePlantHasShape[*RotatedRhombusGridShape](
		stager,
		func() *RotatedRhombusGridShape { return new(RotatedRhombusGridShape) },
		func(p *Plant) *RotatedRhombusGridShape { return p.RhombusStuff.RotatedRhombusGridShape2 },
		func(p *Plant, shape *RotatedRhombusGridShape) { p.RhombusStuff.RotatedRhombusGridShape2 = shape },
		func(p *Plant, shape *RotatedRhombusGridShape) bool {
			return p.RhombusStuff.RotatedRhombusGridShape2 == shape
		},
		"RotatedRhombusGridShape",
	)

	n5 := enforcePlantHasShape[*GrowthCurveRhombusGridShape](
		stager,
		func() *GrowthCurveRhombusGridShape { return new(GrowthCurveRhombusGridShape) },
		func(p *Plant) *GrowthCurveRhombusGridShape { return p.RhombusStuff.GrowthCurveRhombusGridShape },
		func(p *Plant, shape *GrowthCurveRhombusGridShape) { p.RhombusStuff.GrowthCurveRhombusGridShape = shape },
		func(p *Plant, shape *GrowthCurveRhombusGridShape) bool {
			return p.RhombusStuff.GrowthCurveRhombusGridShape == shape
		},
		"GrowthCurveRhombusGridShape",
	)

	n6 := enforcePlantHasShape[*GrowthVectorShape](
		stager,
		func() *GrowthVectorShape { return new(GrowthVectorShape) },
		func(p *Plant) *GrowthVectorShape { return p.GrowthVectorShape },
		func(p *Plant, shape *GrowthVectorShape) { p.GrowthVectorShape = shape },
		func(p *Plant, shape *GrowthVectorShape) bool {
			return p.GrowthVectorShape == shape
		},
		"GrowthVectorShape",
	)

	n7 := enforcePlantHasShape[*PerpendicularVectorGrid](
		stager,
		func() *PerpendicularVectorGrid { return new(PerpendicularVectorGrid) },
		func(p *Plant) *PerpendicularVectorGrid { return p.PerpendicularVectorGrid },
		func(p *Plant, shape *PerpendicularVectorGrid) { p.PerpendicularVectorGrid = shape },
		func(p *Plant, shape *PerpendicularVectorGrid) bool {
			return p.PerpendicularVectorGrid == shape
		},
		"PerpendicularVectorGrid",
	)
	needCommit = n7 || needCommit

	n7_halfway := enforcePlantHasShape[*PerpendicularVectorGridHalfway](
		stager,
		func() *PerpendicularVectorGridHalfway { return new(PerpendicularVectorGridHalfway) },
		func(p *Plant) *PerpendicularVectorGridHalfway { return p.PerpendicularVectorGridHalfway },
		func(p *Plant, shape *PerpendicularVectorGridHalfway) { p.PerpendicularVectorGridHalfway = shape },
		func(p *Plant, shape *PerpendicularVectorGridHalfway) bool {
			return p.PerpendicularVectorGridHalfway == shape
		},
		"PerpendicularVectorGridHalfway",
	)
	needCommit = n7_halfway || needCommit

	n7_base := enforcePlantHasShape[*BaseVectorShapeGrid](
		stager,
		func() *BaseVectorShapeGrid { return new(BaseVectorShapeGrid) },
		func(p *Plant) *BaseVectorShapeGrid { return p.BaseVectorShapeGrid },
		func(p *Plant, shape *BaseVectorShapeGrid) { p.BaseVectorShapeGrid = shape },
		func(p *Plant, shape *BaseVectorShapeGrid) bool {
			return p.BaseVectorShapeGrid == shape
		},
		"BaseVectorShapeGrid",
	)
	needCommit = n7_base || needCommit

	n7_arc_normal := enforcePlantHasShape[*ArcNormalVectorShapeGrid](
		stager,
		func() *ArcNormalVectorShapeGrid { return new(ArcNormalVectorShapeGrid) },
		func(p *Plant) *ArcNormalVectorShapeGrid { return p.ArcNormalVectorShapeGrid },
		func(p *Plant, shape *ArcNormalVectorShapeGrid) { p.ArcNormalVectorShapeGrid = shape },
		func(p *Plant, shape *ArcNormalVectorShapeGrid) bool {
			return p.ArcNormalVectorShapeGrid == shape
		},
		"ArcNormalVectorShapeGrid",
	)
	needCommit = n7_arc_normal || needCommit

	n7_arc_v2 := enforcePlantHasShape[*StartArcShapeGrid](
		stager,
		func() *StartArcShapeGrid { return new(StartArcShapeGrid) },
		func(p *Plant) *StartArcShapeGrid { return p.StartArcShapeGrid },
		func(p *Plant, shape *StartArcShapeGrid) { p.StartArcShapeGrid = shape },
		func(p *Plant, shape *StartArcShapeGrid) bool {
			return p.StartArcShapeGrid == shape
		},
		"StartArcShapeV2Grid",
	)
	needCommit = n7_arc_v2 || needCommit

	n7_top_arc_v2 := enforcePlantHasShape[*TopStartArcShapeGrid](
		stager,
		func() *TopStartArcShapeGrid { return new(TopStartArcShapeGrid) },
		func(p *Plant) *TopStartArcShapeGrid { return p.TopStartArcShapeGrid },
		func(p *Plant, shape *TopStartArcShapeGrid) { p.TopStartArcShapeGrid = shape },
		func(p *Plant, shape *TopStartArcShapeGrid) bool {
			return p.TopStartArcShapeGrid == shape
		},
		"TopStartArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2 || needCommit

	n7_arc_v2_end := enforcePlantHasShape[*EndArcShapeGrid](
		stager,
		func() *EndArcShapeGrid { return new(EndArcShapeGrid) },
		func(p *Plant) *EndArcShapeGrid { return p.EndArcShapeGrid },
		func(p *Plant, shape *EndArcShapeGrid) { p.EndArcShapeGrid = shape },
		func(p *Plant, shape *EndArcShapeGrid) bool {
			return p.EndArcShapeGrid == shape
		},
		"EndArcShapeV2Grid",
	)
	needCommit = n7_arc_v2_end || needCommit

	n7_top_arc_v2_end := enforcePlantHasShape[*TopEndArcShapeGrid](
		stager,
		func() *TopEndArcShapeGrid { return new(TopEndArcShapeGrid) },
		func(p *Plant) *TopEndArcShapeGrid { return p.TopEndArcShapeGrid },
		func(p *Plant, shape *TopEndArcShapeGrid) { p.TopEndArcShapeGrid = shape },
		func(p *Plant, shape *TopEndArcShapeGrid) bool {
			return p.TopEndArcShapeGrid == shape
		},
		"TopEndArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2_end || needCommit

	n10 := enforcePlantHasShape[*StackOfRotatedGrowthCurve2D](
		stager,
		func() *StackOfRotatedGrowthCurve2D { return new(StackOfRotatedGrowthCurve2D) },
		func(p *Plant) *StackOfRotatedGrowthCurve2D { return p.StackOfRotatedGrowthCurve2D },
		func(p *Plant, shape *StackOfRotatedGrowthCurve2D) { p.StackOfRotatedGrowthCurve2D = shape },
		func(p *Plant, shape *StackOfRotatedGrowthCurve2D) bool {
			return p.StackOfRotatedGrowthCurve2D == shape
		},
		"StackOfGrowthCurveV2",
	)
	needCommit = n10 || needCommit

	n11 := enforcePlantHasShape[*TopStackOfRotatedGrowthCurve2D](
		stager,
		func() *TopStackOfRotatedGrowthCurve2D { return new(TopStackOfRotatedGrowthCurve2D) },
		func(p *Plant) *TopStackOfRotatedGrowthCurve2D { return p.TopStackOfRotatedGrowthCurve2D },
		func(p *Plant, shape *TopStackOfRotatedGrowthCurve2D) { p.TopStackOfRotatedGrowthCurve2D = shape },
		func(p *Plant, shape *TopStackOfRotatedGrowthCurve2D) bool {
			return p.TopStackOfRotatedGrowthCurve2D == shape
		},
		"TopStackOfGrowthCurveV2",
	)
	needCommit = n11 || needCommit

	n15 := enforcePlantHasShape[*MidArcVectorShapeGrid](
		stager,
		func() *MidArcVectorShapeGrid { return new(MidArcVectorShapeGrid) },
		func(p *Plant) *MidArcVectorShapeGrid { return p.MidArcVectorShapeGrid },
		func(p *Plant, shape *MidArcVectorShapeGrid) { p.MidArcVectorShapeGrid = shape },
		func(p *Plant, shape *MidArcVectorShapeGrid) bool {
			return p.MidArcVectorShapeGrid == shape
		},
		"MidArcVectorShapeGrid",
	)
	needCommit = n15 || needCommit

	n16 := enforcePlantHasShape[*TopMidArcVectorShapeGrid](
		stager,
		func() *TopMidArcVectorShapeGrid { return new(TopMidArcVectorShapeGrid) },
		func(p *Plant) *TopMidArcVectorShapeGrid { return p.TopMidArcVectorShapeGrid },
		func(p *Plant, shape *TopMidArcVectorShapeGrid) { p.TopMidArcVectorShapeGrid = shape },
		func(p *Plant, shape *TopMidArcVectorShapeGrid) bool {
			return p.TopMidArcVectorShapeGrid == shape
		},
		"TopMidArcVectorShapeGrid",
	)
	needCommit = n16 || needCommit

	n14 := enforcePlantHasShape[*ShiftedBottomTopStartArcShapeGrid](
		stager,
		func() *ShiftedBottomTopStartArcShapeGrid { return new(ShiftedBottomTopStartArcShapeGrid) },
		func(p *Plant) *ShiftedBottomTopStartArcShapeGrid { return p.ShiftedBottomTopStartArcShapeGrid },
		func(p *Plant, shape *ShiftedBottomTopStartArcShapeGrid) { p.ShiftedBottomTopStartArcShapeGrid = shape },
		func(p *Plant, shape *ShiftedBottomTopStartArcShapeGrid) bool {
			return p.ShiftedBottomTopStartArcShapeGrid == shape
		},
		"ShiftedBottomTopStartArcShapeGrid",
	)
	needCommit = n14 || needCommit

	n17 := enforcePlantHasShape[*StackOfGrowthCurve2D](
		stager,
		func() *StackOfGrowthCurve2D { return new(StackOfGrowthCurve2D) },
		func(p *Plant) *StackOfGrowthCurve2D { return p.StackOfGrowthCurve2D },
		func(p *Plant, shape *StackOfGrowthCurve2D) { p.StackOfGrowthCurve2D = shape },
		func(p *Plant, shape *StackOfGrowthCurve2D) bool {
			return p.StackOfGrowthCurve2D == shape
		},
		"StackOfGrowthCurve2D",
	)
	needCommit = n17 || needCommit

	n18 := enforcePlantHasShape[*TopStackOfGrowthCurve2D](
		stager,
		func() *TopStackOfGrowthCurve2D { return new(TopStackOfGrowthCurve2D) },
		func(p *Plant) *TopStackOfGrowthCurve2D { return p.TopStackOfGrowthCurve2D },
		func(p *Plant, shape *TopStackOfGrowthCurve2D) { p.TopStackOfGrowthCurve2D = shape },
		func(p *Plant, shape *TopStackOfGrowthCurve2D) bool {
			return p.TopStackOfGrowthCurve2D == shape
		},
		"TopStackOfGrowthCurve2D",
	)
	needCommit = n18 || needCommit

	n19 := enforcePlantHasShape[*StackOfGrowthCurve2DRibbon](
		stager,
		func() *StackOfGrowthCurve2DRibbon { return new(StackOfGrowthCurve2DRibbon) },
		func(p *Plant) *StackOfGrowthCurve2DRibbon { return p.StackOfGrowthCurve2DRibbon },
		func(p *Plant, shape *StackOfGrowthCurve2DRibbon) { p.StackOfGrowthCurve2DRibbon = shape },
		func(p *Plant, shape *StackOfGrowthCurve2DRibbon) bool {
			return p.StackOfGrowthCurve2DRibbon == shape
		},
		"StackOfGrowthCurve2DRibbon",
	)
	needCommit = n19 || needCommit

	return n1 || n2 || n3 || n4 || n5 || n6 || n7 || n7_halfway || n7_base || n7_arc_normal || n7_arc_v2 || n7_top_arc_v2 || n7_arc_v2_end || n7_top_arc_v2_end || n10 || n11 || n14 || n15 || n16 || n17 || n18 || n19
}

// enforceReferenceRhombusName ensures that the name of the ReferenceRhombus matches its owning Plant
func (stager *Stager) enforceReferenceRhombusName() (needCommit bool) {
	return enforcePlantShapeName[*RhombusShape](
		stager,
		func(p *Plant) *RhombusShape { return p.RhombusStuff.ReferenceRhombus },
		"ReferenceRhombus",
	)
}

// enforceInitialRhombusGridShapeName ensures that the name of the InitialRhombusGridShape matches its owning Plant
func (stager *Stager) enforceInitialRhombusGridShapeName() (needCommit bool) {
	return enforcePlantShapeName[*InitialRhombusGridShape](
		stager,
		func(p *Plant) *InitialRhombusGridShape { return p.RhombusStuff.InitialRhombusGridShape },
		"InitialRhombusGridShape",
	)
}

// enforceExplanationTextShapeName ensures that the name of the ExplanationTextShape matches its owning Plant
func (stager *Stager) enforceExplanationTextShapeName() (needCommit bool) {
	return enforcePlantShapeName[*ExplanationTextShape](
		stager,
		func(p *Plant) *ExplanationTextShape { return p.RhombusStuff.ExplanationTextShape },
		"ExplanationTextShape",
	)
}

// enforceRotatedShapesNames ensures that the name of the Rotated shapes match their owning Plant
func (stager *Stager) enforceRotatedShapesNames() (needCommit bool) {
	n1 := enforcePlantShapeName[*RhombusShape](
		stager,
		func(p *Plant) *RhombusShape { return p.RhombusStuff.RotatedReferenceRhombus },
		"RotatedReferenceRhombus",
	)

	n2 := enforcePlantShapeName[*PlantCircumferenceShape](
		stager,
		func(p *Plant) *PlantCircumferenceShape { return p.RhombusStuff.RotatedPlantCircumferenceShape },
		"RotatedPlantCircumferenceShape",
	)

	n3 := enforcePlantShapeName[*GridPathShape](
		stager,
		func(p *Plant) *GridPathShape { return p.RhombusStuff.RotatedGridPathShape },
		"RotatedGridPathShape",
	)

	n4 := enforcePlantShapeName[*RotatedRhombusGridShape](
		stager,
		func(p *Plant) *RotatedRhombusGridShape { return p.RhombusStuff.RotatedRhombusGridShape2 },
		"RotatedRhombusGridShape",
	)

	n5 := enforcePlantShapeName[*GrowthCurveRhombusGridShape](
		stager,
		func(p *Plant) *GrowthCurveRhombusGridShape { return p.RhombusStuff.GrowthCurveRhombusGridShape },
		"GrowthCurveRhombusGridShape",
	)

	n6 := enforcePlantShapeName[*GrowthVectorShape](
		stager,
		func(p *Plant) *GrowthVectorShape { return p.GrowthVectorShape },
		"GrowthVectorShape",
	)

	n7 := enforcePlantShapeName[*PerpendicularVectorGrid](
		stager,
		func(p *Plant) *PerpendicularVectorGrid { return p.PerpendicularVectorGrid },
		"PerpendicularVectorGrid",
	)
	needCommit = n7 || needCommit

	n7_halfway := enforcePlantShapeName[*PerpendicularVectorGridHalfway](
		stager,
		func(p *Plant) *PerpendicularVectorGridHalfway { return p.PerpendicularVectorGridHalfway },
		"PerpendicularVectorGridHalfway",
	)
	needCommit = n7_halfway || needCommit

	n7_base := enforcePlantShapeName[*BaseVectorShapeGrid](
		stager,
		func(p *Plant) *BaseVectorShapeGrid { return p.BaseVectorShapeGrid },
		"BaseVectorShapeGrid",
	)
	needCommit = n7_base || needCommit

	n7_arc_normal := enforcePlantShapeName[*ArcNormalVectorShapeGrid](
		stager,
		func(p *Plant) *ArcNormalVectorShapeGrid { return p.ArcNormalVectorShapeGrid },
		"ArcNormalVectorShapeGrid",
	)
	needCommit = n7_arc_normal || needCommit

	n7_arc_v2 := enforcePlantShapeName[*StartArcShapeGrid](
		stager,
		func(p *Plant) *StartArcShapeGrid { return p.StartArcShapeGrid },
		"StartArcShapeV2Grid",
	)
	needCommit = n7_arc_v2 || needCommit

	n7_top_arc_v2 := enforcePlantShapeName[*TopStartArcShapeGrid](
		stager,
		func(p *Plant) *TopStartArcShapeGrid { return p.TopStartArcShapeGrid },
		"TopStartArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2 || needCommit

	n7_arc_v2_end := enforcePlantShapeName[*EndArcShapeGrid](
		stager,
		func(p *Plant) *EndArcShapeGrid { return p.EndArcShapeGrid },
		"EndArcShapeV2Grid",
	)
	needCommit = n7_arc_v2_end || needCommit

	n7_top_arc_v2_end := enforcePlantShapeName[*TopEndArcShapeGrid](
		stager,
		func(p *Plant) *TopEndArcShapeGrid { return p.TopEndArcShapeGrid },
		"TopEndArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2_end || needCommit

	n10 := enforcePlantShapeName[*StackOfRotatedGrowthCurve2D](
		stager,
		func(p *Plant) *StackOfRotatedGrowthCurve2D { return p.StackOfRotatedGrowthCurve2D },
		"StackOfGrowthCurveV2",
	)
	needCommit = n10 || needCommit

	n11 := enforcePlantShapeName[*TopStackOfRotatedGrowthCurve2D](
		stager,
		func(p *Plant) *TopStackOfRotatedGrowthCurve2D { return p.TopStackOfRotatedGrowthCurve2D },
		"TopStackOfGrowthCurveV2",
	)
	needCommit = n11 || needCommit

	n12 := enforcePlantShapeName[*StackOfGrowthCurve2D](
		stager,
		func(p *Plant) *StackOfGrowthCurve2D { return p.StackOfGrowthCurve2D },
		"StackOfGrowthCurve2D",
	)
	needCommit = n12 || needCommit

	n13 := enforcePlantShapeName[*TopStackOfGrowthCurve2D](
		stager,
		func(p *Plant) *TopStackOfGrowthCurve2D { return p.TopStackOfGrowthCurve2D },
		"TopStackOfGrowthCurve2D",
	)
	needCommit = n13 || needCommit

	n14 := enforcePlantShapeName[*StackOfGrowthCurve2DRibbon](
		stager,
		func(p *Plant) *StackOfGrowthCurve2DRibbon { return p.StackOfGrowthCurve2DRibbon },
		"StackOfGrowthCurve2DRibbon",
	)
	needCommit = n14 || needCommit

	return n1 || n2 || n3 || n4 || n5 || n6 || n7 || n7_halfway || n7_base || n7_arc_normal || n7_arc_v2 || n7_top_arc_v2 || n7_arc_v2_end || n7_top_arc_v2_end || n10 || n11 || n12 || n13 || n14
}

// enforcePlantRhombusGridShapeHasRhombuses ensures that each RhombusGridShape has the correct number of RhombusShapes and their X,Y fields are correctly computed
func isRhombusShapeOwnedByPlant(p *Plant, shape *RhombusShape) bool {
	if p.RhombusStuff.ReferenceRhombus == shape || p.RhombusStuff.RotatedReferenceRhombus == shape {
		return true
	}
	// Initial, Rotated and Growth grids no longer use generic RhombusShape.
	return false
}
