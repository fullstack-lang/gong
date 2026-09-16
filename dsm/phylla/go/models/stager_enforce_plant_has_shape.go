package models

import (
	"fmt"
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
			stager.logAndNotify(fmt.Sprintf("Plant %s: created missing %s", plant.Name, shapeName))

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
			stager.logAndNotify(fmt.Sprintf("Removed orphaned %s %s", shapeName, shape.GetName()))
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
				oldName := shape.GetName()
				shape.SetName(expectedName)
				stager.logAndNotify(fmt.Sprintf("Renamed %s from '%s' to '%s'", shapeNameSuffix, oldName, expectedName))
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
		func() *AxesShape { return &AxesShape{LengthX: 200.0, LengthY: 200.0} },
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
		"StartArcShapeGrid",
	)
	needCommit = n7_arc_v2 || needCommit

	n7_arc_v2_end := enforcePlantHasShape[*EndArcShapeGrid](
		stager,
		func() *EndArcShapeGrid { return new(EndArcShapeGrid) },
		func(p *PlantAbstract) *EndArcShapeGrid { return p.EndArcShapeGrid },
		func(p *PlantAbstract, shape *EndArcShapeGrid) { p.EndArcShapeGrid = shape },
		func(p *PlantAbstract, shape *EndArcShapeGrid) bool {
			return p.EndArcShapeGrid == shape
		},
		"EndArcShapeGrid",
	)
	needCommit = n7_arc_v2_end || needCommit

	n7_mid_arc := enforcePlantHasShape[*MidArcVectorShapeGrid](
		stager,
		func() *MidArcVectorShapeGrid { return new(MidArcVectorShapeGrid) },
		func(p *PlantAbstract) *MidArcVectorShapeGrid { return p.MidArcVectorShapeGrid },
		func(p *PlantAbstract, shape *MidArcVectorShapeGrid) { p.MidArcVectorShapeGrid = shape },
		func(p *PlantAbstract, shape *MidArcVectorShapeGrid) bool {
			return p.MidArcVectorShapeGrid == shape
		},
		"MidArcVectorShapeGrid",
	)
	needCommit = n7_mid_arc || needCommit

	n7_growth_curve := enforcePlantHasShape[*GrowthCurve2D](
		stager,
		func() *GrowthCurve2D { return new(GrowthCurve2D) },
		func(p *PlantAbstract) *GrowthCurve2D { return p.GrowthCurve2D },
		func(p *PlantAbstract, shape *GrowthCurve2D) { p.GrowthCurve2D = shape },
		func(p *PlantAbstract, shape *GrowthCurve2D) bool {
			return p.GrowthCurve2D == shape
		},
		"GrowthCurve2D",
	)
	needCommit = n7_growth_curve || needCommit

	n7_stack_by_growth_vector := enforcePlantHasShape[*StackOfGrowthCurve2DByGrowthVector](
		stager,
		func() *StackOfGrowthCurve2DByGrowthVector { return new(StackOfGrowthCurve2DByGrowthVector) },
		func(p *PlantAbstract) *StackOfGrowthCurve2DByGrowthVector {
			return p.StackOfGrowthCurve2DByGrowthVector
		},
		func(p *PlantAbstract, shape *StackOfGrowthCurve2DByGrowthVector) {
			p.StackOfGrowthCurve2DByGrowthVector = shape
		},
		func(p *PlantAbstract, shape *StackOfGrowthCurve2DByGrowthVector) bool {
			return p.StackOfGrowthCurve2DByGrowthVector == shape
		},
		"StackOfGrowthCurve2DByGrowthVector",
	)
	needCommit = n7_stack_by_growth_vector || needCommit

	return n1 || n2 || n3 || n4 || n5 || n6 || n7 || n7_base || n7_arc_normal || n7_arc_v2 || n7_arc_v2_end || n7_mid_arc || n7_growth_curve || n7_stack_by_growth_vector || needCommit
}

