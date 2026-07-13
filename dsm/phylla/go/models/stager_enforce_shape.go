package models

import (
	"fmt"
	"math"
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
		func(p *Plant) *GridPathShape { return p.GridPathShape },
		"GridPathShape",
	)
}

// enforcePlantCircumferenceShapeName ensures that the name of the PlantCircumferenceShape matches its owning Plant
func (stager *Stager) enforcePlantCircumferenceShapeName() (needCommit bool) {
	return enforcePlantShapeName[*PlantCircumferenceShape](
		stager,
		func(p *Plant) *PlantCircumferenceShape { return p.PlantCircumferenceShape },
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
		func(p *Plant) *GridPathShape { return p.GridPathShape },
		func(p *Plant, shape *GridPathShape) { p.GridPathShape = shape },
		func(p *Plant, shape *GridPathShape) bool {
			return p.GridPathShape == shape || p.RotatedGridPathShape == shape
		},
		"GridPathShape",
	)
}

// enforcePlantHasInitialRhombusGridShape ensures that each Plant has one and only one InitialRhombusGridShape that belong to it
func (stager *Stager) enforcePlantHasInitialRhombusGridShape() (needCommit bool) {
	return enforcePlantHasShape[*InitialRhombusGridShape](
		stager,
		func() *InitialRhombusGridShape { return new(InitialRhombusGridShape) },
		func(p *Plant) *InitialRhombusGridShape { return p.InitialRhombusGridShape },
		func(p *Plant, shape *InitialRhombusGridShape) { p.InitialRhombusGridShape = shape },
		func(p *Plant, shape *InitialRhombusGridShape) bool {
			return p.InitialRhombusGridShape == shape
		},
		"InitialRhombusGridShape",
	)
}

// enforcePlantHasExplanationTextShape ensures that each Plant has one and only one ExplanationTextShape that belong to it
func (stager *Stager) enforcePlantHasExplanationTextShape() (needCommit bool) {
	return enforcePlantHasShape[*ExplanationTextShape](
		stager,
		func() *ExplanationTextShape { return new(ExplanationTextShape) },
		func(p *Plant) *ExplanationTextShape { return p.ExplanationTextShape },
		func(p *Plant, shape *ExplanationTextShape) { p.ExplanationTextShape = shape },
		func(p *Plant, shape *ExplanationTextShape) bool {
			return p.ExplanationTextShape == shape
		},
		"ExplanationTextShape",
	)
}

// enforcePlantHasPlantCircumferenceShape ensures that each Plant has one and only one PlantCircumferenceShape that belong to it
func (stager *Stager) enforcePlantHasPlantCircumferenceShape() (needCommit bool) {
	return enforcePlantHasShape[*PlantCircumferenceShape](
		stager,
		func() *PlantCircumferenceShape { return new(PlantCircumferenceShape) },
		func(p *Plant) *PlantCircumferenceShape { return p.PlantCircumferenceShape },
		func(p *Plant, shape *PlantCircumferenceShape) { p.PlantCircumferenceShape = shape },
		func(p *Plant, shape *PlantCircumferenceShape) bool {
			return p.PlantCircumferenceShape == shape || p.RotatedPlantCircumferenceShape == shape
		},
		"PlantCircumferenceShape",
	)
}

