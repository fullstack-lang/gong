package models

import (
	"fmt"
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

// enforceGrowthVectorShapeName ensures that the name of the GrowthVectorShape matches its owning PlantDiagram
func (stager *Stager) enforceGrowthVectorShapeName() (needCommit bool) {
	return enforcePlantDiagramShapeName[*GrowthVectorShape](
		stager,
		func(pd *PlantDiagram) *GrowthVectorShape { return pd.GrowthVectorShape },
		"GrowthVectorShape",
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

// enforcePlantDiagramHasGrowthVectorShape ensures that each PlantDiagram has one and only one GrowthVectorShape that belong to it
func (stager *Stager) enforcePlantDiagramHasGrowthVectorShape() (needCommit bool) {
	return enforcePlantDiagramHasShape[*GrowthVectorShape](
		stager,
		func() *GrowthVectorShape { return new(GrowthVectorShape) },
		func(pd *PlantDiagram) *GrowthVectorShape { return pd.GrowthVectorShape },
		func(pd *PlantDiagram, shape *GrowthVectorShape) { pd.GrowthVectorShape = shape },
		func(pd *PlantDiagram, shape *GrowthVectorShape) bool {
			return pd.GrowthVectorShape == shape || pd.RotatedGrowthVectorShape == shape
		},
		"GrowthVectorShape",
	)
}

// enforcePlantDiagramHasReferenceRhombus ensures that each PlantDiagram has one and only one ReferenceRhombus that belong to it
func (stager *Stager) enforcePlantDiagramHasReferenceRhombus() (needCommit bool) {
	return enforcePlantDiagramHasShape[*ReferenceRhombus](
		stager,
		func() *ReferenceRhombus { return new(ReferenceRhombus) },
		func(pd *PlantDiagram) *ReferenceRhombus { return pd.ReferenceRhombus },
		func(pd *PlantDiagram, shape *ReferenceRhombus) { pd.ReferenceRhombus = shape },
		func(pd *PlantDiagram, shape *ReferenceRhombus) bool {
			return pd.ReferenceRhombus == shape || pd.RotatedReferenceRhombus == shape
		},
		"ReferenceRhombus",
	)
}

// enforcePlantDiagramHasRotatedShapes ensures that each PlantDiagram has its Rotated shapes
func (stager *Stager) enforcePlantDiagramHasRotatedShapes() (needCommit bool) {
	n1 := enforcePlantDiagramHasShape[*ReferenceRhombus](
		stager,
		func() *ReferenceRhombus { return new(ReferenceRhombus) },
		func(pd *PlantDiagram) *ReferenceRhombus { return pd.RotatedReferenceRhombus },
		func(pd *PlantDiagram, shape *ReferenceRhombus) { pd.RotatedReferenceRhombus = shape },
		func(pd *PlantDiagram, shape *ReferenceRhombus) bool {
			return pd.ReferenceRhombus == shape || pd.RotatedReferenceRhombus == shape
		},
		"RotatedReferenceRhombus",
	)

	n2 := enforcePlantDiagramHasShape[*GrowthVectorShape](
		stager,
		func() *GrowthVectorShape { return new(GrowthVectorShape) },
		func(pd *PlantDiagram) *GrowthVectorShape { return pd.RotatedGrowthVectorShape },
		func(pd *PlantDiagram, shape *GrowthVectorShape) { pd.RotatedGrowthVectorShape = shape },
		func(pd *PlantDiagram, shape *GrowthVectorShape) bool {
			return pd.GrowthVectorShape == shape || pd.RotatedGrowthVectorShape == shape
		},
		"RotatedGrowthVectorShape",
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
	return enforcePlantDiagramShapeName[*ReferenceRhombus](
		stager,
		func(pd *PlantDiagram) *ReferenceRhombus { return pd.ReferenceRhombus },
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

// enforceRotatedShapesNames ensures that the name of the Rotated shapes match their owning PlantDiagram
func (stager *Stager) enforceRotatedShapesNames() (needCommit bool) {
	n1 := enforcePlantDiagramShapeName[*ReferenceRhombus](
		stager,
		func(pd *PlantDiagram) *ReferenceRhombus { return pd.RotatedReferenceRhombus },
		"RotatedReferenceRhombus",
	)

	n2 := enforcePlantDiagramShapeName[*GrowthVectorShape](
		stager,
		func(pd *PlantDiagram) *GrowthVectorShape { return pd.RotatedGrowthVectorShape },
		"RotatedGrowthVectorShape",
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
