package models

import (
	"fmt"
	"math"
	"time"
)

func enforcePlantDiagramHasShape[ShapePointerType PointerToGongstruct](
	stager *Stager,
	newShape func() ShapePointerType,
	getShape func(plantDiagram *PlantDiagram) ShapePointerType,
	setShape func(plantDiagram *PlantDiagram, shape ShapePointerType),
	isOwned func(plantDiagram *PlantDiagram, shape ShapePointerType) bool,
	shapeName string,
) (needCommit bool) {
	stage := stager.stage

	// 1. Ensure each PlantDiagram has the shape
	for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		// Since we can't compare ShapePointerType (interface constraint) directly to nil if it's nil pointer inside interface,
		// Wait, ShapePointerType is a type parameter, so checking against a zero value is better.
		// A zero value of a pointer is nil.
		var zero ShapePointerType
		if getShape(plantDiagram) == zero {
			shapePointer := newShape()
			shapePointer.StageVoid(stage)

			setShape(plantDiagram, shapePointer)

			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Added default %s for %s", shapeName, plantDiagram.Name))
			}

			needCommit = true
		}
	}

	// 2. Ensure each Shape belongs to exactly one PlantDiagram. If orphaned, remove it.
	for shape := range *GetGongstructInstancesSetFromPointerType[ShapePointerType](stage) {
		hasOwner := false
		for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
			if isOwned(plantDiagram, shape) {
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

func enforcePlantDiagramShapeName[ShapePointerType PointerToGongstruct](
	stager *Stager,
	getShape func(plantDiagram *PlantDiagram) ShapePointerType,
	shapeNameSuffix string,
) (needCommit bool) {
	stage := stager.stage

	for plantDiagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		var zero ShapePointerType
		shape := getShape(plantDiagram)
		if shape != zero {
			expectedName := plantDiagram.Name + "-" + shapeNameSuffix
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

// enforceAxesShapeName ensures that the name of the AxesShape matches its owning PlantDiagram
func (stager *Stager) enforceAxesShapeName() (needCommit bool) {
	return enforcePlantDiagramShapeName[*AxesShape](
		stager,
		func(pd *PlantDiagram) *AxesShape { return pd.AxesShape },
		"AxesShape",
	)
}

// enforceGridPathShapeName ensures that the name of the GridPathShape matches its owning PlantDiagram
func (stager *Stager) enforceGridPathShapeName() (needCommit bool) {
	return enforcePlantDiagramShapeName[*GridPathShape](
		stager,
		func(pd *PlantDiagram) *GridPathShape { return pd.GridPathShape },
		"GridPathShape",
	)
}

// enforcePlantCircumferenceShapeName ensures that the name of the PlantCircumferenceShape matches its owning PlantDiagram
func (stager *Stager) enforcePlantCircumferenceShapeName() (needCommit bool) {
	return enforcePlantDiagramShapeName[*PlantCircumferenceShape](
		stager,
		func(pd *PlantDiagram) *PlantCircumferenceShape { return pd.PlantCircumferenceShape },
		"PlantCircumferenceShape",
	)
}

// enforcePlantDiagramHasAxes ensures that each PlantDiagram has one and only one Axes that belong to it
func (stager *Stager) enforcePlantDiagramHasAxes() (needCommit bool) {
	return enforcePlantDiagramHasShape[*AxesShape](
		stager,
		func() *AxesShape { return new(AxesShape) },
		func(pd *PlantDiagram) *AxesShape { return pd.AxesShape },
		func(pd *PlantDiagram, shape *AxesShape) { pd.AxesShape = shape },
		func(pd *PlantDiagram, shape *AxesShape) bool { return pd.AxesShape == shape },
		"AxesShape",
	)
}

// enforcePlantDiagramHasGridPathShape ensures that each PlantDiagram has one and only one GridPathShape that belong to it
func (stager *Stager) enforcePlantDiagramHasGridPathShape() (needCommit bool) {
	return enforcePlantDiagramHasShape[*GridPathShape](
		stager,
		func() *GridPathShape { return new(GridPathShape) },
		func(pd *PlantDiagram) *GridPathShape { return pd.GridPathShape },
		func(pd *PlantDiagram, shape *GridPathShape) { pd.GridPathShape = shape },
		func(pd *PlantDiagram, shape *GridPathShape) bool {
			return pd.GridPathShape == shape || pd.RotatedGridPathShape == shape
		},
		"GridPathShape",
	)
}

// enforcePlantDiagramHasRhombusGridShape ensures that each PlantDiagram has one and only one RhombusGridShape that belong to it
func (stager *Stager) enforcePlantDiagramHasRhombusGridShape() (needCommit bool) {
	return enforcePlantDiagramHasShape[*RhombusGridShape](
		stager,
		func() *RhombusGridShape { return new(RhombusGridShape) },
		func(pd *PlantDiagram) *RhombusGridShape { return pd.RhombusGridShape },
		func(pd *PlantDiagram, shape *RhombusGridShape) { pd.RhombusGridShape = shape },
		func(pd *PlantDiagram, shape *RhombusGridShape) bool {
			return pd.RhombusGridShape == shape || pd.RotatedRhombusGridShape == shape
		},
		"RhombusGridShape",
	)
}

// enforcePlantDiagramHasExplanationTextShape ensures that each PlantDiagram has one and only one ExplanationTextShape that belong to it
func (stager *Stager) enforcePlantDiagramHasExplanationTextShape() (needCommit bool) {
	return enforcePlantDiagramHasShape[*ExplanationTextShape](
		stager,
		func() *ExplanationTextShape { return new(ExplanationTextShape) },
		func(pd *PlantDiagram) *ExplanationTextShape { return pd.ExplanationTextShape },
		func(pd *PlantDiagram, shape *ExplanationTextShape) { pd.ExplanationTextShape = shape },
		func(pd *PlantDiagram, shape *ExplanationTextShape) bool {
			return pd.ExplanationTextShape == shape
		},
		"ExplanationTextShape",
	)
}

// enforcePlantDiagramHasPlantCircumferenceShape ensures that each PlantDiagram has one and only one PlantCircumferenceShape that belong to it
func (stager *Stager) enforcePlantDiagramHasPlantCircumferenceShape() (needCommit bool) {
	return enforcePlantDiagramHasShape[*PlantCircumferenceShape](
		stager,
		func() *PlantCircumferenceShape { return new(PlantCircumferenceShape) },
		func(pd *PlantDiagram) *PlantCircumferenceShape { return pd.PlantCircumferenceShape },
		func(pd *PlantDiagram, shape *PlantCircumferenceShape) { pd.PlantCircumferenceShape = shape },
		func(pd *PlantDiagram, shape *PlantCircumferenceShape) bool {
			return pd.PlantCircumferenceShape == shape || pd.RotatedPlantCircumferenceShape == shape
		},
		"PlantCircumferenceShape",
	)
}

// enforcePlantDiagramHasReferenceRhombus ensures that each PlantDiagram has one and only one ReferenceRhombus that belong to it
func (stager *Stager) enforcePlantDiagramHasReferenceRhombus() (needCommit bool) {
	return enforcePlantDiagramHasShape[*RhombusShape](
		stager,
		func() *RhombusShape { return new(RhombusShape) },
		func(pd *PlantDiagram) *RhombusShape { return pd.ReferenceRhombus },
		func(pd *PlantDiagram, shape *RhombusShape) { pd.ReferenceRhombus = shape },
		func(pd *PlantDiagram, shape *RhombusShape) bool {
			return isRhombusShapeOwnedByPlantDiagram(pd, shape)
		},
		"ReferenceRhombus",
	)
}

// enforcePlantDiagramHasRotatedShapes ensures that each PlantDiagram has its Rotated shapes
func (stager *Stager) enforcePlantDiagramHasRotatedShapes() (needCommit bool) {
	n1 := enforcePlantDiagramHasShape[*RhombusShape](
		stager,
		func() *RhombusShape { return new(RhombusShape) },
		func(pd *PlantDiagram) *RhombusShape { return pd.RotatedReferenceRhombus },
		func(pd *PlantDiagram, shape *RhombusShape) { pd.RotatedReferenceRhombus = shape },
		func(pd *PlantDiagram, shape *RhombusShape) bool {
			return isRhombusShapeOwnedByPlantDiagram(pd, shape)
		},
		"RotatedReferenceRhombus",
	)

	n2 := enforcePlantDiagramHasShape[*PlantCircumferenceShape](
		stager,
		func() *PlantCircumferenceShape { return new(PlantCircumferenceShape) },
		func(pd *PlantDiagram) *PlantCircumferenceShape { return pd.RotatedPlantCircumferenceShape },
		func(pd *PlantDiagram, shape *PlantCircumferenceShape) { pd.RotatedPlantCircumferenceShape = shape },
		func(pd *PlantDiagram, shape *PlantCircumferenceShape) bool {
			return pd.PlantCircumferenceShape == shape || pd.RotatedPlantCircumferenceShape == shape
		},
		"RotatedPlantCircumferenceShape",
	)

	n3 := enforcePlantDiagramHasShape[*GridPathShape](
		stager,
		func() *GridPathShape { return new(GridPathShape) },
		func(pd *PlantDiagram) *GridPathShape { return pd.RotatedGridPathShape },
		func(pd *PlantDiagram, shape *GridPathShape) { pd.RotatedGridPathShape = shape },
		func(pd *PlantDiagram, shape *GridPathShape) bool {
			return pd.GridPathShape == shape || pd.RotatedGridPathShape == shape
		},
		"RotatedGridPathShape",
	)

	n4 := enforcePlantDiagramHasShape[*RhombusGridShape](
		stager,
		func() *RhombusGridShape { return new(RhombusGridShape) },
		func(pd *PlantDiagram) *RhombusGridShape { return pd.RotatedRhombusGridShape },
		func(pd *PlantDiagram, shape *RhombusGridShape) { pd.RotatedRhombusGridShape = shape },
		func(pd *PlantDiagram, shape *RhombusGridShape) bool {
			return pd.RhombusGridShape == shape || pd.RotatedRhombusGridShape == shape
		},
		"RotatedRhombusGridShape",
	)

	return n1 || n2 || n3 || n4
}

// enforceReferenceRhombusName ensures that the name of the ReferenceRhombus matches its owning PlantDiagram
func (stager *Stager) enforceReferenceRhombusName() (needCommit bool) {
	return enforcePlantDiagramShapeName[*RhombusShape](
		stager,
		func(pd *PlantDiagram) *RhombusShape { return pd.ReferenceRhombus },
		"ReferenceRhombus",
	)
}

// enforceRhombusGridShapeName ensures that the name of the RhombusGridShape matches its owning PlantDiagram
func (stager *Stager) enforceRhombusGridShapeName() (needCommit bool) {
	return enforcePlantDiagramShapeName[*RhombusGridShape](
		stager,
		func(pd *PlantDiagram) *RhombusGridShape { return pd.RhombusGridShape },
		"RhombusGridShape",
	)
}

// enforceExplanationTextShapeName ensures that the name of the ExplanationTextShape matches its owning PlantDiagram
func (stager *Stager) enforceExplanationTextShapeName() (needCommit bool) {
	return enforcePlantDiagramShapeName[*ExplanationTextShape](
		stager,
		func(pd *PlantDiagram) *ExplanationTextShape { return pd.ExplanationTextShape },
		"ExplanationTextShape",
	)
}

// enforceRotatedShapesNames ensures that the name of the Rotated shapes match their owning PlantDiagram
func (stager *Stager) enforceRotatedShapesNames() (needCommit bool) {
	n1 := enforcePlantDiagramShapeName[*RhombusShape](
		stager,
		func(pd *PlantDiagram) *RhombusShape { return pd.RotatedReferenceRhombus },
		"RotatedReferenceRhombus",
	)

	n2 := enforcePlantDiagramShapeName[*PlantCircumferenceShape](
		stager,
		func(pd *PlantDiagram) *PlantCircumferenceShape { return pd.RotatedPlantCircumferenceShape },
		"RotatedPlantCircumferenceShape",
	)

	n3 := enforcePlantDiagramShapeName[*GridPathShape](
		stager,
		func(pd *PlantDiagram) *GridPathShape { return pd.RotatedGridPathShape },
		"RotatedGridPathShape",
	)

	n4 := enforcePlantDiagramShapeName[*RhombusGridShape](
		stager,
		func(pd *PlantDiagram) *RhombusGridShape { return pd.RotatedRhombusGridShape },
		"RotatedRhombusGridShape",
	)

	return n1 || n2 || n3 || n4
}

// enforcePlantDiagramRhombusGridShapeHasRhombuses ensures that each RhombusGridShape has the correct number of RhombusShapes and their X,Y fields are correctly computed
func (stager *Stager) enforcePlantDiagramRhombusGridShapeHasRhombuses() (needCommit bool) {
	stage := stager.stage

	// We need plant N and M and angles. They are accessible via plant.PlantDiagrams
	for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
		angleRad := plant.RhombusInsideAngle * math.Pi / 180.0
		length := plant.RhombusSideLength

		// SVG Y-axis is inverted
		v1x := length * math.Cos(angleRad/2.0)
		v1y := -length * math.Sin(angleRad/2.0)

		v2x := -length * math.Cos(angleRad/2.0)
		v2y := -length * math.Sin(angleRad/2.0)

		for _, pd := range plant.PlantDiagrams {
			if pd.RhombusGridShape != nil {
				needCommit = enforceRhombusGridShapeHasRhombuses(stage, pd.RhombusGridShape, plant.N, plant.M, v1x, v1y, v2x, v2y) || needCommit
			}
			if pd.RotatedRhombusGridShape != nil {
				needCommit = enforceRhombusGridShapeHasRhombuses(stage, pd.RotatedRhombusGridShape, plant.N, plant.M, v1x, v1y, v2x, v2y) || needCommit
			}
		}
	}
	return
}

func enforceRhombusGridShapeHasRhombuses(stage *Stage, grid *RhombusGridShape, N, M int, v1x, v1y, v2x, v2y float64) (needCommit bool) {
	valid := true
	if len(grid.RhombusShapes) != (N+1)*M {
		valid = false
	} else {
		seen := make(map[*RhombusShape]bool)
		idx := 0
		for i := -1; i < N; i++ {
			for j := 0; j < M; j++ {
				r := grid.RhombusShapes[idx]
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
		grid.RhombusShapes = make([]*RhombusShape, 0, (N+1)*M)
		for i := -1; i < N; i++ {
			for j := 0; j < M; j++ {
				r := new(RhombusShape).Stage(stage)
				r.Name = fmt.Sprintf("%s-%d-%d", grid.Name, i, j)
				r.X = float64(i)*v1x + float64(j)*v2x
				r.Y = float64(i)*v1y + float64(j)*v2y
				grid.RhombusShapes = append(grid.RhombusShapes, r)
			}
		}
		needCommit = true
	} else {
		idx := 0
		for i := -1; i < N; i++ {
			for j := 0; j < M; j++ {
				r := grid.RhombusShapes[idx]
				expectedX := float64(i)*v1x + float64(j)*v2x
				expectedY := float64(i)*v1y + float64(j)*v2y
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

func isRhombusShapeOwnedByPlantDiagram(pd *PlantDiagram, shape *RhombusShape) bool {
	if pd.ReferenceRhombus == shape || pd.RotatedReferenceRhombus == shape {
		return true
	}
	if pd.RhombusGridShape != nil {
		for _, r := range pd.RhombusGridShape.RhombusShapes {
			if r == shape {
				return true
			}
		}
	}
	if pd.RotatedRhombusGridShape != nil {
		for _, r := range pd.RotatedRhombusGridShape.RhombusShapes {
			if r == shape {
				return true
			}
		}
	}
	return false
}