// enforcePlantHasReferenceRhombus ensures that each Plant has one and only one ReferenceRhombus that belong to it
func (stager *Stager) enforcePlantHasReferenceRhombus() (needCommit bool) {
	return enforcePlantHasShape[*RhombusShape](
		stager,
		func() *RhombusShape { return new(RhombusShape) },
		func(p *Plant) *RhombusShape { return p.ReferenceRhombus },
		func(p *Plant, shape *RhombusShape) { p.ReferenceRhombus = shape },
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
		func(p *Plant) *RhombusShape { return p.RotatedReferenceRhombus },
		func(p *Plant, shape *RhombusShape) { p.RotatedReferenceRhombus = shape },
		func(p *Plant, shape *RhombusShape) bool {
			return isRhombusShapeOwnedByPlant(p, shape)
		},
		"RotatedReferenceRhombus",
	)

	n2 := enforcePlantHasShape[*PlantCircumferenceShape](
		stager,
		func() *PlantCircumferenceShape { return new(PlantCircumferenceShape) },
		func(p *Plant) *PlantCircumferenceShape { return p.RotatedPlantCircumferenceShape },
		func(p *Plant, shape *PlantCircumferenceShape) { p.RotatedPlantCircumferenceShape = shape },
		func(p *Plant, shape *PlantCircumferenceShape) bool {
			return p.PlantCircumferenceShape == shape || p.RotatedPlantCircumferenceShape == shape
		},
		"RotatedPlantCircumferenceShape",
	)

	n3 := enforcePlantHasShape[*GridPathShape](
		stager,
		func() *GridPathShape { return new(GridPathShape) },
		func(p *Plant) *GridPathShape { return p.RotatedGridPathShape },
		func(p *Plant, shape *GridPathShape) { p.RotatedGridPathShape = shape },
		func(p *Plant, shape *GridPathShape) bool {
			return p.GridPathShape == shape || p.RotatedGridPathShape == shape
		},
		"RotatedGridPathShape",
	)

	n4 := enforcePlantHasShape[*RotatedRhombusGridShape](
		stager,
		func() *RotatedRhombusGridShape { return new(RotatedRhombusGridShape) },
		func(p *Plant) *RotatedRhombusGridShape { return p.RotatedRhombusGridShape2 },
		func(p *Plant, shape *RotatedRhombusGridShape) { p.RotatedRhombusGridShape2 = shape },
		func(p *Plant, shape *RotatedRhombusGridShape) bool {
			return p.RotatedRhombusGridShape2 == shape
		},
		"RotatedRhombusGridShape",
	)

	n5 := enforcePlantHasShape[*GrowthCurveRhombusGridShape](
		stager,
		func() *GrowthCurveRhombusGridShape { return new(GrowthCurveRhombusGridShape) },
		func(p *Plant) *GrowthCurveRhombusGridShape { return p.GrowthCurveRhombusGridShape },
		func(p *Plant, shape *GrowthCurveRhombusGridShape) { p.GrowthCurveRhombusGridShape = shape },
		func(p *Plant, shape *GrowthCurveRhombusGridShape) bool {
			return p.GrowthCurveRhombusGridShape == shape
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

	n7_bottom_arc_v2 := enforcePlantHasShape[*BottomStartArcShapeGrid](
		stager,
		func() *BottomStartArcShapeGrid { return new(BottomStartArcShapeGrid) },
		func(p *Plant) *BottomStartArcShapeGrid { return p.BottomStartArcShapeGrid },
		func(p *Plant, shape *BottomStartArcShapeGrid) { p.BottomStartArcShapeGrid = shape },
		func(p *Plant, shape *BottomStartArcShapeGrid) bool {
			return p.BottomStartArcShapeGrid == shape
		},
		"BottomStartArcShapeV2Grid",
	)
	needCommit = n7_bottom_arc_v2 || needCommit

	n7_bottom_arc_v2_end := enforcePlantHasShape[*BottomEndArcShapeGrid](
		stager,
		func() *BottomEndArcShapeGrid { return new(BottomEndArcShapeGrid) },
		func(p *Plant) *BottomEndArcShapeGrid { return p.BottomEndArcShapeGrid },
		func(p *Plant, shape *BottomEndArcShapeGrid) { p.BottomEndArcShapeGrid = shape },
		func(p *Plant, shape *BottomEndArcShapeGrid) bool {
			return p.BottomEndArcShapeGrid == shape
		},
		"BottomEndArcShapeV2Grid",
	)
	needCommit = n7_bottom_arc_v2_end || needCommit

	n8 := enforcePlantHasShape[*GrowthCurveBezierShapeGrid](
		stager,
		func() *GrowthCurveBezierShapeGrid { return new(GrowthCurveBezierShapeGrid) },
		func(p *Plant) *GrowthCurveBezierShapeGrid { return p.GrowthCurveBezierShapeGrid },
		func(p *Plant, shape *GrowthCurveBezierShapeGrid) { p.GrowthCurveBezierShapeGrid = shape },
		func(p *Plant, shape *GrowthCurveBezierShapeGrid) bool {
			return p.GrowthCurveBezierShapeGrid == shape
		},
		"GrowthCurveBezierShapeGrid",
	)
	needCommit = n8 || needCommit

	n10 := enforcePlantHasShape[*StackOfGrowthCurve](
		stager,
		func() *StackOfGrowthCurve { return new(StackOfGrowthCurve) },
		func(p *Plant) *StackOfGrowthCurve { return p.StackOfGrowthCurve },
		func(p *Plant, shape *StackOfGrowthCurve) { p.StackOfGrowthCurve = shape },
		func(p *Plant, shape *StackOfGrowthCurve) bool {
			return p.StackOfGrowthCurve == shape
		},
		"StackOfGrowthCurveV2",
	)
	needCommit = n10 || needCommit

	n11 := enforcePlantHasShape[*TopStackOfGrowthCurve](
		stager,
		func() *TopStackOfGrowthCurve { return new(TopStackOfGrowthCurve) },
		func(p *Plant) *TopStackOfGrowthCurve { return p.TopStackOfGrowthCurve },
		func(p *Plant, shape *TopStackOfGrowthCurve) { p.TopStackOfGrowthCurve = shape },
		func(p *Plant, shape *TopStackOfGrowthCurve) bool {
			return p.TopStackOfGrowthCurve == shape
		},
		"TopStackOfGrowthCurveV2",
	)
	needCommit = n11 || needCommit

	n12 := enforcePlantHasShape[*BottomStackOfGrowthCurve](
		stager,
		func() *BottomStackOfGrowthCurve { return new(BottomStackOfGrowthCurve) },
		func(p *Plant) *BottomStackOfGrowthCurve { return p.BottomStackOfGrowthCurve },
		func(p *Plant, shape *BottomStackOfGrowthCurve) { p.BottomStackOfGrowthCurve = shape },
		func(p *Plant, shape *BottomStackOfGrowthCurve) bool {
			return p.BottomStackOfGrowthCurve == shape
		},
		"BottomStackOfGrowthCurveV2",
	)
	needCommit = n12 || needCommit

	return n1 || n2 || n3 || n4 || n5 || n6 || n7 || n7_halfway || n7_base || n7_arc_normal || n7_arc_v2 || n7_top_arc_v2 || n7_arc_v2_end || n7_top_arc_v2_end || n7_bottom_arc_v2 || n7_bottom_arc_v2_end || n8 || n10 || n11 || n12
}

// enforceReferenceRhombusName ensures that the name of the ReferenceRhombus matches its owning Plant
func (stager *Stager) enforceReferenceRhombusName() (needCommit bool) {
	return enforcePlantShapeName[*RhombusShape](
		stager,
		func(p *Plant) *RhombusShape { return p.ReferenceRhombus },
		"ReferenceRhombus",
	)
}

// enforceInitialRhombusGridShapeName ensures that the name of the InitialRhombusGridShape matches its owning Plant
func (stager *Stager) enforceInitialRhombusGridShapeName() (needCommit bool) {
	return enforcePlantShapeName[*InitialRhombusGridShape](
		stager,
		func(p *Plant) *InitialRhombusGridShape { return p.InitialRhombusGridShape },
		"InitialRhombusGridShape",
	)
}

// enforceExplanationTextShapeName ensures that the name of the ExplanationTextShape matches its owning Plant
func (stager *Stager) enforceExplanationTextShapeName() (needCommit bool) {
	return enforcePlantShapeName[*ExplanationTextShape](
		stager,
		func(p *Plant) *ExplanationTextShape { return p.ExplanationTextShape },
		"ExplanationTextShape",
	)
}

// enforceRotatedShapesNames ensures that the name of the Rotated shapes match their owning Plant
func (stager *Stager) enforceRotatedShapesNames() (needCommit bool) {
	n1 := enforcePlantShapeName[*RhombusShape](
		stager,
		func(p *Plant) *RhombusShape { return p.RotatedReferenceRhombus },
		"RotatedReferenceRhombus",
	)

	n2 := enforcePlantShapeName[*PlantCircumferenceShape](
		stager,
		func(p *Plant) *PlantCircumferenceShape { return p.RotatedPlantCircumferenceShape },
		"RotatedPlantCircumferenceShape",
	)

	n3 := enforcePlantShapeName[*GridPathShape](
		stager,
		func(p *Plant) *GridPathShape { return p.RotatedGridPathShape },
		"RotatedGridPathShape",
	)

	n4 := enforcePlantShapeName[*RotatedRhombusGridShape](
		stager,
		func(p *Plant) *RotatedRhombusGridShape { return p.RotatedRhombusGridShape2 },
		"RotatedRhombusGridShape",
	)

	n5 := enforcePlantShapeName[*GrowthCurveRhombusGridShape](
		stager,
		func(p *Plant) *GrowthCurveRhombusGridShape { return p.GrowthCurveRhombusGridShape },
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

	n7_bottom_arc_v2 := enforcePlantShapeName[*BottomStartArcShapeGrid](
		stager,
		func(p *Plant) *BottomStartArcShapeGrid { return p.BottomStartArcShapeGrid },
		"BottomStartArcShapeV2Grid",
	)
	needCommit = n7_bottom_arc_v2 || needCommit

	n7_bottom_arc_v2_end := enforcePlantShapeName[*BottomEndArcShapeGrid](
		stager,
		func(p *Plant) *BottomEndArcShapeGrid { return p.BottomEndArcShapeGrid },
		"BottomEndArcShapeV2Grid",
	)
	needCommit = n7_bottom_arc_v2_end || needCommit

	n8 := enforcePlantShapeName[*GrowthCurveBezierShapeGrid](
		stager,
		func(p *Plant) *GrowthCurveBezierShapeGrid { return p.GrowthCurveBezierShapeGrid },
		"GrowthCurveBezierShapeGrid",
	)
	needCommit = n8 || needCommit

	n10 := enforcePlantShapeName[*StackOfGrowthCurve](
		stager,
		func(p *Plant) *StackOfGrowthCurve { return p.StackOfGrowthCurve },
		"StackOfGrowthCurveV2",
	)
	needCommit = n10 || needCommit

	n11 := enforcePlantShapeName[*TopStackOfGrowthCurve](
		stager,
		func(p *Plant) *TopStackOfGrowthCurve { return p.TopStackOfGrowthCurve },
		"TopStackOfGrowthCurveV2",
	)
	needCommit = n11 || needCommit

	n12 := enforcePlantShapeName[*BottomStackOfGrowthCurve](
		stager,
		func(p *Plant) *BottomStackOfGrowthCurve { return p.BottomStackOfGrowthCurve },
		"BottomStackOfGrowthCurveV2",
	)
	needCommit = n12 || needCommit

	return n1 || n2 || n3 || n4 || n5 || n6 || n7 || n7_halfway || n7_base || n7_arc_normal || n7_arc_v2 || n7_top_arc_v2 || n7_arc_v2_end || n7_top_arc_v2_end || n7_bottom_arc_v2 || n7_bottom_arc_v2_end || n8 || n10 || n11 || n12
}

// enforcePlantRhombusGridShapeHasRhombuses ensures that each RhombusGridShape has the correct number of RhombusShapes and their X,Y fields are correctly computed
func (stager *Stager) enforcePlantRhombusGridShapeHasRhombuses() (needCommit bool) {
	stage := stager.stage

	for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
		angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
		length := plant.RhombusSideLength

		// Cartesian Y-axis points UP
		v1x := length * math.Cos(angleRad/2.0)
		v1y := length * math.Sin(angleRad/2.0)

		v2x := -length * math.Cos(angleRad/2.0)
		v2y := length * math.Sin(angleRad/2.0)

		{
			needCommit = enforceInitialRhombusGridShapeHasRhombuses(stage, plant.InitialRhombusGridShape, plant.N, plant.M, v1x, v1y, v2x, v2y, 0.0) || needCommit
		}
		rotRad := 0.0
		{
			rotRad = -plant.PlantCircumferenceShape.AngleDegree * math.Pi / 180.0
		}
		{
			needCommit = enforceRotatedRhombusGridShapeHasRhombuses(stage, plant.RotatedRhombusGridShape2, plant.N, plant.M, v1x, v1y, v2x, v2y, rotRad) || needCommit
		}
		{
			needCommit = enforceGrowthPathRhombusGridShapeHasRhombuses(stage, plant.GrowthCurveRhombusGridShape, plant.RotatedRhombusGridShape2, plant.PlantCircumferenceShape.AngleDegree, plant.RhombusSideLength, plant.RhombusInsideAngle, plant.PlantCircumferenceShape.Length) || needCommit
		}

		{
			rhombuses := plant.GrowthCurveRhombusGridShape.GrowthCurveRhombusShapes
			if len(rhombuses) >= 2 {
				first := rhombuses[0]
				var endRhombus *GrowthCurveRhombusShape
				minDiffY := math.MaxFloat64
				for i := 1; i < len(rhombuses); i++ {
					r := rhombuses[i]
					if r.Y > first.Y {
						diff := r.Y - first.Y
						if diff < minDiffY {
							minDiffY = diff
							endRhombus = r
						}
					}
				}

				if endRhombus != nil {
					vx := endRhombus.X - first.X
					vy := endRhombus.Y - first.Y
					if plant.GrowthVectorShape.X != vx || plant.GrowthVectorShape.Y != vy {
						plant.GrowthVectorShape.X = vx
						plant.GrowthVectorShape.Y = vy
						needCommit = true
					}
				} else {
					if plant.GrowthVectorShape.X != 0 || plant.GrowthVectorShape.Y != 0 {
						plant.GrowthVectorShape.X = 0
						plant.GrowthVectorShape.Y = 0
						needCommit = true
					}
				}
			} else {
				if plant.GrowthVectorShape.X != 0 || plant.GrowthVectorShape.Y != 0 {
					plant.GrowthVectorShape.X = 0
					plant.GrowthVectorShape.Y = 0
					needCommit = true
				}
			}
		}

		{
			needCommit = enforcePerpendicularVectorGridHasVectors(stage, plant.PerpendicularVectorGrid, plant.GrowthCurveRhombusGridShape, v1x, v1y, v2x, v2y, rotRad) || needCommit
		}

		{
			needCommit = enforcePerpendicularVectorGridHalfwayHasVectors(stage, plant.PerpendicularVectorGridHalfway, plant.PerpendicularVectorGrid) || needCommit
		}

		{
			needCommit = enforceBaseVectorShapeGridHasShapes(stage, plant.BaseVectorShapeGrid, plant.PerpendicularVectorGrid) || needCommit
		}

		{
			needCommit = enforceArcNormalVectorShapeGridHasShapes(stage, plant.ArcNormalVectorShapeGrid, plant.PerpendicularVectorGrid) || needCommit
		}

		{
			needCommit = enforceStartArcShapeV2GridHasShapes(stage, plant.StartArcShapeGrid, plant.PerpendicularVectorGrid) || needCommit
		}

		{
			needCommit = enforceTopStartArcShapeV2GridHasShapes(stage, plant.TopStartArcShapeGrid, plant.PerpendicularVectorGrid, plant.Thickness) || needCommit
		}

		{
			needCommit = enforceEndArcShapeV2GridHasShapes(stage, plant.EndArcShapeGrid, plant.PerpendicularVectorGrid) || needCommit
		}

		{
			needCommit = enforceTopEndArcShapeV2GridHasShapes(stage, plant.TopEndArcShapeGrid, plant.PerpendicularVectorGrid, plant.Thickness) || needCommit
		}

		if plant.GrowthCurve2D == nil {
			plant.GrowthCurve2D = new(GrowthCurve2D).Stage(stage)
			plant.GrowthCurve2D.Name = plant.Name + "-GrowthCurve2D"
			needCommit = true
		}
		{
			if plant.GrowthCurve2D.StartArcShapeGrid != plant.StartArcShapeGrid {
				plant.GrowthCurve2D.StartArcShapeGrid = plant.StartArcShapeGrid
				needCommit = true
			}
			if plant.GrowthCurve2D.EndArcShapeGrid != plant.EndArcShapeGrid {
				plant.GrowthCurve2D.EndArcShapeGrid = plant.EndArcShapeGrid
				needCommit = true
			}
		}

		if plant.TopGrowthCurve2D == nil {
			plant.TopGrowthCurve2D = new(TopGrowthCurve2D).Stage(stage)
			plant.TopGrowthCurve2D.Name = plant.Name + "-TopGrowthCurve2D"
			needCommit = true
		}
		{
			if plant.TopGrowthCurve2D.TopStartArcShapeGrid != plant.TopStartArcShapeGrid {
				plant.TopGrowthCurve2D.TopStartArcShapeGrid = plant.TopStartArcShapeGrid
				needCommit = true
			}
			if plant.TopGrowthCurve2D.TopEndArcShapeGrid != plant.TopEndArcShapeGrid {
				plant.TopGrowthCurve2D.TopEndArcShapeGrid = plant.TopEndArcShapeGrid
				needCommit = true
			}
		}

		{
			needCommit = enforceBottomStartArcShapeV2GridHasShapes(stage, plant.BottomStartArcShapeGrid, plant.PerpendicularVectorGrid, plant.Thickness) || needCommit
		}

		{
			needCommit = enforceBottomEndArcShapeV2GridHasShapes(stage, plant.BottomEndArcShapeGrid, plant.PerpendicularVectorGrid, plant.Thickness) || needCommit
		}

		circLen := 0.0
		{
			circLen = plant.PlantCircumferenceShape.Length
		}

		{
			needCommit = enforceGrowthCurveBezierShapeGridHasShapes(stage, plant.GrowthCurveBezierShapeGrid, plant.PerpendicularVectorGrid, plant.RhombusSideLength, circLen) || needCommit
		}

		{
			needCommit = enforceStackOfGrowthCurveV2HasShapes(stage, plant.StackOfGrowthCurve, plant.PerpendicularVectorGrid, plant.GrowthVectorShape, plant.StackHeight, circLen, plant.Thickness) || needCommit
		}
		{
			needCommit = enforceTopStackOfGrowthCurveV2HasShapes(stage, plant.TopStackOfGrowthCurve, plant.PerpendicularVectorGrid, plant.GrowthVectorShape, plant.StackHeight, circLen, plant.Thickness) || needCommit
		}
		{
			needCommit = enforceBottomStackOfGrowthCurveV2HasShapes(stage, plant.BottomStackOfGrowthCurve, plant.PerpendicularVectorGrid, plant.GrowthVectorShape, plant.StackHeight, circLen, plant.Thickness) || needCommit
		}
	}
	return
}

func enforcePerpendicularVectorGridHasVectors(stage *Stage, grid *PerpendicularVectorGrid, sourceGrid *GrowthCurveRhombusGridShape, v1x, v1y, v2x, v2y, rotAngleRad float64) (needCommit bool) {
	// The four vertices relative to the center before rotation
	p1x, p1y := -(v1x+v2x)/2.0, -(v1y+v2y)/2.0
	p2x, p2y := (v1x-v2x)/2.0, (v1y-v2y)/2.0
	p3x, p3y := (v2x-v1x)/2.0, (v2y-v1y)/2.0
	p4x, p4y := (v1x+v2x)/2.0, (v1y+v2y)/2.0

	cosA := math.Cos(rotAngleRad)
	sinA := math.Sin(rotAngleRad)

	pts := []struct{ x, y float64 }{
		{p1x*cosA - p1y*sinA, p1x*sinA + p1y*cosA},
		{p2x*cosA - p2y*sinA, p2x*sinA + p2y*cosA},
		{p3x*cosA - p3y*sinA, p3x*sinA + p3y*cosA},
		{p4x*cosA - p4y*sinA, p4x*sinA + p4y*cosA},
	}

	minY := pts[0].y
	bottomPt := pts[0]
	for i := 1; i < len(pts); i++ {
		if pts[i].y < minY {
			minY = pts[i].y
			bottomPt = pts[i]
		}
	}

	dx := bottomPt.x
	dy := bottomPt.y

	valid := true
	if len(grid.PerpendicularVectors) != len(sourceGrid.GrowthCurveRhombusShapes) {
		valid = false
	} else {
		for i, r := range sourceGrid.GrowthCurveRhombusShapes {
			vec := grid.PerpendicularVectors[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if vec == nil || vec.Name != expectedName {
				valid = false
				break
			}
			expectedStartX := r.X + dx
			expectedStartY := r.Y + dy
			if math.Abs(vec.StartX-expectedStartX) > 1e-4 || math.Abs(vec.StartY-expectedStartY) > 1e-4 ||
				math.Abs(vec.EndX-r.X) > 1e-4 || math.Abs(vec.EndY-r.Y) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		grid.PerpendicularVectors = make([]*PerpendicularVector, len(sourceGrid.GrowthCurveRhombusShapes))
		for i, r := range sourceGrid.GrowthCurveRhombusShapes {
			vec := new(PerpendicularVector).Stage(stage)
			vec.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			vec.StartX = r.X + dx
			vec.StartY = r.Y + dy
			vec.EndX = r.X
			vec.EndY = r.Y
			grid.PerpendicularVectors[i] = vec
		}
		needCommit = true
	} else {
		for i, r := range sourceGrid.GrowthCurveRhombusShapes {
			vec := grid.PerpendicularVectors[i]
			expectedStartX := r.X + dx
			expectedStartY := r.Y + dy
			if vec.StartX != expectedStartX || vec.StartY != expectedStartY || vec.EndX != r.X || vec.EndY != r.Y {
				vec.StartX = expectedStartX
				vec.StartY = expectedStartY
				vec.EndX = r.X
				vec.EndY = r.Y
				needCommit = true
			}
		}
	}

	return
}

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

func isRhombusShapeOwnedByPlant(p *Plant, shape *RhombusShape) bool {
	if p.ReferenceRhombus == shape || p.RotatedReferenceRhombus == shape {
		return true
	}
	// Initial, Rotated and Growth grids no longer use generic RhombusShape.
	return false
}

func enforceBaseVectorShapeGridHasShapes(stage *Stage, grid *BaseVectorShapeGrid, pGrid *PerpendicularVectorGrid) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.BaseVectorShapes) > 0 {
			grid.BaseVectorShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.BaseVectorShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			base := grid.BaseVectorShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if base == nil || base.Name != expectedName {
				valid = false
				break
			}

			expectedStartX := v1.StartX
			expectedStartY := v1.StartY
			expectedEndX := v2.StartX
			expectedEndY := v2.StartY

			if math.Abs(base.StartX-expectedStartX) > 1e-4 || math.Abs(base.StartY-expectedStartY) > 1e-4 ||
				math.Abs(base.EndX-expectedEndX) > 1e-4 || math.Abs(base.EndY-expectedEndY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.BaseVectorShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.BaseVectorShapes = make([]*BaseVectorShape, expectedLen)
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			expectedStartX := v1.StartX
			expectedStartY := v1.StartY
			expectedEndX := v2.StartX
			expectedEndY := v2.StartY

			newBase := new(BaseVectorShape).Stage(stage)
			newBase.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newBase.StartX = expectedStartX
			newBase.StartY = expectedStartY
			newBase.EndX = expectedEndX
			newBase.EndY = expectedEndY

			grid.BaseVectorShapes[i] = newBase
		}
		needCommit = true
	}
	return needCommit
}

func enforceArcNormalVectorShapeGridHasShapes(stage *Stage, grid *ArcNormalVectorShapeGrid, pGrid *PerpendicularVectorGrid) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.ArcNormalVectorShapes) > 0 {
			grid.ArcNormalVectorShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.ArcNormalVectorShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arcNorm := grid.ArcNormalVectorShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if arcNorm == nil || arcNorm.Name != expectedName {
				valid = false
				break
			}

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
			origSweepFlag := (model_cross < 0)

			radX := midX - cx
			radY := midY - cy
			var tx, ty float64
			if origSweepFlag {
				tx = -radY
				ty = radX
			} else {
				tx = radY
				ty = -radX
			}

			tLen := math.Hypot(tx, ty)
			if tLen != 0 {
				tx /= tLen
				ty /= tLen
			}

			dirX := ty
			dirY := -tx

			vLen := length

			expectedStartX := midX
			expectedStartY := midY
			expectedEndX := midX + dirX*vLen
			expectedEndY := midY + dirY*vLen

			if math.Abs(arcNorm.StartX-expectedStartX) > 1e-4 || math.Abs(arcNorm.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arcNorm.EndX-expectedEndX) > 1e-4 || math.Abs(arcNorm.EndY-expectedEndY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.ArcNormalVectorShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.ArcNormalVectorShapes = make([]*ArcNormalVectorShape, expectedLen)
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
			origSweepFlag := (model_cross < 0)

			radX := midX - cx
			radY := midY - cy
			var tx, ty float64
			if origSweepFlag {
				tx = -radY
				ty = radX
			} else {
				tx = radY
				ty = -radX
			}

			tLen := math.Hypot(tx, ty)
			if tLen != 0 {
				tx /= tLen
				ty /= tLen
			}

			dirX := ty
			dirY := -tx

			vLen := length

			expectedStartX := midX
			expectedStartY := midY
			expectedEndX := midX + dirX*vLen
			expectedEndY := midY + dirY*vLen

			newVec := new(ArcNormalVectorShape).Stage(stage)
			newVec.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newVec.StartX = expectedStartX
			newVec.StartY = expectedStartY
			newVec.EndX = expectedEndX
			newVec.EndY = expectedEndY

			grid.ArcNormalVectorShapes[i] = newVec
		}
		needCommit = true
	}
	return needCommit
}

