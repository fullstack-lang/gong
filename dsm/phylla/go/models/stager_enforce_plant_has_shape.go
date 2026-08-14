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

func enforceVaseHasShape[ShapePointerType PointerToGongstruct](
	stager *Stager,
	newShape func() ShapePointerType,
	getShape func(vase *VaseAbstract) ShapePointerType,
	setShape func(vase *VaseAbstract, shape ShapePointerType),
	isOwned func(vase *VaseAbstract, shape ShapePointerType) bool,
	shapeName string,
) (needCommit bool) {
	stage := stager.stage

	// 1. Ensure each Vase has the shape
	for vase := range *GetGongstructInstancesSetFromPointerType[*VaseAbstract](stage) {
		var zero ShapePointerType
		if getShape(vase) == zero {
			shapePointer := newShape()
			shapePointer.StageVoid(stage)

			setShape(vase, shapePointer)

			needCommit = true
		}
	}

	// 2. Ensure each Shape belongs to exactly one Vase. If orphaned, remove it.
	for shape := range *GetGongstructInstancesSetFromPointerType[ShapePointerType](stage) {
		hasOwner := false
		for vase := range *GetGongstructInstancesSetFromPointerType[*VaseAbstract](stage) {
			if isOwned(vase, shape) {
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

func enforceVaseShapeName[ShapePointerType PointerToGongstruct](
	stager *Stager,
	getShape func(vase *VaseAbstract) ShapePointerType,
	shapeNameSuffix string,
) (needCommit bool) {
	stage := stager.stage

	for vase := range *GetGongstructInstancesSetFromPointerType[*VaseAbstract](stage) {
		var zero ShapePointerType
		shape := getShape(vase)
		if shape != zero {
			expectedName := vase.Name + "-" + shapeNameSuffix
			if shape.GetName() != expectedName {
				shape.SetName(expectedName)
				needCommit = true
			}
		}
	}

	return
}

func (stager *Stager) enforceVaseHasShapes() (needCommit bool) {
	n7_halfway := enforceVaseHasShape[*PerpendicularVectorGridHalfway](
		stager,
		func() *PerpendicularVectorGridHalfway { return new(PerpendicularVectorGridHalfway) },
		func(v *VaseAbstract) *PerpendicularVectorGridHalfway { return v.PerpendicularVectorGridHalfway },
		func(v *VaseAbstract, shape *PerpendicularVectorGridHalfway) {
			v.PerpendicularVectorGridHalfway = shape
		},
		func(v *VaseAbstract, shape *PerpendicularVectorGridHalfway) bool {
			return v.PerpendicularVectorGridHalfway == shape
		},
		"PerpendicularVectorGridHalfway",
	)
	needCommit = n7_halfway || needCommit

	n7_top_arc_v2 := enforceVaseHasShape[*TopStartArcShapeGrid](
		stager,
		func() *TopStartArcShapeGrid { return new(TopStartArcShapeGrid) },
		func(v *VaseAbstract) *TopStartArcShapeGrid { return v.TopStartArcShapeGrid },
		func(v *VaseAbstract, shape *TopStartArcShapeGrid) { v.TopStartArcShapeGrid = shape },
		func(v *VaseAbstract, shape *TopStartArcShapeGrid) bool {
			return v.TopStartArcShapeGrid == shape
		},
		"TopStartArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2 || needCommit

	n7_top_arc_v2_end := enforceVaseHasShape[*TopEndArcShapeGrid](
		stager,
		func() *TopEndArcShapeGrid { return new(TopEndArcShapeGrid) },
		func(v *VaseAbstract) *TopEndArcShapeGrid { return v.TopEndArcShapeGrid },
		func(v *VaseAbstract, shape *TopEndArcShapeGrid) { v.TopEndArcShapeGrid = shape },
		func(v *VaseAbstract, shape *TopEndArcShapeGrid) bool {
			return v.TopEndArcShapeGrid == shape
		},
		"TopEndArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2_end || needCommit

	n10 := enforceVaseHasShape[*StackOfRotatedGrowthCurve2D](
		stager,
		func() *StackOfRotatedGrowthCurve2D { return new(StackOfRotatedGrowthCurve2D) },
		func(v *VaseAbstract) *StackOfRotatedGrowthCurve2D { return v.StackOfRotatedGrowthCurve2D },
		func(v *VaseAbstract, shape *StackOfRotatedGrowthCurve2D) { v.StackOfRotatedGrowthCurve2D = shape },
		func(v *VaseAbstract, shape *StackOfRotatedGrowthCurve2D) bool {
			return v.StackOfRotatedGrowthCurve2D == shape
		},
		"StackOfGrowthCurveV2",
	)
	needCommit = n10 || needCommit

	n11 := enforceVaseHasShape[*TopStackOfRotatedGrowthCurve2D](
		stager,
		func() *TopStackOfRotatedGrowthCurve2D { return new(TopStackOfRotatedGrowthCurve2D) },
		func(v *VaseAbstract) *TopStackOfRotatedGrowthCurve2D { return v.TopStackOfRotatedGrowthCurve2D },
		func(v *VaseAbstract, shape *TopStackOfRotatedGrowthCurve2D) {
			v.TopStackOfRotatedGrowthCurve2D = shape
		},
		func(v *VaseAbstract, shape *TopStackOfRotatedGrowthCurve2D) bool {
			return v.TopStackOfRotatedGrowthCurve2D == shape
		},
		"TopStackOfGrowthCurveV2",
	)
	needCommit = n11 || needCommit

	n16 := enforceVaseHasShape[*TopMidArcVectorShapeGrid](
		stager,
		func() *TopMidArcVectorShapeGrid { return new(TopMidArcVectorShapeGrid) },
		func(v *VaseAbstract) *TopMidArcVectorShapeGrid { return v.TopMidArcVectorShapeGrid },
		func(v *VaseAbstract, shape *TopMidArcVectorShapeGrid) { v.TopMidArcVectorShapeGrid = shape },
		func(v *VaseAbstract, shape *TopMidArcVectorShapeGrid) bool {
			return v.TopMidArcVectorShapeGrid == shape
		},
		"TopMidArcVectorShapeGrid",
	)
	needCommit = n16 || needCommit

	n14 := enforceVaseHasShape[*ShiftedBottomTopStartArcShapeGrid](
		stager,
		func() *ShiftedBottomTopStartArcShapeGrid { return new(ShiftedBottomTopStartArcShapeGrid) },
		func(v *VaseAbstract) *ShiftedBottomTopStartArcShapeGrid { return v.ShiftedBottomTopStartArcShapeGrid },
		func(v *VaseAbstract, shape *ShiftedBottomTopStartArcShapeGrid) {
			v.ShiftedBottomTopStartArcShapeGrid = shape
		},
		func(v *VaseAbstract, shape *ShiftedBottomTopStartArcShapeGrid) bool {
			return v.ShiftedBottomTopStartArcShapeGrid == shape
		},
		"ShiftedBottomTopStartArcShapeGrid",
	)
	needCommit = n14 || needCommit

	n_halfway_start := enforceVaseHasShape[*StartHalfwayArcShapeGrid](
		stager,
		func() *StartHalfwayArcShapeGrid { return new(StartHalfwayArcShapeGrid) },
		func(v *VaseAbstract) *StartHalfwayArcShapeGrid { return v.StartHalfwayArcShapeGrid },
		func(v *VaseAbstract, shape *StartHalfwayArcShapeGrid) { v.StartHalfwayArcShapeGrid = shape },
		func(v *VaseAbstract, shape *StartHalfwayArcShapeGrid) bool {
			return v.StartHalfwayArcShapeGrid == shape
		},
		"StartHalfwayArcShapeGrid",
	)
	needCommit = n_halfway_start || needCommit

	n_top_halfway_start := enforceVaseHasShape[*TopStartHalfwayArcShapeGrid](
		stager,
		func() *TopStartHalfwayArcShapeGrid { return new(TopStartHalfwayArcShapeGrid) },
		func(v *VaseAbstract) *TopStartHalfwayArcShapeGrid { return v.TopStartHalfwayArcShapeGrid },
		func(v *VaseAbstract, shape *TopStartHalfwayArcShapeGrid) { v.TopStartHalfwayArcShapeGrid = shape },
		func(v *VaseAbstract, shape *TopStartHalfwayArcShapeGrid) bool {
			return v.TopStartHalfwayArcShapeGrid == shape
		},
		"TopStartHalfwayArcShapeGrid",
	)
	needCommit = n_top_halfway_start || needCommit

	n_halfway_end := enforceVaseHasShape[*EndHalfwayArcShapeGrid](
		stager,
		func() *EndHalfwayArcShapeGrid { return new(EndHalfwayArcShapeGrid) },
		func(v *VaseAbstract) *EndHalfwayArcShapeGrid { return v.EndHalfwayArcShapeGrid },
		func(v *VaseAbstract, shape *EndHalfwayArcShapeGrid) { v.EndHalfwayArcShapeGrid = shape },
		func(v *VaseAbstract, shape *EndHalfwayArcShapeGrid) bool {
			return v.EndHalfwayArcShapeGrid == shape
		},
		"EndHalfwayArcShapeGrid",
	)
	needCommit = n_halfway_end || needCommit

	n_top_halfway_end := enforceVaseHasShape[*TopEndHalfwayArcShapeGrid](
		stager,
		func() *TopEndHalfwayArcShapeGrid { return new(TopEndHalfwayArcShapeGrid) },
		func(v *VaseAbstract) *TopEndHalfwayArcShapeGrid { return v.TopEndHalfwayArcShapeGrid },
		func(v *VaseAbstract, shape *TopEndHalfwayArcShapeGrid) { v.TopEndHalfwayArcShapeGrid = shape },
		func(v *VaseAbstract, shape *TopEndHalfwayArcShapeGrid) bool {
			return v.TopEndHalfwayArcShapeGrid == shape
		},
		"TopEndHalfwayArcShapeGrid",
	)
	needCommit = n_top_halfway_end || needCommit

	n_top_gc := enforceVaseHasShape[*TopGrowthCurve2D](
		stager,
		func() *TopGrowthCurve2D { return new(TopGrowthCurve2D) },
		func(v *VaseAbstract) *TopGrowthCurve2D { return v.TopGrowthCurve2D },
		func(v *VaseAbstract, shape *TopGrowthCurve2D) { v.TopGrowthCurve2D = shape },
		func(v *VaseAbstract, shape *TopGrowthCurve2D) bool {
			return v.TopGrowthCurve2D == shape
		},
		"TopGrowthCurve2D",
	)
	needCommit = n_top_gc || needCommit

	n17 := enforceVaseHasShape[*StackOfGrowthCurve2D](
		stager,
		func() *StackOfGrowthCurve2D { return new(StackOfGrowthCurve2D) },
		func(v *VaseAbstract) *StackOfGrowthCurve2D { return v.StackOfGrowthCurve2D },
		func(v *VaseAbstract, shape *StackOfGrowthCurve2D) { v.StackOfGrowthCurve2D = shape },
		func(v *VaseAbstract, shape *StackOfGrowthCurve2D) bool {
			return v.StackOfGrowthCurve2D == shape
		},
		"StackOfGrowthCurve2D",
	)
	needCommit = n17 || needCommit

	n18 := enforceVaseHasShape[*TopStackOfGrowthCurve2D](
		stager,
		func() *TopStackOfGrowthCurve2D { return new(TopStackOfGrowthCurve2D) },
		func(v *VaseAbstract) *TopStackOfGrowthCurve2D { return v.TopStackOfGrowthCurve2D },
		func(v *VaseAbstract, shape *TopStackOfGrowthCurve2D) { v.TopStackOfGrowthCurve2D = shape },
		func(v *VaseAbstract, shape *TopStackOfGrowthCurve2D) bool {
			return v.TopStackOfGrowthCurve2D == shape
		},
		"TopStackOfGrowthCurve2D",
	)
	needCommit = n18 || needCommit

	n19 := enforceVaseHasShape[*StackOfGrowthCurve2DRibbon](
		stager,
		func() *StackOfGrowthCurve2DRibbon { return new(StackOfGrowthCurve2DRibbon) },
		func(v *VaseAbstract) *StackOfGrowthCurve2DRibbon { return v.StackOfGrowthCurve2DRibbon },
		func(v *VaseAbstract, shape *StackOfGrowthCurve2DRibbon) { v.StackOfGrowthCurve2DRibbon = shape },
		func(v *VaseAbstract, shape *StackOfGrowthCurve2DRibbon) bool {
			return v.StackOfGrowthCurve2DRibbon == shape
		},
		"StackOfGrowthCurve2DRibbon",
	)
	needCommit = n19 || needCommit

	n20 := enforceVaseHasShape[*StackOfRotatedGrowthCurve2DRibbon](
		stager,
		func() *StackOfRotatedGrowthCurve2DRibbon { return new(StackOfRotatedGrowthCurve2DRibbon) },
		func(v *VaseAbstract) *StackOfRotatedGrowthCurve2DRibbon { return v.StackOfRotatedGrowthCurve2DRibbon },
		func(v *VaseAbstract, shape *StackOfRotatedGrowthCurve2DRibbon) {
			v.StackOfRotatedGrowthCurve2DRibbon = shape
		},
		func(v *VaseAbstract, shape *StackOfRotatedGrowthCurve2DRibbon) bool {
			return v.StackOfRotatedGrowthCurve2DRibbon == shape
		},
		"StackOfRotatedGrowthCurve2DRibbon",
	)
	needCommit = n20 || needCommit

	n21 := enforceVaseHasShape[*PartiallyGrowthCurve2DRibbon](
		stager,
		func() *PartiallyGrowthCurve2DRibbon { return new(PartiallyGrowthCurve2DRibbon) },
		func(v *VaseAbstract) *PartiallyGrowthCurve2DRibbon { return v.PartiallyGrowthCurve2DRibbon },
		func(v *VaseAbstract, shape *PartiallyGrowthCurve2DRibbon) { v.PartiallyGrowthCurve2DRibbon = shape },
		func(v *VaseAbstract, shape *PartiallyGrowthCurve2DRibbon) bool {
			return v.PartiallyGrowthCurve2DRibbon == shape
		},
		"PartiallyGrowthCurve2DRibbon",
	)
	needCommit = n21 || needCommit

	n21_shiftedleft_partially := enforceVaseHasShape[*ShiftedLeftPartiallyGrowthCurve2DRibbon](
		stager,
		func() *ShiftedLeftPartiallyGrowthCurve2DRibbon { return new(ShiftedLeftPartiallyGrowthCurve2DRibbon) },
		func(v *VaseAbstract) *ShiftedLeftPartiallyGrowthCurve2DRibbon {
			return v.ShiftedLeftPartiallyGrowthCurve2DRibbon
		},
		func(v *VaseAbstract, shape *ShiftedLeftPartiallyGrowthCurve2DRibbon) {
			v.ShiftedLeftPartiallyGrowthCurve2DRibbon = shape
		},
		func(v *VaseAbstract, shape *ShiftedLeftPartiallyGrowthCurve2DRibbon) bool {
			return v.ShiftedLeftPartiallyGrowthCurve2DRibbon == shape
		},
		"ShiftedLeftPartiallyGrowthCurve2DRibbon",
	)
	needCommit = n21_shiftedleft_partially || needCommit

	n21_traj := enforceVaseHasShape[*PartiallyGrowthCurve2DTrajectory](
		stager,
		func() *PartiallyGrowthCurve2DTrajectory { return new(PartiallyGrowthCurve2DTrajectory) },
		func(v *VaseAbstract) *PartiallyGrowthCurve2DTrajectory { return v.PartiallyGrowthCurve2DTrajectory },
		func(v *VaseAbstract, shape *PartiallyGrowthCurve2DTrajectory) {
			v.PartiallyGrowthCurve2DTrajectory = shape
		},
		func(v *VaseAbstract, shape *PartiallyGrowthCurve2DTrajectory) bool {
			return v.PartiallyGrowthCurve2DTrajectory == shape
		},
		"PartiallyGrowthCurve2DTrajectory",
	)
	needCommit = n21_traj || needCommit

	n21_trajP1P2 := enforceVaseHasShape[*PartiallyGrowthCurve2DTrajectoryP1P2](
		stager,
		func() *PartiallyGrowthCurve2DTrajectoryP1P2 { return new(PartiallyGrowthCurve2DTrajectoryP1P2) },
		func(v *VaseAbstract) *PartiallyGrowthCurve2DTrajectoryP1P2 {
			return v.PartiallyGrowthCurve2DTrajectoryP1P2
		},
		func(v *VaseAbstract, shape *PartiallyGrowthCurve2DTrajectoryP1P2) {
			v.PartiallyGrowthCurve2DTrajectoryP1P2 = shape
		},
		func(v *VaseAbstract, shape *PartiallyGrowthCurve2DTrajectoryP1P2) bool {
			return v.PartiallyGrowthCurve2DTrajectoryP1P2 == shape
		},
		"PartiallyGrowthCurve2DTrajectoryP1P2",
	)
	needCommit = n21_trajP1P2 || needCommit

	n21_px := enforceVaseHasShape[*PxShape](
		stager,
		func() *PxShape { return new(PxShape) },
		func(v *VaseAbstract) *PxShape { return v.PxShape },
		func(v *VaseAbstract, shape *PxShape) { v.PxShape = shape },
		func(v *VaseAbstract, shape *PxShape) bool {
			return v.PxShape == shape
		},
		"PxShape",
	)
	needCommit = n21_px || needCommit

	n21_chosenP1P2 := enforceVaseHasShape[*ChosenP1P2PairShape](
		stager,
		func() *ChosenP1P2PairShape { return new(ChosenP1P2PairShape) },
		func(v *VaseAbstract) *ChosenP1P2PairShape { return v.ChosenP1P2PairShape },
		func(v *VaseAbstract, shape *ChosenP1P2PairShape) { v.ChosenP1P2PairShape = shape },
		func(v *VaseAbstract, shape *ChosenP1P2PairShape) bool {
			return v.ChosenP1P2PairShape == shape
		},
		"ChosenP1P2PairShape",
	)
	needCommit = n21_chosenP1P2 || needCommit

	n21_keyHole := enforceVaseHasShape[*KeyHoleShape](
		stager,
		func() *KeyHoleShape { return new(KeyHoleShape) },
		func(v *VaseAbstract) *KeyHoleShape { return v.KeyHoleShape },
		func(v *VaseAbstract, shape *KeyHoleShape) { v.KeyHoleShape = shape },
		func(v *VaseAbstract, shape *KeyHoleShape) bool {
			return v.KeyHoleShape == shape
		},
		"KeyHoleShape",
	)
	needCommit = n21_keyHole || needCommit

	n22 := enforceVaseHasShape[*GrowthCurve2DRibbon](
		stager,
		func() *GrowthCurve2DRibbon { return new(GrowthCurve2DRibbon) },
		func(v *VaseAbstract) *GrowthCurve2DRibbon { return v.GrowthCurve2DRibbon },
		func(v *VaseAbstract, shape *GrowthCurve2DRibbon) { v.GrowthCurve2DRibbon = shape },
		func(v *VaseAbstract, shape *GrowthCurve2DRibbon) bool {
			return v.GrowthCurve2DRibbon == shape
		},
		"GrowthCurve2DRibbon",
	)
	needCommit = n22 || needCommit

	n23 := enforceVaseHasShape[*ShiftedRightGrowthCurve2DRibbon](
		stager,
		func() *ShiftedRightGrowthCurve2DRibbon { return new(ShiftedRightGrowthCurve2DRibbon) },
		func(v *VaseAbstract) *ShiftedRightGrowthCurve2DRibbon { return v.ShiftedRightGrowthCurve2DRibbon },
		func(v *VaseAbstract, shape *ShiftedRightGrowthCurve2DRibbon) {
			v.ShiftedRightGrowthCurve2DRibbon = shape
		},
		func(v *VaseAbstract, shape *ShiftedRightGrowthCurve2DRibbon) bool {
			return v.ShiftedRightGrowthCurve2DRibbon == shape
		},
		"ShiftedRightGrowthCurve2DRibbon",
	)
	needCommit = n23 || needCommit

	n24 := enforceVaseHasShape[*ShiftedLeftGrowthCurve2DRibbon](
		stager,
		func() *ShiftedLeftGrowthCurve2DRibbon { return new(ShiftedLeftGrowthCurve2DRibbon) },
		func(v *VaseAbstract) *ShiftedLeftGrowthCurve2DRibbon { return v.ShiftedLeftGrowthCurve2DRibbon },
		func(v *VaseAbstract, shape *ShiftedLeftGrowthCurve2DRibbon) {
			v.ShiftedLeftGrowthCurve2DRibbon = shape
		},
		func(v *VaseAbstract, shape *ShiftedLeftGrowthCurve2DRibbon) bool {
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

func (stager *Stager) enforceVaseShapeNames() (needCommit bool) {
	n7_halfway := enforceVaseShapeName[*PerpendicularVectorGridHalfway](
		stager,
		func(v *VaseAbstract) *PerpendicularVectorGridHalfway { return v.PerpendicularVectorGridHalfway },
		"PerpendicularVectorGridHalfway",
	)
	needCommit = n7_halfway || needCommit

	n7_top_arc_v2 := enforceVaseShapeName[*TopStartArcShapeGrid](
		stager,
		func(v *VaseAbstract) *TopStartArcShapeGrid { return v.TopStartArcShapeGrid },
		"TopStartArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2 || needCommit

	n7_top_arc_v2_end := enforceVaseShapeName[*TopEndArcShapeGrid](
		stager,
		func(v *VaseAbstract) *TopEndArcShapeGrid { return v.TopEndArcShapeGrid },
		"TopEndArcShapeV2Grid",
	)
	needCommit = n7_top_arc_v2_end || needCommit

	n10 := enforceVaseShapeName[*StackOfRotatedGrowthCurve2D](
		stager,
		func(v *VaseAbstract) *StackOfRotatedGrowthCurve2D { return v.StackOfRotatedGrowthCurve2D },
		"StackOfGrowthCurveV2",
	)
	needCommit = n10 || needCommit

	n11 := enforceVaseShapeName[*TopStackOfRotatedGrowthCurve2D](
		stager,
		func(v *VaseAbstract) *TopStackOfRotatedGrowthCurve2D { return v.TopStackOfRotatedGrowthCurve2D },
		"TopStackOfGrowthCurveV2",
	)
	needCommit = n11 || needCommit

	n12 := enforceVaseShapeName[*StackOfGrowthCurve2D](
		stager,
		func(v *VaseAbstract) *StackOfGrowthCurve2D { return v.StackOfGrowthCurve2D },
		"StackOfGrowthCurve2D",
	)
	needCommit = n12 || needCommit

	n13 := enforceVaseShapeName[*TopStackOfGrowthCurve2D](
		stager,
		func(v *VaseAbstract) *TopStackOfGrowthCurve2D { return v.TopStackOfGrowthCurve2D },
		"TopStackOfGrowthCurve2D",
	)
	needCommit = n13 || needCommit

	n14 := enforceVaseShapeName[*StackOfGrowthCurve2DRibbon](
		stager,
		func(v *VaseAbstract) *StackOfGrowthCurve2DRibbon { return v.StackOfGrowthCurve2DRibbon },
		"StackOfGrowthCurve2DRibbon",
	)
	needCommit = n14 || needCommit

	n15 := enforceVaseShapeName[*StackOfRotatedGrowthCurve2DRibbon](
		stager,
		func(v *VaseAbstract) *StackOfRotatedGrowthCurve2DRibbon { return v.StackOfRotatedGrowthCurve2DRibbon },
		"StackOfRotatedGrowthCurve2DRibbon",
	)
	needCommit = n15 || needCommit

	n16_r := enforceVaseShapeName[*GrowthCurve2DRibbon](
		stager,
		func(v *VaseAbstract) *GrowthCurve2DRibbon { return v.GrowthCurve2DRibbon },
		"GrowthCurve2DRibbon",
	)
	needCommit = n16_r || needCommit

	n17_r := enforceVaseShapeName[*ShiftedRightGrowthCurve2DRibbon](
		stager,
		func(v *VaseAbstract) *ShiftedRightGrowthCurve2DRibbon { return v.ShiftedRightGrowthCurve2DRibbon },
		"ShiftedRightGrowthCurve2DRibbon",
	)
	needCommit = n17_r || needCommit

	n18_r := enforceVaseShapeName[*ShiftedLeftGrowthCurve2DRibbon](
		stager,
		func(v *VaseAbstract) *ShiftedLeftGrowthCurve2DRibbon { return v.ShiftedLeftGrowthCurve2DRibbon },
		"ShiftedLeftGrowthCurve2DRibbon",
	)
	needCommit = n18_r || needCommit

	n19_r := enforceVaseShapeName[*ShiftedLeftPartiallyGrowthCurve2DRibbon](
		stager,
		func(v *VaseAbstract) *ShiftedLeftPartiallyGrowthCurve2DRibbon {
			return v.ShiftedLeftPartiallyGrowthCurve2DRibbon
		},
		"ShiftedLeftPartiallyGrowthCurve2DRibbon",
	)
	needCommit = n19_r || needCommit

	n20_keyHole := enforceVaseShapeName[*KeyHoleShape](
		stager,
		func(v *VaseAbstract) *KeyHoleShape { return v.KeyHoleShape },
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
