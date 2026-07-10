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

	n9 := enforcePlantHasShape[*StackOfGrowthCurve](
		stager,
		func() *StackOfGrowthCurve { return new(StackOfGrowthCurve) },
		func(p *Plant) *StackOfGrowthCurve { return p.StackOfGrowthCurve },
		func(p *Plant, shape *StackOfGrowthCurve) { p.StackOfGrowthCurve = shape },
		func(p *Plant, shape *StackOfGrowthCurve) bool {
			return p.StackOfGrowthCurve == shape
		},
		"StackOfGrowthCurve",
	)
	needCommit = n9 || needCommit

	return n1 || n2 || n3 || n4 || n5 || n6 || n7 || n8 || n9
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

	n8 := enforcePlantShapeName[*GrowthCurveBezierShapeGrid](
		stager,
		func(p *Plant) *GrowthCurveBezierShapeGrid { return p.GrowthCurveBezierShapeGrid },
		"GrowthCurveBezierShapeGrid",
	)
	needCommit = n8 || needCommit

	n9 := enforcePlantShapeName[*StackOfGrowthCurve](
		stager,
		func(p *Plant) *StackOfGrowthCurve { return p.StackOfGrowthCurve },
		"StackOfGrowthCurve",
	)
	needCommit = n9 || needCommit

	return n1 || n2 || n3 || n4 || n5 || n6 || n7 || n8 || n9
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

		if plant.InitialRhombusGridShape != nil {
			needCommit = enforceInitialRhombusGridShapeHasRhombuses(stage, plant.InitialRhombusGridShape, plant.N, plant.M, v1x, v1y, v2x, v2y, 0.0) || needCommit
		}
		rotRad := 0.0
		if plant.PlantCircumferenceShape != nil {
			rotRad = -plant.PlantCircumferenceShape.AngleDegree * math.Pi / 180.0
		}
		if plant.RotatedRhombusGridShape2 != nil {
			needCommit = enforceRotatedRhombusGridShapeHasRhombuses(stage, plant.RotatedRhombusGridShape2, plant.N, plant.M, v1x, v1y, v2x, v2y, rotRad) || needCommit
		}
		if plant.GrowthCurveRhombusGridShape != nil && plant.RotatedRhombusGridShape2 != nil && plant.PlantCircumferenceShape != nil {
			needCommit = enforceGrowthPathRhombusGridShapeHasRhombuses(stage, plant.GrowthCurveRhombusGridShape, plant.RotatedRhombusGridShape2, plant.PlantCircumferenceShape.AngleDegree, plant.RhombusSideLength, plant.RhombusInsideAngle, plant.PlantCircumferenceShape.Length) || needCommit
		}

		if plant.GrowthVectorShape != nil && plant.GrowthCurveRhombusGridShape != nil {
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
		
		if plant.PerpendicularVectorGrid != nil && plant.GrowthCurveRhombusGridShape != nil {
			needCommit = enforcePerpendicularVectorGridHasVectors(stage, plant.PerpendicularVectorGrid, plant.GrowthCurveRhombusGridShape, v1x, v1y, v2x, v2y, rotRad) || needCommit
		}
		
		circLen := 0.0
		if plant.PlantCircumferenceShape != nil {
			circLen = plant.PlantCircumferenceShape.Length
		}

		if plant.GrowthCurveBezierShapeGrid != nil && plant.PerpendicularVectorGrid != nil {
			needCommit = enforceGrowthCurveBezierShapeGridHasShapes(stage, plant.GrowthCurveBezierShapeGrid, plant.PerpendicularVectorGrid, plant.RhombusSideLength, circLen) || needCommit
		}

		if plant.StackOfGrowthCurve != nil && plant.GrowthCurveBezierShapeGrid != nil && plant.GrowthVectorShape != nil {
			needCommit = enforceStackOfGrowthCurveHasShapes(stage, plant.StackOfGrowthCurve, plant.GrowthCurveBezierShapeGrid, plant.GrowthVectorShape, plant.StackHeight, circLen) || needCommit
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

func enforceStackOfGrowthCurveHasShapes(stage *Stage, stack *StackOfGrowthCurve, grid *GrowthCurveBezierShapeGrid, vector *GrowthVectorShape, stackHeight int, circLen float64) (needCommit bool) {
	if stack == nil || grid == nil || vector == nil || stackHeight < 1 || circLen <= 0 {
		if len(stack.StackGrowthCurveBezierShapes) > 0 {
			stack.StackGrowthCurveBezierShapes = nil
			return true
		}
		return false
	}
	
	baseCurves := grid.GrowthCurveBezierShapes
	if len(baseCurves) == 0 {
		return false
	}
	
	// baseCurves contains both the original and the +circLen translated curves.
	// We only want the original ones (the first half).
	numBase := len(baseCurves) / 2
	if numBase == 0 {
		numBase = len(baseCurves)
	}
	
	type expectedShape struct {
		name string
		startX, startY, endX, endY float64
		cpStartX, cpStartY, cpEndX, cpEndY float64
	}
	var expected []expectedShape
	
	for h := 0; h < stackHeight; h++ {
		dx := float64(h) * vector.X
		dy := float64(h) * vector.Y
		
		for i := 0; i < numBase; i++ {
			baseCurve := baseCurves[i]
			currentDX := dx
			
			for (baseCurve.StartX + currentDX) >= 0 {
				currentDX -= circLen
			}
			
			currentDX += circLen
			
			startX := baseCurve.StartX + currentDX
			expected = append(expected, expectedShape{
				name: fmt.Sprintf("%s-layer-%d-%d", stack.Name, h, i),
				startX: startX,
				startY: baseCurve.StartY + dy,
				endX: baseCurve.EndX + currentDX,
				endY: baseCurve.EndY + dy,
				cpStartX: baseCurve.ControlPointStartX + currentDX,
				cpStartY: baseCurve.ControlPointStartY + dy,
				cpEndX: baseCurve.ControlPointEndX + currentDX,
				cpEndY: baseCurve.ControlPointEndY + dy,
			})
		}
	}
	
	valid := true
	if len(stack.StackGrowthCurveBezierShapes) != len(expected) {
		valid = false
	} else {
		for i, exp := range expected {
			b := stack.StackGrowthCurveBezierShapes[i]
			if b == nil || b.Name != exp.name {
				valid = false
				break
			}
			
			if math.Abs(b.StartX-exp.startX) > 1e-4 || math.Abs(b.StartY-exp.startY) > 1e-4 ||
				math.Abs(b.EndX-exp.endX) > 1e-4 || math.Abs(b.EndY-exp.endY) > 1e-4 ||
				math.Abs(b.ControlPointStartX-exp.cpStartX) > 1e-4 || math.Abs(b.ControlPointStartY-exp.cpStartY) > 1e-4 ||
				math.Abs(b.ControlPointEndX-exp.cpEndX) > 1e-4 || math.Abs(b.ControlPointEndY-exp.cpEndY) > 1e-4 {
				valid = false
				break
			}
		}
	}
	
	if !valid {
		stack.StackGrowthCurveBezierShapes = make([]*StackGrowthCurveBezierShape, len(expected))
		for i, exp := range expected {
			b := new(StackGrowthCurveBezierShape).Stage(stage)
			b.Name = exp.name
			
			b.StartX = exp.startX
			b.StartY = exp.startY
			b.EndX = exp.endX
			b.EndY = exp.endY
			b.ControlPointStartX = exp.cpStartX
			b.ControlPointStartY = exp.cpStartY
			b.ControlPointEndX = exp.cpEndX
			b.ControlPointEndY = exp.cpEndY
			
			stack.StackGrowthCurveBezierShapes[i] = b
		}
		needCommit = true
	}
	
	return needCommit
}