func enforceStartArcShapeV2GridHasShapes(stage *Stage, grid *StartArcShapeGrid, pGrid *PerpendicularVectorGrid) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.StartArcShapes) > 0 {
			grid.StartArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.StartArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arc := grid.StartArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if arc == nil || arc.Name != expectedName {
				valid = false
				break
			}

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			expectedStartX := v1.StartX
			expectedStartY := v1.StartY

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)
			R := math.Abs(R_val)
			expectedRadiusX := R
			expectedRadiusY := R

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			expectedEndX := midX
			expectedEndY := midY

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x

			expectedSweepFlag := (model_cross < 0)

			if math.Abs(arc.StartX-expectedStartX) > 1e-4 || math.Abs(arc.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arc.EndX-expectedEndX) > 1e-4 || math.Abs(arc.EndY-expectedEndY) > 1e-4 ||
				math.Abs(arc.RadiusX-expectedRadiusX) > 1e-4 || math.Abs(arc.RadiusY-expectedRadiusY) > 1e-4 ||
				arc.SweepFlag != expectedSweepFlag {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.StartArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.StartArcShapes = make([]*StartArcShape, expectedLen)
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			expectedStartX := v1.StartX
			expectedStartY := v1.StartY

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)
			R := math.Abs(R_val)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			expectedEndX := midX
			expectedEndY := midY

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x

			sweepFlag := (model_cross < 0)

			newArc := new(StartArcShape).Stage(stage)
			newArc.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newArc.StartX = expectedStartX
			newArc.StartY = expectedStartY
			newArc.EndX = expectedEndX
			newArc.EndY = expectedEndY
			newArc.RadiusX = R
			newArc.RadiusY = R
			newArc.SweepFlag = sweepFlag
			newArc.LargeArcFlag = false
			newArc.XAxisRotation = 0

			grid.StartArcShapes[i] = newArc
		}
		needCommit = true
	}
	return needCommit
}