func enforceTubeVaseHasShape[ShapePointerType PointerToGongstruct](
	stager *Stager,
	newShape func() ShapePointerType,
	getShape func(vase *TubeVaseAbstract) ShapePointerType,
	setShape func(vase *TubeVaseAbstract, shape ShapePointerType),
	isOwned func(vase *TubeVaseAbstract, shape ShapePointerType) bool,
	shapeName string,
) (needCommit bool) {
	stage := stager.stage

	// 1. Ensure each Vase has the shape
	for vase := range *GetGongstructInstancesSetFromPointerType[*TubeVaseAbstract](stage) {
		var zero ShapePointerType
		if getShape(vase) == zero {
			shapePointer := newShape()
			shapePointer.StageVoid(stage)

			setShape(vase, shapePointer)
			stager.logAndNotify(fmt.Sprintf("TubeVase %s: created missing %s", vase.Name, shapeName))

			needCommit = true
		}
	}

	// 2. Ensure each Shape belongs to exactly one Vase. If orphaned, remove it.
	for shape := range *GetGongstructInstancesSetFromPointerType[ShapePointerType](stage) {
		hasOwner := false
		for vase := range *GetGongstructInstancesSetFromPointerType[*TubeVaseAbstract](stage) {
			if isOwned(vase, shape) {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			shape.UnstageVoid(stage)
			stager.logAndNotify(fmt.Sprintf("Removed orphaned %s %s", shapeName, shape.GetName()))
			needCommit = true
		}
	}

	return
}

func enforceTubeVaseShapeName[ShapePointerType PointerToGongstruct](
	stager *Stager,
	getShape func(vase *TubeVaseAbstract) ShapePointerType,
	shapeNameSuffix string,
) (needCommit bool) {
	stage := stager.stage

	for vase := range *GetGongstructInstancesSetFromPointerType[*TubeVaseAbstract](stage) {
		var zero ShapePointerType
		shape := getShape(vase)
		if shape != zero {
			expectedName := vase.Name + "-" + shapeNameSuffix
			if shape.GetName() != expectedName {
				oldName := shape.GetName()
				shape.SetName(expectedName)
				stager.logAndNotify(fmt.Sprintf("Renamed %s from '%s' to '%s'", shapeNameSuffix, oldName, expectedName))
				needCommit = true
			}
		}
	}

	return
}

func (stager *Stager) enforceTubeVaseHasShapes() (needCommit bool) {
	n7_halfway := enforceTubeVaseHasShape[*PerpendicularVectorGridHalfway](
		stager,
		func() *PerpendicularVectorGridHalfway { return new(PerpendicularVectorGridHalfway) },
		func(v *TubeVaseAbstract) *PerpendicularVectorGridHalfway { return v.PerpendicularVectorGridHalfway },
		func(v *TubeVaseAbstract, shape *PerpendicularVectorGridHalfway) {
			v.PerpendicularVectorGridHalfway = shape
		},
		func(v *TubeVaseAbstract, shape *PerpendicularVectorGridHalfway) bool {
			return v.PerpendicularVectorGridHalfway == shape
		},
		"PerpendicularVectorGridHalfway",
	)
	needCommit = n7_halfway || needCommit

	n7_top_arc_v2 := enforceTubeVaseHasShape[*TopStartArcShapeGrid](
		stager,
		func() *TopStartArcShapeGrid { return new(TopStartArcShapeGrid) },
		func(v *TubeVaseAbstract) *TopStartArcShapeGrid { return v.TopStartArcShapeGrid },
		func(v *TubeVaseAbstract, shape *TopStartArcShapeGrid) { v.TopStartArcShapeGrid = shape },
		func(v *TubeVaseAbstract, shape *TopStartArcShapeGrid) bool {
			return v.TopStartArcShapeGrid == shape
		},
		"TopStartArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2 || needCommit

	n7_top_arc_v2_end := enforceTubeVaseHasShape[*TopEndArcShapeGrid](
		stager,
		func() *TopEndArcShapeGrid { return new(TopEndArcShapeGrid) },
		func(v *TubeVaseAbstract) *TopEndArcShapeGrid { return v.TopEndArcShapeGrid },
		func(v *TubeVaseAbstract, shape *TopEndArcShapeGrid) { v.TopEndArcShapeGrid = shape },
		func(v *TubeVaseAbstract, shape *TopEndArcShapeGrid) bool {
			return v.TopEndArcShapeGrid == shape
		},
		"TopEndArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2_end || needCommit

	n10 := enforceTubeVaseHasShape[*StackOfRotatedGrowthCurve2D](
		stager,
		func() *StackOfRotatedGrowthCurve2D { return new(StackOfRotatedGrowthCurve2D) },
		func(v *TubeVaseAbstract) *StackOfRotatedGrowthCurve2D { return v.StackOfRotatedGrowthCurve2D },
		func(v *TubeVaseAbstract, shape *StackOfRotatedGrowthCurve2D) { v.StackOfRotatedGrowthCurve2D = shape },
		func(v *TubeVaseAbstract, shape *StackOfRotatedGrowthCurve2D) bool {
			return v.StackOfRotatedGrowthCurve2D == shape
		},
		"StackOfGrowthCurveV2",
	)
	needCommit = n10 || needCommit

	n11 := enforceTubeVaseHasShape[*TopStackOfRotatedGrowthCurve2D](
		stager,
		func() *TopStackOfRotatedGrowthCurve2D { return new(TopStackOfRotatedGrowthCurve2D) },
		func(v *TubeVaseAbstract) *TopStackOfRotatedGrowthCurve2D { return v.TopStackOfRotatedGrowthCurve2D },
		func(v *TubeVaseAbstract, shape *TopStackOfRotatedGrowthCurve2D) {
			v.TopStackOfRotatedGrowthCurve2D = shape
		},
		func(v *TubeVaseAbstract, shape *TopStackOfRotatedGrowthCurve2D) bool {
			return v.TopStackOfRotatedGrowthCurve2D == shape
		},
		"TopStackOfGrowthCurveV2",
	)
	needCommit = n11 || needCommit

	n16 := enforceTubeVaseHasShape[*TopMidArcVectorShapeGrid](
		stager,
		func() *TopMidArcVectorShapeGrid { return new(TopMidArcVectorShapeGrid) },
		func(v *TubeVaseAbstract) *TopMidArcVectorShapeGrid { return v.TopMidArcVectorShapeGrid },
		func(v *TubeVaseAbstract, shape *TopMidArcVectorShapeGrid) { v.TopMidArcVectorShapeGrid = shape },
		func(v *TubeVaseAbstract, shape *TopMidArcVectorShapeGrid) bool {
			return v.TopMidArcVectorShapeGrid == shape
		},
		"TopMidArcVectorShapeGrid",
	)
	needCommit = n16 || needCommit

	n14 := enforceTubeVaseHasShape[*ShiftedBottomTopStartArcShapeGrid](
		stager,
		func() *ShiftedBottomTopStartArcShapeGrid { return new(ShiftedBottomTopStartArcShapeGrid) },
		func(v *TubeVaseAbstract) *ShiftedBottomTopStartArcShapeGrid { return v.ShiftedBottomTopStartArcShapeGrid },
		func(v *TubeVaseAbstract, shape *ShiftedBottomTopStartArcShapeGrid) {
			v.ShiftedBottomTopStartArcShapeGrid = shape
		},
		func(v *TubeVaseAbstract, shape *ShiftedBottomTopStartArcShapeGrid) bool {
			return v.ShiftedBottomTopStartArcShapeGrid == shape
		},
		"ShiftedBottomTopStartArcShapeGrid",
	)
	needCommit = n14 || needCommit

	n_halfway_start := enforceTubeVaseHasShape[*StartHalfwayArcShapeGrid](
		stager,
		func() *StartHalfwayArcShapeGrid { return new(StartHalfwayArcShapeGrid) },
		func(v *TubeVaseAbstract) *StartHalfwayArcShapeGrid { return v.StartHalfwayArcShapeGrid },
		func(v *TubeVaseAbstract, shape *StartHalfwayArcShapeGrid) { v.StartHalfwayArcShapeGrid = shape },
		func(v *TubeVaseAbstract, shape *StartHalfwayArcShapeGrid) bool {
			return v.StartHalfwayArcShapeGrid == shape
		},
		"StartHalfwayArcShapeGrid",
	)
	needCommit = n_halfway_start || needCommit

	n_top_halfway_start := enforceTubeVaseHasShape[*TopStartHalfwayArcShapeGrid](
		stager,
		func() *TopStartHalfwayArcShapeGrid { return new(TopStartHalfwayArcShapeGrid) },
		func(v *TubeVaseAbstract) *TopStartHalfwayArcShapeGrid { return v.TopStartHalfwayArcShapeGrid },
		func(v *TubeVaseAbstract, shape *TopStartHalfwayArcShapeGrid) { v.TopStartHalfwayArcShapeGrid = shape },
		func(v *TubeVaseAbstract, shape *TopStartHalfwayArcShapeGrid) bool {
			return v.TopStartHalfwayArcShapeGrid == shape
		},
		"TopStartHalfwayArcShapeGrid",
	)
	needCommit = n_top_halfway_start || needCommit

	n_halfway_end := enforceTubeVaseHasShape[*EndHalfwayArcShapeGrid](
		stager,
		func() *EndHalfwayArcShapeGrid { return new(EndHalfwayArcShapeGrid) },
		func(v *TubeVaseAbstract) *EndHalfwayArcShapeGrid { return v.EndHalfwayArcShapeGrid },
		func(v *TubeVaseAbstract, shape *EndHalfwayArcShapeGrid) { v.EndHalfwayArcShapeGrid = shape },
		func(v *TubeVaseAbstract, shape *EndHalfwayArcShapeGrid) bool {
			return v.EndHalfwayArcShapeGrid == shape
		},
		"EndHalfwayArcShapeGrid",
	)
	needCommit = n_halfway_end || needCommit

	n_top_halfway_end := enforceTubeVaseHasShape[*TopEndHalfwayArcShapeGrid](
		stager,
		func() *TopEndHalfwayArcShapeGrid { return new(TopEndHalfwayArcShapeGrid) },
		func(v *TubeVaseAbstract) *TopEndHalfwayArcShapeGrid { return v.TopEndHalfwayArcShapeGrid },
		func(v *TubeVaseAbstract, shape *TopEndHalfwayArcShapeGrid) { v.TopEndHalfwayArcShapeGrid = shape },
		func(v *TubeVaseAbstract, shape *TopEndHalfwayArcShapeGrid) bool {
			return v.TopEndHalfwayArcShapeGrid == shape
		},
		"TopEndHalfwayArcShapeGrid",
	)
	needCommit = n_top_halfway_end || needCommit

	n_top_gc := enforceTubeVaseHasShape[*TopGrowthCurve2D](
		stager,
		func() *TopGrowthCurve2D { return new(TopGrowthCurve2D) },
		func(v *TubeVaseAbstract) *TopGrowthCurve2D { return v.TopGrowthCurve2D },
		func(v *TubeVaseAbstract, shape *TopGrowthCurve2D) { v.TopGrowthCurve2D = shape },
		func(v *TubeVaseAbstract, shape *TopGrowthCurve2D) bool {
			return v.TopGrowthCurve2D == shape
		},
		"TopGrowthCurve2D",
	)
	needCommit = n_top_gc || needCommit

	n17 := enforceTubeVaseHasShape[*StackOfGrowthCurve2D](
		stager,
		func() *StackOfGrowthCurve2D { return new(StackOfGrowthCurve2D) },
		func(v *TubeVaseAbstract) *StackOfGrowthCurve2D { return v.StackOfGrowthCurve2D },
		func(v *TubeVaseAbstract, shape *StackOfGrowthCurve2D) { v.StackOfGrowthCurve2D = shape },
		func(v *TubeVaseAbstract, shape *StackOfGrowthCurve2D) bool {
			return v.StackOfGrowthCurve2D == shape
		},
		"StackOfGrowthCurve2D",
	)
	needCommit = n17 || needCommit

	n18 := enforceTubeVaseHasShape[*TopStackOfGrowthCurve2D](
		stager,
		func() *TopStackOfGrowthCurve2D { return new(TopStackOfGrowthCurve2D) },
		func(v *TubeVaseAbstract) *TopStackOfGrowthCurve2D { return v.TopStackOfGrowthCurve2D },
		func(v *TubeVaseAbstract, shape *TopStackOfGrowthCurve2D) { v.TopStackOfGrowthCurve2D = shape },
		func(v *TubeVaseAbstract, shape *TopStackOfGrowthCurve2D) bool {
			return v.TopStackOfGrowthCurve2D == shape
		},
		"TopStackOfGrowthCurve2D",
	)
	needCommit = n18 || needCommit

	n19 := enforceTubeVaseHasShape[*StackOfGrowthCurve2DRibbon](
		stager,
		func() *StackOfGrowthCurve2DRibbon { return new(StackOfGrowthCurve2DRibbon) },
		func(v *TubeVaseAbstract) *StackOfGrowthCurve2DRibbon { return v.StackOfGrowthCurve2DRibbon },
		func(v *TubeVaseAbstract, shape *StackOfGrowthCurve2DRibbon) { v.StackOfGrowthCurve2DRibbon = shape },
		func(v *TubeVaseAbstract, shape *StackOfGrowthCurve2DRibbon) bool {
			return v.StackOfGrowthCurve2DRibbon == shape
		},
		"StackOfGrowthCurve2DRibbon",
	)
	needCommit = n19 || needCommit

	n20 := enforceTubeVaseHasShape[*StackOfRotatedGrowthCurve2DRibbon](
		stager,
		func() *StackOfRotatedGrowthCurve2DRibbon { return new(StackOfRotatedGrowthCurve2DRibbon) },
		func(v *TubeVaseAbstract) *StackOfRotatedGrowthCurve2DRibbon { return v.StackOfRotatedGrowthCurve2DRibbon },
		func(v *TubeVaseAbstract, shape *StackOfRotatedGrowthCurve2DRibbon) {
			v.StackOfRotatedGrowthCurve2DRibbon = shape
		},
		func(v *TubeVaseAbstract, shape *StackOfRotatedGrowthCurve2DRibbon) bool {
			return v.StackOfRotatedGrowthCurve2DRibbon == shape
		},
		"StackOfRotatedGrowthCurve2DRibbon",
	)
	needCommit = n20 || needCommit

	n21 := enforceTubeVaseHasShape[*PartiallyGrowthCurve2DRibbon](
		stager,
		func() *PartiallyGrowthCurve2DRibbon { return new(PartiallyGrowthCurve2DRibbon) },
		func(v *TubeVaseAbstract) *PartiallyGrowthCurve2DRibbon { return v.PartiallyGrowthCurve2DRibbon },
		func(v *TubeVaseAbstract, shape *PartiallyGrowthCurve2DRibbon) { v.PartiallyGrowthCurve2DRibbon = shape },
		func(v *TubeVaseAbstract, shape *PartiallyGrowthCurve2DRibbon) bool {
			return v.PartiallyGrowthCurve2DRibbon == shape
		},
		"PartiallyGrowthCurve2DRibbon",
	)
	needCommit = n21 || needCommit

	n21_shiftedleft_partially := enforceTubeVaseHasShape[*ShiftedLeftPartiallyGrowthCurve2DRibbon](
		stager,
		func() *ShiftedLeftPartiallyGrowthCurve2DRibbon { return new(ShiftedLeftPartiallyGrowthCurve2DRibbon) },
		func(v *TubeVaseAbstract) *ShiftedLeftPartiallyGrowthCurve2DRibbon {
			return v.ShiftedLeftPartiallyGrowthCurve2DRibbon
		},
		func(v *TubeVaseAbstract, shape *ShiftedLeftPartiallyGrowthCurve2DRibbon) {
			v.ShiftedLeftPartiallyGrowthCurve2DRibbon = shape
		},
		func(v *TubeVaseAbstract, shape *ShiftedLeftPartiallyGrowthCurve2DRibbon) bool {
			return v.ShiftedLeftPartiallyGrowthCurve2DRibbon == shape
		},
		"ShiftedLeftPartiallyGrowthCurve2DRibbon",
	)
	needCommit = n21_shiftedleft_partially || needCommit

	n21_traj := enforceTubeVaseHasShape[*PartiallyGrowthCurve2DTrajectory](
		stager,
		func() *PartiallyGrowthCurve2DTrajectory { return new(PartiallyGrowthCurve2DTrajectory) },
		func(v *TubeVaseAbstract) *PartiallyGrowthCurve2DTrajectory { return v.PartiallyGrowthCurve2DTrajectory },
		func(v *TubeVaseAbstract, shape *PartiallyGrowthCurve2DTrajectory) {
			v.PartiallyGrowthCurve2DTrajectory = shape
		},
		func(v *TubeVaseAbstract, shape *PartiallyGrowthCurve2DTrajectory) bool {
			return v.PartiallyGrowthCurve2DTrajectory == shape
		},
		"PartiallyGrowthCurve2DTrajectory",
	)
	needCommit = n21_traj || needCommit

	n21_trajP1P2 := enforceTubeVaseHasShape[*PartiallyGrowthCurve2DTrajectoryP1P2](
		stager,
		func() *PartiallyGrowthCurve2DTrajectoryP1P2 { return new(PartiallyGrowthCurve2DTrajectoryP1P2) },
		func(v *TubeVaseAbstract) *PartiallyGrowthCurve2DTrajectoryP1P2 {
			return v.PartiallyGrowthCurve2DTrajectoryP1P2
		},
		func(v *TubeVaseAbstract, shape *PartiallyGrowthCurve2DTrajectoryP1P2) {
			v.PartiallyGrowthCurve2DTrajectoryP1P2 = shape
		},
		func(v *TubeVaseAbstract, shape *PartiallyGrowthCurve2DTrajectoryP1P2) bool {
			return v.PartiallyGrowthCurve2DTrajectoryP1P2 == shape
		},
		"PartiallyGrowthCurve2DTrajectoryP1P2",
	)
	needCommit = n21_trajP1P2 || needCommit

	n21_px := enforceTubeVaseHasShape[*PxShape](
		stager,
		func() *PxShape { return new(PxShape) },
		func(v *TubeVaseAbstract) *PxShape { return v.PxShape },
		func(v *TubeVaseAbstract, shape *PxShape) { v.PxShape = shape },
		func(v *TubeVaseAbstract, shape *PxShape) bool {
			return v.PxShape == shape
		},
		"PxShape",
	)
	needCommit = n21_px || needCommit

	n21_chosenP1P2 := enforceTubeVaseHasShape[*ChosenP1P2PairShape](
		stager,
		func() *ChosenP1P2PairShape { return new(ChosenP1P2PairShape) },
		func(v *TubeVaseAbstract) *ChosenP1P2PairShape { return v.ChosenP1P2PairShape },
		func(v *TubeVaseAbstract, shape *ChosenP1P2PairShape) { v.ChosenP1P2PairShape = shape },
		func(v *TubeVaseAbstract, shape *ChosenP1P2PairShape) bool {
			return v.ChosenP1P2PairShape == shape
		},
		"ChosenP1P2PairShape",
	)
	needCommit = n21_chosenP1P2 || needCommit

	n21_keyHole := enforceTubeVaseHasShape[*KeyHoleShape](
		stager,
		func() *KeyHoleShape { return new(KeyHoleShape) },
		func(v *TubeVaseAbstract) *KeyHoleShape { return v.KeyHoleShape },
		func(v *TubeVaseAbstract, shape *KeyHoleShape) { v.KeyHoleShape = shape },
		func(v *TubeVaseAbstract, shape *KeyHoleShape) bool {
			return v.KeyHoleShape == shape
		},
		"KeyHoleShape",
	)
	needCommit = n21_keyHole || needCommit

	n22 := enforceTubeVaseHasShape[*GrowthCurve2DRibbon](
		stager,
		func() *GrowthCurve2DRibbon { return new(GrowthCurve2DRibbon) },
		func(v *TubeVaseAbstract) *GrowthCurve2DRibbon { return v.GrowthCurve2DRibbon },
		func(v *TubeVaseAbstract, shape *GrowthCurve2DRibbon) { v.GrowthCurve2DRibbon = shape },
		func(v *TubeVaseAbstract, shape *GrowthCurve2DRibbon) bool {
			return v.GrowthCurve2DRibbon == shape
		},
		"GrowthCurve2DRibbon",
	)
	needCommit = n22 || needCommit

	n23 := enforceTubeVaseHasShape[*ShiftedRightGrowthCurve2DRibbon](
		stager,
		func() *ShiftedRightGrowthCurve2DRibbon { return new(ShiftedRightGrowthCurve2DRibbon) },
		func(v *TubeVaseAbstract) *ShiftedRightGrowthCurve2DRibbon { return v.ShiftedRightGrowthCurve2DRibbon },
		func(v *TubeVaseAbstract, shape *ShiftedRightGrowthCurve2DRibbon) {
			v.ShiftedRightGrowthCurve2DRibbon = shape
		},
		func(v *TubeVaseAbstract, shape *ShiftedRightGrowthCurve2DRibbon) bool {
			return v.ShiftedRightGrowthCurve2DRibbon == shape
		},
		"ShiftedRightGrowthCurve2DRibbon",
	)
	needCommit = n23 || needCommit

	n24 := enforceTubeVaseHasShape[*ShiftedLeftGrowthCurve2DRibbon](
		stager,
		func() *ShiftedLeftGrowthCurve2DRibbon { return new(ShiftedLeftGrowthCurve2DRibbon) },
		func(v *TubeVaseAbstract) *ShiftedLeftGrowthCurve2DRibbon { return v.ShiftedLeftGrowthCurve2DRibbon },
		func(v *TubeVaseAbstract, shape *ShiftedLeftGrowthCurve2DRibbon) {
			v.ShiftedLeftGrowthCurve2DRibbon = shape
		},
		func(v *TubeVaseAbstract, shape *ShiftedLeftGrowthCurve2DRibbon) bool {
			return v.ShiftedLeftGrowthCurve2DRibbon == shape
		},
		"ShiftedLeftGrowthCurve2DRibbon",
	)
	needCommit = n24 || needCommit

	return
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
		"StartArcShapeGrid",
	)
	needCommit = n7_arc_v2 || needCommit

	n7_arc_v2_end := enforcePlantShapeName[*EndArcShapeGrid](
		stager,
		func(p *PlantAbstract) *EndArcShapeGrid { return p.EndArcShapeGrid },
		"EndArcShapeGrid",
	)
	needCommit = n7_arc_v2_end || needCommit

	n7_mid_arc := enforcePlantShapeName[*MidArcVectorShapeGrid](
		stager,
		func(p *PlantAbstract) *MidArcVectorShapeGrid { return p.MidArcVectorShapeGrid },
		"MidArcVectorShapeGrid",
	)
	needCommit = n7_mid_arc || needCommit

	n7_growth_curve := enforcePlantShapeName[*GrowthCurve2D](
		stager,
		func(p *PlantAbstract) *GrowthCurve2D { return p.GrowthCurve2D },
		"GrowthCurve2D",
	)
	needCommit = n7_growth_curve || needCommit

	return n1 || n2 || n3 || n4 || n5 || n6 || n7 || n7_base || n7_arc_normal || n7_arc_v2 || n7_arc_v2_end || n7_mid_arc || n7_growth_curve || needCommit
}

func (stager *Stager) enforceTubeVaseShapeNames() (needCommit bool) {
	n7_halfway := enforceTubeVaseShapeName[*PerpendicularVectorGridHalfway](
		stager,
		func(v *TubeVaseAbstract) *PerpendicularVectorGridHalfway { return v.PerpendicularVectorGridHalfway },
		"PerpendicularVectorGridHalfway",
	)
	needCommit = n7_halfway || needCommit

	n7_top_arc_v2 := enforceTubeVaseShapeName[*TopStartArcShapeGrid](
		stager,
		func(v *TubeVaseAbstract) *TopStartArcShapeGrid { return v.TopStartArcShapeGrid },
		"TopStartArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2 || needCommit

	n7_top_arc_v2_end := enforceTubeVaseShapeName[*TopEndArcShapeGrid](
		stager,
		func(v *TubeVaseAbstract) *TopEndArcShapeGrid { return v.TopEndArcShapeGrid },
		"TopEndArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2_end || needCommit

	n10 := enforceTubeVaseShapeName[*StackOfRotatedGrowthCurve2D](
		stager,
		func(v *TubeVaseAbstract) *StackOfRotatedGrowthCurve2D { return v.StackOfRotatedGrowthCurve2D },
		"StackOfGrowthCurveV2",
	)
	needCommit = n10 || needCommit

	n11 := enforceTubeVaseShapeName[*TopStackOfRotatedGrowthCurve2D](
		stager,
		func(v *TubeVaseAbstract) *TopStackOfRotatedGrowthCurve2D { return v.TopStackOfRotatedGrowthCurve2D },
		"TopStackOfGrowthCurveV2",
	)
	needCommit = n11 || needCommit

	n12 := enforceTubeVaseShapeName[*StackOfGrowthCurve2D](
		stager,
		func(v *TubeVaseAbstract) *StackOfGrowthCurve2D { return v.StackOfGrowthCurve2D },
		"StackOfGrowthCurve2D",
	)
	needCommit = n12 || needCommit

	n13 := enforceTubeVaseShapeName[*TopStackOfGrowthCurve2D](
		stager,
		func(v *TubeVaseAbstract) *TopStackOfGrowthCurve2D { return v.TopStackOfGrowthCurve2D },
		"TopStackOfGrowthCurve2D",
	)
	needCommit = n13 || needCommit

	n14 := enforceTubeVaseShapeName[*StackOfGrowthCurve2DRibbon](
		stager,
		func(v *TubeVaseAbstract) *StackOfGrowthCurve2DRibbon { return v.StackOfGrowthCurve2DRibbon },
		"StackOfGrowthCurve2DRibbon",
	)
	needCommit = n14 || needCommit

	n15 := enforceTubeVaseShapeName[*StackOfRotatedGrowthCurve2DRibbon](
		stager,
		func(v *TubeVaseAbstract) *StackOfRotatedGrowthCurve2DRibbon { return v.StackOfRotatedGrowthCurve2DRibbon },
		"StackOfRotatedGrowthCurve2DRibbon",
	)
	needCommit = n15 || needCommit

	n16_r := enforceTubeVaseShapeName[*GrowthCurve2DRibbon](
		stager,
		func(v *TubeVaseAbstract) *GrowthCurve2DRibbon { return v.GrowthCurve2DRibbon },
		"GrowthCurve2DRibbon",
	)
	needCommit = n16_r || needCommit

	n17_r := enforceTubeVaseShapeName[*ShiftedRightGrowthCurve2DRibbon](
		stager,
		func(v *TubeVaseAbstract) *ShiftedRightGrowthCurve2DRibbon { return v.ShiftedRightGrowthCurve2DRibbon },
		"ShiftedRightGrowthCurve2DRibbon",
	)
	needCommit = n17_r || needCommit

	n18_r := enforceTubeVaseShapeName[*ShiftedLeftGrowthCurve2DRibbon](
		stager,
		func(v *TubeVaseAbstract) *ShiftedLeftGrowthCurve2DRibbon { return v.ShiftedLeftGrowthCurve2DRibbon },
		"ShiftedLeftGrowthCurve2DRibbon",
	)
	needCommit = n18_r || needCommit

	n19_r := enforceTubeVaseShapeName[*ShiftedLeftPartiallyGrowthCurve2DRibbon](
		stager,
		func(v *TubeVaseAbstract) *ShiftedLeftPartiallyGrowthCurve2DRibbon {
			return v.ShiftedLeftPartiallyGrowthCurve2DRibbon
		},
		"ShiftedLeftPartiallyGrowthCurve2DRibbon",
	)
	needCommit = n19_r || needCommit

	n20_keyHole := enforceTubeVaseShapeName[*KeyHoleShape](
		stager,
		func(v *TubeVaseAbstract) *KeyHoleShape { return v.KeyHoleShape },
		"KeyHoleShape",
	)
	needCommit = n20_keyHole || needCommit

	return needCommit
}

// enforcePlantRhombusGridShapeHasRhombuses ensures that each RhombusGridShape has the correct number of RhombusShapes and their X,Y fields are correctly computed
func isRhombusShapeOwnedByPlant(p *PlantAbstract, shape *RhombusShape) bool {
	if p.RhombusStuff.ReferenceRhombus == shape || p.RhombusStuff.RotatedReferenceRhombus == shape {
		return true
	}
	// Initial, Rotated and Growth grids no longer use generic RhombusShape.
	return false
}
