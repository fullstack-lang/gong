package models

type Plant struct {
	Name string

	N int
	M int
	Z int // number of rhombus

	// InsideAngle is set by the user. It represents the angle in degree of the diamond (rhombus) at the origin 0,0.
	// Together with the spiral indices (N and M) and the SideLength, this angle determines
	// the Cartesian X and Y coordinates of the plant's GrowthVector (or Initial Axis).
	// The GrowthVector represents the fundamental geometrical shift (the growth step)
	// from one leaf/cell to the next corresponding cell in the repeating spiral pattern.
	InsideAngle float64

	// SideLength is set by the user. It represents the length of the side of the
	// fundamental diamond (rhombus) shape that tiles the surface of the plant.
	// Like InsideAngle, it is a key parameter in computing the exact Cartesian coordinates
	// (length and angle) of the GrowthVector.
	SideLength float64

	// how many circle to go around for the front curve
	// the front curve goes from one circle to the nearest
	ShiftToNearestCircle int

	LibraryAbstractFields
	AbstractTypeFields

	IsPlantDiagramsNodeExpanded      bool
	PlantDiagramsWhoseNodeIsExpanded []*PlantDiagram
	PlantDiagrams                    []*PlantDiagram
}