func enforceTopStartArcShapeV2GridHasShapes(stage *Stage, grid *TopStartArcShapeGrid, pGrid *PerpendicularVectorGrid, thickness float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.TopStartArcShapes) > 0 {
			grid.TopStartArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.TopStartArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arc := grid.TopStartArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if arc == nil || arc.Name != expectedName {
				valid = false
				break
			}

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)
			R := math.Abs(R_val)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x

			sweepFlag := (model_cross < 0)

			scale := 1.0 - thickness/R_val
			R_new := R * math.Abs(scale)

			expectedStartX := cx + scale*(v1.StartX-cx)
			expectedStartY := cy + scale*(v1.StartY-cy)
			expectedEndX := cx + scale*(midX-cx)
			expectedEndY := cy + scale*(midY-cy)

			expectedRadiusX := R_new
			expectedRadiusY := R_new

			if math.Abs(arc.StartX-expectedStartX) > 1e-4 || math.Abs(arc.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arc.EndX-expectedEndX) > 1e-4 || math.Abs(arc.EndY-expectedEndY) > 1e-4 ||
				math.Abs(arc.RadiusX-expectedRadiusX) > 1e-4 || math.Abs(arc.RadiusY-expectedRadiusY) > 1e-4 ||
				arc.SweepFlag != sweepFlag {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.TopStartArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.TopStartArcShapes = make([]*TopStartArcShape, expectedLen)
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)
			R := math.Abs(R_val)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x

			sweepFlag := (model_cross < 0)

			scale := 1.0 - thickness/R_val
			R_new := R * math.Abs(scale)

			expectedStartX := cx + scale*(v1.StartX-cx)
			expectedStartY := cy + scale*(v1.StartY-cy)
			expectedEndX := cx + scale*(midX-cx)
			expectedEndY := cy + scale*(midY-cy)

			newArc := new(TopStartArcShape).Stage(stage)
			newArc.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newArc.StartX = expectedStartX
			newArc.StartY = expectedStartY
			newArc.EndX = expectedEndX
			newArc.EndY = expectedEndY
			newArc.RadiusX = R_new
			newArc.RadiusY = R_new
			newArc.SweepFlag = sweepFlag
			newArc.LargeArcFlag = false
			newArc.XAxisRotation = 0

			grid.TopStartArcShapes[i] = newArc
		}
		needCommit = true
	}
	return needCommit
}

func enforcePerpendicularVectorGridHalfwayHasVectors(stage *Stage, grid *PerpendicularVectorGridHalfway, sourceGrid *PerpendicularVectorGrid) (needCommit bool) {
	if sourceGrid == nil || grid == nil || len(sourceGrid.PerpendicularVectors) < 2 {
		if len(grid.PerpendicularVectorHalfways) > 0 {
			grid.PerpendicularVectorHalfways = nil
			return true
		}
		return false
	}

	expectedLen := len(sourceGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.PerpendicularVectorHalfways) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := sourceGrid.PerpendicularVectors[i]
			v2 := sourceGrid.PerpendicularVectors[i+1]
			vHalfway := grid.PerpendicularVectorHalfways[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if vHalfway == nil || vHalfway.Name != expectedName {
				valid = false
				break
			}
			expectedStartX := (v1.StartX + v2.StartX) / 2.0
			expectedStartY := (v1.StartY + v2.StartY) / 2.0

			// Translation of v1
			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY

			expectedEndX := expectedStartX + dx
			expectedEndY := expectedStartY + dy

			if math.Abs(vHalfway.StartX-expectedStartX) > 1e-4 || math.Abs(vHalfway.StartY-expectedStartY) > 1e-4 ||
				math.Abs(vHalfway.EndX-expectedEndX) > 1e-4 || math.Abs(vHalfway.EndY-expectedEndY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, v := range grid.PerpendicularVectorHalfways {
			if v != nil {
				v.Unstage(stage)
			}
		}
		grid.PerpendicularVectorHalfways = make([]*PerpendicularVectorHalfway, expectedLen)
		for i := 0; i < expectedLen; i++ {
			v1 := sourceGrid.PerpendicularVectors[i]
			v2 := sourceGrid.PerpendicularVectors[i+1]

			expectedStartX := (v1.StartX + v2.StartX) / 2.0
			expectedStartY := (v1.StartY + v2.StartY) / 2.0

			// Translation of v1
			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY

			expectedEndX := expectedStartX + dx
			expectedEndY := expectedStartY + dy

			newVec := new(PerpendicularVectorHalfway).Stage(stage)
			newVec.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newVec.StartX = expectedStartX
			newVec.StartY = expectedStartY
			newVec.EndX = expectedEndX
			newVec.EndY = expectedEndY
			grid.PerpendicularVectorHalfways[i] = newVec
		}
		needCommit = true
	}
	return needCommit
}

func enforceEndArcShapeV2GridHasShapes(stage *Stage, grid *EndArcShapeGrid, pGrid *PerpendicularVectorGrid) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.EndArcShapes) > 0 {
			grid.EndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.EndArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arc := grid.EndArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if arc == nil || arc.Name != expectedName {
				valid = false
				break
			}

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)
			R := math.Abs(R_val)
			expectedRadiusX := R
			expectedRadiusY := R

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
			origSweepFlag := (model_cross < 0)

			expectedStartX := 2*midX - v1.StartX
			expectedStartY := 2*midY - v1.StartY
			expectedEndX := midX
			expectedEndY := midY
			expectedSweepFlag := origSweepFlag

			if math.Abs(arc.StartX-expectedStartX) > 1e-4 || math.Abs(arc.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arc.EndX-expectedEndX) > 1e-4 || math.Abs(arc.EndY-expectedEndY) > 1e-4 ||
				math.Abs(arc.RadiusX-expectedRadiusX) > 1e-4 || math.Abs(arc.RadiusY-expectedRadiusY) > 1e-4 ||
				arc.SweepFlag != expectedSweepFlag {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.EndArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.EndArcShapes = make([]*EndArcShape, expectedLen)
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)
			R := math.Abs(R_val)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
			origSweepFlag := (model_cross < 0)

			expectedStartX := 2*midX - v1.StartX
			expectedStartY := 2*midY - v1.StartY
			expectedEndX := midX
			expectedEndY := midY

			newArc := new(EndArcShape).Stage(stage)
			newArc.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newArc.StartX = expectedStartX
			newArc.StartY = expectedStartY
			newArc.EndX = expectedEndX
			newArc.EndY = expectedEndY
			newArc.RadiusX = R
			newArc.RadiusY = R
			newArc.SweepFlag = origSweepFlag
			newArc.LargeArcFlag = false
			newArc.XAxisRotation = 0

			grid.EndArcShapes[i] = newArc
		}
		needCommit = true
	}
	return needCommit
}

func enforceTopEndArcShapeV2GridHasShapes(stage *Stage, grid *TopEndArcShapeGrid, pGrid *PerpendicularVectorGrid, thickness float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.TopEndArcShapes) > 0 {
			grid.TopEndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.TopEndArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arc := grid.TopEndArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if arc == nil || arc.Name != expectedName {
				valid = false
				break
			}

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)
			R := math.Abs(R_val)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
			origSweepFlag := (model_cross < 0)

			cx_new := 2*midX - cx
			cy_new := 2*midY - cy

			scale := 1.0 + thickness/R_val
			R_new := R * math.Abs(scale)

			expectedStartX := cx_new - scale*(v1.StartX-cx)
			expectedStartY := cy_new - scale*(v1.StartY-cy)
			expectedEndX := cx_new + scale*(midX-cx_new)
			expectedEndY := cy_new + scale*(midY-cy_new)

			expectedRadiusX := R_new
			expectedRadiusY := R_new

			if math.Abs(arc.StartX-expectedStartX) > 1e-4 || math.Abs(arc.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arc.EndX-expectedEndX) > 1e-4 || math.Abs(arc.EndY-expectedEndY) > 1e-4 ||
				math.Abs(arc.RadiusX-expectedRadiusX) > 1e-4 || math.Abs(arc.RadiusY-expectedRadiusY) > 1e-4 ||
				arc.SweepFlag != origSweepFlag {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.TopEndArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.TopEndArcShapes = make([]*TopEndArcShape, expectedLen)
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)
			R := math.Abs(R_val)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
			origSweepFlag := (model_cross < 0)

			cx_new := 2*midX - cx
			cy_new := 2*midY - cy

			scale := 1.0 + thickness/R_val
			R_new := R * math.Abs(scale)

			expectedStartX := cx_new - scale*(v1.StartX-cx)
			expectedStartY := cy_new - scale*(v1.StartY-cy)
			expectedEndX := cx_new + scale*(midX-cx_new)
			expectedEndY := cy_new + scale*(midY-cy_new)

			newArc := new(TopEndArcShape).Stage(stage)
			newArc.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newArc.StartX = expectedStartX
			newArc.StartY = expectedStartY
			newArc.EndX = expectedEndX
			newArc.EndY = expectedEndY
			newArc.RadiusX = R_new
			newArc.RadiusY = R_new
			newArc.SweepFlag = origSweepFlag
			newArc.LargeArcFlag = false
			newArc.XAxisRotation = 0

			grid.TopEndArcShapes[i] = newArc
		}
		needCommit = true
	}
	return needCommit
}

func enforceGrowthCurveBezierShapeGridHasShapes(stage *Stage, grid *GrowthCurveBezierShapeGrid, pGrid *PerpendicularVectorGrid, sideLength float64, circumferenceLength float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.GrowthCurveBezierShapes) > 0 {
			grid.GrowthCurveBezierShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	expectedTotalLen := expectedLen * 2
	valid := true
	if len(grid.GrowthCurveBezierShapes) != expectedTotalLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			b := grid.GrowthCurveBezierShapes[i]
			bTrans := grid.GrowthCurveBezierShapes[i+expectedLen]

			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			expectedTransName := fmt.Sprintf("%s-translated-%d", grid.Name, i)
			if b == nil || b.Name != expectedName || bTrans == nil || bTrans.Name != expectedTransName {
				valid = false
				break
			}
			angleRad := math.Atan2(v1.EndY-v1.StartY, v1.EndX-v1.StartX) - math.Pi/2.0

			// wait, if we consider length ratio = 0.33
			ratio := 0.33
			expectedCtrlStartX := v1.StartX + sideLength*ratio*math.Cos(angleRad)
			expectedCtrlStartY := v1.StartY + sideLength*ratio*math.Sin(angleRad)
			expectedCtrlEndX := v2.StartX + sideLength*ratio*math.Cos(angleRad+math.Pi)
			expectedCtrlEndY := v2.StartY + sideLength*ratio*math.Sin(angleRad+math.Pi)

			if math.Abs(b.StartX-v1.StartX) > 1e-4 || math.Abs(b.StartY-v1.StartY) > 1e-4 ||
				math.Abs(b.EndX-v2.StartX) > 1e-4 || math.Abs(b.EndY-v2.StartY) > 1e-4 ||
				math.Abs(b.ControlPointStartX-expectedCtrlStartX) > 1e-4 || math.Abs(b.ControlPointStartY-expectedCtrlStartY) > 1e-4 ||
				math.Abs(b.ControlPointEndX-expectedCtrlEndX) > 1e-4 || math.Abs(b.ControlPointEndY-expectedCtrlEndY) > 1e-4 {
				valid = false
				break
			}

			if math.Abs(bTrans.StartX-(v1.StartX+circumferenceLength)) > 1e-4 || math.Abs(bTrans.StartY-v1.StartY) > 1e-4 ||
				math.Abs(bTrans.EndX-(v2.StartX+circumferenceLength)) > 1e-4 || math.Abs(bTrans.EndY-v2.StartY) > 1e-4 ||
				math.Abs(bTrans.ControlPointStartX-(expectedCtrlStartX+circumferenceLength)) > 1e-4 || math.Abs(bTrans.ControlPointStartY-expectedCtrlStartY) > 1e-4 ||
				math.Abs(bTrans.ControlPointEndX-(expectedCtrlEndX+circumferenceLength)) > 1e-4 || math.Abs(bTrans.ControlPointEndY-expectedCtrlEndY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		grid.GrowthCurveBezierShapes = make([]*GrowthCurveBezierShape, expectedTotalLen)
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			b := new(GrowthCurveBezierShape).Stage(stage)
			b.Name = fmt.Sprintf("%s-%d", grid.Name, i)

			b.StartX = v1.StartX
			b.StartY = v1.StartY
			b.EndX = v2.StartX
			b.EndY = v2.StartY

			angleRad := math.Atan2(v1.EndY-v1.StartY, v1.EndX-v1.StartX) - math.Pi/2.0
			ratio := 0.33
			b.ControlPointStartX = b.StartX + sideLength*ratio*math.Cos(angleRad)
			b.ControlPointStartY = b.StartY + sideLength*ratio*math.Sin(angleRad)
			b.ControlPointEndX = b.EndX + sideLength*ratio*math.Cos(angleRad+math.Pi)
			b.ControlPointEndY = b.EndY + sideLength*ratio*math.Sin(angleRad+math.Pi)

			grid.GrowthCurveBezierShapes[i] = b

			bTrans := new(GrowthCurveBezierShape).Stage(stage)
			bTrans.Name = fmt.Sprintf("%s-translated-%d", grid.Name, i)

			bTrans.StartX = b.StartX + circumferenceLength
			bTrans.StartY = b.StartY
			bTrans.EndX = b.EndX + circumferenceLength
			bTrans.EndY = b.EndY
			bTrans.ControlPointStartX = b.ControlPointStartX + circumferenceLength
			bTrans.ControlPointStartY = b.ControlPointStartY
			bTrans.ControlPointEndX = b.ControlPointEndX + circumferenceLength
			bTrans.ControlPointEndY = b.ControlPointEndY

			grid.GrowthCurveBezierShapes[i+expectedLen] = bTrans
		}
		needCommit = true
	}

	return needCommit
}

func enforceStackOfGrowthCurveV2HasShapes(stage *Stage, stack *StackOfGrowthCurve, pGrid *PerpendicularVectorGrid, vector *GrowthVectorShape, stackHeight int, circLen float64, thickness float64) (needCommit bool) {
	if stack == nil || pGrid == nil || vector == nil || stackHeight < 1 || circLen <= 0 || len(pGrid.PerpendicularVectors) < 2 {
		if len(stack.StackGrowthCurveStartArcShapes) > 0 || len(stack.StackGrowthCurveEndArcShapes) > 0 {
			stack.StackGrowthCurveStartArcShapes = nil
			stack.StackGrowthCurveEndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1

	type expectedStartShape struct {
		name                       string
		startX, startY, endX, endY float64
		radiusX, radiusY           float64
		xAxisRotation              float64
		largeArcFlag, sweepFlag    bool
	}
	var expectedStart []expectedStartShape
	var expectedEnd []expectedStartShape

	for h := 0; h < stackHeight; h++ {
		dx := float64(h) * vector.X
		dy := float64(h) * vector.Y
		offset := float64(h) * thickness

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			currentDX := math.Mod(dx, circLen)
			if currentDX < 0 {
				currentDX += circLen
			}

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp := computeArcV2Geometry(v1, v2, offset, false)
			expectedStart = append(expectedStart, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-start-%d-%d", stack.Name, h, i),
				startX: sX + currentDX, startY: sY + dy,
				endX: eX + currentDX, endY: eY + dy,
				radiusX: rX, radiusY: rY,
				xAxisRotation: xRot, largeArcFlag: lArc, sweepFlag: swp,
			})

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp = computeArcV2Geometry(v1, v2, offset, true)
			expectedEnd = append(expectedEnd, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-end-%d-%d", stack.Name, h, i),
				startX: sX + currentDX, startY: sY + dy,
				endX: eX + currentDX, endY: eY + dy,
				radiusX: rX, radiusY: rY,
				xAxisRotation: xRot, largeArcFlag: lArc, sweepFlag: swp,
			})
		}
	}

	valid := true
	if len(stack.StackGrowthCurveStartArcShapes) != len(expectedStart) || len(stack.StackGrowthCurveEndArcShapes) != len(expectedEnd) {
		valid = false
	} else {
		for i, exp := range expectedStart {
			b := stack.StackGrowthCurveStartArcShapes[i]
			if b == nil || b.Name != exp.name {
				valid = false
				break
			}
			if math.Abs(b.StartX-exp.startX) > 1e-4 || math.Abs(b.StartY-exp.startY) > 1e-4 ||
				math.Abs(b.EndX-exp.endX) > 1e-4 || math.Abs(b.EndY-exp.endY) > 1e-4 {
				valid = false
				break
			}
		}
		for i, exp := range expectedEnd {
			b := stack.StackGrowthCurveEndArcShapes[i]
			if b == nil || b.Name != exp.name {
				valid = false
				break
			}
			if math.Abs(b.StartX-exp.startX) > 1e-4 || math.Abs(b.StartY-exp.startY) > 1e-4 ||
				math.Abs(b.EndX-exp.endX) > 1e-4 || math.Abs(b.EndY-exp.endY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		stack.StackGrowthCurveStartArcShapes = make([]*StackGrowthCurveStartArcShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(StackGrowthCurveStartArcShape).Stage(stage)
			b.Name = exp.name
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			b.RadiusX = exp.radiusX
			b.RadiusY = exp.radiusY
			b.XAxisRotation = exp.xAxisRotation
			b.LargeArcFlag = exp.largeArcFlag
			b.SweepFlag = exp.sweepFlag
			stack.StackGrowthCurveStartArcShapes[i] = b
		}
		stack.StackGrowthCurveEndArcShapes = make([]*StackGrowthCurveEndArcShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(StackGrowthCurveEndArcShape).Stage(stage)
			b.Name = exp.name
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			b.RadiusX = exp.radiusX
			b.RadiusY = exp.radiusY
			b.XAxisRotation = exp.xAxisRotation
			b.LargeArcFlag = exp.largeArcFlag
			b.SweepFlag = exp.sweepFlag
			stack.StackGrowthCurveEndArcShapes[i] = b
		}
		needCommit = true
	}

	return needCommit
}

func enforceTopStackOfGrowthCurveV2HasShapes(stage *Stage, stack *TopStackOfGrowthCurve, pGrid *PerpendicularVectorGrid, vector *GrowthVectorShape, stackHeight int, circLen float64, thickness float64) (needCommit bool) {
	if stack == nil || pGrid == nil || vector == nil || stackHeight < 1 || circLen <= 0 || len(pGrid.PerpendicularVectors) < 2 {
		if len(stack.TopStackGrowthCurveStartArcShapes) > 0 || len(stack.TopStackGrowthCurveEndArcShapes) > 0 {
			stack.TopStackGrowthCurveStartArcShapes = nil
			stack.TopStackGrowthCurveEndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1

	type expectedStartShape struct {
		name                       string
		startX, startY, endX, endY float64
		radiusX, radiusY           float64
		xAxisRotation              float64
		largeArcFlag, sweepFlag    bool
	}
	var expectedStart []expectedStartShape
	var expectedEnd []expectedStartShape

	for h := 0; h < stackHeight; h++ {
		dx := float64(h) * vector.X
		dy := float64(h) * vector.Y
		offset := thickness + float64(h)*thickness

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			currentDX := math.Mod(dx, circLen)
			if currentDX < 0 {
				currentDX += circLen
			}

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp := computeArcV2Geometry(v1, v2, offset, false)
			expectedStart = append(expectedStart, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-start-%d-%d", stack.Name, h, i),
				startX: sX + currentDX, startY: sY + dy,
				endX: eX + currentDX, endY: eY + dy,
				radiusX: rX, radiusY: rY,
				xAxisRotation: xRot, largeArcFlag: lArc, sweepFlag: swp,
			})

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp = computeArcV2Geometry(v1, v2, offset, true)
			expectedEnd = append(expectedEnd, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-end-%d-%d", stack.Name, h, i),
				startX: sX + currentDX, startY: sY + dy,
				endX: eX + currentDX, endY: eY + dy,
				radiusX: rX, radiusY: rY,
				xAxisRotation: xRot, largeArcFlag: lArc, sweepFlag: swp,
			})
		}
	}

	valid := true
	if len(stack.TopStackGrowthCurveStartArcShapes) != len(expectedStart) || len(stack.TopStackGrowthCurveEndArcShapes) != len(expectedEnd) {
		valid = false
	} else {
		for i, exp := range expectedStart {
			b := stack.TopStackGrowthCurveStartArcShapes[i]
			if b == nil || b.Name != exp.name {
				valid = false
				break
			}
			if math.Abs(b.StartX-exp.startX) > 1e-4 || math.Abs(b.StartY-exp.startY) > 1e-4 ||
				math.Abs(b.EndX-exp.endX) > 1e-4 || math.Abs(b.EndY-exp.endY) > 1e-4 {
				valid = false
				break
			}
		}
		for i, exp := range expectedEnd {
			b := stack.TopStackGrowthCurveEndArcShapes[i]
			if b == nil || b.Name != exp.name {
				valid = false
				break
			}
			if math.Abs(b.StartX-exp.startX) > 1e-4 || math.Abs(b.StartY-exp.startY) > 1e-4 ||
				math.Abs(b.EndX-exp.endX) > 1e-4 || math.Abs(b.EndY-exp.endY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		stack.TopStackGrowthCurveStartArcShapes = make([]*TopStackGrowthCurveStartArcShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(TopStackGrowthCurveStartArcShape).Stage(stage)
			b.Name = exp.name
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			b.RadiusX = exp.radiusX
			b.RadiusY = exp.radiusY
			b.XAxisRotation = exp.xAxisRotation
			b.LargeArcFlag = exp.largeArcFlag
			b.SweepFlag = exp.sweepFlag
			stack.TopStackGrowthCurveStartArcShapes[i] = b
		}
		stack.TopStackGrowthCurveEndArcShapes = make([]*TopStackGrowthCurveEndArcShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(TopStackGrowthCurveEndArcShape).Stage(stage)
			b.Name = exp.name
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			b.RadiusX = exp.radiusX
			b.RadiusY = exp.radiusY
			b.XAxisRotation = exp.xAxisRotation
			b.LargeArcFlag = exp.largeArcFlag
			b.SweepFlag = exp.sweepFlag
			stack.TopStackGrowthCurveEndArcShapes[i] = b
		}
		needCommit = true
	}

	return needCommit
}

func enforceBottomStackOfGrowthCurveV2HasShapes(stage *Stage, stack *BottomStackOfGrowthCurve, pGrid *PerpendicularVectorGrid, vector *GrowthVectorShape, stackHeight int, circLen float64, thickness float64) (needCommit bool) {
	if stack == nil || pGrid == nil || vector == nil || stackHeight < 1 || circLen <= 0 || len(pGrid.PerpendicularVectors) < 2 {
		if len(stack.BottomStackGrowthCurveStartArcShapes) > 0 || len(stack.BottomStackGrowthCurveEndArcShapes) > 0 {
			stack.BottomStackGrowthCurveStartArcShapes = nil
			stack.BottomStackGrowthCurveEndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1

	type expectedStartShape struct {
		name                       string
		startX, startY, endX, endY float64
		radiusX, radiusY           float64
		xAxisRotation              float64
		largeArcFlag, sweepFlag    bool
	}
	var expectedStart []expectedStartShape
	var expectedEnd []expectedStartShape

	for h := 0; h < stackHeight; h++ {
		dx := float64(h) * vector.X
		dy := float64(h) * vector.Y
		offset := -thickness + float64(h)*thickness

		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			currentDX := math.Mod(dx, circLen)
			if currentDX < 0 {
				currentDX += circLen
			}

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp := computeArcV2Geometry(v1, v2, offset, false)
			expectedStart = append(expectedStart, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-start-%d-%d", stack.Name, h, i),
				startX: sX + currentDX, startY: sY + dy,
				endX: eX + currentDX, endY: eY + dy,
				radiusX: rX, radiusY: rY,
				xAxisRotation: xRot, largeArcFlag: lArc, sweepFlag: swp,
			})

			sX, sY, eX, eY, rX, rY, xRot, lArc, swp = computeArcV2Geometry(v1, v2, offset, true)
			expectedEnd = append(expectedEnd, expectedStartShape{
				name:   fmt.Sprintf("%s-layer-end-%d-%d", stack.Name, h, i),
				startX: sX + currentDX, startY: sY + dy,
				endX: eX + currentDX, endY: eY + dy,
				radiusX: rX, radiusY: rY,
				xAxisRotation: xRot, largeArcFlag: lArc, sweepFlag: swp,
			})
		}
	}

	valid := true
	if len(stack.BottomStackGrowthCurveStartArcShapes) != len(expectedStart) || len(stack.BottomStackGrowthCurveEndArcShapes) != len(expectedEnd) {
		valid = false
	} else {
		for i, exp := range expectedStart {
			b := stack.BottomStackGrowthCurveStartArcShapes[i]
			if b == nil || b.Name != exp.name {
				valid = false
				break
			}
			if math.Abs(b.StartX-exp.startX) > 1e-4 || math.Abs(b.StartY-exp.startY) > 1e-4 ||
				math.Abs(b.EndX-exp.endX) > 1e-4 || math.Abs(b.EndY-exp.endY) > 1e-4 {
				valid = false
				break
			}
		}
		for i, exp := range expectedEnd {
			b := stack.BottomStackGrowthCurveEndArcShapes[i]
			if b == nil || b.Name != exp.name {
				valid = false
				break
			}
			if math.Abs(b.StartX-exp.startX) > 1e-4 || math.Abs(b.StartY-exp.startY) > 1e-4 ||
				math.Abs(b.EndX-exp.endX) > 1e-4 || math.Abs(b.EndY-exp.endY) > 1e-4 {
				valid = false
				break
			}
		}
	}

	if !valid {
		stack.BottomStackGrowthCurveStartArcShapes = make([]*BottomStackGrowthCurveStartArcShape, len(expectedStart))
		for i, exp := range expectedStart {
			b := new(BottomStackGrowthCurveStartArcShape).Stage(stage)
			b.Name = exp.name
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			b.RadiusX = exp.radiusX
			b.RadiusY = exp.radiusY
			b.XAxisRotation = exp.xAxisRotation
			b.LargeArcFlag = exp.largeArcFlag
			b.SweepFlag = exp.sweepFlag
			stack.BottomStackGrowthCurveStartArcShapes[i] = b
		}
		stack.BottomStackGrowthCurveEndArcShapes = make([]*BottomStackGrowthCurveEndArcShape, len(expectedEnd))
		for i, exp := range expectedEnd {
			b := new(BottomStackGrowthCurveEndArcShape).Stage(stage)
			b.Name = exp.name
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			b.RadiusX = exp.radiusX
			b.RadiusY = exp.radiusY
			b.XAxisRotation = exp.xAxisRotation
			b.LargeArcFlag = exp.largeArcFlag
			b.SweepFlag = exp.sweepFlag
			stack.BottomStackGrowthCurveEndArcShapes[i] = b
		}
		needCommit = true
	}

	return needCommit
}

func enforceBottomStartArcShapeV2GridHasShapes(stage *Stage, grid *BottomStartArcShapeGrid, pGrid *PerpendicularVectorGrid, thickness float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.BottomStartArcShapes) > 0 {
			grid.BottomStartArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.BottomStartArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arc := grid.BottomStartArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if arc == nil || arc.Name != expectedName {
				valid = false
				break
			}

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)
			R := math.Abs(R_val)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x

			sweepFlag := (model_cross < 0)

			scale := 1.0 + thickness/R_val
			R_new := R * math.Abs(scale)

			expectedStartX := cx + scale*(v1.StartX-cx)
			expectedStartY := cy + scale*(v1.StartY-cy)
			expectedEndX := cx + scale*(midX-cx)
			expectedEndY := cy + scale*(midY-cy)

			expectedRadiusX := R_new
			expectedRadiusY := R_new

			if math.Abs(arc.StartX-expectedStartX) > 1e-4 || math.Abs(arc.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arc.EndX-expectedEndX) > 1e-4 || math.Abs(arc.EndY-expectedEndY) > 1e-4 ||
				math.Abs(arc.RadiusX-expectedRadiusX) > 1e-4 || math.Abs(arc.RadiusY-expectedRadiusY) > 1e-4 ||
				arc.SweepFlag != sweepFlag {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.BottomStartArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.BottomStartArcShapes = make([]*BottomStartArcShape, expectedLen)
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)
			R := math.Abs(R_val)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x

			sweepFlag := (model_cross < 0)

			scale := 1.0 + thickness/R_val
			R_new := R * math.Abs(scale)

			expectedStartX := cx + scale*(v1.StartX-cx)
			expectedStartY := cy + scale*(v1.StartY-cy)
			expectedEndX := cx + scale*(midX-cx)
			expectedEndY := cy + scale*(midY-cy)

			newArc := new(BottomStartArcShape).Stage(stage)
			newArc.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newArc.StartX = expectedStartX
			newArc.StartY = expectedStartY
			newArc.EndX = expectedEndX
			newArc.EndY = expectedEndY
			newArc.RadiusX = R_new
			newArc.RadiusY = R_new
			newArc.SweepFlag = sweepFlag
			newArc.LargeArcFlag = false
			newArc.XAxisRotation = 0

			grid.BottomStartArcShapes[i] = newArc
		}
		needCommit = true
	}
	return needCommit
}

func enforceBottomEndArcShapeV2GridHasShapes(stage *Stage, grid *BottomEndArcShapeGrid, pGrid *PerpendicularVectorGrid, thickness float64) (needCommit bool) {
	if pGrid == nil || grid == nil || len(pGrid.PerpendicularVectors) < 2 {
		if len(grid.BottomEndArcShapes) > 0 {
			grid.BottomEndArcShapes = nil
			return true
		}
		return false
	}

	expectedLen := len(pGrid.PerpendicularVectors) - 1
	valid := true
	if len(grid.BottomEndArcShapes) != expectedLen {
		valid = false
	} else {
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]
			arc := grid.BottomEndArcShapes[i]
			expectedName := fmt.Sprintf("%s-%d", grid.Name, i)
			if arc == nil || arc.Name != expectedName {
				valid = false
				break
			}

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)
			R := math.Abs(R_val)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
			origSweepFlag := (model_cross < 0)

			cx_new := 2*midX - cx
			cy_new := 2*midY - cy

			scale := 1.0 - thickness/R_val
			R_new := R * math.Abs(scale)

			expectedStartX := cx_new - scale*(v1.StartX-cx)
			expectedStartY := cy_new - scale*(v1.StartY-cy)
			expectedEndX := cx_new + scale*(midX-cx_new)
			expectedEndY := cy_new + scale*(midY-cy_new)

			expectedRadiusX := R_new
			expectedRadiusY := R_new

			if math.Abs(arc.StartX-expectedStartX) > 1e-4 || math.Abs(arc.StartY-expectedStartY) > 1e-4 ||
				math.Abs(arc.EndX-expectedEndX) > 1e-4 || math.Abs(arc.EndY-expectedEndY) > 1e-4 ||
				math.Abs(arc.RadiusX-expectedRadiusX) > 1e-4 || math.Abs(arc.RadiusY-expectedRadiusY) > 1e-4 ||
				arc.SweepFlag != origSweepFlag {
				valid = false
				break
			}
		}
	}

	if !valid {
		for _, s := range grid.BottomEndArcShapes {
			if s != nil {
				s.Unstage(stage)
			}
		}
		grid.BottomEndArcShapes = make([]*BottomEndArcShape, expectedLen)
		for i := 0; i < expectedLen; i++ {
			v1 := pGrid.PerpendicularVectors[i]
			v2 := pGrid.PerpendicularVectors[i+1]

			dx := v1.EndX - v1.StartX
			dy := v1.EndY - v1.StartY
			length := math.Hypot(dx, dy)
			if length == 0 {
				length = 1
			}
			ux, uy := dx/length, dy/length

			midX := (v1.StartX + v2.StartX) / 2.0
			midY := (v1.StartY + v2.StartY) / 2.0

			Vx := v1.StartX - midX
			Vy := v1.StartY - midY
			V_sq := Vx*Vx + Vy*Vy
			V_dot_u := Vx*ux + Vy*uy
			if math.Abs(V_dot_u) < 1e-6 {
				if V_dot_u >= 0 {
					V_dot_u = 1e-6
				} else {
					V_dot_u = -1e-6
				}
			}

			R_val := -V_sq / (2.0 * V_dot_u)
			R := math.Abs(R_val)

			cx := v1.StartX + R_val*ux
			cy := v1.StartY + R_val*uy

			AB_model_x := midX - v1.StartX
			AB_model_y := midY - v1.StartY
			AC_model_x := cx - v1.StartX
			AC_model_y := cy - v1.StartY
			model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
			origSweepFlag := (model_cross < 0)

			cx_new := 2*midX - cx
			cy_new := 2*midY - cy

			scale := 1.0 - thickness/R_val
			R_new := R * math.Abs(scale)

			expectedStartX := cx_new - scale*(v1.StartX-cx)
			expectedStartY := cy_new - scale*(v1.StartY-cy)
			expectedEndX := cx_new + scale*(midX-cx_new)
			expectedEndY := cy_new + scale*(midY-cy_new)

			newArc := new(BottomEndArcShape).Stage(stage)
			newArc.Name = fmt.Sprintf("%s-%d", grid.Name, i)
			newArc.StartX = expectedStartX
			newArc.StartY = expectedStartY
			newArc.EndX = expectedEndX
			newArc.EndY = expectedEndY
			newArc.RadiusX = R_new
			newArc.RadiusY = R_new
			newArc.SweepFlag = origSweepFlag
			newArc.LargeArcFlag = false
			newArc.XAxisRotation = 0

			grid.BottomEndArcShapes[i] = newArc
		}
		needCommit = true
	}
	return needCommit
}

func computeArcV2Geometry(v1, v2 *PerpendicularVector, offset float64, isEndArc bool) (expectedStartX, expectedStartY, expectedEndX, expectedEndY, expectedRadiusX, expectedRadiusY, xAxisRotation float64, largeArcFlag, sweepFlag bool) {
	dx := v1.EndX - v1.StartX
	dy := v1.EndY - v1.StartY
	length := math.Hypot(dx, dy)
	if length == 0 {
		length = 1
	}
	ux, uy := dx/length, dy/length

	midX := (v1.StartX + v2.StartX) / 2.0
	midY := (v1.StartY + v2.StartY) / 2.0

	Vx := v1.StartX - midX
	Vy := v1.StartY - midY
	V_sq := Vx*Vx + Vy*Vy
	V_dot_u := Vx*ux + Vy*uy
	if math.Abs(V_dot_u) < 1e-6 {
		if V_dot_u >= 0 {
			V_dot_u = 1e-6
		} else {
			V_dot_u = -1e-6
		}
	}

	R_val := -V_sq / (2.0 * V_dot_u)
	R := math.Abs(R_val)

	cx := v1.StartX + R_val*ux
	cy := v1.StartY + R_val*uy

	AB_model_x := midX - v1.StartX
	AB_model_y := midY - v1.StartY
	AC_model_x := cx - v1.StartX
	AC_model_y := cy - v1.StartY
	model_cross := AB_model_x*AC_model_y - AB_model_y*AC_model_x
	baseSweepFlag := (model_cross < 0)

	if !isEndArc {
		scale := 1.0 - offset/R_val
		R_new := R * math.Abs(scale)

		expectedStartX = cx + scale*(v1.StartX-cx)
		expectedStartY = cy + scale*(v1.StartY-cy)
		expectedEndX = cx + scale*(midX-cx)
		expectedEndY = cy + scale*(midY-cy)

		expectedRadiusX = R_new
		expectedRadiusY = R_new
		xAxisRotation = 0.0
		largeArcFlag = false
		sweepFlag = baseSweepFlag
	} else {
		cx_new := 2*midX - cx
		cy_new := 2*midY - cy

		scale := 1.0 + offset/R_val
		R_new := R * math.Abs(scale)

		expectedStartX = cx_new - scale*(v1.StartX-cx)
		expectedStartY = cy_new - scale*(v1.StartY-cy)
		expectedEndX = cx_new + scale*(midX-cx_new)
		expectedEndY = cy_new + scale*(midY-cy_new)

		expectedRadiusX = R_new
		expectedRadiusY = R_new
		xAxisRotation = 0.0
		largeArcFlag = false
		sweepFlag = baseSweepFlag
	}

	return
}
